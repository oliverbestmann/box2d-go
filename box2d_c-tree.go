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
		*(*int32_t)(unsafe.Add(unsafe.Pointer(&*(*b2TreeNode)(unsafe.Add(nil, tree.Nodes+uintptr(i)*40))), 32)) = i + int32(1)
		goto _1
	_1:
		;
		i = i + 1
	}
	*(*int32_t)(unsafe.Add(unsafe.Pointer(&*(*b2TreeNode)(unsafe.Add(nil, tree.Nodes+uintptr(tree.NodeCapacity-int32(1))*40))), 32)) = -int32FromInt32(1)
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
	*(*int32)(unsafe.Add(nil, tree+24)) += int32(1)
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
	*(*int32)(unsafe.Add(nil, tree+24)) -= int32(1)
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
	(*(*b2TreeNode)(unsafe.Add(nil, (*DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr(proxyId)*40))).Aabb = aabb
	shouldRotate = boolUint8(false1 != 0)
	b2InsertLeaf(tls, tree, proxyId, shouldRotate)
}

func b2DynamicTree_EnlargeProxy(tls *_Stack, tree uintptr, proxyId int32, aabb AABB) {
	var changed, changed1, s, v3 uint8
	var nodes, v5, v9 uintptr
	var parentIndex int32
	var v1, v2 AABB
	_, _, _, _, _, _, _, _, _, _ = changed, changed1, nodes, parentIndex, s, v1, v2, v3, v5, v9
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
	v1 = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(proxyId)*40))).Aabb
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
	(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(proxyId)*40))).Aabb = aabb
	parentIndex = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(proxyId)*40))).ｆ__ccgo3_32.Parent
	for parentIndex != -int32(1) {
		v5 = nodes + uintptr(parentIndex)*40
		v1 = aabb
		changed = boolUint8(false1 != 0)
		if v1.LowerBound.X < (*AABB)(unsafe.Pointer(v5)).LowerBound.X {
			(*AABB)(unsafe.Pointer(v5)).LowerBound.X = v1.LowerBound.X
			changed = boolUint8(true1 != 0)
		}
		if v1.LowerBound.Y < (*AABB)(unsafe.Pointer(v5)).LowerBound.Y {
			(*AABB)(unsafe.Pointer(v5)).LowerBound.Y = v1.LowerBound.Y
			changed = boolUint8(true1 != 0)
		}
		if (*AABB)(unsafe.Pointer(v5)).UpperBound.X < v1.UpperBound.X {
			(*AABB)(unsafe.Pointer(v5)).UpperBound.X = v1.UpperBound.X
			changed = boolUint8(true1 != 0)
		}
		if (*AABB)(unsafe.Pointer(v5)).UpperBound.Y < v1.UpperBound.Y {
			(*AABB)(unsafe.Pointer(v5)).UpperBound.Y = v1.UpperBound.Y
			changed = boolUint8(true1 != 0)
		}
		v3 = changed
		goto _8
	_8:
		changed1 = v3
		v9 = nodes + uintptr(parentIndex)*40 + 38
		*(*uint16_t)(unsafe.Pointer(v9)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v9))) | int32(b2_enlargedNode))
		parentIndex = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(parentIndex)*40))).ｆ__ccgo3_32.Parent
		if int32FromUint8(changed1) == false1 {
			break
		}
	}
	for parentIndex != -int32(1) {
		if int32FromUint16((*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(parentIndex)*40))).Flags)&int32(b2_enlargedNode) != 0 {
			// early out because this ancestor was previously ascended and marked as enlarged
			break
		}
		v5 = nodes + uintptr(parentIndex)*40 + 38
		*(*uint16_t)(unsafe.Pointer(v5)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v5))) | int32(b2_enlargedNode))
		parentIndex = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(parentIndex)*40))).ｆ__ccgo3_32.Parent
	}
}

func b2DynamicTree_SetCategoryBits(tls *_Stack, tree uintptr, proxyId int32, categoryBits uint64) {
	var child1, child2, nodeIndex int32
	var node, nodes uintptr
	_, _, _, _, _ = child1, child2, node, nodeIndex, nodes
	nodes = (*DynamicTree)(unsafe.Pointer(tree)).Nodes
	if !((*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(proxyId)*40))).ｆ__ccgo2_24.Children.Child1 == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+5740, __ccgo_ts+4746, int32FromInt32(872)) != 0 {
		__builtin_trap(tls)
	}
	if !((*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(proxyId)*40))).ｆ__ccgo2_24.Children.Child2 == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+5788, __ccgo_ts+4746, int32FromInt32(873)) != 0 {
		__builtin_trap(tls)
	}
	if !(int32FromUint16((*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(proxyId)*40))).Flags)&int32(b2_leafNode) == int32(b2_leafNode)) && b2InternalAssertFcn(tls, __ccgo_ts+5836, __ccgo_ts+4746, int32FromInt32(874)) != 0 {
		__builtin_trap(tls)
	}
	(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(proxyId)*40))).CategoryBits = categoryBits
	// Fix up category bits in ancestor internal nodes
	nodeIndex = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(proxyId)*40))).ｆ__ccgo3_32.Parent
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
		(*b2TreeNode)(unsafe.Pointer(node)).CategoryBits = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(child1)*40))).CategoryBits | (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(child2)*40))).CategoryBits
		nodeIndex = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo3_32.Parent
	}
}

func b2DynamicTree_GetCategoryBits(tls *_Stack, tree uintptr, proxyId int32) (r uint64) {
	if !(0 <= proxyId && proxyId < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) && b2InternalAssertFcn(tls, __ccgo_ts+5468, __ccgo_ts+4746, int32FromInt32(895)) != 0 {
		__builtin_trap(tls)
	}
	return (*(*b2TreeNode)(unsafe.Add(nil, (*DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr(proxyId)*40))).CategoryBits
}

func b2DynamicTree_GetHeight(tls *_Stack, tree uintptr) (r int32) {
	if (*DynamicTree)(unsafe.Pointer(tree)).Root == -int32(1) {
		return 0
	}
	return int32FromUint16((*(*b2TreeNode)(unsafe.Add(nil, (*DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr((*DynamicTree)(unsafe.Pointer(tree)).Root)*40))).Height)
}

func b2DynamicTree_GetAreaRatio(tls *_Stack, tree uintptr) (r float32) {
	var i int32
	var node, root uintptr
	var rootArea, totalArea, wx, wy, v2 float32
	var v1 AABB
	_, _, _, _, _, _, _, _, _ = i, node, root, rootArea, totalArea, wx, wy, v1, v2
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
		v1 = (*b2TreeNode)(unsafe.Pointer(node)).Aabb
		wx = v1.UpperBound.X - v1.LowerBound.X
		wy = v1.UpperBound.Y - v1.LowerBound.Y
		v2 = float32(float32FromFloat32(2) * (wx + wy))
		goto _7
	_7:
		totalArea = totalArea + v2
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
		return (*(*b2TreeNode)(unsafe.Add(nil, (*DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr((*DynamicTree)(unsafe.Pointer(tree)).Root)*40))).Aabb
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
	return *(*uint64_t)(unsafe.Add(unsafe.Pointer(&*(*b2TreeNode)(unsafe.Add(nil, (*DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr(proxyId)*40))), 24))
}

func b2DynamicTree_GetAABB(tls *_Stack, tree uintptr, proxyId int32) (r AABB) {
	if !(0 <= proxyId && proxyId < (*DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) && b2InternalAssertFcn(tls, __ccgo_ts+5468, __ccgo_ts+4746, int32FromInt32(1110)) != 0 {
		__builtin_trap(tls)
	}
	return (*(*b2TreeNode)(unsafe.Add(nil, (*DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr(proxyId)*40))).Aabb
}

func b2DynamicTree_Query(tls *_Stack, tree uintptr, aabb AABB, maskBits uint64, __ccgo_fp_callback uintptr, context uintptr) (r TreeStats) {
	var node uintptr
	var nodeId, stackCount, v1 int32
	var proceed, v5 uint8
	var result TreeStats
	var stack [1024]int32
	var v3, v4 AABB
	_, _, _, _, _, _, _, _, _, _ = node, nodeId, proceed, result, stack, stackCount, v1, v3, v4, v5
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
		v1 = stackCount
		nodeId = stack[v1]
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
					v1 = stackCount
					stackCount = stackCount + 1
					stack[v1] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child1
					v1 = stackCount
					stackCount = stackCount + 1
					stack[v1] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child2
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
	var abs_v, b10, b5, b9, c, c1, c11, c2, c21, c3, d, h, n, p1, p2, r, v2, v1, v17, v19, v21, v211, v22, v24, v25, v36, v38, v39, v5, v50, v6, v8 Vec2
	var invLength, length, maxFraction, term1, term2, value, v10, v12, v13, v14, v16, v20, v26, v27, v28, v30, v31, v32, v33, v35, v4, v40, v41, v42, v44, v45, v46, v47, v49, v9 float32
	var node, nodes uintptr
	var nodeAABB, segmentAABB, v54, v55 AABB
	var nodeId, stackCount, v52 int32
	var result TreeStats
	var stack [1024]int32
	var v56 uint8
	var v58 bool
	var _ /* subInput at bp+0 */ RayCastInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = abs_v, b10, b5, b9, c, c1, c11, c2, c21, c3, d, h, invLength, length, maxFraction, n, node, nodeAABB, nodeId, nodes, p1, p2, r, result, segmentAABB, stack, stackCount, term1, term2, v2, value, v1, v10, v12, v13, v14, v16, v17, v19, v21, v20, v211, v22, v24, v25, v26, v27, v28, v30, v31, v32, v33, v35, v36, v38, v39, v4, v40, v41, v42, v44, v45, v46, v47, v49, v5, v50, v52, v54, v55, v56, v58, v6, v8, v9
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
	*(*RayCastInput)(unsafe.Add(nil, bp)) = *(*RayCastInput)(unsafe.Add(nil, input))
	for stackCount > 0 {
		stackCount = stackCount - 1
		v52 = stackCount
		nodeId = stack[v52]
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
		v54 = nodeAABB
		b9 = Vec2{
			X: float32(float32FromFloat32(0.5) * (v54.LowerBound.X + v54.UpperBound.X)),
			Y: float32(float32FromFloat32(0.5) * (v54.LowerBound.Y + v54.UpperBound.Y)),
		}
		v1 = b9
		goto _61
	_61:
		// Separating axis for segment (Gino, p80).
		// |dot(v, p1 - c)| > dot(|v|, h)
		// radius extension is added to the node in this case
		c3 = v1
		v55 = nodeAABB
		b10 = Vec2{
			X: float32(float32FromFloat32(0.5) * (v55.UpperBound.X - v55.LowerBound.X)),
			Y: float32(float32FromFloat32(0.5) * (v55.UpperBound.Y - v55.LowerBound.Y)),
		}
		v21 = b10
		goto _64
	_64:
		h = v21
		v5 = p1
		v6 = c3
		v8 = Vec2{
			X: v5.X - v6.X,
			Y: v5.Y - v6.Y,
		}
		goto _68
	_68:
		v17 = v2
		v19 = v8
		v4 = float32(v17.X*v19.X) + float32(v17.Y*v19.Y)
		goto _72
	_72:
		v9 = v4
		if v9 < float32FromInt32(0) {
			v12 = -v9
		} else {
			v12 = v9
		}
		v10 = v12
		goto _75
	_75:
		term1 = v10
		v211 = abs_v
		v22 = h
		v13 = float32(v211.X*v22.X) + float32(v211.Y*v22.Y)
		goto _80
	_80:
		term2 = v13
		if term2 < term1 {
			continue
		}
		if b2IsLeaf(tls, node) != 0 {
			(*(*RayCastInput)(unsafe.Add(nil, bp))).MaxFraction = maxFraction
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
				v1 = p1
				v4 = maxFraction
				v21 = d
				v5 = Vec2{
					X: v1.X + float32(v4*v21.X),
					Y: v1.Y + float32(v4*v21.Y),
				}
				goto _85
			_85:
				p2 = v5
				v1 = p1
				v21 = p2
				v4 = v1.X
				v9 = v21.X
				if v4 < v9 {
					v12 = v4
				} else {
					v12 = v9
				}
				v10 = v12
				goto _91
			_91:
				c.X = v10
				v13 = v1.Y
				v14 = v21.Y
				if v13 < v14 {
					v20 = v13
				} else {
					v20 = v14
				}
				v16 = v20
				goto _96
			_96:
				c.Y = v16
				v5 = c
				goto _99
			_99:
				segmentAABB.LowerBound = v5
				v1 = p1
				v21 = p2
				v4 = v1.X
				v9 = v21.X
				if v4 > v9 {
					v12 = v4
				} else {
					v12 = v9
				}
				v10 = v12
				goto _105
			_105:
				c1.X = v10
				v13 = v1.Y
				v14 = v21.Y
				if v13 > v14 {
					v20 = v13
				} else {
					v20 = v14
				}
				v16 = v20
				goto _110
			_110:
				c1.Y = v16
				v5 = c1
				goto _113
			_113:
				segmentAABB.UpperBound = v5
			}
		} else {
			if stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)-int32FromInt32(1) {
				v54 = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr((*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child1)*40))).Aabb
				b9 = Vec2{
					X: float32(float32FromFloat32(0.5) * (v54.LowerBound.X + v54.UpperBound.X)),
					Y: float32(float32FromFloat32(0.5) * (v54.LowerBound.Y + v54.UpperBound.Y)),
				}
				v1 = b9
				goto _116
			_116:
				c11 = v1
				v55 = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr((*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child2)*40))).Aabb
				b9 = Vec2{
					X: float32(float32FromFloat32(0.5) * (v55.LowerBound.X + v55.UpperBound.X)),
					Y: float32(float32FromFloat32(0.5) * (v55.LowerBound.Y + v55.UpperBound.Y)),
				}
				v21 = b9
				goto _119
			_119:
				c21 = v21
				v5 = c11
				v6 = p1
				c2 = Vec2{
					X: v6.X - v5.X,
					Y: v6.Y - v5.Y,
				}
				v4 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
				goto _123
			_123:
				v8 = c21
				v17 = p1
				c2 = Vec2{
					X: v17.X - v8.X,
					Y: v17.Y - v8.Y,
				}
				v9 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
				goto _127
			_127:
				if v4 < v9 {
					v52 = stackCount
					stackCount = stackCount + 1
					stack[v52] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child2
					v52 = stackCount
					stackCount = stackCount + 1
					stack[v52] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child1
				} else {
					v52 = stackCount
					stackCount = stackCount + 1
					stack[v52] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child1
					v52 = stackCount
					stackCount = stackCount + 1
					stack[v52] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child2
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
	var abs_v, b10, b5, b9, c, c1, c11, c2, c21, c3, extension, h, p1, r, radius, t, v2, v14, v16, v17, v21, v28, v3, v30, v31, v32, v34, v35, v36, v39, v42, v45, v46, v48, v57, v60, v61 Vec2
	var i, nodeId, stackCount, v99 int32
	var maxFraction, term1, term2, value, v10, v11, v13, v18, v19, v20, v22, v23, v24, v25, v27, v4, v44, v49, v5, v50, v52, v53, v54, v56, v59, v6, v8, v9 float32
	var node, nodes uintptr
	var originAABB, totalAABB, v38, v41 AABB
	var stack [1024]int32
	var stats TreeStats
	var v103 uint8
	var v105 bool
	var _ /* subInput at bp+0 */ ShapeCastInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = abs_v, b10, b5, b9, c, c1, c11, c2, c21, c3, extension, h, i, maxFraction, node, nodeId, nodes, originAABB, p1, r, radius, stack, stackCount, stats, t, term1, term2, totalAABB, v2, value, v10, v103, v105, v11, v13, v14, v16, v17, v18, v19, v21, v20, v22, v23, v24, v25, v27, v28, v3, v30, v31, v32, v34, v35, v36, v38, v39, v4, v41, v42, v44, v45, v46, v48, v49, v5, v50, v52, v53, v54, v56, v57, v59, v6, v60, v61, v8, v9, v99
	stats = TreeStats{}
	if (*DynamicTree)(unsafe.Pointer(tree)).NodeCount == 0 || (*ShapeCastInput)(unsafe.Pointer(input)).Proxy.Count == 0 {
		return stats
	}
	originAABB = AABB{
		LowerBound: *(*Vec2)(unsafe.Add(nil, input)),
		UpperBound: *(*Vec2)(unsafe.Add(nil, input)),
	}
	i = int32(1)
	for {
		if !(i < (*ShapeCastInput)(unsafe.Pointer(input)).Proxy.Count) {
			break
		}
		v21 = originAABB.LowerBound
		v3 = *(*Vec2)(unsafe.Add(nil, input+uintptr(i)*8))
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
		v21 = originAABB.UpperBound
		v3 = *(*Vec2)(unsafe.Add(nil, input+uintptr(i)*8))
		v4 = v21.X
		v5 = v3.X
		if v4 > v5 {
			v8 = v4
		} else {
			v8 = v5
		}
		v6 = v8
		goto _21
	_21:
		c1.X = v6
		v9 = v21.Y
		v10 = v3.Y
		if v9 > v10 {
			v13 = v9
		} else {
			v13 = v10
		}
		v11 = v13
		goto _26
	_26:
		c1.Y = v11
		v14 = c1
		goto _29
	_29:
		originAABB.UpperBound = v14
		goto _1
	_1:
		;
		i = i + 1
	}
	radius = Vec2{
		X: (*ShapeCastInput)(unsafe.Pointer(input)).Proxy.Radius,
		Y: (*ShapeCastInput)(unsafe.Pointer(input)).Proxy.Radius,
	}
	v21 = originAABB.LowerBound
	v3 = radius
	v14 = Vec2{
		X: v21.X - v3.X,
		Y: v21.Y - v3.Y,
	}
	goto _33
_33:
	originAABB.LowerBound = v14
	v21 = originAABB.UpperBound
	v3 = radius
	v14 = Vec2{
		X: v21.X + v3.X,
		Y: v21.Y + v3.Y,
	}
	goto _37
_37:
	originAABB.UpperBound = v14
	v38 = originAABB
	b9 = Vec2{
		X: float32(float32FromFloat32(0.5) * (v38.LowerBound.X + v38.UpperBound.X)),
		Y: float32(float32FromFloat32(0.5) * (v38.LowerBound.Y + v38.UpperBound.Y)),
	}
	v21 = b9
	goto _40
_40:
	p1 = v21
	v41 = originAABB
	b10 = Vec2{
		X: float32(float32FromFloat32(0.5) * (v41.UpperBound.X - v41.LowerBound.X)),
		Y: float32(float32FromFloat32(0.5) * (v41.UpperBound.Y - v41.LowerBound.Y)),
	}
	v3 = b10
	goto _43
_43:
	extension = v3
	// v is perpendicular to the segment.
	r = (*ShapeCastInput)(unsafe.Pointer(input)).Translation
	v4 = float32FromFloat32(1)
	v14 = r
	v16 = Vec2{
		X: float32(-v4 * v14.Y),
		Y: float32(v4 * v14.X),
	}
	goto _47
_47:
	v2 = v16
	v17 = v2
	v5 = v17.X
	if v5 < float32FromInt32(0) {
		v8 = -v5
	} else {
		v8 = v5
	}
	v6 = v8
	goto _51
_51:
	b5.X = v6
	v9 = v17.Y
	if v9 < float32FromInt32(0) {
		v11 = -v9
	} else {
		v11 = v9
	}
	v10 = v11
	goto _55
_55:
	b5.Y = v10
	v28 = b5
	goto _58
_58:
	abs_v = v28
	// Separating axis for segment (Gino, p80).
	// |dot(v, p1 - c)| > dot(|v|, h)
	maxFraction = (*ShapeCastInput)(unsafe.Pointer(input)).MaxFraction
	v13 = maxFraction
	v30 = (*ShapeCastInput)(unsafe.Pointer(input)).Translation
	v31 = Vec2{
		X: float32(v13 * v30.X),
		Y: float32(v13 * v30.Y),
	}
	goto _62
_62:
	// Build total box for the shape cast
	t = v31
	v32 = originAABB.LowerBound
	v34 = t
	v35 = Vec2{
		X: v32.X + v34.X,
		Y: v32.Y + v34.Y,
	}
	goto _66
_66:
	v36 = originAABB.LowerBound
	v39 = v35
	v18 = v36.X
	v19 = v39.X
	if v18 < v19 {
		v22 = v18
	} else {
		v22 = v19
	}
	v20 = v22
	goto _72
_72:
	c.X = v20
	v23 = v36.Y
	v24 = v39.Y
	if v23 < v24 {
		v27 = v23
	} else {
		v27 = v24
	}
	v25 = v27
	goto _77
_77:
	c.Y = v25
	v42 = c
	goto _80
_80:
	v45 = originAABB.UpperBound
	v46 = t
	v48 = Vec2{
		X: v45.X + v46.X,
		Y: v45.Y + v46.Y,
	}
	goto _84
_84:
	v57 = originAABB.UpperBound
	v60 = v48
	v44 = v57.X
	v49 = v60.X
	if v44 > v49 {
		v52 = v44
	} else {
		v52 = v49
	}
	v50 = v52
	goto _90
_90:
	c1.X = v50
	v53 = v57.Y
	v54 = v60.Y
	if v53 > v54 {
		v59 = v53
	} else {
		v59 = v54
	}
	v56 = v59
	goto _95
_95:
	c1.Y = v56
	v61 = c1
	goto _98
_98:
	totalAABB = AABB{
		LowerBound: v42,
		UpperBound: v61,
	}
	*(*ShapeCastInput)(unsafe.Add(nil, bp)) = *(*ShapeCastInput)(unsafe.Add(nil, input))
	nodes = (*DynamicTree)(unsafe.Pointer(tree)).Nodes
	stackCount = 0
	v99 = stackCount
	stackCount = stackCount + 1
	stack[v99] = (*DynamicTree)(unsafe.Pointer(tree)).Root
	for stackCount > 0 {
		stackCount = stackCount - 1
		v99 = stackCount
		nodeId = stack[v99]
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
			v38 = (*b2TreeNode)(unsafe.Pointer(node)).Aabb
			v41 = totalAABB
			v103 = boolUint8(!(v41.LowerBound.X > v38.UpperBound.X || v41.LowerBound.Y > v38.UpperBound.Y || v38.LowerBound.X > v41.UpperBound.X || v38.LowerBound.Y > v41.UpperBound.Y))
			goto _104
		_104:
		}
		if v105 || int32FromUint8(v103) == false1 {
			continue
		}
		v38 = (*b2TreeNode)(unsafe.Pointer(node)).Aabb
		b9 = Vec2{
			X: float32(float32FromFloat32(0.5) * (v38.LowerBound.X + v38.UpperBound.X)),
			Y: float32(float32FromFloat32(0.5) * (v38.LowerBound.Y + v38.UpperBound.Y)),
		}
		v21 = b9
		goto _108
	_108:
		// Separating axis for segment (Gino, p80).
		// |dot(v, p1 - c)| > dot(|v|, h)
		// radius extension is added to the node in this case
		c3 = v21
		v41 = (*b2TreeNode)(unsafe.Pointer(node)).Aabb
		b10 = Vec2{
			X: float32(float32FromFloat32(0.5) * (v41.UpperBound.X - v41.LowerBound.X)),
			Y: float32(float32FromFloat32(0.5) * (v41.UpperBound.Y - v41.LowerBound.Y)),
		}
		v3 = b10
		goto _111
	_111:
		v14 = v3
		v16 = extension
		v17 = Vec2{
			X: v14.X + v16.X,
			Y: v14.Y + v16.Y,
		}
		goto _115
	_115:
		h = v17
		v28 = p1
		v30 = c3
		v31 = Vec2{
			X: v28.X - v30.X,
			Y: v28.Y - v30.Y,
		}
		goto _119
	_119:
		v32 = v2
		v34 = v31
		v4 = float32(v32.X*v34.X) + float32(v32.Y*v34.Y)
		goto _123
	_123:
		v5 = v4
		if v5 < float32FromInt32(0) {
			v8 = -v5
		} else {
			v8 = v5
		}
		v6 = v8
		goto _126
	_126:
		term1 = v6
		v35 = abs_v
		v36 = h
		v9 = float32(v35.X*v36.X) + float32(v35.Y*v36.Y)
		goto _131
	_131:
		term2 = v9
		if term2 < term1 {
			continue
		}
		if b2IsLeaf(tls, node) != 0 {
			(*(*ShapeCastInput)(unsafe.Add(nil, bp))).MaxFraction = maxFraction
			value = (*(*func(*_Stack, uintptr, int32, uint64, uintptr) float32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_callback})))(tls, bp, nodeId, *(*uint64_t)(unsafe.Add(unsafe.Pointer(node), 24)), context)
			stats.LeafVisits += int32(1)
			if value == float32FromFloat32(0) {
				// The client has terminated the ray cast.
				return stats
			}
			if float32FromFloat32(0) < value && value < maxFraction {
				// Update segment bounding box.
				maxFraction = value
				v4 = maxFraction
				v21 = (*ShapeCastInput)(unsafe.Pointer(input)).Translation
				v3 = Vec2{
					X: float32(v4 * v21.X),
					Y: float32(v4 * v21.Y),
				}
				goto _135
			_135:
				t = v3
				v21 = originAABB.LowerBound
				v3 = t
				v14 = Vec2{
					X: v21.X + v3.X,
					Y: v21.Y + v3.Y,
				}
				goto _139
			_139:
				v16 = originAABB.LowerBound
				v17 = v14
				v4 = v16.X
				v5 = v17.X
				if v4 < v5 {
					v8 = v4
				} else {
					v8 = v5
				}
				v6 = v8
				goto _145
			_145:
				c.X = v6
				v9 = v16.Y
				v10 = v17.Y
				if v9 < v10 {
					v13 = v9
				} else {
					v13 = v10
				}
				v11 = v13
				goto _150
			_150:
				c.Y = v11
				v28 = c
				goto _153
			_153:
				totalAABB.LowerBound = v28
				v21 = originAABB.UpperBound
				v3 = t
				v14 = Vec2{
					X: v21.X + v3.X,
					Y: v21.Y + v3.Y,
				}
				goto _157
			_157:
				v16 = originAABB.UpperBound
				v17 = v14
				v4 = v16.X
				v5 = v17.X
				if v4 > v5 {
					v8 = v4
				} else {
					v8 = v5
				}
				v6 = v8
				goto _163
			_163:
				c1.X = v6
				v9 = v16.Y
				v10 = v17.Y
				if v9 > v10 {
					v13 = v9
				} else {
					v13 = v10
				}
				v11 = v13
				goto _168
			_168:
				c1.Y = v11
				v28 = c1
				goto _171
			_171:
				totalAABB.UpperBound = v28
			}
		} else {
			if stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)-int32FromInt32(1) {
				v38 = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr((*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child1)*40))).Aabb
				b9 = Vec2{
					X: float32(float32FromFloat32(0.5) * (v38.LowerBound.X + v38.UpperBound.X)),
					Y: float32(float32FromFloat32(0.5) * (v38.LowerBound.Y + v38.UpperBound.Y)),
				}
				v21 = b9
				goto _174
			_174:
				c11 = v21
				v41 = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr((*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child2)*40))).Aabb
				b9 = Vec2{
					X: float32(float32FromFloat32(0.5) * (v41.LowerBound.X + v41.UpperBound.X)),
					Y: float32(float32FromFloat32(0.5) * (v41.LowerBound.Y + v41.UpperBound.Y)),
				}
				v3 = b9
				goto _177
			_177:
				c21 = v3
				v14 = c11
				v16 = p1
				c2 = Vec2{
					X: v16.X - v14.X,
					Y: v16.Y - v14.Y,
				}
				v4 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
				goto _181
			_181:
				v17 = c21
				v28 = p1
				c2 = Vec2{
					X: v28.X - v17.X,
					Y: v28.Y - v17.Y,
				}
				v5 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
				goto _185
			_185:
				if v4 < v5 {
					v99 = stackCount
					stackCount = stackCount + 1
					stack[v99] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child2
					v99 = stackCount
					stackCount = stackCount + 1
					stack[v99] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child1
				} else {
					v99 = stackCount
					stackCount = stackCount + 1
					stack[v99] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child1
					v99 = stackCount
					stackCount = stackCount + 1
					stack[v99] = (*b2TreeNode)(unsafe.Pointer(node)).ｆ__ccgo2_24.Children.Child2
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
	var doomedNodeIndex, leafCount, newCapacity, nodeIndex, proxyCount, stackCount, v4 int32
	var leafCenters, leafIndices, node, nodes uintptr
	var stack [1024]int32
	var v1 AABB
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b, doomedNodeIndex, leafCenters, leafCount, leafIndices, newCapacity, node, nodeIndex, nodes, proxyCount, stack, stackCount, v1, v2, v4
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
			*(*int32)(unsafe.Add(nil, leafIndices+uintptr(leafCount)*4)) = nodeIndex
			v1 = (*b2TreeNode)(unsafe.Pointer(node)).Aabb
			b = Vec2{
				X: float32(float32FromFloat32(0.5) * (v1.LowerBound.X + v1.UpperBound.X)),
				Y: float32(float32FromFloat32(0.5) * (v1.LowerBound.Y + v1.UpperBound.Y)),
			}
			v2 = b
			goto _3
		_3:
			*(*Vec2)(unsafe.Add(nil, leafCenters+uintptr(leafCount)*8)) = v2
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
		v4 = stackCount
		nodeIndex = stack[v4]
		node = nodes + uintptr(nodeIndex)*40
	}
	if !(leafCount <= proxyCount) && b2InternalAssertFcn(tls, __ccgo_ts+6481, __ccgo_ts+4746, int32FromInt32(1982)) != 0 {
		__builtin_trap(tls)
	}
	(*DynamicTree)(unsafe.Pointer(tree)).Root = b2BuildTree(tls, tree, leafCount)
	b2DynamicTree_Validate(tls, tree)
	return leafCount
}
