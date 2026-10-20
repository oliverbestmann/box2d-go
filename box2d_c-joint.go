package b2

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

// C documentation
//
//	// Notice that a joint cannot share the same color as a contact between the same two bodies. This means I can solve contacts and
//	// joints in parallel with each other within each color.
func b2AssignJointColor(tls *_Stack, graph uintptr, bodyIdA int32, bodyIdB int32, typeA BodyType, typeB BodyType) (r int32) {
	var blockIndex, blockIndex1, v4, v8 uint32_t
	var color, color1, color2, v3, v7 uintptr
	var i, i1, i2 int32
	var v1 bool
	var v5, v9 uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = blockIndex, blockIndex1, color, color1, color2, i, i1, i2, v1, v3, v4, v5, v7, v8, v9
	if v1 = !!(typeA == int32(b2_dynamicBody) || typeB == int32(b2_dynamicBody)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+5201, __ccgo_ts+5049, int32FromInt32(244))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if typeA == int32(b2_dynamicBody) && typeB == int32(b2_dynamicBody) {
		// Dynamic constraint colors cannot encroach on colors reserved for static constraints
		i = 0
		for {
			if !(i < int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(4)) {
				break
			}
			color = graph + uintptr(i)*64
			v3 = color
			v4 = uint32FromInt32(bodyIdA)
			blockIndex1 = v4 / uint32(64)
			if blockIndex1 >= (*b2BitSet)(unsafe.Pointer(v3)).BlockCount {
				v5 = boolUint8(false1 != 0)
				goto _6
			}
			v5 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v3)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v4%uint32FromInt32(64))) != uint64(0))
			goto _6
		_6:
			;
			if v1 = v5 != 0; !v1 {
				v7 = color
				v8 = uint32FromInt32(bodyIdB)
				blockIndex1 = v8 / uint32(64)
				if blockIndex1 >= (*b2BitSet)(unsafe.Pointer(v7)).BlockCount {
					v9 = boolUint8(false1 != 0)
					goto _10
				}
				v9 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v7)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v8%uint32FromInt32(64))) != uint64(0))
				goto _10
			_10:
			}
			if v1 || v9 != 0 {
				goto _2
			}
			v3 = color
			v4 = uint32FromInt32(bodyIdA)
			blockIndex = v4 / uint32(64)
			if blockIndex >= (*b2BitSet)(unsafe.Pointer(v3)).BlockCount {
				b2GrowBitSet(tls, v3, blockIndex+uint32(1))
			}
			*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v3)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v4 % uint32FromInt32(64))
			v3 = color
			v4 = uint32FromInt32(bodyIdB)
			blockIndex = v4 / uint32(64)
			if blockIndex >= (*b2BitSet)(unsafe.Pointer(v3)).BlockCount {
				b2GrowBitSet(tls, v3, blockIndex+uint32(1))
			}
			*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v3)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v4 % uint32FromInt32(64))
			return i
		_2:
			;
			i = i + 1
		}
	} else {
		if typeA == int32(b2_dynamicBody) {
			// Static constraint colors build from the end to get higher priority than dyn-dyn constraints
			i1 = int32FromInt32(_B2_GRAPH_COLOR_COUNT) - int32FromInt32(1) - int32FromInt32(1)
			for {
				if !(i1 >= int32(1)) {
					break
				}
				color1 = graph + uintptr(i1)*64
				v3 = color1
				v4 = uint32FromInt32(bodyIdA)
				blockIndex1 = v4 / uint32(64)
				if blockIndex1 >= (*b2BitSet)(unsafe.Pointer(v3)).BlockCount {
					v5 = boolUint8(false1 != 0)
					goto _20
				}
				v5 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v3)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v4%uint32FromInt32(64))) != uint64(0))
				goto _20
			_20:
				if v5 != 0 {
					goto _16
				}
				v3 = color1
				v4 = uint32FromInt32(bodyIdA)
				blockIndex = v4 / uint32(64)
				if blockIndex >= (*b2BitSet)(unsafe.Pointer(v3)).BlockCount {
					b2GrowBitSet(tls, v3, blockIndex+uint32(1))
				}
				*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v3)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v4 % uint32FromInt32(64))
				return i1
			_16:
				;
				i1 = i1 - 1
			}
		} else {
			if typeB == int32(b2_dynamicBody) {
				// Static constraint colors build from the end to get higher priority than dyn-dyn constraints
				i2 = int32FromInt32(_B2_GRAPH_COLOR_COUNT) - int32FromInt32(1) - int32FromInt32(1)
				for {
					if !(i2 >= int32(1)) {
						break
					}
					color2 = graph + uintptr(i2)*64
					v3 = color2
					v4 = uint32FromInt32(bodyIdB)
					blockIndex1 = v4 / uint32(64)
					if blockIndex1 >= (*b2BitSet)(unsafe.Pointer(v3)).BlockCount {
						v5 = boolUint8(false1 != 0)
						goto _27
					}
					v5 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v3)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v4%uint32FromInt32(64))) != uint64(0))
					goto _27
				_27:
					if v5 != 0 {
						goto _23
					}
					v3 = color2
					v4 = uint32FromInt32(bodyIdB)
					blockIndex = v4 / uint32(64)
					if blockIndex >= (*b2BitSet)(unsafe.Pointer(v3)).BlockCount {
						b2GrowBitSet(tls, v3, blockIndex+uint32(1))
					}
					*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v3)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v4 % uint32FromInt32(64))
					return i2
				_23:
					;
					i2 = i2 - 1
				}
			}
		}
	}
	return int32FromInt32(_B2_GRAPH_COLOR_COUNT) - int32FromInt32(1)
}

func b2CreateJointInGraph(tls *_Stack, world uintptr, joint uintptr) (r uintptr) {
	var bodyA, bodyB, graph, jointSim, v11, v3, v4, v5, v8 uintptr
	var bodyIdA, bodyIdB, colorIndex, newCapacity, newSize, oldCapacity, oldSize, v10, v6, v7 int32
	var v1, v2 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, bodyIdA, bodyIdB, colorIndex, graph, jointSim, newCapacity, newSize, oldCapacity, oldSize, v1, v10, v11, v2, v3, v4, v5, v6, v7, v8
	graph = world + 376
	bodyIdA = (*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId
	bodyIdB = (*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId
	if v1 = !!(0 <= bodyIdA && bodyIdA < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+4679, __ccgo_ts+5049, int32FromInt32(306))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(bodyIdA)*104
	if v2 = !!(0 <= bodyIdB && bodyIdB < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+4737, __ccgo_ts+5049, int32FromInt32(307))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(bodyIdB)*104
	colorIndex = b2AssignJointColor(tls, graph, bodyIdA, bodyIdB, (*b2Body)(unsafe.Pointer(bodyA)).Type1, (*b2Body)(unsafe.Pointer(bodyB)).Type1)
	v3 = graph + uintptr(colorIndex)*64 + 32
	v4 = graph + uintptr(colorIndex)*64 + 32 + 8
	v5 = graph + uintptr(colorIndex)*64 + 32 + 12
	v6 = int32(252)
	if *(*int32)(unsafe.Add(nil, v4)) >= *(*int32)(unsafe.Add(nil, v5)) {
		oldCapacity = *(*int32)(unsafe.Add(nil, v5))
		oldSize = oldCapacity * v6
		if oldCapacity == 0 {
			v7 = int32(8)
		} else {
			v7 = int32(2) * oldCapacity
		}
		newCapacity = v7
		newSize = newCapacity * v6
		*(*uintptr)(unsafe.Add(nil, v3)) = b2GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v3)), uint64FromInt32(oldSize), uint64FromInt32(newSize))
		*(*int32)(unsafe.Add(nil, v5)) = newCapacity
	}
	v11 = v4
	v10 = *(*int32)(unsafe.Pointer(v11))
	*(*int32)(unsafe.Pointer(v11)) = *(*int32)(unsafe.Pointer(v11)) + 1
	v8 = *(*uintptr)(unsafe.Add(nil, v3)) + uintptr(v10*v6)
	goto _9
_9:
	jointSim = v8
	memset(tls, jointSim, 0, uint64(252))
	(*b2Joint)(unsafe.Pointer(joint)).ColorIndex = colorIndex
	(*b2Joint)(unsafe.Pointer(joint)).LocalIndex = (*(*b2GraphColor)(unsafe.Add(nil, graph+uintptr(colorIndex)*64))).JointSims.Count - int32(1)
	return jointSim
}

func b2AddJointToGraph(tls *_Stack, world uintptr, jointSim uintptr, joint uintptr) {
	var jointDst uintptr
	_ = jointDst
	jointDst = b2CreateJointInGraph(tls, world, joint)
	memcpy(tls, jointDst, jointSim, uint64(252))
}

func b2RemoveJointFromGraph(tls *_Stack, world uintptr, bodyIdA int32, bodyIdB int32, colorIndex int32, localIndex int32) {
	var blockIndex, v3 uint32_t
	var color, graph, movedJoint, movedJointSim, v2, v5 uintptr
	var movedId, movedIndex, v10, v11, v13 int32
	var v1, v12, v15 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = blockIndex, color, graph, movedId, movedIndex, movedJoint, movedJointSim, v1, v10, v11, v12, v13, v15, v2, v3, v5
	graph = world + 376
	if v1 = !!(0 <= colorIndex && colorIndex < int32(_B2_GRAPH_COLOR_COUNT)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+5383, __ccgo_ts+5049, int32FromInt32(329))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	color = graph + uintptr(colorIndex)*64
	if colorIndex != int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1) {
		// May clear static bodies, no effect
		v2 = color
		v3 = uint32FromInt32(bodyIdA)
		blockIndex = v3 / uint32(64)
		if blockIndex >= (*b2BitSet)(unsafe.Pointer(v2)).BlockCount {
			goto _4
		}
		*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v2)).Bits+uintptr(blockIndex)*8)) &= ^(uint64FromInt32(1) << (v3 % uint32FromInt32(64)))
	_4:
		;
		v2 = color
		v3 = uint32FromInt32(bodyIdB)
		blockIndex = v3 / uint32(64)
		if blockIndex >= (*b2BitSet)(unsafe.Pointer(v2)).BlockCount {
			goto _7
		}
		*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v2)).Bits+uintptr(blockIndex)*8)) &= ^(uint64FromInt32(1) << (v3 % uint32FromInt32(64)))
	_7:
	}
	v2 = (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Data
	v5 = color + 32 + 8
	v10 = localIndex
	v11 = int32(252)
	if v1 = !!(0 <= v10 && v10 < *(*int32)(unsafe.Add(nil, v5)) && bool(__ccgo_ts+2029 != 0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+2055, __ccgo_ts+2115, int32FromInt32(153))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*int32)(unsafe.Add(nil, v5)) = *(*int32)(unsafe.Add(nil, v5)) - 1
	if v10 != *(*int32)(unsafe.Add(nil, v5)) {
		memcpy(tls, v2+uintptr(v10*v11), v2+uintptr(*(*int32)(unsafe.Add(nil, v5))*v11), uint64FromInt32(v11))
		v13 = *(*int32)(unsafe.Add(nil, v5))
		goto _14
	}
	v13 = -int32(1)
	goto _14
_14:
	movedIndex = v13
	if movedIndex != -int32(1) {
		// Fix moved joint
		movedJointSim = (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Data + uintptr(localIndex)*252
		movedId = (*b2JointSim)(unsafe.Pointer(movedJointSim)).JointId
		if v12 = !!(0 <= movedId && movedId < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v12 {
			b2InternalAssert(tls, __ccgo_ts+5612, __ccgo_ts+5049, int32FromInt32(345))
		}
		_ = v12 || bool(int32FromInt32(0) != 0)
		movedJoint = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr(movedId)*72
		if v15 = !!((*b2Joint)(unsafe.Pointer(movedJoint)).SetIndex == int32(b2_awakeSet)); !v15 {
			b2InternalAssert(tls, __ccgo_ts+5670, __ccgo_ts+5049, int32FromInt32(346))
		}
		_ = v15 || bool(int32FromInt32(0) != 0)
		if v1 = !!((*b2Joint)(unsafe.Pointer(movedJoint)).ColorIndex == colorIndex); !v1 {
			b2InternalAssert(tls, __ccgo_ts+5706, __ccgo_ts+5049, int32FromInt32(347))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		if v1 = !!((*b2Joint)(unsafe.Pointer(movedJoint)).LocalIndex == movedIndex); !v1 {
			b2InternalAssert(tls, __ccgo_ts+5743, __ccgo_ts+5049, int32FromInt32(348))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		(*b2Joint)(unsafe.Pointer(movedJoint)).LocalIndex = localIndex
	}
}

/** @} */

func b2DistanceJoint_SetRestLength(tls *_Stack, jointId JointId, length float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var base, joint, world uintptr
	var v1 bool
	var v2, v3, v4, v5, v7, v8 float32
	var _ /* _a at bp+0 */ b2RecArgs_DistanceJointSetRestLength
	_, _, _, _, _, _, _, _, _, _ = base, joint, world, v1, v2, v3, v4, v5, v7, v8
	if int32FromUint8(b2IsValidFloat(tls, length)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8172, uintptr(unsafe.Pointer(&__func__17))))
		if v1 = !!(b2IsValidFloat(tls, length) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8172, __ccgo_ts+8197, int32FromInt32(21))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_DistanceJointSetRestLength)(unsafe.Add(nil, bp)) = b2RecArgs_DistanceJointSetRestLength{
			Joint:  jointId,
			Length: length,
		}
		b2RecWrite_DistanceJointSetRestLength(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	joint = base + 92
	v2 = length
	v3 = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
	v4 = float32(float32(float32FromFloat32(100)*float32FromFloat32(1000)) * b2_lengthUnitsPerMeter)
	if v2 < v3 {
		v7 = v3
	} else {
		if v2 > v4 {
			v8 = v4
		} else {
			v8 = v2
		}
		v7 = v8
	}
	v5 = v7
	goto _6
_6:
	(*b2DistanceJoint)(unsafe.Pointer(joint)).Length = v5
	(*b2DistanceJoint)(unsafe.Pointer(joint)).Impulse = float32FromFloat32(0)
	(*b2DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse = float32FromFloat32(0)
	(*b2DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse = float32FromFloat32(0)
}

func b2DistanceJoint_GetRestLength(tls *_Stack, jointId JointId) (r float32) {
	var base, joint uintptr
	_, _ = base, joint
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	joint = base + 92
	return (*b2DistanceJoint)(unsafe.Pointer(joint)).Length
}

func b2DistanceJoint_EnableLimit(tls *_Stack, jointId JointId, enableLimit uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, joint, world uintptr
	var _ /* _a at bp+0 */ b2RecArgs_DistanceJointEnableLimit
	_, _, _ = base, joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_DistanceJointEnableLimit)(unsafe.Add(nil, bp)) = b2RecArgs_DistanceJointEnableLimit{
			Joint:       jointId,
			EnableLimit: enableLimit,
		}
		b2RecWrite_DistanceJointEnableLimit(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	joint = base + 92
	(*b2DistanceJoint)(unsafe.Pointer(joint)).EnableLimit = enableLimit
}

func b2DistanceJoint_IsLimitEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	return (*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.EnableLimit
}

func b2DistanceJoint_SetLimits(tls *_Stack, jointId JointId, minLength float32, maxLength float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var base, joint, world uintptr
	var v1 bool
	var v3, v4, v5, v6, v8, v9 float32
	var _ /* _a at bp+0 */ b2RecArgs_DistanceJointSetLimits
	_, _, _, _, _, _, _, _, _, _ = base, joint, world, v1, v3, v4, v5, v6, v8, v9
	if int32FromUint8(b2IsValidFloat(tls, minLength)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8229, uintptr(unsafe.Pointer(&__func__32))))
		if v1 = !!(b2IsValidFloat(tls, minLength) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8229, __ccgo_ts+8197, int32FromInt32(58))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if int32FromUint8(b2IsValidFloat(tls, maxLength)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8257, uintptr(unsafe.Pointer(&__func__32))))
		if v1 = !!(b2IsValidFloat(tls, maxLength) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8257, __ccgo_ts+8197, int32FromInt32(59))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_DistanceJointSetLimits)(unsafe.Add(nil, bp)) = b2RecArgs_DistanceJointSetLimits{
			Joint:     jointId,
			MinLength: minLength,
			MaxLength: maxLength,
		}
		b2RecWrite_DistanceJointSetLimits(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	joint = base + 92
	v3 = minLength
	v4 = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
	v5 = float32(float32(float32FromFloat32(100)*float32FromFloat32(1000)) * b2_lengthUnitsPerMeter)
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
	minLength = v6
	v3 = maxLength
	v4 = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
	v5 = float32(float32(float32FromFloat32(100)*float32FromFloat32(1000)) * b2_lengthUnitsPerMeter)
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
	goto _14
_14:
	maxLength = v6
	v3 = minLength
	v4 = maxLength
	if v3 < v4 {
		v6 = v3
	} else {
		v6 = v4
	}
	v5 = v6
	goto _20
_20:
	(*b2DistanceJoint)(unsafe.Pointer(joint)).MinLength = v5
	v3 = minLength
	v4 = maxLength
	if v3 > v4 {
		v6 = v3
	} else {
		v6 = v4
	}
	v5 = v6
	goto _25
_25:
	(*b2DistanceJoint)(unsafe.Pointer(joint)).MaxLength = v5
}

func b2DistanceJoint_GetLowerLimit(tls *_Stack, jointId JointId) (r float32) {
	var base, joint uintptr
	_, _ = base, joint
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	joint = base + 92
	return (*b2DistanceJoint)(unsafe.Pointer(joint)).MinLength
}

func b2DistanceJoint_GetUpperLimit(tls *_Stack, jointId JointId) (r float32) {
	var base, joint uintptr
	_, _ = base, joint
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	joint = base + 92
	return (*b2DistanceJoint)(unsafe.Pointer(joint)).MaxLength
}

func b2DistanceJoint_GetCurrentLength(tls *_Stack, jointId JointId) (r1 float32) {
	var base1, world uintptr
	var d, pA, pB, v11, v12, v15, v16, v18, v19, v20, v22 Vec2
	var length, x, y, v23 float32
	var r, transformA, transformB, v10, v14, v4, v8 Transform
	var wxfA, v2, v6 WorldTransform
	var v1 bool
	var v3, v7 Pos
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = base1, d, length, pA, pB, r, transformA, transformB, world, wxfA, x, y, v1, v10, v11, v12, v14, v15, v16, v18, v19, v2, v20, v22, v23, v3, v4, v6, v7, v8
	base1 = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+8197, int32FromInt32(91))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return float32FromFloat32(0)
	}
	// Relative to body A so the difference stays in float precision far from the origin
	wxfA = b2GetBodyTransform(tls, world, (*b2JointSim)(unsafe.Pointer(base1)).BodyIdA)
	v2 = wxfA
	v3 = wxfA.P
	r.Q = v2.Q
	r.P = Vec2{
		X: v2.P.X - v3.X,
		Y: v2.P.Y - v3.Y,
	}
	v4 = r
	goto _5
_5:
	transformA = v4
	v6 = b2GetBodyTransform(tls, world, (*b2JointSim)(unsafe.Pointer(base1)).BodyIdB)
	v7 = wxfA.P
	r.Q = v6.Q
	r.P = Vec2{
		X: v6.P.X - v7.X,
		Y: v6.P.Y - v7.Y,
	}
	v8 = r
	goto _9
_9:
	transformB = v8
	v10 = transformA
	v11 = (*b2JointSim)(unsafe.Pointer(base1)).LocalFrameA.P
	x = float32(v10.Q.C*v11.X) - float32(v10.Q.S*v11.Y) + v10.P.X
	y = float32(v10.Q.S*v11.X) + float32(v10.Q.C*v11.Y) + v10.P.Y
	v12 = Vec2{
		X: x,
		Y: y,
	}
	goto _13
_13:
	pA = v12
	v14 = transformB
	v15 = (*b2JointSim)(unsafe.Pointer(base1)).LocalFrameB.P
	x = float32(v14.Q.C*v15.X) - float32(v14.Q.S*v15.Y) + v14.P.X
	y = float32(v14.Q.S*v15.X) + float32(v14.Q.C*v15.Y) + v14.P.Y
	v16 = Vec2{
		X: x,
		Y: y,
	}
	goto _17
_17:
	pB = v16
	v18 = pB
	v19 = pA
	v20 = Vec2{
		X: v18.X - v19.X,
		Y: v18.Y - v19.Y,
	}
	goto _21
_21:
	d = v20
	v22 = d
	v23 = sqrtf(tls, float32(v22.X*v22.X)+float32(v22.Y*v22.Y))
	goto _24
_24:
	length = v23
	return length
}

func b2DistanceJoint_EnableSpring(tls *_Stack, jointId JointId, enableSpring uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var base, world uintptr
	var _ /* _a at bp+0 */ b2RecArgs_DistanceJointEnableSpring
	_, _ = base, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_DistanceJointEnableSpring)(unsafe.Add(nil, bp)) = b2RecArgs_DistanceJointEnableSpring{
			Joint:        jointId,
			EnableSpring: enableSpring,
		}
		b2RecWrite_DistanceJointEnableSpring(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	(*b2JointSim)(unsafe.Pointer(base)).ｆ__ccgo15_92.DistanceJoint.EnableSpring = enableSpring
}

func b2DistanceJoint_IsSpringEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var base uintptr
	_ = base
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	return (*b2JointSim)(unsafe.Pointer(base)).ｆ__ccgo15_92.DistanceJoint.EnableSpring
}

func b2DistanceJoint_SetSpringForceRange(tls *_Stack, jointId JointId, lowerForce float32, upperForce float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var base, world uintptr
	var v1 bool
	var v4, v5, v6, v8 float32
	var _ /* _a at bp+0 */ b2RecArgs_DistanceJointSetSpringForceRange
	_, _, _, _, _, _, _ = base, world, v1, v4, v5, v6, v8
	if int32FromUint8(b2IsValidFloat(tls, lowerForce)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8285, uintptr(unsafe.Pointer(&__func__22))))
		if v1 = !!(b2IsValidFloat(tls, lowerForce) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8285, __ccgo_ts+8197, int32FromInt32(125))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if int32FromUint8(b2IsValidFloat(tls, upperForce)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8314, uintptr(unsafe.Pointer(&__func__22))))
		if v1 = !!(b2IsValidFloat(tls, upperForce) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8314, __ccgo_ts+8197, int32FromInt32(126))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_DistanceJointSetSpringForceRange)(unsafe.Add(nil, bp)) = b2RecArgs_DistanceJointSetSpringForceRange{
			Joint:      jointId,
			LowerForce: lowerForce,
			UpperForce: upperForce,
		}
		b2RecWrite_DistanceJointSetSpringForceRange(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	if v1 = !!(lowerForce <= upperForce); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8343, __ccgo_ts+8197, int32FromInt32(130))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	v4 = lowerForce
	v5 = upperForce
	if v4 < v5 {
		v8 = v4
	} else {
		v8 = v5
	}
	v6 = v8
	goto _7
_7:
	(*b2JointSim)(unsafe.Pointer(base)).ｆ__ccgo15_92.DistanceJoint.LowerSpringForce = v6
	v4 = lowerForce
	v5 = upperForce
	if v4 > v5 {
		v8 = v4
	} else {
		v8 = v5
	}
	v6 = v8
	goto _12
_12:
	(*b2JointSim)(unsafe.Pointer(base)).ｆ__ccgo15_92.DistanceJoint.UpperSpringForce = v6
}

func b2DistanceJoint_GetSpringForceRange(tls *_Stack, jointId JointId, lowerForce uintptr, upperForce uintptr) {
	var base uintptr
	_ = base
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	*(*float32)(unsafe.Add(nil, lowerForce)) = (*b2JointSim)(unsafe.Pointer(base)).ｆ__ccgo15_92.DistanceJoint.LowerSpringForce
	*(*float32)(unsafe.Add(nil, upperForce)) = (*b2JointSim)(unsafe.Pointer(base)).ｆ__ccgo15_92.DistanceJoint.UpperSpringForce
}

func b2DistanceJoint_SetSpringHertz(tls *_Stack, jointId JointId, hertz float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var base, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_DistanceJointSetSpringHertz
	_, _, _ = base, world, v1
	if int32FromUint8(b2IsValidFloat(tls, hertz)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8368, uintptr(unsafe.Pointer(&__func__9))))
		if v1 = !!(b2IsValidFloat(tls, hertz) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8368, __ccgo_ts+8197, int32FromInt32(145))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_DistanceJointSetSpringHertz)(unsafe.Add(nil, bp)) = b2RecArgs_DistanceJointSetSpringHertz{
			Joint: jointId,
			Hertz: hertz,
		}
		b2RecWrite_DistanceJointSetSpringHertz(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	(*b2JointSim)(unsafe.Pointer(base)).ｆ__ccgo15_92.DistanceJoint.Hertz = hertz
}

func b2DistanceJoint_SetSpringDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var base, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_DistanceJointSetSpringDampingRatio
	_, _, _ = base, world, v1
	if int32FromUint8(b2IsValidFloat(tls, dampingRatio)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8392, uintptr(unsafe.Pointer(&__func__10))))
		if v1 = !!(b2IsValidFloat(tls, dampingRatio) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8392, __ccgo_ts+8197, int32FromInt32(155))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_DistanceJointSetSpringDampingRatio)(unsafe.Add(nil, bp)) = b2RecArgs_DistanceJointSetSpringDampingRatio{
			Joint:        jointId,
			DampingRatio: dampingRatio,
		}
		b2RecWrite_DistanceJointSetSpringDampingRatio(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	(*b2JointSim)(unsafe.Pointer(base)).ｆ__ccgo15_92.DistanceJoint.DampingRatio = dampingRatio
}

func b2DistanceJoint_GetSpringHertz(tls *_Stack, jointId JointId) (r float32) {
	var base, joint uintptr
	_, _ = base, joint
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	joint = base + 92
	return (*b2DistanceJoint)(unsafe.Pointer(joint)).Hertz
}

func b2DistanceJoint_GetSpringDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var base, joint uintptr
	_, _ = base, joint
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	joint = base + 92
	return (*b2DistanceJoint)(unsafe.Pointer(joint)).DampingRatio
}

func b2DistanceJoint_GetSpringForce(tls *_Stack, jointId JointId) (r float32) {
	var base, joint, world uintptr
	_, _, _ = base, joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	joint = base + 92
	if (*b2DistanceJoint)(unsafe.Pointer(joint)).EnableSpring != 0 {
		return float32((*b2DistanceJoint)(unsafe.Pointer(joint)).Impulse * (*b2World)(unsafe.Pointer(world)).Inv_h)
	}
	return float32FromFloat32(0)
}

func b2DistanceJoint_EnableMotor(tls *_Stack, jointId JointId, enableMotor uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* _a at bp+0 */ b2RecArgs_DistanceJointEnableMotor
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_DistanceJointEnableMotor)(unsafe.Add(nil, bp)) = b2RecArgs_DistanceJointEnableMotor{
			Joint:       jointId,
			EnableMotor: enableMotor,
		}
		b2RecWrite_DistanceJointEnableMotor(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	if int32FromUint8(enableMotor) != int32FromUint8((*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.EnableMotor) {
		(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.EnableMotor = enableMotor
		(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.MotorImpulse = float32FromFloat32(0)
	}
}

func b2DistanceJoint_IsMotorEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	return (*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.EnableMotor
}

func b2DistanceJoint_SetTargetMotorSpeed(tls *_Stack, jointId JointId, motorSpeed float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_DistanceJointSetTargetMotorSpeed
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, motorSpeed)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8423, uintptr(unsafe.Pointer(&__func__18))))
		if v1 = !!(b2IsValidFloat(tls, motorSpeed) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8423, __ccgo_ts+8197, int32FromInt32(210))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_DistanceJointSetTargetMotorSpeed)(unsafe.Add(nil, bp)) = b2RecArgs_DistanceJointSetTargetMotorSpeed{
			Joint:      jointId,
			MotorSpeed: motorSpeed,
		}
		b2RecWrite_DistanceJointSetTargetMotorSpeed(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.MotorSpeed = motorSpeed
}

func b2DistanceJoint_GetTargetMotorSpeed(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	return (*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.MotorSpeed
}

func b2DistanceJoint_GetCurrentMotorForce(tls *_Stack, jointId JointId) (r float32) {
	var base, world uintptr
	_, _ = base, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	return float32((*b2World)(unsafe.Pointer(world)).Inv_h * (*b2JointSim)(unsafe.Pointer(base)).ｆ__ccgo15_92.DistanceJoint.MotorImpulse)
}

func b2DistanceJoint_SetMaxMotorForce(tls *_Stack, jointId JointId, force float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_DistanceJointSetMaxMotorForce
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, force)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8452, uintptr(unsafe.Pointer(&__func__19))))
		if v1 = !!(b2IsValidFloat(tls, force) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8452, __ccgo_ts+8197, int32FromInt32(233))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_DistanceJointSetMaxMotorForce)(unsafe.Add(nil, bp)) = b2RecArgs_DistanceJointSetMaxMotorForce{
			Joint: jointId,
			Force: force,
		}
		b2RecWrite_DistanceJointSetMaxMotorForce(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.MaxMotorForce = force
}

func b2DistanceJoint_GetMaxMotorForce(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	return (*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.MaxMotorForce
}

func b2GetDistanceJointForce(tls *_Stack, world uintptr, base1 uintptr) (r1 Vec2) {
	var axis, d, pA, pB, u, v10, v11, v14, v15, v17, v18, v19, v21, v22, v25, v26 Vec2
	var force, lengthSquared, s1, x, y, v24 float32
	var joint uintptr
	var r, transformA, transformB, v13, v3, v7, v9 Transform
	var worldTransformA, worldTransformB, v1, v5 WorldTransform
	var v2, v6 Pos
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axis, d, force, joint, lengthSquared, pA, pB, r, s1, transformA, transformB, u, worldTransformA, worldTransformB, x, y, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v24, v25, v26, v3, v5, v6, v7, v9
	joint = base1 + 92
	worldTransformA = b2GetBodyTransform(tls, world, (*b2JointSim)(unsafe.Pointer(base1)).BodyIdA)
	v1 = worldTransformA
	v2 = worldTransformA.P
	r.Q = v1.Q
	r.P = Vec2{
		X: v1.P.X - v2.X,
		Y: v1.P.Y - v2.Y,
	}
	v3 = r
	goto _4
_4:
	transformA = v3
	worldTransformB = b2GetBodyTransform(tls, world, (*b2JointSim)(unsafe.Pointer(base1)).BodyIdB)
	v5 = worldTransformB
	v6 = worldTransformA.P
	r.Q = v5.Q
	r.P = Vec2{
		X: v5.P.X - v6.X,
		Y: v5.P.Y - v6.Y,
	}
	v7 = r
	goto _8
_8:
	transformB = v7
	v9 = transformA
	v10 = (*b2JointSim)(unsafe.Pointer(base1)).LocalFrameA.P
	x = float32(v9.Q.C*v10.X) - float32(v9.Q.S*v10.Y) + v9.P.X
	y = float32(v9.Q.S*v10.X) + float32(v9.Q.C*v10.Y) + v9.P.Y
	v11 = Vec2{
		X: x,
		Y: y,
	}
	goto _12
_12:
	pA = v11
	v13 = transformB
	v14 = (*b2JointSim)(unsafe.Pointer(base1)).LocalFrameB.P
	x = float32(v13.Q.C*v14.X) - float32(v13.Q.S*v14.Y) + v13.P.X
	y = float32(v13.Q.S*v14.X) + float32(v13.Q.C*v14.Y) + v13.P.Y
	v15 = Vec2{
		X: x,
		Y: y,
	}
	goto _16
_16:
	pB = v15
	v17 = pB
	v18 = pA
	v19 = Vec2{
		X: v17.X - v18.X,
		Y: v17.Y - v18.Y,
	}
	goto _20
_20:
	d = v19
	v21 = d
	lengthSquared = float32(v21.X*v21.X) + float32(v21.Y*v21.Y)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s1 = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec2{
			X: float32(s1 * v21.X),
			Y: float32(s1 * v21.Y),
		}
		v22 = u
		goto _23
	}
	v22 = Vec2{}
	goto _23
_23:
	axis = v22
	force = float32(((*b2DistanceJoint)(unsafe.Pointer(joint)).Impulse + (*b2DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b2DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse + (*b2DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse) * (*b2World)(unsafe.Pointer(world)).Inv_h)
	v24 = force
	v25 = axis
	v26 = Vec2{
		X: float32(v24 * v25.X),
		Y: float32(v24 * v25.Y),
	}
	goto _27
_27:
	return v26
}

// 1-D constrained system
// m (v2 - v1) = lambda
// v2 + (beta/h) * x1 + gamma * lambda = 0, gamma has units of inverse mass.
// x2 = x1 + h * v2

// 1-D mass-damper-spring system
// m (v2 - v1) + h * d * v2 + h * k *

// C = norm(p2 - p1) - L
// u = (p2 - p1) / norm(p2 - p1)
// Cdot = dot(u, v2 + cross(w2, r2) - v1 - cross(w1, r1))
// J = [-u -cross(r1, u) u cross(r2, u)]
// K = J * invM * JT
//   = invMass1 + invI1 * cross(r1, u)^2 + invMass2 + invI2 * cross(r2, u)^2

func b2PrepareDistanceJoint(tls *_Stack, base uintptr, context uintptr) {
	var a11, a21, a31, crA, crB, iA, iB, k, lengthSquared, mA, mB, omega, s, v44, v48, v50 float32
	var axis, rA, rB, separation, u, v11, v12, v13, v16, v17, v19, v20, v21, v24, v25, v29, v31 Vec2
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world uintptr
	var idA, idB, localIndexA, localIndexB, v9 int32
	var v1, v2, v3, v4 bool
	var v15 Rot
	var v27, v28 Pos
	var v53 b2Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a21, a31, axis, bodyA, bodyB, bodySimA, bodySimB, crA, crB, iA, iB, idA, idB, joint, k, lengthSquared, localIndexA, localIndexB, mA, mB, omega, rA, rB, s, separation, setA, setB, u, world, v1, v11, v12, v13, v15, v16, v17, v19, v2, v20, v21, v24, v25, v27, v28, v29, v3, v31, v4, v44, v48, v50, v53, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_distanceJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8476, __ccgo_ts+8197, int32FromInt32(281))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// chase body id to the solver set where the body lives
	idA = (*b2JointSim)(unsafe.Pointer(base)).BodyIdA
	idB = (*b2JointSim)(unsafe.Pointer(base)).BodyIdB
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	if v1 = !!(0 <= idA && idA < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8507, __ccgo_ts+8197, int32FromInt32(288))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idA)*104
	if v2 = !!(0 <= idB && idB < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+8557, __ccgo_ts+8197, int32FromInt32(289))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idB)*104
	if v3 = !!((*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet)); !v3 {
		b2InternalAssert(tls, __ccgo_ts+8607, __ccgo_ts+8197, int32FromInt32(291))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(bodyA)).SetIndex && (*b2Body)(unsafe.Pointer(bodyA)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8672, __ccgo_ts+8197, int32FromInt32(293))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	setA = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(bodyA)).SetIndex)*88
	if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(bodyB)).SetIndex && (*b2Body)(unsafe.Pointer(bodyB)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+8750, __ccgo_ts+8197, int32FromInt32(294))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	setB = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(bodyB)).SetIndex)*88
	localIndexA = (*b2Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b2Body)(unsafe.Pointer(bodyB)).LocalIndex
	if v3 = !!(0 <= localIndexA && localIndexA < (*b2SolverSet)(unsafe.Pointer(setA)).BodySims.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+8828, __ccgo_ts+8197, int32FromInt32(299))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	bodySimA = (*b2SolverSet)(unsafe.Pointer(setA)).BodySims.Data + uintptr(localIndexA)*96
	if v4 = !!(0 <= localIndexB && localIndexB < (*b2SolverSet)(unsafe.Pointer(setB)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+8895, __ccgo_ts+8197, int32FromInt32(300))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySimB = (*b2SolverSet)(unsafe.Pointer(setB)).BodySims.Data + uintptr(localIndexB)*96
	mA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvMass
	iA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvInertia
	mB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvMass
	iB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvInertia
	(*b2JointSim)(unsafe.Pointer(base)).InvMassA = mA
	(*b2JointSim)(unsafe.Pointer(base)).InvMassB = mB
	(*b2JointSim)(unsafe.Pointer(base)).InvIA = iA
	(*b2JointSim)(unsafe.Pointer(base)).InvIB = iB
	joint = base + 92
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) {
		v9 = localIndexA
	} else {
		v9 = -int32(1)
	}
	(*b2DistanceJoint)(unsafe.Pointer(joint)).IndexA = v9
	if (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) {
		v9 = localIndexB
	} else {
		v9 = -int32(1)
	}
	(*b2DistanceJoint)(unsafe.Pointer(joint)).IndexB = v9
	// initial anchors in world space
	v11 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA.P
	v12 = (*b2BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v13 = Vec2{
		X: v11.X - v12.X,
		Y: v11.Y - v12.Y,
	}
	goto _14
_14:
	v15 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v16 = v13
	v17 = Vec2{
		X: float32(v15.C*v16.X) - float32(v15.S*v16.Y),
		Y: float32(v15.S*v16.X) + float32(v15.C*v16.Y),
	}
	goto _18
_18:
	(*b2DistanceJoint)(unsafe.Pointer(joint)).AnchorA = v17
	v11 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v12 = (*b2BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v13 = Vec2{
		X: v11.X - v12.X,
		Y: v11.Y - v12.Y,
	}
	goto _22
_22:
	v15 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v16 = v13
	v17 = Vec2{
		X: float32(v15.C*v16.X) - float32(v15.S*v16.Y),
		Y: float32(v15.S*v16.X) + float32(v15.C*v16.Y),
	}
	goto _26
_26:
	(*b2DistanceJoint)(unsafe.Pointer(joint)).AnchorB = v17
	v27 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Center
	v28 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Center
	v11 = Vec2{
		X: v27.X - v28.X,
		Y: v27.Y - v28.Y,
	}
	goto _30
_30:
	(*b2DistanceJoint)(unsafe.Pointer(joint)).DeltaCenter = v11
	rA = (*b2DistanceJoint)(unsafe.Pointer(joint)).AnchorA
	rB = (*b2DistanceJoint)(unsafe.Pointer(joint)).AnchorB
	v11 = rB
	v12 = rA
	v13 = Vec2{
		X: v11.X - v12.X,
		Y: v11.Y - v12.Y,
	}
	goto _34
_34:
	v16 = v13
	v17 = (*b2DistanceJoint)(unsafe.Pointer(joint)).DeltaCenter
	v19 = Vec2{
		X: v16.X + v17.X,
		Y: v16.Y + v17.Y,
	}
	goto _38
_38:
	separation = v19
	v20 = separation
	lengthSquared = float32(v20.X*v20.X) + float32(v20.Y*v20.Y)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec2{
			X: float32(s * v20.X),
			Y: float32(s * v20.Y),
		}
		v21 = u
		goto _41
	}
	v21 = Vec2{}
	goto _41
_41:
	axis = v21
	v24 = rA
	v25 = axis
	v44 = float32(v24.X*v25.Y) - float32(v24.Y*v25.X)
	goto _45
_45:
	// compute effective mass
	crA = v44
	v29 = rB
	v31 = axis
	v48 = float32(v29.X*v31.Y) - float32(v29.Y*v31.X)
	goto _49
_49:
	crB = v48
	k = mA + mB + float32(float32(iA*crA)*crA) + float32(float32(iB*crB)*crB)
	if k > float32FromFloat32(0) {
		v50 = float32FromFloat32(1) / k
	} else {
		v50 = float32FromFloat32(0)
	}
	(*b2DistanceJoint)(unsafe.Pointer(joint)).AxialMass = v50
	v44 = (*b2DistanceJoint)(unsafe.Pointer(joint)).Hertz
	v48 = (*b2StepContext)(unsafe.Pointer(context)).H
	if v44 == float32FromFloat32(0) {
		v53 = b2Softness{}
		goto _54
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(_B2_PI)) * v44)
	a11 = float32(float32FromFloat32(2)*(*b2DistanceJoint)(unsafe.Pointer(joint)).DampingRatio) + float32(v48*omega)
	a21 = float32(float32(v48*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v53 = b2Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _54
_54:
	(*b2DistanceJoint)(unsafe.Pointer(joint)).DistanceSoftness = v53
	if int32FromUint8((*b2StepContext)(unsafe.Pointer(context)).EnableWarmStarting) == false1 {
		(*b2DistanceJoint)(unsafe.Pointer(joint)).Impulse = float32FromFloat32(0)
		(*b2DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse = float32FromFloat32(0)
		(*b2DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse = float32FromFloat32(0)
		(*b2DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse = float32FromFloat32(0)
	}
}

func b2WarmStartDistanceJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var P, axis, ds, rA, rB, separation, u, v10, v12, v13, v14, v16, v17, v18, v20, v21, v22, v24, v25, v26, v28, v29, v32, v33, v35, v37, v38, v5, v6, v9 Vec2
	var axialImpulse, iA, iB, lengthSquared, mA, mB, s3, v31, v36 float32
	var joint, stateA, stateB, v2, v3 uintptr
	var v1 bool
	var v4, v8 Rot
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = P, axialImpulse, axis, ds, iA, iB, joint, lengthSquared, mA, mB, rA, rB, s3, separation, stateA, stateB, u, v1, v10, v12, v13, v14, v16, v17, v18, v2, v20, v21, v22, v24, v25, v26, v28, v29, v3, v31, v32, v33, v35, v36, v37, v38, v4, v5, v6, v8, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_distanceJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8476, __ccgo_ts+8197, int32FromInt32(346))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Add(nil, bp)) = b2_identityBodyState6
	joint = base + 92
	if (*b2DistanceJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2DistanceJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v2
	if (*b2DistanceJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2DistanceJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v3
	v4 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b2DistanceJoint)(unsafe.Pointer(joint)).AnchorA
	v6 = Vec2{
		X: float32(v4.C*v5.X) - float32(v4.S*v5.Y),
		Y: float32(v4.S*v5.X) + float32(v4.C*v5.Y),
	}
	goto _7
_7:
	rA = v6
	v8 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v9 = (*b2DistanceJoint)(unsafe.Pointer(joint)).AnchorB
	v10 = Vec2{
		X: float32(v8.C*v9.X) - float32(v8.S*v9.Y),
		Y: float32(v8.S*v9.X) + float32(v8.C*v9.Y),
	}
	goto _11
_11:
	rB = v10
	v12 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
	v13 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaPosition
	v14 = Vec2{
		X: v12.X - v13.X,
		Y: v12.Y - v13.Y,
	}
	goto _15
_15:
	v16 = rB
	v17 = rA
	v18 = Vec2{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
	}
	goto _19
_19:
	v20 = v14
	v21 = v18
	v22 = Vec2{
		X: v20.X + v21.X,
		Y: v20.Y + v21.Y,
	}
	goto _23
_23:
	ds = v22
	v24 = (*b2DistanceJoint)(unsafe.Pointer(joint)).DeltaCenter
	v25 = ds
	v26 = Vec2{
		X: v24.X + v25.X,
		Y: v24.Y + v25.Y,
	}
	goto _27
_27:
	separation = v26
	v28 = separation
	lengthSquared = float32(v28.X*v28.X) + float32(v28.Y*v28.Y)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s3 = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec2{
			X: float32(s3 * v28.X),
			Y: float32(s3 * v28.Y),
		}
		v29 = u
		goto _30
	}
	v29 = Vec2{}
	goto _30
_30:
	axis = v29
	axialImpulse = (*b2DistanceJoint)(unsafe.Pointer(joint)).Impulse + (*b2DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b2DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse + (*b2DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse
	v31 = axialImpulse
	v32 = axis
	v33 = Vec2{
		X: float32(v31 * v32.X),
		Y: float32(v31 * v32.Y),
	}
	goto _34
_34:
	P = v33
	if (*b2BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b2_dynamicFlag) != 0 {
		v35 = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
		v36 = mA
		v37 = P
		v38 = Vec2{
			X: v35.X - float32(v36*v37.X),
			Y: v35.Y - float32(v36*v37.Y),
		}
		goto _39
	_39:
		(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = v38
		v5 = rA
		v6 = P
		v31 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		goto _43
	_43:
		*(*float32)(unsafe.Add(nil, stateA+8)) -= float32(iA * v31)
	}
	if (*b2BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b2_dynamicFlag) != 0 {
		v5 = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
		v31 = mB
		v6 = P
		v9 = Vec2{
			X: v5.X + float32(v31*v6.X),
			Y: v5.Y + float32(v31*v6.Y),
		}
		goto _48
	_48:
		(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = v9
		v5 = rB
		v6 = P
		v31 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		goto _52
	_52:
		*(*float32)(unsafe.Add(nil, stateB+8)) += float32(iB * v31)
	}
}

func b2SolveDistanceJoint(tls *_Stack, base uintptr, context uintptr, useBias uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var C, C1, Cdot, Cdot1, Cdot2, bias, bias1, bias2, bias3, c, c1, cdot, cdot1, h, iA, iB, impulse, impulse1, impulse2, impulse3, impulse4, impulseScale, impulseScale1, impulseScale2, length, lengthSquared, m, mA, mB, massScale, massScale1, massScale2, maxImpulse, newImpulse, newImpulse1, oldImpulse, oldImpulse1, s4, wA, wB, v29, v38, v42, v56, v58, v59, v60, v61, v63, v64 float32
	var P, P1, P2, P3, P4, axis, ds, rA, rB, separation, u, vA, vB, vr, vr1, vr2, vr3, vr4, v10, v12, v13, v14, v16, v17, v18, v20, v21, v22, v24, v25, v26, v28, v31, v32, v34, v35, v36, v39, v40, v43, v44, v46, v47, v48, v5, v50, v51, v52, v54, v55, v6, v9 Vec2
	var joint, stateA, stateB, v2, v3 uintptr
	var v1 bool
	var v4, v8 Rot
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, C1, Cdot, Cdot1, Cdot2, P, P1, P2, P3, P4, axis, bias, bias1, bias2, bias3, c, c1, cdot, cdot1, ds, h, iA, iB, impulse, impulse1, impulse2, impulse3, impulse4, impulseScale, impulseScale1, impulseScale2, joint, length, lengthSquared, m, mA, mB, massScale, massScale1, massScale2, maxImpulse, newImpulse, newImpulse1, oldImpulse, oldImpulse1, rA, rB, s4, separation, stateA, stateB, u, vA, vB, vr, vr1, vr2, vr3, vr4, wA, wB, v1, v10, v12, v13, v14, v16, v17, v18, v2, v20, v21, v22, v24, v25, v26, v28, v29, v3, v31, v32, v34, v35, v36, v38, v39, v4, v40, v42, v43, v44, v46, v47, v48, v5, v50, v51, v52, v54, v55, v56, v58, v59, v6, v60, v61, v63, v64, v8, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_distanceJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8476, __ccgo_ts+8197, int32FromInt32(385))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Add(nil, bp)) = b2_identityBodyState6
	joint = base + 92
	if (*b2DistanceJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2DistanceJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v2
	if (*b2DistanceJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2DistanceJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v3
	vA = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	v4 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b2DistanceJoint)(unsafe.Pointer(joint)).AnchorA
	v6 = Vec2{
		X: float32(v4.C*v5.X) - float32(v4.S*v5.Y),
		Y: float32(v4.S*v5.X) + float32(v4.C*v5.Y),
	}
	goto _7
_7:
	// current anchors
	rA = v6
	v8 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v9 = (*b2DistanceJoint)(unsafe.Pointer(joint)).AnchorB
	v10 = Vec2{
		X: float32(v8.C*v9.X) - float32(v8.S*v9.Y),
		Y: float32(v8.S*v9.X) + float32(v8.C*v9.Y),
	}
	goto _11
_11:
	rB = v10
	v12 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
	v13 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaPosition
	v14 = Vec2{
		X: v12.X - v13.X,
		Y: v12.Y - v13.Y,
	}
	goto _15
_15:
	v16 = rB
	v17 = rA
	v18 = Vec2{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
	}
	goto _19
_19:
	v20 = v14
	v21 = v18
	v22 = Vec2{
		X: v20.X + v21.X,
		Y: v20.Y + v21.Y,
	}
	goto _23
_23:
	// current separation
	ds = v22
	v24 = (*b2DistanceJoint)(unsafe.Pointer(joint)).DeltaCenter
	v25 = ds
	v26 = Vec2{
		X: v24.X + v25.X,
		Y: v24.Y + v25.Y,
	}
	goto _27
_27:
	separation = v26
	v28 = separation
	v29 = sqrtf(tls, float32(v28.X*v28.X)+float32(v28.Y*v28.Y))
	goto _30
_30:
	length = v29
	v31 = separation
	lengthSquared = float32(v31.X*v31.X) + float32(v31.Y*v31.Y)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s4 = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec2{
			X: float32(s4 * v31.X),
			Y: float32(s4 * v31.Y),
		}
		v32 = u
		goto _33
	}
	v32 = Vec2{}
	goto _33
_33:
	axis = v32
	// joint is soft if
	// - spring is enabled
	// - and (joint limit is disabled or limits are not equal)
	if (*b2DistanceJoint)(unsafe.Pointer(joint)).EnableSpring != 0 && ((*b2DistanceJoint)(unsafe.Pointer(joint)).MinLength < (*b2DistanceJoint)(unsafe.Pointer(joint)).MaxLength || int32FromUint8((*b2DistanceJoint)(unsafe.Pointer(joint)).EnableLimit) == false1) {
		// spring
		if (*b2DistanceJoint)(unsafe.Pointer(joint)).Hertz > float32FromFloat32(0) {
			v34 = vB
			v35 = vA
			v36 = Vec2{
				X: v34.X - v35.X,
				Y: v34.Y - v35.Y,
			}
			goto _37
		_37:
			v38 = wB
			v39 = rB
			v40 = Vec2{
				X: float32(-v38 * v39.Y),
				Y: float32(v38 * v39.X),
			}
			goto _41
		_41:
			v42 = wA
			v43 = rA
			v44 = Vec2{
				X: float32(-v42 * v43.Y),
				Y: float32(v42 * v43.X),
			}
			goto _45
		_45:
			v46 = v40
			v47 = v44
			v48 = Vec2{
				X: v46.X - v47.X,
				Y: v46.Y - v47.Y,
			}
			goto _49
		_49:
			v50 = v36
			v51 = v48
			v52 = Vec2{
				X: v50.X + v51.X,
				Y: v50.Y + v51.Y,
			}
			goto _53
		_53:
			// Cdot = dot(u, v + cross(w, r))
			vr = v52
			v54 = axis
			v55 = vr
			v56 = float32(v54.X*v55.X) + float32(v54.Y*v55.Y)
			goto _57
		_57:
			cdot = v56
			c = length - (*b2DistanceJoint)(unsafe.Pointer(joint)).Length
			bias = float32((*b2DistanceJoint)(unsafe.Pointer(joint)).DistanceSoftness.BiasRate * c)
			m = float32((*b2DistanceJoint)(unsafe.Pointer(joint)).DistanceSoftness.MassScale * (*b2DistanceJoint)(unsafe.Pointer(joint)).AxialMass)
			oldImpulse = (*b2DistanceJoint)(unsafe.Pointer(joint)).Impulse
			impulse = float32(-m*(cdot+bias)) - float32((*b2DistanceJoint)(unsafe.Pointer(joint)).DistanceSoftness.ImpulseScale*oldImpulse)
			h = (*b2StepContext)(unsafe.Pointer(context)).H
			v58 = (*b2DistanceJoint)(unsafe.Pointer(joint)).Impulse + impulse
			v59 = float32((*b2DistanceJoint)(unsafe.Pointer(joint)).LowerSpringForce * h)
			v60 = float32((*b2DistanceJoint)(unsafe.Pointer(joint)).UpperSpringForce * h)
			if v58 < v59 {
				v63 = v59
			} else {
				if v58 > v60 {
					v64 = v60
				} else {
					v64 = v58
				}
				v63 = v64
			}
			v61 = v63
			goto _62
		_62:
			(*b2DistanceJoint)(unsafe.Pointer(joint)).Impulse = v61
			impulse = (*b2DistanceJoint)(unsafe.Pointer(joint)).Impulse - oldImpulse
			v29 = impulse
			v5 = axis
			v6 = Vec2{
				X: float32(v29 * v5.X),
				Y: float32(v29 * v5.Y),
			}
			goto _68
		_68:
			P = v6
			v9 = vA
			v38 = mA
			v10 = P
			v12 = Vec2{
				X: v9.X - float32(v38*v10.X),
				Y: v9.Y - float32(v38*v10.Y),
			}
			goto _73
		_73:
			vA = v12
			v5 = rA
			v6 = P
			v29 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			goto _77
		_77:
			wA = wA - float32(iA*v29)
			v5 = vB
			v29 = mB
			v6 = P
			v9 = Vec2{
				X: v5.X + float32(v29*v6.X),
				Y: v5.Y + float32(v29*v6.Y),
			}
			goto _82
		_82:
			vB = v9
			v5 = rB
			v6 = P
			v29 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			goto _86
		_86:
			wB = wB + float32(iB*v29)
		}
		if (*b2DistanceJoint)(unsafe.Pointer(joint)).EnableMotor != 0 {
			v5 = vB
			v6 = vA
			v9 = Vec2{
				X: v5.X - v6.X,
				Y: v5.Y - v6.Y,
			}
			goto _90
		_90:
			v29 = wB
			v10 = rB
			v12 = Vec2{
				X: float32(-v29 * v10.Y),
				Y: float32(v29 * v10.X),
			}
			goto _94
		_94:
			v38 = wA
			v13 = rA
			v14 = Vec2{
				X: float32(-v38 * v13.Y),
				Y: float32(v38 * v13.X),
			}
			goto _98
		_98:
			v16 = v12
			v17 = v14
			v18 = Vec2{
				X: v16.X - v17.X,
				Y: v16.Y - v17.Y,
			}
			goto _102
		_102:
			v20 = v9
			v21 = v18
			v22 = Vec2{
				X: v20.X + v21.X,
				Y: v20.Y + v21.Y,
			}
			goto _106
		_106:
			vr1 = v22
			v24 = axis
			v25 = vr1
			v42 = float32(v24.X*v25.X) + float32(v24.Y*v25.Y)
			goto _110
		_110:
			Cdot = v42
			impulse1 = float32((*b2DistanceJoint)(unsafe.Pointer(joint)).AxialMass * ((*b2DistanceJoint)(unsafe.Pointer(joint)).MotorSpeed - Cdot))
			oldImpulse1 = (*b2DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse
			maxImpulse = float32((*b2StepContext)(unsafe.Pointer(context)).H * (*b2DistanceJoint)(unsafe.Pointer(joint)).MaxMotorForce)
			v56 = (*b2DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse + impulse1
			v58 = -maxImpulse
			v59 = maxImpulse
			if v56 < v58 {
				v61 = v58
			} else {
				if v56 > v59 {
					v63 = v59
				} else {
					v63 = v56
				}
				v61 = v63
			}
			v60 = v61
			goto _115
		_115:
			(*b2DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse = v60
			impulse1 = (*b2DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse - oldImpulse1
			v29 = impulse1
			v5 = axis
			v6 = Vec2{
				X: float32(v29 * v5.X),
				Y: float32(v29 * v5.Y),
			}
			goto _121
		_121:
			P1 = v6
			v9 = vA
			v38 = mA
			v10 = P1
			v12 = Vec2{
				X: v9.X - float32(v38*v10.X),
				Y: v9.Y - float32(v38*v10.Y),
			}
			goto _126
		_126:
			vA = v12
			v5 = rA
			v6 = P1
			v29 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			goto _130
		_130:
			wA = wA - float32(iA*v29)
			v5 = vB
			v29 = mB
			v6 = P1
			v9 = Vec2{
				X: v5.X + float32(v29*v6.X),
				Y: v5.Y + float32(v29*v6.Y),
			}
			goto _135
		_135:
			vB = v9
			v5 = rB
			v6 = P1
			v29 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			goto _139
		_139:
			wB = wB + float32(iB*v29)
		}
		if (*b2DistanceJoint)(unsafe.Pointer(joint)).EnableLimit != 0 {
			// lower limit
			v5 = vB
			v6 = vA
			v9 = Vec2{
				X: v5.X - v6.X,
				Y: v5.Y - v6.Y,
			}
			goto _143
		_143:
			v29 = wB
			v10 = rB
			v12 = Vec2{
				X: float32(-v29 * v10.Y),
				Y: float32(v29 * v10.X),
			}
			goto _147
		_147:
			v38 = wA
			v13 = rA
			v14 = Vec2{
				X: float32(-v38 * v13.Y),
				Y: float32(v38 * v13.X),
			}
			goto _151
		_151:
			v16 = v12
			v17 = v14
			v18 = Vec2{
				X: v16.X - v17.X,
				Y: v16.Y - v17.Y,
			}
			goto _155
		_155:
			v20 = v9
			v21 = v18
			v22 = Vec2{
				X: v20.X + v21.X,
				Y: v20.Y + v21.Y,
			}
			goto _159
		_159:
			vr2 = v22
			v24 = axis
			v25 = vr2
			v42 = float32(v24.X*v25.X) + float32(v24.Y*v25.Y)
			goto _163
		_163:
			cdot1 = v42
			c1 = length - (*b2DistanceJoint)(unsafe.Pointer(joint)).MinLength
			bias1 = float32FromFloat32(0)
			massScale = float32FromFloat32(1)
			impulseScale = float32FromFloat32(0)
			if c1 > float32FromFloat32(0) {
				// speculative
				bias1 = float32(c1 * (*b2StepContext)(unsafe.Pointer(context)).Inv_h)
			} else {
				if useBias != 0 {
					bias1 = float32((*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * c1)
					massScale = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
					impulseScale = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
				}
			}
			impulse2 = float32(float32(-massScale*(*b2DistanceJoint)(unsafe.Pointer(joint)).AxialMass)*(cdot1+bias1)) - float32(impulseScale*(*b2DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse)
			v29 = float32FromFloat32(0)
			v38 = (*b2DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse + impulse2
			if v29 > v38 {
				v56 = v29
			} else {
				v56 = v38
			}
			v42 = v56
			goto _167
		_167:
			newImpulse = v42
			impulse2 = newImpulse - (*b2DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse
			(*b2DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse = newImpulse
			v29 = impulse2
			v5 = axis
			v6 = Vec2{
				X: float32(v29 * v5.X),
				Y: float32(v29 * v5.Y),
			}
			goto _172
		_172:
			P2 = v6
			v9 = vA
			v38 = mA
			v10 = P2
			v12 = Vec2{
				X: v9.X - float32(v38*v10.X),
				Y: v9.Y - float32(v38*v10.Y),
			}
			goto _177
		_177:
			vA = v12
			v5 = rA
			v6 = P2
			v29 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			goto _181
		_181:
			wA = wA - float32(iA*v29)
			v5 = vB
			v29 = mB
			v6 = P2
			v9 = Vec2{
				X: v5.X + float32(v29*v6.X),
				Y: v5.Y + float32(v29*v6.Y),
			}
			goto _186
		_186:
			vB = v9
			v5 = rB
			v6 = P2
			v29 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			goto _190
		_190:
			wB = wB + float32(iB*v29)
			// upper
			v5 = vA
			v6 = vB
			v9 = Vec2{
				X: v5.X - v6.X,
				Y: v5.Y - v6.Y,
			}
			goto _194
		_194:
			v29 = wA
			v10 = rA
			v12 = Vec2{
				X: float32(-v29 * v10.Y),
				Y: float32(v29 * v10.X),
			}
			goto _198
		_198:
			v38 = wB
			v13 = rB
			v14 = Vec2{
				X: float32(-v38 * v13.Y),
				Y: float32(v38 * v13.X),
			}
			goto _202
		_202:
			v16 = v12
			v17 = v14
			v18 = Vec2{
				X: v16.X - v17.X,
				Y: v16.Y - v17.Y,
			}
			goto _206
		_206:
			v20 = v9
			v21 = v18
			v22 = Vec2{
				X: v20.X + v21.X,
				Y: v20.Y + v21.Y,
			}
			goto _210
		_210:
			vr3 = v22
			v24 = axis
			v25 = vr3
			v42 = float32(v24.X*v25.X) + float32(v24.Y*v25.Y)
			goto _214
		_214:
			Cdot1 = v42
			C = (*b2DistanceJoint)(unsafe.Pointer(joint)).MaxLength - length
			bias2 = float32FromFloat32(0)
			massScale1 = float32FromFloat32(1)
			impulseScale1 = float32FromFloat32(0)
			if C > float32FromFloat32(0) {
				// speculative
				bias2 = float32(C * (*b2StepContext)(unsafe.Pointer(context)).Inv_h)
			} else {
				if useBias != 0 {
					bias2 = float32((*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * C)
					massScale1 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
					impulseScale1 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
				}
			}
			impulse3 = float32(float32(-massScale1*(*b2DistanceJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot1+bias2)) - float32(impulseScale1*(*b2DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse)
			v29 = float32FromFloat32(0)
			v38 = (*b2DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse + impulse3
			if v29 > v38 {
				v56 = v29
			} else {
				v56 = v38
			}
			v42 = v56
			goto _218
		_218:
			newImpulse1 = v42
			impulse3 = newImpulse1 - (*b2DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse
			(*b2DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse = newImpulse1
			v29 = -impulse3
			v5 = axis
			v6 = Vec2{
				X: float32(v29 * v5.X),
				Y: float32(v29 * v5.Y),
			}
			goto _223
		_223:
			P3 = v6
			v9 = vA
			v38 = mA
			v10 = P3
			v12 = Vec2{
				X: v9.X - float32(v38*v10.X),
				Y: v9.Y - float32(v38*v10.Y),
			}
			goto _228
		_228:
			vA = v12
			v5 = rA
			v6 = P3
			v29 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			goto _232
		_232:
			wA = wA - float32(iA*v29)
			v5 = vB
			v29 = mB
			v6 = P3
			v9 = Vec2{
				X: v5.X + float32(v29*v6.X),
				Y: v5.Y + float32(v29*v6.Y),
			}
			goto _237
		_237:
			vB = v9
			v5 = rB
			v6 = P3
			v29 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
			goto _241
		_241:
			wB = wB + float32(iB*v29)
		}
	} else {
		v5 = vB
		v6 = vA
		v9 = Vec2{
			X: v5.X - v6.X,
			Y: v5.Y - v6.Y,
		}
		goto _245
	_245:
		v29 = wB
		v10 = rB
		v12 = Vec2{
			X: float32(-v29 * v10.Y),
			Y: float32(v29 * v10.X),
		}
		goto _249
	_249:
		v38 = wA
		v13 = rA
		v14 = Vec2{
			X: float32(-v38 * v13.Y),
			Y: float32(v38 * v13.X),
		}
		goto _253
	_253:
		v16 = v12
		v17 = v14
		v18 = Vec2{
			X: v16.X - v17.X,
			Y: v16.Y - v17.Y,
		}
		goto _257
	_257:
		v20 = v9
		v21 = v18
		v22 = Vec2{
			X: v20.X + v21.X,
			Y: v20.Y + v21.Y,
		}
		goto _261
	_261:
		// rigid constraint
		vr4 = v22
		v24 = axis
		v25 = vr4
		v42 = float32(v24.X*v25.X) + float32(v24.Y*v25.Y)
		goto _265
	_265:
		Cdot2 = v42
		C1 = length - (*b2DistanceJoint)(unsafe.Pointer(joint)).Length
		bias3 = float32FromFloat32(0)
		massScale2 = float32FromFloat32(1)
		impulseScale2 = float32FromFloat32(0)
		if useBias != 0 {
			bias3 = float32((*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * C1)
			massScale2 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
			impulseScale2 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
		}
		impulse4 = float32(float32(-massScale2*(*b2DistanceJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot2+bias3)) - float32(impulseScale2*(*b2DistanceJoint)(unsafe.Pointer(joint)).Impulse)
		*(*float32)(unsafe.Add(nil, joint+36)) += impulse4
		v29 = impulse4
		v5 = axis
		v6 = Vec2{
			X: float32(v29 * v5.X),
			Y: float32(v29 * v5.Y),
		}
		goto _269
	_269:
		P4 = v6
		v9 = vA
		v38 = mA
		v10 = P4
		v12 = Vec2{
			X: v9.X - float32(v38*v10.X),
			Y: v9.Y - float32(v38*v10.Y),
		}
		goto _274
	_274:
		vA = v12
		v5 = rA
		v6 = P4
		v29 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		goto _278
	_278:
		wA = wA - float32(iA*v29)
		v5 = vB
		v29 = mB
		v6 = P4
		v9 = Vec2{
			X: v5.X + float32(v29*v6.X),
			Y: v5.Y + float32(v29*v6.Y),
		}
		goto _283
	_283:
		vB = v9
		v5 = rB
		v6 = P4
		v29 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		goto _287
	_287:
		wB = wB + float32(iB*v29)
	}
	if (*b2BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b2_dynamicFlag) != 0 {
		(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
		(*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	}
	if (*b2BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b2_dynamicFlag) != 0 {
		(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
		(*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
	}
}

func b2DrawDistanceJoint(tls *_Stack, draw uintptr, base uintptr, transformA WorldTransform, transformB WorldTransform) {
	var axis, offset, u, v12, v14, v15, v18, v19, v22, v26, v27, v3, v30, v33, v34, v37, v38, v40, v41, v44, v48, v7 Vec2
	var joint uintptr
	var lengthSquared, rx, ry, s1, v17, v25, v36 float32
	var pA, pB, pMax, pMin, pRest, v10, v11, v21, v23, v29, v31, v4, v43, v45, v47, v49, v8 Pos
	var v1 bool
	var v2, v6 WorldTransform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axis, joint, lengthSquared, offset, pA, pB, pMax, pMin, pRest, rx, ry, s1, u, v1, v10, v11, v12, v14, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v29, v3, v30, v31, v33, v34, v36, v37, v38, v4, v40, v41, v43, v44, v45, v47, v48, v49, v6, v7, v8
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_distanceJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8476, __ccgo_ts+8197, int32FromInt32(575))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	joint = base + 92
	v2 = transformA
	v3 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA.P
	rx = float32(v2.Q.C*v3.X) - float32(v2.Q.S*v3.Y)
	ry = float32(v2.Q.S*v3.X) + float32(v2.Q.C*v3.Y)
	v4 = Pos{
		X: v2.P.X + rx,
		Y: v2.P.Y + ry,
	}
	goto _5
_5:
	pA = v4
	v6 = transformB
	v7 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	rx = float32(v6.Q.C*v7.X) - float32(v6.Q.S*v7.Y)
	ry = float32(v6.Q.S*v7.X) + float32(v6.Q.C*v7.Y)
	v8 = Pos{
		X: v6.P.X + rx,
		Y: v6.P.Y + ry,
	}
	goto _9
_9:
	pB = v8
	v10 = pB
	v11 = pA
	v12 = Vec2{
		X: v10.X - v11.X,
		Y: v10.Y - v11.Y,
	}
	goto _13
_13:
	v14 = v12
	lengthSquared = float32(v14.X*v14.X) + float32(v14.Y*v14.Y)
	if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s1 = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
		u = Vec2{
			X: float32(s1 * v14.X),
			Y: float32(s1 * v14.Y),
		}
		v15 = u
		goto _16
	}
	v15 = Vec2{}
	goto _16
_16:
	axis = v15
	if (*b2DistanceJoint)(unsafe.Pointer(joint)).MinLength < (*b2DistanceJoint)(unsafe.Pointer(joint)).MaxLength && (*b2DistanceJoint)(unsafe.Pointer(joint)).EnableLimit != 0 {
		v17 = (*b2DistanceJoint)(unsafe.Pointer(joint)).MinLength
		v18 = axis
		v19 = Vec2{
			X: float32(v17 * v18.X),
			Y: float32(v17 * v18.Y),
		}
		goto _20
	_20:
		v21 = pA
		v22 = v19
		v23 = Pos{
			X: v21.X + v22.X,
			Y: v21.Y + v22.Y,
		}
		goto _24
	_24:
		pMin = v23
		v25 = (*b2DistanceJoint)(unsafe.Pointer(joint)).MaxLength
		v26 = axis
		v27 = Vec2{
			X: float32(v25 * v26.X),
			Y: float32(v25 * v26.Y),
		}
		goto _28
	_28:
		v29 = pA
		v30 = v27
		v31 = Pos{
			X: v29.X + v30.X,
			Y: v29.Y + v30.Y,
		}
		goto _32
	_32:
		pMax = v31
		v33 = axis
		v34 = Vec2{
			X: v33.Y,
			Y: -v33.X,
		}
		goto _35
	_35:
		v36 = float32(float32FromFloat32(0.05) * b2GetLengthUnitsPerMeter(tls))
		v37 = v34
		v38 = Vec2{
			X: float32(v36 * v37.X),
			Y: float32(v36 * v37.Y),
		}
		goto _39
	_39:
		offset = v38
		if (*b2DistanceJoint)(unsafe.Pointer(joint)).MinLength > float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter) {
			v40 = offset
			v41 = Vec2{
				X: -v40.X,
				Y: -v40.Y,
			}
			goto _42
		_42:
			v43 = pMin
			v44 = v41
			v45 = Pos{
				X: v43.X + v44.X,
				Y: v43.Y + v44.Y,
			}
			goto _46
		_46:
			v47 = pMin
			v48 = offset
			v49 = Pos{
				X: v47.X + v48.X,
				Y: v47.Y + v48.Y,
			}
			goto _50
		_50:
			(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, v45, v49, int32(b2_colorLightGreen), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		}
		if (*b2DistanceJoint)(unsafe.Pointer(joint)).MaxLength < float32(float32(float32FromFloat32(100)*float32FromFloat32(1000))*b2_lengthUnitsPerMeter) {
			v3 = offset
			v7 = Vec2{
				X: -v3.X,
				Y: -v3.Y,
			}
			goto _53
		_53:
			v4 = pMax
			v12 = v7
			v8 = Pos{
				X: v4.X + v12.X,
				Y: v4.Y + v12.Y,
			}
			goto _57
		_57:
			v10 = pMax
			v14 = offset
			v11 = Pos{
				X: v10.X + v14.X,
				Y: v10.Y + v14.Y,
			}
			goto _61
		_61:
			(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, v8, v11, int32(b2_colorRed), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		}
		if (*b2DistanceJoint)(unsafe.Pointer(joint)).MinLength > float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter) && (*b2DistanceJoint)(unsafe.Pointer(joint)).MaxLength < float32(float32(float32FromFloat32(100)*float32FromFloat32(1000))*b2_lengthUnitsPerMeter) {
			(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, pMin, pMax, int32(b2_colorGray), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		}
	}
	(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, pA, pB, int32(b2_colorWhite), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pA, float32FromFloat32(4), int32(b2_colorWhite), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pB, float32FromFloat32(4), int32(b2_colorWhite), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	if (*b2DistanceJoint)(unsafe.Pointer(joint)).Hertz > float32FromFloat32(0) && (*b2DistanceJoint)(unsafe.Pointer(joint)).EnableSpring != 0 {
		v17 = (*b2DistanceJoint)(unsafe.Pointer(joint)).Length
		v3 = axis
		v7 = Vec2{
			X: float32(v17 * v3.X),
			Y: float32(v17 * v3.Y),
		}
		goto _65
	_65:
		v4 = pA
		v12 = v7
		v8 = Pos{
			X: v4.X + v12.X,
			Y: v4.Y + v12.Y,
		}
		goto _69
	_69:
		pRest = v8
		(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pRest, float32FromFloat32(4), int32(b2_colorBlue), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	}
}

func b2AddJointToIsland(tls *_Stack, world uintptr, islandId int32, joint uintptr) {
	var elementSize, newCapacity, newSize, oldSize, v4 int32
	var island uintptr
	var link b2JointLink
	var v1 bool
	_, _, _, _, _, _, _, _ = elementSize, island, link, newCapacity, newSize, oldSize, v1, v4
	if v1 = !!((*b2Joint)(unsafe.Pointer(joint)).IslandId == -int32FromInt32(1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+4382, __ccgo_ts+10214, int32FromInt32(266))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*b2Joint)(unsafe.Pointer(joint)).IslandIndex == -int32FromInt32(1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+11331, __ccgo_ts+10214, int32FromInt32(267))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= islandId && islandId < (*b2World)(unsafe.Pointer(world)).Islands.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+962, __ccgo_ts+10214, int32FromInt32(269))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	island = (*b2World)(unsafe.Pointer(world)).Islands.Data + uintptr(islandId)*64
	(*b2Joint)(unsafe.Pointer(joint)).IslandId = islandId
	(*b2Joint)(unsafe.Pointer(joint)).IslandIndex = (*b2Island)(unsafe.Pointer(island)).Joints.Count
	link.JointId = (*b2Joint)(unsafe.Pointer(joint)).JointId
	link.BodyIdA = (*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId
	link.BodyIdB = (*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId
	elementSize = int32(12)
	if (*b2Island)(unsafe.Pointer(island)).Joints.Count >= (*b2Island)(unsafe.Pointer(island)).Joints.Capacity {
		oldSize = (*b2Island)(unsafe.Pointer(island)).Joints.Capacity * elementSize
		if (*b2Island)(unsafe.Pointer(island)).Joints.Capacity == 0 {
			v4 = int32(8)
		} else {
			v4 = int32(2) * (*b2Island)(unsafe.Pointer(island)).Joints.Capacity
		}
		newCapacity = v4
		newSize = newCapacity * elementSize
		(*b2Island)(unsafe.Pointer(island)).Joints.Data = b2GrowAlloc(tls, (*b2Island)(unsafe.Pointer(island)).Joints.Data, uint64FromInt32(oldSize), uint64FromInt32(newSize))
		(*b2Island)(unsafe.Pointer(island)).Joints.Capacity = newCapacity
	}
	*(*b2JointLink)(unsafe.Add(nil, (*b2Island)(unsafe.Pointer(island)).Joints.Data+uintptr((*b2Island)(unsafe.Pointer(island)).Joints.Count)*12)) = link
	(*b2Island)(unsafe.Pointer(island)).Joints.Count += int32(1)
	b2ValidateIsland(tls, world, islandId)
}

func b2LinkJoint(tls *_Stack, world uintptr, joint uintptr) {
	var bodyA, bodyB uintptr
	var finalIslandId, islandIdA, islandIdB int32
	var v1, v2, v3 bool
	_, _, _, _, _, _, _, _ = bodyA, bodyB, finalIslandId, islandIdA, islandIdB, v1, v2, v3
	if v1 = !!(0 <= (*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId && (*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+3036, __ccgo_ts+10214, int32FromInt32(285))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId)*104
	if v2 = !!(0 <= (*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId && (*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+3124, __ccgo_ts+10214, int32FromInt32(286))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId)*104
	if v3 = !!((*b2Body)(unsafe.Pointer(bodyA)).Type1 == int32(b2_dynamicBody) || (*b2Body)(unsafe.Pointer(bodyB)).Type1 == int32(b2_dynamicBody)); !v3 {
		b2InternalAssert(tls, __ccgo_ts+11367, __ccgo_ts+10214, int32FromInt32(288))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) && (*b2Body)(unsafe.Pointer(bodyB)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeSolverSet(tls, world, (*b2Body)(unsafe.Pointer(bodyB)).SetIndex)
	} else {
		if (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) && (*b2Body)(unsafe.Pointer(bodyA)).SetIndex >= int32(b2_firstSleepingSet) {
			b2WakeSolverSet(tls, world, (*b2Body)(unsafe.Pointer(bodyA)).SetIndex)
		}
	}
	islandIdA = (*b2Body)(unsafe.Pointer(bodyA)).IslandId
	islandIdB = (*b2Body)(unsafe.Pointer(bodyB)).IslandId
	if v1 = !!(islandIdA != -int32(1) || islandIdB != -int32(1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+10989, __ccgo_ts+10214, int32FromInt32(302))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// Merge islands. This will destroy one of the islands.
	finalIslandId = b2MergeIslands(tls, world, islandIdA, islandIdB)
	// Add joint the island that survived
	b2AddJointToIsland(tls, world, finalIslandId, joint)
}

func b2UnlinkJoint(tls *_Stack, world uintptr, joint uintptr) {
	var island, movedJoint, movedLink, v4, v5 uintptr
	var islandId, movedIndex, removeIndex, v6, v7, v9 int32
	var v1, v2, v3 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = island, islandId, movedIndex, movedJoint, movedLink, removeIndex, v1, v2, v3, v4, v5, v6, v7, v9
	if (*b2Joint)(unsafe.Pointer(joint)).IslandId == -int32(1) {
		return
	}
	// remove from island
	islandId = (*b2Joint)(unsafe.Pointer(joint)).IslandId
	if v1 = !!(0 <= islandId && islandId < (*b2World)(unsafe.Pointer(world)).Islands.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+962, __ccgo_ts+10214, int32FromInt32(320))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	island = (*b2World)(unsafe.Pointer(world)).Islands.Data + uintptr(islandId)*64
	removeIndex = (*b2Joint)(unsafe.Pointer(joint)).IslandIndex
	if v2 = !!(0 <= removeIndex && removeIndex < (*b2Island)(unsafe.Pointer(island)).Joints.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+11430, __ccgo_ts+10214, int32FromInt32(323))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*(*b2JointLink)(unsafe.Add(nil, (*b2Island)(unsafe.Pointer(island)).Joints.Data+uintptr(removeIndex)*12))).JointId == (*b2Joint)(unsafe.Pointer(joint)).JointId); !v1 {
		b2InternalAssert(tls, __ccgo_ts+11485, __ccgo_ts+10214, int32FromInt32(324))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	v4 = (*b2Island)(unsafe.Pointer(island)).Joints.Data
	v5 = island + 48 + 8
	v6 = removeIndex
	v7 = int32(12)
	if v1 = !!(0 <= v6 && v6 < *(*int32)(unsafe.Add(nil, v5)) && bool(__ccgo_ts+2029 != 0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+2055, __ccgo_ts+2115, int32FromInt32(153))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*int32)(unsafe.Add(nil, v5)) = *(*int32)(unsafe.Add(nil, v5)) - 1
	if v6 != *(*int32)(unsafe.Add(nil, v5)) {
		memcpy(tls, v4+uintptr(v6*v7), v4+uintptr(*(*int32)(unsafe.Add(nil, v5))*v7), uint64FromInt32(v7))
		v9 = *(*int32)(unsafe.Add(nil, v5))
		goto _10
	}
	v9 = -int32(1)
	goto _10
_10:
	movedIndex = v9
	if movedIndex != -int32(1) {
		// Fix islandIndex on the joint that was swapped into removeIndex
		movedLink = (*b2Island)(unsafe.Pointer(island)).Joints.Data + uintptr(removeIndex)*12
		if v2 = !!(0 <= (*b2JointLink)(unsafe.Pointer(movedLink)).JointId && (*b2JointLink)(unsafe.Pointer(movedLink)).JointId < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v2 {
			b2InternalAssert(tls, __ccgo_ts+11544, __ccgo_ts+10214, int32FromInt32(331))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		movedJoint = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr((*b2JointLink)(unsafe.Pointer(movedLink)).JointId)*72
		if v3 = !!((*b2Joint)(unsafe.Pointer(movedJoint)).IslandIndex == movedIndex); !v3 {
			b2InternalAssert(tls, __ccgo_ts+11624, __ccgo_ts+10214, int32FromInt32(332))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		(*b2Joint)(unsafe.Pointer(movedJoint)).IslandIndex = removeIndex
	}
	(*b2Joint)(unsafe.Pointer(joint)).IslandId = -int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).IslandIndex = -int32(1)
	*(*int32)(unsafe.Add(nil, island+12)) += int32(1)
	b2ValidateIsland(tls, world, islandId)
}

func b2DefaultJointDef(tls *_Stack) (r JointDef) {
	var def JointDef
	_ = def
	def = JointDef{}
	def.LocalFrameA.Q = b2Rot_identity
	def.LocalFrameB.Q = b2Rot_identity
	def.ForceThreshold = float32FromFloat32(3.4028234663852886e+38)
	def.TorqueThreshold = float32FromFloat32(3.4028234663852886e+38)
	def.ConstraintHertz = float32FromFloat32(60)
	def.ConstraintDampingRatio = float32FromFloat32(2)
	def.DrawScale = b2GetLengthUnitsPerMeter(tls)
	return def
}

func b2DefaultDistanceJointDef(tls *_Stack) (r DistanceJointDef) {
	var def DistanceJointDef
	_ = def
	def = DistanceJointDef{}
	def.Base = b2DefaultJointDef(tls)
	def.LowerSpringForce = -float32FromFloat32(3.4028234663852886e+38)
	def.UpperSpringForce = float32FromFloat32(3.4028234663852886e+38)
	def.Length = float32FromFloat32(1)
	def.MaxLength = float32(float32(float32FromFloat32(100)*float32FromFloat32(1000)) * b2_lengthUnitsPerMeter)
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}

func b2DefaultFilterJointDef(tls *_Stack) (r FilterJointDef) {
	var def FilterJointDef
	_ = def
	def = FilterJointDef{}
	def.Base = b2DefaultJointDef(tls)
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}

func b2DefaultMotorJointDef(tls *_Stack) (r MotorJointDef) {
	var def MotorJointDef
	_ = def
	def = MotorJointDef{}
	def.Base = b2DefaultJointDef(tls)
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}

func b2DefaultMoverJointDef(tls *_Stack) (r MoverJointDef) {
	var def MoverJointDef
	_ = def
	def = MoverJointDef{}
	def.Base = b2DefaultJointDef(tls)
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}

func b2DefaultPogoJointDef(tls *_Stack) (r PogoJointDef) {
	var def PogoJointDef
	_ = def
	def = PogoJointDef{}
	def.Base = b2DefaultJointDef(tls)
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}

func b2DefaultPrismaticJointDef(tls *_Stack) (r PrismaticJointDef) {
	var def PrismaticJointDef
	_ = def
	def = PrismaticJointDef{}
	def.Base = b2DefaultJointDef(tls)
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}

func b2DefaultRevoluteJointDef(tls *_Stack) (r RevoluteJointDef) {
	var def RevoluteJointDef
	_ = def
	def = RevoluteJointDef{}
	def.Base = b2DefaultJointDef(tls)
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}

func b2DefaultWeldJointDef(tls *_Stack) (r WeldJointDef) {
	var def WeldJointDef
	_ = def
	def = WeldJointDef{}
	def.Base = b2DefaultJointDef(tls)
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}

func b2DefaultWheelJointDef(tls *_Stack) (r WheelJointDef) {
	var def WheelJointDef
	_ = def
	def = WheelJointDef{}
	def.Base = b2DefaultJointDef(tls)
	def.EnableSpring = boolUint8(true1 != 0)
	def.Hertz = float32FromFloat32(1)
	def.DampingRatio = float32FromFloat32(0.7)
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}

func b2GetJointFullId(tls *_Stack, world uintptr, jointId JointId) (r uintptr) {
	var id int32
	var joint uintptr
	var v1, v2 bool
	_, _, _, _ = id, joint, v1, v2
	id = jointId.Index1 - int32(1)
	if v1 = !!(0 <= id && id < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+12273, __ccgo_ts+12321, int32FromInt32(125))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	joint = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr(id)*72
	if v2 = !!((*b2Joint)(unsafe.Pointer(joint)).JointId == id && int32FromUint16((*b2Joint)(unsafe.Pointer(joint)).Generation) == int32FromUint16(jointId.Generation)); !v2 {
		b2InternalAssert(tls, __ccgo_ts+12344, __ccgo_ts+12321, int32FromInt32(126))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	return joint
}

func b2GetJointSim(tls *_Stack, world uintptr, joint uintptr) (r uintptr) {
	var color, set uintptr
	var v1, v2 bool
	_, _, _, _ = color, set, v1, v2
	if (*b2Joint)(unsafe.Pointer(joint)).SetIndex == int32(b2_awakeSet) {
		if v1 = !!(0 <= (*b2Joint)(unsafe.Pointer(joint)).ColorIndex && (*b2Joint)(unsafe.Pointer(joint)).ColorIndex < int32(_B2_GRAPH_COLOR_COUNT)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+12408, __ccgo_ts+12321, int32FromInt32(134))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		color = world + 376 + uintptr((*b2Joint)(unsafe.Pointer(joint)).ColorIndex)*64
		if v1 = !!(0 <= (*b2Joint)(unsafe.Pointer(joint)).LocalIndex && (*b2Joint)(unsafe.Pointer(joint)).LocalIndex < (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+12475, __ccgo_ts+12321, int32FromInt32(136))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Data + uintptr((*b2Joint)(unsafe.Pointer(joint)).LocalIndex)*252
	}
	if v1 = !!(0 <= (*b2Joint)(unsafe.Pointer(joint)).SetIndex && (*b2Joint)(unsafe.Pointer(joint)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+3212, __ccgo_ts+12321, int32FromInt32(139))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Joint)(unsafe.Pointer(joint)).SetIndex)*88
	if v2 = !!(0 <= (*b2Joint)(unsafe.Pointer(joint)).LocalIndex && (*b2Joint)(unsafe.Pointer(joint)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).JointSims.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+12556, __ccgo_ts+12321, int32FromInt32(140))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	return (*b2SolverSet)(unsafe.Pointer(set)).JointSims.Data + uintptr((*b2Joint)(unsafe.Pointer(joint)).LocalIndex)*252
}

func b2GetJointSimCheckType(tls *_Stack, jointId JointId, type1 JointType) (r uintptr) {
	var joint, jointSim, world uintptr
	var v1 bool
	_, _, _, _ = joint, jointSim, world, v1
	_ = uint64FromInt64(4)
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+12321, int32FromInt32(148))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return uintptrFromInt32(0)
	}
	joint = b2GetJointFullId(tls, world, jointId)
	if v1 = !!((*b2Joint)(unsafe.Pointer(joint)).Type1 == type1); !v1 {
		b2InternalAssert(tls, __ccgo_ts+12635, __ccgo_ts+12321, int32FromInt32(155))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	jointSim = b2GetJointSim(tls, world, joint)
	if v1 = !!((*b2JointSim)(unsafe.Pointer(jointSim)).Type1 == type1); !v1 {
		b2InternalAssert(tls, __ccgo_ts+12655, __ccgo_ts+12321, int32FromInt32(157))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return jointSim
}

func b2IsValidJointDef(tls *_Stack, def uintptr) (r uint8) {
	return boolUint8(b2IsValidTransform(tls, (*JointDef)(unsafe.Pointer(def)).LocalFrameA) != 0 && b2IsValidTransform(tls, (*JointDef)(unsafe.Pointer(def)).LocalFrameB) != 0 && b2IsValidFloat(tls, (*JointDef)(unsafe.Pointer(def)).ForceThreshold) != 0 && (*JointDef)(unsafe.Pointer(def)).ForceThreshold >= float32FromFloat32(0) && b2IsValidFloat(tls, (*JointDef)(unsafe.Pointer(def)).TorqueThreshold) != 0 && (*JointDef)(unsafe.Pointer(def)).TorqueThreshold >= float32FromFloat32(0) && b2IsValidFloat(tls, (*JointDef)(unsafe.Pointer(def)).ConstraintHertz) != 0 && (*JointDef)(unsafe.Pointer(def)).ConstraintHertz >= float32FromFloat32(0) && b2IsValidFloat(tls, (*JointDef)(unsafe.Pointer(def)).ConstraintDampingRatio) != 0 && (*JointDef)(unsafe.Pointer(def)).ConstraintDampingRatio >= float32FromFloat32(0) && b2IsValidFloat(tls, (*JointDef)(unsafe.Pointer(def)).DrawScale) != 0)
}

func b2CreateJoint(tls *_Stack, world uintptr, def uintptr, type1 JointType) (r b2JointPair) {
	var bodyA, bodyB, edgeA, edgeB, joint, jointA, jointB, jointSim, mergedSet, set, set1, set2, v13, v17, v18, v19, v22 uintptr
	var bodyIdA, bodyIdB, elementSize1, jointId, keyA, keyB, maxSetIndex, newCapacity, newCapacity1, newSize, newSize1, oldCapacity, oldSize, oldSize1, setIndex, v10, v11, v6, v7, v8 int32
	var v1, v2 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, bodyIdA, bodyIdB, edgeA, edgeB, elementSize1, joint, jointA, jointB, jointId, jointSim, keyA, keyB, maxSetIndex, mergedSet, newCapacity, newCapacity1, newSize, newSize1, oldCapacity, oldSize, oldSize1, set, set1, set2, setIndex, v1, v10, v11, v13, v17, v18, v19, v2, v22, v6, v7, v8
	if v1 = !!(b2IsValidTransform(tls, (*JointDef)(unsafe.Pointer(def)).LocalFrameA) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+12678, __ccgo_ts+12321, int32FromInt32(215))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b2IsValidTransform(tls, (*JointDef)(unsafe.Pointer(def)).LocalFrameB) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+12717, __ccgo_ts+12321, int32FromInt32(216))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(int32FromUint16((*b2World)(unsafe.Pointer(world)).WorldId) == int32FromUint16((*JointDef)(unsafe.Pointer(def)).BodyIdA.World0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+12756, __ccgo_ts+12321, int32FromInt32(217))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(int32FromUint16((*b2World)(unsafe.Pointer(world)).WorldId) == int32FromUint16((*JointDef)(unsafe.Pointer(def)).BodyIdB.World0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+12794, __ccgo_ts+12321, int32FromInt32(218))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(boolInt32((*JointDef)(unsafe.Pointer(def)).BodyIdA.Index1 == (*JointDef)(unsafe.Pointer(def)).BodyIdB.Index1 && int32FromUint16((*JointDef)(unsafe.Pointer(def)).BodyIdA.World0) == int32FromUint16((*JointDef)(unsafe.Pointer(def)).BodyIdB.World0) && int32FromUint16((*JointDef)(unsafe.Pointer(def)).BodyIdA.Generation) == int32FromUint16((*JointDef)(unsafe.Pointer(def)).BodyIdB.Generation)) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+12832, __ccgo_ts+12321, int32FromInt32(219))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = b2GetBodyFullId(tls, world, (*JointDef)(unsafe.Pointer(def)).BodyIdA)
	bodyB = b2GetBodyFullId(tls, world, (*JointDef)(unsafe.Pointer(def)).BodyIdB)
	// If the joint prevents collisions, then destroy all contacts between attached bodies.
	// Do this first so world arrays are not disturbed in the following code.
	if int32FromUint8((*JointDef)(unsafe.Pointer(def)).CollideConnected) == false1 {
		b2DestroyContactsBetweenBodies(tls, world, bodyA, bodyB)
	}
	bodyIdA = (*b2Body)(unsafe.Pointer(bodyA)).Id
	bodyIdB = (*b2Body)(unsafe.Pointer(bodyB)).Id
	v6 = (*b2Body)(unsafe.Pointer(bodyA)).SetIndex
	v7 = (*b2Body)(unsafe.Pointer(bodyB)).SetIndex
	if v6 > v7 {
		v10 = v6
	} else {
		v10 = v7
	}
	v8 = v10
	goto _9
_9:
	maxSetIndex = v8
	// Create joint id and joint
	jointId = b2AllocId(tls, world+1992)
	if jointId == (*b2World)(unsafe.Pointer(world)).Joints.Count {
		elementSize1 = int32(72)
		if (*b2World)(unsafe.Pointer(world)).Joints.Count >= (*b2World)(unsafe.Pointer(world)).Joints.Capacity {
			oldSize1 = (*b2World)(unsafe.Pointer(world)).Joints.Capacity * elementSize1
			if (*b2World)(unsafe.Pointer(world)).Joints.Capacity == 0 {
				v11 = int32(8)
			} else {
				v11 = int32(2) * (*b2World)(unsafe.Pointer(world)).Joints.Capacity
			}
			newCapacity1 = v11
			newSize1 = newCapacity1 * elementSize1
			(*b2World)(unsafe.Pointer(world)).Joints.Data = b2GrowAlloc(tls, (*b2World)(unsafe.Pointer(world)).Joints.Data, uint64FromInt32(oldSize1), uint64FromInt32(newSize1))
			(*b2World)(unsafe.Pointer(world)).Joints.Capacity = newCapacity1
		}
		*(*b2Joint)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).Joints.Data+uintptr((*b2World)(unsafe.Pointer(world)).Joints.Count)*72)) = b2Joint{}
		(*b2World)(unsafe.Pointer(world)).Joints.Count += int32(1)
	}
	if v1 = !!(0 <= jointId && jointId < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1850, __ccgo_ts+12321, int32FromInt32(242))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	joint = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr(jointId)*72
	(*b2Joint)(unsafe.Pointer(joint)).JointId = jointId
	(*b2Joint)(unsafe.Pointer(joint)).UserData = uintptr((*JointDef)(unsafe.Pointer(def)).UserData)
	v13 = joint + 64
	*(*uint16_t)(unsafe.Pointer(v13)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(v13))) + int32FromInt32(1))
	(*b2Joint)(unsafe.Pointer(joint)).SetIndex = -int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).ColorIndex = -int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).LocalIndex = -int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).IslandId = -int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).IslandIndex = -int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).DrawScale = (*JointDef)(unsafe.Pointer(def)).DrawScale
	(*b2Joint)(unsafe.Pointer(joint)).Type1 = type1
	(*b2Joint)(unsafe.Pointer(joint)).CollideConnected = (*JointDef)(unsafe.Pointer(def)).CollideConnected
	// Doubly linked list on bodyA
	(*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId = bodyIdA
	(*(*b2JointEdge)(unsafe.Add(nil, joint+20))).PrevKey = -int32(1)
	(*(*b2JointEdge)(unsafe.Add(nil, joint+20))).NextKey = (*b2Body)(unsafe.Pointer(bodyA)).HeadJointKey
	keyA = jointId<<int32(1) | 0
	if (*b2Body)(unsafe.Pointer(bodyA)).HeadJointKey != -int32(1) {
		if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(bodyA)).HeadJointKey>>int32(1) && (*b2Body)(unsafe.Pointer(bodyA)).HeadJointKey>>int32(1) < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+12884, __ccgo_ts+12321, int32FromInt32(263))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		jointA = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr((*b2Body)(unsafe.Pointer(bodyA)).HeadJointKey>>int32FromInt32(1))*72
		edgeA = jointA + 20 + uintptr((*b2Body)(unsafe.Pointer(bodyA)).HeadJointKey&int32FromInt32(1))*12
		(*b2JointEdge)(unsafe.Pointer(edgeA)).PrevKey = keyA
	}
	(*b2Body)(unsafe.Pointer(bodyA)).HeadJointKey = keyA
	*(*int32)(unsafe.Add(nil, bodyA+40)) += int32(1)
	// Doubly linked list on bodyB
	(*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId = bodyIdB
	(*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).PrevKey = -int32(1)
	(*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).NextKey = (*b2Body)(unsafe.Pointer(bodyB)).HeadJointKey
	keyB = jointId<<int32(1) | int32(1)
	if (*b2Body)(unsafe.Pointer(bodyB)).HeadJointKey != -int32(1) {
		if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(bodyB)).HeadJointKey>>int32(1) && (*b2Body)(unsafe.Pointer(bodyB)).HeadJointKey>>int32(1) < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+12976, __ccgo_ts+12321, int32FromInt32(278))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		jointB = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr((*b2Body)(unsafe.Pointer(bodyB)).HeadJointKey>>int32FromInt32(1))*72
		edgeB = jointB + 20 + uintptr((*b2Body)(unsafe.Pointer(bodyB)).HeadJointKey&int32FromInt32(1))*12
		(*b2JointEdge)(unsafe.Pointer(edgeB)).PrevKey = keyB
	}
	(*b2Body)(unsafe.Pointer(bodyB)).HeadJointKey = keyB
	*(*int32)(unsafe.Add(nil, bodyB+40)) += int32(1)
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_disabledSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_disabledSet) {
		if v1 = !!(bool(0 <= int32(b2_disabledSet)) && int32(b2_disabledSet) < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+4171, __ccgo_ts+12321, int32FromInt32(290))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		// if either body is disabled, create in disabled set
		set = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b2_disabledSet))*88
		(*b2Joint)(unsafe.Pointer(joint)).SetIndex = int32(b2_disabledSet)
		(*b2Joint)(unsafe.Pointer(joint)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(set)).JointSims.Count
		v13 = set + 32
		v17 = set + 32 + 8
		v18 = set + 32 + 12
		v6 = int32(252)
		if *(*int32)(unsafe.Add(nil, v17)) >= *(*int32)(unsafe.Add(nil, v18)) {
			oldCapacity = *(*int32)(unsafe.Add(nil, v18))
			oldSize = oldCapacity * v6
			if oldCapacity == 0 {
				v7 = int32(8)
			} else {
				v7 = int32(2) * oldCapacity
			}
			newCapacity = v7
			newSize = newCapacity * v6
			*(*uintptr)(unsafe.Add(nil, v13)) = b2GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v13)), uint64FromInt32(oldSize), uint64FromInt32(newSize))
			*(*int32)(unsafe.Add(nil, v18)) = newCapacity
		}
		v22 = v17
		v8 = *(*int32)(unsafe.Pointer(v22))
		*(*int32)(unsafe.Pointer(v22)) = *(*int32)(unsafe.Pointer(v22)) + 1
		v19 = *(*uintptr)(unsafe.Add(nil, v13)) + uintptr(v8*v6)
		goto _23
	_23:
		jointSim = v19
		memset(tls, jointSim, 0, uint64(252))
		(*b2JointSim)(unsafe.Pointer(jointSim)).JointId = jointId
		(*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdA = bodyIdA
		(*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdB = bodyIdB
	} else {
		if (*b2Body)(unsafe.Pointer(bodyA)).Type1 != int32(b2_dynamicBody) && (*b2Body)(unsafe.Pointer(bodyB)).Type1 != int32(b2_dynamicBody) {
			if v1 = !!(bool(0 <= int32(b2_staticSet)) && int32(b2_staticSet) < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
				b2InternalAssert(tls, __ccgo_ts+2964, __ccgo_ts+12321, int32FromInt32(304))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			// joint is not attached to a dynamic body
			set1 = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b2_staticSet))*88
			(*b2Joint)(unsafe.Pointer(joint)).SetIndex = int32(b2_staticSet)
			(*b2Joint)(unsafe.Pointer(joint)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(set1)).JointSims.Count
			v13 = set1 + 32
			v17 = set1 + 32 + 8
			v18 = set1 + 32 + 12
			v6 = int32(252)
			if *(*int32)(unsafe.Add(nil, v17)) >= *(*int32)(unsafe.Add(nil, v18)) {
				oldCapacity = *(*int32)(unsafe.Add(nil, v18))
				oldSize = oldCapacity * v6
				if oldCapacity == 0 {
					v7 = int32(8)
				} else {
					v7 = int32(2) * oldCapacity
				}
				newCapacity = v7
				newSize = newCapacity * v6
				*(*uintptr)(unsafe.Add(nil, v13)) = b2GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v13)), uint64FromInt32(oldSize), uint64FromInt32(newSize))
				*(*int32)(unsafe.Add(nil, v18)) = newCapacity
			}
			v22 = v17
			v8 = *(*int32)(unsafe.Pointer(v22))
			*(*int32)(unsafe.Pointer(v22)) = *(*int32)(unsafe.Pointer(v22)) + 1
			v19 = *(*uintptr)(unsafe.Add(nil, v13)) + uintptr(v8*v6)
			goto _33
		_33:
			jointSim = v19
			memset(tls, jointSim, 0, uint64(252))
			(*b2JointSim)(unsafe.Pointer(jointSim)).JointId = jointId
			(*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdA = bodyIdA
			(*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdB = bodyIdB
		} else {
			if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) {
				// if either body is sleeping, wake it
				if maxSetIndex >= int32(b2_firstSleepingSet) {
					b2WakeSolverSet(tls, world, maxSetIndex)
				}
				(*b2Joint)(unsafe.Pointer(joint)).SetIndex = int32(b2_awakeSet)
				jointSim = b2CreateJointInGraph(tls, world, joint)
				(*b2JointSim)(unsafe.Pointer(jointSim)).JointId = jointId
				(*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdA = bodyIdA
				(*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdB = bodyIdB
			} else {
				// joint connected between sleeping and/or static bodies
				if v1 = !!((*b2Body)(unsafe.Pointer(bodyA)).SetIndex >= int32(b2_firstSleepingSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex >= int32(b2_firstSleepingSet)); !v1 {
					b2InternalAssert(tls, __ccgo_ts+13068, __ccgo_ts+12321, int32FromInt32(333))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				if v1 = !!((*b2Body)(unsafe.Pointer(bodyA)).SetIndex != int32(b2_staticSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex != int32(b2_staticSet)); !v1 {
					b2InternalAssert(tls, __ccgo_ts+6121, __ccgo_ts+12321, int32FromInt32(334))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				// joint should go into the sleeping set (not static set)
				setIndex = maxSetIndex
				if v1 = !!(0 <= setIndex && setIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
					b2InternalAssert(tls, __ccgo_ts+6188, __ccgo_ts+12321, int32FromInt32(339))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				set2 = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setIndex)*88
				(*b2Joint)(unsafe.Pointer(joint)).SetIndex = setIndex
				(*b2Joint)(unsafe.Pointer(joint)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(set2)).JointSims.Count
				v13 = set2 + 32
				v17 = set2 + 32 + 8
				v18 = set2 + 32 + 12
				v6 = int32(252)
				if *(*int32)(unsafe.Add(nil, v17)) >= *(*int32)(unsafe.Add(nil, v18)) {
					oldCapacity = *(*int32)(unsafe.Add(nil, v18))
					oldSize = oldCapacity * v6
					if oldCapacity == 0 {
						v7 = int32(8)
					} else {
						v7 = int32(2) * oldCapacity
					}
					newCapacity = v7
					newSize = newCapacity * v6
					*(*uintptr)(unsafe.Add(nil, v13)) = b2GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v13)), uint64FromInt32(oldSize), uint64FromInt32(newSize))
					*(*int32)(unsafe.Add(nil, v18)) = newCapacity
				}
				v22 = v17
				v8 = *(*int32)(unsafe.Pointer(v22))
				*(*int32)(unsafe.Pointer(v22)) = *(*int32)(unsafe.Pointer(v22)) + 1
				v19 = *(*uintptr)(unsafe.Add(nil, v13)) + uintptr(v8*v6)
				goto _45
			_45:
				jointSim = v19
				memset(tls, jointSim, 0, uint64(252))
				// These must be set to accommodate the merge below
				(*b2JointSim)(unsafe.Pointer(jointSim)).JointId = jointId
				(*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdA = bodyIdA
				(*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdB = bodyIdB
				if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex != (*b2Body)(unsafe.Pointer(bodyB)).SetIndex && (*b2Body)(unsafe.Pointer(bodyA)).SetIndex >= int32(b2_firstSleepingSet) && (*b2Body)(unsafe.Pointer(bodyB)).SetIndex >= int32(b2_firstSleepingSet) {
					// merge sleeping sets
					b2MergeSolverSets(tls, world, (*b2Body)(unsafe.Pointer(bodyA)).SetIndex, (*b2Body)(unsafe.Pointer(bodyB)).SetIndex)
					if v1 = !!((*b2Body)(unsafe.Pointer(bodyA)).SetIndex == (*b2Body)(unsafe.Pointer(bodyB)).SetIndex); !v1 {
						b2InternalAssert(tls, __ccgo_ts+13149, __ccgo_ts+12321, int32FromInt32(356))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					// fix potentially invalid set index
					setIndex = (*b2Body)(unsafe.Pointer(bodyA)).SetIndex
					if v1 = !!(0 <= setIndex && setIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
						b2InternalAssert(tls, __ccgo_ts+6188, __ccgo_ts+12321, int32FromInt32(361))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					mergedSet = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setIndex)*88
					// Careful! The joint sim pointer was orphaned by the set merge.
					if v2 = !!(0 <= (*b2Joint)(unsafe.Pointer(joint)).LocalIndex && (*b2Joint)(unsafe.Pointer(joint)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(mergedSet)).JointSims.Count); !v2 {
						b2InternalAssert(tls, __ccgo_ts+13184, __ccgo_ts+12321, int32FromInt32(364))
					}
					_ = v2 || bool(int32FromInt32(0) != 0)
					jointSim = (*b2SolverSet)(unsafe.Pointer(mergedSet)).JointSims.Data + uintptr((*b2Joint)(unsafe.Pointer(joint)).LocalIndex)*252
				}
				if v1 = !!((*b2Joint)(unsafe.Pointer(joint)).SetIndex == setIndex); !v1 {
					b2InternalAssert(tls, __ccgo_ts+13269, __ccgo_ts+12321, int32FromInt32(367))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
			}
		}
	}
	(*b2JointSim)(unsafe.Pointer(jointSim)).LocalFrameA = (*JointDef)(unsafe.Pointer(def)).LocalFrameA
	(*b2JointSim)(unsafe.Pointer(jointSim)).LocalFrameB = (*JointDef)(unsafe.Pointer(def)).LocalFrameB
	(*b2JointSim)(unsafe.Pointer(jointSim)).Type1 = type1
	(*b2JointSim)(unsafe.Pointer(jointSim)).ConstraintHertz = (*JointDef)(unsafe.Pointer(def)).ConstraintHertz
	(*b2JointSim)(unsafe.Pointer(jointSim)).ConstraintDampingRatio = (*JointDef)(unsafe.Pointer(def)).ConstraintDampingRatio
	(*b2JointSim)(unsafe.Pointer(jointSim)).ConstraintSoftness = b2Softness{
		MassScale: float32FromFloat32(1),
	}
	if v1 = !!(b2IsValidFloat(tls, (*JointDef)(unsafe.Pointer(def)).ForceThreshold) != 0 && (*JointDef)(unsafe.Pointer(def)).ForceThreshold >= float32FromFloat32(0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+13297, __ccgo_ts+12321, int32FromInt32(381))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b2IsValidFloat(tls, (*JointDef)(unsafe.Pointer(def)).TorqueThreshold) != 0 && (*JointDef)(unsafe.Pointer(def)).TorqueThreshold >= float32FromFloat32(0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+13366, __ccgo_ts+12321, int32FromInt32(382))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	(*b2JointSim)(unsafe.Pointer(jointSim)).ForceThreshold = (*JointDef)(unsafe.Pointer(def)).ForceThreshold
	(*b2JointSim)(unsafe.Pointer(jointSim)).TorqueThreshold = (*JointDef)(unsafe.Pointer(def)).TorqueThreshold
	if v1 = !!((*b2JointSim)(unsafe.Pointer(jointSim)).JointId == jointId); !v1 {
		b2InternalAssert(tls, __ccgo_ts+13437, __ccgo_ts+12321, int32FromInt32(387))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdA == bodyIdA); !v1 {
		b2InternalAssert(tls, __ccgo_ts+13466, __ccgo_ts+12321, int32FromInt32(388))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!((*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdB == bodyIdB); !v1 {
		b2InternalAssert(tls, __ccgo_ts+13495, __ccgo_ts+12321, int32FromInt32(389))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2Joint)(unsafe.Pointer(joint)).SetIndex > int32(b2_disabledSet) {
		// Add edge to island graph
		b2LinkJoint(tls, world, joint)
	}
	b2ValidateSolverSets(tls, world)
	return b2JointPair{
		Joint:    joint,
		JointSim: jointSim,
	}
}

func b2CreateDistanceJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var empty b2DistanceJoint
	var joint, world uintptr
	var jointId JointId
	var pair b2JointPair
	var v1 bool
	var v14, v15, v16, v18 float32
	var _ /* _ca at bp+0 */ b2RecArgs_CreateDistanceJoint
	_, _, _, _, _, _, _, _, _, _ = empty, joint, jointId, pair, world, v1, v14, v15, v16, v18
	if v1 = !!((*DistanceJointDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1077, __ccgo_ts+12321, int32FromInt32(404))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if int32FromUint8(b2IsValidJointDef(tls, def)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+152, __ccgo_ts+13524, uintptr(unsafe.Pointer(&__func__23))))
		if v1 = !!(b2IsValidJointDef(tls, def) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13524, __ccgo_ts+12321, int32FromInt32(405))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if boolInt32(b2IsValidFloat(tls, (*DistanceJointDef)(unsafe.Pointer(def)).Length) != 0 && (*DistanceJointDef)(unsafe.Pointer(def)).Length > float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+152, __ccgo_ts+13556, uintptr(unsafe.Pointer(&__func__23))))
		if v1 = !!(b2IsValidFloat(tls, (*DistanceJointDef)(unsafe.Pointer(def)).Length) != 0 && (*DistanceJointDef)(unsafe.Pointer(def)).Length > float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13556, __ccgo_ts+12321, int32FromInt32(406))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*DistanceJointDef)(unsafe.Pointer(def)).Hertz)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+152, __ccgo_ts+13608, uintptr(unsafe.Pointer(&__func__23))))
		if v1 = !!(b2IsValidFloat(tls, (*DistanceJointDef)(unsafe.Pointer(def)).Hertz) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13608, __ccgo_ts+12321, int32FromInt32(407))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*DistanceJointDef)(unsafe.Pointer(def)).DampingRatio)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+152, __ccgo_ts+13637, uintptr(unsafe.Pointer(&__func__23))))
		if v1 = !!(b2IsValidFloat(tls, (*DistanceJointDef)(unsafe.Pointer(def)).DampingRatio) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13637, __ccgo_ts+12321, int32FromInt32(408))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*DistanceJointDef)(unsafe.Pointer(def)).MinLength)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+152, __ccgo_ts+13673, uintptr(unsafe.Pointer(&__func__23))))
		if v1 = !!(b2IsValidFloat(tls, (*DistanceJointDef)(unsafe.Pointer(def)).MinLength) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13673, __ccgo_ts+12321, int32FromInt32(409))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*DistanceJointDef)(unsafe.Pointer(def)).MaxLength)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+152, __ccgo_ts+13706, uintptr(unsafe.Pointer(&__func__23))))
		if v1 = !!(b2IsValidFloat(tls, (*DistanceJointDef)(unsafe.Pointer(def)).MaxLength) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13706, __ccgo_ts+12321, int32FromInt32(410))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*DistanceJointDef)(unsafe.Pointer(def)).MaxMotorForce)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+152, __ccgo_ts+13739, uintptr(unsafe.Pointer(&__func__23))))
		if v1 = !!(b2IsValidFloat(tls, (*DistanceJointDef)(unsafe.Pointer(def)).MaxMotorForce) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13739, __ccgo_ts+12321, int32FromInt32(411))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*DistanceJointDef)(unsafe.Pointer(def)).MotorSpeed)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+152, __ccgo_ts+13776, uintptr(unsafe.Pointer(&__func__23))))
		if v1 = !!(b2IsValidFloat(tls, (*DistanceJointDef)(unsafe.Pointer(def)).MotorSpeed) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13776, __ccgo_ts+12321, int32FromInt32(412))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*DistanceJointDef)(unsafe.Pointer(def)).LowerSpringForce)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+152, __ccgo_ts+13810, uintptr(unsafe.Pointer(&__func__23))))
		if v1 = !!(b2IsValidFloat(tls, (*DistanceJointDef)(unsafe.Pointer(def)).LowerSpringForce) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13810, __ccgo_ts+12321, int32FromInt32(413))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*DistanceJointDef)(unsafe.Pointer(def)).UpperSpringForce)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+152, __ccgo_ts+13850, uintptr(unsafe.Pointer(&__func__23))))
		if v1 = !!(b2IsValidFloat(tls, (*DistanceJointDef)(unsafe.Pointer(def)).UpperSpringForce) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13850, __ccgo_ts+12321, int32FromInt32(414))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if v1 = !!((*DistanceJointDef)(unsafe.Pointer(def)).LowerSpringForce <= (*DistanceJointDef)(unsafe.Pointer(def)).UpperSpringForce); !v1 {
		b2InternalAssert(tls, __ccgo_ts+13890, __ccgo_ts+12321, int32FromInt32(415))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+12321, int32FromInt32(419))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return JointId{}
	}
	pair = b2CreateJoint(tls, world, def, int32(b2_distanceJoint))
	joint = pair.JointSim
	empty = b2DistanceJoint{}
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint = empty
	v14 = (*DistanceJointDef)(unsafe.Pointer(def)).Length
	v15 = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
	if v14 > v15 {
		v18 = v14
	} else {
		v18 = v15
	}
	v16 = v18
	goto _17
_17:
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.Length = v16
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.Hertz = (*DistanceJointDef)(unsafe.Pointer(def)).Hertz
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.DampingRatio = (*DistanceJointDef)(unsafe.Pointer(def)).DampingRatio
	v14 = (*DistanceJointDef)(unsafe.Pointer(def)).MinLength
	v15 = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
	if v14 > v15 {
		v18 = v14
	} else {
		v18 = v15
	}
	v16 = v18
	goto _22
_22:
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.MinLength = v16
	v14 = (*DistanceJointDef)(unsafe.Pointer(def)).MinLength
	v15 = (*DistanceJointDef)(unsafe.Pointer(def)).MaxLength
	if v14 > v15 {
		v18 = v14
	} else {
		v18 = v15
	}
	v16 = v18
	goto _27
_27:
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.MaxLength = v16
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.MaxMotorForce = (*DistanceJointDef)(unsafe.Pointer(def)).MaxMotorForce
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.MotorSpeed = (*DistanceJointDef)(unsafe.Pointer(def)).MotorSpeed
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.EnableSpring = (*DistanceJointDef)(unsafe.Pointer(def)).EnableSpring
	v14 = (*DistanceJointDef)(unsafe.Pointer(def)).LowerSpringForce
	v15 = (*DistanceJointDef)(unsafe.Pointer(def)).UpperSpringForce
	if v14 < v15 {
		v18 = v14
	} else {
		v18 = v15
	}
	v16 = v18
	goto _32
_32:
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.LowerSpringForce = v16
	v14 = (*DistanceJointDef)(unsafe.Pointer(def)).LowerSpringForce
	v15 = (*DistanceJointDef)(unsafe.Pointer(def)).UpperSpringForce
	if v14 > v15 {
		v18 = v14
	} else {
		v18 = v15
	}
	v16 = v18
	goto _37
_37:
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.UpperSpringForce = v16
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.EnableLimit = (*DistanceJointDef)(unsafe.Pointer(def)).EnableLimit
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.EnableMotor = (*DistanceJointDef)(unsafe.Pointer(def)).EnableMotor
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.Impulse = float32FromFloat32(0)
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.LowerImpulse = float32FromFloat32(0)
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.UpperImpulse = float32FromFloat32(0)
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo15_92.DistanceJoint.MotorImpulse = float32FromFloat32(0)
	jointId = JointId{
		Index1:     (*b2JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_CreateDistanceJoint)(unsafe.Add(nil, bp)) = b2RecArgs_CreateDistanceJoint{
			World: worldId,
			Def:   *(*DistanceJointDef)(unsafe.Add(nil, def)),
		}
		b2RecWriteRet_CreateDistanceJoint(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp, jointId)
	}
	return jointId
}

func b2CreateFilterJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var joint, world uintptr
	var jointId JointId
	var pair b2JointPair
	var v1 bool
	var _ /* _ca at bp+0 */ b2RecArgs_CreateFilterJoint
	_, _, _, _, _ = joint, jointId, pair, world, v1
	if v1 = !!((*FilterJointDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1077, __ccgo_ts+12321, int32FromInt32(458))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if int32FromUint8(b2IsValidJointDef(tls, def)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+104, __ccgo_ts+13524, uintptr(unsafe.Pointer(&__func__20))))
		if v1 = !!(b2IsValidJointDef(tls, def) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13524, __ccgo_ts+12321, int32FromInt32(459))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+12321, int32FromInt32(463))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return JointId{}
	}
	pair = b2CreateJoint(tls, world, def, int32(b2_filterJoint))
	joint = pair.JointSim
	jointId = JointId{
		Index1:     (*b2JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_CreateFilterJoint)(unsafe.Add(nil, bp)) = b2RecArgs_CreateFilterJoint{
			World: worldId,
			Def:   *(*FilterJointDef)(unsafe.Add(nil, def)),
		}
		b2RecWriteRet_CreateFilterJoint(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp, jointId)
	}
	return jointId
}

func b2CreateMotorJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var joint, world uintptr
	var jointId JointId
	var pair b2JointPair
	var v1 bool
	var _ /* _ca at bp+0 */ b2RecArgs_CreateMotorJoint
	_, _, _, _, _ = joint, jointId, pair, world, v1
	if v1 = !!((*MotorJointDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1077, __ccgo_ts+12321, int32FromInt32(483))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if int32FromUint8(b2IsValidJointDef(tls, def)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+13524, uintptr(unsafe.Pointer(&__func__24))))
		if v1 = !!(b2IsValidJointDef(tls, def) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13524, __ccgo_ts+12321, int32FromInt32(484))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidVec2(tls, (*MotorJointDef)(unsafe.Pointer(def)).LinearVelocity)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+1201, uintptr(unsafe.Pointer(&__func__24))))
		if v1 = !!(b2IsValidVec2(tls, (*MotorJointDef)(unsafe.Pointer(def)).LinearVelocity) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1201, __ccgo_ts+12321, int32FromInt32(485))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*MotorJointDef)(unsafe.Pointer(def)).MaxVelocityForce)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+13937, uintptr(unsafe.Pointer(&__func__24))))
		if v1 = !!(b2IsValidFloat(tls, (*MotorJointDef)(unsafe.Pointer(def)).MaxVelocityForce) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13937, __ccgo_ts+12321, int32FromInt32(486))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*MotorJointDef)(unsafe.Pointer(def)).AngularVelocity)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+1238, uintptr(unsafe.Pointer(&__func__24))))
		if v1 = !!(b2IsValidFloat(tls, (*MotorJointDef)(unsafe.Pointer(def)).AngularVelocity) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1238, __ccgo_ts+12321, int32FromInt32(487))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*MotorJointDef)(unsafe.Pointer(def)).MaxVelocityTorque)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+13977, uintptr(unsafe.Pointer(&__func__24))))
		if v1 = !!(b2IsValidFloat(tls, (*MotorJointDef)(unsafe.Pointer(def)).MaxVelocityTorque) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13977, __ccgo_ts+12321, int32FromInt32(488))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*MotorJointDef)(unsafe.Pointer(def)).LinearHertz)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+14018, uintptr(unsafe.Pointer(&__func__24))))
		if v1 = !!(b2IsValidFloat(tls, (*MotorJointDef)(unsafe.Pointer(def)).LinearHertz) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14018, __ccgo_ts+12321, int32FromInt32(489))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*MotorJointDef)(unsafe.Pointer(def)).LinearDampingRatio)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+14053, uintptr(unsafe.Pointer(&__func__24))))
		if v1 = !!(b2IsValidFloat(tls, (*MotorJointDef)(unsafe.Pointer(def)).LinearDampingRatio) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14053, __ccgo_ts+12321, int32FromInt32(490))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*MotorJointDef)(unsafe.Pointer(def)).MaxSpringForce)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+14095, uintptr(unsafe.Pointer(&__func__24))))
		if v1 = !!(b2IsValidFloat(tls, (*MotorJointDef)(unsafe.Pointer(def)).MaxSpringForce) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14095, __ccgo_ts+12321, int32FromInt32(491))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*MotorJointDef)(unsafe.Pointer(def)).AngularHertz)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+14133, uintptr(unsafe.Pointer(&__func__24))))
		if v1 = !!(b2IsValidFloat(tls, (*MotorJointDef)(unsafe.Pointer(def)).AngularHertz) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14133, __ccgo_ts+12321, int32FromInt32(492))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*MotorJointDef)(unsafe.Pointer(def)).AngularDampingRatio)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+14169, uintptr(unsafe.Pointer(&__func__24))))
		if v1 = !!(b2IsValidFloat(tls, (*MotorJointDef)(unsafe.Pointer(def)).AngularDampingRatio) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14169, __ccgo_ts+12321, int32FromInt32(493))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*MotorJointDef)(unsafe.Pointer(def)).MaxSpringTorque)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+14212, uintptr(unsafe.Pointer(&__func__24))))
		if v1 = !!(b2IsValidFloat(tls, (*MotorJointDef)(unsafe.Pointer(def)).MaxSpringTorque) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14212, __ccgo_ts+12321, int32FromInt32(494))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+12321, int32FromInt32(498))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return JointId{}
	}
	pair = b2CreateJoint(tls, world, def, int32(b2_motorJoint))
	joint = pair.JointSim
	*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92)) = b2MotorJoint{}
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LinearVelocity = (*MotorJointDef)(unsafe.Pointer(def)).LinearVelocity
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxVelocityForce = (*MotorJointDef)(unsafe.Pointer(def)).MaxVelocityForce
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).AngularVelocity = (*MotorJointDef)(unsafe.Pointer(def)).AngularVelocity
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxVelocityTorque = (*MotorJointDef)(unsafe.Pointer(def)).MaxVelocityTorque
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LinearHertz = (*MotorJointDef)(unsafe.Pointer(def)).LinearHertz
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LinearDampingRatio = (*MotorJointDef)(unsafe.Pointer(def)).LinearDampingRatio
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxSpringForce = (*MotorJointDef)(unsafe.Pointer(def)).MaxSpringForce
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).AngularHertz = (*MotorJointDef)(unsafe.Pointer(def)).AngularHertz
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).AngularDampingRatio = (*MotorJointDef)(unsafe.Pointer(def)).AngularDampingRatio
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxSpringTorque = (*MotorJointDef)(unsafe.Pointer(def)).MaxSpringTorque
	jointId = JointId{
		Index1:     (*b2JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_CreateMotorJoint)(unsafe.Add(nil, bp)) = b2RecArgs_CreateMotorJoint{
			World: worldId,
			Def:   *(*MotorJointDef)(unsafe.Add(nil, def)),
		}
		b2RecWriteRet_CreateMotorJoint(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp, jointId)
	}
	return jointId
}

func b2CreateMoverJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var joint, world uintptr
	var jointId JointId
	var pair b2JointPair
	var v1 bool
	var _ /* _ca at bp+0 */ b2RecArgs_CreateMoverJoint
	_, _, _, _, _ = joint, jointId, pair, world, v1
	if v1 = !!((*MoverJointDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1077, __ccgo_ts+12321, int32FromInt32(529))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if int32FromUint8(b2IsValidJointDef(tls, def)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+120, __ccgo_ts+13524, uintptr(unsafe.Pointer(&__func__25))))
		if v1 = !!(b2IsValidJointDef(tls, def) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13524, __ccgo_ts+12321, int32FromInt32(530))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidVec2(tls, (*MoverJointDef)(unsafe.Pointer(def)).LinearVelocity)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+120, __ccgo_ts+1201, uintptr(unsafe.Pointer(&__func__25))))
		if v1 = !!(b2IsValidVec2(tls, (*MoverJointDef)(unsafe.Pointer(def)).LinearVelocity) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+1201, __ccgo_ts+12321, int32FromInt32(531))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidVec2(tls, (*MoverJointDef)(unsafe.Pointer(def)).MaxVelocityForce)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+120, __ccgo_ts+14251, uintptr(unsafe.Pointer(&__func__25))))
		if v1 = !!(b2IsValidVec2(tls, (*MoverJointDef)(unsafe.Pointer(def)).MaxVelocityForce) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14251, __ccgo_ts+12321, int32FromInt32(532))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+12321, int32FromInt32(536))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return JointId{}
	}
	pair = b2CreateJoint(tls, world, def, int32(b2_moverJoint))
	joint = pair.JointSim
	*(*b2MoverJoint)(unsafe.Add(unsafe.Pointer(joint), 92)) = b2MoverJoint{}
	(*(*b2MoverJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LinearVelocity = (*MoverJointDef)(unsafe.Pointer(def)).LinearVelocity
	(*(*b2MoverJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxVelocityForce = (*MoverJointDef)(unsafe.Pointer(def)).MaxVelocityForce
	jointId = JointId{
		Index1:     (*b2JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_CreateMoverJoint)(unsafe.Add(nil, bp)) = b2RecArgs_CreateMoverJoint{
			World: worldId,
			Def:   *(*MoverJointDef)(unsafe.Add(nil, def)),
		}
		b2RecWriteRet_CreateMoverJoint(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp, jointId)
	}
	return jointId
}

func b2CreatePogoJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var joint, world uintptr
	var jointId JointId
	var pair b2JointPair
	var v1 bool
	var _ /* _ca at bp+0 */ b2RecArgs_CreatePogoJoint
	_, _, _, _, _ = joint, jointId, pair, world, v1
	if v1 = !!((*PogoJointDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1077, __ccgo_ts+12321, int32FromInt32(559))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if int32FromUint8(b2IsValidJointDef(tls, def)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+136, __ccgo_ts+13524, uintptr(unsafe.Pointer(&__func__26))))
		if v1 = !!(b2IsValidJointDef(tls, def) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13524, __ccgo_ts+12321, int32FromInt32(560))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidVec2(tls, (*PogoJointDef)(unsafe.Pointer(def)).Normal)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+136, __ccgo_ts+14290, uintptr(unsafe.Pointer(&__func__26))))
		if v1 = !!(b2IsValidVec2(tls, (*PogoJointDef)(unsafe.Pointer(def)).Normal) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14290, __ccgo_ts+12321, int32FromInt32(561))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if boolInt32(b2IsValidFloat(tls, (*PogoJointDef)(unsafe.Pointer(def)).RestLength) != 0 && (*PogoJointDef)(unsafe.Pointer(def)).RestLength >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+136, __ccgo_ts+14319, uintptr(unsafe.Pointer(&__func__26))))
		if v1 = !!(b2IsValidFloat(tls, (*PogoJointDef)(unsafe.Pointer(def)).RestLength) != 0 && (*PogoJointDef)(unsafe.Pointer(def)).RestLength >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14319, __ccgo_ts+12321, int32FromInt32(562))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if boolInt32(b2IsValidFloat(tls, (*PogoJointDef)(unsafe.Pointer(def)).Hertz) != 0 && (*PogoJointDef)(unsafe.Pointer(def)).Hertz >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+136, __ccgo_ts+14380, uintptr(unsafe.Pointer(&__func__26))))
		if v1 = !!(b2IsValidFloat(tls, (*PogoJointDef)(unsafe.Pointer(def)).Hertz) != 0 && (*PogoJointDef)(unsafe.Pointer(def)).Hertz >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14380, __ccgo_ts+12321, int32FromInt32(563))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if boolInt32(b2IsValidFloat(tls, (*PogoJointDef)(unsafe.Pointer(def)).DampingRatio) != 0 && (*PogoJointDef)(unsafe.Pointer(def)).DampingRatio >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+136, __ccgo_ts+14431, uintptr(unsafe.Pointer(&__func__26))))
		if v1 = !!(b2IsValidFloat(tls, (*PogoJointDef)(unsafe.Pointer(def)).DampingRatio) != 0 && (*PogoJointDef)(unsafe.Pointer(def)).DampingRatio >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14431, __ccgo_ts+12321, int32FromInt32(564))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if boolInt32(b2IsValidFloat(tls, (*PogoJointDef)(unsafe.Pointer(def)).MaxTensionForce) != 0 && (*PogoJointDef)(unsafe.Pointer(def)).MaxTensionForce >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+136, __ccgo_ts+14496, uintptr(unsafe.Pointer(&__func__26))))
		if v1 = !!(b2IsValidFloat(tls, (*PogoJointDef)(unsafe.Pointer(def)).MaxTensionForce) != 0 && (*PogoJointDef)(unsafe.Pointer(def)).MaxTensionForce >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14496, __ccgo_ts+12321, int32FromInt32(565))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if boolInt32(b2IsValidFloat(tls, (*PogoJointDef)(unsafe.Pointer(def)).MaxCompressionForce) != 0 && (*PogoJointDef)(unsafe.Pointer(def)).MaxCompressionForce >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+136, __ccgo_ts+14567, uintptr(unsafe.Pointer(&__func__26))))
		if v1 = !!(b2IsValidFloat(tls, (*PogoJointDef)(unsafe.Pointer(def)).MaxCompressionForce) != 0 && (*PogoJointDef)(unsafe.Pointer(def)).MaxCompressionForce >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14567, __ccgo_ts+12321, int32FromInt32(566))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*PogoJointDef)(unsafe.Pointer(def)).Impulse)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+136, __ccgo_ts+14646, uintptr(unsafe.Pointer(&__func__26))))
		if v1 = !!(b2IsValidFloat(tls, (*PogoJointDef)(unsafe.Pointer(def)).Impulse) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14646, __ccgo_ts+12321, int32FromInt32(567))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*PogoJointDef)(unsafe.Pointer(def)).Velocity)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+136, __ccgo_ts+14677, uintptr(unsafe.Pointer(&__func__26))))
		if v1 = !!(b2IsValidFloat(tls, (*PogoJointDef)(unsafe.Pointer(def)).Velocity) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14677, __ccgo_ts+12321, int32FromInt32(568))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+12321, int32FromInt32(572))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return JointId{}
	}
	pair = b2CreateJoint(tls, world, def, int32(b2_pogoJoint))
	joint = pair.JointSim
	*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(joint), 92)) = b2PogoJoint{}
	(*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).RestLength = (*PogoJointDef)(unsafe.Pointer(def)).RestLength
	(*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).Hertz = (*PogoJointDef)(unsafe.Pointer(def)).Hertz
	(*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).DampingRatio = (*PogoJointDef)(unsafe.Pointer(def)).DampingRatio
	(*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxTensionForce = (*PogoJointDef)(unsafe.Pointer(def)).MaxTensionForce
	(*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxCompressionForce = (*PogoJointDef)(unsafe.Pointer(def)).MaxCompressionForce
	(*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).Normal = (*PogoJointDef)(unsafe.Pointer(def)).Normal
	(*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).Impulse = (*PogoJointDef)(unsafe.Pointer(def)).Impulse
	(*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).Velocity = (*PogoJointDef)(unsafe.Pointer(def)).Velocity
	jointId = JointId{
		Index1:     (*b2JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_CreatePogoJoint)(unsafe.Add(nil, bp)) = b2RecArgs_CreatePogoJoint{
			World: worldId,
			Def:   *(*PogoJointDef)(unsafe.Add(nil, def)),
		}
		b2RecWriteRet_CreatePogoJoint(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp, jointId)
	}
	return jointId
}

func b2CreatePrismaticJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var joint, world uintptr
	var jointId JointId
	var pair b2JointPair
	var v1 bool
	var v12, v13, v14, v16 float32
	var _ /* _ca at bp+0 */ b2RecArgs_CreatePrismaticJoint
	_, _, _, _, _, _, _, _, _ = joint, jointId, pair, world, v1, v12, v13, v14, v16
	if v1 = !!((*PrismaticJointDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1077, __ccgo_ts+12321, int32FromInt32(606))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if int32FromUint8(b2IsValidJointDef(tls, def)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+13524, uintptr(unsafe.Pointer(&__func__27))))
		if v1 = !!(b2IsValidJointDef(tls, def) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13524, __ccgo_ts+12321, int32FromInt32(607))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*PrismaticJointDef)(unsafe.Pointer(def)).Hertz)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+13608, uintptr(unsafe.Pointer(&__func__27))))
		if v1 = !!(b2IsValidFloat(tls, (*PrismaticJointDef)(unsafe.Pointer(def)).Hertz) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13608, __ccgo_ts+12321, int32FromInt32(608))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*PrismaticJointDef)(unsafe.Pointer(def)).DampingRatio)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+13637, uintptr(unsafe.Pointer(&__func__27))))
		if v1 = !!(b2IsValidFloat(tls, (*PrismaticJointDef)(unsafe.Pointer(def)).DampingRatio) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13637, __ccgo_ts+12321, int32FromInt32(609))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*PrismaticJointDef)(unsafe.Pointer(def)).TargetTranslation)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+14709, uintptr(unsafe.Pointer(&__func__27))))
		if v1 = !!(b2IsValidFloat(tls, (*PrismaticJointDef)(unsafe.Pointer(def)).TargetTranslation) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14709, __ccgo_ts+12321, int32FromInt32(610))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*PrismaticJointDef)(unsafe.Pointer(def)).LowerTranslation)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+14750, uintptr(unsafe.Pointer(&__func__27))))
		if v1 = !!(b2IsValidFloat(tls, (*PrismaticJointDef)(unsafe.Pointer(def)).LowerTranslation) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14750, __ccgo_ts+12321, int32FromInt32(611))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*PrismaticJointDef)(unsafe.Pointer(def)).UpperTranslation)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+14790, uintptr(unsafe.Pointer(&__func__27))))
		if v1 = !!(b2IsValidFloat(tls, (*PrismaticJointDef)(unsafe.Pointer(def)).UpperTranslation) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14790, __ccgo_ts+12321, int32FromInt32(612))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*PrismaticJointDef)(unsafe.Pointer(def)).MaxMotorForce)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+13739, uintptr(unsafe.Pointer(&__func__27))))
		if v1 = !!(b2IsValidFloat(tls, (*PrismaticJointDef)(unsafe.Pointer(def)).MaxMotorForce) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13739, __ccgo_ts+12321, int32FromInt32(613))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*PrismaticJointDef)(unsafe.Pointer(def)).MotorSpeed)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+13776, uintptr(unsafe.Pointer(&__func__27))))
		if v1 = !!(b2IsValidFloat(tls, (*PrismaticJointDef)(unsafe.Pointer(def)).MotorSpeed) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13776, __ccgo_ts+12321, int32FromInt32(614))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if v1 = !!((*PrismaticJointDef)(unsafe.Pointer(def)).LowerTranslation <= (*PrismaticJointDef)(unsafe.Pointer(def)).UpperTranslation); !v1 {
		b2InternalAssert(tls, __ccgo_ts+14830, __ccgo_ts+12321, int32FromInt32(615))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+12321, int32FromInt32(619))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return JointId{}
	}
	pair = b2CreateJoint(tls, world, def, int32(b2_prismaticJoint))
	joint = pair.JointSim
	*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92)) = b2PrismaticJoint{}
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).Hertz = (*PrismaticJointDef)(unsafe.Pointer(def)).Hertz
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).DampingRatio = (*PrismaticJointDef)(unsafe.Pointer(def)).DampingRatio
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).TargetTranslation = (*PrismaticJointDef)(unsafe.Pointer(def)).TargetTranslation
	v12 = (*PrismaticJointDef)(unsafe.Pointer(def)).LowerTranslation
	v13 = (*PrismaticJointDef)(unsafe.Pointer(def)).UpperTranslation
	if v12 < v13 {
		v16 = v12
	} else {
		v16 = v13
	}
	v14 = v16
	goto _15
_15:
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LowerTranslation = v14
	v12 = (*PrismaticJointDef)(unsafe.Pointer(def)).LowerTranslation
	v13 = (*PrismaticJointDef)(unsafe.Pointer(def)).UpperTranslation
	if v12 > v13 {
		v16 = v12
	} else {
		v16 = v13
	}
	v14 = v16
	goto _20
_20:
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).UpperTranslation = v14
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxMotorForce = (*PrismaticJointDef)(unsafe.Pointer(def)).MaxMotorForce
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MotorSpeed = (*PrismaticJointDef)(unsafe.Pointer(def)).MotorSpeed
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableSpring = (*PrismaticJointDef)(unsafe.Pointer(def)).EnableSpring
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableLimit = (*PrismaticJointDef)(unsafe.Pointer(def)).EnableLimit
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableMotor = (*PrismaticJointDef)(unsafe.Pointer(def)).EnableMotor
	jointId = JointId{
		Index1:     (*b2JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_CreatePrismaticJoint)(unsafe.Add(nil, bp)) = b2RecArgs_CreatePrismaticJoint{
			World: worldId,
			Def:   *(*PrismaticJointDef)(unsafe.Add(nil, def)),
		}
		b2RecWriteRet_CreatePrismaticJoint(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp, jointId)
	}
	return jointId
}

func b2CreateRevoluteJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var empty b2RevoluteJoint
	var joint, world uintptr
	var jointId JointId
	var lowerAngle, upperAngle, v12, v13, v14, v15, v17, v18, v19, v20, v21, v23, v24, v25, v26, v28 float32
	var pair b2JointPair
	var v1 bool
	var _ /* _ca at bp+0 */ b2RecArgs_CreateRevoluteJoint
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = empty, joint, jointId, lowerAngle, pair, upperAngle, world, v1, v12, v13, v14, v15, v17, v18, v19, v20, v21, v23, v24, v25, v26, v28
	if v1 = !!((*RevoluteJointDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1077, __ccgo_ts+12321, int32FromInt32(655))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if int32FromUint8(b2IsValidJointDef(tls, def)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+13524, uintptr(unsafe.Pointer(&__func__33))))
		if v1 = !!(b2IsValidJointDef(tls, def) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13524, __ccgo_ts+12321, int32FromInt32(656))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*RevoluteJointDef)(unsafe.Pointer(def)).TargetAngle)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+14877, uintptr(unsafe.Pointer(&__func__33))))
		if v1 = !!(b2IsValidFloat(tls, (*RevoluteJointDef)(unsafe.Pointer(def)).TargetAngle) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14877, __ccgo_ts+12321, int32FromInt32(657))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*RevoluteJointDef)(unsafe.Pointer(def)).Hertz)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+13608, uintptr(unsafe.Pointer(&__func__33))))
		if v1 = !!(b2IsValidFloat(tls, (*RevoluteJointDef)(unsafe.Pointer(def)).Hertz) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13608, __ccgo_ts+12321, int32FromInt32(658))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*RevoluteJointDef)(unsafe.Pointer(def)).DampingRatio)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+13637, uintptr(unsafe.Pointer(&__func__33))))
		if v1 = !!(b2IsValidFloat(tls, (*RevoluteJointDef)(unsafe.Pointer(def)).DampingRatio) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13637, __ccgo_ts+12321, int32FromInt32(659))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*RevoluteJointDef)(unsafe.Pointer(def)).LowerAngle)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+14912, uintptr(unsafe.Pointer(&__func__33))))
		if v1 = !!(b2IsValidFloat(tls, (*RevoluteJointDef)(unsafe.Pointer(def)).LowerAngle) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14912, __ccgo_ts+12321, int32FromInt32(660))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*RevoluteJointDef)(unsafe.Pointer(def)).UpperAngle)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+14946, uintptr(unsafe.Pointer(&__func__33))))
		if v1 = !!(b2IsValidFloat(tls, (*RevoluteJointDef)(unsafe.Pointer(def)).UpperAngle) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14946, __ccgo_ts+12321, int32FromInt32(661))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*RevoluteJointDef)(unsafe.Pointer(def)).MaxMotorTorque)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+14980, uintptr(unsafe.Pointer(&__func__33))))
		if v1 = !!(b2IsValidFloat(tls, (*RevoluteJointDef)(unsafe.Pointer(def)).MaxMotorTorque) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14980, __ccgo_ts+12321, int32FromInt32(662))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*RevoluteJointDef)(unsafe.Pointer(def)).MotorSpeed)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+144, __ccgo_ts+13776, uintptr(unsafe.Pointer(&__func__33))))
		if v1 = !!(b2IsValidFloat(tls, (*RevoluteJointDef)(unsafe.Pointer(def)).MotorSpeed) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13776, __ccgo_ts+12321, int32FromInt32(663))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if v1 = !!((*RevoluteJointDef)(unsafe.Pointer(def)).LowerAngle <= (*RevoluteJointDef)(unsafe.Pointer(def)).UpperAngle); !v1 {
		b2InternalAssert(tls, __ccgo_ts+15018, __ccgo_ts+12321, int32FromInt32(664))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+12321, int32FromInt32(668))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return JointId{}
	}
	pair = b2CreateJoint(tls, world, def, int32(b2_revoluteJoint))
	joint = pair.JointSim
	empty = b2RevoluteJoint{}
	*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92)) = empty
	v12 = (*RevoluteJointDef)(unsafe.Pointer(def)).TargetAngle
	v13 = -float32FromFloat32(3.14159265359)
	v14 = float32FromFloat32(3.14159265359)
	if v12 < v13 {
		v17 = v13
	} else {
		if v12 > v14 {
			v18 = v14
		} else {
			v18 = v12
		}
		v17 = v18
	}
	v15 = v17
	goto _16
_16:
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).TargetAngle = v15
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).Hertz = (*RevoluteJointDef)(unsafe.Pointer(def)).Hertz
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).DampingRatio = (*RevoluteJointDef)(unsafe.Pointer(def)).DampingRatio
	v12 = (*RevoluteJointDef)(unsafe.Pointer(def)).LowerAngle
	v13 = (*RevoluteJointDef)(unsafe.Pointer(def)).UpperAngle
	if v12 < v13 {
		v15 = v12
	} else {
		v15 = v13
	}
	v14 = v15
	goto _22
_22:
	lowerAngle = v14
	v17 = (*RevoluteJointDef)(unsafe.Pointer(def)).LowerAngle
	v18 = (*RevoluteJointDef)(unsafe.Pointer(def)).UpperAngle
	if v17 > v18 {
		v20 = v17
	} else {
		v20 = v18
	}
	v19 = v20
	goto _27
_27:
	upperAngle = v19
	v21 = lowerAngle
	v23 = float32(-float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	v24 = float32(float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	if v21 < v23 {
		v26 = v23
	} else {
		if v21 > v24 {
			v28 = v24
		} else {
			v28 = v21
		}
		v26 = v28
	}
	v25 = v26
	goto _33
_33:
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LowerAngle = v25
	v12 = upperAngle
	v13 = float32(-float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	v14 = float32(float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	if v12 < v13 {
		v17 = v13
	} else {
		if v12 > v14 {
			v18 = v14
		} else {
			v18 = v12
		}
		v17 = v18
	}
	v15 = v17
	goto _40
_40:
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).UpperAngle = v15
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxMotorTorque = (*RevoluteJointDef)(unsafe.Pointer(def)).MaxMotorTorque
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MotorSpeed = (*RevoluteJointDef)(unsafe.Pointer(def)).MotorSpeed
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableSpring = (*RevoluteJointDef)(unsafe.Pointer(def)).EnableSpring
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableLimit = (*RevoluteJointDef)(unsafe.Pointer(def)).EnableLimit
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableMotor = (*RevoluteJointDef)(unsafe.Pointer(def)).EnableMotor
	jointId = JointId{
		Index1:     (*b2JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_CreateRevoluteJoint)(unsafe.Add(nil, bp)) = b2RecArgs_CreateRevoluteJoint{
			World: worldId,
			Def:   *(*RevoluteJointDef)(unsafe.Add(nil, def)),
		}
		b2RecWriteRet_CreateRevoluteJoint(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp, jointId)
	}
	return jointId
}

func b2CreateWeldJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var empty b2WeldJoint
	var joint, world uintptr
	var jointId JointId
	var pair b2JointPair
	var v1 bool
	var _ /* _ca at bp+0 */ b2RecArgs_CreateWeldJoint
	_, _, _, _, _, _ = empty, joint, jointId, pair, world, v1
	if v1 = !!((*WeldJointDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1077, __ccgo_ts+12321, int32FromInt32(706))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if int32FromUint8(b2IsValidJointDef(tls, def)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+120, __ccgo_ts+13524, uintptr(unsafe.Pointer(&__func__28))))
		if v1 = !!(b2IsValidJointDef(tls, def) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13524, __ccgo_ts+12321, int32FromInt32(707))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*WeldJointDef)(unsafe.Pointer(def)).LinearHertz)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+120, __ccgo_ts+14018, uintptr(unsafe.Pointer(&__func__28))))
		if v1 = !!(b2IsValidFloat(tls, (*WeldJointDef)(unsafe.Pointer(def)).LinearHertz) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14018, __ccgo_ts+12321, int32FromInt32(708))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*WeldJointDef)(unsafe.Pointer(def)).AngularHertz)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+120, __ccgo_ts+14133, uintptr(unsafe.Pointer(&__func__28))))
		if v1 = !!(b2IsValidFloat(tls, (*WeldJointDef)(unsafe.Pointer(def)).AngularHertz) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14133, __ccgo_ts+12321, int32FromInt32(709))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*WeldJointDef)(unsafe.Pointer(def)).LinearDampingRatio)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+120, __ccgo_ts+14053, uintptr(unsafe.Pointer(&__func__28))))
		if v1 = !!(b2IsValidFloat(tls, (*WeldJointDef)(unsafe.Pointer(def)).LinearDampingRatio) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14053, __ccgo_ts+12321, int32FromInt32(710))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*WeldJointDef)(unsafe.Pointer(def)).AngularDampingRatio)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+120, __ccgo_ts+14169, uintptr(unsafe.Pointer(&__func__28))))
		if v1 = !!(b2IsValidFloat(tls, (*WeldJointDef)(unsafe.Pointer(def)).AngularDampingRatio) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14169, __ccgo_ts+12321, int32FromInt32(711))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+12321, int32FromInt32(715))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return JointId{}
	}
	pair = b2CreateJoint(tls, world, def, int32(b2_weldJoint))
	joint = pair.JointSim
	empty = b2WeldJoint{}
	*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 92)) = empty
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LinearHertz = (*WeldJointDef)(unsafe.Pointer(def)).LinearHertz
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LinearDampingRatio = (*WeldJointDef)(unsafe.Pointer(def)).LinearDampingRatio
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).AngularHertz = (*WeldJointDef)(unsafe.Pointer(def)).AngularHertz
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).AngularDampingRatio = (*WeldJointDef)(unsafe.Pointer(def)).AngularDampingRatio
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LinearImpulse = b2Vec2_zero
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).AngularImpulse = float32FromFloat32(0)
	jointId = JointId{
		Index1:     (*b2JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_CreateWeldJoint)(unsafe.Add(nil, bp)) = b2RecArgs_CreateWeldJoint{
			World: worldId,
			Def:   *(*WeldJointDef)(unsafe.Add(nil, def)),
		}
		b2RecWriteRet_CreateWeldJoint(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp, jointId)
	}
	return jointId
}

func b2CreateWheelJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var joint, world uintptr
	var jointId JointId
	var pair b2JointPair
	var v1 bool
	var v11, v12, v13, v15 float32
	var _ /* _ca at bp+0 */ b2RecArgs_CreateWheelJoint
	_, _, _, _, _, _, _, _, _ = joint, jointId, pair, world, v1, v11, v12, v13, v15
	if v1 = !!((*WheelJointDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1077, __ccgo_ts+12321, int32FromInt32(744))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if int32FromUint8(b2IsValidJointDef(tls, def)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+136, __ccgo_ts+13524, uintptr(unsafe.Pointer(&__func__29))))
		if v1 = !!(b2IsValidJointDef(tls, def) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13524, __ccgo_ts+12321, int32FromInt32(745))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*WheelJointDef)(unsafe.Pointer(def)).Hertz)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+136, __ccgo_ts+13608, uintptr(unsafe.Pointer(&__func__29))))
		if v1 = !!(b2IsValidFloat(tls, (*WheelJointDef)(unsafe.Pointer(def)).Hertz) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13608, __ccgo_ts+12321, int32FromInt32(746))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*WheelJointDef)(unsafe.Pointer(def)).DampingRatio)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+136, __ccgo_ts+13637, uintptr(unsafe.Pointer(&__func__29))))
		if v1 = !!(b2IsValidFloat(tls, (*WheelJointDef)(unsafe.Pointer(def)).DampingRatio) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13637, __ccgo_ts+12321, int32FromInt32(747))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*WheelJointDef)(unsafe.Pointer(def)).LowerTranslation)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+136, __ccgo_ts+14750, uintptr(unsafe.Pointer(&__func__29))))
		if v1 = !!(b2IsValidFloat(tls, (*WheelJointDef)(unsafe.Pointer(def)).LowerTranslation) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14750, __ccgo_ts+12321, int32FromInt32(748))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*WheelJointDef)(unsafe.Pointer(def)).UpperTranslation)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+136, __ccgo_ts+14790, uintptr(unsafe.Pointer(&__func__29))))
		if v1 = !!(b2IsValidFloat(tls, (*WheelJointDef)(unsafe.Pointer(def)).UpperTranslation) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14790, __ccgo_ts+12321, int32FromInt32(749))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*WheelJointDef)(unsafe.Pointer(def)).MaxMotorTorque)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+136, __ccgo_ts+14980, uintptr(unsafe.Pointer(&__func__29))))
		if v1 = !!(b2IsValidFloat(tls, (*WheelJointDef)(unsafe.Pointer(def)).MaxMotorTorque) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+14980, __ccgo_ts+12321, int32FromInt32(750))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if int32FromUint8(b2IsValidFloat(tls, (*WheelJointDef)(unsafe.Pointer(def)).MotorSpeed)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+136, __ccgo_ts+13776, uintptr(unsafe.Pointer(&__func__29))))
		if v1 = !!(b2IsValidFloat(tls, (*WheelJointDef)(unsafe.Pointer(def)).MotorSpeed) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+13776, __ccgo_ts+12321, int32FromInt32(751))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return JointId{}
	}
	if v1 = !!((*WheelJointDef)(unsafe.Pointer(def)).LowerTranslation <= (*WheelJointDef)(unsafe.Pointer(def)).UpperTranslation); !v1 {
		b2InternalAssert(tls, __ccgo_ts+14830, __ccgo_ts+12321, int32FromInt32(752))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b2GetWorldFromId(tls, worldId)
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+12321, int32FromInt32(756))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return JointId{}
	}
	pair = b2CreateJoint(tls, world, def, int32(b2_wheelJoint))
	joint = pair.JointSim
	*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92)) = b2WheelJoint{}
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).PerpMass = float32FromFloat32(0)
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).AxialMass = float32FromFloat32(0)
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MotorImpulse = float32FromFloat32(0)
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LowerImpulse = float32FromFloat32(0)
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).UpperImpulse = float32FromFloat32(0)
	v11 = (*WheelJointDef)(unsafe.Pointer(def)).LowerTranslation
	v12 = (*WheelJointDef)(unsafe.Pointer(def)).UpperTranslation
	if v11 < v12 {
		v15 = v11
	} else {
		v15 = v12
	}
	v13 = v15
	goto _14
_14:
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LowerTranslation = v13
	v11 = (*WheelJointDef)(unsafe.Pointer(def)).LowerTranslation
	v12 = (*WheelJointDef)(unsafe.Pointer(def)).UpperTranslation
	if v11 > v12 {
		v15 = v11
	} else {
		v15 = v12
	}
	v13 = v15
	goto _19
_19:
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).UpperTranslation = v13
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxMotorTorque = (*WheelJointDef)(unsafe.Pointer(def)).MaxMotorTorque
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MotorSpeed = (*WheelJointDef)(unsafe.Pointer(def)).MotorSpeed
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).Hertz = (*WheelJointDef)(unsafe.Pointer(def)).Hertz
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).DampingRatio = (*WheelJointDef)(unsafe.Pointer(def)).DampingRatio
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableSpring = (*WheelJointDef)(unsafe.Pointer(def)).EnableSpring
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableLimit = (*WheelJointDef)(unsafe.Pointer(def)).EnableLimit
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableMotor = (*WheelJointDef)(unsafe.Pointer(def)).EnableMotor
	jointId = JointId{
		Index1:     (*b2JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_CreateWheelJoint)(unsafe.Add(nil, bp)) = b2RecArgs_CreateWheelJoint{
			World: worldId,
			Def:   *(*WheelJointDef)(unsafe.Add(nil, def)),
		}
		b2RecWriteRet_CreateWheelJoint(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp, jointId)
	}
	return jointId
}

func b2DestroyJointInternal(tls *_Stack, world uintptr, joint uintptr) {
	var bodyA, bodyB, edgeA, edgeB, movedJoint, movedJointSim, nextEdge, nextEdge1, nextJoint, nextJoint1, prevEdge, prevEdge1, prevJoint, prevJoint1, set, v10, v11 uintptr
	var edgeKeyA, edgeKeyB, idA, idB, jointId, localIndex, movedId, movedIndex, setIndex, v12, v13, v15 int32
	var v1, v2, v3, v4 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, edgeA, edgeB, edgeKeyA, edgeKeyB, idA, idB, jointId, localIndex, movedId, movedIndex, movedJoint, movedJointSim, nextEdge, nextEdge1, nextJoint, nextJoint1, prevEdge, prevEdge1, prevJoint, prevJoint1, set, setIndex, v1, v10, v11, v12, v13, v15, v2, v3, v4
	jointId = (*b2Joint)(unsafe.Pointer(joint)).JointId
	edgeA = joint + 20 + uintptr(0)*12
	edgeB = joint + 20 + uintptr(1)*12
	idA = (*b2JointEdge)(unsafe.Pointer(edgeA)).BodyId
	idB = (*b2JointEdge)(unsafe.Pointer(edgeB)).BodyId
	if v1 = !!(0 <= idA && idA < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8507, __ccgo_ts+12321, int32FromInt32(799))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idA)*104
	if v2 = !!(0 <= idB && idB < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+8557, __ccgo_ts+12321, int32FromInt32(800))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idB)*104
	// Remove from body A
	if (*b2JointEdge)(unsafe.Pointer(edgeA)).PrevKey != -int32(1) {
		if v3 = !!(0 <= (*b2JointEdge)(unsafe.Pointer(edgeA)).PrevKey>>int32(1) && (*b2JointEdge)(unsafe.Pointer(edgeA)).PrevKey>>int32(1) < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v3 {
			b2InternalAssert(tls, __ccgo_ts+15053, __ccgo_ts+12321, int32FromInt32(805))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		prevJoint = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr((*b2JointEdge)(unsafe.Pointer(edgeA)).PrevKey>>int32FromInt32(1))*72
		prevEdge = prevJoint + 20 + uintptr((*b2JointEdge)(unsafe.Pointer(edgeA)).PrevKey&int32FromInt32(1))*12
		(*b2JointEdge)(unsafe.Pointer(prevEdge)).NextKey = (*b2JointEdge)(unsafe.Pointer(edgeA)).NextKey
	}
	if (*b2JointEdge)(unsafe.Pointer(edgeA)).NextKey != -int32(1) {
		if v1 = !!(0 <= (*b2JointEdge)(unsafe.Pointer(edgeA)).NextKey>>int32(1) && (*b2JointEdge)(unsafe.Pointer(edgeA)).NextKey>>int32(1) < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+15135, __ccgo_ts+12321, int32FromInt32(812))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		nextJoint = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr((*b2JointEdge)(unsafe.Pointer(edgeA)).NextKey>>int32FromInt32(1))*72
		nextEdge = nextJoint + 20 + uintptr((*b2JointEdge)(unsafe.Pointer(edgeA)).NextKey&int32FromInt32(1))*12
		(*b2JointEdge)(unsafe.Pointer(nextEdge)).PrevKey = (*b2JointEdge)(unsafe.Pointer(edgeA)).PrevKey
	}
	edgeKeyA = jointId<<int32(1) | 0
	if (*b2Body)(unsafe.Pointer(bodyA)).HeadJointKey == edgeKeyA {
		(*b2Body)(unsafe.Pointer(bodyA)).HeadJointKey = (*b2JointEdge)(unsafe.Pointer(edgeA)).NextKey
	}
	*(*int32)(unsafe.Add(nil, bodyA+40)) -= int32(1)
	// Remove from body B
	if (*b2JointEdge)(unsafe.Pointer(edgeB)).PrevKey != -int32(1) {
		if v1 = !!(0 <= (*b2JointEdge)(unsafe.Pointer(edgeB)).PrevKey>>int32(1) && (*b2JointEdge)(unsafe.Pointer(edgeB)).PrevKey>>int32(1) < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+15217, __ccgo_ts+12321, int32FromInt32(828))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		prevJoint1 = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr((*b2JointEdge)(unsafe.Pointer(edgeB)).PrevKey>>int32FromInt32(1))*72
		prevEdge1 = prevJoint1 + 20 + uintptr((*b2JointEdge)(unsafe.Pointer(edgeB)).PrevKey&int32FromInt32(1))*12
		(*b2JointEdge)(unsafe.Pointer(prevEdge1)).NextKey = (*b2JointEdge)(unsafe.Pointer(edgeB)).NextKey
	}
	if (*b2JointEdge)(unsafe.Pointer(edgeB)).NextKey != -int32(1) {
		if v1 = !!(0 <= (*b2JointEdge)(unsafe.Pointer(edgeB)).NextKey>>int32(1) && (*b2JointEdge)(unsafe.Pointer(edgeB)).NextKey>>int32(1) < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+15299, __ccgo_ts+12321, int32FromInt32(835))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		nextJoint1 = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr((*b2JointEdge)(unsafe.Pointer(edgeB)).NextKey>>int32FromInt32(1))*72
		nextEdge1 = nextJoint1 + 20 + uintptr((*b2JointEdge)(unsafe.Pointer(edgeB)).NextKey&int32FromInt32(1))*12
		(*b2JointEdge)(unsafe.Pointer(nextEdge1)).PrevKey = (*b2JointEdge)(unsafe.Pointer(edgeB)).PrevKey
	}
	edgeKeyB = jointId<<int32(1) | int32(1)
	if (*b2Body)(unsafe.Pointer(bodyB)).HeadJointKey == edgeKeyB {
		(*b2Body)(unsafe.Pointer(bodyB)).HeadJointKey = (*b2JointEdge)(unsafe.Pointer(edgeB)).NextKey
	}
	*(*int32)(unsafe.Add(nil, bodyB+40)) -= int32(1)
	if (*b2Joint)(unsafe.Pointer(joint)).IslandId != -int32(1) {
		if v1 = !!((*b2Joint)(unsafe.Pointer(joint)).SetIndex > int32(b2_disabledSet)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+15381, __ccgo_ts+12321, int32FromInt32(850))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		b2UnlinkJoint(tls, world, joint)
	} else {
		if v1 = !!((*b2Joint)(unsafe.Pointer(joint)).SetIndex <= int32(b2_disabledSet)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+15414, __ccgo_ts+12321, int32FromInt32(855))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
	}
	// Remove joint from solver set that owns it
	setIndex = (*b2Joint)(unsafe.Pointer(joint)).SetIndex
	localIndex = (*b2Joint)(unsafe.Pointer(joint)).LocalIndex
	if setIndex == int32(b2_awakeSet) {
		b2RemoveJointFromGraph(tls, world, (*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId, (*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId, (*b2Joint)(unsafe.Pointer(joint)).ColorIndex, localIndex)
	} else {
		if v1 = !!(0 <= setIndex && setIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+6188, __ccgo_ts+12321, int32FromInt32(868))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		set = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setIndex)*88
		v10 = (*b2SolverSet)(unsafe.Pointer(set)).JointSims.Data
		v11 = set + 32 + 8
		v12 = localIndex
		v13 = int32(252)
		if v2 = !!(0 <= v12 && v12 < *(*int32)(unsafe.Add(nil, v11)) && bool(__ccgo_ts+2029 != 0)); !v2 {
			b2InternalAssert(tls, __ccgo_ts+2055, __ccgo_ts+2115, int32FromInt32(153))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		*(*int32)(unsafe.Add(nil, v11)) = *(*int32)(unsafe.Add(nil, v11)) - 1
		if v12 != *(*int32)(unsafe.Add(nil, v11)) {
			memcpy(tls, v10+uintptr(v12*v13), v10+uintptr(*(*int32)(unsafe.Add(nil, v11))*v13), uint64FromInt32(v13))
			v15 = *(*int32)(unsafe.Add(nil, v11))
			goto _16
		}
		v15 = -int32(1)
		goto _16
	_16:
		movedIndex = v15
		if movedIndex != -int32(1) {
			// Fix moved joint
			movedJointSim = (*b2SolverSet)(unsafe.Pointer(set)).JointSims.Data + uintptr(localIndex)*252
			movedId = (*b2JointSim)(unsafe.Pointer(movedJointSim)).JointId
			if v3 = !!(0 <= movedId && movedId < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v3 {
				b2InternalAssert(tls, __ccgo_ts+5612, __ccgo_ts+12321, int32FromInt32(875))
			}
			_ = v3 || bool(int32FromInt32(0) != 0)
			movedJoint = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr(movedId)*72
			if v4 = !!((*b2Joint)(unsafe.Pointer(movedJoint)).LocalIndex == movedIndex); !v4 {
				b2InternalAssert(tls, __ccgo_ts+5743, __ccgo_ts+12321, int32FromInt32(876))
			}
			_ = v4 || bool(int32FromInt32(0) != 0)
			(*b2Joint)(unsafe.Pointer(movedJoint)).LocalIndex = localIndex
		}
	}
	// Free joint and id (preserve joint generation)
	(*b2Joint)(unsafe.Pointer(joint)).SetIndex = -int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).LocalIndex = -int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).ColorIndex = -int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).JointId = -int32(1)
	b2FreeId(tls, world+1992, jointId)
	b2WakeBody(tls, world, bodyA)
	b2WakeBody(tls, world, bodyB)
	b2ValidateSolverSets(tls, world)
}

func b2DestroyJoint(tls *_Stack, jointId JointId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_DestroyJoint
	_, _, _ = joint, world, v1
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if v1 = !!(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1583, __ccgo_ts+12321, int32FromInt32(897))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_DestroyJoint)(unsafe.Add(nil, bp)) = b2RecArgs_DestroyJoint{
			Joint: jointId,
		}
		b2RecWrite_DestroyJoint(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointFullId(tls, world, jointId)
	b2DestroyJointInternal(tls, world, joint)
}

func b2Joint_GetType(tls *_Stack, jointId JointId) (r JointType) {
	var joint, world uintptr
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	return (*b2Joint)(unsafe.Pointer(joint)).Type1
}

func b2Joint_GetBodyA(tls *_Stack, jointId JointId) (r BodyId) {
	var joint, world uintptr
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	return b2MakeBodyId(tls, world, (*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId)
}

func b2Joint_GetBodyB(tls *_Stack, jointId JointId) (r BodyId) {
	var joint, world uintptr
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	return b2MakeBodyId(tls, world, (*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId)
}

func b2Joint_GetWorld(tls *_Stack, jointId JointId) (r WorldId) {
	var world uintptr
	_ = world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	return WorldId{
		Index1:     uint16FromInt32(int32FromUint16(jointId.World0) + int32(1)),
		Generation: (*b2World)(unsafe.Pointer(world)).Generation,
	}
}

func b2Joint_SetLocalFrameA(tls *_Stack, jointId JointId, localFrame Transform) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, jointSim, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_JointSetLocalFrameA
	_, _, _, _ = joint, jointSim, world, v1
	if int32FromUint8(b2IsValidTransform(tls, localFrame)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+32, __ccgo_ts+15448, uintptr(unsafe.Pointer(&__func__30))))
		if v1 = !!(b2IsValidTransform(tls, localFrame) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+15448, __ccgo_ts+12321, int32FromInt32(940))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_JointSetLocalFrameA)(unsafe.Add(nil, bp)) = b2RecArgs_JointSetLocalFrameA{
			Joint:      jointId,
			LocalFrame: localFrame,
		}
		b2RecWrite_JointSetLocalFrameA(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointFullId(tls, world, jointId)
	jointSim = b2GetJointSim(tls, world, joint)
	(*b2JointSim)(unsafe.Pointer(jointSim)).LocalFrameA = localFrame
}

func b2Joint_GetLocalFrameA(tls *_Stack, jointId JointId) (r Transform) {
	var joint, jointSim, world uintptr
	_, _, _ = joint, jointSim, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	jointSim = b2GetJointSim(tls, world, joint)
	return (*b2JointSim)(unsafe.Pointer(jointSim)).LocalFrameA
}

func b2Joint_SetLocalFrameB(tls *_Stack, jointId JointId, localFrame Transform) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, jointSim, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_JointSetLocalFrameB
	_, _, _, _ = joint, jointSim, world, v1
	if int32FromUint8(b2IsValidTransform(tls, localFrame)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+32, __ccgo_ts+15448, uintptr(unsafe.Pointer(&__func__34))))
		if v1 = !!(b2IsValidTransform(tls, localFrame) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+15448, __ccgo_ts+12321, int32FromInt32(959))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_JointSetLocalFrameB)(unsafe.Add(nil, bp)) = b2RecArgs_JointSetLocalFrameB{
			Joint:      jointId,
			LocalFrame: localFrame,
		}
		b2RecWrite_JointSetLocalFrameB(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointFullId(tls, world, jointId)
	jointSim = b2GetJointSim(tls, world, joint)
	(*b2JointSim)(unsafe.Pointer(jointSim)).LocalFrameB = localFrame
}

func b2Joint_GetLocalFrameB(tls *_Stack, jointId JointId) (r Transform) {
	var joint, jointSim, world uintptr
	_, _, _ = joint, jointSim, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	jointSim = b2GetJointSim(tls, world, joint)
	return (*b2JointSim)(unsafe.Pointer(jointSim)).LocalFrameB
}

func b2Joint_SetCollideConnected(tls *_Stack, jointId JointId, shouldCollide uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var bodyA, bodyB, joint, shape, world uintptr
	var proxyId, shapeId, v3 int32
	var proxyType, typeA, typeB b2BodyType1
	var v1, v2 bool
	var _ /* _a at bp+0 */ b2RecArgs_JointSetCollideConnected
	_, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, joint, proxyId, proxyType, shape, shapeId, typeA, typeB, world, v1, v2, v3
	world = b2GetWorldLocked(tls, int32FromUint16(jointId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_JointSetCollideConnected)(unsafe.Add(nil, bp)) = b2RecArgs_JointSetCollideConnected{
			Joint:         jointId,
			ShouldCollide: shouldCollide,
		}
		b2RecWrite_JointSetCollideConnected(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointFullId(tls, world, jointId)
	if int32FromUint8((*b2Joint)(unsafe.Pointer(joint)).CollideConnected) == int32FromUint8(shouldCollide) {
		return
	}
	(*b2Joint)(unsafe.Pointer(joint)).CollideConnected = shouldCollide
	if v1 = !!(0 <= (*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId && (*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+3036, __ccgo_ts+12321, int32FromInt32(994))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId)*104
	if v2 = !!(0 <= (*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId && (*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+3124, __ccgo_ts+12321, int32FromInt32(995))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId)*104
	if shouldCollide != 0 {
		typeA = (*b2Body)(unsafe.Pointer(bodyA)).Type1
		typeB = (*b2Body)(unsafe.Pointer(bodyB)).Type1
		shapeId = -int32(1)
		if typeA == int32(b2_dynamicBody) && typeB == int32(b2_dynamicBody) {
			if (*b2Body)(unsafe.Pointer(bodyA)).ShapeCount < (*b2Body)(unsafe.Pointer(bodyB)).ShapeCount {
				v3 = (*b2Body)(unsafe.Pointer(bodyA)).HeadShapeId
			} else {
				v3 = (*b2Body)(unsafe.Pointer(bodyB)).HeadShapeId
			}
			shapeId = v3
		} else {
			if typeA == int32(b2_dynamicBody) {
				shapeId = (*b2Body)(unsafe.Pointer(bodyA)).HeadShapeId
			} else {
				if typeB == int32(b2_dynamicBody) {
					shapeId = (*b2Body)(unsafe.Pointer(bodyB)).HeadShapeId
				}
			}
		}
		for shapeId != -int32(1) {
			if v1 = !!(0 <= shapeId && shapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
				b2InternalAssert(tls, __ccgo_ts+1908, __ccgo_ts+12321, int32FromInt32(1018))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			shape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*280
			if (*b2Shape)(unsafe.Pointer(shape)).ProxyKey != -int32(1) {
				v3 = (*b2Shape)(unsafe.Pointer(shape)).ProxyKey
				proxyType = v3 & int32FromInt32(3)
				proxyId = v3 >> int32(2)
				b2DynamicTree_MarkProxyMovedSerial(tls, world+40+uintptr(proxyType)*104, proxyId)
			}
			shapeId = (*b2Shape)(unsafe.Pointer(shape)).NextShapeId
		}
	} else {
		b2DestroyContactsBetweenBodies(tls, world, bodyA, bodyB)
	}
}

func b2Joint_GetCollideConnected(tls *_Stack, jointId JointId) (r uint8) {
	var joint, world uintptr
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	return (*b2Joint)(unsafe.Pointer(joint)).CollideConnected
}

func b2Joint_SetUserData(tls *_Stack, jointId JointId, userData uintptr) {
	var joint, world uintptr
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	(*b2Joint)(unsafe.Pointer(joint)).UserData = userData
}

func b2Joint_GetUserData(tls *_Stack, jointId JointId) (r uintptr) {
	var joint, world uintptr
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	return (*b2Joint)(unsafe.Pointer(joint)).UserData
}

func b2Joint_WakeBodies(tls *_Stack, jointId JointId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var bodyA, bodyB, joint, world uintptr
	var v1, v2 bool
	var _ /* _a at bp+0 */ b2RecArgs_JointWakeBodies
	_, _, _, _, _, _ = bodyA, bodyB, joint, world, v1, v2
	world = b2GetWorldLocked(tls, int32FromUint16(jointId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_JointWakeBodies)(unsafe.Add(nil, bp)) = b2RecArgs_JointWakeBodies{
			Joint: jointId,
		}
		b2RecWrite_JointWakeBodies(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointFullId(tls, world, jointId)
	if v1 = !!(0 <= (*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId && (*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+3036, __ccgo_ts+12321, int32FromInt32(1066))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId)*104
	if v2 = !!(0 <= (*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId && (*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+3124, __ccgo_ts+12321, int32FromInt32(1067))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId)*104
	b2WakeBody(tls, world, bodyA)
	b2WakeBody(tls, world, bodyB)
}

func b2GetJointReaction(tls *_Stack, sim uintptr, invTimeStep float32, force uintptr, torque uintptr) {
	var angularImpulse, axialImpulse, axialImpulse1, linearImpulse, perpImpulse, perpImpulse1, v1, v2, v4 float32
	var joint, joint1, joint2, joint3, joint4, joint5, joint6, joint7 uintptr
	var v5, v6, v7, v9 Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angularImpulse, axialImpulse, axialImpulse1, joint, joint1, joint2, joint3, joint4, joint5, joint6, joint7, linearImpulse, perpImpulse, perpImpulse1, v1, v2, v4, v5, v6, v7, v9
	linearImpulse = float32FromFloat32(0)
	angularImpulse = float32FromFloat32(0)
	switch (*b2JointSim)(unsafe.Pointer(sim)).Type1 {
	case int32(b2_distanceJoint):
		joint = sim + 92
		v1 = (*b2DistanceJoint)(unsafe.Pointer(joint)).Impulse + (*b2DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b2DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse + (*b2DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse
		if v1 < float32FromInt32(0) {
			v4 = -v1
		} else {
			v4 = v1
		}
		v2 = v4
		goto _3
	_3:
		linearImpulse = v2
	case int32(b2_motorJoint):
		joint1 = sim + 92
		v5 = (*b2MotorJoint)(unsafe.Pointer(joint1)).LinearVelocityImpulse
		v6 = (*b2MotorJoint)(unsafe.Pointer(joint1)).LinearSpringImpulse
		v7 = Vec2{
			X: v5.X + v6.X,
			Y: v5.Y + v6.Y,
		}
		goto _8
	_8:
		v9 = v7
		v1 = sqrtf(tls, float32(v9.X*v9.X)+float32(v9.Y*v9.Y))
		goto _11
	_11:
		linearImpulse = v1
		v1 = (*b2MotorJoint)(unsafe.Pointer(joint1)).AngularVelocityImpulse + (*b2MotorJoint)(unsafe.Pointer(joint1)).AngularSpringImpulse
		if v1 < float32FromInt32(0) {
			v4 = -v1
		} else {
			v4 = v1
		}
		v2 = v4
		goto _14
	_14:
		angularImpulse = v2
	case int32(b2_moverJoint):
		joint2 = sim + 92
		v5 = (*b2MoverJoint)(unsafe.Pointer(joint2)).LinearVelocityImpulse
		v1 = sqrtf(tls, float32(v5.X*v5.X)+float32(v5.Y*v5.Y))
		goto _18
	_18:
		linearImpulse = v1
	case int32(b2_pogoJoint):
		joint3 = sim + 92
		v1 = (*b2PogoJoint)(unsafe.Pointer(joint3)).Impulse
		if v1 < float32FromInt32(0) {
			v4 = -v1
		} else {
			v4 = v1
		}
		v2 = v4
		goto _21
	_21:
		linearImpulse = v2
	case int32(b2_prismaticJoint):
		joint4 = sim + 92
		perpImpulse = (*b2PrismaticJoint)(unsafe.Pointer(joint4)).Impulse.X
		axialImpulse = (*b2PrismaticJoint)(unsafe.Pointer(joint4)).MotorImpulse + (*b2PrismaticJoint)(unsafe.Pointer(joint4)).LowerImpulse - (*b2PrismaticJoint)(unsafe.Pointer(joint4)).UpperImpulse
		linearImpulse = sqrtf(tls, float32(perpImpulse*perpImpulse)+float32(axialImpulse*axialImpulse))
		v1 = (*b2PrismaticJoint)(unsafe.Pointer(joint4)).Impulse.Y
		if v1 < float32FromInt32(0) {
			v4 = -v1
		} else {
			v4 = v1
		}
		v2 = v4
		goto _25
	_25:
		angularImpulse = v2
	case int32(b2_revoluteJoint):
		joint5 = sim + 92
		v5 = (*b2RevoluteJoint)(unsafe.Pointer(joint5)).LinearImpulse
		v1 = sqrtf(tls, float32(v5.X*v5.X)+float32(v5.Y*v5.Y))
		goto _29
	_29:
		linearImpulse = v1
		v1 = (*b2RevoluteJoint)(unsafe.Pointer(joint5)).MotorImpulse + (*b2RevoluteJoint)(unsafe.Pointer(joint5)).LowerImpulse - (*b2RevoluteJoint)(unsafe.Pointer(joint5)).UpperImpulse
		if v1 < float32FromInt32(0) {
			v4 = -v1
		} else {
			v4 = v1
		}
		v2 = v4
		goto _32
	_32:
		angularImpulse = v2
	case int32(b2_weldJoint):
		joint6 = sim + 92
		v5 = (*b2WeldJoint)(unsafe.Pointer(joint6)).LinearImpulse
		v1 = sqrtf(tls, float32(v5.X*v5.X)+float32(v5.Y*v5.Y))
		goto _36
	_36:
		linearImpulse = v1
		v1 = (*b2WeldJoint)(unsafe.Pointer(joint6)).AngularImpulse
		if v1 < float32FromInt32(0) {
			v4 = -v1
		} else {
			v4 = v1
		}
		v2 = v4
		goto _39
	_39:
		angularImpulse = v2
	case int32(b2_wheelJoint):
		joint7 = sim + 92
		perpImpulse1 = (*b2WheelJoint)(unsafe.Pointer(joint7)).PerpImpulse
		axialImpulse1 = (*b2WheelJoint)(unsafe.Pointer(joint7)).SpringImpulse + (*b2WheelJoint)(unsafe.Pointer(joint7)).LowerImpulse - (*b2WheelJoint)(unsafe.Pointer(joint7)).UpperImpulse
		linearImpulse = sqrtf(tls, float32(perpImpulse1*perpImpulse1)+float32(axialImpulse1*axialImpulse1))
		v1 = (*b2WheelJoint)(unsafe.Pointer(joint7)).MotorImpulse
		if v1 < float32FromInt32(0) {
			v4 = -v1
		} else {
			v4 = v1
		}
		v2 = v4
		goto _43
	_43:
		angularImpulse = v2
	default:
		break
	}
	*(*float32)(unsafe.Add(nil, force)) = float32(linearImpulse * invTimeStep)
	*(*float32)(unsafe.Add(nil, torque)) = float32(angularImpulse * invTimeStep)
}

func b2GetJointConstraintForce(tls *_Stack, world uintptr, joint uintptr) (r Vec2) {
	var base uintptr
	var v1 bool
	_, _ = base, v1
	base = b2GetJointSim(tls, world, joint)
	switch (*b2Joint)(unsafe.Pointer(joint)).Type1 {
	case int32(b2_distanceJoint):
		return b2GetDistanceJointForce(tls, world, base)
	case int32(b2_motorJoint):
		return b2GetMotorJointForce(tls, world, base)
	case int32(b2_moverJoint):
		return b2GetMoverJointForce(tls, world, base)
	case int32(b2_filterJoint):
		return b2Vec2_zero
	case int32(b2_pogoJoint):
		return b2GetPogoJointForce(tls, world, base)
	case int32(b2_prismaticJoint):
		return b2GetPrismaticJointForce(tls, world, base)
	case int32(b2_revoluteJoint):
		return b2GetRevoluteJointForce(tls, world, base)
	case int32(b2_weldJoint):
		return b2GetWeldJointForce(tls, world, base)
	case int32(b2_wheelJoint):
		return b2GetWheelJointForce(tls, world, base)
	default:
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+7715, __ccgo_ts+12321, int32FromInt32(1188))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return b2Vec2_zero
	}
}

func b2GetJointConstraintTorque(tls *_Stack, world uintptr, joint uintptr) (r float32) {
	var base uintptr
	var v1 bool
	_, _ = base, v1
	base = b2GetJointSim(tls, world, joint)
	switch (*b2Joint)(unsafe.Pointer(joint)).Type1 {
	case int32(b2_distanceJoint):
		return float32FromFloat32(0)
	case int32(b2_motorJoint):
		return b2GetMotorJointTorque(tls, world, base)
	case int32(b2_filterJoint):
		return float32FromFloat32(0)
	case int32(b2_prismaticJoint):
		return b2GetPrismaticJointTorque(tls, world, base)
	case int32(b2_revoluteJoint):
		return b2GetRevoluteJointTorque(tls, world, base)
	case int32(b2_weldJoint):
		return b2GetWeldJointTorque(tls, world, base)
	case int32(b2_wheelJoint):
		return b2GetWheelJointTorque(tls, world, base)
	default:
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+7715, __ccgo_ts+12321, int32FromInt32(1221))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return float32FromFloat32(0)
	}
}

func b2Joint_GetConstraintForce(tls *_Stack, jointId JointId) (r Vec2) {
	var joint, world uintptr
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	return b2GetJointConstraintForce(tls, world, joint)
}

func b2Joint_GetConstraintTorque(tls *_Stack, jointId JointId) (r float32) {
	var joint, world uintptr
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	return b2GetJointConstraintTorque(tls, world, joint)
}

func b2Joint_GetLinearSeparation(tls *_Stack, jointId JointId) (r1 float32) {
	var axisA, axisA1, dp, pA, pB, perpA, perpA1, v10, v11, v14, v15, v17, v18, v19, v21 Vec2
	var base1, distanceJoint, joint, prismaticJoint, weldJoint, wheelJoint, world uintptr
	var length, limitSeparation, limitSeparation1, perpendicularSeparation, perpendicularSeparation1, translation, translation1, x, y, v22, v24, v25, v27, v37 float32
	var r, xfA, xfB, v13, v3, v7, v9 Transform
	var wxfA, v1, v5 WorldTransform
	var v2, v6 Pos
	var v28 Rot
	var v72 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axisA, axisA1, base1, distanceJoint, dp, joint, length, limitSeparation, limitSeparation1, pA, pB, perpA, perpA1, perpendicularSeparation, perpendicularSeparation1, prismaticJoint, r, translation, translation1, weldJoint, wheelJoint, world, wxfA, x, xfA, xfB, y, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v24, v25, v27, v28, v3, v37, v5, v6, v7, v72, v9
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	base1 = b2GetJointSim(tls, world, joint)
	// Relative to body A so the difference stays in float precision far from the origin
	wxfA = b2GetBodyTransform(tls, world, (*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId)
	v1 = wxfA
	v2 = wxfA.P
	r.Q = v1.Q
	r.P = Vec2{
		X: v1.P.X - v2.X,
		Y: v1.P.Y - v2.Y,
	}
	v3 = r
	goto _4
_4:
	xfA = v3
	v5 = b2GetBodyTransform(tls, world, (*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId)
	v6 = wxfA.P
	r.Q = v5.Q
	r.P = Vec2{
		X: v5.P.X - v6.X,
		Y: v5.P.Y - v6.Y,
	}
	v7 = r
	goto _8
_8:
	xfB = v7
	v9 = xfA
	v10 = (*b2JointSim)(unsafe.Pointer(base1)).LocalFrameA.P
	x = float32(v9.Q.C*v10.X) - float32(v9.Q.S*v10.Y) + v9.P.X
	y = float32(v9.Q.S*v10.X) + float32(v9.Q.C*v10.Y) + v9.P.Y
	v11 = Vec2{
		X: x,
		Y: y,
	}
	goto _12
_12:
	pA = v11
	v13 = xfB
	v14 = (*b2JointSim)(unsafe.Pointer(base1)).LocalFrameB.P
	x = float32(v13.Q.C*v14.X) - float32(v13.Q.S*v14.Y) + v13.P.X
	y = float32(v13.Q.S*v14.X) + float32(v13.Q.C*v14.Y) + v13.P.Y
	v15 = Vec2{
		X: x,
		Y: y,
	}
	goto _16
_16:
	pB = v15
	v17 = pB
	v18 = pA
	v19 = Vec2{
		X: v17.X - v18.X,
		Y: v17.Y - v18.Y,
	}
	goto _20
_20:
	dp = v19
	switch (*b2Joint)(unsafe.Pointer(joint)).Type1 {
	case int32(b2_distanceJoint):
		distanceJoint = base1 + 92
		v21 = dp
		v22 = sqrtf(tls, float32(v21.X*v21.X)+float32(v21.Y*v21.Y))
		goto _23
	_23:
		length = v22
		if (*b2DistanceJoint)(unsafe.Pointer(distanceJoint)).EnableSpring != 0 {
			if (*b2DistanceJoint)(unsafe.Pointer(distanceJoint)).EnableLimit != 0 {
				if length < (*b2DistanceJoint)(unsafe.Pointer(distanceJoint)).MinLength {
					return (*b2DistanceJoint)(unsafe.Pointer(distanceJoint)).MinLength - length
				}
				if length > (*b2DistanceJoint)(unsafe.Pointer(distanceJoint)).MaxLength {
					return length - (*b2DistanceJoint)(unsafe.Pointer(distanceJoint)).MaxLength
				}
				return float32FromFloat32(0)
			}
			return float32FromFloat32(0)
		}
		v22 = length - (*b2DistanceJoint)(unsafe.Pointer(distanceJoint)).Length
		if v22 < float32FromInt32(0) {
			v25 = -v22
		} else {
			v25 = v22
		}
		v24 = v25
		goto _26
	_26:
		return v24
	case int32(b2_motorJoint):
		return float32FromFloat32(0)
	case int32(b2_filterJoint):
		return float32FromFloat32(0)
	case int32(b2_prismaticJoint):
		prismaticJoint = base1 + 92
		v28 = xfA.Q
		v10 = Vec2{
			X: float32FromFloat32(1),
		}
		v11 = Vec2{
			X: float32(v28.C*v10.X) - float32(v28.S*v10.Y),
			Y: float32(v28.S*v10.X) + float32(v28.C*v10.Y),
		}
		goto _31
	_31:
		axisA = v11
		v14 = axisA
		v15 = Vec2{
			X: -v14.Y,
			Y: v14.X,
		}
		goto _34
	_34:
		perpA = v15
		v17 = perpA
		v18 = dp
		v22 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y)
		goto _38
	_38:
		v24 = v22
		if v24 < float32FromInt32(0) {
			v27 = -v24
		} else {
			v27 = v24
		}
		v25 = v27
		goto _41
	_41:
		perpendicularSeparation = v25
		limitSeparation = float32FromFloat32(0)
		if (*b2PrismaticJoint)(unsafe.Pointer(prismaticJoint)).EnableLimit != 0 {
			v19 = axisA
			v21 = dp
			v37 = float32(v19.X*v21.X) + float32(v19.Y*v21.Y)
			goto _46
		_46:
			translation = v37
			if translation < (*b2PrismaticJoint)(unsafe.Pointer(prismaticJoint)).LowerTranslation {
				limitSeparation = (*b2PrismaticJoint)(unsafe.Pointer(prismaticJoint)).LowerTranslation - translation
			}
			if (*b2PrismaticJoint)(unsafe.Pointer(prismaticJoint)).UpperTranslation < translation {
				limitSeparation = translation - (*b2PrismaticJoint)(unsafe.Pointer(prismaticJoint)).UpperTranslation
			}
		}
		return sqrtf(tls, float32(perpendicularSeparation*perpendicularSeparation)+float32(limitSeparation*limitSeparation))
	case int32(b2_revoluteJoint):
		v10 = dp
		v22 = sqrtf(tls, float32(v10.X*v10.X)+float32(v10.Y*v10.Y))
		goto _49
	_49:
		return v22
	case int32(b2_weldJoint):
		weldJoint = base1 + 92
		if (*b2WeldJoint)(unsafe.Pointer(weldJoint)).LinearHertz == float32FromFloat32(0) {
			v10 = dp
			v22 = sqrtf(tls, float32(v10.X*v10.X)+float32(v10.Y*v10.Y))
			goto _52
		_52:
			return v22
		}
		return float32FromFloat32(0)
	case int32(b2_wheelJoint):
		wheelJoint = base1 + 92
		v28 = xfA.Q
		v10 = Vec2{
			X: float32FromFloat32(1),
		}
		v11 = Vec2{
			X: float32(v28.C*v10.X) - float32(v28.S*v10.Y),
			Y: float32(v28.S*v10.X) + float32(v28.C*v10.Y),
		}
		goto _56
	_56:
		axisA1 = v11
		v14 = axisA1
		v15 = Vec2{
			X: -v14.Y,
			Y: v14.X,
		}
		goto _59
	_59:
		perpA1 = v15
		v17 = perpA1
		v18 = dp
		v22 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y)
		goto _63
	_63:
		v24 = v22
		if v24 < float32FromInt32(0) {
			v27 = -v24
		} else {
			v27 = v24
		}
		v25 = v27
		goto _66
	_66:
		perpendicularSeparation1 = v25
		limitSeparation1 = float32FromFloat32(0)
		if (*b2WheelJoint)(unsafe.Pointer(wheelJoint)).EnableLimit != 0 {
			v19 = axisA1
			v21 = dp
			v37 = float32(v19.X*v21.X) + float32(v19.Y*v21.Y)
			goto _71
		_71:
			translation1 = v37
			if translation1 < (*b2WheelJoint)(unsafe.Pointer(wheelJoint)).LowerTranslation {
				limitSeparation1 = (*b2WheelJoint)(unsafe.Pointer(wheelJoint)).LowerTranslation - translation1
			}
			if (*b2WheelJoint)(unsafe.Pointer(wheelJoint)).UpperTranslation < translation1 {
				limitSeparation1 = translation1 - (*b2WheelJoint)(unsafe.Pointer(wheelJoint)).UpperTranslation
			}
		}
		return sqrtf(tls, float32(perpendicularSeparation1*perpendicularSeparation1)+float32(limitSeparation1*limitSeparation1))
	default:
		if v72 = bool(!!(int32FromInt32(false1) != 0)); !v72 {
			b2InternalAssert(tls, __ccgo_ts+7715, __ccgo_ts+12321, int32FromInt32(1355))
		}
		_ = v72 || bool(int32FromInt32(0) != 0)
		return float32FromFloat32(0)
	}
}

func b2Joint_GetAngularSeparation(tls *_Stack, jointId JointId) (r float32) {
	var angle, c, relativeAngle, s, v3 float32
	var base, joint, revoluteJoint, weldJoint, world uintptr
	var qA, qB, v1, v2 Rot
	var v5 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angle, base, c, joint, qA, qB, relativeAngle, revoluteJoint, s, weldJoint, world, v1, v2, v3, v5
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	base = b2GetJointSim(tls, world, joint)
	qA = b2GetBodyTransform(tls, world, (*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId).Q
	qB = b2GetBodyTransform(tls, world, (*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId).Q
	v1 = qA
	v2 = qB
	s = float32(v1.C*v2.S) - float32(v1.S*v2.C)
	c = float32(v1.C*v2.C) + float32(v1.S*v2.S)
	v3 = b2Atan2(tls, s, c)
	goto _4
_4:
	relativeAngle = v3
	switch (*b2Joint)(unsafe.Pointer(joint)).Type1 {
	case int32(b2_distanceJoint):
		return float32FromFloat32(0)
	case int32(b2_motorJoint):
		return float32FromFloat32(0)
	case int32(b2_filterJoint):
		return float32FromFloat32(0)
	case int32(b2_prismaticJoint):
		return relativeAngle
	case int32(b2_revoluteJoint):
		revoluteJoint = base + 92
		if (*b2RevoluteJoint)(unsafe.Pointer(revoluteJoint)).EnableLimit != 0 {
			angle = relativeAngle
			if angle < (*b2RevoluteJoint)(unsafe.Pointer(revoluteJoint)).LowerAngle {
				return (*b2RevoluteJoint)(unsafe.Pointer(revoluteJoint)).LowerAngle - angle
			}
			if (*b2RevoluteJoint)(unsafe.Pointer(revoluteJoint)).UpperAngle < angle {
				return angle - (*b2RevoluteJoint)(unsafe.Pointer(revoluteJoint)).UpperAngle
			}
		}
		return float32FromFloat32(0)
	case int32(b2_weldJoint):
		weldJoint = base + 92
		if (*b2WeldJoint)(unsafe.Pointer(weldJoint)).AngularHertz == float32FromFloat32(0) {
			return relativeAngle
		}
		return float32FromFloat32(0)
	case int32(b2_wheelJoint):
		return float32FromFloat32(0)
	default:
		if v5 = bool(!!(int32FromInt32(false1) != 0)); !v5 {
			b2InternalAssert(tls, __ccgo_ts+7715, __ccgo_ts+12321, int32FromInt32(1421))
		}
		_ = v5 || bool(int32FromInt32(0) != 0)
		return float32FromFloat32(0)
	}
}

func b2Joint_SetConstraintTuning(tls *_Stack, jointId JointId, hertz float32, dampingRatio float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var base, joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_JointSetConstraintTuning
	_, _, _, _ = base, joint, world, v1
	if boolInt32(b2IsValidFloat(tls, hertz) != 0 && hertz >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+15481, uintptr(unsafe.Pointer(&__func__35))))
		if v1 = !!(b2IsValidFloat(tls, hertz) != 0 && hertz >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+15481, __ccgo_ts+12321, int32FromInt32(1428))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if boolInt32(b2IsValidFloat(tls, dampingRatio) != 0 && dampingRatio >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+15522, uintptr(unsafe.Pointer(&__func__35))))
		if v1 = !!(b2IsValidFloat(tls, dampingRatio) != 0 && dampingRatio >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+15522, __ccgo_ts+12321, int32FromInt32(1429))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_JointSetConstraintTuning)(unsafe.Add(nil, bp)) = b2RecArgs_JointSetConstraintTuning{
			Joint:        jointId,
			Hertz:        hertz,
			DampingRatio: dampingRatio,
		}
		b2RecWrite_JointSetConstraintTuning(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointFullId(tls, world, jointId)
	base = b2GetJointSim(tls, world, joint)
	(*b2JointSim)(unsafe.Pointer(base)).ConstraintHertz = hertz
	(*b2JointSim)(unsafe.Pointer(base)).ConstraintDampingRatio = dampingRatio
}

func b2Joint_GetConstraintTuning(tls *_Stack, jointId JointId, hertz uintptr, dampingRatio uintptr) {
	var base, joint, world uintptr
	_, _, _ = base, joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	base = b2GetJointSim(tls, world, joint)
	*(*float32)(unsafe.Add(nil, hertz)) = (*b2JointSim)(unsafe.Pointer(base)).ConstraintHertz
	*(*float32)(unsafe.Add(nil, dampingRatio)) = (*b2JointSim)(unsafe.Pointer(base)).ConstraintDampingRatio
}

func b2Joint_SetForceThreshold(tls *_Stack, jointId JointId, threshold float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var base, joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_JointSetForceThreshold
	_, _, _, _ = base, joint, world, v1
	if boolInt32(b2IsValidFloat(tls, threshold) != 0 && threshold >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+15577, uintptr(unsafe.Pointer(&__func__36))))
		if v1 = !!(b2IsValidFloat(tls, threshold) != 0 && threshold >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+15577, __ccgo_ts+12321, int32FromInt32(1450))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_JointSetForceThreshold)(unsafe.Add(nil, bp)) = b2RecArgs_JointSetForceThreshold{
			Joint:     jointId,
			Threshold: threshold,
		}
		b2RecWrite_JointSetForceThreshold(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointFullId(tls, world, jointId)
	base = b2GetJointSim(tls, world, joint)
	(*b2JointSim)(unsafe.Pointer(base)).ForceThreshold = threshold
}

func b2Joint_GetForceThreshold(tls *_Stack, jointId JointId) (r float32) {
	var base, joint, world uintptr
	_, _, _ = base, joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	base = b2GetJointSim(tls, world, joint)
	return (*b2JointSim)(unsafe.Pointer(base)).ForceThreshold
}

func b2Joint_SetTorqueThreshold(tls *_Stack, jointId JointId, threshold float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var base, joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_JointSetTorqueThreshold
	_, _, _, _ = base, joint, world, v1
	if boolInt32(b2IsValidFloat(tls, threshold) != 0 && threshold >= float32FromFloat32(0)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+15577, uintptr(unsafe.Pointer(&__func__37))))
		if v1 = !!(b2IsValidFloat(tls, threshold) != 0 && threshold >= float32FromFloat32(0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+15577, __ccgo_ts+12321, int32FromInt32(1469))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_JointSetTorqueThreshold)(unsafe.Add(nil, bp)) = b2RecArgs_JointSetTorqueThreshold{
			Joint:     jointId,
			Threshold: threshold,
		}
		b2RecWrite_JointSetTorqueThreshold(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointFullId(tls, world, jointId)
	base = b2GetJointSim(tls, world, joint)
	(*b2JointSim)(unsafe.Pointer(base)).TorqueThreshold = threshold
}

func b2Joint_GetTorqueThreshold(tls *_Stack, jointId JointId) (r float32) {
	var base, joint, world uintptr
	_, _, _ = base, joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	base = b2GetJointSim(tls, world, joint)
	return (*b2JointSim)(unsafe.Pointer(base)).TorqueThreshold
}

func b2PrepareJoint(tls *_Stack, joint uintptr, context uintptr) {
	var a1, a2, a3, hertz1, omega, v1, v2, v3, v5, v6, v7 float32
	var v10 bool
	var v8 b2Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _ = a1, a2, a3, hertz1, omega, v1, v10, v2, v3, v5, v6, v7, v8
	v1 = (*b2JointSim)(unsafe.Pointer(joint)).ConstraintHertz
	v2 = float32(float32FromFloat32(0.25) * (*b2StepContext)(unsafe.Pointer(context)).Inv_h)
	if v1 < v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	// Clamp joint hertz based on the time step to reduce jitter.
	hertz1 = v3
	v6 = hertz1
	v7 = (*b2StepContext)(unsafe.Pointer(context)).H
	if v6 == float32FromFloat32(0) {
		v8 = b2Softness{}
		goto _9
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v6)
	a1 = float32(float32FromFloat32(2)*(*b2JointSim)(unsafe.Pointer(joint)).ConstraintDampingRatio) + float32(v7*omega)
	a2 = float32(float32(v7*omega) * a1)
	a3 = float32FromFloat32(1) / (float32FromFloat32(1) + a2)
	v8 = b2Softness{
		BiasRate:     omega / a1,
		MassScale:    float32(a2 * a3),
		ImpulseScale: a3,
	}
	goto _9
_9:
	(*b2JointSim)(unsafe.Pointer(joint)).ConstraintSoftness = v8
	switch (*b2JointSim)(unsafe.Pointer(joint)).Type1 {
	case int32(b2_distanceJoint):
		b2PrepareDistanceJoint(tls, joint, context)
	case int32(b2_motorJoint):
		b2PrepareMotorJoint(tls, joint, context)
	case int32(b2_moverJoint):
		b2PrepareMoverJoint(tls, joint, context)
	case int32(b2_filterJoint):
	case int32(b2_pogoJoint):
		b2PreparePogoJoint(tls, joint, context)
	case int32(b2_prismaticJoint):
		b2PreparePrismaticJoint(tls, joint, context)
	case int32(b2_revoluteJoint):
		b2PrepareRevoluteJoint(tls, joint, context)
	case int32(b2_weldJoint):
		b2PrepareWeldJoint(tls, joint, context)
	case int32(b2_wheelJoint):
		b2PrepareWheelJoint(tls, joint, context)
	default:
		if v10 = bool(!!(int32FromInt32(false1) != 0)); !v10 {
			b2InternalAssert(tls, __ccgo_ts+7715, __ccgo_ts+12321, int32FromInt32(1530))
		}
		_ = v10 || bool(int32FromInt32(0) != 0)
	}
}

func b2WarmStartJoint(tls *_Stack, joint uintptr, context uintptr) {
	var v1 bool
	_ = v1
	switch (*b2JointSim)(unsafe.Pointer(joint)).Type1 {
	case int32(b2_distanceJoint):
		b2WarmStartDistanceJoint(tls, joint, context)
	case int32(b2_motorJoint):
		b2WarmStartMotorJoint(tls, joint, context)
	case int32(b2_moverJoint):
		b2WarmStartMoverJoint(tls, joint, context)
	case int32(b2_filterJoint):
	case int32(b2_pogoJoint):
		b2WarmStartPogoJoint(tls, joint, context)
	case int32(b2_prismaticJoint):
		b2WarmStartPrismaticJoint(tls, joint, context)
	case int32(b2_revoluteJoint):
		b2WarmStartRevoluteJoint(tls, joint, context)
	case int32(b2_weldJoint):
		b2WarmStartWeldJoint(tls, joint, context)
	case int32(b2_wheelJoint):
		b2WarmStartWheelJoint(tls, joint, context)
	default:
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+7715, __ccgo_ts+12321, int32FromInt32(1574))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
	}
}

func b2SolveJoint(tls *_Stack, joint uintptr, context uintptr, useBias uint8) {
	var v1 bool
	_ = v1
	switch (*b2JointSim)(unsafe.Pointer(joint)).Type1 {
	case int32(b2_distanceJoint):
		b2SolveDistanceJoint(tls, joint, context, useBias)
	case int32(b2_motorJoint):
		b2SolveMotorJoint(tls, joint, context)
	case int32(b2_moverJoint):
		b2SolveMoverJoint(tls, joint, context)
	case int32(b2_filterJoint):
	case int32(b2_pogoJoint):
		b2SolvePogoJoint(tls, joint, context, useBias)
	case int32(b2_prismaticJoint):
		b2SolvePrismaticJoint(tls, joint, context, useBias)
	case int32(b2_revoluteJoint):
		b2SolveRevoluteJoint(tls, joint, context, useBias)
	case int32(b2_weldJoint):
		b2SolveWeldJoint(tls, joint, context, useBias)
	case int32(b2_wheelJoint):
		b2SolveWheelJoint(tls, joint, context, useBias)
	default:
		if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+7715, __ccgo_ts+12321, int32FromInt32(1618))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
	}
}

func b2PrepareJoints_Overflow(tls *_Stack, context uintptr) {
	var graph, joint, joints uintptr
	var i, jointCount int32
	_, _, _, _, _ = graph, i, joint, jointCount, joints
	graph = (*b2StepContext)(unsafe.Pointer(context)).Graph
	joints = (*(*b2GraphColor)(unsafe.Add(nil, graph+uintptr(int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1))*64))).JointSims.Data
	jointCount = (*(*b2GraphColor)(unsafe.Add(nil, graph+uintptr(int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1))*64))).JointSims.Count
	i = 0
	for {
		if !(i < jointCount) {
			break
		}
		joint = joints + uintptr(i)*252
		b2PrepareJoint(tls, joint, context)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b2WarmStartJoints_Overflow(tls *_Stack, context uintptr) {
	var graph, joint, joints uintptr
	var i, jointCount int32
	_, _, _, _, _ = graph, i, joint, jointCount, joints
	graph = (*b2StepContext)(unsafe.Pointer(context)).Graph
	joints = (*(*b2GraphColor)(unsafe.Add(nil, graph+uintptr(int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1))*64))).JointSims.Data
	jointCount = (*(*b2GraphColor)(unsafe.Add(nil, graph+uintptr(int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1))*64))).JointSims.Count
	i = 0
	for {
		if !(i < jointCount) {
			break
		}
		joint = joints + uintptr(i)*252
		b2WarmStartJoint(tls, joint, context)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b2SolveJoints_Overflow(tls *_Stack, context uintptr, useBias uint8) {
	var graph, joint, joints uintptr
	var i, jointCount int32
	_, _, _, _, _ = graph, i, joint, jointCount, joints
	graph = (*b2StepContext)(unsafe.Pointer(context)).Graph
	joints = (*(*b2GraphColor)(unsafe.Add(nil, graph+uintptr(int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1))*64))).JointSims.Data
	jointCount = (*(*b2GraphColor)(unsafe.Add(nil, graph+uintptr(int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1))*64))).JointSims.Count
	i = 0
	for {
		if !(i < jointCount) {
			break
		}
		joint = joints + uintptr(i)*252
		b2SolveJoint(tls, joint, context, useBias)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b2PrepareJointsTask(tls *_Stack, block b2SolverBlock, context uintptr) {
	var colorEndIndex, colorIndex, colorStart, endIndex, index, v1, v2, v3, v5 int32
	var joint, joints, spans uintptr
	var v7 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _ = colorEndIndex, colorIndex, colorStart, endIndex, index, joint, joints, spans, v1, v2, v3, v5, v7
	spans = (*b2StepContext)(unsafe.Pointer(context)).JointPrepareSpans
	index = block.StartIndex
	endIndex = block.StartIndex + int32FromUint16(block.Count)
	// Find color for start index. Linear search but fast.
	colorIndex = 0
	for (*(*b2JointPrepareSpan)(unsafe.Add(nil, spans+uintptr(colorIndex+int32(1))*16))).Start <= index {
		colorIndex = colorIndex + int32(1)
	}
	// Loop over block
	for index < endIndex {
		colorStart = (*(*b2JointPrepareSpan)(unsafe.Add(nil, spans+uintptr(colorIndex)*16))).Start
		v1 = (*(*b2JointPrepareSpan)(unsafe.Add(nil, spans+uintptr(colorIndex+int32(1))*16))).Start
		v2 = endIndex
		if v1 < v2 {
			v5 = v1
		} else {
			v5 = v2
		}
		v3 = v5
		goto _4
	_4:
		colorEndIndex = v3
		joints = (*(*b2JointPrepareSpan)(unsafe.Add(nil, spans+uintptr(colorIndex)*16))).Joints
		// Loop over color
		for {
			if !(index < colorEndIndex) {
				break
			}
			if v7 = !!(0 <= index-colorStart && index-colorStart < (*(*b2JointPrepareSpan)(unsafe.Add(nil, spans+uintptr(colorIndex)*16))).Count); !v7 {
				b2InternalAssert(tls, __ccgo_ts+15626, __ccgo_ts+12321, int32FromInt32(1699))
			}
			_ = v7 || bool(int32FromInt32(0) != 0)
			joint = joints + uintptr(index-colorStart)*252
			b2PrepareJoint(tls, joint, context)
			goto _6
		_6:
			;
			index = index + 1
		}
		// Advance to next color
		colorIndex = colorIndex + int32(1)
	}
}

func b2WarmStartJointsTask(tls *_Stack, block b2SolverBlock, context uintptr) {
	var color, joint, joints uintptr
	var i int32
	_, _, _, _ = color, i, joint, joints
	color = (*b2StepContext)(unsafe.Pointer(context)).Graph + uintptr(block.ColorIndex)*64
	joints = (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Data
	i = block.StartIndex
	for {
		if !(i < block.StartIndex+int32FromUint16(block.Count)) {
			break
		}
		joint = joints + uintptr(i)*252
		b2WarmStartJoint(tls, joint, context)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b2SolveJointsTask(tls *_Stack, block b2SolverBlock, context uintptr, useBias uint8, workerIndex int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var blockIndex, blockIndex1, v4 uint32_t
	var color, joint, jointStateBitSet, joints, v3 uintptr
	var i int32
	var v1 bool
	var v5 uint8
	var _ /* force at bp+0 */ float32
	var _ /* torque at bp+4 */ float32
	_, _, _, _, _, _, _, _, _, _, _ = blockIndex, blockIndex1, color, i, joint, jointStateBitSet, joints, v1, v3, v4, v5
	color = (*b2StepContext)(unsafe.Pointer(context)).Graph + uintptr(block.ColorIndex)*64
	joints = (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Data
	if v1 = !!(0 <= block.StartIndex && block.StartIndex+int32FromUint16(block.Count) <= (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+15698, __ccgo_ts+12321, int32FromInt32(1734))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	jointStateBitSet = (*b2World)(unsafe.Pointer((*b2StepContext)(unsafe.Pointer(context)).World)).TaskContexts.Data + uintptr(workerIndex)*136 + 72
	i = block.StartIndex
	for {
		if !(i < block.StartIndex+int32FromUint16(block.Count)) {
			break
		}
		joint = joints + uintptr(i)*252
		b2SolveJoint(tls, joint, context, useBias)
		if v1 = useBias != 0 && ((*b2JointSim)(unsafe.Pointer(joint)).ForceThreshold < float32FromFloat32(3.4028234663852886e+38) || (*b2JointSim)(unsafe.Pointer(joint)).TorqueThreshold < float32FromFloat32(3.4028234663852886e+38)); v1 {
			v3 = jointStateBitSet
			v4 = uint32FromInt32((*b2JointSim)(unsafe.Pointer(joint)).JointId)
			blockIndex1 = v4 / uint32(64)
			if blockIndex1 >= (*b2BitSet)(unsafe.Pointer(v3)).BlockCount {
				v5 = boolUint8(false1 != 0)
				goto _6
			}
			v5 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v3)).Bits+uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v4%uint32FromInt32(64))) != uint64(0))
			goto _6
		_6:
		}
		if v1 && int32FromUint8(v5) == false1 {
			b2GetJointReaction(tls, joint, (*b2StepContext)(unsafe.Pointer(context)).Inv_h, bp, bp+4)
			// Check thresholds. A zero threshold means all awake joints get reported.
			if *(*float32)(unsafe.Add(nil, bp)) >= (*b2JointSim)(unsafe.Pointer(joint)).ForceThreshold || *(*float32)(unsafe.Add(nil, bp+4)) >= (*b2JointSim)(unsafe.Pointer(joint)).TorqueThreshold {
				// Flag this joint for processing.
				v3 = jointStateBitSet
				v4 = uint32FromInt32((*b2JointSim)(unsafe.Pointer(joint)).JointId)
				blockIndex = v4 / uint32(64)
				if v1 = !!(blockIndex < (*b2BitSet)(unsafe.Pointer(v3)).BlockCount); !v1 {
					b2InternalAssert(tls, __ccgo_ts+7356, __ccgo_ts+7388, int32FromInt32(29))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v3)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v4 % uint32FromInt32(64))
			}
		}
		goto _2
	_2:
		;
		i = i + 1
	}
}

func b2DrawJoint(tls *_Stack, draw uintptr, world uintptr, joint uintptr) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var bodyA, bodyB, jointSim uintptr
	var colorIndex int32
	var force, v27, v4, v8 Vec2
	var p2, p3, pA, pB, v16, v17, v19, v5, v9 Pos
	var rx, ry, scale, torque, v11, v12, v13, v15 float32
	var xfA, xfB, v3, v7 WorldTransform
	var v1, v2 bool
	var _ /* buffer at bp+0 */ [64]uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, colorIndex, force, jointSim, p2, p3, pA, pB, rx, ry, scale, torque, xfA, xfB, v1, v11, v12, v13, v15, v16, v17, v19, v2, v27, v3, v4, v5, v7, v8, v9
	if v1 = !!(0 <= (*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId && (*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+3036, __ccgo_ts+12321, int32FromInt32(1763))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId)*104
	if v2 = !!(0 <= (*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId && (*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+3124, __ccgo_ts+12321, int32FromInt32(1764))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId)*104
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_disabledSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_disabledSet) {
		return
	}
	jointSim = b2GetJointSim(tls, world, joint)
	xfA = b2GetBodyTransformQuick(tls, world, bodyA)
	xfB = b2GetBodyTransformQuick(tls, world, bodyB)
	v3 = xfA
	v4 = (*b2JointSim)(unsafe.Pointer(jointSim)).LocalFrameA.P
	rx = float32(v3.Q.C*v4.X) - float32(v3.Q.S*v4.Y)
	ry = float32(v3.Q.S*v4.X) + float32(v3.Q.C*v4.Y)
	v5 = Pos{
		X: v3.P.X + rx,
		Y: v3.P.Y + ry,
	}
	goto _6
_6:
	pA = v5
	v7 = xfB
	v8 = (*b2JointSim)(unsafe.Pointer(jointSim)).LocalFrameB.P
	rx = float32(v7.Q.C*v8.X) - float32(v7.Q.S*v8.Y)
	ry = float32(v7.Q.S*v8.X) + float32(v7.Q.C*v8.Y)
	v9 = Pos{
		X: v7.P.X + rx,
		Y: v7.P.Y + ry,
	}
	goto _10
_10:
	pB = v9
	v11 = float32FromFloat32(0.0001)
	v12 = float32((*b2DebugDraw)(unsafe.Pointer(draw)).JointScale * (*b2Joint)(unsafe.Pointer(joint)).DrawScale)
	if v11 > v12 {
		v15 = v11
	} else {
		v15 = v12
	}
	v13 = v15
	goto _14
_14:
	scale = v13
	switch (*b2Joint)(unsafe.Pointer(joint)).Type1 {
	case int32(b2_distanceJoint):
		b2DrawDistanceJoint(tls, draw, jointSim, xfA, xfB)
	case int32(b2_filterJoint):
		(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pA, float32FromFloat32(8), int32(b2_colorLightSkyBlue), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pB, float32FromFloat32(8), int32(b2_colorLightSkyBlue), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawJointExtras != 0 {
			(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, pA, pB, int32(b2_colorGold), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		}
	case int32(b2_motorJoint):
		(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pA, float32FromFloat32(8), int32(b2_colorYellowGreen), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pB, float32FromFloat32(8), int32(b2_colorPlum), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawJointExtras != 0 {
			(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, pA, pB, int32(b2_colorLightGray), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		}
	case int32(b2_moverJoint):
		(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pB, float32FromFloat32(8), int32(b2_colorPlum), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	case int32(b2_pogoJoint):
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, pA, pB, int32(b2_colorLightGray), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	case int32(b2_prismaticJoint):
		b2DrawPrismaticJoint(tls, draw, jointSim, xfA, xfB, scale)
	case int32(b2_revoluteJoint):
		b2DrawRevoluteJoint(tls, draw, jointSim, xfA, xfB, scale)
	case int32(b2_weldJoint):
		b2DrawWeldJoint(tls, draw, jointSim, xfA, xfB, scale)
	case int32(b2_wheelJoint):
		b2DrawWheelJoint(tls, draw, jointSim, xfA, xfB, scale)
	default:
		break
	}
	if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawGraphColors != 0 {
		colorIndex = (*b2Joint)(unsafe.Pointer(joint)).ColorIndex
		if colorIndex != -int32(1) {
			v5 = pA
			v9 = pB
			v11 = float32FromFloat32(0.5)
			v16 = Pos{
				X: float32((float32FromFloat32(1)-v11)*v5.X) + float32(v11*v9.X),
				Y: float32((float32FromFloat32(1)-v11)*v5.Y) + float32(v11*v9.Y),
			}
			goto _20
		_20:
			p2 = v16
			(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p2, float32FromFloat32(5), b2GetGraphColor(tls, colorIndex), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		}
	}
	if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawJointExtras != 0 {
		force = b2GetJointConstraintForce(tls, world, joint)
		torque = b2GetJointConstraintTorque(tls, world, joint)
		v5 = pA
		v9 = pB
		v11 = float32FromFloat32(0.5)
		v16 = Pos{
			X: float32((float32FromFloat32(1)-v11)*v5.X) + float32(v11*v9.X),
			Y: float32((float32FromFloat32(1)-v11)*v5.Y) + float32(v11*v9.Y),
		}
		goto _25
	_25:
		p3 = v16
		v12 = float32FromFloat32(0.001)
		v4 = force
		v8 = Vec2{
			X: float32(v12 * v4.X),
			Y: float32(v12 * v4.Y),
		}
		goto _29
	_29:
		v17 = p3
		v27 = v8
		v19 = Pos{
			X: v17.X + v27.X,
			Y: v17.Y + v27.Y,
		}
		goto _33
	_33:
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, p3, v19, int32(b2_colorAzure), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		__builtin_snprintf(tls, bp, uint64(64), __ccgo_ts+15780, vaList(bp+72, float64(force.X), float64(force.Y), float64(torque)))
		(*(*func(*_Stack, Pos, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, p3, bp, int32(b2_colorAzure), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	}
}

/** @} */

func b2MotorJoint_SetLinearVelocity(tls *_Stack, jointId JointId, velocity Vec2) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_MotorJointSetLinearVelocity
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidVec2(tls, velocity)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+16197, uintptr(unsafe.Pointer(&__func__38))))
		if v1 = !!(b2IsValidVec2(tls, velocity) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+16197, __ccgo_ts+16223, int32FromInt32(16))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_MotorJointSetLinearVelocity)(unsafe.Add(nil, bp)) = b2RecArgs_MotorJointSetLinearVelocity{
			Joint:    jointId,
			Velocity: velocity,
		}
		b2RecWrite_MotorJointSetLinearVelocity(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LinearVelocity = velocity
}

func b2MotorJoint_GetLinearVelocity(tls *_Stack, jointId JointId) (r Vec2) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	return (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LinearVelocity
}

func b2MotorJoint_SetAngularVelocity(tls *_Stack, jointId JointId, velocity float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_MotorJointSetAngularVelocity
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, velocity)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+16252, uintptr(unsafe.Pointer(&__func__39))))
		if v1 = !!(b2IsValidFloat(tls, velocity) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+16252, __ccgo_ts+16223, int32FromInt32(32))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_MotorJointSetAngularVelocity)(unsafe.Add(nil, bp)) = b2RecArgs_MotorJointSetAngularVelocity{
			Joint:    jointId,
			Velocity: velocity,
		}
		b2RecWrite_MotorJointSetAngularVelocity(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).AngularVelocity = velocity
}

func b2MotorJoint_GetAngularVelocity(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	return (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).AngularVelocity
}

func b2MotorJoint_SetMaxVelocityTorque(tls *_Stack, jointId JointId, maxTorque float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_MotorJointSetMaxVelocityTorque
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, maxTorque)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+16279, uintptr(unsafe.Pointer(&__func__40))))
		if v1 = !!(b2IsValidFloat(tls, maxTorque) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+16279, __ccgo_ts+16223, int32FromInt32(48))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_MotorJointSetMaxVelocityTorque)(unsafe.Add(nil, bp)) = b2RecArgs_MotorJointSetMaxVelocityTorque{
			Joint:     jointId,
			MaxTorque: maxTorque,
		}
		b2RecWrite_MotorJointSetMaxVelocityTorque(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxVelocityTorque = maxTorque
}

func b2MotorJoint_GetMaxVelocityTorque(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	return (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxVelocityTorque
}

func b2MotorJoint_SetMaxVelocityForce(tls *_Stack, jointId JointId, maxForce float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_MotorJointSetMaxVelocityForce
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, maxForce)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+16307, uintptr(unsafe.Pointer(&__func__41))))
		if v1 = !!(b2IsValidFloat(tls, maxForce) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+16307, __ccgo_ts+16223, int32FromInt32(64))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_MotorJointSetMaxVelocityForce)(unsafe.Add(nil, bp)) = b2RecArgs_MotorJointSetMaxVelocityForce{
			Joint:    jointId,
			MaxForce: maxForce,
		}
		b2RecWrite_MotorJointSetMaxVelocityForce(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxVelocityForce = maxForce
}

func b2MotorJoint_GetMaxVelocityForce(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	return (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxVelocityForce
}

func b2MotorJoint_SetLinearHertz(tls *_Stack, jointId JointId, hertz float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_MotorJointSetLinearHertz
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, hertz)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8368, uintptr(unsafe.Pointer(&__func__42))))
		if v1 = !!(b2IsValidFloat(tls, hertz) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8368, __ccgo_ts+16223, int32FromInt32(80))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_MotorJointSetLinearHertz)(unsafe.Add(nil, bp)) = b2RecArgs_MotorJointSetLinearHertz{
			Joint: jointId,
			Hertz: hertz,
		}
		b2RecWrite_MotorJointSetLinearHertz(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LinearHertz = hertz
}

func b2MotorJoint_GetLinearHertz(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	return (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LinearHertz
}

func b2MotorJoint_SetLinearDampingRatio(tls *_Stack, jointId JointId, damping float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_MotorJointSetLinearDampingRatio
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, damping)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+16334, uintptr(unsafe.Pointer(&__func__43))))
		if v1 = !!(b2IsValidFloat(tls, damping) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+16334, __ccgo_ts+16223, int32FromInt32(96))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_MotorJointSetLinearDampingRatio)(unsafe.Add(nil, bp)) = b2RecArgs_MotorJointSetLinearDampingRatio{
			Joint:   jointId,
			Damping: damping,
		}
		b2RecWrite_MotorJointSetLinearDampingRatio(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LinearDampingRatio = damping
}

func b2MotorJoint_GetLinearDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	return (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LinearDampingRatio
}

func b2MotorJoint_SetAngularHertz(tls *_Stack, jointId JointId, hertz float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_MotorJointSetAngularHertz
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, hertz)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8368, uintptr(unsafe.Pointer(&__func__44))))
		if v1 = !!(b2IsValidFloat(tls, hertz) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8368, __ccgo_ts+16223, int32FromInt32(112))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_MotorJointSetAngularHertz)(unsafe.Add(nil, bp)) = b2RecArgs_MotorJointSetAngularHertz{
			Joint: jointId,
			Hertz: hertz,
		}
		b2RecWrite_MotorJointSetAngularHertz(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).AngularHertz = hertz
}

func b2MotorJoint_GetAngularHertz(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	return (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).AngularHertz
}

func b2MotorJoint_SetAngularDampingRatio(tls *_Stack, jointId JointId, damping float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_MotorJointSetAngularDampingRatio
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, damping)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+16334, uintptr(unsafe.Pointer(&__func__45))))
		if v1 = !!(b2IsValidFloat(tls, damping) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+16334, __ccgo_ts+16223, int32FromInt32(128))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_MotorJointSetAngularDampingRatio)(unsafe.Add(nil, bp)) = b2RecArgs_MotorJointSetAngularDampingRatio{
			Joint:   jointId,
			Damping: damping,
		}
		b2RecWrite_MotorJointSetAngularDampingRatio(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).AngularDampingRatio = damping
}

func b2MotorJoint_GetAngularDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	return (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).AngularDampingRatio
}

func b2MotorJoint_SetMaxSpringForce(tls *_Stack, jointId JointId, maxForce float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var v2, v3, v4, v6 float32
	var _ /* _a at bp+0 */ b2RecArgs_MotorJointSetMaxSpringForce
	_, _, _, _, _, _, _ = joint, world, v1, v2, v3, v4, v6
	if int32FromUint8(b2IsValidFloat(tls, maxForce)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+16307, uintptr(unsafe.Pointer(&__func__110))))
		if v1 = !!(b2IsValidFloat(tls, maxForce) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+16307, __ccgo_ts+16223, int32FromInt32(144))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_MotorJointSetMaxSpringForce)(unsafe.Add(nil, bp)) = b2RecArgs_MotorJointSetMaxSpringForce{
			Joint:    jointId,
			MaxForce: maxForce,
		}
		b2RecWrite_MotorJointSetMaxSpringForce(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	v2 = float32FromFloat32(0)
	v3 = maxForce
	if v2 > v3 {
		v6 = v2
	} else {
		v6 = v3
	}
	v4 = v6
	goto _5
_5:
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxSpringForce = v4
}

func b2MotorJoint_GetMaxSpringForce(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	return (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxSpringForce
}

func b2MotorJoint_SetMaxSpringTorque(tls *_Stack, jointId JointId, maxTorque float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var v2, v3, v4, v6 float32
	var _ /* _a at bp+0 */ b2RecArgs_MotorJointSetMaxSpringTorque
	_, _, _, _, _, _, _ = joint, world, v1, v2, v3, v4, v6
	if int32FromUint8(b2IsValidFloat(tls, maxTorque)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+16279, uintptr(unsafe.Pointer(&__func__111))))
		if v1 = !!(b2IsValidFloat(tls, maxTorque) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+16279, __ccgo_ts+16223, int32FromInt32(160))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_MotorJointSetMaxSpringTorque)(unsafe.Add(nil, bp)) = b2RecArgs_MotorJointSetMaxSpringTorque{
			Joint:     jointId,
			MaxTorque: maxTorque,
		}
		b2RecWrite_MotorJointSetMaxSpringTorque(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	v2 = float32FromFloat32(0)
	v3 = maxTorque
	if v2 > v3 {
		v6 = v2
	} else {
		v6 = v3
	}
	v4 = v6
	goto _5
_5:
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxSpringTorque = v4
}

func b2MotorJoint_GetMaxSpringTorque(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	return (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxSpringTorque
}

func b2GetMotorJointForce(tls *_Stack, world uintptr, base uintptr) (r Vec2) {
	var force, v1, v2, v3, v6, v7 Vec2
	var v5 float32
	_, _, _, _, _, _, _ = force, v1, v2, v3, v5, v6, v7
	v1 = (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(base), 92))).LinearVelocityImpulse
	v2 = (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(base), 92))).LinearSpringImpulse
	v3 = Vec2{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
	goto _4
_4:
	v5 = (*b2World)(unsafe.Pointer(world)).Inv_h
	v6 = v3
	v7 = Vec2{
		X: float32(v5 * v6.X),
		Y: float32(v5 * v6.Y),
	}
	goto _8
_8:
	force = v7
	return force
}

func b2GetMotorJointTorque(tls *_Stack, world uintptr, base uintptr) (r float32) {
	return float32((*b2World)(unsafe.Pointer(world)).Inv_h * ((*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(base), 92))).AngularVelocityImpulse + (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(base), 92))).AngularSpringImpulse))
}

// Point-to-point constraint
// C = p2 - p1
// Cdot = v2 - v1
//      = v2 + cross(w2, r2) - v1 - cross(w1, r1)
// J = [-I -r1_skew I r2_skew ]
// Identity used:
// w k % (rx i + ry j) = w * (-ry i + rx j)

// Angle constraint
// C = angle2 - angle1 - referenceAngle
// Cdot = w2 - w1
// J = [0 0 -1 0 0 1]
// K = invI1 + invI2

func b2PrepareMotorJoint(tls *_Stack, base uintptr, context uintptr) {
	var B, kl, v47, v48 Mat22
	var a11, a2, a21, a3, b2, c, d, det, iA, iB, ka, mA, mB, omega, v39, v40 float32
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world uintptr
	var idA, idB, localIndexA, localIndexB, v9 int32
	var qr, v11, v12, v13 Rot
	var rA, rB, v15, v16, v17, v20, v21 Vec2
	var v1, v2, v3, v4 bool
	var v35, v36 Pos
	var v41 b2Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = B, a11, a2, a21, a3, b2, bodyA, bodyB, bodySimA, bodySimB, c, d, det, iA, iB, idA, idB, joint, ka, kl, localIndexA, localIndexB, mA, mB, omega, qr, rA, rB, setA, setB, world, v1, v11, v12, v13, v15, v16, v17, v2, v20, v21, v3, v35, v36, v39, v4, v40, v41, v47, v48, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_motorJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+16360, __ccgo_ts+16223, int32FromInt32(201))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// chase body id to the solver set where the body lives
	idA = (*b2JointSim)(unsafe.Pointer(base)).BodyIdA
	idB = (*b2JointSim)(unsafe.Pointer(base)).BodyIdB
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	if v1 = !!(0 <= idA && idA < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8507, __ccgo_ts+16223, int32FromInt32(209))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idA)*104
	if v2 = !!(0 <= idB && idB < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+8557, __ccgo_ts+16223, int32FromInt32(210))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idB)*104
	if v3 = !!((*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet)); !v3 {
		b2InternalAssert(tls, __ccgo_ts+8607, __ccgo_ts+16223, int32FromInt32(212))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(bodyA)).SetIndex && (*b2Body)(unsafe.Pointer(bodyA)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8672, __ccgo_ts+16223, int32FromInt32(214))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	setA = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(bodyA)).SetIndex)*88
	if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(bodyB)).SetIndex && (*b2Body)(unsafe.Pointer(bodyB)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+8750, __ccgo_ts+16223, int32FromInt32(215))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	setB = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(bodyB)).SetIndex)*88
	localIndexA = (*b2Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b2Body)(unsafe.Pointer(bodyB)).LocalIndex
	if v3 = !!(0 <= localIndexA && localIndexA < (*b2SolverSet)(unsafe.Pointer(setA)).BodySims.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+8828, __ccgo_ts+16223, int32FromInt32(220))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	bodySimA = (*b2SolverSet)(unsafe.Pointer(setA)).BodySims.Data + uintptr(localIndexA)*96
	if v4 = !!(0 <= localIndexB && localIndexB < (*b2SolverSet)(unsafe.Pointer(setB)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+8895, __ccgo_ts+16223, int32FromInt32(221))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySimB = (*b2SolverSet)(unsafe.Pointer(setB)).BodySims.Data + uintptr(localIndexB)*96
	mA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvMass
	iA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvInertia
	mB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvMass
	iB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvInertia
	(*b2JointSim)(unsafe.Pointer(base)).InvMassA = mA
	(*b2JointSim)(unsafe.Pointer(base)).InvMassB = mB
	(*b2JointSim)(unsafe.Pointer(base)).InvIA = iA
	(*b2JointSim)(unsafe.Pointer(base)).InvIB = iB
	joint = base + 92
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) {
		v9 = localIndexA
	} else {
		v9 = -int32(1)
	}
	(*b2MotorJoint)(unsafe.Pointer(joint)).IndexA = v9
	if (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) {
		v9 = localIndexB
	} else {
		v9 = -int32(1)
	}
	(*b2MotorJoint)(unsafe.Pointer(joint)).IndexB = v9
	// Compute joint anchor frames with world space rotation, relative to center of mass
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v12 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	qr.S = float32(v11.S*v12.C) + float32(v11.C*v12.S)
	qr.C = float32(v11.C*v12.C) - float32(v11.S*v12.S)
	v13 = qr
	goto _14
_14:
	(*b2MotorJoint)(unsafe.Pointer(joint)).FrameA.Q = v13
	v15 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA.P
	v16 = (*b2BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v17 = Vec2{
		X: v15.X - v16.X,
		Y: v15.Y - v16.Y,
	}
	goto _18
_18:
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v20 = v17
	v21 = Vec2{
		X: float32(v11.C*v20.X) - float32(v11.S*v20.Y),
		Y: float32(v11.S*v20.X) + float32(v11.C*v20.Y),
	}
	goto _22
_22:
	(*b2MotorJoint)(unsafe.Pointer(joint)).FrameA.P = v21
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v12 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameB.Q
	qr.S = float32(v11.S*v12.C) + float32(v11.C*v12.S)
	qr.C = float32(v11.C*v12.C) - float32(v11.S*v12.S)
	v13 = qr
	goto _26
_26:
	(*b2MotorJoint)(unsafe.Pointer(joint)).FrameB.Q = v13
	v15 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v16 = (*b2BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v17 = Vec2{
		X: v15.X - v16.X,
		Y: v15.Y - v16.Y,
	}
	goto _30
_30:
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v20 = v17
	v21 = Vec2{
		X: float32(v11.C*v20.X) - float32(v11.S*v20.Y),
		Y: float32(v11.S*v20.X) + float32(v11.C*v20.Y),
	}
	goto _34
_34:
	(*b2MotorJoint)(unsafe.Pointer(joint)).FrameB.P = v21
	// Compute the initial center delta. Incremental position updates are relative to this.
	v35 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Center
	v36 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Center
	v15 = Vec2{
		X: v35.X - v36.X,
		Y: v35.Y - v36.Y,
	}
	goto _38
_38:
	(*b2MotorJoint)(unsafe.Pointer(joint)).DeltaCenter = v15
	rA = (*b2MotorJoint)(unsafe.Pointer(joint)).FrameA.P
	rB = (*b2MotorJoint)(unsafe.Pointer(joint)).FrameB.P
	v39 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearHertz
	v40 = (*b2StepContext)(unsafe.Pointer(context)).H
	if v39 == float32FromFloat32(0) {
		v41 = b2Softness{}
		goto _42
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(_B2_PI)) * v39)
	a11 = float32(float32FromFloat32(2)*(*b2MotorJoint)(unsafe.Pointer(joint)).LinearDampingRatio) + float32(v40*omega)
	a21 = float32(float32(v40*omega) * a11)
	a3 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v41 = b2Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a3),
		ImpulseScale: a3,
	}
	goto _42
_42:
	(*b2MotorJoint)(unsafe.Pointer(joint)).LinearSpring = v41
	v39 = (*b2MotorJoint)(unsafe.Pointer(joint)).AngularHertz
	v40 = (*b2StepContext)(unsafe.Pointer(context)).H
	if v39 == float32FromFloat32(0) {
		v41 = b2Softness{}
		goto _46
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(_B2_PI)) * v39)
	a11 = float32(float32FromFloat32(2)*(*b2MotorJoint)(unsafe.Pointer(joint)).AngularDampingRatio) + float32(v40*omega)
	a21 = float32(float32(v40*omega) * a11)
	a3 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v41 = b2Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a3),
		ImpulseScale: a3,
	}
	goto _46
_46:
	(*b2MotorJoint)(unsafe.Pointer(joint)).AngularSpring = v41
	kl.Cx.X = mA + mB + float32(float32(rA.Y*rA.Y)*iA) + float32(float32(rB.Y*rB.Y)*iB)
	kl.Cx.Y = float32(float32(-rA.Y*rA.X)*iA) - float32(float32(rB.Y*rB.X)*iB)
	kl.Cy.X = kl.Cx.Y
	kl.Cy.Y = mA + mB + float32(float32(rA.X*rA.X)*iA) + float32(float32(rB.X*rB.X)*iB)
	v47 = kl
	a2 = v47.Cx.X
	b2 = v47.Cy.X
	c = v47.Cx.Y
	d = v47.Cy.Y
	det = float32(a2*d) - float32(b2*c)
	if det != float32FromFloat32(0) {
		det = float32FromFloat32(1) / det
	}
	B = Mat22{
		Cx: Vec2{
			X: float32(det * d),
			Y: float32(-det * c),
		},
		Cy: Vec2{
			X: float32(-det * b2),
			Y: float32(det * a2),
		},
	}
	v48 = B
	goto _49
_49:
	(*b2MotorJoint)(unsafe.Pointer(joint)).LinearMass = v48
	ka = iA + iB
	if ka > float32FromFloat32(0) {
		v39 = float32FromFloat32(1) / ka
	} else {
		v39 = float32FromFloat32(0)
	}
	(*b2MotorJoint)(unsafe.Pointer(joint)).AngularMass = v39
	if int32FromUint8((*b2StepContext)(unsafe.Pointer(context)).EnableWarmStarting) == false1 {
		(*b2MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse = b2Vec2_zero
		(*b2MotorJoint)(unsafe.Pointer(joint)).AngularVelocityImpulse = float32FromFloat32(0)
		(*b2MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse = b2Vec2_zero
		(*b2MotorJoint)(unsafe.Pointer(joint)).AngularSpringImpulse = float32FromFloat32(0)
	}
}

func b2WarmStartMotorJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var angularImpulse, iA, iB, mA, mB, v17 float32
	var joint, stateA, stateB, v2, v3 uintptr
	var linearImpulse, rA, rB, v10, v12, v13, v14, v16, v18, v19, v5, v6, v9 Vec2
	var v1 bool
	var v4, v8 Rot
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angularImpulse, iA, iB, joint, linearImpulse, mA, mB, rA, rB, stateA, stateB, v1, v10, v12, v13, v14, v16, v17, v18, v19, v2, v3, v4, v5, v6, v8, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_motorJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+16360, __ccgo_ts+16223, int32FromInt32(273))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	joint = base + 92
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Add(nil, bp)) = b2_identityBodyState9
	if (*b2MotorJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2MotorJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v2
	if (*b2MotorJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2MotorJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v3
	v4 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b2MotorJoint)(unsafe.Pointer(joint)).FrameA.P
	v6 = Vec2{
		X: float32(v4.C*v5.X) - float32(v4.S*v5.Y),
		Y: float32(v4.S*v5.X) + float32(v4.C*v5.Y),
	}
	goto _7
_7:
	rA = v6
	v8 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v9 = (*b2MotorJoint)(unsafe.Pointer(joint)).FrameB.P
	v10 = Vec2{
		X: float32(v8.C*v9.X) - float32(v8.S*v9.Y),
		Y: float32(v8.S*v9.X) + float32(v8.C*v9.Y),
	}
	goto _11
_11:
	rB = v10
	v12 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse
	v13 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse
	v14 = Vec2{
		X: v12.X + v13.X,
		Y: v12.Y + v13.Y,
	}
	goto _15
_15:
	linearImpulse = v14
	angularImpulse = (*b2MotorJoint)(unsafe.Pointer(joint)).AngularVelocityImpulse + (*b2MotorJoint)(unsafe.Pointer(joint)).AngularSpringImpulse
	if (*b2BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b2_dynamicFlag) != 0 {
		v16 = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
		v17 = mA
		v18 = linearImpulse
		v19 = Vec2{
			X: v16.X - float32(v17*v18.X),
			Y: v16.Y - float32(v17*v18.Y),
		}
		goto _20
	_20:
		(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = v19
		v5 = rA
		v6 = linearImpulse
		v17 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		goto _24
	_24:
		*(*float32)(unsafe.Add(nil, stateA+8)) -= float32(iA * (v17 + angularImpulse))
	}
	if (*b2BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b2_dynamicFlag) != 0 {
		v5 = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
		v17 = mB
		v6 = linearImpulse
		v9 = Vec2{
			X: v5.X + float32(v17*v6.X),
			Y: v5.Y + float32(v17*v6.Y),
		}
		goto _29
	_29:
		(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = v9
		v5 = rB
		v6 = linearImpulse
		v17 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		goto _33
	_33:
		*(*float32)(unsafe.Add(nil, stateB+8)) += float32(iB * (v17 + angularImpulse))
	}
}

func b2SolveMotorJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var B, kl, v85, v86 Mat22
	var a8, b6, bias, c, c1, cdot, cdot1, d, det, iA, iB, impulse, impulse1, impulseScale, impulseScale1, lengthSquared, mA, mB, massScale, massScale1, maxImpulse, maxImpulse1, maxImpulse2, maxImpulse3, oldImpulse, oldImpulse1, s4, wA, wB, v17, v19, v20, v21, v22, v24, v25 float32
	var b7, b8, bias1, c2, cdot2, cdot3, dcA, dcB, impulse2, impulse3, oldImpulse2, oldImpulse3, rA, rB, u, u1, vA, vB, v34, v35, v38, v39, v41, v42, v43, v45, v46, v47, v49, v50, v51, v53, v54, v55, v58, v59, v62, v63, v65, v66, v67, v70, v71, v73, v74, v75, v77, v78, v79, v81, v82, v83 Vec2
	var joint, stateA, stateB, v2, v3 uintptr
	var qA, qB, qr, r1, relQ, v10, v12, v13, v14, v16, v4, v5, v6, v8, v9 Rot
	var v1 bool
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = B, a8, b6, b7, b8, bias, bias1, c, c1, c2, cdot, cdot1, cdot2, cdot3, d, dcA, dcB, det, iA, iB, impulse, impulse1, impulse2, impulse3, impulseScale, impulseScale1, joint, kl, lengthSquared, mA, mB, massScale, massScale1, maxImpulse, maxImpulse1, maxImpulse2, maxImpulse3, oldImpulse, oldImpulse1, oldImpulse2, oldImpulse3, qA, qB, qr, r1, rA, rB, relQ, s4, stateA, stateB, u, u1, vA, vB, wA, wB, v1, v10, v12, v13, v14, v16, v17, v19, v2, v20, v21, v22, v24, v25, v3, v34, v35, v38, v39, v4, v41, v42, v43, v45, v46, v47, v49, v5, v50, v51, v53, v54, v55, v58, v59, v6, v62, v63, v65, v66, v67, v70, v71, v73, v74, v75, v77, v78, v79, v8, v81, v82, v83, v85, v86, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_motorJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+16360, __ccgo_ts+16223, int32FromInt32(309))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Add(nil, bp)) = b2_identityBodyState9
	joint = base + 92
	if (*b2MotorJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2MotorJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v2
	if (*b2MotorJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2MotorJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v3
	vA = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	// angular spring
	if (*b2MotorJoint)(unsafe.Pointer(joint)).MaxSpringTorque > float32FromFloat32(0) && (*b2MotorJoint)(unsafe.Pointer(joint)).AngularHertz > float32FromFloat32(0) {
		v4 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
		v5 = (*b2MotorJoint)(unsafe.Pointer(joint)).FrameA.Q
		qr.S = float32(v4.S*v5.C) + float32(v4.C*v5.S)
		qr.C = float32(v4.C*v5.C) - float32(v4.S*v5.S)
		v6 = qr
		goto _7
	_7:
		qA = v6
		v8 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
		v9 = (*b2MotorJoint)(unsafe.Pointer(joint)).FrameB.Q
		qr.S = float32(v8.S*v9.C) + float32(v8.C*v9.S)
		qr.C = float32(v8.C*v9.C) - float32(v8.S*v9.S)
		v10 = qr
		goto _11
	_11:
		qB = v10
		v12 = qA
		v13 = qB
		r1.S = float32(v12.C*v13.S) - float32(v12.S*v13.C)
		r1.C = float32(v12.C*v13.C) + float32(v12.S*v13.S)
		v14 = r1
		goto _15
	_15:
		relQ = v14
		v16 = relQ
		v17 = b2Atan2(tls, v16.S, v16.C)
		goto _18
	_18:
		c1 = v17
		bias = float32((*b2MotorJoint)(unsafe.Pointer(joint)).AngularSpring.BiasRate * c1)
		massScale = (*b2MotorJoint)(unsafe.Pointer(joint)).AngularSpring.MassScale
		impulseScale = (*b2MotorJoint)(unsafe.Pointer(joint)).AngularSpring.ImpulseScale
		cdot = wB - wA
		maxImpulse = float32((*b2StepContext)(unsafe.Pointer(context)).H * (*b2MotorJoint)(unsafe.Pointer(joint)).MaxSpringTorque)
		oldImpulse = (*b2MotorJoint)(unsafe.Pointer(joint)).AngularSpringImpulse
		impulse = float32(float32(-massScale*(*b2MotorJoint)(unsafe.Pointer(joint)).AngularMass)*(cdot+bias)) - float32(impulseScale*oldImpulse)
		v19 = oldImpulse + impulse
		v20 = -maxImpulse
		v21 = maxImpulse
		if v19 < v20 {
			v24 = v20
		} else {
			if v19 > v21 {
				v25 = v21
			} else {
				v25 = v19
			}
			v24 = v25
		}
		v22 = v24
		goto _23
	_23:
		(*b2MotorJoint)(unsafe.Pointer(joint)).AngularSpringImpulse = v22
		impulse = (*b2MotorJoint)(unsafe.Pointer(joint)).AngularSpringImpulse - oldImpulse
		wA = wA - float32(iA*impulse)
		wB = wB + float32(iB*impulse)
	}
	// angular velocity
	if (*b2MotorJoint)(unsafe.Pointer(joint)).MaxVelocityTorque > float32FromFloat32(0) {
		cdot1 = wB - wA - (*b2MotorJoint)(unsafe.Pointer(joint)).AngularVelocity
		impulse1 = float32(-(*b2MotorJoint)(unsafe.Pointer(joint)).AngularMass * cdot1)
		maxImpulse1 = float32((*b2StepContext)(unsafe.Pointer(context)).H * (*b2MotorJoint)(unsafe.Pointer(joint)).MaxVelocityTorque)
		oldImpulse1 = (*b2MotorJoint)(unsafe.Pointer(joint)).AngularVelocityImpulse
		v17 = oldImpulse1 + impulse1
		v19 = -maxImpulse1
		v20 = maxImpulse1
		if v17 < v19 {
			v22 = v19
		} else {
			if v17 > v20 {
				v24 = v20
			} else {
				v24 = v17
			}
			v22 = v24
		}
		v21 = v22
		goto _30
	_30:
		(*b2MotorJoint)(unsafe.Pointer(joint)).AngularVelocityImpulse = v21
		impulse1 = (*b2MotorJoint)(unsafe.Pointer(joint)).AngularVelocityImpulse - oldImpulse1
		wA = wA - float32(iA*impulse1)
		wB = wB + float32(iB*impulse1)
	}
	v4 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v34 = (*b2MotorJoint)(unsafe.Pointer(joint)).FrameA.P
	v35 = Vec2{
		X: float32(v4.C*v34.X) - float32(v4.S*v34.Y),
		Y: float32(v4.S*v34.X) + float32(v4.C*v34.Y),
	}
	goto _36
_36:
	rA = v35
	v5 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v38 = (*b2MotorJoint)(unsafe.Pointer(joint)).FrameB.P
	v39 = Vec2{
		X: float32(v5.C*v38.X) - float32(v5.S*v38.Y),
		Y: float32(v5.S*v38.X) + float32(v5.C*v38.Y),
	}
	goto _40
_40:
	rB = v39
	// linear spring
	if (*b2MotorJoint)(unsafe.Pointer(joint)).MaxSpringForce > float32FromFloat32(0) && (*b2MotorJoint)(unsafe.Pointer(joint)).LinearHertz > float32FromFloat32(0) {
		dcA = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaPosition
		dcB = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
		v41 = dcB
		v42 = dcA
		v43 = Vec2{
			X: v41.X - v42.X,
			Y: v41.Y - v42.Y,
		}
		goto _44
	_44:
		v45 = rB
		v46 = rA
		v47 = Vec2{
			X: v45.X - v46.X,
			Y: v45.Y - v46.Y,
		}
		goto _48
	_48:
		v49 = v43
		v50 = v47
		v51 = Vec2{
			X: v49.X + v50.X,
			Y: v49.Y + v50.Y,
		}
		goto _52
	_52:
		v53 = v51
		v54 = (*b2MotorJoint)(unsafe.Pointer(joint)).DeltaCenter
		v55 = Vec2{
			X: v53.X + v54.X,
			Y: v53.Y + v54.Y,
		}
		goto _56
	_56:
		c2 = v55
		v17 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearSpring.BiasRate
		v58 = c2
		v59 = Vec2{
			X: float32(v17 * v58.X),
			Y: float32(v17 * v58.Y),
		}
		goto _60
	_60:
		bias1 = v59
		massScale1 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearSpring.MassScale
		impulseScale1 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearSpring.ImpulseScale
		v19 = wB
		v62 = rB
		v63 = Vec2{
			X: float32(-v19 * v62.Y),
			Y: float32(v19 * v62.X),
		}
		goto _64
	_64:
		v65 = vB
		v66 = v63
		v67 = Vec2{
			X: v65.X + v66.X,
			Y: v65.Y + v66.Y,
		}
		goto _68
	_68:
		v20 = wA
		v70 = rA
		v71 = Vec2{
			X: float32(-v20 * v70.Y),
			Y: float32(v20 * v70.X),
		}
		goto _72
	_72:
		v73 = vA
		v74 = v71
		v75 = Vec2{
			X: v73.X + v74.X,
			Y: v73.Y + v74.Y,
		}
		goto _76
	_76:
		v77 = v67
		v78 = v75
		v79 = Vec2{
			X: v77.X - v78.X,
			Y: v77.Y - v78.Y,
		}
		goto _80
	_80:
		cdot2 = v79
		v81 = cdot2
		v82 = bias1
		v83 = Vec2{
			X: v81.X + v82.X,
			Y: v81.Y + v82.Y,
		}
		goto _84
	_84:
		cdot2 = v83
		kl.Cx.X = mA + mB + float32(float32(rA.Y*rA.Y)*iA) + float32(float32(rB.Y*rB.Y)*iB)
		kl.Cx.Y = float32(float32(-rA.Y*rA.X)*iA) - float32(float32(rB.Y*rB.X)*iB)
		kl.Cy.X = kl.Cx.Y
		kl.Cy.Y = mA + mB + float32(float32(rA.X*rA.X)*iA) + float32(float32(rB.X*rB.X)*iB)
		v85 = kl
		a8 = v85.Cx.X
		b6 = v85.Cy.X
		c = v85.Cx.Y
		d = v85.Cy.Y
		det = float32(a8*d) - float32(b6*c)
		if det != float32FromFloat32(0) {
			det = float32FromFloat32(1) / det
		}
		B = Mat22{
			Cx: Vec2{
				X: float32(det * d),
				Y: float32(-det * c),
			},
			Cy: Vec2{
				X: float32(-det * b6),
				Y: float32(det * a8),
			},
		}
		v86 = B
		goto _87
	_87:
		(*b2MotorJoint)(unsafe.Pointer(joint)).LinearMass = v86
		v85 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearMass
		v34 = cdot2
		u1 = Vec2{
			X: float32(v85.Cx.X*v34.X) + float32(v85.Cy.X*v34.Y),
			Y: float32(v85.Cx.Y*v34.X) + float32(v85.Cy.Y*v34.Y),
		}
		v35 = u1
		goto _91
	_91:
		b7 = v35
		oldImpulse2 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse
		impulse2 = Vec2{
			X: float32(-massScale1*b7.X) - float32(impulseScale1*oldImpulse2.X),
			Y: float32(-massScale1*b7.Y) - float32(impulseScale1*oldImpulse2.Y),
		}
		maxImpulse2 = float32((*b2StepContext)(unsafe.Pointer(context)).H * (*b2MotorJoint)(unsafe.Pointer(joint)).MaxSpringForce)
		v38 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse
		v39 = impulse2
		v41 = Vec2{
			X: v38.X + v39.X,
			Y: v38.Y + v39.Y,
		}
		goto _95
	_95:
		(*b2MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse = v41
		v34 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse
		v17 = float32(v34.X*v34.X) + float32(v34.Y*v34.Y)
		goto _98
	_98:
		if v17 > float32(maxImpulse2*maxImpulse2) {
			v35 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse
			lengthSquared = float32(v35.X*v35.X) + float32(v35.Y*v35.Y)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s4 = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec2{
					X: float32(s4 * v35.X),
					Y: float32(s4 * v35.Y),
				}
				v38 = u
				goto _101
			}
			v38 = Vec2{}
			goto _101
		_101:
			(*b2MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse = v38
			(*b2MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse.X = float32((*b2MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse.X * maxImpulse2)
			(*b2MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse.Y = float32((*b2MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse.Y * maxImpulse2)
		}
		v34 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearSpringImpulse
		v35 = oldImpulse2
		v38 = Vec2{
			X: v34.X - v35.X,
			Y: v34.Y - v35.Y,
		}
		goto _105
	_105:
		impulse2 = v38
		v34 = vA
		v17 = mA
		v35 = impulse2
		v38 = Vec2{
			X: v34.X - float32(v17*v35.X),
			Y: v34.Y - float32(v17*v35.Y),
		}
		goto _110
	_110:
		vA = v38
		v34 = rA
		v35 = impulse2
		v17 = float32(v34.X*v35.Y) - float32(v34.Y*v35.X)
		goto _114
	_114:
		wA = wA - float32(iA*v17)
		v34 = vB
		v17 = mB
		v35 = impulse2
		v38 = Vec2{
			X: v34.X + float32(v17*v35.X),
			Y: v34.Y + float32(v17*v35.Y),
		}
		goto _119
	_119:
		vB = v38
		v34 = rB
		v35 = impulse2
		v17 = float32(v34.X*v35.Y) - float32(v34.Y*v35.X)
		goto _123
	_123:
		wB = wB + float32(iB*v17)
	}
	// linear velocity
	if (*b2MotorJoint)(unsafe.Pointer(joint)).MaxVelocityForce > float32FromFloat32(0) {
		v17 = wB
		v34 = rB
		v35 = Vec2{
			X: float32(-v17 * v34.Y),
			Y: float32(v17 * v34.X),
		}
		goto _127
	_127:
		v38 = vB
		v39 = v35
		v41 = Vec2{
			X: v38.X + v39.X,
			Y: v38.Y + v39.Y,
		}
		goto _131
	_131:
		v19 = wA
		v42 = rA
		v43 = Vec2{
			X: float32(-v19 * v42.Y),
			Y: float32(v19 * v42.X),
		}
		goto _135
	_135:
		v45 = vA
		v46 = v43
		v47 = Vec2{
			X: v45.X + v46.X,
			Y: v45.Y + v46.Y,
		}
		goto _139
	_139:
		v49 = v41
		v50 = v47
		v51 = Vec2{
			X: v49.X - v50.X,
			Y: v49.Y - v50.Y,
		}
		goto _143
	_143:
		cdot3 = v51
		v53 = cdot3
		v54 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearVelocity
		v55 = Vec2{
			X: v53.X - v54.X,
			Y: v53.Y - v54.Y,
		}
		goto _147
	_147:
		cdot3 = v55
		v85 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearMass
		v34 = cdot3
		u1 = Vec2{
			X: float32(v85.Cx.X*v34.X) + float32(v85.Cy.X*v34.Y),
			Y: float32(v85.Cx.Y*v34.X) + float32(v85.Cy.Y*v34.Y),
		}
		v35 = u1
		goto _151
	_151:
		b8 = v35
		impulse3 = Vec2{
			X: -b8.X,
			Y: -b8.Y,
		}
		oldImpulse3 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse
		maxImpulse3 = float32((*b2StepContext)(unsafe.Pointer(context)).H * (*b2MotorJoint)(unsafe.Pointer(joint)).MaxVelocityForce)
		v38 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse
		v39 = impulse3
		v41 = Vec2{
			X: v38.X + v39.X,
			Y: v38.Y + v39.Y,
		}
		goto _155
	_155:
		(*b2MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse = v41
		v34 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse
		v17 = float32(v34.X*v34.X) + float32(v34.Y*v34.Y)
		goto _158
	_158:
		if v17 > float32(maxImpulse3*maxImpulse3) {
			v35 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse
			lengthSquared = float32(v35.X*v35.X) + float32(v35.Y*v35.Y)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s4 = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec2{
					X: float32(s4 * v35.X),
					Y: float32(s4 * v35.Y),
				}
				v38 = u
				goto _161
			}
			v38 = Vec2{}
			goto _161
		_161:
			(*b2MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse = v38
			(*b2MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse.X = float32((*b2MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse.X * maxImpulse3)
			(*b2MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse.Y = float32((*b2MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse.Y * maxImpulse3)
		}
		v34 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse
		v35 = oldImpulse3
		v38 = Vec2{
			X: v34.X - v35.X,
			Y: v34.Y - v35.Y,
		}
		goto _165
	_165:
		impulse3 = v38
		v34 = vA
		v17 = mA
		v35 = impulse3
		v38 = Vec2{
			X: v34.X - float32(v17*v35.X),
			Y: v34.Y - float32(v17*v35.Y),
		}
		goto _170
	_170:
		vA = v38
		v34 = rA
		v35 = impulse3
		v17 = float32(v34.X*v35.Y) - float32(v34.Y*v35.X)
		goto _174
	_174:
		wA = wA - float32(iA*v17)
		v34 = vB
		v17 = mB
		v35 = impulse3
		v38 = Vec2{
			X: v34.X + float32(v17*v35.X),
			Y: v34.Y + float32(v17*v35.Y),
		}
		goto _179
	_179:
		vB = v38
		v34 = rB
		v35 = impulse3
		v17 = float32(v34.X*v35.Y) - float32(v34.Y*v35.X)
		goto _183
	_183:
		wB = wB + float32(iB*v17)
	}
	if (*b2BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b2_dynamicFlag) != 0 {
		(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
		(*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	}
	if (*b2BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b2_dynamicFlag) != 0 {
		(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
		(*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
	}
}

/** @} */

func b2MoverJoint_SetLinearVelocity(tls *_Stack, jointId JointId, velocity Vec2) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_MoverJointSetLinearVelocity
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidVec2(tls, velocity)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+16197, uintptr(unsafe.Pointer(&__func__46))))
		if v1 = !!(b2IsValidVec2(tls, velocity) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+16197, __ccgo_ts+16388, int32FromInt32(16))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_MoverJointSetLinearVelocity)(unsafe.Add(nil, bp)) = b2RecArgs_MoverJointSetLinearVelocity{
			Joint:    jointId,
			Velocity: velocity,
		}
		b2RecWrite_MoverJointSetLinearVelocity(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_moverJoint))
	(*(*b2MoverJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LinearVelocity = velocity
}

func b2MoverJoint_GetLinearVelocity(tls *_Stack, jointId JointId) (r Vec2) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_moverJoint))
	return (*(*b2MoverJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LinearVelocity
}

func b2MoverJoint_SetMaxVelocityForce(tls *_Stack, jointId JointId, maxForce Vec2) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_MoverJointSetMaxVelocityForce
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidVec2(tls, maxForce)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+16417, uintptr(unsafe.Pointer(&__func__47))))
		if v1 = !!(b2IsValidVec2(tls, maxForce) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+16417, __ccgo_ts+16388, int32FromInt32(33))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_MoverJointSetMaxVelocityForce)(unsafe.Add(nil, bp)) = b2RecArgs_MoverJointSetMaxVelocityForce{
			Joint:    jointId,
			MaxForce: maxForce,
		}
		b2RecWrite_MoverJointSetMaxVelocityForce(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_moverJoint))
	(*(*b2MoverJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxVelocityForce = maxForce
}

func b2MoverJoint_GetMaxVelocityForce(tls *_Stack, jointId JointId) (r Vec2) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_moverJoint))
	return (*(*b2MoverJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxVelocityForce
}

func b2GetMoverJointForce(tls *_Stack, world uintptr, base uintptr) (r Vec2) {
	var force, v2, v3 Vec2
	var v1 float32
	_, _, _, _ = force, v1, v2, v3
	v1 = (*b2World)(unsafe.Pointer(world)).Inv_h
	v2 = (*(*b2MoverJoint)(unsafe.Add(unsafe.Pointer(base), 92))).LinearVelocityImpulse
	v3 = Vec2{
		X: float32(v1 * v2.X),
		Y: float32(v1 * v2.Y),
	}
	goto _4
_4:
	force = v3
	return force
}

func b2PrepareMoverJoint(tls *_Stack, base uintptr, context uintptr) {
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world uintptr
	var idA, idB, localIndexA, localIndexB, v9 int32
	var k, mA, mB, v11 float32
	var v1, v2, v3, v4 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, bodySimA, bodySimB, idA, idB, joint, k, localIndexA, localIndexB, mA, mB, setA, setB, world, v1, v11, v2, v3, v4, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_moverJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+16443, __ccgo_ts+16388, int32FromInt32(56))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	idA = (*b2JointSim)(unsafe.Pointer(base)).BodyIdA
	idB = (*b2JointSim)(unsafe.Pointer(base)).BodyIdB
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	if v1 = !!(0 <= idA && idA < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8507, __ccgo_ts+16388, int32FromInt32(63))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idA)*104
	if v2 = !!(0 <= idB && idB < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+8557, __ccgo_ts+16388, int32FromInt32(64))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idB)*104
	if v3 = !!((*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet)); !v3 {
		b2InternalAssert(tls, __ccgo_ts+8607, __ccgo_ts+16388, int32FromInt32(66))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(bodyA)).SetIndex && (*b2Body)(unsafe.Pointer(bodyA)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8672, __ccgo_ts+16388, int32FromInt32(68))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	setA = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(bodyA)).SetIndex)*88
	if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(bodyB)).SetIndex && (*b2Body)(unsafe.Pointer(bodyB)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+8750, __ccgo_ts+16388, int32FromInt32(69))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	setB = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(bodyB)).SetIndex)*88
	localIndexA = (*b2Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b2Body)(unsafe.Pointer(bodyB)).LocalIndex
	if v3 = !!(0 <= localIndexA && localIndexA < (*b2SolverSet)(unsafe.Pointer(setA)).BodySims.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+8828, __ccgo_ts+16388, int32FromInt32(74))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	bodySimA = (*b2SolverSet)(unsafe.Pointer(setA)).BodySims.Data + uintptr(localIndexA)*96
	if v4 = !!(0 <= localIndexB && localIndexB < (*b2SolverSet)(unsafe.Pointer(setB)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+8895, __ccgo_ts+16388, int32FromInt32(75))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySimB = (*b2SolverSet)(unsafe.Pointer(setB)).BodySims.Data + uintptr(localIndexB)*96
	mA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvMass
	mB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvMass
	(*b2JointSim)(unsafe.Pointer(base)).InvMassA = mA
	(*b2JointSim)(unsafe.Pointer(base)).InvMassB = mB
	(*b2JointSim)(unsafe.Pointer(base)).InvIA = float32FromFloat32(0)
	(*b2JointSim)(unsafe.Pointer(base)).InvIB = float32FromFloat32(0)
	joint = base + 92
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) {
		v9 = localIndexA
	} else {
		v9 = -int32(1)
	}
	(*b2MoverJoint)(unsafe.Pointer(joint)).IndexA = v9
	if (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) {
		v9 = localIndexB
	} else {
		v9 = -int32(1)
	}
	(*b2MoverJoint)(unsafe.Pointer(joint)).IndexB = v9
	k = mA + mB
	if k > float32FromFloat32(0) {
		v11 = float32FromFloat32(1) / k
	} else {
		v11 = float32FromFloat32(0)
	}
	(*b2MoverJoint)(unsafe.Pointer(joint)).LinearMass = v11
	if int32FromUint8((*b2StepContext)(unsafe.Pointer(context)).EnableWarmStarting) == false1 {
		(*b2MoverJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse = b2Vec2_zero
	}
}

func b2WarmStartMoverJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var joint, stateA, stateB, v2, v3 uintptr
	var mA, mB, v5 float32
	var v1 bool
	var v4, v6, v7 Vec2
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _ = joint, mA, mB, stateA, stateB, v1, v2, v3, v4, v5, v6, v7
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_moverJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+16443, __ccgo_ts+16388, int32FromInt32(100))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	joint = base + 92
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Add(nil, bp)) = b2_identityBodyState10
	if (*b2MoverJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2MoverJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v2
	if (*b2MoverJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2MoverJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v3
	if (*b2BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b2_dynamicFlag) != 0 {
		v4 = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
		v5 = mA
		v6 = (*b2MoverJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse
		v7 = Vec2{
			X: v4.X - float32(v5*v6.X),
			Y: v4.Y - float32(v5*v6.Y),
		}
		goto _8
	_8:
		(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = v7
	}
	if (*b2BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b2_dynamicFlag) != 0 {
		v4 = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
		v5 = mB
		v6 = (*b2MoverJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse
		v7 = Vec2{
			X: v4.X + float32(v5*v6.X),
			Y: v4.Y + float32(v5*v6.Y),
		}
		goto _13
	_13:
		(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = v7
	}
}

func b2SolveMoverJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var b5, c, cdot, impulse, maxImpulse, oldImpulse, vA, vB, v10, v13, v14, v4, v5, v6, v8, v9 Vec2
	var joint, stateA, stateB, v2, v3 uintptr
	var mA, mB, v12, v20, v30, v31, v32, v33, v35, v36, v37, v38, v39, v40, v42 float32
	var v1 bool
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b5, c, cdot, impulse, joint, mA, mB, maxImpulse, oldImpulse, stateA, stateB, vA, vB, v1, v10, v12, v13, v14, v2, v20, v3, v30, v31, v32, v33, v35, v36, v37, v38, v39, v4, v40, v42, v5, v6, v8, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_moverJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+16443, __ccgo_ts+16388, int32FromInt32(126))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Add(nil, bp)) = b2_identityBodyState10
	joint = base + 92
	if (*b2MoverJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2MoverJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v2
	if (*b2MoverJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2MoverJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v3
	vA = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	// linear velocity
	if (*b2MoverJoint)(unsafe.Pointer(joint)).MaxVelocityForce.X > float32FromFloat32(0) || (*b2MoverJoint)(unsafe.Pointer(joint)).MaxVelocityForce.Y > float32FromFloat32(0) {
		v4 = vB
		v5 = vA
		v6 = Vec2{
			X: v4.X - v5.X,
			Y: v4.Y - v5.Y,
		}
		goto _7
	_7:
		cdot = v6
		v8 = cdot
		v9 = (*b2MoverJoint)(unsafe.Pointer(joint)).LinearVelocity
		v10 = Vec2{
			X: v8.X - v9.X,
			Y: v8.Y - v9.Y,
		}
		goto _11
	_11:
		cdot = v10
		v12 = (*b2MoverJoint)(unsafe.Pointer(joint)).LinearMass
		v4 = cdot
		v5 = Vec2{
			X: float32(v12 * v4.X),
			Y: float32(v12 * v4.Y),
		}
		goto _15
	_15:
		b5 = v5
		impulse = Vec2{
			X: -b5.X,
			Y: -b5.Y,
		}
		oldImpulse = (*b2MoverJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse
		v6 = (*b2MoverJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse
		v8 = impulse
		v9 = Vec2{
			X: v6.X + v8.X,
			Y: v6.Y + v8.Y,
		}
		goto _19
	_19:
		(*b2MoverJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse = v9
		v12 = (*b2StepContext)(unsafe.Pointer(context)).H
		v4 = (*b2MoverJoint)(unsafe.Pointer(joint)).MaxVelocityForce
		v5 = Vec2{
			X: float32(v12 * v4.X),
			Y: float32(v12 * v4.Y),
		}
		goto _23
	_23:
		maxImpulse = v5
		// clamp x and y separately
		v6 = maxImpulse
		v8 = Vec2{
			X: -v6.X,
			Y: -v6.Y,
		}
		goto _26
	_26:
		v9 = (*b2MoverJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse
		v10 = v8
		v13 = maxImpulse
		v20 = v9.X
		v30 = v10.X
		v31 = v13.X
		if v20 < v30 {
			v33 = v30
		} else {
			if v20 > v31 {
				v35 = v31
			} else {
				v35 = v20
			}
			v33 = v35
		}
		v32 = v33
		goto _34
	_34:
		c.X = v32
		v36 = v9.Y
		v37 = v10.Y
		v38 = v13.Y
		if v36 < v37 {
			v40 = v37
		} else {
			if v36 > v38 {
				v42 = v38
			} else {
				v42 = v36
			}
			v40 = v42
		}
		v39 = v40
		goto _41
	_41:
		c.Y = v39
		v14 = c
		goto _45
	_45:
		(*b2MoverJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse = v14
		v4 = (*b2MoverJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse
		v5 = oldImpulse
		v6 = Vec2{
			X: v4.X - v5.X,
			Y: v4.Y - v5.Y,
		}
		goto _49
	_49:
		impulse = v6
		v4 = vA
		v12 = mA
		v5 = impulse
		v6 = Vec2{
			X: v4.X - float32(v12*v5.X),
			Y: v4.Y - float32(v12*v5.Y),
		}
		goto _54
	_54:
		vA = v6
		v4 = vB
		v12 = mB
		v5 = impulse
		v6 = Vec2{
			X: v4.X + float32(v12*v5.X),
			Y: v4.Y + float32(v12*v5.Y),
		}
		goto _59
	_59:
		vB = v6
	} else {
		(*b2MoverJoint)(unsafe.Pointer(joint)).LinearVelocityImpulse = b2Vec2_zero
	}
	if (*b2BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b2_dynamicFlag) != 0 {
		(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
	}
	if (*b2BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b2_dynamicFlag) != 0 {
		(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
	}
}

func b2Joint_IsValid(tls *_Stack, id JointId) (r uint8) {
	var joint, world uintptr
	var jointId int32
	var v1 bool
	_, _, _, _ = joint, jointId, world, v1
	if int32(_B2_MAX_WORLDS) <= int32FromUint16(id.World0) {
		return boolUint8(false1 != 0)
	}
	world = uintptr(unsafe.Pointer(&b2_worlds)) + uintptr(id.World0)*2792
	if int32FromUint16((*b2World)(unsafe.Pointer(world)).WorldId) != int32FromUint16(id.World0) {
		// world is free
		return boolUint8(false1 != 0)
	}
	jointId = id.Index1 - int32(1)
	if jointId < 0 || (*b2World)(unsafe.Pointer(world)).Joints.Count <= jointId {
		return boolUint8(false1 != 0)
	}
	joint = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr(jointId)*72
	if (*b2Joint)(unsafe.Pointer(joint)).JointId == -int32(1) {
		// joint is free
		return boolUint8(false1 != 0)
	}
	if v1 = !!((*b2Joint)(unsafe.Pointer(joint)).JointId == jointId); !v1 {
		b2InternalAssert(tls, __ccgo_ts+18476, __ccgo_ts+16660, int32FromInt32(1708))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return boolUint8(int32FromUint16(id.Generation) == int32FromUint16((*b2Joint)(unsafe.Pointer(joint)).Generation))
}

func b2PogoJoint_SetRestLength(tls *_Stack, jointId JointId, length float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_PogoJointSetRestLength
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, length)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8172, uintptr(unsafe.Pointer(&__func__52))))
		if v1 = !!(b2IsValidFloat(tls, length) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8172, __ccgo_ts+20167, int32FromInt32(18))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_PogoJointSetRestLength)(unsafe.Add(nil, bp)) = b2RecArgs_PogoJointSetRestLength{
			Joint:  jointId,
			Length: length,
		}
		b2RecWrite_PogoJointSetRestLength(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_pogoJoint))
	(*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).RestLength = length
}

func b2PogoJoint_GetRestLength(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_pogoJoint))
	return (*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).RestLength
}

func b2PogoJoint_GetSpringHertz(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_pogoJoint))
	return (*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).Hertz
}

func b2PogoJoint_SetSpringHertz(tls *_Stack, jointId JointId, hertz float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_PogoJointSetSpringHertz
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, hertz)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8368, uintptr(unsafe.Pointer(&__func__53))))
		if v1 = !!(b2IsValidFloat(tls, hertz) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8368, __ccgo_ts+20167, int32FromInt32(41))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_PogoJointSetSpringHertz)(unsafe.Add(nil, bp)) = b2RecArgs_PogoJointSetSpringHertz{
			Joint: jointId,
			Hertz: hertz,
		}
		b2RecWrite_PogoJointSetSpringHertz(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_pogoJoint))
	(*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).Hertz = hertz
}

func b2PogoJoint_SetSpringDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_PogoJointSetSpringDampingRatio
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, dampingRatio)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8392, uintptr(unsafe.Pointer(&__func__54))))
		if v1 = !!(b2IsValidFloat(tls, dampingRatio) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8392, __ccgo_ts+20167, int32FromInt32(52))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_PogoJointSetSpringDampingRatio)(unsafe.Add(nil, bp)) = b2RecArgs_PogoJointSetSpringDampingRatio{
			Joint:        jointId,
			DampingRatio: dampingRatio,
		}
		b2RecWrite_PogoJointSetSpringDampingRatio(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_pogoJoint))
	(*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).DampingRatio = dampingRatio
}

func b2PogoJoint_GetSpringDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_pogoJoint))
	return (*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).DampingRatio
}

func b2PogoJoint_GetCurrentLength(tls *_Stack, jointId JointId) (r1 float32) {
	var axis, d, pA, pB, v10, v11, v14, v15, v18, v19, v22, v23, v25 Vec2
	var jointSim, world uintptr
	var length, x, y, v31 float32
	var r, transformA, transformB, v3, v7 Transform
	var wxfA, v1, v5 WorldTransform
	var v13, v9 Rot
	var v2, v6 Pos
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axis, d, jointSim, length, pA, pB, r, transformA, transformB, world, wxfA, x, y, v1, v10, v11, v13, v14, v15, v18, v19, v2, v22, v23, v25, v3, v31, v5, v6, v7, v9
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	jointSim = b2GetJointSimCheckType(tls, jointId, int32(b2_pogoJoint))
	// Relative to body A so the difference stays in float precision far from the origin
	wxfA = b2GetBodyTransform(tls, world, (*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdA)
	v1 = wxfA
	v2 = wxfA.P
	r.Q = v1.Q
	r.P = Vec2{
		X: v1.P.X - v2.X,
		Y: v1.P.Y - v2.Y,
	}
	v3 = r
	goto _4
_4:
	transformA = v3
	v5 = b2GetBodyTransform(tls, world, (*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdB)
	v6 = wxfA.P
	r.Q = v5.Q
	r.P = Vec2{
		X: v5.P.X - v6.X,
		Y: v5.P.Y - v6.Y,
	}
	v7 = r
	goto _8
_8:
	transformB = v7
	v9 = (*b2JointSim)(unsafe.Pointer(jointSim)).LocalFrameB.Q
	v10 = b2_pogoAxis
	v11 = Vec2{
		X: float32(v9.C*v10.X) - float32(v9.S*v10.Y),
		Y: float32(v9.S*v10.X) + float32(v9.C*v10.Y),
	}
	goto _12
_12:
	axis = v11
	v13 = transformB.Q
	v14 = axis
	v15 = Vec2{
		X: float32(v13.C*v14.X) - float32(v13.S*v14.Y),
		Y: float32(v13.S*v14.X) + float32(v13.C*v14.Y),
	}
	goto _16
_16:
	axis = v15
	v3 = transformA
	v10 = (*b2JointSim)(unsafe.Pointer(jointSim)).LocalFrameA.P
	x = float32(v3.Q.C*v10.X) - float32(v3.Q.S*v10.Y) + v3.P.X
	y = float32(v3.Q.S*v10.X) + float32(v3.Q.C*v10.Y) + v3.P.Y
	v11 = Vec2{
		X: x,
		Y: y,
	}
	goto _20
_20:
	pA = v11
	v7 = transformB
	v14 = (*b2JointSim)(unsafe.Pointer(jointSim)).LocalFrameB.P
	x = float32(v7.Q.C*v14.X) - float32(v7.Q.S*v14.Y) + v7.P.X
	y = float32(v7.Q.S*v14.X) + float32(v7.Q.C*v14.Y) + v7.P.Y
	v15 = Vec2{
		X: x,
		Y: y,
	}
	goto _24
_24:
	pB = v15
	v18 = pB
	v19 = pA
	v22 = Vec2{
		X: v18.X - v19.X,
		Y: v18.Y - v19.Y,
	}
	goto _28
_28:
	d = v22
	v23 = d
	v25 = axis
	v31 = float32(v23.X*v25.X) + float32(v23.Y*v25.Y)
	goto _32
_32:
	length = v31
	return length
}

func b2PogoJoint_GetImpulse(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_pogoJoint))
	return (*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).Impulse
}

func b2PogoJoint_GetCurrentSpeed(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_pogoJoint))
	return (*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).Velocity
}

func b2GetPogoJointForce(tls *_Stack, world uintptr, base uintptr) (r Vec2) {
	var force, v2, v3 Vec2
	var v1 float32
	_, _, _, _ = force, v1, v2, v3
	v1 = float32((*b2World)(unsafe.Pointer(world)).Inv_h * (*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(base), 92))).Impulse)
	v2 = (*(*b2PogoJoint)(unsafe.Add(unsafe.Pointer(base), 92))).Normal
	v3 = Vec2{
		X: float32(v1 * v2.X),
		Y: float32(v1 * v2.Y),
	}
	goto _4
_4:
	force = v3
	return force
}

func b2PreparePogoJoint(tls *_Stack, base uintptr, context uintptr) {
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world uintptr
	var crA, crB, iA, iB, k, mA, mB, v41, v45, v47 float32
	var idA, idB, localIndexA, localIndexB, v9 int32
	var qr, v11, v12, v13 Rot
	var rA, rB, v15, v16, v17, v20, v21 Vec2
	var v1, v2, v3, v4 bool
	var v35, v36 Pos
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, bodySimA, bodySimB, crA, crB, iA, iB, idA, idB, joint, k, localIndexA, localIndexB, mA, mB, qr, rA, rB, setA, setB, world, v1, v11, v12, v13, v15, v16, v17, v2, v20, v21, v3, v35, v36, v4, v41, v45, v47, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_pogoJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20195, __ccgo_ts+20167, int32FromInt32(106))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	// chase body id to the solver set where the body lives
	idA = (*b2JointSim)(unsafe.Pointer(base)).BodyIdA
	idB = (*b2JointSim)(unsafe.Pointer(base)).BodyIdB
	if v1 = !!(0 <= idA && idA < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8507, __ccgo_ts+20167, int32FromInt32(114))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idA)*104
	if v2 = !!(0 <= idB && idB < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+8557, __ccgo_ts+20167, int32FromInt32(115))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idB)*104
	if v3 = !!((*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet)); !v3 {
		b2InternalAssert(tls, __ccgo_ts+8607, __ccgo_ts+20167, int32FromInt32(117))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(bodyA)).SetIndex && (*b2Body)(unsafe.Pointer(bodyA)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8672, __ccgo_ts+20167, int32FromInt32(119))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	setA = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(bodyA)).SetIndex)*88
	if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(bodyB)).SetIndex && (*b2Body)(unsafe.Pointer(bodyB)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+8750, __ccgo_ts+20167, int32FromInt32(120))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	setB = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(bodyB)).SetIndex)*88
	localIndexA = (*b2Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b2Body)(unsafe.Pointer(bodyB)).LocalIndex
	if v3 = !!(0 <= localIndexA && localIndexA < (*b2SolverSet)(unsafe.Pointer(setA)).BodySims.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+8828, __ccgo_ts+20167, int32FromInt32(125))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	bodySimA = (*b2SolverSet)(unsafe.Pointer(setA)).BodySims.Data + uintptr(localIndexA)*96
	if v4 = !!(0 <= localIndexB && localIndexB < (*b2SolverSet)(unsafe.Pointer(setB)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+8895, __ccgo_ts+20167, int32FromInt32(126))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySimB = (*b2SolverSet)(unsafe.Pointer(setB)).BodySims.Data + uintptr(localIndexB)*96
	mA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvMass
	iA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvInertia
	mB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvMass
	iB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvInertia
	(*b2JointSim)(unsafe.Pointer(base)).InvMassA = mA
	(*b2JointSim)(unsafe.Pointer(base)).InvMassB = mB
	(*b2JointSim)(unsafe.Pointer(base)).InvIA = iA
	(*b2JointSim)(unsafe.Pointer(base)).InvIB = iB
	joint = base + 92
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) {
		v9 = localIndexA
	} else {
		v9 = -int32(1)
	}
	(*b2PogoJoint)(unsafe.Pointer(joint)).IndexA = v9
	if (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) {
		v9 = localIndexB
	} else {
		v9 = -int32(1)
	}
	(*b2PogoJoint)(unsafe.Pointer(joint)).IndexB = v9
	// Compute joint anchor frames with world space rotation, relative to center of mass
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v12 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	qr.S = float32(v11.S*v12.C) + float32(v11.C*v12.S)
	qr.C = float32(v11.C*v12.C) - float32(v11.S*v12.S)
	v13 = qr
	goto _14
_14:
	(*b2PogoJoint)(unsafe.Pointer(joint)).FrameA.Q = v13
	v15 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA.P
	v16 = (*b2BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v17 = Vec2{
		X: v15.X - v16.X,
		Y: v15.Y - v16.Y,
	}
	goto _18
_18:
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v20 = v17
	v21 = Vec2{
		X: float32(v11.C*v20.X) - float32(v11.S*v20.Y),
		Y: float32(v11.S*v20.X) + float32(v11.C*v20.Y),
	}
	goto _22
_22:
	(*b2PogoJoint)(unsafe.Pointer(joint)).FrameA.P = v21
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v12 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameB.Q
	qr.S = float32(v11.S*v12.C) + float32(v11.C*v12.S)
	qr.C = float32(v11.C*v12.C) - float32(v11.S*v12.S)
	v13 = qr
	goto _26
_26:
	(*b2PogoJoint)(unsafe.Pointer(joint)).FrameB.Q = v13
	v15 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v16 = (*b2BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v17 = Vec2{
		X: v15.X - v16.X,
		Y: v15.Y - v16.Y,
	}
	goto _30
_30:
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v20 = v17
	v21 = Vec2{
		X: float32(v11.C*v20.X) - float32(v11.S*v20.Y),
		Y: float32(v11.S*v20.X) + float32(v11.C*v20.Y),
	}
	goto _34
_34:
	(*b2PogoJoint)(unsafe.Pointer(joint)).FrameB.P = v21
	// Compute the initial center delta. Incremental position updates are relative to this.
	v35 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Center
	v36 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Center
	v15 = Vec2{
		X: v35.X - v36.X,
		Y: v35.Y - v36.Y,
	}
	goto _38
_38:
	(*b2PogoJoint)(unsafe.Pointer(joint)).DeltaCenter = v15
	rA = (*b2PogoJoint)(unsafe.Pointer(joint)).FrameA.P
	rB = (*b2PogoJoint)(unsafe.Pointer(joint)).FrameB.P
	v15 = rA
	v16 = (*b2PogoJoint)(unsafe.Pointer(joint)).Normal
	v41 = float32(v15.X*v16.Y) - float32(v15.Y*v16.X)
	goto _42
_42:
	// compute effective mass
	crA = v41
	v17 = rB
	v20 = (*b2PogoJoint)(unsafe.Pointer(joint)).Normal
	v45 = float32(v17.X*v20.Y) - float32(v17.Y*v20.X)
	goto _46
_46:
	crB = v45
	k = mA + mB + float32(float32(iA*crA)*crA) + float32(float32(iB*crB)*crB)
	if k > float32FromFloat32(0) {
		v47 = float32FromFloat32(1) / k
	} else {
		v47 = float32FromFloat32(0)
	}
	(*b2PogoJoint)(unsafe.Pointer(joint)).LinearMass = v47
	if int32FromUint8((*b2StepContext)(unsafe.Pointer(context)).EnableWarmStarting) == false1 {
		(*b2PogoJoint)(unsafe.Pointer(joint)).Impulse = float32FromFloat32(0)
	}
}

func b2WarmStartPogoJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var iA, iB, mA, mB, v12, v17 float32
	var joint, stateA, stateB, v2, v3 uintptr
	var linearImpulse, rA, rB, v10, v13, v14, v16, v18, v19, v5, v6, v9 Vec2
	var v1 bool
	var v4, v8 Rot
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = iA, iB, joint, linearImpulse, mA, mB, rA, rB, stateA, stateB, v1, v10, v12, v13, v14, v16, v17, v18, v19, v2, v3, v4, v5, v6, v8, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_pogoJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20195, __ccgo_ts+20167, int32FromInt32(168))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	joint = base + 92
	if (*b2PogoJoint)(unsafe.Pointer(joint)).Hertz == float32FromFloat32(0) {
		(*b2PogoJoint)(unsafe.Pointer(joint)).Impulse = float32FromFloat32(0)
		return
	}
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Add(nil, bp)) = b2_identityBodyState12
	if (*b2PogoJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2PogoJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v2
	if (*b2PogoJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2PogoJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v3
	v4 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b2PogoJoint)(unsafe.Pointer(joint)).FrameA.P
	v6 = Vec2{
		X: float32(v4.C*v5.X) - float32(v4.S*v5.Y),
		Y: float32(v4.S*v5.X) + float32(v4.C*v5.Y),
	}
	goto _7
_7:
	rA = v6
	v8 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v9 = (*b2PogoJoint)(unsafe.Pointer(joint)).FrameB.P
	v10 = Vec2{
		X: float32(v8.C*v9.X) - float32(v8.S*v9.Y),
		Y: float32(v8.S*v9.X) + float32(v8.C*v9.Y),
	}
	goto _11
_11:
	rB = v10
	v12 = (*b2PogoJoint)(unsafe.Pointer(joint)).Impulse
	v13 = (*b2PogoJoint)(unsafe.Pointer(joint)).Normal
	v14 = Vec2{
		X: float32(v12 * v13.X),
		Y: float32(v12 * v13.Y),
	}
	goto _15
_15:
	linearImpulse = v14
	if (*b2BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b2_dynamicFlag) != 0 {
		v16 = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
		v17 = mA
		v18 = linearImpulse
		v19 = Vec2{
			X: v16.X - float32(v17*v18.X),
			Y: v16.Y - float32(v17*v18.Y),
		}
		goto _20
	_20:
		(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = v19
		v5 = rA
		v6 = linearImpulse
		v12 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		goto _24
	_24:
		*(*float32)(unsafe.Add(nil, stateA+8)) -= float32(iA * v12)
	}
	if (*b2BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b2_dynamicFlag) != 0 {
		v5 = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
		v12 = mB
		v6 = linearImpulse
		v9 = Vec2{
			X: v5.X + float32(v12*v6.X),
			Y: v5.Y + float32(v12*v6.Y),
		}
		goto _29
	_29:
		(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = v9
		v5 = rB
		v6 = linearImpulse
		v12 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		goto _33
	_33:
		*(*float32)(unsafe.Add(nil, stateB+8)) += float32(iB * v12)
	}
}

func b2SolvePogoJoint(tls *_Stack, base uintptr, context uintptr, useBias uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var P, d, dcA, dcB, pogoAxis, rA, rB, vA, vB, vr, v10, v12, v13, v14, v16, v17, v18, v20, v21, v22, v24, v25, v26, v29, v30, v32, v33, v5, v6, v9 Vec2
	var bias, c, cdot, iA, iB, impulse, mA, mB, maxCompressionImpulse, maxTensionImpulse, oldImpulse, omega, omegaH, wA, wB, v34, v36, v38, v46, v60, v62, v63, v64, v65 float32
	var joint, stateA, stateB, v2, v3 uintptr
	var v1 bool
	var v28, v4, v8 Rot
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = P, bias, c, cdot, d, dcA, dcB, iA, iB, impulse, joint, mA, mB, maxCompressionImpulse, maxTensionImpulse, oldImpulse, omega, omegaH, pogoAxis, rA, rB, stateA, stateB, vA, vB, vr, wA, wB, v1, v10, v12, v13, v14, v16, v17, v18, v2, v20, v21, v22, v24, v25, v26, v28, v29, v3, v30, v32, v33, v34, v36, v38, v4, v46, v5, v6, v60, v62, v63, v64, v65, v8, v9
	_ = uint64FromInt64(4)
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_pogoJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20195, __ccgo_ts+20167, int32FromInt32(210))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Add(nil, bp)) = b2_identityBodyState12
	joint = base + 92
	if (*b2PogoJoint)(unsafe.Pointer(joint)).Hertz == float32FromFloat32(0) {
		(*b2PogoJoint)(unsafe.Pointer(joint)).Impulse = float32FromFloat32(0)
		return
	}
	if (*b2PogoJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2PogoJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v2
	if (*b2PogoJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2PogoJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v3
	vA = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	v4 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b2PogoJoint)(unsafe.Pointer(joint)).FrameA.P
	v6 = Vec2{
		X: float32(v4.C*v5.X) - float32(v4.S*v5.Y),
		Y: float32(v4.S*v5.X) + float32(v4.C*v5.Y),
	}
	goto _7
_7:
	rA = v6
	v8 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v9 = (*b2PogoJoint)(unsafe.Pointer(joint)).FrameB.P
	v10 = Vec2{
		X: float32(v8.C*v9.X) - float32(v8.S*v9.Y),
		Y: float32(v8.S*v9.X) + float32(v8.C*v9.Y),
	}
	goto _11
_11:
	rB = v10
	bias = float32FromFloat32(0)
	if useBias != 0 {
		// This is a custom pogo position correction designed to avoid feeding velocity into
		// the mover body. Otherwise the mover can make huge jumps going up stairs. This still
		// works with continuous collision.
		dcA = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaPosition
		dcB = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
		v12 = dcB
		v13 = dcA
		v14 = Vec2{
			X: v12.X - v13.X,
			Y: v12.Y - v13.Y,
		}
		goto _15
	_15:
		v16 = rB
		v17 = rA
		v18 = Vec2{
			X: v16.X - v17.X,
			Y: v16.Y - v17.Y,
		}
		goto _19
	_19:
		v20 = v14
		v21 = v18
		v22 = Vec2{
			X: v20.X + v21.X,
			Y: v20.Y + v21.Y,
		}
		goto _23
	_23:
		v24 = v22
		v25 = (*b2PogoJoint)(unsafe.Pointer(joint)).DeltaCenter
		v26 = Vec2{
			X: v24.X + v25.X,
			Y: v24.Y + v25.Y,
		}
		goto _27
	_27:
		d = v26
		v28 = (*b2PogoJoint)(unsafe.Pointer(joint)).FrameB.Q
		v29 = b2_pogoAxis
		v30 = Vec2{
			X: float32(v28.C*v29.X) - float32(v28.S*v29.Y),
			Y: float32(v28.S*v29.X) + float32(v28.C*v29.Y),
		}
		goto _31
	_31:
		// The pogo axis is usually fixed as the world y-axis. Different than the ground normal.
		pogoAxis = v30
		v32 = pogoAxis
		v33 = d
		v34 = float32(v32.X*v33.X) + float32(v32.Y*v33.Y)
		goto _35
	_35:
		c = v34 - (*b2PogoJoint)(unsafe.Pointer(joint)).RestLength
		// This could be divided by dot(joint->normal, pogoAxis) to account for the constraint
		// direction, but I'd rather diminish the pogo recovery rate than risk making it huge.
		omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(_B2_PI)) * (*b2PogoJoint)(unsafe.Pointer(joint)).Hertz)
		omegaH = float32(omega * (*b2StepContext)(unsafe.Pointer(context)).H)
		v36 = ((*b2PogoJoint)(unsafe.Pointer(joint)).Velocity - float32(float32(omega*omegaH)*c)) / (float32FromFloat32(1) + float32(float32(float32FromFloat32(2)*(*b2PogoJoint)(unsafe.Pointer(joint)).DampingRatio)*omegaH) + float32(omegaH*omegaH))
		goto _37
	_37:
		(*b2PogoJoint)(unsafe.Pointer(joint)).Velocity = v36
		bias = -(*b2PogoJoint)(unsafe.Pointer(joint)).Velocity
	}
	v34 = wB
	v5 = rB
	v6 = Vec2{
		X: float32(-v34 * v5.Y),
		Y: float32(v34 * v5.X),
	}
	goto _41
_41:
	v9 = vB
	v10 = v6
	v12 = Vec2{
		X: v9.X + v10.X,
		Y: v9.Y + v10.Y,
	}
	goto _45
_45:
	v36 = wA
	v13 = rA
	v14 = Vec2{
		X: float32(-v36 * v13.Y),
		Y: float32(v36 * v13.X),
	}
	goto _49
_49:
	v16 = vA
	v17 = v14
	v18 = Vec2{
		X: v16.X + v17.X,
		Y: v16.Y + v17.Y,
	}
	goto _53
_53:
	v20 = v12
	v21 = v18
	v22 = Vec2{
		X: v20.X - v21.X,
		Y: v20.Y - v21.Y,
	}
	goto _57
_57:
	vr = v22
	v24 = (*b2PogoJoint)(unsafe.Pointer(joint)).Normal
	v25 = vr
	v38 = float32(v24.X*v25.X) + float32(v24.Y*v25.Y)
	goto _61
_61:
	cdot = v38
	maxTensionImpulse = float32((*b2StepContext)(unsafe.Pointer(context)).H * (*b2PogoJoint)(unsafe.Pointer(joint)).MaxTensionForce)
	maxCompressionImpulse = float32((*b2StepContext)(unsafe.Pointer(context)).H * (*b2PogoJoint)(unsafe.Pointer(joint)).MaxCompressionForce)
	oldImpulse = (*b2PogoJoint)(unsafe.Pointer(joint)).Impulse
	impulse = float32(-(*b2PogoJoint)(unsafe.Pointer(joint)).LinearMass * (cdot + bias))
	v46 = (*b2PogoJoint)(unsafe.Pointer(joint)).Impulse + impulse
	v60 = -maxTensionImpulse
	v62 = maxCompressionImpulse
	if v46 < v60 {
		v64 = v60
	} else {
		if v46 > v62 {
			v65 = v62
		} else {
			v65 = v46
		}
		v64 = v65
	}
	v63 = v64
	goto _66
_66:
	(*b2PogoJoint)(unsafe.Pointer(joint)).Impulse = v63
	impulse = (*b2PogoJoint)(unsafe.Pointer(joint)).Impulse - oldImpulse
	v34 = impulse
	v5 = (*b2PogoJoint)(unsafe.Pointer(joint)).Normal
	v6 = Vec2{
		X: float32(v34 * v5.X),
		Y: float32(v34 * v5.Y),
	}
	goto _72
_72:
	P = v6
	v9 = vA
	v36 = mA
	v10 = P
	v12 = Vec2{
		X: v9.X - float32(v36*v10.X),
		Y: v9.Y - float32(v36*v10.Y),
	}
	goto _77
_77:
	vA = v12
	v5 = rA
	v6 = P
	v34 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	goto _81
_81:
	wA = wA - float32(iA*v34)
	v5 = vB
	v34 = mB
	v6 = P
	v9 = Vec2{
		X: v5.X + float32(v34*v6.X),
		Y: v5.Y + float32(v34*v6.Y),
	}
	goto _86
_86:
	vB = v9
	v5 = rB
	v6 = P
	v34 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
	goto _90
_90:
	wB = wB + float32(iB*v34)
	if (*b2BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b2_dynamicFlag) != 0 {
		(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
		(*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	}
	if (*b2BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b2_dynamicFlag) != 0 {
		(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
		(*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
	}
}

/** @} */

func b2PrismaticJoint_EnableSpring(tls *_Stack, jointId JointId, enableSpring uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* _a at bp+0 */ b2RecArgs_PrismaticJointEnableSpring
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_PrismaticJointEnableSpring)(unsafe.Add(nil, bp)) = b2RecArgs_PrismaticJointEnableSpring{
			Joint:        jointId,
			EnableSpring: enableSpring,
		}
		b2RecWrite_PrismaticJointEnableSpring(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	if int32FromUint8(enableSpring) != int32FromUint8((*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableSpring) {
		(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableSpring = enableSpring
		(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).SpringImpulse = float32FromFloat32(0)
	}
}

func b2PrismaticJoint_IsSpringEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableSpring
}

func b2PrismaticJoint_SetSpringHertz(tls *_Stack, jointId JointId, hertz float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_PrismaticJointSetSpringHertz
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, hertz)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8368, uintptr(unsafe.Pointer(&__func__55))))
		if v1 = !!(b2IsValidFloat(tls, hertz) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8368, __ccgo_ts+20222, int32FromInt32(35))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_PrismaticJointSetSpringHertz)(unsafe.Add(nil, bp)) = b2RecArgs_PrismaticJointSetSpringHertz{
			Joint: jointId,
			Hertz: hertz,
		}
		b2RecWrite_PrismaticJointSetSpringHertz(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).Hertz = hertz
}

func b2PrismaticJoint_GetSpringHertz(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).Hertz
}

func b2PrismaticJoint_SetSpringDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_PrismaticJointSetSpringDampingRatio
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, dampingRatio)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8392, uintptr(unsafe.Pointer(&__func__56))))
		if v1 = !!(b2IsValidFloat(tls, dampingRatio) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8392, __ccgo_ts+20222, int32FromInt32(51))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_PrismaticJointSetSpringDampingRatio)(unsafe.Add(nil, bp)) = b2RecArgs_PrismaticJointSetSpringDampingRatio{
			Joint:        jointId,
			DampingRatio: dampingRatio,
		}
		b2RecWrite_PrismaticJointSetSpringDampingRatio(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).DampingRatio = dampingRatio
}

func b2PrismaticJoint_GetSpringDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).DampingRatio
}

func b2PrismaticJoint_SetTargetTranslation(tls *_Stack, jointId JointId, translation float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_PrismaticJointSetTargetTranslation
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, translation)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+20255, uintptr(unsafe.Pointer(&__func__57))))
		if v1 = !!(b2IsValidFloat(tls, translation) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+20255, __ccgo_ts+20222, int32FromInt32(67))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_PrismaticJointSetTargetTranslation)(unsafe.Add(nil, bp)) = b2RecArgs_PrismaticJointSetTargetTranslation{
			Joint:       jointId,
			Translation: translation,
		}
		b2RecWrite_PrismaticJointSetTargetTranslation(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).TargetTranslation = translation
}

func b2PrismaticJoint_GetTargetTranslation(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).TargetTranslation
}

func b2PrismaticJoint_EnableLimit(tls *_Stack, jointId JointId, enableLimit uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* _a at bp+0 */ b2RecArgs_PrismaticJointEnableLimit
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_PrismaticJointEnableLimit)(unsafe.Add(nil, bp)) = b2RecArgs_PrismaticJointEnableLimit{
			Joint:       jointId,
			EnableLimit: enableLimit,
		}
		b2RecWrite_PrismaticJointEnableLimit(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	if int32FromUint8(enableLimit) != int32FromUint8((*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableLimit) {
		(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableLimit = enableLimit
		(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LowerImpulse = float32FromFloat32(0)
		(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).UpperImpulse = float32FromFloat32(0)
	}
}

func b2PrismaticJoint_IsLimitEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableLimit
}

func b2PrismaticJoint_GetLowerLimit(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LowerTranslation
}

func b2PrismaticJoint_GetUpperLimit(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).UpperTranslation
}

func b2PrismaticJoint_SetLimits(tls *_Stack, jointId JointId, lower float32, upper float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var v4, v5, v6, v8 float32
	var _ /* _a at bp+0 */ b2RecArgs_PrismaticJointSetLimits
	_, _, _, _, _, _, _ = joint, world, v1, v4, v5, v6, v8
	if int32FromUint8(b2IsValidFloat(tls, lower)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+20285, uintptr(unsafe.Pointer(&__func__210))))
		if v1 = !!(b2IsValidFloat(tls, lower) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+20285, __ccgo_ts+20222, int32FromInt32(114))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if int32FromUint8(b2IsValidFloat(tls, upper)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+20309, uintptr(unsafe.Pointer(&__func__210))))
		if v1 = !!(b2IsValidFloat(tls, upper) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+20309, __ccgo_ts+20222, int32FromInt32(115))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_PrismaticJointSetLimits)(unsafe.Add(nil, bp)) = b2RecArgs_PrismaticJointSetLimits{
			Joint: jointId,
			Lower: lower,
			Upper: upper,
		}
		b2RecWrite_PrismaticJointSetLimits(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	if v1 = !!(lower <= upper); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20333, __ccgo_ts+20222, int32FromInt32(119))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	v4 = lower
	v5 = upper
	if v4 < v5 {
		v8 = v4
	} else {
		v8 = v5
	}
	v6 = v8
	goto _7
_7:
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LowerTranslation = v6
	v4 = lower
	v5 = upper
	if v4 > v5 {
		v8 = v4
	} else {
		v8 = v5
	}
	v6 = v8
	goto _12
_12:
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).UpperTranslation = v6
}

func b2PrismaticJoint_EnableMotor(tls *_Stack, jointId JointId, enableMotor uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* _a at bp+0 */ b2RecArgs_PrismaticJointEnableMotor
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_PrismaticJointEnableMotor)(unsafe.Add(nil, bp)) = b2RecArgs_PrismaticJointEnableMotor{
			Joint:       jointId,
			EnableMotor: enableMotor,
		}
		b2RecWrite_PrismaticJointEnableMotor(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	if int32FromUint8(enableMotor) != int32FromUint8((*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableMotor) {
		(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableMotor = enableMotor
		(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MotorImpulse = float32FromFloat32(0)
	}
}

func b2PrismaticJoint_IsMotorEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableMotor
}

func b2PrismaticJoint_SetTargetMotorSpeed(tls *_Stack, jointId JointId, motorSpeed float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_PrismaticJointSetTargetMotorSpeed
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, motorSpeed)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8423, uintptr(unsafe.Pointer(&__func__58))))
		if v1 = !!(b2IsValidFloat(tls, motorSpeed) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8423, __ccgo_ts+20222, int32FromInt32(146))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_PrismaticJointSetTargetMotorSpeed)(unsafe.Add(nil, bp)) = b2RecArgs_PrismaticJointSetTargetMotorSpeed{
			Joint:      jointId,
			MotorSpeed: motorSpeed,
		}
		b2RecWrite_PrismaticJointSetTargetMotorSpeed(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MotorSpeed = motorSpeed
}

func b2PrismaticJoint_GetTargetMotorSpeed(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MotorSpeed
}

func b2PrismaticJoint_GetCurrentMotorForce(tls *_Stack, jointId JointId) (r float32) {
	var base, world uintptr
	_, _ = base, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return float32((*b2World)(unsafe.Pointer(world)).Inv_h * (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 92))).MotorImpulse)
}

func b2PrismaticJoint_SetMaxMotorForce(tls *_Stack, jointId JointId, force float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_PrismaticJointSetMaxMotorForce
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, force)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8452, uintptr(unsafe.Pointer(&__func__59))))
		if v1 = !!(b2IsValidFloat(tls, force) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8452, __ccgo_ts+20222, int32FromInt32(169))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_PrismaticJointSetMaxMotorForce)(unsafe.Add(nil, bp)) = b2RecArgs_PrismaticJointSetMaxMotorForce{
			Joint: jointId,
			Force: force,
		}
		b2RecWrite_PrismaticJointSetMaxMotorForce(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxMotorForce = force
}

func b2PrismaticJoint_GetMaxMotorForce(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxMotorForce
}

func b2PrismaticJoint_GetCurrentTranslation(tls *_Stack, jointId JointId) (r1 float32) {
	var axisA, d, localAxisA, pA, pB, v10, v11, v14, v15, v18, v19, v22, v23, v25, v26, v27, v29, v30 Vec2
	var jointSim, world uintptr
	var r, transformA, transformB, v17, v21, v3, v7 Transform
	var translation, x, y, v31 float32
	var wxfA, v1, v5 WorldTransform
	var v13, v9 Rot
	var v2, v6 Pos
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axisA, d, jointSim, localAxisA, pA, pB, r, transformA, transformB, translation, world, wxfA, x, y, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v29, v3, v30, v31, v5, v6, v7, v9
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	jointSim = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	// Relative to body A so the difference stays in float precision far from the origin
	wxfA = b2GetBodyTransform(tls, world, (*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdA)
	v1 = wxfA
	v2 = wxfA.P
	r.Q = v1.Q
	r.P = Vec2{
		X: v1.P.X - v2.X,
		Y: v1.P.Y - v2.Y,
	}
	v3 = r
	goto _4
_4:
	transformA = v3
	v5 = b2GetBodyTransform(tls, world, (*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdB)
	v6 = wxfA.P
	r.Q = v5.Q
	r.P = Vec2{
		X: v5.P.X - v6.X,
		Y: v5.P.Y - v6.Y,
	}
	v7 = r
	goto _8
_8:
	transformB = v7
	v9 = (*b2JointSim)(unsafe.Pointer(jointSim)).LocalFrameA.Q
	v10 = Vec2{
		X: float32FromFloat32(1),
	}
	v11 = Vec2{
		X: float32(v9.C*v10.X) - float32(v9.S*v10.Y),
		Y: float32(v9.S*v10.X) + float32(v9.C*v10.Y),
	}
	goto _12
_12:
	localAxisA = v11
	v13 = transformA.Q
	v14 = localAxisA
	v15 = Vec2{
		X: float32(v13.C*v14.X) - float32(v13.S*v14.Y),
		Y: float32(v13.S*v14.X) + float32(v13.C*v14.Y),
	}
	goto _16
_16:
	axisA = v15
	v17 = transformA
	v18 = (*b2JointSim)(unsafe.Pointer(jointSim)).LocalFrameA.P
	x = float32(v17.Q.C*v18.X) - float32(v17.Q.S*v18.Y) + v17.P.X
	y = float32(v17.Q.S*v18.X) + float32(v17.Q.C*v18.Y) + v17.P.Y
	v19 = Vec2{
		X: x,
		Y: y,
	}
	goto _20
_20:
	pA = v19
	v21 = transformB
	v22 = (*b2JointSim)(unsafe.Pointer(jointSim)).LocalFrameB.P
	x = float32(v21.Q.C*v22.X) - float32(v21.Q.S*v22.Y) + v21.P.X
	y = float32(v21.Q.S*v22.X) + float32(v21.Q.C*v22.Y) + v21.P.Y
	v23 = Vec2{
		X: x,
		Y: y,
	}
	goto _24
_24:
	pB = v23
	v25 = pB
	v26 = pA
	v27 = Vec2{
		X: v25.X - v26.X,
		Y: v25.Y - v26.Y,
	}
	goto _28
_28:
	d = v27
	v29 = d
	v30 = axisA
	v31 = float32(v29.X*v30.X) + float32(v29.Y*v30.Y)
	goto _32
_32:
	translation = v31
	return translation
}

func b2PrismaticJoint_GetCurrentSpeed(tls *_Stack, jointId JointId) (r float32) {
	var axisA, d, dc, localAxisA, rA, rB, vA, vB, vRel, v30, v31, v34, v35, v37, v38, v39, v42, v43, v45, v46, v47, v50, v51, v55, v57, v58, v59, v61, v62, v63, v65, v66, v70, v71, v73, v74, v75, v78, v79, v81, v82, v83, v85, v86, v87, v90, v91, v93, v94, v97, v98 Vec2
	var base, bodyA, bodyB, bodySim, bodySimA, bodySimB, bodyStateA, bodyStateB, joint, set, set1, world2, v11, v12, v15, v17, v18, v20, v23, v24, v26, v5, v6, v9 uintptr
	var qA, qB, v29, v33, v41, v49 Rot
	var speed, wA, wB, v67, v68, v69, v77, v89, v95, v99 float32
	var v1, v13, v14, v19, v2, v22, v3, v4, v7, v8 bool
	var v53, v54 Pos
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axisA, base, bodyA, bodyB, bodySim, bodySimA, bodySimB, bodyStateA, bodyStateB, d, dc, joint, localAxisA, qA, qB, rA, rB, set, set1, speed, vA, vB, vRel, wA, wB, world2, v1, v11, v12, v13, v14, v15, v17, v18, v19, v2, v20, v22, v23, v24, v26, v29, v3, v30, v31, v33, v34, v35, v37, v38, v39, v4, v41, v42, v43, v45, v46, v47, v49, v5, v50, v51, v53, v54, v55, v57, v58, v59, v6, v61, v62, v63, v65, v66, v67, v68, v69, v7, v70, v71, v73, v74, v75, v77, v78, v79, v8, v81, v82, v83, v85, v86, v87, v89, v9, v90, v91, v93, v94, v95, v97, v98, v99
	world2 = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world2, jointId)
	if v1 = !!((*b2Joint)(unsafe.Pointer(joint)).Type1 == int32(b2_prismaticJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20348, __ccgo_ts+20222, int32FromInt32(205))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	base = b2GetJointSim(tls, world2, joint)
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_prismaticJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20381, __ccgo_ts+20222, int32FromInt32(207))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= (*b2JointSim)(unsafe.Pointer(base)).BodyIdA && (*b2JointSim)(unsafe.Pointer(base)).BodyIdA < (*b2World)(unsafe.Pointer(world2)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20413, __ccgo_ts+20222, int32FromInt32(209))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b2World)(unsafe.Pointer(world2)).Bodies.Data + uintptr((*b2JointSim)(unsafe.Pointer(base)).BodyIdA)*104
	if v2 = !!(0 <= (*b2JointSim)(unsafe.Pointer(base)).BodyIdB && (*b2JointSim)(unsafe.Pointer(base)).BodyIdB < (*b2World)(unsafe.Pointer(world2)).Bodies.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+20483, __ccgo_ts+20222, int32FromInt32(210))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b2World)(unsafe.Pointer(world2)).Bodies.Data + uintptr((*b2JointSim)(unsafe.Pointer(base)).BodyIdB)*104
	v5 = world2
	v6 = bodyA
	if v3 = !!(0 <= (*b2Body)(unsafe.Pointer(v6)).SetIndex && (*b2Body)(unsafe.Pointer(v6)).SetIndex < (*b2World)(unsafe.Pointer(v5)).SolverSets.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v5)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v6)).SetIndex)*88
	if v4 = !!(0 <= (*b2Body)(unsafe.Pointer(v6)).LocalIndex && (*b2Body)(unsafe.Pointer(v6)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v6)).LocalIndex)*96
	v9 = bodySim
	goto _10
_10:
	bodySimA = v9
	v11 = world2
	v12 = bodyB
	if v7 = !!(0 <= (*b2Body)(unsafe.Pointer(v12)).SetIndex && (*b2Body)(unsafe.Pointer(v12)).SetIndex < (*b2World)(unsafe.Pointer(v11)).SolverSets.Count); !v7 {
		b2InternalAssert(tls, __ccgo_ts+490, __ccgo_ts+698, int32FromInt32(245))
	}
	_ = v7 || bool(int32FromInt32(0) != 0)
	set = (*b2World)(unsafe.Pointer(v11)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(v12)).SetIndex)*88
	if v8 = !!(0 <= (*b2Body)(unsafe.Pointer(v12)).LocalIndex && (*b2Body)(unsafe.Pointer(v12)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Count); !v8 {
		b2InternalAssert(tls, __ccgo_ts+566, __ccgo_ts+698, int32FromInt32(246))
	}
	_ = v8 || bool(int32FromInt32(0) != 0)
	bodySim = (*b2SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr((*b2Body)(unsafe.Pointer(v12)).LocalIndex)*96
	v15 = bodySim
	goto _16
_16:
	bodySimB = v15
	v17 = world2
	v18 = bodyA
	if (*b2Body)(unsafe.Pointer(v18)).SetIndex == int32(b2_awakeSet) {
		if v13 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(v17)).SolverSets.Count); !v13 {
			b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+698, int32FromInt32(254))
		}
		_ = v13 || bool(int32FromInt32(0) != 0)
		set1 = (*b2World)(unsafe.Pointer(v17)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
		if v14 = !!(0 <= (*b2Body)(unsafe.Pointer(v18)).LocalIndex && (*b2Body)(unsafe.Pointer(v18)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set1)).BodyStates.Count); !v14 {
			b2InternalAssert(tls, __ccgo_ts+790, __ccgo_ts+698, int32FromInt32(255))
		}
		_ = v14 || bool(int32FromInt32(0) != 0)
		v20 = (*b2SolverSet)(unsafe.Pointer(set1)).BodyStates.Data + uintptr((*b2Body)(unsafe.Pointer(v18)).LocalIndex)*32
		goto _21
	}
	v20 = uintptrFromInt32(0)
	goto _21
_21:
	bodyStateA = v20
	v23 = world2
	v24 = bodyB
	if (*b2Body)(unsafe.Pointer(v24)).SetIndex == int32(b2_awakeSet) {
		if v19 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(v23)).SolverSets.Count); !v19 {
			b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+698, int32FromInt32(254))
		}
		_ = v19 || bool(int32FromInt32(0) != 0)
		set1 = (*b2World)(unsafe.Pointer(v23)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
		if v22 = !!(0 <= (*b2Body)(unsafe.Pointer(v24)).LocalIndex && (*b2Body)(unsafe.Pointer(v24)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set1)).BodyStates.Count); !v22 {
			b2InternalAssert(tls, __ccgo_ts+790, __ccgo_ts+698, int32FromInt32(255))
		}
		_ = v22 || bool(int32FromInt32(0) != 0)
		v26 = (*b2SolverSet)(unsafe.Pointer(set1)).BodyStates.Data + uintptr((*b2Body)(unsafe.Pointer(v24)).LocalIndex)*32
		goto _27
	}
	v26 = uintptrFromInt32(0)
	goto _27
_27:
	bodyStateB = v26
	qA = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	qB = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v29 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v30 = Vec2{
		X: float32FromFloat32(1),
	}
	v31 = Vec2{
		X: float32(v29.C*v30.X) - float32(v29.S*v30.Y),
		Y: float32(v29.S*v30.X) + float32(v29.C*v30.Y),
	}
	goto _32
_32:
	localAxisA = v31
	v33 = qA
	v34 = localAxisA
	v35 = Vec2{
		X: float32(v33.C*v34.X) - float32(v33.S*v34.Y),
		Y: float32(v33.S*v34.X) + float32(v33.C*v34.Y),
	}
	goto _36
_36:
	axisA = v35
	v37 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA.P
	v38 = (*b2BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v39 = Vec2{
		X: v37.X - v38.X,
		Y: v37.Y - v38.Y,
	}
	goto _40
_40:
	v41 = qA
	v42 = v39
	v43 = Vec2{
		X: float32(v41.C*v42.X) - float32(v41.S*v42.Y),
		Y: float32(v41.S*v42.X) + float32(v41.C*v42.Y),
	}
	goto _44
_44:
	rA = v43
	v45 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v46 = (*b2BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v47 = Vec2{
		X: v45.X - v46.X,
		Y: v45.Y - v46.Y,
	}
	goto _48
_48:
	v49 = qB
	v50 = v47
	v51 = Vec2{
		X: float32(v49.C*v50.X) - float32(v49.S*v50.Y),
		Y: float32(v49.S*v50.X) + float32(v49.C*v50.Y),
	}
	goto _52
_52:
	rB = v51
	v53 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Center
	v54 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Center
	v55 = Vec2{
		X: v53.X - v54.X,
		Y: v53.Y - v54.Y,
	}
	goto _56
_56:
	// Difference the centers in double so the speed stays exact far from the origin
	dc = v55
	v57 = rB
	v58 = rA
	v59 = Vec2{
		X: v57.X - v58.X,
		Y: v57.Y - v58.Y,
	}
	goto _60
_60:
	v61 = dc
	v62 = v59
	v63 = Vec2{
		X: v61.X + v62.X,
		Y: v61.Y + v62.Y,
	}
	goto _64
_64:
	d = v63
	if bodyStateA != 0 {
		v65 = (*b2BodyState)(unsafe.Pointer(bodyStateA)).LinearVelocity
	} else {
		v65 = b2Vec2_zero
	}
	vA = v65
	if bodyStateB != 0 {
		v66 = (*b2BodyState)(unsafe.Pointer(bodyStateB)).LinearVelocity
	} else {
		v66 = b2Vec2_zero
	}
	vB = v66
	if bodyStateA != 0 {
		v67 = (*b2BodyState)(unsafe.Pointer(bodyStateA)).AngularVelocity
	} else {
		v67 = float32FromFloat32(0)
	}
	wA = v67
	if bodyStateB != 0 {
		v68 = (*b2BodyState)(unsafe.Pointer(bodyStateB)).AngularVelocity
	} else {
		v68 = float32FromFloat32(0)
	}
	wB = v68
	v69 = wB
	v70 = rB
	v71 = Vec2{
		X: float32(-v69 * v70.Y),
		Y: float32(v69 * v70.X),
	}
	goto _72
_72:
	v73 = vB
	v74 = v71
	v75 = Vec2{
		X: v73.X + v74.X,
		Y: v73.Y + v74.Y,
	}
	goto _76
_76:
	v77 = wA
	v78 = rA
	v79 = Vec2{
		X: float32(-v77 * v78.Y),
		Y: float32(v77 * v78.X),
	}
	goto _80
_80:
	v81 = vA
	v82 = v79
	v83 = Vec2{
		X: v81.X + v82.X,
		Y: v81.Y + v82.Y,
	}
	goto _84
_84:
	v85 = v75
	v86 = v83
	v87 = Vec2{
		X: v85.X - v86.X,
		Y: v85.Y - v86.Y,
	}
	goto _88
_88:
	vRel = v87
	v89 = wA
	v90 = axisA
	v91 = Vec2{
		X: float32(-v89 * v90.Y),
		Y: float32(v89 * v90.X),
	}
	goto _92
_92:
	v93 = d
	v94 = v91
	v95 = float32(v93.X*v94.X) + float32(v93.Y*v94.Y)
	goto _96
_96:
	v97 = axisA
	v98 = vRel
	v99 = float32(v97.X*v98.X) + float32(v97.Y*v98.Y)
	goto _100
_100:
	speed = v95 + v99
	return speed
}

func b2GetPrismaticJointForce(tls *_Stack, world uintptr, base uintptr) (r Vec2) {
	var axialForce, inv_h, perpForce, v12, v16 float32
	var axisA, force, localAxisA, perpA, v10, v13, v14, v17, v18, v2, v20, v21, v22, v3, v6, v7, v9 Vec2
	var joint uintptr
	var qA, v1, v5 Rot
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axialForce, axisA, force, inv_h, joint, localAxisA, perpA, perpForce, qA, v1, v10, v12, v13, v14, v16, v17, v18, v2, v20, v21, v22, v3, v5, v6, v7, v9
	qA = b2GetBodyTransform(tls, world, (*b2JointSim)(unsafe.Pointer(base)).BodyIdA).Q
	joint = base + 92
	v1 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v2 = Vec2{
		X: float32FromFloat32(1),
	}
	v3 = Vec2{
		X: float32(v1.C*v2.X) - float32(v1.S*v2.Y),
		Y: float32(v1.S*v2.X) + float32(v1.C*v2.Y),
	}
	goto _4
_4:
	localAxisA = v3
	v5 = qA
	v6 = localAxisA
	v7 = Vec2{
		X: float32(v5.C*v6.X) - float32(v5.S*v6.Y),
		Y: float32(v5.S*v6.X) + float32(v5.C*v6.Y),
	}
	goto _8
_8:
	axisA = v7
	v9 = axisA
	v10 = Vec2{
		X: -v9.Y,
		Y: v9.X,
	}
	goto _11
_11:
	perpA = v10
	inv_h = (*b2World)(unsafe.Pointer(world)).Inv_h
	perpForce = float32(inv_h * (*b2PrismaticJoint)(unsafe.Pointer(joint)).Impulse.X)
	axialForce = float32(inv_h * ((*b2PrismaticJoint)(unsafe.Pointer(joint)).MotorImpulse + (*b2PrismaticJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b2PrismaticJoint)(unsafe.Pointer(joint)).UpperImpulse))
	v12 = perpForce
	v13 = perpA
	v14 = Vec2{
		X: float32(v12 * v13.X),
		Y: float32(v12 * v13.Y),
	}
	goto _15
_15:
	v16 = axialForce
	v17 = axisA
	v18 = Vec2{
		X: float32(v16 * v17.X),
		Y: float32(v16 * v17.Y),
	}
	goto _19
_19:
	v20 = v14
	v21 = v18
	v22 = Vec2{
		X: v20.X + v21.X,
		Y: v20.Y + v21.Y,
	}
	goto _23
_23:
	force = v22
	return force
}

func b2GetPrismaticJointTorque(tls *_Stack, world uintptr, base uintptr) (r float32) {
	return float32((*b2World)(unsafe.Pointer(world)).Inv_h * (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 92))).Impulse.Y)
}

// Linear constraint (point-to-line)
// d = pB - pA = xB + rB - xA - rA
// C = dot(perp, d)
// Cdot = dot(d, cross(wA, perp)) + dot(perp, vB + cross(wB, rB) - vA - cross(wA, rA))
//      = -dot(perp, vA) - dot(cross(rA + d, perp), wA) + dot(perp, vB) + dot(cross(rB, perp), vB)
// J = [-perp, -cross(rA + d, perp), perp, cross(rB, perp)]
//
// Angular constraint
// C = aB - aA + a_initial
// Cdot = wB - wA
// J = [0 0 -1 0 0 1]
//
// K = J * invM * JT
//
// J = [-a -sA a sB]
//     [0  -1  0  1]
// a = perp
// sA = cross(rA + d, a) = cross(pB - xA, a)
// sB = cross(rB, a) = cross(pB - xB, a)

// Motor/Limit linear constraint
// C = dot(axA, d)
// Cdot = -dot(axA, vA) - dot(cross(rA + d, axA), wA) + dot(axA, vB) + dot(cross(rB, axA), vB)
// J = [-axA -cross(rA + d, axA) axA cross(rB, ax1)]

// Predictive limit is applied even when the limit is not active.
// Prevents a constraint speed that can lead to a constraint error in one time step.
// Want C2 = C1 + h * Cdot >= 0
// Or:
// Cdot + C1/h >= 0
// I do not apply a negative constraint error because that is handled in position correction.
// So:
// Cdot + max(C1, 0)/h >= 0

// Block Solver
// We develop a block solver that includes the angular and linear constraints. This makes the limit stiffer.
//
// The Jacobian has 2 rows:
// J = [-uT -s1 uT s2] // linear
//     [0   -1   0  1] // angular
//
// u = perp
// s1 = cross(d + r1, u), s2 = cross(r2, u)
// a1 = cross(d + r1, v), a2 = cross(r2, v)

func b2PreparePrismaticJoint(tls *_Stack, base uintptr, context uintptr) {
	var a11, a2, a3, iA, iB, mA, mB, omega, v39, v40 float32
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world uintptr
	var idA, idB, localIndexA, localIndexB, v9 int32
	var qr, v11, v12, v13 Rot
	var v1, v2, v3, v4 bool
	var v15, v16, v17, v20, v21 Vec2
	var v35, v36 Pos
	var v41 b2Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a2, a3, bodyA, bodyB, bodySimA, bodySimB, iA, iB, idA, idB, joint, localIndexA, localIndexB, mA, mB, omega, qr, setA, setB, world, v1, v11, v12, v13, v15, v16, v17, v2, v20, v21, v3, v35, v36, v39, v4, v40, v41, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_prismaticJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20381, __ccgo_ts+20222, int32FromInt32(308))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// chase body id to the solver set where the body lives
	idA = (*b2JointSim)(unsafe.Pointer(base)).BodyIdA
	idB = (*b2JointSim)(unsafe.Pointer(base)).BodyIdB
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	if v1 = !!(0 <= idA && idA < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8507, __ccgo_ts+20222, int32FromInt32(316))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idA)*104
	if v2 = !!(0 <= idB && idB < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+8557, __ccgo_ts+20222, int32FromInt32(317))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idB)*104
	if v3 = !!((*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet)); !v3 {
		b2InternalAssert(tls, __ccgo_ts+8607, __ccgo_ts+20222, int32FromInt32(319))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(bodyA)).SetIndex && (*b2Body)(unsafe.Pointer(bodyA)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8672, __ccgo_ts+20222, int32FromInt32(320))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	setA = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(bodyA)).SetIndex)*88
	if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(bodyB)).SetIndex && (*b2Body)(unsafe.Pointer(bodyB)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+8750, __ccgo_ts+20222, int32FromInt32(321))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	setB = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(bodyB)).SetIndex)*88
	localIndexA = (*b2Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b2Body)(unsafe.Pointer(bodyB)).LocalIndex
	if v3 = !!(0 <= localIndexA && localIndexA < (*b2SolverSet)(unsafe.Pointer(setA)).BodySims.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+8828, __ccgo_ts+20222, int32FromInt32(326))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	bodySimA = (*b2SolverSet)(unsafe.Pointer(setA)).BodySims.Data + uintptr(localIndexA)*96
	if v4 = !!(0 <= localIndexB && localIndexB < (*b2SolverSet)(unsafe.Pointer(setB)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+8895, __ccgo_ts+20222, int32FromInt32(327))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySimB = (*b2SolverSet)(unsafe.Pointer(setB)).BodySims.Data + uintptr(localIndexB)*96
	mA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvMass
	iA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvInertia
	mB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvMass
	iB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvInertia
	(*b2JointSim)(unsafe.Pointer(base)).InvMassA = mA
	(*b2JointSim)(unsafe.Pointer(base)).InvMassB = mB
	(*b2JointSim)(unsafe.Pointer(base)).InvIA = iA
	(*b2JointSim)(unsafe.Pointer(base)).InvIB = iB
	joint = base + 92
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) {
		v9 = localIndexA
	} else {
		v9 = -int32(1)
	}
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).IndexA = v9
	if (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) {
		v9 = localIndexB
	} else {
		v9 = -int32(1)
	}
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).IndexB = v9
	// Compute joint anchor frames with world space rotation, relative to center of mass
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v12 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	qr.S = float32(v11.S*v12.C) + float32(v11.C*v12.S)
	qr.C = float32(v11.C*v12.C) - float32(v11.S*v12.S)
	v13 = qr
	goto _14
_14:
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).FrameA.Q = v13
	v15 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA.P
	v16 = (*b2BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v17 = Vec2{
		X: v15.X - v16.X,
		Y: v15.Y - v16.Y,
	}
	goto _18
_18:
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v20 = v17
	v21 = Vec2{
		X: float32(v11.C*v20.X) - float32(v11.S*v20.Y),
		Y: float32(v11.S*v20.X) + float32(v11.C*v20.Y),
	}
	goto _22
_22:
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).FrameA.P = v21
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v12 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameB.Q
	qr.S = float32(v11.S*v12.C) + float32(v11.C*v12.S)
	qr.C = float32(v11.C*v12.C) - float32(v11.S*v12.S)
	v13 = qr
	goto _26
_26:
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).FrameB.Q = v13
	v15 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v16 = (*b2BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v17 = Vec2{
		X: v15.X - v16.X,
		Y: v15.Y - v16.Y,
	}
	goto _30
_30:
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v20 = v17
	v21 = Vec2{
		X: float32(v11.C*v20.X) - float32(v11.S*v20.Y),
		Y: float32(v11.S*v20.X) + float32(v11.C*v20.Y),
	}
	goto _34
_34:
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).FrameB.P = v21
	// Compute the initial center delta. Incremental position updates are relative to this.
	v35 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Center
	v36 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Center
	v15 = Vec2{
		X: v35.X - v36.X,
		Y: v35.Y - v36.Y,
	}
	goto _38
_38:
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).DeltaCenter = v15
	v39 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).Hertz
	v40 = (*b2StepContext)(unsafe.Pointer(context)).H
	if v39 == float32FromFloat32(0) {
		v41 = b2Softness{}
		goto _42
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(_B2_PI)) * v39)
	a11 = float32(float32FromFloat32(2)*(*b2PrismaticJoint)(unsafe.Pointer(joint)).DampingRatio) + float32(v40*omega)
	a2 = float32(float32(v40*omega) * a11)
	a3 = float32FromFloat32(1) / (float32FromFloat32(1) + a2)
	v41 = b2Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a2 * a3),
		ImpulseScale: a3,
	}
	goto _42
_42:
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).SpringSoftness = v41
	if int32FromUint8((*b2StepContext)(unsafe.Pointer(context)).EnableWarmStarting) == false1 {
		(*b2PrismaticJoint)(unsafe.Pointer(joint)).Impulse = b2Vec2_zero
		(*b2PrismaticJoint)(unsafe.Pointer(joint)).SpringImpulse = float32FromFloat32(0)
		(*b2PrismaticJoint)(unsafe.Pointer(joint)).MotorImpulse = float32FromFloat32(0)
		(*b2PrismaticJoint)(unsafe.Pointer(joint)).LowerImpulse = float32FromFloat32(0)
		(*b2PrismaticJoint)(unsafe.Pointer(joint)).UpperImpulse = float32FromFloat32(0)
	}
}

func b2WarmStartPrismaticJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var LA, LB, a11, a21, angleImpulse, axialImpulse, iA, iB, mA, mB, perpImpulse, s11, s21, v42, v46, v57, v61, v63, v67, v76 float32
	var P, axisA, d, perpA, rA, rB, v10, v12, v13, v14, v16, v17, v18, v20, v21, v22, v24, v25, v26, v29, v30, v33, v34, v36, v37, v38, v40, v41, v44, v5, v6, v9 Vec2
	var joint, stateA, stateB, v2, v3 uintptr
	var v1 bool
	var v28, v32, v4, v8 Rot
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = LA, LB, P, a11, a21, angleImpulse, axialImpulse, axisA, d, iA, iB, joint, mA, mB, perpA, perpImpulse, rA, rB, s11, s21, stateA, stateB, v1, v10, v12, v13, v14, v16, v17, v18, v2, v20, v21, v22, v24, v25, v26, v28, v29, v3, v30, v32, v33, v34, v36, v37, v38, v4, v40, v41, v42, v44, v46, v5, v57, v6, v61, v63, v67, v76, v8, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_prismaticJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20381, __ccgo_ts+20222, int32FromInt32(366))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Add(nil, bp)) = b2_identityBodyState13
	joint = base + 92
	if (*b2PrismaticJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2PrismaticJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v2
	if (*b2PrismaticJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2PrismaticJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v3
	v4 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).FrameA.P
	v6 = Vec2{
		X: float32(v4.C*v5.X) - float32(v4.S*v5.Y),
		Y: float32(v4.S*v5.X) + float32(v4.C*v5.Y),
	}
	goto _7
_7:
	rA = v6
	v8 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v9 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).FrameB.P
	v10 = Vec2{
		X: float32(v8.C*v9.X) - float32(v8.S*v9.Y),
		Y: float32(v8.S*v9.X) + float32(v8.C*v9.Y),
	}
	goto _11
_11:
	rB = v10
	v12 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
	v13 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaPosition
	v14 = Vec2{
		X: v12.X - v13.X,
		Y: v12.Y - v13.Y,
	}
	goto _15
_15:
	v16 = v14
	v17 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).DeltaCenter
	v18 = Vec2{
		X: v16.X + v17.X,
		Y: v16.Y + v17.Y,
	}
	goto _19
_19:
	v20 = rB
	v21 = rA
	v22 = Vec2{
		X: v20.X - v21.X,
		Y: v20.Y - v21.Y,
	}
	goto _23
_23:
	v24 = v18
	v25 = v22
	v26 = Vec2{
		X: v24.X + v25.X,
		Y: v24.Y + v25.Y,
	}
	goto _27
_27:
	d = v26
	v28 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).FrameA.Q
	v29 = Vec2{
		X: float32FromFloat32(1),
	}
	v30 = Vec2{
		X: float32(v28.C*v29.X) - float32(v28.S*v29.Y),
		Y: float32(v28.S*v29.X) + float32(v28.C*v29.Y),
	}
	goto _31
_31:
	axisA = v30
	v32 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v33 = axisA
	v34 = Vec2{
		X: float32(v32.C*v33.X) - float32(v32.S*v33.Y),
		Y: float32(v32.S*v33.X) + float32(v32.C*v33.Y),
	}
	goto _35
_35:
	axisA = v34
	v5 = rA
	v6 = d
	v9 = Vec2{
		X: v5.X + v6.X,
		Y: v5.Y + v6.Y,
	}
	goto _39
_39:
	v10 = v9
	v12 = axisA
	v42 = float32(v10.X*v12.Y) - float32(v10.Y*v12.X)
	goto _43
_43:
	// impulse is applied at anchor point on body B
	a11 = v42
	v13 = rB
	v14 = axisA
	v46 = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
	goto _47
_47:
	a21 = v46
	axialImpulse = (*b2PrismaticJoint)(unsafe.Pointer(joint)).SpringImpulse + (*b2PrismaticJoint)(unsafe.Pointer(joint)).MotorImpulse + (*b2PrismaticJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b2PrismaticJoint)(unsafe.Pointer(joint)).UpperImpulse
	v16 = axisA
	v17 = Vec2{
		X: -v16.Y,
		Y: v16.X,
	}
	goto _50
_50:
	// perpendicular constraint
	perpA = v17
	v18 = rA
	v20 = d
	v21 = Vec2{
		X: v18.X + v20.X,
		Y: v18.Y + v20.Y,
	}
	goto _54
_54:
	v22 = v21
	v24 = perpA
	v57 = float32(v22.X*v24.Y) - float32(v22.Y*v24.X)
	goto _58
_58:
	s11 = v57
	v25 = rB
	v26 = perpA
	v61 = float32(v25.X*v26.Y) - float32(v25.Y*v26.X)
	goto _62
_62:
	s21 = v61
	perpImpulse = (*b2PrismaticJoint)(unsafe.Pointer(joint)).Impulse.X
	angleImpulse = (*b2PrismaticJoint)(unsafe.Pointer(joint)).Impulse.Y
	v63 = axialImpulse
	v29 = axisA
	v30 = Vec2{
		X: float32(v63 * v29.X),
		Y: float32(v63 * v29.Y),
	}
	goto _66
_66:
	v67 = perpImpulse
	v33 = perpA
	v34 = Vec2{
		X: float32(v67 * v33.X),
		Y: float32(v67 * v33.Y),
	}
	goto _70
_70:
	v36 = v30
	v37 = v34
	v38 = Vec2{
		X: v36.X + v37.X,
		Y: v36.Y + v37.Y,
	}
	goto _74
_74:
	P = v38
	LA = float32(axialImpulse*a11) + float32(perpImpulse*s11) + angleImpulse
	LB = float32(axialImpulse*a21) + float32(perpImpulse*s21) + angleImpulse
	if (*b2BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b2_dynamicFlag) != 0 {
		v40 = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
		v76 = mA
		v41 = P
		v44 = Vec2{
			X: v40.X - float32(v76*v41.X),
			Y: v40.Y - float32(v76*v41.Y),
		}
		goto _79
	_79:
		(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = v44
		*(*float32)(unsafe.Add(nil, stateA+8)) -= float32(iA * LA)
	}
	if (*b2BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b2_dynamicFlag) != 0 {
		v5 = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
		v42 = mB
		v6 = P
		v9 = Vec2{
			X: v5.X + float32(v42*v6.X),
			Y: v5.Y + float32(v42*v6.Y),
		}
		goto _84
	_84:
		(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = v9
		*(*float32)(unsafe.Add(nil, stateB+8)) += float32(iB * LB)
	}
}

func b2SolvePrismaticJoint(tls *_Stack, base uintptr, context uintptr, useBias uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var C, C1, C2, Cdot, Cdot1, Cdot2, Cdot3, LA, LA1, LA2, LA3, LA4, LB, LB1, LB2, LB3, LB4, a11, a111, a12, a21, a211, a22, axialMass, bias, bias1, bias2, deltaImpulse, deltaImpulse1, deltaImpulse2, det, iA, iB, impulse, impulseScale, impulseScale1, impulseScale2, impulseScale3, k, k11, k12, k22, mA, mB, massScale, massScale1, massScale2, massScale3, maxImpulse, oldImpulse, oldImpulse1, oldImpulse2, s11, s21, safe, safe1, speculativeDistance, translation, wA, wB, v50, v58, v62, v64, v71, v73, v78 float32
	var C3, Cdot4, P, P1, P2, P3, P4, axisA, b10, bias3, d, deltaImpulse3, perpA, rA, rB, vA, vB, x, v17, v18, v21, v22, v24, v25, v26, v28, v29, v30, v32, v33, v34, v36, v37, v38, v41, v42, v45, v46 Vec2
	var K, v218 Mat22
	var joint, stateA, stateB, v2, v3 uintptr
	var qA, qB, qr, r1, relQ, v10, v12, v13, v14, v16, v20, v4, v40, v44, v5, v6, v8, v9 Rot
	var softness b2Softness
	var v1 bool
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, C1, C2, C3, Cdot, Cdot1, Cdot2, Cdot3, Cdot4, K, LA, LA1, LA2, LA3, LA4, LB, LB1, LB2, LB3, LB4, P, P1, P2, P3, P4, a11, a111, a12, a21, a211, a22, axialMass, axisA, b10, bias, bias1, bias2, bias3, d, deltaImpulse, deltaImpulse1, deltaImpulse2, deltaImpulse3, det, iA, iB, impulse, impulseScale, impulseScale1, impulseScale2, impulseScale3, joint, k, k11, k12, k22, mA, mB, massScale, massScale1, massScale2, massScale3, maxImpulse, oldImpulse, oldImpulse1, oldImpulse2, perpA, qA, qB, qr, r1, rA, rB, relQ, s11, s21, safe, safe1, softness, speculativeDistance, stateA, stateB, translation, vA, vB, wA, wB, x, v1, v10, v12, v13, v14, v16, v17, v18, v2, v20, v21, v218, v22, v24, v25, v26, v28, v29, v3, v30, v32, v33, v34, v36, v37, v38, v4, v40, v41, v42, v44, v45, v46, v5, v50, v58, v6, v62, v64, v71, v73, v78, v8, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_prismaticJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20381, __ccgo_ts+20222, int32FromInt32(420))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Add(nil, bp)) = b2_identityBodyState13
	joint = base + 92
	if (*b2PrismaticJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2PrismaticJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v2
	if (*b2PrismaticJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2PrismaticJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v3
	vA = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	v4 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).FrameA.Q
	qr.S = float32(v4.S*v5.C) + float32(v4.C*v5.S)
	qr.C = float32(v4.C*v5.C) - float32(v4.S*v5.S)
	v6 = qr
	goto _7
_7:
	qA = v6
	v8 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v9 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).FrameB.Q
	qr.S = float32(v8.S*v9.C) + float32(v8.C*v9.S)
	qr.C = float32(v8.C*v9.C) - float32(v8.S*v9.S)
	v10 = qr
	goto _11
_11:
	qB = v10
	v12 = qA
	v13 = qB
	r1.S = float32(v12.C*v13.S) - float32(v12.S*v13.C)
	r1.C = float32(v12.C*v13.C) + float32(v12.S*v13.S)
	v14 = r1
	goto _15
_15:
	relQ = v14
	v16 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v17 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).FrameA.P
	v18 = Vec2{
		X: float32(v16.C*v17.X) - float32(v16.S*v17.Y),
		Y: float32(v16.S*v17.X) + float32(v16.C*v17.Y),
	}
	goto _19
_19:
	// current anchors
	rA = v18
	v20 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v21 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).FrameB.P
	v22 = Vec2{
		X: float32(v20.C*v21.X) - float32(v20.S*v21.Y),
		Y: float32(v20.S*v21.X) + float32(v20.C*v21.Y),
	}
	goto _23
_23:
	rB = v22
	v24 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
	v25 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaPosition
	v26 = Vec2{
		X: v24.X - v25.X,
		Y: v24.Y - v25.Y,
	}
	goto _27
_27:
	v28 = v26
	v29 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).DeltaCenter
	v30 = Vec2{
		X: v28.X + v29.X,
		Y: v28.Y + v29.Y,
	}
	goto _31
_31:
	v32 = rB
	v33 = rA
	v34 = Vec2{
		X: v32.X - v33.X,
		Y: v32.Y - v33.Y,
	}
	goto _35
_35:
	v36 = v30
	v37 = v34
	v38 = Vec2{
		X: v36.X + v37.X,
		Y: v36.Y + v37.Y,
	}
	goto _39
_39:
	d = v38
	v40 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).FrameA.Q
	v41 = Vec2{
		X: float32FromFloat32(1),
	}
	v42 = Vec2{
		X: float32(v40.C*v41.X) - float32(v40.S*v41.Y),
		Y: float32(v40.S*v41.X) + float32(v40.C*v41.Y),
	}
	goto _43
_43:
	axisA = v42
	v44 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v45 = axisA
	v46 = Vec2{
		X: float32(v44.C*v45.X) - float32(v44.S*v45.Y),
		Y: float32(v44.S*v45.X) + float32(v44.C*v45.Y),
	}
	goto _47
_47:
	axisA = v46
	v17 = axisA
	v18 = d
	v50 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y)
	goto _51
_51:
	translation = v50
	v21 = rA
	v22 = d
	v24 = Vec2{
		X: v21.X + v22.X,
		Y: v21.Y + v22.Y,
	}
	goto _55
_55:
	v25 = v24
	v26 = axisA
	v58 = float32(v25.X*v26.Y) - float32(v25.Y*v26.X)
	goto _59
_59:
	// These scalars are for torques generated by axial forces
	a11 = v58
	v28 = rB
	v29 = axisA
	v62 = float32(v28.X*v29.Y) - float32(v28.Y*v29.X)
	goto _63
_63:
	a21 = v62
	k = mA + mB + float32(float32(iA*a11)*a11) + float32(float32(iB*a21)*a21)
	if k > float32FromFloat32(0) {
		v64 = float32FromFloat32(1) / k
	} else {
		v64 = float32FromFloat32(0)
	}
	axialMass = v64
	softness = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness
	// spring constraint
	if (*b2PrismaticJoint)(unsafe.Pointer(joint)).EnableSpring != 0 {
		// This is a real spring and should be applied even during relax
		C = translation - (*b2PrismaticJoint)(unsafe.Pointer(joint)).TargetTranslation
		bias = float32((*b2PrismaticJoint)(unsafe.Pointer(joint)).SpringSoftness.BiasRate * C)
		massScale = (*b2PrismaticJoint)(unsafe.Pointer(joint)).SpringSoftness.MassScale
		impulseScale = (*b2PrismaticJoint)(unsafe.Pointer(joint)).SpringSoftness.ImpulseScale
		v30 = vB
		v32 = vA
		v33 = Vec2{
			X: v30.X - v32.X,
			Y: v30.Y - v32.Y,
		}
		goto _68
	_68:
		v34 = axisA
		v36 = v33
		v71 = float32(v34.X*v36.X) + float32(v34.Y*v36.Y)
		goto _72
	_72:
		Cdot = v71 + float32(a21*wB) - float32(a11*wA)
		deltaImpulse = float32(float32(-massScale*axialMass)*(Cdot+bias)) - float32(impulseScale*(*b2PrismaticJoint)(unsafe.Pointer(joint)).SpringImpulse)
		*(*float32)(unsafe.Add(nil, joint+8)) += deltaImpulse
		v50 = deltaImpulse
		v17 = axisA
		v18 = Vec2{
			X: float32(v50 * v17.X),
			Y: float32(v50 * v17.Y),
		}
		goto _76
	_76:
		P = v18
		LA = float32(deltaImpulse * a11)
		LB = float32(deltaImpulse * a21)
		v21 = vA
		v58 = mA
		v22 = P
		v24 = Vec2{
			X: v21.X - float32(v58*v22.X),
			Y: v21.Y - float32(v58*v22.Y),
		}
		goto _81
	_81:
		vA = v24
		wA = wA - float32(iA*LA)
		v17 = vB
		v50 = mB
		v18 = P
		v21 = Vec2{
			X: v17.X + float32(v50*v18.X),
			Y: v17.Y + float32(v50*v18.Y),
		}
		goto _86
	_86:
		vB = v21
		wB = wB + float32(iB*LB)
	}
	// Solve motor constraint
	if (*b2PrismaticJoint)(unsafe.Pointer(joint)).EnableMotor != 0 {
		v17 = vB
		v18 = vA
		v21 = Vec2{
			X: v17.X - v18.X,
			Y: v17.Y - v18.Y,
		}
		goto _90
	_90:
		v22 = axisA
		v24 = v21
		v50 = float32(v22.X*v24.X) + float32(v22.Y*v24.Y)
		goto _94
	_94:
		Cdot1 = v50 + float32(a21*wB) - float32(a11*wA)
		impulse = float32(axialMass * ((*b2PrismaticJoint)(unsafe.Pointer(joint)).MotorSpeed - Cdot1))
		oldImpulse = (*b2PrismaticJoint)(unsafe.Pointer(joint)).MotorImpulse
		maxImpulse = float32((*b2StepContext)(unsafe.Pointer(context)).H * (*b2PrismaticJoint)(unsafe.Pointer(joint)).MaxMotorForce)
		v58 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).MotorImpulse + impulse
		v62 = -maxImpulse
		v64 = maxImpulse
		if v58 < v62 {
			v73 = v62
		} else {
			if v58 > v64 {
				v78 = v64
			} else {
				v78 = v58
			}
			v73 = v78
		}
		v71 = v73
		goto _99
	_99:
		(*b2PrismaticJoint)(unsafe.Pointer(joint)).MotorImpulse = v71
		impulse = (*b2PrismaticJoint)(unsafe.Pointer(joint)).MotorImpulse - oldImpulse
		v50 = impulse
		v17 = axisA
		v18 = Vec2{
			X: float32(v50 * v17.X),
			Y: float32(v50 * v17.Y),
		}
		goto _105
	_105:
		P1 = v18
		LA1 = float32(impulse * a11)
		LB1 = float32(impulse * a21)
		v21 = vA
		v58 = mA
		v22 = P1
		v24 = Vec2{
			X: v21.X - float32(v58*v22.X),
			Y: v21.Y - float32(v58*v22.Y),
		}
		goto _110
	_110:
		vA = v24
		wA = wA - float32(iA*LA1)
		v17 = vB
		v50 = mB
		v18 = P1
		v21 = Vec2{
			X: v17.X + float32(v50*v18.X),
			Y: v17.Y + float32(v50*v18.Y),
		}
		goto _115
	_115:
		vB = v21
		wB = wB + float32(iB*LB1)
	}
	if (*b2PrismaticJoint)(unsafe.Pointer(joint)).EnableLimit != 0 {
		// Clamp the speculative distance to a reasonable value
		speculativeDistance = float32(float32FromFloat32(0.25) * ((*b2PrismaticJoint)(unsafe.Pointer(joint)).UpperTranslation - (*b2PrismaticJoint)(unsafe.Pointer(joint)).LowerTranslation))
		// Lower limit
		C1 = translation - (*b2PrismaticJoint)(unsafe.Pointer(joint)).LowerTranslation
		if C1 < speculativeDistance {
			bias1 = float32FromFloat32(0)
			massScale1 = float32FromFloat32(1)
			impulseScale1 = float32FromFloat32(0)
			if C1 > float32FromFloat32(0) {
				// speculation
				safe = b2GetLengthUnitsPerMeter(tls)
				v50 = C1
				v58 = safe
				if v50 < v58 {
					v64 = v50
				} else {
					v64 = v58
				}
				v62 = v64
				goto _119
			_119:
				bias1 = float32(v62 * (*b2StepContext)(unsafe.Pointer(context)).Inv_h)
			} else {
				if useBias != 0 {
					bias1 = float32(softness.BiasRate * C1)
					massScale1 = softness.MassScale
					impulseScale1 = softness.ImpulseScale
				}
			}
			oldImpulse1 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).LowerImpulse
			v17 = vB
			v18 = vA
			v21 = Vec2{
				X: v17.X - v18.X,
				Y: v17.Y - v18.Y,
			}
			goto _124
		_124:
			v22 = axisA
			v24 = v21
			v50 = float32(v22.X*v24.X) + float32(v22.Y*v24.Y)
			goto _128
		_128:
			Cdot2 = v50 + float32(a21*wB) - float32(a11*wA)
			deltaImpulse1 = float32(float32(-axialMass*massScale1)*(Cdot2+bias1)) - float32(impulseScale1*oldImpulse1)
			v58 = oldImpulse1 + deltaImpulse1
			v62 = float32FromFloat32(0)
			if v58 > v62 {
				v71 = v58
			} else {
				v71 = v62
			}
			v64 = v71
			goto _132
		_132:
			(*b2PrismaticJoint)(unsafe.Pointer(joint)).LowerImpulse = v64
			deltaImpulse1 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).LowerImpulse - oldImpulse1
			v50 = deltaImpulse1
			v17 = axisA
			v18 = Vec2{
				X: float32(v50 * v17.X),
				Y: float32(v50 * v17.Y),
			}
			goto _137
		_137:
			P2 = v18
			LA2 = float32(deltaImpulse1 * a11)
			LB2 = float32(deltaImpulse1 * a21)
			v21 = vA
			v58 = mA
			v22 = P2
			v24 = Vec2{
				X: v21.X - float32(v58*v22.X),
				Y: v21.Y - float32(v58*v22.Y),
			}
			goto _142
		_142:
			vA = v24
			wA = wA - float32(iA*LA2)
			v17 = vB
			v50 = mB
			v18 = P2
			v21 = Vec2{
				X: v17.X + float32(v50*v18.X),
				Y: v17.Y + float32(v50*v18.Y),
			}
			goto _147
		_147:
			vB = v21
			wB = wB + float32(iB*LB2)
		} else {
			(*b2PrismaticJoint)(unsafe.Pointer(joint)).LowerImpulse = float32FromFloat32(0)
		}
		// Upper limit
		// Note: signs are flipped to keep C positive when the constraint is satisfied.
		// This also keeps the impulse positive when the limit is active.
		// sign flipped
		C2 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).UpperTranslation - translation
		if C2 < speculativeDistance {
			bias2 = float32FromFloat32(0)
			massScale2 = float32FromFloat32(1)
			impulseScale2 = float32FromFloat32(0)
			if C2 > float32FromFloat32(0) {
				// speculation
				safe1 = b2GetLengthUnitsPerMeter(tls)
				v50 = C2
				v58 = safe1
				if v50 < v58 {
					v64 = v50
				} else {
					v64 = v58
				}
				v62 = v64
				goto _151
			_151:
				bias2 = float32(v62 * (*b2StepContext)(unsafe.Pointer(context)).Inv_h)
			} else {
				if useBias != 0 {
					bias2 = float32(softness.BiasRate * C2)
					massScale2 = softness.MassScale
					impulseScale2 = softness.ImpulseScale
				}
			}
			oldImpulse2 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).UpperImpulse
			v17 = vA
			v18 = vB
			v21 = Vec2{
				X: v17.X - v18.X,
				Y: v17.Y - v18.Y,
			}
			goto _156
		_156:
			v22 = axisA
			v24 = v21
			v50 = float32(v22.X*v24.X) + float32(v22.Y*v24.Y)
			goto _160
		_160:
			// sign flipped
			Cdot3 = v50 + float32(a11*wA) - float32(a21*wB)
			deltaImpulse2 = float32(float32(-axialMass*massScale2)*(Cdot3+bias2)) - float32(impulseScale2*oldImpulse2)
			v58 = oldImpulse2 + deltaImpulse2
			v62 = float32FromFloat32(0)
			if v58 > v62 {
				v71 = v58
			} else {
				v71 = v62
			}
			v64 = v71
			goto _164
		_164:
			(*b2PrismaticJoint)(unsafe.Pointer(joint)).UpperImpulse = v64
			deltaImpulse2 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).UpperImpulse - oldImpulse2
			v50 = deltaImpulse2
			v17 = axisA
			v18 = Vec2{
				X: float32(v50 * v17.X),
				Y: float32(v50 * v17.Y),
			}
			goto _169
		_169:
			P3 = v18
			LA3 = float32(deltaImpulse2 * a11)
			LB3 = float32(deltaImpulse2 * a21)
			// sign flipped
			v21 = vA
			v58 = mA
			v22 = P3
			v24 = Vec2{
				X: v21.X + float32(v58*v22.X),
				Y: v21.Y + float32(v58*v22.Y),
			}
			goto _174
		_174:
			vA = v24
			wA = wA + float32(iA*LA3)
			v17 = vB
			v50 = mB
			v18 = P3
			v21 = Vec2{
				X: v17.X - float32(v50*v18.X),
				Y: v17.Y - float32(v50*v18.Y),
			}
			goto _179
		_179:
			vB = v21
			wB = wB - float32(iB*LB3)
		} else {
			(*b2PrismaticJoint)(unsafe.Pointer(joint)).UpperImpulse = float32FromFloat32(0)
		}
	}
	// Solve the prismatic constraint in block form
	v17 = axisA
	v18 = Vec2{
		X: -v17.Y,
		Y: v17.X,
	}
	goto _182
_182:
	perpA = v18
	v21 = d
	v22 = rA
	v24 = Vec2{
		X: v21.X + v22.X,
		Y: v21.Y + v22.Y,
	}
	goto _186
_186:
	v25 = v24
	v26 = perpA
	v50 = float32(v25.X*v26.Y) - float32(v25.Y*v26.X)
	goto _190
_190:
	// These scalars are for torques generated by the perpendicular constraint force
	s11 = v50
	v28 = rB
	v29 = perpA
	v58 = float32(v28.X*v29.Y) - float32(v28.Y*v29.X)
	goto _194
_194:
	s21 = v58
	v30 = vB
	v32 = vA
	v33 = Vec2{
		X: v30.X - v32.X,
		Y: v30.Y - v32.Y,
	}
	goto _198
_198:
	v34 = perpA
	v36 = v33
	v62 = float32(v34.X*v36.X) + float32(v34.Y*v36.Y)
	goto _202
_202:
	Cdot4.X = v62 + float32(s21*wB) - float32(s11*wA)
	Cdot4.Y = wB - wA
	bias3 = b2Vec2_zero
	massScale3 = float32FromFloat32(1)
	impulseScale3 = float32FromFloat32(0)
	if useBias != 0 {
		v17 = perpA
		v18 = d
		v50 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y)
		goto _206
	_206:
		C3.X = v50
		v4 = relQ
		v50 = b2Atan2(tls, v4.S, v4.C)
		goto _209
	_209:
		C3.Y = v50
		v50 = softness.BiasRate
		v17 = C3
		v18 = Vec2{
			X: float32(v50 * v17.X),
			Y: float32(v50 * v17.Y),
		}
		goto _213
	_213:
		bias3 = v18
		massScale3 = softness.MassScale
		impulseScale3 = softness.ImpulseScale
	}
	k11 = mA + mB + float32(float32(iA*s11)*s11) + float32(float32(iB*s21)*s21)
	k12 = float32(iA*s11) + float32(iB*s21)
	k22 = iA + iB
	if k22 == float32FromFloat32(0) {
		// For bodies with fixed rotation.
		k22 = float32FromFloat32(1)
	}
	K = Mat22{
		Cx: Vec2{
			X: k11,
			Y: k12,
		},
		Cy: Vec2{
			X: k12,
			Y: k22,
		},
	}
	v17 = Cdot4
	v18 = bias3
	v21 = Vec2{
		X: v17.X + v18.X,
		Y: v17.Y + v18.Y,
	}
	goto _217
_217:
	v218 = K
	v22 = v21
	a111 = v218.Cx.X
	a12 = v218.Cy.X
	a211 = v218.Cx.Y
	a22 = v218.Cy.Y
	det = float32(a111*a22) - float32(a12*a211)
	if det != float32FromFloat32(0) {
		det = float32FromFloat32(1) / det
	}
	x = Vec2{
		X: float32(det * (float32(a22*v22.X) - float32(a12*v22.Y))),
		Y: float32(det * (float32(a111*v22.Y) - float32(a211*v22.X))),
	}
	v24 = x
	goto _221
_221:
	b10 = v24
	deltaImpulse3.X = float32(-massScale3*b10.X) - float32(impulseScale3*(*b2PrismaticJoint)(unsafe.Pointer(joint)).Impulse.X)
	deltaImpulse3.Y = float32(-massScale3*b10.Y) - float32(impulseScale3*(*b2PrismaticJoint)(unsafe.Pointer(joint)).Impulse.Y)
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).Impulse.X += deltaImpulse3.X
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).Impulse.Y += deltaImpulse3.Y
	v50 = deltaImpulse3.X
	v17 = perpA
	v18 = Vec2{
		X: float32(v50 * v17.X),
		Y: float32(v50 * v17.Y),
	}
	goto _225
_225:
	P4 = v18
	LA4 = float32(deltaImpulse3.X*s11) + deltaImpulse3.Y
	LB4 = float32(deltaImpulse3.X*s21) + deltaImpulse3.Y
	v21 = vA
	v58 = mA
	v22 = P4
	v24 = Vec2{
		X: v21.X - float32(v58*v22.X),
		Y: v21.Y - float32(v58*v22.Y),
	}
	goto _230
_230:
	vA = v24
	wA = wA - float32(iA*LA4)
	v17 = vB
	v50 = mB
	v18 = P4
	v21 = Vec2{
		X: v17.X + float32(v50*v18.X),
		Y: v17.Y + float32(v50*v18.Y),
	}
	goto _235
_235:
	vB = v21
	wB = wB + float32(iB*LB4)
	if v1 = !!(b2IsValidVec2(tls, vA) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20553, __ccgo_ts+20222, int32FromInt32(661))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b2IsValidFloat(tls, wA) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20573, __ccgo_ts+20222, int32FromInt32(662))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b2IsValidVec2(tls, vB) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20594, __ccgo_ts+20222, int32FromInt32(663))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b2IsValidFloat(tls, wB) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20614, __ccgo_ts+20222, int32FromInt32(664))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b2_dynamicFlag) != 0 {
		(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
		(*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	}
	if (*b2BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b2_dynamicFlag) != 0 {
		(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
		(*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
	}
}

func b2DrawPrismaticJoint(tls *_Stack, draw uintptr, base uintptr, transformA WorldTransform, transformB WorldTransform, drawScale float32) {
	var C, frameA, frameB, v16, v18, v2, v32 WorldTransform
	var axisA, perp, v10, v13, v25, v26, v29, v35, v36, v9 Vec2
	var b, v38, v46 float32
	var joint uintptr
	var lower, p1, upper, v12, v14, v28, v30 Pos
	var qr, v20, v21, v22, v24, v34, v4, v5, v6, v8 Rot
	var v1 bool
	var v19, v3 Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, axisA, b, frameA, frameB, joint, lower, p1, perp, qr, upper, v1, v10, v12, v13, v14, v16, v18, v19, v2, v20, v21, v22, v24, v25, v26, v28, v29, v3, v30, v32, v34, v35, v36, v38, v4, v46, v5, v6, v8, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_prismaticJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20381, __ccgo_ts+20222, int32FromInt32(682))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	joint = base + 92
	v2 = transformA
	v3 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA
	v4 = v2.Q
	v5 = v3.Q
	qr.S = float32(v4.S*v5.C) + float32(v4.C*v5.S)
	qr.C = float32(v4.C*v5.C) - float32(v4.S*v5.S)
	v6 = qr
	goto _7
_7:
	C.Q = v6
	v8 = v2.Q
	v9 = v3.P
	v10 = Vec2{
		X: float32(v8.C*v9.X) - float32(v8.S*v9.Y),
		Y: float32(v8.S*v9.X) + float32(v8.C*v9.Y),
	}
	goto _11
_11:
	v12 = v2.P
	v13 = v10
	v14 = Pos{
		X: v12.X + v13.X,
		Y: v12.Y + v13.Y,
	}
	goto _15
_15:
	C.P = v14
	v16 = C
	goto _17
_17:
	frameA = v16
	v18 = transformB
	v19 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameB
	v20 = v18.Q
	v21 = v19.Q
	qr.S = float32(v20.S*v21.C) + float32(v20.C*v21.S)
	qr.C = float32(v20.C*v21.C) - float32(v20.S*v21.S)
	v22 = qr
	goto _23
_23:
	C.Q = v22
	v24 = v18.Q
	v25 = v19.P
	v26 = Vec2{
		X: float32(v24.C*v25.X) - float32(v24.S*v25.Y),
		Y: float32(v24.S*v25.X) + float32(v24.C*v25.Y),
	}
	goto _27
_27:
	v28 = v18.P
	v29 = v26
	v30 = Pos{
		X: v28.X + v29.X,
		Y: v28.Y + v29.Y,
	}
	goto _31
_31:
	C.P = v30
	v32 = C
	goto _33
_33:
	frameB = v32
	v34 = frameA.Q
	v35 = Vec2{
		X: float32FromFloat32(1),
	}
	v36 = Vec2{
		X: float32(v34.C*v35.X) - float32(v34.S*v35.Y),
		Y: float32(v34.S*v35.X) + float32(v34.C*v35.Y),
	}
	goto _37
_37:
	axisA = v36
	(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, frameA.P, frameB.P, int32(b2_colorDimGray), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	if (*b2PrismaticJoint)(unsafe.Pointer(joint)).EnableLimit != 0 {
		b = float32(float32FromFloat32(0.25) * drawScale)
		v38 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).LowerTranslation
		v9 = axisA
		v10 = Vec2{
			X: float32(v38 * v9.X),
			Y: float32(v38 * v9.Y),
		}
		goto _41
	_41:
		v12 = frameA.P
		v13 = v10
		v14 = Pos{
			X: v12.X + v13.X,
			Y: v12.Y + v13.Y,
		}
		goto _45
	_45:
		lower = v14
		v46 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).UpperTranslation
		v25 = axisA
		v26 = Vec2{
			X: float32(v46 * v25.X),
			Y: float32(v46 * v25.Y),
		}
		goto _49
	_49:
		v28 = frameA.P
		v29 = v26
		v30 = Pos{
			X: v28.X + v29.X,
			Y: v28.Y + v29.Y,
		}
		goto _53
	_53:
		upper = v30
		v35 = axisA
		v36 = Vec2{
			X: -v35.Y,
			Y: v35.X,
		}
		goto _56
	_56:
		perp = v36
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, lower, upper, int32(b2_colorGray), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		v38 = -b
		v9 = perp
		v10 = Vec2{
			X: float32(v38 * v9.X),
			Y: float32(v38 * v9.Y),
		}
		goto _60
	_60:
		v12 = lower
		v13 = v10
		v14 = Pos{
			X: v12.X + v13.X,
			Y: v12.Y + v13.Y,
		}
		goto _64
	_64:
		v46 = b
		v25 = perp
		v26 = Vec2{
			X: float32(v46 * v25.X),
			Y: float32(v46 * v25.Y),
		}
		goto _68
	_68:
		v28 = lower
		v29 = v26
		v30 = Pos{
			X: v28.X + v29.X,
			Y: v28.Y + v29.Y,
		}
		goto _72
	_72:
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, v14, v30, int32(b2_colorGreen), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		v38 = -b
		v9 = perp
		v10 = Vec2{
			X: float32(v38 * v9.X),
			Y: float32(v38 * v9.Y),
		}
		goto _76
	_76:
		v12 = upper
		v13 = v10
		v14 = Pos{
			X: v12.X + v13.X,
			Y: v12.Y + v13.Y,
		}
		goto _80
	_80:
		v46 = b
		v25 = perp
		v26 = Vec2{
			X: float32(v46 * v25.X),
			Y: float32(v46 * v25.Y),
		}
		goto _84
	_84:
		v28 = upper
		v29 = v26
		v30 = Pos{
			X: v28.X + v29.X,
			Y: v28.Y + v29.Y,
		}
		goto _88
	_88:
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, v14, v30, int32(b2_colorRed), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	} else {
		v9 = axisA
		v10 = Vec2{
			X: -v9.X,
			Y: -v9.Y,
		}
		goto _91
	_91:
		v12 = frameA.P
		v13 = v10
		v14 = Pos{
			X: v12.X + v13.X,
			Y: v12.Y + v13.Y,
		}
		goto _95
	_95:
		v28 = frameA.P
		v25 = axisA
		v30 = Pos{
			X: v28.X + v25.X,
			Y: v28.Y + v25.Y,
		}
		goto _99
	_99:
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, v14, v30, int32(b2_colorGray), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	}
	if (*b2PrismaticJoint)(unsafe.Pointer(joint)).EnableSpring != 0 {
		v38 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).TargetTranslation
		v9 = axisA
		v10 = Vec2{
			X: float32(v38 * v9.X),
			Y: float32(v38 * v9.Y),
		}
		goto _103
	_103:
		v12 = frameA.P
		v13 = v10
		v14 = Pos{
			X: v12.X + v13.X,
			Y: v12.Y + v13.Y,
		}
		goto _107
	_107:
		p1 = v14
		(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p1, float32FromFloat32(8), int32(b2_colorViolet), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	}
	(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, frameA.P, float32FromFloat32(5), int32(b2_colorGray), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, frameB.P, float32FromFloat32(5), int32(b2_colorBlue), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
}

func b2RecMakeJointId(tls *_Stack, rdr uintptr, recorded JointId) (r JointId) {
	var id JointId
	_ = id
	id.Index1 = recorded.Index1
	id.World0 = uint16(uint32((*b2RecReader)(unsafe.Pointer(rdr)).ReplayWorldId.Index1) - uint32FromUint32(1))
	id.Generation = recorded.Generation
	return id
}

func b2RecCheckJointId(tls *_Stack, rdr uintptr, got JointId, rec JointId) {
	b2RecCheckId(tls, rdr, __ccgo_ts+20845, got.Index1, uint32(got.Generation), rec.Index1, uint32(rec.Generation))
}

// Point-to-point constraint
// C = pB - pA
// Cdot = vB - vA
//      = vB + cross(wB, rB) - vA - cross(wA, rA)
// J = [-E -skew(rA) E skew(rB) ]

// Identity used:
// w k % (rx i + ry j) = w * (-ry i + rx j)

// Motor constraint
// Cdot = wB - wA
// J = [0 0 -1 0 0 1]
// K = invIA + invIB

func b2RevoluteJoint_EnableSpring(tls *_Stack, jointId JointId, enableSpring uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* _a at bp+0 */ b2RecArgs_RevoluteJointEnableSpring
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_RevoluteJointEnableSpring)(unsafe.Add(nil, bp)) = b2RecArgs_RevoluteJointEnableSpring{
			Joint:        jointId,
			EnableSpring: enableSpring,
		}
		b2RecWrite_RevoluteJointEnableSpring(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	if int32FromUint8(enableSpring) != int32FromUint8((*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableSpring) {
		(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableSpring = enableSpring
		(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).SpringImpulse = float32FromFloat32(0)
	}
}

func b2RevoluteJoint_IsSpringEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableSpring
}

func b2RevoluteJoint_SetSpringHertz(tls *_Stack, jointId JointId, hertz float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_RevoluteJointSetSpringHertz
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, hertz)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8368, uintptr(unsafe.Pointer(&__func__60))))
		if v1 = !!(b2IsValidFloat(tls, hertz) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8368, __ccgo_ts+21321, int32FromInt32(55))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_RevoluteJointSetSpringHertz)(unsafe.Add(nil, bp)) = b2RecArgs_RevoluteJointSetSpringHertz{
			Joint: jointId,
			Hertz: hertz,
		}
		b2RecWrite_RevoluteJointSetSpringHertz(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).Hertz = hertz
}

func b2RevoluteJoint_GetSpringHertz(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).Hertz
}

func b2RevoluteJoint_SetSpringDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_RevoluteJointSetSpringDampingRatio
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, dampingRatio)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8392, uintptr(unsafe.Pointer(&__func__61))))
		if v1 = !!(b2IsValidFloat(tls, dampingRatio) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8392, __ccgo_ts+21321, int32FromInt32(71))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_RevoluteJointSetSpringDampingRatio)(unsafe.Add(nil, bp)) = b2RecArgs_RevoluteJointSetSpringDampingRatio{
			Joint:        jointId,
			DampingRatio: dampingRatio,
		}
		b2RecWrite_RevoluteJointSetSpringDampingRatio(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).DampingRatio = dampingRatio
}

func b2RevoluteJoint_GetSpringDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).DampingRatio
}

func b2RevoluteJoint_SetTargetAngle(tls *_Stack, jointId JointId, angle float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_RevoluteJointSetTargetAngle
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, angle)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+21353, uintptr(unsafe.Pointer(&__func__62))))
		if v1 = !!(b2IsValidFloat(tls, angle) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+21353, __ccgo_ts+21321, int32FromInt32(87))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_RevoluteJointSetTargetAngle)(unsafe.Add(nil, bp)) = b2RecArgs_RevoluteJointSetTargetAngle{
			Joint: jointId,
			Angle: angle,
		}
		b2RecWrite_RevoluteJointSetTargetAngle(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).TargetAngle = angle
}

func b2RevoluteJoint_GetTargetAngle(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).TargetAngle
}

func b2RevoluteJoint_EnableLimit(tls *_Stack, jointId JointId, enableLimit uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* _a at bp+0 */ b2RecArgs_RevoluteJointEnableLimit
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_RevoluteJointEnableLimit)(unsafe.Add(nil, bp)) = b2RecArgs_RevoluteJointEnableLimit{
			Joint:       jointId,
			EnableLimit: enableLimit,
		}
		b2RecWrite_RevoluteJointEnableLimit(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	if int32FromUint8(enableLimit) != int32FromUint8((*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableLimit) {
		(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableLimit = enableLimit
		(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LowerImpulse = float32FromFloat32(0)
		(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).UpperImpulse = float32FromFloat32(0)
	}
}

func b2RevoluteJoint_IsLimitEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableLimit
}

func b2RevoluteJoint_GetLowerLimit(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LowerAngle
}

func b2RevoluteJoint_GetUpperLimit(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).UpperAngle
}

func b2RevoluteJoint_SetLimits(tls *_Stack, jointId JointId, lower1 float32, upper1 float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var lowerAngle, upperAngle, v10, v11, v13, v14, v15, v16, v17, v19, v20, v4, v5, v6, v8, v9 float32
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_RevoluteJointSetLimits
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = joint, lowerAngle, upperAngle, world, v1, v10, v11, v13, v14, v15, v16, v17, v19, v20, v4, v5, v6, v8, v9
	if int32FromUint8(b2IsValidFloat(tls, lower1)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+20285, uintptr(unsafe.Pointer(&__func__311))))
		if v1 = !!(b2IsValidFloat(tls, lower1) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+20285, __ccgo_ts+21321, int32FromInt32(134))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if int32FromUint8(b2IsValidFloat(tls, upper1)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+20309, uintptr(unsafe.Pointer(&__func__311))))
		if v1 = !!(b2IsValidFloat(tls, upper1) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+20309, __ccgo_ts+21321, int32FromInt32(135))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if v1 = !!(lower1 <= upper1); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20333, __ccgo_ts+21321, int32FromInt32(137))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_RevoluteJointSetLimits)(unsafe.Add(nil, bp)) = b2RecArgs_RevoluteJointSetLimits{
			Joint: jointId,
			Lower: lower1,
			Upper: upper1,
		}
		b2RecWrite_RevoluteJointSetLimits(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	v4 = lower1
	v5 = upper1
	if v4 < v5 {
		v8 = v4
	} else {
		v8 = v5
	}
	v6 = v8
	goto _7
_7:
	lowerAngle = v6
	v9 = lower1
	v10 = upper1
	if v9 > v10 {
		v13 = v9
	} else {
		v13 = v10
	}
	v11 = v13
	goto _12
_12:
	upperAngle = v11
	v14 = lowerAngle
	v15 = float32(-float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	v16 = float32(float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	if v14 < v15 {
		v19 = v15
	} else {
		if v14 > v16 {
			v20 = v16
		} else {
			v20 = v14
		}
		v19 = v20
	}
	v17 = v19
	goto _18
_18:
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LowerAngle = v17
	v4 = upperAngle
	v5 = float32(-float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	v6 = float32(float32FromFloat32(0.99) * float32FromFloat32(3.14159265359))
	if v4 < v5 {
		v9 = v5
	} else {
		if v4 > v6 {
			v10 = v6
		} else {
			v10 = v4
		}
		v9 = v10
	}
	v8 = v9
	goto _25
_25:
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).UpperAngle = v8
}

func b2RevoluteJoint_EnableMotor(tls *_Stack, jointId JointId, enableMotor uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* _a at bp+0 */ b2RecArgs_RevoluteJointEnableMotor
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_RevoluteJointEnableMotor)(unsafe.Add(nil, bp)) = b2RecArgs_RevoluteJointEnableMotor{
			Joint:       jointId,
			EnableMotor: enableMotor,
		}
		b2RecWrite_RevoluteJointEnableMotor(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	if int32FromUint8(enableMotor) != int32FromUint8((*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableMotor) {
		(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableMotor = enableMotor
		(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MotorImpulse = float32FromFloat32(0)
	}
}

func b2RevoluteJoint_IsMotorEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableMotor
}

func b2RevoluteJoint_SetTargetMotorSpeed(tls *_Stack, jointId JointId, motorSpeed float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_RevoluteJointSetTargetMotorSpeed
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, motorSpeed)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8423, uintptr(unsafe.Pointer(&__func__63))))
		if v1 = !!(b2IsValidFloat(tls, motorSpeed) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8423, __ccgo_ts+21321, int32FromInt32(169))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_RevoluteJointSetTargetMotorSpeed)(unsafe.Add(nil, bp)) = b2RecArgs_RevoluteJointSetTargetMotorSpeed{
			Joint:      jointId,
			MotorSpeed: motorSpeed,
		}
		b2RecWrite_RevoluteJointSetTargetMotorSpeed(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MotorSpeed = motorSpeed
}

func b2RevoluteJoint_GetTargetMotorSpeed(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MotorSpeed
}

func b2RevoluteJoint_SetMaxMotorTorque(tls *_Stack, jointId JointId, torque float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_RevoluteJointSetMaxMotorTorque
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, torque)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+2758, uintptr(unsafe.Pointer(&__func__64))))
		if v1 = !!(b2IsValidFloat(tls, torque) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2758, __ccgo_ts+21321, int32FromInt32(185))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_RevoluteJointSetMaxMotorTorque)(unsafe.Add(nil, bp)) = b2RecArgs_RevoluteJointSetMaxMotorTorque{
			Joint:  jointId,
			Torque: torque,
		}
		b2RecWrite_RevoluteJointSetMaxMotorTorque(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxMotorTorque = torque
}

func b2RevoluteJoint_GetMaxMotorTorque(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxMotorTorque
}

func b2RevoluteJoint_GetCurrentAngle(tls *_Stack, jointId JointId) (r float32) {
	var angle, c, s, v11 float32
	var jointSim, world uintptr
	var qA, qB, qr, v1, v10, v2, v3, v5, v6, v7, v9 Rot
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angle, c, jointSim, qA, qB, qr, s, world, v1, v10, v11, v2, v3, v5, v6, v7, v9
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	jointSim = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	v1 = b2GetBodyTransform(tls, world, (*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdA).Q
	v2 = (*b2JointSim)(unsafe.Pointer(jointSim)).LocalFrameA.Q
	qr.S = float32(v1.S*v2.C) + float32(v1.C*v2.S)
	qr.C = float32(v1.C*v2.C) - float32(v1.S*v2.S)
	v3 = qr
	goto _4
_4:
	qA = v3
	v5 = b2GetBodyTransform(tls, world, (*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdB).Q
	v6 = (*b2JointSim)(unsafe.Pointer(jointSim)).LocalFrameB.Q
	qr.S = float32(v5.S*v6.C) + float32(v5.C*v6.S)
	qr.C = float32(v5.C*v6.C) - float32(v5.S*v6.S)
	v7 = qr
	goto _8
_8:
	qB = v7
	v9 = qA
	v10 = qB
	s = float32(v9.C*v10.S) - float32(v9.S*v10.C)
	c = float32(v9.C*v10.C) + float32(v9.S*v10.S)
	v11 = b2Atan2(tls, s, c)
	goto _12
_12:
	angle = v11
	return angle
}

func b2RevoluteJoint_GetCurrentSpeed(tls *_Stack, jointId JointId) (r float32) {
	var base, bodyA, bodyB, bodyStateA, bodyStateB, set, world1, v10, v12, v3, v4, v6, v9 uintptr
	var wA, wB, v15, v16 float32
	var v1, v11, v14, v2, v5, v8 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = base, bodyA, bodyB, bodyStateA, bodyStateB, set, wA, wB, world1, v1, v10, v11, v12, v14, v15, v16, v2, v3, v4, v5, v6, v8, v9
	world1 = b2GetWorld(tls, int32FromUint16(jointId.World0))
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	if v1 = !!(0 <= (*b2JointSim)(unsafe.Pointer(base)).BodyIdA && (*b2JointSim)(unsafe.Pointer(base)).BodyIdA < (*b2World)(unsafe.Pointer(world1)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20413, __ccgo_ts+21321, int32FromInt32(215))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b2World)(unsafe.Pointer(world1)).Bodies.Data + uintptr((*b2JointSim)(unsafe.Pointer(base)).BodyIdA)*104
	if v2 = !!(0 <= (*b2JointSim)(unsafe.Pointer(base)).BodyIdB && (*b2JointSim)(unsafe.Pointer(base)).BodyIdB < (*b2World)(unsafe.Pointer(world1)).Bodies.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+20483, __ccgo_ts+21321, int32FromInt32(216))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b2World)(unsafe.Pointer(world1)).Bodies.Data + uintptr((*b2JointSim)(unsafe.Pointer(base)).BodyIdB)*104
	v3 = world1
	v4 = bodyA
	if (*b2Body)(unsafe.Pointer(v4)).SetIndex == int32(b2_awakeSet) {
		if v5 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(v3)).SolverSets.Count); !v5 {
			b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+698, int32FromInt32(254))
		}
		_ = v5 || bool(int32FromInt32(0) != 0)
		set = (*b2World)(unsafe.Pointer(v3)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
		if v8 = !!(0 <= (*b2Body)(unsafe.Pointer(v4)).LocalIndex && (*b2Body)(unsafe.Pointer(v4)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Count); !v8 {
			b2InternalAssert(tls, __ccgo_ts+790, __ccgo_ts+698, int32FromInt32(255))
		}
		_ = v8 || bool(int32FromInt32(0) != 0)
		v6 = (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Data + uintptr((*b2Body)(unsafe.Pointer(v4)).LocalIndex)*32
		goto _7
	}
	v6 = uintptrFromInt32(0)
	goto _7
_7:
	bodyStateA = v6
	v9 = world1
	v10 = bodyB
	if (*b2Body)(unsafe.Pointer(v10)).SetIndex == int32(b2_awakeSet) {
		if v11 = !!(bool(0 <= int32(b2_awakeSet)) && int32(b2_awakeSet) < (*b2World)(unsafe.Pointer(v9)).SolverSets.Count); !v11 {
			b2InternalAssert(tls, __ccgo_ts+720, __ccgo_ts+698, int32FromInt32(254))
		}
		_ = v11 || bool(int32FromInt32(0) != 0)
		set = (*b2World)(unsafe.Pointer(v9)).SolverSets.Data + uintptr(int32(b2_awakeSet))*88
		if v14 = !!(0 <= (*b2Body)(unsafe.Pointer(v10)).LocalIndex && (*b2Body)(unsafe.Pointer(v10)).LocalIndex < (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Count); !v14 {
			b2InternalAssert(tls, __ccgo_ts+790, __ccgo_ts+698, int32FromInt32(255))
		}
		_ = v14 || bool(int32FromInt32(0) != 0)
		v12 = (*b2SolverSet)(unsafe.Pointer(set)).BodyStates.Data + uintptr((*b2Body)(unsafe.Pointer(v10)).LocalIndex)*32
		goto _13
	}
	v12 = uintptrFromInt32(0)
	goto _13
_13:
	bodyStateB = v12
	if bodyStateA != 0 {
		v15 = (*b2BodyState)(unsafe.Pointer(bodyStateA)).AngularVelocity
	} else {
		v15 = float32FromFloat32(0)
	}
	wA = v15
	if bodyStateB != 0 {
		v16 = (*b2BodyState)(unsafe.Pointer(bodyStateB)).AngularVelocity
	} else {
		v16 = float32FromFloat32(0)
	}
	wB = v16
	return wB - wA
}

func b2RevoluteJoint_GetCurrentMotorTorque(tls *_Stack, jointId JointId) (r float32) {
	var joint, world uintptr
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return float32((*b2World)(unsafe.Pointer(world)).Inv_h * (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MotorImpulse)
}

func b2GetRevoluteJointForce(tls *_Stack, world uintptr, base uintptr) (r Vec2) {
	var force, v2, v3 Vec2
	var v1 float32
	_, _, _, _ = force, v1, v2, v3
	v1 = (*b2World)(unsafe.Pointer(world)).Inv_h
	v2 = (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 92))).LinearImpulse
	v3 = Vec2{
		X: float32(v1 * v2.X),
		Y: float32(v1 * v2.Y),
	}
	goto _4
_4:
	force = v3
	return force
}

func b2GetRevoluteJointTorque(tls *_Stack, world uintptr, base uintptr) (r float32) {
	var revolute uintptr
	var torque float32
	_, _ = revolute, torque
	revolute = base + 92
	torque = float32((*b2World)(unsafe.Pointer(world)).Inv_h * ((*b2RevoluteJoint)(unsafe.Pointer(revolute)).MotorImpulse + (*b2RevoluteJoint)(unsafe.Pointer(revolute)).LowerImpulse - (*b2RevoluteJoint)(unsafe.Pointer(revolute)).UpperImpulse))
	return torque
}

func b2PrepareRevoluteJoint(tls *_Stack, base uintptr, context uintptr) {
	var a11, a2, a3, iA, iB, k, mA, mB, omega, v39, v40 float32
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world uintptr
	var idA, idB, localIndexA, localIndexB, v9 int32
	var qr, v11, v12, v13 Rot
	var v1, v2, v3, v4 bool
	var v15, v16, v17, v20, v21 Vec2
	var v35, v36 Pos
	var v42 b2Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a2, a3, bodyA, bodyB, bodySimA, bodySimB, iA, iB, idA, idB, joint, k, localIndexA, localIndexB, mA, mB, omega, qr, setA, setB, world, v1, v11, v12, v13, v15, v16, v17, v2, v20, v21, v3, v35, v36, v39, v4, v40, v42, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_revoluteJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+21377, __ccgo_ts+21321, int32FromInt32(248))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// chase body id to the solver set where the body lives
	idA = (*b2JointSim)(unsafe.Pointer(base)).BodyIdA
	idB = (*b2JointSim)(unsafe.Pointer(base)).BodyIdB
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	if v1 = !!(0 <= idA && idA < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8507, __ccgo_ts+21321, int32FromInt32(256))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idA)*104
	if v2 = !!(0 <= idB && idB < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+8557, __ccgo_ts+21321, int32FromInt32(257))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idB)*104
	if v3 = !!((*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet)); !v3 {
		b2InternalAssert(tls, __ccgo_ts+8607, __ccgo_ts+21321, int32FromInt32(259))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(bodyA)).SetIndex && (*b2Body)(unsafe.Pointer(bodyA)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8672, __ccgo_ts+21321, int32FromInt32(260))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	setA = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(bodyA)).SetIndex)*88
	if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(bodyB)).SetIndex && (*b2Body)(unsafe.Pointer(bodyB)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+8750, __ccgo_ts+21321, int32FromInt32(261))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	setB = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(bodyB)).SetIndex)*88
	localIndexA = (*b2Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b2Body)(unsafe.Pointer(bodyB)).LocalIndex
	if v3 = !!(0 <= localIndexA && localIndexA < (*b2SolverSet)(unsafe.Pointer(setA)).BodySims.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+8828, __ccgo_ts+21321, int32FromInt32(266))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	bodySimA = (*b2SolverSet)(unsafe.Pointer(setA)).BodySims.Data + uintptr(localIndexA)*96
	if v4 = !!(0 <= localIndexB && localIndexB < (*b2SolverSet)(unsafe.Pointer(setB)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+8895, __ccgo_ts+21321, int32FromInt32(267))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySimB = (*b2SolverSet)(unsafe.Pointer(setB)).BodySims.Data + uintptr(localIndexB)*96
	mA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvMass
	iA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvInertia
	mB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvMass
	iB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvInertia
	(*b2JointSim)(unsafe.Pointer(base)).InvMassA = mA
	(*b2JointSim)(unsafe.Pointer(base)).InvMassB = mB
	(*b2JointSim)(unsafe.Pointer(base)).InvIA = iA
	(*b2JointSim)(unsafe.Pointer(base)).InvIB = iB
	joint = base + 92
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) {
		v9 = localIndexA
	} else {
		v9 = -int32(1)
	}
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).IndexA = v9
	if (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) {
		v9 = localIndexB
	} else {
		v9 = -int32(1)
	}
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).IndexB = v9
	// Compute joint anchor frames with world space rotation, relative to center of mass.
	// Avoid round-off here as much as possible.
	// b2Vec2 pf = (xf.p - c) + rot(xf.q, f.p)
	// pf = xf.p - (xf.p + rot(xf.q, lc)) + rot(xf.q, f.p)
	// pf = rot(xf.q, f.p - lc)
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v12 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	qr.S = float32(v11.S*v12.C) + float32(v11.C*v12.S)
	qr.C = float32(v11.C*v12.C) - float32(v11.S*v12.S)
	v13 = qr
	goto _14
_14:
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).FrameA.Q = v13
	v15 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA.P
	v16 = (*b2BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v17 = Vec2{
		X: v15.X - v16.X,
		Y: v15.Y - v16.Y,
	}
	goto _18
_18:
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v20 = v17
	v21 = Vec2{
		X: float32(v11.C*v20.X) - float32(v11.S*v20.Y),
		Y: float32(v11.S*v20.X) + float32(v11.C*v20.Y),
	}
	goto _22
_22:
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).FrameA.P = v21
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v12 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameB.Q
	qr.S = float32(v11.S*v12.C) + float32(v11.C*v12.S)
	qr.C = float32(v11.C*v12.C) - float32(v11.S*v12.S)
	v13 = qr
	goto _26
_26:
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).FrameB.Q = v13
	v15 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v16 = (*b2BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v17 = Vec2{
		X: v15.X - v16.X,
		Y: v15.Y - v16.Y,
	}
	goto _30
_30:
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v20 = v17
	v21 = Vec2{
		X: float32(v11.C*v20.X) - float32(v11.S*v20.Y),
		Y: float32(v11.S*v20.X) + float32(v11.C*v20.Y),
	}
	goto _34
_34:
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).FrameB.P = v21
	// Compute the initial center delta. Incremental position updates are relative to this.
	v35 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Center
	v36 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Center
	v15 = Vec2{
		X: v35.X - v36.X,
		Y: v35.Y - v36.Y,
	}
	goto _38
_38:
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).DeltaCenter = v15
	k = iA + iB
	if k > float32FromFloat32(0) {
		v39 = float32FromFloat32(1) / k
	} else {
		v39 = float32FromFloat32(0)
	}
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).AxialMass = v39
	v39 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).Hertz
	v40 = (*b2StepContext)(unsafe.Pointer(context)).H
	if v39 == float32FromFloat32(0) {
		v42 = b2Softness{}
		goto _43
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v39)
	a11 = float32(float32FromFloat32(2)*(*b2RevoluteJoint)(unsafe.Pointer(joint)).DampingRatio) + float32(v40*omega)
	a2 = float32(float32(v40*omega) * a11)
	a3 = float32FromFloat32(1) / (float32FromFloat32(1) + a2)
	v42 = b2Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a2 * a3),
		ImpulseScale: a3,
	}
	goto _43
_43:
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).SpringSoftness = v42
	if int32FromUint8((*b2StepContext)(unsafe.Pointer(context)).EnableWarmStarting) == false1 {
		(*b2RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse = b2Vec2_zero
		(*b2RevoluteJoint)(unsafe.Pointer(joint)).SpringImpulse = float32FromFloat32(0)
		(*b2RevoluteJoint)(unsafe.Pointer(joint)).MotorImpulse = float32FromFloat32(0)
		(*b2RevoluteJoint)(unsafe.Pointer(joint)).LowerImpulse = float32FromFloat32(0)
		(*b2RevoluteJoint)(unsafe.Pointer(joint)).UpperImpulse = float32FromFloat32(0)
	}
}

func b2WarmStartRevoluteJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var axialImpulse, iA, iB, mA, mB, v13 float32
	var joint, stateA, stateB, v2, v3 uintptr
	var rA, rB, v10, v12, v14, v15, v5, v6, v9 Vec2
	var v1 bool
	var v4, v8 Rot
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axialImpulse, iA, iB, joint, mA, mB, rA, rB, stateA, stateB, v1, v10, v12, v13, v14, v15, v2, v3, v4, v5, v6, v8, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_revoluteJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+21377, __ccgo_ts+21321, int32FromInt32(314))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Add(nil, bp)) = b2_identityBodyState16
	joint = base + 92
	if (*b2RevoluteJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2RevoluteJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v2
	if (*b2RevoluteJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2RevoluteJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v3
	v4 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).FrameA.P
	v6 = Vec2{
		X: float32(v4.C*v5.X) - float32(v4.S*v5.Y),
		Y: float32(v4.S*v5.X) + float32(v4.C*v5.Y),
	}
	goto _7
_7:
	rA = v6
	v8 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v9 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).FrameB.P
	v10 = Vec2{
		X: float32(v8.C*v9.X) - float32(v8.S*v9.Y),
		Y: float32(v8.S*v9.X) + float32(v8.C*v9.Y),
	}
	goto _11
_11:
	rB = v10
	axialImpulse = (*b2RevoluteJoint)(unsafe.Pointer(joint)).SpringImpulse + (*b2RevoluteJoint)(unsafe.Pointer(joint)).MotorImpulse + (*b2RevoluteJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b2RevoluteJoint)(unsafe.Pointer(joint)).UpperImpulse
	if (*b2BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b2_dynamicFlag) != 0 {
		v12 = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
		v13 = mA
		v14 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse
		v15 = Vec2{
			X: v12.X - float32(v13*v14.X),
			Y: v12.Y - float32(v13*v14.Y),
		}
		goto _16
	_16:
		(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = v15
		v5 = rA
		v6 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse
		v13 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		goto _20
	_20:
		*(*float32)(unsafe.Add(nil, stateA+8)) -= float32(iA * (v13 + axialImpulse))
	}
	if (*b2BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b2_dynamicFlag) != 0 {
		v5 = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
		v13 = mB
		v6 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse
		v9 = Vec2{
			X: v5.X + float32(v13*v6.X),
			Y: v5.Y + float32(v13*v6.Y),
		}
		goto _25
	_25:
		(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = v9
		v5 = rB
		v6 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse
		v13 = float32(v5.X*v6.Y) - float32(v5.Y*v6.X)
		goto _29
	_29:
		*(*float32)(unsafe.Add(nil, stateB+8)) += float32(iB * (v13 + axialImpulse))
	}
}

func b2SolveRevoluteJoint(tls *_Stack, base uintptr, context uintptr, useBias uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var C, C1, C2, Cdot, Cdot1, Cdot2, Cdot3, a11, a12, a21, a22, angularCdot, axialImpulse, bias, bias1, bias2, det, iA, iB, impulse, impulse1, impulse2, impulse3, impulseScale, impulseScale1, impulseScale2, impulseScale3, jointAngle, jointAngle1, jointAngleDelta, k1, mA, mB, massScale, massScale1, massScale2, massScale3, maxImpulse, oldImpulse, oldImpulse1, oldImpulse2, wA, wB, x, v24, v32, v45, v47, v48, v49, v50, v52, v53, v54 float32
	var Cdot4, b8, bias3, dcA, dcB, impulse4, linearCdot, rA, rB, separation, vA, vB, x1, v17, v18, v21, v22, v25, v26, v28, v29, v30, v33, v34, v36, v37, v38, v40, v41, v42, v77, v78, v81, v82, v84, v85, v86 Vec2
	var K, v117 Mat22
	var a8, k, roundToNearest, twoPi float64
	var fixedRotation uint8
	var joint, stateA, stateB, v2, v3 uintptr
	var qA, qB, qr, r1, relQ, v10, v12, v13, v14, v16, v20, v4, v44, v5, v6, v8, v9 Rot
	var v1 bool
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, C1, C2, Cdot, Cdot1, Cdot2, Cdot3, Cdot4, K, a11, a12, a21, a22, a8, angularCdot, axialImpulse, b8, bias, bias1, bias2, bias3, dcA, dcB, det, fixedRotation, iA, iB, impulse, impulse1, impulse2, impulse3, impulse4, impulseScale, impulseScale1, impulseScale2, impulseScale3, joint, jointAngle, jointAngle1, jointAngleDelta, k, k1, linearCdot, mA, mB, massScale, massScale1, massScale2, massScale3, maxImpulse, oldImpulse, oldImpulse1, oldImpulse2, qA, qB, qr, r1, rA, rB, relQ, roundToNearest, separation, stateA, stateB, twoPi, vA, vB, wA, wB, x, x1, v1, v10, v117, v12, v13, v14, v16, v17, v18, v2, v20, v21, v22, v24, v25, v26, v28, v29, v3, v30, v32, v33, v34, v36, v37, v38, v4, v40, v41, v42, v44, v45, v47, v48, v49, v5, v50, v52, v53, v54, v6, v77, v78, v8, v81, v82, v84, v85, v86, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_revoluteJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+21377, __ccgo_ts+21321, int32FromInt32(348))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Add(nil, bp)) = b2_identityBodyState16
	joint = base + 92
	if (*b2RevoluteJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2RevoluteJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v2
	if (*b2RevoluteJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2RevoluteJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v3
	vA = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	v4 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).FrameA.Q
	qr.S = float32(v4.S*v5.C) + float32(v4.C*v5.S)
	qr.C = float32(v4.C*v5.C) - float32(v4.S*v5.S)
	v6 = qr
	goto _7
_7:
	qA = v6
	v8 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v9 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).FrameB.Q
	qr.S = float32(v8.S*v9.C) + float32(v8.C*v9.S)
	qr.C = float32(v8.C*v9.C) - float32(v8.S*v9.S)
	v10 = qr
	goto _11
_11:
	qB = v10
	v12 = qA
	v13 = qB
	r1.S = float32(v12.C*v13.S) - float32(v12.S*v13.C)
	r1.C = float32(v12.C*v13.C) + float32(v12.S*v13.S)
	v14 = r1
	goto _15
_15:
	relQ = v14
	v16 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v17 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).FrameA.P
	v18 = Vec2{
		X: float32(v16.C*v17.X) - float32(v16.S*v17.Y),
		Y: float32(v16.S*v17.X) + float32(v16.C*v17.Y),
	}
	goto _19
_19:
	// current anchors
	rA = v18
	v20 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v21 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).FrameB.P
	v22 = Vec2{
		X: float32(v20.C*v21.X) - float32(v20.S*v21.Y),
		Y: float32(v20.S*v21.X) + float32(v20.C*v21.Y),
	}
	goto _23
_23:
	rB = v22
	v24 = wB
	v25 = rB
	v26 = Vec2{
		X: float32(-v24 * v25.Y),
		Y: float32(v24 * v25.X),
	}
	goto _27
_27:
	v28 = vB
	v29 = v26
	v30 = Vec2{
		X: v28.X + v29.X,
		Y: v28.Y + v29.Y,
	}
	goto _31
_31:
	v32 = wA
	v33 = rA
	v34 = Vec2{
		X: float32(-v32 * v33.Y),
		Y: float32(v32 * v33.X),
	}
	goto _35
_35:
	v36 = vA
	v37 = v34
	v38 = Vec2{
		X: v36.X + v37.X,
		Y: v36.Y + v37.Y,
	}
	goto _39
_39:
	v40 = v30
	v41 = v38
	v42 = Vec2{
		X: v40.X - v41.X,
		Y: v40.Y - v41.Y,
	}
	goto _43
_43:
	linearCdot = v42
	angularCdot = wB - wA
	axialImpulse = float32FromFloat32(0)
	k1 = iA + iB
	fixedRotation = boolUint8(k1 == float32FromFloat32(0))
	// Solve spring.
	if (*b2RevoluteJoint)(unsafe.Pointer(joint)).EnableSpring != 0 && int32FromUint8(fixedRotation) == false1 {
		v44 = relQ
		v45 = b2Atan2(tls, v44.S, v44.C)
		goto _46
	_46:
		jointAngle = v45
		v47 = jointAngle - (*b2RevoluteJoint)(unsafe.Pointer(joint)).TargetAngle
		v48 = -float32FromFloat32(1e+06)
		v49 = float32FromFloat32(1e+06)
		if v47 < v48 {
			v52 = v48
		} else {
			if v47 > v49 {
				v53 = v49
			} else {
				v53 = v47
			}
			v52 = v53
		}
		v50 = v52
		goto _51
	_51:
		x = v50
		twoPi = float64(float32FromFloat32(2) * float32FromFloat32(3.14159265359))
		roundToNearest = float64(6.755399441055744e+15)
		a8 = float64(x)
		k = a8/twoPi + roundToNearest - roundToNearest
		v54 = float32(a8 - float64(k*twoPi))
		goto _55
	_55:
		jointAngleDelta = v54
		C = jointAngleDelta
		bias = float32((*b2RevoluteJoint)(unsafe.Pointer(joint)).SpringSoftness.BiasRate * C)
		massScale = (*b2RevoluteJoint)(unsafe.Pointer(joint)).SpringSoftness.MassScale
		impulseScale = (*b2RevoluteJoint)(unsafe.Pointer(joint)).SpringSoftness.ImpulseScale
		Cdot = angularCdot
		impulse = float32(float32(-massScale*(*b2RevoluteJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot+bias)) - float32(impulseScale*(*b2RevoluteJoint)(unsafe.Pointer(joint)).SpringImpulse)
		*(*float32)(unsafe.Add(nil, joint+8)) += impulse
		angularCdot = angularCdot + float32(k1*impulse)
		axialImpulse = axialImpulse + impulse
	}
	// Solve motor constraint.
	if (*b2RevoluteJoint)(unsafe.Pointer(joint)).EnableMotor != 0 && int32FromUint8(fixedRotation) == false1 {
		Cdot1 = angularCdot - (*b2RevoluteJoint)(unsafe.Pointer(joint)).MotorSpeed
		oldImpulse = (*b2RevoluteJoint)(unsafe.Pointer(joint)).MotorImpulse
		maxImpulse = float32((*b2StepContext)(unsafe.Pointer(context)).H * (*b2RevoluteJoint)(unsafe.Pointer(joint)).MaxMotorTorque)
		v24 = float32(-(*b2RevoluteJoint)(unsafe.Pointer(joint)).AxialMass * Cdot1)
		v32 = -maxImpulse - oldImpulse
		v45 = maxImpulse - oldImpulse
		if v24 < v32 {
			v48 = v32
		} else {
			if v24 > v45 {
				v49 = v45
			} else {
				v49 = v24
			}
			v48 = v49
		}
		v47 = v48
		goto _60
	_60:
		impulse1 = v47
		(*b2RevoluteJoint)(unsafe.Pointer(joint)).MotorImpulse = oldImpulse + impulse1
		angularCdot = angularCdot + float32(k1*impulse1)
		axialImpulse = axialImpulse + impulse1
	}
	if (*b2RevoluteJoint)(unsafe.Pointer(joint)).EnableLimit != 0 && int32FromUint8(fixedRotation) == false1 {
		v4 = relQ
		v24 = b2Atan2(tls, v4.S, v4.C)
		goto _65
	_65:
		jointAngle1 = v24
		// Lower limit
		C1 = jointAngle1 - (*b2RevoluteJoint)(unsafe.Pointer(joint)).LowerAngle
		bias1 = float32FromFloat32(0)
		massScale1 = float32FromFloat32(1)
		impulseScale1 = float32FromFloat32(0)
		if C1 > float32FromFloat32(0) {
			// speculation
			bias1 = float32(C1 * (*b2StepContext)(unsafe.Pointer(context)).Inv_h)
		} else {
			if useBias != 0 {
				bias1 = float32((*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * C1)
				massScale1 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
				impulseScale1 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
			}
		}
		Cdot2 = angularCdot
		oldImpulse1 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).LowerImpulse
		v24 = float32(float32(-massScale1*(*b2RevoluteJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot2+bias1)) - float32(impulseScale1*oldImpulse1)
		v32 = -oldImpulse1
		if v24 > v32 {
			v47 = v24
		} else {
			v47 = v32
		}
		v45 = v47
		goto _69
	_69:
		impulse2 = v45
		(*b2RevoluteJoint)(unsafe.Pointer(joint)).LowerImpulse = oldImpulse1 + impulse2
		angularCdot = angularCdot + float32(k1*impulse2)
		axialImpulse = axialImpulse + impulse2
		// Upper limit
		// Note: signs are flipped to keep C positive when the constraint is satisfied.
		// This also keeps the impulse positive when the limit is active.
		C2 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).UpperAngle - jointAngle1
		bias2 = float32FromFloat32(0)
		massScale2 = float32FromFloat32(1)
		impulseScale2 = float32FromFloat32(0)
		if C2 > float32FromFloat32(0) {
			// speculation
			bias2 = float32(C2 * (*b2StepContext)(unsafe.Pointer(context)).Inv_h)
		} else {
			if useBias != 0 {
				bias2 = float32((*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * C2)
				massScale2 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
				impulseScale2 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
			}
		}
		// sign flipped on Cdot
		Cdot3 = -angularCdot
		oldImpulse2 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).UpperImpulse
		v24 = float32(float32(-massScale2*(*b2RevoluteJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot3+bias2)) - float32(impulseScale2*oldImpulse2)
		v32 = -oldImpulse2
		if v24 > v32 {
			v47 = v24
		} else {
			v47 = v32
		}
		v45 = v47
		goto _74
	_74:
		impulse3 = v45
		(*b2RevoluteJoint)(unsafe.Pointer(joint)).UpperImpulse = oldImpulse2 + impulse3
		// sign flipped on applied impulse
		axialImpulse = axialImpulse - impulse3
	}
	// Solve point-to-point constraint
	v24 = iA
	v17 = rA
	v18 = Vec2{
		X: float32(-v24 * v17.Y),
		Y: float32(v24 * v17.X),
	}
	goto _79
_79:
	v32 = iB
	v21 = rB
	v22 = Vec2{
		X: float32(-v32 * v21.Y),
		Y: float32(v32 * v21.X),
	}
	goto _83
_83:
	v25 = v18
	v26 = v22
	v28 = Vec2{
		X: v25.X + v26.X,
		Y: v25.Y + v26.Y,
	}
	goto _87
_87:
	v29 = linearCdot
	v45 = axialImpulse
	v30 = v28
	v33 = Vec2{
		X: v29.X + float32(v45*v30.X),
		Y: v29.Y + float32(v45*v30.Y),
	}
	goto _92
_92:
	// J = [-I -r1_skew I r2_skew]
	// r_skew = [-ry; rx]
	// K = [ mA+r1y^2*iA+mB+r2y^2*iB,  -r1y*iA*r1x-r2y*iB*r2x]
	//     [  -r1y*iA*r1x-r2y*iB*r2x, mA+r1x^2*iA+mB+r2x^2*iB]
	Cdot4 = v33
	bias3 = b2Vec2_zero
	massScale3 = float32FromFloat32(1)
	impulseScale3 = float32FromFloat32(0)
	if useBias != 0 {
		dcA = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaPosition
		dcB = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
		v34 = dcB
		v36 = dcA
		v37 = Vec2{
			X: v34.X - v36.X,
			Y: v34.Y - v36.Y,
		}
		goto _96
	_96:
		v38 = rB
		v40 = rA
		v41 = Vec2{
			X: v38.X - v40.X,
			Y: v38.Y - v40.Y,
		}
		goto _100
	_100:
		v42 = v37
		v77 = v41
		v78 = Vec2{
			X: v42.X + v77.X,
			Y: v42.Y + v77.Y,
		}
		goto _104
	_104:
		v81 = v78
		v82 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).DeltaCenter
		v84 = Vec2{
			X: v81.X + v82.X,
			Y: v81.Y + v82.Y,
		}
		goto _108
	_108:
		separation = v84
		v47 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate
		v85 = separation
		v86 = Vec2{
			X: float32(v47 * v85.X),
			Y: float32(v47 * v85.Y),
		}
		goto _112
	_112:
		bias3 = v86
		massScale3 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
		impulseScale3 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
	}
	K.Cx.X = mA + mB + float32(float32(rA.Y*rA.Y)*iA) + float32(float32(rB.Y*rB.Y)*iB)
	K.Cy.X = float32(float32(-rA.Y*rA.X)*iA) - float32(float32(rB.Y*rB.X)*iB)
	K.Cx.Y = K.Cy.X
	K.Cy.Y = mA + mB + float32(float32(rA.X*rA.X)*iA) + float32(float32(rB.X*rB.X)*iB)
	v17 = Cdot4
	v18 = bias3
	v21 = Vec2{
		X: v17.X + v18.X,
		Y: v17.Y + v18.Y,
	}
	goto _116
_116:
	v117 = K
	v22 = v21
	a11 = v117.Cx.X
	a12 = v117.Cy.X
	a21 = v117.Cx.Y
	a22 = v117.Cy.Y
	det = float32(a11*a22) - float32(a12*a21)
	if det != float32FromFloat32(0) {
		det = float32FromFloat32(1) / det
	}
	x1 = Vec2{
		X: float32(det * (float32(a22*v22.X) - float32(a12*v22.Y))),
		Y: float32(det * (float32(a11*v22.Y) - float32(a21*v22.X))),
	}
	v25 = x1
	goto _120
_120:
	b8 = v25
	impulse4.X = float32(-massScale3*b8.X) - float32(impulseScale3*(*b2RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse.X)
	impulse4.Y = float32(-massScale3*b8.Y) - float32(impulseScale3*(*b2RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse.Y)
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse.X += impulse4.X
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse.Y += impulse4.Y
	v17 = vA
	v24 = mA
	v18 = impulse4
	v21 = Vec2{
		X: v17.X - float32(v24*v18.X),
		Y: v17.Y - float32(v24*v18.Y),
	}
	goto _125
_125:
	vA = v21
	v17 = rA
	v18 = impulse4
	v24 = float32(v17.X*v18.Y) - float32(v17.Y*v18.X)
	goto _129
_129:
	wA = wA - float32(iA*(v24+axialImpulse))
	v17 = vB
	v24 = mB
	v18 = impulse4
	v21 = Vec2{
		X: v17.X + float32(v24*v18.X),
		Y: v17.Y + float32(v24*v18.Y),
	}
	goto _134
_134:
	vB = v21
	v17 = rB
	v18 = impulse4
	v24 = float32(v17.X*v18.Y) - float32(v17.Y*v18.X)
	goto _138
_138:
	wB = wB + float32(iB*(v24+axialImpulse))
	if (*b2BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b2_dynamicFlag) != 0 {
		(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
		(*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	}
	if (*b2BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b2_dynamicFlag) != 0 {
		(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
		(*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
	}
}

func b2DrawRevoluteJoint(tls *_Stack, draw uintptr, base uintptr, transformA WorldTransform, transformB WorldTransform, drawScale float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var C, frameA, frameB, v16, v18, v2, v32 WorldTransform
	var c, jointAngle, lowerAngle, radius, s, upperAngle, v52 float32
	var color b2HexColor1
	var cs CosSin
	var joint uintptr
	var q2, qr, rotHi, rotLo, v20, v21, v22, v24, v34, v4, v42, v5, v6, v8 Rot
	var r1, rhi, rlo, rx, v1, v10, v13, v25, v26, v29, v9 Vec2
	var v11 bool
	var v12, v14, v28, v30 Pos
	var v19, v3 Transform
	var _ /* buffer at bp+0 */ [32]uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, c, color, cs, frameA, frameB, joint, jointAngle, lowerAngle, q2, qr, r1, radius, rhi, rlo, rotHi, rotLo, rx, s, upperAngle, v1, v11, v10, v12, v13, v14, v16, v18, v19, v2, v20, v21, v22, v24, v25, v26, v28, v29, v3, v30, v32, v34, v4, v42, v5, v52, v6, v8, v9
	if v11 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_revoluteJoint)); !v11 {
		b2InternalAssert(tls, __ccgo_ts+21377, __ccgo_ts+21321, int32FromInt32(535))
	}
	_ = v11 || bool(int32FromInt32(0) != 0)
	joint = base + 92
	v2 = transformA
	v3 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA
	v4 = v2.Q
	v5 = v3.Q
	qr.S = float32(v4.S*v5.C) + float32(v4.C*v5.S)
	qr.C = float32(v4.C*v5.C) - float32(v4.S*v5.S)
	v6 = qr
	goto _7
_7:
	C.Q = v6
	v8 = v2.Q
	v9 = v3.P
	v10 = Vec2{
		X: float32(v8.C*v9.X) - float32(v8.S*v9.Y),
		Y: float32(v8.S*v9.X) + float32(v8.C*v9.Y),
	}
	goto _11
_11:
	v12 = v2.P
	v13 = v10
	v14 = Pos{
		X: v12.X + v13.X,
		Y: v12.Y + v13.Y,
	}
	goto _15
_15:
	C.P = v14
	v16 = C
	goto _17
_17:
	frameA = v16
	v18 = transformB
	v19 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameB
	v20 = v18.Q
	v21 = v19.Q
	qr.S = float32(v20.S*v21.C) + float32(v20.C*v21.S)
	qr.C = float32(v20.C*v21.C) - float32(v20.S*v21.S)
	v22 = qr
	goto _23
_23:
	C.Q = v22
	v24 = v18.Q
	v25 = v19.P
	v26 = Vec2{
		X: float32(v24.C*v25.X) - float32(v24.S*v25.Y),
		Y: float32(v24.S*v25.X) + float32(v24.C*v25.Y),
	}
	goto _27
_27:
	v28 = v18.P
	v29 = v26
	v30 = Pos{
		X: v28.X + v29.X,
		Y: v28.Y + v29.Y,
	}
	goto _31
_31:
	C.P = v30
	v32 = C
	goto _33
_33:
	frameB = v32
	radius = float32(float32FromFloat32(0.25) * drawScale)
	(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawCircleFcn})))(tls, frameB.P, radius, int32(b2_colorGray), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	rx = Vec2{
		X: radius,
	}
	v4 = frameA.Q
	v9 = rx
	v10 = Vec2{
		X: float32(v4.C*v9.X) - float32(v4.S*v9.Y),
		Y: float32(v4.S*v9.X) + float32(v4.C*v9.Y),
	}
	goto _37
_37:
	r1 = v10
	v12 = frameA.P
	v13 = r1
	v14 = Pos{
		X: v12.X + v13.X,
		Y: v12.Y + v13.Y,
	}
	goto _41
_41:
	(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, frameA.P, v14, int32(b2_colorGray), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	v4 = frameB.Q
	v9 = rx
	v10 = Vec2{
		X: float32(v4.C*v9.X) - float32(v4.S*v9.Y),
		Y: float32(v4.S*v9.X) + float32(v4.C*v9.Y),
	}
	goto _45
_45:
	r1 = v10
	v12 = frameB.P
	v9 = r1
	v14 = Pos{
		X: v12.X + v9.X,
		Y: v12.Y + v9.Y,
	}
	goto _49
_49:
	(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, frameB.P, v14, int32(b2_colorBlue), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawJointExtras != 0 {
		v4 = frameA.Q
		v5 = frameB.Q
		s = float32(v4.C*v5.S) - float32(v4.S*v5.C)
		c = float32(v4.C*v5.C) + float32(v4.S*v5.S)
		v52 = b2Atan2(tls, s, c)
		goto _53
	_53:
		jointAngle = v52
		__builtin_snprintf(tls, bp, uint64(32), __ccgo_ts+21408, vaList(bp+40, float64(float32(float32FromFloat32(180)*jointAngle)/float32FromFloat32(3.14159265359))))
		v12 = frameA.P
		v9 = r1
		v14 = Pos{
			X: v12.X + v9.X,
			Y: v12.Y + v9.Y,
		}
		goto _57
	_57:
		(*(*func(*_Stack, Pos, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, v14, bp, int32(b2_colorWhite), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	}
	lowerAngle = (*b2RevoluteJoint)(unsafe.Pointer(joint)).LowerAngle
	upperAngle = (*b2RevoluteJoint)(unsafe.Pointer(joint)).UpperAngle
	if (*b2RevoluteJoint)(unsafe.Pointer(joint)).EnableLimit != 0 {
		cs = b2ComputeCosSin(tls, lowerAngle)
		v4 = Rot{
			C: cs.Cosine,
			S: cs.Sine,
		}
		goto _59
	_59:
		v5 = frameA.Q
		v6 = v4
		qr.S = float32(v5.S*v6.C) + float32(v5.C*v6.S)
		qr.C = float32(v5.C*v6.C) - float32(v5.S*v6.S)
		v8 = qr
		goto _63
	_63:
		rotLo = v8
		v20 = rotLo
		v9 = rx
		v10 = Vec2{
			X: float32(v20.C*v9.X) - float32(v20.S*v9.Y),
			Y: float32(v20.S*v9.X) + float32(v20.C*v9.Y),
		}
		goto _67
	_67:
		rlo = v10
		cs = b2ComputeCosSin(tls, upperAngle)
		v21 = Rot{
			C: cs.Cosine,
			S: cs.Sine,
		}
		goto _69
	_69:
		v22 = frameA.Q
		v24 = v21
		qr.S = float32(v22.S*v24.C) + float32(v22.C*v24.S)
		qr.C = float32(v22.C*v24.C) - float32(v22.S*v24.S)
		v34 = qr
		goto _73
	_73:
		rotHi = v34
		v42 = rotHi
		v13 = rx
		v25 = Vec2{
			X: float32(v42.C*v13.X) - float32(v42.S*v13.Y),
			Y: float32(v42.S*v13.X) + float32(v42.C*v13.Y),
		}
		goto _77
	_77:
		rhi = v25
		v12 = frameB.P
		v26 = rlo
		v14 = Pos{
			X: v12.X + v26.X,
			Y: v12.Y + v26.Y,
		}
		goto _81
	_81:
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, frameB.P, v14, int32(b2_colorGreen), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		v12 = frameB.P
		v9 = rhi
		v14 = Pos{
			X: v12.X + v9.X,
			Y: v12.Y + v9.Y,
		}
		goto _85
	_85:
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, frameB.P, v14, int32(b2_colorRed), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	}
	if (*b2RevoluteJoint)(unsafe.Pointer(joint)).EnableSpring != 0 {
		cs = b2ComputeCosSin(tls, (*b2RevoluteJoint)(unsafe.Pointer(joint)).TargetAngle)
		v4 = Rot{
			C: cs.Cosine,
			S: cs.Sine,
		}
		goto _87
	_87:
		v5 = frameA.Q
		v6 = v4
		qr.S = float32(v5.S*v6.C) + float32(v5.C*v6.S)
		qr.C = float32(v5.C*v6.C) - float32(v5.S*v6.S)
		v8 = qr
		goto _91
	_91:
		q2 = v8
		v20 = q2
		v9 = rx
		v10 = Vec2{
			X: float32(v20.C*v9.X) - float32(v20.S*v9.Y),
			Y: float32(v20.S*v9.X) + float32(v20.C*v9.Y),
		}
		goto _95
	_95:
		v1 = v10
		v12 = frameB.P
		v13 = v1
		v14 = Pos{
			X: v12.X + v13.X,
			Y: v12.Y + v13.Y,
		}
		goto _99
	_99:
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, frameB.P, v14, int32(b2_colorViolet), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	}
	color = int32(b2_colorGold)
	(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, frameA.P, frameB.P, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawJointExtras != 0 {
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, transformA.P, frameA.P, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, transformB.P, frameB.P, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	}
	// char buffer[32];
	// sprintf(buffer, "%.1f", b2Length(joint->impulse));
	// draw->DrawString(pA, buffer, draw->context);
}

func b2TransferJoint(tls *_Stack, world uintptr, targetSet uintptr, sourceSet uintptr, joint uintptr) {
	var color, movedJoint, movedJointSim, sourceSim, targetSim, v10, v13, v5, v6, v7 uintptr
	var colorIndex, localIndex, movedId, movedIndex, newCapacity, newSize, oldCapacity, oldSize, v12, v8, v9 int32
	var v1, v2 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = color, colorIndex, localIndex, movedId, movedIndex, movedJoint, movedJointSim, newCapacity, newSize, oldCapacity, oldSize, sourceSim, targetSim, v1, v10, v12, v13, v2, v5, v6, v7, v8, v9
	if targetSet == sourceSet {
		return
	}
	localIndex = (*b2Joint)(unsafe.Pointer(joint)).LocalIndex
	colorIndex = (*b2Joint)(unsafe.Pointer(joint)).ColorIndex
	if (*b2SolverSet)(unsafe.Pointer(sourceSet)).SetIndex == int32(b2_awakeSet) {
		if v1 = !!(0 <= colorIndex && colorIndex < int32(_B2_GRAPH_COLOR_COUNT)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+5383, __ccgo_ts+26452, int32FromInt32(634))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		color = world + 376 + uintptr(colorIndex)*64
		if v1 = !!(0 <= localIndex && localIndex < (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+27769, __ccgo_ts+26452, int32FromInt32(637))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		sourceSim = (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Data + uintptr(localIndex)*252
	} else {
		if v1 = !!(colorIndex == -int32FromInt32(1)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+28679, __ccgo_ts+26452, int32FromInt32(641))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		if v1 = !!(0 <= localIndex && localIndex < (*b2SolverSet)(unsafe.Pointer(sourceSet)).JointSims.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+28707, __ccgo_ts+26452, int32FromInt32(642))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		sourceSim = (*b2SolverSet)(unsafe.Pointer(sourceSet)).JointSims.Data + uintptr(localIndex)*252
	}
	// Create target and copy. Fix joint.
	if (*b2SolverSet)(unsafe.Pointer(targetSet)).SetIndex == int32(b2_awakeSet) {
		b2AddJointToGraph(tls, world, sourceSim, joint)
		(*b2Joint)(unsafe.Pointer(joint)).SetIndex = int32(b2_awakeSet)
	} else {
		(*b2Joint)(unsafe.Pointer(joint)).SetIndex = (*b2SolverSet)(unsafe.Pointer(targetSet)).SetIndex
		(*b2Joint)(unsafe.Pointer(joint)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(targetSet)).JointSims.Count
		(*b2Joint)(unsafe.Pointer(joint)).ColorIndex = -int32(1)
		v5 = targetSet + 32
		v6 = targetSet + 32 + 8
		v7 = targetSet + 32 + 12
		v8 = int32(252)
		if *(*int32)(unsafe.Add(nil, v6)) >= *(*int32)(unsafe.Add(nil, v7)) {
			oldCapacity = *(*int32)(unsafe.Add(nil, v7))
			oldSize = oldCapacity * v8
			if oldCapacity == 0 {
				v9 = int32(8)
			} else {
				v9 = int32(2) * oldCapacity
			}
			newCapacity = v9
			newSize = newCapacity * v8
			*(*uintptr)(unsafe.Add(nil, v5)) = b2GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v5)), uint64FromInt32(oldSize), uint64FromInt32(newSize))
			*(*int32)(unsafe.Add(nil, v7)) = newCapacity
		}
		v13 = v6
		v12 = *(*int32)(unsafe.Pointer(v13))
		*(*int32)(unsafe.Pointer(v13)) = *(*int32)(unsafe.Pointer(v13)) + 1
		v10 = *(*uintptr)(unsafe.Add(nil, v5)) + uintptr(v12*v8)
		goto _11
	_11:
		targetSim = v10
		memcpy(tls, targetSim, sourceSim, uint64(252))
	}
	// Destroy source.
	if (*b2SolverSet)(unsafe.Pointer(sourceSet)).SetIndex == int32(b2_awakeSet) {
		b2RemoveJointFromGraph(tls, world, (*(*b2JointEdge)(unsafe.Add(nil, joint+20))).BodyId, (*(*b2JointEdge)(unsafe.Add(nil, joint+20+1*12))).BodyId, colorIndex, localIndex)
	} else {
		v5 = (*b2SolverSet)(unsafe.Pointer(sourceSet)).JointSims.Data
		v6 = sourceSet + 32 + 8
		v8 = localIndex
		v9 = int32(252)
		if v1 = !!(0 <= v8 && v8 < *(*int32)(unsafe.Add(nil, v6)) && bool(__ccgo_ts+2029 != 0)); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2055, __ccgo_ts+2115, int32FromInt32(153))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		*(*int32)(unsafe.Add(nil, v6)) = *(*int32)(unsafe.Add(nil, v6)) - 1
		if v8 != *(*int32)(unsafe.Add(nil, v6)) {
			memcpy(tls, v5+uintptr(v8*v9), v5+uintptr(*(*int32)(unsafe.Add(nil, v6))*v9), uint64FromInt32(v9))
			v12 = *(*int32)(unsafe.Add(nil, v6))
			goto _20
		}
		v12 = -int32(1)
		goto _20
	_20:
		movedIndex = v12
		if movedIndex != -int32(1) {
			// fix swapped element
			movedJointSim = (*b2SolverSet)(unsafe.Pointer(sourceSet)).JointSims.Data + uintptr(localIndex)*252
			movedId = (*b2JointSim)(unsafe.Pointer(movedJointSim)).JointId
			if v2 = !!(0 <= movedId && movedId < (*b2World)(unsafe.Pointer(world)).Joints.Count); !v2 {
				b2InternalAssert(tls, __ccgo_ts+5612, __ccgo_ts+26452, int32FromInt32(674))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
			movedJoint = (*b2World)(unsafe.Pointer(world)).Joints.Data + uintptr(movedId)*72
			(*b2Joint)(unsafe.Pointer(movedJoint)).LocalIndex = localIndex
		}
	}
}

/** @} */

func b2WeldJoint_SetLinearHertz(tls *_Stack, jointId JointId, hertz float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_WeldJointSetLinearHertz
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, hertz)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8368, uintptr(unsafe.Pointer(&__func__77))))
		if v1 = !!(b2IsValidFloat(tls, hertz) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8368, __ccgo_ts+28909, int32FromInt32(57))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WeldJointSetLinearHertz)(unsafe.Add(nil, bp)) = b2RecArgs_WeldJointSetLinearHertz{
			Joint: jointId,
			Hertz: hertz,
		}
		b2RecWrite_WeldJointSetLinearHertz(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	if v1 = !!(b2IsValidFloat(tls, hertz) != 0 && hertz >= float32FromFloat32(0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+15481, __ccgo_ts+28909, int32FromInt32(61))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_weldJoint))
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LinearHertz = hertz
}

func b2WeldJoint_GetLinearHertz(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_weldJoint))
	return (*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LinearHertz
}

func b2WeldJoint_SetLinearDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_WeldJointSetLinearDampingRatio
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, dampingRatio)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8392, uintptr(unsafe.Pointer(&__func__78))))
		if v1 = !!(b2IsValidFloat(tls, dampingRatio) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8392, __ccgo_ts+28909, int32FromInt32(74))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WeldJointSetLinearDampingRatio)(unsafe.Add(nil, bp)) = b2RecArgs_WeldJointSetLinearDampingRatio{
			Joint:        jointId,
			DampingRatio: dampingRatio,
		}
		b2RecWrite_WeldJointSetLinearDampingRatio(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	if v1 = !!(b2IsValidFloat(tls, dampingRatio) != 0 && dampingRatio >= float32FromFloat32(0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+15522, __ccgo_ts+28909, int32FromInt32(78))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_weldJoint))
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LinearDampingRatio = dampingRatio
}

func b2WeldJoint_GetLinearDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_weldJoint))
	return (*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LinearDampingRatio
}

func b2WeldJoint_SetAngularHertz(tls *_Stack, jointId JointId, hertz float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_WeldJointSetAngularHertz
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, hertz)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8368, uintptr(unsafe.Pointer(&__func__79))))
		if v1 = !!(b2IsValidFloat(tls, hertz) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8368, __ccgo_ts+28909, int32FromInt32(91))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WeldJointSetAngularHertz)(unsafe.Add(nil, bp)) = b2RecArgs_WeldJointSetAngularHertz{
			Joint: jointId,
			Hertz: hertz,
		}
		b2RecWrite_WeldJointSetAngularHertz(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	if v1 = !!(b2IsValidFloat(tls, hertz) != 0 && hertz >= float32FromFloat32(0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+15481, __ccgo_ts+28909, int32FromInt32(95))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_weldJoint))
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).AngularHertz = hertz
}

func b2WeldJoint_GetAngularHertz(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_weldJoint))
	return (*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).AngularHertz
}

func b2WeldJoint_SetAngularDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_WeldJointSetAngularDampingRatio
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, dampingRatio)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8392, uintptr(unsafe.Pointer(&__func__80))))
		if v1 = !!(b2IsValidFloat(tls, dampingRatio) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8392, __ccgo_ts+28909, int32FromInt32(108))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WeldJointSetAngularDampingRatio)(unsafe.Add(nil, bp)) = b2RecArgs_WeldJointSetAngularDampingRatio{
			Joint:        jointId,
			DampingRatio: dampingRatio,
		}
		b2RecWrite_WeldJointSetAngularDampingRatio(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	if v1 = !!(b2IsValidFloat(tls, dampingRatio) != 0 && dampingRatio >= float32FromFloat32(0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+15522, __ccgo_ts+28909, int32FromInt32(112))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_weldJoint))
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).AngularDampingRatio = dampingRatio
}

func b2WeldJoint_GetAngularDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_weldJoint))
	return (*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).AngularDampingRatio
}

func b2GetWeldJointForce(tls *_Stack, world uintptr, base uintptr) (r Vec2) {
	var force, v2, v3 Vec2
	var v1 float32
	_, _, _, _ = force, v1, v2, v3
	v1 = (*b2World)(unsafe.Pointer(world)).Inv_h
	v2 = (*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(base), 92))).LinearImpulse
	v3 = Vec2{
		X: float32(v1 * v2.X),
		Y: float32(v1 * v2.Y),
	}
	goto _4
_4:
	force = v3
	return force
}

func b2GetWeldJointTorque(tls *_Stack, world uintptr, base uintptr) (r float32) {
	return float32((*b2World)(unsafe.Pointer(world)).Inv_h * (*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(base), 92))).AngularImpulse)
}

// Point-to-point constraint
// C = p2 - p1
// Cdot = v2 - v1
//      = v2 + cross(w2, r2) - v1 - cross(w1, r1)
// J = [-E -r1_skew E r2_skew ]
// Identity used:
// w k % (rx i + ry j) = w * (-ry i + rx j)

// Angle constraint
// C = angle2 - angle1 - referenceAngle
// Cdot = w2 - w1
// J = [0 0 -1 0 0 1]
// K = invI1 + invI2

// 3x3 Block
// K = [J1] * invM * [J1T J2T]
//     [J2]
//   = [J1] * [invM * J1T invM * J2T]
//     [J2]
//   = [J1 * invM * J1T J1 * invM * J2T]
//     [J2 * invM * J1T J2 * invM * J2T]

func b2PrepareWeldJoint(tls *_Stack, base uintptr, context uintptr) {
	var a11, a2, a3, iA, iB, ka, mA, mB, omega, v39, v40 float32
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world uintptr
	var idA, idB, localIndexA, localIndexB, v9 int32
	var qr, v11, v12, v13 Rot
	var v1, v2, v3, v4 bool
	var v15, v16, v17, v20, v21 Vec2
	var v35, v36 Pos
	var v42 b2Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a2, a3, bodyA, bodyB, bodySimA, bodySimB, iA, iB, idA, idB, joint, ka, localIndexA, localIndexB, mA, mB, omega, qr, setA, setB, world, v1, v11, v12, v13, v15, v16, v17, v2, v20, v21, v3, v35, v36, v39, v4, v40, v42, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_weldJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+28937, __ccgo_ts+28909, int32FromInt32(158))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// chase body id to the solver set where the body lives
	idA = (*b2JointSim)(unsafe.Pointer(base)).BodyIdA
	idB = (*b2JointSim)(unsafe.Pointer(base)).BodyIdB
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	if v1 = !!(0 <= idA && idA < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8507, __ccgo_ts+28909, int32FromInt32(166))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idA)*104
	if v2 = !!(0 <= idB && idB < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+8557, __ccgo_ts+28909, int32FromInt32(167))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idB)*104
	if v3 = !!((*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet)); !v3 {
		b2InternalAssert(tls, __ccgo_ts+8607, __ccgo_ts+28909, int32FromInt32(169))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(bodyA)).SetIndex && (*b2Body)(unsafe.Pointer(bodyA)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8672, __ccgo_ts+28909, int32FromInt32(170))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	setA = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(bodyA)).SetIndex)*88
	if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(bodyB)).SetIndex && (*b2Body)(unsafe.Pointer(bodyB)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+8750, __ccgo_ts+28909, int32FromInt32(171))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	setB = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(bodyB)).SetIndex)*88
	localIndexA = (*b2Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b2Body)(unsafe.Pointer(bodyB)).LocalIndex
	if v3 = !!(0 <= localIndexA && localIndexA < (*b2SolverSet)(unsafe.Pointer(setA)).BodySims.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+8828, __ccgo_ts+28909, int32FromInt32(176))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	bodySimA = (*b2SolverSet)(unsafe.Pointer(setA)).BodySims.Data + uintptr(localIndexA)*96
	if v4 = !!(0 <= localIndexB && localIndexB < (*b2SolverSet)(unsafe.Pointer(setB)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+8895, __ccgo_ts+28909, int32FromInt32(177))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySimB = (*b2SolverSet)(unsafe.Pointer(setB)).BodySims.Data + uintptr(localIndexB)*96
	mA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvMass
	iA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvInertia
	mB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvMass
	iB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvInertia
	(*b2JointSim)(unsafe.Pointer(base)).InvMassA = mA
	(*b2JointSim)(unsafe.Pointer(base)).InvMassB = mB
	(*b2JointSim)(unsafe.Pointer(base)).InvIA = iA
	(*b2JointSim)(unsafe.Pointer(base)).InvIB = iB
	joint = base + 92
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) {
		v9 = localIndexA
	} else {
		v9 = -int32(1)
	}
	(*b2WeldJoint)(unsafe.Pointer(joint)).IndexA = v9
	if (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) {
		v9 = localIndexB
	} else {
		v9 = -int32(1)
	}
	(*b2WeldJoint)(unsafe.Pointer(joint)).IndexB = v9
	// Compute joint anchor frames with world space rotation, relative to center of mass
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v12 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	qr.S = float32(v11.S*v12.C) + float32(v11.C*v12.S)
	qr.C = float32(v11.C*v12.C) - float32(v11.S*v12.S)
	v13 = qr
	goto _14
_14:
	(*b2WeldJoint)(unsafe.Pointer(joint)).FrameA.Q = v13
	v15 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA.P
	v16 = (*b2BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v17 = Vec2{
		X: v15.X - v16.X,
		Y: v15.Y - v16.Y,
	}
	goto _18
_18:
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v20 = v17
	v21 = Vec2{
		X: float32(v11.C*v20.X) - float32(v11.S*v20.Y),
		Y: float32(v11.S*v20.X) + float32(v11.C*v20.Y),
	}
	goto _22
_22:
	(*b2WeldJoint)(unsafe.Pointer(joint)).FrameA.P = v21
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v12 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameB.Q
	qr.S = float32(v11.S*v12.C) + float32(v11.C*v12.S)
	qr.C = float32(v11.C*v12.C) - float32(v11.S*v12.S)
	v13 = qr
	goto _26
_26:
	(*b2WeldJoint)(unsafe.Pointer(joint)).FrameB.Q = v13
	v15 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v16 = (*b2BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v17 = Vec2{
		X: v15.X - v16.X,
		Y: v15.Y - v16.Y,
	}
	goto _30
_30:
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v20 = v17
	v21 = Vec2{
		X: float32(v11.C*v20.X) - float32(v11.S*v20.Y),
		Y: float32(v11.S*v20.X) + float32(v11.C*v20.Y),
	}
	goto _34
_34:
	(*b2WeldJoint)(unsafe.Pointer(joint)).FrameB.P = v21
	// Compute the initial center delta. Incremental position updates are relative to this.
	v35 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Center
	v36 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Center
	v15 = Vec2{
		X: v35.X - v36.X,
		Y: v35.Y - v36.Y,
	}
	goto _38
_38:
	(*b2WeldJoint)(unsafe.Pointer(joint)).DeltaCenter = v15
	ka = iA + iB
	if ka > float32FromFloat32(0) {
		v39 = float32FromFloat32(1) / ka
	} else {
		v39 = float32FromFloat32(0)
	}
	(*b2WeldJoint)(unsafe.Pointer(joint)).AxialMass = v39
	if (*b2WeldJoint)(unsafe.Pointer(joint)).LinearHertz == float32FromFloat32(0) {
		(*b2WeldJoint)(unsafe.Pointer(joint)).LinearSpring = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness
	} else {
		v39 = (*b2WeldJoint)(unsafe.Pointer(joint)).LinearHertz
		v40 = (*b2StepContext)(unsafe.Pointer(context)).H
		if v39 == float32FromFloat32(0) {
			v42 = b2Softness{}
			goto _43
		}
		omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(_B2_PI)) * v39)
		a11 = float32(float32FromFloat32(2)*(*b2WeldJoint)(unsafe.Pointer(joint)).LinearDampingRatio) + float32(v40*omega)
		a2 = float32(float32(v40*omega) * a11)
		a3 = float32FromFloat32(1) / (float32FromFloat32(1) + a2)
		v42 = b2Softness{
			BiasRate:     omega / a11,
			MassScale:    float32(a2 * a3),
			ImpulseScale: a3,
		}
		goto _43
	_43:
		(*b2WeldJoint)(unsafe.Pointer(joint)).LinearSpring = v42
	}
	if (*b2WeldJoint)(unsafe.Pointer(joint)).AngularHertz == float32FromFloat32(0) {
		(*b2WeldJoint)(unsafe.Pointer(joint)).AngularSpring = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness
	} else {
		v39 = (*b2WeldJoint)(unsafe.Pointer(joint)).AngularHertz
		v40 = (*b2StepContext)(unsafe.Pointer(context)).H
		if v39 == float32FromFloat32(0) {
			v42 = b2Softness{}
			goto _47
		}
		omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(_B2_PI)) * v39)
		a11 = float32(float32FromFloat32(2)*(*b2WeldJoint)(unsafe.Pointer(joint)).AngularDampingRatio) + float32(v40*omega)
		a2 = float32(float32(v40*omega) * a11)
		a3 = float32FromFloat32(1) / (float32FromFloat32(1) + a2)
		v42 = b2Softness{
			BiasRate:     omega / a11,
			MassScale:    float32(a2 * a3),
			ImpulseScale: a3,
		}
		goto _47
	_47:
		(*b2WeldJoint)(unsafe.Pointer(joint)).AngularSpring = v42
	}
	if int32FromUint8((*b2StepContext)(unsafe.Pointer(context)).EnableWarmStarting) == false1 {
		(*b2WeldJoint)(unsafe.Pointer(joint)).LinearImpulse = b2Vec2_zero
		(*b2WeldJoint)(unsafe.Pointer(joint)).AngularImpulse = float32FromFloat32(0)
	}
}

func b2WarmStartWeldJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var iA, iB, mA, mB, v12 float32
	var joint, stateA, stateB, v1, v2 uintptr
	var rA, rB, v11, v13, v14, v4, v5, v8, v9 Vec2
	var v3, v7 Rot
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = iA, iB, joint, mA, mB, rA, rB, stateA, stateB, v1, v11, v12, v13, v14, v2, v3, v4, v5, v7, v8, v9
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Add(nil, bp)) = b2_identityBodyState22
	joint = base + 92
	if (*b2WeldJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v1 = bp
	} else {
		v1 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2WeldJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v1
	if (*b2WeldJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2WeldJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v2
	v3 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v4 = (*b2WeldJoint)(unsafe.Pointer(joint)).FrameA.P
	v5 = Vec2{
		X: float32(v3.C*v4.X) - float32(v3.S*v4.Y),
		Y: float32(v3.S*v4.X) + float32(v3.C*v4.Y),
	}
	goto _6
_6:
	rA = v5
	v7 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v8 = (*b2WeldJoint)(unsafe.Pointer(joint)).FrameB.P
	v9 = Vec2{
		X: float32(v7.C*v8.X) - float32(v7.S*v8.Y),
		Y: float32(v7.S*v8.X) + float32(v7.C*v8.Y),
	}
	goto _10
_10:
	rB = v9
	if (*b2BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b2_dynamicFlag) != 0 {
		v11 = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
		v12 = mA
		v13 = (*b2WeldJoint)(unsafe.Pointer(joint)).LinearImpulse
		v14 = Vec2{
			X: v11.X - float32(v12*v13.X),
			Y: v11.Y - float32(v12*v13.Y),
		}
		goto _15
	_15:
		(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = v14
		v4 = rA
		v5 = (*b2WeldJoint)(unsafe.Pointer(joint)).LinearImpulse
		v12 = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
		goto _19
	_19:
		*(*float32)(unsafe.Add(nil, stateA+8)) -= float32(iA * (v12 + (*b2WeldJoint)(unsafe.Pointer(joint)).AngularImpulse))
	}
	if (*b2BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b2_dynamicFlag) != 0 {
		v4 = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
		v12 = mB
		v5 = (*b2WeldJoint)(unsafe.Pointer(joint)).LinearImpulse
		v8 = Vec2{
			X: v4.X + float32(v12*v5.X),
			Y: v4.Y + float32(v12*v5.Y),
		}
		goto _24
	_24:
		(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = v8
		v4 = rB
		v5 = (*b2WeldJoint)(unsafe.Pointer(joint)).LinearImpulse
		v12 = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
		goto _28
	_28:
		*(*float32)(unsafe.Add(nil, stateB+8)) += float32(iB * (v12 + (*b2WeldJoint)(unsafe.Pointer(joint)).AngularImpulse))
	}
}

func b2SolveWeldJoint(tls *_Stack, base uintptr, context uintptr, useBias uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var C, Cdot, a11, a12, a21, a22, bias, det, iA, iB, impulse, impulseScale, impulseScale1, jointAngle, mA, mB, massScale, massScale1, wA, wB, v17, v43 float32
	var C1, Cdot1, b7, bias1, dcA, dcB, impulse1, rA, rB, vA, vB, x, v20, v21, v24, v25, v27, v28, v29, v31, v32, v33, v35, v36, v37, v39, v40, v41, v44, v45 Vec2
	var K, v71 Mat22
	var joint, stateA, stateB, v2, v3 uintptr
	var qA, qB, qr, r1, relQ, v10, v12, v13, v14, v16, v4, v5, v6, v8, v9 Rot
	var v1 bool
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, C1, Cdot, Cdot1, K, a11, a12, a21, a22, b7, bias, bias1, dcA, dcB, det, iA, iB, impulse, impulse1, impulseScale, impulseScale1, joint, jointAngle, mA, mB, massScale, massScale1, qA, qB, qr, r1, rA, rB, relQ, stateA, stateB, vA, vB, wA, wB, x, v1, v10, v12, v13, v14, v16, v17, v2, v20, v21, v24, v25, v27, v28, v29, v3, v31, v32, v33, v35, v36, v37, v39, v4, v40, v41, v43, v44, v45, v5, v6, v71, v8, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_weldJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+28937, __ccgo_ts+28909, int32FromInt32(263))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Add(nil, bp)) = b2_identityBodyState22
	joint = base + 92
	if (*b2WeldJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2WeldJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v2
	if (*b2WeldJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2WeldJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v3
	vA = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	// Block solve doesn't work correctly with mixed stiffness values
	// angular constraint
	v4 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b2WeldJoint)(unsafe.Pointer(joint)).FrameA.Q
	qr.S = float32(v4.S*v5.C) + float32(v4.C*v5.S)
	qr.C = float32(v4.C*v5.C) - float32(v4.S*v5.S)
	v6 = qr
	goto _7
_7:
	qA = v6
	v8 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v9 = (*b2WeldJoint)(unsafe.Pointer(joint)).FrameB.Q
	qr.S = float32(v8.S*v9.C) + float32(v8.C*v9.S)
	qr.C = float32(v8.C*v9.C) - float32(v8.S*v9.S)
	v10 = qr
	goto _11
_11:
	qB = v10
	v12 = qA
	v13 = qB
	r1.S = float32(v12.C*v13.S) - float32(v12.S*v13.C)
	r1.C = float32(v12.C*v13.C) + float32(v12.S*v13.S)
	v14 = r1
	goto _15
_15:
	relQ = v14
	v16 = relQ
	v17 = b2Atan2(tls, v16.S, v16.C)
	goto _18
_18:
	jointAngle = v17
	bias = float32FromFloat32(0)
	massScale = float32FromFloat32(1)
	impulseScale = float32FromFloat32(0)
	if useBias != 0 || (*b2WeldJoint)(unsafe.Pointer(joint)).AngularHertz > float32FromFloat32(0) {
		C = jointAngle
		bias = float32((*b2WeldJoint)(unsafe.Pointer(joint)).AngularSpring.BiasRate * C)
		massScale = (*b2WeldJoint)(unsafe.Pointer(joint)).AngularSpring.MassScale
		impulseScale = (*b2WeldJoint)(unsafe.Pointer(joint)).AngularSpring.ImpulseScale
	}
	Cdot = wB - wA
	impulse = float32(float32(-massScale*(*b2WeldJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot+bias)) - float32(impulseScale*(*b2WeldJoint)(unsafe.Pointer(joint)).AngularImpulse)
	*(*float32)(unsafe.Add(nil, joint+48)) += impulse
	wA = wA - float32(iA*impulse)
	wB = wB + float32(iB*impulse)
	// linear constraint
	v4 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v20 = (*b2WeldJoint)(unsafe.Pointer(joint)).FrameA.P
	v21 = Vec2{
		X: float32(v4.C*v20.X) - float32(v4.S*v20.Y),
		Y: float32(v4.S*v20.X) + float32(v4.C*v20.Y),
	}
	goto _22
_22:
	rA = v21
	v5 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v24 = (*b2WeldJoint)(unsafe.Pointer(joint)).FrameB.P
	v25 = Vec2{
		X: float32(v5.C*v24.X) - float32(v5.S*v24.Y),
		Y: float32(v5.S*v24.X) + float32(v5.C*v24.Y),
	}
	goto _26
_26:
	rB = v25
	bias1 = b2Vec2_zero
	massScale1 = float32FromFloat32(1)
	impulseScale1 = float32FromFloat32(0)
	if useBias != 0 || (*b2WeldJoint)(unsafe.Pointer(joint)).LinearHertz > float32FromFloat32(0) {
		dcA = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaPosition
		dcB = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
		v27 = dcB
		v28 = dcA
		v29 = Vec2{
			X: v27.X - v28.X,
			Y: v27.Y - v28.Y,
		}
		goto _30
	_30:
		v31 = rB
		v32 = rA
		v33 = Vec2{
			X: v31.X - v32.X,
			Y: v31.Y - v32.Y,
		}
		goto _34
	_34:
		v35 = v29
		v36 = v33
		v37 = Vec2{
			X: v35.X + v36.X,
			Y: v35.Y + v36.Y,
		}
		goto _38
	_38:
		v39 = v37
		v40 = (*b2WeldJoint)(unsafe.Pointer(joint)).DeltaCenter
		v41 = Vec2{
			X: v39.X + v40.X,
			Y: v39.Y + v40.Y,
		}
		goto _42
	_42:
		C1 = v41
		v17 = (*b2WeldJoint)(unsafe.Pointer(joint)).LinearSpring.BiasRate
		v44 = C1
		v45 = Vec2{
			X: float32(v17 * v44.X),
			Y: float32(v17 * v44.Y),
		}
		goto _46
	_46:
		bias1 = v45
		massScale1 = (*b2WeldJoint)(unsafe.Pointer(joint)).LinearSpring.MassScale
		impulseScale1 = (*b2WeldJoint)(unsafe.Pointer(joint)).LinearSpring.ImpulseScale
	}
	v17 = wB
	v20 = rB
	v21 = Vec2{
		X: float32(-v17 * v20.Y),
		Y: float32(v17 * v20.X),
	}
	goto _50
_50:
	v24 = vB
	v25 = v21
	v27 = Vec2{
		X: v24.X + v25.X,
		Y: v24.Y + v25.Y,
	}
	goto _54
_54:
	v43 = wA
	v28 = rA
	v29 = Vec2{
		X: float32(-v43 * v28.Y),
		Y: float32(v43 * v28.X),
	}
	goto _58
_58:
	v31 = vA
	v32 = v29
	v33 = Vec2{
		X: v31.X + v32.X,
		Y: v31.Y + v32.Y,
	}
	goto _62
_62:
	v35 = v27
	v36 = v33
	v37 = Vec2{
		X: v35.X - v36.X,
		Y: v35.Y - v36.Y,
	}
	goto _66
_66:
	Cdot1 = v37
	K.Cx.X = mA + mB + float32(float32(rA.Y*rA.Y)*iA) + float32(float32(rB.Y*rB.Y)*iB)
	K.Cy.X = float32(float32(-rA.Y*rA.X)*iA) - float32(float32(rB.Y*rB.X)*iB)
	K.Cx.Y = K.Cy.X
	K.Cy.Y = mA + mB + float32(float32(rA.X*rA.X)*iA) + float32(float32(rB.X*rB.X)*iB)
	v20 = Cdot1
	v21 = bias1
	v24 = Vec2{
		X: v20.X + v21.X,
		Y: v20.Y + v21.Y,
	}
	goto _70
_70:
	v71 = K
	v25 = v24
	a11 = v71.Cx.X
	a12 = v71.Cy.X
	a21 = v71.Cx.Y
	a22 = v71.Cy.Y
	det = float32(a11*a22) - float32(a12*a21)
	if det != float32FromFloat32(0) {
		det = float32FromFloat32(1) / det
	}
	x = Vec2{
		X: float32(det * (float32(a22*v25.X) - float32(a12*v25.Y))),
		Y: float32(det * (float32(a11*v25.Y) - float32(a21*v25.X))),
	}
	v27 = x
	goto _74
_74:
	b7 = v27
	impulse1 = Vec2{
		X: float32(-massScale1*b7.X) - float32(impulseScale1*(*b2WeldJoint)(unsafe.Pointer(joint)).LinearImpulse.X),
		Y: float32(-massScale1*b7.Y) - float32(impulseScale1*(*b2WeldJoint)(unsafe.Pointer(joint)).LinearImpulse.Y),
	}
	v28 = (*b2WeldJoint)(unsafe.Pointer(joint)).LinearImpulse
	v29 = impulse1
	v31 = Vec2{
		X: v28.X + v29.X,
		Y: v28.Y + v29.Y,
	}
	goto _78
_78:
	(*b2WeldJoint)(unsafe.Pointer(joint)).LinearImpulse = v31
	v20 = vA
	v17 = mA
	v21 = impulse1
	v24 = Vec2{
		X: v20.X - float32(v17*v21.X),
		Y: v20.Y - float32(v17*v21.Y),
	}
	goto _83
_83:
	vA = v24
	v20 = rA
	v21 = impulse1
	v17 = float32(v20.X*v21.Y) - float32(v20.Y*v21.X)
	goto _87
_87:
	wA = wA - float32(iA*v17)
	v20 = vB
	v17 = mB
	v21 = impulse1
	v24 = Vec2{
		X: v20.X + float32(v17*v21.X),
		Y: v20.Y + float32(v17*v21.Y),
	}
	goto _92
_92:
	vB = v24
	v20 = rB
	v21 = impulse1
	v17 = float32(v20.X*v21.Y) - float32(v20.Y*v21.X)
	goto _96
_96:
	wB = wB + float32(iB*v17)
	if v1 = !!(b2IsValidVec2(tls, vA) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20553, __ccgo_ts+28909, int32FromInt32(445))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b2IsValidFloat(tls, wA) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20573, __ccgo_ts+28909, int32FromInt32(446))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b2IsValidVec2(tls, vB) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20594, __ccgo_ts+28909, int32FromInt32(447))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(b2IsValidFloat(tls, wB) != 0); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20614, __ccgo_ts+28909, int32FromInt32(448))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b2BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b2_dynamicFlag) != 0 {
		(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
		(*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	}
	if (*b2BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b2_dynamicFlag) != 0 {
		(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
		(*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
	}
}

func b2DrawWeldJoint(tls *_Stack, draw uintptr, base uintptr, transformA WorldTransform, transformB WorldTransform, drawScale float32) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var C, frameA, frameB, v16, v18, v2, v32 WorldTransform
	var qr, v20, v21, v22, v24, v4, v5, v6, v8 Rot
	var v1 bool
	var v10, v13, v25, v26, v29, v9 Vec2
	var v12, v14, v28, v30 Pos
	var v19, v3 Transform
	var _ /* box at bp+0 */ Polygon
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, frameA, frameB, qr, v1, v10, v12, v13, v14, v16, v18, v19, v2, v20, v21, v22, v24, v25, v26, v28, v29, v3, v30, v32, v4, v5, v6, v8, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_weldJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+28937, __ccgo_ts+28909, int32FromInt32(466))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	v2 = transformA
	v3 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA
	v4 = v2.Q
	v5 = v3.Q
	qr.S = float32(v4.S*v5.C) + float32(v4.C*v5.S)
	qr.C = float32(v4.C*v5.C) - float32(v4.S*v5.S)
	v6 = qr
	goto _7
_7:
	C.Q = v6
	v8 = v2.Q
	v9 = v3.P
	v10 = Vec2{
		X: float32(v8.C*v9.X) - float32(v8.S*v9.Y),
		Y: float32(v8.S*v9.X) + float32(v8.C*v9.Y),
	}
	goto _11
_11:
	v12 = v2.P
	v13 = v10
	v14 = Pos{
		X: v12.X + v13.X,
		Y: v12.Y + v13.Y,
	}
	goto _15
_15:
	C.P = v14
	v16 = C
	goto _17
_17:
	frameA = v16
	v18 = transformB
	v19 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameB
	v20 = v18.Q
	v21 = v19.Q
	qr.S = float32(v20.S*v21.C) + float32(v20.C*v21.S)
	qr.C = float32(v20.C*v21.C) - float32(v20.S*v21.S)
	v22 = qr
	goto _23
_23:
	C.Q = v22
	v24 = v18.Q
	v25 = v19.P
	v26 = Vec2{
		X: float32(v24.C*v25.X) - float32(v24.S*v25.Y),
		Y: float32(v24.S*v25.X) + float32(v24.C*v25.Y),
	}
	goto _27
_27:
	v28 = v18.P
	v29 = v26
	v30 = Pos{
		X: v28.X + v29.X,
		Y: v28.Y + v29.Y,
	}
	goto _31
_31:
	C.P = v30
	v32 = C
	goto _33
_33:
	frameB = v32
	*(*Polygon)(unsafe.Add(nil, bp)) = b2MakeBox(tls, float32(float32FromFloat32(0.25)*drawScale), float32(float32FromFloat32(0.125)*drawScale))
	(*(*func(*_Stack, WorldTransform, uintptr, int32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPolygonFcn})))(tls, frameA, bp, int32(4), int32(b2_colorDarkOrange), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	(*(*func(*_Stack, WorldTransform, uintptr, int32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPolygonFcn})))(tls, frameB, bp, int32(4), int32(b2_colorDarkCyan), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
}

/** @} */

func b2WheelJoint_EnableSpring(tls *_Stack, jointId JointId, enableSpring uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* _a at bp+0 */ b2RecArgs_WheelJointEnableSpring
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WheelJointEnableSpring)(unsafe.Add(nil, bp)) = b2RecArgs_WheelJointEnableSpring{
			Joint:        jointId,
			EnableSpring: enableSpring,
		}
		b2RecWrite_WheelJointEnableSpring(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	if int32FromUint8(enableSpring) != int32FromUint8((*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableSpring) {
		(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableSpring = enableSpring
		(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).SpringImpulse = float32FromFloat32(0)
	}
}

func b2WheelJoint_IsSpringEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	return (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableSpring
}

func b2WheelJoint_SetSpringHertz(tls *_Stack, jointId JointId, hertz float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_WheelJointSetSpringHertz
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, hertz)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8368, uintptr(unsafe.Pointer(&__func__82))))
		if v1 = !!(b2IsValidFloat(tls, hertz) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8368, __ccgo_ts+28964, int32FromInt32(36))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WheelJointSetSpringHertz)(unsafe.Add(nil, bp)) = b2RecArgs_WheelJointSetSpringHertz{
			Joint: jointId,
			Hertz: hertz,
		}
		b2RecWrite_WheelJointSetSpringHertz(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).Hertz = hertz
}

func b2WheelJoint_GetSpringHertz(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	return (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).Hertz
}

func b2WheelJoint_SetSpringDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_WheelJointSetSpringDampingRatio
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, dampingRatio)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8392, uintptr(unsafe.Pointer(&__func__83))))
		if v1 = !!(b2IsValidFloat(tls, dampingRatio) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8392, __ccgo_ts+28964, int32FromInt32(52))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WheelJointSetSpringDampingRatio)(unsafe.Add(nil, bp)) = b2RecArgs_WheelJointSetSpringDampingRatio{
			Joint:        jointId,
			DampingRatio: dampingRatio,
		}
		b2RecWrite_WheelJointSetSpringDampingRatio(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).DampingRatio = dampingRatio
}

func b2WheelJoint_GetSpringDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	return (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).DampingRatio
}

func b2WheelJoint_EnableLimit(tls *_Stack, jointId JointId, enableLimit uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* _a at bp+0 */ b2RecArgs_WheelJointEnableLimit
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WheelJointEnableLimit)(unsafe.Add(nil, bp)) = b2RecArgs_WheelJointEnableLimit{
			Joint:       jointId,
			EnableLimit: enableLimit,
		}
		b2RecWrite_WheelJointEnableLimit(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	if int32FromUint8((*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableLimit) != int32FromUint8(enableLimit) {
		(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LowerImpulse = float32FromFloat32(0)
		(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).UpperImpulse = float32FromFloat32(0)
		(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableLimit = enableLimit
	}
}

func b2WheelJoint_IsLimitEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	return (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableLimit
}

func b2WheelJoint_GetLowerLimit(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	return (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LowerTranslation
}

func b2WheelJoint_GetUpperLimit(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	return (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).UpperTranslation
}

func b2WheelJoint_SetLimits(tls *_Stack, jointId JointId, lower float32, upper float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var v4, v5, v6, v8 float32
	var _ /* _a at bp+0 */ b2RecArgs_WheelJointSetLimits
	_, _, _, _, _, _, _ = joint, world, v1, v4, v5, v6, v8
	if int32FromUint8(b2IsValidFloat(tls, lower)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+20285, uintptr(unsafe.Pointer(&__func__211))))
		if v1 = !!(b2IsValidFloat(tls, lower) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+20285, __ccgo_ts+28964, int32FromInt32(99))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	if int32FromUint8(b2IsValidFloat(tls, upper)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+20309, uintptr(unsafe.Pointer(&__func__211))))
		if v1 = !!(b2IsValidFloat(tls, upper) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+20309, __ccgo_ts+28964, int32FromInt32(100))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WheelJointSetLimits)(unsafe.Add(nil, bp)) = b2RecArgs_WheelJointSetLimits{
			Joint: jointId,
			Lower: lower,
			Upper: upper,
		}
		b2RecWrite_WheelJointSetLimits(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	if v1 = !!(lower <= upper); !v1 {
		b2InternalAssert(tls, __ccgo_ts+20333, __ccgo_ts+28964, int32FromInt32(104))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	v4 = lower
	v5 = upper
	if v4 < v5 {
		v8 = v4
	} else {
		v8 = v5
	}
	v6 = v8
	goto _7
_7:
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).LowerTranslation = v6
	v4 = lower
	v5 = upper
	if v4 > v5 {
		v8 = v4
	} else {
		v8 = v5
	}
	v6 = v8
	goto _12
_12:
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).UpperTranslation = v6
}

func b2WheelJoint_EnableMotor(tls *_Stack, jointId JointId, enableMotor uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var joint, world uintptr
	var _ /* _a at bp+0 */ b2RecArgs_WheelJointEnableMotor
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WheelJointEnableMotor)(unsafe.Add(nil, bp)) = b2RecArgs_WheelJointEnableMotor{
			Joint:       jointId,
			EnableMotor: enableMotor,
		}
		b2RecWrite_WheelJointEnableMotor(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	if int32FromUint8((*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableMotor) != int32FromUint8(enableMotor) {
		(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MotorImpulse = float32FromFloat32(0)
		(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableMotor = enableMotor
	}
}

func b2WheelJoint_IsMotorEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	return (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).EnableMotor
}

func b2WheelJoint_SetTargetMotorSpeed(tls *_Stack, jointId JointId, motorSpeed float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_WheelJointSetTargetMotorSpeed
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, motorSpeed)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+8423, uintptr(unsafe.Pointer(&__func__84))))
		if v1 = !!(b2IsValidFloat(tls, motorSpeed) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+8423, __ccgo_ts+28964, int32FromInt32(131))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WheelJointSetTargetMotorSpeed)(unsafe.Add(nil, bp)) = b2RecArgs_WheelJointSetTargetMotorSpeed{
			Joint:      jointId,
			MotorSpeed: motorSpeed,
		}
		b2RecWrite_WheelJointSetTargetMotorSpeed(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MotorSpeed = motorSpeed
}

func b2WheelJoint_GetTargetMotorSpeed(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	return (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MotorSpeed
}

func b2WheelJoint_GetCurrentMotorTorque(tls *_Stack, jointId JointId) (r float32) {
	var joint, world uintptr
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	return float32((*b2World)(unsafe.Pointer(world)).Inv_h * (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MotorImpulse)
}

func b2WheelJoint_SetMaxMotorTorque(tls *_Stack, jointId JointId, torque float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var joint, world uintptr
	var v1 bool
	var _ /* _a at bp+0 */ b2RecArgs_WheelJointSetMaxMotorTorque
	_, _, _ = joint, world, v1
	if int32FromUint8(b2IsValidFloat(tls, torque)) == false1 {
		b2Log(tls, __ccgo_ts+1116, vaList(bp+24, __ccgo_ts+2758, uintptr(unsafe.Pointer(&__func__85))))
		if v1 = !!(b2IsValidFloat(tls, torque) != 0); !v1 {
			b2InternalAssert(tls, __ccgo_ts+2758, __ccgo_ts+28964, int32FromInt32(154))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		return
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if (*b2World)(unsafe.Pointer(world)).Recording != uintptrFromInt32(0) {
		*(*b2RecArgs_WheelJointSetMaxMotorTorque)(unsafe.Add(nil, bp)) = b2RecArgs_WheelJointSetMaxMotorTorque{
			Joint:  jointId,
			Torque: torque,
		}
		b2RecWrite_WheelJointSetMaxMotorTorque(tls, (*b2World)(unsafe.Pointer(world)).Recording, bp)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxMotorTorque = torque
}

func b2WheelJoint_GetMaxMotorTorque(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	return (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 92))).MaxMotorTorque
}

func b2GetWheelJointForce(tls *_Stack, world uintptr, base uintptr) (r Vec2) {
	var axialForce, perpForce, v12, v16 float32
	var axisA, force, localAxisA, perpA, v10, v13, v14, v17, v18, v2, v20, v21, v22, v3, v6, v7, v9 Vec2
	var joint uintptr
	var qA, v1, v5 Rot
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axialForce, axisA, force, joint, localAxisA, perpA, perpForce, qA, v1, v10, v12, v13, v14, v16, v17, v18, v2, v20, v21, v22, v3, v5, v6, v7, v9
	qA = b2GetBodyTransform(tls, world, (*b2JointSim)(unsafe.Pointer(base)).BodyIdA).Q
	v1 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	v2 = Vec2{
		X: float32FromFloat32(1),
	}
	v3 = Vec2{
		X: float32(v1.C*v2.X) - float32(v1.S*v2.Y),
		Y: float32(v1.S*v2.X) + float32(v1.C*v2.Y),
	}
	goto _4
_4:
	localAxisA = v3
	v5 = qA
	v6 = localAxisA
	v7 = Vec2{
		X: float32(v5.C*v6.X) - float32(v5.S*v6.Y),
		Y: float32(v5.S*v6.X) + float32(v5.C*v6.Y),
	}
	goto _8
_8:
	axisA = v7
	v9 = axisA
	v10 = Vec2{
		X: -v9.Y,
		Y: v9.X,
	}
	goto _11
_11:
	perpA = v10
	joint = base + 92
	perpForce = float32((*b2World)(unsafe.Pointer(world)).Inv_h * (*b2WheelJoint)(unsafe.Pointer(joint)).PerpImpulse)
	axialForce = float32((*b2World)(unsafe.Pointer(world)).Inv_h * ((*b2WheelJoint)(unsafe.Pointer(joint)).SpringImpulse + (*b2WheelJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b2WheelJoint)(unsafe.Pointer(joint)).UpperImpulse))
	v12 = perpForce
	v13 = perpA
	v14 = Vec2{
		X: float32(v12 * v13.X),
		Y: float32(v12 * v13.Y),
	}
	goto _15
_15:
	v16 = axialForce
	v17 = axisA
	v18 = Vec2{
		X: float32(v16 * v17.X),
		Y: float32(v16 * v17.Y),
	}
	goto _19
_19:
	v20 = v14
	v21 = v18
	v22 = Vec2{
		X: v20.X + v21.X,
		Y: v20.Y + v21.Y,
	}
	goto _23
_23:
	force = v22
	return force
}

func b2GetWheelJointTorque(tls *_Stack, world uintptr, base uintptr) (r float32) {
	return float32((*b2World)(unsafe.Pointer(world)).Inv_h * (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(base), 92))).MotorImpulse)
}

// Linear constraint (point-to-line)
// d = pB - pA = xB + rB - xA - rA
// C = dot(ay, d)
// Cdot = dot(d, cross(wA, ay)) + dot(ay, vB + cross(wB, rB) - vA - cross(wA, rA))
//      = -dot(ay, vA) - dot(cross(d + rA, ay), wA) + dot(ay, vB) + dot(cross(rB, ay), vB)
// J = [-ay, -cross(d + rA, ay), ay, cross(rB, ay)]

// Spring linear constraint
// C = dot(ax, d)
// Cdot = = -dot(ax, vA) - dot(cross(d + rA, ax), wA) + dot(ax, vB) + dot(cross(rB, ax), vB)
// J = [-ax -cross(d+rA, ax) ax cross(rB, ax)]

// Motor rotational constraint
// Cdot = wB - wA
// J = [0 0 -1 0 0 1]

func b2PrepareWheelJoint(tls *_Stack, base uintptr, context uintptr) {
	var a11, a12, a21, a22, a31, iA, iB, ka, km, kp, mA, mB, omega, s1, s2, v60, v64, v66 float32
	var axisA, d, perpA, rA, rB, v15, v16, v17, v20, v21, v27, v28, v29, v32, v33, v37, v39, v40, v41, v43, v44, v45 Vec2
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world uintptr
	var idA, idB, localIndexA, localIndexB, v9 int32
	var qr, v11, v12, v13 Rot
	var v1, v2, v3, v4 bool
	var v35, v36 Pos
	var v82 b2Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a12, a21, a22, a31, axisA, bodyA, bodyB, bodySimA, bodySimB, d, iA, iB, idA, idB, joint, ka, km, kp, localIndexA, localIndexB, mA, mB, omega, perpA, qr, rA, rB, s1, s2, setA, setB, world, v1, v11, v12, v13, v15, v16, v17, v2, v20, v21, v27, v28, v29, v3, v32, v33, v35, v36, v37, v39, v4, v40, v41, v43, v44, v45, v60, v64, v66, v82, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_wheelJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+28993, __ccgo_ts+28964, int32FromInt32(208))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// chase body id to the solver set where the body lives
	idA = (*b2JointSim)(unsafe.Pointer(base)).BodyIdA
	idB = (*b2JointSim)(unsafe.Pointer(base)).BodyIdB
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	if v1 = !!(0 <= idA && idA < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8507, __ccgo_ts+28964, int32FromInt32(216))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	bodyA = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idA)*104
	if v2 = !!(0 <= idB && idB < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+8557, __ccgo_ts+28964, int32FromInt32(217))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	bodyB = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr(idB)*104
	if v3 = !!((*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet)); !v3 {
		b2InternalAssert(tls, __ccgo_ts+8607, __ccgo_ts+28964, int32FromInt32(219))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= (*b2Body)(unsafe.Pointer(bodyA)).SetIndex && (*b2Body)(unsafe.Pointer(bodyA)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+8672, __ccgo_ts+28964, int32FromInt32(220))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	setA = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(bodyA)).SetIndex)*88
	if v2 = !!(0 <= (*b2Body)(unsafe.Pointer(bodyB)).SetIndex && (*b2Body)(unsafe.Pointer(bodyB)).SetIndex < (*b2World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b2InternalAssert(tls, __ccgo_ts+8750, __ccgo_ts+28964, int32FromInt32(221))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	setB = (*b2World)(unsafe.Pointer(world)).SolverSets.Data + uintptr((*b2Body)(unsafe.Pointer(bodyB)).SetIndex)*88
	localIndexA = (*b2Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b2Body)(unsafe.Pointer(bodyB)).LocalIndex
	if v3 = !!(0 <= localIndexA && localIndexA < (*b2SolverSet)(unsafe.Pointer(setA)).BodySims.Count); !v3 {
		b2InternalAssert(tls, __ccgo_ts+8828, __ccgo_ts+28964, int32FromInt32(226))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	bodySimA = (*b2SolverSet)(unsafe.Pointer(setA)).BodySims.Data + uintptr(localIndexA)*96
	if v4 = !!(0 <= localIndexB && localIndexB < (*b2SolverSet)(unsafe.Pointer(setB)).BodySims.Count); !v4 {
		b2InternalAssert(tls, __ccgo_ts+8895, __ccgo_ts+28964, int32FromInt32(227))
	}
	_ = v4 || bool(int32FromInt32(0) != 0)
	bodySimB = (*b2SolverSet)(unsafe.Pointer(setB)).BodySims.Data + uintptr(localIndexB)*96
	mA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvMass
	iA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvInertia
	mB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvMass
	iB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvInertia
	(*b2JointSim)(unsafe.Pointer(base)).InvMassA = mA
	(*b2JointSim)(unsafe.Pointer(base)).InvMassB = mB
	(*b2JointSim)(unsafe.Pointer(base)).InvIA = iA
	(*b2JointSim)(unsafe.Pointer(base)).InvIB = iB
	joint = base + 92
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) {
		v9 = localIndexA
	} else {
		v9 = -int32(1)
	}
	(*b2WheelJoint)(unsafe.Pointer(joint)).IndexA = v9
	if (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) {
		v9 = localIndexB
	} else {
		v9 = -int32(1)
	}
	(*b2WheelJoint)(unsafe.Pointer(joint)).IndexB = v9
	// Compute joint anchor frames with world space rotation, relative to center of mass
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v12 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA.Q
	qr.S = float32(v11.S*v12.C) + float32(v11.C*v12.S)
	qr.C = float32(v11.C*v12.C) - float32(v11.S*v12.S)
	v13 = qr
	goto _14
_14:
	(*b2WheelJoint)(unsafe.Pointer(joint)).FrameA.Q = v13
	v15 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA.P
	v16 = (*b2BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v17 = Vec2{
		X: v15.X - v16.X,
		Y: v15.Y - v16.Y,
	}
	goto _18
_18:
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v20 = v17
	v21 = Vec2{
		X: float32(v11.C*v20.X) - float32(v11.S*v20.Y),
		Y: float32(v11.S*v20.X) + float32(v11.C*v20.Y),
	}
	goto _22
_22:
	(*b2WheelJoint)(unsafe.Pointer(joint)).FrameA.P = v21
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v12 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameB.Q
	qr.S = float32(v11.S*v12.C) + float32(v11.C*v12.S)
	qr.C = float32(v11.C*v12.C) - float32(v11.S*v12.S)
	v13 = qr
	goto _26
_26:
	(*b2WheelJoint)(unsafe.Pointer(joint)).FrameB.Q = v13
	v15 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameB.P
	v16 = (*b2BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v17 = Vec2{
		X: v15.X - v16.X,
		Y: v15.Y - v16.Y,
	}
	goto _30
_30:
	v11 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v20 = v17
	v21 = Vec2{
		X: float32(v11.C*v20.X) - float32(v11.S*v20.Y),
		Y: float32(v11.S*v20.X) + float32(v11.C*v20.Y),
	}
	goto _34
_34:
	(*b2WheelJoint)(unsafe.Pointer(joint)).FrameB.P = v21
	// Compute the initial center delta. Incremental position updates are relative to this.
	v35 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Center
	v36 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Center
	v15 = Vec2{
		X: v35.X - v36.X,
		Y: v35.Y - v36.Y,
	}
	goto _38
_38:
	(*b2WheelJoint)(unsafe.Pointer(joint)).DeltaCenter = v15
	rA = (*b2WheelJoint)(unsafe.Pointer(joint)).FrameA.P
	rB = (*b2WheelJoint)(unsafe.Pointer(joint)).FrameB.P
	v15 = rB
	v16 = rA
	v17 = Vec2{
		X: v15.X - v16.X,
		Y: v15.Y - v16.Y,
	}
	goto _42
_42:
	v20 = (*b2WheelJoint)(unsafe.Pointer(joint)).DeltaCenter
	v21 = v17
	v27 = Vec2{
		X: v20.X + v21.X,
		Y: v20.Y + v21.Y,
	}
	goto _46
_46:
	d = v27
	v11 = (*b2WheelJoint)(unsafe.Pointer(joint)).FrameA.Q
	v28 = Vec2{
		X: float32FromFloat32(1),
	}
	v29 = Vec2{
		X: float32(v11.C*v28.X) - float32(v11.S*v28.Y),
		Y: float32(v11.S*v28.X) + float32(v11.C*v28.Y),
	}
	goto _50
_50:
	axisA = v29
	v32 = axisA
	v33 = Vec2{
		X: -v32.Y,
		Y: v32.X,
	}
	goto _53
_53:
	perpA = v33
	v37 = d
	v39 = rA
	v40 = Vec2{
		X: v37.X + v39.X,
		Y: v37.Y + v39.Y,
	}
	goto _57
_57:
	v41 = v40
	v43 = perpA
	v60 = float32(v41.X*v43.Y) - float32(v41.Y*v43.X)
	goto _61
_61:
	// perpendicular constraint (keep wheel on line)
	s1 = v60
	v44 = rB
	v45 = perpA
	v64 = float32(v44.X*v45.Y) - float32(v44.Y*v45.X)
	goto _65
_65:
	s2 = v64
	kp = mA + mB + float32(float32(iA*s1)*s1) + float32(float32(iB*s2)*s2)
	if kp > float32FromFloat32(0) {
		v66 = float32FromFloat32(1) / kp
	} else {
		v66 = float32FromFloat32(0)
	}
	(*b2WheelJoint)(unsafe.Pointer(joint)).PerpMass = v66
	v15 = d
	v16 = rA
	v17 = Vec2{
		X: v15.X + v16.X,
		Y: v15.Y + v16.Y,
	}
	goto _70
_70:
	v20 = v17
	v21 = axisA
	v60 = float32(v20.X*v21.Y) - float32(v20.Y*v21.X)
	goto _74
_74:
	// spring constraint
	a12 = v60
	v27 = rB
	v28 = axisA
	v64 = float32(v27.X*v28.Y) - float32(v27.Y*v28.X)
	goto _78
_78:
	a22 = v64
	ka = mA + mB + float32(float32(iA*a12)*a12) + float32(float32(iB*a22)*a22)
	if ka > float32FromFloat32(0) {
		v66 = float32FromFloat32(1) / ka
	} else {
		v66 = float32FromFloat32(0)
	}
	(*b2WheelJoint)(unsafe.Pointer(joint)).AxialMass = v66
	v60 = (*b2WheelJoint)(unsafe.Pointer(joint)).Hertz
	v64 = (*b2StepContext)(unsafe.Pointer(context)).H
	if v60 == float32FromFloat32(0) {
		v82 = b2Softness{}
		goto _83
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(_B2_PI)) * v60)
	a11 = float32(float32FromFloat32(2)*(*b2WheelJoint)(unsafe.Pointer(joint)).DampingRatio) + float32(v64*omega)
	a21 = float32(float32(v64*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v82 = b2Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _83
_83:
	(*b2WheelJoint)(unsafe.Pointer(joint)).SpringSoftness = v82
	km = iA + iB
	if km > float32FromFloat32(0) {
		v60 = float32FromFloat32(1) / km
	} else {
		v60 = float32FromFloat32(0)
	}
	(*b2WheelJoint)(unsafe.Pointer(joint)).MotorMass = v60
	if int32FromUint8((*b2StepContext)(unsafe.Pointer(context)).EnableWarmStarting) == false1 {
		(*b2WheelJoint)(unsafe.Pointer(joint)).PerpImpulse = float32FromFloat32(0)
		(*b2WheelJoint)(unsafe.Pointer(joint)).SpringImpulse = float32FromFloat32(0)
		(*b2WheelJoint)(unsafe.Pointer(joint)).MotorImpulse = float32FromFloat32(0)
		(*b2WheelJoint)(unsafe.Pointer(joint)).LowerImpulse = float32FromFloat32(0)
		(*b2WheelJoint)(unsafe.Pointer(joint)).UpperImpulse = float32FromFloat32(0)
	}
}

func b2WarmStartWheelJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var LA, LB, a11, a21, axialImpulse, iA, iB, mA, mB, s11, s21, v45, v49, v57, v61, v63, v67, v76 float32
	var P, axisA, d, perpA, rA, rB, v10, v12, v13, v14, v16, v17, v18, v20, v21, v22, v24, v25, v26, v29, v30, v33, v34, v36, v37, v39, v40, v41, v43, v5, v6, v9 Vec2
	var joint, stateA, stateB, v2, v3 uintptr
	var v1 bool
	var v28, v32, v4, v8 Rot
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = LA, LB, P, a11, a21, axialImpulse, axisA, d, iA, iB, joint, mA, mB, perpA, rA, rB, s11, s21, stateA, stateB, v1, v10, v12, v13, v14, v16, v17, v18, v2, v20, v21, v22, v24, v25, v26, v28, v29, v3, v30, v32, v33, v34, v36, v37, v39, v4, v40, v41, v43, v45, v49, v5, v57, v6, v61, v63, v67, v76, v8, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_wheelJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+28993, __ccgo_ts+28964, int32FromInt32(291))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Add(nil, bp)) = b2_identityBodyState23
	joint = base + 92
	if (*b2WheelJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2WheelJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v2
	if (*b2WheelJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2WheelJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v3
	v4 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b2WheelJoint)(unsafe.Pointer(joint)).FrameA.P
	v6 = Vec2{
		X: float32(v4.C*v5.X) - float32(v4.S*v5.Y),
		Y: float32(v4.S*v5.X) + float32(v4.C*v5.Y),
	}
	goto _7
_7:
	rA = v6
	v8 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v9 = (*b2WheelJoint)(unsafe.Pointer(joint)).FrameB.P
	v10 = Vec2{
		X: float32(v8.C*v9.X) - float32(v8.S*v9.Y),
		Y: float32(v8.S*v9.X) + float32(v8.C*v9.Y),
	}
	goto _11
_11:
	rB = v10
	v12 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
	v13 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaPosition
	v14 = Vec2{
		X: v12.X - v13.X,
		Y: v12.Y - v13.Y,
	}
	goto _15
_15:
	v16 = v14
	v17 = (*b2WheelJoint)(unsafe.Pointer(joint)).DeltaCenter
	v18 = Vec2{
		X: v16.X + v17.X,
		Y: v16.Y + v17.Y,
	}
	goto _19
_19:
	v20 = rB
	v21 = rA
	v22 = Vec2{
		X: v20.X - v21.X,
		Y: v20.Y - v21.Y,
	}
	goto _23
_23:
	v24 = v18
	v25 = v22
	v26 = Vec2{
		X: v24.X + v25.X,
		Y: v24.Y + v25.Y,
	}
	goto _27
_27:
	d = v26
	v28 = (*b2WheelJoint)(unsafe.Pointer(joint)).FrameA.Q
	v29 = Vec2{
		X: float32FromFloat32(1),
	}
	v30 = Vec2{
		X: float32(v28.C*v29.X) - float32(v28.S*v29.Y),
		Y: float32(v28.S*v29.X) + float32(v28.C*v29.Y),
	}
	goto _31
_31:
	axisA = v30
	v32 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v33 = axisA
	v34 = Vec2{
		X: float32(v32.C*v33.X) - float32(v32.S*v33.Y),
		Y: float32(v32.S*v33.X) + float32(v32.C*v33.Y),
	}
	goto _35
_35:
	axisA = v34
	v5 = axisA
	v6 = Vec2{
		X: -v5.Y,
		Y: v5.X,
	}
	goto _38
_38:
	perpA = v6
	v9 = d
	v10 = rA
	v12 = Vec2{
		X: v9.X + v10.X,
		Y: v9.Y + v10.Y,
	}
	goto _42
_42:
	v13 = v12
	v14 = axisA
	v45 = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
	goto _46
_46:
	a11 = v45
	v16 = rB
	v17 = axisA
	v49 = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	goto _50
_50:
	a21 = v49
	v18 = d
	v20 = rA
	v21 = Vec2{
		X: v18.X + v20.X,
		Y: v18.Y + v20.Y,
	}
	goto _54
_54:
	v22 = v21
	v24 = perpA
	v57 = float32(v22.X*v24.Y) - float32(v22.Y*v24.X)
	goto _58
_58:
	s11 = v57
	v25 = rB
	v26 = perpA
	v61 = float32(v25.X*v26.Y) - float32(v25.Y*v26.X)
	goto _62
_62:
	s21 = v61
	axialImpulse = (*b2WheelJoint)(unsafe.Pointer(joint)).SpringImpulse + (*b2WheelJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b2WheelJoint)(unsafe.Pointer(joint)).UpperImpulse
	v63 = axialImpulse
	v29 = axisA
	v30 = Vec2{
		X: float32(v63 * v29.X),
		Y: float32(v63 * v29.Y),
	}
	goto _66
_66:
	v67 = (*b2WheelJoint)(unsafe.Pointer(joint)).PerpImpulse
	v33 = perpA
	v34 = Vec2{
		X: float32(v67 * v33.X),
		Y: float32(v67 * v33.Y),
	}
	goto _70
_70:
	v36 = v30
	v37 = v34
	v39 = Vec2{
		X: v36.X + v37.X,
		Y: v36.Y + v37.Y,
	}
	goto _74
_74:
	P = v39
	LA = float32(axialImpulse*a11) + float32((*b2WheelJoint)(unsafe.Pointer(joint)).PerpImpulse*s11) + (*b2WheelJoint)(unsafe.Pointer(joint)).MotorImpulse
	LB = float32(axialImpulse*a21) + float32((*b2WheelJoint)(unsafe.Pointer(joint)).PerpImpulse*s21) + (*b2WheelJoint)(unsafe.Pointer(joint)).MotorImpulse
	if (*b2BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b2_dynamicFlag) != 0 {
		v40 = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
		v76 = mA
		v41 = P
		v43 = Vec2{
			X: v40.X - float32(v76*v41.X),
			Y: v40.Y - float32(v76*v41.Y),
		}
		goto _79
	_79:
		(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = v43
		*(*float32)(unsafe.Add(nil, stateA+8)) -= float32(iA * LA)
	}
	if (*b2BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b2_dynamicFlag) != 0 {
		v5 = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
		v45 = mB
		v6 = P
		v9 = Vec2{
			X: v5.X + float32(v45*v6.X),
			Y: v5.Y + float32(v45*v6.Y),
		}
		goto _84
	_84:
		(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = v9
		*(*float32)(unsafe.Add(nil, stateB+8)) += float32(iB * LB)
	}
}

func b2SolveWheelJoint(tls *_Stack, base uintptr, context uintptr, useBias uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var C, C1, C2, C3, Cdot, Cdot1, Cdot2, Cdot3, Cdot4, LA, LA1, LA2, LA3, LB, LB1, LB2, LB3, a11, a21, bias, bias1, bias2, bias3, iA, iB, impulse, impulse1, impulse2, impulse3, impulse4, impulseScale, impulseScale1, impulseScale2, impulseScale3, mA, mB, massScale, massScale1, massScale2, massScale3, maxImpulse, oldImpulse, oldImpulse1, oldImpulse2, s11, s21, translation, wA, wB, v38, v46, v50, v52, v53, v54, v55, v57, v58 float32
	var P, P1, P2, P3, axisA, d, perpA, rA, rB, vA, vB, v10, v12, v13, v14, v16, v17, v18, v20, v21, v22, v24, v25, v26, v29, v30, v33, v34, v5, v6, v9 Vec2
	var fixedRotation uint8
	var joint, stateA, stateB, v2, v3 uintptr
	var v1 bool
	var v28, v32, v4, v8 Rot
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, C1, C2, C3, Cdot, Cdot1, Cdot2, Cdot3, Cdot4, LA, LA1, LA2, LA3, LB, LB1, LB2, LB3, P, P1, P2, P3, a11, a21, axisA, bias, bias1, bias2, bias3, d, fixedRotation, iA, iB, impulse, impulse1, impulse2, impulse3, impulse4, impulseScale, impulseScale1, impulseScale2, impulseScale3, joint, mA, mB, massScale, massScale1, massScale2, massScale3, maxImpulse, oldImpulse, oldImpulse1, oldImpulse2, perpA, rA, rB, s11, s21, stateA, stateB, translation, vA, vB, wA, wB, v1, v10, v12, v13, v14, v16, v17, v18, v2, v20, v21, v22, v24, v25, v26, v28, v29, v3, v30, v32, v33, v34, v38, v4, v46, v5, v50, v52, v53, v54, v55, v57, v58, v6, v8, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_wheelJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+28993, __ccgo_ts+28964, int32FromInt32(340))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Add(nil, bp)) = b2_identityBodyState23
	joint = base + 92
	if (*b2WheelJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2WheelJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v2
	if (*b2WheelJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v3 = bp
	} else {
		v3 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2WheelJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v3
	vA = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	fixedRotation = boolUint8(iA+iB == float32FromFloat32(0))
	v4 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v5 = (*b2WheelJoint)(unsafe.Pointer(joint)).FrameA.P
	v6 = Vec2{
		X: float32(v4.C*v5.X) - float32(v4.S*v5.Y),
		Y: float32(v4.S*v5.X) + float32(v4.C*v5.Y),
	}
	goto _7
_7:
	// current anchors
	rA = v6
	v8 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v9 = (*b2WheelJoint)(unsafe.Pointer(joint)).FrameB.P
	v10 = Vec2{
		X: float32(v8.C*v9.X) - float32(v8.S*v9.Y),
		Y: float32(v8.S*v9.X) + float32(v8.C*v9.Y),
	}
	goto _11
_11:
	rB = v10
	v12 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
	v13 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaPosition
	v14 = Vec2{
		X: v12.X - v13.X,
		Y: v12.Y - v13.Y,
	}
	goto _15
_15:
	v16 = v14
	v17 = (*b2WheelJoint)(unsafe.Pointer(joint)).DeltaCenter
	v18 = Vec2{
		X: v16.X + v17.X,
		Y: v16.Y + v17.Y,
	}
	goto _19
_19:
	v20 = rB
	v21 = rA
	v22 = Vec2{
		X: v20.X - v21.X,
		Y: v20.Y - v21.Y,
	}
	goto _23
_23:
	v24 = v18
	v25 = v22
	v26 = Vec2{
		X: v24.X + v25.X,
		Y: v24.Y + v25.Y,
	}
	goto _27
_27:
	d = v26
	v28 = (*b2WheelJoint)(unsafe.Pointer(joint)).FrameA.Q
	v29 = Vec2{
		X: float32FromFloat32(1),
	}
	v30 = Vec2{
		X: float32(v28.C*v29.X) - float32(v28.S*v29.Y),
		Y: float32(v28.S*v29.X) + float32(v28.C*v29.Y),
	}
	goto _31
_31:
	axisA = v30
	v32 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v33 = axisA
	v34 = Vec2{
		X: float32(v32.C*v33.X) - float32(v32.S*v33.Y),
		Y: float32(v32.S*v33.X) + float32(v32.C*v33.Y),
	}
	goto _35
_35:
	axisA = v34
	v5 = axisA
	v6 = d
	v38 = float32(v5.X*v6.X) + float32(v5.Y*v6.Y)
	goto _39
_39:
	translation = v38
	v9 = d
	v10 = rA
	v12 = Vec2{
		X: v9.X + v10.X,
		Y: v9.Y + v10.Y,
	}
	goto _43
_43:
	v13 = v12
	v14 = axisA
	v46 = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
	goto _47
_47:
	a11 = v46
	v16 = rB
	v17 = axisA
	v50 = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	goto _51
_51:
	a21 = v50
	// motor constraint
	if (*b2WheelJoint)(unsafe.Pointer(joint)).EnableMotor != 0 && int32FromUint8(fixedRotation) == false1 {
		Cdot = wB - wA - (*b2WheelJoint)(unsafe.Pointer(joint)).MotorSpeed
		impulse = float32(-(*b2WheelJoint)(unsafe.Pointer(joint)).MotorMass * Cdot)
		oldImpulse = (*b2WheelJoint)(unsafe.Pointer(joint)).MotorImpulse
		maxImpulse = float32((*b2StepContext)(unsafe.Pointer(context)).H * (*b2WheelJoint)(unsafe.Pointer(joint)).MaxMotorTorque)
		v52 = (*b2WheelJoint)(unsafe.Pointer(joint)).MotorImpulse + impulse
		v53 = -maxImpulse
		v54 = maxImpulse
		if v52 < v53 {
			v57 = v53
		} else {
			if v52 > v54 {
				v58 = v54
			} else {
				v58 = v52
			}
			v57 = v58
		}
		v55 = v57
		goto _56
	_56:
		(*b2WheelJoint)(unsafe.Pointer(joint)).MotorImpulse = v55
		impulse = (*b2WheelJoint)(unsafe.Pointer(joint)).MotorImpulse - oldImpulse
		wA = wA - float32(iA*impulse)
		wB = wB + float32(iB*impulse)
	}
	// spring constraint
	if (*b2WheelJoint)(unsafe.Pointer(joint)).EnableSpring != 0 {
		// This is a real spring and should be applied even during relax
		C = translation
		bias = float32((*b2WheelJoint)(unsafe.Pointer(joint)).SpringSoftness.BiasRate * C)
		massScale = (*b2WheelJoint)(unsafe.Pointer(joint)).SpringSoftness.MassScale
		impulseScale = (*b2WheelJoint)(unsafe.Pointer(joint)).SpringSoftness.ImpulseScale
		v5 = vB
		v6 = vA
		v9 = Vec2{
			X: v5.X - v6.X,
			Y: v5.Y - v6.Y,
		}
		goto _62
	_62:
		v10 = axisA
		v12 = v9
		v38 = float32(v10.X*v12.X) + float32(v10.Y*v12.Y)
		goto _66
	_66:
		Cdot1 = v38 + float32(a21*wB) - float32(a11*wA)
		impulse1 = float32(float32(-massScale*(*b2WheelJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot1+bias)) - float32(impulseScale*(*b2WheelJoint)(unsafe.Pointer(joint)).SpringImpulse)
		*(*float32)(unsafe.Add(nil, joint+8)) += impulse1
		v38 = impulse1
		v5 = axisA
		v6 = Vec2{
			X: float32(v38 * v5.X),
			Y: float32(v38 * v5.Y),
		}
		goto _70
	_70:
		P = v6
		LA = float32(impulse1 * a11)
		LB = float32(impulse1 * a21)
		v9 = vA
		v46 = mA
		v10 = P
		v12 = Vec2{
			X: v9.X - float32(v46*v10.X),
			Y: v9.Y - float32(v46*v10.Y),
		}
		goto _75
	_75:
		vA = v12
		wA = wA - float32(iA*LA)
		v5 = vB
		v38 = mB
		v6 = P
		v9 = Vec2{
			X: v5.X + float32(v38*v6.X),
			Y: v5.Y + float32(v38*v6.Y),
		}
		goto _80
	_80:
		vB = v9
		wB = wB + float32(iB*LB)
	}
	if (*b2WheelJoint)(unsafe.Pointer(joint)).EnableLimit != 0 {
		// Lower limit
		C1 = translation - (*b2WheelJoint)(unsafe.Pointer(joint)).LowerTranslation
		bias1 = float32FromFloat32(0)
		massScale1 = float32FromFloat32(1)
		impulseScale1 = float32FromFloat32(0)
		if C1 > float32FromFloat32(0) {
			// speculation
			bias1 = float32(C1 * (*b2StepContext)(unsafe.Pointer(context)).Inv_h)
		} else {
			if useBias != 0 {
				bias1 = float32((*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * C1)
				massScale1 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
				impulseScale1 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
			}
		}
		v5 = vB
		v6 = vA
		v9 = Vec2{
			X: v5.X - v6.X,
			Y: v5.Y - v6.Y,
		}
		goto _84
	_84:
		v10 = axisA
		v12 = v9
		v38 = float32(v10.X*v12.X) + float32(v10.Y*v12.Y)
		goto _88
	_88:
		Cdot2 = v38 + float32(a21*wB) - float32(a11*wA)
		impulse2 = float32(float32(-massScale1*(*b2WheelJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot2+bias1)) - float32(impulseScale1*(*b2WheelJoint)(unsafe.Pointer(joint)).LowerImpulse)
		oldImpulse1 = (*b2WheelJoint)(unsafe.Pointer(joint)).LowerImpulse
		v46 = oldImpulse1 + impulse2
		v50 = float32FromFloat32(0)
		if v46 > v50 {
			v53 = v46
		} else {
			v53 = v50
		}
		v52 = v53
		goto _92
	_92:
		(*b2WheelJoint)(unsafe.Pointer(joint)).LowerImpulse = v52
		impulse2 = (*b2WheelJoint)(unsafe.Pointer(joint)).LowerImpulse - oldImpulse1
		v38 = impulse2
		v5 = axisA
		v6 = Vec2{
			X: float32(v38 * v5.X),
			Y: float32(v38 * v5.Y),
		}
		goto _97
	_97:
		P1 = v6
		LA1 = float32(impulse2 * a11)
		LB1 = float32(impulse2 * a21)
		v9 = vA
		v46 = mA
		v10 = P1
		v12 = Vec2{
			X: v9.X - float32(v46*v10.X),
			Y: v9.Y - float32(v46*v10.Y),
		}
		goto _102
	_102:
		vA = v12
		wA = wA - float32(iA*LA1)
		v5 = vB
		v38 = mB
		v6 = P1
		v9 = Vec2{
			X: v5.X + float32(v38*v6.X),
			Y: v5.Y + float32(v38*v6.Y),
		}
		goto _107
	_107:
		vB = v9
		wB = wB + float32(iB*LB1)
		// Upper limit
		// Note: signs are flipped to keep C positive when the constraint is satisfied.
		// This also keeps the impulse positive when the limit is active.
		// sign flipped
		C2 = (*b2WheelJoint)(unsafe.Pointer(joint)).UpperTranslation - translation
		bias2 = float32FromFloat32(0)
		massScale2 = float32FromFloat32(1)
		impulseScale2 = float32FromFloat32(0)
		if C2 > float32FromFloat32(0) {
			// speculation
			bias2 = float32(C2 * (*b2StepContext)(unsafe.Pointer(context)).Inv_h)
		} else {
			if useBias != 0 {
				bias2 = float32((*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * C2)
				massScale2 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
				impulseScale2 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
			}
		}
		v5 = vA
		v6 = vB
		v9 = Vec2{
			X: v5.X - v6.X,
			Y: v5.Y - v6.Y,
		}
		goto _111
	_111:
		v10 = axisA
		v12 = v9
		v38 = float32(v10.X*v12.X) + float32(v10.Y*v12.Y)
		goto _115
	_115:
		// sign flipped on Cdot
		Cdot3 = v38 + float32(a11*wA) - float32(a21*wB)
		impulse3 = float32(float32(-massScale2*(*b2WheelJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot3+bias2)) - float32(impulseScale2*(*b2WheelJoint)(unsafe.Pointer(joint)).UpperImpulse)
		oldImpulse2 = (*b2WheelJoint)(unsafe.Pointer(joint)).UpperImpulse
		v46 = oldImpulse2 + impulse3
		v50 = float32FromFloat32(0)
		if v46 > v50 {
			v53 = v46
		} else {
			v53 = v50
		}
		v52 = v53
		goto _119
	_119:
		(*b2WheelJoint)(unsafe.Pointer(joint)).UpperImpulse = v52
		impulse3 = (*b2WheelJoint)(unsafe.Pointer(joint)).UpperImpulse - oldImpulse2
		v38 = impulse3
		v5 = axisA
		v6 = Vec2{
			X: float32(v38 * v5.X),
			Y: float32(v38 * v5.Y),
		}
		goto _124
	_124:
		P2 = v6
		LA2 = float32(impulse3 * a11)
		LB2 = float32(impulse3 * a21)
		// sign flipped on applied impulse
		v9 = vA
		v46 = mA
		v10 = P2
		v12 = Vec2{
			X: v9.X + float32(v46*v10.X),
			Y: v9.Y + float32(v46*v10.Y),
		}
		goto _129
	_129:
		vA = v12
		wA = wA + float32(iA*LA2)
		v5 = vB
		v38 = mB
		v6 = P2
		v9 = Vec2{
			X: v5.X - float32(v38*v6.X),
			Y: v5.Y - float32(v38*v6.Y),
		}
		goto _134
	_134:
		vB = v9
		wB = wB - float32(iB*LB2)
	}
	// point to line constraint
	v5 = axisA
	v6 = Vec2{
		X: -v5.Y,
		Y: v5.X,
	}
	goto _137
_137:
	perpA = v6
	bias3 = float32FromFloat32(0)
	massScale3 = float32FromFloat32(1)
	impulseScale3 = float32FromFloat32(0)
	if useBias != 0 {
		v9 = perpA
		v10 = d
		v38 = float32(v9.X*v10.X) + float32(v9.Y*v10.Y)
		goto _141
	_141:
		C3 = v38
		bias3 = float32((*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * C3)
		massScale3 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
		impulseScale3 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
	}
	v5 = d
	v6 = rA
	v9 = Vec2{
		X: v5.X + v6.X,
		Y: v5.Y + v6.Y,
	}
	goto _145
_145:
	v10 = v9
	v12 = perpA
	v38 = float32(v10.X*v12.Y) - float32(v10.Y*v12.X)
	goto _149
_149:
	s11 = v38
	v13 = rB
	v14 = perpA
	v46 = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
	goto _153
_153:
	s21 = v46
	v16 = vB
	v17 = vA
	v18 = Vec2{
		X: v16.X - v17.X,
		Y: v16.Y - v17.Y,
	}
	goto _157
_157:
	v20 = perpA
	v21 = v18
	v50 = float32(v20.X*v21.X) + float32(v20.Y*v21.Y)
	goto _161
_161:
	Cdot4 = v50 + float32(s21*wB) - float32(s11*wA)
	impulse4 = float32(float32(-massScale3*(*b2WheelJoint)(unsafe.Pointer(joint)).PerpMass)*(Cdot4+bias3)) - float32(impulseScale3*(*b2WheelJoint)(unsafe.Pointer(joint)).PerpImpulse)
	*(*float32)(unsafe.Add(nil, joint)) += impulse4
	v38 = impulse4
	v5 = perpA
	v6 = Vec2{
		X: float32(v38 * v5.X),
		Y: float32(v38 * v5.Y),
	}
	goto _165
_165:
	P3 = v6
	LA3 = float32(impulse4 * s11)
	LB3 = float32(impulse4 * s21)
	v9 = vA
	v46 = mA
	v10 = P3
	v12 = Vec2{
		X: v9.X - float32(v46*v10.X),
		Y: v9.Y - float32(v46*v10.Y),
	}
	goto _170
_170:
	vA = v12
	wA = wA - float32(iA*LA3)
	v5 = vB
	v38 = mB
	v6 = P3
	v9 = Vec2{
		X: v5.X + float32(v38*v6.X),
		Y: v5.Y + float32(v38*v6.Y),
	}
	goto _175
_175:
	vB = v9
	wB = wB + float32(iB*LB3)
	if (*b2BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b2_dynamicFlag) != 0 {
		(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
		(*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	}
	if (*b2BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b2_dynamicFlag) != 0 {
		(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
		(*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
	}
}

func b2DrawWheelJoint(tls *_Stack, draw uintptr, base uintptr, transformA WorldTransform, transformB WorldTransform, drawScale float32) {
	var C, frameA, frameB, v16, v18, v2, v32 WorldTransform
	var axisA, perp, v10, v13, v25, v26, v29, v35, v36, v9 Vec2
	var c1, c2, c3, c4, c5 b2HexColor1
	var joint uintptr
	var lower, upper, v12, v14, v28, v30 Pos
	var qr, v20, v21, v22, v24, v34, v4, v5, v6, v8 Rot
	var v1 bool
	var v19, v3 Transform
	var v38, v46 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, axisA, c1, c2, c3, c4, c5, frameA, frameB, joint, lower, perp, qr, upper, v1, v10, v12, v13, v14, v16, v18, v19, v2, v20, v21, v22, v24, v25, v26, v28, v29, v3, v30, v32, v34, v35, v36, v38, v4, v46, v5, v6, v8, v9
	if v1 = !!((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_wheelJoint)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+28993, __ccgo_ts+28964, int32FromInt32(560))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	joint = base + 92
	v2 = transformA
	v3 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameA
	v4 = v2.Q
	v5 = v3.Q
	qr.S = float32(v4.S*v5.C) + float32(v4.C*v5.S)
	qr.C = float32(v4.C*v5.C) - float32(v4.S*v5.S)
	v6 = qr
	goto _7
_7:
	C.Q = v6
	v8 = v2.Q
	v9 = v3.P
	v10 = Vec2{
		X: float32(v8.C*v9.X) - float32(v8.S*v9.Y),
		Y: float32(v8.S*v9.X) + float32(v8.C*v9.Y),
	}
	goto _11
_11:
	v12 = v2.P
	v13 = v10
	v14 = Pos{
		X: v12.X + v13.X,
		Y: v12.Y + v13.Y,
	}
	goto _15
_15:
	C.P = v14
	v16 = C
	goto _17
_17:
	frameA = v16
	v18 = transformB
	v19 = (*b2JointSim)(unsafe.Pointer(base)).LocalFrameB
	v20 = v18.Q
	v21 = v19.Q
	qr.S = float32(v20.S*v21.C) + float32(v20.C*v21.S)
	qr.C = float32(v20.C*v21.C) - float32(v20.S*v21.S)
	v22 = qr
	goto _23
_23:
	C.Q = v22
	v24 = v18.Q
	v25 = v19.P
	v26 = Vec2{
		X: float32(v24.C*v25.X) - float32(v24.S*v25.Y),
		Y: float32(v24.S*v25.X) + float32(v24.C*v25.Y),
	}
	goto _27
_27:
	v28 = v18.P
	v29 = v26
	v30 = Pos{
		X: v28.X + v29.X,
		Y: v28.Y + v29.Y,
	}
	goto _31
_31:
	C.P = v30
	v32 = C
	goto _33
_33:
	frameB = v32
	v34 = frameA.Q
	v35 = Vec2{
		X: float32FromFloat32(1),
	}
	v36 = Vec2{
		X: float32(v34.C*v35.X) - float32(v34.S*v35.Y),
		Y: float32(v34.S*v35.X) + float32(v34.C*v35.Y),
	}
	goto _37
_37:
	axisA = v36
	c1 = int32(b2_colorGray)
	c2 = int32(b2_colorGreen)
	c3 = int32(b2_colorRed)
	c4 = int32(b2_colorDimGray)
	c5 = int32(b2_colorBlue)
	(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, frameA.P, frameB.P, c5, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	if (*b2WheelJoint)(unsafe.Pointer(joint)).EnableLimit != 0 {
		v38 = (*b2WheelJoint)(unsafe.Pointer(joint)).LowerTranslation
		v9 = axisA
		v10 = Vec2{
			X: float32(v38 * v9.X),
			Y: float32(v38 * v9.Y),
		}
		goto _41
	_41:
		v12 = frameA.P
		v13 = v10
		v14 = Pos{
			X: v12.X + v13.X,
			Y: v12.Y + v13.Y,
		}
		goto _45
	_45:
		lower = v14
		v46 = (*b2WheelJoint)(unsafe.Pointer(joint)).UpperTranslation
		v25 = axisA
		v26 = Vec2{
			X: float32(v46 * v25.X),
			Y: float32(v46 * v25.Y),
		}
		goto _49
	_49:
		v28 = frameA.P
		v29 = v26
		v30 = Pos{
			X: v28.X + v29.X,
			Y: v28.Y + v29.Y,
		}
		goto _53
	_53:
		upper = v30
		v35 = axisA
		v36 = Vec2{
			X: -v35.Y,
			Y: v35.X,
		}
		goto _56
	_56:
		perp = v36
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, lower, upper, c1, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		v38 = float32(-float32FromFloat32(0.1) * drawScale)
		v9 = perp
		v10 = Vec2{
			X: float32(v38 * v9.X),
			Y: float32(v38 * v9.Y),
		}
		goto _60
	_60:
		v12 = lower
		v13 = v10
		v14 = Pos{
			X: v12.X + v13.X,
			Y: v12.Y + v13.Y,
		}
		goto _64
	_64:
		v46 = float32(float32FromFloat32(0.1) * drawScale)
		v25 = perp
		v26 = Vec2{
			X: float32(v46 * v25.X),
			Y: float32(v46 * v25.Y),
		}
		goto _68
	_68:
		v28 = lower
		v29 = v26
		v30 = Pos{
			X: v28.X + v29.X,
			Y: v28.Y + v29.Y,
		}
		goto _72
	_72:
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, v14, v30, c2, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		v38 = float32(-float32FromFloat32(0.1) * drawScale)
		v9 = perp
		v10 = Vec2{
			X: float32(v38 * v9.X),
			Y: float32(v38 * v9.Y),
		}
		goto _76
	_76:
		v12 = upper
		v13 = v10
		v14 = Pos{
			X: v12.X + v13.X,
			Y: v12.Y + v13.Y,
		}
		goto _80
	_80:
		v46 = float32(float32FromFloat32(0.1) * drawScale)
		v25 = perp
		v26 = Vec2{
			X: float32(v46 * v25.X),
			Y: float32(v46 * v25.Y),
		}
		goto _84
	_84:
		v28 = upper
		v29 = v26
		v30 = Pos{
			X: v28.X + v29.X,
			Y: v28.Y + v29.Y,
		}
		goto _88
	_88:
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, v14, v30, c3, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	} else {
		v9 = axisA
		v10 = Vec2{
			X: -v9.X,
			Y: -v9.Y,
		}
		goto _91
	_91:
		v12 = frameA.P
		v13 = v10
		v14 = Pos{
			X: v12.X + v13.X,
			Y: v12.Y + v13.Y,
		}
		goto _95
	_95:
		v28 = frameA.P
		v25 = axisA
		v30 = Pos{
			X: v28.X + v25.X,
			Y: v28.Y + v25.Y,
		}
		goto _99
	_99:
		(*(*func(*_Stack, Pos, Pos, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawLineFcn})))(tls, v14, v30, c1, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	}
	(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, frameA.P, float32FromFloat32(5), c1, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	(*(*func(*_Stack, Pos, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, frameB.P, float32FromFloat32(5), c4, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
}
