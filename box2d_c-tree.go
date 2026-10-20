package b2

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b2CreateDynamicTreeInternal(tls *_Stack, tree uintptr, proxyCapacity int32) {
	var capacity, i, v1, v2, v3, v5 int32
	var v6 b2TreeNode
	_, _, _, _, _, _, _ = capacity, i, v1, v2, v3, v5, v6
	v1 = proxyCapacity
	v2 = int32(16)
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	capacity = v3
	// memset needed for deterministic serialization.
	memset(tls, tree, 0, uint64(104))
	// A tree of n proxies has 2n - 1 nodes plus the empty node beside the root
	(*b2DynamicTree)(unsafe.Pointer(tree)).NodeCapacity = int32(2) * capacity
	(*b2DynamicTree)(unsafe.Pointer(tree)).Nodes = b2AllocZero(tls, uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).NodeCapacity)*uint64(32))
	(*b2DynamicTree)(unsafe.Pointer(tree)).Parents = b2AllocZero(tls, uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).NodeCapacity)*uint64(4))
	(*b2DynamicTree)(unsafe.Pointer(tree)).PairFreeList = -int32FromInt32(1)
	// The root and the empty node always exist
	v6 = b2TreeNode{
		Aabb: AABB{
			LowerBound: Vec2{
				X: __builtin_inff(tls),
				Y: __builtin_inff(tls),
			},
			UpperBound: Vec2{
				X: -__builtin_inff(tls),
				Y: -__builtin_inff(tls),
			},
		},
		FlagIndex: uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30)|uint32FromUint32(1)<<int32FromInt32(31)) | uint32FromUint32(1)<<int32FromInt32(31),
	}
	goto _7
_7:
	*(*b2TreeNode)(unsafe.Add(nil, (*b2DynamicTree)(unsafe.Pointer(tree)).Nodes)) = v6
	v6 = b2TreeNode{
		Aabb: AABB{
			LowerBound: Vec2{
				X: __builtin_inff(tls),
				Y: __builtin_inff(tls),
			},
			UpperBound: Vec2{
				X: -__builtin_inff(tls),
				Y: -__builtin_inff(tls),
			},
		},
		FlagIndex: uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30)|uint32FromUint32(1)<<int32FromInt32(31)) | uint32FromUint32(1)<<int32FromInt32(31),
	}
	goto _9
_9:
	*(*b2TreeNode)(unsafe.Add(nil, (*b2DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr(int32FromInt32(_B2_ROOT_NODE)+int32FromInt32(1))*32)) = v6
	*(*int32_t)(unsafe.Add(nil, (*b2DynamicTree)(unsafe.Pointer(tree)).Parents)) = -int32FromInt32(1)
	*(*int32_t)(unsafe.Add(nil, (*b2DynamicTree)(unsafe.Pointer(tree)).Parents+uintptr(int32FromInt32(_B2_ROOT_NODE)+int32FromInt32(1))*4)) = -int32FromInt32(1)
	// This is the bump index.
	(*b2DynamicTree)(unsafe.Pointer(tree)).NodeEnd = int32(2)
	(*b2DynamicTree)(unsafe.Pointer(tree)).DfsOrdered = boolUint8(true1 != 0)
	(*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity = capacity
	(*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCount = 0
	(*b2DynamicTree)(unsafe.Pointer(tree)).Proxies = b2AllocZero(tls, uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity)*uint64(24))
	// Build a linked list for the free list.
	i = 0
	for {
		if !(i < (*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity-int32(1)) {
			break
		}
		(*(*b2TreeProxy)(unsafe.Add(nil, (*b2DynamicTree)(unsafe.Pointer(tree)).Proxies+uintptr(i)*24))).Node = -int32FromInt32(1)
		(*(*b2TreeProxy)(unsafe.Add(nil, (*b2DynamicTree)(unsafe.Pointer(tree)).Proxies+uintptr(i)*24))).Next = i + int32(1)
		goto _10
	_10:
		;
		i = i + 1
	}
	(*(*b2TreeProxy)(unsafe.Add(nil, (*b2DynamicTree)(unsafe.Pointer(tree)).Proxies+uintptr((*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity-int32(1))*24))).Node = -int32FromInt32(1)
	(*(*b2TreeProxy)(unsafe.Add(nil, (*b2DynamicTree)(unsafe.Pointer(tree)).Proxies+uintptr((*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity-int32(1))*24))).Next = -int32FromInt32(1)
	(*b2DynamicTree)(unsafe.Pointer(tree)).ProxyFreeList = 0
	(*b2DynamicTree)(unsafe.Pointer(tree)).LeafIndices = uintptrFromInt32(0)
	(*b2DynamicTree)(unsafe.Pointer(tree)).LeafNodes = uintptrFromInt32(0)
	(*b2DynamicTree)(unsafe.Pointer(tree)).LeafBoxes = uintptrFromInt32(0)
	(*b2DynamicTree)(unsafe.Pointer(tree)).LeafCenters = uintptrFromInt32(0)
	(*b2DynamicTree)(unsafe.Pointer(tree)).BinIndices = uintptrFromInt32(0)
	(*b2DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity = 0
}

func b2DestroyDynamicTreeInternal(tls *_Stack, tree uintptr) {
	b2Free(tls, (*b2DynamicTree)(unsafe.Pointer(tree)).Nodes, uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).NodeCapacity)*uint64(32))
	b2Free(tls, (*b2DynamicTree)(unsafe.Pointer(tree)).Parents, uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).NodeCapacity)*uint64(4))
	b2Free(tls, (*b2DynamicTree)(unsafe.Pointer(tree)).Proxies, uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity)*uint64(24))
	b2Free(tls, (*b2DynamicTree)(unsafe.Pointer(tree)).SwapNodes, uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).NodeCapacity)*uint64(32))
	b2Free(tls, (*b2DynamicTree)(unsafe.Pointer(tree)).LeafIndices, uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(4))
	b2Free(tls, (*b2DynamicTree)(unsafe.Pointer(tree)).LeafNodes, uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(32))
	b2Free(tls, (*b2DynamicTree)(unsafe.Pointer(tree)).LeafBoxes, uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(16))
	b2Free(tls, (*b2DynamicTree)(unsafe.Pointer(tree)).LeafCenters, uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(8))
	b2Free(tls, (*b2DynamicTree)(unsafe.Pointer(tree)).BinIndices, uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(4))
	memset(tls, tree, 0, uint64(104))
}

func b2CreateDynamicTree(tls *_Stack, proxyCapacity int32) (r uintptr) {
	var tree uintptr
	_ = tree
	tree = b2Alloc(tls, uint64(104))
	b2CreateDynamicTreeInternal(tls, tree, proxyCapacity)
	return tree
}

func b2DestroyDynamicTree(tls *_Stack, tree uintptr) {
	if tree == uintptrFromInt32(0) {
		return
	}
	b2DestroyDynamicTreeInternal(tls, tree)
	b2Free(tls, tree, uint64(104))
}

func b2DynamicTree_GetProxyCount(tls *_Stack, tree uintptr) (r int32) {
	return (*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCount
}

func b2DynamicTree_MoveProxyInternal(tls *_Stack, tree uintptr, proxyId int32, aabb AABB, markMoved uint8) {
	var shouldRotate uint8
	var v1 bool
	_, _ = shouldRotate, v1
	if v1 = !!(0 <= proxyId && proxyId < (*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity); !v1 {
		b2InternalAssert(tls, __ccgo_ts+9087, __ccgo_ts+9017, int32FromInt32(620))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	b2RemoveLeaf(tls, tree, proxyId)
	shouldRotate = boolUint8(false1 != 0)
	b2InsertLeaf(tls, tree, aabb, proxyId, markMoved, shouldRotate)
}

func b2DynamicTree_MoveProxy(tls *_Stack, tree uintptr, proxyId int32, aabb AABB) {
	b2DynamicTree_MoveProxyInternal(tls, tree, proxyId, aabb, boolUint8(false1 != 0))
}

func b2DynamicTree_EnlargeProxy(tls *_Stack, tree uintptr, proxyId int32, aabb AABB) {
	var changed, changed1, v2 uint8
	var index int32
	var node1, nodes, parents, v5 uintptr
	var v1 bool
	var v6 AABB
	_, _, _, _, _, _, _, _, _, _ = changed, changed1, index, node1, nodes, parents, v1, v2, v5, v6
	if v1 = !!(0 <= proxyId && proxyId < (*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity); !v1 {
		b2InternalAssert(tls, __ccgo_ts+9087, __ccgo_ts+9017, int32FromInt32(638))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	nodes = (*b2DynamicTree)(unsafe.Pointer(tree)).Nodes
	parents = (*b2DynamicTree)(unsafe.Pointer(tree)).Parents
	index = (*(*b2TreeProxy)(unsafe.Add(nil, (*b2DynamicTree)(unsafe.Pointer(tree)).Proxies+uintptr(proxyId)*24))).Node
	node1 = nodes + uintptr(index)*32
	v2 = boolUint8((*b2TreeNode)(unsafe.Pointer(node1)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
	goto _3
_3:
	;
	if v1 = !!(v2 != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+9321, __ccgo_ts+9017, int32FromInt32(645))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// Caller must ensure this
	(*b2TreeNode)(unsafe.Pointer(node1)).Aabb = aabb
	*(*uint32_t)(unsafe.Add(nil, node1+24)) |= uint32FromUint32(1) << int32FromInt32(30)
	index = *(*int32_t)(unsafe.Add(nil, parents+uintptr(index)*4))
	for index != -int32(1) {
		node1 = nodes + uintptr(index)*32
		v5 = node1
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
		v2 = changed
		goto _8
	_8:
		changed1 = v2
		// This is marked to ensure the root is marked in this loop or the one below.
		*(*uint32_t)(unsafe.Add(nil, node1+24)) |= uint32FromUint32(1) << int32FromInt32(30)
		index = *(*int32_t)(unsafe.Add(nil, parents+uintptr(index)*4))
		if int32FromUint8(changed1) == false1 {
			break
		}
	}
	// Mark all the way up to the root.
	for index != -int32(1) {
		node1 = nodes + uintptr(index)*32
		if (*b2TreeNode)(unsafe.Pointer(node1)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) != 0 {
			// Early out because this ancestor was previously ascended and marked as moved.
			break
		}
		*(*uint32_t)(unsafe.Add(nil, node1+24)) |= uint32FromUint32(1) << int32FromInt32(30)
		index = *(*int32_t)(unsafe.Add(nil, parents+uintptr(index)*4))
	}
}

func b2DynamicTree_SetCategoryBits(tls *_Stack, tree uintptr, proxyId int32, categoryBits uint64) {
	var v1 bool
	_ = v1
	if v1 = !!(0 <= proxyId && proxyId < (*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity); !v1 {
		b2InternalAssert(tls, __ccgo_ts+9087, __ccgo_ts+9017, int32FromInt32(687))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	(*(*b2TreeProxy)(unsafe.Add(nil, (*b2DynamicTree)(unsafe.Pointer(tree)).Proxies+uintptr(proxyId)*24))).CategoryBits = categoryBits
}

func b2DynamicTree_GetCategoryBits(tls *_Stack, tree uintptr, proxyId int32) (r uint64) {
	var v1 bool
	_ = v1
	if v1 = !!(0 <= proxyId && proxyId < (*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity); !v1 {
		b2InternalAssert(tls, __ccgo_ts+9087, __ccgo_ts+9017, int32FromInt32(693))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return (*(*b2TreeProxy)(unsafe.Add(nil, (*b2DynamicTree)(unsafe.Pointer(tree)).Proxies+uintptr(proxyId)*24))).CategoryBits
}

func b2DynamicTree_GetHeight(tls *_Stack, tree uintptr) (r int32) {
	return b2GetNodeHeight(tls, (*b2DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr(_B2_ROOT_NODE)*32)
}

// C documentation
//
//	// The area ratio is the thing that SAH seeks to minimize. SAH
//	// cannot do anything about leaf boxes or the root box. It seeks
//	// to minimize the area of all non-root internal nodes. Divide this
//	// by the root area to make the metric non-dimensional.
//	// So this becomes a meaningful measure of tree quality.
func b2DynamicTree_GetAreaRatio(tls *_Stack, tree uintptr) (r float32) {
	var i, nodeEnd int32
	var internalArea, rootArea, wx, wy, v2 float32
	var nodes uintptr
	var v1 AABB
	var v5 uint8
	_, _, _, _, _, _, _, _, _, _ = i, internalArea, nodeEnd, nodes, rootArea, wx, wy, v1, v2, v5
	if (*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCount == 0 {
		return float32FromFloat32(0)
	}
	nodes = (*b2DynamicTree)(unsafe.Pointer(tree)).Nodes
	v1 = (*(*b2TreeNode)(unsafe.Add(nil, nodes))).Aabb
	wx = v1.UpperBound.X - v1.LowerBound.X
	wy = v1.UpperBound.Y - v1.LowerBound.Y
	v2 = float32(float32FromFloat32(2) * (wx + wy))
	goto _3
_3:
	rootArea = v2
	if rootArea <= float32FromFloat32(0) {
		return float32FromFloat32(0)
	}
	// Free nodes and the empty node are leaf tagged.
	internalArea = float32FromFloat32(0)
	nodeEnd = (*b2DynamicTree)(unsafe.Pointer(tree)).NodeEnd
	i = int32(2)
	for {
		if !(i < nodeEnd) {
			break
		}
		v5 = boolUint8((*b2TreeNode)(unsafe.Pointer(nodes+uintptr(i)*32)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
		goto _6
	_6:
		if int32FromUint8(v5) == false1 {
			v1 = (*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(i)*32))).Aabb
			wx = v1.UpperBound.X - v1.LowerBound.X
			wy = v1.UpperBound.Y - v1.LowerBound.Y
			v2 = float32(float32FromFloat32(2) * (wx + wy))
			goto _9
		_9:
			internalArea = internalArea + v2
		}
		goto _4
	_4:
		;
		i = i + 1
	}
	return internalArea / rootArea
}

func b2DynamicTree_GetRootBounds(tls *_Stack, tree uintptr) (r AABB) {
	if (*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCount == 0 {
		return AABB{
			LowerBound: b2Vec2_zero,
			UpperBound: b2Vec2_zero,
		}
	}
	return (*(*b2TreeNode)(unsafe.Add(nil, (*b2DynamicTree)(unsafe.Pointer(tree)).Nodes))).Aabb
}

func b2DynamicTree_Validate(tls *_Stack, tree uintptr) {
	_ = uint64FromInt64(4)
}

func b2DynamicTree_ValidateNoMoved(tls *_Stack, tree uintptr) {
	_ = uint64FromInt64(4)
}

func b2DynamicTree_GetByteCount(tls *_Stack, tree uintptr) (r size_t) {
	var size size_t
	var v1 uint64
	_, _ = size, v1
	size = uint64(104)
	size = size + uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).NodeCapacity)*uint64(32)
	size = size + uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).NodeCapacity)*uint64(4)
	size = size + uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity)*uint64(24)
	if (*b2DynamicTree)(unsafe.Pointer(tree)).SwapNodes == uintptrFromInt32(0) {
		v1 = uint64(0)
	} else {
		v1 = uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).NodeCapacity) * uint64(32)
	}
	size = size + v1
	// leafIndices
	size = size + uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(4)
	// leafNodes
	size = size + uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(32)
	// leafBoxes
	size = size + uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(16)
	// leafCenters
	size = size + uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(8)
	// binIndices
	size = size + uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity)*uint64(4)
	return size
}

func b2DynamicTree_GetUserData(tls *_Stack, tree uintptr, proxyId int32) (r uint64) {
	var v1 bool
	_ = v1
	if v1 = !!(0 <= proxyId && proxyId < (*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity); !v1 {
		b2InternalAssert(tls, __ccgo_ts+9087, __ccgo_ts+9017, int32FromInt32(872))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return (*(*b2TreeProxy)(unsafe.Add(nil, (*b2DynamicTree)(unsafe.Pointer(tree)).Proxies+uintptr(proxyId)*24))).UserData
}

func b2DynamicTree_GetProxyBounds(tls *_Stack, tree uintptr, proxyId int32) (r AABB) {
	var nodeIndex int32
	var v1 bool
	_, _ = nodeIndex, v1
	if v1 = !!(0 <= proxyId && proxyId < (*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCapacity); !v1 {
		b2InternalAssert(tls, __ccgo_ts+9087, __ccgo_ts+9017, int32FromInt32(878))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	nodeIndex = (*(*b2TreeProxy)(unsafe.Add(nil, (*b2DynamicTree)(unsafe.Pointer(tree)).Proxies+uintptr(proxyId)*24))).Node
	if v1 = !!(0 <= nodeIndex && nodeIndex < (*b2DynamicTree)(unsafe.Pointer(tree)).NodeEnd); !v1 {
		b2InternalAssert(tls, __ccgo_ts+9338, __ccgo_ts+9017, int32FromInt32(880))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return (*(*b2TreeNode)(unsafe.Add(nil, (*b2DynamicTree)(unsafe.Pointer(tree)).Nodes+uintptr(nodeIndex)*32))).Aabb
}

func b2DynamicTree_OverlapAABB(tls *_Stack, tree uintptr, _aabb AABB, maskBits uint64, __ccgo_fp_callback uintptr, context uintptr) (r TreeStats) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*AABB)(unsafe.Pointer(bp)) = _aabb
	var boxv, v9 b2AABBV
	var bv, node4, nodes1, proxy, root uintptr
	var i, pair, proxyId, stackCount, v1, v2, v4, v7 int32
	var proceed, v14, v5 uint8
	var result TreeStats
	var stack [512]int32
	var v23 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = boxv, bv, i, node4, nodes1, pair, proceed, proxy, proxyId, result, root, stack, stackCount, v1, v14, v2, v23, v4, v5, v7, v9
	result = TreeStats{}
	if (*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCount == 0 {
		return result
	}
	nodes1 = (*b2DynamicTree)(unsafe.Pointer(tree)).Nodes
	stackCount = 0
	v1 = stackCount
	stackCount = stackCount + 1
	root = nodes1 + uintptr(_B2_ROOT_NODE)*32
	v5 = boolUint8((*b2TreeNode)(unsafe.Pointer(root)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
	goto _6
_6:
	if v5 != 0 {
		v4 = _B2_ROOT_NODE
	} else {
		v7 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(root)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
		goto _8
	_8:
		v4 = v7
	}
	v2 = v4
	goto _3
_3:
	stack[v1] = v2
	v9 = *(*AABB)(unsafe.Add(nil, bp))
	goto _10
_10:
	boxv = v9
	for stackCount > 0 {
		stackCount = stackCount - 1
		v1 = stackCount
		pair = stack[v1]
		result.NodeVisits += int32(1)
		i = 0
		for {
			if !(i < int32(2)) {
				break
			}
			node4 = nodes1 + uintptr(pair)*32 + uintptr(i)*32
			v9 = boxv
			bv = node4
			v5 = boolUint8(v9.LowerBound.X <= (*AABB)(unsafe.Pointer(bv)).UpperBound.X && v9.LowerBound.Y <= (*AABB)(unsafe.Pointer(bv)).UpperBound.Y && (*AABB)(unsafe.Pointer(bv)).LowerBound.X <= v9.UpperBound.X && (*AABB)(unsafe.Pointer(bv)).LowerBound.Y <= v9.UpperBound.Y)
			goto _15
		_15:
			if v5 != 0 {
				v14 = boolUint8((*b2TreeNode)(unsafe.Pointer(node4)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
				goto _17
			_17:
				if v14 != 0 {
					v1 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(node4)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
					goto _19
				_19:
					// callback to user code with proxy id
					proxyId = v1
					proxy = (*b2DynamicTree)(unsafe.Pointer(tree)).Proxies + uintptr(proxyId)*24
					if (*b2TreeProxy)(unsafe.Pointer(proxy)).CategoryBits&maskBits != 0 {
						proceed = (*(*func(*_Stack, int32, uint64, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_callback})))(tls, proxyId, (*b2TreeProxy)(unsafe.Pointer(proxy)).UserData, context)
						result.LeafVisits += int32(1)
						if int32FromUint8(proceed) == false1 {
							return result
						}
					}
				} else {
					if stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)-int32FromInt32(1) {
						v1 = stackCount
						stackCount = stackCount + 1
						v2 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(node4)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
						goto _22
					_22:
						stack[v1] = v2
					} else {
						if v23 = !!(stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)-int32FromInt32(1)); !v23 {
							b2InternalAssert(tls, __ccgo_ts+9382, __ccgo_ts+9017, int32FromInt32(936))
						}
						_ = v23 || bool(int32FromInt32(0) != 0)
					}
				}
			}
			goto _12
		_12:
			;
			i = i + 1
		}
	}
	return result
}

// C documentation
//
//	// This ignores the category bits. Useful for debug draw.
func b2DynamicTree_OverlapAABBAll(tls *_Stack, tree uintptr, _aabb AABB, __ccgo_fp_callback uintptr, context uintptr) (r TreeStats) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*AABB)(unsafe.Pointer(bp)) = _aabb
	var boxv, v9 b2AABBV
	var bv, node4, nodes1, proxy, root uintptr
	var i, pair, proxyId, stackCount, v1, v2, v4, v7 int32
	var proceed, v14, v5 uint8
	var result TreeStats
	var stack [512]int32
	var v23 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = boxv, bv, i, node4, nodes1, pair, proceed, proxy, proxyId, result, root, stack, stackCount, v1, v14, v2, v23, v4, v5, v7, v9
	result = TreeStats{}
	if (*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCount == 0 {
		return result
	}
	nodes1 = (*b2DynamicTree)(unsafe.Pointer(tree)).Nodes
	stackCount = 0
	v1 = stackCount
	stackCount = stackCount + 1
	root = nodes1 + uintptr(_B2_ROOT_NODE)*32
	v5 = boolUint8((*b2TreeNode)(unsafe.Pointer(root)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
	goto _6
_6:
	if v5 != 0 {
		v4 = _B2_ROOT_NODE
	} else {
		v7 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(root)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
		goto _8
	_8:
		v4 = v7
	}
	v2 = v4
	goto _3
_3:
	stack[v1] = v2
	v9 = *(*AABB)(unsafe.Add(nil, bp))
	goto _10
_10:
	boxv = v9
	for stackCount > 0 {
		stackCount = stackCount - 1
		v1 = stackCount
		pair = stack[v1]
		result.NodeVisits += int32(1)
		i = 0
		for {
			if !(i < int32(2)) {
				break
			}
			node4 = nodes1 + uintptr(pair)*32 + uintptr(i)*32
			v9 = boxv
			bv = node4
			v5 = boolUint8(v9.LowerBound.X <= (*AABB)(unsafe.Pointer(bv)).UpperBound.X && v9.LowerBound.Y <= (*AABB)(unsafe.Pointer(bv)).UpperBound.Y && (*AABB)(unsafe.Pointer(bv)).LowerBound.X <= v9.UpperBound.X && (*AABB)(unsafe.Pointer(bv)).LowerBound.Y <= v9.UpperBound.Y)
			goto _15
		_15:
			if v5 != 0 {
				v14 = boolUint8((*b2TreeNode)(unsafe.Pointer(node4)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
				goto _17
			_17:
				if v14 != 0 {
					v1 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(node4)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
					goto _19
				_19:
					// callback to user code with proxy id
					proxyId = v1
					proxy = (*b2DynamicTree)(unsafe.Pointer(tree)).Proxies + uintptr(proxyId)*24
					proceed = (*(*func(*_Stack, int32, uint64, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_callback})))(tls, proxyId, (*b2TreeProxy)(unsafe.Pointer(proxy)).UserData, context)
					result.LeafVisits += int32(1)
					if int32FromUint8(proceed) == false1 {
						return result
					}
				} else {
					if stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)-int32FromInt32(1) {
						v1 = stackCount
						stackCount = stackCount + 1
						v2 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(node4)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
						goto _22
					_22:
						stack[v1] = v2
					} else {
						if v23 = !!(stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)-int32FromInt32(1)); !v23 {
							b2InternalAssert(tls, __ccgo_ts+9382, __ccgo_ts+9017, int32FromInt32(995))
						}
						_ = v23 || bool(int32FromInt32(0) != 0)
					}
				}
			}
			goto _12
		_12:
			;
			i = i + 1
		}
	}
	return result
}

// C documentation
//
//	// A lot of optimization work went into this. It beats the slab test by a significant margin.
//	// It is faster than having category bits in the nodes because of the cache line friendly node
//	// size (64 bytes).
func b2DynamicTree_CastRay(tls *_Stack, tree uintptr, input uintptr, maskBits uint64, __ccgo_fp_callback uintptr, context uintptr) (r1 TreeStats) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var B2_SWAP_TEMP, bv, node4, nodes1, proxy, root uintptr
	var abs_v, b10, b5, b9, c, c1, c2, c3, center1, center2, d, h, p1, p2, r, u, v1, v11, v17, v19, v2, v21, v22, v24, v25, v36, v38, v39, v5, v50, v6, v8 Vec2
	var boxv, v52 b2AABBV
	var d1, d2, lengthSquared, maxFraction, s2, term1, term2, value, v10, v12, v13, v14, v16, v20, v26, v27, v28, v30, v31, v32, v33, v35, v4, v40, v41, v42, v44, v45, v46, v47, v49, v9 float32
	var hit [2]uintptr
	var hitCount, i, i1, pair, proxyId, stackCount, v54, v55, v57, v60 int32
	var isLeaf [2]uint8
	var nodeAABB, v67, v70 AABB
	var result TreeStats
	var stack [512]int32
	var v146 bool
	var v58 uint8
	var _ /* segmentAABB at bp+0 */ AABB
	var _ /* subInput at bp+16 */ RayCastInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = B2_SWAP_TEMP, abs_v, b10, b5, b9, boxv, bv, c, c1, c2, c3, center1, center2, d, d1, d2, h, hit, hitCount, i, i1, isLeaf, lengthSquared, maxFraction, node4, nodeAABB, nodes1, p1, p2, pair, proxy, proxyId, r, result, root, s2, stack, stackCount, term1, term2, u, v1, value, v11, v10, v12, v13, v14, v146, v16, v17, v19, v2, v20, v21, v22, v24, v25, v26, v27, v28, v30, v31, v32, v33, v35, v36, v38, v39, v4, v40, v41, v42, v44, v45, v46, v47, v49, v5, v50, v52, v54, v55, v57, v58, v6, v60, v67, v70, v8, v9
	result = TreeStats{}
	if (*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCount == 0 {
		return result
	}
	p1 = (*RayCastInput)(unsafe.Pointer(input)).Origin
	d = (*RayCastInput)(unsafe.Pointer(input)).Translation
	v11 = d
	lengthSquared = float32(v11.X*v11.X) + float32(v11.Y*v11.Y)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s2 = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec2{
			X: float32(s2 * v11.X),
			Y: float32(s2 * v11.Y),
		}
		v2 = u
		goto _3
	}
	v2 = Vec2{}
	goto _3
_3:
	r = v2
	v4 = float32FromFloat32(1)
	v5 = r
	v6 = Vec2{
		X: float32(-v4 * v5.Y),
		Y: float32(v4 * v5.X),
	}
	goto _7
_7:
	// v is perpendicular to the segment.
	v1 = v6
	v8 = v1
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
	v21 = d
	v22 = Vec2{
		X: v19.X + float32(v20*v21.X),
		Y: v19.Y + float32(v20*v21.Y),
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
	*(*AABB)(unsafe.Add(nil, bp)) = AABB{
		LowerBound: v36,
		UpperBound: v50,
	}
	v52 = *(*AABB)(unsafe.Add(nil, bp))
	goto _53
_53:
	boxv = v52
	nodes1 = (*b2DynamicTree)(unsafe.Pointer(tree)).Nodes
	stackCount = 0
	v54 = stackCount
	stackCount = stackCount + 1
	root = nodes1 + uintptr(_B2_ROOT_NODE)*32
	v58 = boolUint8((*b2TreeNode)(unsafe.Pointer(root)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
	goto _59
_59:
	if v58 != 0 {
		v57 = _B2_ROOT_NODE
	} else {
		v60 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(root)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
		goto _61
	_61:
		v57 = v60
	}
	v55 = v57
	goto _56
_56:
	stack[v54] = v55
	*(*RayCastInput)(unsafe.Add(nil, bp+16)) = *(*RayCastInput)(unsafe.Add(nil, input))
	for stackCount > 0 {
		stackCount = stackCount - 1
		v54 = stackCount
		pair = stack[v54]
		result.NodeVisits += int32(1)
		hitCount = 0
		i = 0
		for {
			if !(i < int32(2)) {
				break
			}
			node4 = nodes1 + uintptr(pair)*32 + uintptr(i)*32
			v52 = boxv
			bv = node4
			v58 = boolUint8(v52.LowerBound.X <= (*AABB)(unsafe.Pointer(bv)).UpperBound.X && v52.LowerBound.Y <= (*AABB)(unsafe.Pointer(bv)).UpperBound.Y && (*AABB)(unsafe.Pointer(bv)).LowerBound.X <= v52.UpperBound.X && (*AABB)(unsafe.Pointer(bv)).LowerBound.Y <= v52.UpperBound.Y)
			goto _66
		_66:
			if int32FromUint8(v58) == false1 {
				goto _63
			}
			// Separating axis for segment (Gino, p80).
			// |dot(v, p1 - c)| > dot(|v|, h)
			nodeAABB = (*b2TreeNode)(unsafe.Pointer(node4)).Aabb
			v67 = nodeAABB
			b9 = Vec2{
				X: float32(float32FromFloat32(0.5) * (v67.LowerBound.X + v67.UpperBound.X)),
				Y: float32(float32FromFloat32(0.5) * (v67.LowerBound.Y + v67.UpperBound.Y)),
			}
			v11 = b9
			goto _69
		_69:
			c3 = v11
			v70 = nodeAABB
			b10 = Vec2{
				X: float32(float32FromFloat32(0.5) * (v70.UpperBound.X - v70.LowerBound.X)),
				Y: float32(float32FromFloat32(0.5) * (v70.UpperBound.Y - v70.LowerBound.Y)),
			}
			v2 = b10
			goto _72
		_72:
			h = v2
			v5 = p1
			v6 = c3
			v8 = Vec2{
				X: v5.X - v6.X,
				Y: v5.Y - v6.Y,
			}
			goto _76
		_76:
			v17 = v1
			v19 = v8
			v4 = float32(v17.X*v19.X) + float32(v17.Y*v19.Y)
			goto _80
		_80:
			v9 = v4
			if v9 < float32FromInt32(0) {
				v12 = -v9
			} else {
				v12 = v9
			}
			v10 = v12
			goto _83
		_83:
			term1 = v10
			v21 = abs_v
			v22 = h
			v13 = float32(v21.X*v22.X) + float32(v21.Y*v22.Y)
			goto _88
		_88:
			term2 = v13
			if term2 < term1 {
				goto _63
			}
			v58 = boolUint8((*b2TreeNode)(unsafe.Pointer(node4)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
			goto _90
		_90:
			isLeaf[hitCount] = v58
			hit[hitCount] = node4
			hitCount = hitCount + int32(1)
			goto _63
		_63:
			;
			i = i + 1
		}
		if hitCount == int32(2) && int32FromUint8(isLeaf[0]) == false1 && int32FromUint8(isLeaf[int32(1)]) == false1 {
			v67 = (*b2TreeNode)(unsafe.Pointer(hit[0])).Aabb
			b9 = Vec2{
				X: float32(float32FromFloat32(0.5) * (v67.LowerBound.X + v67.UpperBound.X)),
				Y: float32(float32FromFloat32(0.5) * (v67.LowerBound.Y + v67.UpperBound.Y)),
			}
			v11 = b9
			goto _93
		_93:
			center1 = v11
			v70 = (*b2TreeNode)(unsafe.Pointer(hit[int32(1)])).Aabb
			b9 = Vec2{
				X: float32(float32FromFloat32(0.5) * (v70.LowerBound.X + v70.UpperBound.X)),
				Y: float32(float32FromFloat32(0.5) * (v70.LowerBound.Y + v70.UpperBound.Y)),
			}
			v2 = b9
			goto _96
		_96:
			center2 = v2
			v5 = center1
			v6 = p1
			c2 = Vec2{
				X: v6.X - v5.X,
				Y: v6.Y - v5.Y,
			}
			v4 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
			goto _100
		_100:
			d1 = v4
			v8 = center2
			v17 = p1
			c2 = Vec2{
				X: v17.X - v8.X,
				Y: v17.Y - v8.Y,
			}
			v9 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
			goto _104
		_104:
			d2 = v9
			// Want to push the closest one last. Both have the same isLeaf, so they don't swap.
			if d1 < d2 {
				B2_SWAP_TEMP = hit[0]
				hit[0] = hit[int32(1)]
				hit[int32(1)] = B2_SWAP_TEMP
			}
		}
		i1 = 0
		for {
			if !(i1 < hitCount) {
				break
			}
			if isLeaf[i1] != 0 {
				v54 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(hit[i1])).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
				goto _107
			_107:
				proxyId = v54
				proxy = (*b2DynamicTree)(unsafe.Pointer(tree)).Proxies + uintptr(proxyId)*24
				if (*b2TreeProxy)(unsafe.Pointer(proxy)).CategoryBits&maskBits == uint64(0) {
					goto _105
				}
				(*(*RayCastInput)(unsafe.Add(nil, bp+16))).MaxFraction = maxFraction
				value = (*(*func(*_Stack, uintptr, int32, uint64, uintptr) float32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_callback})))(tls, bp+16, proxyId, (*b2TreeProxy)(unsafe.Pointer(proxy)).UserData, context)
				result.LeafVisits += int32(1)
				// The user may return -1 to indicate this shape should be skipped
				if value == float32FromFloat32(0) {
					// The client has terminated the ray cast.
					return result
				}
				if float32FromFloat32(0) < value && value <= maxFraction {
					// Update segment bounding box.
					maxFraction = value
					v11 = p1
					v4 = maxFraction
					v2 = d
					v5 = Vec2{
						X: v11.X + float32(v4*v2.X),
						Y: v11.Y + float32(v4*v2.Y),
					}
					goto _112
				_112:
					p2 = v5
					v11 = p1
					v2 = p2
					v4 = v11.X
					v9 = v2.X
					if v4 < v9 {
						v12 = v4
					} else {
						v12 = v9
					}
					v10 = v12
					goto _118
				_118:
					c.X = v10
					v13 = v11.Y
					v14 = v2.Y
					if v13 < v14 {
						v20 = v13
					} else {
						v20 = v14
					}
					v16 = v20
					goto _123
				_123:
					c.Y = v16
					v5 = c
					goto _126
				_126:
					(*(*AABB)(unsafe.Add(nil, bp))).LowerBound = v5
					v11 = p1
					v2 = p2
					v4 = v11.X
					v9 = v2.X
					if v4 > v9 {
						v12 = v4
					} else {
						v12 = v9
					}
					v10 = v12
					goto _132
				_132:
					c1.X = v10
					v13 = v11.Y
					v14 = v2.Y
					if v13 > v14 {
						v20 = v13
					} else {
						v20 = v14
					}
					v16 = v20
					goto _137
				_137:
					c1.Y = v16
					v5 = c1
					goto _140
				_140:
					(*(*AABB)(unsafe.Add(nil, bp))).UpperBound = v5
					v52 = *(*AABB)(unsafe.Add(nil, bp))
					goto _142
				_142:
					boxv = v52
				}
			} else {
				if stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)-int32FromInt32(1) {
					v54 = stackCount
					stackCount = stackCount + 1
					v55 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(hit[i1])).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
					goto _145
				_145:
					stack[v54] = v55
				} else {
					if v146 = !!(stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)-int32FromInt32(1)); !v146 {
						b2InternalAssert(tls, __ccgo_ts+9382, __ccgo_ts+9017, int32FromInt32(1138))
					}
					_ = v146 || bool(int32FromInt32(0) != 0)
				}
			}
			goto _105
		_105:
			;
			i1 = i1 + 1
		}
	}
	return result
}

// C documentation
//
//	// Follows structure of ray cast with small tweaks to handle a swept box.
func b2DynamicTree_CastBox(tls *_Stack, tree uintptr, input uintptr, maskBits uint64, __ccgo_fp_callback uintptr, context uintptr) (r1 TreeStats) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var B2_SWAP_TEMP, bv, node4, nodes1, proxy, root uintptr
	var abs_v, b10, b5, b9, c, c1, c2, c3, center1, center2, extension, h, p1, r, t, v2, v11, v21, v20, v23, v24, v26, v27, v28, v30, v31, v42, v44, v45, v46, v48, v49, v5, v60, v8, v9 Vec2
	var boxv, v62 b2AABBV
	var d1, d2, maxFraction, term1, term2, value, v12, v13, v15, v16, v17, v19, v22, v32, v33, v34, v36, v37, v38, v39, v41, v50, v51, v52, v54, v55, v56, v57, v59, v7 float32
	var hit [2]uintptr
	var hitCount, i, i1, pair, proxyId, stackCount, v64, v65, v67, v70 int32
	var isLeaf [2]uint8
	var nodeAABB, originAABB, v1, v4 AABB
	var result TreeStats
	var stack [512]int32
	var v68 uint8
	var v167 bool
	var _ /* subInput at bp+16 */ BoxCastInput
	var _ /* totalAABB at bp+0 */ AABB
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = B2_SWAP_TEMP, abs_v, b10, b5, b9, boxv, bv, c, c1, c2, c3, center1, center2, d1, d2, extension, h, hit, hitCount, i, i1, isLeaf, maxFraction, node4, nodeAABB, nodes1, originAABB, p1, pair, proxy, proxyId, r, result, root, stack, stackCount, t, term1, term2, v2, value, v1, v11, v12, v13, v15, v16, v167, v17, v19, v21, v20, v22, v23, v24, v26, v27, v28, v30, v31, v32, v33, v34, v36, v37, v38, v39, v4, v41, v42, v44, v45, v46, v48, v49, v5, v50, v51, v52, v54, v55, v56, v57, v59, v60, v62, v64, v65, v67, v68, v7, v70, v8, v9
	result = TreeStats{}
	if (*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCount == 0 {
		return result
	}
	originAABB = (*BoxCastInput)(unsafe.Pointer(input)).Box
	v1 = originAABB
	b9 = Vec2{
		X: float32(float32FromFloat32(0.5) * (v1.LowerBound.X + v1.UpperBound.X)),
		Y: float32(float32FromFloat32(0.5) * (v1.LowerBound.Y + v1.UpperBound.Y)),
	}
	v21 = b9
	goto _3
_3:
	p1 = v21
	v4 = originAABB
	b10 = Vec2{
		X: float32(float32FromFloat32(0.5) * (v4.UpperBound.X - v4.LowerBound.X)),
		Y: float32(float32FromFloat32(0.5) * (v4.UpperBound.Y - v4.LowerBound.Y)),
	}
	v5 = b10
	goto _6
_6:
	extension = v5
	// v is perpendicular to the segment.
	r = (*BoxCastInput)(unsafe.Pointer(input)).Translation
	v7 = float32FromFloat32(1)
	v8 = r
	v9 = Vec2{
		X: float32(-v7 * v8.Y),
		Y: float32(v7 * v8.X),
	}
	goto _10
_10:
	v2 = v9
	v11 = v2
	v12 = v11.X
	if v12 < float32FromInt32(0) {
		v15 = -v12
	} else {
		v15 = v12
	}
	v13 = v15
	goto _14
_14:
	b5.X = v13
	v16 = v11.Y
	if v16 < float32FromInt32(0) {
		v19 = -v16
	} else {
		v19 = v16
	}
	v17 = v19
	goto _18
_18:
	b5.Y = v17
	v20 = b5
	goto _21
_21:
	abs_v = v20
	// Separating axis for segment (Gino, p80).
	// |dot(v, p1 - c)| > dot(|v|, h)
	maxFraction = (*BoxCastInput)(unsafe.Pointer(input)).MaxFraction
	v22 = maxFraction
	v23 = (*BoxCastInput)(unsafe.Pointer(input)).Translation
	v24 = Vec2{
		X: float32(v22 * v23.X),
		Y: float32(v22 * v23.Y),
	}
	goto _25
_25:
	// Build total box for the cast
	t = v24
	v26 = originAABB.LowerBound
	v27 = t
	v28 = Vec2{
		X: v26.X + v27.X,
		Y: v26.Y + v27.Y,
	}
	goto _29
_29:
	v30 = originAABB.LowerBound
	v31 = v28
	v32 = v30.X
	v33 = v31.X
	if v32 < v33 {
		v36 = v32
	} else {
		v36 = v33
	}
	v34 = v36
	goto _35
_35:
	c.X = v34
	v37 = v30.Y
	v38 = v31.Y
	if v37 < v38 {
		v41 = v37
	} else {
		v41 = v38
	}
	v39 = v41
	goto _40
_40:
	c.Y = v39
	v42 = c
	goto _43
_43:
	v44 = originAABB.UpperBound
	v45 = t
	v46 = Vec2{
		X: v44.X + v45.X,
		Y: v44.Y + v45.Y,
	}
	goto _47
_47:
	v48 = originAABB.UpperBound
	v49 = v46
	v50 = v48.X
	v51 = v49.X
	if v50 > v51 {
		v54 = v50
	} else {
		v54 = v51
	}
	v52 = v54
	goto _53
_53:
	c1.X = v52
	v55 = v48.Y
	v56 = v49.Y
	if v55 > v56 {
		v59 = v55
	} else {
		v59 = v56
	}
	v57 = v59
	goto _58
_58:
	c1.Y = v57
	v60 = c1
	goto _61
_61:
	*(*AABB)(unsafe.Add(nil, bp)) = AABB{
		LowerBound: v42,
		UpperBound: v60,
	}
	v62 = *(*AABB)(unsafe.Add(nil, bp))
	goto _63
_63:
	boxv = v62
	nodes1 = (*b2DynamicTree)(unsafe.Pointer(tree)).Nodes
	stackCount = 0
	v64 = stackCount
	stackCount = stackCount + 1
	root = nodes1 + uintptr(_B2_ROOT_NODE)*32
	v68 = boolUint8((*b2TreeNode)(unsafe.Pointer(root)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
	goto _69
_69:
	if v68 != 0 {
		v67 = _B2_ROOT_NODE
	} else {
		v70 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(root)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
		goto _71
	_71:
		v67 = v70
	}
	v65 = v67
	goto _66
_66:
	stack[v64] = v65
	*(*BoxCastInput)(unsafe.Add(nil, bp+16)) = *(*BoxCastInput)(unsafe.Add(nil, input))
	for stackCount > 0 {
		stackCount = stackCount - 1
		v64 = stackCount
		pair = stack[v64]
		result.NodeVisits += int32(1)
		hitCount = 0
		i = 0
		for {
			if !(i < int32(2)) {
				break
			}
			node4 = nodes1 + uintptr(pair)*32 + uintptr(i)*32
			v62 = boxv
			bv = node4
			v68 = boolUint8(v62.LowerBound.X <= (*AABB)(unsafe.Pointer(bv)).UpperBound.X && v62.LowerBound.Y <= (*AABB)(unsafe.Pointer(bv)).UpperBound.Y && (*AABB)(unsafe.Pointer(bv)).LowerBound.X <= v62.UpperBound.X && (*AABB)(unsafe.Pointer(bv)).LowerBound.Y <= v62.UpperBound.Y)
			goto _76
		_76:
			if int32FromUint8(v68) == false1 {
				goto _73
			}
			// Separating axis for segment (Gino, p80).
			// |dot(v, p1 - c)| > dot(|v|, h)
			// radius extension is added to the node in this case
			nodeAABB = (*b2TreeNode)(unsafe.Pointer(node4)).Aabb
			v1 = nodeAABB
			b9 = Vec2{
				X: float32(float32FromFloat32(0.5) * (v1.LowerBound.X + v1.UpperBound.X)),
				Y: float32(float32FromFloat32(0.5) * (v1.LowerBound.Y + v1.UpperBound.Y)),
			}
			v21 = b9
			goto _79
		_79:
			c3 = v21
			v4 = nodeAABB
			b10 = Vec2{
				X: float32(float32FromFloat32(0.5) * (v4.UpperBound.X - v4.LowerBound.X)),
				Y: float32(float32FromFloat32(0.5) * (v4.UpperBound.Y - v4.LowerBound.Y)),
			}
			v5 = b10
			goto _82
		_82:
			v8 = v5
			v9 = extension
			v11 = Vec2{
				X: v8.X + v9.X,
				Y: v8.Y + v9.Y,
			}
			goto _86
		_86:
			h = v11
			v20 = p1
			v23 = c3
			v24 = Vec2{
				X: v20.X - v23.X,
				Y: v20.Y - v23.Y,
			}
			goto _90
		_90:
			v26 = v2
			v27 = v24
			v7 = float32(v26.X*v27.X) + float32(v26.Y*v27.Y)
			goto _94
		_94:
			v12 = v7
			if v12 < float32FromInt32(0) {
				v15 = -v12
			} else {
				v15 = v12
			}
			v13 = v15
			goto _97
		_97:
			term1 = v13
			v28 = abs_v
			v30 = h
			v16 = float32(v28.X*v30.X) + float32(v28.Y*v30.Y)
			goto _102
		_102:
			term2 = v16
			if term2 < term1 {
				goto _73
			}
			v68 = boolUint8((*b2TreeNode)(unsafe.Pointer(node4)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
			goto _104
		_104:
			isLeaf[hitCount] = v68
			hit[hitCount] = node4
			hitCount = hitCount + int32(1)
			goto _73
		_73:
			;
			i = i + 1
		}
		if hitCount == int32(2) && int32FromUint8(isLeaf[0]) == false1 && int32FromUint8(isLeaf[int32(1)]) == false1 {
			v1 = (*b2TreeNode)(unsafe.Pointer(hit[0])).Aabb
			b9 = Vec2{
				X: float32(float32FromFloat32(0.5) * (v1.LowerBound.X + v1.UpperBound.X)),
				Y: float32(float32FromFloat32(0.5) * (v1.LowerBound.Y + v1.UpperBound.Y)),
			}
			v21 = b9
			goto _107
		_107:
			center1 = v21
			v4 = (*b2TreeNode)(unsafe.Pointer(hit[int32(1)])).Aabb
			b9 = Vec2{
				X: float32(float32FromFloat32(0.5) * (v4.LowerBound.X + v4.UpperBound.X)),
				Y: float32(float32FromFloat32(0.5) * (v4.LowerBound.Y + v4.UpperBound.Y)),
			}
			v5 = b9
			goto _110
		_110:
			center2 = v5
			v8 = center1
			v9 = p1
			c2 = Vec2{
				X: v9.X - v8.X,
				Y: v9.Y - v8.Y,
			}
			v7 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
			goto _114
		_114:
			d1 = v7
			v11 = center2
			v20 = p1
			c2 = Vec2{
				X: v20.X - v11.X,
				Y: v20.Y - v11.Y,
			}
			v12 = float32(c2.X*c2.X) + float32(c2.Y*c2.Y)
			goto _118
		_118:
			d2 = v12
			// Want to push the closest one last. Both have the same isLeaf, so they don't swap.
			if d1 < d2 {
				B2_SWAP_TEMP = hit[0]
				hit[0] = hit[int32(1)]
				hit[int32(1)] = B2_SWAP_TEMP
			}
		}
		i1 = 0
		for {
			if !(i1 < hitCount) {
				break
			}
			if isLeaf[i1] != 0 {
				v64 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(hit[i1])).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
				goto _121
			_121:
				proxyId = v64
				proxy = (*b2DynamicTree)(unsafe.Pointer(tree)).Proxies + uintptr(proxyId)*24
				if (*b2TreeProxy)(unsafe.Pointer(proxy)).CategoryBits&maskBits == uint64(0) {
					goto _119
				}
				(*(*BoxCastInput)(unsafe.Add(nil, bp+16))).MaxFraction = maxFraction
				value = (*(*func(*_Stack, uintptr, int32, uint64, uintptr) float32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_callback})))(tls, bp+16, proxyId, (*b2TreeProxy)(unsafe.Pointer(proxy)).UserData, context)
				result.LeafVisits += int32(1)
				// The user may return -1 to indicate this shape should be skipped
				if value == float32FromFloat32(0) {
					// The client has terminated the ray cast.
					return result
				}
				if float32FromFloat32(0) < value && value < maxFraction {
					// Update cast bounding box.
					maxFraction = value
					v7 = maxFraction
					v21 = (*BoxCastInput)(unsafe.Pointer(input)).Translation
					v5 = Vec2{
						X: float32(v7 * v21.X),
						Y: float32(v7 * v21.Y),
					}
					goto _125
				_125:
					t = v5
					v21 = originAABB.LowerBound
					v5 = t
					v8 = Vec2{
						X: v21.X + v5.X,
						Y: v21.Y + v5.Y,
					}
					goto _129
				_129:
					v9 = originAABB.LowerBound
					v11 = v8
					v7 = v9.X
					v12 = v11.X
					if v7 < v12 {
						v15 = v7
					} else {
						v15 = v12
					}
					v13 = v15
					goto _135
				_135:
					c.X = v13
					v16 = v9.Y
					v17 = v11.Y
					if v16 < v17 {
						v22 = v16
					} else {
						v22 = v17
					}
					v19 = v22
					goto _140
				_140:
					c.Y = v19
					v20 = c
					goto _143
				_143:
					(*(*AABB)(unsafe.Add(nil, bp))).LowerBound = v20
					v21 = originAABB.UpperBound
					v5 = t
					v8 = Vec2{
						X: v21.X + v5.X,
						Y: v21.Y + v5.Y,
					}
					goto _147
				_147:
					v9 = originAABB.UpperBound
					v11 = v8
					v7 = v9.X
					v12 = v11.X
					if v7 > v12 {
						v15 = v7
					} else {
						v15 = v12
					}
					v13 = v15
					goto _153
				_153:
					c1.X = v13
					v16 = v9.Y
					v17 = v11.Y
					if v16 > v17 {
						v22 = v16
					} else {
						v22 = v17
					}
					v19 = v22
					goto _158
				_158:
					c1.Y = v19
					v20 = c1
					goto _161
				_161:
					(*(*AABB)(unsafe.Add(nil, bp))).UpperBound = v20
					v62 = *(*AABB)(unsafe.Add(nil, bp))
					goto _163
				_163:
					boxv = v62
				}
			} else {
				if stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)-int32FromInt32(1) {
					v64 = stackCount
					stackCount = stackCount + 1
					v65 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(hit[i1])).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
					goto _166
				_166:
					stack[v64] = v65
				} else {
					if v167 = !!(stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)-int32FromInt32(1)); !v167 {
						b2InternalAssert(tls, __ccgo_ts+9382, __ccgo_ts+9017, int32FromInt32(1282))
					}
					_ = v167 || bool(int32FromInt32(0) != 0)
				}
			}
			goto _119
		_119:
			;
			i1 = i1 + 1
		}
	}
	return result
}

// C documentation
//
//	// Rebuild the stale parts of the tree. The entire tree is put into DFS order. This makes
//	// refitting much faster. This is done async with threading, so the cost is hidden.
//	// Not safe to access tree during this operation.
func b2DynamicTree_Rebuild(tls *_Stack, tree uintptr, fullBuild uint8) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var B2_SWAP_TEMP, leafCenters, leafIndices, leaves, nodes, root uintptr
	var b, v14 Vec2
	var i, leafCount, newCapacity, oldCapacity, pair, proxyCount, stackCount, v10, v11 int32
	var node3 b2TreeNode
	var stack [512]int32
	var v1, v4 uint8
	var v13 AABB
	var v3, v8, v9 bool
	var _ /* node at bp+0 */ b2TreeNode
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = B2_SWAP_TEMP, b, i, leafCenters, leafCount, leafIndices, leaves, newCapacity, node3, nodes, oldCapacity, pair, proxyCount, root, stack, stackCount, v1, v10, v11, v13, v14, v3, v4, v8, v9
	proxyCount = (*b2DynamicTree)(unsafe.Pointer(tree)).ProxyCount
	if proxyCount == 0 {
		return 0
	}
	// An unordered tree is rebuilt even when nothing moved, the sweep refit needs the order
	nodes = (*b2DynamicTree)(unsafe.Pointer(tree)).Nodes
	root = nodes + uintptr(_B2_ROOT_NODE)*32
	if v3 = int32FromUint8(fullBuild) == false1; v3 {
		v1 = boolUint8((*b2TreeNode)(unsafe.Pointer(root)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
		goto _2
	_2:
	}
	if v3 && int32FromUint8(v1) == false1 && (*b2DynamicTree)(unsafe.Pointer(tree)).DfsOrdered != 0 {
		return 0
	}
	if (*b2DynamicTree)(unsafe.Pointer(tree)).SwapNodes == uintptrFromInt32(0) {
		(*b2DynamicTree)(unsafe.Pointer(tree)).SwapNodes = b2AllocZero(tls, uint64FromInt32((*b2DynamicTree)(unsafe.Pointer(tree)).NodeCapacity)*uint64(32))
	}
	if proxyCount > (*b2DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity {
		oldCapacity = (*b2DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity
		newCapacity = proxyCount + proxyCount/int32(2)
		b2Free(tls, (*b2DynamicTree)(unsafe.Pointer(tree)).LeafIndices, uint64FromInt32(oldCapacity)*uint64(4))
		(*b2DynamicTree)(unsafe.Pointer(tree)).LeafIndices = b2Alloc(tls, uint64FromInt32(newCapacity)*uint64(4))
		b2Free(tls, (*b2DynamicTree)(unsafe.Pointer(tree)).LeafNodes, uint64FromInt32(oldCapacity)*uint64(32))
		(*b2DynamicTree)(unsafe.Pointer(tree)).LeafNodes = b2Alloc(tls, uint64FromInt32(newCapacity)*uint64(32))
		b2Free(tls, (*b2DynamicTree)(unsafe.Pointer(tree)).LeafCenters, uint64FromInt32(oldCapacity)*uint64(8))
		(*b2DynamicTree)(unsafe.Pointer(tree)).LeafCenters = b2Alloc(tls, uint64FromInt32(newCapacity)*uint64(8))
		(*b2DynamicTree)(unsafe.Pointer(tree)).RebuildCapacity = newCapacity
	}
	leafIndices = (*b2DynamicTree)(unsafe.Pointer(tree)).LeafIndices
	leaves = (*b2DynamicTree)(unsafe.Pointer(tree)).LeafNodes
	leafCenters = (*b2DynamicTree)(unsafe.Pointer(tree)).LeafCenters
	// Gather build leaves. A marked internal node is descended and abandoned in the old array.
	// Everything else is a build leaf, kept subtrees included. A leaf root, or an unmarked root
	// on an unordered array, is the one build leaf.
	leafCount = 0
	stackCount = 0
	v1 = boolUint8((*b2TreeNode)(unsafe.Pointer(root)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
	goto _5
_5:
	;
	if v8 = int32FromUint8(v1) == false1; v8 {
		if v3 = fullBuild != 0; !v3 {
			v4 = boolUint8((*b2TreeNode)(unsafe.Pointer(root)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
			goto _7
		_7:
		}
	}
	if v8 && (v3 || v4 != 0) {
		v10 = stackCount
		stackCount = stackCount + 1
		v11 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(root)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
		goto _12
	_12:
		stack[v10] = v11
	} else {
		node3 = *(*b2TreeNode)(unsafe.Add(nil, root))
		node3.FlagIndex &= ^(uint32FromUint32(1) << int32FromInt32(30))
		*(*int32)(unsafe.Add(nil, leafIndices)) = 0
		*(*b2TreeNode)(unsafe.Add(nil, leaves)) = node3
		v13 = node3.Aabb
		b = Vec2{
			X: float32(float32FromFloat32(0.5) * (v13.LowerBound.X + v13.UpperBound.X)),
			Y: float32(float32FromFloat32(0.5) * (v13.LowerBound.Y + v13.UpperBound.Y)),
		}
		v14 = b
		goto _15
	_15:
		*(*Vec2)(unsafe.Add(nil, leafCenters)) = v14
		leafCount = int32(1)
	}
	for stackCount > 0 {
		stackCount = stackCount - 1
		v10 = stackCount
		pair = stack[v10]
		i = 0
		for {
			if !(i < int32(2)) {
				break
			}
			*(*b2TreeNode)(unsafe.Pointer(bp)) = b2TreeNode{}
			*(*b2TreeNode)(unsafe.Pointer(bp)) = *(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(pair+i)*32))
			v1 = boolUint8((*b2TreeNode)(unsafe.Pointer(bp)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
			goto _19
		_19:
			;
			if v8 = int32FromUint8(v1) == false1; v8 {
				if v3 = fullBuild != 0; !v3 {
					v4 = boolUint8((*b2TreeNode)(unsafe.Pointer(bp)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
					goto _21
				_21:
				}
			}
			if v8 && (v3 || v4 != 0) {
				// todo fail gracefully if the stack runs out in release
				if v9 = !!(stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)); !v9 {
					b2InternalAssert(tls, __ccgo_ts+4795, __ccgo_ts+9017, int32FromInt32(1872))
				}
				_ = v9 || bool(int32FromInt32(0) != 0)
				v10 = stackCount
				stackCount = stackCount + 1
				v11 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(bp)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
				goto _27
			_27:
				stack[v10] = v11
				goto _17
			}
			(*(*b2TreeNode)(unsafe.Add(nil, bp))).FlagIndex &= ^(uint32FromUint32(1) << int32FromInt32(30))
			*(*int32)(unsafe.Add(nil, leafIndices+uintptr(leafCount)*4)) = leafCount
			*(*b2TreeNode)(unsafe.Add(nil, leaves+uintptr(leafCount)*32)) = *(*b2TreeNode)(unsafe.Add(nil, bp))
			v13 = (*(*b2TreeNode)(unsafe.Add(nil, bp))).Aabb
			b = Vec2{
				X: float32(float32FromFloat32(0.5) * (v13.LowerBound.X + v13.UpperBound.X)),
				Y: float32(float32FromFloat32(0.5) * (v13.LowerBound.Y + v13.UpperBound.Y)),
			}
			v14 = b
			goto _30
		_30:
			*(*Vec2)(unsafe.Add(nil, leafCenters+uintptr(leafCount)*8)) = v14
			leafCount = leafCount + int32(1)
			goto _17
		_17:
			;
			i = i + 1
		}
	}
	if v3 = !!(0 < leafCount && leafCount <= proxyCount); !v3 {
		b2InternalAssert(tls, __ccgo_ts+9497, __ccgo_ts+9017, int32FromInt32(1889))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	b2BuildTree(tls, tree, leafCount)
	// The spare is now the tree and the old array is the spare. The build is dense, so the
	// stale tail of the old array sits above the end and is never read.
	B2_SWAP_TEMP = (*b2DynamicTree)(unsafe.Pointer(tree)).Nodes
	(*b2DynamicTree)(unsafe.Pointer(tree)).Nodes = (*b2DynamicTree)(unsafe.Pointer(tree)).SwapNodes
	(*b2DynamicTree)(unsafe.Pointer(tree)).SwapNodes = B2_SWAP_TEMP
	(*b2DynamicTree)(unsafe.Pointer(tree)).PairFreeList = -int32FromInt32(1)
	(*b2DynamicTree)(unsafe.Pointer(tree)).DfsOrdered = boolUint8(true1 != 0)
	b2DynamicTree_Validate(tls, tree)
	b2DynamicTree_ValidateNoMoved(tls, tree)
	return leafCount
}

// C documentation
//
//	// Set the moved flag on the ancestors of the proxy. Serial use case.
func b2DynamicTree_MarkProxyMovedSerial(tls *_Stack, tree uintptr, proxyId int32) {
	var index int32
	var nodes, parents uintptr
	_, _, _ = index, nodes, parents
	nodes = (*b2DynamicTree)(unsafe.Pointer(tree)).Nodes
	parents = (*b2DynamicTree)(unsafe.Pointer(tree)).Parents
	index = (*(*b2TreeProxy)(unsafe.Add(nil, (*b2DynamicTree)(unsafe.Pointer(tree)).Proxies+uintptr(proxyId)*24))).Node
	for index != -int32(1) {
		(*(*b2TreeNode)(unsafe.Add(nil, nodes+uintptr(index)*32))).FlagIndex |= uint32FromUint32(1) << int32FromInt32(30)
		index = *(*int32_t)(unsafe.Add(nil, parents+uintptr(index)*4))
	}
}

// C documentation
//
//	// Update a proxy AABB and flag the ancestors as moved. Thread-safe using atomics.
func b2DynamicTree_MarkProxyMoved(tls *_Stack, tree uintptr, proxyId int32, aabb AABB) {
	var index int32
	var node, nodes, parents uintptr
	var previousFlags, v1 uint32_t
	_, _, _, _, _, _ = index, node, nodes, parents, previousFlags, v1
	nodes = (*b2DynamicTree)(unsafe.Pointer(tree)).Nodes
	parents = (*b2DynamicTree)(unsafe.Pointer(tree)).Parents
	index = (*(*b2TreeProxy)(unsafe.Add(nil, (*b2DynamicTree)(unsafe.Pointer(tree)).Proxies+uintptr(proxyId)*24))).Node
	node = nodes + uintptr(index)*32
	// This is not raced since it is the leaf.
	(*b2TreeNode)(unsafe.Pointer(node)).Aabb = aabb
	*(*uint32_t)(unsafe.Add(nil, node+24)) |= uint32FromUint32(1) << int32FromInt32(30)
	index = *(*int32_t)(unsafe.Add(nil, parents+uintptr(index)*4))
	for index != -int32(1) {
		node = nodes + uintptr(index)*32
		// Read first to avoid the FetchOr if possible.
		v1 = atomicLoadNUint32(node+24, int32FromInt32(__ATOMIC_RELAXED))
		goto _2
	_2:
		if v1&(uint32FromUint32(1)<<int32FromInt32(30)) != 0 {
			break
		}
		v1 = __atomic_fetch_orUint32(tls, node+24, uint32FromUint32(1)<<int32FromInt32(30), int32FromInt32(__ATOMIC_SEQ_CST))
		goto _4
	_4:
		previousFlags = v1
		if previousFlags&(uint32FromUint32(1)<<int32FromInt32(30)) != 0 {
			// Ancestor already visited.
			break
		}
		index = *(*int32_t)(unsafe.Add(nil, parents+uintptr(index)*4))
	}
}

// C documentation
//
//	// Clear the moved flags from the entire tree.
func b2DynamicTree_ClearMoved(tls *_Stack, tree uintptr) {
	var i, pair, stackCount, v5, v6 int32
	var node3, nodes, root uintptr
	var stack [512]int32
	var v1 uint8
	var v15 bool
	_, _, _, _, _, _, _, _, _, _, _ = i, node3, nodes, pair, root, stack, stackCount, v1, v15, v5, v6
	nodes = (*b2DynamicTree)(unsafe.Pointer(tree)).Nodes
	root = nodes + uintptr(_B2_ROOT_NODE)*32
	v1 = boolUint8((*b2TreeNode)(unsafe.Pointer(root)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
	goto _2
_2:
	if int32FromUint8(v1) == false1 {
		return
	}
	*(*uint32_t)(unsafe.Add(nil, root+24)) &= ^(uint32FromUint32(1) << int32FromInt32(30))
	v1 = boolUint8((*b2TreeNode)(unsafe.Pointer(root)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
	goto _4
_4:
	if v1 != 0 {
		return
	}
	stackCount = 0
	v5 = stackCount
	stackCount = stackCount + 1
	v6 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(root)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
	goto _7
_7:
	stack[v5] = v6
	for stackCount > 0 {
		stackCount = stackCount - 1
		v5 = stackCount
		pair = stack[v5]
		i = 0
		for {
			if !(i < int32(2)) {
				break
			}
			node3 = nodes + uintptr(pair)*32 + uintptr(i)*32
			if (*b2TreeNode)(unsafe.Pointer(node3)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) != 0 {
				*(*uint32_t)(unsafe.Add(nil, node3+24)) &= ^(uint32FromUint32(1) << int32FromInt32(30))
				v1 = boolUint8((*b2TreeNode)(unsafe.Pointer(node3)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
				goto _11
			_11:
				if int32FromUint8(v1) == false1 {
					if stackCount < int32(_B2_TREE_STACK_SIZE) {
						v5 = stackCount
						stackCount = stackCount + 1
						v6 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(node3)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
						goto _14
					_14:
						stack[v5] = v6
					} else {
						// Bad stuff will happen if the moved flags don't get cleared.
						if v15 = !!(stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)); !v15 {
							b2InternalAssert(tls, __ccgo_ts+4795, __ccgo_ts+9017, int32FromInt32(2004))
						}
						_ = v15 || bool(int32FromInt32(0) != 0)
					}
				}
			}
			goto _9
		_9:
			;
			i = i + 1
		}
	}
}

func b2DynamicTree_GatherMovedProxies(tls *_Stack, tree uintptr, proxyIds uintptr) (r int32) {
	var count, i, pair, stackCount, v5, v7 int32
	var node4, nodes, root uintptr
	var stack [512]int32
	var v1 uint8
	var v22 bool
	_, _, _, _, _, _, _, _, _, _, _, _ = count, i, node4, nodes, pair, root, stack, stackCount, v1, v22, v5, v7
	nodes = (*b2DynamicTree)(unsafe.Pointer(tree)).Nodes
	root = nodes + uintptr(_B2_ROOT_NODE)*32
	v1 = boolUint8((*b2TreeNode)(unsafe.Pointer(root)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
	goto _2
_2:
	if int32FromUint8(v1) == false1 {
		return 0
	}
	v1 = boolUint8((*b2TreeNode)(unsafe.Pointer(root)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
	goto _4
_4:
	if v1 != 0 {
		v5 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(root)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
		goto _6
	_6:
		*(*int32)(unsafe.Add(nil, proxyIds)) = v5
		return int32(1)
	}
	count = 0
	stackCount = 0
	v5 = stackCount
	stackCount = stackCount + 1
	v7 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(root)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
	goto _9
_9:
	stack[v5] = v7
	for stackCount > 0 {
		stackCount = stackCount - 1
		v5 = stackCount
		pair = stack[v5]
		i = 0
		for {
			if !(i < int32(2)) {
				break
			}
			node4 = nodes + uintptr(pair)*32 + uintptr(i)*32
			v1 = boolUint8((*b2TreeNode)(unsafe.Pointer(node4)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
			goto _13
		_13:
			if int32FromUint8(v1) == false1 {
				goto _11
			}
			v1 = boolUint8((*b2TreeNode)(unsafe.Pointer(node4)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(31)) == uint32FromUint32(1)<<int32FromInt32(31))
			goto _15
		_15:
			if v1 != 0 {
				v5 = count
				count = count + 1
				v7 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(node4)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
				goto _18
			_18:
				*(*int32)(unsafe.Add(nil, proxyIds+uintptr(v5)*4)) = v7
			} else {
				if stackCount < int32(_B2_TREE_STACK_SIZE) {
					v5 = stackCount
					stackCount = stackCount + 1
					v7 = int32FromUint32((*b2TreeNode)(unsafe.Pointer(node4)).FlagIndex & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
					goto _21
				_21:
					stack[v5] = v7
				} else {
					if v22 = !!(stackCount < int32FromInt32(_B2_TREE_STACK_SIZE)); !v22 {
						b2InternalAssert(tls, __ccgo_ts+4795, __ccgo_ts+9017, int32FromInt32(2055))
					}
					_ = v22 || bool(int32FromInt32(0) != 0)
				}
			}
			goto _11
		_11:
			;
			i = i + 1
		}
	}
	return count
}

func b2DynamicTree_Refit(tls *_Stack, tree1 uintptr) {
	var c, v11, v12, v33 AABB
	var children1, children2, pair1, v6, v7 int32
	var flags1, flags2 uint32_t
	var node1, nodes, v8 uintptr
	var refit1, refit2, v1, v3 uint8
	var v13, v14, v15, v17, v18, v19, v20, v22, v23, v24, v25, v27, v28, v29, v30, v32 float32
	var v9 b2AABBV
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, children1, children2, flags1, flags2, node1, nodes, pair1, refit1, refit2, v1, v11, v12, v13, v14, v15, v17, v18, v19, v20, v22, v23, v24, v25, v27, v28, v29, v3, v30, v32, v33, v6, v7, v8, v9
	v3 = boolUint8((*b2TreeNode)(unsafe.Pointer((*b2DynamicTree)(unsafe.Pointer(tree1)).Nodes+uintptr(_B2_ROOT_NODE)*32)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
	goto _4
_4:
	v1 = v3
	goto _2
_2:
	if int32FromUint8(v1) == false1 {
		return
	}
	nodes = (*b2DynamicTree)(unsafe.Pointer(tree1)).Nodes
	if int32FromUint8((*b2DynamicTree)(unsafe.Pointer(tree1)).DfsOrdered) == false1 {
		// This only happens in unit tests.
		b2RefitSubtree(tls, nodes, _B2_ROOT_NODE)
		return
	}
	pair1 = (*b2DynamicTree)(unsafe.Pointer(tree1)).NodeEnd - int32(2)
	for {
		if !(pair1 >= 0) {
			break
		}
		node1 = nodes + uintptr(pair1)*32
		flags1 = (*(*b2TreeNode)(unsafe.Add(nil, node1))).FlagIndex
		flags2 = (*(*b2TreeNode)(unsafe.Add(nil, node1+1*32))).FlagIndex
		// Did either move?
		if (flags1|flags2)&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32(0) {
			goto _5
		}
		// Is the node internal and moved?
		refit1 = boolUint8(flags1&(uint32FromUint32(1)<<int32FromInt32(31)|uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
		refit2 = boolUint8(flags2&(uint32FromUint32(1)<<int32FromInt32(31)|uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
		if refit1 != 0 {
			v6 = int32FromUint32(flags1 & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
		} else {
			v6 = pair1
		}
		// If the node needs a refit then get the child index. Otherwise use a valid dummy index.
		// The store function won't actually modify the node at the dummy index. This
		// avoids an unpredictable branch. SIMD FTW.
		children1 = v6
		if refit2 != 0 {
			v7 = int32FromUint32(flags2 & (uint32FromUint32(0xFFFFFFFF) & ^(uint32FromUint32(1)<<int32FromInt32(30) | uint32FromUint32(1)<<int32FromInt32(31))))
		} else {
			v7 = pair1
		}
		children2 = v7
		// Conditionally store the union.
		v8 = nodes + uintptr(children1)*32
		v11 = (*(*b2TreeNode)(unsafe.Add(nil, v8))).Aabb
		v12 = (*(*b2TreeNode)(unsafe.Add(nil, v8+1*32))).Aabb
		v13 = v11.LowerBound.X
		v14 = v12.LowerBound.X
		if v13 < v14 {
			v17 = v13
		} else {
			v17 = v14
		}
		v15 = v17
		goto _16
	_16:
		c.LowerBound.X = v15
		v18 = v11.LowerBound.Y
		v19 = v12.LowerBound.Y
		if v18 < v19 {
			v22 = v18
		} else {
			v22 = v19
		}
		v20 = v22
		goto _21
	_21:
		c.LowerBound.Y = v20
		v23 = v11.UpperBound.X
		v24 = v12.UpperBound.X
		if v23 > v24 {
			v27 = v23
		} else {
			v27 = v24
		}
		v25 = v27
		goto _26
	_26:
		c.UpperBound.X = v25
		v28 = v11.UpperBound.Y
		v29 = v12.UpperBound.Y
		if v28 > v29 {
			v32 = v28
		} else {
			v32 = v29
		}
		v30 = v32
		goto _31
	_31:
		c.UpperBound.Y = v30
		v33 = c
		goto _34
	_34:
		v9 = v33
		goto _10
	_10:
		if refit1 != 0 {
			*(*AABB)(unsafe.Add(nil, node1)) = v9
		}
		v8 = nodes + uintptr(children2)*32
		v11 = (*(*b2TreeNode)(unsafe.Add(nil, v8))).Aabb
		v12 = (*(*b2TreeNode)(unsafe.Add(nil, v8+1*32))).Aabb
		v13 = v11.LowerBound.X
		v14 = v12.LowerBound.X
		if v13 < v14 {
			v17 = v13
		} else {
			v17 = v14
		}
		v15 = v17
		goto _43
	_43:
		c.LowerBound.X = v15
		v18 = v11.LowerBound.Y
		v19 = v12.LowerBound.Y
		if v18 < v19 {
			v22 = v18
		} else {
			v22 = v19
		}
		v20 = v22
		goto _48
	_48:
		c.LowerBound.Y = v20
		v23 = v11.UpperBound.X
		v24 = v12.UpperBound.X
		if v23 > v24 {
			v27 = v23
		} else {
			v27 = v24
		}
		v25 = v27
		goto _53
	_53:
		c.UpperBound.X = v25
		v28 = v11.UpperBound.Y
		v29 = v12.UpperBound.Y
		if v28 > v29 {
			v32 = v28
		} else {
			v32 = v29
		}
		v30 = v32
		goto _58
	_58:
		c.UpperBound.Y = v30
		v33 = c
		goto _61
	_61:
		v9 = v33
		goto _37
	_37:
		if refit2 != 0 {
			*(*AABB)(unsafe.Add(nil, node1+1*32)) = v9
		}
		goto _5
	_5:
		;
		pair1 = pair1 - int32(2)
	}
}
