package main

import (
	"go/token"

	"github.com/dave/dst"
	"github.com/dave/dst/dstutil"
)

// removeDeadCode drops statements that can never run, e.g. the `return r` ccgo emits
// after an exhaustive switch. go vet reports them as unreachable code.
func removeDeadCode(root dst.Node) {
	dst.Inspect(root, func(node dst.Node) bool {
		fn, ok := node.(*dst.FuncDecl)
		if !ok || fn.Body == nil {
			return true
		}

		// dropping a goto can leave its label unused, which makes the code after it dead
		for changed := true; changed; {
			changed = pruneDeadStmts(fn.Body)
			changed = dropUnusedLabels(fn) || changed
		}

		return false
	})
}

func pruneDeadStmts(body *dst.BlockStmt) bool {
	changed := false

	prune := func(list []dst.Stmt) []dst.Stmt {
		var result []dst.Stmt
		dead := false
		for _, stmt := range list {
			if _, ok := stmt.(*dst.LabeledStmt); ok {
				// a used label can be reached by a goto
				dead = false
			}

			if dead {
				changed = true
				continue
			}

			result = append(result, stmt)

			// break and continue are not terminating, but nothing after them runs either
			_, branch := stmt.(*dst.BranchStmt)
			dead = branch || isTerminating(stmt, "")
		}

		return result
	}

	dst.Inspect(body, func(node dst.Node) bool {
		switch node := node.(type) {
		case *dst.BlockStmt:
			node.List = prune(node.List)
		case *dst.CaseClause:
			node.Body = prune(node.Body)
		case *dst.CommClause:
			node.Body = prune(node.Body)
		}

		return true
	})

	return changed
}

func dropUnusedLabels(fn *dst.FuncDecl) bool {
	used := map[string]bool{}
	dst.Inspect(fn.Body, func(node dst.Node) bool {
		if branch, ok := node.(*dst.BranchStmt); ok && branch.Label != nil {
			used[branch.Label.Name] = true
		}

		return true
	})

	changed := false
	dstutil.Apply(fn.Body, nil, func(c *dstutil.Cursor) bool {
		labeled, ok := c.Node().(*dst.LabeledStmt)
		if !ok || used[labeled.Label.Name] {
			return true
		}

		changed = true
		if _, empty := labeled.Stmt.(*dst.EmptyStmt); empty && c.Index() >= 0 {
			c.Delete()
		} else {
			c.Replace(labeled.Stmt)
		}

		return true
	})

	return changed
}

// isTerminating implements "terminating statements" from the Go spec.
// label is the label of stmt, if any.
func isTerminating(stmt dst.Stmt, label string) bool {
	switch stmt := stmt.(type) {
	case *dst.ReturnStmt:
		return true

	case *dst.BranchStmt:
		return stmt.Tok == token.GOTO || stmt.Tok == token.FALLTHROUGH

	case *dst.ExprStmt:
		call, ok := stmt.X.(*dst.CallExpr)
		return ok && identOf(call.Fun) == "panic"

	case *dst.BlockStmt:
		return len(stmt.List) > 0 && isTerminating(stmt.List[len(stmt.List)-1], "")

	case *dst.IfStmt:
		return stmt.Else != nil && isTerminating(stmt.Body, "") && isTerminating(stmt.Else, "")

	case *dst.LabeledStmt:
		return isTerminating(stmt.Stmt, stmt.Label.Name)

	case *dst.ForStmt:
		return stmt.Cond == nil && !hasBreak(stmt.Body.List, label)

	case *dst.SwitchStmt:
		return clausesTerminate(stmt.Body, label)

	case *dst.TypeSwitchStmt:
		return clausesTerminate(stmt.Body, label)

	case *dst.SelectStmt:
		return clausesTerminate(stmt.Body, label)
	}

	return false
}

func clausesTerminate(body *dst.BlockStmt, label string) bool {
	hasDefault := false
	for _, clause := range body.List {
		var stmts []dst.Stmt
		switch clause := clause.(type) {
		case *dst.CaseClause:
			hasDefault = hasDefault || clause.List == nil
			stmts = clause.Body
		case *dst.CommClause:
			// a select does not need a default case
			hasDefault = true
			stmts = clause.Body
		}

		if len(stmts) == 0 || !isTerminating(stmts[len(stmts)-1], "") || hasBreak(stmts, label) {
			return false
		}
	}

	return hasDefault
}

// hasBreak reports whether stmts contain a break that leaves the enclosing statement
// labeled with label (may be empty).
func hasBreak(stmts []dst.Stmt, label string) bool {
	found := false

	var visit func(node dst.Node, nested bool)
	visit = func(node dst.Node, nested bool) {
		dst.Inspect(node, func(node dst.Node) bool {
			if found {
				return false
			}

			var nestedBody *dst.BlockStmt
			switch node := node.(type) {
			case *dst.FuncLit:
				return false
			case *dst.BranchStmt:
				if node.Tok == token.BREAK {
					found = node.Label == nil && !nested || node.Label != nil && node.Label.Name == label
				}
			case *dst.ForStmt:
				nestedBody = node.Body
			case *dst.RangeStmt:
				nestedBody = node.Body
			case *dst.SwitchStmt:
				nestedBody = node.Body
			case *dst.TypeSwitchStmt:
				nestedBody = node.Body
			case *dst.SelectStmt:
				nestedBody = node.Body
			}

			if nestedBody != nil && !nested {
				// an unlabeled break in here belongs to the nested statement
				visit(nestedBody, true)
				return false
			}

			return true
		})
	}

	for _, stmt := range stmts {
		visit(stmt, false)
	}

	return found
}
