package b2

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

// C documentation
//
//	// Implement functions for b2BodyArray
func b2BodyArray_Create(tls *_Stack, capacity int32) (r b2BodyArray) {
	var a b2BodyArray
	_ = a
	a = b2BodyArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(128)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2BodyArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2BodyArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2BodyArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2BodyArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2BodyArray)(unsafe.Pointer(a)).Capacity)*uint64(128)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(128)))
	(*b2BodyArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2BodyArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2BodyArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2BodyArray)(unsafe.Pointer(a)).Capacity)*uint64(128)))
	(*b2BodyArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2BodyArray)(unsafe.Pointer(a)).Count = 0
	(*b2BodyArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2BodySimArray_Create(tls *_Stack, capacity int32) (r b2BodySimArray) {
	var a b2BodySimArray
	_ = a
	a = b2BodySimArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(100)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2BodySimArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2BodySimArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2BodySimArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2BodySimArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2BodySimArray)(unsafe.Pointer(a)).Capacity)*uint64(100)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(100)))
	(*b2BodySimArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2BodySimArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2BodySimArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2BodySimArray)(unsafe.Pointer(a)).Capacity)*uint64(100)))
	(*b2BodySimArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2BodySimArray)(unsafe.Pointer(a)).Count = 0
	(*b2BodySimArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2BodyStateArray_Create(tls *_Stack, capacity int32) (r b2BodyStateArray) {
	var a b2BodyStateArray
	_ = a
	a = b2BodyStateArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(32)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2BodyStateArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2BodyStateArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2BodyStateArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2BodyStateArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2BodyStateArray)(unsafe.Pointer(a)).Capacity)*uint64(32)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(32)))
	(*b2BodyStateArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2BodyStateArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2BodyStateArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2BodyStateArray)(unsafe.Pointer(a)).Capacity)*uint64(32)))
	(*b2BodyStateArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2BodyStateArray)(unsafe.Pointer(a)).Count = 0
	(*b2BodyStateArray)(unsafe.Pointer(a)).Capacity = 0
}

// C documentation
//
//	// Get a validated body from a world using an id.
func b2GetBodyFullId(tls *_Stack, world uintptr, bodyId BodyId) (r uintptr) {
	var v1, v3 uintptr
	var v2 int32
	_, _, _ = v1, v2, v3
	if !(b2Body_IsValid(tls, bodyId) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+302, __ccgo_ts+327, int32FromInt32(40)) != 0 {
		__builtin_trap(tls)
	}
	// id index starts at one so that zero can represent null
	v1 = world + 1024
	v2 = bodyId.Index1 - int32(1)
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	return v3
}

func b2GetBodyTransformQuick(tls *_Stack, world uintptr, body uintptr) (r Transform) {
	var bodySim, set, v1, v3, v5, v7 uintptr
	var v2, v6 int32
	_, _, _, _, _, _, _, _ = bodySim, set, v1, v2, v3, v5, v6, v7
	v1 = world + 1064
	v2 = (*b2Body)(unsafe.Pointer(body)).SetIndex
	if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
	goto _4
_4:
	set = v3
	v5 = set
	v6 = (*b2Body)(unsafe.Pointer(body)).LocalIndex
	if !(0 <= v6 && v6 < (*b2BodySimArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodySimArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*100
	goto _8
_8:
	bodySim = v7
	return (*b2BodySim)(unsafe.Pointer(bodySim)).Transform
}

func b2GetBodyTransform(tls *_Stack, world uintptr, bodyId int32) (r Transform) {
	var body, v1, v3 uintptr
	var v2 int32
	_, _, _, _ = body, v1, v2, v3
	v1 = world + 1024
	v2 = bodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	body = v3
	return b2GetBodyTransformQuick(tls, world, body)
}

// C documentation
//
//	// Create a b2BodyId from a raw id.
func b2MakeBodyId(tls *_Stack, world uintptr, bodyId int32) (r BodyId) {
	var body, v1, v3 uintptr
	var v2 int32
	_, _, _, _ = body, v1, v2, v3
	v1 = world + 1024
	v2 = bodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	body = v3
	return BodyId{
		Index1:     bodyId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Body)(unsafe.Pointer(body)).Generation,
	}
}

func b2GetBodySim(tls *_Stack, world uintptr, body uintptr) (r uintptr) {
	var bodySim, set, v1, v3, v5, v7 uintptr
	var v2, v6 int32
	_, _, _, _, _, _, _, _ = bodySim, set, v1, v2, v3, v5, v6, v7
	v1 = world + 1064
	v2 = (*b2Body)(unsafe.Pointer(body)).SetIndex
	if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
	goto _4
_4:
	set = v3
	v5 = set
	v6 = (*b2Body)(unsafe.Pointer(body)).LocalIndex
	if !(0 <= v6 && v6 < (*b2BodySimArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodySimArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*100
	goto _8
_8:
	bodySim = v7
	return bodySim
}

func b2GetBodyState(tls *_Stack, world uintptr, body uintptr) (r uintptr) {
	var set, v1, v3, v5, v7 uintptr
	var v2, v6 int32
	_, _, _, _, _, _, _ = set, v1, v2, v3, v5, v6, v7
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
		v1 = world + 1064
		v2 = int32(b2_awakeSet)
		if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
		goto _4
	_4:
		set = v3
		v5 = set + 16
		v6 = (*b2Body)(unsafe.Pointer(body)).LocalIndex
		if !(0 <= v6 && v6 < (*b2BodyStateArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(194)) != 0 {
			__builtin_trap(tls)
		}
		v7 = (*b2BodyStateArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*32
		goto _8
	_8:
		return v7
	}
	return uintptrFromInt32(0)
}

func b2CreateIslandForBody(tls *_Stack, world uintptr, setIndex int32, body uintptr) {
	var island uintptr
	_ = island
	if !((*b2Body)(unsafe.Pointer(body)).IslandId == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+430, __ccgo_ts+327, int32FromInt32(86)) != 0 {
		__builtin_trap(tls)
	}
	if !((*b2Body)(unsafe.Pointer(body)).IslandPrev == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+462, __ccgo_ts+327, int32FromInt32(87)) != 0 {
		__builtin_trap(tls)
	}
	if !((*b2Body)(unsafe.Pointer(body)).IslandNext == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+496, __ccgo_ts+327, int32FromInt32(88)) != 0 {
		__builtin_trap(tls)
	}
	if !(setIndex != int32(b2_disabledSet)) && b2InternalAssertFcn(tls, __ccgo_ts+530, __ccgo_ts+327, int32FromInt32(89)) != 0 {
		__builtin_trap(tls)
	}
	island = b2CreateIsland(tls, world, setIndex)
	(*b2Body)(unsafe.Pointer(body)).IslandId = (*b2Island)(unsafe.Pointer(island)).IslandId
	(*b2Island)(unsafe.Pointer(island)).HeadBody = (*b2Body)(unsafe.Pointer(body)).Id
	(*b2Island)(unsafe.Pointer(island)).TailBody = (*b2Body)(unsafe.Pointer(body)).Id
	(*b2Island)(unsafe.Pointer(island)).BodyCount = int32(1)
}

func b2RemoveBodyFromIsland(tls *_Stack, world uintptr, body uintptr) {
	var island, nextBody, prevBody, v1, v11, v3, v5, v7, v9 uintptr
	var islandDestroyed uint8
	var islandId, v10, v2, v6 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = island, islandDestroyed, islandId, nextBody, prevBody, v1, v10, v11, v2, v3, v5, v6, v7, v9
	if (*b2Body)(unsafe.Pointer(body)).IslandId == -int32(1) {
		if !((*b2Body)(unsafe.Pointer(body)).IslandPrev == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+462, __ccgo_ts+327, int32FromInt32(103)) != 0 {
			__builtin_trap(tls)
		}
		if !((*b2Body)(unsafe.Pointer(body)).IslandNext == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+496, __ccgo_ts+327, int32FromInt32(104)) != 0 {
			__builtin_trap(tls)
		}
		return
	}
	islandId = (*b2Body)(unsafe.Pointer(body)).IslandId
	v1 = world + 1184
	v2 = islandId
	if !(0 <= v2 && v2 < (*b2IslandArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2IslandArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*56
	goto _4
_4:
	island = v3
	// Fix the island's linked list of sims
	if (*b2Body)(unsafe.Pointer(body)).IslandPrev != -int32(1) {
		v5 = world + 1024
		v6 = (*b2Body)(unsafe.Pointer(body)).IslandPrev
		if !(0 <= v6 && v6 < (*b2BodyArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
			__builtin_trap(tls)
		}
		v7 = (*b2BodyArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*128
		goto _8
	_8:
		prevBody = v7
		(*b2Body)(unsafe.Pointer(prevBody)).IslandNext = (*b2Body)(unsafe.Pointer(body)).IslandNext
	}
	if (*b2Body)(unsafe.Pointer(body)).IslandNext != -int32(1) {
		v9 = world + 1024
		v10 = (*b2Body)(unsafe.Pointer(body)).IslandNext
		if !(0 <= v10 && v10 < (*b2BodyArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
			__builtin_trap(tls)
		}
		v11 = (*b2BodyArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*128
		goto _12
	_12:
		nextBody = v11
		(*b2Body)(unsafe.Pointer(nextBody)).IslandPrev = (*b2Body)(unsafe.Pointer(body)).IslandPrev
	}
	if !((*b2Island)(unsafe.Pointer(island)).BodyCount > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+581, __ccgo_ts+327, int32FromInt32(124)) != 0 {
		__builtin_trap(tls)
	}
	*(*int32)(unsafe.Pointer(island + 20)) -= int32(1)
	islandDestroyed = boolUint8(false1 != 0)
	if (*b2Island)(unsafe.Pointer(island)).HeadBody == (*b2Body)(unsafe.Pointer(body)).Id {
		(*b2Island)(unsafe.Pointer(island)).HeadBody = (*b2Body)(unsafe.Pointer(body)).IslandNext
		if (*b2Island)(unsafe.Pointer(island)).HeadBody == -int32(1) {
			// Destroy empty island
			if !((*b2Island)(unsafe.Pointer(island)).TailBody == (*b2Body)(unsafe.Pointer(body)).Id) && b2InternalAssertFcn(tls, __ccgo_ts+603, __ccgo_ts+327, int32FromInt32(135)) != 0 {
				__builtin_trap(tls)
			}
			if !((*b2Island)(unsafe.Pointer(island)).BodyCount == int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+632, __ccgo_ts+327, int32FromInt32(136)) != 0 {
				__builtin_trap(tls)
			}
			if !((*b2Island)(unsafe.Pointer(island)).ContactCount == int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+655, __ccgo_ts+327, int32FromInt32(137)) != 0 {
				__builtin_trap(tls)
			}
			if !((*b2Island)(unsafe.Pointer(island)).JointCount == int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+681, __ccgo_ts+327, int32FromInt32(138)) != 0 {
				__builtin_trap(tls)
			}
			// Free the island
			b2DestroyIsland(tls, world, (*b2Island)(unsafe.Pointer(island)).IslandId)
			islandDestroyed = boolUint8(true1 != 0)
		}
	} else {
		if (*b2Island)(unsafe.Pointer(island)).TailBody == (*b2Body)(unsafe.Pointer(body)).Id {
			(*b2Island)(unsafe.Pointer(island)).TailBody = (*b2Body)(unsafe.Pointer(body)).IslandPrev
		}
	}
	if int32FromUint8(islandDestroyed) == false1 {
		b2ValidateIsland(tls, world, islandId)
	}
	(*b2Body)(unsafe.Pointer(body)).IslandId = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).IslandPrev = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).IslandNext = -int32(1)
}

func b2DestroyBodyContacts(tls *_Stack, world uintptr, body uintptr, wakeBodies uint8) {
	var contact, v1, v3 uintptr
	var contactId, edgeIndex, edgeKey, v2 int32
	_, _, _, _, _, _, _ = contact, contactId, edgeIndex, edgeKey, v1, v2, v3
	// Destroy the attached contacts
	edgeKey = (*b2Body)(unsafe.Pointer(body)).HeadContactKey
	for edgeKey != -int32(1) {
		contactId = edgeKey >> int32(1)
		edgeIndex = edgeKey & int32(1)
		v1 = world + 1144
		v2 = contactId
		if !(0 <= v2 && v2 < (*b2ContactArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ContactArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*68
		goto _4
	_4:
		contact = v3
		edgeKey = (*(*b2ContactEdge)(unsafe.Pointer(contact + 12 + uintptr(edgeIndex)*12))).NextKey
		b2DestroyContact(tls, world, contact, wakeBodies)
	}
	b2ValidateSolverSets(tls, world)
}

func b2CreateBody(tls *_Stack, worldId WorldId, def uintptr) (r BodyId) {
	var body, bodySim, bodyState, set, world, v1, v11, v13, v15, v17, v19, v3, v5, v7, v9, p21 uintptr
	var bodyId, i, newCapacity, newCapacity1, newCapacity2, newCapacity3, setId, v12, v16, v18, v2, v4, v8 int32
	var id BodyId
	var isAwake uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodyId, bodySim, bodyState, i, id, isAwake, newCapacity, newCapacity1, newCapacity2, newCapacity3, set, setId, world, v1, v11, v12, v13, v15, v16, v17, v18, v19, v2, v3, v4, v5, v7, v8, v9, p21
	if !((*BodyDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)) && b2InternalAssertFcn(tls, __ccgo_ts+730, __ccgo_ts+327, int32FromInt32(179)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsValidVec2(tls, (*BodyDef)(unsafe.Pointer(def)).Position) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+769, __ccgo_ts+327, int32FromInt32(180)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsValidRotation(tls, (*BodyDef)(unsafe.Pointer(def)).Rotation) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+800, __ccgo_ts+327, int32FromInt32(181)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsValidVec2(tls, (*BodyDef)(unsafe.Pointer(def)).LinearVelocity) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+835, __ccgo_ts+327, int32FromInt32(182)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).AngularVelocity) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+872, __ccgo_ts+327, int32FromInt32(183)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).LinearDamping) != 0 && (*BodyDef)(unsafe.Pointer(def)).LinearDamping >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+911, __ccgo_ts+327, int32FromInt32(184)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).AngularDamping) != 0 && (*BodyDef)(unsafe.Pointer(def)).AngularDamping >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+978, __ccgo_ts+327, int32FromInt32(185)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).SleepThreshold) != 0 && (*BodyDef)(unsafe.Pointer(def)).SleepThreshold >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+1047, __ccgo_ts+327, int32FromInt32(186)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).GravityScale) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+1116, __ccgo_ts+327, int32FromInt32(187)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorldFromId(tls, worldId)
	if !(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)) && b2InternalAssertFcn(tls, __ccgo_ts+1152, __ccgo_ts+327, int32FromInt32(190)) != 0 {
		__builtin_trap(tls)
	}
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return b2_nullBodyId
	}
	isAwake = boolUint8(((*BodyDef)(unsafe.Pointer(def)).IsAwake != 0 || int32FromUint8((*BodyDef)(unsafe.Pointer(def)).EnableSleep) == false1) && (*BodyDef)(unsafe.Pointer(def)).IsEnabled != 0)
	if int32FromUint8((*BodyDef)(unsafe.Pointer(def)).IsEnabled) == false1 {
		// any body type can be disabled
		setId = int32(b2_disabledSet)
	} else {
		if (*BodyDef)(unsafe.Pointer(def)).Type1 == int32(b2_staticBody) {
			setId = int32(b2_staticSet)
		} else {
			if int32FromUint8(isAwake) == int32(true1) {
				setId = int32(b2_awakeSet)
			} else {
				// new set for a sleeping body in its own island
				setId = b2AllocId(tls, world+1040)
				if setId == (*b2World)(unsafe.Pointer(world)).SolverSets.Count {
					// Create a zero initialized solver set. All sub-arrays are also zero initialized.
					v1 = world + 1064
					if (*b2SolverSetArray)(unsafe.Pointer(v1)).Count == (*b2SolverSetArray)(unsafe.Pointer(v1)).Capacity {
						if (*b2SolverSetArray)(unsafe.Pointer(v1)).Capacity < int32(2) {
							v2 = int32(2)
						} else {
							v2 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Capacity + (*b2SolverSetArray)(unsafe.Pointer(v1)).Capacity>>int32(1)
						}
						newCapacity3 = v2
						b2SolverSetArray_Reserve(tls, v1, newCapacity3)
					}
					*(*b2SolverSet)(unsafe.Pointer((*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr((*b2SolverSetArray)(unsafe.Pointer(v1)).Count)*88)) = b2SolverSet{}
					*(*int32)(unsafe.Pointer(v1 + 8)) += int32(1)
				} else {
					if !((*(*b2SolverSet)(unsafe.Pointer((*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setId)*88))).SetIndex == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+1175, __ccgo_ts+327, int32FromInt32(225)) != 0 {
						__builtin_trap(tls)
					}
				}
				(*(*b2SolverSet)(unsafe.Pointer((*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setId)*88))).SetIndex = setId
			}
		}
	}
	if !(0 <= setId && setId < (*b2World)(unsafe.Pointer(world)).SolverSets.Count) && b2InternalAssertFcn(tls, __ccgo_ts+1231, __ccgo_ts+327, int32FromInt32(231)) != 0 {
		__builtin_trap(tls)
	}
	bodyId = b2AllocId(tls, world+1000)
	v3 = world + 1064
	v4 = setId
	if !(0 <= v4 && v4 < (*b2SolverSetArray)(unsafe.Pointer(v3)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v5 = (*b2SolverSetArray)(unsafe.Pointer(v3)).Data + uintptr(v4)*88
	goto _6
_6:
	set = v5
	v7 = set
	if (*b2BodySimArray)(unsafe.Pointer(v7)).Count == (*b2BodySimArray)(unsafe.Pointer(v7)).Capacity {
		if (*b2BodySimArray)(unsafe.Pointer(v7)).Capacity < int32(2) {
			v8 = int32(2)
		} else {
			v8 = (*b2BodySimArray)(unsafe.Pointer(v7)).Capacity + (*b2BodySimArray)(unsafe.Pointer(v7)).Capacity>>int32(1)
		}
		newCapacity1 = v8
		b2BodySimArray_Reserve(tls, v7, newCapacity1)
	}
	*(*int32)(unsafe.Pointer(v7 + 8)) += int32(1)
	v9 = (*b2BodySimArray)(unsafe.Pointer(v7)).Data + uintptr((*b2BodySimArray)(unsafe.Pointer(v7)).Count-int32FromInt32(1))*100
	goto _10
_10:
	bodySim = v9
	*(*b2BodySim)(unsafe.Pointer(bodySim)) = b2BodySim{}
	(*b2BodySim)(unsafe.Pointer(bodySim)).Transform.P = (*BodyDef)(unsafe.Pointer(def)).Position
	(*b2BodySim)(unsafe.Pointer(bodySim)).Transform.Q = (*BodyDef)(unsafe.Pointer(def)).Rotation
	(*b2BodySim)(unsafe.Pointer(bodySim)).Center = (*BodyDef)(unsafe.Pointer(def)).Position
	(*b2BodySim)(unsafe.Pointer(bodySim)).Rotation0 = (*b2BodySim)(unsafe.Pointer(bodySim)).Transform.Q
	(*b2BodySim)(unsafe.Pointer(bodySim)).Center0 = (*b2BodySim)(unsafe.Pointer(bodySim)).Center
	(*b2BodySim)(unsafe.Pointer(bodySim)).MinExtent = float32(float32FromFloat32(100000) * b2_lengthUnitsPerMeter)
	(*b2BodySim)(unsafe.Pointer(bodySim)).MaxExtent = float32FromFloat32(0)
	(*b2BodySim)(unsafe.Pointer(bodySim)).LinearDamping = (*BodyDef)(unsafe.Pointer(def)).LinearDamping
	(*b2BodySim)(unsafe.Pointer(bodySim)).AngularDamping = (*BodyDef)(unsafe.Pointer(def)).AngularDamping
	(*b2BodySim)(unsafe.Pointer(bodySim)).GravityScale = (*BodyDef)(unsafe.Pointer(def)).GravityScale
	(*b2BodySim)(unsafe.Pointer(bodySim)).BodyId = bodyId
	(*b2BodySim)(unsafe.Pointer(bodySim)).IsBullet = (*BodyDef)(unsafe.Pointer(def)).IsBullet
	(*b2BodySim)(unsafe.Pointer(bodySim)).AllowFastRotation = (*BodyDef)(unsafe.Pointer(def)).AllowFastRotation
	if setId == int32(b2_awakeSet) {
		v11 = set + 16
		if (*b2BodyStateArray)(unsafe.Pointer(v11)).Count == (*b2BodyStateArray)(unsafe.Pointer(v11)).Capacity {
			if (*b2BodyStateArray)(unsafe.Pointer(v11)).Capacity < int32(2) {
				v12 = int32(2)
			} else {
				v12 = (*b2BodyStateArray)(unsafe.Pointer(v11)).Capacity + (*b2BodyStateArray)(unsafe.Pointer(v11)).Capacity>>int32(1)
			}
			newCapacity2 = v12
			b2BodyStateArray_Reserve(tls, v11, newCapacity2)
		}
		*(*int32)(unsafe.Pointer(v11 + 8)) += int32(1)
		v13 = (*b2BodyStateArray)(unsafe.Pointer(v11)).Data + uintptr((*b2BodyStateArray)(unsafe.Pointer(v11)).Count-int32FromInt32(1))*32
		goto _14
	_14:
		bodyState = v13
		if !(uint64(bodyState)&uint64FromInt32(0x1F) == uint64FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+1277, __ccgo_ts+327, int32FromInt32(255)) != 0 {
			__builtin_trap(tls)
		}
		*(*b2BodyState)(unsafe.Pointer(bodyState)) = b2BodyState{}
		(*b2BodyState)(unsafe.Pointer(bodyState)).LinearVelocity = (*BodyDef)(unsafe.Pointer(def)).LinearVelocity
		(*b2BodyState)(unsafe.Pointer(bodyState)).AngularVelocity = (*BodyDef)(unsafe.Pointer(def)).AngularVelocity
		(*b2BodyState)(unsafe.Pointer(bodyState)).DeltaRotation = b2Rot_identity
	}
	if bodyId == (*b2World)(unsafe.Pointer(world)).Bodies.Count {
		v15 = world + 1024
		if (*b2BodyArray)(unsafe.Pointer(v15)).Count == (*b2BodyArray)(unsafe.Pointer(v15)).Capacity {
			if (*b2BodyArray)(unsafe.Pointer(v15)).Capacity < int32(2) {
				v16 = int32(2)
			} else {
				v16 = (*b2BodyArray)(unsafe.Pointer(v15)).Capacity + (*b2BodyArray)(unsafe.Pointer(v15)).Capacity>>int32(1)
			}
			newCapacity = v16
			b2BodyArray_Reserve(tls, v15, newCapacity)
		}
		*(*b2Body)(unsafe.Pointer((*b2BodyArray)(unsafe.Pointer(v15)).Data + uintptr((*b2BodyArray)(unsafe.Pointer(v15)).Count)*128)) = b2Body{}
		*(*int32)(unsafe.Pointer(v15 + 8)) += int32(1)
	} else {
		if !((*(*b2Body)(unsafe.Pointer((*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(bodyId)*128))).Id == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+1314, __ccgo_ts+327, int32FromInt32(269)) != 0 {
			__builtin_trap(tls)
		}
	}
	v17 = world + 1024
	v18 = bodyId
	if !(0 <= v18 && v18 < (*b2BodyArray)(unsafe.Pointer(v17)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v19 = (*b2BodyArray)(unsafe.Pointer(v17)).Data + uintptr(v18)*128
	goto _20
_20:
	body = v19
	if (*BodyDef)(unsafe.Pointer(def)).Name != 0 {
		i = 0
		for i < int32(31) && int32FromUint8(*(*uint8)(unsafe.Pointer((*BodyDef)(unsafe.Pointer(def)).Name + uintptr(i)))) != 0 {
			*(*uint8)(unsafe.Pointer(body + uintptr(i))) = *(*uint8)(unsafe.Pointer((*BodyDef)(unsafe.Pointer(def)).Name + uintptr(i)))
			i = i + int32(1)
		}
		for i < int32(32) {
			*(*uint8)(unsafe.Pointer(body + uintptr(i))) = uint8(0)
			i = i + int32(1)
		}
	} else {
		memset(tls, body, 0, uint64FromInt32(32)*uint64FromInt64(1))
	}
	(*b2Body)(unsafe.Pointer(body)).UserData = (*BodyDef)(unsafe.Pointer(def)).UserData
	(*b2Body)(unsafe.Pointer(body)).SetIndex = setId
	(*b2Body)(unsafe.Pointer(body)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count - int32(1)
	p21 = body + 116
	*(*uint16_t)(unsafe.Pointer(p21)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(p21))) + int32FromInt32(1))
	(*b2Body)(unsafe.Pointer(body)).HeadShapeId = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).ShapeCount = 0
	(*b2Body)(unsafe.Pointer(body)).HeadChainId = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).HeadContactKey = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).ContactCount = 0
	(*b2Body)(unsafe.Pointer(body)).HeadJointKey = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).JointCount = 0
	(*b2Body)(unsafe.Pointer(body)).IslandId = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).IslandPrev = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).IslandNext = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).BodyMoveIndex = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).Id = bodyId
	(*b2Body)(unsafe.Pointer(body)).Mass = float32FromFloat32(0)
	(*b2Body)(unsafe.Pointer(body)).Inertia = float32FromFloat32(0)
	(*b2Body)(unsafe.Pointer(body)).SleepThreshold = (*BodyDef)(unsafe.Pointer(def)).SleepThreshold
	(*b2Body)(unsafe.Pointer(body)).SleepTime = float32FromFloat32(0)
	(*b2Body)(unsafe.Pointer(body)).Type1 = (*BodyDef)(unsafe.Pointer(def)).Type1
	(*b2Body)(unsafe.Pointer(body)).EnableSleep = (*BodyDef)(unsafe.Pointer(def)).EnableSleep
	(*b2Body)(unsafe.Pointer(body)).FixedRotation = (*BodyDef)(unsafe.Pointer(def)).FixedRotation
	(*b2Body)(unsafe.Pointer(body)).IsSpeedCapped = boolUint8(false1 != 0)
	(*b2Body)(unsafe.Pointer(body)).IsMarked = boolUint8(false1 != 0)
	// dynamic and kinematic bodies that are enabled need a island
	if setId >= int32(b2_awakeSet) {
		b2CreateIslandForBody(tls, world, setId, body)
	}
	b2ValidateSolverSets(tls, world)
	id = BodyId{
		Index1:     bodyId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Body)(unsafe.Pointer(body)).Generation,
	}
	return id
}

func b2WakeBody(tls *_Stack, world uintptr, body uintptr) (r uint8) {
	if (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeSolverSet(tls, world, (*b2Body)(unsafe.Pointer(body)).SetIndex)
		return boolUint8(true1 != 0)
	}
	return boolUint8(false1 != 0)
}

func b2DestroyBody(tls *_Stack, bodyId BodyId) {
	var body, chain, joint, movedBody, movedSim, set, shape, world, v1, v11, v13, v15, v17, v21, v23, v25, v3, v5, v7, v9 uintptr
	var chainId, edgeIndex, edgeKey, jointId, movedId, movedIndex, movedIndex1, movedIndex2, result, shapeId, v10, v14, v18, v19, v2, v22, v26, v27, v6 int32
	var wakeBodies uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, chain, chainId, edgeIndex, edgeKey, joint, jointId, movedBody, movedId, movedIndex, movedIndex1, movedIndex2, movedSim, result, set, shape, shapeId, wakeBodies, world, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v3, v5, v6, v7, v9
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	// Wake bodies attached to this body, even if this body is static.
	wakeBodies = boolUint8(true1 != 0)
	// Destroy the attached joints
	edgeKey = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
	for edgeKey != -int32(1) {
		jointId = edgeKey >> int32(1)
		edgeIndex = edgeKey & int32(1)
		v1 = world + 1104
		v2 = jointId
		if !(0 <= v2 && v2 < (*b2JointArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2JointArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*72
		goto _4
	_4:
		joint = v3
		edgeKey = (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + uintptr(edgeIndex)*12))).NextKey
		// Careful because this modifies the list being traversed
		b2DestroyJointInternal(tls, world, joint, wakeBodies)
	}
	// Destroy all contacts attached to this body.
	b2DestroyBodyContacts(tls, world, body, wakeBodies)
	// Destroy the attached shapes and their broad-phase proxies.
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		v5 = world + 1248
		v6 = shapeId
		if !(0 <= v6 && v6 < (*b2ShapeArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v7 = (*b2ShapeArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*288
		goto _8
	_8:
		shape = v7
		if (*b2Shape)(unsafe.Pointer(shape)).SensorIndex != -int32(1) {
			b2DestroySensor(tls, world, shape)
		}
		b2DestroyShapeProxy(tls, shape, world+40)
		// Return shape to free list.
		b2FreeId(tls, world+1200, shapeId)
		(*b2Shape)(unsafe.Pointer(shape)).Id = -int32(1)
		shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
	}
	// Destroy the attached chains. The associated shapes have already been destroyed above.
	chainId = (*b2Body)(unsafe.Pointer(body)).HeadChainId
	for chainId != -int32(1) {
		v9 = world + 1264
		v10 = chainId
		if !(0 <= v10 && v10 < (*b2ChainShapeArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(137)) != 0 {
			__builtin_trap(tls)
		}
		v11 = (*b2ChainShapeArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*48
		goto _12
	_12:
		chain = v11
		b2FreeChainData(tls, chain)
		// Return chain to free list.
		b2FreeId(tls, world+1224, chainId)
		(*b2ChainShape)(unsafe.Pointer(chain)).Id = -int32(1)
		chainId = (*b2ChainShape)(unsafe.Pointer(chain)).NextChainId
	}
	b2RemoveBodyFromIsland(tls, world, body)
	v13 = world + 1064
	v14 = (*b2Body)(unsafe.Pointer(body)).SetIndex
	if !(0 <= v14 && v14 < (*b2SolverSetArray)(unsafe.Pointer(v13)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v15 = (*b2SolverSetArray)(unsafe.Pointer(v13)).Data + uintptr(v14)*88
	goto _16
_16:
	// Remove body sim from solver set that owns it
	set = v15
	v17 = set
	v18 = (*b2Body)(unsafe.Pointer(body)).LocalIndex
	if !(0 <= v18 && v18 < (*b2BodySimArray)(unsafe.Pointer(v17)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	movedIndex = -int32(1)
	if v18 != (*b2BodySimArray)(unsafe.Pointer(v17)).Count-int32FromInt32(1) {
		movedIndex = (*b2BodySimArray)(unsafe.Pointer(v17)).Count - int32(1)
		*(*b2BodySim)(unsafe.Pointer((*b2BodySimArray)(unsafe.Pointer(v17)).Data + uintptr(v18)*100)) = *(*b2BodySim)(unsafe.Pointer((*b2BodySimArray)(unsafe.Pointer(v17)).Data + uintptr(movedIndex)*100))
	}
	*(*int32)(unsafe.Pointer(v17 + 8)) -= int32(1)
	v19 = movedIndex
	goto _20
_20:
	movedIndex2 = v19
	if movedIndex2 != -int32(1) {
		// Fix moved body index
		movedSim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(body)).LocalIndex)*100
		movedId = (*b2BodySim)(unsafe.Pointer(movedSim)).BodyId
		v21 = world + 1024
		v22 = movedId
		if !(0 <= v22 && v22 < (*b2BodyArray)(unsafe.Pointer(v21)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
			__builtin_trap(tls)
		}
		v23 = (*b2BodyArray)(unsafe.Pointer(v21)).Data + uintptr(v22)*128
		goto _24
	_24:
		movedBody = v23
		if !((*b2Body)(unsafe.Pointer(movedBody)).LocalIndex == movedIndex2) && b2InternalAssertFcn(tls, __ccgo_ts+1407, __ccgo_ts+327, int32FromInt32(419)) != 0 {
			__builtin_trap(tls)
		}
		(*b2Body)(unsafe.Pointer(movedBody)).LocalIndex = (*b2Body)(unsafe.Pointer(body)).LocalIndex
	}
	// Remove body state from awake set
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
		v25 = set + 16
		v26 = (*b2Body)(unsafe.Pointer(body)).LocalIndex
		if !(0 <= v26 && v26 < (*b2BodyStateArray)(unsafe.Pointer(v25)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(194)) != 0 {
			__builtin_trap(tls)
		}
		movedIndex1 = -int32(1)
		if v26 != (*b2BodyStateArray)(unsafe.Pointer(v25)).Count-int32FromInt32(1) {
			movedIndex1 = (*b2BodyStateArray)(unsafe.Pointer(v25)).Count - int32(1)
			*(*b2BodyState)(unsafe.Pointer((*b2BodyStateArray)(unsafe.Pointer(v25)).Data + uintptr(v26)*32)) = *(*b2BodyState)(unsafe.Pointer((*b2BodyStateArray)(unsafe.Pointer(v25)).Data + uintptr(movedIndex1)*32))
		}
		*(*int32)(unsafe.Pointer(v25 + 8)) -= int32(1)
		v27 = movedIndex1
		goto _28
	_28:
		result = v27
		if !(result == movedIndex2) && b2InternalAssertFcn(tls, __ccgo_ts+1443, __ccgo_ts+327, int32FromInt32(427)) != 0 {
			__builtin_trap(tls)
		}
		_ = uint64FromInt64(4)
	} else {
		if (*b2SolverSet)(unsafe.Pointer(set)).SetIndex >= int32(b2_firstSleepingSet) && (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count == 0 {
			// Remove solver set if it's now an orphan.
			b2DestroySolverSet(tls, world, (*b2SolverSet)(unsafe.Pointer(set)).SetIndex)
		}
	}
	// Free body and id (preserve body generation)
	b2FreeId(tls, world+1000, (*b2Body)(unsafe.Pointer(body)).Id)
	(*b2Body)(unsafe.Pointer(body)).SetIndex = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).LocalIndex = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).Id = -int32(1)
	b2ValidateSolverSets(tls, world)
}

func b2Body_GetContactCapacity(tls *_Stack, bodyId BodyId) (r int32) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	// Conservative and fast
	return (*b2Body)(unsafe.Pointer(body)).ContactCount
}

func b2Body_GetContactData(tls *_Stack, bodyId BodyId, contactData uintptr, capacity int32) (r int32) {
	var body, contact, contactSim, shapeA, shapeB, world, v1, v11, v3, v5, v7, v9 uintptr
	var contactId, contactKey, edgeIndex, index2, v10, v2, v6 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, contact, contactId, contactKey, contactSim, edgeIndex, index2, shapeA, shapeB, world, v1, v10, v11, v2, v3, v5, v6, v7, v9
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	contactKey = (*b2Body)(unsafe.Pointer(body)).HeadContactKey
	index2 = 0
	for contactKey != -int32(1) && index2 < capacity {
		contactId = contactKey >> int32(1)
		edgeIndex = contactKey & int32(1)
		v1 = world + 1144
		v2 = contactId
		if !(0 <= v2 && v2 < (*b2ContactArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+705, int32FromInt32(146)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ContactArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*68
		goto _4
	_4:
		contact = v3
		// Is contact touching?
		if (*b2Contact)(unsafe.Pointer(contact)).Flags&uint32(b2_contactTouchingFlag) != 0 {
			v5 = world + 1248
			v6 = (*b2Contact)(unsafe.Pointer(contact)).ShapeIdA
			if !(0 <= v6 && v6 < (*b2ShapeArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
				__builtin_trap(tls)
			}
			v7 = (*b2ShapeArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*288
			goto _8
		_8:
			shapeA = v7
			v9 = world + 1248
			v10 = (*b2Contact)(unsafe.Pointer(contact)).ShapeIdB
			if !(0 <= v10 && v10 < (*b2ShapeArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
				__builtin_trap(tls)
			}
			v11 = (*b2ShapeArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*288
			goto _12
		_12:
			shapeB = v11
			(*(*ContactData)(unsafe.Pointer(contactData + uintptr(index2)*128))).ShapeIdA = ShapeId{
				Index1:     (*b2Shape)(unsafe.Pointer(shapeA)).Id + int32(1),
				World0:     bodyId.World0,
				Generation: (*b2Shape)(unsafe.Pointer(shapeA)).Generation,
			}
			(*(*ContactData)(unsafe.Pointer(contactData + uintptr(index2)*128))).ShapeIdB = ShapeId{
				Index1:     (*b2Shape)(unsafe.Pointer(shapeB)).Id + int32(1),
				World0:     bodyId.World0,
				Generation: (*b2Shape)(unsafe.Pointer(shapeB)).Generation,
			}
			contactSim = b2GetContactSim(tls, world, contact)
			(*(*ContactData)(unsafe.Pointer(contactData + uintptr(index2)*128))).Manifold = (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold
			index2 = index2 + int32(1)
		}
		contactKey = (*(*b2ContactEdge)(unsafe.Pointer(contact + 12 + uintptr(edgeIndex)*12))).NextKey
	}
	if !(index2 <= capacity) && b2InternalAssertFcn(tls, __ccgo_ts+1464, __ccgo_ts+327, int32FromInt32(497)) != 0 {
		__builtin_trap(tls)
	}
	return index2
}

func b2Body_ComputeAABB(tls *_Stack, bodyId BodyId) (r AABB) {
	var aabb, c, v10, v31, v9 AABB
	var body, shape, world, v1, v3, v5, v7 uintptr
	var transform Transform
	var v11, v12, v13, v15, v16, v17, v18, v20, v21, v22, v23, v25, v26, v27, v28, v30 float32
	var v2, v6 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, body, c, shape, transform, world, v1, v10, v11, v12, v13, v15, v16, v17, v18, v2, v20, v21, v22, v23, v25, v26, v27, v28, v3, v30, v31, v5, v6, v7, v9
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return AABB{}
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).HeadShapeId == -int32(1) {
		transform = b2GetBodyTransform(tls, world, (*b2Body)(unsafe.Pointer(body)).Id)
		return AABB{
			LowerBound: transform.P,
			UpperBound: transform.P,
		}
	}
	v1 = world + 1248
	v2 = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
	goto _4
_4:
	shape = v3
	aabb = (*b2Shape)(unsafe.Pointer(shape)).Aabb
	for (*b2Shape)(unsafe.Pointer(shape)).NextShapeId != -int32(1) {
		v5 = world + 1248
		v6 = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
		if !(0 <= v6 && v6 < (*b2ShapeArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v7 = (*b2ShapeArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*288
		goto _8
	_8:
		shape = v7
		v9 = aabb
		v10 = (*b2Shape)(unsafe.Pointer(shape)).Aabb
		v11 = v9.LowerBound.X
		v12 = v10.LowerBound.X
		if v11 < v12 {
			v15 = v11
		} else {
			v15 = v12
		}
		v13 = v15
		goto _14
	_14:
		c.LowerBound.X = v13
		v16 = v9.LowerBound.Y
		v17 = v10.LowerBound.Y
		if v16 < v17 {
			v20 = v16
		} else {
			v20 = v17
		}
		v18 = v20
		goto _19
	_19:
		c.LowerBound.Y = v18
		v21 = v9.UpperBound.X
		v22 = v10.UpperBound.X
		if v21 > v22 {
			v25 = v21
		} else {
			v25 = v22
		}
		v23 = v25
		goto _24
	_24:
		c.UpperBound.X = v23
		v26 = v9.UpperBound.Y
		v27 = v10.UpperBound.Y
		if v26 > v27 {
			v30 = v26
		} else {
			v30 = v27
		}
		v28 = v30
		goto _29
	_29:
		c.UpperBound.Y = v28
		v31 = c
		goto _32
	_32:
		aabb = v31
	}
	return aabb
}

func b2UpdateBodyMassData(tls *_Stack, world uintptr, body uintptr) {
	var bodySim, s3, s4, s5, state, v1, v15, v17, v3, v48, v50 uintptr
	var deltaLinear, localCenter, oldCenter, v19, v21, v22, v25, v26, v28, v29, v33, v34, v36, v37, v38, v41, v42, v44, v45, v46 Vec2
	var extent, extent1 b2ShapeExtent
	var massData MassData
	var shapeId, shapeId1, v16, v2, v49 int32
	var x, y, v10, v11, v12, v14, v20, v24, v30, v40, v5, v52, v53, v54, v56, v57, v58, v59, v6, v61, v7, v9 float32
	var v32 Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodySim, deltaLinear, extent, extent1, localCenter, massData, oldCenter, s3, s4, s5, shapeId, shapeId1, state, x, y, v1, v10, v11, v12, v14, v15, v16, v17, v19, v2, v20, v21, v22, v24, v25, v26, v28, v29, v3, v30, v32, v33, v34, v36, v37, v38, v40, v41, v42, v44, v45, v46, v48, v49, v5, v50, v52, v53, v54, v56, v57, v58, v59, v6, v61, v7, v9
	bodySim = b2GetBodySim(tls, world, body)
	// Compute mass data from shapes. Each shape has its own density.
	(*b2Body)(unsafe.Pointer(body)).Mass = float32FromFloat32(0)
	(*b2Body)(unsafe.Pointer(body)).Inertia = float32FromFloat32(0)
	(*b2BodySim)(unsafe.Pointer(bodySim)).InvMass = float32FromFloat32(0)
	(*b2BodySim)(unsafe.Pointer(bodySim)).InvInertia = float32FromFloat32(0)
	(*b2BodySim)(unsafe.Pointer(bodySim)).LocalCenter = b2Vec2_zero
	(*b2BodySim)(unsafe.Pointer(bodySim)).MinExtent = float32(float32FromFloat32(100000) * b2_lengthUnitsPerMeter)
	(*b2BodySim)(unsafe.Pointer(bodySim)).MaxExtent = float32FromFloat32(0)
	// Static and kinematic sims have zero mass.
	if (*b2Body)(unsafe.Pointer(body)).Type1 != int32(b2_dynamicBody) {
		(*b2BodySim)(unsafe.Pointer(bodySim)).Center = (*b2BodySim)(unsafe.Pointer(bodySim)).Transform.P
		// Need extents for kinematic bodies for sleeping to work correctly.
		if (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_kinematicBody) {
			shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
			for shapeId != -int32(1) {
				v1 = world + 1248
				v2 = shapeId
				if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
					__builtin_trap(tls)
				}
				v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
				goto _4
			_4:
				s3 = v3
				extent = b2ComputeShapeExtent(tls, s3, b2Vec2_zero)
				v5 = (*b2BodySim)(unsafe.Pointer(bodySim)).MinExtent
				v6 = extent.MinExtent
				if v5 < v6 {
					v9 = v5
				} else {
					v9 = v6
				}
				v7 = v9
				goto _8
			_8:
				(*b2BodySim)(unsafe.Pointer(bodySim)).MinExtent = v7
				v10 = (*b2BodySim)(unsafe.Pointer(bodySim)).MaxExtent
				v11 = extent.MaxExtent
				if v10 > v11 {
					v14 = v10
				} else {
					v14 = v11
				}
				v12 = v14
				goto _13
			_13:
				(*b2BodySim)(unsafe.Pointer(bodySim)).MaxExtent = v12
				shapeId = (*b2Shape)(unsafe.Pointer(s3)).NextShapeId
			}
		}
		return
	}
	// Accumulate mass over all shapes.
	localCenter = b2Vec2_zero
	shapeId1 = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId1 != -int32(1) {
		v15 = world + 1248
		v16 = shapeId1
		if !(0 <= v16 && v16 < (*b2ShapeArray)(unsafe.Pointer(v15)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v17 = (*b2ShapeArray)(unsafe.Pointer(v15)).Data + uintptr(v16)*288
		goto _18
	_18:
		s4 = v17
		shapeId1 = (*b2Shape)(unsafe.Pointer(s4)).NextShapeId
		if (*b2Shape)(unsafe.Pointer(s4)).Density == float32FromFloat32(0) {
			continue
		}
		massData = b2ComputeShapeMass(tls, s4)
		*(*float32)(unsafe.Pointer(body + 88)) += massData.Mass
		v19 = localCenter
		v20 = massData.Mass
		v21 = massData.Center
		v22 = Vec2{
			X: v19.X + float32(v20*v21.X),
			Y: v19.Y + float32(v20*v21.Y),
		}
		goto _23
	_23:
		localCenter = v22
		*(*float32)(unsafe.Pointer(body + 92)) += massData.RotationalInertia
	}
	// Compute center of mass.
	if (*b2Body)(unsafe.Pointer(body)).Mass > float32FromFloat32(0) {
		(*b2BodySim)(unsafe.Pointer(bodySim)).InvMass = float32FromFloat32(1) / (*b2Body)(unsafe.Pointer(body)).Mass
		v24 = (*b2BodySim)(unsafe.Pointer(bodySim)).InvMass
		v25 = localCenter
		v26 = Vec2{
			X: float32(v24 * v25.X),
			Y: float32(v24 * v25.Y),
		}
		goto _27
	_27:
		localCenter = v26
	}
	if (*b2Body)(unsafe.Pointer(body)).Inertia > float32FromFloat32(0) && int32FromUint8((*b2Body)(unsafe.Pointer(body)).FixedRotation) == false1 {
		// Center the inertia about the center of mass.
		v28 = localCenter
		v29 = localCenter
		v30 = float32(v28.X*v29.X) + float32(v28.Y*v29.Y)
		goto _31
	_31:
		*(*float32)(unsafe.Pointer(body + 92)) -= float32((*b2Body)(unsafe.Pointer(body)).Mass * v30)
		if !((*b2Body)(unsafe.Pointer(body)).Inertia > float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+1482, __ccgo_ts+327, int32FromInt32(596)) != 0 {
			__builtin_trap(tls)
		}
		(*b2BodySim)(unsafe.Pointer(bodySim)).InvInertia = float32FromFloat32(1) / (*b2Body)(unsafe.Pointer(body)).Inertia
	} else {
		(*b2Body)(unsafe.Pointer(body)).Inertia = float32FromFloat32(0)
		(*b2BodySim)(unsafe.Pointer(bodySim)).InvInertia = float32FromFloat32(0)
	}
	// Move center of mass.
	oldCenter = (*b2BodySim)(unsafe.Pointer(bodySim)).Center
	(*b2BodySim)(unsafe.Pointer(bodySim)).LocalCenter = localCenter
	v32 = (*b2BodySim)(unsafe.Pointer(bodySim)).Transform
	v33 = (*b2BodySim)(unsafe.Pointer(bodySim)).LocalCenter
	x = float32(v32.Q.C*v33.X) - float32(v32.Q.S*v33.Y) + v32.P.X
	y = float32(v32.Q.S*v33.X) + float32(v32.Q.C*v33.Y) + v32.P.Y
	v34 = Vec2{
		X: x,
		Y: y,
	}
	goto _35
_35:
	(*b2BodySim)(unsafe.Pointer(bodySim)).Center = v34
	(*b2BodySim)(unsafe.Pointer(bodySim)).Center0 = (*b2BodySim)(unsafe.Pointer(bodySim)).Center
	// Update center of mass velocity
	state = b2GetBodyState(tls, world, body)
	if state != uintptrFromInt32(0) {
		v36 = (*b2BodySim)(unsafe.Pointer(bodySim)).Center
		v37 = oldCenter
		v38 = Vec2{
			X: v36.X - v37.X,
			Y: v36.Y - v37.Y,
		}
		goto _39
	_39:
		v40 = (*b2BodyState)(unsafe.Pointer(state)).AngularVelocity
		v41 = v38
		v42 = Vec2{
			X: float32(-v40 * v41.Y),
			Y: float32(v40 * v41.X),
		}
		goto _43
	_43:
		deltaLinear = v42
		v44 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
		v45 = deltaLinear
		v46 = Vec2{
			X: v44.X + v45.X,
			Y: v44.Y + v45.Y,
		}
		goto _47
	_47:
		(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity = v46
	}
	// Compute body extents relative to center of mass
	shapeId1 = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId1 != -int32(1) {
		v48 = world + 1248
		v49 = shapeId1
		if !(0 <= v49 && v49 < (*b2ShapeArray)(unsafe.Pointer(v48)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v50 = (*b2ShapeArray)(unsafe.Pointer(v48)).Data + uintptr(v49)*288
		goto _51
	_51:
		s5 = v50
		extent1 = b2ComputeShapeExtent(tls, s5, localCenter)
		v52 = (*b2BodySim)(unsafe.Pointer(bodySim)).MinExtent
		v53 = extent1.MinExtent
		if v52 < v53 {
			v56 = v52
		} else {
			v56 = v53
		}
		v54 = v56
		goto _55
	_55:
		(*b2BodySim)(unsafe.Pointer(bodySim)).MinExtent = v54
		v57 = (*b2BodySim)(unsafe.Pointer(bodySim)).MaxExtent
		v58 = extent1.MaxExtent
		if v57 > v58 {
			v61 = v57
		} else {
			v61 = v58
		}
		v59 = v61
		goto _60
	_60:
		(*b2BodySim)(unsafe.Pointer(bodySim)).MaxExtent = v59
		shapeId1 = (*b2Shape)(unsafe.Pointer(s5)).NextShapeId
	}
}

func b2Body_GetPosition(tls *_Stack, bodyId BodyId) (r Vec2) {
	var body, world uintptr
	var transform Transform
	_, _, _ = body, transform, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	transform = b2GetBodyTransformQuick(tls, world, body)
	return transform.P
}

func b2Body_GetRotation(tls *_Stack, bodyId BodyId) (r Rot) {
	var body, world uintptr
	var transform Transform
	_, _, _ = body, transform, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	transform = b2GetBodyTransformQuick(tls, world, body)
	return transform.Q
}

func b2Body_GetTransform(tls *_Stack, bodyId BodyId) (r Transform) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return b2GetBodyTransformQuick(tls, world, body)
}

func b2Body_GetLocalPoint(tls *_Stack, bodyId BodyId, worldPoint Vec2) (r Vec2) {
	var body, world uintptr
	var transform, v1 Transform
	var vx, vy float32
	var v2, v3 Vec2
	_, _, _, _, _, _, _, _ = body, transform, vx, vy, world, v1, v2, v3
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	transform = b2GetBodyTransformQuick(tls, world, body)
	v1 = transform
	v2 = worldPoint
	vx = v2.X - v1.P.X
	vy = v2.Y - v1.P.Y
	v3 = Vec2{
		X: float32(v1.Q.C*vx) + float32(v1.Q.S*vy),
		Y: float32(-v1.Q.S*vx) + float32(v1.Q.C*vy),
	}
	goto _4
_4:
	return v3
}

func b2Body_GetWorldPoint(tls *_Stack, bodyId BodyId, localPoint Vec2) (r Vec2) {
	var body, world uintptr
	var transform, v1 Transform
	var x, y float32
	var v2, v3 Vec2
	_, _, _, _, _, _, _, _ = body, transform, world, x, y, v1, v2, v3
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	transform = b2GetBodyTransformQuick(tls, world, body)
	v1 = transform
	v2 = localPoint
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	return v3
}

func b2Body_GetLocalVector(tls *_Stack, bodyId BodyId, worldVector Vec2) (r Vec2) {
	var body, world uintptr
	var transform Transform
	var v1 Rot
	var v2, v3 Vec2
	_, _, _, _, _, _ = body, transform, world, v1, v2, v3
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	transform = b2GetBodyTransformQuick(tls, world, body)
	v1 = transform.Q
	v2 = worldVector
	v3 = Vec2{
		X: float32(v1.C*v2.X) + float32(v1.S*v2.Y),
		Y: float32(-v1.S*v2.X) + float32(v1.C*v2.Y),
	}
	goto _4
_4:
	return v3
}

func b2Body_GetWorldVector(tls *_Stack, bodyId BodyId, localVector Vec2) (r Vec2) {
	var body, world uintptr
	var transform Transform
	var v1 Rot
	var v2, v3 Vec2
	_, _, _, _, _, _ = body, transform, world, v1, v2, v3
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	transform = b2GetBodyTransformQuick(tls, world, body)
	v1 = transform.Q
	v2 = localVector
	v3 = Vec2{
		X: float32(v1.C*v2.X) - float32(v1.S*v2.Y),
		Y: float32(v1.S*v2.X) + float32(v1.C*v2.Y),
	}
	goto _4
_4:
	return v3
}

func b2Body_SetTransform(tls *_Stack, bodyId BodyId, position Vec2, rotation Rot) {
	var aabb, fatAABB, v10, v9 AABB
	var body, bodySim, broadPhase, shape, world, v5, v7 uintptr
	var margin, speculativeDistance, x, y float32
	var s, v11 uint8
	var shapeId, v6 int32
	var transform, v1 Transform
	var v2, v3 Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, body, bodySim, broadPhase, fatAABB, margin, s, shape, shapeId, speculativeDistance, transform, world, x, y, v1, v10, v11, v2, v3, v5, v6, v7, v9
	if !(b2IsValidVec2(tls, position) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+1503, __ccgo_ts+327, int32FromInt32(690)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsValidRotation(tls, rotation) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+1529, __ccgo_ts+327, int32FromInt32(691)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2Body_IsValid(tls, bodyId) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+302, __ccgo_ts+327, int32FromInt32(692)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if !(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)) && b2InternalAssertFcn(tls, __ccgo_ts+1152, __ccgo_ts+327, int32FromInt32(694)) != 0 {
		__builtin_trap(tls)
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	(*b2BodySim)(unsafe.Pointer(bodySim)).Transform.P = position
	(*b2BodySim)(unsafe.Pointer(bodySim)).Transform.Q = rotation
	v1 = (*b2BodySim)(unsafe.Pointer(bodySim)).Transform
	v2 = (*b2BodySim)(unsafe.Pointer(bodySim)).LocalCenter
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	(*b2BodySim)(unsafe.Pointer(bodySim)).Center = v3
	(*b2BodySim)(unsafe.Pointer(bodySim)).Rotation0 = (*b2BodySim)(unsafe.Pointer(bodySim)).Transform.Q
	(*b2BodySim)(unsafe.Pointer(bodySim)).Center0 = (*b2BodySim)(unsafe.Pointer(bodySim)).Center
	broadPhase = world + 40
	transform = (*b2BodySim)(unsafe.Pointer(bodySim)).Transform
	margin = float32(float32FromFloat32(0.05) * b2_lengthUnitsPerMeter)
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		v5 = world + 1248
		v6 = shapeId
		if !(0 <= v6 && v6 < (*b2ShapeArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v7 = (*b2ShapeArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*288
		goto _8
	_8:
		shape = v7
		aabb = b2ComputeShapeAABB(tls, shape, transform)
		aabb.LowerBound.X -= speculativeDistance
		aabb.LowerBound.Y -= speculativeDistance
		aabb.UpperBound.X += speculativeDistance
		aabb.UpperBound.Y += speculativeDistance
		(*b2Shape)(unsafe.Pointer(shape)).Aabb = aabb
		v9 = (*b2Shape)(unsafe.Pointer(shape)).FatAABB
		v10 = aabb
		s = boolUint8(true1 != 0)
		s = boolUint8(s != 0 && v9.LowerBound.X <= v10.LowerBound.X)
		s = boolUint8(s != 0 && v9.LowerBound.Y <= v10.LowerBound.Y)
		s = boolUint8(s != 0 && v10.UpperBound.X <= v9.UpperBound.X)
		s = boolUint8(s != 0 && v10.UpperBound.Y <= v9.UpperBound.Y)
		v11 = s
		goto _12
	_12:
		if int32FromUint8(v11) == false1 {
			fatAABB.LowerBound.X = aabb.LowerBound.X - margin
			fatAABB.LowerBound.Y = aabb.LowerBound.Y - margin
			fatAABB.UpperBound.X = aabb.UpperBound.X + margin
			fatAABB.UpperBound.Y = aabb.UpperBound.Y + margin
			(*b2Shape)(unsafe.Pointer(shape)).FatAABB = fatAABB
			// They body could be disabled
			if (*b2Shape)(unsafe.Pointer(shape)).ProxyKey != -int32(1) {
				b2BroadPhase_MoveProxy(tls, broadPhase, (*b2Shape)(unsafe.Pointer(shape)).ProxyKey, fatAABB)
			}
		}
		shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
	}
}

func b2Body_GetLinearVelocity(tls *_Stack, bodyId BodyId) (r Vec2) {
	var body, state, world uintptr
	_, _, _ = body, state, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	state = b2GetBodyState(tls, world, body)
	if state != uintptrFromInt32(0) {
		return (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
	}
	return b2Vec2_zero
}

func b2Body_GetAngularVelocity(tls *_Stack, bodyId BodyId) (r float32) {
	var body, state, world uintptr
	_, _, _ = body, state, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	state = b2GetBodyState(tls, world, body)
	if state != uintptrFromInt32(0) {
		return (*b2BodyState)(unsafe.Pointer(state)).AngularVelocity
	}
	return float32(0)
}

func b2Body_SetLinearVelocity(tls *_Stack, bodyId BodyId, linearVelocity Vec2) {
	var body, state, world uintptr
	var v1 Vec2
	var v2 float32
	_, _, _, _, _ = body, state, world, v1, v2
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_staticBody) {
		return
	}
	v1 = linearVelocity
	v2 = float32(v1.X*v1.X) + float32(v1.Y*v1.Y)
	goto _3
_3:
	if v2 > float32FromFloat32(0) {
		b2WakeBody(tls, world, body)
	}
	state = b2GetBodyState(tls, world, body)
	if state == uintptrFromInt32(0) {
		return
	}
	(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity = linearVelocity
}

func b2Body_SetAngularVelocity(tls *_Stack, bodyId BodyId, angularVelocity float32) {
	var body, state, world uintptr
	_, _, _ = body, state, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_staticBody) || (*b2Body)(unsafe.Pointer(body)).FixedRotation != 0 {
		return
	}
	if angularVelocity != float32FromFloat32(0) {
		b2WakeBody(tls, world, body)
	}
	state = b2GetBodyState(tls, world, body)
	if state == uintptrFromInt32(0) {
		return
	}
	(*b2BodyState)(unsafe.Pointer(state)).AngularVelocity = angularVelocity
}

func b2Body_SetTargetTransform(tls *_Stack, bodyId BodyId, target Transform, timeStep float32) {
	var angularVelocity, c, deltaAngle, invTimeStep, maxVelocity, s1, x, y, v15, v18, v20, v21, v23, v9 float32
	var body, sim, state, world uintptr
	var center1, center2, linearVelocity, v10, v11, v17, v2, v3, v5, v6, v7 Vec2
	var q1, q2, v13, v14 Rot
	var v1 Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angularVelocity, body, c, center1, center2, deltaAngle, invTimeStep, linearVelocity, maxVelocity, q1, q2, s1, sim, state, world, x, y, v1, v10, v11, v13, v14, v15, v17, v18, v2, v20, v21, v23, v3, v5, v6, v7, v9
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_staticBody) || timeStep <= float32FromFloat32(0) {
		return
	}
	sim = b2GetBodySim(tls, world, body)
	// Compute linear velocity
	center1 = (*b2BodySim)(unsafe.Pointer(sim)).Center
	v1 = target
	v2 = (*b2BodySim)(unsafe.Pointer(sim)).LocalCenter
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	center2 = v3
	invTimeStep = float32FromFloat32(1) / timeStep
	v5 = center2
	v6 = center1
	v7 = Vec2{
		X: v5.X - v6.X,
		Y: v5.Y - v6.Y,
	}
	goto _8
_8:
	v9 = invTimeStep
	v10 = v7
	v11 = Vec2{
		X: float32(v9 * v10.X),
		Y: float32(v9 * v10.Y),
	}
	goto _12
_12:
	linearVelocity = v11
	// Compute angular velocity
	angularVelocity = float32FromFloat32(0)
	if int32FromUint8((*b2Body)(unsafe.Pointer(body)).FixedRotation) == false1 {
		q1 = (*b2BodySim)(unsafe.Pointer(sim)).Transform.Q
		q2 = target.Q
		v13 = q2
		v14 = q1
		s1 = float32(v13.S*v14.C) - float32(v13.C*v14.S)
		c = float32(v13.C*v14.C) + float32(v13.S*v14.S)
		v15 = b2Atan2(tls, s1, c)
		goto _16
	_16:
		deltaAngle = v15
		angularVelocity = float32(invTimeStep * deltaAngle)
	}
	v17 = linearVelocity
	v18 = sqrtf(tls, float32(v17.X*v17.X)+float32(v17.Y*v17.Y))
	goto _19
_19:
	v20 = angularVelocity
	if v20 < float32FromInt32(0) {
		v23 = -v20
	} else {
		v23 = v20
	}
	v21 = v23
	goto _22
_22:
	maxVelocity = v18 + float32(v21*(*b2BodySim)(unsafe.Pointer(sim)).MaxExtent)
	// Return if velocity would be sleepy
	if maxVelocity < (*b2Body)(unsafe.Pointer(body)).SleepThreshold {
		return
	}
	// Must wake for state to exist
	b2WakeBody(tls, world, body)
	state = b2GetBodyState(tls, world, body)
	if state == uintptrFromInt32(0) {
		return
	}
	(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity = linearVelocity
	(*b2BodyState)(unsafe.Pointer(state)).AngularVelocity = angularVelocity
}

func b2Body_GetLocalPointVelocity(tls *_Stack, bodyId BodyId, localPoint Vec2) (r1 Vec2) {
	var body, bodySim, set, state, world, v1, v3, v5, v7 uintptr
	var r, v2, v10, v11, v14, v15, v18, v19, v211, v22, v23, v9 Vec2
	var v13 Rot
	var v17 float32
	var v21, v6 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodySim, r, set, state, v2, world, v1, v10, v11, v13, v14, v15, v17, v18, v19, v21, v211, v22, v23, v3, v5, v6, v7, v9
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	state = b2GetBodyState(tls, world, body)
	if state == uintptrFromInt32(0) {
		return b2Vec2_zero
	}
	v1 = world + 1064
	v21 = (*b2Body)(unsafe.Pointer(body)).SetIndex
	if !(0 <= v21 && v21 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v21)*88
	goto _4
_4:
	set = v3
	v5 = set
	v6 = (*b2Body)(unsafe.Pointer(body)).LocalIndex
	if !(0 <= v6 && v6 < (*b2BodySimArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodySimArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*100
	goto _8
_8:
	bodySim = v7
	v9 = localPoint
	v10 = (*b2BodySim)(unsafe.Pointer(bodySim)).LocalCenter
	v11 = Vec2{
		X: v9.X - v10.X,
		Y: v9.Y - v10.Y,
	}
	goto _12
_12:
	v13 = (*b2BodySim)(unsafe.Pointer(bodySim)).Transform.Q
	v14 = v11
	v15 = Vec2{
		X: float32(v13.C*v14.X) - float32(v13.S*v14.Y),
		Y: float32(v13.S*v14.X) + float32(v13.C*v14.Y),
	}
	goto _16
_16:
	r = v15
	v17 = (*b2BodyState)(unsafe.Pointer(state)).AngularVelocity
	v18 = r
	v19 = Vec2{
		X: float32(-v17 * v18.Y),
		Y: float32(v17 * v18.X),
	}
	goto _20
_20:
	v211 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
	v22 = v19
	v23 = Vec2{
		X: v211.X + v22.X,
		Y: v211.Y + v22.Y,
	}
	goto _24
_24:
	v2 = v23
	return v2
}

func b2Body_GetWorldPointVelocity(tls *_Stack, bodyId BodyId, worldPoint Vec2) (r1 Vec2) {
	var body, bodySim, set, state, world, v11, v3, v5, v7 uintptr
	var r, v1, v10, v111, v14, v15, v17, v18, v19, v9 Vec2
	var v13 float32
	var v2, v6 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodySim, r, set, state, v1, world, v11, v10, v111, v13, v14, v15, v17, v18, v19, v2, v3, v5, v6, v7, v9
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	state = b2GetBodyState(tls, world, body)
	if state == uintptrFromInt32(0) {
		return b2Vec2_zero
	}
	v11 = world + 1064
	v2 = (*b2Body)(unsafe.Pointer(body)).SetIndex
	if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v11)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v11)).Data + uintptr(v2)*88
	goto _4
_4:
	set = v3
	v5 = set
	v6 = (*b2Body)(unsafe.Pointer(body)).LocalIndex
	if !(0 <= v6 && v6 < (*b2BodySimArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodySimArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*100
	goto _8
_8:
	bodySim = v7
	v9 = worldPoint
	v10 = (*b2BodySim)(unsafe.Pointer(bodySim)).Center
	v111 = Vec2{
		X: v9.X - v10.X,
		Y: v9.Y - v10.Y,
	}
	goto _12
_12:
	r = v111
	v13 = (*b2BodyState)(unsafe.Pointer(state)).AngularVelocity
	v14 = r
	v15 = Vec2{
		X: float32(-v13 * v14.Y),
		Y: float32(v13 * v14.X),
	}
	goto _16
_16:
	v17 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
	v18 = v15
	v19 = Vec2{
		X: v17.X + v18.X,
		Y: v17.Y + v18.Y,
	}
	goto _20
_20:
	v1 = v19
	return v1
}

func b2Body_ApplyForce(tls *_Stack, bodyId BodyId, force Vec2, point Vec2, wake uint8) {
	var body, bodySim, world uintptr
	var v1, v10, v2, v3, v5, v6, v7, v9 Vec2
	var v11 float32
	_, _, _, _, _, _, _, _, _, _, _, _ = body, bodySim, world, v1, v10, v11, v2, v3, v5, v6, v7, v9
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	if wake != 0 && (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeBody(tls, world, body)
	}
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
		bodySim = b2GetBodySim(tls, world, body)
		v1 = (*b2BodySim)(unsafe.Pointer(bodySim)).Force
		v2 = force
		v3 = Vec2{
			X: v1.X + v2.X,
			Y: v1.Y + v2.Y,
		}
		goto _4
	_4:
		(*b2BodySim)(unsafe.Pointer(bodySim)).Force = v3
		v5 = point
		v6 = (*b2BodySim)(unsafe.Pointer(bodySim)).Center
		v7 = Vec2{
			X: v5.X - v6.X,
			Y: v5.Y - v6.Y,
		}
		goto _8
	_8:
		v9 = v7
		v10 = force
		v11 = float32(v9.X*v10.Y) - float32(v9.Y*v10.X)
		goto _12
	_12:
		*(*float32)(unsafe.Pointer(bodySim + 56)) += v11
	}
}

func b2Body_ApplyForceToCenter(tls *_Stack, bodyId BodyId, force Vec2, wake uint8) {
	var body, bodySim, world uintptr
	var v1, v2, v3 Vec2
	_, _, _, _, _, _ = body, bodySim, world, v1, v2, v3
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	if wake != 0 && (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeBody(tls, world, body)
	}
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
		bodySim = b2GetBodySim(tls, world, body)
		v1 = (*b2BodySim)(unsafe.Pointer(bodySim)).Force
		v2 = force
		v3 = Vec2{
			X: v1.X + v2.X,
			Y: v1.Y + v2.Y,
		}
		goto _4
	_4:
		(*b2BodySim)(unsafe.Pointer(bodySim)).Force = v3
	}
}

func b2Body_ApplyTorque(tls *_Stack, bodyId BodyId, torque float32, wake uint8) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	if wake != 0 && (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeBody(tls, world, body)
	}
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
		bodySim = b2GetBodySim(tls, world, body)
		*(*float32)(unsafe.Pointer(bodySim + 56)) += torque
	}
}

func b2Body_ApplyLinearImpulse(tls *_Stack, bodyId BodyId, impulse Vec2, point Vec2, wake uint8) {
	var body, bodySim, set, state, world, v1, v11, v3, v5, v7, v9 uintptr
	var localIndex, v10, v2, v6 int32
	var v13, v15, v16, v18, v19, v20, v22, v23 Vec2
	var v14, v24 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodySim, localIndex, set, state, world, v1, v10, v11, v13, v14, v15, v16, v18, v19, v2, v20, v22, v23, v24, v3, v5, v6, v7, v9
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	if wake != 0 && (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeBody(tls, world, body)
	}
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
		localIndex = (*b2Body)(unsafe.Pointer(body)).LocalIndex
		v1 = world + 1064
		v2 = int32(b2_awakeSet)
		if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
		goto _4
	_4:
		set = v3
		v5 = set + 16
		v6 = localIndex
		if !(0 <= v6 && v6 < (*b2BodyStateArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(194)) != 0 {
			__builtin_trap(tls)
		}
		v7 = (*b2BodyStateArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*32
		goto _8
	_8:
		state = v7
		v9 = set
		v10 = localIndex
		if !(0 <= v10 && v10 < (*b2BodySimArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
			__builtin_trap(tls)
		}
		v11 = (*b2BodySimArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*100
		goto _12
	_12:
		bodySim = v11
		v13 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
		v14 = (*b2BodySim)(unsafe.Pointer(bodySim)).InvMass
		v15 = impulse
		v16 = Vec2{
			X: v13.X + float32(v14*v15.X),
			Y: v13.Y + float32(v14*v15.Y),
		}
		goto _17
	_17:
		(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity = v16
		v18 = point
		v19 = (*b2BodySim)(unsafe.Pointer(bodySim)).Center
		v20 = Vec2{
			X: v18.X - v19.X,
			Y: v18.Y - v19.Y,
		}
		goto _21
	_21:
		v22 = v20
		v23 = impulse
		v24 = float32(v22.X*v23.Y) - float32(v22.Y*v23.X)
		goto _25
	_25:
		*(*float32)(unsafe.Pointer(state + 8)) += float32((*b2BodySim)(unsafe.Pointer(bodySim)).InvInertia * v24)
		b2LimitVelocity(tls, state, (*b2World)(unsafe.Pointer(world)).MaxLinearSpeed)
	}
}

func b2Body_ApplyLinearImpulseToCenter(tls *_Stack, bodyId BodyId, impulse Vec2, wake uint8) {
	var body, bodySim, set, state, world, v1, v11, v3, v5, v7, v9 uintptr
	var localIndex, v10, v2, v6 int32
	var v13, v15, v16 Vec2
	var v14 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodySim, localIndex, set, state, world, v1, v10, v11, v13, v14, v15, v16, v2, v3, v5, v6, v7, v9
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	if wake != 0 && (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeBody(tls, world, body)
	}
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
		localIndex = (*b2Body)(unsafe.Pointer(body)).LocalIndex
		v1 = world + 1064
		v2 = int32(b2_awakeSet)
		if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
		goto _4
	_4:
		set = v3
		v5 = set + 16
		v6 = localIndex
		if !(0 <= v6 && v6 < (*b2BodyStateArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(194)) != 0 {
			__builtin_trap(tls)
		}
		v7 = (*b2BodyStateArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*32
		goto _8
	_8:
		state = v7
		v9 = set
		v10 = localIndex
		if !(0 <= v10 && v10 < (*b2BodySimArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
			__builtin_trap(tls)
		}
		v11 = (*b2BodySimArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*100
		goto _12
	_12:
		bodySim = v11
		v13 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
		v14 = (*b2BodySim)(unsafe.Pointer(bodySim)).InvMass
		v15 = impulse
		v16 = Vec2{
			X: v13.X + float32(v14*v15.X),
			Y: v13.Y + float32(v14*v15.Y),
		}
		goto _17
	_17:
		(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity = v16
		b2LimitVelocity(tls, state, (*b2World)(unsafe.Pointer(world)).MaxLinearSpeed)
	}
}

func b2Body_ApplyAngularImpulse(tls *_Stack, bodyId BodyId, impulse float32, wake uint8) {
	var body, bodySim, set, state, world, v1, v11, v13, v15, v3, v5, v7, v9 uintptr
	var id, localIndex, v10, v14, v2, v6 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodySim, id, localIndex, set, state, world, v1, v10, v11, v13, v14, v15, v2, v3, v5, v6, v7, v9
	if !(b2Body_IsValid(tls, bodyId) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+302, __ccgo_ts+327, int32FromInt32(999)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	id = bodyId.Index1 - int32(1)
	v1 = world + 1024
	v2 = id
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	body = v3
	if !(int32FromUint16((*b2Body)(unsafe.Pointer(body)).Generation) == int32FromUint16(bodyId.Generation)) && b2InternalAssertFcn(tls, __ccgo_ts+1559, __ccgo_ts+327, int32FromInt32(1004)) != 0 {
		__builtin_trap(tls)
	}
	if wake != 0 && (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		// this will not invalidate body pointer
		b2WakeBody(tls, world, body)
	}
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
		localIndex = (*b2Body)(unsafe.Pointer(body)).LocalIndex
		v5 = world + 1064
		v6 = int32(b2_awakeSet)
		if !(0 <= v6 && v6 < (*b2SolverSetArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
			__builtin_trap(tls)
		}
		v7 = (*b2SolverSetArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*88
		goto _8
	_8:
		set = v7
		v9 = set + 16
		v10 = localIndex
		if !(0 <= v10 && v10 < (*b2BodyStateArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(194)) != 0 {
			__builtin_trap(tls)
		}
		v11 = (*b2BodyStateArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*32
		goto _12
	_12:
		state = v11
		v13 = set
		v14 = localIndex
		if !(0 <= v14 && v14 < (*b2BodySimArray)(unsafe.Pointer(v13)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
			__builtin_trap(tls)
		}
		v15 = (*b2BodySimArray)(unsafe.Pointer(v13)).Data + uintptr(v14)*100
		goto _16
	_16:
		bodySim = v15
		*(*float32)(unsafe.Pointer(state + 8)) += float32((*b2BodySim)(unsafe.Pointer(bodySim)).InvInertia * impulse)
	}
}

func b2Body_GetType(tls *_Stack, bodyId BodyId) (r BodyType) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return (*b2Body)(unsafe.Pointer(body)).Type1
}

// C documentation
//
//	// Changing the body type is quite complex mainly due to joints.
//	// Considerations:
//	// - body and joints must be moved to the correct set
//	// - islands must be updated
//	// - graph coloring must be correct
//	// - any body connected to a joint may be disabled
//	// - joints between static bodies must go into the static set
func b2Body_SetType(tls *_Stack, bodyId BodyId, type1 BodyType) {
	var awakeSet, awakeSet1, body, bodyA, bodyB, bodySim, joint, joint1, joint2, joint3, otherBody, otherBody1, shape, shape1, shape2, staticSet, staticSet1, world, v1, v11, v13, v15, v17, v19, v21, v23, v25, v27, v29, v3, v31, v33, v35, v37, v39, v41, v43, v45, v47, v49, v5, v51, v53, v55, v57, v59, v61, v63, v7, v9 uintptr
	var edgeIndex, edgeIndex1, edgeIndex2, edgeIndex3, jointId, jointId1, jointId2, jointId3, jointKey, jointKey1, jointKey2, jointKey3, otherBodyId, otherEdgeIndex, otherEdgeIndex1, shapeId, shapeId1, shapeId2, v10, v14, v18, v2, v22, v26, v30, v34, v38, v42, v46, v50, v54, v58, v6, v62 int32
	var forcePairCreation, forcePairCreation1, forcePairCreation2, wakeBodies uint8
	var originalType, proxyType, proxyType1 b2BodyType1
	var transform, transform1, transform2 Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = awakeSet, awakeSet1, body, bodyA, bodyB, bodySim, edgeIndex, edgeIndex1, edgeIndex2, edgeIndex3, forcePairCreation, forcePairCreation1, forcePairCreation2, joint, joint1, joint2, joint3, jointId, jointId1, jointId2, jointId3, jointKey, jointKey1, jointKey2, jointKey3, originalType, otherBody, otherBody1, otherBodyId, otherEdgeIndex, otherEdgeIndex1, proxyType, proxyType1, shape, shape1, shape2, shapeId, shapeId1, shapeId2, staticSet, staticSet1, transform, transform1, transform2, wakeBodies, world, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v29, v3, v30, v31, v33, v34, v35, v37, v38, v39, v41, v42, v43, v45, v46, v47, v49, v5, v50, v51, v53, v54, v55, v57, v58, v59, v6, v61, v62, v63, v7, v9
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	originalType = (*b2Body)(unsafe.Pointer(body)).Type1
	if originalType == type1 {
		return
	}
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_disabledSet) {
		// Disabled bodies don't change solver sets or islands when they change type.
		(*b2Body)(unsafe.Pointer(body)).Type1 = type1
		// Body type affects the mass
		b2UpdateBodyMassData(tls, world, body)
		return
	}
	// Destroy all contacts but don't wake bodies.
	wakeBodies = boolUint8(false1 != 0)
	b2DestroyBodyContacts(tls, world, body, wakeBodies)
	// Wake this body because we assume below that it is awake or static.
	b2WakeBody(tls, world, body)
	// Unlink all joints and wake attached bodies.
	jointKey = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
	for jointKey != -int32(1) {
		jointId = jointKey >> int32(1)
		edgeIndex = jointKey & int32(1)
		v1 = world + 1104
		v2 = jointId
		if !(0 <= v2 && v2 < (*b2JointArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2JointArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*72
		goto _4
	_4:
		joint = v3
		if (*b2Joint)(unsafe.Pointer(joint)).IslandId != -int32(1) {
			b2UnlinkJoint(tls, world, joint)
		}
		v5 = world + 1024
		v6 = (*(*b2JointEdge)(unsafe.Pointer(joint + 20))).BodyId
		if !(0 <= v6 && v6 < (*b2BodyArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
			__builtin_trap(tls)
		}
		v7 = (*b2BodyArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*128
		goto _8
	_8:
		// A body going from static to dynamic or kinematic goes to the awake set
		// and other attached bodies must be awake as well. For consistency, this is
		// done for all cases.
		bodyA = v7
		v9 = world + 1024
		v10 = (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + 1*12))).BodyId
		if !(0 <= v10 && v10 < (*b2BodyArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
			__builtin_trap(tls)
		}
		v11 = (*b2BodyArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*128
		goto _12
	_12:
		bodyB = v11
		b2WakeBody(tls, world, bodyA)
		b2WakeBody(tls, world, bodyB)
		jointKey = (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + uintptr(edgeIndex)*12))).NextKey
	}
	(*b2Body)(unsafe.Pointer(body)).Type1 = type1
	if originalType == int32(b2_staticBody) {
		// Body is going from static to dynamic or kinematic. It only makes sense to move it to the awake set.
		if !((*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_staticSet)) && b2InternalAssertFcn(tls, __ccgo_ts+1597, __ccgo_ts+327, int32FromInt32(1095)) != 0 {
			__builtin_trap(tls)
		}
		v13 = world + 1064
		v14 = int32(b2_staticSet)
		if !(0 <= v14 && v14 < (*b2SolverSetArray)(unsafe.Pointer(v13)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
			__builtin_trap(tls)
		}
		v15 = (*b2SolverSetArray)(unsafe.Pointer(v13)).Data + uintptr(v14)*88
		goto _16
	_16:
		staticSet = v15
		v17 = world + 1064
		v18 = int32(b2_awakeSet)
		if !(0 <= v18 && v18 < (*b2SolverSetArray)(unsafe.Pointer(v17)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
			__builtin_trap(tls)
		}
		v19 = (*b2SolverSetArray)(unsafe.Pointer(v17)).Data + uintptr(v18)*88
		goto _20
	_20:
		awakeSet = v19
		// Transfer body to awake set
		b2TransferBody(tls, world, awakeSet, staticSet, body)
		// Create island for body
		b2CreateIslandForBody(tls, world, int32(b2_awakeSet), body)
		// Transfer static joints to awake set
		jointKey1 = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
		for jointKey1 != -int32(1) {
			jointId1 = jointKey1 >> int32(1)
			edgeIndex1 = jointKey1 & int32(1)
			v21 = world + 1104
			v22 = jointId1
			if !(0 <= v22 && v22 < (*b2JointArray)(unsafe.Pointer(v21)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
				__builtin_trap(tls)
			}
			v23 = (*b2JointArray)(unsafe.Pointer(v21)).Data + uintptr(v22)*72
			goto _24
		_24:
			joint1 = v23
			// Transfer the joint if it is in the static set
			if (*b2Joint)(unsafe.Pointer(joint1)).SetIndex == int32(b2_staticSet) {
				b2TransferJoint(tls, world, awakeSet, staticSet, joint1)
			} else {
				if (*b2Joint)(unsafe.Pointer(joint1)).SetIndex == int32(b2_awakeSet) {
					// In this case the joint must be re-inserted into the constraint graph to ensure the correct
					// graph color.
					// First transfer to the static set.
					b2TransferJoint(tls, world, staticSet, awakeSet, joint1)
					// Now transfer it back to the awake set and into the graph coloring.
					b2TransferJoint(tls, world, awakeSet, staticSet, joint1)
				} else {
					// Otherwise the joint must be disabled.
					if !((*b2Joint)(unsafe.Pointer(joint1)).SetIndex == int32(b2_disabledSet)) && b2InternalAssertFcn(tls, __ccgo_ts+1628, __ccgo_ts+327, int32FromInt32(1134)) != 0 {
						__builtin_trap(tls)
					}
				}
			}
			jointKey1 = (*(*b2JointEdge)(unsafe.Pointer(joint1 + 20 + uintptr(edgeIndex1)*12))).NextKey
		}
		// Recreate shape proxies in movable tree.
		transform = b2GetBodyTransformQuick(tls, world, body)
		shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
		for shapeId != -int32(1) {
			v25 = world + 1248
			v26 = shapeId
			if !(0 <= v26 && v26 < (*b2ShapeArray)(unsafe.Pointer(v25)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
				__builtin_trap(tls)
			}
			v27 = (*b2ShapeArray)(unsafe.Pointer(v25)).Data + uintptr(v26)*288
			goto _28
		_28:
			shape = v27
			shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
			b2DestroyShapeProxy(tls, shape, world+40)
			forcePairCreation = boolUint8(true1 != 0)
			proxyType = type1
			b2CreateShapeProxy(tls, shape, world+40, proxyType, transform, forcePairCreation)
		}
	} else {
		if type1 == int32(b2_staticBody) {
			// The body is going from dynamic/kinematic to static. It should be awake.
			if !((*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+1662, __ccgo_ts+327, int32FromInt32(1156)) != 0 {
				__builtin_trap(tls)
			}
			v29 = world + 1064
			v30 = int32(b2_staticSet)
			if !(0 <= v30 && v30 < (*b2SolverSetArray)(unsafe.Pointer(v29)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
				__builtin_trap(tls)
			}
			v31 = (*b2SolverSetArray)(unsafe.Pointer(v29)).Data + uintptr(v30)*88
			goto _32
		_32:
			staticSet1 = v31
			v33 = world + 1064
			v34 = int32(b2_awakeSet)
			if !(0 <= v34 && v34 < (*b2SolverSetArray)(unsafe.Pointer(v33)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
				__builtin_trap(tls)
			}
			v35 = (*b2SolverSetArray)(unsafe.Pointer(v33)).Data + uintptr(v34)*88
			goto _36
		_36:
			awakeSet1 = v35
			// Transfer body to static set
			b2TransferBody(tls, world, staticSet1, awakeSet1, body)
			// Remove body from island.
			b2RemoveBodyFromIsland(tls, world, body)
			v37 = staticSet1
			v38 = (*b2Body)(unsafe.Pointer(body)).LocalIndex
			if !(0 <= v38 && v38 < (*b2BodySimArray)(unsafe.Pointer(v37)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
				__builtin_trap(tls)
			}
			v39 = (*b2BodySimArray)(unsafe.Pointer(v37)).Data + uintptr(v38)*100
			goto _40
		_40:
			bodySim = v39
			(*b2BodySim)(unsafe.Pointer(bodySim)).IsFast = boolUint8(false1 != 0)
			// Maybe transfer joints to static set.
			jointKey2 = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
			for jointKey2 != -int32(1) {
				jointId2 = jointKey2 >> int32(1)
				edgeIndex2 = jointKey2 & int32(1)
				v41 = world + 1104
				v42 = jointId2
				if !(0 <= v42 && v42 < (*b2JointArray)(unsafe.Pointer(v41)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
					__builtin_trap(tls)
				}
				v43 = (*b2JointArray)(unsafe.Pointer(v41)).Data + uintptr(v42)*72
				goto _44
			_44:
				joint2 = v43
				jointKey2 = (*(*b2JointEdge)(unsafe.Pointer(joint2 + 20 + uintptr(edgeIndex2)*12))).NextKey
				otherEdgeIndex = edgeIndex2 ^ int32(1)
				v45 = world + 1024
				v46 = (*(*b2JointEdge)(unsafe.Pointer(joint2 + 20 + uintptr(otherEdgeIndex)*12))).BodyId
				if !(0 <= v46 && v46 < (*b2BodyArray)(unsafe.Pointer(v45)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
					__builtin_trap(tls)
				}
				v47 = (*b2BodyArray)(unsafe.Pointer(v45)).Data + uintptr(v46)*128
				goto _48
			_48:
				otherBody = v47
				// Skip disabled joint
				if (*b2Joint)(unsafe.Pointer(joint2)).SetIndex == int32(b2_disabledSet) {
					// Joint is disable, should be connected to a disabled body
					if !((*b2Body)(unsafe.Pointer(otherBody)).SetIndex == int32(b2_disabledSet)) && b2InternalAssertFcn(tls, __ccgo_ts+1692, __ccgo_ts+327, int32FromInt32(1187)) != 0 {
						__builtin_trap(tls)
					}
					continue
				}
				// Since the body was not static, the joint must be awake.
				if !((*b2Joint)(unsafe.Pointer(joint2)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+1730, __ccgo_ts+327, int32FromInt32(1192)) != 0 {
					__builtin_trap(tls)
				}
				// Only transfer joint to static set if both bodies are static.
				if (*b2Body)(unsafe.Pointer(otherBody)).SetIndex == int32(b2_staticSet) {
					b2TransferJoint(tls, world, staticSet1, awakeSet1, joint2)
				} else {
					// The other body must be awake.
					if !((*b2Body)(unsafe.Pointer(otherBody)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+1761, __ccgo_ts+327, int32FromInt32(1202)) != 0 {
						__builtin_trap(tls)
					}
					// The joint must live in a graph color.
					if !(0 <= (*b2Joint)(unsafe.Pointer(joint2)).ColorIndex && (*b2Joint)(unsafe.Pointer(joint2)).ColorIndex < int32(_B2_GRAPH_COLOR_COUNT)) && b2InternalAssertFcn(tls, __ccgo_ts+1796, __ccgo_ts+327, int32FromInt32(1205)) != 0 {
						__builtin_trap(tls)
					}
					// In this case the joint must be re-inserted into the constraint graph to ensure the correct
					// graph color.
					// First transfer to the static set.
					b2TransferJoint(tls, world, staticSet1, awakeSet1, joint2)
					// Now transfer it back to the awake set and into the graph coloring.
					b2TransferJoint(tls, world, awakeSet1, staticSet1, joint2)
				}
			}
			// Recreate shape proxies in static tree.
			transform1 = b2GetBodyTransformQuick(tls, world, body)
			shapeId1 = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
			for shapeId1 != -int32(1) {
				v49 = world + 1248
				v50 = shapeId1
				if !(0 <= v50 && v50 < (*b2ShapeArray)(unsafe.Pointer(v49)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
					__builtin_trap(tls)
				}
				v51 = (*b2ShapeArray)(unsafe.Pointer(v49)).Data + uintptr(v50)*288
				goto _52
			_52:
				shape1 = v51
				shapeId1 = (*b2Shape)(unsafe.Pointer(shape1)).NextShapeId
				b2DestroyShapeProxy(tls, shape1, world+40)
				forcePairCreation1 = boolUint8(true1 != 0)
				b2CreateShapeProxy(tls, shape1, world+40, int32(b2_staticBody), transform1, forcePairCreation1)
			}
		} else {
			if !(originalType == int32(b2_dynamicBody) || originalType == int32(b2_kinematicBody)) && b2InternalAssertFcn(tls, __ccgo_ts+1863, __ccgo_ts+327, int32FromInt32(1232)) != 0 {
				__builtin_trap(tls)
			}
			if !(type1 == int32(b2_dynamicBody) || type1 == int32(b2_kinematicBody)) && b2InternalAssertFcn(tls, __ccgo_ts+1930, __ccgo_ts+327, int32FromInt32(1233)) != 0 {
				__builtin_trap(tls)
			}
			// Recreate shape proxies in static tree.
			transform2 = b2GetBodyTransformQuick(tls, world, body)
			shapeId2 = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
			for shapeId2 != -int32(1) {
				v53 = world + 1248
				v54 = shapeId2
				if !(0 <= v54 && v54 < (*b2ShapeArray)(unsafe.Pointer(v53)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
					__builtin_trap(tls)
				}
				v55 = (*b2ShapeArray)(unsafe.Pointer(v53)).Data + uintptr(v54)*288
				goto _56
			_56:
				shape2 = v55
				shapeId2 = (*b2Shape)(unsafe.Pointer(shape2)).NextShapeId
				b2DestroyShapeProxy(tls, shape2, world+40)
				proxyType1 = type1
				forcePairCreation2 = boolUint8(true1 != 0)
				b2CreateShapeProxy(tls, shape2, world+40, proxyType1, transform2, forcePairCreation2)
			}
		}
	}
	// Relink all joints
	jointKey3 = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
	for jointKey3 != -int32(1) {
		jointId3 = jointKey3 >> int32(1)
		edgeIndex3 = jointKey3 & int32(1)
		v57 = world + 1104
		v58 = jointId3
		if !(0 <= v58 && v58 < (*b2JointArray)(unsafe.Pointer(v57)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v59 = (*b2JointArray)(unsafe.Pointer(v57)).Data + uintptr(v58)*72
		goto _60
	_60:
		joint3 = v59
		jointKey3 = (*(*b2JointEdge)(unsafe.Pointer(joint3 + 20 + uintptr(edgeIndex3)*12))).NextKey
		otherEdgeIndex1 = edgeIndex3 ^ int32(1)
		otherBodyId = (*(*b2JointEdge)(unsafe.Pointer(joint3 + 20 + uintptr(otherEdgeIndex1)*12))).BodyId
		v61 = world + 1024
		v62 = otherBodyId
		if !(0 <= v62 && v62 < (*b2BodyArray)(unsafe.Pointer(v61)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
			__builtin_trap(tls)
		}
		v63 = (*b2BodyArray)(unsafe.Pointer(v61)).Data + uintptr(v62)*128
		goto _64
	_64:
		otherBody1 = v63
		if (*b2Body)(unsafe.Pointer(otherBody1)).SetIndex == int32(b2_disabledSet) {
			continue
		}
		if (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_staticBody) && (*b2Body)(unsafe.Pointer(otherBody1)).Type1 == int32(b2_staticBody) {
			continue
		}
		b2LinkJoint(tls, world, joint3, boolUint8(false1 != 0))
	}
	b2MergeAwakeIslands(tls, world)
	// Body type affects the mass
	b2UpdateBodyMassData(tls, world, body)
	b2ValidateSolverSets(tls, world)
}

func b2Body_SetName(tls *_Stack, bodyId BodyId, name uintptr) {
	var body, world uintptr
	var i int32
	_, _, _ = body, i, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	if name != uintptrFromInt32(0) {
		i = 0
		for {
			if !(i < int32(31)) {
				break
			}
			*(*uint8)(unsafe.Pointer(body + uintptr(i))) = *(*uint8)(unsafe.Pointer(name + uintptr(i)))
			goto _1
		_1:
			;
			i = i + 1
		}
		*(*uint8)(unsafe.Pointer(body + 31)) = uint8(0)
	} else {
		memset(tls, body, 0, uint64FromInt32(32)*uint64FromInt64(1))
	}
}

func b2Body_GetName(tls *_Stack, bodyId BodyId) (r uintptr) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return body
}

func b2Body_SetUserData(tls *_Stack, bodyId BodyId, userData uintptr) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	(*b2Body)(unsafe.Pointer(body)).UserData = userData
}

func b2Body_GetUserData(tls *_Stack, bodyId BodyId) (r uintptr) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return (*b2Body)(unsafe.Pointer(body)).UserData
}

func b2Body_GetMass(tls *_Stack, bodyId BodyId) (r float32) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return (*b2Body)(unsafe.Pointer(body)).Mass
}

func b2Body_GetRotationalInertia(tls *_Stack, bodyId BodyId) (r float32) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return (*b2Body)(unsafe.Pointer(body)).Inertia
}

func b2Body_GetLocalCenterOfMass(tls *_Stack, bodyId BodyId) (r Vec2) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	return (*b2BodySim)(unsafe.Pointer(bodySim)).LocalCenter
}

func b2Body_GetWorldCenterOfMass(tls *_Stack, bodyId BodyId) (r Vec2) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	return (*b2BodySim)(unsafe.Pointer(bodySim)).Center
}

func b2Body_SetMassData(tls *_Stack, bodyId BodyId, massData MassData) {
	var body, bodySim, world uintptr
	var center, v2, v3 Vec2
	var x, y, v5, v6 float32
	var v1 Transform
	_, _, _, _, _, _, _, _, _, _, _ = body, bodySim, center, world, x, y, v1, v2, v3, v5, v6
	if !(b2IsValidFloat(tls, massData.Mass) != 0 && massData.Mass >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+1981, __ccgo_ts+327, int32FromInt32(1359)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsValidFloat(tls, massData.RotationalInertia) != 0 && massData.RotationalInertia >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+2038, __ccgo_ts+327, int32FromInt32(1360)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsValidVec2(tls, massData.Center) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+2121, __ccgo_ts+327, int32FromInt32(1361)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	(*b2Body)(unsafe.Pointer(body)).Mass = massData.Mass
	(*b2Body)(unsafe.Pointer(body)).Inertia = massData.RotationalInertia
	(*b2BodySim)(unsafe.Pointer(bodySim)).LocalCenter = massData.Center
	v1 = (*b2BodySim)(unsafe.Pointer(bodySim)).Transform
	v2 = massData.Center
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	center = v3
	(*b2BodySim)(unsafe.Pointer(bodySim)).Center = center
	(*b2BodySim)(unsafe.Pointer(bodySim)).Center0 = center
	if (*b2Body)(unsafe.Pointer(body)).Mass > float32FromFloat32(0) {
		v5 = float32FromFloat32(1) / (*b2Body)(unsafe.Pointer(body)).Mass
	} else {
		v5 = float32FromFloat32(0)
	}
	(*b2BodySim)(unsafe.Pointer(bodySim)).InvMass = v5
	if (*b2Body)(unsafe.Pointer(body)).Inertia > float32FromFloat32(0) {
		v6 = float32FromFloat32(1) / (*b2Body)(unsafe.Pointer(body)).Inertia
	} else {
		v6 = float32FromFloat32(0)
	}
	(*b2BodySim)(unsafe.Pointer(bodySim)).InvInertia = v6
}

func b2Body_GetMassData(tls *_Stack, bodyId BodyId) (r MassData) {
	var body, bodySim, world uintptr
	var massData MassData
	_, _, _, _ = body, bodySim, massData, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	massData = MassData{
		Mass:              (*b2Body)(unsafe.Pointer(body)).Mass,
		Center:            (*b2BodySim)(unsafe.Pointer(bodySim)).LocalCenter,
		RotationalInertia: (*b2Body)(unsafe.Pointer(body)).Inertia,
	}
	return massData
}

func b2Body_ApplyMassFromShapes(tls *_Stack, bodyId BodyId) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	b2UpdateBodyMassData(tls, world, body)
}

func b2Body_SetLinearDamping(tls *_Stack, bodyId BodyId, linearDamping float32) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	if !(b2IsValidFloat(tls, linearDamping) != 0 && linearDamping >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+2154, __ccgo_ts+327, int32FromInt32(1407)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	(*b2BodySim)(unsafe.Pointer(bodySim)).LinearDamping = linearDamping
}

func b2Body_GetLinearDamping(tls *_Stack, bodyId BodyId) (r float32) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	return (*b2BodySim)(unsafe.Pointer(bodySim)).LinearDamping
}

func b2Body_SetAngularDamping(tls *_Stack, bodyId BodyId, angularDamping float32) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	if !(b2IsValidFloat(tls, angularDamping) != 0 && angularDamping >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+2211, __ccgo_ts+327, int32FromInt32(1430)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	(*b2BodySim)(unsafe.Pointer(bodySim)).AngularDamping = angularDamping
}

func b2Body_GetAngularDamping(tls *_Stack, bodyId BodyId) (r float32) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	return (*b2BodySim)(unsafe.Pointer(bodySim)).AngularDamping
}

func b2Body_SetGravityScale(tls *_Stack, bodyId BodyId, gravityScale float32) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	if !(b2Body_IsValid(tls, bodyId) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+302, __ccgo_ts+327, int32FromInt32(1453)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsValidFloat(tls, gravityScale) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+2270, __ccgo_ts+327, int32FromInt32(1454)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	(*b2BodySim)(unsafe.Pointer(bodySim)).GravityScale = gravityScale
}

func b2Body_GetGravityScale(tls *_Stack, bodyId BodyId) (r float32) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	if !(b2Body_IsValid(tls, bodyId) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+302, __ccgo_ts+327, int32FromInt32(1469)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	return (*b2BodySim)(unsafe.Pointer(bodySim)).GravityScale
}

func b2Body_IsAwake(tls *_Stack, bodyId BodyId) (r uint8) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return boolUint8((*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet))
}

func b2Body_SetAwake(tls *_Stack, bodyId BodyId, awake uint8) {
	var body, island, world, v1, v3 uintptr
	var v2 int32
	_, _, _, _, _, _ = body, island, world, v1, v2, v3
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	if awake != 0 && (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeBody(tls, world, body)
	} else {
		if int32FromUint8(awake) == false1 && (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
			v1 = world + 1184
			v2 = (*b2Body)(unsafe.Pointer(body)).IslandId
			if !(0 <= v2 && v2 < (*b2IslandArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
				__builtin_trap(tls)
			}
			v3 = (*b2IslandArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*56
			goto _4
		_4:
			island = v3
			if (*b2Island)(unsafe.Pointer(island)).ConstraintRemoveCount > 0 {
				// Must split the island before sleeping. This is expensive.
				b2SplitIsland(tls, world, (*b2Body)(unsafe.Pointer(body)).IslandId)
			}
			b2TrySleepIsland(tls, world, (*b2Body)(unsafe.Pointer(body)).IslandId)
		}
	}
}

func b2Body_IsEnabled(tls *_Stack, bodyId BodyId) (r uint8) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return boolUint8((*b2Body)(unsafe.Pointer(body)).SetIndex != int32(b2_disabledSet))
}

func b2Body_IsSleepEnabled(tls *_Stack, bodyId BodyId) (r uint8) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return (*b2Body)(unsafe.Pointer(body)).EnableSleep
}

func b2Body_SetSleepThreshold(tls *_Stack, bodyId BodyId, sleepThreshold float32) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	(*b2Body)(unsafe.Pointer(body)).SleepThreshold = sleepThreshold
}

func b2Body_GetSleepThreshold(tls *_Stack, bodyId BodyId) (r float32) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return (*b2Body)(unsafe.Pointer(body)).SleepThreshold
}

func b2Body_EnableSleep(tls *_Stack, bodyId BodyId, enableSleep uint8) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	(*b2Body)(unsafe.Pointer(body)).EnableSleep = enableSleep
	if int32FromUint8(enableSleep) == false1 {
		b2WakeBody(tls, world, body)
	}
}

// C documentation
//
//	// Disabling a body requires a lot of detailed bookkeeping, but it is a valuable feature.
//	// The most challenging aspect that joints may connect to bodies that are not disabled.
func b2Body_Disable(tls *_Stack, bodyId BodyId) {
	var body, disabledSet, joint, jointSet, set, shape, world, v1, v11, v13, v15, v17, v19, v3, v5, v7, v9 uintptr
	var edgeIndex, jointId, jointKey, shapeId, v10, v14, v18, v2, v6 int32
	var wakeBodies uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, disabledSet, edgeIndex, joint, jointId, jointKey, jointSet, set, shape, shapeId, wakeBodies, world, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v3, v5, v6, v7, v9
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_disabledSet) {
		return
	}
	// Destroy contacts and wake bodies touching this body. This avoid floating bodies.
	// This is necessary even for static bodies.
	wakeBodies = boolUint8(true1 != 0)
	b2DestroyBodyContacts(tls, world, body, wakeBodies)
	// Disabled bodies are not in an island.
	b2RemoveBodyFromIsland(tls, world, body)
	// Remove shapes from broad-phase
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		v1 = world + 1248
		v2 = shapeId
		if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
		goto _4
	_4:
		shape = v3
		shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
		b2DestroyShapeProxy(tls, shape, world+40)
	}
	v5 = world + 1064
	v6 = (*b2Body)(unsafe.Pointer(body)).SetIndex
	if !(0 <= v6 && v6 < (*b2SolverSetArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2SolverSetArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*88
	goto _8
_8:
	// Transfer simulation data to disabled set
	set = v7
	v9 = world + 1064
	v10 = int32(b2_disabledSet)
	if !(0 <= v10 && v10 < (*b2SolverSetArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v11 = (*b2SolverSetArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*88
	goto _12
_12:
	disabledSet = v11
	// Transfer body sim
	b2TransferBody(tls, world, disabledSet, set, body)
	// Unlink joints and transfer
	jointKey = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
	for jointKey != -int32(1) {
		jointId = jointKey >> int32(1)
		edgeIndex = jointKey & int32(1)
		v13 = world + 1104
		v14 = jointId
		if !(0 <= v14 && v14 < (*b2JointArray)(unsafe.Pointer(v13)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v15 = (*b2JointArray)(unsafe.Pointer(v13)).Data + uintptr(v14)*72
		goto _16
	_16:
		joint = v15
		jointKey = (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + uintptr(edgeIndex)*12))).NextKey
		// joint may already be disabled by other body
		if (*b2Joint)(unsafe.Pointer(joint)).SetIndex == int32(b2_disabledSet) {
			continue
		}
		if !((*b2Joint)(unsafe.Pointer(joint)).SetIndex == (*b2SolverSet)(unsafe.Pointer(set)).SetIndex || (*b2SolverSet)(unsafe.Pointer(set)).SetIndex == int32(b2_staticSet)) && b2InternalAssertFcn(tls, __ccgo_ts+2301, __ccgo_ts+327, int32FromInt32(1611)) != 0 {
			__builtin_trap(tls)
		}
		// Remove joint from island
		if (*b2Joint)(unsafe.Pointer(joint)).IslandId != -int32(1) {
			b2UnlinkJoint(tls, world, joint)
		}
		v17 = world + 1064
		v18 = (*b2Joint)(unsafe.Pointer(joint)).SetIndex
		if !(0 <= v18 && v18 < (*b2SolverSetArray)(unsafe.Pointer(v17)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
			__builtin_trap(tls)
		}
		v19 = (*b2SolverSetArray)(unsafe.Pointer(v17)).Data + uintptr(v18)*88
		goto _20
	_20:
		// Transfer joint to disabled set
		jointSet = v19
		b2TransferJoint(tls, world, disabledSet, jointSet, joint)
	}
	b2ValidateConnectivity(tls, world)
	b2ValidateSolverSets(tls, world)
}

func b2Body_Enable(tls *_Stack, bodyId BodyId) {
	var body, bodyA, bodyB, disabledSet, joint, jointSet, shape, targetSet, world, v1, v10, v12, v14, v16, v18, v20, v22, v24, v26, v28, v3, v6, v8 uintptr
	var edgeIndex, jointId, jointKey, jointSetId, setId, shapeId, v11, v15, v19, v2, v23, v27, v5, v7 int32
	var forcePairCreation, mergeIslands uint8
	var proxyType b2BodyType1
	var transform Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodyA, bodyB, disabledSet, edgeIndex, forcePairCreation, joint, jointId, jointKey, jointSet, jointSetId, mergeIslands, proxyType, setId, shape, shapeId, targetSet, transform, world, v1, v10, v11, v12, v14, v15, v16, v18, v19, v2, v20, v22, v23, v24, v26, v27, v28, v3, v5, v6, v7, v8
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).SetIndex != int32(b2_disabledSet) {
		return
	}
	v1 = world + 1064
	v2 = int32(b2_disabledSet)
	if !(0 <= v2 && v2 < (*b2SolverSetArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SolverSetArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*88
	goto _4
_4:
	disabledSet = v3
	if (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_staticBody) {
		v5 = int32(b2_staticSet)
	} else {
		v5 = int32(b2_awakeSet)
	}
	setId = v5
	v6 = world + 1064
	v7 = setId
	if !(0 <= v7 && v7 < (*b2SolverSetArray)(unsafe.Pointer(v6)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v8 = (*b2SolverSetArray)(unsafe.Pointer(v6)).Data + uintptr(v7)*88
	goto _9
_9:
	targetSet = v8
	b2TransferBody(tls, world, targetSet, disabledSet, body)
	transform = b2GetBodyTransformQuick(tls, world, body)
	// Add shapes to broad-phase
	proxyType = (*b2Body)(unsafe.Pointer(body)).Type1
	forcePairCreation = boolUint8(true1 != 0)
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		v10 = world + 1248
		v11 = shapeId
		if !(0 <= v11 && v11 < (*b2ShapeArray)(unsafe.Pointer(v10)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v12 = (*b2ShapeArray)(unsafe.Pointer(v10)).Data + uintptr(v11)*288
		goto _13
	_13:
		shape = v12
		shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
		b2CreateShapeProxy(tls, shape, world+40, proxyType, transform, forcePairCreation)
	}
	if setId != int32(b2_staticSet) {
		b2CreateIslandForBody(tls, world, setId, body)
	}
	// Transfer joints. If the other body is disabled, don't transfer.
	// If the other body is sleeping, wake it.
	mergeIslands = boolUint8(false1 != 0)
	jointKey = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
	for jointKey != -int32(1) {
		jointId = jointKey >> int32(1)
		edgeIndex = jointKey & int32(1)
		v14 = world + 1104
		v15 = jointId
		if !(0 <= v15 && v15 < (*b2JointArray)(unsafe.Pointer(v14)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v16 = (*b2JointArray)(unsafe.Pointer(v14)).Data + uintptr(v15)*72
		goto _17
	_17:
		joint = v16
		if !((*b2Joint)(unsafe.Pointer(joint)).SetIndex == int32(b2_disabledSet)) && b2InternalAssertFcn(tls, __ccgo_ts+1628, __ccgo_ts+327, int32FromInt32(1677)) != 0 {
			__builtin_trap(tls)
		}
		if !((*b2Joint)(unsafe.Pointer(joint)).IslandId == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+2367, __ccgo_ts+327, int32FromInt32(1678)) != 0 {
			__builtin_trap(tls)
		}
		jointKey = (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + uintptr(edgeIndex)*12))).NextKey
		v18 = world + 1024
		v19 = (*(*b2JointEdge)(unsafe.Pointer(joint + 20))).BodyId
		if !(0 <= v19 && v19 < (*b2BodyArray)(unsafe.Pointer(v18)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
			__builtin_trap(tls)
		}
		v20 = (*b2BodyArray)(unsafe.Pointer(v18)).Data + uintptr(v19)*128
		goto _21
	_21:
		bodyA = v20
		v22 = world + 1024
		v23 = (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + 1*12))).BodyId
		if !(0 <= v23 && v23 < (*b2BodyArray)(unsafe.Pointer(v22)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
			__builtin_trap(tls)
		}
		v24 = (*b2BodyArray)(unsafe.Pointer(v22)).Data + uintptr(v23)*128
		goto _25
	_25:
		bodyB = v24
		if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_disabledSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_disabledSet) {
			// one body is still disabled
			continue
		}
		if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_staticSet) && (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_staticSet) {
			jointSetId = int32(b2_staticSet)
		} else {
			if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_staticSet) {
				jointSetId = (*b2Body)(unsafe.Pointer(bodyB)).SetIndex
			} else {
				jointSetId = (*b2Body)(unsafe.Pointer(bodyA)).SetIndex
			}
		}
		v26 = world + 1064
		v27 = jointSetId
		if !(0 <= v27 && v27 < (*b2SolverSetArray)(unsafe.Pointer(v26)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
			__builtin_trap(tls)
		}
		v28 = (*b2SolverSetArray)(unsafe.Pointer(v26)).Data + uintptr(v27)*88
		goto _29
	_29:
		jointSet = v28
		b2TransferJoint(tls, world, jointSet, disabledSet, joint)
		// Now that the joint is in the correct set, I can link the joint in the island.
		if jointSetId != int32(b2_staticSet) {
			b2LinkJoint(tls, world, joint, mergeIslands)
		}
	}
	// Now merge islands
	b2MergeAwakeIslands(tls, world)
	b2ValidateSolverSets(tls, world)
}

func b2Body_SetFixedRotation(tls *_Stack, bodyId BodyId, flag uint8) {
	var body, state, world uintptr
	_, _, _ = body, state, world
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).FixedRotation != flag {
		(*b2Body)(unsafe.Pointer(body)).FixedRotation = flag
		state = b2GetBodyState(tls, world, body)
		if state != uintptrFromInt32(0) {
			(*b2BodyState)(unsafe.Pointer(state)).AngularVelocity = float32FromFloat32(0)
		}
		b2UpdateBodyMassData(tls, world, body)
	}
}

func b2Body_IsFixedRotation(tls *_Stack, bodyId BodyId) (r uint8) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return (*b2Body)(unsafe.Pointer(body)).FixedRotation
}

func b2Body_SetBullet(tls *_Stack, bodyId BodyId, flag uint8) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	(*b2BodySim)(unsafe.Pointer(bodySim)).IsBullet = flag
}

func b2Body_IsBullet(tls *_Stack, bodyId BodyId) (r uint8) {
	var body, bodySim, world uintptr
	_, _, _ = body, bodySim, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	bodySim = b2GetBodySim(tls, world, body)
	return (*b2BodySim)(unsafe.Pointer(bodySim)).IsBullet
}

func b2Body_EnableContactEvents(tls *_Stack, bodyId BodyId, flag uint8) {
	var body, shape, world, v1, v3 uintptr
	var shapeId, v2 int32
	_, _, _, _, _, _, _ = body, shape, shapeId, world, v1, v2, v3
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		v1 = world + 1248
		v2 = shapeId
		if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
		goto _4
	_4:
		shape = v3
		(*b2Shape)(unsafe.Pointer(shape)).EnableContactEvents = flag
		shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
	}
}

func b2Body_EnableHitEvents(tls *_Stack, bodyId BodyId, flag uint8) {
	var body, shape, world, v1, v3 uintptr
	var shapeId, v2 int32
	_, _, _, _, _, _, _ = body, shape, shapeId, world, v1, v2, v3
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		v1 = world + 1248
		v2 = shapeId
		if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
		goto _4
	_4:
		shape = v3
		(*b2Shape)(unsafe.Pointer(shape)).EnableHitEvents = flag
		shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
	}
}

func b2Body_GetWorld(tls *_Stack, bodyId BodyId) (r WorldId) {
	var world uintptr
	_ = world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	return WorldId{
		Index1:     uint16FromInt32(int32FromUint16(bodyId.World0) + int32(1)),
		Generation: (*b2World)(unsafe.Pointer(world)).Generation,
	}
}

func b2Body_GetShapeCount(tls *_Stack, bodyId BodyId) (r int32) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return (*b2Body)(unsafe.Pointer(body)).ShapeCount
}

func b2Body_GetShapes(tls *_Stack, bodyId BodyId, shapeArray uintptr, capacity int32) (r int32) {
	var body, shape, world, v1, v3 uintptr
	var id ShapeId
	var shapeCount, shapeId, v2 int32
	_, _, _, _, _, _, _, _, _ = body, id, shape, shapeCount, shapeId, world, v1, v2, v3
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	shapeCount = 0
	for shapeId != -int32(1) && shapeCount < capacity {
		v1 = world + 1248
		v2 = shapeId
		if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
		goto _4
	_4:
		shape = v3
		id = ShapeId{
			Index1:     (*b2Shape)(unsafe.Pointer(shape)).Id + int32(1),
			World0:     bodyId.World0,
			Generation: (*b2Shape)(unsafe.Pointer(shape)).Generation,
		}
		*(*ShapeId)(unsafe.Pointer(shapeArray + uintptr(shapeCount)*8)) = id
		shapeCount = shapeCount + int32(1)
		shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
	}
	return shapeCount
}

func b2Body_GetJointCount(tls *_Stack, bodyId BodyId) (r int32) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return (*b2Body)(unsafe.Pointer(body)).JointCount
}

func b2Body_GetJoints(tls *_Stack, bodyId BodyId, jointArray uintptr, capacity int32) (r int32) {
	var body, joint, world, v1, v3 uintptr
	var edgeIndex, jointCount, jointId, jointKey, v2 int32
	var id JointId
	_, _, _, _, _, _, _, _, _, _, _ = body, edgeIndex, id, joint, jointCount, jointId, jointKey, world, v1, v2, v3
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	jointKey = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
	jointCount = 0
	for jointKey != -int32(1) && jointCount < capacity {
		jointId = jointKey >> int32(1)
		edgeIndex = jointKey & int32(1)
		v1 = world + 1104
		v2 = jointId
		if !(0 <= v2 && v2 < (*b2JointArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2JointArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*72
		goto _4
	_4:
		joint = v3
		id = JointId{
			Index1:     jointId + int32(1),
			World0:     bodyId.World0,
			Generation: (*b2Joint)(unsafe.Pointer(joint)).Generation,
		}
		*(*JointId)(unsafe.Pointer(jointArray + uintptr(jointCount)*8)) = id
		jointCount = jointCount + int32(1)
		jointKey = (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + uintptr(edgeIndex)*12))).NextKey
	}
	return jointCount
}

func b2BulletBodyTask(tls *_Stack, startIndex int32, endIndex int32, threadIndex uint32, taskContext uintptr) {
	var i, simIndex int32
	var stepContext uintptr
	_, _, _ = i, simIndex, stepContext
	_ = uint64FromInt64(4)
	stepContext = taskContext
	if !(startIndex <= endIndex) && b2InternalAssertFcn(tls, __ccgo_ts+12545, __ccgo_ts+12460, int32FromInt32(1177)) != 0 {
		__builtin_trap(tls)
	}
	i = startIndex
	for {
		if !(i < endIndex) {
			break
		}
		simIndex = *(*int32)(unsafe.Pointer((*b2StepContext)(unsafe.Pointer(stepContext)).BulletBodies + uintptr(i)*4))
		b2SolveContinuous(tls, (*b2StepContext)(unsafe.Pointer(stepContext)).World, simIndex)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b2TransferBody(tls *_Stack, world uintptr, targetSet uintptr, sourceSet uintptr, body uintptr) {
	var movedBody, movedSim, sourceSim, state, targetSim, v1, v13, v15, v17, v20, v22, v3, v5, v7, v9 uintptr
	var movedId, movedIndex, movedIndex1, movedIndex2, newCapacity, newCapacity1, sourceIndex, targetIndex, v10, v11, v14, v18, v2, v21, v6 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = movedBody, movedId, movedIndex, movedIndex1, movedIndex2, movedSim, newCapacity, newCapacity1, sourceIndex, sourceSim, state, targetIndex, targetSim, v1, v10, v11, v13, v14, v15, v17, v18, v2, v20, v21, v22, v3, v5, v6, v7, v9
	if !(targetSet != sourceSet) && b2InternalAssertFcn(tls, __ccgo_ts+14968, __ccgo_ts+14063, int32FromInt32(523)) != 0 {
		__builtin_trap(tls)
	}
	sourceIndex = (*b2Body)(unsafe.Pointer(body)).LocalIndex
	v1 = sourceSet
	v2 = sourceIndex
	if !(0 <= v2 && v2 < (*b2BodySimArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodySimArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*100
	goto _4
_4:
	sourceSim = v3
	targetIndex = (*b2SolverSet)(unsafe.Pointer(targetSet)).BodySims.Count
	v5 = targetSet
	if (*b2BodySimArray)(unsafe.Pointer(v5)).Count == (*b2BodySimArray)(unsafe.Pointer(v5)).Capacity {
		if (*b2BodySimArray)(unsafe.Pointer(v5)).Capacity < int32(2) {
			v6 = int32(2)
		} else {
			v6 = (*b2BodySimArray)(unsafe.Pointer(v5)).Capacity + (*b2BodySimArray)(unsafe.Pointer(v5)).Capacity>>int32(1)
		}
		newCapacity = v6
		b2BodySimArray_Reserve(tls, v5, newCapacity)
	}
	*(*int32)(unsafe.Pointer(v5 + 8)) += int32(1)
	v7 = (*b2BodySimArray)(unsafe.Pointer(v5)).Data + uintptr((*b2BodySimArray)(unsafe.Pointer(v5)).Count-int32FromInt32(1))*100
	goto _8
_8:
	targetSim = v7
	memcpy(tls, targetSim, sourceSim, uint64(100))
	v9 = sourceSet
	v10 = sourceIndex
	if !(0 <= v10 && v10 < (*b2BodySimArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	movedIndex = -int32(1)
	if v10 != (*b2BodySimArray)(unsafe.Pointer(v9)).Count-int32FromInt32(1) {
		movedIndex = (*b2BodySimArray)(unsafe.Pointer(v9)).Count - int32(1)
		*(*b2BodySim)(unsafe.Pointer((*b2BodySimArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*100)) = *(*b2BodySim)(unsafe.Pointer((*b2BodySimArray)(unsafe.Pointer(v9)).Data + uintptr(movedIndex)*100))
	}
	*(*int32)(unsafe.Pointer(v9 + 8)) -= int32(1)
	v11 = movedIndex
	goto _12
_12:
	// Remove body sim from solver set that owns it
	movedIndex2 = v11
	if movedIndex2 != -int32(1) {
		// Fix moved body index
		movedSim = (*b2SolverSet)(unsafe.Pointer(sourceSet)).BodySims.Data + uintptr(sourceIndex)*100
		movedId = (*b2BodySim)(unsafe.Pointer(movedSim)).BodyId
		v13 = world + 1024
		v14 = movedId
		if !(0 <= v14 && v14 < (*b2BodyArray)(unsafe.Pointer(v13)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
			__builtin_trap(tls)
		}
		v15 = (*b2BodyArray)(unsafe.Pointer(v13)).Data + uintptr(v14)*128
		goto _16
	_16:
		movedBody = v15
		if !((*b2Body)(unsafe.Pointer(movedBody)).LocalIndex == movedIndex2) && b2InternalAssertFcn(tls, __ccgo_ts+1407, __ccgo_ts+14063, int32FromInt32(540)) != 0 {
			__builtin_trap(tls)
		}
		(*b2Body)(unsafe.Pointer(movedBody)).LocalIndex = sourceIndex
	}
	if (*b2SolverSet)(unsafe.Pointer(sourceSet)).SetIndex == int32(b2_awakeSet) {
		v17 = sourceSet + 16
		v18 = sourceIndex
		if !(0 <= v18 && v18 < (*b2BodyStateArray)(unsafe.Pointer(v17)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(194)) != 0 {
			__builtin_trap(tls)
		}
		movedIndex1 = -int32(1)
		if v18 != (*b2BodyStateArray)(unsafe.Pointer(v17)).Count-int32FromInt32(1) {
			movedIndex1 = (*b2BodyStateArray)(unsafe.Pointer(v17)).Count - int32(1)
			*(*b2BodyState)(unsafe.Pointer((*b2BodyStateArray)(unsafe.Pointer(v17)).Data + uintptr(v18)*32)) = *(*b2BodyState)(unsafe.Pointer((*b2BodyStateArray)(unsafe.Pointer(v17)).Data + uintptr(movedIndex1)*32))
		}
		*(*int32)(unsafe.Pointer(v17 + 8)) -= int32(1)
		_ = movedIndex1
		goto _19
	_19:
	} else {
		if (*b2SolverSet)(unsafe.Pointer(targetSet)).SetIndex == int32(b2_awakeSet) {
			v20 = targetSet + 16
			if (*b2BodyStateArray)(unsafe.Pointer(v20)).Count == (*b2BodyStateArray)(unsafe.Pointer(v20)).Capacity {
				if (*b2BodyStateArray)(unsafe.Pointer(v20)).Capacity < int32(2) {
					v21 = int32(2)
				} else {
					v21 = (*b2BodyStateArray)(unsafe.Pointer(v20)).Capacity + (*b2BodyStateArray)(unsafe.Pointer(v20)).Capacity>>int32(1)
				}
				newCapacity1 = v21
				b2BodyStateArray_Reserve(tls, v20, newCapacity1)
			}
			*(*int32)(unsafe.Pointer(v20 + 8)) += int32(1)
			v22 = (*b2BodyStateArray)(unsafe.Pointer(v20)).Data + uintptr((*b2BodyStateArray)(unsafe.Pointer(v20)).Count-int32FromInt32(1))*32
			goto _23
		_23:
			state = v22
			*(*b2BodyState)(unsafe.Pointer(state)) = b2_identityBodyState15
		}
	}
	(*b2Body)(unsafe.Pointer(body)).SetIndex = (*b2SolverSet)(unsafe.Pointer(targetSet)).SetIndex
	(*b2Body)(unsafe.Pointer(body)).LocalIndex = targetIndex
}

func b2DefaultBodyDef(tls *_Stack) (r BodyDef) {
	var def BodyDef
	_ = def
	def = BodyDef{}
	def.Type1 = int32(b2_staticBody)
	def.Rotation = b2Rot_identity
	def.SleepThreshold = float32(float32FromFloat32(0.05) * b2_lengthUnitsPerMeter)
	def.GravityScale = float32FromFloat32(1)
	def.EnableSleep = boolUint8(true1 != 0)
	def.IsAwake = boolUint8(true1 != 0)
	def.IsEnabled = boolUint8(true1 != 0)
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}

func b2BodyMoveEventArray_Create(tls *_Stack, capacity int32) (r b2BodyMoveEventArray) {
	var a b2BodyMoveEventArray
	_ = a
	a = b2BodyMoveEventArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(40)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2BodyMoveEventArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2BodyMoveEventArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2BodyMoveEventArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2BodyMoveEventArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2BodyMoveEventArray)(unsafe.Pointer(a)).Capacity)*uint64(40)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(40)))
	(*b2BodyMoveEventArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2BodyMoveEventArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2BodyMoveEventArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2BodyMoveEventArray)(unsafe.Pointer(a)).Capacity)*uint64(40)))
	(*b2BodyMoveEventArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2BodyMoveEventArray)(unsafe.Pointer(a)).Count = 0
	(*b2BodyMoveEventArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2Body_IsValid(tls *_Stack, id BodyId) (r uint8) {
	var body, world uintptr
	_, _ = body, world
	if int32(_B2_MAX_WORLDS) <= int32FromUint16(id.World0) {
		// invalid world
		return boolUint8(false1 != 0)
	}
	world = uintptr(unsafe.Pointer(&b2_worlds)) + uintptr(id.World0)*1792
	if int32FromUint16((*b2World)(unsafe.Pointer(world)).WorldId) != int32FromUint16(id.World0) {
		// world is free
		return boolUint8(false1 != 0)
	}
	if id.Index1 < int32(1) || (*b2World)(unsafe.Pointer(world)).Bodies.Count < id.Index1 {
		// invalid index
		return boolUint8(false1 != 0)
	}
	body = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(id.Index1-int32FromInt32(1))*128
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == -int32(1) {
		// this was freed
		return boolUint8(false1 != 0)
	}
	if !((*b2Body)(unsafe.Pointer(body)).LocalIndex != -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+16112, __ccgo_ts+15342, int32FromInt32(1603)) != 0 {
		__builtin_trap(tls)
	}
	if int32FromUint16((*b2Body)(unsafe.Pointer(body)).Generation) != int32FromUint16(id.Generation) {
		// this id is orphaned
		return boolUint8(false1 != 0)
	}
	return boolUint8(true1 != 0)
}
