package b2

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

// #include <stdio.h>

// static FILE* s_file = NULL;

func b2CreateBroadPhase(tls *_Stack, bp uintptr) {
	var i int32
	_ = i
	// if (s_file == NULL)
	//{
	//	s_file = fopen("pairs01.txt", "a");
	//	fprintf(s_file, "============\n\n");
	// }
	(*b2BroadPhase)(unsafe.Pointer(bp)).MoveSet = b2CreateSet(tls, int32(16))
	(*b2BroadPhase)(unsafe.Pointer(bp)).MoveArray = b2IntArray_Create(tls, int32(16))
	(*b2BroadPhase)(unsafe.Pointer(bp)).MoveResults = uintptrFromInt32(0)
	(*b2BroadPhase)(unsafe.Pointer(bp)).MovePairs = uintptrFromInt32(0)
	(*b2BroadPhase)(unsafe.Pointer(bp)).MovePairCapacity = 0
	atomicStoreNInt32(bp+268, 0, int32FromInt32(__ATOMIC_SEQ_CST))
	(*b2BroadPhase)(unsafe.Pointer(bp)).PairSet = b2CreateSet(tls, int32(32))
	i = 0
	for {
		if !(i < int32(b2_bodyTypeCount)) {
			break
		}
		*(*DynamicTree)(unsafe.Pointer(bp + uintptr(i)*72)) = b2DynamicTree_Create(tls)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b2DestroyBroadPhase(tls *_Stack, bp uintptr) {
	var i int32
	_ = i
	i = 0
	for {
		if !(i < int32(b2_bodyTypeCount)) {
			break
		}
		b2DynamicTree_Destroy(tls, bp+uintptr(i)*72)
		goto _1
	_1:
		;
		i = i + 1
	}
	b2DestroySet(tls, bp+216)
	b2IntArray_Destroy(tls, bp+232)
	b2DestroySet(tls, bp+272)
	memset(tls, bp, 0, uint64(288))
	// if (s_file != NULL)
	//{
	//	fclose(s_file);
	//	s_file = NULL;
	// }
}

func b2BroadPhase_CreateProxy(tls *_Stack, bp1 uintptr, proxyType BodyType, aabb AABB, categoryBits uint64, shapeIndex int32, forcePairCreation uint8) (r int32) {
	var alreadyAdded uint8
	var newCapacity, proxyId, proxyKey, v2, v4 int32
	var v1, v3 uintptr
	_, _, _, _, _, _, _, _ = alreadyAdded, newCapacity, proxyId, proxyKey, v1, v2, v3, v4
	if !(0 <= proxyType && proxyType < int32(b2_bodyTypeCount)) && b2InternalAssertFcn(tls, __ccgo_ts+2423, __ccgo_ts+2470, int32FromInt32(94)) != 0 {
		__builtin_trap(tls)
	}
	proxyId = b2DynamicTree_CreateProxy(tls, bp1+uintptr(proxyType)*72, aabb, categoryBits, uint64FromInt32(shapeIndex))
	proxyKey = proxyId<<int32(2) | proxyType
	if proxyType != int32(b2_staticBody) || forcePairCreation != 0 {
		v1 = bp1
		v2 = proxyKey
		alreadyAdded = b2AddKey(tls, v1+216, uint64FromInt32(v2+int32(1)))
		if int32FromUint8(alreadyAdded) == int32FromInt32(false1) {
			v3 = v1 + 232
			if (*b2IntArray)(unsafe.Pointer(v3)).Count == (*b2IntArray)(unsafe.Pointer(v3)).Capacity {
				if (*b2IntArray)(unsafe.Pointer(v3)).Capacity < int32(2) {
					v4 = int32(2)
				} else {
					v4 = (*b2IntArray)(unsafe.Pointer(v3)).Capacity + (*b2IntArray)(unsafe.Pointer(v3)).Capacity>>int32(1)
				}
				newCapacity = v4
				b2IntArray_Reserve(tls, v3, newCapacity)
			}
			*(*int32)(unsafe.Pointer((*b2IntArray)(unsafe.Pointer(v3)).Data + uintptr((*b2IntArray)(unsafe.Pointer(v3)).Count)*4)) = v2
			*(*int32)(unsafe.Pointer(v3 + 8)) += int32(1)
		}
	}
	return proxyKey
}

func b2BroadPhase_DestroyProxy(tls *_Stack, bp uintptr, proxyKey int32) {
	var proxyId int32
	var proxyType b2BodyType1
	_, _ = proxyId, proxyType
	if !((*b2BroadPhase)(unsafe.Pointer(bp)).MoveArray.Count == int32FromUint32((*b2BroadPhase)(unsafe.Pointer(bp)).MoveSet.Count)) && b2InternalAssertFcn(tls, __ccgo_ts+2499, __ccgo_ts+2470, int32FromInt32(106)) != 0 {
		__builtin_trap(tls)
	}
	b2UnBufferMove(tls, bp, proxyKey)
	proxyType = proxyKey & int32FromInt32(3)
	proxyId = proxyKey >> int32(2)
	if !(0 <= proxyType && proxyType <= int32(b2_bodyTypeCount)) && b2InternalAssertFcn(tls, __ccgo_ts+2545, __ccgo_ts+2470, int32FromInt32(112)) != 0 {
		__builtin_trap(tls)
	}
	b2DynamicTree_DestroyProxy(tls, bp+uintptr(proxyType)*72, proxyId)
}

func b2BroadPhase_MoveProxy(tls *_Stack, bp1 uintptr, proxyKey int32, aabb AABB) {
	var alreadyAdded uint8
	var newCapacity, proxyId, v2, v4 int32
	var proxyType b2BodyType1
	var v1, v3 uintptr
	_, _, _, _, _, _, _, _ = alreadyAdded, newCapacity, proxyId, proxyType, v1, v2, v3, v4
	proxyType = proxyKey & int32FromInt32(3)
	proxyId = proxyKey >> int32(2)
	b2DynamicTree_MoveProxy(tls, bp1+uintptr(proxyType)*72, proxyId, aabb)
	v1 = bp1
	v2 = proxyKey
	alreadyAdded = b2AddKey(tls, v1+216, uint64FromInt32(v2+int32(1)))
	if int32FromUint8(alreadyAdded) == int32FromInt32(false1) {
		v3 = v1 + 232
		if (*b2IntArray)(unsafe.Pointer(v3)).Count == (*b2IntArray)(unsafe.Pointer(v3)).Capacity {
			if (*b2IntArray)(unsafe.Pointer(v3)).Capacity < int32(2) {
				v4 = int32(2)
			} else {
				v4 = (*b2IntArray)(unsafe.Pointer(v3)).Capacity + (*b2IntArray)(unsafe.Pointer(v3)).Capacity>>int32(1)
			}
			newCapacity = v4
			b2IntArray_Reserve(tls, v3, newCapacity)
		}
		*(*int32)(unsafe.Pointer((*b2IntArray)(unsafe.Pointer(v3)).Data + uintptr((*b2IntArray)(unsafe.Pointer(v3)).Count)*4)) = v2
		*(*int32)(unsafe.Pointer(v3 + 8)) += int32(1)
	}
}

func b2BroadPhase_EnlargeProxy(tls *_Stack, bp1 uintptr, proxyKey int32, aabb AABB) {
	var alreadyAdded uint8
	var newCapacity, proxyId, typeIndex, v2, v4 int32
	var v1, v3 uintptr
	_, _, _, _, _, _, _, _ = alreadyAdded, newCapacity, proxyId, typeIndex, v1, v2, v3, v4
	if !(proxyKey != -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+2593, __ccgo_ts+2470, int32FromInt32(127)) != 0 {
		__builtin_trap(tls)
	}
	typeIndex = proxyKey & int32FromInt32(3)
	proxyId = proxyKey >> int32(2)
	if !(typeIndex != int32(b2_staticBody)) && b2InternalAssertFcn(tls, __ccgo_ts+2619, __ccgo_ts+2470, int32FromInt32(131)) != 0 {
		__builtin_trap(tls)
	}
	b2DynamicTree_EnlargeProxy(tls, bp1+uintptr(typeIndex)*72, proxyId, aabb)
	v1 = bp1
	v2 = proxyKey
	alreadyAdded = b2AddKey(tls, v1+216, uint64FromInt32(v2+int32(1)))
	if int32FromUint8(alreadyAdded) == int32FromInt32(false1) {
		v3 = v1 + 232
		if (*b2IntArray)(unsafe.Pointer(v3)).Count == (*b2IntArray)(unsafe.Pointer(v3)).Capacity {
			if (*b2IntArray)(unsafe.Pointer(v3)).Capacity < int32(2) {
				v4 = int32(2)
			} else {
				v4 = (*b2IntArray)(unsafe.Pointer(v3)).Capacity + (*b2IntArray)(unsafe.Pointer(v3)).Capacity>>int32(1)
			}
			newCapacity = v4
			b2IntArray_Reserve(tls, v3, newCapacity)
		}
		*(*int32)(unsafe.Pointer((*b2IntArray)(unsafe.Pointer(v3)).Data + uintptr((*b2IntArray)(unsafe.Pointer(v3)).Count)*4)) = v2
		*(*int32)(unsafe.Pointer(v3 + 8)) += int32(1)
	}
}

func b2UpdateBroadPhasePairs(tls *_Stack, world uintptr) {
	var alloc, bp, pair, result, shapeA, shapeB, temp, userPairTask, v2, v4, v6, v8 uintptr
	var i, minRange, moveCount, shapeIdA, shapeIdB, v3, v7 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = alloc, bp, i, minRange, moveCount, pair, result, shapeA, shapeB, shapeIdA, shapeIdB, temp, userPairTask, v2, v3, v4, v6, v7, v8
	bp = world + 40
	moveCount = (*b2BroadPhase)(unsafe.Pointer(bp)).MoveArray.Count
	if !(moveCount == int32FromUint32((*b2BroadPhase)(unsafe.Pointer(bp)).MoveSet.Count)) && b2InternalAssertFcn(tls, __ccgo_ts+2673, __ccgo_ts+2470, int32FromInt32(381)) != 0 {
		__builtin_trap(tls)
	}
	if moveCount == 0 {
		return
	}
	alloc = world
	// todo these could be in the step context
	(*b2BroadPhase)(unsafe.Pointer(bp)).MoveResults = b2AllocateArenaItem(tls, alloc, int32FromUint64(uint64FromInt32(moveCount)*uint64(8)), __ccgo_ts+2709)
	(*b2BroadPhase)(unsafe.Pointer(bp)).MovePairCapacity = int32(16) * moveCount
	(*b2BroadPhase)(unsafe.Pointer(bp)).MovePairs = b2AllocateArenaItem(tls, alloc, int32FromUint64(uint64FromInt32((*b2BroadPhase)(unsafe.Pointer(bp)).MovePairCapacity)*uint64(24)), __ccgo_ts+2722)
	atomicStoreNInt32(bp+268, 0, int32FromInt32(__ATOMIC_SEQ_CST))
	minRange = int32(64)
	userPairTask = (*(*func(*_Stack, uintptr, int32, int32, uintptr, uintptr) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).EnqueueTaskFcn})))(tls, __ccgo_fp(b2FindPairsTask), moveCount, minRange, world, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
	if userPairTask != uintptrFromInt32(0) {
		(*(*func(*_Stack, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).FinishTaskFcn})))(tls, userPairTask, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
		*(*int32)(unsafe.Pointer(world + 1776)) += int32(1)
	}
	// todo_erin could start tree rebuild here
	// Single-threaded work
	// - Clear move flags
	// - Create contacts in deterministic order
	i = 0
	for {
		if !(i < moveCount) {
			break
		}
		result = (*b2BroadPhase)(unsafe.Pointer(bp)).MoveResults + uintptr(i)*8
		pair = (*b2MoveResult)(unsafe.Pointer(result)).PairList
		for pair != uintptrFromInt32(0) {
			shapeIdA = (*b2MovePair)(unsafe.Pointer(pair)).ShapeIndexA
			shapeIdB = (*b2MovePair)(unsafe.Pointer(pair)).ShapeIndexB
			v2 = world + 1248
			v3 = shapeIdA
			if !(0 <= v3 && v3 < (*b2ShapeArray)(unsafe.Pointer(v2)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
				__builtin_trap(tls)
			}
			v4 = (*b2ShapeArray)(unsafe.Pointer(v2)).Data + uintptr(v3)*288
			goto _5
		_5:
			// if (s_file != NULL)
			//{
			//	fprintf(s_file, "%d %d\n", shapeIdA, shapeIdB);
			// }
			shapeA = v4
			v6 = world + 1248
			v7 = shapeIdB
			if !(0 <= v7 && v7 < (*b2ShapeArray)(unsafe.Pointer(v6)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
				__builtin_trap(tls)
			}
			v8 = (*b2ShapeArray)(unsafe.Pointer(v6)).Data + uintptr(v7)*288
			goto _9
		_9:
			shapeB = v8
			b2CreateContact(tls, world, shapeA, shapeB)
			if (*b2MovePair)(unsafe.Pointer(pair)).Heap != 0 {
				temp = pair
				pair = (*b2MovePair)(unsafe.Pointer(pair)).Next
				b2Free(tls, temp, int32(24))
			} else {
				pair = (*b2MovePair)(unsafe.Pointer(pair)).Next
			}
		}
		// if (s_file != NULL)
		//{
		//	fprintf(s_file, "\n");
		// }
		goto _1
	_1:
		;
		i = i + 1
	}
	// if (s_file != NULL)
	//{
	//	fprintf(s_file, "count = %d\n\n", pairCount);
	// }
	// Reset move buffer
	(*b2IntArray)(unsafe.Pointer(bp + 232)).Count = 0
	b2ClearSet(tls, bp+216)
	b2FreeArenaItem(tls, alloc, (*b2BroadPhase)(unsafe.Pointer(bp)).MovePairs)
	(*b2BroadPhase)(unsafe.Pointer(bp)).MovePairs = uintptrFromInt32(0)
	b2FreeArenaItem(tls, alloc, (*b2BroadPhase)(unsafe.Pointer(bp)).MoveResults)
	(*b2BroadPhase)(unsafe.Pointer(bp)).MoveResults = uintptrFromInt32(0)
	b2ValidateSolverSets(tls, world)
}

func b2BroadPhase_TestOverlap(tls *_Stack, bp uintptr, proxyKeyA int32, proxyKeyB int32) (r uint8) {
	var aabbA, aabbB, v1, v2 AABB
	var proxyIdA, proxyIdB, typeIndexA, typeIndexB int32
	var v3 uint8
	_, _, _, _, _, _, _, _, _ = aabbA, aabbB, proxyIdA, proxyIdB, typeIndexA, typeIndexB, v1, v2, v3
	typeIndexA = proxyKeyA & int32FromInt32(3)
	proxyIdA = proxyKeyA >> int32(2)
	typeIndexB = proxyKeyB & int32FromInt32(3)
	proxyIdB = proxyKeyB >> int32(2)
	aabbA = b2DynamicTree_GetAABB(tls, bp+uintptr(typeIndexA)*72, proxyIdA)
	aabbB = b2DynamicTree_GetAABB(tls, bp+uintptr(typeIndexB)*72, proxyIdB)
	v1 = aabbA
	v2 = aabbB
	v3 = boolUint8(!(v2.LowerBound.X > v1.UpperBound.X || v2.LowerBound.Y > v1.UpperBound.Y || v1.LowerBound.X > v2.UpperBound.X || v1.LowerBound.Y > v2.UpperBound.Y))
	goto _4
_4:
	return v3
}

func b2BroadPhase_RebuildTrees(tls *_Stack, bp uintptr) {
	b2DynamicTree_Rebuild(tls, bp+uintptr(b2_dynamicBody)*72, boolUint8(false1 != 0))
	b2DynamicTree_Rebuild(tls, bp+uintptr(b2_kinematicBody)*72, boolUint8(false1 != 0))
}

func b2BroadPhase_GetShapeIndex(tls *_Stack, bp uintptr, proxyKey int32) (r int32) {
	var proxyId, typeIndex int32
	_, _ = proxyId, typeIndex
	typeIndex = proxyKey & int32FromInt32(3)
	proxyId = proxyKey >> int32(2)
	return int32FromUint64(b2DynamicTree_GetUserData(tls, bp+uintptr(typeIndex)*72, proxyId))
}
