package b2

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

// C documentation
//
//	// Refresh the body sim index on associated contact sims.
func b2RefreshBodyContactIndices(tls *_Stack, world uintptr, body1 uintptr) {
	var contact, contactSim, v1 uintptr
	var contactId, contactKey, edgeIndex, encodedIndex, v2 int32
	var v4 bool
	_, _, _, _, _, _, _, _, _ = contact, contactId, contactKey, contactSim, edgeIndex, encodedIndex, v1, v2, v4
	v1 = body1
	if (*b2Body)(unsafe.Pointer(v1)).SetIndex == int32(b2_awakeSet) {
		v2 = (*b2Body)(unsafe.Pointer(v1)).LocalIndex
		goto _3
	}
	if (*b2Body)(unsafe.Pointer(v1)).SetIndex == int32(b2_staticSet) {
		v2 = -((*b2Body)(unsafe.Pointer(v1)).LocalIndex + int32(2))
		goto _3
	}
	v2 = -int32(1)
	goto _3
_3:
	encodedIndex = v2
	contactKey = (*b2Body)(unsafe.Pointer(body1)).HeadContactKey
	for contactKey != -int32(1) {
		contactId = contactKey >> int32(1)
		edgeIndex = contactKey & int32(1)
		if v4 = !!(0 <= contactId && contactId < (*b2World)(unsafe.Pointer(world)).Contacts.Count); !v4 {
			b2InternalAssert(tls, __ccgo_ts+301, __ccgo_ts+365, int32FromInt32(45))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		contact = (*b2World)(unsafe.Pointer(world)).Contacts.Data + uintptr(contactId)*64
		contactSim = b2GetContactSim(tls, world, contact)
		if edgeIndex == 0 {
			(*b2ContactSim)(unsafe.Pointer(contactSim)).EncodedBodySimA = encodedIndex
		} else {
			(*b2ContactSim)(unsafe.Pointer(contactSim)).EncodedBodySimB = encodedIndex
		}
		contactKey = (*(*b2ContactEdge)(unsafe.Add(nil, contact+uintptr(edgeIndex)*12))).NextKey
	}
}

// C documentation
//
//	// Get a validated body from a world using an id.
func b2GetBodyFullId(tls *_Stack, world uintptr, bodyId BodyId) (r uintptr) {
	var v1 bool
	_ = v1
	if v1 = !!(b2Body_IsValid(tls, bodyId) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+387, __ccgo_ts+365, int32FromInt32(64))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// id index starts at one so that zero can represent null
	if v1 = !!(0 <= bodyId.Index1-int32(1) && bodyId.Index1-int32(1) < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+412, __ccgo_ts+365, int32FromInt32(67))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(bodyId.Index1-int32FromInt32(1))*104
}

func b2GetBodyTransformQuick(tls *_Stack, world uintptr, body uintptr) (r WorldTransform) {
	var bodySim, set uintptr
	var v1, v2 bool
	_, _, _, _ = bodySim, set, v1, v2
	if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(body)).SetIndex && (*b2Body)(unsafe.Pointer(body)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+365, int32FromInt32(72))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(body)).SetIndex)*88
	if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(body)).LocalIndex && (*b2Body)(unsafe.Pointer(body)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+365, int32FromInt32(73))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(body)).LocalIndex)*96
	return (*b2BodySim)(unsafe.Pointer(bodySim)).Transform
}

func b2GetBodyTransform(tls *_Stack, world uintptr, bodyId int32) (r WorldTransform) {
	var body uintptr
	var v1 bool
	_, _ = body, v1
	if v1 = !!(0 <= bodyId && bodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+642, __ccgo_ts+365, int32FromInt32(79))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	body = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(bodyId)*104
	return b2GetBodyTransformQuick(tls, world, body)
}

// C documentation
//
//	// Create a b2BodyId from a raw id.
func b2MakeBodyId(tls *_Stack, world uintptr, bodyId int32) (r BodyId) {
	var body uintptr
	var v1 bool
	_, _ = body, v1
	if v1 = !!(0 <= bodyId && bodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+642, __ccgo_ts+365, int32FromInt32(86))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	body = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(bodyId)*104
	return BodyId{
		Index1:     bodyId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Body)(unsafe.Pointer(body)).Generation,
	}
}

// C documentation
//
//	// Sync the body flags when something external changes them,
//	// such as the user changing the motion locks.
func b2SyncBodyFlags(tls *_Stack, world2 uintptr, body2 uintptr) {
	var bodySim, bodySim1, bodyState, set, set1, v1, v2, v5 uintptr
	var v3, v4 bool
	_, _, _, _, _, _, _, _, _, _ = bodySim, bodySim1, bodyState, set, set1, v1, v2, v3, v4, v5
	v1 = world2
	v2 = body2
	if v3 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).SetIndex && (*b2Body)(unsafe.Pointer(v2)).SetIndex < (*b2World)(unsafe.Pointer(v1)).SolverSets.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v1)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).SetIndex)*88
	if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).LocalIndex && (*b2Body)(unsafe.Pointer(v2)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).LocalIndex)*96
	v5 = bodySim
	goto _6
_6:
	bodySim1 = v5
	// Preserve the sim only flags: fast for contact recycling, time of impact for debug draw.
	(*b2BodySim)(unsafe.Pointer(bodySim1)).Flags = (*b2BodySim)(unsafe.Pointer(bodySim1)).Flags&uint32FromInt32(int32(b2_isFast)|int32(b2_hadTimeOfImpact)) | (*b2Body)(unsafe.Pointer(body2)).Flags&uint32FromInt32(^int32(b2_bodyTransientFlags))
	v1 = world2
	v2 = body2
	if (*b2Body)(unsafe.Pointer(v2)).SetIndex == int32(b2_awakeSet) {
		if v3 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(v1)).SolverSets.Count); !v3 {
			b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+698, int32FromInt32(254))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		set1 = (*b2World)(unsafe.Pointer(v1)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
		if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).LocalIndex && (*b2Body)(unsafe.Pointer(v2)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set1)).BodyStates.Count); !v4 {
			b2InternalAssert(tls, __ccgo_ts+790, __ccgo_ts+698, int32FromInt32(255))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		v5 = (*b2SolverSet)(unsafe.Pointer(set1)).BodyStates.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).LocalIndex)*32
		goto _11
	}
	v5 = uintptrFromInt32(0)
	goto _11
_11:
	bodyState = v5
	if bodyState != uintptrFromInt32(0) {
		(*b2BodyState)(unsafe.Pointer(bodyState)).Flags = (*b2Body)(unsafe.Pointer(body2)).Flags & uint32FromInt32(^int32(b2_bodyTransientFlags))
	}
}

func b2CreateIslandForBody(tls *_Stack, world uintptr, setIndex int32, body uintptr) {
	var elementSize, newCapacity, newSize, oldSize, v3 int32
	var island uintptr
	var v1 bool
	_, _, _, _, _, _, _ = elementSize, island, newCapacity, newSize, oldSize, v1, v3
	if v1 = !!((*b2Body)(unsafe.Pointer(body)).IslandId == -int32FromInt32(1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+868, __ccgo_ts+365, int32FromInt32(108))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(setIndex != int32(b2_disabledSet)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+900, __ccgo_ts+365, int32FromInt32(109))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	island = b2CreateIsland(tls, world, setIndex)
	elementSize = int32(4)
	if (*b2Island)(unsafe.Pointer(island)).Bodies.Count >= (*b2Island)(unsafe.Pointer(island)).Bodies.Capacity {
		oldSize = (*b2Island)(unsafe.Pointer(island)).Bodies.Capacity * elementSize
		if (*b2Island)(unsafe.Pointer(island)).Bodies.Capacity == 0 {
			v3 = int32(8)
		} else {
			v3 = int32(2) * (*b2Island)(unsafe.Pointer(island)).Bodies.Capacity
		}
		newCapacity = v3
		newSize = newCapacity * elementSize
		(*b2Island)(unsafe.Pointer(island)).Bodies.Data = b2GrowAlloc(tls, (*b2Island)(unsafe.Pointer(island)).Bodies.Data, uint64FromInt32(oldSize), uint64FromInt32(newSize))
		(*b2Island)(unsafe.Pointer(island)).Bodies.Capacity = newCapacity
	}
	*(*int32)(unsafe.Add(nil, (*b2Island)(unsafe.Pointer(island)).Bodies.Data+uintptr((*b2Island)(unsafe.Pointer(island)).Bodies.Count)*4)) = (*b2Body)(unsafe.Pointer(body)).Id
	(*b2Island)(unsafe.Pointer(island)).Bodies.Count += int32(1)
	(*b2Body)(unsafe.Pointer(body)).IslandId = (*b2Island)(unsafe.Pointer(island)).IslandId
	(*b2Body)(unsafe.Pointer(body)).IslandIndex = 0
	b2ValidateIsland(tls, world, (*b2Island)(unsafe.Pointer(island)).IslandId)
}

func b2RemoveBodyFromIsland(tls *_Stack, world uintptr, body uintptr) {
	var island uintptr
	var islandId, localIndex, movedBodyId int32
	var v1 bool
	_, _, _, _, _ = island, islandId, localIndex, movedBodyId, v1
	if (*b2Body)(unsafe.Pointer(body)).IslandId == -int32(1) {
		if v1 = !!((*b2Body)(unsafe.Pointer(body)).IslandIndex == -int32FromInt32(1)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+927, __ccgo_ts+365, int32FromInt32(123))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	islandId = (*b2Body)(unsafe.Pointer(body)).IslandId
	if v1 = !!(0 <= islandId && islandId < (*b2World)(unsafe.Pointer(world)).Islands.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+962, __ccgo_ts+365, int32FromInt32(128))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	island = (*b2World)(unsafe.Pointer(world)).Islands.Data + uintptr(islandId)*64
	localIndex = (*b2Body)(unsafe.Pointer(body)).IslandIndex
	movedBodyId = *(*int32)(unsafe.Add(nil, (*b2Island)(unsafe.Pointer(island)).Bodies.Data+uintptr((*b2Island)(unsafe.Pointer(island)).Bodies.Count-int32(1))*4))
	*(*int32)(unsafe.Add(nil, (*b2Island)(unsafe.Pointer(island)).Bodies.Data+uintptr(localIndex)*4)) = movedBodyId
	(*(*b2Body)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Bodies.Data+uintptr(movedBodyId)*104))).IslandIndex = localIndex
	(*b2Island)(unsafe.Pointer(island)).Bodies.Count -= int32(1)
	if (*b2Island)(unsafe.Pointer(island)).Bodies.Count == 0 {
		// Destroy empty island
		if v1 = !!((*b2Island)(unsafe.Pointer(island)).Contacts.Count == int32FromInt32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1023, __ccgo_ts+365, int32FromInt32(141))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		if v1 = !!((*b2Island)(unsafe.Pointer(island)).Joints.Count == int32FromInt32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1051, __ccgo_ts+365, int32FromInt32(142))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		// Free the island
		b2DestroyIsland(tls, world, (*b2Island)(unsafe.Pointer(island)).IslandId)
	} else {
		b2ValidateIsland(tls, world, islandId)
	}
	(*b2Body)(unsafe.Pointer(body)).IslandId = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).IslandIndex = -int32(1)
}

func b2DestroyBodyContacts(tls *_Stack, world uintptr, body uintptr) {
	var contact uintptr
	var contactId, edgeIndex, edgeKey int32
	var v1 bool
	_, _, _, _, _ = contact, contactId, edgeIndex, edgeKey, v1
	// Destroy the attached contacts
	edgeKey = (*b2Body)(unsafe.Pointer(body)).HeadContactKey
	for edgeKey != -int32(1) {
		contactId = edgeKey >> int32(1)
		edgeIndex = edgeKey & int32(1)
		if v1 = !!(0 <= contactId && contactId < (*b2World)(unsafe.Pointer(world)).Contacts.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+301, __ccgo_ts+365, int32FromInt32(165))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		contact = (*b2World)(unsafe.Pointer(world)).Contacts.Data + uintptr(contactId)*64
		edgeKey = (*(*b2ContactEdge)(unsafe.Add(nil, contact+uintptr(edgeIndex)*12))).NextKey
		b2DestroyContact(tls, world, contact)
	}
	b2ValidateSolverSets(tls, world)
}

func b2CreateBody(tls *_Stack, worldId WorldId, def uintptr) (r BodyId) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var body, bodySim, bodyState, set, world, v19, v20, v21, v24, v27 uintptr
	var bodyId, elementSize1, elementSize2, newCapacity, newCapacity1, newCapacity2, newSize, newSize1, newSize2, oldCapacity, oldSize, oldSize1, oldSize2, setId, v12, v15, v16 int32
	var id BodyId
	var isAwake uint8
	var lockFlags uint32_t
	var v1 bool
	var _ /* _ca at bp+0 */ b2RecArgs_CreateBody
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodyId, bodySim, bodyState, elementSize1, elementSize2, id, isAwake, lockFlags, newCapacity, newCapacity1, newCapacity2, newSize, newSize1, newSize2, oldCapacity, oldSize, oldSize1, oldSize2, set, setId, world, v1, v12, v15, v16, v19, v20, v21, v24, v27
	if v1 = !!((*BodyDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1077, __ccgo_ts+365, int32FromInt32(175))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if int32FromUint8(b2IsValidPos(tls, (*BodyDef)(unsafe.Pointer(def)).Position)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+104, __ccgo_ts+1141, uintptr(unsafe.Pointer(&__func__1))))
		if v1 = !!(b2IsValidPos(tls, (*BodyDef)(unsafe.Pointer(def)).Position) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1141, __ccgo_ts+365, int32FromInt32(176))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return b2_nullBodyId
	}
	if int32FromUint8(b2IsValidRot(tls, (*BodyDef)(unsafe.Pointer(def)).Rotation)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+104, __ccgo_ts+1171, uintptr(unsafe.Pointer(&__func__1))))
		if v1 = !!(b2IsValidRot(tls, (*BodyDef)(unsafe.Pointer(def)).Rotation) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1171, __ccgo_ts+365, int32FromInt32(177))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return b2_nullBodyId
	}
	if int32FromUint8(b2IsValidVec2(tls, (*BodyDef)(unsafe.Pointer(def)).LinearVelocity)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+104, __ccgo_ts+1201, uintptr(unsafe.Pointer(&__func__1))))
		if v1 = !!(b2IsValidVec2(tls, (*BodyDef)(unsafe.Pointer(def)).LinearVelocity) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1201, __ccgo_ts+365, int32FromInt32(178))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return b2_nullBodyId
	}
	if int32FromUint8(b2IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).AngularVelocity)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+104, __ccgo_ts+1238, uintptr(unsafe.Pointer(&__func__1))))
		if v1 = !!(b2IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).AngularVelocity) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1238, __ccgo_ts+365, int32FromInt32(179))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return b2_nullBodyId
	}
	if boolInt32(b2IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).LinearDamping) != 0 && (*BodyDef)(unsafe.Pointer(def)).LinearDamping >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+104, __ccgo_ts+1277, uintptr(unsafe.Pointer(&__func__1))))
		if v1 = !!(b2IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).LinearDamping) != 0 && (*BodyDef)(unsafe.Pointer(def)).LinearDamping >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1277, __ccgo_ts+365, int32FromInt32(180))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return b2_nullBodyId
	}
	if boolInt32(b2IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).AngularDamping) != 0 && (*BodyDef)(unsafe.Pointer(def)).AngularDamping >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+104, __ccgo_ts+1344, uintptr(unsafe.Pointer(&__func__1))))
		if v1 = !!(b2IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).AngularDamping) != 0 && (*BodyDef)(unsafe.Pointer(def)).AngularDamping >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1344, __ccgo_ts+365, int32FromInt32(181))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return b2_nullBodyId
	}
	if boolInt32(b2IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).SleepThreshold) != 0 && (*BodyDef)(unsafe.Pointer(def)).SleepThreshold >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+104, __ccgo_ts+1413, uintptr(unsafe.Pointer(&__func__1))))
		if v1 = !!(b2IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).SleepThreshold) != 0 && (*BodyDef)(unsafe.Pointer(def)).SleepThreshold >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1413, __ccgo_ts+365, int32FromInt32(182))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return b2_nullBodyId
	}
	if boolInt32(b2IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).SafetyFactor) != 0 && (*BodyDef)(unsafe.Pointer(def)).SafetyFactor >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+104, __ccgo_ts+1482, uintptr(unsafe.Pointer(&__func__1))))
		if v1 = !!(b2IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).SafetyFactor) != 0 && (*BodyDef)(unsafe.Pointer(def)).SafetyFactor >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1482, __ccgo_ts+365, int32FromInt32(183))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return b2_nullBodyId
	}
	if int32FromUint8(b2IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).GravityScale)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+104, __ccgo_ts+1547, uintptr(unsafe.Pointer(&__func__1))))
		if v1 = !!(b2IsValidFloat(tls, (*BodyDef)(unsafe.Pointer(def)).GravityScale) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1547, __ccgo_ts+365, int32FromInt32(184))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return b2_nullBodyId
	}
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+365, int32FromInt32(187))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
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
				setId = b2AllocId(tls, world+1952)
				if setId == (*b2World)(unsafe.Pointer(world)).SolverSets.Count {
					// Create a zero initialized solver set. All sub-arrays are also zero initialized.
					elementSize1 = int32(88)
					if (*b2World)(unsafe.Pointer(world)).SolverSets.Count >= (*b2World)(unsafe.Pointer(world)).SolverSets.Capacity {
						oldSize1 = (*b2World)(unsafe.Pointer(world)).SolverSets.Capacity * elementSize1
						if (*b2World)(unsafe.Pointer(world)).SolverSets.Capacity == 0 {
							v12 = int32(8)
						} else {
							v12 = int32(2) * (*b2World)(unsafe.Pointer(world)).SolverSets.Capacity
						}
						newCapacity1 = v12
						newSize1 = newCapacity1 * elementSize1
						(*b2World)(unsafe.Pointer(world)).SolverSets.Data = b2GrowAlloc(tls, (*b2World)(unsafe.Pointer(world)).SolverSets.Data, uint64FromInt32(oldSize1), uint64FromInt32(newSize1))
						(*b2World)(unsafe.Pointer(world)).SolverSets.Capacity = newCapacity1
					}
					*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr((*b2World)(unsafe.Pointer(world)).SolverSets.Count)*88)) = b2SolverSet{}
					(*b2World)(unsafe.Pointer(world)).SolverSets.Count += int32(1)
				} else {
					if v1 = !!((*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(setId)*88))).SetIndex == -int32FromInt32(1)); !v1 {
						b2InternalAssert(tls, __ccgo_ts+1606, __ccgo_ts+365, int32FromInt32(222))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
				}
				(*(*b2SolverSet)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SolverSets.Data+uintptr(setId)*88))).SetIndex = setId
			}
		}
	}
	if v1 = !!(0 <= setId && setId < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1662, __ccgo_ts+365, int32FromInt32(228))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyId = b2AllocId(tls, world+1912)
	lockFlags = uint32(0)
	if (*BodyDef)(unsafe.Pointer(def)).MotionLocks.LinearX != 0 {
		v12 = int32(b2_lockLinearX)
	} else {
		v12 = 0
	}
	lockFlags = lockFlags | uint32FromInt32(v12)
	if (*BodyDef)(unsafe.Pointer(def)).MotionLocks.LinearY != 0 {
		v12 = int32(b2_lockLinearY)
	} else {
		v12 = 0
	}
	lockFlags = lockFlags | uint32FromInt32(v12)
	if (*BodyDef)(unsafe.Pointer(def)).MotionLocks.AngularZ != 0 {
		v12 = int32(b2_lockAngularZ)
	} else {
		v12 = 0
	}
	lockFlags = lockFlags | uint32FromInt32(v12)
	if v1 = !!(0 <= setId && setId < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1708, __ccgo_ts+365, int32FromInt32(237))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setId)*88
	v19 = set
	v20 = set + 8
	v21 = set + 12
	v12 = int32(96)
	if *(*int32)(unsafe.Add(nil, v20)) >= *(*int32)(unsafe.Add(nil, v21)) {
		oldCapacity = *(*int32)(unsafe.Add(nil, v21))
		oldSize = oldCapacity * v12
		if oldCapacity == 0 {
			v15 = int32(8)
		} else {
			v15 = int32(2) * oldCapacity
		}
		newCapacity = v15
		newSize = newCapacity * v12
		*(*uintptr)(unsafe.Add(nil, v19)) = b2GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v19)), uint64FromInt32(oldSize), uint64FromInt32(newSize))
		*(*int32)(unsafe.Add(nil, v21)) = newCapacity
	}
	v27 = v20
	v16 = *(*int32)(unsafe.Pointer(v27))
	*(*int32)(unsafe.Pointer(v27)) = *(*int32)(unsafe.Pointer(v27)) + 1
	v24 = *(*uintptr)(unsafe.Add(nil, v19)) + uintptr(v16*v12)
	goto _25
_25:
	bodySim = v24
	*(*b2BodySim)(unsafe.Add(nil, bodySim)) = b2BodySim{}
	(*b2BodySim)(unsafe.Pointer(bodySim)).Transform.P = (*BodyDef)(unsafe.Pointer(def)).Position
	(*b2BodySim)(unsafe.Pointer(bodySim)).Transform.Q = (*BodyDef)(unsafe.Pointer(def)).Rotation
	(*b2BodySim)(unsafe.Pointer(bodySim)).Center = (*BodyDef)(unsafe.Pointer(def)).Position
	(*b2BodySim)(unsafe.Pointer(bodySim)).Rotation0 = (*b2BodySim)(unsafe.Pointer(bodySim)).Transform.Q
	(*b2BodySim)(unsafe.Pointer(bodySim)).Center0 = (*b2BodySim)(unsafe.Pointer(bodySim)).Center
	(*b2BodySim)(unsafe.Pointer(bodySim)).MinExtent = float32(float32(float32FromFloat32(100)*float32FromFloat32(1000)) * b2_lengthUnitsPerMeter)
	(*b2BodySim)(unsafe.Pointer(bodySim)).MaxExtent = float32FromFloat32(0)
	(*b2BodySim)(unsafe.Pointer(bodySim)).LinearDamping = (*BodyDef)(unsafe.Pointer(def)).LinearDamping
	(*b2BodySim)(unsafe.Pointer(bodySim)).AngularDamping = (*BodyDef)(unsafe.Pointer(def)).AngularDamping
	(*b2BodySim)(unsafe.Pointer(bodySim)).GravityScale = (*BodyDef)(unsafe.Pointer(def)).GravityScale
	(*b2BodySim)(unsafe.Pointer(bodySim)).BodyId = bodyId
	(*b2BodySim)(unsafe.Pointer(bodySim)).Flags = lockFlags
	if (*BodyDef)(unsafe.Pointer(def)).IsBullet != 0 {
		v12 = int32(b2_isBullet)
	} else {
		v12 = 0
	}
	*(*uint32_t)(unsafe.Add(nil, bodySim+92)) |= uint32FromInt32(v12)
	if (*BodyDef)(unsafe.Pointer(def)).AllowFastRotation != 0 {
		v12 = int32(b2_allowFastRotation)
	} else {
		v12 = 0
	}
	*(*uint32_t)(unsafe.Add(nil, bodySim+92)) |= uint32FromInt32(v12)
	if (*BodyDef)(unsafe.Pointer(def)).Type1 == int32(b2_dynamicBody) {
		v12 = int32(b2_dynamicFlag)
	} else {
		v12 = 0
	}
	*(*uint32_t)(unsafe.Add(nil, bodySim+92)) |= uint32FromInt32(v12)
	if (*BodyDef)(unsafe.Pointer(def)).EnableSleep != 0 {
		v12 = int32(b2_enableSleep)
	} else {
		v12 = 0
	}
	*(*uint32_t)(unsafe.Add(nil, bodySim+92)) |= uint32FromInt32(v12)
	if (*BodyDef)(unsafe.Pointer(def)).EnableContactRecycling != 0 {
		v12 = int32(b2_bodyEnableContactRecycling)
	} else {
		v12 = 0
	}
	*(*uint32_t)(unsafe.Add(nil, bodySim+92)) |= uint32FromInt32(v12)
	if setId == int32(b2_awakeSet) {
		v19 = set + 16
		v20 = set + 16 + 8
		v21 = set + 16 + 12
		v12 = int32(32)
		if *(*int32)(unsafe.Add(nil, v20)) >= *(*int32)(unsafe.Add(nil, v21)) {
			oldCapacity = *(*int32)(unsafe.Add(nil, v21))
			oldSize = oldCapacity * v12
			if oldCapacity == 0 {
				v15 = int32(8)
			} else {
				v15 = int32(2) * oldCapacity
			}
			newCapacity = v15
			newSize = newCapacity * v12
			*(*uintptr)(unsafe.Add(nil, v19)) = b2GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v19)), uint64FromInt32(oldSize), uint64FromInt32(newSize))
			*(*int32)(unsafe.Add(nil, v21)) = newCapacity
		}
		v27 = v20
		v16 = *(*int32)(unsafe.Pointer(v27))
		*(*int32)(unsafe.Pointer(v27)) = *(*int32)(unsafe.Pointer(v27)) + 1
		v24 = *(*uintptr)(unsafe.Add(nil, v19)) + uintptr(v16*v12)
		goto _39
	_39:
		bodyState = v24
		if v1 = !!(uint64(bodyState)&uint64FromInt32(0x1F) == uint64FromInt32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1766, __ccgo_ts+365, int32FromInt32(261))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		*(*b2BodyState)(unsafe.Add(nil, bodyState)) = b2BodyState{}
		(*b2BodyState)(unsafe.Pointer(bodyState)).LinearVelocity = (*BodyDef)(unsafe.Pointer(def)).LinearVelocity
		(*b2BodyState)(unsafe.Pointer(bodyState)).AngularVelocity = (*BodyDef)(unsafe.Pointer(def)).AngularVelocity
		(*b2BodyState)(unsafe.Pointer(bodyState)).DeltaRotation = b2Rot_identity
		(*b2BodyState)(unsafe.Pointer(bodyState)).Flags = (*b2BodySim)(unsafe.Pointer(bodySim)).Flags
	}
	if bodyId == (*b2World)(unsafe.Pointer(world)).Bodies.Count {
		elementSize2 = int32(104)
		if (*b2World)(unsafe.Pointer(world)).Bodies.Count >= (*b2World)(unsafe.Pointer(world)).Bodies.Capacity {
			oldSize2 = (*b2World)(unsafe.Pointer(world)).Bodies.Capacity * elementSize2
			if (*b2World)(unsafe.Pointer(world)).Bodies.Capacity == 0 {
				v12 = int32(8)
			} else {
				v12 = int32(2) * (*b2World)(unsafe.Pointer(world)).Bodies.Capacity
			}
			newCapacity2 = v12
			newSize2 = newCapacity2 * elementSize2
			(*b2World)(unsafe.Pointer(world)).Bodies.Data = b2GrowAlloc(tls, (*b2World)(unsafe.Pointer(world)).Bodies.Data, uint64FromInt32(oldSize2), uint64FromInt32(newSize2))
			(*b2World)(unsafe.Pointer(world)).Bodies.Capacity = newCapacity2
		}
		*(*b2Body)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Bodies.Data+uintptr((*b2World)(unsafe.Pointer(world)).Bodies.Count)*104)) = b2Body{}
		(*b2World)(unsafe.Pointer(world)).Bodies.Count += int32(1)
	} else {
		if v1 = !!((*(*b2Body)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Bodies.Data+uintptr(bodyId)*104))).Id == -int32FromInt32(1)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1803, __ccgo_ts+365, int32FromInt32(276))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
	}
	if v1 = !!(0 <= bodyId && bodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+642, __ccgo_ts+365, int32FromInt32(279))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	body = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(bodyId)*104
	if (*BodyDef)(unsafe.Pointer(def)).name != 0 {
		strncpy(tls, body+90, (*BodyDef)(unsafe.Pointer(def)).name, uint64(_B2_NAME_LENGTH))
		*(*uint8)(unsafe.Add(nil, body+90+10)) = uint8(0)
	} else {
		memset(tls, body+90, 0, uint64(11))
	}
	(*b2Body)(unsafe.Pointer(body)).UserData = uintptr((*BodyDef)(unsafe.Pointer(def)).UserData)
	(*b2Body)(unsafe.Pointer(body)).SetIndex = setId
	(*b2Body)(unsafe.Pointer(body)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count - int32(1)
	v19 = body + 88
	*(*uint16_t)(unsafe.Pointer(v19)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v19))) + int32FromInt32(1))
	(*b2Body)(unsafe.Pointer(body)).HeadShapeId = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).ShapeCount = 0
	(*b2Body)(unsafe.Pointer(body)).HeadChainId = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).HeadContactKey = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).ContactCount = 0
	(*b2Body)(unsafe.Pointer(body)).HeadJointKey = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).JointCount = 0
	(*b2Body)(unsafe.Pointer(body)).IslandId = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).IslandIndex = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).BodyMoveIndex = -int32(1)
	(*b2Body)(unsafe.Pointer(body)).Id = bodyId
	(*b2Body)(unsafe.Pointer(body)).Mass = float32FromFloat32(0)
	(*b2Body)(unsafe.Pointer(body)).Inertia = float32FromFloat32(0)
	(*b2Body)(unsafe.Pointer(body)).SleepThreshold = (*BodyDef)(unsafe.Pointer(def)).SleepThreshold
	(*b2Body)(unsafe.Pointer(body)).SleepTime = float32FromFloat32(0)
	(*b2Body)(unsafe.Pointer(body)).SafetyFactor = (*BodyDef)(unsafe.Pointer(def)).SafetyFactor
	(*b2Body)(unsafe.Pointer(body)).Type1 = (*BodyDef)(unsafe.Pointer(def)).Type1
	(*b2Body)(unsafe.Pointer(body)).Flags = (*b2BodySim)(unsafe.Pointer(bodySim)).Flags
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
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_CreateBody)(unsafe.Add(nil, bp)) = b2RecArgs_CreateBody{
			World: worldId,
			Def:   *(*BodyDef)(unsafe.Add(nil, def)),
		}
		b2RecWriteRet_CreateBody(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp, id)
	}
	return id
}

func b2WakeBody(tls *_Stack, world uintptr, body uintptr) (r uint8) {
	if (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeSolverSet(tls, world, (*b2Body)(unsafe.Pointer(body)).SetIndex)
		b2ValidateSolverSets(tls, world)
		return boolUint8(true1 != 0)
	}
	return boolUint8(false1 != 0)
}

func b2DestroyBody(tls *_Stack, bodyId BodyId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, chain, joint, movedBody, set, shape, world, v5, v6 uintptr
	var chainId, edgeIndex, edgeKey, jointId, shapeId, v7, v8 int32
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_DestroyBody
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, chain, chainId, edgeIndex, edgeKey, joint, jointId, movedBody, set, shape, shapeId, world, v1, v5, v6, v7, v8
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	// Record before destroying (body must still be valid)
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_DestroyBody)(unsafe.Add(nil, bp)) = b2RecArgs_DestroyBody{
			Body: bodyId,
		}
		b2RecWrite_DestroyBody(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	// Destroy the attached joints
	edgeKey = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
	for edgeKey != -int32(1) {
		jointId = edgeKey >> int32(1)
		edgeIndex = edgeKey & int32(1)
		if v1 = !!(0 <= jointId && jointId < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1850, __ccgo_ts+365, int32FromInt32(365))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		joint = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr(jointId)*72
		edgeKey = (*(*b2JointEdge)(unsafe.Add(nil, joint+20+uintptr(edgeIndex)*12))).NextKey
		// Careful because this modifies the list being traversed
		b2DestroyJointInternal(tls, world, joint)
	}
	// Destroy all contacts attached to this body.
	b2DestroyBodyContacts(tls, world, body)
	// Destroy the attached shapes and their broad-phase proxies.
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		if v1 = !!(0 <= shapeId && shapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1908, __ccgo_ts+365, int32FromInt32(379))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		shape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*280
		if (*b2Shape)(unsafe.Pointer(shape)).SensorIndex != -int32(1) {
			b2DestroySensor(tls, world, shape)
		}
		b2DestroyShapeProxy(tls, shape, world+40)
		// Return shape to free list.
		b2FreeId(tls, world+2112, shapeId)
		(*b2Shape)(unsafe.Pointer(shape)).Id = -int32(1)
		shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
	}
	// Destroy the attached chains. The associated shapes have already been destroyed above.
	chainId = (*b2Body)(unsafe.Pointer(body)).HeadChainId
	for chainId != -int32(1) {
		if v1 = !!(0 <= chainId && chainId < (*b2World)(unsafe.Pointer(world)).ChainShapes.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1966, __ccgo_ts+365, int32FromInt32(399))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		chain = (*b2World)(unsafe.Pointer(world)).ChainShapes.Data + uintptr(chainId)*32
		b2FreeChainData(tls, chain)
		// Return chain to free list.
		b2FreeId(tls, world+2136, chainId)
		(*b2ChainShape)(unsafe.Pointer(chain)).Id = -int32(1)
		chainId = (*b2ChainShape)(unsafe.Pointer(chain)).NextChainId
	}
	b2RemoveBodyFromIsland(tls, world, body)
	if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(body)).SetIndex && (*b2Body)(unsafe.Pointer(body)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+365, int32FromInt32(413))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// Remove body sim from solver set that owns it
	set = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(body)).SetIndex)*88
	movedBody = b2RemoveBodySim(tls, world, set, (*b2Body)(unsafe.Pointer(body)).LocalIndex)
	if movedBody != uintptrFromInt32(0) {
		b2RefreshBodyContactIndices(tls, world, movedBody)
	}
	// Remove body state from awake set
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
		v5 = (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Data
		v6 = set + 16 + 8
		v7 = (*b2Body)(unsafe.Pointer(body)).LocalIndex
		v8 = int32(32)
		if v1 = !!(0 <= v7 && v7 < *(*int32)(unsafe.Add(nil, v6)) && bool(__ccgo_ts+2029 != 0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2055, __ccgo_ts+2115, int32FromInt32(153))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		*(*int32)(unsafe.Add(nil, v6)) = *(*int32)(unsafe.Add(nil, v6)) - 1
		if v7 != *(*int32)(unsafe.Add(nil, v6)) {
			memcpy(tls, v5+uintptr(v7*v8), v5+uintptr(*(*int32)(unsafe.Add(nil, v6))*v8), uint64FromInt32(v8))
			_ = *(*int32)(unsafe.Add(nil, v6))
			goto _10
		}
		_ = -int32(1)
		goto _10
	_10:
	} else {
		if (*b2SolverSet)(unsafe.Pointer(set)).SetIndex >= int32(b2_firstSleepingSet) && (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count == 0 {
			// Remove solver set if it is empty
			b2DestroySolverSet(tls, world, (*b2SolverSet)(unsafe.Pointer(set)).SetIndex)
		}
	}
	// Free body and id (preserve body generation)
	b2FreeId(tls, world+1912, (*b2Body)(unsafe.Pointer(body)).Id)
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
	var body, contact, contactSim, shapeA, shapeB, world uintptr
	var contactId, contactKey, edgeIndex, index int32
	var v1, v2, v3 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _ = body, contact, contactId, contactKey, contactSim, edgeIndex, index, shapeA, shapeB, world, v1, v2, v3
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return 0
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	contactKey = (*b2Body)(unsafe.Pointer(body)).HeadContactKey
	index = 0
	for contactKey != -int32(1) && index < capacity {
		contactId = contactKey >> int32(1)
		edgeIndex = contactKey & int32(1)
		if v1 = !!(0 <= contactId && contactId < (*b2World)(unsafe.Pointer(world)).Contacts.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+301, __ccgo_ts+365, int32FromInt32(472))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		contact = (*b2World)(unsafe.Pointer(world)).Contacts.Data + uintptr(contactId)*64
		// Is contact touching?
		if (*b2Contact)(unsafe.Pointer(contact)).Flags&uint32(b2_contactTouchingFlag) != 0 {
			if v2 = !!(0 <= (*b2Contact)(unsafe.Pointer(contact)).ShapeIdA && (*b2Contact)(unsafe.Pointer(contact)).ShapeIdA < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v2 {
				b2InternalAssert(tls, __ccgo_ts+2142, __ccgo_ts+365, int32FromInt32(477))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
			shapeA = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b2Contact)(unsafe.Pointer(contact)).ShapeIdA)*280
			if v3 = !!(0 <= (*b2Contact)(unsafe.Pointer(contact)).ShapeIdB && (*b2Contact)(unsafe.Pointer(contact)).ShapeIdB < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v3 {
				b2InternalAssert(tls, __ccgo_ts+2220, __ccgo_ts+365, int32FromInt32(478))
			}
			_ = v3 || bool(int32FromInt32(0) != 0)
			shapeB = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b2Contact)(unsafe.Pointer(contact)).ShapeIdB)*280
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index)*132))).ContactId = ContactId{
				Index1:     (*b2Contact)(unsafe.Pointer(contact)).ContactId + int32(1),
				World0:     bodyId.World0,
				Generation: (*b2Contact)(unsafe.Pointer(contact)).Generation,
			}
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index)*132))).ShapeIdA = ShapeId{
				Index1:     (*b2Shape)(unsafe.Pointer(shapeA)).Id + int32(1),
				World0:     bodyId.World0,
				Generation: (*b2Shape)(unsafe.Pointer(shapeA)).Generation,
			}
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index)*132))).ShapeIdB = ShapeId{
				Index1:     (*b2Shape)(unsafe.Pointer(shapeB)).Id + int32(1),
				World0:     bodyId.World0,
				Generation: (*b2Shape)(unsafe.Pointer(shapeB)).Generation,
			}
			contactSim = b2GetContactSim(tls, world, contact)
			(*(*ContactData)(unsafe.Add(nil, contactData+uintptr(index)*132))).Manifold = (*b2ContactSim)(unsafe.Pointer(contactSim)).Manifold
			index = index + int32(1)
		}
		contactKey = (*(*b2ContactEdge)(unsafe.Add(nil, contact+uintptr(edgeIndex)*12))).NextKey
	}
	if v1 = !!(index <= capacity); !v1 {
		b2InternalAssert(tls, __ccgo_ts+2298, __ccgo_ts+365, int32FromInt32(493))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return index
}

func b2Body_ComputeBounds(tls *_Stack, bodyId BodyId) (r AABB) {
	var aabb, c, v11, v12, v33 AABB
	var body, shape, world uintptr
	var lower, upper Vec2
	var transform WorldTransform
	var v1, v13, v14, v15, v17, v18, v19, v20, v22, v23, v24, v25, v27, v3, v5, v7 float32
	var v10, v9 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, body, c, lower, shape, transform, upper, world, v1, v10, v11, v12, v13, v14, v15, v17, v18, v19, v20, v22, v23, v24, v25, v27, v3, v33, v5, v7, v9
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return AABB{}
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).HeadShapeId == -int32(1) {
		// No shapes, bracket the body origin so the box still contains the true position far away
		transform = b2GetBodyTransform(tls, world, (*b2Body)(unsafe.Pointer(body)).Id)
		v1 = float32(float64(transform.P.X))
		goto _2
	_2:
		v3 = float32(float64(transform.P.Y))
		goto _4
	_4:
		lower = Vec2{
			X: v1,
			Y: v3,
		}
		v5 = float32(float64(transform.P.X))
		goto _6
	_6:
		v7 = float32(float64(transform.P.Y))
		goto _8
	_8:
		upper = Vec2{
			X: v5,
			Y: v7,
		}
		return AABB{
			LowerBound: lower,
			UpperBound: upper,
		}
	}
	if v9 = !!(0 <= (*b2Body)(unsafe.Pointer(body)).HeadShapeId && (*b2Body)(unsafe.Pointer(body)).HeadShapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v9 {
		b2InternalAssert(tls, __ccgo_ts+2316, __ccgo_ts+365, int32FromInt32(516))
	}
	_ = v9 || bool(int32FromInt32(0) != 0)
	shape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b2Body)(unsafe.Pointer(body)).HeadShapeId)*280
	aabb = (*b2Shape)(unsafe.Pointer(shape)).Aabb
	for (*b2Shape)(unsafe.Pointer(shape)).NextShapeId != -int32(1) {
		if v10 = !!(0 <= (*b2Shape)(unsafe.Pointer(shape)).NextShapeId && (*b2Shape)(unsafe.Pointer(shape)).NextShapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v10 {
			b2InternalAssert(tls, __ccgo_ts+2394, __ccgo_ts+365, int32FromInt32(520))
		}
		_ = v10 || bool(int32FromInt32(0) != 0)
		shape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b2Shape)(unsafe.Pointer(shape)).NextShapeId)*280
		v11 = aabb
		v12 = (*b2Shape)(unsafe.Pointer(shape)).Aabb
		v1 = v11.LowerBound.X
		v3 = v12.LowerBound.X
		if v1 < v3 {
			v7 = v1
		} else {
			v7 = v3
		}
		v5 = v7
		goto _16
	_16:
		c.LowerBound.X = v5
		v13 = v11.LowerBound.Y
		v14 = v12.LowerBound.Y
		if v13 < v14 {
			v17 = v13
		} else {
			v17 = v14
		}
		v15 = v17
		goto _21
	_21:
		c.LowerBound.Y = v15
		v18 = v11.UpperBound.X
		v19 = v12.UpperBound.X
		if v18 > v19 {
			v22 = v18
		} else {
			v22 = v19
		}
		v20 = v22
		goto _26
	_26:
		c.UpperBound.X = v20
		v23 = v11.UpperBound.Y
		v24 = v12.UpperBound.Y
		if v23 > v24 {
			v27 = v23
		} else {
			v27 = v24
		}
		v25 = v27
		goto _31
	_31:
		c.UpperBound.Y = v25
		v33 = c
		goto _34
	_34:
		aabb = v33
	}
	return aabb
}

func b2UpdateBodyMassData(tls *_Stack, world2 uintptr, body2 uintptr) {
	var bodySim, bodySim1, contact, contactSim, masses, s3, s4, s5, set, set1, state, v1, v2, v5 uintptr
	var contactId, edgeIndex, edgeKey, shapeCount, shapeId, shapeId1, shapeIndex int32
	var deltaLinear, localCenter, offset, v19, v21, v22, v25, v26, v29 Vec2
	var extent, extent1 b2ShapeExtent
	var inertia, rx, ry, v10, v12, v8, v9 float32
	var massData, massData1 MassData
	var oldCenter, v40, v48 Pos
	var v3, v4 bool
	var v38 WorldTransform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodySim, bodySim1, contact, contactId, contactSim, deltaLinear, edgeIndex, edgeKey, extent, extent1, inertia, localCenter, massData, massData1, masses, offset, oldCenter, rx, ry, s3, s4, s5, set, set1, shapeCount, shapeId, shapeId1, shapeIndex, state, v1, v10, v12, v19, v2, v21, v22, v25, v26, v29, v3, v38, v4, v40, v48, v5, v8, v9
	v1 = world2
	v2 = body2
	if v3 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).SetIndex && (*b2Body)(unsafe.Pointer(v2)).SetIndex < (*b2World)(unsafe.Pointer(v1)).SolverSets.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v1)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).SetIndex)*88
	if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).LocalIndex && (*b2Body)(unsafe.Pointer(v2)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).LocalIndex)*96
	v5 = bodySim
	goto _6
_6:
	bodySim1 = v5
	// Mass is no longer dirty
	*(*uint32_t)(unsafe.Add(nil, body2+80)) &= uint32FromInt32(^int32(b2_dirtyMass))
	// Compute mass data from shapes. Each shape has its own density.
	(*b2Body)(unsafe.Pointer(body2)).Mass = float32FromFloat32(0)
	(*b2Body)(unsafe.Pointer(body2)).Inertia = float32FromFloat32(0)
	(*b2BodySim)(unsafe.Pointer(bodySim1)).InvMass = float32FromFloat32(0)
	(*b2BodySim)(unsafe.Pointer(bodySim1)).InvInertia = float32FromFloat32(0)
	(*b2BodySim)(unsafe.Pointer(bodySim1)).LocalCenter = b2Vec2_zero
	(*b2BodySim)(unsafe.Pointer(bodySim1)).MinExtent = float32(float32(float32FromFloat32(100)*float32FromFloat32(1000)) * b2_lengthUnitsPerMeter)
	(*b2BodySim)(unsafe.Pointer(bodySim1)).MaxExtent = float32FromFloat32(0)
	// Static and kinematic sims have zero mass.
	if (*b2Body)(unsafe.Pointer(body2)).Type1 != int32(b2_dynamicBody) {
		(*b2BodySim)(unsafe.Pointer(bodySim1)).Center = (*b2BodySim)(unsafe.Pointer(bodySim1)).Transform.P
		(*b2BodySim)(unsafe.Pointer(bodySim1)).Center0 = (*b2BodySim)(unsafe.Pointer(bodySim1)).Center
		// Need extents for kinematic bodies for sleeping to work correctly.
		if (*b2Body)(unsafe.Pointer(body2)).Type1 == int32(b2_kinematicBody) {
			shapeId = (*b2Body)(unsafe.Pointer(body2)).HeadShapeId
			for shapeId != -int32(1) {
				if v3 = !!(0 <= shapeId && shapeId < (*b2World)(unsafe.Pointer(world2)).Shapes.Count); !v3 {
					b2InternalAssert(tls, __ccgo_ts+1908, __ccgo_ts+365, int32FromInt32(556))
				}
				_ = v3 || bool(int32FromInt32(0) != 0)
				s3 = (*b2World)(unsafe.Pointer(world2)).Shapes.Data + uintptr(shapeId)*280
				extent = b2ComputeShapeExtent(tls, s3, b2Vec2_zero)
				v8 = (*b2BodySim)(unsafe.Pointer(bodySim1)).MinExtent
				v9 = extent.MinExtent
				if v8 < v9 {
					v12 = v8
				} else {
					v12 = v9
				}
				v10 = v12
				goto _11
			_11:
				(*b2BodySim)(unsafe.Pointer(bodySim1)).MinExtent = v10
				v8 = (*b2BodySim)(unsafe.Pointer(bodySim1)).MaxExtent
				v9 = extent.MaxExtent
				if v8 > v9 {
					v12 = v8
				} else {
					v12 = v9
				}
				v10 = v12
				goto _16
			_16:
				(*b2BodySim)(unsafe.Pointer(bodySim1)).MaxExtent = v10
				shapeId = (*b2Shape)(unsafe.Pointer(s3)).NextShapeId
			}
		}
		return
	}
	shapeCount = (*b2Body)(unsafe.Pointer(body2)).ShapeCount
	masses = b2StackAlloc(tls, world2, int32FromUint64(uint64FromInt32(shapeCount)*uint64(16)), __ccgo_ts+2474)
	// Accumulate mass over all shapes.
	localCenter = b2Vec2_zero
	shapeId1 = (*b2Body)(unsafe.Pointer(body2)).HeadShapeId
	shapeIndex = 0
	for shapeId1 != -int32(1) {
		if v3 = !!(0 <= shapeId1 && shapeId1 < (*b2World)(unsafe.Pointer(world2)).Shapes.Count); !v3 {
			b2InternalAssert(tls, __ccgo_ts+1908, __ccgo_ts+365, int32FromInt32(578))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		s4 = (*b2World)(unsafe.Pointer(world2)).Shapes.Data + uintptr(shapeId1)*280
		shapeId1 = (*b2Shape)(unsafe.Pointer(s4)).NextShapeId
		if (*b2Shape)(unsafe.Pointer(s4)).Density == float32FromFloat32(0) {
			*(*MassData)(unsafe.Add(nil, masses+uintptr(shapeIndex)*16)) = MassData{}
			shapeIndex = shapeIndex + int32(1)
			continue
		}
		massData = b2ComputeShapeMass(tls, s4)
		*(*float32)(unsafe.Add(nil, body2+52)) += massData.Mass
		v19 = localCenter
		v8 = massData.Mass
		v21 = massData.Center
		v22 = Vec2{
			X: v19.X + float32(v8*v21.X),
			Y: v19.Y + float32(v8*v21.Y),
		}
		goto _23
	_23:
		localCenter = v22
		*(*MassData)(unsafe.Add(nil, masses+uintptr(shapeIndex)*16)) = massData
		shapeIndex = shapeIndex + int32(1)
	}
	// Compute center of mass.
	if (*b2Body)(unsafe.Pointer(body2)).Mass > float32FromFloat32(0) {
		(*b2BodySim)(unsafe.Pointer(bodySim1)).InvMass = float32FromFloat32(1) / (*b2Body)(unsafe.Pointer(body2)).Mass
		v8 = (*b2BodySim)(unsafe.Pointer(bodySim1)).InvMass
		v19 = localCenter
		v21 = Vec2{
			X: float32(v8 * v19.X),
			Y: float32(v8 * v19.Y),
		}
		goto _27
	_27:
		localCenter = v21
	}
	// Second loop to accumulate the rotational inertia about the center of mass
	shapeIndex = 0
	for {
		if !(shapeIndex < shapeCount) {
			break
		}
		massData1 = *(*MassData)(unsafe.Add(nil, masses+uintptr(shapeIndex)*16))
		if massData1.Mass == float32FromFloat32(0) {
			goto _28
		}
		v19 = localCenter
		v21 = massData1.Center
		v22 = Vec2{
			X: v19.X - v21.X,
			Y: v19.Y - v21.Y,
		}
		goto _32
	_32:
		// Shift to center of mass. This is safe because it can only increase.
		offset = v22
		v25 = offset
		v26 = offset
		v8 = float32(v25.X*v26.X) + float32(v25.Y*v26.Y)
		goto _36
	_36:
		inertia = massData1.RotationalInertia + float32(massData1.Mass*v8)
		*(*float32)(unsafe.Add(nil, body2+56)) += inertia
		goto _28
	_28:
		;
		shapeIndex = shapeIndex + 1
	}
	b2StackFree(tls, world2, masses)
	masses = uintptrFromInt32(0)
	if v3 = !!((*b2Body)(unsafe.Pointer(body2)).Inertia >= float32FromFloat32(0)); !v3 {
		b2InternalAssert(tls, __ccgo_ts+2484, __ccgo_ts+365, int32FromInt32(621))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	if (*b2Body)(unsafe.Pointer(body2)).Inertia > float32FromFloat32(0) && (*b2Body)(unsafe.Pointer(body2)).Flags&uint32(b2_fixedRotation) == uint32(0) {
		(*b2BodySim)(unsafe.Pointer(bodySim1)).InvInertia = float32FromFloat32(1) / (*b2Body)(unsafe.Pointer(body2)).Inertia
	} else {
		(*b2Body)(unsafe.Pointer(body2)).Inertia = float32FromFloat32(0)
		(*b2BodySim)(unsafe.Pointer(bodySim1)).InvInertia = float32FromFloat32(0)
	}
	// Move center of mass.
	oldCenter = (*b2BodySim)(unsafe.Pointer(bodySim1)).Center
	(*b2BodySim)(unsafe.Pointer(bodySim1)).LocalCenter = localCenter
	v38 = (*b2BodySim)(unsafe.Pointer(bodySim1)).Transform
	v19 = (*b2BodySim)(unsafe.Pointer(bodySim1)).LocalCenter
	rx = float32(v38.Q.C*v19.X) - float32(v38.Q.S*v19.Y)
	ry = float32(v38.Q.S*v19.X) + float32(v38.Q.C*v19.Y)
	v40 = Pos{
		X: v38.P.X + rx,
		Y: v38.P.Y + ry,
	}
	goto _41
_41:
	(*b2BodySim)(unsafe.Pointer(bodySim1)).Center = v40
	(*b2BodySim)(unsafe.Pointer(bodySim1)).Center0 = (*b2BodySim)(unsafe.Pointer(bodySim1)).Center
	v1 = world2
	v2 = body2
	if (*b2Body)(unsafe.Pointer(v2)).SetIndex == int32(b2_awakeSet) {
		if v3 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(v1)).SolverSets.Count); !v3 {
			b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+698, int32FromInt32(254))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		set1 = (*b2World)(unsafe.Pointer(v1)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
		if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).LocalIndex && (*b2Body)(unsafe.Pointer(v2)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set1)).BodyStates.Count); !v4 {
			b2InternalAssert(tls, __ccgo_ts+790, __ccgo_ts+698, int32FromInt32(255))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		v5 = (*b2SolverSet)(unsafe.Pointer(set1)).BodyStates.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).LocalIndex)*32
		goto _46
	}
	v5 = uintptrFromInt32(0)
	goto _46
_46:
	// Update center of mass velocity
	state = v5
	if state != uintptrFromInt32(0) {
		v40 = (*b2BodySim)(unsafe.Pointer(bodySim1)).Center
		v48 = oldCenter
		v19 = Vec2{
			X: v40.X - v48.X,
			Y: v40.Y - v48.Y,
		}
		goto _51
	_51:
		v8 = (*b2BodyState)(unsafe.Pointer(state)).AngularVelocity
		v21 = v19
		v22 = Vec2{
			X: float32(-v8 * v21.Y),
			Y: float32(v8 * v21.X),
		}
		goto _55
	_55:
		deltaLinear = v22
		v25 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
		v26 = deltaLinear
		v29 = Vec2{
			X: v25.X + v26.X,
			Y: v25.Y + v26.Y,
		}
		goto _59
	_59:
		(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity = v29
	}
	// Compute body extents relative to center of mass
	shapeId1 = (*b2Body)(unsafe.Pointer(body2)).HeadShapeId
	for shapeId1 != -int32(1) {
		if v3 = !!(0 <= shapeId1 && shapeId1 < (*b2World)(unsafe.Pointer(world2)).Shapes.Count); !v3 {
			b2InternalAssert(tls, __ccgo_ts+1908, __ccgo_ts+365, int32FromInt32(651))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		s5 = (*b2World)(unsafe.Pointer(world2)).Shapes.Data + uintptr(shapeId1)*280
		extent1 = b2ComputeShapeExtent(tls, s5, localCenter)
		v8 = (*b2BodySim)(unsafe.Pointer(bodySim1)).MinExtent
		v9 = extent1.MinExtent
		if v8 < v9 {
			v12 = v8
		} else {
			v12 = v9
		}
		v10 = v12
		goto _64
	_64:
		(*b2BodySim)(unsafe.Pointer(bodySim1)).MinExtent = v10
		v8 = (*b2BodySim)(unsafe.Pointer(bodySim1)).MaxExtent
		v9 = extent1.MaxExtent
		if v8 > v9 {
			v12 = v8
		} else {
			v12 = v9
		}
		v10 = v12
		goto _69
	_69:
		(*b2BodySim)(unsafe.Pointer(bodySim1)).MaxExtent = v10
		shapeId1 = (*b2Shape)(unsafe.Pointer(s5)).NextShapeId
	}
	// When the center of mass changes, any cached contact manifold becomes invalid.
	edgeKey = (*b2Body)(unsafe.Pointer(body2)).HeadContactKey
	for edgeKey != -int32(1) {
		contactId = edgeKey >> int32(1)
		edgeIndex = edgeKey & int32(1)
		if v3 = !!(0 <= contactId && contactId < (*b2World)(unsafe.Pointer(world2)).Contacts.Count); !v3 {
			b2InternalAssert(tls, __ccgo_ts+301, __ccgo_ts+365, int32FromInt32(667))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		contact = (*b2World)(unsafe.Pointer(world2)).Contacts.Data + uintptr(contactId)*64
		contactSim = b2GetContactSim(tls, world2, contact)
		*(*uint32_t)(unsafe.Add(nil, contactSim+44)) &= uint32FromInt32(^int32(b2_simRelativeTransformValid))
		edgeKey = (*(*b2ContactEdge)(unsafe.Add(nil, contact+uintptr(edgeIndex)*12))).NextKey
	}
}

func b2Body_GetPosition(tls *_Stack, bodyId BodyId) (r Pos) {
	var body, world uintptr
	var transform WorldTransform
	_, _, _ = body, transform, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	transform = b2GetBodyTransformQuick(tls, world, body)
	return transform.P
}

func b2Body_GetRotation(tls *_Stack, bodyId BodyId) (r Rot) {
	var body, world uintptr
	var transform WorldTransform
	_, _, _ = body, transform, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	transform = b2GetBodyTransformQuick(tls, world, body)
	return transform.Q
}

func b2Body_GetTransform(tls *_Stack, bodyId BodyId) (r WorldTransform) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return b2GetBodyTransformQuick(tls, world, body)
}

func b2Body_GetLocalPoint(tls *_Stack, bodyId BodyId, worldPoint Pos) (r Vec2) {
	var body, world uintptr
	var transform, v1 WorldTransform
	var vx, vy float32
	var v2 Pos
	var v3 Vec2
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

func b2Body_GetWorldPoint(tls *_Stack, bodyId BodyId, localPoint Vec2) (r Pos) {
	var body, world uintptr
	var rx, ry float32
	var transform, v1 WorldTransform
	var v2 Vec2
	var v3 Pos
	_, _, _, _, _, _, _, _ = body, rx, ry, transform, world, v1, v2, v3
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	transform = b2GetBodyTransformQuick(tls, world, body)
	v1 = transform
	v2 = localPoint
	rx = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y)
	ry = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y)
	v3 = Pos{
		X: v1.P.X + rx,
		Y: v1.P.Y + ry,
	}
	goto _4
_4:
	return v3
}

func b2Body_GetLocalVector(tls *_Stack, bodyId BodyId, worldVector Vec2) (r Vec2) {
	var body, world uintptr
	var transform WorldTransform
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
	var transform WorldTransform
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

func b2Body_SetTransform(tls *_Stack, bodyId BodyId, position Pos, rotation Rot) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var aabb, fatAABB, v16, v17 AABB
	var body1, bodySim, bodySim1, broadPhase, set, shape, shapeFatAABB, world1, v5, v6, v9 uintptr
	var margin, rx, ry, speculativeDistance float32
	var s, v18 uint8
	var shapeId int32
	var transform, v11 WorldTransform
	var v1, v2 bool
	var v12 Vec2
	var v13 Pos
	var _ /* _a at bp+0 */ b2RecArgs_BodySetTransform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, body1, bodySim, bodySim1, broadPhase, fatAABB, margin, rx, ry, s, set, shape, shapeFatAABB, shapeId, speculativeDistance, transform, world1, v1, v11, v12, v13, v16, v17, v18, v2, v5, v6, v9
	if int32FromUint8(b2IsValidPos(tls, position)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+32, __ccgo_ts+2506, uintptr(unsafe.Pointer(&__func__3))))
		if v1 = !!(b2IsValidPos(tls, position) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2506, __ccgo_ts+365, int32FromInt32(732))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if int32FromUint8(b2IsValidRot(tls, rotation)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+32, __ccgo_ts+2531, uintptr(unsafe.Pointer(&__func__3))))
		if v1 = !!(b2IsValidRot(tls, rotation) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2531, __ccgo_ts+365, int32FromInt32(733))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if v1 = !!(b2Body_IsValid(tls, bodyId) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+387, __ccgo_ts+365, int32FromInt32(734))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world1)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+365, int32FromInt32(736))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world1)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodySetTransform)(unsafe.Add(nil, bp)) = b2RecArgs_BodySetTransform{
			Body:     bodyId,
			Position: position,
			Rotation: rotation,
		}
		b2RecWrite_BodySetTransform(tls, (*b2World)(unsafe.Pointer(world1)).Recording, bp)
	}
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	v5 = world1
	v6 = body1
	if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(v6)).SetIndex && (*b2Body)(unsafe.Pointer(v6)).SetIndex < (*b2World)(unsafe.Pointer(v5)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v5)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v6)).SetIndex)*88
	if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(v6)).LocalIndex && (*b2Body)(unsafe.Pointer(v6)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v6)).LocalIndex)*96
	v9 = bodySim
	goto _10
_10:
	bodySim1 = v9
	(*b2BodySim)(unsafe.Pointer(bodySim1)).Transform.P = position
	(*b2BodySim)(unsafe.Pointer(bodySim1)).Transform.Q = rotation
	v11 = (*b2BodySim)(unsafe.Pointer(bodySim1)).Transform
	v12 = (*b2BodySim)(unsafe.Pointer(bodySim1)).LocalCenter
	rx = float32(v11.Q.C*v12.X) - float32(v11.Q.S*v12.Y)
	ry = float32(v11.Q.S*v12.X) + float32(v11.Q.C*v12.Y)
	v13 = Pos{
		X: v11.P.X + rx,
		Y: v11.P.Y + ry,
	}
	goto _14
_14:
	(*b2BodySim)(unsafe.Pointer(bodySim1)).Center = v13
	(*b2BodySim)(unsafe.Pointer(bodySim1)).Rotation0 = (*b2BodySim)(unsafe.Pointer(bodySim1)).Transform.Q
	(*b2BodySim)(unsafe.Pointer(bodySim1)).Center0 = (*b2BodySim)(unsafe.Pointer(bodySim1)).Center
	broadPhase = world1 + 40
	transform = (*b2BodySim)(unsafe.Pointer(bodySim1)).Transform
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter))
	shapeId = (*b2Body)(unsafe.Pointer(body1)).HeadShapeId
	for shapeId != -int32(1) {
		if v1 = !!(0 <= shapeId && shapeId < (*b2World)(unsafe.Pointer(world1)).Shapes.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1908, __ccgo_ts+365, int32FromInt32(758))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		shape = (*b2World)(unsafe.Pointer(world1)).Shapes.Data + uintptr(shapeId)*280
		aabb = b2ComputeFatShapeAABB(tls, shape, transform, speculativeDistance)
		(*b2Shape)(unsafe.Pointer(shape)).Aabb = aabb
		shapeFatAABB = (*b2World)(unsafe.Pointer(world1)).FatAABBs.Data + uintptr(shapeId)*16
		v16 = *(*AABB)(unsafe.Add(nil, shapeFatAABB))
		v17 = aabb
		s = boolUint8(true1 != 0)
		s = boolUint8(s != 0 && v16.LowerBound.X <= v17.LowerBound.X)
		s = boolUint8(s != 0 && v16.LowerBound.Y <= v17.LowerBound.Y)
		s = boolUint8(s != 0 && v17.UpperBound.X <= v16.UpperBound.X)
		s = boolUint8(s != 0 && v17.UpperBound.Y <= v16.UpperBound.Y)
		v18 = s
		goto _19
	_19:
		if int32FromUint8(v18) == false1 {
			margin = (*b2Shape)(unsafe.Pointer(shape)).AabbMargin
			fatAABB.LowerBound.X = aabb.LowerBound.X - margin
			fatAABB.LowerBound.Y = aabb.LowerBound.Y - margin
			fatAABB.UpperBound.X = aabb.UpperBound.X + margin
			fatAABB.UpperBound.Y = aabb.UpperBound.Y + margin
			*(*AABB)(unsafe.Add(nil, shapeFatAABB)) = fatAABB
			// They body could be disabled
			if (*b2Shape)(unsafe.Pointer(shape)).ProxyKey != -int32(1) {
				b2BroadPhase_MoveProxy(tls, broadPhase, (*b2Shape)(unsafe.Pointer(shape)).ProxyKey, fatAABB)
			}
		}
		shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
	}
}

func b2Body_GetLinearVelocity(tls *_Stack, bodyId BodyId) (r Vec2) {
	var body1, set, state, world1, v1, v2, v4 uintptr
	var v3, v6 bool
	_, _, _, _, _, _, _, _, _ = body1, set, state, world1, v1, v2, v3, v4, v6
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	v1 = world1
	v2 = body1
	if (*b2Body)(unsafe.Pointer(v2)).SetIndex == int32(b2_awakeSet) {
		if v3 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(v1)).SolverSets.Count); !v3 {
			b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+698, int32FromInt32(254))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		set = (*b2World)(unsafe.Pointer(v1)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
		if v6 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).LocalIndex && (*b2Body)(unsafe.Pointer(v2)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Count); !v6 {
			b2InternalAssert(tls, __ccgo_ts+790, __ccgo_ts+698, int32FromInt32(255))
		}
		_ = v6 || bool(int32FromInt32(0) != 0)
		v4 = (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).LocalIndex)*32
		goto _5
	}
	v4 = uintptrFromInt32(0)
	goto _5
_5:
	state = v4
	if state != uintptrFromInt32(0) {
		return (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
	}
	return b2Vec2_zero
}

func b2Body_GetAngularVelocity(tls *_Stack, bodyId BodyId) (r float32) {
	var body1, set, state, world1, v1, v2, v4 uintptr
	var v3, v6 bool
	_, _, _, _, _, _, _, _, _ = body1, set, state, world1, v1, v2, v3, v4, v6
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	v1 = world1
	v2 = body1
	if (*b2Body)(unsafe.Pointer(v2)).SetIndex == int32(b2_awakeSet) {
		if v3 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(v1)).SolverSets.Count); !v3 {
			b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+698, int32FromInt32(254))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		set = (*b2World)(unsafe.Pointer(v1)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
		if v6 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).LocalIndex && (*b2Body)(unsafe.Pointer(v2)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Count); !v6 {
			b2InternalAssert(tls, __ccgo_ts+790, __ccgo_ts+698, int32FromInt32(255))
		}
		_ = v6 || bool(int32FromInt32(0) != 0)
		v4 = (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).LocalIndex)*32
		goto _5
	}
	v4 = uintptrFromInt32(0)
	goto _5
_5:
	state = v4
	if state != uintptrFromInt32(0) {
		return (*b2BodyState)(unsafe.Pointer(state)).AngularVelocity
	}
	return float32(0)
}

func b2Body_SetLinearVelocity(tls *_Stack, bodyId BodyId, linearVelocity Vec2) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var body1, set, state, world1, v5, v6, v8 uintptr
	var v1, v7 bool
	var v2 Vec2
	var v3 float32
	var _ /* _a at bp+0 */ b2RecArgs_BodySetLinearVelocity
	_, _, _, _, _, _, _, _, _, _, _ = body1, set, state, world1, v1, v2, v3, v5, v6, v7, v8
	if int32FromUint8(b2IsValidVec2(tls, linearVelocity)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+2556, uintptr(unsafe.Pointer(&__func__2))))
		if v1 = !!(b2IsValidVec2(tls, linearVelocity) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2556, __ccgo_ts+365, int32FromInt32(810))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b2World)(unsafe.Pointer(world1)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodySetLinearVelocity)(unsafe.Add(nil, bp)) = b2RecArgs_BodySetLinearVelocity{
			Body: bodyId,
			V:    linearVelocity,
		}
		b2RecWrite_BodySetLinearVelocity(tls, (*b2World)(unsafe.Pointer(world1)).Recording, bp)
	}
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	if (*b2Body)(unsafe.Pointer(body1)).Type1 == int32(b2_staticBody) {
		return
	}
	v2 = linearVelocity
	v3 = float32(v2.X*v2.X) + float32(v2.Y*v2.Y)
	goto _4
_4:
	if v3 > float32FromFloat32(0) {
		b2WakeBody(tls, world1, body1)
	}
	v5 = world1
	v6 = body1
	if (*b2Body)(unsafe.Pointer(v6)).SetIndex == int32(b2_awakeSet) {
		if v1 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(v5)).SolverSets.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+698, int32FromInt32(254))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		set = (*b2World)(unsafe.Pointer(v5)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
		if v7 = !!(0 <= (*b2Body)(unsafe.Pointer(v6)).LocalIndex && (*b2Body)(unsafe.Pointer(v6)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Count); !v7 {
			b2InternalAssert(tls, __ccgo_ts+790, __ccgo_ts+698, int32FromInt32(255))
		}
		_ = v7 || bool(int32FromInt32(0) != 0)
		v8 = (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Data + uintptr((*b2Body)(unsafe.Pointer(v6)).LocalIndex)*32
		goto _9
	}
	v8 = uintptrFromInt32(0)
	goto _9
_9:
	state = v8
	if state == uintptrFromInt32(0) {
		return
	}
	(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity = linearVelocity
}

func b2Body_SetAngularVelocity(tls *_Stack, bodyId BodyId, angularVelocity float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var body1, set, state, world1, v2, v3, v5 uintptr
	var v1, v4 bool
	var _ /* _a at bp+0 */ b2RecArgs_BodySetAngularVelocity
	_, _, _, _, _, _, _, _, _ = body1, set, state, world1, v1, v2, v3, v4, v5
	if int32FromUint8(b2IsValidFloat(tls, angularVelocity)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+2588, uintptr(unsafe.Pointer(&__func__11))))
		if v1 = !!(b2IsValidFloat(tls, angularVelocity) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2588, __ccgo_ts+365, int32FromInt32(839))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b2World)(unsafe.Pointer(world1)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodySetAngularVelocity)(unsafe.Add(nil, bp)) = b2RecArgs_BodySetAngularVelocity{
			Body: bodyId,
			W:    angularVelocity,
		}
		b2RecWrite_BodySetAngularVelocity(tls, (*b2World)(unsafe.Pointer(world1)).Recording, bp)
	}
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	if (*b2Body)(unsafe.Pointer(body1)).Type1 == int32(b2_staticBody) || (*b2Body)(unsafe.Pointer(body1)).Flags&uint32(b2_lockAngularZ) != 0 {
		return
	}
	if angularVelocity != float32FromFloat32(0) {
		b2WakeBody(tls, world1, body1)
	}
	v2 = world1
	v3 = body1
	if (*b2Body)(unsafe.Pointer(v3)).SetIndex == int32(b2_awakeSet) {
		if v1 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(v2)).SolverSets.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+698, int32FromInt32(254))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		set = (*b2World)(unsafe.Pointer(v2)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
		if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v3)).LocalIndex && (*b2Body)(unsafe.Pointer(v3)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Count); !v4 {
			b2InternalAssert(tls, __ccgo_ts+790, __ccgo_ts+698, int32FromInt32(255))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		v5 = (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Data + uintptr((*b2Body)(unsafe.Pointer(v3)).LocalIndex)*32
		goto _6
	}
	v5 = uintptrFromInt32(0)
	goto _6
_6:
	state = v5
	if state == uintptrFromInt32(0) {
		return
	}
	(*b2BodyState)(unsafe.Pointer(state)).AngularVelocity = angularVelocity
}

func b2Body_SetTargetTransform(tls *_Stack, bodyId BodyId, target WorldTransform, timeStep float32, wake uint8) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var angularVelocity, c, deltaAngle, invTimeStep, maxVelocity, rx, ry, s1, v17, v23, v26, v28, v29, v31 float32
	var body2, bodySim, set, set1, sim, state, world2, v3, v4, v7 uintptr
	var delta, linearVelocity, v10, v15, v18, v19, v25 Vec2
	var q1, q2, v21, v22 Rot
	var v1, v2 bool
	var v11, v13, v14 Pos
	var v9 WorldTransform
	var _ /* _a at bp+0 */ b2RecArgs_BodySetTargetTransform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angularVelocity, body2, bodySim, c, delta, deltaAngle, invTimeStep, linearVelocity, maxVelocity, q1, q2, rx, ry, s1, set, set1, sim, state, world2, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v28, v29, v3, v31, v4, v7, v9
	if int32FromUint8(b2IsValidWorldTransform(tls, target)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+40, __ccgo_ts+2622, uintptr(unsafe.Pointer(&__func__8))))
		if v1 = !!(b2IsValidWorldTransform(tls, target) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2622, __ccgo_ts+365, int32FromInt32(866))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if int32FromUint8(b2IsValidFloat(tls, timeStep)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+40, __ccgo_ts+2656, uintptr(unsafe.Pointer(&__func__8))))
		if v1 = !!(b2IsValidFloat(tls, timeStep) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2656, __ccgo_ts+365, int32FromInt32(867))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world2 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b2World)(unsafe.Pointer(world2)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodySetTargetTransform)(unsafe.Add(nil, bp)) = b2RecArgs_BodySetTargetTransform{
			Body:     bodyId,
			Target:   target,
			TimeStep: timeStep,
			Wake:     wake,
		}
		b2RecWrite_BodySetTargetTransform(tls, (*b2World)(unsafe.Pointer(world2)).Recording, bp)
	}
	body2 = b2GetBodyFullId(tls, world2, bodyId)
	if (*b2Body)(unsafe.Pointer(body2)).SetIndex == int32(b2_disabledSet) {
		return
	}
	if (*b2Body)(unsafe.Pointer(body2)).Type1 == int32(b2_staticBody) || timeStep <= float32FromFloat32(0) {
		return
	}
	if (*b2Body)(unsafe.Pointer(body2)).SetIndex != int32(b2_awakeSet) && int32FromUint8(wake) == false1 {
		return
	}
	v3 = world2
	v4 = body2
	if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(v4)).SetIndex && (*b2Body)(unsafe.Pointer(v4)).SetIndex < (*b2World)(unsafe.Pointer(v3)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v3)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v4)).SetIndex)*88
	if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(v4)).LocalIndex && (*b2Body)(unsafe.Pointer(v4)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v4)).LocalIndex)*96
	v7 = bodySim
	goto _8
_8:
	sim = v7
	v9 = target
	v10 = (*b2BodySim)(unsafe.Pointer(sim)).LocalCenter
	rx = float32(v9.Q.C*v10.X) - float32(v9.Q.S*v10.Y)
	ry = float32(v9.Q.S*v10.X) + float32(v9.Q.C*v10.Y)
	v11 = Pos{
		X: v9.P.X + rx,
		Y: v9.P.Y + ry,
	}
	goto _12
_12:
	v13 = v11
	v14 = (*b2BodySim)(unsafe.Pointer(sim)).Center
	v15 = Vec2{
		X: v13.X - v14.X,
		Y: v13.Y - v14.Y,
	}
	goto _16
_16:
	// Compute linear velocity. The center difference is taken in world precision then demoted
	delta = v15
	invTimeStep = float32FromFloat32(1) / timeStep
	v17 = invTimeStep
	v18 = delta
	v19 = Vec2{
		X: float32(v17 * v18.X),
		Y: float32(v17 * v18.Y),
	}
	goto _20
_20:
	linearVelocity = v19
	// Compute angular velocity
	q1 = (*b2BodySim)(unsafe.Pointer(sim)).Transform.Q
	q2 = target.Q
	v21 = q1
	v22 = q2
	s1 = float32(v21.C*v22.S) - float32(v21.S*v22.C)
	c = float32(v21.C*v22.C) + float32(v21.S*v22.S)
	v23 = b2Atan2(tls, s1, c)
	goto _24
_24:
	deltaAngle = v23
	angularVelocity = float32(invTimeStep * deltaAngle)
	// Early out if the body is asleep already and the desired movement is small
	if (*b2Body)(unsafe.Pointer(body2)).SetIndex != int32(b2_awakeSet) {
		v25 = linearVelocity
		v26 = sqrtf(tls, float32(v25.X*v25.X)+float32(v25.Y*v25.Y))
		goto _27
	_27:
		v28 = angularVelocity
		if v28 < float32FromInt32(0) {
			v31 = -v28
		} else {
			v31 = v28
		}
		v29 = v31
		goto _30
	_30:
		maxVelocity = v26 + float32(v29*(*b2BodySim)(unsafe.Pointer(sim)).MaxExtent)
		// Return if velocity would be sleepy
		if maxVelocity < (*b2Body)(unsafe.Pointer(body2)).SleepThreshold {
			return
		}
		// Must wake for state to exist
		b2WakeBody(tls, world2, body2)
	}
	if v1 = !!((*b2Body)(unsafe.Pointer(body2)).SetIndex == int32(b2_awakeSet)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+2683, __ccgo_ts+365, int32FromInt32(917))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	v3 = world2
	v4 = body2
	if (*b2Body)(unsafe.Pointer(v4)).SetIndex == int32(b2_awakeSet) {
		if v1 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(v3)).SolverSets.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+698, int32FromInt32(254))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		set1 = (*b2World)(unsafe.Pointer(v3)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
		if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(v4)).LocalIndex && (*b2Body)(unsafe.Pointer(v4)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set1)).BodyStates.Count); !v2 {
			b2InternalAssert(tls, __ccgo_ts+790, __ccgo_ts+698, int32FromInt32(255))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		v7 = (*b2SolverSet)(unsafe.Pointer(set1)).BodyStates.Data + uintptr((*b2Body)(unsafe.Pointer(v4)).LocalIndex)*32
		goto _37
	}
	v7 = uintptrFromInt32(0)
	goto _37
_37:
	state = v7
	(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity = linearVelocity
	(*b2BodyState)(unsafe.Pointer(state)).AngularVelocity = angularVelocity
}

func b2Body_GetLocalPointVelocity(tls *_Stack, bodyId BodyId, localPoint Vec2) (r1 Vec2) {
	var body1, bodySim, set, set1, state, world1, v1, v21, v4 uintptr
	var r, v2, v10, v11, v14, v15, v18, v19, v211, v22, v23, v9 Vec2
	var v13 Rot
	var v17 float32
	var v3, v6 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body1, bodySim, r, set, set1, state, v2, world1, v1, v10, v11, v13, v14, v15, v17, v18, v19, v21, v211, v22, v23, v3, v4, v6, v9
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	v1 = world1
	v21 = body1
	if (*b2Body)(unsafe.Pointer(v21)).SetIndex == int32(b2_awakeSet) {
		if v3 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(v1)).SolverSets.Count); !v3 {
			b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+698, int32FromInt32(254))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		set = (*b2World)(unsafe.Pointer(v1)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
		if v6 = !!(0 <= (*b2Body)(unsafe.Pointer(v21)).LocalIndex && (*b2Body)(unsafe.Pointer(v21)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Count); !v6 {
			b2InternalAssert(tls, __ccgo_ts+790, __ccgo_ts+698, int32FromInt32(255))
		}
		_ = v6 || bool(int32FromInt32(0) != 0)
		v4 = (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Data + uintptr((*b2Body)(unsafe.Pointer(v21)).LocalIndex)*32
		goto _5
	}
	v4 = uintptrFromInt32(0)
	goto _5
_5:
	state = v4
	if state == uintptrFromInt32(0) {
		return b2Vec2_zero
	}
	if v3 = !!(0 <= (*b2Body)(unsafe.Pointer(body1)).SetIndex && (*b2Body)(unsafe.Pointer(body1)).SetIndex < (*b2World)(unsafe.Pointer(world1)).SolverSets.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+365, int32FromInt32(934))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	set1 = (*b2World)(unsafe.Pointer(world1)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(body1)).SetIndex)*88
	if v6 = !!(0 <= (*b2Body)(unsafe.Pointer(body1)).LocalIndex && (*b2Body)(unsafe.Pointer(body1)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set1)).BodySims.Count); !v6 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+365, int32FromInt32(935))
	}
	_ = v6 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set1)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(body1)).LocalIndex)*96
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

func b2Body_GetWorldPointVelocity(tls *_Stack, bodyId BodyId, worldPoint Pos) (r1 Vec2) {
	var body1, bodySim, set, set1, state, world1, v11, v2, v4 uintptr
	var r, v1, v111, v14, v15, v17, v18, v19 Vec2
	var v10, v9 Pos
	var v13 float32
	var v3, v6 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body1, bodySim, r, set, set1, state, v1, world1, v11, v10, v111, v13, v14, v15, v17, v18, v19, v2, v3, v4, v6, v9
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	v11 = world1
	v2 = body1
	if (*b2Body)(unsafe.Pointer(v2)).SetIndex == int32(b2_awakeSet) {
		if v3 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(v11)).SolverSets.Count); !v3 {
			b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+698, int32FromInt32(254))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		set = (*b2World)(unsafe.Pointer(v11)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
		if v6 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).LocalIndex && (*b2Body)(unsafe.Pointer(v2)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Count); !v6 {
			b2InternalAssert(tls, __ccgo_ts+790, __ccgo_ts+698, int32FromInt32(255))
		}
		_ = v6 || bool(int32FromInt32(0) != 0)
		v4 = (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).LocalIndex)*32
		goto _5
	}
	v4 = uintptrFromInt32(0)
	goto _5
_5:
	state = v4
	if state == uintptrFromInt32(0) {
		return b2Vec2_zero
	}
	if v3 = !!(0 <= (*b2Body)(unsafe.Pointer(body1)).SetIndex && (*b2Body)(unsafe.Pointer(body1)).SetIndex < (*b2World)(unsafe.Pointer(world1)).SolverSets.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+365, int32FromInt32(952))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	set1 = (*b2World)(unsafe.Pointer(world1)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(body1)).SetIndex)*88
	if v6 = !!(0 <= (*b2Body)(unsafe.Pointer(body1)).LocalIndex && (*b2Body)(unsafe.Pointer(body1)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set1)).BodySims.Count); !v6 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+365, int32FromInt32(953))
	}
	_ = v6 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set1)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(body1)).LocalIndex)*96
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

func b2Body_ApplyForce(tls *_Stack, bodyId BodyId, force Vec2, point Pos, wake uint8) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var body1, bodySim, bodySim1, set, world1, v3, v4, v7 uintptr
	var v1, v2 bool
	var v10, v11, v9 Vec2
	var v13, v14 Pos
	var v19 float32
	var _ /* _a at bp+0 */ b2RecArgs_BodyApplyForce
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body1, bodySim, bodySim1, set, world1, v1, v10, v11, v13, v14, v19, v2, v3, v4, v7, v9
	if int32FromUint8(b2IsValidVec2(tls, force)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+40, __ccgo_ts+2713, uintptr(unsafe.Pointer(&__func__4))))
		if v1 = !!(b2IsValidVec2(tls, force) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2713, __ccgo_ts+365, int32FromInt32(962))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if int32FromUint8(b2IsValidPos(tls, point)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+40, __ccgo_ts+2736, uintptr(unsafe.Pointer(&__func__4))))
		if v1 = !!(b2IsValidPos(tls, point) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2736, __ccgo_ts+365, int32FromInt32(963))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b2World)(unsafe.Pointer(world1)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodyApplyForce)(unsafe.Add(nil, bp)) = b2RecArgs_BodyApplyForce{
			Body:  bodyId,
			Force: force,
			Point: point,
			Wake:  wake,
		}
		b2RecWrite_BodyApplyForce(tls, (*b2World)(unsafe.Pointer(world1)).Recording, bp)
	}
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	if (*b2Body)(unsafe.Pointer(body1)).Type1 != int32(b2_dynamicBody) || (*b2Body)(unsafe.Pointer(body1)).SetIndex == int32(b2_disabledSet) {
		return
	}
	if wake != 0 && (*b2Body)(unsafe.Pointer(body1)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeBody(tls, world1, body1)
	}
	if (*b2Body)(unsafe.Pointer(body1)).SetIndex == int32(b2_awakeSet) {
		v3 = world1
		v4 = body1
		if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(v4)).SetIndex && (*b2Body)(unsafe.Pointer(v4)).SetIndex < (*b2World)(unsafe.Pointer(v3)).SolverSets.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		set = (*b2World)(unsafe.Pointer(v3)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v4)).SetIndex)*88
		if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(v4)).LocalIndex && (*b2Body)(unsafe.Pointer(v4)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v2 {
			b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v4)).LocalIndex)*96
		v7 = bodySim
		goto _8
	_8:
		bodySim1 = v7
		v9 = (*b2BodySim)(unsafe.Pointer(bodySim1)).Force
		v10 = force
		v11 = Vec2{
			X: v9.X + v10.X,
			Y: v9.Y + v10.Y,
		}
		goto _12
	_12:
		(*b2BodySim)(unsafe.Pointer(bodySim1)).Force = v11
		v13 = point
		v14 = (*b2BodySim)(unsafe.Pointer(bodySim1)).Center
		v9 = Vec2{
			X: v13.X - v14.X,
			Y: v13.Y - v14.Y,
		}
		goto _16
	_16:
		v10 = v9
		v11 = force
		v19 = float32(v10.X*v11.Y) - float32(v10.Y*v11.X)
		goto _20
	_20:
		*(*float32)(unsafe.Add(nil, bodySim1+56)) += v19
	}
}

func b2Body_ApplyForceToCenter(tls *_Stack, bodyId BodyId, force Vec2, wake uint8) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var body1, bodySim, bodySim1, set, world1, v2, v3, v6 uintptr
	var v1, v4 bool
	var v10, v8, v9 Vec2
	var _ /* _a at bp+0 */ b2RecArgs_BodyApplyForceToCenter
	_, _, _, _, _, _, _, _, _, _, _, _, _ = body1, bodySim, bodySim1, set, world1, v1, v10, v2, v3, v4, v6, v8, v9
	if int32FromUint8(b2IsValidVec2(tls, force)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+32, __ccgo_ts+2713, uintptr(unsafe.Pointer(&__func__21))))
		if v1 = !!(b2IsValidVec2(tls, force) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2713, __ccgo_ts+365, int32FromInt32(989))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b2World)(unsafe.Pointer(world1)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodyApplyForceToCenter)(unsafe.Add(nil, bp)) = b2RecArgs_BodyApplyForceToCenter{
			Body:  bodyId,
			Force: force,
			Wake:  wake,
		}
		b2RecWrite_BodyApplyForceToCenter(tls, (*b2World)(unsafe.Pointer(world1)).Recording, bp)
	}
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	if (*b2Body)(unsafe.Pointer(body1)).Type1 != int32(b2_dynamicBody) || (*b2Body)(unsafe.Pointer(body1)).SetIndex == int32(b2_disabledSet) {
		return
	}
	if wake != 0 && (*b2Body)(unsafe.Pointer(body1)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeBody(tls, world1, body1)
	}
	if (*b2Body)(unsafe.Pointer(body1)).SetIndex == int32(b2_awakeSet) {
		v2 = world1
		v3 = body1
		if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(v3)).SetIndex && (*b2Body)(unsafe.Pointer(v3)).SetIndex < (*b2World)(unsafe.Pointer(v2)).SolverSets.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		set = (*b2World)(unsafe.Pointer(v2)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v3)).SetIndex)*88
		if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v3)).LocalIndex && (*b2Body)(unsafe.Pointer(v3)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v4 {
			b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v3)).LocalIndex)*96
		v6 = bodySim
		goto _7
	_7:
		bodySim1 = v6
		v8 = (*b2BodySim)(unsafe.Pointer(bodySim1)).Force
		v9 = force
		v10 = Vec2{
			X: v8.X + v9.X,
			Y: v8.Y + v9.Y,
		}
		goto _11
	_11:
		(*b2BodySim)(unsafe.Pointer(bodySim1)).Force = v10
	}
}

func b2Body_ApplyTorque(tls *_Stack, bodyId BodyId, torque float32, wake uint8) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var body1, bodySim, bodySim1, set, world1, v2, v3, v6 uintptr
	var v1, v4 bool
	var _ /* _a at bp+0 */ b2RecArgs_BodyApplyTorque
	_, _, _, _, _, _, _, _, _, _ = body1, bodySim, bodySim1, set, world1, v1, v2, v3, v4, v6
	if int32FromUint8(b2IsValidFloat(tls, torque)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+2758, uintptr(unsafe.Pointer(&__func__12))))
		if v1 = !!(b2IsValidFloat(tls, torque) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2758, __ccgo_ts+365, int32FromInt32(1014))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b2World)(unsafe.Pointer(world1)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodyApplyTorque)(unsafe.Add(nil, bp)) = b2RecArgs_BodyApplyTorque{
			Body:   bodyId,
			Torque: torque,
			Wake:   wake,
		}
		b2RecWrite_BodyApplyTorque(tls, (*b2World)(unsafe.Pointer(world1)).Recording, bp)
	}
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	if (*b2Body)(unsafe.Pointer(body1)).Type1 != int32(b2_dynamicBody) || (*b2Body)(unsafe.Pointer(body1)).SetIndex == int32(b2_disabledSet) {
		return
	}
	if wake != 0 && (*b2Body)(unsafe.Pointer(body1)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeBody(tls, world1, body1)
	}
	if (*b2Body)(unsafe.Pointer(body1)).SetIndex == int32(b2_awakeSet) {
		v2 = world1
		v3 = body1
		if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(v3)).SetIndex && (*b2Body)(unsafe.Pointer(v3)).SetIndex < (*b2World)(unsafe.Pointer(v2)).SolverSets.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		set = (*b2World)(unsafe.Pointer(v2)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v3)).SetIndex)*88
		if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v3)).LocalIndex && (*b2Body)(unsafe.Pointer(v3)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v4 {
			b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v3)).LocalIndex)*96
		v6 = bodySim
		goto _7
	_7:
		bodySim1 = v6
		*(*float32)(unsafe.Add(nil, bodySim1+56)) += torque
	}
}

func b2Body_ClearForces(tls *_Stack, bodyId BodyId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body1, bodySim, bodySim1, set, world1, v1, v2, v5 uintptr
	var v3, v4 bool
	var _ /* _a at bp+0 */ b2RecArgs_BodyClearForces
	_, _, _, _, _, _, _, _, _, _ = body1, bodySim, bodySim1, set, world1, v1, v2, v3, v4, v5
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b2World)(unsafe.Pointer(world1)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodyClearForces)(unsafe.Add(nil, bp)) = b2RecArgs_BodyClearForces{
			Body: bodyId,
		}
		b2RecWrite_BodyClearForces(tls, (*b2World)(unsafe.Pointer(world1)).Recording, bp)
	}
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	v1 = world1
	v2 = body1
	if v3 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).SetIndex && (*b2Body)(unsafe.Pointer(v2)).SetIndex < (*b2World)(unsafe.Pointer(v1)).SolverSets.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v1)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).SetIndex)*88
	if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).LocalIndex && (*b2Body)(unsafe.Pointer(v2)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).LocalIndex)*96
	v5 = bodySim
	goto _6
_6:
	bodySim1 = v5
	(*b2BodySim)(unsafe.Pointer(bodySim1)).Force = b2Vec2_zero
	(*b2BodySim)(unsafe.Pointer(bodySim1)).Torque = float32FromFloat32(0)
}

func b2Body_ApplyLinearImpulse(tls *_Stack, bodyId BodyId, impulse Vec2, point Pos, wake uint8) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var body, bodySim, set, state, world uintptr
	var localIndex int32
	var v1, v2, v3 bool
	var v11, v12 Pos
	var v6, v8, v9 Vec2
	var v7 float32
	var _ /* _a at bp+0 */ b2RecArgs_BodyApplyLinearImpulse
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodySim, localIndex, set, state, world, v1, v11, v12, v2, v3, v6, v7, v8, v9
	if int32FromUint8(b2IsValidVec2(tls, impulse)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+40, __ccgo_ts+2783, uintptr(unsafe.Pointer(&__func__31))))
		if v1 = !!(b2IsValidVec2(tls, impulse) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2783, __ccgo_ts+365, int32FromInt32(1049))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if int32FromUint8(b2IsValidPos(tls, point)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+40, __ccgo_ts+2736, uintptr(unsafe.Pointer(&__func__31))))
		if v1 = !!(b2IsValidPos(tls, point) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2736, __ccgo_ts+365, int32FromInt32(1050))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodyApplyLinearImpulse)(unsafe.Add(nil, bp)) = b2RecArgs_BodyApplyLinearImpulse{
			Body:    bodyId,
			Impulse: impulse,
			Point:   point,
			Wake:    wake,
		}
		b2RecWrite_BodyApplyLinearImpulse(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).Type1 != int32(b2_dynamicBody) || (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_disabledSet) {
		return
	}
	if wake != 0 && (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeBody(tls, world, body)
	}
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
		localIndex = (*b2Body)(unsafe.Pointer(body)).LocalIndex
		if v1 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+365, int32FromInt32(1069))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		set = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
		if v2 = !!(0 <= localIndex && localIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Count); !v2 {
			b2InternalAssert(tls, __ccgo_ts+2808, __ccgo_ts+365, int32FromInt32(1070))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		state = (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Data + uintptr(localIndex)*32
		if v3 = !!(0 <= localIndex && localIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v3 {
			b2InternalAssert(tls, __ccgo_ts+2874, __ccgo_ts+365, int32FromInt32(1071))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr(localIndex)*96
		v6 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
		v7 = (*b2BodySim)(unsafe.Pointer(bodySim)).InvMass
		v8 = impulse
		v9 = Vec2{
			X: v6.X + float32(v7*v8.X),
			Y: v6.Y + float32(v7*v8.Y),
		}
		goto _10
	_10:
		(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity = v9
		v11 = point
		v12 = (*b2BodySim)(unsafe.Pointer(bodySim)).Center
		v6 = Vec2{
			X: v11.X - v12.X,
			Y: v11.Y - v12.Y,
		}
		goto _14
	_14:
		v8 = v6
		v9 = impulse
		v7 = float32(v8.X*v9.Y) - float32(v8.Y*v9.X)
		goto _18
	_18:
		*(*float32)(unsafe.Add(nil, state+8)) += float32((*b2BodySim)(unsafe.Pointer(bodySim)).InvInertia * v7)
		b2LimitVelocity(tls, state, (*b2World)(unsafe.Pointer(world)).MaxLinearSpeed)
	}
}

func b2Body_ApplyLinearImpulseToCenter(tls *_Stack, bodyId BodyId, impulse Vec2, wake uint8) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var body, bodySim, set, state, world uintptr
	var localIndex int32
	var v1, v2, v3 bool
	var v5, v7, v8 Vec2
	var v6 float32
	var _ /* _a at bp+0 */ b2RecArgs_BodyApplyLinearImpulseToCenter
	_, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodySim, localIndex, set, state, world, v1, v2, v3, v5, v6, v7, v8
	if int32FromUint8(b2IsValidVec2(tls, impulse)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+32, __ccgo_ts+2783, uintptr(unsafe.Pointer(&__func__13))))
		if v1 = !!(b2IsValidVec2(tls, impulse) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2783, __ccgo_ts+365, int32FromInt32(1081))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodyApplyLinearImpulseToCenter)(unsafe.Add(nil, bp)) = b2RecArgs_BodyApplyLinearImpulseToCenter{
			Body:    bodyId,
			Impulse: impulse,
			Wake:    wake,
		}
		b2RecWrite_BodyApplyLinearImpulseToCenter(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).Type1 != int32(b2_dynamicBody) || (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_disabledSet) {
		return
	}
	if wake != 0 && (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeBody(tls, world, body)
	}
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
		localIndex = (*b2Body)(unsafe.Pointer(body)).LocalIndex
		if v1 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+365, int32FromInt32(1100))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		set = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
		if v2 = !!(0 <= localIndex && localIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Count); !v2 {
			b2InternalAssert(tls, __ccgo_ts+2808, __ccgo_ts+365, int32FromInt32(1101))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		state = (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Data + uintptr(localIndex)*32
		if v3 = !!(0 <= localIndex && localIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v3 {
			b2InternalAssert(tls, __ccgo_ts+2874, __ccgo_ts+365, int32FromInt32(1102))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr(localIndex)*96
		v5 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
		v6 = (*b2BodySim)(unsafe.Pointer(bodySim)).InvMass
		v7 = impulse
		v8 = Vec2{
			X: v5.X + float32(v6*v7.X),
			Y: v5.Y + float32(v6*v7.Y),
		}
		goto _9
	_9:
		(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity = v8
		b2LimitVelocity(tls, state, (*b2World)(unsafe.Pointer(world)).MaxLinearSpeed)
	}
}

func b2Body_ApplyAngularImpulse(tls *_Stack, bodyId BodyId, impulse float32, wake uint8) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var body, bodySim, set, state, world uintptr
	var localIndex int32
	var v1, v2, v3 bool
	var _ /* _a at bp+0 */ b2RecArgs_BodyApplyAngularImpulse
	_, _, _, _, _, _, _, _, _ = body, bodySim, localIndex, set, state, world, v1, v2, v3
	if int32FromUint8(b2IsValidFloat(tls, impulse)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+2938, uintptr(unsafe.Pointer(&__func__))))
		if v1 = !!(b2IsValidFloat(tls, impulse) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2938, __ccgo_ts+365, int32FromInt32(1111))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if v1 = !!(b2Body_IsValid(tls, bodyId) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+387, __ccgo_ts+365, int32FromInt32(1112))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodyApplyAngularImpulse)(unsafe.Add(nil, bp)) = b2RecArgs_BodyApplyAngularImpulse{
			Body:    bodyId,
			Impulse: impulse,
			Wake:    wake,
		}
		b2RecWrite_BodyApplyAngularImpulse(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).Type1 != int32(b2_dynamicBody) || (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_disabledSet) {
		return
	}
	if wake != 0 && (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		// this will not invalidate body pointer
		b2WakeBody(tls, world, body)
	}
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
		localIndex = (*b2Body)(unsafe.Pointer(body)).LocalIndex
		if v1 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+365, int32FromInt32(1131))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		set = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
		if v2 = !!(0 <= localIndex && localIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Count); !v2 {
			b2InternalAssert(tls, __ccgo_ts+2808, __ccgo_ts+365, int32FromInt32(1132))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		state = (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Data + uintptr(localIndex)*32
		if v3 = !!(0 <= localIndex && localIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v3 {
			b2InternalAssert(tls, __ccgo_ts+2874, __ccgo_ts+365, int32FromInt32(1133))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr(localIndex)*96
		*(*float32)(unsafe.Add(nil, state+8)) += float32((*b2BodySim)(unsafe.Pointer(bodySim)).InvInertia * impulse)
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
//	// This should follow similar steps as you would get destroying and recreating the body, shapes, and joints.
//	// Contacts are difficult to preserve because the broad-phase pairs change, so I just destroy them.
//	// todo with a bit more effort I could support an option to let the body sleep
//	//
//	// Revised steps:
//	// 1 Skip disabled bodies
//	// 2 Destroy all contacts on the body
//	// 3 Wake the body
//	// 4 For all joints attached to the body
//	//  - wake attached bodies
//	//  - remove from island
//	//  - move to static set temporarily
//	// 5 Change the body type and transfer the body
//	// 6 If the body was static
//	//   - create an island for the body
//	//   Else if the body is becoming static
//	//   - remove it from the island
//	// 7 For all joints
//	//  - if either body is non-static
//	//    - link into island
//	//    - transfer to constraint graph
//	// 8 For all shapes
//	//  - Destroy proxy in old tree
//	//  - Create proxy in new tree
//	// Notes:
//	// - the implementation below tries to minimize the number of predicates, so some
//	//   operations may have no effect, such as transferring a joint to the same set
func b2Body_SetType(tls *_Stack, bodyId BodyId, type1 BodyType) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var awakeSet, body, bodyA, bodyA1, bodyB, bodyB1, joint, joint1, joint2, jointSourceSet, otherBody, shape, sourceSet, staticSet, targetSet, world, v8 uintptr
	var edgeIndex, edgeIndex1, edgeIndex2, jointId, jointId1, jointId2, jointKey, otherBodyId, otherEdgeIndex, shapeId int32
	var forcePairCreation uint8
	var originalType b2BodyType1
	var transform WorldTransform
	var v1, v2, v3 bool
	var _ /* _a at bp+0 */ b2RecArgs_BodySetType
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = awakeSet, body, bodyA, bodyA1, bodyB, bodyB1, edgeIndex, edgeIndex1, edgeIndex2, forcePairCreation, joint, joint1, joint2, jointId, jointId1, jointId2, jointKey, jointSourceSet, originalType, otherBody, otherBodyId, otherEdgeIndex, shape, shapeId, sourceSet, staticSet, targetSet, transform, world, v1, v2, v3, v8
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodySetType)(unsafe.Add(nil, bp)) = b2RecArgs_BodySetType{
			Body:  bodyId,
			Type1: type1,
		}
		b2RecWrite_BodySetType(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	originalType = (*b2Body)(unsafe.Pointer(body)).Type1
	if originalType == type1 {
		return
	}
	// Stage 1: skip disabled bodies
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_disabledSet) {
		// Disabled bodies don't change solver sets or islands when they change type.
		(*b2Body)(unsafe.Pointer(body)).Type1 = type1
		if type1 == int32(b2_dynamicBody) {
			*(*uint32_t)(unsafe.Add(nil, body+80)) |= uint32(b2_dynamicFlag)
		} else {
			*(*uint32_t)(unsafe.Add(nil, body+80)) &= uint32FromInt32(^int32(b2_dynamicFlag))
		}
		b2SyncBodyFlags(tls, world, body)
		// Body type affects the mass properties
		b2UpdateBodyMassData(tls, world, body)
		return
	}
	// Stage 2: destroy all contacts
	b2DestroyBodyContacts(tls, world, body)
	// Stage 3: wake this body (does nothing if body is static), otherwise it will also wake
	// all bodies in the same sleeping solver set.
	b2WakeBody(tls, world, body)
	if v1 = !!(bool(0 <= int32(b2_staticSet)) && int32(b2_staticSet) < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+2964, __ccgo_ts+365, int32FromInt32(1214))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// Stage 4: move joints to temporary storage
	staticSet = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b2_staticSet))*88
	jointKey = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
	for jointKey != -int32(1) {
		jointId = jointKey >> int32(1)
		edgeIndex = jointKey & int32(1)
		if v2 = !!(0 <= jointId && jointId < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v2 {
			b2InternalAssert(tls, __ccgo_ts+1850, __ccgo_ts+365, int32FromInt32(1222))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		joint = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr(jointId)*72
		jointKey = (*(*b2JointEdge)(unsafe.Add(nil, joint+20+uintptr(edgeIndex)*12))).NextKey
		// Joint may be disabled by other body
		if (*b2Joint)(unsafe.Pointer(joint)).SetIndex == int32(b2_disabledSet) {
			continue
		}
		if v1 = !!(0 <= (*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId && (*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+3036, __ccgo_ts+365, int32FromInt32(1234))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		// Wake attached bodies. The b2WakeBody call above does not wake bodies
		// attached to a static body. But it is necessary because the body may have
		// no joints.
		bodyA = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId)*104
		if v2 = !!(0 <= (*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId && (*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
			b2InternalAssert(tls, __ccgo_ts+3124, __ccgo_ts+365, int32FromInt32(1235))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		bodyB = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId)*104
		b2WakeBody(tls, world, bodyA)
		b2WakeBody(tls, world, bodyB)
		// Remove joint from island
		b2UnlinkJoint(tls, world, joint)
		if v1 = !!(0 <= (*b2Joint)(unsafe.Pointer(joint)).SetIndex && (*b2Joint)(unsafe.Pointer(joint)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+3212, __ccgo_ts+365, int32FromInt32(1244))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		// It is necessary to transfer all joints to the static set
		// so they can be added to the constraint graph below and acquire consistent colors.
		jointSourceSet = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Joint)(unsafe.Pointer(joint)).SetIndex)*88
		b2TransferJoint(tls, world, staticSet, jointSourceSet, joint)
	}
	// Stage 5: change the body type and transfer body
	(*b2Body)(unsafe.Pointer(body)).Type1 = type1
	if type1 == int32(b2_dynamicBody) {
		*(*uint32_t)(unsafe.Add(nil, body+80)) |= uint32(b2_dynamicFlag)
	} else {
		*(*uint32_t)(unsafe.Add(nil, body+80)) &= uint32FromInt32(^int32(b2_dynamicFlag))
	}
	if v1 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+365, int32FromInt32(1260))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	awakeSet = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
	if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(body)).SetIndex && (*b2Body)(unsafe.Pointer(body)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+365, int32FromInt32(1261))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	sourceSet = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(body)).SetIndex)*88
	if type1 == int32(b2_staticBody) {
		v8 = staticSet
	} else {
		v8 = awakeSet
	}
	targetSet = v8
	// Transfer body
	b2TransferBody(tls, world, targetSet, sourceSet, body)
	// Stage 6: update island participation for the body
	if originalType == int32(b2_staticBody) {
		// Create island for body
		b2CreateIslandForBody(tls, world, int32(b2_awakeSet), body)
	} else {
		if type1 == int32(b2_staticBody) {
			// Remove body from island.
			b2RemoveBodyFromIsland(tls, world, body)
		}
	}
	// Stage 7: Transfer joints to the target set
	jointKey = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
	for jointKey != -int32(1) {
		jointId1 = jointKey >> int32(1)
		edgeIndex1 = jointKey & int32(1)
		if v1 = !!(0 <= jointId1 && jointId1 < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1850, __ccgo_ts+365, int32FromInt32(1286))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		joint1 = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr(jointId1)*72
		jointKey = (*(*b2JointEdge)(unsafe.Add(nil, joint1+20+uintptr(edgeIndex1)*12))).NextKey
		// Joint may be disabled by other body
		if (*b2Joint)(unsafe.Pointer(joint1)).SetIndex == int32(b2_disabledSet) {
			continue
		}
		// All joints were transferred to the static set in an earlier stage
		if v1 = !!((*b2Joint)(unsafe.Pointer(joint1)).SetIndex == int32(b2_staticSet)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+3290, __ccgo_ts+365, int32FromInt32(1297))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		if v1 = !!(0 <= (*(*b2JointEdge)(unsafe.Add(nil, joint1+20))).BodyId && (*(*b2JointEdge)(unsafe.Add(nil, joint1+20))).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+3036, __ccgo_ts+365, int32FromInt32(1299))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		bodyA1 = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b2JointEdge)(unsafe.Add(nil, joint1+20))).BodyId)*104
		if v2 = !!(0 <= (*(*b2JointEdge)(unsafe.Add(nil, joint1+20+1*12))).BodyId && (*(*b2JointEdge)(unsafe.Add(nil, joint1+20+1*12))).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
			b2InternalAssert(tls, __ccgo_ts+3124, __ccgo_ts+365, int32FromInt32(1300))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		bodyB1 = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b2JointEdge)(unsafe.Add(nil, joint1+20+1*12))).BodyId)*104
		if v3 = !!((*b2Body)(unsafe.Pointer(bodyA1)).SetIndex == int32(b2_staticSet) || (*b2Body)(unsafe.Pointer(bodyA1)).SetIndex == int32(b2_awakeSet)); !v3 {
			b2InternalAssert(tls, __ccgo_ts+3322, __ccgo_ts+365, int32FromInt32(1301))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		if v1 = !!((*b2Body)(unsafe.Pointer(bodyB1)).SetIndex == int32(b2_staticSet) || (*b2Body)(unsafe.Pointer(bodyB1)).SetIndex == int32(b2_awakeSet)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+3388, __ccgo_ts+365, int32FromInt32(1302))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		if (*b2Body)(unsafe.Pointer(bodyA1)).Type1 == int32(b2_dynamicBody) || (*b2Body)(unsafe.Pointer(bodyB1)).Type1 == int32(b2_dynamicBody) {
			b2TransferJoint(tls, world, awakeSet, staticSet, joint1)
		}
	}
	// Recreate shape proxies in broadphase
	transform = b2GetBodyTransformQuick(tls, world, body)
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		if v1 = !!(0 <= shapeId && shapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1908, __ccgo_ts+365, int32FromInt32(1315))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		shape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*280
		shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
		b2DestroyShapeProxy(tls, shape, world+40)
		forcePairCreation = boolUint8(true1 != 0)
		b2CreateShapeProxy(tls, world, shape, type1, transform, forcePairCreation)
	}
	// Relink all joints
	jointKey = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
	for jointKey != -int32(1) {
		jointId2 = jointKey >> int32(1)
		edgeIndex2 = jointKey & int32(1)
		if v1 = !!(0 <= jointId2 && jointId2 < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1850, __ccgo_ts+365, int32FromInt32(1329))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		joint2 = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr(jointId2)*72
		jointKey = (*(*b2JointEdge)(unsafe.Add(nil, joint2+20+uintptr(edgeIndex2)*12))).NextKey
		otherEdgeIndex = edgeIndex2 ^ int32(1)
		otherBodyId = (*(*b2JointEdge)(unsafe.Add(nil, joint2+20+uintptr(otherEdgeIndex)*12))).BodyId
		if v1 = !!(0 <= otherBodyId && otherBodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+3454, __ccgo_ts+365, int32FromInt32(1334))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		otherBody = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(otherBodyId)*104
		if (*b2Body)(unsafe.Pointer(otherBody)).SetIndex == int32(b2_disabledSet) {
			continue
		}
		if (*b2Body)(unsafe.Pointer(body)).Type1 != int32(b2_dynamicBody) && (*b2Body)(unsafe.Pointer(otherBody)).Type1 != int32(b2_dynamicBody) {
			continue
		}
		b2LinkJoint(tls, world, joint2)
	}
	b2SyncBodyFlags(tls, world, body)
	// Body type affects the mass
	b2UpdateBodyMassData(tls, world, body)
	b2ValidateSolverSets(tls, world)
	b2ValidateIsland(tls, world, (*b2Body)(unsafe.Pointer(body)).IslandId)
}

func b2Body_SetName(tls *_Stack, bodyId BodyId, name uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, world uintptr
	var _ /* _a at bp+0 */ b2RecArgs_BodySetName
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodySetName)(unsafe.Add(nil, bp)) = b2RecArgs_BodySetName{
			Body: bodyId,
			Name: name,
		}
		b2RecWrite_BodySetName(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	if name != 0 {
		strncpy(tls, body+90, name, uint64(_B2_NAME_LENGTH))
		*(*uint8)(unsafe.Add(nil, body+90+10)) = uint8(0)
	} else {
		memset(tls, body+90, 0, uint64(11))
	}
}

func b2Body_GetName(tls *_Stack, bodyId BodyId) (r uintptr) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return body + 90
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

func b2Body_GetLocalCenter(tls *_Stack, bodyId BodyId) (r Vec2) {
	var body1, bodySim, bodySim1, set, world1, v1, v2, v5 uintptr
	var v3, v4 bool
	_, _, _, _, _, _, _, _, _, _ = body1, bodySim, bodySim1, set, world1, v1, v2, v3, v4, v5
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	v1 = world1
	v2 = body1
	if v3 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).SetIndex && (*b2Body)(unsafe.Pointer(v2)).SetIndex < (*b2World)(unsafe.Pointer(v1)).SolverSets.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v1)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).SetIndex)*88
	if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).LocalIndex && (*b2Body)(unsafe.Pointer(v2)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).LocalIndex)*96
	v5 = bodySim
	goto _6
_6:
	bodySim1 = v5
	return (*b2BodySim)(unsafe.Pointer(bodySim1)).LocalCenter
}

func b2Body_GetWorldCenter(tls *_Stack, bodyId BodyId) (r Pos) {
	var body1, bodySim, bodySim1, set, world1, v1, v2, v5 uintptr
	var v3, v4 bool
	_, _, _, _, _, _, _, _, _, _ = body1, bodySim, bodySim1, set, world1, v1, v2, v3, v4, v5
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	v1 = world1
	v2 = body1
	if v3 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).SetIndex && (*b2Body)(unsafe.Pointer(v2)).SetIndex < (*b2World)(unsafe.Pointer(v1)).SolverSets.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v1)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).SetIndex)*88
	if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).LocalIndex && (*b2Body)(unsafe.Pointer(v2)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).LocalIndex)*96
	v5 = bodySim
	goto _6
_6:
	bodySim1 = v5
	return (*b2BodySim)(unsafe.Pointer(bodySim1)).Center
}

func b2Body_SetMassData(tls *_Stack, bodyId BodyId, massData MassData) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var body2, bodySim, bodySim1, contact, contactSim, s1, set, set1, state, world2, v4, v5, v8 uintptr
	var center, oldCenter, v12, v20 Pos
	var contactId, edgeIndex, edgeKey, shapeId int32
	var deltaLinear, v11, v22, v25, v26, v28, v29 Vec2
	var extent b2ShapeExtent
	var rx, ry, v24, v32, v33, v35 float32
	var v1, v2 bool
	var v10 WorldTransform
	var _ /* _a at bp+0 */ b2RecArgs_BodySetMassData
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body2, bodySim, bodySim1, center, contact, contactId, contactSim, deltaLinear, edgeIndex, edgeKey, extent, oldCenter, rx, ry, s1, set, set1, shapeId, state, world2, v1, v10, v11, v12, v2, v20, v22, v24, v25, v26, v28, v29, v32, v33, v35, v4, v5, v8
	if boolInt32(b2IsValidFloat(tls, massData.Mass) != 0 && massData.Mass >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+32, __ccgo_ts+3520, uintptr(unsafe.Pointer(&__func__81))))
		if v1 = !!(b2IsValidFloat(tls, massData.Mass) != 0 && massData.Mass >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+3520, __ccgo_ts+365, int32FromInt32(1432))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if boolInt32(b2IsValidFloat(tls, massData.RotationalInertia) != 0 && massData.RotationalInertia >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+32, __ccgo_ts+3577, uintptr(unsafe.Pointer(&__func__81))))
		if v1 = !!(b2IsValidFloat(tls, massData.RotationalInertia) != 0 && massData.RotationalInertia >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+3577, __ccgo_ts+365, int32FromInt32(1433))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if int32FromUint8(b2IsValidVec2(tls, massData.Center)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+32, __ccgo_ts+3660, uintptr(unsafe.Pointer(&__func__81))))
		if v1 = !!(b2IsValidVec2(tls, massData.Center) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+3660, __ccgo_ts+365, int32FromInt32(1434))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world2 = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world2 == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world2)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodySetMassData)(unsafe.Add(nil, bp)) = b2RecArgs_BodySetMassData{
			Body:     bodyId,
			MassData: massData,
		}
		b2RecWrite_BodySetMassData(tls, (*b2World)(unsafe.Pointer(world2)).Recording, bp)
	}
	body2 = b2GetBodyFullId(tls, world2, bodyId)
	v4 = world2
	v5 = body2
	if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(v5)).SetIndex && (*b2Body)(unsafe.Pointer(v5)).SetIndex < (*b2World)(unsafe.Pointer(v4)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v4)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v5)).SetIndex)*88
	if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(v5)).LocalIndex && (*b2Body)(unsafe.Pointer(v5)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v5)).LocalIndex)*96
	v8 = bodySim
	goto _9
_9:
	bodySim1 = v8
	(*b2Body)(unsafe.Pointer(body2)).Mass = massData.Mass
	(*b2Body)(unsafe.Pointer(body2)).Inertia = massData.RotationalInertia
	(*b2BodySim)(unsafe.Pointer(bodySim1)).LocalCenter = massData.Center
	oldCenter = (*b2BodySim)(unsafe.Pointer(bodySim1)).Center
	v10 = (*b2BodySim)(unsafe.Pointer(bodySim1)).Transform
	v11 = massData.Center
	rx = float32(v10.Q.C*v11.X) - float32(v10.Q.S*v11.Y)
	ry = float32(v10.Q.S*v11.X) + float32(v10.Q.C*v11.Y)
	v12 = Pos{
		X: v10.P.X + rx,
		Y: v10.P.Y + ry,
	}
	goto _13
_13:
	center = v12
	(*b2BodySim)(unsafe.Pointer(bodySim1)).Center = center
	(*b2BodySim)(unsafe.Pointer(bodySim1)).Center0 = center
	v4 = world2
	v5 = body2
	if (*b2Body)(unsafe.Pointer(v5)).SetIndex == int32(b2_awakeSet) {
		if v1 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(v4)).SolverSets.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+698, int32FromInt32(254))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		set1 = (*b2World)(unsafe.Pointer(v4)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
		if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(v5)).LocalIndex && (*b2Body)(unsafe.Pointer(v5)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set1)).BodyStates.Count); !v2 {
			b2InternalAssert(tls, __ccgo_ts+790, __ccgo_ts+698, int32FromInt32(255))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		v8 = (*b2SolverSet)(unsafe.Pointer(set1)).BodyStates.Data + uintptr((*b2Body)(unsafe.Pointer(v5)).LocalIndex)*32
		goto _18
	}
	v8 = uintptrFromInt32(0)
	goto _18
_18:
	// Update center of mass velocity
	state = v8
	if state != uintptrFromInt32(0) {
		v12 = (*b2BodySim)(unsafe.Pointer(bodySim1)).Center
		v20 = oldCenter
		v11 = Vec2{
			X: v12.X - v20.X,
			Y: v12.Y - v20.Y,
		}
		goto _23
	_23:
		v24 = (*b2BodyState)(unsafe.Pointer(state)).AngularVelocity
		v22 = v11
		v25 = Vec2{
			X: float32(-v24 * v22.Y),
			Y: float32(v24 * v22.X),
		}
		goto _27
	_27:
		deltaLinear = v25
		v26 = (*b2BodyState)(unsafe.Pointer(state)).LinearVelocity
		v28 = deltaLinear
		v29 = Vec2{
			X: v26.X + v28.X,
			Y: v26.Y + v28.Y,
		}
		goto _31
	_31:
		(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity = v29
	}
	if (*b2Body)(unsafe.Pointer(body2)).Mass > float32FromFloat32(0) {
		v24 = float32FromFloat32(1) / (*b2Body)(unsafe.Pointer(body2)).Mass
	} else {
		v24 = float32FromFloat32(0)
	}
	(*b2BodySim)(unsafe.Pointer(bodySim1)).InvMass = v24
	if (*b2Body)(unsafe.Pointer(body2)).Inertia > float32FromFloat32(0) {
		v24 = float32FromFloat32(1) / (*b2Body)(unsafe.Pointer(body2)).Inertia
	} else {
		v24 = float32FromFloat32(0)
	}
	(*b2BodySim)(unsafe.Pointer(bodySim1)).InvInertia = v24
	// Update extents using supplied mass center.
	(*b2BodySim)(unsafe.Pointer(bodySim1)).MinExtent = float32(float32(float32FromFloat32(100)*float32FromFloat32(1000)) * b2_lengthUnitsPerMeter)
	(*b2BodySim)(unsafe.Pointer(bodySim1)).MaxExtent = float32FromFloat32(0)
	shapeId = (*b2Body)(unsafe.Pointer(body2)).HeadShapeId
	for shapeId != -int32(1) {
		if v1 = !!(0 <= shapeId && shapeId < (*b2World)(unsafe.Pointer(world2)).Shapes.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1908, __ccgo_ts+365, int32FromInt32(1473))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		s1 = (*b2World)(unsafe.Pointer(world2)).Shapes.Data + uintptr(shapeId)*280
		extent = b2ComputeShapeExtent(tls, s1, massData.Center)
		v24 = (*b2BodySim)(unsafe.Pointer(bodySim1)).MinExtent
		v32 = extent.MinExtent
		if v24 < v32 {
			v35 = v24
		} else {
			v35 = v32
		}
		v33 = v35
		goto _38
	_38:
		(*b2BodySim)(unsafe.Pointer(bodySim1)).MinExtent = v33
		v24 = (*b2BodySim)(unsafe.Pointer(bodySim1)).MaxExtent
		v32 = extent.MaxExtent
		if v24 > v32 {
			v35 = v24
		} else {
			v35 = v32
		}
		v33 = v35
		goto _43
	_43:
		(*b2BodySim)(unsafe.Pointer(bodySim1)).MaxExtent = v33
		shapeId = (*b2Shape)(unsafe.Pointer(s1)).NextShapeId
	}
	// When the center of mass changes, any cached contact manifold becomes invalid.
	edgeKey = (*b2Body)(unsafe.Pointer(body2)).HeadContactKey
	for edgeKey != -int32(1) {
		contactId = edgeKey >> int32(1)
		edgeIndex = edgeKey & int32(1)
		if v1 = !!(0 <= contactId && contactId < (*b2World)(unsafe.Pointer(world2)).Contacts.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+301, __ccgo_ts+365, int32FromInt32(1487))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		contact = (*b2World)(unsafe.Pointer(world2)).Contacts.Data + uintptr(contactId)*64
		contactSim = b2GetContactSim(tls, world2, contact)
		*(*uint32_t)(unsafe.Add(nil, contactSim+44)) &= uint32FromInt32(^int32(b2_simRelativeTransformValid))
		edgeKey = (*(*b2ContactEdge)(unsafe.Add(nil, contact+uintptr(edgeIndex)*12))).NextKey
	}
	// Motion locks take priority over mass data.
	if (*b2Body)(unsafe.Pointer(body2)).Flags&uint32(b2_fixedRotation) != 0 {
		(*b2Body)(unsafe.Pointer(body2)).Inertia = float32FromFloat32(0)
		(*b2BodySim)(unsafe.Pointer(bodySim1)).InvInertia = float32FromFloat32(0)
	}
	*(*uint32_t)(unsafe.Add(nil, body2+80)) &= uint32FromInt32(^int32(b2_dirtyMass))
}

func b2Body_GetMassData(tls *_Stack, bodyId BodyId) (r MassData) {
	var body1, bodySim, bodySim1, set, world1, v1, v2, v5 uintptr
	var massData MassData
	var v3, v4 bool
	_, _, _, _, _, _, _, _, _, _, _ = body1, bodySim, bodySim1, massData, set, world1, v1, v2, v3, v4, v5
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	v1 = world1
	v2 = body1
	if v3 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).SetIndex && (*b2Body)(unsafe.Pointer(v2)).SetIndex < (*b2World)(unsafe.Pointer(v1)).SolverSets.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v1)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).SetIndex)*88
	if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).LocalIndex && (*b2Body)(unsafe.Pointer(v2)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).LocalIndex)*96
	v5 = bodySim
	goto _6
_6:
	bodySim1 = v5
	massData = MassData{
		Mass:              (*b2Body)(unsafe.Pointer(body1)).Mass,
		Center:            (*b2BodySim)(unsafe.Pointer(bodySim1)).LocalCenter,
		RotationalInertia: (*b2Body)(unsafe.Pointer(body1)).Inertia,
	}
	return massData
}

func b2Body_UpdateMassFromShapes(tls *_Stack, bodyId BodyId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, world uintptr
	var _ /* _a at bp+0 */ b2RecArgs_BodyUpdateMassFromShapes
	_, _ = body, world
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodyUpdateMassFromShapes)(unsafe.Add(nil, bp)) = b2RecArgs_BodyUpdateMassFromShapes{
			Body: bodyId,
		}
		b2RecWrite_BodyUpdateMassFromShapes(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	b2UpdateBodyMassData(tls, world, body)
}

func b2Body_SetLinearDamping(tls *_Stack, bodyId BodyId, linearDamping float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var body1, bodySim, bodySim1, set, world1, v2, v3, v6 uintptr
	var v1, v4 bool
	var _ /* _a at bp+0 */ b2RecArgs_BodySetLinearDamping
	_, _, _, _, _, _, _, _, _, _ = body1, bodySim, bodySim1, set, world1, v1, v2, v3, v4, v6
	if boolInt32(b2IsValidFloat(tls, linearDamping) != 0 && linearDamping >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+3693, uintptr(unsafe.Pointer(&__func__14))))
		if v1 = !!(b2IsValidFloat(tls, linearDamping) != 0 && linearDamping >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+3693, __ccgo_ts+365, int32FromInt32(1529))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world1 = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world1 == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world1)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodySetLinearDamping)(unsafe.Add(nil, bp)) = b2RecArgs_BodySetLinearDamping{
			Body:    bodyId,
			Damping: linearDamping,
		}
		b2RecWrite_BodySetLinearDamping(tls, (*b2World)(unsafe.Pointer(world1)).Recording, bp)
	}
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	v2 = world1
	v3 = body1
	if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(v3)).SetIndex && (*b2Body)(unsafe.Pointer(v3)).SetIndex < (*b2World)(unsafe.Pointer(v2)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v2)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v3)).SetIndex)*88
	if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v3)).LocalIndex && (*b2Body)(unsafe.Pointer(v3)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v3)).LocalIndex)*96
	v6 = bodySim
	goto _7
_7:
	bodySim1 = v6
	(*b2BodySim)(unsafe.Pointer(bodySim1)).LinearDamping = linearDamping
}

func b2Body_GetLinearDamping(tls *_Stack, bodyId BodyId) (r float32) {
	var body1, bodySim, bodySim1, set, world1, v1, v2, v5 uintptr
	var v3, v4 bool
	_, _, _, _, _, _, _, _, _, _ = body1, bodySim, bodySim1, set, world1, v1, v2, v3, v4, v5
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	v1 = world1
	v2 = body1
	if v3 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).SetIndex && (*b2Body)(unsafe.Pointer(v2)).SetIndex < (*b2World)(unsafe.Pointer(v1)).SolverSets.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v1)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).SetIndex)*88
	if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).LocalIndex && (*b2Body)(unsafe.Pointer(v2)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).LocalIndex)*96
	v5 = bodySim
	goto _6
_6:
	bodySim1 = v5
	return (*b2BodySim)(unsafe.Pointer(bodySim1)).LinearDamping
}

func b2Body_SetAngularDamping(tls *_Stack, bodyId BodyId, angularDamping float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var body1, bodySim, bodySim1, set, world1, v2, v3, v6 uintptr
	var v1, v4 bool
	var _ /* _a at bp+0 */ b2RecArgs_BodySetAngularDamping
	_, _, _, _, _, _, _, _, _, _ = body1, bodySim, bodySim1, set, world1, v1, v2, v3, v4, v6
	if boolInt32(b2IsValidFloat(tls, angularDamping) != 0 && angularDamping >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+3750, uintptr(unsafe.Pointer(&__func__15))))
		if v1 = !!(b2IsValidFloat(tls, angularDamping) != 0 && angularDamping >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+3750, __ccgo_ts+365, int32FromInt32(1554))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world1 = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world1 == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world1)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodySetAngularDamping)(unsafe.Add(nil, bp)) = b2RecArgs_BodySetAngularDamping{
			Body:    bodyId,
			Damping: angularDamping,
		}
		b2RecWrite_BodySetAngularDamping(tls, (*b2World)(unsafe.Pointer(world1)).Recording, bp)
	}
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	v2 = world1
	v3 = body1
	if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(v3)).SetIndex && (*b2Body)(unsafe.Pointer(v3)).SetIndex < (*b2World)(unsafe.Pointer(v2)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v2)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v3)).SetIndex)*88
	if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v3)).LocalIndex && (*b2Body)(unsafe.Pointer(v3)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v3)).LocalIndex)*96
	v6 = bodySim
	goto _7
_7:
	bodySim1 = v6
	(*b2BodySim)(unsafe.Pointer(bodySim1)).AngularDamping = angularDamping
}

func b2Body_GetAngularDamping(tls *_Stack, bodyId BodyId) (r float32) {
	var body1, bodySim, bodySim1, set, world1, v1, v2, v5 uintptr
	var v3, v4 bool
	_, _, _, _, _, _, _, _, _, _ = body1, bodySim, bodySim1, set, world1, v1, v2, v3, v4, v5
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	v1 = world1
	v2 = body1
	if v3 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).SetIndex && (*b2Body)(unsafe.Pointer(v2)).SetIndex < (*b2World)(unsafe.Pointer(v1)).SolverSets.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v1)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).SetIndex)*88
	if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).LocalIndex && (*b2Body)(unsafe.Pointer(v2)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).LocalIndex)*96
	v5 = bodySim
	goto _6
_6:
	bodySim1 = v5
	return (*b2BodySim)(unsafe.Pointer(bodySim1)).AngularDamping
}

func b2Body_SetGravityScale(tls *_Stack, bodyId BodyId, gravityScale float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var body1, bodySim, bodySim1, set, world1, v3, v4, v7 uintptr
	var v1, v2 bool
	var _ /* _a at bp+0 */ b2RecArgs_BodySetGravityScale
	_, _, _, _, _, _, _, _, _, _ = body1, bodySim, bodySim1, set, world1, v1, v2, v3, v4, v7
	if v1 = !!(b2Body_IsValid(tls, bodyId) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+387, __ccgo_ts+365, int32FromInt32(1579))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if int32FromUint8(b2IsValidFloat(tls, gravityScale)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+3809, uintptr(unsafe.Pointer(&__func__16))))
		if v1 = !!(b2IsValidFloat(tls, gravityScale) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+3809, __ccgo_ts+365, int32FromInt32(1580))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world1 = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world1 == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world1)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodySetGravityScale)(unsafe.Add(nil, bp)) = b2RecArgs_BodySetGravityScale{
			Body:  bodyId,
			Scale: gravityScale,
		}
		b2RecWrite_BodySetGravityScale(tls, (*b2World)(unsafe.Pointer(world1)).Recording, bp)
	}
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	v3 = world1
	v4 = body1
	if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(v4)).SetIndex && (*b2Body)(unsafe.Pointer(v4)).SetIndex < (*b2World)(unsafe.Pointer(v3)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v3)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v4)).SetIndex)*88
	if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(v4)).LocalIndex && (*b2Body)(unsafe.Pointer(v4)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v4)).LocalIndex)*96
	v7 = bodySim
	goto _8
_8:
	bodySim1 = v7
	(*b2BodySim)(unsafe.Pointer(bodySim1)).GravityScale = gravityScale
}

func b2Body_GetGravityScale(tls *_Stack, bodyId BodyId) (r float32) {
	var body1, bodySim, bodySim1, set, world1, v2, v3, v6 uintptr
	var v1, v4 bool
	_, _, _, _, _, _, _, _, _, _ = body1, bodySim, bodySim1, set, world1, v1, v2, v3, v4, v6
	if v1 = !!(b2Body_IsValid(tls, bodyId) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+387, __ccgo_ts+365, int32FromInt32(1597))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	v2 = world1
	v3 = body1
	if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(v3)).SetIndex && (*b2Body)(unsafe.Pointer(v3)).SetIndex < (*b2World)(unsafe.Pointer(v2)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v2)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v3)).SetIndex)*88
	if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v3)).LocalIndex && (*b2Body)(unsafe.Pointer(v3)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v3)).LocalIndex)*96
	v6 = bodySim
	goto _7
_7:
	bodySim1 = v6
	return (*b2BodySim)(unsafe.Pointer(bodySim1)).GravityScale
}

func b2Body_IsAwake(tls *_Stack, bodyId BodyId) (r uint8) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return boolUint8((*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet))
}

func b2Body_SetAwake(tls *_Stack, bodyId BodyId, awake uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, island, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_BodySetAwake
	_, _, _, _ = body, island, world, v1
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodySetAwake)(unsafe.Add(nil, bp)) = b2RecArgs_BodySetAwake{
			Body:  bodyId,
			Awake: awake,
		}
		b2RecWrite_BodySetAwake(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	if awake != 0 && (*b2Body)(unsafe.Pointer(body)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeBody(tls, world, body)
	} else {
		if int32FromUint8(awake) == false1 && (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_awakeSet) {
			if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(body)).IslandId && (*b2Body)(unsafe.Pointer(body)).IslandId < (*b2World)(unsafe.Pointer(world)).Islands.Count); !v1 {
				b2InternalAssert(tls, __ccgo_ts+3840, __ccgo_ts+365, int32FromInt32(1629))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			island = (*b2World)(unsafe.Pointer(world)).Islands.Data + uintptr((*b2Body)(unsafe.Pointer(body)).IslandId)*64
			if (*b2Island)(unsafe.Pointer(island)).ConstraintRemoveCount > 0 {
				// Must split the island before sleeping. This is expensive.
				b2SplitIsland(tls, world, (*b2Body)(unsafe.Pointer(body)).IslandId)
			}
			b2TrySleepIsland(tls, world, (*b2Body)(unsafe.Pointer(body)).IslandId)
		}
	}
}

func b2Body_WakeTouching(tls *_Stack, bodyId BodyId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, contact, otherBody, otherBody1, shapeA, shapeB, world uintptr
	var contactId, contactKey, edgeIndex int32
	var v1, v2, v3, v4 bool
	var _ /* _a at bp+0 */ b2RecArgs_BodyWakeTouching
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, contact, contactId, contactKey, edgeIndex, otherBody, otherBody1, shapeA, shapeB, world, v1, v2, v3, v4
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodyWakeTouching)(unsafe.Add(nil, bp)) = b2RecArgs_BodyWakeTouching{
			Body: bodyId,
		}
		b2RecWrite_BodyWakeTouching(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	contactKey = (*b2Body)(unsafe.Pointer(body)).HeadContactKey
	for contactKey != -int32(1) {
		contactId = contactKey >> int32(1)
		edgeIndex = contactKey & int32(1)
		if v1 = !!(0 <= contactId && contactId < (*b2World)(unsafe.Pointer(world)).Contacts.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+301, __ccgo_ts+365, int32FromInt32(1657))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		contact = (*b2World)(unsafe.Pointer(world)).Contacts.Data + uintptr(contactId)*64
		if v2 = !!(0 <= (*b2Contact)(unsafe.Pointer(contact)).ShapeIdA && (*b2Contact)(unsafe.Pointer(contact)).ShapeIdA < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v2 {
			b2InternalAssert(tls, __ccgo_ts+2142, __ccgo_ts+365, int32FromInt32(1658))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		shapeA = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b2Contact)(unsafe.Pointer(contact)).ShapeIdA)*280
		if v3 = !!(0 <= (*b2Contact)(unsafe.Pointer(contact)).ShapeIdB && (*b2Contact)(unsafe.Pointer(contact)).ShapeIdB < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v3 {
			b2InternalAssert(tls, __ccgo_ts+2220, __ccgo_ts+365, int32FromInt32(1659))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		shapeB = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b2Contact)(unsafe.Pointer(contact)).ShapeIdB)*280
		if (*b2Shape)(unsafe.Pointer(shapeA)).BodyId == bodyId.Index1-int32(1) {
			if v4 = !!(0 <= (*b2Shape)(unsafe.Pointer(shapeB)).BodyId && (*b2Shape)(unsafe.Pointer(shapeB)).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v4 {
				b2InternalAssert(tls, __ccgo_ts+3913, __ccgo_ts+365, int32FromInt32(1663))
			}
			_ = v4 || bool(int32FromInt32(0) != 0)
			otherBody = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b2Shape)(unsafe.Pointer(shapeB)).BodyId)*104
			b2WakeBody(tls, world, otherBody)
		} else {
			if v1 = !!(0 <= (*b2Shape)(unsafe.Pointer(shapeA)).BodyId && (*b2Shape)(unsafe.Pointer(shapeA)).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
				b2InternalAssert(tls, __ccgo_ts+3985, __ccgo_ts+365, int32FromInt32(1668))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			otherBody1 = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b2Shape)(unsafe.Pointer(shapeA)).BodyId)*104
			b2WakeBody(tls, world, otherBody1)
		}
		contactKey = (*(*b2ContactEdge)(unsafe.Add(nil, contact+uintptr(edgeIndex)*12))).NextKey
	}
}

func b2Body_IsEnabled(tls *_Stack, bodyId BodyId) (r uint8) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return boolUint8((*b2Body)(unsafe.Pointer(body)).SetIndex != int32(b2_disabledSet))
}

func b2Body_IsSleepingEnabled(tls *_Stack, bodyId BodyId) (r uint8) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return boolUint8((*b2Body)(unsafe.Pointer(body)).Flags&uint32(b2_enableSleep) == uint32(b2_enableSleep))
}

func b2Body_SetSleepThreshold(tls *_Stack, bodyId BodyId, sleepThreshold float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var body, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_BodySetSleepThreshold
	_, _, _ = body, world, v1
	if boolInt32(b2IsValidFloat(tls, sleepThreshold) != 0 && sleepThreshold >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+4057, uintptr(unsafe.Pointer(&__func__5))))
		if v1 = !!(b2IsValidFloat(tls, sleepThreshold) != 0 && sleepThreshold >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+4057, __ccgo_ts+365, int32FromInt32(1692))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodySetSleepThreshold)(unsafe.Add(nil, bp)) = b2RecArgs_BodySetSleepThreshold{
			Body:      bodyId,
			Threshold: sleepThreshold,
		}
		b2RecWrite_BodySetSleepThreshold(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
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

func b2Body_SetSafetyFactor(tls *_Stack, bodyId BodyId, safetyFactor float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var body, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_BodySetSafetyFactor
	_, _, _ = body, world, v1
	if boolInt32(b2IsValidFloat(tls, safetyFactor) != 0 && safetyFactor >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+4116, uintptr(unsafe.Pointer(&__func__6))))
		if v1 = !!(b2IsValidFloat(tls, safetyFactor) != 0 && safetyFactor >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+4116, __ccgo_ts+365, int32FromInt32(1714))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodySetSafetyFactor)(unsafe.Add(nil, bp)) = b2RecArgs_BodySetSafetyFactor{
			Body:  bodyId,
			Value: safetyFactor,
		}
		b2RecWrite_BodySetSafetyFactor(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	(*b2Body)(unsafe.Pointer(body)).SafetyFactor = safetyFactor
}

func b2Body_GetSafetyFactor(tls *_Stack, bodyId BodyId) (r float32) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return (*b2Body)(unsafe.Pointer(body)).SafetyFactor
}

func b2Body_EnableSleeping(tls *_Stack, bodyId BodyId, enableSleep uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, world uintptr
	var flag uint8
	var v1 uint32
	var _ /* _a at bp+0 */ b2RecArgs_BodyEnableSleeping
	_, _, _, _ = body, flag, world, v1
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodyEnableSleeping)(unsafe.Add(nil, bp)) = b2RecArgs_BodyEnableSleeping{
			Body: bodyId,
			Flag: enableSleep,
		}
		b2RecWrite_BodyEnableSleeping(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	flag = boolUint8((*b2Body)(unsafe.Pointer(body)).Flags&uint32(b2_enableSleep) == uint32(b2_enableSleep))
	if int32FromUint8(enableSleep) == int32FromUint8(flag) {
		return
	}
	if enableSleep != 0 {
		v1 = (*b2Body)(unsafe.Pointer(body)).Flags | uint32(b2_enableSleep)
	} else {
		v1 = (*b2Body)(unsafe.Pointer(body)).Flags & uint32FromInt32(^int32(b2_enableSleep))
	}
	(*b2Body)(unsafe.Pointer(body)).Flags = v1
	b2SyncBodyFlags(tls, world, body)
	if int32FromUint8(enableSleep) == false1 {
		b2WakeBody(tls, world, body)
	}
}

// C documentation
//
//	// Disabling a body requires a lot of detailed bookkeeping, but it is a valuable feature.
//	// The most challenging aspect is that joints may connect to bodies that are not disabled.
func b2Body_Disable(tls *_Stack, bodyId BodyId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, disabledSet, joint, jointSet, set, shape, world uintptr
	var edgeIndex, jointId, jointKey, shapeId int32
	var v1, v2, v3 bool
	var _ /* _a at bp+0 */ b2RecArgs_BodyDisable
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, disabledSet, edgeIndex, joint, jointId, jointKey, jointSet, set, shape, shapeId, world, v1, v2, v3
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodyDisable)(unsafe.Add(nil, bp)) = b2RecArgs_BodyDisable{
			Body: bodyId,
		}
		b2RecWrite_BodyDisable(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_disabledSet) {
		return
	}
	// Destroy contacts and wake bodies touching this body. This avoid floating bodies.
	// This is necessary even for static bodies.
	b2DestroyBodyContacts(tls, world, body)
	if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(body)).SetIndex && (*b2Body)(unsafe.Pointer(body)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+365, int32FromInt32(1783))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// The current solver set of the body
	set = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(body)).SetIndex)*88
	if v2 = !!(bool(0 <= int32(b2_disabledSet)) && int32(b2_disabledSet) < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+4171, __ccgo_ts+365, int32FromInt32(1786))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	// Disabled bodies and connected joints are moved to the disabled set
	disabledSet = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b2_disabledSet))*88
	// Unlink joints and transfer them to the disabled set
	jointKey = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
	for jointKey != -int32(1) {
		jointId = jointKey >> int32(1)
		edgeIndex = jointKey & int32(1)
		if v3 = !!(0 <= jointId && jointId < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v3 {
			b2InternalAssert(tls, __ccgo_ts+1850, __ccgo_ts+365, int32FromInt32(1795))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		joint = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr(jointId)*72
		jointKey = (*(*b2JointEdge)(unsafe.Add(nil, joint+20+uintptr(edgeIndex)*12))).NextKey
		// joint may already be disabled by other body
		if (*b2Joint)(unsafe.Pointer(joint)).SetIndex == int32(b2_disabledSet) {
			continue
		}
		if v1 = !!((*b2Joint)(unsafe.Pointer(joint)).SetIndex == (*b2SolverSet)(unsafe.Pointer(set)).SetIndex || (*b2SolverSet)(unsafe.Pointer(set)).SetIndex == int32(b2_staticSet) || (*b2Joint)(unsafe.Pointer(joint)).SetIndex == int32(b2_staticSet)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+4247, __ccgo_ts+365, int32FromInt32(1804))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		// Remove joint from island
		b2UnlinkJoint(tls, world, joint)
		if v1 = !!(0 <= (*b2Joint)(unsafe.Pointer(joint)).SetIndex && (*b2Joint)(unsafe.Pointer(joint)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+3212, __ccgo_ts+365, int32FromInt32(1810))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		// Transfer joint to disabled set
		jointSet = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Joint)(unsafe.Pointer(joint)).SetIndex)*88
		b2TransferJoint(tls, world, disabledSet, jointSet, joint)
	}
	// Remove shapes from broad-phase
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		if v1 = !!(0 <= shapeId && shapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1908, __ccgo_ts+365, int32FromInt32(1818))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		shape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*280
		shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
		b2DestroyShapeProxy(tls, shape, world+40)
	}
	// Disabled bodies are not in an island. If the island becomes empty it will be destroyed.
	b2RemoveBodyFromIsland(tls, world, body)
	// Transfer body sim
	b2TransferBody(tls, world, disabledSet, set, body)
	b2ValidateConnectivity(tls, world)
	b2ValidateSolverSets(tls, world)
}

func b2Body_Enable(tls *_Stack, bodyId BodyId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, bodyA, bodyB, disabledSet, joint, jointSet, shape, targetSet, world uintptr
	var edgeIndex, jointId, jointKey, jointSetId, setId, shapeId, v2 int32
	var forcePairCreation uint8
	var proxyType b2BodyType1
	var transform WorldTransform
	var v1, v3 bool
	var _ /* _a at bp+0 */ b2RecArgs_BodyEnable
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = body, bodyA, bodyB, disabledSet, edgeIndex, forcePairCreation, joint, jointId, jointKey, jointSet, jointSetId, proxyType, setId, shape, shapeId, targetSet, transform, world, v1, v2, v3
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodyEnable)(unsafe.Add(nil, bp)) = b2RecArgs_BodyEnable{
			Body: bodyId,
		}
		b2RecWrite_BodyEnable(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).SetIndex != int32(b2_disabledSet) {
		return
	}
	if v1 = !!(bool(0 <= int32(b2_disabledSet)) && int32(b2_disabledSet) < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+4171, __ccgo_ts+365, int32FromInt32(1849))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	disabledSet = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b2_disabledSet))*88
	if (*b2Body)(unsafe.Pointer(body)).Type1 == int32(b2_staticBody) {
		v2 = int32(b2_staticSet)
	} else {
		v2 = int32(b2_awakeSet)
	}
	setId = v2
	if v3 = !!(0 <= setId && setId < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+1708, __ccgo_ts+365, int32FromInt32(1851))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	targetSet = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setId)*88
	b2TransferBody(tls, world, targetSet, disabledSet, body)
	transform = b2GetBodyTransformQuick(tls, world, body)
	// Add shapes to broad-phase
	proxyType = (*b2Body)(unsafe.Pointer(body)).Type1
	forcePairCreation = boolUint8(true1 != 0)
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		if v1 = !!(0 <= shapeId && shapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1908, __ccgo_ts+365, int32FromInt32(1863))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		shape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*280
		shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
		b2CreateShapeProxy(tls, world, shape, proxyType, transform, forcePairCreation)
	}
	if setId != int32(b2_staticSet) {
		b2CreateIslandForBody(tls, world, setId, body)
	}
	// Transfer joints. If the other body is disabled, don't transfer.
	// If the other body is sleeping, wake it.
	jointKey = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
	for jointKey != -int32(1) {
		jointId = jointKey >> int32(1)
		edgeIndex = jointKey & int32(1)
		if v1 = !!(0 <= jointId && jointId < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1850, __ccgo_ts+365, int32FromInt32(1882))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		joint = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr(jointId)*72
		if v3 = !!((*b2Joint)(unsafe.Pointer(joint)).SetIndex == int32(b2_disabledSet)); !v3 {
			b2InternalAssert(tls, __ccgo_ts+4348, __ccgo_ts+365, int32FromInt32(1883))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		if v1 = !!((*b2Joint)(unsafe.Pointer(joint)).IslandId == -int32FromInt32(1)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+4382, __ccgo_ts+365, int32FromInt32(1884))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		jointKey = (*(*b2JointEdge)(unsafe.Add(nil, joint+20+uintptr(edgeIndex)*12))).NextKey
		if v1 = !!(0 <= (*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId && (*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+3036, __ccgo_ts+365, int32FromInt32(1888))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		bodyA = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId)*104
		if v3 = !!(0 <= (*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId && (*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v3 {
			b2InternalAssert(tls, __ccgo_ts+3124, __ccgo_ts+365, int32FromInt32(1889))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		bodyB = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId)*104
		if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_disabledSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_disabledSet) {
			// one body is still disabled
			continue
		}
		if (*b2Body)(unsafe.Pointer(bodyA)).Type1 != int32(b2_dynamicBody) && (*b2Body)(unsafe.Pointer(bodyB)).Type1 != int32(b2_dynamicBody) {
			jointSetId = int32(b2_staticSet)
		} else {
			if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_staticSet) {
				jointSetId = (*b2Body)(unsafe.Pointer(bodyB)).SetIndex
			} else {
				jointSetId = (*b2Body)(unsafe.Pointer(bodyA)).SetIndex
			}
		}
		if v1 = !!(0 <= jointSetId && jointSetId < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+4415, __ccgo_ts+365, int32FromInt32(1912))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		jointSet = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(jointSetId)*88
		b2TransferJoint(tls, world, jointSet, disabledSet, joint)
		// Now that the joint is in the correct set, I can link the joint in the island.
		if jointSetId != int32(b2_staticSet) {
			b2LinkJoint(tls, world, joint)
		}
	}
	b2ValidateSolverSets(tls, world)
}

func b2Body_SetMotionLocks(tls *_Stack, bodyId BodyId, locks MotionLocks) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body1, set, state, world1, v4, v5, v7 uintptr
	var newFlags uint32_t
	var v1 int32
	var v6, v9 bool
	var _ /* _a at bp+0 */ b2RecArgs_BodySetMotionLocks
	_, _, _, _, _, _, _, _, _, _, _ = body1, newFlags, set, state, world1, v1, v4, v5, v6, v7, v9
	world1 = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world1 == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world1)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodySetMotionLocks)(unsafe.Add(nil, bp)) = b2RecArgs_BodySetMotionLocks{
			Body:  bodyId,
			Locks: locks,
		}
		b2RecWrite_BodySetMotionLocks(tls, (*b2World)(unsafe.Pointer(world1)).Recording, bp)
	}
	newFlags = uint32(0)
	if locks.LinearX != 0 {
		v1 = int32(b2_lockLinearX)
	} else {
		v1 = 0
	}
	newFlags = newFlags | uint32FromInt32(v1)
	if locks.LinearY != 0 {
		v1 = int32(b2_lockLinearY)
	} else {
		v1 = 0
	}
	newFlags = newFlags | uint32FromInt32(v1)
	if locks.AngularZ != 0 {
		v1 = int32(b2_lockAngularZ)
	} else {
		v1 = 0
	}
	newFlags = newFlags | uint32FromInt32(v1)
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	if (*b2Body)(unsafe.Pointer(body1)).Flags&uint32(b2_allLocks) != newFlags {
		*(*uint32_t)(unsafe.Add(nil, body1+80)) &= uint32FromInt32(^int32(b2_allLocks))
		*(*uint32_t)(unsafe.Add(nil, body1+80)) |= newFlags
		b2SyncBodyFlags(tls, world1, body1)
		v4 = world1
		v5 = body1
		if (*b2Body)(unsafe.Pointer(v5)).SetIndex == int32(b2_awakeSet) {
			if v6 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(v4)).SolverSets.Count); !v6 {
				b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+698, int32FromInt32(254))
			}
			_ = v6 || bool(int32FromInt32(0) != 0)
			set = (*b2World)(unsafe.Pointer(v4)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
			if v9 = !!(0 <= (*b2Body)(unsafe.Pointer(v5)).LocalIndex && (*b2Body)(unsafe.Pointer(v5)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Count); !v9 {
				b2InternalAssert(tls, __ccgo_ts+790, __ccgo_ts+698, int32FromInt32(255))
			}
			_ = v9 || bool(int32FromInt32(0) != 0)
			v7 = (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Data + uintptr((*b2Body)(unsafe.Pointer(v5)).LocalIndex)*32
			goto _8
		}
		v7 = uintptrFromInt32(0)
		goto _8
	_8:
		state = v7
		if state != uintptrFromInt32(0) {
			if locks.LinearX != 0 {
				(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity.X = float32FromFloat32(0)
			}
			if locks.LinearY != 0 {
				(*b2BodyState)(unsafe.Pointer(state)).LinearVelocity.Y = float32FromFloat32(0)
			}
			if locks.AngularZ != 0 {
				(*b2BodyState)(unsafe.Pointer(state)).AngularVelocity = float32FromFloat32(0)
			}
		}
		// Motion locks can affect mass properties.
		b2UpdateBodyMassData(tls, world1, body1)
	}
}

func b2Body_GetMotionLocks(tls *_Stack, bodyId BodyId) (r MotionLocks) {
	var body, world uintptr
	var locks MotionLocks
	_, _, _ = body, locks, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	locks.LinearX = uint8(boolUint32((*b2Body)(unsafe.Pointer(body)).Flags&uint32(b2_lockLinearX) != 0))
	locks.LinearY = uint8(boolUint32((*b2Body)(unsafe.Pointer(body)).Flags&uint32(b2_lockLinearY) != 0))
	locks.AngularZ = uint8(boolUint32((*b2Body)(unsafe.Pointer(body)).Flags&uint32(b2_lockAngularZ) != 0))
	return locks
}

func b2Body_SetBullet(tls *_Stack, bodyId BodyId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, world uintptr
	var newFlag uint32_t
	var v1 int32
	var _ /* _a at bp+0 */ b2RecArgs_BodySetBullet
	_, _, _, _ = body, newFlag, world, v1
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodySetBullet)(unsafe.Add(nil, bp)) = b2RecArgs_BodySetBullet{
			Body: bodyId,
			Flag: flag,
		}
		b2RecWrite_BodySetBullet(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	if flag != 0 {
		v1 = int32(b2_isBullet)
	} else {
		v1 = 0
	}
	newFlag = uint32FromInt32(v1)
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).Flags&uint32(b2_isBullet) == newFlag {
		return
	}
	*(*uint32_t)(unsafe.Add(nil, body+80)) &= uint32FromInt32(^int32(b2_isBullet))
	*(*uint32_t)(unsafe.Add(nil, body+80)) |= newFlag
	b2SyncBodyFlags(tls, world, body)
}

func b2Body_IsBullet(tls *_Stack, bodyId BodyId) (r uint8) {
	var body1, bodySim, bodySim1, set, world1, v1, v2, v5 uintptr
	var v3, v4 bool
	_, _, _, _, _, _, _, _, _, _ = body1, bodySim, bodySim1, set, world1, v1, v2, v3, v4, v5
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	v1 = world1
	v2 = body1
	if v3 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).SetIndex && (*b2Body)(unsafe.Pointer(v2)).SetIndex < (*b2World)(unsafe.Pointer(v1)).SolverSets.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v1)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).SetIndex)*88
	if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).LocalIndex && (*b2Body)(unsafe.Pointer(v2)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).LocalIndex)*96
	v5 = bodySim
	goto _6
_6:
	bodySim1 = v5
	return boolUint8((*b2BodySim)(unsafe.Pointer(bodySim1)).Flags&uint32(b2_isBullet) != uint32(0))
}

func b2Body_EnableContactRecycling(tls *_Stack, bodyId BodyId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, world uintptr
	var newFlag uint32_t
	var v1 int32
	var _ /* _a at bp+0 */ b2RecArgs_BodyEnableContactRecycling
	_, _, _, _ = body, newFlag, world, v1
	world = b2GetWorldLocked(tls, int32FromUint16(bodyId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodyEnableContactRecycling)(unsafe.Add(nil, bp)) = b2RecArgs_BodyEnableContactRecycling{
			Body: bodyId,
			Flag: flag,
		}
		b2RecWrite_BodyEnableContactRecycling(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	if flag != 0 {
		v1 = int32(b2_bodyEnableContactRecycling)
	} else {
		v1 = 0
	}
	newFlag = uint32FromInt32(v1)
	body = b2GetBodyFullId(tls, world, bodyId)
	if (*b2Body)(unsafe.Pointer(body)).Flags&uint32(b2_bodyEnableContactRecycling) == newFlag {
		return
	}
	*(*uint32_t)(unsafe.Add(nil, body+80)) &= uint32FromInt32(^int32(b2_bodyEnableContactRecycling))
	*(*uint32_t)(unsafe.Add(nil, body+80)) |= newFlag
	b2SyncBodyFlags(tls, world, body)
}

func b2Body_IsContactRecyclingEnabled(tls *_Stack, bodyId BodyId) (r uint8) {
	var body, world uintptr
	_, _ = body, world
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	return boolUint8((*b2Body)(unsafe.Pointer(body)).Flags&uint32(b2_bodyEnableContactRecycling) != uint32(0))
}

func b2Body_EnableContactEvents(tls *_Stack, bodyId BodyId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, shape, world uintptr
	var shapeId int32
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_BodyEnableContactEvents
	_, _, _, _, _ = body, shape, shapeId, world, v1
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodyEnableContactEvents)(unsafe.Add(nil, bp)) = b2RecArgs_BodyEnableContactEvents{
			Body: bodyId,
			Flag: flag,
		}
		b2RecWrite_BodyEnableContactEvents(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		if v1 = !!(0 <= shapeId && shapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1908, __ccgo_ts+365, int32FromInt32(2056))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		shape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*280
		(*b2Shape)(unsafe.Pointer(shape)).EnableContactEvents = flag
		shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
	}
}

func b2Body_EnableHitEvents(tls *_Stack, bodyId BodyId, flag uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var body, shape, world uintptr
	var shapeId int32
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_BodyEnableHitEvents
	_, _, _, _, _ = body, shape, shapeId, world, v1
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_BodyEnableHitEvents)(unsafe.Add(nil, bp)) = b2RecArgs_BodyEnableHitEvents{
			Body: bodyId,
			Flag: flag,
		}
		b2RecWrite_BodyEnableHitEvents(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	body = b2GetBodyFullId(tls, world, bodyId)
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	for shapeId != -int32(1) {
		if v1 = !!(0 <= shapeId && shapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1908, __ccgo_ts+365, int32FromInt32(2070))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		shape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*280
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
	var body, shape, world uintptr
	var id ShapeId
	var shapeCount, shapeId int32
	var v1 bool
	_, _, _, _, _, _, _ = body, id, shape, shapeCount, shapeId, world, v1
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	shapeId = (*b2Body)(unsafe.Pointer(body)).HeadShapeId
	shapeCount = 0
	for shapeId != -int32(1) && shapeCount < capacity {
		if v1 = !!(0 <= shapeId && shapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1908, __ccgo_ts+365, int32FromInt32(2097))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		shape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*280
		id = ShapeId{
			Index1:     (*b2Shape)(unsafe.Pointer(shape)).Id + int32(1),
			World0:     bodyId.World0,
			Generation: (*b2Shape)(unsafe.Pointer(shape)).Generation,
		}
		*(*ShapeId)(unsafe.Add(nil, shapeArray+uintptr(shapeCount)*8)) = id
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
	var body, joint, world uintptr
	var edgeIndex, jointCount, jointId, jointKey int32
	var id JointId
	var v1 bool
	_, _, _, _, _, _, _, _, _ = body, edgeIndex, id, joint, jointCount, jointId, jointKey, world, v1
	world = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body = b2GetBodyFullId(tls, world, bodyId)
	jointKey = (*b2Body)(unsafe.Pointer(body)).HeadJointKey
	jointCount = 0
	for jointKey != -int32(1) && jointCount < capacity {
		jointId = jointKey >> int32(1)
		edgeIndex = jointKey & int32(1)
		if v1 = !!(0 <= jointId && jointId < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1850, __ccgo_ts+365, int32FromInt32(2127))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		joint = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr(jointId)*72
		id = JointId{
			Index1:     jointId + int32(1),
			World0:     bodyId.World0,
			Generation: (*b2Joint)(unsafe.Pointer(joint)).Generation,
		}
		*(*JointId)(unsafe.Add(nil, jointArray+uintptr(jointCount)*8)) = id
		jointCount = jointCount + int32(1)
		jointKey = (*(*b2JointEdge)(unsafe.Add(nil, joint+20+uintptr(edgeIndex)*12))).NextKey
	}
	return jointCount
}

func b2Body_GetMinExtent(tls *_Stack, bodyId BodyId) (r float32) {
	var body1, bodySim, bodySim1, set, world1, v1, v2, v5 uintptr
	var v3, v4 bool
	_, _, _, _, _, _, _, _, _, _ = body1, bodySim, bodySim1, set, world1, v1, v2, v3, v4, v5
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	v1 = world1
	v2 = body1
	if v3 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).SetIndex && (*b2Body)(unsafe.Pointer(v2)).SetIndex < (*b2World)(unsafe.Pointer(v1)).SolverSets.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v1)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).SetIndex)*88
	if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).LocalIndex && (*b2Body)(unsafe.Pointer(v2)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).LocalIndex)*96
	v5 = bodySim
	goto _6
_6:
	bodySim1 = v5
	return (*b2BodySim)(unsafe.Pointer(bodySim1)).MinExtent
}

func b2Body_GetMaxExtent(tls *_Stack, bodyId BodyId) (r float32) {
	var body1, bodySim, bodySim1, set, world1, v1, v2, v5 uintptr
	var v3, v4 bool
	_, _, _, _, _, _, _, _, _, _ = body1, bodySim, bodySim1, set, world1, v1, v2, v3, v4, v5
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	v1 = world1
	v2 = body1
	if v3 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).SetIndex && (*b2Body)(unsafe.Pointer(v2)).SetIndex < (*b2World)(unsafe.Pointer(v1)).SolverSets.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v1)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).SetIndex)*88
	if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).LocalIndex && (*b2Body)(unsafe.Pointer(v2)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).LocalIndex)*96
	v5 = bodySim
	goto _6
_6:
	bodySim1 = v5
	return (*b2BodySim)(unsafe.Pointer(bodySim1)).MaxExtent
}

func b2Body_GetMaxExtentFromOrigin(tls *_Stack, bodyId BodyId) (r float32) {
	var body1, bodySim, bodySim1, set, world1, v1, v2, v5 uintptr
	var v3, v4 bool
	var v7 Vec2
	var v8 float32
	_, _, _, _, _, _, _, _, _, _, _, _ = body1, bodySim, bodySim1, set, world1, v1, v2, v3, v4, v5, v7, v8
	world1 = b2GetWorld(tls, int32FromUint16(bodyId.World0))
	body1 = b2GetBodyFullId(tls, world1, bodyId)
	v1 = world1
	v2 = body1
	if v3 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).SetIndex && (*b2Body)(unsafe.Pointer(v2)).SetIndex < (*b2World)(unsafe.Pointer(v1)).SolverSets.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v1)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).SetIndex)*88
	if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v2)).LocalIndex && (*b2Body)(unsafe.Pointer(v2)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v2)).LocalIndex)*96
	v5 = bodySim
	goto _6
_6:
	bodySim1 = v5
	v7 = (*b2BodySim)(unsafe.Pointer(bodySim1)).LocalCenter
	v8 = sqrtf(tls, float32(v7.X*v7.X)+float32(v7.Y*v7.Y))
	goto _9
_9:
	return (*b2BodySim)(unsafe.Pointer(bodySim1)).MaxExtent + v8
}

func b2ResolveContactBodySim(tls *_Stack, world1 uintptr, awakeSims uintptr, staticSims uintptr, encodedIndex int32, bodyId int32) (r uintptr) {
	var bodySim, set, v4, v5, v8 uintptr
	var v1 uint8
	var v3, v6, v7 bool
	_, _, _, _, _, _, _, _, _ = bodySim, set, v1, v3, v4, v5, v6, v7, v8
	if encodedIndex >= 0 {
		return awakeSims + uintptr(encodedIndex)*96
	}
	v1 = boolUint8(encodedIndex < -int32(1))
	goto _2
_2:
	if v1 != 0 {
		return staticSims - uintptr(encodedIndex)*96 - uintptr(2)*96
	}
	if v3 = !!(0 <= bodyId && bodyId < (*b2World)(unsafe.Pointer(world1)).Bodies.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+642, __ccgo_ts+16660, int32FromInt32(450))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	v4 = world1
	v5 = (*b2World)(unsafe.Pointer(world1)).Bodies.Data + uintptr(bodyId)*104
	if v6 = !!(0 <= (*b2Body)(unsafe.Pointer(v5)).SetIndex && (*b2Body)(unsafe.Pointer(v5)).SetIndex < (*b2World)(unsafe.Pointer(v4)).SolverSets.Count); !v6 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v6 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v4)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v5)).SetIndex)*88
	if v7 = !!(0 <= (*b2Body)(unsafe.Pointer(v5)).LocalIndex && (*b2Body)(unsafe.Pointer(v5)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v7 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v7 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v5)).LocalIndex)*96
	v8 = bodySim
	goto _9
_9:
	return v8
}

func b2Body_IsValid(tls *_Stack, id BodyId) (r uint8) {
	var body, world uintptr
	var v1 bool
	_, _, _ = body, world, v1
	if int32(_B2_MAX_WORLDS) <= int32FromUint16(id.World0) {
		// invalid world
		return boolUint8(false1 != 0)
	}
	world = uintptr(unsafe.Pointer(&b2_worlds)) + uintptr(id.World0)*2792
	if int32FromUint16((*b2World)(unsafe.Pointer(world)).WorldId) != int32FromUint16(id.World0) {
		// world is free
		return boolUint8(false1 != 0)
	}
	if id.Index1 < int32(1) || (*b2World)(unsafe.Pointer(world)).Bodies.Count < id.Index1 {
		// invalid index
		return boolUint8(false1 != 0)
	}
	body = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(id.Index1-int32FromInt32(1))*104
	if (*b2Body)(unsafe.Pointer(body)).SetIndex == -int32(1) {
		// this was freed
		return boolUint8(false1 != 0)
	}
	if v1 = !!((*b2Body)(unsafe.Pointer(body)).LocalIndex != -int32FromInt32(1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+18421, __ccgo_ts+16660, int32FromInt32(1606))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if int32FromUint16((*b2Body)(unsafe.Pointer(body)).Generation) != int32FromUint16(id.Generation) {
		// this id is orphaned
		return boolUint8(false1 != 0)
	}
	return boolUint8(true1 != 0)
}

// Per op dispatch, the only place real public API names appear
// Body and shape ids have world0 replaced with the replay world's slot index

func b2RecMakeBodyId(tls *_Stack, rdr uintptr, recorded BodyId) (r BodyId) {
	var id BodyId
	_ = id
	id.Index1 = recorded.Index1
	id.World0 = uint16(uint32((*b2RecReader)(unsafe.Pointer(rdr)).ReplayWorldId.Index1) - uint32FromUint32(1))
	id.Generation = recorded.Generation
	return id
}

func b2RecCheckBodyId(tls *_Stack, rdr uintptr, got BodyId, rec BodyId) {
	b2RecCheckId(tls, rdr, __ccgo_ts+20828, got.Index1, uint32(got.Generation), rec.Index1, uint32(rec.Generation))
}

// C documentation
//
//	// Append a created body to the outliner tracking list. Ordinals are creation order and never reused.
func b2RecTrackBodyCreate(tls *_Stack, player uintptr, id BodyId) {
	b2RecGrow(tls, player+192, player+204, (*b2Replay)(unsafe.Pointer(player)).BodyIdCount+int32(1), (*b2Replay)(unsafe.Pointer(player)).BodyIdCount, int32FromInt64(8))
	*(*BodyId)(unsafe.Add(nil, (*b2Replay)(unsafe.Pointer(player)).BodyIds+uintptr((*b2Replay)(unsafe.Pointer(player)).BodyIdCount)*8)) = id
	*(*int32)(unsafe.Add(nil, player+200)) += int32(1)
}

// C documentation
//
//	// Leave a hole so later ordinals do not shift, keeping a stored selection stable across the playthrough
func b2RecTrackBodyDestroy(tls *_Stack, player uintptr, id BodyId) {
	var i int32
	_ = i
	i = 0
	for {
		if !(i < (*b2Replay)(unsafe.Pointer(player)).BodyIdCount) {
			break
		}
		if (*(*BodyId)(unsafe.Add(nil, (*b2Replay)(unsafe.Pointer(player)).BodyIds+uintptr(i)*8))).Index1 == id.Index1 && int32FromUint16((*(*BodyId)(unsafe.Add(nil, (*b2Replay)(unsafe.Pointer(player)).BodyIds+uintptr(i)*8))).World0) == int32FromUint16(id.World0) && int32FromUint16((*(*BodyId)(unsafe.Add(nil, (*b2Replay)(unsafe.Pointer(player)).BodyIds+uintptr(i)*8))).Generation) == int32FromUint16(id.Generation) {
			*(*BodyId)(unsafe.Add(nil, (*b2Replay)(unsafe.Pointer(player)).BodyIds+uintptr(i)*8)) = b2_nullBodyId
			return
		}
		goto _1
	_1:
		;
		i = i + 1
	}
}

// C documentation
//
//	// Snapshot bodies are restored as a struct image and never hit the CreateBody hook the tracker keys
//	// on, so the seed world must be walked once to populate the outliner list. Slot order is stable.
func b2RecSeedBodyIds(tls *_Stack, player uintptr) {
	var count, i int32
	var world uintptr
	_, _, _ = count, i, world
	world = b2GetWorldFromId(tls, (*b2Replay)(unsafe.Pointer(player)).Rdr.ReplayWorldId)
	(*b2Replay)(unsafe.Pointer(player)).BodyIdCount = 0
	count = (*b2World)(unsafe.Pointer(world)).Bodies.Count
	i = 0
	for {
		if !(i < count) {
			break
		}
		if (*(*b2Body)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Bodies.Data+uintptr(i)*104))).Id != i {
			goto _1 // free slot
		}
		b2RecTrackBodyCreate(tls, player, b2MakeBodyId(tls, world, i))
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b2BulletBodyTask(tls *_Stack, startIndex int32, endIndex int32, workerIndex int32, context uintptr) {
	var i, simIndex int32
	var stepContext, taskContext uintptr
	var v1 bool
	_, _, _, _, _ = i, simIndex, stepContext, taskContext, v1
	stepContext = context
	taskContext = (*b2World)(unsafe.Pointer((*b2StepContext)(unsafe.Pointer(stepContext)).World)).TaskContexts.Data + uintptr(workerIndex)*136
	if v1 = !!(startIndex <= endIndex); !v1 {
		b2InternalAssert(tls, __ccgo_ts+25604, __ccgo_ts+24283, int32FromInt32(1276))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	i = startIndex
	for {
		if !(i < endIndex) {
			break
		}
		simIndex = *(*int32)(unsafe.Add(nil, (*b2StepContext)(unsafe.Pointer(stepContext)).BulletBodies+uintptr(i)*4))
		b2SolveContinuous(tls, (*b2StepContext)(unsafe.Pointer(stepContext)).World, simIndex, taskContext, (*b2StepContext)(unsafe.Pointer(stepContext)).Dt)
		goto _2
	_2:
		;
		i = i + 1
	}
}

// C documentation
//
//	// Remove a body from a solver set and return the body swapped into the slot.
func b2RemoveBodySim(tls *_Stack, world uintptr, set uintptr, localIndex int32) (r uintptr) {
	var bodySims, movedBody uintptr
	var lastIndex int32
	var v1, v2 bool
	_, _, _, _, _ = bodySims, lastIndex, movedBody, v1, v2
	bodySims = set
	if v1 = !!(0 <= localIndex && localIndex < (*b2DynamicArray_b2BodySim)(unsafe.Pointer(bodySims)).Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+28376, __ccgo_ts+26452, int32FromInt32(556))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	lastIndex = (*b2DynamicArray_b2BodySim)(unsafe.Pointer(bodySims)).Count - int32(1)
	*(*int32)(unsafe.Add(nil, bodySims+8)) -= int32(1)
	if localIndex == lastIndex {
		return uintptrFromInt32(0)
	}
	*(*b2BodySim)(unsafe.Add(nil, (*b2DynamicArray_b2BodySim)(unsafe.Pointer(bodySims)).Data+uintptr(localIndex)*96)) = *(*b2BodySim)(unsafe.Add(nil, (*b2DynamicArray_b2BodySim)(unsafe.Pointer(bodySims)).Data+uintptr(lastIndex)*96))
	if v1 = !!(0 <= (*(*b2BodySim)(unsafe.Add(nil, (*b2DynamicArray_b2BodySim)(unsafe.Pointer(bodySims)).Data+uintptr(localIndex)*96))).BodyId && (*(*b2BodySim)(unsafe.Add(nil, (*b2DynamicArray_b2BodySim)(unsafe.Pointer(bodySims)).Data+uintptr(localIndex)*96))).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+28424, __ccgo_ts+26452, int32FromInt32(566))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	movedBody = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b2BodySim)(unsafe.Add(nil, (*b2DynamicArray_b2BodySim)(unsafe.Pointer(bodySims)).Data+uintptr(localIndex)*96))).BodyId)*104
	if v2 = !!((*b2Body)(unsafe.Pointer(movedBody)).LocalIndex == lastIndex); !v2 {
		b2InternalAssert(tls, __ccgo_ts+28534, __ccgo_ts+26452, int32FromInt32(567))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	(*b2Body)(unsafe.Pointer(movedBody)).LocalIndex = localIndex
	return movedBody
}

func b2TransferBody(tls *_Stack, world uintptr, targetSet uintptr, sourceSet uintptr, body uintptr) {
	var movedBody, sourceSim, state, targetSim, v12, v4, v5, v6, v9 uintptr
	var newCapacity, newSize, oldCapacity, oldSize, sourceIndex, targetIndex, v11, v7, v8 int32
	var v1 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = movedBody, newCapacity, newSize, oldCapacity, oldSize, sourceIndex, sourceSim, state, targetIndex, targetSim, v1, v11, v12, v4, v5, v6, v7, v8, v9
	if v1 = !!((*b2Body)(unsafe.Pointer(body)).HeadContactKey == -int32FromInt32(1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+28569, __ccgo_ts+26452, int32FromInt32(575))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if targetSet == sourceSet {
		return
	}
	// The use cases for transfer body (set body type, disable, etc) destroy all the contacts
	// first. Double check here because this function wouldn't work correctly otherwise.
	if v1 = !!((*b2Body)(unsafe.Pointer(body)).HeadContactKey == -int32FromInt32(1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+28569, __ccgo_ts+26452, int32FromInt32(584))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	sourceIndex = (*b2Body)(unsafe.Pointer(body)).LocalIndex
	if v1 = !!(0 <= sourceIndex && sourceIndex < (*b2SolverSet)(unsafe.Pointer(sourceSet)).BodySims.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+28607, __ccgo_ts+26452, int32FromInt32(587))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	sourceSim = (*b2SolverSet)(unsafe.Pointer(sourceSet)).BodySims.Data + uintptr(sourceIndex)*96
	targetIndex = (*b2SolverSet)(unsafe.Pointer(targetSet)).BodySims.Count
	v4 = targetSet
	v5 = targetSet + 8
	v6 = targetSet + 12
	v7 = int32(96)
	if *(*int32)(unsafe.Add(nil, v5)) >= *(*int32)(unsafe.Add(nil, v6)) {
		oldCapacity = *(*int32)(unsafe.Add(nil, v6))
		oldSize = oldCapacity * v7
		if oldCapacity == 0 {
			v8 = int32(8)
		} else {
			v8 = int32(2) * oldCapacity
		}
		newCapacity = v8
		newSize = newCapacity * v7
		*(*uintptr)(unsafe.Add(nil, v4)) = b2GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v4)), uint64FromInt32(oldSize), uint64FromInt32(newSize))
		*(*int32)(unsafe.Add(nil, v6)) = newCapacity
	}
	v12 = v5
	v11 = *(*int32)(unsafe.Pointer(v12))
	*(*int32)(unsafe.Pointer(v12)) = *(*int32)(unsafe.Pointer(v12)) + 1
	v9 = *(*uintptr)(unsafe.Add(nil, v4)) + uintptr(v11*v7)
	goto _10
_10:
	targetSim = v9
	memcpy(tls, targetSim, sourceSim, uint64(96))
	// Clear transient body flags
	*(*uint32_t)(unsafe.Add(nil, body+80)) &= uint32FromInt32(^int32(b2_bodyTransientFlags))
	*(*uint32_t)(unsafe.Add(nil, targetSim+92)) &= uint32FromInt32(^(int32(b2_isFast) | int32(b2_bodyTransientFlags)))
	// Remove body sim from solver set that owns it
	movedBody = b2RemoveBodySim(tls, world, sourceSet, sourceIndex)
	if movedBody != uintptrFromInt32(0) {
		// Fixup the moved body's contact sims.
		b2RefreshBodyContactIndices(tls, world, movedBody)
	}
	if (*b2SolverSet)(unsafe.Pointer(sourceSet)).SetIndex == int32(b2_awakeSet) {
		v4 = (*b2SolverSet)(unsafe.Pointer(sourceSet)).BodyStates.Data
		v5 = sourceSet + 16 + 8
		v7 = sourceIndex
		v8 = int32(32)
		if v1 = !!(0 <= v7 && v7 < *(*int32)(unsafe.Add(nil, v5)) && bool(__ccgo_ts+2029 != 0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2055, __ccgo_ts+2115, int32FromInt32(153))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		*(*int32)(unsafe.Add(nil, v5)) = *(*int32)(unsafe.Add(nil, v5)) - 1
		if v7 != *(*int32)(unsafe.Add(nil, v5)) {
			memcpy(tls, v4+uintptr(v7*v8), v4+uintptr(*(*int32)(unsafe.Add(nil, v5))*v8), uint64FromInt32(v8))
			_ = *(*int32)(unsafe.Add(nil, v5))
			goto _18
		}
		_ = -int32(1)
		goto _18
	_18:
	} else {
		if (*b2SolverSet)(unsafe.Pointer(targetSet)).SetIndex == int32(b2_awakeSet) {
			v4 = targetSet + 16
			v5 = targetSet + 16 + 8
			v6 = targetSet + 16 + 12
			v7 = int32(32)
			if *(*int32)(unsafe.Add(nil, v5)) >= *(*int32)(unsafe.Add(nil, v6)) {
				oldCapacity = *(*int32)(unsafe.Add(nil, v6))
				oldSize = oldCapacity * v7
				if oldCapacity == 0 {
					v8 = int32(8)
				} else {
					v8 = int32(2) * oldCapacity
				}
				newCapacity = v8
				newSize = newCapacity * v7
				*(*uintptr)(unsafe.Add(nil, v4)) = b2GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v4)), uint64FromInt32(oldSize), uint64FromInt32(newSize))
				*(*int32)(unsafe.Add(nil, v6)) = newCapacity
			}
			v12 = v5
			v11 = *(*int32)(unsafe.Pointer(v12))
			*(*int32)(unsafe.Pointer(v12)) = *(*int32)(unsafe.Pointer(v12)) + 1
			v9 = *(*uintptr)(unsafe.Add(nil, v4)) + uintptr(v11*v7)
			goto _25
		_25:
			state = v9
			*(*b2BodyState)(unsafe.Add(nil, state)) = b2_identityBodyState21
			(*b2BodyState)(unsafe.Pointer(state)).Flags = (*b2Body)(unsafe.Pointer(body)).Flags
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
	def.SafetyFactor = float32FromFloat32(0.5)
	def.SleepThreshold = float32(float32FromFloat32(0.05) * b2GetLengthUnitsPerMeter(tls))
	def.GravityScale = float32FromFloat32(1)
	def.EnableSleep = boolUint8(true1 != 0)
	def.EnableContactRecycling = boolUint8(true1 != 0)
	def.IsAwake = boolUint8(true1 != 0)
	def.IsEnabled = boolUint8(true1 != 0)
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}
