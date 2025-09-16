package b2

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b2DynamicTree_Create(tls *_Stack) (r DynamicTree) {
	var i int32
	var tree DynamicTree
	_, _ = i, tree
	tree.Root = -int32(1)
	tree.NodeCapacity = int32(16)
	tree.NodeCount = 0
	tree.Nodes = b2Alloc(tls, int32FromUint64(uint64FromInt32(tree.NodeCapacity)*uint64(40)))
	memset(tls, tree.Nodes, 0, uint64FromInt32(tree.NodeCapacity)*uint64(40))
	// Build a linked list for the free list.
	i = 0
	for {
		if !(i < tree.NodeCapacity-int32(1)) {
			break
		}
		*(*int32_t)(unsafe.Add(unsafe.Pointer(&*(*b2TreeNode)(unsafe.Pointer(tree.Nodes + uintptr(i)*40))), 32)) = i + int32(1)
		goto _1
	_1:
		;
		i = i + 1
	}
	*(*int32_t)(unsafe.Add(unsafe.Pointer(&*(*b2TreeNode)(unsafe.Pointer(tree.Nodes + uintptr(tree.NodeCapacity-int32(1))*40))), 32)) = -int32FromInt32(1)
	tree.FreeList = 0
	tree.ProxyCount = 0
	tree.LeafIndices = uintptrFromInt32(0)
	tree.LeafBoxes = uintptrFromInt32(0)
	tree.LeafCenters = uintptrFromInt32(0)
	tree.BinIndices = uintptrFromInt32(0)
	tree.RebuildCapacity = 0
	return tree
}

func b2DynamicTree_Destroy(tls *_Stack, tree uintptr) {
	b2Free(tls, (*DynamicTree)(unsafe.Pointer(tree)).Nodes, int32FromUint64(uint64FromInt32((*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity)*uint64(40)))
	b2Free(tls, (*DynamicTree)(unsafe.Pointer(tree)).LeafIndices, int32FromUint64(uint64FromInt32((*DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(4)))
	b2Free(tls, (*DynamicTree)(unsafe.Pointer(tree)).LeafBoxes, int32FromUint64(uint64FromInt32((*DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(16)))
	b2Free(tls, (*DynamicTree)(unsafe.Pointer(tree)).LeafCenters, int32FromUint64(uint64FromInt32((*DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(8)))
	b2Free(tls, (*DynamicTree)(unsafe.Pointer(tree)).BinIndices, int32FromUint64(uint64FromInt32((*DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(4)))
	memset(tls, tree, 0, uint64(72))
}

// C documentation
//
//	// Create a proxy in the tree as a leaf node. We return the index of the node instead of a pointer so that we can grow
//	// the node pool.
func b2DynamicTree_CreateProxy(tls *_Stack, tree uintptr, aabb AABB, categoryBits uint64, userData uint64) (r int32) {
	var node uintptr
	var proxyId int32
	var shouldRotate uint8
	_, _, _ = node, proxyId, shouldRotate
	if !(-float32(float32FromFloat32(100000)*b2_lengthUnitsPerMeter) < aabb.LowerBound.X && aabb.LowerBound.X < float32(float32FromFloat32(100000)*b2_lengthUnitsPerMeter)) && b2InternalAssertFcn(tls, __ccgo_ts+5228, __ccgo_ts+4746, int32FromInt32(772)) != 0 {
		__builtin_trap(tls)
	}
	if !(-float32(float32FromFloat32(100000)*b2_lengthUnitsPerMeter) < aabb.LowerBound.Y && aabb.LowerBound.Y < float32(float32FromFloat32(100000)*b2_lengthUnitsPerMeter)) && b2InternalAssertFcn(tls, __ccgo_ts+5288, __ccgo_ts+4746, int32FromInt32(773)) != 0 {
		__builtin_trap(tls)
	}
	if !(-float32(float32FromFloat32(100000)*b2_lengthUnitsPerMeter) < aabb.UpperBound.X && aabb.UpperBound.X < float32(float32FromFloat32(100000)*b2_lengthUnitsPerMeter)) && b2InternalAssertFcn(tls, __ccgo_ts+5348, __ccgo_ts+4746, int32FromInt32(774)) != 0 {
		__builtin_trap(tls)
	}
	if !(-float32(float32FromFloat32(100000)*b2_lengthUnitsPerMeter) < aabb.UpperBound.Y && aabb.UpperBound.Y < float32(float32FromFloat32(100000)*b2_lengthUnitsPerMeter)) && b2InternalAssertFcn(tls, __ccgo_ts+5408, __ccgo_ts+4746, int32FromInt32(775)) != 0 {
		__builtin_trap(tls)
	}
	proxyId = b2AllocateNode(tls, tree)
	node = (*DynamicTree)(unsafe.Pointer(tree)).Nodes + uintptr(proxyId)*40
	(*b2TreeNode)(unsafe.Pointer(node)).Aabb = aabb
	*(*uint64_t)(unsafe.Add(unsafe.Pointer(node), 24)) = userData
	(*b2TreeNode)(unsafe.Pointer(node)).CategoryBits = categoryBits
	(*b2TreeNode)(unsafe.Pointer(node)).Height = uint16(0)
	(*b2TreeNode)(unsafe.Pointer(node)).Flags = uint16FromInt32(int32(b2_allocatedNode) | int32(b2_leafNode))
	shouldRotate = boolUint8(true1 != 0)
	b2InsertLeaf(tls, tree, proxyId, shouldRotate)
	*(*int32)(unsafe.Pointer(tree + 24)) += int32(1)
	return proxyId
}

func b2DynamicTree_DestroyProxy(tls *_Stack, tree uintptr, proxyId int32) {
	if !(0 <= proxyId && proxyId < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) && b2InternalAssertFcn(tls, __ccgo_ts+5468, __ccgo_ts+4746, int32FromInt32(796)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsLeaf(tls, (*DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr(proxyId)*40) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+5513, __ccgo_ts+4746, int32FromInt32(797)) != 0 {
		__builtin_trap(tls)
	}
	b2RemoveLeaf(tls, tree, proxyId)
	b2FreeNode(tls, tree, proxyId)
	if !((*DynamicTree)(unsafe.Pointer(tree)).ProxyCount > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+5547, __ccgo_ts+4746, int32FromInt32(802)) != 0 {
		__builtin_trap(tls)
	}
	*(*int32)(unsafe.Pointer(tree + 24)) -= int32(1)
}

func b2DynamicTree_GetProxyCount(tls *_Stack, tree uintptr) (r int32) {
	return (*DynamicTree)(unsafe.Pointer(tree)).ProxyCount
}

func b2DynamicTree_MoveProxy(tls *_Stack, tree uintptr, proxyId int32, aabb AABB) {
	var shouldRotate uint8
	_ = shouldRotate
	if !(b2IsValidAABB(tls, aabb) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+5568, __ccgo_ts+4746, int32FromInt32(813)) != 0 {
		__builtin_trap(tls)
	}
	if !(aabb.UpperBound.X-aabb.LowerBound.X < float32(float32FromFloat32(100000)*b2_lengthUnitsPerMeter)) && b2InternalAssertFcn(tls, __ccgo_ts+5590, __ccgo_ts+4746, int32FromInt32(814)) != 0 {
		__builtin_trap(tls)
	}
	if !(aabb.UpperBound.Y-aabb.LowerBound.Y < float32(float32FromFloat32(100000)*b2_lengthUnitsPerMeter)) && b2InternalAssertFcn(tls, __ccgo_ts+5638, __ccgo_ts+4746, int32FromInt32(815)) != 0 {
		__builtin_trap(tls)
	}
	if !(0 <= proxyId && proxyId < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) && b2InternalAssertFcn(tls, __ccgo_ts+5468, __ccgo_ts+4746, int32FromInt32(816)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsLeaf(tls, (*DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr(proxyId)*40) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+5513, __ccgo_ts+4746, int32FromInt32(817)) != 0 {
		__builtin_trap(tls)
	}
	b2RemoveLeaf(tls, tree, proxyId)
	(*(*b2TreeNode)(unsafe.Pointer((*DynamicTree)(unsafe.Pointer(tree)).Nodes + uintptr(proxyId)*40))).Aabb = aabb
	shouldRotate = boolUint8(false1 != 0)
	b2InsertLeaf(tls, tree, proxyId, shouldRotate)
}

func b2DynamicTree_EnlargeProxy(tls *_Stack, tree uintptr, proxyId int32, aabb AABB) {
	var changed, changed1, s, v3, v7 uint8
	var nodes, v5, p10, p9 uintptr
	var parentIndex int32
	var v1, v2, v6 AABB
	_, _, _, _, _, _, _, _, _, _, _, _, _ = changed, changed1, nodes, parentIndex, s, v1, v2, v3, v5, v6, v7, p10, p9
	nodes = (*DynamicTree)(unsafe.Pointer(tree)).Nodes
	if !(b2IsValidAABB(tls, aabb) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+5568, __ccgo_ts+4746, int32FromInt32(831)) != 0 {
		__builtin_trap(tls)
	}
	if !(aabb.UpperBound.X-aabb.LowerBound.X < float32(float32FromFloat32(100000)*b2_lengthUnitsPerMeter)) && b2InternalAssertFcn(tls, __ccgo_ts+5590, __ccgo_ts+4746, int32FromInt32(832)) != 0 {
		__builtin_trap(tls)
	}
	if !(aabb.UpperBound.Y-aabb.LowerBound.Y < float32(float32FromFloat32(100000)*b2_lengthUnitsPerMeter)) && b2InternalAssertFcn(tls, __ccgo_ts+5638, __ccgo_ts+4746, int32FromInt32(833)) != 0 {
		__builtin_trap(tls)
	}
	if !(0 <= proxyId && proxyId < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) && b2InternalAssertFcn(tls, __ccgo_ts+5468, __ccgo_ts+4746, int32FromInt32(834)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsLeaf(tls, (*DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr(proxyId)*40) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+5513, __ccgo_ts+4746, int32FromInt32(835)) != 0 {
		__builtin_trap(tls)
	}
	// Caller must ensure this
	v1 = (*(*b2TreeNode)(unsafe.Pointer(nodes + uintptr(proxyId)*40))).Aabb
	v2 = aabb
	s = boolUint8(true1 != 0)
	s = boolUint8(s != 0 && v1.LowerBound.X <= v2.LowerBound.X)
	s = boolUint8(s != 0 && v1.LowerBound.Y <= v2.LowerBound.Y)
	s = boolUint8(s != 0 && v2.UpperBound.X <= v1.UpperBound.X)
	s = boolUint8(s != 0 && v2.UpperBound.Y <= v1.UpperBound.Y)
	v3 = s
	goto _4
_4:
	;
	if !(int32FromUint8(v3) == int32FromInt32(false1)) && b2InternalAssertFcn(tls, __ccgo_ts+5686, __ccgo_ts+4746, int32FromInt32(838)) != 0 {
		__builtin_trap(tls)
	}
	(*(*b2TreeNode)(unsafe.Pointer(nodes + uintptr(proxyId)*40))).Aabb = aabb
	parentIndex = (*(*b2TreeNode)(unsafe.Pointer(nodes + uintptr(proxyId)*40))).ｆ__ccgo3_32.Parent
	for parentIndex != -int32(1) {
		v5 = nodes + uintptr(parentIndex)*40
		v6 = aabb
		changed = boolUint8(false1 != 0)
		if v6.LowerBound.X < (*AABB)(unsafe.Pointer(v5)).LowerBound.X {
			(*AABB)(unsafe.Pointer(v5)).LowerBound.X = v6.LowerBound.X
			changed = boolUint8(true1 != 0)
		}
		if v6.LowerBound.Y < (*AABB)(unsafe.Pointer(v5)).LowerBound.Y {
			(*AABB)(unsafe.Pointer(v5)).LowerBound.Y = v6.LowerBound.Y
			changed = boolUint8(true1 != 0)
		}
		if (*AABB)(unsafe.Pointer(v5)).UpperBound.X < v6.UpperBound.X {
			(*AABB)(unsafe.Pointer(v5)).UpperBound.X = v6.UpperBound.X
			changed = boolUint8(true1 != 0)
		}
		if (*AABB)(unsafe.Pointer(v5)).UpperBound.Y < v6.UpperBound.Y {
			(*AABB)(unsafe.Pointer(v5)).UpperBound.Y = v6.UpperBound.Y
			changed = boolUint8(true1 != 0)
		}
		v7 = changed
		goto _8
	_8:
		changed1 = v7
		p9 = nodes + uintptr(parentIndex)*40 + 38
		*(*uint16_t)(unsafe.Pointer(p9)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(p9))) | int32(b2_enlargedNode))
		parentIndex = (*(*b2TreeNode)(unsafe.Pointer(nodes + uintptr(parentIndex)*40))).ｆ__ccgo3_32.Parent
		if int32FromUint8(changed1) == false1 {
			break
		}
	}
	for parentIndex != -int32(1) {
		if int32FromUint16((*(*b2TreeNode)(unsafe.Pointer(nodes + uintptr(parentIndex)*40))).Flags)&int32(b2_enlargedNode) != 0 {
			// early out because this ancestor was previously ascended and marked as enlarged
			break
		}
		p10 = nodes + uintptr(parentIndex)*40 + 38
		*(*uint16_t)(unsafe.Pointer(p10)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(p10))) | int32(b2_enlargedNode))
		parentIndex = (*(*b2TreeNode)(unsafe.Pointer(nodes + uintptr(parentIndex)*40))).ｆ__ccgo3_32.Parent
	}
}

func b2DynamicTree_SetCategoryBits(tls *_Stack, tree uintptr, proxyId int32, categoryBits uint64) {
	var child1, child2, nodeIndex int32
	var node, nodes uintptr
	_, _, _, _, _ = child1, child2, node, nodeIndex, nodes
	nodes = (*DynamicTree)(unsafe.Pointer(tree)).Nodes
	if !((*(*b2TreeNode)(unsafe.Pointer(nodes + uintptr(proxyId)*40))).ｆ__ccgo2_24.Children.Child1 == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+5740, __ccgo_ts+4746, int32FromInt32(872)) != 0 {
		__builtin_trap(tls)
	}
	if !((*(*b2TreeNode)(unsafe.Pointer(nodes + uintptr(proxyId)*40))).ｆ__ccgo2_24.Children.Child2 == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+5788, __ccgo_ts+4746, int32FromInt32(873)) != 0 {
		__builtin_trap(tls)
	}
	if !(int32FromUint16((*(*b2TreeNode)(unsafe.Pointer(nodes + uintptr(proxyId)*40))).Flags)&int32(b2_leafNode) == int32(b2_leafNode)) && b2InternalAssertFcn(tls, __ccgo_ts+5836, __ccgo_ts+4746, int32FromInt32(874)) != 0 {
		__builtin_trap(tls)
	}
	(*(*b2TreeNode)(unsafe.Pointer(nodes + uintptr(proxyId)*40))).CategoryBits = categoryBits
	// Fix up category bits in ancestor internal nodes
	nodeIndex = (*(*b2TreeNode)(unsafe.Pointer(nodes + uintptr(proxyId)*40))).ｆ__ccgo3_32.Parent
	for nodeIndex != -int32(1) {
		node = nodes + uintptr(nodeIndex)*40
		child1 = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child1
		if !(child1 != -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+5180, __ccgo_ts+4746, int32FromInt32(884)) != 0 {
			__builtin_trap(tls)
		}
		child2 = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child2
		if !(child2 != -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+5204, __ccgo_ts+4746, int32FromInt32(886)) != 0 {
			__builtin_trap(tls)
		}
		(*b2TreeNode)(unsafe.Pointer(node)).CategoryBits = (*(*b2TreeNode)(unsafe.Pointer(nodes + uintptr(child1)*40))).CategoryBits | (*(*b2TreeNode)(unsafe.Pointer(nodes + uintptr(child2)*40))).CategoryBits
		nodeIndex = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo3_32.Parent
	}
}

func b2DynamicTree_GetCategoryBits(tls *_Stack, tree uintptr, proxyId int32) (r uint64) {
	if !(0 <= proxyId && proxyId < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) && b2InternalAssertFcn(tls, __ccgo_ts+5468, __ccgo_ts+4746, int32FromInt32(895)) != 0 {
		__builtin_trap(tls)
	}
	return (*(*b2TreeNode)(unsafe.Pointer((*DynamicTree)(unsafe.Pointer(tree)).Nodes + uintptr(proxyId)*40))).CategoryBits
}

func b2DynamicTree_GetHeight(tls *_Stack, tree uintptr) (r int32) {
	if (*DynamicTree)(unsafe.Pointer(tree)).Root == -int32(1) {
		return 0
	}
	return int32FromUint16((*(*b2TreeNode)(unsafe.Pointer((*DynamicTree)(unsafe.Pointer(tree)).Nodes + uintptr((*DynamicTree)(unsafe.Pointer(tree)).Root)*40))).Height)
}

func b2DynamicTree_GetAreaRatio(tls *_Stack, tree uintptr) (r float32) {
	var i int32
	var node, root uintptr
	var rootArea, totalArea, wx, wy, v2, v6 float32
	var v1, v5 AABB
	_, _, _, _, _, _, _, _, _, _, _ = i, node, root, rootArea, totalArea, wx, wy, v1, v2, v5, v6
	if (*DynamicTree)(unsafe.Pointer(tree)).Root == -int32(1) {
		return float32FromFloat32(0)
	}
	root = (*DynamicTree)(unsafe.Pointer(tree)).Nodes + uintptr((*DynamicTree)(unsafe.Pointer(tree)).Root)*40
	v1 = (*b2TreeNode)(unsafe.Pointer(root)).Aabb
	wx = v1.UpperBound.X - v1.LowerBound.X
	wy = v1.UpperBound.Y - v1.LowerBound.Y
	v2 = float32(float32FromFloat32(2) * (wx + wy))
	goto _3
_3:
	rootArea = v2
	totalArea = float32FromFloat32(0)
	i = 0
	for {
		if !(i < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) {
			break
		}
		node = (*DynamicTree)(unsafe.Pointer(tree)).Nodes + uintptr(i)*40
		if int32FromUint8(b2IsAllocated(tls, node)) == false1 || b2IsLeaf(tls, node) != 0 || i == (*DynamicTree)(unsafe.Pointer(tree)).Root {
			goto _4
		}
		v5 = (*b2TreeNode)(unsafe.Pointer(node)).Aabb
		wx = v5.UpperBound.X - v5.LowerBound.X
		wy = v5.UpperBound.Y - v5.LowerBound.Y
		v6 = float32(float32FromFloat32(2) * (wx + wy))
		goto _7
	_7:
		totalArea = totalArea + v6
		goto _4
	_4:
		;
		i = i + 1
	}
	return totalArea / rootArea
}

func b2DynamicTree_GetRootBounds(tls *_Stack, tree uintptr) (r AABB) {
	var empty AABB
	_ = empty
	if (*DynamicTree)(unsafe.Pointer(tree)).Root != -int32(1) {
		return (*(*b2TreeNode)(unsafe.Pointer((*DynamicTree)(unsafe.Pointer(tree)).Nodes + uintptr((*DynamicTree)(unsafe.Pointer(tree)).Root)*40))).Aabb
	}
	empty = AABB{
		LowerBound: b2Vec2_zero,
		UpperBound: b2Vec2_zero,
	}
	return empty
}

func b2DynamicTree_Validate(tls *_Stack, tree uintptr) {
	_ = uint64FromInt64(4)
}

func b2DynamicTree_ValidateNoEnlarged(tls *_Stack, tree uintptr) {
	_ = uint64FromInt64(4)
}

func b2DynamicTree_GetByteCount(tls *_Stack, tree uintptr) (r int32) {
	var size size_t
	_ = size
	size = uint64(72) + uint64(40)*uint64FromInt32((*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) + uint64FromInt32((*DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*(uint64FromInt64(4)+uint64FromInt64(16)+uint64FromInt64(8)+uint64FromInt64(4))
	return int32FromUint64(size)
}

func b2DynamicTree_GetUserData(tls *_Stack, tree uintptr, proxyId int32) (r uint64) {
	if !(0 <= proxyId && proxyId < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) && b2InternalAssertFcn(tls, __ccgo_ts+5468, __ccgo_ts+4746, int32FromInt32(1104)) != 0 {
		__builtin_trap(tls)
	}
	return *(*uint64_t)(unsafe.Add(unsafe.Pointer(&*(*b2TreeNode)(unsafe.Pointer((*DynamicTree)(unsafe.Pointer(tree)).Nodes + uintptr(proxyId)*40))), 24))
}

func b2DynamicTree_GetAABB(tls *_Stack, tree uintptr, proxyId int32) (r AABB) {
	if !(0 <= proxyId && proxyId < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) && b2InternalAssertFcn(tls, __ccgo_ts+5468, __ccgo_ts+4746, int32FromInt32(1110)) != 0 {
		__builtin_trap(tls)
	}
	return (*(*b2TreeNode)(unsafe.Pointer((*DynamicTree)(unsafe.Pointer(tree)).Nodes + uintptr(proxyId)*40))).Aabb
}

func b2DynamicTree_Query(tls *_Stack, tree uintptr, aabb AABB, maskBits uint64, __ccgo_fp_callback uintptr, context uintptr) (r TreeStats) {
	var node uintptr
	var nodeId, stackCount, v1, v2, v7, v8 int32
	var proceed, v5 uint8
	var result TreeStats
	var stack [1024]int32
	var v3, v4 AABB
	_, _, _, _, _, _, _, _, _, _, _, _, _ = node, nodeId, proceed, result, stack, stackCount, v1, v2, v3, v4, v5, v7, v8
	result = TreeStats{}
	if (*DynamicTree)(unsafe.Pointer(tree)).NodeCount == 0 {
		return result
	}
	stackCount = 0
	v1 = stackCount
	stackCount = stackCount + 1
	stack[v1] = (*DynamicTree)(unsafe.Pointer(tree)).Root
	for stackCount > 0 {
		stackCount = stackCount - 1
		v2 = stackCount
		nodeId = stack[v2]
		if nodeId == -int32(1) {
			// todo huh?
			if bool(!(int32FromInt32(false1) != 0)) && b2InternalAssertFcn(tls, __ccgo_ts+4123, __ccgo_ts+4746, int32FromInt32(1134)) != 0 {
				__builtin_trap(tls)
			}
			continue
		}
		node = (*DynamicTree)(unsafe.Pointer(tree)).Nodes + uintptr(nodeId)*40
		result.NodeVisits += int32(1)
		v3 = (*b2TreeNode)(unsafe.Pointer(node)).Aabb
		v4 = aabb
		v5 = boolUint8(!(v4.LowerBound.X > v3.UpperBound.X || v4.LowerBound.Y > v3.UpperBound.Y || v3.LowerBound.X > v4.UpperBound.X || v3.LowerBound.Y > v4.UpperBound.Y))
		goto _6
	_6:
		;
		if v5 != 0 && (*b2TreeNode)(unsafe.Pointer(node)).CategoryBits&maskBits != uint64(0) {
			if b2IsLeaf(tls, node) != 0 {
				// callback to user code with proxy id
				proceed = (*(*func(*_Stack, int32, uint64, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_callback})))(tls, nodeId, *(*uint64_t)(unsafe.Add(unsafe.Pointer(node), 24)), context)
				result.LeafVisits += int32(1)
				if int32FromUint8(proceed) == false1 {
					return result
				}
			} else {
				if stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)-int32FromInt32(1) {
					v7 = stackCount
					stackCount = stackCount + 1
					stack[v7] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child1
					v8 = stackCount
					stackCount = stackCount + 1
					stack[v8] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child2
				} else {
					if !(stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)-int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+5888, __ccgo_ts+4746, int32FromInt32(1163)) != 0 {
						__builtin_trap(tls)
					}
				}
			}
		}
	}
	return result
}

func b2DynamicTree_RayCast(tls *_Stack, tree uintptr, input uintptr, maskBits uint64, __ccgo_fp_callback uintptr, context uintptr) (r1 TreeStats) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var abs_v, b10, b5, b9, c, c1, c11, c2, c21, c3, d, h, n, p1, p2, r, v2, v1, v100, v101, v112, v115, v118, v120, v121, v124, v125, v17, v19, v21, v211, v22, v24, v25, v36, v38, v39, v5, v50, v6, v60, v63, v65, v66, v67, v69, v70, v77, v78, v8, v81, v83, v84, v86, v87, v98 Vec2
	var invLength, length, maxFraction, term1, term2, value, v10, v102, v103, v104, v106, v107, v108, v109, v111, v12, v122, v126, v13, v14, v16, v20, v26, v27, v28, v30, v31, v32, v33, v35, v4, v40, v41, v42, v44, v45, v46, v47, v49, v71, v73, v74, v76, v79, v82, v88, v89, v9, v90, v92, v93, v94, v95, v97 float32
	var node, nodes uintptr
	var nodeAABB, segmentAABB, v114, v117, v54, v55, v59, v62 AABB
	var nodeId, stackCount, v128, v129, v130, v131, v52, v53 int32
	var result TreeStats
	var stack [1024]int32
	var v56 uint8
	var v58 bool
	var _ /* subInput at bp+0 */ RayCastInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = abs_v, b10, b5, b9, c, c1, c11, c2, c21, c3, d, h, invLength, length, maxFraction, n, node, nodeAABB, nodeId, nodes, p1, p2, r, result, segmentAABB, stack, stackCount, term1, term2, v2, value, v1, v10, v100, v101, v102, v103, v104, v106, v107, v108, v109, v111, v112, v114, v115, v117, v118, v12, v120, v121, v122, v124, v125, v126, v128, v129, v13, v130, v131, v14, v16, v17, v19, v21, v20, v211, v22, v24, v25, v26, v27, v28, v30, v31, v32, v33, v35, v36, v38, v39, v4, v40, v41, v42, v44, v45, v46, v47, v49, v5, v50, v52, v53, v54, v55, v56, v58, v59, v6, v60, v62, v63, v65, v66, v67, v69, v70, v71, v73, v74, v76, v77, v78, v79, v8, v81, v82, v83, v84, v86, v87, v88, v89, v9, v90, v92, v93, v94, v95, v97, v98
	result = TreeStats{}
	if (*DynamicTree)(unsafe.Pointer(tree)).NodeCount == 0 {
		return result
	}
	p1 = (*RayCastInput)(unsafe.Pointer(input)).Origin
	d = (*RayCastInput)(unsafe.Pointer(input)).Translation
	v1 = d
	length = sqrtf(tls, float32(v1.X*v1.X)+float32(v1.Y*v1.Y))
	if length < float32FromFloat32(1.1920928955078125e-07) {
		v21 = Vec2{}
		goto _3
	}
	invLength = float32FromFloat32(1) / length
	n = Vec2{
		X: float32(invLength * v1.X),
		Y: float32(invLength * v1.Y),
	}
	v21 = n
	goto _3
_3:
	r = v21
	v4 = float32FromFloat32(1)
	v5 = r
	v6 = Vec2{
		X: float32(-v4 * v5.Y),
		Y: float32(v4 * v5.X),
	}
	goto _7
_7:
	// v is perpendicular to the segment.
	v2 = v6
	v8 = v2
	v9 = v8.X
	if v9 < float32FromInt32(0) {
		v12 = -v9
	} else {
		v12 = v9
	}
	v10 = v12
	goto _11
_11:
	b5.X = v10
	v13 = v8.Y
	if v13 < float32FromInt32(0) {
		v16 = -v13
	} else {
		v16 = v13
	}
	v14 = v16
	goto _15
_15:
	b5.Y = v14
	v17 = b5
	goto _18
_18:
	abs_v = v17
	// Separating axis for segment (Gino, p80).
	// |dot(v, p1 - c)| > dot(|v|, h)
	maxFraction = (*RayCastInput)(unsafe.Pointer(input)).MaxFraction
	v19 = p1
	v20 = maxFraction
	v211 = d
	v22 = Vec2{
		X: v19.X + float32(v20*v211.X),
		Y: v19.Y + float32(v20*v211.Y),
	}
	goto _23
_23:
	p2 = v22
	v24 = p1
	v25 = p2
	v26 = v24.X
	v27 = v25.X
	if v26 < v27 {
		v30 = v26
	} else {
		v30 = v27
	}
	v28 = v30
	goto _29
_29:
	c.X = v28
	v31 = v24.Y
	v32 = v25.Y
	if v31 < v32 {
		v35 = v31
	} else {
		v35 = v32
	}
	v33 = v35
	goto _34
_34:
	c.Y = v33
	v36 = c
	goto _37
_37:
	v38 = p1
	v39 = p2
	v40 = v38.X
	v41 = v39.X
	if v40 > v41 {
		v44 = v40
	} else {
		v44 = v41
	}
	v42 = v44
	goto _43
_43:
	c1.X = v42
	v45 = v38.Y
	v46 = v39.Y
	if v45 > v46 {
		v49 = v45
	} else {
		v49 = v46
	}
	v47 = v49
	goto _48
_48:
	c1.Y = v47
	v50 = c1
	goto _51
_51:
	// Build a bounding box for the segment.
	segmentAABB = AABB{
		LowerBound: v36,
		UpperBound: v50,
	}
	stackCount = 0
	v52 = stackCount
	stackCount = stackCount + 1
	stack[v52] = (*DynamicTree)(unsafe.Pointer(tree)).Root
	nodes = (*DynamicTree)(unsafe.Pointer(tree)).Nodes
	*(*RayCastInput)(unsafe.Pointer(bp)) = *(*RayCastInput)(unsafe.Pointer(input))
	for stackCount > 0 {
		stackCount = stackCount - 1
		v53 = stackCount
		nodeId = stack[v53]
		if nodeId == -int32(1) {
			// todo is this possible?
			if bool(!(int32FromInt32(false1) != 0)) && b2InternalAssertFcn(tls, __ccgo_ts+4123, __ccgo_ts+4746, int32FromInt32(1215)) != 0 {
				__builtin_trap(tls)
			}
			continue
		}
		node = nodes + uintptr(nodeId)*40
		result.NodeVisits += int32(1)
		nodeAABB = (*b2TreeNode)(unsafe.Pointer(node)).Aabb
		if v58 = (*b2TreeNode)(unsafe.Pointer(node)).CategoryBits&maskBits == uint64(0); !v58 {
			v54 = nodeAABB
			v55 = segmentAABB
			v56 = boolUint8(!(v55.LowerBound.X > v54.UpperBound.X || v55.LowerBound.Y > v54.UpperBound.Y || v54.LowerBound.X > v55.UpperBound.X || v54.LowerBound.Y > v55.UpperBound.Y))
			goto _57
		_57:
		}
		if v58 || int32FromUint8(v56) == false1 {
			continue
		}
		v59 = nodeAABB
		b9 = Vec2{
			X: float32(float32FromFloat32(0.5) * (v59.LowerBound.X + v59.UpperBound.X)),
			Y: float32(float32FromFloat32(0.5) * (v59.LowerBound.Y + v59.UpperBound.Y)),
		}
		v60 = b9
		goto _61
	_61:
		// Separating axis for segment (Gino, p80).
		// |dot(v, p1 - c)| > dot(|v|, h)
		// radius extension is added to the node in this case
		c3 = v60
		v62 = nodeAABB
		b10 = Vec2{
			X: float32(float32FromFloat32(0.5) * (v62.UpperBound.X - v62.LowerBound.X)),
			Y: float32(float32FromFloat32(0.5) * (v62.UpperBound.Y - v62.LowerBound.Y)),
		}
		v63 = b10
		goto _64
	_64:
		h = v63
		v65 = p1
		v66 = c3
		v67 = Vec2{
			X: v65.X - v66.X,
			Y: v65.Y - v66.Y,
		}
		goto _68
	_68:
		v69 = v2
		v70 = v67
		v71 = float32(v69.X*v70.X) + float32(v69.Y*v70.Y)
		goto _72
	_72:
		v73 = v71
		if v73 < float32FromInt32(0) {
			v76 = -v73
		} else {
			v76 = v73
		}
		v74 = v76
		goto _75
	_75:
		term1 = v74
		v77 = abs_v
		v78 = h
		v79 = float32(v77.X*v78.X) + float32(v77.Y*v78.Y)
		goto _80
	_80:
		term2 = v79
		if term2 < term1 {
			continue
		}
		if b2IsLeaf(tls, node) != 0 {
			(*(*RayCastInput)(unsafe.Pointer(bp))).MaxFraction = maxFraction
			value = (*(*func(*_Stack, uintptr, int32, uint64, uintptr) float32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_callback})))(tls, bp, nodeId, *(*uint64_t)(unsafe.Add(unsafe.Pointer(node), 24)), context)
			result.LeafVisits += int32(1)
			// The user may return -1 to indicate this shape should be skipped
			if value == float32FromFloat32(0) {
				// The client has terminated the ray cast.
				return result
			}
			if float32FromFloat32(0) < value && value <= maxFraction {
				// Update segment bounding box.
				maxFraction = value
				v81 = p1
				v82 = maxFraction
				v83 = d
				v84 = Vec2{
					X: v81.X + float32(v82*v83.X),
					Y: v81.Y + float32(v82*v83.Y),
				}
				goto _85
			_85:
				p2 = v84
				v86 = p1
				v87 = p2
				v88 = v86.X
				v89 = v87.X
				if v88 < v89 {
					v92 = v88
				} else {
					v92 = v89
				}
				v90 = v92
				goto _91
			_91:
				c.X = v90
				v93 = v86.Y
				v94 = v87.Y
				if v93 < v94 {
					v97 = v93
				} else {
					v97 = v94
				}
				v95 = v97
				goto _96
			_96:
				c.Y = v95
				v98 = c
				goto _99
			_99:
				segmentAABB.LowerBound = v98
				v100 = p1
				v101 = p2
				v102 = v100.X
				v103 = v101.X
				if v102 > v103 {
					v106 = v102
				} else {
					v106 = v103
				}
				v104 = v106
				goto _105
			_105:
				c1.X = v104
				v107 = v100.Y
				v108 = v101.Y
				if v107 > v108 {
					v111 = v107
				} else {
					v111 = v108
				}
				v109 = v111
				goto _110
			_110:
				c1.Y = v109
				v112 = c1
				goto _113
			_113:
				segmentAABB.UpperBound = v112
			}
		} else {
			if stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)-int32FromInt32(1) {
				v114 = (*(*b2TreeNode)(unsafe.Pointer(nodes + uintptr((*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child1)*40))).Aabb
				b9 = Vec2{
					X: float32(float32FromFloat32(0.5) * (v114.LowerBound.X + v114.UpperBound.X)),
					Y: float32(float32FromFloat32(0.5) * (v114.LowerBound.Y + v114.UpperBound.Y)),
				}
				v115 = b9
				goto _116
			_116:
				c11 = v115
				v117 = (*(*b2TreeNode)(unsafe.Pointer(nodes + uintptr((*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child2)*40))).Aabb
				b9 = Vec2{
					X: float32(float32FromFloat32(0.5) * (v117.LowerBound.X + v117.UpperBound.X)),
					Y: float32(float32FromFloat32(0.5) * (v117.LowerBound.Y + v117.UpperBound.Y)),
				}
				v118 = b9
				goto _119
			_119:
				c21 = v118
				v120 = c11
				v121 = p1
				c2 = Vec2{
					X: v121.X - v120.X,
					Y: v121.Y - v120.Y,
				}
				v122 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
				goto _123
			_123:
				v124 = c21
				v125 = p1
				c2 = Vec2{
					X: v125.X - v124.X,
					Y: v125.Y - v124.Y,
				}
				v126 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
				goto _127
			_127:
				if v122 < v126 {
					v128 = stackCount
					stackCount = stackCount + 1
					stack[v128] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child2
					v129 = stackCount
					stackCount = stackCount + 1
					stack[v129] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child1
				} else {
					v130 = stackCount
					stackCount = stackCount + 1
					stack[v130] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child1
					v131 = stackCount
					stackCount = stackCount + 1
					stack[v131] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child2
				}
			} else {
				if !(stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)-int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+5888, __ccgo_ts+4746, int32FromInt32(1284)) != 0 {
					__builtin_trap(tls)
				}
			}
		}
	}
	return result
}

func b2DynamicTree_ShapeCast(tls *_Stack, tree uintptr, input uintptr, maskBits uint64, __ccgo_fp_callback uintptr, context uintptr) (r1 TreeStats) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var abs_v, b10, b5, b9, c, c1, c11, c2, c21, c3, extension, h, p1, r, radius, t, v2, v107, v110, v112, v113, v114, v116, v117, v118, v120, v121, v128, v129, v133, v134, v136, v137, v138, v14, v140, v141, v152, v154, v155, v156, v158, v159, v16, v17, v170, v173, v176, v178, v179, v182, v183, v21, v28, v3, v30, v31, v32, v34, v35, v36, v39, v42, v45, v46, v48, v57, v60, v61, v63, v64, v65, v67, v68, v79, v81, v82, v83, v85, v86, v97 Vec2
	var i, nodeId, stackCount, v100, v186, v187, v188, v189, v99 int32
	var maxFraction, term1, term2, value, v10, v11, v122, v124, v125, v127, v13, v130, v132, v142, v143, v144, v146, v147, v148, v149, v151, v160, v161, v162, v164, v165, v166, v167, v169, v18, v180, v184, v19, v20, v22, v23, v24, v25, v27, v4, v44, v49, v5, v50, v52, v53, v54, v56, v59, v6, v69, v70, v71, v73, v74, v75, v76, v78, v8, v87, v88, v89, v9, v91, v92, v93, v94, v96 float32
	var node, nodes uintptr
	var originAABB, totalAABB, v101, v102, v106, v109, v172, v175, v38, v41 AABB
	var stack [1024]int32
	var stats TreeStats
	var v103 uint8
	var v105 bool
	var _ /* subInput at bp+0 */ ShapeCastInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = abs_v, b10, b5, b9, c, c1, c11, c2, c21, c3, extension, h, i, maxFraction, node, nodeId, nodes, originAABB, p1, r, radius, stack, stackCount, stats, t, term1, term2, totalAABB, v2, value, v10, v100, v101, v102, v103, v105, v106, v107, v109, v11, v110, v112, v113, v114, v116, v117, v118, v120, v121, v122, v124, v125, v127, v128, v129, v13, v130, v132, v133, v134, v136, v137, v138, v14, v140, v141, v142, v143, v144, v146, v147, v148, v149, v151, v152, v154, v155, v156, v158, v159, v16, v160, v161, v162, v164, v165, v166, v167, v169, v17, v170, v172, v173, v175, v176, v178, v179, v18, v180, v182, v183, v184, v186, v187, v188, v189, v19, v21, v20, v22, v23, v24, v25, v27, v28, v3, v30, v31, v32, v34, v35, v36, v38, v39, v4, v41, v42, v44, v45, v46, v48, v49, v5, v50, v52, v53, v54, v56, v57, v59, v6, v60, v61, v63, v64, v65, v67, v68, v69, v70, v71, v73, v74, v75, v76, v78, v79, v8, v81, v82, v83, v85, v86, v87, v88, v89, v9, v91, v92, v93, v94, v96, v97, v99
	stats = TreeStats{}
	if (*DynamicTree)(unsafe.Pointer(tree)).NodeCount == 0 || (*ShapeCastInput)(unsafe.Pointer(input)).Proxy.Count == 0 {
		return stats
	}
	originAABB = AABB{
		LowerBound: *(*Vec2)(unsafe.Pointer(input)),
		UpperBound: *(*Vec2)(unsafe.Pointer(input)),
	}
	i = int32(1)
	for {
		if !(i < (*ShapeCastInput)(unsafe.Pointer(input)).Proxy.Count) {
			break
		}
		v21 = originAABB.LowerBound
		v3 = *(*Vec2)(unsafe.Pointer(input + uintptr(i)*8))
		v4 = v21.X
		v5 = v3.X
		if v4 < v5 {
			v8 = v4
		} else {
			v8 = v5
		}
		v6 = v8
		goto _7
	_7:
		c.X = v6
		v9 = v21.Y
		v10 = v3.Y
		if v9 < v10 {
			v13 = v9
		} else {
			v13 = v10
		}
		v11 = v13
		goto _12
	_12:
		c.Y = v11
		v14 = c
		goto _15
	_15:
		originAABB.LowerBound = v14
		v16 = originAABB.UpperBound
		v17 = *(*Vec2)(unsafe.Pointer(input + uintptr(i)*8))
		v18 = v16.X
		v19 = v17.X
		if v18 > v19 {
			v22 = v18
		} else {
			v22 = v19
		}
		v20 = v22
		goto _21
	_21:
		c1.X = v20
		v23 = v16.Y
		v24 = v17.Y
		if v23 > v24 {
			v27 = v23
		} else {
			v27 = v24
		}
		v25 = v27
		goto _26
	_26:
		c1.Y = v25
		v28 = c1
		goto _29
	_29:
		originAABB.UpperBound = v28
		goto _1
	_1:
		;
		i = i + 1
	}
	radius = Vec2{
		X: (*ShapeCastInput)(unsafe.Pointer(input)).Proxy.Radius,
		Y: (*ShapeCastInput)(unsafe.Pointer(input)).Proxy.Radius,
	}
	v30 = originAABB.LowerBound
	v31 = radius
	v32 = Vec2{
		X: v30.X - v31.X,
		Y: v30.Y - v31.Y,
	}
	goto _33
_33:
	originAABB.LowerBound = v32
	v34 = originAABB.UpperBound
	v35 = radius
	v36 = Vec2{
		X: v34.X + v35.X,
		Y: v34.Y + v35.Y,
	}
	goto _37
_37:
	originAABB.UpperBound = v36
	v38 = originAABB
	b9 = Vec2{
		X: float32(float32FromFloat32(0.5) * (v38.LowerBound.X + v38.UpperBound.X)),
		Y: float32(float32FromFloat32(0.5) * (v38.LowerBound.Y + v38.UpperBound.Y)),
	}
	v39 = b9
	goto _40
_40:
	p1 = v39
	v41 = originAABB
	b10 = Vec2{
		X: float32(float32FromFloat32(0.5) * (v41.UpperBound.X - v41.LowerBound.X)),
		Y: float32(float32FromFloat32(0.5) * (v41.UpperBound.Y - v41.LowerBound.Y)),
	}
	v42 = b10
	goto _43
_43:
	extension = v42
	// v is perpendicular to the segment.
	r = (*ShapeCastInput)(unsafe.Pointer(input)).Translation
	v44 = float32FromFloat32(1)
	v45 = r
	v46 = Vec2{
		X: float32(-v44 * v45.Y),
		Y: float32(v44 * v45.X),
	}
	goto _47
_47:
	v2 = v46
	v48 = v2
	v49 = v48.X
	if v49 < float32FromInt32(0) {
		v52 = -v49
	} else {
		v52 = v49
	}
	v50 = v52
	goto _51
_51:
	b5.X = v50
	v53 = v48.Y
	if v53 < float32FromInt32(0) {
		v56 = -v53
	} else {
		v56 = v53
	}
	v54 = v56
	goto _55
_55:
	b5.Y = v54
	v57 = b5
	goto _58
_58:
	abs_v = v57
	// Separating axis for segment (Gino, p80).
	// |dot(v, p1 - c)| > dot(|v|, h)
	maxFraction = (*ShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	v59 = maxFraction
	v60 = (*ShapeCastInput)(unsafe.Pointer(input)).Translation
	v61 = Vec2{
		X: float32(v59 * v60.X),
		Y: float32(v59 * v60.Y),
	}
	goto _62
_62:
	// Build total box for the shape cast
	t = v61
	v63 = originAABB.LowerBound
	v64 = t
	v65 = Vec2{
		X: v63.X + v64.X,
		Y: v63.Y + v64.Y,
	}
	goto _66
_66:
	v67 = originAABB.LowerBound
	v68 = v65
	v69 = v67.X
	v70 = v68.X
	if v69 < v70 {
		v73 = v69
	} else {
		v73 = v70
	}
	v71 = v73
	goto _72
_72:
	c.X = v71
	v74 = v67.Y
	v75 = v68.Y
	if v74 < v75 {
		v78 = v74
	} else {
		v78 = v75
	}
	v76 = v78
	goto _77
_77:
	c.Y = v76
	v79 = c
	goto _80
_80:
	v81 = originAABB.UpperBound
	v82 = t
	v83 = Vec2{
		X: v81.X + v82.X,
		Y: v81.Y + v82.Y,
	}
	goto _84
_84:
	v85 = originAABB.UpperBound
	v86 = v83
	v87 = v85.X
	v88 = v86.X
	if v87 > v88 {
		v91 = v87
	} else {
		v91 = v88
	}
	v89 = v91
	goto _90
_90:
	c1.X = v89
	v92 = v85.Y
	v93 = v86.Y
	if v92 > v93 {
		v96 = v92
	} else {
		v96 = v93
	}
	v94 = v96
	goto _95
_95:
	c1.Y = v94
	v97 = c1
	goto _98
_98:
	totalAABB = AABB{
		LowerBound: v79,
		UpperBound: v97,
	}
	*(*ShapeCastInput)(unsafe.Pointer(bp)) = *(*ShapeCastInput)(unsafe.Pointer(input))
	nodes = (*DynamicTree)(unsafe.Pointer(tree)).Nodes
	stackCount = 0
	v99 = stackCount
	stackCount = stackCount + 1
	stack[v99] = (*DynamicTree)(unsafe.Pointer(tree)).Root
	for stackCount > 0 {
		stackCount = stackCount - 1
		v100 = stackCount
		nodeId = stack[v100]
		if nodeId == -int32(1) {
			// todo is this possible?
			if bool(!(int32FromInt32(false1) != 0)) && b2InternalAssertFcn(tls, __ccgo_ts+4123, __ccgo_ts+4746, int32FromInt32(1347)) != 0 {
				__builtin_trap(tls)
			}
			continue
		}
		node = nodes + uintptr(nodeId)*40
		stats.NodeVisits += int32(1)
		if v105 = (*b2TreeNode)(unsafe.Pointer(node)).CategoryBits&maskBits == uint64(0); !v105 {
			v101 = (*b2TreeNode)(unsafe.Pointer(node)).Aabb
			v102 = totalAABB
			v103 = boolUint8(!(v102.LowerBound.X > v101.UpperBound.X || v102.LowerBound.Y > v101.UpperBound.Y || v101.LowerBound.X > v102.UpperBound.X || v101.LowerBound.Y > v102.UpperBound.Y))
			goto _104
		_104:
		}
		if v105 || int32FromUint8(v103) == false1 {
			continue
		}
		v106 = (*b2TreeNode)(unsafe.Pointer(node)).Aabb
		b9 = Vec2{
			X: float32(float32FromFloat32(0.5) * (v106.LowerBound.X + v106.UpperBound.X)),
			Y: float32(float32FromFloat32(0.5) * (v106.LowerBound.Y + v106.UpperBound.Y)),
		}
		v107 = b9
		goto _108
	_108:
		// Separating axis for segment (Gino, p80).
		// |dot(v, p1 - c)| > dot(|v|, h)
		// radius extension is added to the node in this case
		c3 = v107
		v109 = (*b2TreeNode)(unsafe.Pointer(node)).Aabb
		b10 = Vec2{
			X: float32(float32FromFloat32(0.5) * (v109.UpperBound.X - v109.LowerBound.X)),
			Y: float32(float32FromFloat32(0.5) * (v109.UpperBound.Y - v109.LowerBound.Y)),
		}
		v110 = b10
		goto _111
	_111:
		v112 = v110
		v113 = extension
		v114 = Vec2{
			X: v112.X + v113.X,
			Y: v112.Y + v113.Y,
		}
		goto _115
	_115:
		h = v114
		v116 = p1
		v117 = c3
		v118 = Vec2{
			X: v116.X - v117.X,
			Y: v116.Y - v117.Y,
		}
		goto _119
	_119:
		v120 = v2
		v121 = v118
		v122 = float32(v120.X*v121.X) + float32(v120.Y*v121.Y)
		goto _123
	_123:
		v124 = v122
		if v124 < float32FromInt32(0) {
			v127 = -v124
		} else {
			v127 = v124
		}
		v125 = v127
		goto _126
	_126:
		term1 = v125
		v128 = abs_v
		v129 = h
		v130 = float32(v128.X*v129.X) + float32(v128.Y*v129.Y)
		goto _131
	_131:
		term2 = v130
		if term2 < term1 {
			continue
		}
		if b2IsLeaf(tls, node) != 0 {
			(*(*ShapeCastInput)(unsafe.Pointer(bp))).MaxFraction = maxFraction
			value = (*(*func(*_Stack, uintptr, int32, uint64, uintptr) float32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_callback})))(tls, bp, nodeId, *(*uint64_t)(unsafe.Add(unsafe.Pointer(node), 24)), context)
			stats.LeafVisits += int32(1)
			if value == float32FromFloat32(0) {
				// The client has terminated the ray cast.
				return stats
			}
			if float32FromFloat32(0) < value && value < maxFraction {
				// Update segment bounding box.
				maxFraction = value
				v132 = maxFraction
				v133 = (*ShapeCastInput)(unsafe.Pointer(input)).Translation
				v134 = Vec2{
					X: float32(v132 * v133.X),
					Y: float32(v132 * v133.Y),
				}
				goto _135
			_135:
				t = v134
				v136 = originAABB.LowerBound
				v137 = t
				v138 = Vec2{
					X: v136.X + v137.X,
					Y: v136.Y + v137.Y,
				}
				goto _139
			_139:
				v140 = originAABB.LowerBound
				v141 = v138
				v142 = v140.X
				v143 = v141.X
				if v142 < v143 {
					v146 = v142
				} else {
					v146 = v143
				}
				v144 = v146
				goto _145
			_145:
				c.X = v144
				v147 = v140.Y
				v148 = v141.Y
				if v147 < v148 {
					v151 = v147
				} else {
					v151 = v148
				}
				v149 = v151
				goto _150
			_150:
				c.Y = v149
				v152 = c
				goto _153
			_153:
				totalAABB.LowerBound = v152
				v154 = originAABB.UpperBound
				v155 = t
				v156 = Vec2{
					X: v154.X + v155.X,
					Y: v154.Y + v155.Y,
				}
				goto _157
			_157:
				v158 = originAABB.UpperBound
				v159 = v156
				v160 = v158.X
				v161 = v159.X
				if v160 > v161 {
					v164 = v160
				} else {
					v164 = v161
				}
				v162 = v164
				goto _163
			_163:
				c1.X = v162
				v165 = v158.Y
				v166 = v159.Y
				if v165 > v166 {
					v169 = v165
				} else {
					v169 = v166
				}
				v167 = v169
				goto _168
			_168:
				c1.Y = v167
				v170 = c1
				goto _171
			_171:
				totalAABB.UpperBound = v170
			}
		} else {
			if stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)-int32FromInt32(1) {
				v172 = (*(*b2TreeNode)(unsafe.Pointer(nodes + uintptr((*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child1)*40))).Aabb
				b9 = Vec2{
					X: float32(float32FromFloat32(0.5) * (v172.LowerBound.X + v172.UpperBound.X)),
					Y: float32(float32FromFloat32(0.5) * (v172.LowerBound.Y + v172.UpperBound.Y)),
				}
				v173 = b9
				goto _174
			_174:
				c11 = v173
				v175 = (*(*b2TreeNode)(unsafe.Pointer(nodes + uintptr((*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child2)*40))).Aabb
				b9 = Vec2{
					X: float32(float32FromFloat32(0.5) * (v175.LowerBound.X + v175.UpperBound.X)),
					Y: float32(float32FromFloat32(0.5) * (v175.LowerBound.Y + v175.UpperBound.Y)),
				}
				v176 = b9
				goto _177
			_177:
				c21 = v176
				v178 = c11
				v179 = p1
				c2 = Vec2{
					X: v179.X - v178.X,
					Y: v179.Y - v178.Y,
				}
				v180 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
				goto _181
			_181:
				v182 = c21
				v183 = p1
				c2 = Vec2{
					X: v183.X - v182.X,
					Y: v183.Y - v182.Y,
				}
				v184 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
				goto _185
			_185:
				if v180 < v184 {
					v186 = stackCount
					stackCount = stackCount + 1
					stack[v186] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child2
					v187 = stackCount
					stackCount = stackCount + 1
					stack[v187] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child1
				} else {
					v188 = stackCount
					stackCount = stackCount + 1
					stack[v188] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child1
					v189 = stackCount
					stackCount = stackCount + 1
					stack[v189] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child2
				}
			} else {
				if !(stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)-int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+5888, __ccgo_ts+4746, int32FromInt32(1412)) != 0 {
					__builtin_trap(tls)
				}
			}
		}
	}
	return stats
}

// C documentation
//
//	// Not safe to access tree during this operation because it may grow
func b2DynamicTree_Rebuild(tls *_Stack, tree uintptr, fullBuild uint8) (r int32) {
	var b, v2 Vec2
	var doomedNodeIndex, leafCount, newCapacity, nodeIndex, proxyCount, stackCount, v4, v5 int32
	var leafCenters, leafIndices, node, nodes uintptr
	var stack [1024]int32
	var v1 AABB
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b, doomedNodeIndex, leafCenters, leafCount, leafIndices, newCapacity, node, nodeIndex, nodes, proxyCount, stack, stackCount, v1, v2, v4, v5
	proxyCount = (*DynamicTree)(unsafe.Pointer(tree)).ProxyCount
	if proxyCount == 0 {
		return 0
	}
	// Ensure capacity for rebuild space
	if proxyCount > (*DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity {
		newCapacity = proxyCount + proxyCount/int32(2)
		b2Free(tls, (*DynamicTree)(unsafe.Pointer(tree)).LeafIndices, int32FromUint64(uint64FromInt32((*DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(4)))
		(*DynamicTree)(unsafe.Pointer(tree)).LeafIndices = b2Alloc(tls, int32FromUint64(uint64FromInt32(newCapacity)*uint64(4)))
		b2Free(tls, (*DynamicTree)(unsafe.Pointer(tree)).LeafCenters, int32FromUint64(uint64FromInt32((*DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(8)))
		(*DynamicTree)(unsafe.Pointer(tree)).LeafCenters = b2Alloc(tls, int32FromUint64(uint64FromInt32(newCapacity)*uint64(8)))
		(*DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity = newCapacity
	}
	leafCount = 0
	stackCount = 0
	nodeIndex = (*DynamicTree)(unsafe.Pointer(tree)).Root
	nodes = (*DynamicTree)(unsafe.Pointer(tree)).Nodes
	node = nodes + uintptr(nodeIndex)*40
	// These are the nodes that get sorted to rebuild the tree.
	// I'm using indices because the node pool may grow during the build.
	leafIndices = (*DynamicTree)(unsafe.Pointer(tree)).LeafIndices
	leafCenters = (*DynamicTree)(unsafe.Pointer(tree)).LeafCenters
	// Gather all proxy nodes that have grown and all internal nodes that haven't grown. Both are
	// considered leaves in the tree rebuild.
	// Free all internal nodes that have grown.
	// todo use a node growth metric instead of simply enlarged to reduce rebuild size and frequency
	// this should be weighed against B2_AABB_MARGIN
	for int32(true1) != 0 {
		if int32FromUint16((*b2TreeNode)(unsafe.Pointer(node)).Height) == 0 || int32FromUint16((*b2TreeNode)(unsafe.Pointer(node)).Flags)&int32(b2_enlargedNode) == 0 && int32FromUint8(fullBuild) == false1 {
			*(*int32)(unsafe.Pointer(leafIndices + uintptr(leafCount)*4)) = nodeIndex
			v1 = (*b2TreeNode)(unsafe.Pointer(node)).Aabb
			b = Vec2{
				X: float32(float32FromFloat32(0.5) * (v1.LowerBound.X + v1.UpperBound.X)),
				Y: float32(float32FromFloat32(0.5) * (v1.LowerBound.Y + v1.UpperBound.Y)),
			}
			v2 = b
			goto _3
		_3:
			*(*Vec2)(unsafe.Pointer(leafCenters + uintptr(leafCount)*8)) = v2
			leafCount = leafCount + int32(1)
			// Detach
			(*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo3_32.Parent = -int32FromInt32(1)
		} else {
			doomedNodeIndex = nodeIndex
			// Handle children
			nodeIndex = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child1
			if stackCount < int32(_B2_TREE_STACK_SIZE) {
				v4 = stackCount
				stackCount = stackCount + 1
				stack[v4] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child2
			} else {
				if !(stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)) && b2InternalAssertFcn(tls, __ccgo_ts+6449, __ccgo_ts+4746, int32FromInt32(1951)) != 0 {
					__builtin_trap(tls)
				}
			}
			node = nodes + uintptr(nodeIndex)*40
			// Remove doomed node
			b2FreeNode(tls, tree, doomedNodeIndex)
			continue
		}
		if stackCount == 0 {
			break
		}
		stackCount = stackCount - 1
		v5 = stackCount
		nodeIndex = stack[v5]
		node = nodes + uintptr(nodeIndex)*40
	}
	if !(leafCount <= proxyCount) && b2InternalAssertFcn(tls, __ccgo_ts+6481, __ccgo_ts+4746, int32FromInt32(1982)) != 0 {
		__builtin_trap(tls)
	}
	(*DynamicTree)(unsafe.Pointer(tree)).Root = b2BuildTree(tls, tree, leafCount)
	b2DynamicTree_Validate(tls, tree)
	return leafCount
}
