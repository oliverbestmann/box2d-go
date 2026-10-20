package b2

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b2CreateBroadPhase(tls *_Stack, bp uintptr, capacity uintptr) {
	var dynamicCapacity, kinematicCapacity, staticCapacity, v1, v2, v3, v5 int32
	_, _, _, _, _, _, _ = dynamicCapacity, kinematicCapacity, staticCapacity, v1, v2, v3, v5
	v1 = int32(32)
	v2 = int32(2) * (*Capacity)(unsafe.Pointer(capacity)).ContactCount
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	(*b2BroadPhase)(unsafe.Pointer(bp)).PairSet = b2CreateSet(tls, v3)
	v1 = int32(16)
	v2 = (*Capacity)(unsafe.Pointer(capacity)).StaticShapeCount
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _9
_9:
	staticCapacity = v3
	b2CreateDynamicTreeInternal(tls, bp+uintptr(b2_staticBody)*104, staticCapacity)
	kinematicCapacity = int32(16)
	b2CreateDynamicTreeInternal(tls, bp+uintptr(b2_kinematicBody)*104, kinematicCapacity)
	v1 = int32(16)
	v2 = (*Capacity)(unsafe.Pointer(capacity)).DynamicShapeCount
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _14
_14:
	dynamicCapacity = v3
	b2CreateDynamicTreeInternal(tls, bp+uintptr(b2_dynamicBody)*104, dynamicCapacity)
	(*b2BroadPhase)(unsafe.Pointer(bp)).MovedSiblings = uintptrFromInt32(0)
}

func b2DestroyBroadPhase(tls *_Stack, bp uintptr) {
	var i int32
	_ = i
	i = 0
	for {
		if !(i < int32(b2_bodyTypeCount)) {
			break
		}
		b2DestroyDynamicTreeInternal(tls, bp+uintptr(i)*104)
		goto _1
	_1:
		;
		i = i + 1
	}
	b2DestroySet(tls, bp+320)
	memset(tls, bp, 0, uint64(336))
}

func b2BroadPhase_CreateProxy(tls *_Stack, bp uintptr, proxyType BodyType, aabb AABB, categoryBits uint64, shapeIndex int32, forcePairCreation uint8) (r int32) {
	var mark uint8
	var proxyId, proxyKey int32
	var v1 bool
	_, _, _, _ = mark, proxyId, proxyKey, v1
	if v1 = !!(0 <= proxyType && proxyType < int32(b2_bodyTypeCount)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+4483, __ccgo_ts+4530, int32FromInt32(58))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mark = boolUint8(proxyType != int32(b2_staticBody) || forcePairCreation != 0)
	proxyId = b2CreateTreeProxyInternal(tls, bp+uintptr(proxyType)*104, aabb, categoryBits, uint64FromInt32(shapeIndex), mark)
	proxyKey = proxyId<<int32(2) | proxyType
	return proxyKey
}

func b2BroadPhase_DestroyProxy(tls *_Stack, bp uintptr, proxyKey int32) {
	var proxyId int32
	var proxyType b2BodyType1
	var v1 bool
	_, _, _ = proxyId, proxyType, v1
	proxyType = proxyKey & int32FromInt32(3)
	proxyId = proxyKey >> int32(2)
	if v1 = !!(0 <= proxyType && proxyType < int32(b2_bodyTypeCount)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+4483, __ccgo_ts+4530, int32FromInt32(72))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	b2DestroyTreeProxy(tls, bp+uintptr(proxyType)*104, proxyId)
}

func b2BroadPhase_MoveProxy(tls *_Stack, bp uintptr, proxyKey int32, aabb AABB) {
	var mark uint8
	var proxyId int32
	var proxyType b2BodyType1
	_, _, _ = mark, proxyId, proxyType
	proxyType = proxyKey & int32FromInt32(3)
	proxyId = proxyKey >> int32(2)
	mark = boolUint8(true1 != 0)
	b2DynamicTree_MoveProxyInternal(tls, bp+uintptr(proxyType)*104, proxyId, aabb, mark)
}

func b2UpdateBroadPhasePairs(tls *_Stack, world uintptr) {
	bp1 := tls.Alloc(8224)
	defer tls.Free(8224)
	var alloc, bp, dynamicTree, pairKeys, shapeA, shapeB, tempKeys, workerKeys, v5 uintptr
	var crossMoveCount, dynamicMoveCount, i, i1, i2, i3, keyCount, kinematicSeedCount, pairCapacity, pairCount, shapeIdA, shapeIdB, staticSeedCount, v21, v22, v23, v25 int32
	var needUpdate, v1, v3, v6, v8 uint8
	var v10, v16 bool
	var _ /* crossContext at bp+8192 */ b2CrossContext
	var _ /* crossSeeds at bp+0 */ [128]b2NodePair
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = alloc, bp, crossMoveCount, dynamicMoveCount, dynamicTree, i, i1, i2, i3, keyCount, kinematicSeedCount, needUpdate, pairCapacity, pairCount, pairKeys, shapeA, shapeB, shapeIdA, shapeIdB, staticSeedCount, tempKeys, workerKeys, v1, v10, v16, v21, v22, v23, v25, v3, v5, v6, v8
	bp = world + 40
	v3 = boolUint8((*b2TreeNode)(unsafe.Pointer((*b2DynamicTree)(unsafe.Pointer(bp+uintptr(b2_staticBody)*104)).Nodes+uintptr(_B2_ROOT_NODE)*32)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
	goto _4
_4:
	v1 = v3
	goto _2
_2:
	needUpdate = v1
	if v10 = needUpdate != 0; !v10 {
		v5 = bp + uintptr(b2_kinematicBody)*104
		v8 = boolUint8((*b2TreeNode)(unsafe.Pointer((*b2DynamicTree)(unsafe.Pointer(v5)).Nodes+uintptr(_B2_ROOT_NODE)*32)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
		goto _9
	_9:
		;
		v6 = boolUint8(v8 != 0 || int32FromUint8((*b2DynamicTree)(unsafe.Pointer(v5)).DfsOrdered) == false1)
		goto _7
	_7:
	}
	needUpdate = boolUint8(v10 || v6 != 0)
	if v10 = needUpdate != 0; !v10 {
		v5 = bp + uintptr(b2_dynamicBody)*104
		v3 = boolUint8((*b2TreeNode)(unsafe.Pointer((*b2DynamicTree)(unsafe.Pointer(v5)).Nodes+uintptr(_B2_ROOT_NODE)*32)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
		goto _15
	_15:
		;
		v1 = boolUint8(v3 != 0 || int32FromUint8((*b2DynamicTree)(unsafe.Pointer(v5)).DfsOrdered) == false1)
		goto _13
	_13:
	}
	needUpdate = boolUint8(v10 || v1 != 0)
	if int32FromUint8(needUpdate) == false1 {
		return
	}
	alloc = world
	i = 0
	for {
		if !(i < (*b2World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		(*(*b2TaskContext)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i)*136))).PairKeys.Count = 0
		goto _17
	_17:
		;
		i = i + 1
	}
	// Generate pairs by querying the dynamic body tree against itself and against
	// the kinematic and static trees.
	// Get the sibling pairs of the dynamic body tree that have moved.
	dynamicTree = bp + uintptr(b2_dynamicBody)*104
	pairCapacity = (*b2DynamicTree)(unsafe.Pointer(dynamicTree)).NodeEnd / int32(2)
	(*b2BroadPhase)(unsafe.Pointer(bp)).MovedSiblings = b2StackAlloc(tls, alloc, int32FromUint64(uint64FromInt32(pairCapacity)*uint64(4)), __ccgo_ts+4860)
	dynamicMoveCount = b2GatherMovedSiblings(tls, dynamicTree, (*b2BroadPhase)(unsafe.Pointer(bp)).MovedSiblings)
	staticSeedCount = b2GatherCrossSeeds(tls, dynamicTree, bp+uintptr(b2_staticBody)*104, bp1)
	if v10 = !!(staticSeedCount <= int32FromInt32(_B2_CROSS_SEED_COUNT)); !v10 {
		b2InternalAssert(tls, __ccgo_ts+4872, __ccgo_ts+4530, int32FromInt32(628))
	}
	_ = v10 || bool(int32FromInt32(0) != 0)
	kinematicSeedCount = b2GatherCrossSeeds(tls, dynamicTree, bp+uintptr(b2_kinematicBody)*104, bp1+uintptr(staticSeedCount)*64)
	if v10 = !!(kinematicSeedCount <= int32FromInt32(_B2_CROSS_SEED_COUNT)); !v10 {
		b2InternalAssert(tls, __ccgo_ts+4911, __ccgo_ts+4530, int32FromInt32(630))
	}
	_ = v10 || bool(int32FromInt32(0) != 0)
	crossMoveCount = staticSeedCount + kinematicSeedCount
	// Collide the dynamic body tree against the static and kinematic trees.
	*(*b2CrossContext)(unsafe.Add(nil, bp1+8192)) = b2CrossContext{
		World:           world,
		Seeds:           bp1,
		StaticSeedCount: staticSeedCount,
	}
	b2ParallelFor(tls, world, __ccgo_fp(b2CrossPairsTask), crossMoveCount, int32(1), bp1+8192)
	// Collide the dynamic body tree against itself.
	b2ParallelFor(tls, world, __ccgo_fp(b2SelfPairsTask), dynamicMoveCount, int32(64), world)
	b2DynamicTree_ClearMoved(tls, bp+uintptr(b2_staticBody)*104)
	// Update stale trees.
	b2EnqueueTreeUpdate(tls, world)
	// Pairs arrive in deterministic order but scrambled relative to body and shape order
	// sorting them here improves solver performance.
	pairCount = 0
	i1 = 0
	for {
		if !(i1 < (*b2World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		pairCount = pairCount + (*(*b2TaskContext)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i1)*136))).PairKeys.Count
		goto _20
	_20:
		;
		i1 = i1 + 1
	}
	v21 = pairCount
	v22 = int32(1)
	if v21 > v22 {
		v25 = v21
	} else {
		v25 = v22
	}
	v23 = v25
	goto _24
_24:
	pairKeys = b2StackAlloc(tls, alloc, int32FromUint64(uint64FromInt32(v23)*uint64(8)), __ccgo_ts+4953)
	keyCount = 0
	i2 = 0
	for {
		if !(i2 < (*b2World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		workerKeys = (*b2World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(i2)*136 + 16
		if (*b2DynamicArray_uint64_t)(unsafe.Pointer(workerKeys)).Count > 0 {
			memcpy(tls, pairKeys+uintptr(keyCount)*8, (*b2DynamicArray_uint64_t)(unsafe.Pointer(workerKeys)).Data, uint64FromInt32((*b2DynamicArray_uint64_t)(unsafe.Pointer(workerKeys)).Count)*uint64(8))
			keyCount = keyCount + (*b2DynamicArray_uint64_t)(unsafe.Pointer(workerKeys)).Count
		}
		goto _26
	_26:
		;
		i2 = i2 + 1
	}
	if v10 = !!(keyCount == pairCount); !v10 {
		b2InternalAssert(tls, __ccgo_ts+4963, __ccgo_ts+4530, int32FromInt32(674))
	}
	_ = v10 || bool(int32FromInt32(0) != 0)
	if pairCount > int32(1) {
		tempKeys = b2StackAlloc(tls, alloc, int32FromUint64(uint64FromInt32(pairCount)*uint64(8)), __ccgo_ts+4985)
		b2RadixSortKeys(tls, pairKeys, tempKeys, pairCount)
		b2StackFree(tls, alloc, tempKeys)
	}
	i3 = 0
	for {
		if !(i3 < keyCount) {
			break
		}
		shapeIdA = int32FromUint64(*(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(i3)*8)) >> int32FromInt32(32))
		shapeIdB = int32FromUint64(*(*uint64_t)(unsafe.Add(nil, pairKeys+uintptr(i3)*8)) & uint64FromUint32(0xFFFFFFFF))
		if v10 = !!(0 <= shapeIdA && shapeIdA < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v10 {
			b2InternalAssert(tls, __ccgo_ts+4559, __ccgo_ts+4530, int32FromInt32(687))
		}
		_ = v10 || bool(int32FromInt32(0) != 0)
		shapeA = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeIdA)*280
		if v16 = !!(0 <= shapeIdB && shapeIdB < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v16 {
			b2InternalAssert(tls, __ccgo_ts+4619, __ccgo_ts+4530, int32FromInt32(688))
		}
		_ = v16 || bool(int32FromInt32(0) != 0)
		shapeB = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeIdB)*280
		b2CreateContact(tls, world, shapeA, shapeB)
		goto _28
	_28:
		;
		i3 = i3 + 1
	}
	b2StackFree(tls, alloc, pairKeys)
	b2StackFree(tls, alloc, (*b2BroadPhase)(unsafe.Pointer(bp)).MovedSiblings)
	(*b2BroadPhase)(unsafe.Pointer(bp)).MovedSiblings = uintptrFromInt32(0)
	b2ValidateSolverSets(tls, world)
}
