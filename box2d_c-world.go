package b2

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b2IsValidWorldTransform(tls *_Stack, t WorldTransform) (r uint8) {
	if int32FromUint8(b2IsValidPos(tls, t.P)) == false1 {
		return boolUint8(false1 != 0)
	}
	return b2IsValidRot(tls, t.Q)
}

func b2GetUnlockedWorldFromId(tls *_Stack, id WorldId) (r uintptr) {
	var world uintptr
	var v1 bool
	_, _ = world, v1
	if v1 = !!(int32(1) <= int32FromUint16(id.Index1) && int32FromUint16(id.Index1) <= int32(_B2_MAX_WORLDS)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+16615, __ccgo_ts+16660, int32FromInt32(44))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = uintptr(unsafe.Pointer(&b2_worlds)) + uintptr(int32FromUint16(id.Index1)-int32FromInt32(1))*2792
	if v1 = !!(int32FromUint16(id.Index1) == int32FromUint16((*b2World)(unsafe.Pointer(world)).WorldId)+int32FromInt32(1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+16691, __ccgo_ts+16660, int32FromInt32(46))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(int32FromUint16(id.Generation) == int32FromUint16((*b2World)(unsafe.Pointer(world)).Generation)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+16723, __ccgo_ts+16660, int32FromInt32(47))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// A world accessed from an id should not be locked
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+7715, __ccgo_ts+16660, int32FromInt32(52))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return uintptrFromInt32(0)
	}
	return world
}

func b2GetWorldFromId(tls *_Stack, id WorldId) (r uintptr) {
	var world uintptr
	var v1 bool
	_, _ = world, v1
	if v1 = !!(int32(1) <= int32FromUint16(id.Index1) && int32FromUint16(id.Index1) <= int32(_B2_MAX_WORLDS)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+16615, __ccgo_ts+16660, int32FromInt32(60))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = uintptr(unsafe.Pointer(&b2_worlds)) + uintptr(int32FromUint16(id.Index1)-int32FromInt32(1))*2792
	if v1 = !!(int32FromUint16(id.Index1) == int32FromUint16((*b2World)(unsafe.Pointer(world)).WorldId)+int32FromInt32(1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+16691, __ccgo_ts+16660, int32FromInt32(62))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(int32FromUint16(id.Generation) == int32FromUint16((*b2World)(unsafe.Pointer(world)).Generation)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+16723, __ccgo_ts+16660, int32FromInt32(63))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return world
}

func b2GetWorld(tls *_Stack, index int32) (r uintptr) {
	var world uintptr
	var v1 bool
	_, _ = world, v1
	if v1 = !!(0 <= index && index < int32(_B2_MAX_WORLDS)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+16758, __ccgo_ts+16660, int32FromInt32(69))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = uintptr(unsafe.Pointer(&b2_worlds)) + uintptr(index)*2792
	if v1 = !!(int32FromUint16((*b2World)(unsafe.Pointer(world)).WorldId) == index); !v1 {
		b2InternalAssert(tls, __ccgo_ts+16794, __ccgo_ts+16660, int32FromInt32(71))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return world
}

func b2GetWorldLocked(tls *_Stack, index int32) (r uintptr) {
	var world uintptr
	var v1 bool
	_, _ = world, v1
	if v1 = !!(0 <= index && index < int32(_B2_MAX_WORLDS)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+16758, __ccgo_ts+16660, int32FromInt32(77))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = uintptr(unsafe.Pointer(&b2_worlds)) + uintptr(index)*2792
	if v1 = !!(int32FromUint16((*b2World)(unsafe.Pointer(world)).WorldId) == index); !v1 {
		b2InternalAssert(tls, __ccgo_ts+16794, __ccgo_ts+16660, int32FromInt32(79))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+7715, __ccgo_ts+16660, int32FromInt32(82))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return uintptrFromInt32(0)
	}
	return world
}

func b2CreateWorld(tls *_Stack, def uintptr) (r WorldId) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var bodyCapacity, elementSize, elementSize1, elementSize2, i, newCapacity, newCapacity1, newCapacity2, newSize, newSize1, newSize2, newSize3, newSize4, newSize5, newSize6, oldSize, oldSize1, oldSize2, oldSize3, oldSize4, oldSize5, oldSize6, shapeCapacity, worldId, v13, v14, v15, v17, v18, v19, v20, v22 int32
	var generation uint16_t
	var set b2SolverSet
	var world, v23 uintptr
	var v1 bool
	var v10 uint8
	var v129, v130, v131, v132, v134, v135 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyCapacity, elementSize, elementSize1, elementSize2, generation, i, newCapacity, newCapacity1, newCapacity2, newSize, newSize1, newSize2, newSize3, newSize4, newSize5, newSize6, oldSize, oldSize1, oldSize2, oldSize3, oldSize4, oldSize5, oldSize6, set, shapeCapacity, world, worldId, v1, v10, v129, v13, v130, v131, v132, v134, v135, v14, v15, v17, v18, v19, v20, v22, v23
	if v1 = !!((*WorldDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1077, __ccgo_ts+16660, int32FromInt32(147))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if int32FromUint8(b2IsValidVec2(tls, (*WorldDef)(unsafe.Pointer(def)).Gravity)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+8, __ccgo_ts+16818, uintptr(unsafe.Pointer(&__func__51))))
		if v1 = !!(b2IsValidVec2(tls, (*WorldDef)(unsafe.Pointer(def)).Gravity) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+16818, __ccgo_ts+16660, int32FromInt32(148))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return WorldId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*WorldDef)(unsafe.Pointer(def)).RestitutionThreshold)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+8, __ccgo_ts+16848, uintptr(unsafe.Pointer(&__func__51))))
		if v1 = !!(b2IsValidFloat(tls, (*WorldDef)(unsafe.Pointer(def)).RestitutionThreshold) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+16848, __ccgo_ts+16660, int32FromInt32(149))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return WorldId{}
	}
	if boolInt32((*WorldDef)(unsafe.Pointer(def)).RestitutionIterations >= 0) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+8, __ccgo_ts+16892, uintptr(unsafe.Pointer(&__func__51))))
		if v1 = !!((*WorldDef)(unsafe.Pointer(def)).RestitutionIterations >= int32FromInt32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+16892, __ccgo_ts+16660, int32FromInt32(150))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return WorldId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*WorldDef)(unsafe.Pointer(def)).HitEventThreshold)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+8, __ccgo_ts+16924, uintptr(unsafe.Pointer(&__func__51))))
		if v1 = !!(b2IsValidFloat(tls, (*WorldDef)(unsafe.Pointer(def)).HitEventThreshold) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+16924, __ccgo_ts+16660, int32FromInt32(151))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return WorldId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*WorldDef)(unsafe.Pointer(def)).ContactHertz)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+8, __ccgo_ts+16965, uintptr(unsafe.Pointer(&__func__51))))
		if v1 = !!(b2IsValidFloat(tls, (*WorldDef)(unsafe.Pointer(def)).ContactHertz) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+16965, __ccgo_ts+16660, int32FromInt32(152))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return WorldId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*WorldDef)(unsafe.Pointer(def)).ContactDampingRatio)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+8, __ccgo_ts+17001, uintptr(unsafe.Pointer(&__func__51))))
		if v1 = !!(b2IsValidFloat(tls, (*WorldDef)(unsafe.Pointer(def)).ContactDampingRatio) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+17001, __ccgo_ts+16660, int32FromInt32(153))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return WorldId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*WorldDef)(unsafe.Pointer(def)).ContactSpeed)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+8, __ccgo_ts+17044, uintptr(unsafe.Pointer(&__func__51))))
		if v1 = !!(b2IsValidFloat(tls, (*WorldDef)(unsafe.Pointer(def)).ContactSpeed) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+17044, __ccgo_ts+16660, int32FromInt32(154))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return WorldId{}
	}
	if boolInt32(b2IsValidFloat(tls, (*WorldDef)(unsafe.Pointer(def)).MaximumLinearSpeed) != 0 && (*WorldDef)(unsafe.Pointer(def)).MaximumLinearSpeed > float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+8, __ccgo_ts+17080, uintptr(unsafe.Pointer(&__func__51))))
		if v1 = !!(b2IsValidFloat(tls, (*WorldDef)(unsafe.Pointer(def)).MaximumLinearSpeed) != 0 && (*WorldDef)(unsafe.Pointer(def)).MaximumLinearSpeed > float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+17080, __ccgo_ts+16660, int32FromInt32(155))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return WorldId{}
	}
	v10 = boolUint8(false1 != 0)
	goto _11
_11:
	if v10 != 0 {
		b2Log(tls, __ccgo_ts+17156, 0)
	}
	worldId = -int32(1)
	i = 0
	for {
		if !(i < int32(_B2_MAX_WORLDS)) {
			break
		}
		if int32FromUint8(b2_worlds[i].InUse) == false1 {
			worldId = i
			break
		}
		goto _12
	_12:
		;
		i = i + 1
	}
	if worldId == -int32(1) {
		return WorldId{}
	}
	world = uintptr(unsafe.Pointer(&b2_worlds)) + uintptr(worldId)*2792
	generation = (*b2World)(unsafe.Pointer(world)).Generation
	*(*b2World)(unsafe.Add(nil, world)) = b2World{}
	(*b2World)(unsafe.Pointer(world)).WorldId = uint16FromInt32(worldId)
	(*b2World)(unsafe.Pointer(world)).Generation = generation
	(*b2World)(unsafe.Pointer(world)).InUse = boolUint8(true1 != 0)
	v13 = int32(2048)
	v14 = (*WorldDef)(unsafe.Pointer(def)).Capacity.ArenaByteCount
	if v13 > v14 {
		v17 = v13
	} else {
		v17 = v14
	}
	v15 = v17
	goto _16
_16:
	(*b2World)(unsafe.Pointer(world)).Stack = b2CreateStack(tls, v15)
	b2CreateBroadPhase(tls, world+40, def+96)
	b2CreateGraph(tls, world+376, def+96)
	// pools
	(*b2World)(unsafe.Pointer(world)).BodyIdPool = b2CreateIdPool(tls)
	v13 = int32(16)
	v14 = (*WorldDef)(unsafe.Pointer(def)).Capacity.StaticBodyCount + (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicBodyCount
	if v13 > v14 {
		v17 = v13
	} else {
		v17 = v14
	}
	v15 = v17
	goto _21
_21:
	bodyCapacity = v15
	if bodyCapacity > 0 {
		v23 = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(bodyCapacity)*uint64(104))
	} else {
		v23 = uintptrFromInt32(0)
	}
	(*b2World)(unsafe.Pointer(world)).Bodies.Data = v23
	(*b2World)(unsafe.Pointer(world)).Bodies.Count = 0
	(*b2World)(unsafe.Pointer(world)).Bodies.Capacity = bodyCapacity
	(*b2World)(unsafe.Pointer(world)).SolverSets.Data = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(int32FromInt32(8))*uint64FromInt64(88))
	(*b2World)(unsafe.Pointer(world)).SolverSets.Count = 0
	(*b2World)(unsafe.Pointer(world)).SolverSets.Capacity = int32(8)
	// add empty static, active, and disabled body sets
	(*b2World)(unsafe.Pointer(world)).SolverSetIdPool = b2CreateIdPool(tls)
	set = b2SolverSet{}
	// static set
	set.SetIndex = b2AllocId(tls, world+1952)
	elementSize = int32(88)
	if (*b2World)(unsafe.Pointer(world)).SolverSets.Count >= (*b2World)(unsafe.Pointer(world)).SolverSets.Capacity {
		oldSize = (*b2World)(unsafe.Pointer(world)).SolverSets.Capacity * elementSize
		if (*b2World)(unsafe.Pointer(world)).SolverSets.Capacity == 0 {
			v13 = int32(8)
		} else {
			v13 = int32(2) * (*b2World)(unsafe.Pointer(world)).SolverSets.Capacity
		}
		newCapacity = v13
		newSize = newCapacity * elementSize
		(*b2World)(unsafe.Pointer(world)).SolverSets.Data = b2GrowAlloc(tls, (*b2World)(unsafe.Pointer(world)).SolverSets.Data, uint64FromInt32(oldSize), uint64FromInt32(newSize))
		(*b2World)(unsafe.Pointer(world)).SolverSets.Capacity = newCapacity
	}
	*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr((*b2World)(unsafe.Pointer(world)).SolverSets.Count)*88)) = set
	(*b2World)(unsafe.Pointer(world)).SolverSets.Count += int32(1)
	v13 = int32(16)
	v14 = (*WorldDef)(unsafe.Pointer(def)).Capacity.StaticBodyCount
	if v13 > v14 {
		v17 = v13
	} else {
		v17 = v14
	}
	v15 = v17
	goto _28
_28:
	if (*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_staticSet)*88))).BodySims.Capacity < v15 {
		oldSize1 = int32FromUint64(uint64FromInt32((*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_staticSet)*88))).BodySims.Capacity) * uint64(96))
		v18 = int32(16)
		v19 = (*WorldDef)(unsafe.Pointer(def)).Capacity.StaticBodyCount
		if v18 > v19 {
			v22 = v18
		} else {
			v22 = v19
		}
		v20 = v22
		goto _33
	_33:
		newSize1 = int32FromUint64(uint64FromInt32(v20) * uint64(96))
		(*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_staticSet)*88))).BodySims.Data = b2GrowAlloc(tls, (*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_staticSet)*88))).BodySims.Data, uint64FromInt32(oldSize1), uint64FromInt32(newSize1))
		v13 = int32(16)
		v14 = (*WorldDef)(unsafe.Pointer(def)).Capacity.StaticBodyCount
		if v13 > v14 {
			v17 = v13
		} else {
			v17 = v14
		}
		v15 = v17
		goto _38
	_38:
		(*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_staticSet)*88))).BodySims.Capacity = v15
	}
	if v1 = !!((*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_staticSet)*88))).SetIndex == int32(b2_staticSet)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+17212, __ccgo_ts+16660, int32FromInt32(205))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// disabled set
	set.SetIndex = b2AllocId(tls, world+1952)
	elementSize1 = int32(88)
	if (*b2World)(unsafe.Pointer(world)).SolverSets.Count >= (*b2World)(unsafe.Pointer(world)).SolverSets.Capacity {
		oldSize2 = (*b2World)(unsafe.Pointer(world)).SolverSets.Capacity * elementSize1
		if (*b2World)(unsafe.Pointer(world)).SolverSets.Capacity == 0 {
			v13 = int32(8)
		} else {
			v13 = int32(2) * (*b2World)(unsafe.Pointer(world)).SolverSets.Capacity
		}
		newCapacity1 = v13
		newSize2 = newCapacity1 * elementSize1
		(*b2World)(unsafe.Pointer(world)).SolverSets.Data = b2GrowAlloc(tls, (*b2World)(unsafe.Pointer(world)).SolverSets.Data, uint64FromInt32(oldSize2), uint64FromInt32(newSize2))
		(*b2World)(unsafe.Pointer(world)).SolverSets.Capacity = newCapacity1
	}
	*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr((*b2World)(unsafe.Pointer(world)).SolverSets.Count)*88)) = set
	(*b2World)(unsafe.Pointer(world)).SolverSets.Count += int32(1)
	if v1 = !!((*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_disabledSet)*88))).SetIndex == int32(b2_disabledSet)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+17274, __ccgo_ts+16660, int32FromInt32(210))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// awake set
	set.SetIndex = b2AllocId(tls, world+1952)
	elementSize2 = int32(88)
	if (*b2World)(unsafe.Pointer(world)).SolverSets.Count >= (*b2World)(unsafe.Pointer(world)).SolverSets.Capacity {
		oldSize3 = (*b2World)(unsafe.Pointer(world)).SolverSets.Capacity * elementSize2
		if (*b2World)(unsafe.Pointer(world)).SolverSets.Capacity == 0 {
			v13 = int32(8)
		} else {
			v13 = int32(2) * (*b2World)(unsafe.Pointer(world)).SolverSets.Capacity
		}
		newCapacity2 = v13
		newSize3 = newCapacity2 * elementSize2
		(*b2World)(unsafe.Pointer(world)).SolverSets.Data = b2GrowAlloc(tls, (*b2World)(unsafe.Pointer(world)).SolverSets.Data, uint64FromInt32(oldSize3), uint64FromInt32(newSize3))
		(*b2World)(unsafe.Pointer(world)).SolverSets.Capacity = newCapacity2
	}
	*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr((*b2World)(unsafe.Pointer(world)).SolverSets.Count)*88)) = set
	(*b2World)(unsafe.Pointer(world)).SolverSets.Count += int32(1)
	v13 = int32(16)
	v14 = (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicBodyCount
	if v13 > v14 {
		v17 = v13
	} else {
		v17 = v14
	}
	v15 = v17
	goto _47
_47:
	if (*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_awakeSet)*88))).BodySims.Capacity < v15 {
		oldSize4 = int32FromUint64(uint64FromInt32((*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_awakeSet)*88))).BodySims.Capacity) * uint64(96))
		v18 = int32(16)
		v19 = (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicBodyCount
		if v18 > v19 {
			v22 = v18
		} else {
			v22 = v19
		}
		v20 = v22
		goto _52
	_52:
		newSize4 = int32FromUint64(uint64FromInt32(v20) * uint64(96))
		(*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_awakeSet)*88))).BodySims.Data = b2GrowAlloc(tls, (*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_awakeSet)*88))).BodySims.Data, uint64FromInt32(oldSize4), uint64FromInt32(newSize4))
		v13 = int32(16)
		v14 = (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicBodyCount
		if v13 > v14 {
			v17 = v13
		} else {
			v17 = v14
		}
		v15 = v17
		goto _57
	_57:
		(*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_awakeSet)*88))).BodySims.Capacity = v15
	}
	v13 = int32(16)
	v14 = (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicBodyCount
	if v13 > v14 {
		v17 = v13
	} else {
		v17 = v14
	}
	v15 = v17
	goto _62
_62:
	if (*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_awakeSet)*88))).BodyStates.Capacity < v15 {
		oldSize5 = int32FromUint64(uint64FromInt32((*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_awakeSet)*88))).BodyStates.Capacity) * uint64(32))
		v18 = int32(16)
		v19 = (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicBodyCount
		if v18 > v19 {
			v22 = v18
		} else {
			v22 = v19
		}
		v20 = v22
		goto _67
	_67:
		newSize5 = int32FromUint64(uint64FromInt32(v20) * uint64(32))
		(*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_awakeSet)*88))).BodyStates.Data = b2GrowAlloc(tls, (*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_awakeSet)*88))).BodyStates.Data, uint64FromInt32(oldSize5), uint64FromInt32(newSize5))
		v13 = int32(16)
		v14 = (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicBodyCount
		if v13 > v14 {
			v17 = v13
		} else {
			v17 = v14
		}
		v15 = v17
		goto _72
	_72:
		(*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_awakeSet)*88))).BodyStates.Capacity = v15
	}
	v13 = int32(16)
	v14 = (*WorldDef)(unsafe.Pointer(def)).Capacity.ContactCount
	if v13 > v14 {
		v17 = v13
	} else {
		v17 = v14
	}
	v15 = v17
	goto _77
_77:
	if (*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_awakeSet)*88))).ContactSims.Capacity < v15 {
		oldSize6 = int32FromUint64(uint64FromInt32((*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_awakeSet)*88))).ContactSims.Capacity) * uint64(200))
		v18 = int32(16)
		v19 = (*WorldDef)(unsafe.Pointer(def)).Capacity.ContactCount
		if v18 > v19 {
			v22 = v18
		} else {
			v22 = v19
		}
		v20 = v22
		goto _82
	_82:
		newSize6 = int32FromUint64(uint64FromInt32(v20) * uint64(200))
		(*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_awakeSet)*88))).ContactSims.Data = b2GrowAlloc(tls, (*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_awakeSet)*88))).ContactSims.Data, uint64FromInt32(oldSize6), uint64FromInt32(newSize6))
		v13 = int32(16)
		v14 = (*WorldDef)(unsafe.Pointer(def)).Capacity.ContactCount
		if v13 > v14 {
			v17 = v13
		} else {
			v17 = v14
		}
		v15 = v17
		goto _87
	_87:
		(*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_awakeSet)*88))).ContactSims.Capacity = v15
	}
	if v1 = !!((*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_awakeSet)*88))).SetIndex == int32(b2_awakeSet)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+17340, __ccgo_ts+16660, int32FromInt32(218))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	(*b2World)(unsafe.Pointer(world)).ShapeIdPool = b2CreateIdPool(tls)
	v13 = int32(16)
	v14 = (*WorldDef)(unsafe.Pointer(def)).Capacity.StaticShapeCount + (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicShapeCount
	if v13 > v14 {
		v17 = v13
	} else {
		v17 = v14
	}
	v15 = v17
	goto _93
_93:
	shapeCapacity = v15
	if shapeCapacity > 0 {
		v23 = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(shapeCapacity)*uint64(280))
	} else {
		v23 = uintptrFromInt32(0)
	}
	(*b2World)(unsafe.Pointer(world)).Shapes.Data = v23
	(*b2World)(unsafe.Pointer(world)).Shapes.Count = 0
	(*b2World)(unsafe.Pointer(world)).Shapes.Capacity = shapeCapacity
	if shapeCapacity > 0 {
		v23 = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(shapeCapacity)*uint64(16))
	} else {
		v23 = uintptrFromInt32(0)
	}
	(*b2World)(unsafe.Pointer(world)).FatAABBs.Data = v23
	(*b2World)(unsafe.Pointer(world)).FatAABBs.Count = 0
	(*b2World)(unsafe.Pointer(world)).FatAABBs.Capacity = shapeCapacity
	(*b2World)(unsafe.Pointer(world)).ChainIdPool = b2CreateIdPool(tls)
	(*b2World)(unsafe.Pointer(world)).ChainShapes.Data = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(int32FromInt32(4))*uint64FromInt64(32))
	(*b2World)(unsafe.Pointer(world)).ChainShapes.Count = 0
	(*b2World)(unsafe.Pointer(world)).ChainShapes.Capacity = int32(4)
	(*b2World)(unsafe.Pointer(world)).ContactIdPool = b2CreateIdPool(tls)
	v13 = int32(16)
	v14 = (*WorldDef)(unsafe.Pointer(def)).Capacity.ContactCount
	if v13 > v14 {
		v17 = v13
	} else {
		v17 = v14
	}
	v15 = v17
	goto _101
_101:
	if v15 > 0 {
		v18 = int32(16)
		v19 = (*WorldDef)(unsafe.Pointer(def)).Capacity.ContactCount
		if v18 > v19 {
			v22 = v18
		} else {
			v22 = v19
		}
		v20 = v22
		goto _106
	_106:
		v23 = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(v20)*uint64(64))
	} else {
		v23 = uintptrFromInt32(0)
	}
	(*b2World)(unsafe.Pointer(world)).Contacts.Data = v23
	(*b2World)(unsafe.Pointer(world)).Contacts.Count = 0
	v13 = int32(16)
	v14 = (*WorldDef)(unsafe.Pointer(def)).Capacity.ContactCount
	if v13 > v14 {
		v17 = v13
	} else {
		v17 = v14
	}
	v15 = v17
	goto _111
_111:
	(*b2World)(unsafe.Pointer(world)).Contacts.Capacity = v15
	(*b2World)(unsafe.Pointer(world)).JointIdPool = b2CreateIdPool(tls)
	(*b2World)(unsafe.Pointer(world)).Joints.Data = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(int32FromInt32(16))*uint64FromInt64(72))
	(*b2World)(unsafe.Pointer(world)).Joints.Count = 0
	(*b2World)(unsafe.Pointer(world)).Joints.Capacity = int32(16)
	(*b2World)(unsafe.Pointer(world)).IslandIdPool = b2CreateIdPool(tls)
	v13 = int32(16)
	v14 = (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicBodyCount
	if v13 > v14 {
		v17 = v13
	} else {
		v17 = v14
	}
	v15 = v17
	goto _117
_117:
	if v15 > 0 {
		v18 = int32(16)
		v19 = (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicBodyCount
		if v18 > v19 {
			v22 = v18
		} else {
			v22 = v19
		}
		v20 = v22
		goto _122
	_122:
		v23 = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(v20)*uint64(64))
	} else {
		v23 = uintptrFromInt32(0)
	}
	(*b2World)(unsafe.Pointer(world)).Islands.Data = v23
	(*b2World)(unsafe.Pointer(world)).Islands.Count = 0
	v13 = int32(16)
	v14 = (*WorldDef)(unsafe.Pointer(def)).Capacity.DynamicBodyCount
	if v13 > v14 {
		v17 = v13
	} else {
		v17 = v14
	}
	v15 = v17
	goto _127
_127:
	(*b2World)(unsafe.Pointer(world)).Islands.Capacity = v15
	(*b2World)(unsafe.Pointer(world)).Sensors.Data = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(int32FromInt32(4))*uint64FromInt64(56))
	(*b2World)(unsafe.Pointer(world)).Sensors.Count = 0
	(*b2World)(unsafe.Pointer(world)).Sensors.Capacity = int32(4)
	(*b2World)(unsafe.Pointer(world)).BodyMoveEvents.Data = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(int32FromInt32(4))*uint64FromInt64(40))
	(*b2World)(unsafe.Pointer(world)).BodyMoveEvents.Count = 0
	(*b2World)(unsafe.Pointer(world)).BodyMoveEvents.Capacity = int32(4)
	(*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Data = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(int32FromInt32(4))*uint64FromInt64(16))
	(*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Count = 0
	(*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity = int32(4)
	(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304))).Data = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(int32FromInt32(4))*uint64FromInt64(16))
	(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304))).Count = 0
	(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304))).Capacity = int32(4)
	(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+1*16))).Data = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(int32FromInt32(4))*uint64FromInt64(16))
	(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+1*16))).Count = 0
	(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+1*16))).Capacity = int32(4)
	(*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Data = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(int32FromInt32(4))*uint64FromInt64(28))
	(*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Count = 0
	(*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Capacity = int32(4)
	(*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336))).Data = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(int32FromInt32(4))*uint64FromInt64(28))
	(*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336))).Count = 0
	(*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336))).Capacity = int32(4)
	(*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+1*16))).Data = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(int32FromInt32(4))*uint64FromInt64(28))
	(*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+1*16))).Count = 0
	(*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+1*16))).Capacity = int32(4)
	(*b2World)(unsafe.Pointer(world)).ContactHitEvents.Data = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(int32FromInt32(4))*uint64FromInt64(48))
	(*b2World)(unsafe.Pointer(world)).ContactHitEvents.Count = 0
	(*b2World)(unsafe.Pointer(world)).ContactHitEvents.Capacity = int32(4)
	(*b2World)(unsafe.Pointer(world)).JointEvents.Data = b2GrowAlloc(tls, uintptrFromInt32(0), uint64(0), uint64FromInt32(int32FromInt32(4))*uint64FromInt64(16))
	(*b2World)(unsafe.Pointer(world)).JointEvents.Count = 0
	(*b2World)(unsafe.Pointer(world)).JointEvents.Capacity = int32(4)
	(*b2World)(unsafe.Pointer(world)).EndEventArrayIndex = 0
	(*b2World)(unsafe.Pointer(world)).StepIndex = uint64(0)
	(*b2World)(unsafe.Pointer(world)).SplitIslandId = -int32(1)
	(*b2World)(unsafe.Pointer(world)).ActiveTaskCount = 0
	(*b2World)(unsafe.Pointer(world)).TaskCount = 0
	(*b2World)(unsafe.Pointer(world)).Gravity = (*WorldDef)(unsafe.Pointer(def)).Gravity
	v129 = (*WorldDef)(unsafe.Pointer(def)).HitEventThreshold
	v130 = float32FromFloat32(0)
	v131 = float32FromFloat32(3.4028234663852886e+38)
	if v129 < v130 {
		v134 = v130
	} else {
		if v129 > v131 {
			v135 = v131
		} else {
			v135 = v129
		}
		v134 = v135
	}
	v132 = v134
	goto _133
_133:
	(*b2World)(unsafe.Pointer(world)).HitEventThreshold = v132
	v129 = (*WorldDef)(unsafe.Pointer(def)).RestitutionThreshold
	v130 = float32FromFloat32(0)
	v131 = float32FromFloat32(3.4028234663852886e+38)
	if v129 < v130 {
		v134 = v130
	} else {
		if v129 > v131 {
			v135 = v131
		} else {
			v135 = v129
		}
		v134 = v135
	}
	v132 = v134
	goto _140
_140:
	(*b2World)(unsafe.Pointer(world)).RestitutionThreshold = v132
	// Clamp this to avoid overflowing the solver sync flags.
	v13 = (*WorldDef)(unsafe.Pointer(def)).RestitutionIterations
	v14 = 0
	v15 = int32(_B2_MAX_RESTITUTION_ITERATIONS)
	if v13 < v14 {
		v18 = v14
	} else {
		if v13 > v15 {
			v19 = v15
		} else {
			v19 = v13
		}
		v18 = v19
	}
	v17 = v18
	goto _147
_147:
	(*b2World)(unsafe.Pointer(world)).RestitutionIterations = v17
	(*b2World)(unsafe.Pointer(world)).EnableRestitutionPropagation = (*WorldDef)(unsafe.Pointer(def)).EnableRestitutionPropagation
	(*b2World)(unsafe.Pointer(world)).MaxLinearSpeed = (*WorldDef)(unsafe.Pointer(def)).MaximumLinearSpeed
	v129 = (*WorldDef)(unsafe.Pointer(def)).ContactSpeed
	v130 = float32FromFloat32(0)
	v131 = float32FromFloat32(3.4028234663852886e+38)
	if v129 < v130 {
		v134 = v130
	} else {
		if v129 > v131 {
			v135 = v131
		} else {
			v135 = v129
		}
		v134 = v135
	}
	v132 = v134
	goto _154
_154:
	(*b2World)(unsafe.Pointer(world)).ContactSpeed = v132
	v129 = (*WorldDef)(unsafe.Pointer(def)).ContactHertz
	v130 = float32FromFloat32(0)
	v131 = float32FromFloat32(3.4028234663852886e+38)
	if v129 < v130 {
		v134 = v130
	} else {
		if v129 > v131 {
			v135 = v131
		} else {
			v135 = v129
		}
		v134 = v135
	}
	v132 = v134
	goto _161
_161:
	(*b2World)(unsafe.Pointer(world)).ContactHertz = v132
	v129 = (*WorldDef)(unsafe.Pointer(def)).ContactDampingRatio
	v130 = float32FromFloat32(0)
	v131 = float32FromFloat32(3.4028234663852886e+38)
	if v129 < v130 {
		v134 = v130
	} else {
		if v129 > v131 {
			v135 = v131
		} else {
			v135 = v129
		}
		v134 = v135
	}
	v132 = v134
	goto _168
_168:
	(*b2World)(unsafe.Pointer(world)).ContactDampingRatio = v132
	(*b2World)(unsafe.Pointer(world)).ContactRecycleDistance = float32(float32FromFloat32(10) * float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))
	if (*WorldDef)(unsafe.Pointer(def)).FrictionCallback != uintptrFromInt32(0) {
		(*b2World)(unsafe.Pointer(world)).FrictionCallback = (*WorldDef)(unsafe.Pointer(def)).FrictionCallback
	}
	if (*WorldDef)(unsafe.Pointer(def)).RestitutionCallback != uintptrFromInt32(0) {
		(*b2World)(unsafe.Pointer(world)).RestitutionCallback = (*WorldDef)(unsafe.Pointer(def)).RestitutionCallback
	}
	(*b2World)(unsafe.Pointer(world)).EnableSleep = (*WorldDef)(unsafe.Pointer(def)).EnableSleep
	(*b2World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
	(*b2World)(unsafe.Pointer(world)).EnableWarmStarting = boolUint8(true1 != 0)
	(*b2World)(unsafe.Pointer(world)).EnableContinuous = (*WorldDef)(unsafe.Pointer(def)).EnableContinuous
	(*b2World)(unsafe.Pointer(world)).SimdWidth = b2GetSIMDWidth(tls)
	(*b2World)(unsafe.Pointer(world)).UserTreeTask = uintptrFromInt32(0)
	(*b2World)(unsafe.Pointer(world)).UserData = (*WorldDef)(unsafe.Pointer(def)).UserData
	if (*WorldDef)(unsafe.Pointer(def)).WorkerCount > 0 && (*WorldDef)(unsafe.Pointer(def)).EnqueueTask != uintptrFromInt32(0) && (*WorldDef)(unsafe.Pointer(def)).FinishTask != uintptrFromInt32(0) {
		// External task system
		v13 = (*WorldDef)(unsafe.Pointer(def)).WorkerCount
		v14 = int32(_B2_MAX_WORKERS)
		if v13 < v14 {
			v17 = v13
		} else {
			v17 = v14
		}
		v15 = v17
		goto _174
	_174:
		(*b2World)(unsafe.Pointer(world)).WorkerCount = v15
		(*b2World)(unsafe.Pointer(world)).EnqueueTaskFcn = (*WorldDef)(unsafe.Pointer(def)).EnqueueTask
		(*b2World)(unsafe.Pointer(world)).FinishTaskFcn = (*WorldDef)(unsafe.Pointer(def)).FinishTask
		(*b2World)(unsafe.Pointer(world)).UserTaskContext = (*WorldDef)(unsafe.Pointer(def)).UserTaskContext
		(*b2World)(unsafe.Pointer(world)).Scheduler = uintptrFromInt32(0)
	} else {
		if (*WorldDef)(unsafe.Pointer(def)).WorkerCount > int32(1) {
			// Built-in scheduler
			v13 = (*WorldDef)(unsafe.Pointer(def)).WorkerCount
			v14 = int32(_B2_MAX_WORKERS)
			if v13 < v14 {
				v17 = v13
			} else {
				v17 = v14
			}
			v15 = v17
			goto _179
		_179:
			(*b2World)(unsafe.Pointer(world)).WorkerCount = v15
			(*b2World)(unsafe.Pointer(world)).Scheduler = b2CreateScheduler(tls, (*b2World)(unsafe.Pointer(world)).WorkerCount)
			(*b2World)(unsafe.Pointer(world)).EnqueueTaskFcn = __ccgo_fp(b2SchedulerEnqueueTask)
			(*b2World)(unsafe.Pointer(world)).FinishTaskFcn = __ccgo_fp(b2SchedulerFinishTask)
			(*b2World)(unsafe.Pointer(world)).UserTaskContext = (*b2World)(unsafe.Pointer(world)).Scheduler
		} else {
			// Serial fallback
			(*b2World)(unsafe.Pointer(world)).WorkerCount = int32(1)
			(*b2World)(unsafe.Pointer(world)).EnqueueTaskFcn = __ccgo_fp(b2DefaultAddTaskFcn)
			(*b2World)(unsafe.Pointer(world)).FinishTaskFcn = __ccgo_fp(b2DefaultFinishTaskFcn)
			(*b2World)(unsafe.Pointer(world)).UserTaskContext = uintptrFromInt32(0)
			(*b2World)(unsafe.Pointer(world)).Scheduler = uintptrFromInt32(0)
		}
	}
	b2CreateWorkerContexts(tls, world)
	(*b2World)(unsafe.Pointer(world)).DebugBodySet = b2CreateBitSet(tls, uint32(256))
	(*b2World)(unsafe.Pointer(world)).DebugJointSet = b2CreateBitSet(tls, uint32(256))
	(*b2World)(unsafe.Pointer(world)).DebugContactSet = b2CreateBitSet(tls, uint32(256))
	(*b2World)(unsafe.Pointer(world)).DebugIslandSet = b2CreateBitSet(tls, uint32(256))
	// Recording is started by the host with b2World_StartRecording, never from the world def
	(*b2World)(unsafe.Pointer(world)).Recording = uintptrFromInt32(0)
	// add one to worldId so that 0 represents a null b2WorldId
	return WorldId{
		Index1:     uint16FromInt32(worldId + int32FromInt32(1)),
		Generation: (*b2World)(unsafe.Pointer(world)).Generation,
	}
}

func b2DestroyWorld(tls *_Stack, worldId WorldId) {
	var chain, set, world uintptr
	var chainCapacity, i, i1, i2, i3, sensorCount, setCapacity int32
	var generation uint16_t
	var v2 bool
	_, _, _, _, _, _, _, _, _, _, _, _ = chain, chainCapacity, generation, i, i1, i2, i3, sensorCount, set, setCapacity, world, v2
	world = b2GetWorldFromId(tls, worldId)
	// Detach any recording before teardown. The user owns and frees the recording buffer.
	b2StopRecordingInternal(tls, world)
	if (*b2World)(unsafe.Pointer(world)).Scheduler != uintptrFromInt32(0) {
		b2DestroyScheduler(tls, (*b2World)(unsafe.Pointer(world)).Scheduler)
		(*b2World)(unsafe.Pointer(world)).Scheduler = uintptrFromInt32(0)
	}
	b2DestroyBitSet(tls, world+2408)
	b2DestroyBitSet(tls, world+2424)
	b2DestroyBitSet(tls, world+2440)
	b2DestroyBitSet(tls, world+2456)
	b2DestroyWorkerContexts(tls, world)
	b2Free(tls, (*b2World)(unsafe.Pointer(world)).BodyMoveEvents.Data, uint64FromInt32((*b2World)(unsafe.Pointer(world)).BodyMoveEvents.Capacity)*uint64(40))
	(*b2World)(unsafe.Pointer(world)).BodyMoveEvents.Data = uintptrFromInt32(0)
	(*b2World)(unsafe.Pointer(world)).BodyMoveEvents.Count = 0
	(*b2World)(unsafe.Pointer(world)).BodyMoveEvents.Capacity = 0
	b2Free(tls, (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Data, uint64FromInt32((*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity)*uint64(16))
	(*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Data = uintptrFromInt32(0)
	(*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Count = 0
	(*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity = 0
	b2Free(tls, (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304))).Data, uint64FromInt32((*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304))).Capacity)*uint64(16))
	(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304))).Data = uintptrFromInt32(0)
	(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304))).Count = 0
	(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304))).Capacity = 0
	b2Free(tls, (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+1*16))).Data, uint64FromInt32((*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+1*16))).Capacity)*uint64(16))
	(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+1*16))).Data = uintptrFromInt32(0)
	(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+1*16))).Count = 0
	(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+1*16))).Capacity = 0
	b2Free(tls, (*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Data, uint64FromInt32((*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Capacity)*uint64(28))
	(*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Data = uintptrFromInt32(0)
	(*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Count = 0
	(*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Capacity = 0
	b2Free(tls, (*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336))).Data, uint64FromInt32((*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336))).Capacity)*uint64(28))
	(*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336))).Data = uintptrFromInt32(0)
	(*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336))).Count = 0
	(*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336))).Capacity = 0
	b2Free(tls, (*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+1*16))).Data, uint64FromInt32((*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+1*16))).Capacity)*uint64(28))
	(*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+1*16))).Data = uintptrFromInt32(0)
	(*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+1*16))).Count = 0
	(*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+1*16))).Capacity = 0
	b2Free(tls, (*b2World)(unsafe.Pointer(world)).ContactHitEvents.Data, uint64FromInt32((*b2World)(unsafe.Pointer(world)).ContactHitEvents.Capacity)*uint64(48))
	(*b2World)(unsafe.Pointer(world)).ContactHitEvents.Data = uintptrFromInt32(0)
	(*b2World)(unsafe.Pointer(world)).ContactHitEvents.Count = 0
	(*b2World)(unsafe.Pointer(world)).ContactHitEvents.Capacity = 0
	b2Free(tls, (*b2World)(unsafe.Pointer(world)).JointEvents.Data, uint64FromInt32((*b2World)(unsafe.Pointer(world)).JointEvents.Capacity)*uint64(16))
	(*b2World)(unsafe.Pointer(world)).JointEvents.Data = uintptrFromInt32(0)
	(*b2World)(unsafe.Pointer(world)).JointEvents.Count = 0
	(*b2World)(unsafe.Pointer(world)).JointEvents.Capacity = 0
	chainCapacity = (*b2World)(unsafe.Pointer(world)).ChainShapes.Count
	i = 0
	for {
		if !(i < chainCapacity) {
			break
		}
		chain = (*b2World)(unsafe.Pointer(world)).ChainShapes.Data + uintptr(i)*32
		if (*b2ChainShape)(unsafe.Pointer(chain)).Id != -int32(1) {
			b2FreeChainData(tls, chain)
		} else {
			if v2 = !!((*b2ChainShape)(unsafe.Pointer(chain)).ShapeIndices == uintptrFromInt32(0)); !v2 {
				b2InternalAssert(tls, __ccgo_ts+17400, __ccgo_ts+16660, int32FromInt32(370))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	sensorCount = (*b2World)(unsafe.Pointer(world)).Sensors.Count
	i1 = 0
	for {
		if !(i1 < sensorCount) {
			break
		}
		b2Free(tls, (*(*b2Sensor)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i1)*56))).Hits.Data, uint64FromInt32((*(*b2Sensor)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i1)*56))).Hits.Capacity)*uint64(8))
		(*(*b2Sensor)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i1)*56))).Hits.Data = uintptrFromInt32(0)
		(*(*b2Sensor)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i1)*56))).Hits.Count = 0
		(*(*b2Sensor)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i1)*56))).Hits.Capacity = 0
		b2Free(tls, (*(*b2Sensor)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i1)*56))).Overlaps1.Data, uint64FromInt32((*(*b2Sensor)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i1)*56))).Overlaps1.Capacity)*uint64(8))
		(*(*b2Sensor)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i1)*56))).Overlaps1.Data = uintptrFromInt32(0)
		(*(*b2Sensor)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i1)*56))).Overlaps1.Count = 0
		(*(*b2Sensor)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i1)*56))).Overlaps1.Capacity = 0
		b2Free(tls, (*(*b2Sensor)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i1)*56))).Overlaps2.Data, uint64FromInt32((*(*b2Sensor)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i1)*56))).Overlaps2.Capacity)*uint64(8))
		(*(*b2Sensor)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i1)*56))).Overlaps2.Data = uintptrFromInt32(0)
		(*(*b2Sensor)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i1)*56))).Overlaps2.Count = 0
		(*(*b2Sensor)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Sensors.Data+uintptr(i1)*56))).Overlaps2.Capacity = 0
		goto _3
	_3:
		;
		i1 = i1 + 1
	}
	b2Free(tls, (*b2World)(unsafe.Pointer(world)).Sensors.Data, uint64FromInt32((*b2World)(unsafe.Pointer(world)).Sensors.Capacity)*uint64(56))
	(*b2World)(unsafe.Pointer(world)).Sensors.Data = uintptrFromInt32(0)
	(*b2World)(unsafe.Pointer(world)).Sensors.Count = 0
	(*b2World)(unsafe.Pointer(world)).Sensors.Capacity = 0
	b2Free(tls, (*b2World)(unsafe.Pointer(world)).Bodies.Data, uint64FromInt32((*b2World)(unsafe.Pointer(world)).Bodies.Capacity)*uint64(104))
	(*b2World)(unsafe.Pointer(world)).Bodies.Data = uintptrFromInt32(0)
	(*b2World)(unsafe.Pointer(world)).Bodies.Count = 0
	(*b2World)(unsafe.Pointer(world)).Bodies.Capacity = 0
	b2Free(tls, (*b2World)(unsafe.Pointer(world)).Shapes.Data, uint64FromInt32((*b2World)(unsafe.Pointer(world)).Shapes.Capacity)*uint64(280))
	(*b2World)(unsafe.Pointer(world)).Shapes.Data = uintptrFromInt32(0)
	(*b2World)(unsafe.Pointer(world)).Shapes.Count = 0
	(*b2World)(unsafe.Pointer(world)).Shapes.Capacity = 0
	b2Free(tls, (*b2World)(unsafe.Pointer(world)).FatAABBs.Data, uint64FromInt32((*b2World)(unsafe.Pointer(world)).FatAABBs.Capacity)*uint64(16))
	(*b2World)(unsafe.Pointer(world)).FatAABBs.Data = uintptrFromInt32(0)
	(*b2World)(unsafe.Pointer(world)).FatAABBs.Count = 0
	(*b2World)(unsafe.Pointer(world)).FatAABBs.Capacity = 0
	b2Free(tls, (*b2World)(unsafe.Pointer(world)).ChainShapes.Data, uint64FromInt32((*b2World)(unsafe.Pointer(world)).ChainShapes.Capacity)*uint64(32))
	(*b2World)(unsafe.Pointer(world)).ChainShapes.Data = uintptrFromInt32(0)
	(*b2World)(unsafe.Pointer(world)).ChainShapes.Count = 0
	(*b2World)(unsafe.Pointer(world)).ChainShapes.Capacity = 0
	b2Free(tls, (*b2World)(unsafe.Pointer(world)).Contacts.Data, uint64FromInt32((*b2World)(unsafe.Pointer(world)).Contacts.Capacity)*uint64(64))
	(*b2World)(unsafe.Pointer(world)).Contacts.Data = uintptrFromInt32(0)
	(*b2World)(unsafe.Pointer(world)).Contacts.Count = 0
	(*b2World)(unsafe.Pointer(world)).Contacts.Capacity = 0
	b2Free(tls, (*b2World)(unsafe.Pointer(world)).Joints.Data, uint64FromInt32((*b2World)(unsafe.Pointer(world)).Joints.Capacity)*uint64(72))
	(*b2World)(unsafe.Pointer(world)).Joints.Data = uintptrFromInt32(0)
	(*b2World)(unsafe.Pointer(world)).Joints.Count = 0
	(*b2World)(unsafe.Pointer(world)).Joints.Capacity = 0
	i2 = 0
	for {
		if !(i2 < (*b2World)(unsafe.Pointer(world)).Islands.Count) {
			break
		}
		b2Free(tls, (*(*b2Island)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Islands.Data+uintptr(i2)*64))).Bodies.Data, uint64FromInt32((*(*b2Island)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Islands.Data+uintptr(i2)*64))).Bodies.Capacity)*uint64(4))
		(*(*b2Island)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Islands.Data+uintptr(i2)*64))).Bodies.Data = uintptrFromInt32(0)
		(*(*b2Island)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Islands.Data+uintptr(i2)*64))).Bodies.Count = 0
		(*(*b2Island)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Islands.Data+uintptr(i2)*64))).Bodies.Capacity = 0
		b2Free(tls, (*(*b2Island)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Islands.Data+uintptr(i2)*64))).Contacts.Data, uint64FromInt32((*(*b2Island)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Islands.Data+uintptr(i2)*64))).Contacts.Capacity)*uint64(12))
		(*(*b2Island)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Islands.Data+uintptr(i2)*64))).Contacts.Data = uintptrFromInt32(0)
		(*(*b2Island)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Islands.Data+uintptr(i2)*64))).Contacts.Count = 0
		(*(*b2Island)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Islands.Data+uintptr(i2)*64))).Contacts.Capacity = 0
		b2Free(tls, (*(*b2Island)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Islands.Data+uintptr(i2)*64))).Joints.Data, uint64FromInt32((*(*b2Island)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Islands.Data+uintptr(i2)*64))).Joints.Capacity)*uint64(12))
		(*(*b2Island)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Islands.Data+uintptr(i2)*64))).Joints.Data = uintptrFromInt32(0)
		(*(*b2Island)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Islands.Data+uintptr(i2)*64))).Joints.Count = 0
		(*(*b2Island)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Islands.Data+uintptr(i2)*64))).Joints.Capacity = 0
		goto _4
	_4:
		;
		i2 = i2 + 1
	}
	b2Free(tls, (*b2World)(unsafe.Pointer(world)).Islands.Data, uint64FromInt32((*b2World)(unsafe.Pointer(world)).Islands.Capacity)*uint64(64))
	(*b2World)(unsafe.Pointer(world)).Islands.Data = uintptrFromInt32(0)
	(*b2World)(unsafe.Pointer(world)).Islands.Count = 0
	(*b2World)(unsafe.Pointer(world)).Islands.Capacity = 0
	// Destroy solver sets
	setCapacity = (*b2World)(unsafe.Pointer(world)).SolverSets.Count
	i3 = 0
	for {
		if !(i3 < setCapacity) {
			break
		}
		set = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(i3)*88
		if (*b2SolverSet)(unsafe.Pointer(set)).SetIndex != -int32(1) {
			b2DestroySolverSet(tls, world, i3)
		}
		goto _5
	_5:
		;
		i3 = i3 + 1
	}
	b2Free(tls, (*b2World)(unsafe.Pointer(world)).SolverSets.Data, uint64FromInt32((*b2World)(unsafe.Pointer(world)).SolverSets.Capacity)*uint64(88))
	(*b2World)(unsafe.Pointer(world)).SolverSets.Data = uintptrFromInt32(0)
	(*b2World)(unsafe.Pointer(world)).SolverSets.Count = 0
	(*b2World)(unsafe.Pointer(world)).SolverSets.Capacity = 0
	b2DestroyGraph(tls, world+376)
	b2DestroyBroadPhase(tls, world+40)
	b2DestroyIdPool(tls, world+1912)
	b2DestroyIdPool(tls, world+2112)
	b2DestroyIdPool(tls, world+2136)
	b2DestroyIdPool(tls, world+2032)
	b2DestroyIdPool(tls, world+1992)
	b2DestroyIdPool(tls, world+2072)
	b2DestroyIdPool(tls, world+1952)
	b2DestroyStack(tls, world)
	// Wipe world but preserve generation
	generation = (*b2World)(unsafe.Pointer(world)).Generation
	*(*b2World)(unsafe.Add(nil, world)) = b2World{}
	(*b2World)(unsafe.Pointer(world)).WorldId = uint16(0)
	(*b2World)(unsafe.Pointer(world)).Generation = uint16FromInt32(int32FromUint16(generation) + int32(1))
}

func b2World_Step(tls *_Stack, worldId WorldId, timeStep float32, subStepCount int32) {
	bp := tls.Alloc(464)
	defer tls.Free(464)
	var a11, a2, a3, contactHertz, omega, v40, v41, v42, v44, v45, v46 float32
	var c, world, v19 uintptr
	var collideTicks, hash, pairTicks, sensorTicks, solveTicks, stepTicks uint64_t
	var staticBodyCount, totalBodyCount, totalContactCount, v4, v5, v6, v8, v9 int32
	var v1 bool
	var v47 b2Softness
	var _ /* _a at bp+0 */ b2RecArgs_Step
	var _ /* bounds at bp+424 */ AABB
	var _ /* context at bp+16 */ b2StepContext
	var _ /* stateHash at bp+408 */ b2RecArgs_StateHash
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a2, a3, c, collideTicks, contactHertz, hash, omega, pairTicks, sensorTicks, solveTicks, staticBodyCount, stepTicks, totalBodyCount, totalContactCount, world, v1, v19, v4, v40, v41, v42, v44, v45, v46, v47, v5, v6, v8, v9
	if int32FromUint8(b2IsValidFloat(tls, timeStep)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+448, __ccgo_ts+2656, uintptr(unsafe.Pointer(&__func__48))))
		if v1 = !!(b2IsValidFloat(tls, timeStep) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2656, __ccgo_ts+16660, int32FromInt32(886))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if v1 = !!(int32FromInt32(0) < subStepCount); !v1 {
		b2InternalAssert(tls, __ccgo_ts+17866, __ccgo_ts+16660, int32FromInt32(887))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(890))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	// Record step inputs before simulation runs
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_Step)(unsafe.Add(nil, bp)) = b2RecArgs_Step{
			World:        worldId,
			Dt:           timeStep,
			SubStepCount: subStepCount,
		}
		b2RecWrite_Step(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	// Prepare to capture events
	// Ensure user does not access stale data if there is an early return
	(*b2World)(unsafe.Pointer(world)).BodyMoveEvents.Count = 0
	(*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Count = 0
	(*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Count = 0
	(*b2World)(unsafe.Pointer(world)).ContactHitEvents.Count = 0
	(*b2World)(unsafe.Pointer(world)).JointEvents.Count = 0
	(*b2World)(unsafe.Pointer(world)).Profile = Profile{}
	(*b2World)(unsafe.Pointer(world)).Locked = boolUint8(true1 != 0)
	(*b2World)(unsafe.Pointer(world)).ActiveTaskCount = 0
	(*b2World)(unsafe.Pointer(world)).TaskCount = 0
	if (*b2World)(unsafe.Pointer(world)).Scheduler != uintptrFromInt32(0) {
		b2ResetScheduler(tls, (*b2World)(unsafe.Pointer(world)).Scheduler)
	}
	stepTicks = b2GetTicks(tls)
	c = world + 2640
	v4 = (*Capacity)(unsafe.Pointer(c)).StaticShapeCount
	v5 = (*(*b2DynamicTree)(unsafe.Add(nil, world+40+uintptr(b2_staticBody)*104))).ProxyCount
	if v4 > v5 {
		v8 = v4
	} else {
		v8 = v5
	}
	v6 = v8
	goto _7
_7:
	(*Capacity)(unsafe.Pointer(c)).StaticShapeCount = v6
	v4 = (*Capacity)(unsafe.Pointer(c)).DynamicShapeCount
	v5 = (*(*b2DynamicTree)(unsafe.Add(nil, world+40+uintptr(b2_dynamicBody)*104))).ProxyCount
	if v4 > v5 {
		v8 = v4
	} else {
		v8 = v5
	}
	v6 = v8
	goto _12
_12:
	(*Capacity)(unsafe.Pointer(c)).DynamicShapeCount = v6
	staticBodyCount = (*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_staticSet)*88))).BodySims.Count
	v4 = (*Capacity)(unsafe.Pointer(c)).StaticBodyCount
	v5 = staticBodyCount
	if v4 > v5 {
		v8 = v4
	} else {
		v8 = v5
	}
	v6 = v8
	goto _17
_17:
	(*Capacity)(unsafe.Pointer(c)).StaticBodyCount = v6
	v19 = world + 1912
	v4 = (*b2IdPool)(unsafe.Pointer(v19)).NextIndex - (*b2IdPool)(unsafe.Pointer(v19)).FreeArray.Count
	goto _21
_21:
	// this includes kinematic bodies
	totalBodyCount = v4
	v5 = (*Capacity)(unsafe.Pointer(c)).DynamicBodyCount
	v6 = totalBodyCount - staticBodyCount
	if v5 > v6 {
		v9 = v5
	} else {
		v9 = v6
	}
	v8 = v9
	goto _25
_25:
	(*Capacity)(unsafe.Pointer(c)).DynamicBodyCount = v8
	v19 = world + 2032
	v4 = (*b2IdPool)(unsafe.Pointer(v19)).NextIndex - (*b2IdPool)(unsafe.Pointer(v19)).FreeArray.Count
	goto _29
_29:
	totalContactCount = v4
	v5 = (*Capacity)(unsafe.Pointer(c)).ContactCount
	v6 = totalContactCount
	if v5 > v6 {
		v9 = v5
	} else {
		v9 = v6
	}
	v8 = v9
	goto _33
_33:
	(*Capacity)(unsafe.Pointer(c)).ContactCount = v8
	// Update collision pairs and create contacts
	pairTicks = b2GetTicks(tls)
	b2UpdateBroadPhasePairs(tls, world)
	(*b2World)(unsafe.Pointer(world)).Profile.Pairs = b2GetMilliseconds(tls, pairTicks)
	*(*b2StepContext)(unsafe.Add(nil, bp+16)) = b2StepContext{}
	(*(*b2StepContext)(unsafe.Add(nil, bp+16))).World = world
	(*(*b2StepContext)(unsafe.Add(nil, bp+16))).Dt = timeStep
	v4 = int32(1)
	v5 = subStepCount
	if v4 > v5 {
		v8 = v4
	} else {
		v8 = v5
	}
	v6 = v8
	goto _38
_38:
	(*(*b2StepContext)(unsafe.Add(nil, bp+16))).SubStepCount = v6
	if timeStep > float32FromFloat32(0) {
		(*(*b2StepContext)(unsafe.Add(nil, bp+16))).Inv_dt = float32FromFloat32(1) / timeStep
		(*(*b2StepContext)(unsafe.Add(nil, bp+16))).H = timeStep / float32((*(*b2StepContext)(unsafe.Add(nil, bp+16))).SubStepCount)
		(*(*b2StepContext)(unsafe.Add(nil, bp+16))).Inv_h = float32(float32((*(*b2StepContext)(unsafe.Add(nil, bp+16))).SubStepCount) * (*(*b2StepContext)(unsafe.Add(nil, bp+16))).Inv_dt)
	} else {
		(*(*b2StepContext)(unsafe.Add(nil, bp+16))).Inv_dt = float32FromFloat32(0)
		(*(*b2StepContext)(unsafe.Add(nil, bp+16))).H = float32FromFloat32(0)
		(*(*b2StepContext)(unsafe.Add(nil, bp+16))).Inv_h = float32FromFloat32(0)
	}
	(*b2World)(unsafe.Pointer(world)).Inv_h = (*(*b2StepContext)(unsafe.Add(nil, bp+16))).Inv_h
	(*b2World)(unsafe.Pointer(world)).Inv_dt = (*(*b2StepContext)(unsafe.Add(nil, bp+16))).Inv_dt
	v40 = (*b2World)(unsafe.Pointer(world)).ContactHertz
	v41 = float32(float32FromFloat32(0.125) * (*(*b2StepContext)(unsafe.Add(nil, bp+16))).Inv_h)
	if v40 < v41 {
		v44 = v40
	} else {
		v44 = v41
	}
	v42 = v44
	goto _43
_43:
	// Hertz values get reduced for large time steps
	contactHertz = v42
	v45 = contactHertz
	v46 = (*(*b2StepContext)(unsafe.Add(nil, bp+16))).H
	if v45 == float32FromFloat32(0) {
		v47 = b2Softness{}
		goto _48
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v45)
	a11 = float32(float32FromFloat32(2)*(*b2World)(unsafe.Pointer(world)).ContactDampingRatio) + float32(v46*omega)
	a2 = float32(float32(v46*omega) * a11)
	a3 = float32FromFloat32(1) / (float32FromFloat32(1) + a2)
	v47 = b2Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a2 * a3),
		ImpulseScale: a3,
	}
	goto _48
_48:
	(*(*b2StepContext)(unsafe.Add(nil, bp+16))).ContactSoftness = v47
	v40 = float32(float32FromFloat32(2) * contactHertz)
	v41 = (*(*b2StepContext)(unsafe.Add(nil, bp+16))).H
	if v40 == float32FromFloat32(0) {
		v47 = b2Softness{}
		goto _52
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v40)
	a11 = float32(float32FromFloat32(2)*(*b2World)(unsafe.Pointer(world)).ContactDampingRatio) + float32(v41*omega)
	a2 = float32(float32(v41*omega) * a11)
	a3 = float32FromFloat32(1) / (float32FromFloat32(1) + a2)
	v47 = b2Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a2 * a3),
		ImpulseScale: a3,
	}
	goto _52
_52:
	(*(*b2StepContext)(unsafe.Add(nil, bp+16))).StaticSoftness = v47
	(*(*b2StepContext)(unsafe.Add(nil, bp+16))).RestitutionThreshold = (*b2World)(unsafe.Pointer(world)).RestitutionThreshold
	(*(*b2StepContext)(unsafe.Add(nil, bp+16))).MaxLinearVelocity = (*b2World)(unsafe.Pointer(world)).MaxLinearSpeed
	(*(*b2StepContext)(unsafe.Add(nil, bp+16))).EnableWarmStarting = (*b2World)(unsafe.Pointer(world)).EnableWarmStarting
	// Narrow phase : update contacts
	collideTicks = b2GetTicks(tls)
	b2Collide(tls, bp+16)
	(*b2World)(unsafe.Pointer(world)).Profile.Collide = b2GetMilliseconds(tls, collideTicks)
	// Integrate velocities, solve velocity constraints, and integrate positions.
	if timeStep > float32FromFloat32(0) {
		solveTicks = b2GetTicks(tls)
		b2Solve(tls, world, bp+16)
		(*b2World)(unsafe.Pointer(world)).Profile.Solve = b2GetMilliseconds(tls, solveTicks)
	}
	// Finish the tree task in case b2Solve didn't finish it
	if (*b2World)(unsafe.Pointer(world)).UserTreeTask != 0 {
		(*(*func(*_Stack, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).FinishTaskFcn})))(tls, (*b2World)(unsafe.Pointer(world)).UserTreeTask, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
		(*b2World)(unsafe.Pointer(world)).UserTreeTask = uintptrFromInt32(0)
		*(*int32)(unsafe.Add(nil, world+2776)) -= int32(1)
		b2ValidateNoEnlarged(tls, world+40)
	}
	// Update sensors
	sensorTicks = b2GetTicks(tls)
	b2OverlapSensors(tls, world)
	(*b2World)(unsafe.Pointer(world)).Profile.Sensors = b2GetMilliseconds(tls, sensorTicks)
	(*b2World)(unsafe.Pointer(world)).Profile.Step = b2GetMilliseconds(tls, stepTicks)
	if v1 = !!(b2GetStackAllocation(tls, world) == int32FromInt32(0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+17883, __ccgo_ts+16660, int32FromInt32(1009))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	v4 = (*b2World)(unsafe.Pointer(world)).MaxCapacity.ArenaByteCount
	v5 = b2GetMaxStackAllocation(tls, world)
	if v4 > v5 {
		v8 = v4
	} else {
		v8 = v5
	}
	v6 = v8
	goto _57
_57:
	(*b2World)(unsafe.Pointer(world)).MaxCapacity.ArenaByteCount = v6
	// Ensure stack is large enough
	b2GrowStack(tls, world)
	// Make sure all tasks that were started were also finished
	if v1 = !!((*b2World)(unsafe.Pointer(world)).ActiveTaskCount == int32FromInt32(0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+17926, __ccgo_ts+16660, int32FromInt32(1017))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// Swap end event array buffers
	(*b2World)(unsafe.Pointer(world)).EndEventArrayIndex = int32(1) - (*b2World)(unsafe.Pointer(world)).EndEventArrayIndex
	(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count = 0
	(*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count = 0
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		// Write the per-step StateHash while the world is still locked. Queries early return while
		// locked, so this keeps the shared recording buffer single-writer without a lock. StateHash
		// proves the simulation reproduced exactly on replay.
		hash = b2HashWorldState(tls, world)
		*(*b2RecArgs_StateHash)(unsafe.Add(nil, bp+408)) = b2RecArgs_StateHash{
			World: worldId,
			Hash:  hash,
		}
		b2RecWrite_StateHash(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp+408)
		if b2ComputeWorldBounds(tls, world, bp+424) != 0 {
			b2RecAccumulateBounds(tls, (*b2World)(unsafe.Pointer(world)).Recording, *(*AABB)(unsafe.Add(nil, bp+424)))
		}
	}
	(*b2World)(unsafe.Pointer(world)).Locked = boolUint8(false1 != 0)
}

// C documentation
//
//	// todo this has varying order for moving shapes, causing flicker when overlapping shapes are moving
//	// solution: display order by shape id modulus 3, keep 3 buckets in GLSolid* and flush in 3 passes.
func b2World_Draw(tls *_Stack, worldId WorldId, draw uintptr) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var aabb, c, v111, v112, v133 AABB
	var addColor, frictionColor, impulseColor, normalColor, persistColor, speculativeColor b2HexColor1
	var bits, body1, bodyA, bodyB, bodySim, bodySim1, bodySim2, bodySimA, bodySimB, contact, contactSim, island, islandBody, joint, mp, set, shape, world1, v16, v17, v20, v26, v27, v30, v37 uintptr
	var blockIndex, blockIndex1, bodyId, ctz, k, wordCount, v13 uint32_t
	var bodyCapacity, bodyIndex, contactCapacity, contactId, contactKey, edgeIndex, edgeIndex1, i, islandBodyId, islandCapacity, islandId, j, jointCapacity, jointId, jointKey, pointCount, shapeCount, shapeId, v3 int32
	var force, force1, k_axisScale, linearSlop, pointSize, rx, ry, v113, v114, v115, v117, v118, v119, v120, v122, v123, v124, v125, v127, v72, v73, v81, v92 float32
	var normal, offset, offset1, tangent, v23, v33, v65, v69, v74 Vec2
	var p11, p12, p13, p2, p21, p22, p23, p3, p4, v24, v34 Pos
	var transform, transform1, v22 WorldTransform
	var word uint64_t
	var v1, v15, v18, v19, v2, v28 bool
	var v39 uint8
	var _ /* buffer at bp+16 */ [32]uint8
	var _ /* buffer at bp+48 */ [32]uint8
	var _ /* drawContext at bp+0 */ DrawContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, addColor, bits, blockIndex, blockIndex1, body1, bodyA, bodyB, bodyCapacity, bodyId, bodyIndex, bodySim, bodySim1, bodySim2, bodySimA, bodySimB, c, contact, contactCapacity, contactId, contactKey, contactSim, ctz, edgeIndex, edgeIndex1, force, force1, frictionColor, i, impulseColor, island, islandBody, islandBodyId, islandCapacity, islandId, j, joint, jointCapacity, jointId, jointKey, k, k_axisScale, linearSlop, mp, normal, normalColor, offset, offset1, p11, p12, p13, p2, p21, p22, p23, p3, p4, persistColor, pointCount, pointSize, rx, ry, set, shape, shapeCount, shapeId, speculativeColor, tangent, transform, transform1, word, wordCount, world1, v1, v111, v112, v113, v114, v115, v117, v118, v119, v120, v122, v123, v124, v125, v127, v13, v133, v15, v16, v17, v18, v19, v2, v20, v22, v23, v24, v26, v27, v28, v3, v30, v33, v34, v37, v39, v65, v69, v72, v73, v74, v81, v92
	world1 = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world1)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(1201))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world1)).Locked != 0 {
		return
	}
	if v1 = !!(b2IsValidAABB(tls, (*b2DebugDraw)(unsafe.Pointer(draw)).DrawingBounds) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+18045, __ccgo_ts+16660, int32FromInt32(1207))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	k_axisScale = float32FromFloat32(0.3)
	speculativeColor = int32(b2_colorGainsboro)
	addColor = int32(b2_colorGreen)
	persistColor = int32(b2_colorBlue)
	normalColor = int32(b2_colorDimGray)
	impulseColor = int32(b2_colorMagenta)
	frictionColor = int32(b2_colorYellow)
	v3 = (*b2IdPool)(unsafe.Pointer(world1 + 1912)).NextIndex
	goto _4
_4:
	bodyCapacity = v3
	b2SetBitCountAndClear(tls, world1+2408, uint32FromInt32(bodyCapacity))
	v3 = (*b2IdPool)(unsafe.Pointer(world1 + 1992)).NextIndex
	goto _6
_6:
	jointCapacity = v3
	b2SetBitCountAndClear(tls, world1+2424, uint32FromInt32(jointCapacity))
	v3 = (*b2IdPool)(unsafe.Pointer(world1 + 2032)).NextIndex
	goto _8
_8:
	contactCapacity = v3
	b2SetBitCountAndClear(tls, world1+2440, uint32FromInt32(contactCapacity))
	v3 = (*b2IdPool)(unsafe.Pointer(world1 + 2072)).NextIndex
	goto _10
_10:
	islandCapacity = v3
	b2SetBitCountAndClear(tls, world1+2456, uint32FromInt32(islandCapacity))
	*(*DrawContext)(unsafe.Add(nil, bp)) = DrawContext{
		World: world1,
		Draw:  draw,
	}
	i = 0
	for {
		if !(i < int32(b2_bodyTypeCount)) {
			break
		}
		b2DynamicTree_OverlapAABBAll(tls, world1+40+uintptr(i)*104, (*b2DebugDraw)(unsafe.Pointer(draw)).DrawingBounds, __ccgo_fp(DrawQueryCallback), bp)
		goto _11
	_11:
		;
		i = i + 1
	}
	wordCount = (*b2World)(unsafe.Pointer(world1)).DebugBodySet.BlockCount
	bits = (*b2World)(unsafe.Pointer(world1)).DebugBodySet.Bits
	k = uint32(0)
	for {
		if !(k < wordCount) {
			break
		}
		word = *(*uint64_t)(unsafe.Add(nil, bits+uintptr(k)*8))
		for word != uint64(0) {
			v13 = uint32FromInt32(__builtin_ctzll(tls, word))
			goto _14
		_14:
			ctz = v13
			bodyId = uint32(64)*k + ctz
			if v1 = !!(0 <= int32FromUint32(bodyId) && int32FromUint32(bodyId) < (*b2World)(unsafe.Pointer(world1)).Bodies.Count); !v1 {
				b2InternalAssert(tls, __ccgo_ts+18082, __ccgo_ts+16660, int32FromInt32(1246))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			body1 = (*b2World)(unsafe.Pointer(world1)).Bodies.Data + uintptr(int32FromUint32(bodyId))*104
			if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawBodyNames != 0 && int32FromUint8(*(*uint8)(unsafe.Add(nil, body1+90))) != 0 {
				offset = Vec2{
					X: float32FromFloat32(0.1),
					Y: float32FromFloat32(0.1),
				}
				v16 = world1
				v17 = body1
				if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(v17)).SetIndex && (*b2Body)(unsafe.Pointer(v17)).SetIndex < (*b2World)(unsafe.Pointer(v16)).SolverSets.Count); !v2 {
					b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
				}
				_ = v2 || bool(int32FromInt32(0) != 0)
				set = (*b2World)(unsafe.Pointer(v16)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v17)).SetIndex)*88
				if v15 = !!(0 <= (*b2Body)(unsafe.Pointer(v17)).LocalIndex && (*b2Body)(unsafe.Pointer(v17)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v15 {
					b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
				}
				_ = v15 || bool(int32FromInt32(0) != 0)
				bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v17)).LocalIndex)*96
				v20 = bodySim
				goto _21
			_21:
				bodySim1 = v20
				transform = WorldTransform{
					P: (*b2BodySim)(unsafe.Pointer(bodySim1)).Center,
					Q: (*b2BodySim)(unsafe.Pointer(bodySim1)).Transform.Q,
				}
				v22 = transform
				v23 = offset
				rx = float32(v22.Q.C*v23.X) - float32(v22.Q.S*v23.Y)
				ry = float32(v22.Q.S*v23.X) + float32(v22.Q.C*v23.Y)
				v24 = Pos{
					X: v22.P.X + rx,
					Y: v22.P.Y + ry,
				}
				goto _25
			_25:
				p2 = v24
				(*(*func(*_Stack, Pos, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, p2, body1+90, int32(b2_colorWhiteSmoke), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
			}
			if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawMass != 0 && (*b2Body)(unsafe.Pointer(body1)).Type1 == int32(b2_dynamicBody) {
				offset1 = Vec2{
					X: float32FromFloat32(0.1),
					Y: float32FromFloat32(0.1),
				}
				v16 = world1
				v17 = body1
				if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(v17)).SetIndex && (*b2Body)(unsafe.Pointer(v17)).SetIndex < (*b2World)(unsafe.Pointer(v16)).SolverSets.Count); !v1 {
					b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				set = (*b2World)(unsafe.Pointer(v16)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v17)).SetIndex)*88
				if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(v17)).LocalIndex && (*b2Body)(unsafe.Pointer(v17)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v2 {
					b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
				}
				_ = v2 || bool(int32FromInt32(0) != 0)
				bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v17)).LocalIndex)*96
				v20 = bodySim
				goto _31
			_31:
				bodySim2 = v20
				transform1 = WorldTransform{
					P: (*b2BodySim)(unsafe.Pointer(bodySim2)).Center,
					Q: (*b2BodySim)(unsafe.Pointer(bodySim2)).Transform.Q,
				}
				(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, (*b2BodySim)(unsafe.Pointer(bodySim2)).Center0, (*b2BodySim)(unsafe.Pointer(bodySim2)).Center, int32(b2_colorWhiteSmoke), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
				(*(*func(*_Stack, WorldTransform, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawTransformFcn})))(tls, transform1, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
				v22 = transform1
				v23 = offset1
				rx = float32(v22.Q.C*v23.X) - float32(v22.Q.S*v23.Y)
				ry = float32(v22.Q.S*v23.X) + float32(v22.Q.C*v23.Y)
				v24 = Pos{
					X: v22.P.X + rx,
					Y: v22.P.Y + ry,
				}
				goto _35
			_35:
				p3 = v24
				__builtin_snprintf(tls, bp+16, uint64(32), __ccgo_ts+18148, vaList(bp+88, float64((*b2Body)(unsafe.Pointer(body1)).Mass)))
				(*(*func(*_Stack, Pos, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, p3, bp+16, int32(b2_colorWhite), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
			}
			if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawJoints != 0 {
				jointKey = (*b2Body)(unsafe.Pointer(body1)).HeadJointKey
				for jointKey != -int32(1) {
					jointId = jointKey >> int32(1)
					edgeIndex = jointKey & int32(1)
					if v1 = !!(0 <= jointId && jointId < (*b2World)(unsafe.Pointer(world1)).Joints.Count); !v1 {
						b2InternalAssert(tls, __ccgo_ts+1850, __ccgo_ts+16660, int32FromInt32(1280))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					joint = (*b2World)(unsafe.Pointer(world1)).Joints.Data + uintptr(jointId)*72
					// avoid double draw
					v16 = world1 + 2424
					v13 = uint32FromInt32(jointId)
					blockIndex1 = v13 / uint32(64)
					if blockIndex1 >= (*b2BitSet)(unsafe.Pointer(v16)).BlockCount {
						v39 = boolUint8(false1 != 0)
						goto _40
					}
					v39 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v16)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v13%uint32FromInt32(64))) != uint64(0))
					goto _40
				_40:
					if int32FromUint8(v39) == false1 {
						b2DrawJoint(tls, draw, world1, joint)
						v16 = world1 + 2424
						v13 = uint32FromInt32(jointId)
						blockIndex = v13 / uint32(64)
						if v1 = !!(blockIndex < (*b2BitSet)(unsafe.Pointer(v16)).BlockCount); !v1 {
							b2InternalAssert(tls, __ccgo_ts+7356, __ccgo_ts+7388, int32FromInt32(29))
						}
						_ = v1 || bool(int32FromInt32(0) != 0)
						*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v16)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v13 % uint32FromInt32(64))
					}
					jointKey = (*(*b2JointEdge)(unsafe.Add(nil, joint+20+uintptr(edgeIndex)*12))).NextKey
				}
			}
			linearSlop = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
			if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawContacts != 0 && (*b2Body)(unsafe.Pointer(body1)).Type1 == int32(b2_dynamicBody) {
				contactKey = (*b2Body)(unsafe.Pointer(body1)).HeadContactKey
				for contactKey != -int32(1) {
					contactId = contactKey >> int32(1)
					edgeIndex1 = contactKey & int32(1)
					if v1 = !!(0 <= contactId && contactId < (*b2World)(unsafe.Pointer(world1)).Contacts.Count); !v1 {
						b2InternalAssert(tls, __ccgo_ts+301, __ccgo_ts+16660, int32FromInt32(1301))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					contact = (*b2World)(unsafe.Pointer(world1)).Contacts.Data + uintptr(contactId)*64
					contactKey = (*(*b2ContactEdge)(unsafe.Add(nil, contact+uintptr(edgeIndex1)*12))).NextKey
					// avoid double draw
					v16 = world1 + 2440
					v13 = uint32FromInt32(contactId)
					blockIndex1 = v13 / uint32(64)
					if blockIndex1 >= (*b2BitSet)(unsafe.Pointer(v16)).BlockCount {
						v39 = boolUint8(false1 != 0)
						goto _48
					}
					v39 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v16)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v13%uint32FromInt32(64))) != uint64(0))
					goto _48
				_48:
					if int32FromUint8(v39) == false1 {
						contactSim = b2GetContactSim(tls, world1, contact)
						if v1 = !!(0 <= (*(*b2ContactEdge)(unsafe.Add(nil, contact))).BodyId && (*(*b2ContactEdge)(unsafe.Add(nil, contact))).BodyId < (*b2World)(unsafe.Pointer(world1)).Bodies.Count); !v1 {
							b2InternalAssert(tls, __ccgo_ts+18155, __ccgo_ts+16660, int32FromInt32(1308))
						}
						_ = v1 || bool(int32FromInt32(0) != 0)
						bodyA = (*b2World)(unsafe.Pointer(world1)).Bodies.Data + uintptr((*(*b2ContactEdge)(unsafe.Add(nil, contact))).BodyId)*104
						v17 = world1
						v20 = bodyA
						if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(v20)).SetIndex && (*b2Body)(unsafe.Pointer(v20)).SetIndex < (*b2World)(unsafe.Pointer(v17)).SolverSets.Count); !v2 {
							b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
						}
						_ = v2 || bool(int32FromInt32(0) != 0)
						set = (*b2World)(unsafe.Pointer(v17)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v20)).SetIndex)*88
						if v15 = !!(0 <= (*b2Body)(unsafe.Pointer(v20)).LocalIndex && (*b2Body)(unsafe.Pointer(v20)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v15 {
							b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
						}
						_ = v15 || bool(int32FromInt32(0) != 0)
						bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v20)).LocalIndex)*96
						v26 = bodySim
						goto _55
					_55:
						bodySimA = v26
						if v18 = !!(0 <= (*(*b2ContactEdge)(unsafe.Add(nil, contact+1*12))).BodyId && (*(*b2ContactEdge)(unsafe.Add(nil, contact+1*12))).BodyId < (*b2World)(unsafe.Pointer(world1)).Bodies.Count); !v18 {
							b2InternalAssert(tls, __ccgo_ts+18247, __ccgo_ts+16660, int32FromInt32(1310))
						}
						_ = v18 || bool(int32FromInt32(0) != 0)
						bodyB = (*b2World)(unsafe.Pointer(world1)).Bodies.Data + uintptr((*(*b2ContactEdge)(unsafe.Add(nil, contact+1*12))).BodyId)*104
						v27 = world1
						v30 = bodyB
						if v19 = !!(0 <= (*b2Body)(unsafe.Pointer(v30)).SetIndex && (*b2Body)(unsafe.Pointer(v30)).SetIndex < (*b2World)(unsafe.Pointer(v27)).SolverSets.Count); !v19 {
							b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
						}
						_ = v19 || bool(int32FromInt32(0) != 0)
						set = (*b2World)(unsafe.Pointer(v27)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v30)).SetIndex)*88
						if v28 = !!(0 <= (*b2Body)(unsafe.Pointer(v30)).LocalIndex && (*b2Body)(unsafe.Pointer(v30)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v28 {
							b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
						}
						_ = v28 || bool(int32FromInt32(0) != 0)
						bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v30)).LocalIndex)*96
						v37 = bodySim
						goto _62
					_62:
						bodySimB = v37
						pointCount = (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.PointCount
						normal = (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold.Normal
						j = 0
						for {
							if !(j < pointCount) {
								break
							}
							mp = contactSim + 72 + 12 + uintptr(j)*44
							if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawAnchorA != 0 {
								v24 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Center
								v23 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
								v34 = Pos{
									X: v24.X + v23.X,
									Y: v24.Y + v23.Y,
								}
								goto _67
							_67:
								p4 = v34
							} else {
								v24 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Center
								v23 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorB
								v34 = Pos{
									X: v24.X + v23.X,
									Y: v24.Y + v23.Y,
								}
								goto _71
							_71:
								p4 = v34
							}
							if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawGraphColors != 0 && (*b2Contact)(unsafe.Pointer(contact)).ColorIndex != -int32(1) {
								if (*b2Contact)(unsafe.Pointer(contact)).ColorIndex == int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1) {
									v72 = float32FromFloat32(7.5)
								} else {
									v72 = float32FromFloat32(5)
								}
								// graph color
								pointSize = v72
								(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p4, pointSize, b2GetGraphColor(tls, (*b2Contact)(unsafe.Pointer(contact)).ColorIndex), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
								// m_context->draw.DrawString(point->position, "%d", point->color);
							} else {
								if (*ManifoldPoint)(unsafe.Pointer(mp)).Separation > linearSlop {
									// Speculative
									(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p4, float32FromFloat32(5), speculativeColor, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
								} else {
									if int32FromUint8((*ManifoldPoint)(unsafe.Pointer(mp)).Persisted) == false1 {
										// Add
										(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p4, float32FromFloat32(10), addColor, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
									} else {
										if int32FromUint8((*ManifoldPoint)(unsafe.Pointer(mp)).Persisted) == int32(true1) {
											// Persist
											(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p4, float32FromFloat32(5), persistColor, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
										}
									}
								}
							}
							if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawContactNormals != 0 {
								p11 = p4
								v72 = k_axisScale
								v23 = normal
								v33 = Vec2{
									X: float32(v72 * v23.X),
									Y: float32(v72 * v23.Y),
								}
								goto _76
							_76:
								v24 = p11
								v65 = v33
								v34 = Pos{
									X: v24.X + v65.X,
									Y: v24.Y + v65.Y,
								}
								goto _80
							_80:
								p21 = v34
								(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, p11, p21, normalColor, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
								__builtin_snprintf(tls, bp+48, uint64FromInt32(int32FromUint64(uint64FromInt64(32)/uint64FromInt64(1))), __ccgo_ts+18339, vaList(bp+88, float64((*ManifoldPoint)(unsafe.Pointer(mp)).Separation)))
								(*(*func(*_Stack, Pos, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, p11, bp+48, int32(b2_colorWhite), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
							} else {
								if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawContactForces != 0 {
									// todo validate
									// multiply by one-half due to relax iteration
									force = float32(float32(float32FromFloat32(0.5)*(*ManifoldPoint)(unsafe.Pointer(mp)).TotalNormalImpulse) * (*b2World)(unsafe.Pointer(world1)).Inv_dt)
									p12 = p4
									v72 = float32((*b2DebugDraw)(unsafe.Pointer(draw)).ForceScale * force)
									v23 = normal
									v33 = Vec2{
										X: float32(v72 * v23.X),
										Y: float32(v72 * v23.Y),
									}
									goto _84
								_84:
									v24 = p12
									v65 = v33
									v34 = Pos{
										X: v24.X + v65.X,
										Y: v24.Y + v65.Y,
									}
									goto _88
								_88:
									p22 = v34
									(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, p12, p22, impulseColor, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
									__builtin_snprintf(tls, bp+48, uint64FromInt32(int32FromUint64(uint64FromInt64(32)/uint64FromInt64(1))), __ccgo_ts+18345, vaList(bp+88, float64(force)))
									(*(*func(*_Stack, Pos, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, p12, bp+48, int32(b2_colorWhite), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
								}
							}
							if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawContactFeatures != 0 {
								__builtin_snprintf(tls, bp+48, uint64FromInt32(int32FromUint64(uint64FromInt64(32)/uint64FromInt64(1))), __ccgo_ts+18350, vaList(bp+88, int32FromUint16((*ManifoldPoint)(unsafe.Pointer(mp)).Id)))
								(*(*func(*_Stack, Pos, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, p4, bp+48, int32(b2_colorOrange), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
							}
							if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawFrictionForces != 0 {
								force1 = float32(float32(float32FromFloat32(0.5)*(*ManifoldPoint)(unsafe.Pointer(mp)).TangentImpulse) * (*b2World)(unsafe.Pointer(world1)).Inv_h)
								v23 = normal
								v33 = Vec2{
									X: v23.Y,
									Y: -v23.X,
								}
								goto _91
							_91:
								tangent = v33
								p13 = p4
								v72 = float32((*b2DebugDraw)(unsafe.Pointer(draw)).ForceScale * force1)
								v65 = tangent
								v69 = Vec2{
									X: float32(v72 * v65.X),
									Y: float32(v72 * v65.Y),
								}
								goto _95
							_95:
								v24 = p13
								v74 = v69
								v34 = Pos{
									X: v24.X + v74.X,
									Y: v24.Y + v74.Y,
								}
								goto _99
							_99:
								p23 = v34
								(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, p13, p23, frictionColor, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
								__builtin_snprintf(tls, bp+48, uint64FromInt32(int32FromUint64(uint64FromInt64(32)/uint64FromInt64(1))), __ccgo_ts+18345, vaList(bp+88, float64(force1)))
								(*(*func(*_Stack, Pos, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, p13, bp+48, int32(b2_colorWhite), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
							}
							goto _63
						_63:
							;
							j = j + 1
						}
						v16 = world1 + 2440
						v13 = uint32FromInt32(contactId)
						blockIndex = v13 / uint32(64)
						if v1 = !!(blockIndex < (*b2BitSet)(unsafe.Pointer(v16)).BlockCount); !v1 {
							b2InternalAssert(tls, __ccgo_ts+7356, __ccgo_ts+7388, int32FromInt32(29))
						}
						_ = v1 || bool(int32FromInt32(0) != 0)
						*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v16)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v13 % uint32FromInt32(64))
					}
					contactKey = (*(*b2ContactEdge)(unsafe.Add(nil, contact+uintptr(edgeIndex1)*12))).NextKey
				}
			}
			if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawIslands != 0 {
				islandId = (*b2Body)(unsafe.Pointer(body1)).IslandId
				if v1 = islandId != -int32(1); v1 {
					v16 = world1 + 2456
					v13 = uint32FromInt32(islandId)
					blockIndex1 = v13 / uint32(64)
					if blockIndex1 >= (*b2BitSet)(unsafe.Pointer(v16)).BlockCount {
						v39 = boolUint8(false1 != 0)
						goto _106
					}
					v39 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v16)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v13%uint32FromInt32(64))) != uint64(0))
					goto _106
				_106:
				}
				if v1 && int32FromUint8(v39) == false1 {
					island = (*b2World)(unsafe.Pointer(world1)).Islands.Data + uintptr(islandId)*64
					if (*b2Island)(unsafe.Pointer(island)).SetIndex == -int32(1) {
						continue
					}
					shapeCount = 0
					aabb = AABB{
						LowerBound: Vec2{
							X: float32FromFloat32(3.4028234663852886e+38),
							Y: float32FromFloat32(3.4028234663852886e+38),
						},
						UpperBound: Vec2{
							X: -float32FromFloat32(3.4028234663852886e+38),
							Y: -float32FromFloat32(3.4028234663852886e+38),
						},
					}
					bodyIndex = 0
					for {
						if !(bodyIndex < (*b2Island)(unsafe.Pointer(island)).Bodies.Count) {
							break
						}
						islandBodyId = *(*int32)(unsafe.Add(nil, (*b2Island)(unsafe.Pointer(island)).Bodies.Data+uintptr(bodyIndex)*4))
						if v1 = !!(0 <= islandBodyId && islandBodyId < (*b2World)(unsafe.Pointer(world1)).Bodies.Count); !v1 {
							b2InternalAssert(tls, __ccgo_ts+18353, __ccgo_ts+16660, int32FromInt32(1419))
						}
						_ = v1 || bool(int32FromInt32(0) != 0)
						islandBody = (*b2World)(unsafe.Pointer(world1)).Bodies.Data + uintptr(islandBodyId)*104
						shapeId = (*b2Body)(unsafe.Pointer(islandBody)).HeadShapeId
						for shapeId != -int32(1) {
							if v2 = !!(0 <= shapeId && shapeId < (*b2World)(unsafe.Pointer(world1)).Shapes.Count); !v2 {
								b2InternalAssert(tls, __ccgo_ts+1908, __ccgo_ts+16660, int32FromInt32(1423))
							}
							_ = v2 || bool(int32FromInt32(0) != 0)
							shape = (*b2World)(unsafe.Pointer(world1)).Shapes.Data + uintptr(shapeId)*280
							v111 = aabb
							v112 = *(*AABB)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world1)).FatAABBs.Data+uintptr(shapeId)*16))
							v72 = v111.LowerBound.X
							v73 = v112.LowerBound.X
							if v72 < v73 {
								v92 = v72
							} else {
								v92 = v73
							}
							v81 = v92
							goto _116
						_116:
							c.LowerBound.X = v81
							v113 = v111.LowerBound.Y
							v114 = v112.LowerBound.Y
							if v113 < v114 {
								v117 = v113
							} else {
								v117 = v114
							}
							v115 = v117
							goto _121
						_121:
							c.LowerBound.Y = v115
							v118 = v111.UpperBound.X
							v119 = v112.UpperBound.X
							if v118 > v119 {
								v122 = v118
							} else {
								v122 = v119
							}
							v120 = v122
							goto _126
						_126:
							c.UpperBound.X = v120
							v123 = v111.UpperBound.Y
							v124 = v112.UpperBound.Y
							if v123 > v124 {
								v127 = v123
							} else {
								v127 = v124
							}
							v125 = v127
							goto _131
						_131:
							c.UpperBound.Y = v125
							v133 = c
							goto _134
						_134:
							aabb = v133
							shapeCount = shapeCount + int32(1)
							shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
						}
						goto _108
					_108:
						;
						bodyIndex = bodyIndex + 1
					}
					if shapeCount > 0 {
						(*(*func(*_Stack, AABB, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawBoundsFcn})))(tls, aabb, int32(b2_colorOrangeRed), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
					}
					v16 = world1 + 2456
					v13 = uint32FromInt32(islandId)
					blockIndex = v13 / uint32(64)
					if v1 = !!(blockIndex < (*b2BitSet)(unsafe.Pointer(v16)).BlockCount); !v1 {
						b2InternalAssert(tls, __ccgo_ts+7356, __ccgo_ts+7388, int32FromInt32(29))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v16)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v13 % uint32FromInt32(64))
				}
			}
			// Clear the smallest set bit
			word = word & (word - uint64(1))
		}
		goto _12
	_12:
		;
		k = k + 1
	}
}

func b2ComputeWorldBounds(tls *_Stack, world uintptr, bounds uintptr) (r uint8) {
	var c, treeBounds, worldBounds, v2, v25, v3, v4 AABB
	var haveBounds uint8
	var i int32
	var tree uintptr
	var v10, v11, v12, v14, v15, v16, v17, v19, v20, v21, v22, v24, v5, v6, v7, v9 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, haveBounds, i, tree, treeBounds, worldBounds, v10, v11, v12, v14, v15, v16, v17, v19, v2, v20, v21, v22, v24, v25, v3, v4, v5, v6, v7, v9
	worldBounds = AABB{}
	haveBounds = boolUint8(false1 != 0)
	i = 0
	for {
		if !(i < int32(b2_bodyTypeCount)) {
			break
		}
		tree = world + 40 + uintptr(i)*104
		if b2DynamicTree_GetProxyCount(tls, tree) == 0 {
			goto _1
		}
		treeBounds = b2DynamicTree_GetRootBounds(tls, tree)
		if haveBounds != 0 {
			v3 = worldBounds
			v4 = treeBounds
			v5 = v3.LowerBound.X
			v6 = v4.LowerBound.X
			if v5 < v6 {
				v9 = v5
			} else {
				v9 = v6
			}
			v7 = v9
			goto _8
		_8:
			c.LowerBound.X = v7
			v10 = v3.LowerBound.Y
			v11 = v4.LowerBound.Y
			if v10 < v11 {
				v14 = v10
			} else {
				v14 = v11
			}
			v12 = v14
			goto _13
		_13:
			c.LowerBound.Y = v12
			v15 = v3.UpperBound.X
			v16 = v4.UpperBound.X
			if v15 > v16 {
				v19 = v15
			} else {
				v19 = v16
			}
			v17 = v19
			goto _18
		_18:
			c.UpperBound.X = v17
			v20 = v3.UpperBound.Y
			v21 = v4.UpperBound.Y
			if v20 > v21 {
				v24 = v20
			} else {
				v24 = v21
			}
			v22 = v24
			goto _23
		_23:
			c.UpperBound.Y = v22
			v25 = c
			goto _26
		_26:
			v2 = v25
		} else {
			v2 = treeBounds
		}
		worldBounds = v2
		haveBounds = boolUint8(true1 != 0)
		goto _1
	_1:
		;
		i = i + 1
	}
	*(*AABB)(unsafe.Add(nil, bounds)) = worldBounds
	return haveBounds
}

func b2World_GetBounds(tls *_Stack, worldId WorldId) (r AABB) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var world uintptr
	var _ /* bounds at bp+0 */ AABB
	_ = world
	world = b2GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return AABB{}
	}
	b2ComputeWorldBounds(tls, world, bp)
	return *(*AABB)(unsafe.Add(nil, bp))
}

func b2World_GetBodyEvents(tls *_Stack, worldId WorldId) (r b2BodyEvents) {
	var count int32
	var events b2BodyEvents
	var world uintptr
	var v1 bool
	_, _, _, _ = count, events, world, v1
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(1483))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return b2BodyEvents{}
	}
	count = (*b2World)(unsafe.Pointer(world)).BodyMoveEvents.Count
	events = b2BodyEvents{
		MoveEvents: (*b2World)(unsafe.Pointer(world)).BodyMoveEvents.Data,
		MoveCount:  count,
	}
	return events
}

func b2World_GetSensorEvents(tls *_Stack, worldId WorldId) (r b2SensorEvents) {
	var beginCount, endCount, endEventArrayIndex int32
	var events b2SensorEvents
	var world uintptr
	var v1 bool
	_, _, _, _, _, _ = beginCount, endCount, endEventArrayIndex, events, world, v1
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(1497))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return b2SensorEvents{}
	}
	// Careful to use previous buffer
	endEventArrayIndex = int32(1) - (*b2World)(unsafe.Pointer(world)).EndEventArrayIndex
	beginCount = (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Count
	endCount = (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr(endEventArrayIndex)*16))).Count
	events = b2SensorEvents{
		BeginEvents: (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Data,
		EndEvents:   (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr(endEventArrayIndex)*16))).Data,
		BeginCount:  beginCount,
		EndCount:    endCount,
	}
	return events
}

func b2World_GetContactEvents(tls *_Stack, worldId WorldId) (r b2ContactEvents) {
	var beginCount, endCount, endEventArrayIndex, hitCount int32
	var events b2ContactEvents
	var world uintptr
	var v1 bool
	_, _, _, _, _, _, _ = beginCount, endCount, endEventArrayIndex, events, hitCount, world, v1
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(1521))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return b2ContactEvents{}
	}
	// Careful to use previous buffer
	endEventArrayIndex = int32(1) - (*b2World)(unsafe.Pointer(world)).EndEventArrayIndex
	beginCount = (*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Count
	endCount = (*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+uintptr(endEventArrayIndex)*16))).Count
	hitCount = (*b2World)(unsafe.Pointer(world)).ContactHitEvents.Count
	events = b2ContactEvents{
		BeginEvents: (*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Data,
		EndEvents:   (*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+uintptr(endEventArrayIndex)*16))).Data,
		HitEvents:   (*b2World)(unsafe.Pointer(world)).ContactHitEvents.Data,
		BeginCount:  beginCount,
		EndCount:    endCount,
		HitCount:    hitCount,
	}
	return events
}

func b2World_GetJointEvents(tls *_Stack, worldId WorldId) (r b2JointEvents) {
	var count int32
	var events b2JointEvents
	var world uintptr
	var v1 bool
	_, _, _, _ = count, events, world, v1
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(1549))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return b2JointEvents{}
	}
	count = (*b2World)(unsafe.Pointer(world)).JointEvents.Count
	events = b2JointEvents{
		JointEvents: (*b2World)(unsafe.Pointer(world)).JointEvents.Data,
		Count:       count,
	}
	return events
}

func b2World_IsValid(tls *_Stack, id WorldId) (r uint8) {
	var world uintptr
	_ = world
	if int32FromUint16(id.Index1) < int32(1) || int32(_B2_MAX_WORLDS) < int32FromUint16(id.Index1) {
		return boolUint8(false1 != 0)
	}
	world = uintptr(unsafe.Pointer(&b2_worlds)) + uintptr(int32FromUint16(id.Index1)-int32FromInt32(1))*2792
	if int32FromUint16((*b2World)(unsafe.Pointer(world)).WorldId) != int32FromUint16(id.Index1)-int32(1) {
		// world is not allocated
		return boolUint8(false1 != 0)
	}
	return boolUint8(int32FromUint16(id.Generation) == int32FromUint16((*b2World)(unsafe.Pointer(world)).Generation))
}

func b2World_EnableSleeping(tls *_Stack, worldId WorldId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i, setCount int32
	var set, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_WorldEnableSleeping
	_, _, _, _, _ = i, set, setCount, world, v1
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(1748))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WorldEnableSleeping)(unsafe.Add(nil, bp)) = b2RecArgs_WorldEnableSleeping{
			World: worldId,
			Flag:  flag,
		}
		b2RecWrite_WorldEnableSleeping(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	if int32FromUint8(flag) == int32FromUint8((*b2World)(unsafe.Pointer(world)).EnableSleep) {
		return
	}
	(*b2World)(unsafe.Pointer(world)).EnableSleep = flag
	if int32FromUint8(flag) == false1 {
		setCount = (*b2World)(unsafe.Pointer(world)).SolverSets.Count
		i = int32(b2_firstSleepingSet)
		for {
			if !(i < setCount) {
				break
			}
			if v1 = !!(0 <= i && i < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
				b2InternalAssert(tls, __ccgo_ts+18534, __ccgo_ts+16660, int32FromInt32(1768))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			set = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(i)*88
			if (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count > 0 {
				b2WakeSolverSet(tls, world, i)
			}
			goto _2
		_2:
			;
			i = i + 1
		}
	}
}

func b2World_IsSleepingEnabled(tls *_Stack, worldId WorldId) (r uint8) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).EnableSleep
}

func b2World_EnableWarmStarting(tls *_Stack, worldId WorldId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_WorldEnableWarmStarting
	_, _ = world, v1
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(1786))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WorldEnableWarmStarting)(unsafe.Add(nil, bp)) = b2RecArgs_WorldEnableWarmStarting{
			World: worldId,
			Flag:  flag,
		}
		b2RecWrite_WorldEnableWarmStarting(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b2World)(unsafe.Pointer(world)).EnableWarmStarting = flag
}

func b2World_IsWarmStartingEnabled(tls *_Stack, worldId WorldId) (r uint8) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).EnableWarmStarting
}

func b2World_EnableSSE2Fallback(tls *_Stack, worldId WorldId, flag uint8) {
	var world uintptr
	var v1 bool
	var v2 int32
	_, _, _ = world, v1, v2
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(1806))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if flag != 0 {
		v2 = int32(4)
	} else {
		v2 = b2GetSIMDWidth(tls)
	}
	(*b2World)(unsafe.Pointer(world)).SimdWidth = v2
}

func b2World_GetAwakeBodyCount(tls *_Stack, worldId WorldId) (r int32) {
	var awakeSet, world uintptr
	var v1 bool
	_, _, _ = awakeSet, world, v1
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+16660, int32FromInt32(1818))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	awakeSet = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
	return (*b2SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Count
}

func b2World_EnableContinuous(tls *_Stack, worldId WorldId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_WorldEnableContinuous
	_, _ = world, v1
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(1825))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WorldEnableContinuous)(unsafe.Add(nil, bp)) = b2RecArgs_WorldEnableContinuous{
			World: worldId,
			Flag:  flag,
		}
		b2RecWrite_WorldEnableContinuous(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b2World)(unsafe.Pointer(world)).EnableContinuous = flag
}

func b2World_IsContinuousEnabled(tls *_Stack, worldId WorldId) (r uint8) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).EnableContinuous
}

func b2World_SetRestitutionThreshold(tls *_Stack, worldId WorldId, value float32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var world uintptr
	var v1 bool
	var v3, v4, v5, v6, v8, v9 float32
	var _ /* _a at bp+0 */ b2RecArgs_WorldSetRestitutionThreshold
	_, _, _, _, _, _, _, _ = world, v1, v3, v4, v5, v6, v8, v9
	if int32FromUint8(b2IsValidFloat(tls, value)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+16, __ccgo_ts+18584, uintptr(unsafe.Pointer(&__func__112))))
		if v1 = !!(b2IsValidFloat(tls, value) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+18584, __ccgo_ts+16660, int32FromInt32(1844))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(1847))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WorldSetRestitutionThreshold)(unsafe.Add(nil, bp)) = b2RecArgs_WorldSetRestitutionThreshold{
			World: worldId,
			Value: value,
		}
		b2RecWrite_WorldSetRestitutionThreshold(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	v3 = value
	v4 = float32FromFloat32(0)
	v5 = float32FromFloat32(3.4028234663852886e+38)
	if v3 < v4 {
		v8 = v4
	} else {
		if v3 > v5 {
			v9 = v5
		} else {
			v9 = v3
		}
		v8 = v9
	}
	v6 = v8
	goto _7
_7:
	(*b2World)(unsafe.Pointer(world)).RestitutionThreshold = v6
}

func b2World_GetRestitutionThreshold(tls *_Stack, worldId WorldId) (r float32) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).RestitutionThreshold
}

func b2World_SetRestitutionIterations(tls *_Stack, worldId WorldId, iterations int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var world uintptr
	var v1 bool
	var v3, v4, v5, v6, v8, v9 int32
	var _ /* _a at bp+0 */ b2RecArgs_WorldSetRestitutionIterations
	_, _, _, _, _, _, _, _ = world, v1, v3, v4, v5, v6, v8, v9
	if boolInt32(iterations >= 0) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+16, __ccgo_ts+18608, uintptr(unsafe.Pointer(&__func__113))))
		if v1 = !!(iterations >= int32FromInt32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+18608, __ccgo_ts+16660, int32FromInt32(1866))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(1869))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WorldSetRestitutionIterations)(unsafe.Add(nil, bp)) = b2RecArgs_WorldSetRestitutionIterations{
			World:      worldId,
			Iterations: iterations,
		}
		b2RecWrite_WorldSetRestitutionIterations(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	v3 = iterations
	v4 = 0
	v5 = int32(_B2_MAX_RESTITUTION_ITERATIONS)
	if v3 < v4 {
		v8 = v4
	} else {
		if v3 > v5 {
			v9 = v5
		} else {
			v9 = v3
		}
		v8 = v9
	}
	v6 = v8
	goto _7
_7:
	(*b2World)(unsafe.Pointer(world)).RestitutionIterations = v6
}

func b2World_GetRestitutionIterations(tls *_Stack, worldId WorldId) (r int32) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).RestitutionIterations
}

func b2World_EnableRestitutionPropagation(tls *_Stack, worldId WorldId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_WorldEnableRestitutionPropagation
	_, _ = world, v1
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(1889))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WorldEnableRestitutionPropagation)(unsafe.Add(nil, bp)) = b2RecArgs_WorldEnableRestitutionPropagation{
			World: worldId,
			Flag:  flag,
		}
		b2RecWrite_WorldEnableRestitutionPropagation(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b2World)(unsafe.Pointer(world)).EnableRestitutionPropagation = flag
}

func b2World_IsRestitutionPropagationEnabled(tls *_Stack, worldId WorldId) (r uint8) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).EnableRestitutionPropagation
}

func b2World_SetHitEventThreshold(tls *_Stack, worldId WorldId, value float32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var world uintptr
	var v1 bool
	var v3, v4, v5, v6, v8, v9 float32
	var _ /* _a at bp+0 */ b2RecArgs_WorldSetHitEventThreshold
	_, _, _, _, _, _, _, _ = world, v1, v3, v4, v5, v6, v8, v9
	if int32FromUint8(b2IsValidFloat(tls, value)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+16, __ccgo_ts+18584, uintptr(unsafe.Pointer(&__func__114))))
		if v1 = !!(b2IsValidFloat(tls, value) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+18584, __ccgo_ts+16660, int32FromInt32(1908))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(1911))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WorldSetHitEventThreshold)(unsafe.Add(nil, bp)) = b2RecArgs_WorldSetHitEventThreshold{
			World: worldId,
			Value: value,
		}
		b2RecWrite_WorldSetHitEventThreshold(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	v3 = value
	v4 = float32FromFloat32(0)
	v5 = float32FromFloat32(3.4028234663852886e+38)
	if v3 < v4 {
		v8 = v4
	} else {
		if v3 > v5 {
			v9 = v5
		} else {
			v9 = v3
		}
		v8 = v9
	}
	v6 = v8
	goto _7
_7:
	(*b2World)(unsafe.Pointer(world)).HitEventThreshold = v6
}

func b2World_GetHitEventThreshold(tls *_Stack, worldId WorldId) (r float32) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).HitEventThreshold
}

func b2World_SetContactTuning(tls *_Stack, worldId WorldId, hertz float32, dampingRatio float32, pushSpeed float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var world uintptr
	var v1 bool
	var v10, v11, v5, v6, v7, v8 float32
	var _ /* _a at bp+0 */ b2RecArgs_WorldSetContactTuning
	_, _, _, _, _, _, _, _ = world, v1, v10, v11, v5, v6, v7, v8
	if int32FromUint8(b2IsValidFloat(tls, hertz)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8368, uintptr(unsafe.Pointer(&__func__115))))
		if v1 = !!(b2IsValidFloat(tls, hertz) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8368, __ccgo_ts+16660, int32FromInt32(1930))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if int32FromUint8(b2IsValidFloat(tls, dampingRatio)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8392, uintptr(unsafe.Pointer(&__func__115))))
		if v1 = !!(b2IsValidFloat(tls, dampingRatio) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8392, __ccgo_ts+16660, int32FromInt32(1931))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if int32FromUint8(b2IsValidFloat(tls, pushSpeed)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+18624, uintptr(unsafe.Pointer(&__func__115))))
		if v1 = !!(b2IsValidFloat(tls, pushSpeed) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+18624, __ccgo_ts+16660, int32FromInt32(1932))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(1935))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WorldSetContactTuning)(unsafe.Add(nil, bp)) = b2RecArgs_WorldSetContactTuning{
			World:        worldId,
			Hertz:        hertz,
			DampingRatio: dampingRatio,
			PushSpeed:    pushSpeed,
		}
		b2RecWrite_WorldSetContactTuning(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	v5 = hertz
	v6 = float32FromFloat32(0)
	v7 = float32FromFloat32(3.4028234663852886e+38)
	if v5 < v6 {
		v10 = v6
	} else {
		if v5 > v7 {
			v11 = v7
		} else {
			v11 = v5
		}
		v10 = v11
	}
	v8 = v10
	goto _9
_9:
	(*b2World)(unsafe.Pointer(world)).ContactHertz = v8
	v5 = dampingRatio
	v6 = float32FromFloat32(0)
	v7 = float32FromFloat32(3.4028234663852886e+38)
	if v5 < v6 {
		v10 = v6
	} else {
		if v5 > v7 {
			v11 = v7
		} else {
			v11 = v5
		}
		v10 = v11
	}
	v8 = v10
	goto _16
_16:
	(*b2World)(unsafe.Pointer(world)).ContactDampingRatio = v8
	v5 = pushSpeed
	v6 = float32FromFloat32(0)
	v7 = float32FromFloat32(3.4028234663852886e+38)
	if v5 < v6 {
		v10 = v6
	} else {
		if v5 > v7 {
			v11 = v7
		} else {
			v11 = v5
		}
		v10 = v11
	}
	v8 = v10
	goto _23
_23:
	(*b2World)(unsafe.Pointer(world)).ContactSpeed = v8
}

func b2World_GetContactTuning(tls *_Stack, worldId WorldId, hertz uintptr, dampingRatio uintptr, pushSpeed uintptr) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	*(*float32)(unsafe.Add(nil, hertz)) = (*b2World)(unsafe.Pointer(world)).ContactHertz
	*(*float32)(unsafe.Add(nil, dampingRatio)) = (*b2World)(unsafe.Pointer(world)).ContactDampingRatio
	*(*float32)(unsafe.Add(nil, pushSpeed)) = (*b2World)(unsafe.Pointer(world)).ContactSpeed
}

func b2World_SetContactRecycleDistance(tls *_Stack, worldId WorldId, recycleDistance float32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var world uintptr
	var v1 bool
	var v3, v4, v5, v6, v8, v9 float32
	var _ /* _a at bp+0 */ b2RecArgs_WorldSetContactRecycleDistance
	_, _, _, _, _, _, _, _ = world, v1, v3, v4, v5, v6, v8, v9
	if int32FromUint8(b2IsValidFloat(tls, recycleDistance)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+16, __ccgo_ts+18652, uintptr(unsafe.Pointer(&__func__116))))
		if v1 = !!(b2IsValidFloat(tls, recycleDistance) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+18652, __ccgo_ts+16660, int32FromInt32(1958))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(1961))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WorldSetContactRecycleDistance)(unsafe.Add(nil, bp)) = b2RecArgs_WorldSetContactRecycleDistance{
			World:           worldId,
			RecycleDistance: recycleDistance,
		}
		b2RecWrite_WorldSetContactRecycleDistance(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	v3 = recycleDistance
	v4 = float32FromFloat32(0)
	v5 = float32FromFloat32(3.4028234663852886e+38)
	if v3 < v4 {
		v8 = v4
	} else {
		if v3 > v5 {
			v9 = v5
		} else {
			v9 = v3
		}
		v8 = v9
	}
	v6 = v8
	goto _7
_7:
	(*b2World)(unsafe.Pointer(world)).ContactRecycleDistance = v6
}

func b2World_GetContactRecycleDistance(tls *_Stack, worldId WorldId) (r float32) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).ContactRecycleDistance
}

func b2World_SetMaxLinearSpeed(tls *_Stack, worldId WorldId, maximumLinearSpeed float32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_WorldSetMaxLinearSpeed
	_, _ = world, v1
	if boolInt32(b2IsValidFloat(tls, maximumLinearSpeed) != 0 && maximumLinearSpeed > float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+16, __ccgo_ts+18686, uintptr(unsafe.Pointer(&__func__49))))
		if v1 = !!(b2IsValidFloat(tls, maximumLinearSpeed) != 0 && maximumLinearSpeed > float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+18686, __ccgo_ts+16660, int32FromInt32(1980))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(1983))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WorldSetMaxLinearSpeed)(unsafe.Add(nil, bp)) = b2RecArgs_WorldSetMaxLinearSpeed{
			World:              worldId,
			MaximumLinearSpeed: maximumLinearSpeed,
		}
		b2RecWrite_WorldSetMaxLinearSpeed(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b2World)(unsafe.Pointer(world)).MaxLinearSpeed = maximumLinearSpeed
}

func b2World_GetMaxLinearSpeed(tls *_Stack, worldId WorldId) (r float32) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).MaxLinearSpeed
}

func b2World_GetProfile(tls *_Stack, worldId WorldId) (r Profile) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).Profile
}

func b2World_GetCounters(tls *_Stack, worldId WorldId) (r Counters) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var color, dynamicTree, kinematicTree, staticTree, world, v1 uintptr
	var i, i1, v11, v2, v5, v8 int32
	var _ /* s at bp+0 */ Counters
	_, _, _, _, _, _, _, _, _, _, _, _ = color, dynamicTree, i, i1, kinematicTree, staticTree, world, v1, v11, v2, v5, v8
	world = b2GetWorldFromId(tls, worldId)
	*(*Counters)(unsafe.Add(nil, bp)) = Counters{}
	v1 = world + 1912
	v2 = (*b2IdPool)(unsafe.Pointer(v1)).NextIndex - (*b2IdPool)(unsafe.Pointer(v1)).FreeArray.Count
	goto _3
_3:
	(*(*Counters)(unsafe.Add(nil, bp))).BodyCount = v2
	v1 = world + 2112
	v2 = (*b2IdPool)(unsafe.Pointer(v1)).NextIndex - (*b2IdPool)(unsafe.Pointer(v1)).FreeArray.Count
	goto _6
_6:
	(*(*Counters)(unsafe.Add(nil, bp))).ShapeCount = v2
	v1 = world + 2032
	v2 = (*b2IdPool)(unsafe.Pointer(v1)).NextIndex - (*b2IdPool)(unsafe.Pointer(v1)).FreeArray.Count
	goto _9
_9:
	(*(*Counters)(unsafe.Add(nil, bp))).ContactCount = v2
	v1 = world + 1992
	v2 = (*b2IdPool)(unsafe.Pointer(v1)).NextIndex - (*b2IdPool)(unsafe.Pointer(v1)).FreeArray.Count
	goto _12
_12:
	(*(*Counters)(unsafe.Add(nil, bp))).JointCount = v2
	v1 = world + 2072
	v2 = (*b2IdPool)(unsafe.Pointer(v1)).NextIndex - (*b2IdPool)(unsafe.Pointer(v1)).FreeArray.Count
	goto _15
_15:
	(*(*Counters)(unsafe.Add(nil, bp))).IslandCount = v2
	staticTree = world + 40 + uintptr(b2_staticBody)*104
	(*(*Counters)(unsafe.Add(nil, bp))).StaticTreeHeight = b2DynamicTree_GetHeight(tls, staticTree)
	dynamicTree = world + 40 + uintptr(b2_dynamicBody)*104
	kinematicTree = world + 40 + uintptr(b2_kinematicBody)*104
	v2 = b2DynamicTree_GetHeight(tls, dynamicTree)
	v5 = b2DynamicTree_GetHeight(tls, kinematicTree)
	if v2 > v5 {
		v11 = v2
	} else {
		v11 = v5
	}
	v8 = v11
	goto _19
_19:
	(*(*Counters)(unsafe.Add(nil, bp))).TreeHeight = v8
	(*(*Counters)(unsafe.Add(nil, bp))).StackUsed = b2GetMaxStackAllocation(tls, world)
	(*(*Counters)(unsafe.Add(nil, bp))).ByteCount = b2GetByteCount(tls)
	(*(*Counters)(unsafe.Add(nil, bp))).TaskCount = (*b2World)(unsafe.Pointer(world)).TaskCount
	(*(*Counters)(unsafe.Add(nil, bp))).RecycledContactCount = 0
	i = 0
	for {
		if !(i < (*b2World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		(*(*Counters)(unsafe.Add(nil, bp))).RecycledContactCount += (*(*b2TaskContext)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i)*136))).RecycledContactCount
		goto _21
	_21:
		;
		i = i + 1
	}
	(*(*Counters)(unsafe.Add(nil, bp))).AwakeContactCount = 0
	i1 = 0
	for {
		if !(i1 < int32(_B2_GRAPH_COLOR_COUNT)) {
			break
		}
		color = world + 376 + uintptr(i1)*64
		*(*int32)(unsafe.Add(nil, bp+44+uintptr(i1)*4)) = (*b2GraphColor)(unsafe.Pointer(color)).ContactSims.Count + (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Count
		(*(*Counters)(unsafe.Add(nil, bp))).AwakeContactCount += (*b2GraphColor)(unsafe.Pointer(color)).ContactSims.Count
		goto _22
	_22:
		;
		i1 = i1 + 1
	}
	(*(*Counters)(unsafe.Add(nil, bp))).AwakeContactCount += (*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(b2_awakeSet)*88))).ContactSims.Count
	return *(*Counters)(unsafe.Add(nil, bp))
}

func b2World_GetMaxCapacity(tls *_Stack, worldId WorldId) (r Capacity) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).MaxCapacity
}

func b2World_SetUserData(tls *_Stack, worldId WorldId, userData uintptr) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	(*b2World)(unsafe.Pointer(world)).UserData = userData
}

func b2World_GetUserData(tls *_Stack, worldId WorldId) (r uintptr) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).UserData
}

func b2World_SetFrictionCallback(tls *_Stack, worldId WorldId, __ccgo_fp_callback uintptr) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	(*b2World)(unsafe.Pointer(world)).FrictionCallback = __ccgo_fp_callback
}

func b2World_SetRestitutionCallback(tls *_Stack, worldId WorldId, __ccgo_fp_callback uintptr) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	(*b2World)(unsafe.Pointer(world)).RestitutionCallback = __ccgo_fp_callback
}

func b2World_SetWorkerCount(tls *_Stack, worldId WorldId, count int32) {
	var world uintptr
	var v1, v2, v3, v4, v6, v7 int32
	_, _, _, _, _, _, _ = world, v1, v2, v3, v4, v6, v7
	world = b2GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	if count == (*b2World)(unsafe.Pointer(world)).WorkerCount {
		return
	}
	b2DestroyWorkerContexts(tls, world)
	v1 = count
	v2 = int32(1)
	v3 = int32(_B2_MAX_WORKERS)
	if v1 < v2 {
		v6 = v2
	} else {
		if v1 > v3 {
			v7 = v3
		} else {
			v7 = v1
		}
		v6 = v7
	}
	v4 = v6
	goto _5
_5:
	(*b2World)(unsafe.Pointer(world)).WorkerCount = v4
	b2CreateWorkerContexts(tls, world)
}

func b2World_GetWorkerCount(tls *_Stack, worldId WorldId) (r int32) {
	var world uintptr
	_ = world
	world = b2GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return 0
	}
	return (*b2World)(unsafe.Pointer(world)).WorkerCount
}

func b2World_StartRecording(tls *_Stack, worldId WorldId, recording uintptr) {
	var world uintptr
	_ = world
	// Must be a step boundary, so refuse a locked world
	world = b2GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) || recording == uintptrFromInt32(0) || (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).PreSolveFcn != uintptrFromInt32(0) || (*b2World)(unsafe.Pointer(world)).PreContinuousFcn != uintptrFromInt32(0) || (*b2World)(unsafe.Pointer(world)).CustomFilterFcn != uintptrFromInt32(0) {
		printf(tls, __ccgo_ts+18752, 0)
		return
	}
	b2StartRecordingIntoBuffer(tls, world, recording)
}

func b2World_StopRecording(tls *_Stack, worldId WorldId) {
	var world uintptr
	_ = world
	world = b2GetUnlockedWorldFromId(tls, worldId)
	if world == uintptrFromInt32(0) {
		return
	}
	b2StopRecordingInternal(tls, world)
}

func b2World_DumpMemoryStats(tls *_Stack, worldId WorldId) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var bodyArrayBytes, bodyBitSetBytes, bodyIdBytes, bodySimCapacity, bodyStateCapacity, chainArrayBytes, chainDataBytes, chainIdBytes, contactArrayBytes, contactIdBytes, contactSimCapacity, debugBytes, dynamicTreeBytes, eventBytes, graphContactSimBytes, graphJointSimBytes, i, i1, i2, i3, i4, i5, i6, islandArrayBytes, islandIdBytes, islandLinkBytes, islandSimCapacity, jointArrayBytes, jointIdBytes, jointSimCapacity, kinematicTreeBytes, pairSetBytes, sensorArrayBytes, sensorOverlapBytes, sensorTaskContextBytes, setBodySimBytes, setBodyStateBytes, setContactSimBytes, setIslandSimBytes, setJointSimBytes, shapeArrayBytes, shapeIdBytes, solverSetArrayBytes, solverSetCapacity, solverSetIdBytes, staticTreeBytes, taskContextBytes, total, v11, v14, v17, v2, v20, v5, v8 int32
	var c, chain, file, island, pairSet, sensor, set, taskContext, taskContext1, world, v1, v10, v13, v16, v19, v4, v7 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyArrayBytes, bodyBitSetBytes, bodyIdBytes, bodySimCapacity, bodyStateCapacity, c, chain, chainArrayBytes, chainDataBytes, chainIdBytes, contactArrayBytes, contactIdBytes, contactSimCapacity, debugBytes, dynamicTreeBytes, eventBytes, file, graphContactSimBytes, graphJointSimBytes, i, i1, i2, i3, i4, i5, i6, island, islandArrayBytes, islandIdBytes, islandLinkBytes, islandSimCapacity, jointArrayBytes, jointIdBytes, jointSimCapacity, kinematicTreeBytes, pairSet, pairSetBytes, sensor, sensorArrayBytes, sensorOverlapBytes, sensorTaskContextBytes, set, setBodySimBytes, setBodyStateBytes, setContactSimBytes, setIslandSimBytes, setJointSimBytes, shapeArrayBytes, shapeIdBytes, solverSetArrayBytes, solverSetCapacity, solverSetIdBytes, staticTreeBytes, taskContext, taskContext1, taskContextBytes, total, world, v1, v10, v11, v13, v14, v16, v17, v19, v2, v20, v4, v5, v7, v8
	file = fopen(tls, __ccgo_ts+18840, __ccgo_ts+18857)
	if file == uintptrFromInt32(0) {
		return
	}
	world = b2GetWorldFromId(tls, worldId)
	total = 0
	v1 = world + 1912
	v2 = (*b2IdPool)(unsafe.Pointer(v1)).FreeArray.Capacity * int32FromInt64(4)
	goto _3
_3:
	// id pools
	bodyIdBytes = v2
	v4 = world + 1952
	v5 = (*b2IdPool)(unsafe.Pointer(v4)).FreeArray.Capacity * int32FromInt64(4)
	goto _6
_6:
	solverSetIdBytes = v5
	v7 = world + 1992
	v8 = (*b2IdPool)(unsafe.Pointer(v7)).FreeArray.Capacity * int32FromInt64(4)
	goto _9
_9:
	jointIdBytes = v8
	v10 = world + 2032
	v11 = (*b2IdPool)(unsafe.Pointer(v10)).FreeArray.Capacity * int32FromInt64(4)
	goto _12
_12:
	contactIdBytes = v11
	v13 = world + 2072
	v14 = (*b2IdPool)(unsafe.Pointer(v13)).FreeArray.Capacity * int32FromInt64(4)
	goto _15
_15:
	islandIdBytes = v14
	v16 = world + 2112
	v17 = (*b2IdPool)(unsafe.Pointer(v16)).FreeArray.Capacity * int32FromInt64(4)
	goto _18
_18:
	shapeIdBytes = v17
	v19 = world + 2136
	v20 = (*b2IdPool)(unsafe.Pointer(v19)).FreeArray.Capacity * int32FromInt64(4)
	goto _21
_21:
	chainIdBytes = v20
	total = total + (bodyIdBytes + solverSetIdBytes + jointIdBytes + contactIdBytes + islandIdBytes + shapeIdBytes + chainIdBytes)
	fprintf(tls, file, __ccgo_ts+18859, 0)
	fprintf(tls, file, __ccgo_ts+18869, vaList(bp+8, bodyIdBytes))
	fprintf(tls, file, __ccgo_ts+18883, vaList(bp+8, solverSetIdBytes))
	fprintf(tls, file, __ccgo_ts+18903, vaList(bp+8, jointIdBytes))
	fprintf(tls, file, __ccgo_ts+18918, vaList(bp+8, contactIdBytes))
	fprintf(tls, file, __ccgo_ts+18935, vaList(bp+8, islandIdBytes))
	fprintf(tls, file, __ccgo_ts+18951, vaList(bp+8, shapeIdBytes))
	fprintf(tls, file, __ccgo_ts+18966, vaList(bp+8, chainIdBytes))
	fprintf(tls, file, __ccgo_ts+18981, 0)
	// Islands own per-island body/contact/joint link arrays
	islandLinkBytes = 0
	i = 0
	for {
		if !(i < (*b2World)(unsafe.Pointer(world)).Islands.Count) {
			break
		}
		island = (*b2World)(unsafe.Pointer(world)).Islands.Data + uintptr(i)*64
		islandLinkBytes = islandLinkBytes + (*b2Island)(unsafe.Pointer(island)).Bodies.Capacity*int32FromInt64(4)
		islandLinkBytes = islandLinkBytes + (*b2Island)(unsafe.Pointer(island)).Contacts.Capacity*int32FromInt64(12)
		islandLinkBytes = islandLinkBytes + (*b2Island)(unsafe.Pointer(island)).Joints.Capacity*int32FromInt64(12)
		goto _22
	_22:
		;
		i = i + 1
	}
	// world arrays
	bodyArrayBytes = (*b2World)(unsafe.Pointer(world)).Bodies.Capacity * int32FromInt64(104)
	solverSetArrayBytes = (*b2World)(unsafe.Pointer(world)).SolverSets.Capacity * int32FromInt64(88)
	jointArrayBytes = (*b2World)(unsafe.Pointer(world)).Joints.Capacity * int32FromInt64(72)
	contactArrayBytes = (*b2World)(unsafe.Pointer(world)).Contacts.Capacity * int32FromInt64(64)
	islandArrayBytes = (*b2World)(unsafe.Pointer(world)).Islands.Capacity * int32FromInt64(64)
	shapeArrayBytes = (*b2World)(unsafe.Pointer(world)).Shapes.Capacity*int32FromInt64(280) + (*b2World)(unsafe.Pointer(world)).FatAABBs.Capacity*int32FromInt64(16)
	chainArrayBytes = (*b2World)(unsafe.Pointer(world)).ChainShapes.Capacity * int32FromInt64(32)
	sensorArrayBytes = (*b2World)(unsafe.Pointer(world)).Sensors.Capacity * int32FromInt64(56)
	total = total + (bodyArrayBytes + solverSetArrayBytes + jointArrayBytes + contactArrayBytes + islandArrayBytes + islandLinkBytes + shapeArrayBytes + chainArrayBytes + sensorArrayBytes)
	fprintf(tls, file, __ccgo_ts+18983, 0)
	fprintf(tls, file, __ccgo_ts+18997, vaList(bp+8, bodyArrayBytes))
	fprintf(tls, file, __ccgo_ts+19009, vaList(bp+8, solverSetArrayBytes))
	fprintf(tls, file, __ccgo_ts+19026, vaList(bp+8, jointArrayBytes))
	fprintf(tls, file, __ccgo_ts+19038, vaList(bp+8, contactArrayBytes))
	fprintf(tls, file, __ccgo_ts+19052, vaList(bp+8, islandArrayBytes))
	fprintf(tls, file, __ccgo_ts+19065, vaList(bp+8, islandLinkBytes))
	fprintf(tls, file, __ccgo_ts+19083, vaList(bp+8, shapeArrayBytes))
	fprintf(tls, file, __ccgo_ts+19095, vaList(bp+8, chainArrayBytes))
	fprintf(tls, file, __ccgo_ts+19107, vaList(bp+8, sensorArrayBytes))
	fprintf(tls, file, __ccgo_ts+18981, 0)
	// Chain shapes own index and surface material arrays
	chainDataBytes = 0
	i1 = 0
	for {
		if !(i1 < (*b2World)(unsafe.Pointer(world)).ChainShapes.Count) {
			break
		}
		chain = (*b2World)(unsafe.Pointer(world)).ChainShapes.Data + uintptr(i1)*32
		if (*b2ChainShape)(unsafe.Pointer(chain)).Id == -int32(1) {
			goto _23
		}
		chainDataBytes = chainDataBytes + (*b2ChainShape)(unsafe.Pointer(chain)).SegmentCount*int32FromInt64(4)
		goto _23
	_23:
		;
		i1 = i1 + 1
	}
	// Sensors own overlap tracking arrays. The sensor array is dense.
	sensorOverlapBytes = 0
	i2 = 0
	for {
		if !(i2 < (*b2World)(unsafe.Pointer(world)).Sensors.Count) {
			break
		}
		sensor = (*b2World)(unsafe.Pointer(world)).Sensors.Data + uintptr(i2)*56
		sensorOverlapBytes = sensorOverlapBytes + (*b2Sensor)(unsafe.Pointer(sensor)).Hits.Capacity*int32FromInt64(8)
		sensorOverlapBytes = sensorOverlapBytes + (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Capacity*int32FromInt64(8)
		sensorOverlapBytes = sensorOverlapBytes + (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Capacity*int32FromInt64(8)
		goto _24
	_24:
		;
		i2 = i2 + 1
	}
	total = total + (chainDataBytes + sensorOverlapBytes)
	fprintf(tls, file, __ccgo_ts+19120, 0)
	fprintf(tls, file, __ccgo_ts+19134, vaList(bp+8, chainDataBytes))
	fprintf(tls, file, __ccgo_ts+19150, vaList(bp+8, sensorOverlapBytes))
	fprintf(tls, file, __ccgo_ts+18981, 0)
	// broad-phase
	staticTreeBytes = int32FromUint64(b2DynamicTree_GetByteCount(tls, world+40+uintptr(b2_staticBody)*104))
	kinematicTreeBytes = int32FromUint64(b2DynamicTree_GetByteCount(tls, world+40+uintptr(b2_kinematicBody)*104))
	dynamicTreeBytes = int32FromUint64(b2DynamicTree_GetByteCount(tls, world+40+uintptr(b2_dynamicBody)*104))
	pairSet = world + 40 + 320
	pairSetBytes = b2GetHashSetBytes(tls, pairSet)
	total = total + (staticTreeBytes + kinematicTreeBytes + dynamicTreeBytes + pairSetBytes)
	fprintf(tls, file, __ccgo_ts+19171, 0)
	fprintf(tls, file, __ccgo_ts+19184, vaList(bp+8, staticTreeBytes))
	fprintf(tls, file, __ccgo_ts+19201, vaList(bp+8, kinematicTreeBytes))
	fprintf(tls, file, __ccgo_ts+19221, vaList(bp+8, dynamicTreeBytes))
	fprintf(tls, file, __ccgo_ts+19239, vaList(bp+8, pairSetBytes, (*b2HashSet)(unsafe.Pointer(pairSet)).Count, (*b2HashSet)(unsafe.Pointer(pairSet)).Capacity))
	fprintf(tls, file, __ccgo_ts+18981, 0)
	// solver sets
	bodySimCapacity = 0
	bodyStateCapacity = 0
	jointSimCapacity = 0
	contactSimCapacity = 0
	islandSimCapacity = 0
	solverSetCapacity = (*b2World)(unsafe.Pointer(world)).SolverSets.Count
	i3 = 0
	for {
		if !(i3 < solverSetCapacity) {
			break
		}
		set = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(i3)*88
		if (*b2SolverSet)(unsafe.Pointer(set)).SetIndex == -int32(1) {
			goto _25
		}
		bodySimCapacity = bodySimCapacity + (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Capacity
		bodyStateCapacity = bodyStateCapacity + (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Capacity
		jointSimCapacity = jointSimCapacity + (*b2SolverSet)(unsafe.Pointer(set)).JointSims.Capacity
		contactSimCapacity = contactSimCapacity + (*b2SolverSet)(unsafe.Pointer(set)).ContactSims.Capacity
		islandSimCapacity = islandSimCapacity + (*b2SolverSet)(unsafe.Pointer(set)).IslandSims.Capacity
		goto _25
	_25:
		;
		i3 = i3 + 1
	}
	setBodySimBytes = bodySimCapacity * int32FromInt64(96)
	setBodyStateBytes = bodyStateCapacity * int32FromInt64(32)
	setJointSimBytes = jointSimCapacity * int32FromInt64(252)
	setContactSimBytes = contactSimCapacity * int32FromInt64(200)
	setIslandSimBytes = islandSimCapacity * int32FromInt64(4)
	total = total + (setBodySimBytes + setBodyStateBytes + setJointSimBytes + setContactSimBytes + setIslandSimBytes)
	fprintf(tls, file, __ccgo_ts+19261, 0)
	fprintf(tls, file, __ccgo_ts+19274, vaList(bp+8, setBodySimBytes))
	fprintf(tls, file, __ccgo_ts+19288, vaList(bp+8, setBodyStateBytes))
	fprintf(tls, file, __ccgo_ts+19304, vaList(bp+8, setJointSimBytes))
	fprintf(tls, file, __ccgo_ts+19319, vaList(bp+8, setContactSimBytes))
	fprintf(tls, file, __ccgo_ts+19336, vaList(bp+8, setIslandSimBytes))
	fprintf(tls, file, __ccgo_ts+18981, 0)
	// constraint graph
	bodyBitSetBytes = 0
	contactSimCapacity = 0
	jointSimCapacity = 0
	i4 = 0
	for {
		if !(i4 < int32(_B2_GRAPH_COLOR_COUNT)) {
			break
		}
		c = world + 376 + uintptr(i4)*64
		v2 = int32FromUint64(uint64((*b2BitSet)(unsafe.Pointer(c)).BlockCapacity) * uint64(8))
		goto _28
	_28:
		bodyBitSetBytes = bodyBitSetBytes + v2
		contactSimCapacity = contactSimCapacity + (*b2GraphColor)(unsafe.Pointer(c)).ContactSims.Capacity
		jointSimCapacity = jointSimCapacity + (*b2GraphColor)(unsafe.Pointer(c)).JointSims.Capacity
		goto _26
	_26:
		;
		i4 = i4 + 1
	}
	graphJointSimBytes = jointSimCapacity * int32FromInt64(252)
	graphContactSimBytes = contactSimCapacity * int32FromInt64(200)
	total = total + (bodyBitSetBytes + graphJointSimBytes + graphContactSimBytes)
	fprintf(tls, file, __ccgo_ts+19352, 0)
	fprintf(tls, file, __ccgo_ts+19370, vaList(bp+8, bodyBitSetBytes))
	fprintf(tls, file, __ccgo_ts+19304, vaList(bp+8, graphJointSimBytes))
	fprintf(tls, file, __ccgo_ts+19319, vaList(bp+8, graphContactSimBytes))
	fprintf(tls, file, __ccgo_ts+18981, 0)
	// Per worker task storage and its bit sets
	taskContextBytes = (*b2World)(unsafe.Pointer(world)).TaskContexts.Capacity * int32FromInt64(136)
	i5 = 0
	for {
		if !(i5 < (*b2World)(unsafe.Pointer(world)).TaskContexts.Count) {
			break
		}
		taskContext = (*b2World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(i5)*136
		taskContextBytes = taskContextBytes + (*b2TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Capacity*int32FromInt64(8)
		taskContextBytes = taskContextBytes + (*b2TaskContext)(unsafe.Pointer(taskContext)).PairKeys.Capacity*int32FromInt64(8)
		v2 = int32FromUint64(uint64((*b2BitSet)(unsafe.Pointer(taskContext+32)).BlockCapacity) * uint64(8))
		goto _31
	_31:
		taskContextBytes = taskContextBytes + v2
		v2 = int32FromUint64(uint64((*b2BitSet)(unsafe.Pointer(taskContext+48)).BlockCapacity) * uint64(8))
		goto _33
	_33:
		taskContextBytes = taskContextBytes + v2
		v2 = int32FromUint64(uint64((*b2BitSet)(unsafe.Pointer(taskContext+72)).BlockCapacity) * uint64(8))
		goto _35
	_35:
		taskContextBytes = taskContextBytes + v2
		v2 = int32FromUint64(uint64((*b2BitSet)(unsafe.Pointer(taskContext+88)).BlockCapacity) * uint64(8))
		goto _37
	_37:
		taskContextBytes = taskContextBytes + v2
		v2 = int32FromUint64(uint64((*b2BitSet)(unsafe.Pointer(taskContext+104)).BlockCapacity) * uint64(8))
		goto _39
	_39:
		taskContextBytes = taskContextBytes + v2
		goto _29
	_29:
		;
		i5 = i5 + 1
	}
	sensorTaskContextBytes = (*b2World)(unsafe.Pointer(world)).SensorTaskContexts.Capacity * int32FromInt64(16)
	i6 = 0
	for {
		if !(i6 < (*b2World)(unsafe.Pointer(world)).SensorTaskContexts.Count) {
			break
		}
		taskContext1 = (*b2World)(unsafe.Pointer(world)).SensorTaskContexts.Data + uintptr(i6)*16
		v2 = int32FromUint64(uint64((*b2BitSet)(unsafe.Pointer(taskContext1)).BlockCapacity) * uint64(8))
		goto _42
	_42:
		sensorTaskContextBytes = sensorTaskContextBytes + v2
		goto _40
	_40:
		;
		i6 = i6 + 1
	}
	total = total + (taskContextBytes + sensorTaskContextBytes)
	fprintf(tls, file, __ccgo_ts+19389, 0)
	fprintf(tls, file, __ccgo_ts+19404, vaList(bp+8, taskContextBytes))
	fprintf(tls, file, __ccgo_ts+19416, vaList(bp+8, sensorTaskContextBytes))
	fprintf(tls, file, __ccgo_ts+18981, 0)
	// Double buffered event arrays
	eventBytes = 0
	eventBytes = eventBytes + (*b2World)(unsafe.Pointer(world)).BodyMoveEvents.Capacity*int32FromInt64(40)
	eventBytes = eventBytes + (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity*int32FromInt64(16)
	eventBytes = eventBytes + (*b2World)(unsafe.Pointer(world)).ContactBeginEvents.Capacity*int32FromInt64(28)
	eventBytes = eventBytes + (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304))).Capacity*int32FromInt64(16)
	eventBytes = eventBytes + (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+1*16))).Capacity*int32FromInt64(16)
	eventBytes = eventBytes + (*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336))).Capacity*int32FromInt64(28)
	eventBytes = eventBytes + (*(*b2DynamicArray_b2ContactEndTouchEvent)(unsafe.Add(nil, world+2336+1*16))).Capacity*int32FromInt64(28)
	eventBytes = eventBytes + (*b2World)(unsafe.Pointer(world)).ContactHitEvents.Capacity*int32FromInt64(48)
	eventBytes = eventBytes + (*b2World)(unsafe.Pointer(world)).JointEvents.Capacity*int32FromInt64(16)
	total = total + eventBytes
	fprintf(tls, file, __ccgo_ts+19428, vaList(bp+8, eventBytes))
	// Debug draw bit sets
	debugBytes = 0
	v2 = int32FromUint64(uint64((*b2BitSet)(unsafe.Pointer(world+2408)).BlockCapacity) * uint64(8))
	goto _44
_44:
	debugBytes = debugBytes + v2
	v2 = int32FromUint64(uint64((*b2BitSet)(unsafe.Pointer(world+2424)).BlockCapacity) * uint64(8))
	goto _46
_46:
	debugBytes = debugBytes + v2
	v2 = int32FromUint64(uint64((*b2BitSet)(unsafe.Pointer(world+2440)).BlockCapacity) * uint64(8))
	goto _48
_48:
	debugBytes = debugBytes + v2
	v2 = int32FromUint64(uint64((*b2BitSet)(unsafe.Pointer(world+2456)).BlockCapacity) * uint64(8))
	goto _50
_50:
	debugBytes = debugBytes + v2
	total = total + debugBytes
	fprintf(tls, file, __ccgo_ts+19441, vaList(bp+8, debugBytes))
	// stack allocator
	total = total + (*b2World)(unsafe.Pointer(world)).Stack.Capacity
	fprintf(tls, file, __ccgo_ts+19458, vaList(bp+8, (*b2World)(unsafe.Pointer(world)).Stack.Capacity))
	fprintf(tls, file, __ccgo_ts+19480, vaList(bp+8, total))
	fclose(tls, file)
}

func b2World_OverlapAABB(tls *_Stack, worldId WorldId, origin1 Pos, aabb AABB, filter QueryFilter, __ccgo_fp_fcn uintptr, context uintptr) (r TreeStats) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var i int32
	var result, worldBox, v14, v4 AABB
	var treeResult, treeStats TreeStats
	var world uintptr
	var v1 bool
	var v10, v12, v6, v8 float32
	var v5 Pos
	var _ /* recWriter at bp+0 */ b2RecQueryWriter
	var _ /* worldContext at bp+48 */ WorldQueryContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = i, result, treeResult, treeStats, world, worldBox, v1, v10, v12, v14, v4, v5, v6, v8
	treeStats = TreeStats{}
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(2409))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return treeStats
	}
	if v1 = !!(b2IsValidPos(tls, origin1) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+19491, __ccgo_ts+16660, int32FromInt32(2415))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b2IsValidAABB(tls, aabb) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+19514, __ccgo_ts+16660, int32FromInt32(2416))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*b2RecQueryWriter)(unsafe.Pointer(bp)) = b2RecQueryWriter{}
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b2RecQueryBegin(tls, bp, context)
		*(*uintptr)(unsafe.Pointer(bp)) = __ccgo_fp_fcn
		b2RecW_WORLDID(tls, bp+16, worldId)
		b2RecW_POSITION(tls, bp+16, origin1)
		b2RecW_AABB(tls, bp+16, aabb)
		b2RecW_QUERYFILTER(tls, bp+16, filter)
		(*(*b2RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset = b2RecReserveU32(tls, bp+16)
		__ccgo_fp_fcn = __ccgo_fp(b2RecOverlapTrampoline)
		context = bp
	}
	v4 = aabb
	v5 = origin1
	v6 = float32(float64(v5.X) + float64(v4.LowerBound.X))
	goto _7
_7:
	result.LowerBound.X = v6
	v8 = float32(float64(v5.Y) + float64(v4.LowerBound.Y))
	goto _9
_9:
	result.LowerBound.Y = v8
	v10 = float32(float64(v5.X) + float64(v4.UpperBound.X))
	goto _11
_11:
	result.UpperBound.X = v10
	v12 = float32(float64(v5.Y) + float64(v4.UpperBound.Y))
	goto _13
_13:
	result.UpperBound.Y = v12
	v14 = result
	goto _15
_15:
	// Lift to a world float box with outward rounding so the conservative tree test never misses
	worldBox = v14
	*(*WorldQueryContext)(unsafe.Add(nil, bp+48)) = WorldQueryContext{
		World:       world,
		Fcn:         __ccgo_fp_fcn,
		Filter:      filter,
		UserContext: context,
	}
	i = 0
	for {
		if !(i < int32(b2_bodyTypeCount)) {
			break
		}
		treeResult = b2DynamicTree_OverlapAABB(tls, world+40+uintptr(i)*104, worldBox, filter.MaskBits, __ccgo_fp(TreeQueryCallback), bp+48)
		treeStats.NodeVisits += treeResult.NodeVisits
		treeStats.LeafVisits += treeResult.LeafVisits
		goto _16
	_16:
		;
		i = i + 1
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b2RecPatchU32(tls, bp+16, (*(*b2RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset, (*(*b2RecQueryWriter)(unsafe.Add(nil, bp))).HitCount)
		b2RecW_TREESTATS(tls, bp+16, treeStats)
		b2RecQueryCommit(tls, (*b2World)(unsafe.Pointer(world)).Recording, uint8(0xE0), bp)
	}
	return treeStats
}

func b2World_OverlapShape(tls *_Stack, worldId WorldId, origin1 Pos, proxy uintptr, filter QueryFilter, __ccgo_fp_fcn uintptr, context uintptr) (r1 TreeStats) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var a6, aabb, result, v44, v46, v56 AABB
	var c, c1, r, v20, v22, v23, v34, v36, v37, v38, v40, v41, v42, v8, v9 Vec2
	var i, i1, v4 int32
	var treeResult, treeStats TreeStats
	var world, v3 uintptr
	var v1 bool
	var v10, v11, v12, v14, v15, v16, v17, v19, v24, v25, v26, v28, v29, v30, v31, v33, v48, v5, v50, v52, v54 float32
	var v47 Pos
	var _ /* recWriter at bp+0 */ b2RecQueryWriter
	var _ /* worldContext at bp+48 */ WorldOverlapContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a6, aabb, c, c1, i, i1, r, result, treeResult, treeStats, world, v1, v10, v11, v12, v14, v15, v16, v17, v19, v20, v22, v23, v24, v25, v26, v28, v29, v3, v30, v31, v33, v34, v36, v37, v38, v4, v40, v41, v42, v44, v46, v47, v48, v5, v50, v52, v54, v56, v8, v9
	treeStats = TreeStats{}
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(2512))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return treeStats
	}
	if v1 = !!(b2IsValidPos(tls, origin1) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+19491, __ccgo_ts+16660, int32FromInt32(2518))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*b2RecQueryWriter)(unsafe.Pointer(bp)) = b2RecQueryWriter{}
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b2RecQueryBegin(tls, bp, context)
		*(*uintptr)(unsafe.Pointer(bp)) = __ccgo_fp_fcn
		b2RecW_WORLDID(tls, bp+16, worldId)
		b2RecW_POSITION(tls, bp+16, origin1)
		b2RecW_SHAPEPROXY(tls, bp+16, *(*ShapeProxy)(unsafe.Add(nil, proxy)))
		b2RecW_QUERYFILTER(tls, bp+16, filter)
		(*(*b2RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset = b2RecReserveU32(tls, bp+16)
		__ccgo_fp_fcn = __ccgo_fp(b2RecOverlapTrampoline)
		context = bp
	}
	v3 = proxy
	v4 = (*ShapeProxy)(unsafe.Pointer(proxy)).Count
	v5 = (*ShapeProxy)(unsafe.Pointer(proxy)).Radius
	if v1 = !!(v4 > int32FromInt32(0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+9458, __ccgo_ts+19536, int32FromInt32(708))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	a6 = AABB{
		LowerBound: *(*Vec2)(unsafe.Add(nil, v3)),
		UpperBound: *(*Vec2)(unsafe.Add(nil, v3)),
	}
	i = int32(1)
	for {
		if !(i < v4) {
			break
		}
		v8 = a6.LowerBound
		v9 = *(*Vec2)(unsafe.Add(nil, v3+uintptr(i)*8))
		v10 = v8.X
		v11 = v9.X
		if v10 < v11 {
			v14 = v10
		} else {
			v14 = v11
		}
		v12 = v14
		goto _13
	_13:
		c.X = v12
		v15 = v8.Y
		v16 = v9.Y
		if v15 < v16 {
			v19 = v15
		} else {
			v19 = v16
		}
		v17 = v19
		goto _18
	_18:
		c.Y = v17
		v20 = c
		goto _21
	_21:
		a6.LowerBound = v20
		v22 = a6.UpperBound
		v23 = *(*Vec2)(unsafe.Add(nil, v3+uintptr(i)*8))
		v24 = v22.X
		v25 = v23.X
		if v24 > v25 {
			v28 = v24
		} else {
			v28 = v25
		}
		v26 = v28
		goto _27
	_27:
		c1.X = v26
		v29 = v22.Y
		v30 = v23.Y
		if v29 > v30 {
			v33 = v29
		} else {
			v33 = v30
		}
		v31 = v33
		goto _32
	_32:
		c1.Y = v31
		v34 = c1
		goto _35
	_35:
		a6.UpperBound = v34
		goto _7
	_7:
		;
		i = i + 1
	}
	r = Vec2{
		X: v5,
		Y: v5,
	}
	v36 = a6.LowerBound
	v37 = r
	v38 = Vec2{
		X: v36.X - v37.X,
		Y: v36.Y - v37.Y,
	}
	goto _39
_39:
	a6.LowerBound = v38
	v40 = a6.UpperBound
	v41 = r
	v42 = Vec2{
		X: v40.X + v41.X,
		Y: v40.Y + v41.Y,
	}
	goto _43
_43:
	a6.UpperBound = v42
	v44 = a6
	goto _45
_45:
	v46 = v44
	v47 = origin1
	v48 = float32(float64(v47.X) + float64(v46.LowerBound.X))
	goto _49
_49:
	result.LowerBound.X = v48
	v50 = float32(float64(v47.Y) + float64(v46.LowerBound.Y))
	goto _51
_51:
	result.LowerBound.Y = v50
	v52 = float32(float64(v47.X) + float64(v46.UpperBound.X))
	goto _53
_53:
	result.UpperBound.X = v52
	v54 = float32(float64(v47.Y) + float64(v46.UpperBound.Y))
	goto _55
_55:
	result.UpperBound.Y = v54
	v56 = result
	goto _57
_57:
	// Relative box lifted to world float with outward rounding, conservative for the tree
	aabb = v56
	*(*WorldOverlapContext)(unsafe.Add(nil, bp+48)) = WorldOverlapContext{
		World:       world,
		Fcn:         __ccgo_fp_fcn,
		Filter:      filter,
		Proxy:       proxy,
		Origin:      origin1,
		UserContext: context,
	}
	i1 = 0
	for {
		if !(i1 < int32(b2_bodyTypeCount)) {
			break
		}
		treeResult = b2DynamicTree_OverlapAABB(tls, world+40+uintptr(i1)*104, aabb, filter.MaskBits, __ccgo_fp(TreeOverlapCallback), bp+48)
		treeStats.NodeVisits += treeResult.NodeVisits
		treeStats.LeafVisits += treeResult.LeafVisits
		goto _58
	_58:
		;
		i1 = i1 + 1
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b2RecPatchU32(tls, bp+16, (*(*b2RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset, (*(*b2RecQueryWriter)(unsafe.Add(nil, bp))).HitCount)
		b2RecW_TREESTATS(tls, bp+16, treeStats)
		b2RecQueryCommit(tls, (*b2World)(unsafe.Pointer(world)).Recording, uint8(0xE1), bp)
	}
	return treeStats
}

func b2World_CastRay(tls *_Stack, worldId WorldId, origin Pos, translation Vec2, filter QueryFilter, __ccgo_fp_fcn uintptr, context uintptr) (r TreeStats) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var i int32
	var treeResult, treeStats TreeStats
	var world uintptr
	var v1 bool
	var v4 Pos
	var v5 Vec2
	var _ /* input at bp+48 */ RayCastInput
	var _ /* recWriter at bp+0 */ b2RecQueryWriter
	var _ /* worldContext at bp+72 */ WorldRayCastContext
	_, _, _, _, _, _, _ = i, treeResult, treeStats, world, v1, v4, v5
	treeStats = TreeStats{}
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(2617))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return treeStats
	}
	if v1 = !!(b2IsValidPos(tls, origin) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+19491, __ccgo_ts+16660, int32FromInt32(2623))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b2IsValidVec2(tls, translation) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+19578, __ccgo_ts+16660, int32FromInt32(2624))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*b2RecQueryWriter)(unsafe.Pointer(bp)) = b2RecQueryWriter{}
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b2RecQueryBegin(tls, bp, context)
		*(*uintptr)(unsafe.Pointer(bp)) = __ccgo_fp_fcn
		b2RecW_WORLDID(tls, bp+16, worldId)
		b2RecW_POSITION(tls, bp+16, origin)
		b2RecW_VEC2(tls, bp+16, translation)
		b2RecW_QUERYFILTER(tls, bp+16, filter)
		(*(*b2RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset = b2RecReserveU32(tls, bp+16)
		__ccgo_fp_fcn = __ccgo_fp(b2RecCastTrampoline)
		context = bp
	}
	v4 = origin
	v5 = Vec2{
		X: v4.X,
		Y: v4.Y,
	}
	goto _6
_6:
	// Tree traversal sees the origin truncated to float, displacing the ray by up to one
	// coordinate ULP, a graze sized miss tolerance at extreme range. Per-shape casts
	// re-difference against the full precision origin carried on the context.
	*(*RayCastInput)(unsafe.Add(nil, bp+48)) = RayCastInput{
		Origin:      v5,
		Translation: translation,
		MaxFraction: float32FromFloat32(1),
	}
	*(*WorldRayCastContext)(unsafe.Add(nil, bp+72)) = WorldRayCastContext{
		World:       world,
		Fcn:         __ccgo_fp_fcn,
		Filter:      filter,
		Fraction:    float32FromFloat32(1),
		Origin:      origin,
		UserContext: context,
	}
	i = 0
	for {
		if !(i < int32(b2_bodyTypeCount)) {
			break
		}
		treeResult = b2DynamicTree_CastRay(tls, world+40+uintptr(i)*104, bp+48, filter.MaskBits, __ccgo_fp(RayCastCallback), bp+72)
		treeStats.NodeVisits += treeResult.NodeVisits
		treeStats.LeafVisits += treeResult.LeafVisits
		if (*(*WorldRayCastContext)(unsafe.Add(nil, bp+72))).Fraction == float32FromFloat32(0) {
			break
		}
		(*(*RayCastInput)(unsafe.Add(nil, bp+48))).MaxFraction = (*(*WorldRayCastContext)(unsafe.Add(nil, bp+72))).Fraction
		goto _7
	_7:
		;
		i = i + 1
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b2RecPatchU32(tls, bp+16, (*(*b2RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset, (*(*b2RecQueryWriter)(unsafe.Add(nil, bp))).HitCount)
		b2RecW_TREESTATS(tls, bp+16, treeStats)
		b2RecQueryCommit(tls, (*b2World)(unsafe.Pointer(world)).Recording, uint8(0xE2), bp)
	}
	return treeStats
}

func b2World_CastRayClosest(tls *_Stack, worldId WorldId, origin Pos, translation Vec2, filter QueryFilter) (r RayResult) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var i int32
	var treeResult TreeStats
	var world uintptr
	var v1 bool
	var v4 Pos
	var v5 Vec2
	var _ /* input at bp+40 */ RayCastInput
	var _ /* recBuf at bp+120 */ b2RecBuffer
	var _ /* result at bp+0 */ RayResult
	var _ /* worldContext at bp+64 */ WorldRayCastContext
	_, _, _, _, _, _ = i, treeResult, world, v1, v4, v5
	*(*RayResult)(unsafe.Add(nil, bp)) = RayResult{}
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(2695))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return *(*RayResult)(unsafe.Add(nil, bp))
	}
	if v1 = !!(b2IsValidPos(tls, origin) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+19491, __ccgo_ts+16660, int32FromInt32(2701))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b2IsValidVec2(tls, translation) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+19578, __ccgo_ts+16660, int32FromInt32(2702))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	v4 = origin
	v5 = Vec2{
		X: v4.X,
		Y: v4.Y,
	}
	goto _6
_6:
	*(*RayCastInput)(unsafe.Add(nil, bp+40)) = RayCastInput{
		Origin:      v5,
		Translation: translation,
		MaxFraction: float32FromFloat32(1),
	}
	*(*WorldRayCastContext)(unsafe.Add(nil, bp+64)) = WorldRayCastContext{
		World:       world,
		Fcn:         __ccgo_fp(b2RayCastClosestFcn),
		Filter:      filter,
		Fraction:    float32FromFloat32(1),
		Origin:      origin,
		UserContext: bp,
	}
	i = 0
	for {
		if !(i < int32(b2_bodyTypeCount)) {
			break
		}
		treeResult = b2DynamicTree_CastRay(tls, world+40+uintptr(i)*104, bp+40, filter.MaskBits, __ccgo_fp(RayCastCallback), bp+64)
		(*(*RayResult)(unsafe.Add(nil, bp))).NodeVisits += treeResult.NodeVisits
		(*(*RayResult)(unsafe.Add(nil, bp))).LeafVisits += treeResult.LeafVisits
		if (*(*WorldRayCastContext)(unsafe.Add(nil, bp+64))).Fraction == float32FromFloat32(0) {
			break
		}
		(*(*RayCastInput)(unsafe.Add(nil, bp+40))).MaxFraction = (*(*WorldRayCastContext)(unsafe.Add(nil, bp+64))).Fraction
		goto _7
	_7:
		;
		i = i + 1
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecBuffer)(unsafe.Add(nil, bp+120)) = b2RecBuffer{}
		b2RecW_WORLDID(tls, bp+120, worldId)
		b2RecW_POSITION(tls, bp+120, origin)
		b2RecW_VEC2(tls, bp+120, translation)
		b2RecW_QUERYFILTER(tls, bp+120, filter)
		b2RecW_RAYRESULT(tls, bp+120, *(*RayResult)(unsafe.Add(nil, bp)))
		b2RecCommitRecord(tls, (*b2World)(unsafe.Pointer(world)).Recording, uint8(0xE5), (*(*b2RecBuffer)(unsafe.Add(nil, bp+120))).Data, (*(*b2RecBuffer)(unsafe.Add(nil, bp+120))).Size)
		b2RecBufFree(tls, bp+120)
	}
	return *(*RayResult)(unsafe.Add(nil, bp))
}

func b2World_CastShape(tls *_Stack, worldId WorldId, origin1 Pos, proxy uintptr, translation Vec2, filter QueryFilter, __ccgo_fp_fcn uintptr, context uintptr) (r1 TreeStats) {
	bp := tls.Alloc(224)
	defer tls.Free(224)
	var a6, box1, localBox, result, v45, v47, v57 AABB
	var c, c1, r, v10, v21, v23, v24, v35, v37, v38, v39, v41, v42, v43, v9 Vec2
	var i, i1, v5 int32
	var treeResult, treeStats TreeStats
	var world, v4 uintptr
	var v1 bool
	var v11, v12, v13, v15, v16, v17, v18, v20, v25, v26, v27, v29, v30, v31, v32, v34, v49, v51, v53, v55, v6 float32
	var v48 Pos
	var _ /* recWriter at bp+0 */ b2RecQueryWriter
	var _ /* treeInput at bp+192 */ BoxCastInput
	var _ /* worldContext at bp+48 */ WorldShapeCastContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a6, box1, c, c1, i, i1, localBox, r, result, treeResult, treeStats, world, v1, v10, v11, v12, v13, v15, v16, v17, v18, v20, v21, v23, v24, v25, v26, v27, v29, v30, v31, v32, v34, v35, v37, v38, v39, v4, v41, v42, v43, v45, v47, v48, v49, v5, v51, v53, v55, v57, v6, v9
	treeStats = TreeStats{}
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(2801))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return treeStats
	}
	if v1 = !!(b2IsValidPos(tls, origin1) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+19491, __ccgo_ts+16660, int32FromInt32(2807))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b2IsValidVec2(tls, translation) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+19578, __ccgo_ts+16660, int32FromInt32(2808))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*b2RecQueryWriter)(unsafe.Pointer(bp)) = b2RecQueryWriter{}
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b2RecQueryBegin(tls, bp, context)
		*(*uintptr)(unsafe.Pointer(bp)) = __ccgo_fp_fcn
		b2RecW_WORLDID(tls, bp+16, worldId)
		b2RecW_POSITION(tls, bp+16, origin1)
		b2RecW_SHAPEPROXY(tls, bp+16, *(*ShapeProxy)(unsafe.Add(nil, proxy)))
		b2RecW_VEC2(tls, bp+16, translation)
		b2RecW_QUERYFILTER(tls, bp+16, filter)
		(*(*b2RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset = b2RecReserveU32(tls, bp+16)
		__ccgo_fp_fcn = __ccgo_fp(b2RecCastTrampoline)
		context = bp
	}
	*(*WorldShapeCastContext)(unsafe.Add(nil, bp+48)) = WorldShapeCastContext{}
	(*(*WorldShapeCastContext)(unsafe.Add(nil, bp+48))).World = world
	(*(*WorldShapeCastContext)(unsafe.Add(nil, bp+48))).Fcn = __ccgo_fp_fcn
	(*(*WorldShapeCastContext)(unsafe.Add(nil, bp+48))).Filter = filter
	(*(*WorldShapeCastContext)(unsafe.Add(nil, bp+48))).Fraction = float32FromFloat32(1)
	(*(*WorldShapeCastContext)(unsafe.Add(nil, bp+48))).Origin = origin1
	(*(*WorldShapeCastContext)(unsafe.Add(nil, bp+48))).Input.Proxy = *(*ShapeProxy)(unsafe.Add(nil, proxy))
	(*(*WorldShapeCastContext)(unsafe.Add(nil, bp+48))).Input.Translation = translation
	(*(*WorldShapeCastContext)(unsafe.Add(nil, bp+48))).Input.MaxFraction = float32FromFloat32(1)
	(*(*WorldShapeCastContext)(unsafe.Add(nil, bp+48))).UserContext = context
	v4 = proxy
	v5 = (*ShapeProxy)(unsafe.Pointer(proxy)).Count
	v6 = (*ShapeProxy)(unsafe.Pointer(proxy)).Radius
	if v1 = !!(v5 > int32FromInt32(0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+9458, __ccgo_ts+19536, int32FromInt32(708))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	a6 = AABB{
		LowerBound: *(*Vec2)(unsafe.Add(nil, v4)),
		UpperBound: *(*Vec2)(unsafe.Add(nil, v4)),
	}
	i = int32(1)
	for {
		if !(i < v5) {
			break
		}
		v9 = a6.LowerBound
		v10 = *(*Vec2)(unsafe.Add(nil, v4+uintptr(i)*8))
		v11 = v9.X
		v12 = v10.X
		if v11 < v12 {
			v15 = v11
		} else {
			v15 = v12
		}
		v13 = v15
		goto _14
	_14:
		c.X = v13
		v16 = v9.Y
		v17 = v10.Y
		if v16 < v17 {
			v20 = v16
		} else {
			v20 = v17
		}
		v18 = v20
		goto _19
	_19:
		c.Y = v18
		v21 = c
		goto _22
	_22:
		a6.LowerBound = v21
		v23 = a6.UpperBound
		v24 = *(*Vec2)(unsafe.Add(nil, v4+uintptr(i)*8))
		v25 = v23.X
		v26 = v24.X
		if v25 > v26 {
			v29 = v25
		} else {
			v29 = v26
		}
		v27 = v29
		goto _28
	_28:
		c1.X = v27
		v30 = v23.Y
		v31 = v24.Y
		if v30 > v31 {
			v34 = v30
		} else {
			v34 = v31
		}
		v32 = v34
		goto _33
	_33:
		c1.Y = v32
		v35 = c1
		goto _36
	_36:
		a6.UpperBound = v35
		goto _8
	_8:
		;
		i = i + 1
	}
	r = Vec2{
		X: v6,
		Y: v6,
	}
	v37 = a6.LowerBound
	v38 = r
	v39 = Vec2{
		X: v37.X - v38.X,
		Y: v37.Y - v38.Y,
	}
	goto _40
_40:
	a6.LowerBound = v39
	v41 = a6.UpperBound
	v42 = r
	v43 = Vec2{
		X: v41.X + v42.X,
		Y: v41.Y + v42.Y,
	}
	goto _44
_44:
	a6.UpperBound = v43
	v45 = a6
	goto _46
_46:
	// Create a conservative world space box from the proxy and origin.
	localBox = v45
	v47 = localBox
	v48 = origin1
	v49 = float32(float64(v48.X) + float64(v47.LowerBound.X))
	goto _50
_50:
	result.LowerBound.X = v49
	v51 = float32(float64(v48.Y) + float64(v47.LowerBound.Y))
	goto _52
_52:
	result.LowerBound.Y = v51
	v53 = float32(float64(v48.X) + float64(v47.UpperBound.X))
	goto _54
_54:
	result.UpperBound.X = v53
	v55 = float32(float64(v48.Y) + float64(v47.UpperBound.Y))
	goto _56
_56:
	result.UpperBound.Y = v55
	v57 = result
	goto _58
_58:
	box1 = v57
	*(*BoxCastInput)(unsafe.Add(nil, bp+192)) = BoxCastInput{
		Box:         box1,
		Translation: translation,
		MaxFraction: float32FromFloat32(1),
	}
	i1 = 0
	for {
		if !(i1 < int32(b2_bodyTypeCount)) {
			break
		}
		treeResult = b2DynamicTree_CastBox(tls, world+40+uintptr(i1)*104, bp+192, filter.MaskBits, __ccgo_fp(ShapeCastCallback), bp+48)
		treeStats.NodeVisits += treeResult.NodeVisits
		treeStats.LeafVisits += treeResult.LeafVisits
		if (*(*WorldShapeCastContext)(unsafe.Add(nil, bp+48))).Fraction == float32FromFloat32(0) {
			break
		}
		(*(*BoxCastInput)(unsafe.Add(nil, bp+192))).MaxFraction = (*(*WorldShapeCastContext)(unsafe.Add(nil, bp+48))).Fraction
		goto _59
	_59:
		;
		i1 = i1 + 1
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b2RecPatchU32(tls, bp+16, (*(*b2RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset, (*(*b2RecQueryWriter)(unsafe.Add(nil, bp))).HitCount)
		b2RecW_TREESTATS(tls, bp+16, treeStats)
		b2RecQueryCommit(tls, (*b2World)(unsafe.Pointer(world)).Recording, uint8(0xE3), bp)
	}
	return treeStats
}

func b2World_CastMover(tls *_Stack, worldId WorldId, origin1 Pos, mover uintptr, translation Vec2, filter QueryFilter) (r1 float32) {
	bp := tls.Alloc(208)
	defer tls.Free(208)
	var a6, box1, result, v46, v48, v58 AABB
	var c, c1, r, v10, v11, v22, v24, v25, v36, v38, v39, v40, v42, v43, v44 Vec2
	var i, i1, v6 int32
	var world, v5 uintptr
	var v1 bool
	var v12, v13, v14, v16, v17, v18, v19, v21, v26, v27, v28, v30, v31, v32, v33, v35, v50, v52, v54, v56, v7 float32
	var v49 Pos
	var _ /* centers at bp+128 */ [2]Vec2
	var _ /* recBuf at bp+176 */ b2RecBuffer
	var _ /* treeInput at bp+144 */ BoxCastInput
	var _ /* worldContext at bp+0 */ WorldMoverCastContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a6, box1, c, c1, i, i1, r, result, world, v1, v10, v11, v12, v13, v14, v16, v17, v18, v19, v21, v22, v24, v25, v26, v27, v28, v30, v31, v32, v33, v35, v36, v38, v39, v40, v42, v43, v44, v46, v48, v49, v5, v50, v52, v54, v56, v58, v6, v7
	if v1 = !!(b2IsValidPos(tls, origin1) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+19491, __ccgo_ts+16660, int32FromInt32(2912))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b2IsValidVec2(tls, translation) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+19578, __ccgo_ts+16660, int32FromInt32(2913))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*Capsule)(unsafe.Pointer(mover)).Radius > float32(float32FromFloat32(2)*float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))); !v1 {
		b2InternalAssert(tls, __ccgo_ts+19607, __ccgo_ts+16660, int32FromInt32(2914))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(2917))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return float32FromFloat32(1)
	}
	*(*WorldMoverCastContext)(unsafe.Add(nil, bp)) = WorldMoverCastContext{}
	(*(*WorldMoverCastContext)(unsafe.Add(nil, bp))).World = world
	(*(*WorldMoverCastContext)(unsafe.Add(nil, bp))).Filter = filter
	(*(*WorldMoverCastContext)(unsafe.Add(nil, bp))).Fraction = float32FromFloat32(1)
	(*(*WorldMoverCastContext)(unsafe.Add(nil, bp))).Origin = origin1
	*(*Vec2)(unsafe.Add(nil, bp+36)) = (*Capsule)(unsafe.Pointer(mover)).Center1
	*(*Vec2)(unsafe.Add(nil, bp+36+1*8)) = (*Capsule)(unsafe.Pointer(mover)).Center2
	(*(*WorldMoverCastContext)(unsafe.Add(nil, bp))).Input.Proxy.Count = int32(2)
	(*(*WorldMoverCastContext)(unsafe.Add(nil, bp))).Input.Proxy.Radius = (*Capsule)(unsafe.Pointer(mover)).Radius
	(*(*WorldMoverCastContext)(unsafe.Add(nil, bp))).Input.Translation = translation
	(*(*WorldMoverCastContext)(unsafe.Add(nil, bp))).Input.MaxFraction = float32FromFloat32(1)
	(*(*WorldMoverCastContext)(unsafe.Add(nil, bp))).Input.CanEncroach = boolUint8(true1 != 0)
	// Bound the capsule in origin relative space then lift to a conservative world float box
	*(*[2]Vec2)(unsafe.Add(nil, bp+128)) = [2]Vec2{
		0: (*Capsule)(unsafe.Pointer(mover)).Center1,
		1: (*Capsule)(unsafe.Pointer(mover)).Center2,
	}
	v5 = bp + 128
	v6 = int32(2)
	v7 = (*Capsule)(unsafe.Pointer(mover)).Radius
	if v1 = !!(v6 > int32FromInt32(0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+9458, __ccgo_ts+19536, int32FromInt32(708))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	a6 = AABB{
		LowerBound: *(*Vec2)(unsafe.Add(nil, v5)),
		UpperBound: *(*Vec2)(unsafe.Add(nil, v5)),
	}
	i = int32(1)
	for {
		if !(i < v6) {
			break
		}
		v10 = a6.LowerBound
		v11 = *(*Vec2)(unsafe.Add(nil, v5+uintptr(i)*8))
		v12 = v10.X
		v13 = v11.X
		if v12 < v13 {
			v16 = v12
		} else {
			v16 = v13
		}
		v14 = v16
		goto _15
	_15:
		c.X = v14
		v17 = v10.Y
		v18 = v11.Y
		if v17 < v18 {
			v21 = v17
		} else {
			v21 = v18
		}
		v19 = v21
		goto _20
	_20:
		c.Y = v19
		v22 = c
		goto _23
	_23:
		a6.LowerBound = v22
		v24 = a6.UpperBound
		v25 = *(*Vec2)(unsafe.Add(nil, v5+uintptr(i)*8))
		v26 = v24.X
		v27 = v25.X
		if v26 > v27 {
			v30 = v26
		} else {
			v30 = v27
		}
		v28 = v30
		goto _29
	_29:
		c1.X = v28
		v31 = v24.Y
		v32 = v25.Y
		if v31 > v32 {
			v35 = v31
		} else {
			v35 = v32
		}
		v33 = v35
		goto _34
	_34:
		c1.Y = v33
		v36 = c1
		goto _37
	_37:
		a6.UpperBound = v36
		goto _9
	_9:
		;
		i = i + 1
	}
	r = Vec2{
		X: v7,
		Y: v7,
	}
	v38 = a6.LowerBound
	v39 = r
	v40 = Vec2{
		X: v38.X - v39.X,
		Y: v38.Y - v39.Y,
	}
	goto _41
_41:
	a6.LowerBound = v40
	v42 = a6.UpperBound
	v43 = r
	v44 = Vec2{
		X: v42.X + v43.X,
		Y: v42.Y + v43.Y,
	}
	goto _45
_45:
	a6.UpperBound = v44
	v46 = a6
	goto _47
_47:
	v48 = v46
	v49 = origin1
	v50 = float32(float64(v49.X) + float64(v48.LowerBound.X))
	goto _51
_51:
	result.LowerBound.X = v50
	v52 = float32(float64(v49.Y) + float64(v48.LowerBound.Y))
	goto _53
_53:
	result.LowerBound.Y = v52
	v54 = float32(float64(v49.X) + float64(v48.UpperBound.X))
	goto _55
_55:
	result.UpperBound.X = v54
	v56 = float32(float64(v49.Y) + float64(v48.UpperBound.Y))
	goto _57
_57:
	result.UpperBound.Y = v56
	v58 = result
	goto _59
_59:
	box1 = v58
	*(*BoxCastInput)(unsafe.Add(nil, bp+144)) = BoxCastInput{
		Box:         box1,
		Translation: translation,
		MaxFraction: float32FromFloat32(1),
	}
	i1 = 0
	for {
		if !(i1 < int32(b2_bodyTypeCount)) {
			break
		}
		b2DynamicTree_CastBox(tls, world+40+uintptr(i1)*104, bp+144, filter.MaskBits, __ccgo_fp(MoverCastCallback), bp)
		if (*(*WorldMoverCastContext)(unsafe.Add(nil, bp))).Fraction == float32FromFloat32(0) {
			break
		}
		(*(*BoxCastInput)(unsafe.Add(nil, bp+144))).MaxFraction = (*(*WorldMoverCastContext)(unsafe.Add(nil, bp))).Fraction
		goto _60
	_60:
		;
		i1 = i1 + 1
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecBuffer)(unsafe.Add(nil, bp+176)) = b2RecBuffer{}
		b2RecW_WORLDID(tls, bp+176, worldId)
		b2RecW_POSITION(tls, bp+176, origin1)
		b2RecW_CAPSULE(tls, bp+176, *(*Capsule)(unsafe.Add(nil, mover)))
		b2RecW_VEC2(tls, bp+176, translation)
		b2RecW_QUERYFILTER(tls, bp+176, filter)
		b2RecW_F32(tls, bp+176, (*(*WorldMoverCastContext)(unsafe.Add(nil, bp))).Fraction)
		b2RecCommitRecord(tls, (*b2World)(unsafe.Pointer(world)).Recording, uint8(0xE6), (*(*b2RecBuffer)(unsafe.Add(nil, bp+176))).Data, (*(*b2RecBuffer)(unsafe.Add(nil, bp+176))).Size)
		b2RecBufFree(tls, bp+176)
	}
	return (*(*WorldMoverCastContext)(unsafe.Add(nil, bp))).Fraction
}

// C documentation
//
//	// It is tempting to use a shape proxy for the mover, but this makes handling deep overlap difficult and the generality may
//	// not be worth it.
func b2World_CollideMover(tls *_Stack, worldId WorldId, origin1 Pos, mover uintptr, filter QueryFilter, __ccgo_fp_fcn uintptr, context uintptr) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var aabb, relBox, result, v39, v49 AABB
	var c, c1, r, v15, v17, v18, v19, v3, v4 Vec2
	var i int32
	var world uintptr
	var v1 bool
	var v10, v11, v12, v14, v5, v6, v7, v9 float32
	var v40 Pos
	var _ /* recWriter at bp+0 */ b2RecQueryWriter
	var _ /* worldContext at bp+48 */ WorldMoverContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, c, c1, i, r, relBox, result, world, v1, v10, v11, v12, v14, v15, v17, v18, v19, v3, v39, v4, v40, v49, v5, v6, v7, v9
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(3016))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if v1 = !!(b2IsValidPos(tls, origin1) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+19491, __ccgo_ts+16660, int32FromInt32(3022))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*b2RecQueryWriter)(unsafe.Pointer(bp)) = b2RecQueryWriter{}
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b2RecQueryBegin(tls, bp, context)
		*(*uintptr)(unsafe.Pointer(bp)) = __ccgo_fp_fcn
		b2RecW_WORLDID(tls, bp+16, worldId)
		b2RecW_POSITION(tls, bp+16, origin1)
		b2RecW_CAPSULE(tls, bp+16, *(*Capsule)(unsafe.Add(nil, mover)))
		b2RecW_QUERYFILTER(tls, bp+16, filter)
		(*(*b2RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset = b2RecReserveU32(tls, bp+16)
		__ccgo_fp_fcn = __ccgo_fp(b2RecPlaneTrampoline)
		context = bp
	}
	r = Vec2{
		X: (*Capsule)(unsafe.Pointer(mover)).Radius,
		Y: (*Capsule)(unsafe.Pointer(mover)).Radius,
	}
	v3 = (*Capsule)(unsafe.Pointer(mover)).Center1
	v4 = (*Capsule)(unsafe.Pointer(mover)).Center2
	v5 = v3.X
	v6 = v4.X
	if v5 < v6 {
		v9 = v5
	} else {
		v9 = v6
	}
	v7 = v9
	goto _8
_8:
	c.X = v7
	v10 = v3.Y
	v11 = v4.Y
	if v10 < v11 {
		v14 = v10
	} else {
		v14 = v11
	}
	v12 = v14
	goto _13
_13:
	c.Y = v12
	v15 = c
	goto _16
_16:
	v17 = v15
	v18 = r
	v19 = Vec2{
		X: v17.X - v18.X,
		Y: v17.Y - v18.Y,
	}
	goto _20
_20:
	relBox.LowerBound = v19
	v3 = (*Capsule)(unsafe.Pointer(mover)).Center1
	v4 = (*Capsule)(unsafe.Pointer(mover)).Center2
	v5 = v3.X
	v6 = v4.X
	if v5 > v6 {
		v9 = v5
	} else {
		v9 = v6
	}
	v7 = v9
	goto _26
_26:
	c1.X = v7
	v10 = v3.Y
	v11 = v4.Y
	if v10 > v11 {
		v14 = v10
	} else {
		v14 = v11
	}
	v12 = v14
	goto _31
_31:
	c1.Y = v12
	v15 = c1
	goto _34
_34:
	v17 = v15
	v18 = r
	v19 = Vec2{
		X: v17.X + v18.X,
		Y: v17.Y + v18.Y,
	}
	goto _38
_38:
	relBox.UpperBound = v19
	v39 = relBox
	v40 = origin1
	v5 = float32(float64(v40.X) + float64(v39.LowerBound.X))
	goto _42
_42:
	result.LowerBound.X = v5
	v6 = float32(float64(v40.Y) + float64(v39.LowerBound.Y))
	goto _44
_44:
	result.LowerBound.Y = v6
	v7 = float32(float64(v40.X) + float64(v39.UpperBound.X))
	goto _46
_46:
	result.UpperBound.X = v7
	v9 = float32(float64(v40.Y) + float64(v39.UpperBound.Y))
	goto _48
_48:
	result.UpperBound.Y = v9
	v49 = result
	goto _50
_50:
	aabb = v49
	*(*WorldMoverContext)(unsafe.Add(nil, bp+48)) = WorldMoverContext{
		World:       world,
		Fcn:         __ccgo_fp_fcn,
		Filter:      filter,
		Mover:       *(*Capsule)(unsafe.Add(nil, mover)),
		Origin:      origin1,
		UserContext: context,
	}
	i = 0
	for {
		if !(i < int32(b2_bodyTypeCount)) {
			break
		}
		b2DynamicTree_OverlapAABB(tls, world+40+uintptr(i)*104, aabb, filter.MaskBits, __ccgo_fp(TreeCollideCallback), bp+48)
		goto _51
	_51:
		;
		i = i + 1
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b2RecPatchU32(tls, bp+16, (*(*b2RecQueryWriter)(unsafe.Add(nil, bp))).CountOffset, (*(*b2RecQueryWriter)(unsafe.Add(nil, bp))).HitCount)
		// CollideMover returns void; no TREESTATS tail
		b2RecQueryCommit(tls, (*b2World)(unsafe.Pointer(world)).Recording, uint8(0xE4), bp)
	}
}

func b2World_SetCustomFilterCallback(tls *_Stack, worldId WorldId, __ccgo_fp_fcn uintptr, context uintptr) {
	var world uintptr
	var v1 bool
	_, _ = world, v1
	world = b2GetWorldFromId(tls, worldId)
	if __ccgo_fp_fcn != uintptrFromInt32(0) && (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		printf(tls, __ccgo_ts+19645, 0)
		if v1 = bool(!!(bool(false1 != 0) && bool(__ccgo_ts+19722 != 0))); !v1 {
			b2InternalAssert(tls, __ccgo_ts+19779, __ccgo_ts+16660, int32FromInt32(3069))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	(*b2World)(unsafe.Pointer(world)).CustomFilterFcn = __ccgo_fp_fcn
	(*b2World)(unsafe.Pointer(world)).CustomFilterContext = context
}

func b2World_SetPreSolveCallback(tls *_Stack, worldId WorldId, __ccgo_fp_preSolveFcn uintptr, __ccgo_fp_preContinuousFcn uintptr, context uintptr) {
	var world uintptr
	var v1 bool
	_, _ = world, v1
	world = b2GetWorldFromId(tls, worldId)
	if (__ccgo_fp_preSolveFcn != uintptrFromInt32(0) || __ccgo_fp_preContinuousFcn != uintptrFromInt32(0)) && (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		printf(tls, __ccgo_ts+19847, 0)
		if v1 = bool(!!(bool(false1 != 0) && bool(__ccgo_ts+19916 != 0))); !v1 {
			b2InternalAssert(tls, __ccgo_ts+19969, __ccgo_ts+16660, int32FromInt32(3084))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	(*b2World)(unsafe.Pointer(world)).PreSolveFcn = __ccgo_fp_preSolveFcn
	(*b2World)(unsafe.Pointer(world)).PreContinuousFcn = __ccgo_fp_preContinuousFcn
	(*b2World)(unsafe.Pointer(world)).PreSolveContext = context
}

func b2World_SetGravity(tls *_Stack, worldId WorldId, gravity Vec2) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_WorldSetGravity
	_, _ = world, v1
	if int32FromUint8(b2IsValidVec2(tls, gravity)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+20033, uintptr(unsafe.Pointer(&__func__50))))
		if v1 = !!(b2IsValidVec2(tls, gravity) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+20033, __ccgo_ts+16660, int32FromInt32(3095))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorldFromId(tls, worldId)
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WorldSetGravity)(unsafe.Add(nil, bp)) = b2RecArgs_WorldSetGravity{
			World:   worldId,
			Gravity: gravity,
		}
		b2RecWrite_WorldSetGravity(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	(*b2World)(unsafe.Pointer(world)).Gravity = gravity
}

func b2World_GetGravity(tls *_Stack, worldId WorldId) (r Vec2) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return (*b2World)(unsafe.Pointer(world)).Gravity
}

func b2World_Explode(tls *_Stack, worldId WorldId, explosionDef uintptr) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var aabb, localBox, result, v16, v6 AABB
	var extent, falloff, impulsePerLength, radius, v10, v12, v14, v8 float32
	var maskBits uint64_t
	var position, v7 Pos
	var world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_WorldExplode
	var _ /* explosionContext at bp+40 */ ExplosionContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, extent, falloff, impulsePerLength, localBox, maskBits, position, radius, result, world, v1, v10, v12, v14, v16, v6, v7, v8
	maskBits = (*ExplosionDef)(unsafe.Pointer(explosionDef)).MaskBits
	position = (*ExplosionDef)(unsafe.Pointer(explosionDef)).Position
	radius = (*ExplosionDef)(unsafe.Pointer(explosionDef)).Radius
	falloff = (*ExplosionDef)(unsafe.Pointer(explosionDef)).Falloff
	impulsePerLength = (*ExplosionDef)(unsafe.Pointer(explosionDef)).ImpulsePerLength
	if int32FromUint8(b2IsValidPos(tls, position)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+80, __ccgo_ts+2506, uintptr(unsafe.Pointer(&__func__310))))
		if v1 = !!(b2IsValidPos(tls, position) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2506, __ccgo_ts+16660, int32FromInt32(3207))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if boolInt32(b2IsValidFloat(tls, radius) != 0 && radius >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+80, __ccgo_ts+9789, uintptr(unsafe.Pointer(&__func__310))))
		if v1 = !!(b2IsValidFloat(tls, radius) != 0 && radius >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+9789, __ccgo_ts+16660, int32FromInt32(3208))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if boolInt32(b2IsValidFloat(tls, falloff) != 0 && falloff >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+80, __ccgo_ts+20087, uintptr(unsafe.Pointer(&__func__310))))
		if v1 = !!(b2IsValidFloat(tls, falloff) != 0 && falloff >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+20087, __ccgo_ts+16660, int32FromInt32(3209))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if int32FromUint8(b2IsValidFloat(tls, impulsePerLength)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+80, __ccgo_ts+20132, uintptr(unsafe.Pointer(&__func__310))))
		if v1 = !!(b2IsValidFloat(tls, impulsePerLength) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+20132, __ccgo_ts+16660, int32FromInt32(3210))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(3213))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WorldExplode)(unsafe.Add(nil, bp)) = b2RecArgs_WorldExplode{
			World: worldId,
			Def:   *(*ExplosionDef)(unsafe.Add(nil, explosionDef)),
		}
		b2RecWrite_WorldExplode(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	*(*ExplosionContext)(unsafe.Add(nil, bp+40)) = ExplosionContext{
		World:            world,
		Position:         position,
		Radius:           radius,
		Falloff:          falloff,
		ImpulsePerLength: impulsePerLength,
	}
	// The broad-phase tree is float, so translate a local query box out to world with outward rounding
	extent = radius + falloff
	localBox = AABB{
		LowerBound: Vec2{
			X: -extent,
			Y: -extent,
		},
		UpperBound: Vec2{
			X: extent,
			Y: extent,
		},
	}
	v6 = localBox
	v7 = position
	v8 = float32(float64(v7.X) + float64(v6.LowerBound.X))
	goto _9
_9:
	result.LowerBound.X = v8
	v10 = float32(float64(v7.Y) + float64(v6.LowerBound.Y))
	goto _11
_11:
	result.LowerBound.Y = v10
	v12 = float32(float64(v7.X) + float64(v6.UpperBound.X))
	goto _13
_13:
	result.UpperBound.X = v12
	v14 = float32(float64(v7.Y) + float64(v6.UpperBound.Y))
	goto _15
_15:
	result.UpperBound.Y = v14
	v16 = result
	goto _17
_17:
	aabb = v16
	b2DynamicTree_OverlapAABB(tls, world+40+uintptr(b2_dynamicBody)*104, aabb, maskBits, __ccgo_fp(ExplosionCallback), bp+40)
}

// C documentation
//
//	// This is for internal testing. Not optimized.
func b2World_RebuildStaticTree(tls *_Stack, worldId WorldId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i, movedCount, proxyCount int32
	var movedProxies, staticTree, world uintptr
	var v1 bool
	var v2, v4 uint8
	var _ /* _a at bp+0 */ b2RecArgs_WorldRebuildStaticTree
	_, _, _, _, _, _, _, _, _ = i, movedCount, movedProxies, proxyCount, staticTree, world, v1, v2, v4
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+16660, int32FromInt32(3235))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WorldRebuildStaticTree)(unsafe.Add(nil, bp)) = b2RecArgs_WorldRebuildStaticTree{
			World: worldId,
		}
		b2RecWrite_WorldRebuildStaticTree(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	staticTree = world + 40 + uintptr(b2_staticBody)*104
	// Moved proxies must retain the flag in the rebuild.
	movedCount = 0
	movedProxies = uintptrFromInt32(0)
	proxyCount = b2DynamicTree_GetProxyCount(tls, staticTree)
	v4 = boolUint8((*b2TreeNode)(unsafe.Pointer((*b2DynamicTree)(unsafe.Pointer(staticTree)).Nodes+uintptr(_B2_ROOT_NODE)*32)).FlagIndex&(uint32FromUint32(1)<<int32FromInt32(30)) == uint32FromUint32(1)<<int32FromInt32(30))
	goto _5
_5:
	v2 = v4
	goto _3
_3:
	if v2 != 0 {
		movedProxies = b2Alloc(tls, uint64FromInt32(proxyCount)*uint64(4))
		movedCount = b2DynamicTree_GatherMovedProxies(tls, staticTree, movedProxies)
	}
	b2DynamicTree_Rebuild(tls, staticTree, boolUint8(true1 != 0))
	i = 0
	for {
		if !(i < movedCount) {
			break
		}
		b2DynamicTree_MarkProxyMovedSerial(tls, staticTree, *(*int32)(unsafe.Add(nil, movedProxies+uintptr(i)*4)))
		goto _6
	_6:
		;
		i = i + 1
	}
	if movedProxies != uintptrFromInt32(0) {
		b2Free(tls, movedProxies, uint64FromInt32(proxyCount)*uint64(4))
	}
}

// C documentation
//
//	// Hash transforms and velocities.
func b2HashWorldState(tls *_Stack, world2 uintptr) (r uint64) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var body2, bodySim, set, set1, sim, state, v2, v3, v6 uintptr
	var bodyCount, i int32
	var bx, by, hash1, prime, v10, v8 uint64_t
	var v4, v5 bool
	var v9 Pos
	var _ /* bits at bp+16 */ uint32_t
	var _ /* fx at bp+8 */ uint32_t
	var _ /* fy at bp+12 */ uint32_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body2, bodyCount, bodySim, bx, by, hash1, i, prime, set, set1, sim, state, v10, v2, v3, v4, v5, v6, v8, v9
	hash1 = uint64(14695981039346656037)
	prime = uint64(1099511628211)
	bodyCount = (*b2World)(unsafe.Pointer(world2)).Bodies.Count
	i = 0
	for {
		if !(i < bodyCount) {
			break
		}
		body2 = (*b2World)(unsafe.Pointer(world2)).Bodies.Data + uintptr(i)*104
		if (*b2Body)(unsafe.Pointer(body2)).Id != i {
			// Free or never-used slot
			goto _1
		}
		v2 = world2
		v3 = body2
		if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v3)).SetIndex && (*b2Body)(unsafe.Pointer(v3)).SetIndex < (*b2World)(unsafe.Pointer(v2)).SolverSets.Count); !v4 {
			b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		set = (*b2World)(unsafe.Pointer(v2)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v3)).SetIndex)*88
		if v5 = !!(0 <= (*b2Body)(unsafe.Pointer(v3)).LocalIndex && (*b2Body)(unsafe.Pointer(v3)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v5 {
			b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
		}
		_ = v5 || bool(int32FromInt32(0) != 0)
		bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v3)).LocalIndex)*96
		v6 = bodySim
		goto _7
	_7:
		sim = v6
		v8 = hash1
		v9 = (*b2BodySim)(unsafe.Pointer(sim)).Transform.P
		*(*Pos)(unsafe.Pointer(bp)) = v9
		memcpy(tls, bp+8, bp, uint64(4))
		memcpy(tls, bp+12, bp+4, uint64(4))
		bx = uint64(*(*uint32_t)(unsafe.Add(nil, bp+8)))
		by = uint64(*(*uint32_t)(unsafe.Add(nil, bp+12)))
		v8 = (v8 ^ bx) * uint64(1099511628211)
		v8 = (v8 ^ by) * uint64(1099511628211)
		v10 = v8
		goto _11
	_11:
		hash1 = v10
		memcpy(tls, bp+16, sim+8, uint64(4))
		hash1 = (hash1 ^ uint64(*(*uint32_t)(unsafe.Add(nil, bp+16)))) * prime
		memcpy(tls, bp+16, sim+8+4, uint64(4))
		hash1 = (hash1 ^ uint64(*(*uint32_t)(unsafe.Add(nil, bp+16)))) * prime
		v2 = world2
		v3 = body2
		if (*b2Body)(unsafe.Pointer(v3)).SetIndex == int32(b2_awakeSet) {
			if v4 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(v2)).SolverSets.Count); !v4 {
				b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+698, int32FromInt32(254))
			}
			_ = v4 || bool(int32FromInt32(0) != 0)
			set1 = (*b2World)(unsafe.Pointer(v2)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
			if v5 = !!(0 <= (*b2Body)(unsafe.Pointer(v3)).LocalIndex && (*b2Body)(unsafe.Pointer(v3)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set1)).BodyStates.Count); !v5 {
				b2InternalAssert(tls, __ccgo_ts+790, __ccgo_ts+698, int32FromInt32(255))
			}
			_ = v5 || bool(int32FromInt32(0) != 0)
			v6 = (*b2SolverSet)(unsafe.Pointer(set1)).BodyStates.Data + uintptr((*b2Body)(unsafe.Pointer(v3)).LocalIndex)*32
			goto _16
		}
		v6 = uintptrFromInt32(0)
		goto _16
	_16:
		state = v6
		if state != uintptrFromInt32(0) {
			memcpy(tls, bp+16, state, uint64(4))
			hash1 = (hash1 ^ uint64(*(*uint32_t)(unsafe.Add(nil, bp+16)))) * prime
			memcpy(tls, bp+16, state+4, uint64(4))
			hash1 = (hash1 ^ uint64(*(*uint32_t)(unsafe.Add(nil, bp+16)))) * prime
			memcpy(tls, bp+16, state+8, uint64(4))
			hash1 = (hash1 ^ uint64(*(*uint32_t)(unsafe.Add(nil, bp+16)))) * prime
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return hash1
}

// C documentation
//
//	// World scalar config (simulation settings only, no runtime/shell fields)
//	// Only simulation scalars belong here, never host or worker state (workerCount,
//	// scheduler, callbacks, user data). b2World_RestoreFromSnapshot relies on that so an in-place
//	// restore preserves the live world's wiring.
func b2SerWorldConfig(tls *_Stack, buf uintptr, world uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var v1 uint32
	var _ /* flags at bp+0 */ uint8_t
	_ = v1
	b2SnapW_Bytes(tls, buf, world+2484, int32(8))
	b2SnapW_Bytes(tls, buf, world+2492, int32(4))
	b2SnapW_Bytes(tls, buf, world+2496, int32(4))
	b2SnapW_I32(tls, buf, (*b2World)(unsafe.Pointer(world)).RestitutionIterations)
	b2SnapW_Bytes(tls, buf, world+2508, int32(4))
	b2SnapW_Bytes(tls, buf, world+2512, int32(4))
	b2SnapW_Bytes(tls, buf, world+2516, int32(4))
	b2SnapW_Bytes(tls, buf, world+2520, int32(4))
	b2SnapW_Bytes(tls, buf, world+2524, int32(4))
	b2SnapW_Bytes(tls, buf, world+2472, int32(8))
	b2SnapW_I32(tls, buf, (*b2World)(unsafe.Pointer(world)).SplitIslandId)
	// Step scaling cached for the force/torque reporting getters, which run between steps
	b2SnapW_Bytes(tls, buf, world+2768, int32(4))
	b2SnapW_Bytes(tls, buf, world+2772, int32(4))
	// End-event double-buffer parity, so the first post-restore event query reads the right half
	b2SnapW_I32(tls, buf, (*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)
	// maxCapacity (b2Capacity struct)
	b2SnapW_Bytes(tls, buf, world+2640, int32(24))
	// bool flags packed as individual bytes for layout stability
	*(*uint8_t)(unsafe.Add(nil, bp)) = uint8(0)
	if (*b2World)(unsafe.Pointer(world)).EnableSleep != 0 {
		v1 = uint32(0x01)
	} else {
		v1 = 0
	}
	*(*uint8_t)(unsafe.Add(nil, bp)) = uint8(uint32(*(*uint8_t)(unsafe.Add(nil, bp))) | v1)
	if (*b2World)(unsafe.Pointer(world)).EnableWarmStarting != 0 {
		v1 = uint32(0x02)
	} else {
		v1 = 0
	}
	*(*uint8_t)(unsafe.Add(nil, bp)) = uint8(uint32(*(*uint8_t)(unsafe.Add(nil, bp))) | v1)
	if (*b2World)(unsafe.Pointer(world)).EnableContinuous != 0 {
		v1 = uint32(0x08)
	} else {
		v1 = 0
	}
	*(*uint8_t)(unsafe.Add(nil, bp)) = uint8(uint32(*(*uint8_t)(unsafe.Add(nil, bp))) | v1)
	if (*b2World)(unsafe.Pointer(world)).EnableRestitutionPropagation != 0 {
		v1 = uint32(0x10)
	} else {
		v1 = 0
	}
	*(*uint8_t)(unsafe.Add(nil, bp)) = uint8(uint32(*(*uint8_t)(unsafe.Add(nil, bp))) | v1)
	b2RecBufAppend(tls, buf, bp, int32(1))
}

func b2DesWorldConfig(tls *_Stack, r uintptr, world uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* flags at bp+0 */ uint8_t
	b2SnapR_Bytes(tls, r, world+2484, int32(8))
	b2SnapR_Bytes(tls, r, world+2492, int32(4))
	b2SnapR_Bytes(tls, r, world+2496, int32(4))
	(*b2World)(unsafe.Pointer(world)).RestitutionIterations = b2SnapR_I32(tls, r)
	b2SnapR_Bytes(tls, r, world+2508, int32(4))
	b2SnapR_Bytes(tls, r, world+2512, int32(4))
	b2SnapR_Bytes(tls, r, world+2516, int32(4))
	b2SnapR_Bytes(tls, r, world+2520, int32(4))
	b2SnapR_Bytes(tls, r, world+2524, int32(4))
	b2SnapR_Bytes(tls, r, world+2472, int32(8))
	(*b2World)(unsafe.Pointer(world)).SplitIslandId = b2SnapR_I32(tls, r)
	b2SnapR_Bytes(tls, r, world+2768, int32(4))
	b2SnapR_Bytes(tls, r, world+2772, int32(4))
	(*b2World)(unsafe.Pointer(world)).EndEventArrayIndex = b2SnapR_I32(tls, r)
	b2SnapR_Bytes(tls, r, world+2640, int32(24))
	*(*uint8_t)(unsafe.Add(nil, bp)) = uint8(0)
	b2SnapR_Bytes(tls, r, bp, int32(1))
	(*b2World)(unsafe.Pointer(world)).EnableSleep = boolUint8(uint32(*(*uint8_t)(unsafe.Add(nil, bp)))&uint32(0x01) != uint32(0))
	(*b2World)(unsafe.Pointer(world)).EnableWarmStarting = boolUint8(uint32(*(*uint8_t)(unsafe.Add(nil, bp)))&uint32(0x02) != uint32(0))
	(*b2World)(unsafe.Pointer(world)).EnableContinuous = boolUint8(uint32(*(*uint8_t)(unsafe.Add(nil, bp)))&uint32(0x08) != uint32(0))
	(*b2World)(unsafe.Pointer(world)).EnableRestitutionPropagation = boolUint8(uint32(*(*uint8_t)(unsafe.Add(nil, bp)))&uint32(0x10) != uint32(0))
}

func b2SerializeWorld(tls *_Stack, world uintptr, buf uintptr) {
	bp1 := tls.Alloc(480)
	defer tls.Free(480)
	var bp, chain, graph, island, s uintptr
	var c, chainCount, i, i1, i2, i3, islandCount, sensorCount, setCount, slot, slot1, slot2, t int32
	var _ /* elem at bp+120 */ b2Shape
	var _ /* elem at bp+16 */ b2Body
	var _ /* elem at bp+400 */ b2Joint
	var _ /* hdr at bp+0 */ b2SnapHeader
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bp, c, chain, chainCount, graph, i, i1, i2, i3, island, islandCount, s, sensorCount, setCount, slot, slot1, slot2, t
	(*(*b2SnapHeader)(unsafe.Add(nil, bp1))).Magic = uint32(0x32534E42)
	(*(*b2SnapHeader)(unsafe.Add(nil, bp1))).Version = uint32(13)
	(*(*b2SnapHeader)(unsafe.Add(nil, bp1))).LayoutHash = b2ComputeLayoutHash(tls)
	(*(*b2SnapHeader)(unsafe.Add(nil, bp1))).Flags = 0
	if b2IsDoublePrecision(tls) != 0 {
		(*(*b2SnapHeader)(unsafe.Add(nil, bp1))).Flags |= uint32(0x2)
	}
	b2SnapW_Bytes(tls, buf, bp1, int32FromInt64(16))
	// World config
	b2SerWorldConfig(tls, buf, world)
	// 7 id pools
	b2SerIdPool(tls, buf, world+1912)
	b2SerIdPool(tls, buf, world+2112)
	b2SerIdPool(tls, buf, world+2136)
	b2SerIdPool(tls, buf, world+2032)
	b2SerIdPool(tls, buf, world+1992)
	b2SerIdPool(tls, buf, world+2072)
	b2SerIdPool(tls, buf, world+1952)
	// Solver sets
	setCount = (*b2World)(unsafe.Pointer(world)).SolverSets.Count
	b2SnapW_I32(tls, buf, setCount)
	i = 0
	for {
		if !(i < setCount) {
			break
		}
		b2SerSolverSet(tls, buf, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(i)*88)
		goto _1
	_1:
		;
		i = i + 1
	}
	// Sparse arrays. Bodies, shapes and joints carry a host userData pointer scrubbed to NULL on
	// write. Contacts have no userData, so they go out as raw POD.
	b2SnapW_I32(tls, buf, (*b2World)(unsafe.Pointer(world)).Bodies.Count)
	slot = 0
	for {
		if !(slot < (*b2World)(unsafe.Pointer(world)).Bodies.Count) {
			break
		}
		*(*b2Body)(unsafe.Add(nil, bp1+16)) = *(*b2Body)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Bodies.Data+uintptr(slot)*104))
		(*(*b2Body)(unsafe.Add(nil, bp1+16))).UserData = uintptrFromInt32(0)
		b2SnapW_Bytes(tls, buf, bp1+16, int32FromInt64(104))
		goto _2
	_2:
		;
		slot = slot + 1
	}
	b2SnapW_I32(tls, buf, (*b2World)(unsafe.Pointer(world)).Shapes.Count)
	slot1 = 0
	for {
		if !(slot1 < (*b2World)(unsafe.Pointer(world)).Shapes.Count) {
			break
		}
		*(*b2Shape)(unsafe.Pointer(bp1 + 120)) = b2Shape{}
		*(*b2Shape)(unsafe.Pointer(bp1 + 120)) = *(*b2Shape)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Shapes.Data+uintptr(slot1)*280))
		(*(*b2Shape)(unsafe.Add(nil, bp1+120))).UserData = uintptrFromInt32(0)
		b2SnapW_Bytes(tls, buf, bp1+120, int32FromInt64(280))
		goto _3
	_3:
		;
		slot1 = slot1 + 1
	}
	b2SnapW_I32(tls, buf, (*b2World)(unsafe.Pointer(world)).FatAABBs.Count)
	if (*b2World)(unsafe.Pointer(world)).FatAABBs.Count > 0 {
		b2SnapW_Bytes(tls, buf, (*b2World)(unsafe.Pointer(world)).FatAABBs.Data, (*b2World)(unsafe.Pointer(world)).FatAABBs.Count*int32FromInt64(16))
	}
	b2SnapW_I32(tls, buf, (*b2World)(unsafe.Pointer(world)).Contacts.Count)
	if (*b2World)(unsafe.Pointer(world)).Contacts.Count > 0 {
		b2SnapW_Bytes(tls, buf, (*b2World)(unsafe.Pointer(world)).Contacts.Data, (*b2World)(unsafe.Pointer(world)).Contacts.Count*int32FromInt64(64))
	}
	b2SnapW_I32(tls, buf, (*b2World)(unsafe.Pointer(world)).Joints.Count)
	slot2 = 0
	for {
		if !(slot2 < (*b2World)(unsafe.Pointer(world)).Joints.Count) {
			break
		}
		*(*b2Joint)(unsafe.Add(nil, bp1+400)) = *(*b2Joint)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Joints.Data+uintptr(slot2)*72))
		(*(*b2Joint)(unsafe.Add(nil, bp1+400))).UserData = uintptrFromInt32(0)
		b2SnapW_Bytes(tls, buf, bp1+400, int32FromInt64(72))
		goto _4
	_4:
		;
		slot2 = slot2 + 1
	}
	// Chain shapes: POD scalars then per-live-slot heap arrays
	chainCount = (*b2World)(unsafe.Pointer(world)).ChainShapes.Count
	b2SnapW_I32(tls, buf, chainCount)
	i1 = 0
	for {
		if !(i1 < chainCount) {
			break
		}
		chain = (*b2World)(unsafe.Pointer(world)).ChainShapes.Data + uintptr(i1)*32
		// Write POD scalars
		b2SnapW_I32(tls, buf, (*b2ChainShape)(unsafe.Pointer(chain)).Id)
		b2SnapW_I32(tls, buf, (*b2ChainShape)(unsafe.Pointer(chain)).BodyId)
		b2SnapW_I32(tls, buf, (*b2ChainShape)(unsafe.Pointer(chain)).NextChainId)
		b2SnapW_I32(tls, buf, (*b2ChainShape)(unsafe.Pointer(chain)).SegmentCount)
		b2SnapW_Bytes(tls, buf, chain+16, int32(2))
		if (*b2ChainShape)(unsafe.Pointer(chain)).Id != -int32(1) {
			// Live slot: write the two heap arrays
			b2SnapW_Bytes(tls, buf, (*b2ChainShape)(unsafe.Pointer(chain)).ShapeIndices, (*b2ChainShape)(unsafe.Pointer(chain)).SegmentCount*int32FromInt64(4))
		}
		goto _5
	_5:
		;
		i1 = i1 + 1
	}
	// Sensors: shapeId + 3 visitor arrays per slot
	sensorCount = (*b2World)(unsafe.Pointer(world)).Sensors.Count
	b2SnapW_I32(tls, buf, sensorCount)
	i2 = 0
	for {
		if !(i2 < sensorCount) {
			break
		}
		s = (*b2World)(unsafe.Pointer(world)).Sensors.Data + uintptr(i2)*56
		b2SnapW_I32(tls, buf, (*b2Sensor)(unsafe.Pointer(s)).ShapeId)
		b2SnapW_I32(tls, buf, (*b2Sensor)(unsafe.Pointer(s)).Hits.Count)
		if (*b2Sensor)(unsafe.Pointer(s)).Hits.Count > 0 {
			b2SnapW_Bytes(tls, buf, (*b2Sensor)(unsafe.Pointer(s)).Hits.Data, (*b2Sensor)(unsafe.Pointer(s)).Hits.Count*int32FromInt64(8))
		}
		b2SnapW_I32(tls, buf, (*b2Sensor)(unsafe.Pointer(s)).Overlaps1.Count)
		if (*b2Sensor)(unsafe.Pointer(s)).Overlaps1.Count > 0 {
			b2SnapW_Bytes(tls, buf, (*b2Sensor)(unsafe.Pointer(s)).Overlaps1.Data, (*b2Sensor)(unsafe.Pointer(s)).Overlaps1.Count*int32FromInt64(8))
		}
		b2SnapW_I32(tls, buf, (*b2Sensor)(unsafe.Pointer(s)).Overlaps2.Count)
		if (*b2Sensor)(unsafe.Pointer(s)).Overlaps2.Count > 0 {
			b2SnapW_Bytes(tls, buf, (*b2Sensor)(unsafe.Pointer(s)).Overlaps2.Data, (*b2Sensor)(unsafe.Pointer(s)).Overlaps2.Count*int32FromInt64(8))
		}
		goto _6
	_6:
		;
		i2 = i2 + 1
	}
	// Islands: POD scalars + 3 inner arrays per slot
	islandCount = (*b2World)(unsafe.Pointer(world)).Islands.Count
	b2SnapW_I32(tls, buf, islandCount)
	i3 = 0
	for {
		if !(i3 < islandCount) {
			break
		}
		island = (*b2World)(unsafe.Pointer(world)).Islands.Data + uintptr(i3)*64
		b2SnapW_I32(tls, buf, (*b2Island)(unsafe.Pointer(island)).SetIndex)
		b2SnapW_I32(tls, buf, (*b2Island)(unsafe.Pointer(island)).LocalIndex)
		b2SnapW_I32(tls, buf, (*b2Island)(unsafe.Pointer(island)).IslandId)
		b2SnapW_I32(tls, buf, (*b2Island)(unsafe.Pointer(island)).ConstraintRemoveCount)
		b2SnapW_I32(tls, buf, (*b2Island)(unsafe.Pointer(island)).Bodies.Count)
		if (*b2Island)(unsafe.Pointer(island)).Bodies.Count > 0 {
			b2SnapW_Bytes(tls, buf, (*b2Island)(unsafe.Pointer(island)).Bodies.Data, (*b2Island)(unsafe.Pointer(island)).Bodies.Count*int32FromInt64(4))
		}
		b2SnapW_I32(tls, buf, (*b2Island)(unsafe.Pointer(island)).Contacts.Count)
		if (*b2Island)(unsafe.Pointer(island)).Contacts.Count > 0 {
			b2SnapW_Bytes(tls, buf, (*b2Island)(unsafe.Pointer(island)).Contacts.Data, (*b2Island)(unsafe.Pointer(island)).Contacts.Count*int32FromInt64(12))
		}
		b2SnapW_I32(tls, buf, (*b2Island)(unsafe.Pointer(island)).Joints.Count)
		if (*b2Island)(unsafe.Pointer(island)).Joints.Count > 0 {
			b2SnapW_Bytes(tls, buf, (*b2Island)(unsafe.Pointer(island)).Joints.Data, (*b2Island)(unsafe.Pointer(island)).Joints.Count*int32FromInt64(12))
		}
		goto _7
	_7:
		;
		i3 = i3 + 1
	}
	// Broad phase
	bp = world + 40
	t = 0
	for {
		if !(t < int32(b2_bodyTypeCount)) {
			break
		}
		b2SerTree(tls, buf, bp+uintptr(t)*104)
		goto _8
	_8:
		;
		t = t + 1
	}
	b2SerHashSet(tls, buf, bp+320)
	// Constraint graph: B2_GRAPH_COLOR_COUNT colors
	graph = world + 376
	c = 0
	for {
		if !(c < int32(_B2_GRAPH_COLOR_COUNT)) {
			break
		}
		b2SerGraphColor(tls, buf, graph+uintptr(c)*64, boolUint8(c == int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1)))
		goto _9
	_9:
		;
		c = c + 1
	}
}

func b2CreateWorldFromSnapshot(tls *_Stack, image uintptr, size int32, workerCount int32) (r1 WorldId) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var id, nullId WorldId
	var r, world uintptr
	var _ /* def at bp+24 */ WorldDef
	var _ /* readerStorage at bp+0 */ b2SnapReader
	_, _, _, _ = id, nullId, r, world
	nullId = b2_nullWorldId
	r = bp
	if int32FromUint8(b2OpenSnapshotImage(tls, image, size, r)) == false1 {
		return nullId
	}
	// Build a minimal valid def so b2CreateWorld produces a fully valid shell
	*(*WorldDef)(unsafe.Add(nil, bp+24)) = b2DefaultWorldDef(tls)
	(*(*WorldDef)(unsafe.Add(nil, bp+24))).WorkerCount = workerCount
	(*(*WorldDef)(unsafe.Add(nil, bp+24))).EnqueueTask = uintptrFromInt32(0)
	(*(*WorldDef)(unsafe.Add(nil, bp+24))).FinishTask = uintptrFromInt32(0)
	(*(*WorldDef)(unsafe.Add(nil, bp+24))).UserTaskContext = uintptrFromInt32(0)
	// Capacity is only a sizing hint. Every container is resized from the image below
	id = b2CreateWorld(tls, bp+24)
	if !(b2World_IsValid(tls, id) != 0) {
		return nullId
	}
	world = b2GetWorldFromId(tls, id)
	if int32FromUint8(b2DeserializeIntoShell(tls, r, world)) == false1 {
		// Image was corrupt; clean up by destroying the world
		b2DestroyWorld(tls, id)
		return nullId
	}
	// A world loaded from scratch carries no host pointers
	(*b2World)(unsafe.Pointer(world)).PreSolveFcn = uintptrFromInt32(0)
	(*b2World)(unsafe.Pointer(world)).PreSolveContext = uintptrFromInt32(0)
	(*b2World)(unsafe.Pointer(world)).CustomFilterFcn = uintptrFromInt32(0)
	(*b2World)(unsafe.Pointer(world)).CustomFilterContext = uintptrFromInt32(0)
	(*b2World)(unsafe.Pointer(world)).UserData = uintptrFromInt32(0)
	return id
}

func b2World_RestoreFromSnapshot(tls *_Stack, worldId WorldId, image uintptr, size int32) (r1 uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var r, world uintptr
	var v1 bool
	var _ /* readerStorage at bp+0 */ b2SnapReader
	_, _, _ = r, world, v1
	r = bp
	if int32FromUint8(b2OpenSnapshotImage(tls, image, size, r)) == false1 {
		return boolUint8(false1 != 0)
	}
	world = b2GetWorldFromId(tls, worldId)
	// Restoring mid step would corrupt an in-flight solve
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+24157, int32FromInt32(964))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return boolUint8(false1 != 0)
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		b2Log(tls, __ccgo_ts+24183, 0)
		return boolUint8(false1 != 0)
	}
	// Point of no return. The slot, generation, and host wiring are kept, so held ids
	// resolve into the rebuilt world. A truncated payload past here leaves the world
	// unusable and the caller must destroy it.
	b2FreeLiveSimElements(tls, world)
	return b2DeserializeIntoShell(tls, r, world)
}

func b2World_GetSnapshot(tls *_Stack, worldId WorldId, image uintptr, capacity int32) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var size int32
	var world uintptr
	var v1 bool
	var _ /* buf at bp+24 */ b2RecBuffer
	var _ /* counter at bp+0 */ b2RecBuffer
	_, _, _ = size, world, v1
	world = b2GetWorldFromId(tls, worldId)
	// Serializing mid step would capture an inconsistent, in-flight world
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+24157, int32FromInt32(989))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return 0
	}
	// Size query: count the bytes without allocating or copying the whole image
	if image == uintptrFromInt32(0) {
		*(*b2RecBuffer)(unsafe.Add(nil, bp)) = b2RecBuffer{}
		(*(*b2RecBuffer)(unsafe.Add(nil, bp))).CountOnly = boolUint8(true1 != 0)
		b2SerializeWorld(tls, world, bp)
		return (*(*b2RecBuffer)(unsafe.Add(nil, bp))).Size
	}
	*(*b2RecBuffer)(unsafe.Add(nil, bp+24)) = b2RecBuffer{}
	b2SerializeWorld(tls, world, bp+24)
	size = (*(*b2RecBuffer)(unsafe.Add(nil, bp+24))).Size
	if size <= capacity {
		memcpy(tls, image, (*(*b2RecBuffer)(unsafe.Add(nil, bp+24))).Data, uint64FromInt32(size))
	}
	b2RecBufFree(tls, bp+24)
	return size
}

func b2HashWorldStateDeep(tls *_Stack, world2 uintptr) (r uint64) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body2, bodySim, contact, joint, m, set, set1, sim, sim1, sim2, state, v2, v3, v6 uintptr
	var bodyCount, contactCount, i, i1, i2, jointCount, p1, v22 int32
	var bx, by, hash1, v10, v8 uint64_t
	var v4, v5 bool
	var v9 Pos
	var _ /* fx at bp+8 */ uint32_t
	var _ /* fy at bp+12 */ uint32_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body2, bodyCount, bodySim, bx, by, contact, contactCount, hash1, i, i1, i2, joint, jointCount, m, p1, set, set1, sim, sim1, sim2, state, v10, v2, v22, v3, v4, v5, v6, v8, v9
	hash1 = uint64(14695981039346656037)
	// Bodies: same iteration order as b2HashWorldState (sparse array, skip free slots)
	bodyCount = (*b2World)(unsafe.Pointer(world2)).Bodies.Count
	i = 0
	for {
		if !(i < bodyCount) {
			break
		}
		body2 = (*b2World)(unsafe.Pointer(world2)).Bodies.Data + uintptr(i)*104
		if (*b2Body)(unsafe.Pointer(body2)).Id != i {
			goto _1
		}
		v2 = world2
		v3 = body2
		if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v3)).SetIndex && (*b2Body)(unsafe.Pointer(v3)).SetIndex < (*b2World)(unsafe.Pointer(v2)).SolverSets.Count); !v4 {
			b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		set = (*b2World)(unsafe.Pointer(v2)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v3)).SetIndex)*88
		if v5 = !!(0 <= (*b2Body)(unsafe.Pointer(v3)).LocalIndex && (*b2Body)(unsafe.Pointer(v3)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v5 {
			b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
		}
		_ = v5 || bool(int32FromInt32(0) != 0)
		bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v3)).LocalIndex)*96
		v6 = bodySim
		goto _7
	_7:
		sim = v6
		v8 = hash1
		v9 = (*b2BodySim)(unsafe.Pointer(sim)).Transform.P
		*(*Pos)(unsafe.Pointer(bp)) = v9
		memcpy(tls, bp+8, bp, uint64(4))
		memcpy(tls, bp+12, bp+4, uint64(4))
		bx = uint64(*(*uint32_t)(unsafe.Add(nil, bp+8)))
		by = uint64(*(*uint32_t)(unsafe.Add(nil, bp+12)))
		v8 = (v8 ^ bx) * uint64(1099511628211)
		v8 = (v8 ^ by) * uint64(1099511628211)
		v10 = v8
		goto _11
	_11:
		hash1 = v10
		hash1 = b2FnvMixFloat(tls, hash1, (*b2BodySim)(unsafe.Pointer(sim)).Transform.Q.C)
		hash1 = b2FnvMixFloat(tls, hash1, (*b2BodySim)(unsafe.Pointer(sim)).Transform.Q.S)
		v2 = world2
		v3 = body2
		if (*b2Body)(unsafe.Pointer(v3)).SetIndex == int32(b2_awakeSet) {
			if v4 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(v2)).SolverSets.Count); !v4 {
				b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+698, int32FromInt32(254))
			}
			_ = v4 || bool(int32FromInt32(0) != 0)
			set1 = (*b2World)(unsafe.Pointer(v2)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
			if v5 = !!(0 <= (*b2Body)(unsafe.Pointer(v3)).LocalIndex && (*b2Body)(unsafe.Pointer(v3)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set1)).BodyStates.Count); !v5 {
				b2InternalAssert(tls, __ccgo_ts+790, __ccgo_ts+698, int32FromInt32(255))
			}
			_ = v5 || bool(int32FromInt32(0) != 0)
			v6 = (*b2SolverSet)(unsafe.Pointer(set1)).BodyStates.Data + uintptr((*b2Body)(unsafe.Pointer(v3)).LocalIndex)*32
			goto _16
		}
		v6 = uintptrFromInt32(0)
		goto _16
	_16:
		state = v6
		if state != uintptrFromInt32(0) {
			hash1 = b2FnvMixFloat(tls, hash1, (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity.X)
			hash1 = b2FnvMixFloat(tls, hash1, (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity.Y)
			hash1 = b2FnvMixFloat(tls, hash1, (*b2BodyState)(unsafe.Pointer(state)).AngularVelocity)
		}
		// Index bookkeeping
		hash1 = b2FnvMixInt(tls, hash1, (*b2Body)(unsafe.Pointer(body2)).SetIndex)
		hash1 = b2FnvMixInt(tls, hash1, (*b2Body)(unsafe.Pointer(body2)).LocalIndex)
		goto _1
	_1:
		;
		i = i + 1
	}
	// Contacts: sparse array, skip free slots (colorIndex == B2_NULL_INDEX and setIndex ==
	// B2_NULL_INDEX for free)
	contactCount = (*b2World)(unsafe.Pointer(world2)).Contacts.Count
	i1 = 0
	for {
		if !(i1 < contactCount) {
			break
		}
		contact = (*b2World)(unsafe.Pointer(world2)).Contacts.Data + uintptr(i1)*64
		if (*b2Contact)(unsafe.Pointer(contact)).ContactId != i1 {
			goto _18
		}
		hash1 = b2FnvMixInt(tls, hash1, (*b2Contact)(unsafe.Pointer(contact)).SetIndex)
		hash1 = b2FnvMixInt(tls, hash1, (*b2Contact)(unsafe.Pointer(contact)).ColorIndex)
		hash1 = b2FnvMixInt(tls, hash1, (*b2Contact)(unsafe.Pointer(contact)).LocalIndex)
		// Manifold points + impulses from the contact sim
		// Exclude B2_ENABLE_VALIDATION-gated bodyIdA/bodyIdB for build-config stability
		sim1 = b2GetContactSim(tls, world2, contact)
		if sim1 != uintptrFromInt32(0) {
			m = sim1 + 72
			hash1 = b2FnvMixInt(tls, hash1, (*Manifold)(unsafe.Pointer(m)).PointCount)
			p1 = 0
			for {
				if !(p1 < (*Manifold)(unsafe.Pointer(m)).PointCount) {
					break
				}
				hash1 = b2FnvMixFloat(tls, hash1, (*(*ManifoldPoint)(unsafe.Add(nil, m+12+uintptr(p1)*44))).NormalImpulse)
				hash1 = b2FnvMixFloat(tls, hash1, (*(*ManifoldPoint)(unsafe.Add(nil, m+12+uintptr(p1)*44))).TangentImpulse)
				hash1 = b2FnvMixFloat(tls, hash1, (*(*ManifoldPoint)(unsafe.Add(nil, m+12+uintptr(p1)*44))).TotalNormalImpulse)
				goto _19
			_19:
				;
				p1 = p1 + 1
			}
		}
		goto _18
	_18:
		;
		i1 = i1 + 1
	}
	// Joints: sparse array, skip free slots
	jointCount = (*b2World)(unsafe.Pointer(world2)).Joints.Count
	i2 = 0
	for {
		if !(i2 < jointCount) {
			break
		}
		joint = (*b2World)(unsafe.Pointer(world2)).Joints.Data + uintptr(i2)*72
		if (*b2Joint)(unsafe.Pointer(joint)).JointId != i2 {
			goto _20
		}
		hash1 = b2FnvMixInt(tls, hash1, (*b2Joint)(unsafe.Pointer(joint)).SetIndex)
		hash1 = b2FnvMixInt(tls, hash1, (*b2Joint)(unsafe.Pointer(joint)).ColorIndex)
		hash1 = b2FnvMixInt(tls, hash1, (*b2Joint)(unsafe.Pointer(joint)).LocalIndex)
		sim2 = b2GetJointSim(tls, world2, joint)
		if sim2 != uintptrFromInt32(0) {
			// Hash accumulated impulses per joint type
			switch (*b2JointSim)(unsafe.Pointer(sim2)).Type1 {
			case int32(b2_distanceJoint):
				hash1 = b2FnvMixFloat(tls, hash1, (*b2JointSim)(unsafe.Pointer(sim2)).ｆ__ccgo15_92.DistanceJoint.Impulse)
				hash1 = b2FnvMixFloat(tls, hash1, (*b2JointSim)(unsafe.Pointer(sim2)).ｆ__ccgo15_92.DistanceJoint.LowerImpulse)
				hash1 = b2FnvMixFloat(tls, hash1, (*b2JointSim)(unsafe.Pointer(sim2)).ｆ__ccgo15_92.DistanceJoint.UpperImpulse)
				hash1 = b2FnvMixFloat(tls, hash1, (*b2JointSim)(unsafe.Pointer(sim2)).ｆ__ccgo15_92.DistanceJoint.MotorImpulse)
			case int32(b2_moverJoint):
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2MoverJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).LinearVelocityImpulse.X)
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2MoverJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).LinearVelocityImpulse.Y)
			case int32(b2_motorJoint):
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).LinearVelocityImpulse.X)
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).LinearVelocityImpulse.Y)
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).AngularVelocityImpulse)
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).LinearSpringImpulse.X)
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).LinearSpringImpulse.Y)
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).AngularSpringImpulse)
			case int32(b2_pogoJoint):
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).Impulse)
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).Velocity)
			case int32(b2_prismaticJoint):
				hash1 = b2FnvMixBytes(tls, hash1, sim2+92, int32(8))
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).SpringImpulse)
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).MotorImpulse)
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).LowerImpulse)
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).UpperImpulse)
			case int32(b2_revoluteJoint):
				hash1 = b2FnvMixBytes(tls, hash1, sim2+92, int32(8))
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).SpringImpulse)
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).MotorImpulse)
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).LowerImpulse)
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).UpperImpulse)
			case int32(b2_weldJoint):
				hash1 = b2FnvMixBytes(tls, hash1, sim2+92+40, int32(8))
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).AngularImpulse)
			case int32(b2_wheelJoint):
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).PerpImpulse)
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).MotorImpulse)
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).SpringImpulse)
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).LowerImpulse)
				hash1 = b2FnvMixFloat(tls, hash1, (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(sim2), 92))).UpperImpulse)
			default:
				break
			}
		}
		goto _20
	_20:
		;
		i2 = i2 + 1
	}
	// 7 id pools: nextIndex + count
	hash1 = b2FnvMixInt(tls, hash1, (*b2World)(unsafe.Pointer(world2)).BodyIdPool.NextIndex)
	v2 = world2 + 1912
	v22 = (*b2IdPool)(unsafe.Pointer(v2)).NextIndex - (*b2IdPool)(unsafe.Pointer(v2)).FreeArray.Count
	goto _23
_23:
	hash1 = b2FnvMixInt(tls, hash1, v22)
	hash1 = b2FnvMixInt(tls, hash1, (*b2World)(unsafe.Pointer(world2)).ShapeIdPool.NextIndex)
	v2 = world2 + 2112
	v22 = (*b2IdPool)(unsafe.Pointer(v2)).NextIndex - (*b2IdPool)(unsafe.Pointer(v2)).FreeArray.Count
	goto _26
_26:
	hash1 = b2FnvMixInt(tls, hash1, v22)
	hash1 = b2FnvMixInt(tls, hash1, (*b2World)(unsafe.Pointer(world2)).ChainIdPool.NextIndex)
	v2 = world2 + 2136
	v22 = (*b2IdPool)(unsafe.Pointer(v2)).NextIndex - (*b2IdPool)(unsafe.Pointer(v2)).FreeArray.Count
	goto _29
_29:
	hash1 = b2FnvMixInt(tls, hash1, v22)
	hash1 = b2FnvMixInt(tls, hash1, (*b2World)(unsafe.Pointer(world2)).ContactIdPool.NextIndex)
	v2 = world2 + 2032
	v22 = (*b2IdPool)(unsafe.Pointer(v2)).NextIndex - (*b2IdPool)(unsafe.Pointer(v2)).FreeArray.Count
	goto _32
_32:
	hash1 = b2FnvMixInt(tls, hash1, v22)
	hash1 = b2FnvMixInt(tls, hash1, (*b2World)(unsafe.Pointer(world2)).JointIdPool.NextIndex)
	v2 = world2 + 1992
	v22 = (*b2IdPool)(unsafe.Pointer(v2)).NextIndex - (*b2IdPool)(unsafe.Pointer(v2)).FreeArray.Count
	goto _35
_35:
	hash1 = b2FnvMixInt(tls, hash1, v22)
	hash1 = b2FnvMixInt(tls, hash1, (*b2World)(unsafe.Pointer(world2)).IslandIdPool.NextIndex)
	v2 = world2 + 2072
	v22 = (*b2IdPool)(unsafe.Pointer(v2)).NextIndex - (*b2IdPool)(unsafe.Pointer(v2)).FreeArray.Count
	goto _38
_38:
	hash1 = b2FnvMixInt(tls, hash1, v22)
	hash1 = b2FnvMixInt(tls, hash1, (*b2World)(unsafe.Pointer(world2)).SolverSetIdPool.NextIndex)
	v2 = world2 + 1952
	v22 = (*b2IdPool)(unsafe.Pointer(v2)).NextIndex - (*b2IdPool)(unsafe.Pointer(v2)).FreeArray.Count
	goto _41
_41:
	hash1 = b2FnvMixInt(tls, hash1, v22)
	// Solver set count
	hash1 = b2FnvMixInt(tls, hash1, (*b2World)(unsafe.Pointer(world2)).SolverSets.Count)
	return hash1
}

func b2World_GetStateHash(tls *_Stack, worldId WorldId) (r uint64) {
	var world uintptr
	_ = world
	world = b2GetWorldFromId(tls, worldId)
	return b2HashWorldStateDeep(tls, world)
}

func b2DefaultWorldDef(tls *_Stack) (r WorldDef) {
	var def WorldDef
	var lengthUnits float32
	_, _ = def, lengthUnits
	lengthUnits = b2GetLengthUnitsPerMeter(tls)
	def = WorldDef{}
	def.Gravity.X = float32FromFloat32(0)
	def.Gravity.Y = -float32FromFloat32(10)
	def.HitEventThreshold = float32(float32FromFloat32(1) * lengthUnits)
	def.RestitutionThreshold = float32(float32FromFloat32(1) * lengthUnits)
	def.RestitutionIterations = int32(2)
	def.EnableRestitutionPropagation = boolUint8(false1 != 0)
	def.ContactSpeed = float32(float32FromFloat32(3) * lengthUnits)
	def.ContactHertz = float32(30)
	def.ContactDampingRatio = float32FromFloat32(10)
	// 400 meters per second, faster than the speed of sound
	def.MaximumLinearSpeed = float32(float32FromFloat32(400) * lengthUnits)
	def.EnableSleep = boolUint8(true1 != 0)
	def.EnableContinuous = boolUint8(true1 != 0)
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}
