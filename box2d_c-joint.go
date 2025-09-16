package b2

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b2AssignJointColor(tls *_Stack, graph uintptr, bodyIdA int32, bodyIdB int32, staticA uint8, staticB uint8) (r int32) {
	var blockIndex, blockIndex1, v12, v14, v17, v21, v24, v28, v3, v7 uint32_t
	var color, color1, color2, v11, v13, v16, v2, v20, v23, v27, v6 uintptr
	var i, i1, i2 int32
	var v10 bool
	var v18, v25, v4, v8 uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = blockIndex, blockIndex1, color, color1, color2, i, i1, i2, v10, v11, v12, v13, v14, v16, v17, v18, v2, v20, v21, v23, v24, v25, v27, v28, v3, v4, v6, v7, v8
	if !(int32FromUint8(staticA) == false1 || int32FromUint8(staticB) == false1) && b2InternalAssertFcn(tls, __ccgo_ts+2939, __ccgo_ts+2787, int32FromInt32(216)) != 0 {
		__builtin_trap(tls)
	}
	if int32FromUint8(staticA) == false1 && int32FromUint8(staticB) == false1 {
		i = 0
		for {
			if !(i < int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1)) {
				break
			}
			color = graph + uintptr(i)*56
			v2 = color
			v3 = uint32FromInt32(bodyIdA)
			blockIndex1 = v3 / uint32(64)
			if blockIndex1 >= (*b2BitSet)(unsafe.Pointer(v2)).BlockCount {
				v4 = boolUint8(false1 != 0)
				goto _5
			}
			v4 = boolUint8(*(*uint64_t)(unsafe.Pointer((*b2BitSet)(unsafe.Pointer(v2)).Bits + uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v3%uint32FromInt32(64))) != uint64(0))
			goto _5
		_5:
			;
			if v10 = v4 != 0; !v10 {
				v6 = color
				v7 = uint32FromInt32(bodyIdB)
				blockIndex1 = v7 / uint32(64)
				if blockIndex1 >= (*b2BitSet)(unsafe.Pointer(v6)).BlockCount {
					v8 = boolUint8(false1 != 0)
					goto _9
				}
				v8 = boolUint8(*(*uint64_t)(unsafe.Pointer((*b2BitSet)(unsafe.Pointer(v6)).Bits + uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v7%uint32FromInt32(64))) != uint64(0))
				goto _9
			_9:
			}
			if v10 || v8 != 0 {
				goto _1
			}
			v11 = color
			v12 = uint32FromInt32(bodyIdA)
			blockIndex = v12 / uint32(64)
			if blockIndex >= (*b2BitSet)(unsafe.Pointer(v11)).BlockCount {
				b2GrowBitSet(tls, v11, blockIndex+uint32(1))
			}
			*(*uint64_t)(unsafe.Pointer((*b2BitSet)(unsafe.Pointer(v11)).Bits + uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v12 % uint32FromInt32(64))
			v13 = color
			v14 = uint32FromInt32(bodyIdB)
			blockIndex = v14 / uint32(64)
			if blockIndex >= (*b2BitSet)(unsafe.Pointer(v13)).BlockCount {
				b2GrowBitSet(tls, v13, blockIndex+uint32(1))
			}
			*(*uint64_t)(unsafe.Pointer((*b2BitSet)(unsafe.Pointer(v13)).Bits + uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v14 % uint32FromInt32(64))
			return i
			goto _1
		_1:
			;
			i = i + 1
		}
	} else {
		if int32FromUint8(staticA) == false1 {
			i1 = 0
			for {
				if !(i1 < int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1)) {
					break
				}
				color1 = graph + uintptr(i1)*56
				v16 = color1
				v17 = uint32FromInt32(bodyIdA)
				blockIndex1 = v17 / uint32(64)
				if blockIndex1 >= (*b2BitSet)(unsafe.Pointer(v16)).BlockCount {
					v18 = boolUint8(false1 != 0)
					goto _19
				}
				v18 = boolUint8(*(*uint64_t)(unsafe.Pointer((*b2BitSet)(unsafe.Pointer(v16)).Bits + uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v17%uint32FromInt32(64))) != uint64(0))
				goto _19
			_19:
				if v18 != 0 {
					goto _15
				}
				v20 = color1
				v21 = uint32FromInt32(bodyIdA)
				blockIndex = v21 / uint32(64)
				if blockIndex >= (*b2BitSet)(unsafe.Pointer(v20)).BlockCount {
					b2GrowBitSet(tls, v20, blockIndex+uint32(1))
				}
				*(*uint64_t)(unsafe.Pointer((*b2BitSet)(unsafe.Pointer(v20)).Bits + uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v21 % uint32FromInt32(64))
				return i1
				goto _15
			_15:
				;
				i1 = i1 + 1
			}
		} else {
			if int32FromUint8(staticB) == false1 {
				i2 = 0
				for {
					if !(i2 < int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1)) {
						break
					}
					color2 = graph + uintptr(i2)*56
					v23 = color2
					v24 = uint32FromInt32(bodyIdB)
					blockIndex1 = v24 / uint32(64)
					if blockIndex1 >= (*b2BitSet)(unsafe.Pointer(v23)).BlockCount {
						v25 = boolUint8(false1 != 0)
						goto _26
					}
					v25 = boolUint8(*(*uint64_t)(unsafe.Pointer((*b2BitSet)(unsafe.Pointer(v23)).Bits + uintptr(blockIndex1)*8))&(uint64FromInt32(1)<<(v24%uint32FromInt32(64))) != uint64(0))
					goto _26
				_26:
					if v25 != 0 {
						goto _22
					}
					v27 = color2
					v28 = uint32FromInt32(bodyIdB)
					blockIndex = v28 / uint32(64)
					if blockIndex >= (*b2BitSet)(unsafe.Pointer(v27)).BlockCount {
						b2GrowBitSet(tls, v27, blockIndex+uint32(1))
					}
					*(*uint64_t)(unsafe.Pointer((*b2BitSet)(unsafe.Pointer(v27)).Bits + uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v28 % uint32FromInt32(64))
					return i2
					goto _22
				_22:
					;
					i2 = i2 + 1
				}
			}
		}
	}
	return int32FromInt32(_B2_GRAPH_COLOR_COUNT) - int32FromInt32(1)
}

func b2CreateJointInGraph(tls *_Stack, world uintptr, joint uintptr) (r uintptr) {
	var bodyA, bodyB, graph, jointSim, v1, v11, v3, v5, v7, v9 uintptr
	var bodyIdA, bodyIdB, colorIndex, newCapacity, v10, v2, v6 int32
	var staticA, staticB uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, bodyIdA, bodyIdB, colorIndex, graph, jointSim, newCapacity, staticA, staticB, v1, v10, v11, v2, v3, v5, v6, v7, v9
	graph = world + 328
	bodyIdA = (*(*b2JointEdge)(unsafe.Pointer(joint + 20))).BodyId
	bodyIdB = (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + 1*12))).BodyId
	v1 = world + 1024
	v2 = bodyIdA
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	bodyA = v3
	v5 = world + 1024
	v6 = bodyIdB
	if !(0 <= v6 && v6 < (*b2BodyArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodyArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*128
	goto _8
_8:
	bodyB = v7
	staticA = boolUint8((*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_staticSet))
	staticB = boolUint8((*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_staticSet))
	colorIndex = b2AssignJointColor(tls, graph, bodyIdA, bodyIdB, staticA, staticB)
	v9 = graph + uintptr(colorIndex)*56 + 32
	if (*b2JointSimArray)(unsafe.Pointer(v9)).Count == (*b2JointSimArray)(unsafe.Pointer(v9)).Capacity {
		if (*b2JointSimArray)(unsafe.Pointer(v9)).Capacity < int32(2) {
			v10 = int32(2)
		} else {
			v10 = (*b2JointSimArray)(unsafe.Pointer(v9)).Capacity + (*b2JointSimArray)(unsafe.Pointer(v9)).Capacity>>int32(1)
		}
		newCapacity = v10
		b2JointSimArray_Reserve(tls, v9, newCapacity)
	}
	*(*int32)(unsafe.Pointer(v9 + 8)) += int32(1)
	v11 = (*b2JointSimArray)(unsafe.Pointer(v9)).Data + uintptr((*b2JointSimArray)(unsafe.Pointer(v9)).Count-int32FromInt32(1))*196
	goto _12
_12:
	jointSim = v11
	memset(tls, jointSim, 0, uint64(196))
	(*b2Joint)(unsafe.Pointer(joint)).ColorIndex = colorIndex
	(*b2Joint)(unsafe.Pointer(joint)).LocalIndex = (*(*b2GraphColor)(unsafe.Pointer(graph + uintptr(colorIndex)*56))).JointSims.Count - int32(1)
	return jointSim
}

func b2AddJointToGraph(tls *_Stack, world uintptr, jointSim uintptr, joint uintptr) {
	var jointDst uintptr
	_ = jointDst
	jointDst = b2CreateJointInGraph(tls, world, joint)
	memcpy(tls, jointDst, jointSim, uint64(196))
}

func b2RemoveJointFromGraph(tls *_Stack, world uintptr, bodyIdA int32, bodyIdB int32, colorIndex int32, localIndex int32) {
	var blockIndex, v2, v5 uint32_t
	var color, graph, movedJoint, movedJointSim, v1, v11, v13, v4, v7 uintptr
	var movedId, movedIndex, movedIndex1, v12, v8, v9 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = blockIndex, color, graph, movedId, movedIndex, movedIndex1, movedJoint, movedJointSim, v1, v11, v12, v13, v2, v4, v5, v7, v8, v9
	graph = world + 328
	if !(0 <= colorIndex && colorIndex < int32(_B2_GRAPH_COLOR_COUNT)) && b2InternalAssertFcn(tls, __ccgo_ts+3038, __ccgo_ts+2787, int32FromInt32(300)) != 0 {
		__builtin_trap(tls)
	}
	color = graph + uintptr(colorIndex)*56
	if colorIndex != int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1) {
		// May clear static bodies, no effect
		v1 = color
		v2 = uint32FromInt32(bodyIdA)
		blockIndex = v2 / uint32(64)
		if blockIndex >= (*b2BitSet)(unsafe.Pointer(v1)).BlockCount {
			goto _3
		}
		*(*uint64_t)(unsafe.Pointer((*b2BitSet)(unsafe.Pointer(v1)).Bits + uintptr(blockIndex)*8)) &= ^(uint64FromInt32(1) << (v2 % uint32FromInt32(64)))
	_3:
		;
		v4 = color
		v5 = uint32FromInt32(bodyIdB)
		blockIndex = v5 / uint32(64)
		if blockIndex >= (*b2BitSet)(unsafe.Pointer(v4)).BlockCount {
			goto _6
		}
		*(*uint64_t)(unsafe.Pointer((*b2BitSet)(unsafe.Pointer(v4)).Bits + uintptr(blockIndex)*8)) &= ^(uint64FromInt32(1) << (v5 % uint32FromInt32(64)))
	_6:
	}
	v7 = color + 32
	v8 = localIndex
	if !(0 <= v8 && v8 < (*b2JointSimArray)(unsafe.Pointer(v7)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(342)) != 0 {
		__builtin_trap(tls)
	}
	movedIndex = -int32(1)
	if v8 != (*b2JointSimArray)(unsafe.Pointer(v7)).Count-int32FromInt32(1) {
		movedIndex = (*b2JointSimArray)(unsafe.Pointer(v7)).Count - int32(1)
		*(*b2JointSim)(unsafe.Pointer((*b2JointSimArray)(unsafe.Pointer(v7)).Data + uintptr(v8)*196)) = *(*b2JointSim)(unsafe.Pointer((*b2JointSimArray)(unsafe.Pointer(v7)).Data + uintptr(movedIndex)*196))
	}
	*(*int32)(unsafe.Pointer(v7 + 8)) -= int32(1)
	v9 = movedIndex
	goto _10
_10:
	movedIndex1 = v9
	if movedIndex1 != -int32(1) {
		// Fix moved joint
		movedJointSim = (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Data + uintptr(localIndex)*196
		movedId = (*b2JointSim)(unsafe.Pointer(movedJointSim)).JointId
		v11 = world + 1104
		v12 = movedId
		if !(0 <= v12 && v12 < (*b2JointArray)(unsafe.Pointer(v11)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v13 = (*b2JointArray)(unsafe.Pointer(v11)).Data + uintptr(v12)*72
		goto _14
	_14:
		movedJoint = v13
		if !((*b2Joint)(unsafe.Pointer(movedJoint)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+3207, __ccgo_ts+2787, int32FromInt32(317)) != 0 {
			__builtin_trap(tls)
		}
		if !((*b2Joint)(unsafe.Pointer(movedJoint)).ColorIndex == colorIndex) && b2InternalAssertFcn(tls, __ccgo_ts+3243, __ccgo_ts+2787, int32FromInt32(318)) != 0 {
			__builtin_trap(tls)
		}
		if !((*b2Joint)(unsafe.Pointer(movedJoint)).LocalIndex == movedIndex1) && b2InternalAssertFcn(tls, __ccgo_ts+3280, __ccgo_ts+2787, int32FromInt32(319)) != 0 {
			__builtin_trap(tls)
		}
		(*b2Joint)(unsafe.Pointer(movedJoint)).LocalIndex = localIndex
	}
}

func b2DistanceJoint_SetLength(tls *_Stack, jointId JointId, length float32) {
	var base, joint uintptr
	var v1, v2, v3, v4, v6, v7 float32
	_, _, _, _, _, _, _, _ = base, joint, v1, v2, v3, v4, v6, v7
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	joint = base + 68
	v1 = length
	v2 = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
	v3 = float32(float32FromFloat32(100000) * b2_lengthUnitsPerMeter)
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
	(*b2DistanceJoint)(unsafe.Pointer(joint)).Length = v4
	(*b2DistanceJoint)(unsafe.Pointer(joint)).Impulse = float32FromFloat32(0)
	(*b2DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse = float32FromFloat32(0)
	(*b2DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse = float32FromFloat32(0)
}

func b2DistanceJoint_GetLength(tls *_Stack, jointId JointId) (r float32) {
	var base, joint uintptr
	_, _ = base, joint
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	joint = base + 68
	return (*b2DistanceJoint)(unsafe.Pointer(joint)).Length
}

func b2DistanceJoint_EnableLimit(tls *_Stack, jointId JointId, enableLimit uint8) {
	var base, joint uintptr
	_, _ = base, joint
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	joint = base + 68
	(*b2DistanceJoint)(unsafe.Pointer(joint)).EnableLimit = enableLimit
}

func b2DistanceJoint_IsLimitEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	return (*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.EnableLimit
}

func b2DistanceJoint_SetLengthRange(tls *_Stack, jointId JointId, minLength float32, maxLength float32) {
	var base, joint uintptr
	var v1, v10, v11, v13, v14, v15, v16, v17, v19, v2, v20, v21, v22, v24, v3, v4, v6, v7, v8, v9 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = base, joint, v1, v10, v11, v13, v14, v15, v16, v17, v19, v2, v20, v21, v22, v24, v3, v4, v6, v7, v8, v9
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	joint = base + 68
	v1 = minLength
	v2 = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
	v3 = float32(float32FromFloat32(100000) * b2_lengthUnitsPerMeter)
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
	minLength = v4
	v8 = maxLength
	v9 = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
	v10 = float32(float32FromFloat32(100000) * b2_lengthUnitsPerMeter)
	if v8 < v9 {
		v13 = v9
	} else {
		if v8 > v10 {
			v14 = v10
		} else {
			v14 = v8
		}
		v13 = v14
	}
	v11 = v13
	goto _12
_12:
	maxLength = v11
	v15 = minLength
	v16 = maxLength
	if v15 < v16 {
		v19 = v15
	} else {
		v19 = v16
	}
	v17 = v19
	goto _18
_18:
	(*b2DistanceJoint)(unsafe.Pointer(joint)).MinLength = v17
	v20 = minLength
	v21 = maxLength
	if v20 > v21 {
		v24 = v20
	} else {
		v24 = v21
	}
	v22 = v24
	goto _23
_23:
	(*b2DistanceJoint)(unsafe.Pointer(joint)).MaxLength = v22
	(*b2DistanceJoint)(unsafe.Pointer(joint)).Impulse = float32FromFloat32(0)
	(*b2DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse = float32FromFloat32(0)
	(*b2DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse = float32FromFloat32(0)
}

func b2DistanceJoint_GetMinLength(tls *_Stack, jointId JointId) (r float32) {
	var base, joint uintptr
	_, _ = base, joint
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	joint = base + 68
	return (*b2DistanceJoint)(unsafe.Pointer(joint)).MinLength
}

func b2DistanceJoint_GetMaxLength(tls *_Stack, jointId JointId) (r float32) {
	var base, joint uintptr
	_, _ = base, joint
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	joint = base + 68
	return (*b2DistanceJoint)(unsafe.Pointer(joint)).MaxLength
}

func b2DistanceJoint_GetCurrentLength(tls *_Stack, jointId JointId) (r float32) {
	var base, world uintptr
	var d, pA, pB, v10, v11, v13, v2, v3, v6, v7, v9 Vec2
	var length, x, y, v14 float32
	var transformA, transformB, v1, v5 Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = base, d, length, pA, pB, transformA, transformB, world, x, y, v1, v10, v11, v13, v14, v2, v3, v5, v6, v7, v9
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if !(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)) && b2InternalAssertFcn(tls, __ccgo_ts+1152, __ccgo_ts+4580, int32FromInt32(84)) != 0 {
		__builtin_trap(tls)
	}
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return float32FromFloat32(0)
	}
	transformA = b2GetBodyTransform(tls, world, (*b2JointSim)(unsafe.Pointer(base)).BodyIdA)
	transformB = b2GetBodyTransform(tls, world, (*b2JointSim)(unsafe.Pointer(base)).BodyIdB)
	v1 = transformA
	v2 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorA
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	pA = v3
	v5 = transformB
	v6 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorB
	x = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y) + v5.P.X
	y = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y) + v5.P.Y
	v7 = Vec2{
		X: x,
		Y: y,
	}
	goto _8
_8:
	pB = v7
	v9 = pB
	v10 = pA
	v11 = Vec2{
		X: v9.X - v10.X,
		Y: v9.Y - v10.Y,
	}
	goto _12
_12:
	d = v11
	v13 = d
	v14 = sqrtf(tls, float32(v13.X*v13.X)+float32(v13.Y*v13.Y))
	goto _15
_15:
	length = v14
	return length
}

func b2DistanceJoint_EnableSpring(tls *_Stack, jointId JointId, enableSpring uint8) {
	var base uintptr
	_ = base
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	(*b2JointSim)(unsafe.Pointer(base)).ｆ__ccgo13_68.DistanceJoint.EnableSpring = enableSpring
}

func b2DistanceJoint_IsSpringEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var base uintptr
	_ = base
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	return (*b2JointSim)(unsafe.Pointer(base)).ｆ__ccgo13_68.DistanceJoint.EnableSpring
}

func b2DistanceJoint_SetSpringHertz(tls *_Stack, jointId JointId, hertz float32) {
	var base uintptr
	_ = base
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	(*b2JointSim)(unsafe.Pointer(base)).ｆ__ccgo13_68.DistanceJoint.Hertz = hertz
}

func b2DistanceJoint_SetSpringDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	var base uintptr
	_ = base
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	(*b2JointSim)(unsafe.Pointer(base)).ｆ__ccgo13_68.DistanceJoint.DampingRatio = dampingRatio
}

func b2DistanceJoint_GetSpringHertz(tls *_Stack, jointId JointId) (r float32) {
	var base, joint uintptr
	_, _ = base, joint
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	joint = base + 68
	return (*b2DistanceJoint)(unsafe.Pointer(joint)).Hertz
}

func b2DistanceJoint_GetSpringDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var base, joint uintptr
	_, _ = base, joint
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	joint = base + 68
	return (*b2DistanceJoint)(unsafe.Pointer(joint)).DampingRatio
}

func b2DistanceJoint_EnableMotor(tls *_Stack, jointId JointId, enableMotor uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	if enableMotor != (*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.EnableMotor {
		(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.EnableMotor = enableMotor
		(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.MotorImpulse = float32FromFloat32(0)
	}
}

func b2DistanceJoint_IsMotorEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	return (*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.EnableMotor
}

func b2DistanceJoint_SetMotorSpeed(tls *_Stack, jointId JointId, motorSpeed float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.MotorSpeed = motorSpeed
}

func b2DistanceJoint_GetMotorSpeed(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	return (*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.MotorSpeed
}

func b2DistanceJoint_GetMotorForce(tls *_Stack, jointId JointId) (r float32) {
	var base, world uintptr
	_, _ = base, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	return float32((*b2World)(unsafe.Pointer(world)).Inv_h * (*b2JointSim)(unsafe.Pointer(base)).ｆ__ccgo13_68.DistanceJoint.MotorImpulse)
}

func b2DistanceJoint_SetMaxMotorForce(tls *_Stack, jointId JointId, force float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.MaxMotorForce = force
}

func b2DistanceJoint_GetMaxMotorForce(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_distanceJoint))
	return (*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.MaxMotorForce
}

func b2GetDistanceJointForce(tls *_Stack, world uintptr, base uintptr) (r Vec2) {
	var axis, d, n, pA, pB, v10, v11, v13, v14, v17, v18, v2, v3, v6, v7, v9 Vec2
	var force, invLength, length, x, y, v16 float32
	var joint uintptr
	var transformA, transformB, v1, v5 Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axis, d, force, invLength, joint, length, n, pA, pB, transformA, transformB, x, y, v1, v10, v11, v13, v14, v16, v17, v18, v2, v3, v5, v6, v7, v9
	joint = base + 68
	transformA = b2GetBodyTransform(tls, world, (*b2JointSim)(unsafe.Pointer(base)).BodyIdA)
	transformB = b2GetBodyTransform(tls, world, (*b2JointSim)(unsafe.Pointer(base)).BodyIdB)
	v1 = transformA
	v2 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorA
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	pA = v3
	v5 = transformB
	v6 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorB
	x = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y) + v5.P.X
	y = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y) + v5.P.Y
	v7 = Vec2{
		X: x,
		Y: y,
	}
	goto _8
_8:
	pB = v7
	v9 = pB
	v10 = pA
	v11 = Vec2{
		X: v9.X - v10.X,
		Y: v9.Y - v10.Y,
	}
	goto _12
_12:
	d = v11
	v13 = d
	length = sqrtf(tls, float32(v13.X*v13.X)+float32(v13.Y*v13.Y))
	if length < float32FromFloat32(1.1920928955078125e-07) {
		v14 = Vec2{}
		goto _15
	}
	invLength = float32FromFloat32(1) / length
	n = Vec2{
		X: float32(invLength * v13.X),
		Y: float32(invLength * v13.Y),
	}
	v14 = n
	goto _15
_15:
	axis = v14
	force = float32(((*b2DistanceJoint)(unsafe.Pointer(joint)).Impulse + (*b2DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b2DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse + (*b2DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse) * (*b2World)(unsafe.Pointer(world)).Inv_h)
	v16 = force
	v17 = axis
	v18 = Vec2{
		X: float32(v16 * v17.X),
		Y: float32(v16 * v17.Y),
	}
	goto _19
_19:
	return v18
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
	var a11, a21, a31, crA, crB, iA, iB, invLength, k, length, mA, mB, omega, v60, v64, v66, v67, v68 float32
	var axis, n, rA, rB, separation, v27, v28, v29, v32, v33, v35, v36, v37, v40, v41, v43, v44, v45, v47, v48, v49, v51, v52, v53, v55, v56, v58, v59, v62, v63 Vec2
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var idA, idB, localIndexA, localIndexB, v10, v14, v18, v2, v22, v25, v26, v6 int32
	var v31, v39 Rot
	var v69 b2Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a21, a31, axis, bodyA, bodyB, bodySimA, bodySimB, crA, crB, iA, iB, idA, idB, invLength, joint, k, length, localIndexA, localIndexB, mA, mB, n, omega, rA, rB, separation, setA, setB, world, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v28, v29, v3, v31, v32, v33, v35, v36, v37, v39, v40, v41, v43, v44, v45, v47, v48, v49, v5, v51, v52, v53, v55, v56, v58, v59, v6, v60, v62, v63, v64, v66, v67, v68, v69, v7, v9
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_distanceJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+4612, __ccgo_ts+4580, int32FromInt32(217)) != 0 {
		__builtin_trap(tls)
	}
	// chase body id to the solver set where the body lives
	idA = (*b2JointSim)(unsafe.Pointer(base)).BodyIdA
	idB = (*b2JointSim)(unsafe.Pointer(base)).BodyIdB
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	v1 = world + 1024
	v2 = idA
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	bodyA = v3
	v5 = world + 1024
	v6 = idB
	if !(0 <= v6 && v6 < (*b2BodyArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodyArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*128
	goto _8
_8:
	bodyB = v7
	if !((*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+4643, __ccgo_ts+4580, int32FromInt32(227)) != 0 {
		__builtin_trap(tls)
	}
	v9 = world + 1064
	v10 = (*b2Body)(unsafe.Pointer(bodyA)).SetIndex
	if !(0 <= v10 && v10 < (*b2SolverSetArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v11 = (*b2SolverSetArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*88
	goto _12
_12:
	setA = v11
	v13 = world + 1064
	v14 = (*b2Body)(unsafe.Pointer(bodyB)).SetIndex
	if !(0 <= v14 && v14 < (*b2SolverSetArray)(unsafe.Pointer(v13)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v15 = (*b2SolverSetArray)(unsafe.Pointer(v13)).Data + uintptr(v14)*88
	goto _16
_16:
	setB = v15
	localIndexA = (*b2Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b2Body)(unsafe.Pointer(bodyB)).LocalIndex
	v17 = setA
	v18 = localIndexA
	if !(0 <= v18 && v18 < (*b2BodySimArray)(unsafe.Pointer(v17)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	v19 = (*b2BodySimArray)(unsafe.Pointer(v17)).Data + uintptr(v18)*100
	goto _20
_20:
	bodySimA = v19
	v21 = setB
	v22 = localIndexB
	if !(0 <= v22 && v22 < (*b2BodySimArray)(unsafe.Pointer(v21)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	v23 = (*b2BodySimArray)(unsafe.Pointer(v21)).Data + uintptr(v22)*100
	goto _24
_24:
	bodySimB = v23
	mA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvMass
	iA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvInertia
	mB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvMass
	iB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvInertia
	(*b2JointSim)(unsafe.Pointer(base)).InvMassA = mA
	(*b2JointSim)(unsafe.Pointer(base)).InvMassB = mB
	(*b2JointSim)(unsafe.Pointer(base)).InvIA = iA
	(*b2JointSim)(unsafe.Pointer(base)).InvIB = iB
	joint = base + 68
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) {
		v25 = localIndexA
	} else {
		v25 = -int32(1)
	}
	(*b2DistanceJoint)(unsafe.Pointer(joint)).IndexA = v25
	if (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) {
		v26 = localIndexB
	} else {
		v26 = -int32(1)
	}
	(*b2DistanceJoint)(unsafe.Pointer(joint)).IndexB = v26
	// initial anchors in world space
	v27 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorA
	v28 = (*b2BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v29 = Vec2{
		X: v27.X - v28.X,
		Y: v27.Y - v28.Y,
	}
	goto _30
_30:
	v31 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v32 = v29
	v33 = Vec2{
		X: float32(v31.C*v32.X) - float32(v31.S*v32.Y),
		Y: float32(v31.S*v32.X) + float32(v31.C*v32.Y),
	}
	goto _34
_34:
	(*b2DistanceJoint)(unsafe.Pointer(joint)).AnchorA = v33
	v35 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorB
	v36 = (*b2BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v37 = Vec2{
		X: v35.X - v36.X,
		Y: v35.Y - v36.Y,
	}
	goto _38
_38:
	v39 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v40 = v37
	v41 = Vec2{
		X: float32(v39.C*v40.X) - float32(v39.S*v40.Y),
		Y: float32(v39.S*v40.X) + float32(v39.C*v40.Y),
	}
	goto _42
_42:
	(*b2DistanceJoint)(unsafe.Pointer(joint)).AnchorB = v41
	v43 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Center
	v44 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Center
	v45 = Vec2{
		X: v43.X - v44.X,
		Y: v43.Y - v44.Y,
	}
	goto _46
_46:
	(*b2DistanceJoint)(unsafe.Pointer(joint)).DeltaCenter = v45
	rA = (*b2DistanceJoint)(unsafe.Pointer(joint)).AnchorA
	rB = (*b2DistanceJoint)(unsafe.Pointer(joint)).AnchorB
	v47 = rB
	v48 = rA
	v49 = Vec2{
		X: v47.X - v48.X,
		Y: v47.Y - v48.Y,
	}
	goto _50
_50:
	v51 = v49
	v52 = (*b2DistanceJoint)(unsafe.Pointer(joint)).DeltaCenter
	v53 = Vec2{
		X: v51.X + v52.X,
		Y: v51.Y + v52.Y,
	}
	goto _54
_54:
	separation = v53
	v55 = separation
	length = sqrtf(tls, float32(v55.X*v55.X)+float32(v55.Y*v55.Y))
	if length < float32FromFloat32(1.1920928955078125e-07) {
		v56 = Vec2{}
		goto _57
	}
	invLength = float32FromFloat32(1) / length
	n = Vec2{
		X: float32(invLength * v55.X),
		Y: float32(invLength * v55.Y),
	}
	v56 = n
	goto _57
_57:
	axis = v56
	v58 = rA
	v59 = axis
	v60 = float32(v58.X*v59.Y) - float32(v58.Y*v59.X)
	goto _61
_61:
	// compute effective mass
	crA = v60
	v62 = rB
	v63 = axis
	v64 = float32(v62.X*v63.Y) - float32(v62.Y*v63.X)
	goto _65
_65:
	crB = v64
	k = mA + mB + float32(float32(iA*crA)*crA) + float32(float32(iB*crB)*crB)
	if k > float32FromFloat32(0) {
		v66 = float32FromFloat32(1) / k
	} else {
		v66 = float32FromFloat32(0)
	}
	(*b2DistanceJoint)(unsafe.Pointer(joint)).AxialMass = v66
	v67 = (*b2DistanceJoint)(unsafe.Pointer(joint)).Hertz
	v68 = (*b2StepContext)(unsafe.Pointer(context)).H
	if v67 == float32FromFloat32(0) {
		v69 = b2Softness{}
		goto _70
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v67)
	a11 = float32(float32FromFloat32(2)*(*b2DistanceJoint)(unsafe.Pointer(joint)).DampingRatio) + float32(v68*omega)
	a21 = float32(float32(v68*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v69 = b2Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _70
_70:
	(*b2DistanceJoint)(unsafe.Pointer(joint)).DistanceSoftness = v69
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
	var P, axis, ds, n, rA, rB, separation, v11, v12, v13, v15, v16, v17, v19, v20, v21, v23, v24, v25, v27, v28, v31, v32, v34, v36, v37, v39, v4, v40, v43, v45, v46, v48, v49, v5, v8, v9 Vec2
	var axialImpulse, iA, iB, invLength, length, mA, mB, v30, v35, v41, v44, v50 float32
	var joint, stateA, stateB, v1, v2 uintptr
	var v3, v7 Rot
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = P, axialImpulse, axis, ds, iA, iB, invLength, joint, length, mA, mB, n, rA, rB, separation, stateA, stateB, v1, v11, v12, v13, v15, v16, v17, v19, v2, v20, v21, v23, v24, v25, v27, v28, v3, v30, v31, v32, v34, v35, v36, v37, v39, v4, v40, v41, v43, v44, v45, v46, v48, v49, v5, v50, v7, v8, v9
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_distanceJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+4612, __ccgo_ts+4580, int32FromInt32(282)) != 0 {
		__builtin_trap(tls)
	}
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Pointer(bp)) = b2_identityBodyState5
	joint = base + 68
	if (*b2DistanceJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v1 = bp
	} else {
		v1 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2DistanceJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v1
	if (*b2DistanceJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2DistanceJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v2
	v3 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v4 = (*b2DistanceJoint)(unsafe.Pointer(joint)).AnchorA
	v5 = Vec2{
		X: float32(v3.C*v4.X) - float32(v3.S*v4.Y),
		Y: float32(v3.S*v4.X) + float32(v3.C*v4.Y),
	}
	goto _6
_6:
	rA = v5
	v7 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v8 = (*b2DistanceJoint)(unsafe.Pointer(joint)).AnchorB
	v9 = Vec2{
		X: float32(v7.C*v8.X) - float32(v7.S*v8.Y),
		Y: float32(v7.S*v8.X) + float32(v7.C*v8.Y),
	}
	goto _10
_10:
	rB = v9
	v11 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
	v12 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaPosition
	v13 = Vec2{
		X: v11.X - v12.X,
		Y: v11.Y - v12.Y,
	}
	goto _14
_14:
	v15 = rB
	v16 = rA
	v17 = Vec2{
		X: v15.X - v16.X,
		Y: v15.Y - v16.Y,
	}
	goto _18
_18:
	v19 = v13
	v20 = v17
	v21 = Vec2{
		X: v19.X + v20.X,
		Y: v19.Y + v20.Y,
	}
	goto _22
_22:
	ds = v21
	v23 = (*b2DistanceJoint)(unsafe.Pointer(joint)).DeltaCenter
	v24 = ds
	v25 = Vec2{
		X: v23.X + v24.X,
		Y: v23.Y + v24.Y,
	}
	goto _26
_26:
	separation = v25
	v27 = separation
	length = sqrtf(tls, float32(v27.X*v27.X)+float32(v27.Y*v27.Y))
	if length < float32FromFloat32(1.1920928955078125e-07) {
		v28 = Vec2{}
		goto _29
	}
	invLength = float32FromFloat32(1) / length
	n = Vec2{
		X: float32(invLength * v27.X),
		Y: float32(invLength * v27.Y),
	}
	v28 = n
	goto _29
_29:
	axis = v28
	axialImpulse = (*b2DistanceJoint)(unsafe.Pointer(joint)).Impulse + (*b2DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b2DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse + (*b2DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse
	v30 = axialImpulse
	v31 = axis
	v32 = Vec2{
		X: float32(v30 * v31.X),
		Y: float32(v30 * v31.Y),
	}
	goto _33
_33:
	P = v32
	v34 = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	v35 = mA
	v36 = P
	v37 = Vec2{
		X: v34.X - float32(v35*v36.X),
		Y: v34.Y - float32(v35*v36.Y),
	}
	goto _38
_38:
	(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = v37
	v39 = rA
	v40 = P
	v41 = float32(v39.X*v40.Y) - float32(v39.Y*v40.X)
	goto _42
_42:
	*(*float32)(unsafe.Pointer(stateA + 8)) -= float32(iA * v41)
	v43 = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	v44 = mB
	v45 = P
	v46 = Vec2{
		X: v43.X + float32(v44*v45.X),
		Y: v43.Y + float32(v44*v45.Y),
	}
	goto _47
_47:
	(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = v46
	v48 = rB
	v49 = P
	v50 = float32(v48.X*v49.Y) - float32(v48.Y*v49.X)
	goto _51
_51:
	*(*float32)(unsafe.Pointer(stateB + 8)) += float32(iB * v50)
}

func b2SolveDistanceJoint(tls *_Stack, base uintptr, context uintptr, useBias uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var C, C1, C2, C3, Cdot, Cdot1, Cdot2, Cdot3, Cdot4, bias, bias1, bias2, bias3, iA, iB, impulse, impulse1, impulse2, impulse3, impulse4, impulseCoeff, impulseScale, impulseScale1, invLength, length, length1, m, mA, mB, massCoeff, massScale, massScale1, maxImpulse, newImpulse, newImpulse1, oldImpulse, wA, wB, v101, v103, v104, v105, v107, v108, v113, v119, v122, v128, v134, v138, v152, v154, v155, v156, v158, v159, v164, v170, v173, v179, v185, v189, v203, v205, v206, v207, v208, v210, v211, v212, v217, v223, v226, v232, v238, v242, v256, v258, v263, v269, v272, v278, v28, v37, v41, v55, v57, v62, v68, v71, v77, v83, v87 float32
	var P, P1, P2, P3, P4, axis, ds, n, rA, rB, separation, vA, vB, vr, vr1, vr2, vr3, vr4, v100, v109, v11, v110, v112, v114, v115, v117, v118, v12, v121, v123, v124, v126, v127, v13, v130, v131, v132, v135, v136, v139, v140, v142, v143, v144, v146, v147, v148, v15, v150, v151, v16, v160, v161, v163, v165, v166, v168, v169, v17, v172, v174, v175, v177, v178, v181, v182, v183, v186, v187, v19, v190, v191, v193, v194, v195, v197, v198, v199, v20, v201, v202, v21, v213, v214, v216, v218, v219, v221, v222, v225, v227, v228, v23, v230, v231, v234, v235, v236, v239, v24, v240, v243, v244, v246, v247, v248, v25, v250, v251, v252, v254, v255, v259, v260, v262, v264, v265, v267, v268, v27, v271, v273, v274, v276, v277, v30, v31, v33, v34, v35, v38, v39, v4, v42, v43, v45, v46, v47, v49, v5, v50, v51, v53, v54, v58, v59, v61, v63, v64, v66, v67, v70, v72, v73, v75, v76, v79, v8, v80, v81, v84, v85, v88, v89, v9, v91, v92, v93, v95, v96, v97, v99 Vec2
	var joint, stateA, stateB, v1, v2 uintptr
	var v3, v7 Rot
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, C1, C2, C3, Cdot, Cdot1, Cdot2, Cdot3, Cdot4, P, P1, P2, P3, P4, axis, bias, bias1, bias2, bias3, ds, iA, iB, impulse, impulse1, impulse2, impulse3, impulse4, impulseCoeff, impulseScale, impulseScale1, invLength, joint, length, length1, m, mA, mB, massCoeff, massScale, massScale1, maxImpulse, n, newImpulse, newImpulse1, oldImpulse, rA, rB, separation, stateA, stateB, vA, vB, vr, vr1, vr2, vr3, vr4, wA, wB, v1, v100, v101, v103, v104, v105, v107, v108, v109, v11, v110, v112, v113, v114, v115, v117, v118, v119, v12, v121, v122, v123, v124, v126, v127, v128, v13, v130, v131, v132, v134, v135, v136, v138, v139, v140, v142, v143, v144, v146, v147, v148, v15, v150, v151, v152, v154, v155, v156, v158, v159, v16, v160, v161, v163, v164, v165, v166, v168, v169, v17, v170, v172, v173, v174, v175, v177, v178, v179, v181, v182, v183, v185, v186, v187, v189, v19, v190, v191, v193, v194, v195, v197, v198, v199, v2, v20, v201, v202, v203, v205, v206, v207, v208, v21, v210, v211, v212, v213, v214, v216, v217, v218, v219, v221, v222, v223, v225, v226, v227, v228, v23, v230, v231, v232, v234, v235, v236, v238, v239, v24, v240, v242, v243, v244, v246, v247, v248, v25, v250, v251, v252, v254, v255, v256, v258, v259, v260, v262, v263, v264, v265, v267, v268, v269, v27, v271, v272, v273, v274, v276, v277, v278, v28, v3, v30, v31, v33, v34, v35, v37, v38, v39, v4, v41, v42, v43, v45, v46, v47, v49, v5, v50, v51, v53, v54, v55, v57, v58, v59, v61, v62, v63, v64, v66, v67, v68, v7, v70, v71, v72, v73, v75, v76, v77, v79, v8, v80, v81, v83, v84, v85, v87, v88, v89, v9, v91, v92, v93, v95, v96, v97, v99
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_distanceJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+4612, __ccgo_ts+4580, int32FromInt32(314)) != 0 {
		__builtin_trap(tls)
	}
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Pointer(bp)) = b2_identityBodyState5
	joint = base + 68
	if (*b2DistanceJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v1 = bp
	} else {
		v1 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2DistanceJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v1
	if (*b2DistanceJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2DistanceJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v2
	vA = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	v3 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v4 = (*b2DistanceJoint)(unsafe.Pointer(joint)).AnchorA
	v5 = Vec2{
		X: float32(v3.C*v4.X) - float32(v3.S*v4.Y),
		Y: float32(v3.S*v4.X) + float32(v3.C*v4.Y),
	}
	goto _6
_6:
	// current anchors
	rA = v5
	v7 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v8 = (*b2DistanceJoint)(unsafe.Pointer(joint)).AnchorB
	v9 = Vec2{
		X: float32(v7.C*v8.X) - float32(v7.S*v8.Y),
		Y: float32(v7.S*v8.X) + float32(v7.C*v8.Y),
	}
	goto _10
_10:
	rB = v9
	v11 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
	v12 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaPosition
	v13 = Vec2{
		X: v11.X - v12.X,
		Y: v11.Y - v12.Y,
	}
	goto _14
_14:
	v15 = rB
	v16 = rA
	v17 = Vec2{
		X: v15.X - v16.X,
		Y: v15.Y - v16.Y,
	}
	goto _18
_18:
	v19 = v13
	v20 = v17
	v21 = Vec2{
		X: v19.X + v20.X,
		Y: v19.Y + v20.Y,
	}
	goto _22
_22:
	// current separation
	ds = v21
	v23 = (*b2DistanceJoint)(unsafe.Pointer(joint)).DeltaCenter
	v24 = ds
	v25 = Vec2{
		X: v23.X + v24.X,
		Y: v23.Y + v24.Y,
	}
	goto _26
_26:
	separation = v25
	v27 = separation
	v28 = sqrtf(tls, float32(v27.X*v27.X)+float32(v27.Y*v27.Y))
	goto _29
_29:
	length1 = v28
	v30 = separation
	length = sqrtf(tls, float32(v30.X*v30.X)+float32(v30.Y*v30.Y))
	if length < float32FromFloat32(1.1920928955078125e-07) {
		v31 = Vec2{}
		goto _32
	}
	invLength = float32FromFloat32(1) / length
	n = Vec2{
		X: float32(invLength * v30.X),
		Y: float32(invLength * v30.Y),
	}
	v31 = n
	goto _32
_32:
	axis = v31
	// joint is soft if
	// - spring is enabled
	// - and (joint limit is disabled or limits are not equal)
	if (*b2DistanceJoint)(unsafe.Pointer(joint)).EnableSpring != 0 && ((*b2DistanceJoint)(unsafe.Pointer(joint)).MinLength < (*b2DistanceJoint)(unsafe.Pointer(joint)).MaxLength || int32FromUint8((*b2DistanceJoint)(unsafe.Pointer(joint)).EnableLimit) == false1) {
		// spring
		if (*b2DistanceJoint)(unsafe.Pointer(joint)).Hertz > float32FromFloat32(0) {
			v33 = vB
			v34 = vA
			v35 = Vec2{
				X: v33.X - v34.X,
				Y: v33.Y - v34.Y,
			}
			goto _36
		_36:
			v37 = wB
			v38 = rB
			v39 = Vec2{
				X: float32(-v37 * v38.Y),
				Y: float32(v37 * v38.X),
			}
			goto _40
		_40:
			v41 = wA
			v42 = rA
			v43 = Vec2{
				X: float32(-v41 * v42.Y),
				Y: float32(v41 * v42.X),
			}
			goto _44
		_44:
			v45 = v39
			v46 = v43
			v47 = Vec2{
				X: v45.X - v46.X,
				Y: v45.Y - v46.Y,
			}
			goto _48
		_48:
			v49 = v35
			v50 = v47
			v51 = Vec2{
				X: v49.X + v50.X,
				Y: v49.Y + v50.Y,
			}
			goto _52
		_52:
			// Cdot = dot(u, v + cross(w, r))
			vr = v51
			v53 = axis
			v54 = vr
			v55 = float32(v53.X*v54.X) + float32(v53.Y*v54.Y)
			goto _56
		_56:
			Cdot = v55
			C = length1 - (*b2DistanceJoint)(unsafe.Pointer(joint)).Length
			bias = float32((*b2DistanceJoint)(unsafe.Pointer(joint)).DistanceSoftness.BiasRate * C)
			m = float32((*b2DistanceJoint)(unsafe.Pointer(joint)).DistanceSoftness.MassScale * (*b2DistanceJoint)(unsafe.Pointer(joint)).AxialMass)
			impulse = float32(-m*(Cdot+bias)) - float32((*b2DistanceJoint)(unsafe.Pointer(joint)).DistanceSoftness.ImpulseScale*(*b2DistanceJoint)(unsafe.Pointer(joint)).Impulse)
			*(*float32)(unsafe.Pointer(joint + 28)) += impulse
			v57 = impulse
			v58 = axis
			v59 = Vec2{
				X: float32(v57 * v58.X),
				Y: float32(v57 * v58.Y),
			}
			goto _60
		_60:
			P = v59
			v61 = vA
			v62 = mA
			v63 = P
			v64 = Vec2{
				X: v61.X - float32(v62*v63.X),
				Y: v61.Y - float32(v62*v63.Y),
			}
			goto _65
		_65:
			vA = v64
			v66 = rA
			v67 = P
			v68 = float32(v66.X*v67.Y) - float32(v66.Y*v67.X)
			goto _69
		_69:
			wA = wA - float32(iA*v68)
			v70 = vB
			v71 = mB
			v72 = P
			v73 = Vec2{
				X: v70.X + float32(v71*v72.X),
				Y: v70.Y + float32(v71*v72.Y),
			}
			goto _74
		_74:
			vB = v73
			v75 = rB
			v76 = P
			v77 = float32(v75.X*v76.Y) - float32(v75.Y*v76.X)
			goto _78
		_78:
			wB = wB + float32(iB*v77)
		}
		if (*b2DistanceJoint)(unsafe.Pointer(joint)).EnableLimit != 0 {
			// lower limit
			v79 = vB
			v80 = vA
			v81 = Vec2{
				X: v79.X - v80.X,
				Y: v79.Y - v80.Y,
			}
			goto _82
		_82:
			v83 = wB
			v84 = rB
			v85 = Vec2{
				X: float32(-v83 * v84.Y),
				Y: float32(v83 * v84.X),
			}
			goto _86
		_86:
			v87 = wA
			v88 = rA
			v89 = Vec2{
				X: float32(-v87 * v88.Y),
				Y: float32(v87 * v88.X),
			}
			goto _90
		_90:
			v91 = v85
			v92 = v89
			v93 = Vec2{
				X: v91.X - v92.X,
				Y: v91.Y - v92.Y,
			}
			goto _94
		_94:
			v95 = v81
			v96 = v93
			v97 = Vec2{
				X: v95.X + v96.X,
				Y: v95.Y + v96.Y,
			}
			goto _98
		_98:
			vr1 = v97
			v99 = axis
			v100 = vr1
			v101 = float32(v99.X*v100.X) + float32(v99.Y*v100.Y)
			goto _102
		_102:
			Cdot1 = v101
			C1 = length1 - (*b2DistanceJoint)(unsafe.Pointer(joint)).MinLength
			bias1 = float32FromFloat32(0)
			massCoeff = float32FromFloat32(1)
			impulseCoeff = float32FromFloat32(0)
			if C1 > float32FromFloat32(0) {
				// speculative
				bias1 = float32(C1 * (*b2StepContext)(unsafe.Pointer(context)).Inv_h)
			} else {
				if useBias != 0 {
					bias1 = float32((*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * C1)
					massCoeff = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
					impulseCoeff = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
				}
			}
			impulse1 = float32(float32(-massCoeff*(*b2DistanceJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot1+bias1)) - float32(impulseCoeff*(*b2DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse)
			v103 = float32FromFloat32(0)
			v104 = (*b2DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse + impulse1
			if v103 > v104 {
				v107 = v103
			} else {
				v107 = v104
			}
			v105 = v107
			goto _106
		_106:
			newImpulse = v105
			impulse1 = newImpulse - (*b2DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse
			(*b2DistanceJoint)(unsafe.Pointer(joint)).LowerImpulse = newImpulse
			v108 = impulse1
			v109 = axis
			v110 = Vec2{
				X: float32(v108 * v109.X),
				Y: float32(v108 * v109.Y),
			}
			goto _111
		_111:
			P1 = v110
			v112 = vA
			v113 = mA
			v114 = P1
			v115 = Vec2{
				X: v112.X - float32(v113*v114.X),
				Y: v112.Y - float32(v113*v114.Y),
			}
			goto _116
		_116:
			vA = v115
			v117 = rA
			v118 = P1
			v119 = float32(v117.X*v118.Y) - float32(v117.Y*v118.X)
			goto _120
		_120:
			wA = wA - float32(iA*v119)
			v121 = vB
			v122 = mB
			v123 = P1
			v124 = Vec2{
				X: v121.X + float32(v122*v123.X),
				Y: v121.Y + float32(v122*v123.Y),
			}
			goto _125
		_125:
			vB = v124
			v126 = rB
			v127 = P1
			v128 = float32(v126.X*v127.Y) - float32(v126.Y*v127.X)
			goto _129
		_129:
			wB = wB + float32(iB*v128)
			// upper
			v130 = vA
			v131 = vB
			v132 = Vec2{
				X: v130.X - v131.X,
				Y: v130.Y - v131.Y,
			}
			goto _133
		_133:
			v134 = wA
			v135 = rA
			v136 = Vec2{
				X: float32(-v134 * v135.Y),
				Y: float32(v134 * v135.X),
			}
			goto _137
		_137:
			v138 = wB
			v139 = rB
			v140 = Vec2{
				X: float32(-v138 * v139.Y),
				Y: float32(v138 * v139.X),
			}
			goto _141
		_141:
			v142 = v136
			v143 = v140
			v144 = Vec2{
				X: v142.X - v143.X,
				Y: v142.Y - v143.Y,
			}
			goto _145
		_145:
			v146 = v132
			v147 = v144
			v148 = Vec2{
				X: v146.X + v147.X,
				Y: v146.Y + v147.Y,
			}
			goto _149
		_149:
			vr2 = v148
			v150 = axis
			v151 = vr2
			v152 = float32(v150.X*v151.X) + float32(v150.Y*v151.Y)
			goto _153
		_153:
			Cdot2 = v152
			C2 = (*b2DistanceJoint)(unsafe.Pointer(joint)).MaxLength - length1
			bias2 = float32FromFloat32(0)
			massScale = float32FromFloat32(1)
			impulseScale = float32FromFloat32(0)
			if C2 > float32FromFloat32(0) {
				// speculative
				bias2 = float32(C2 * (*b2StepContext)(unsafe.Pointer(context)).Inv_h)
			} else {
				if useBias != 0 {
					bias2 = float32((*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * C2)
					massScale = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
					impulseScale = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
				}
			}
			impulse2 = float32(float32(-massScale*(*b2DistanceJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot2+bias2)) - float32(impulseScale*(*b2DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse)
			v154 = float32FromFloat32(0)
			v155 = (*b2DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse + impulse2
			if v154 > v155 {
				v158 = v154
			} else {
				v158 = v155
			}
			v156 = v158
			goto _157
		_157:
			newImpulse1 = v156
			impulse2 = newImpulse1 - (*b2DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse
			(*b2DistanceJoint)(unsafe.Pointer(joint)).UpperImpulse = newImpulse1
			v159 = -impulse2
			v160 = axis
			v161 = Vec2{
				X: float32(v159 * v160.X),
				Y: float32(v159 * v160.Y),
			}
			goto _162
		_162:
			P2 = v161
			v163 = vA
			v164 = mA
			v165 = P2
			v166 = Vec2{
				X: v163.X - float32(v164*v165.X),
				Y: v163.Y - float32(v164*v165.Y),
			}
			goto _167
		_167:
			vA = v166
			v168 = rA
			v169 = P2
			v170 = float32(v168.X*v169.Y) - float32(v168.Y*v169.X)
			goto _171
		_171:
			wA = wA - float32(iA*v170)
			v172 = vB
			v173 = mB
			v174 = P2
			v175 = Vec2{
				X: v172.X + float32(v173*v174.X),
				Y: v172.Y + float32(v173*v174.Y),
			}
			goto _176
		_176:
			vB = v175
			v177 = rB
			v178 = P2
			v179 = float32(v177.X*v178.Y) - float32(v177.Y*v178.X)
			goto _180
		_180:
			wB = wB + float32(iB*v179)
		}
		if (*b2DistanceJoint)(unsafe.Pointer(joint)).EnableMotor != 0 {
			v181 = vB
			v182 = vA
			v183 = Vec2{
				X: v181.X - v182.X,
				Y: v181.Y - v182.Y,
			}
			goto _184
		_184:
			v185 = wB
			v186 = rB
			v187 = Vec2{
				X: float32(-v185 * v186.Y),
				Y: float32(v185 * v186.X),
			}
			goto _188
		_188:
			v189 = wA
			v190 = rA
			v191 = Vec2{
				X: float32(-v189 * v190.Y),
				Y: float32(v189 * v190.X),
			}
			goto _192
		_192:
			v193 = v187
			v194 = v191
			v195 = Vec2{
				X: v193.X - v194.X,
				Y: v193.Y - v194.Y,
			}
			goto _196
		_196:
			v197 = v183
			v198 = v195
			v199 = Vec2{
				X: v197.X + v198.X,
				Y: v197.Y + v198.Y,
			}
			goto _200
		_200:
			vr3 = v199
			v201 = axis
			v202 = vr3
			v203 = float32(v201.X*v202.X) + float32(v201.Y*v202.Y)
			goto _204
		_204:
			Cdot3 = v203
			impulse3 = float32((*b2DistanceJoint)(unsafe.Pointer(joint)).AxialMass * ((*b2DistanceJoint)(unsafe.Pointer(joint)).MotorSpeed - Cdot3))
			oldImpulse = (*b2DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse
			maxImpulse = float32((*b2StepContext)(unsafe.Pointer(context)).H * (*b2DistanceJoint)(unsafe.Pointer(joint)).MaxMotorForce)
			v205 = (*b2DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse + impulse3
			v206 = -maxImpulse
			v207 = maxImpulse
			if v205 < v206 {
				v210 = v206
			} else {
				if v205 > v207 {
					v211 = v207
				} else {
					v211 = v205
				}
				v210 = v211
			}
			v208 = v210
			goto _209
		_209:
			(*b2DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse = v208
			impulse3 = (*b2DistanceJoint)(unsafe.Pointer(joint)).MotorImpulse - oldImpulse
			v212 = impulse3
			v213 = axis
			v214 = Vec2{
				X: float32(v212 * v213.X),
				Y: float32(v212 * v213.Y),
			}
			goto _215
		_215:
			P3 = v214
			v216 = vA
			v217 = mA
			v218 = P3
			v219 = Vec2{
				X: v216.X - float32(v217*v218.X),
				Y: v216.Y - float32(v217*v218.Y),
			}
			goto _220
		_220:
			vA = v219
			v221 = rA
			v222 = P3
			v223 = float32(v221.X*v222.Y) - float32(v221.Y*v222.X)
			goto _224
		_224:
			wA = wA - float32(iA*v223)
			v225 = vB
			v226 = mB
			v227 = P3
			v228 = Vec2{
				X: v225.X + float32(v226*v227.X),
				Y: v225.Y + float32(v226*v227.Y),
			}
			goto _229
		_229:
			vB = v228
			v230 = rB
			v231 = P3
			v232 = float32(v230.X*v231.Y) - float32(v230.Y*v231.X)
			goto _233
		_233:
			wB = wB + float32(iB*v232)
		}
	} else {
		v234 = vB
		v235 = vA
		v236 = Vec2{
			X: v234.X - v235.X,
			Y: v234.Y - v235.Y,
		}
		goto _237
	_237:
		v238 = wB
		v239 = rB
		v240 = Vec2{
			X: float32(-v238 * v239.Y),
			Y: float32(v238 * v239.X),
		}
		goto _241
	_241:
		v242 = wA
		v243 = rA
		v244 = Vec2{
			X: float32(-v242 * v243.Y),
			Y: float32(v242 * v243.X),
		}
		goto _245
	_245:
		v246 = v240
		v247 = v244
		v248 = Vec2{
			X: v246.X - v247.X,
			Y: v246.Y - v247.Y,
		}
		goto _249
	_249:
		v250 = v236
		v251 = v248
		v252 = Vec2{
			X: v250.X + v251.X,
			Y: v250.Y + v251.Y,
		}
		goto _253
	_253:
		// rigid constraint
		vr4 = v252
		v254 = axis
		v255 = vr4
		v256 = float32(v254.X*v255.X) + float32(v254.Y*v255.Y)
		goto _257
	_257:
		Cdot4 = v256
		C3 = length1 - (*b2DistanceJoint)(unsafe.Pointer(joint)).Length
		bias3 = float32FromFloat32(0)
		massScale1 = float32FromFloat32(1)
		impulseScale1 = float32FromFloat32(0)
		if useBias != 0 {
			bias3 = float32((*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * C3)
			massScale1 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
			impulseScale1 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
		}
		impulse4 = float32(float32(-massScale1*(*b2DistanceJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot4+bias3)) - float32(impulseScale1*(*b2DistanceJoint)(unsafe.Pointer(joint)).Impulse)
		*(*float32)(unsafe.Pointer(joint + 28)) += impulse4
		v258 = impulse4
		v259 = axis
		v260 = Vec2{
			X: float32(v258 * v259.X),
			Y: float32(v258 * v259.Y),
		}
		goto _261
	_261:
		P4 = v260
		v262 = vA
		v263 = mA
		v264 = P4
		v265 = Vec2{
			X: v262.X - float32(v263*v264.X),
			Y: v262.Y - float32(v263*v264.Y),
		}
		goto _266
	_266:
		vA = v265
		v267 = rA
		v268 = P4
		v269 = float32(v267.X*v268.Y) - float32(v267.Y*v268.X)
		goto _270
	_270:
		wA = wA - float32(iA*v269)
		v271 = vB
		v272 = mB
		v273 = P4
		v274 = Vec2{
			X: v271.X + float32(v272*v273.X),
			Y: v271.Y + float32(v272*v273.Y),
		}
		goto _275
	_275:
		vB = v274
		v276 = rB
		v277 = P4
		v278 = float32(v276.X*v277.Y) - float32(v276.Y*v277.X)
		goto _279
	_279:
		wB = wB + float32(iB*v278)
	}
	(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
	(*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
	(*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
}

func b2DrawDistanceJoint(tls *_Stack, draw uintptr, base uintptr, transformA Transform, transformB Transform) {
	var axis, n, offset, pA, pB, pMax, pMin, pRest, v10, v11, v13, v14, v16, v18, v19, v2, v21, v23, v24, v26, v27, v3, v30, v31, v33, v34, v35, v37, v38, v39, v41, v42, v43, v45, v46, v47, v49, v51, v52, v6, v7, v9 Vec2
	var invLength, length, x, y, v17, v22, v29, v50 float32
	var joint uintptr
	var v1, v5 Transform
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axis, invLength, joint, length, n, offset, pA, pB, pMax, pMin, pRest, x, y, v1, v10, v11, v13, v14, v16, v17, v18, v19, v2, v21, v22, v23, v24, v26, v27, v29, v3, v30, v31, v33, v34, v35, v37, v38, v39, v41, v42, v43, v45, v46, v47, v49, v5, v50, v51, v52, v6, v7, v9
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_distanceJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+4612, __ccgo_ts+4580, int32FromInt32(514)) != 0 {
		__builtin_trap(tls)
	}
	joint = base + 68
	v1 = transformA
	v2 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorA
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	pA = v3
	v5 = transformB
	v6 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorB
	x = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y) + v5.P.X
	y = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y) + v5.P.Y
	v7 = Vec2{
		X: x,
		Y: y,
	}
	goto _8
_8:
	pB = v7
	v9 = pB
	v10 = pA
	v11 = Vec2{
		X: v9.X - v10.X,
		Y: v9.Y - v10.Y,
	}
	goto _12
_12:
	v13 = v11
	length = sqrtf(tls, float32(v13.X*v13.X)+float32(v13.Y*v13.Y))
	if length < float32FromFloat32(1.1920928955078125e-07) {
		v14 = Vec2{}
		goto _15
	}
	invLength = float32FromFloat32(1) / length
	n = Vec2{
		X: float32(invLength * v13.X),
		Y: float32(invLength * v13.Y),
	}
	v14 = n
	goto _15
_15:
	axis = v14
	if (*b2DistanceJoint)(unsafe.Pointer(joint)).MinLength < (*b2DistanceJoint)(unsafe.Pointer(joint)).MaxLength && (*b2DistanceJoint)(unsafe.Pointer(joint)).EnableLimit != 0 {
		v16 = pA
		v17 = (*b2DistanceJoint)(unsafe.Pointer(joint)).MinLength
		v18 = axis
		v19 = Vec2{
			X: v16.X + float32(v17*v18.X),
			Y: v16.Y + float32(v17*v18.Y),
		}
		goto _20
	_20:
		pMin = v19
		v21 = pA
		v22 = (*b2DistanceJoint)(unsafe.Pointer(joint)).MaxLength
		v23 = axis
		v24 = Vec2{
			X: v21.X + float32(v22*v23.X),
			Y: v21.Y + float32(v22*v23.Y),
		}
		goto _25
	_25:
		pMax = v24
		v26 = axis
		v27 = Vec2{
			X: v26.Y,
			Y: -v26.X,
		}
		goto _28
	_28:
		v29 = float32(float32FromFloat32(0.05) * b2_lengthUnitsPerMeter)
		v30 = v27
		v31 = Vec2{
			X: float32(v29 * v30.X),
			Y: float32(v29 * v30.Y),
		}
		goto _32
	_32:
		offset = v31
		if (*b2DistanceJoint)(unsafe.Pointer(joint)).MinLength > float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter) {
			// draw->DrawPoint(pMin, 4.0f, c2, draw->context);
			v33 = pMin
			v34 = offset
			v35 = Vec2{
				X: v33.X - v34.X,
				Y: v33.Y - v34.Y,
			}
			goto _36
		_36:
			v37 = pMin
			v38 = offset
			v39 = Vec2{
				X: v37.X + v38.X,
				Y: v37.Y + v38.Y,
			}
			goto _40
		_40:
			(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, v35, v39, int32(b2_colorLightGreen), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		}
		if (*b2DistanceJoint)(unsafe.Pointer(joint)).MaxLength < float32(float32FromFloat32(100000)*b2_lengthUnitsPerMeter) {
			// draw->DrawPoint(pMax, 4.0f, c3, draw->context);
			v41 = pMax
			v42 = offset
			v43 = Vec2{
				X: v41.X - v42.X,
				Y: v41.Y - v42.Y,
			}
			goto _44
		_44:
			v45 = pMax
			v46 = offset
			v47 = Vec2{
				X: v45.X + v46.X,
				Y: v45.Y + v46.Y,
			}
			goto _48
		_48:
			(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, v43, v47, int32(b2_colorRed), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		}
		if (*b2DistanceJoint)(unsafe.Pointer(joint)).MinLength > float32(float32FromFloat32(0.005)*b2_lengthUnitsPerMeter) && (*b2DistanceJoint)(unsafe.Pointer(joint)).MaxLength < float32(float32FromFloat32(100000)*b2_lengthUnitsPerMeter) {
			(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, pMin, pMax, int32(b2_colorGray), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		}
	}
	(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, pA, pB, int32(b2_colorWhite), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pA, float32FromFloat32(4), int32(b2_colorWhite), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pB, float32FromFloat32(4), int32(b2_colorWhite), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	if (*b2DistanceJoint)(unsafe.Pointer(joint)).Hertz > float32FromFloat32(0) && (*b2DistanceJoint)(unsafe.Pointer(joint)).EnableSpring != 0 {
		v49 = pA
		v50 = (*b2DistanceJoint)(unsafe.Pointer(joint)).Length
		v51 = axis
		v52 = Vec2{
			X: v49.X + float32(v50*v51.X),
			Y: v49.Y + float32(v50*v51.Y),
		}
		goto _53
	_53:
		pRest = v52
		(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pRest, float32FromFloat32(4), int32(b2_colorBlue), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	}
}

func b2AddJointToIsland(tls *_Stack, world uintptr, islandId int32, joint uintptr) {
	var headJoint, island, v1, v3, v5, v7 uintptr
	var v2, v6 int32
	_, _, _, _, _, _, _, _ = headJoint, island, v1, v2, v3, v5, v6, v7
	if !((*b2Joint)(unsafe.Pointer(joint)).IslandId == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+2367, __ccgo_ts+7079, int32FromInt32(266)) != 0 {
		__builtin_trap(tls)
	}
	if !((*b2Joint)(unsafe.Pointer(joint)).IslandPrev == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+7781, __ccgo_ts+7079, int32FromInt32(267)) != 0 {
		__builtin_trap(tls)
	}
	if !((*b2Joint)(unsafe.Pointer(joint)).IslandNext == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+7816, __ccgo_ts+7079, int32FromInt32(268)) != 0 {
		__builtin_trap(tls)
	}
	v1 = world + 1184
	v2 = islandId
	if !(0 <= v2 && v2 < (*b2IslandArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2IslandArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*56
	goto _4
_4:
	island = v3
	if (*b2Island)(unsafe.Pointer(island)).HeadJoint != -int32(1) {
		(*b2Joint)(unsafe.Pointer(joint)).IslandNext = (*b2Island)(unsafe.Pointer(island)).HeadJoint
		v5 = world + 1104
		v6 = (*b2Island)(unsafe.Pointer(island)).HeadJoint
		if !(0 <= v6 && v6 < (*b2JointArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v7 = (*b2JointArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*72
		goto _8
	_8:
		headJoint = v7
		(*b2Joint)(unsafe.Pointer(headJoint)).IslandPrev = (*b2Joint)(unsafe.Pointer(joint)).JointId
	}
	(*b2Island)(unsafe.Pointer(island)).HeadJoint = (*b2Joint)(unsafe.Pointer(joint)).JointId
	if (*b2Island)(unsafe.Pointer(island)).TailJoint == -int32(1) {
		(*b2Island)(unsafe.Pointer(island)).TailJoint = (*b2Island)(unsafe.Pointer(island)).HeadJoint
	}
	*(*int32)(unsafe.Pointer(island + 44)) += int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).IslandId = islandId
	b2ValidateIsland(tls, world, islandId)
}

func b2LinkJoint(tls *_Stack, world uintptr, joint uintptr, mergeIslands uint8) {
	var bodyA, bodyB, islandA, islandB, parent, parent1, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var islandIdA, islandIdB, v10, v14, v18, v2, v22, v6 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, islandA, islandB, islandIdA, islandIdB, parent, parent1, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v23, v3, v5, v6, v7, v9
	v1 = world + 1024
	v2 = (*(*b2JointEdge)(unsafe.Pointer(joint + 20))).BodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	bodyA = v3
	v5 = world + 1024
	v6 = (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + 1*12))).BodyId
	if !(0 <= v6 && v6 < (*b2BodyArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodyArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*128
	goto _8
_8:
	bodyB = v7
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) && (*b2Body)(unsafe.Pointer(bodyB)).SetIndex >= int32(b2_firstSleepingSet) {
		b2WakeSolverSet(tls, world, (*b2Body)(unsafe.Pointer(bodyB)).SetIndex)
	} else {
		if (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) && (*b2Body)(unsafe.Pointer(bodyA)).SetIndex >= int32(b2_firstSleepingSet) {
			b2WakeSolverSet(tls, world, (*b2Body)(unsafe.Pointer(bodyA)).SetIndex)
		}
	}
	islandIdA = (*b2Body)(unsafe.Pointer(bodyA)).IslandId
	islandIdB = (*b2Body)(unsafe.Pointer(bodyB)).IslandId
	if !(islandIdA != -int32(1) || islandIdB != -int32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+7479, __ccgo_ts+7079, int32FromInt32(308)) != 0 {
		__builtin_trap(tls)
	}
	if islandIdA == islandIdB {
		// Joint in same island
		b2AddJointToIsland(tls, world, islandIdA, joint)
		return
	}
	// Union-find root of islandA
	islandA = uintptrFromInt32(0)
	if islandIdA != -int32(1) {
		v9 = world + 1184
		v10 = islandIdA
		if !(0 <= v10 && v10 < (*b2IslandArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
			__builtin_trap(tls)
		}
		v11 = (*b2IslandArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*56
		goto _12
	_12:
		islandA = v11
		for (*b2Island)(unsafe.Pointer(islandA)).ParentIsland != -int32(1) {
			v13 = world + 1184
			v14 = (*b2Island)(unsafe.Pointer(islandA)).ParentIsland
			if !(0 <= v14 && v14 < (*b2IslandArray)(unsafe.Pointer(v13)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
				__builtin_trap(tls)
			}
			v15 = (*b2IslandArray)(unsafe.Pointer(v13)).Data + uintptr(v14)*56
			goto _16
		_16:
			parent = v15
			if (*b2Island)(unsafe.Pointer(parent)).ParentIsland != -int32(1) {
				// path compression
				(*b2Island)(unsafe.Pointer(islandA)).ParentIsland = (*b2Island)(unsafe.Pointer(parent)).ParentIsland
			}
			islandIdA = (*b2Island)(unsafe.Pointer(islandA)).ParentIsland
			islandA = parent
		}
	}
	// Union-find root of islandB
	islandB = uintptrFromInt32(0)
	if islandIdB != -int32(1) {
		v17 = world + 1184
		v18 = islandIdB
		if !(0 <= v18 && v18 < (*b2IslandArray)(unsafe.Pointer(v17)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
			__builtin_trap(tls)
		}
		v19 = (*b2IslandArray)(unsafe.Pointer(v17)).Data + uintptr(v18)*56
		goto _20
	_20:
		islandB = v19
		for (*b2Island)(unsafe.Pointer(islandB)).ParentIsland != -int32(1) {
			v21 = world + 1184
			v22 = (*b2Island)(unsafe.Pointer(islandB)).ParentIsland
			if !(0 <= v22 && v22 < (*b2IslandArray)(unsafe.Pointer(v21)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
				__builtin_trap(tls)
			}
			v23 = (*b2IslandArray)(unsafe.Pointer(v21)).Data + uintptr(v22)*56
			goto _24
		_24:
			parent1 = v23
			if (*b2Island)(unsafe.Pointer(parent1)).ParentIsland != -int32(1) {
				// path compression
				(*b2Island)(unsafe.Pointer(islandB)).ParentIsland = (*b2Island)(unsafe.Pointer(parent1)).ParentIsland
			}
			islandIdB = (*b2Island)(unsafe.Pointer(islandB)).ParentIsland
			islandB = parent1
		}
	}
	if !(islandA != uintptrFromInt32(0) || islandB != uintptrFromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+7536, __ccgo_ts+7079, int32FromInt32(355)) != 0 {
		__builtin_trap(tls)
	}
	// Union-Find link island roots
	if islandA != islandB && islandA != uintptrFromInt32(0) && islandB != uintptrFromInt32(0) {
		if !(islandA != islandB) && b2InternalAssertFcn(tls, __ccgo_ts+7571, __ccgo_ts+7079, int32FromInt32(360)) != 0 {
			__builtin_trap(tls)
		}
		if !((*b2Island)(unsafe.Pointer(islandB)).ParentIsland == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+7590, __ccgo_ts+7079, int32FromInt32(361)) != 0 {
			__builtin_trap(tls)
		}
		(*b2Island)(unsafe.Pointer(islandB)).ParentIsland = islandIdA
	}
	if islandA != uintptrFromInt32(0) {
		b2AddJointToIsland(tls, world, islandIdA, joint)
	} else {
		b2AddJointToIsland(tls, world, islandIdB, joint)
	}
	// Joints need to have islands merged immediately when they are created
	// to keep the island graph valid.
	// However, when a body type is being changed the merge can be deferred until
	// all joints are linked.
	if mergeIslands != 0 {
		b2MergeAwakeIslands(tls, world)
	}
}

func b2UnlinkJoint(tls *_Stack, world uintptr, joint uintptr) {
	var island, nextJoint, prevJoint, v1, v11, v3, v5, v7, v9 uintptr
	var islandId, v10, v2, v6 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _ = island, islandId, nextJoint, prevJoint, v1, v10, v11, v2, v3, v5, v6, v7, v9
	if !((*b2Joint)(unsafe.Pointer(joint)).IslandId != -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+7851, __ccgo_ts+7079, int32FromInt32(386)) != 0 {
		__builtin_trap(tls)
	}
	// remove from island
	islandId = (*b2Joint)(unsafe.Pointer(joint)).IslandId
	v1 = world + 1184
	v2 = islandId
	if !(0 <= v2 && v2 < (*b2IslandArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+557, int32FromInt32(88)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2IslandArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*56
	goto _4
_4:
	island = v3
	if (*b2Joint)(unsafe.Pointer(joint)).IslandPrev != -int32(1) {
		v5 = world + 1104
		v6 = (*b2Joint)(unsafe.Pointer(joint)).IslandPrev
		if !(0 <= v6 && v6 < (*b2JointArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v7 = (*b2JointArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*72
		goto _8
	_8:
		prevJoint = v7
		if !((*b2Joint)(unsafe.Pointer(prevJoint)).IslandNext == (*b2Joint)(unsafe.Pointer(joint)).JointId) && b2InternalAssertFcn(tls, __ccgo_ts+7884, __ccgo_ts+7079, int32FromInt32(395)) != 0 {
			__builtin_trap(tls)
		}
		(*b2Joint)(unsafe.Pointer(prevJoint)).IslandNext = (*b2Joint)(unsafe.Pointer(joint)).IslandNext
	}
	if (*b2Joint)(unsafe.Pointer(joint)).IslandNext != -int32(1) {
		v9 = world + 1104
		v10 = (*b2Joint)(unsafe.Pointer(joint)).IslandNext
		if !(0 <= v10 && v10 < (*b2JointArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v11 = (*b2JointArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*72
		goto _12
	_12:
		nextJoint = v11
		if !((*b2Joint)(unsafe.Pointer(nextJoint)).IslandPrev == (*b2Joint)(unsafe.Pointer(joint)).JointId) && b2InternalAssertFcn(tls, __ccgo_ts+7924, __ccgo_ts+7079, int32FromInt32(402)) != 0 {
			__builtin_trap(tls)
		}
		(*b2Joint)(unsafe.Pointer(nextJoint)).IslandPrev = (*b2Joint)(unsafe.Pointer(joint)).IslandPrev
	}
	if (*b2Island)(unsafe.Pointer(island)).HeadJoint == (*b2Joint)(unsafe.Pointer(joint)).JointId {
		(*b2Island)(unsafe.Pointer(island)).HeadJoint = (*b2Joint)(unsafe.Pointer(joint)).IslandNext
	}
	if (*b2Island)(unsafe.Pointer(island)).TailJoint == (*b2Joint)(unsafe.Pointer(joint)).JointId {
		(*b2Island)(unsafe.Pointer(island)).TailJoint = (*b2Joint)(unsafe.Pointer(joint)).IslandPrev
	}
	if !((*b2Island)(unsafe.Pointer(island)).JointCount > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+7964, __ccgo_ts+7079, int32FromInt32(416)) != 0 {
		__builtin_trap(tls)
	}
	*(*int32)(unsafe.Pointer(island + 44)) -= int32(1)
	*(*int32)(unsafe.Pointer(island + 52)) += int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).IslandId = -int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).IslandPrev = -int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).IslandNext = -int32(1)
	b2ValidateIsland(tls, world, islandId)
}

func b2JointArray_Create(tls *_Stack, capacity int32) (r b2JointArray) {
	var a b2JointArray
	_ = a
	a = b2JointArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(72)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2JointArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2JointArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2JointArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2JointArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2JointArray)(unsafe.Pointer(a)).Capacity)*uint64(72)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(72)))
	(*b2JointArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2JointArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2JointArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2JointArray)(unsafe.Pointer(a)).Capacity)*uint64(72)))
	(*b2JointArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2JointArray)(unsafe.Pointer(a)).Count = 0
	(*b2JointArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2JointSimArray_Create(tls *_Stack, capacity int32) (r b2JointSimArray) {
	var a b2JointSimArray
	_ = a
	a = b2JointSimArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(196)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2JointSimArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2JointSimArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2JointSimArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2JointSimArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2JointSimArray)(unsafe.Pointer(a)).Capacity)*uint64(196)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(196)))
	(*b2JointSimArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2JointSimArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2JointSimArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2JointSimArray)(unsafe.Pointer(a)).Capacity)*uint64(196)))
	(*b2JointSimArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2JointSimArray)(unsafe.Pointer(a)).Count = 0
	(*b2JointSimArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2DefaultDistanceJointDef(tls *_Stack) (r DistanceJointDef) {
	var def DistanceJointDef
	_ = def
	def = DistanceJointDef{}
	def.Length = float32FromFloat32(1)
	def.MaxLength = float32(float32FromFloat32(100000) * b2_lengthUnitsPerMeter)
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}

func b2DefaultMotorJointDef(tls *_Stack) (r MotorJointDef) {
	var def MotorJointDef
	_ = def
	def = MotorJointDef{}
	def.MaxForce = float32FromFloat32(1)
	def.MaxTorque = float32FromFloat32(1)
	def.CorrectionFactor = float32FromFloat32(0.3)
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}

func b2DefaultMouseJointDef(tls *_Stack) (r MouseJointDef) {
	var def MouseJointDef
	_ = def
	def = MouseJointDef{}
	def.Hertz = float32FromFloat32(4)
	def.DampingRatio = float32FromFloat32(1)
	def.MaxForce = float32FromFloat32(1)
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}

func b2DefaultFilterJointDef(tls *_Stack) (r FilterJointDef) {
	var def FilterJointDef
	_ = def
	def = FilterJointDef{}
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}

func b2DefaultPrismaticJointDef(tls *_Stack) (r PrismaticJointDef) {
	var def PrismaticJointDef
	_ = def
	def = PrismaticJointDef{}
	def.LocalAxisA = Vec2{
		X: float32FromFloat32(1),
	}
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}

func b2DefaultRevoluteJointDef(tls *_Stack) (r RevoluteJointDef) {
	var def RevoluteJointDef
	_ = def
	def = RevoluteJointDef{}
	def.DrawSize = float32FromFloat32(0.25)
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}

func b2DefaultWeldJointDef(tls *_Stack) (r WeldJointDef) {
	var def WeldJointDef
	_ = def
	def = WeldJointDef{}
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}

func b2DefaultWheelJointDef(tls *_Stack) (r WheelJointDef) {
	var def WheelJointDef
	_ = def
	def = WheelJointDef{}
	def.LocalAxisA.Y = float32FromFloat32(1)
	def.EnableSpring = boolUint8(true1 != 0)
	def.Hertz = float32FromFloat32(1)
	def.DampingRatio = float32FromFloat32(0.7)
	def.InternalValue = int32(_B2_SECRET_COOKIE)
	return def
}

func b2GetJointFullId(tls *_Stack, world uintptr, jointId JointId) (r uintptr) {
	var id, v2 int32
	var joint, v1, v3 uintptr
	_, _, _, _, _ = id, joint, v1, v2, v3
	id = jointId.Index1 - int32(1)
	v1 = world + 1104
	v2 = id
	if !(0 <= v2 && v2 < (*b2JointArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2JointArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*72
	goto _4
_4:
	joint = v3
	if !((*b2Joint)(unsafe.Pointer(joint)).JointId == id && int32FromUint16((*b2Joint)(unsafe.Pointer(joint)).Generation) == int32FromUint16(jointId.Generation)) && b2InternalAssertFcn(tls, __ccgo_ts+8997, __ccgo_ts+9061, int32FromInt32(105)) != 0 {
		__builtin_trap(tls)
	}
	return joint
}

func b2GetJointSim(tls *_Stack, world uintptr, joint uintptr) (r uintptr) {
	var color, set, v1, v11, v3, v5, v7, v9 uintptr
	var v10, v2, v6 int32
	_, _, _, _, _, _, _, _, _, _, _ = color, set, v1, v10, v11, v2, v3, v5, v6, v7, v9
	if (*b2Joint)(unsafe.Pointer(joint)).SetIndex == int32(b2_awakeSet) {
		if !(0 <= (*b2Joint)(unsafe.Pointer(joint)).ColorIndex && (*b2Joint)(unsafe.Pointer(joint)).ColorIndex < int32(_B2_GRAPH_COLOR_COUNT)) && b2InternalAssertFcn(tls, __ccgo_ts+1796, __ccgo_ts+9061, int32FromInt32(113)) != 0 {
			__builtin_trap(tls)
		}
		color = world + 328 + uintptr((*b2Joint)(unsafe.Pointer(joint)).ColorIndex)*56
		v1 = color + 32
		v2 = (*b2Joint)(unsafe.Pointer(joint)).LocalIndex
		if !(0 <= v2 && v2 < (*b2JointSimArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(342)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2JointSimArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*196
		goto _4
	_4:
		return v3
	}
	v5 = world + 1064
	v6 = (*b2Joint)(unsafe.Pointer(joint)).SetIndex
	if !(0 <= v6 && v6 < (*b2SolverSetArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2SolverSetArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*88
	goto _8
_8:
	set = v7
	v9 = set + 32
	v10 = (*b2Joint)(unsafe.Pointer(joint)).LocalIndex
	if !(0 <= v10 && v10 < (*b2JointSimArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(342)) != 0 {
		__builtin_trap(tls)
	}
	v11 = (*b2JointSimArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*196
	goto _12
_12:
	return v11
}

func b2GetJointSimCheckType(tls *_Stack, jointId JointId, type1 JointType) (r uintptr) {
	var joint, jointSim, world uintptr
	_, _, _ = joint, jointSim, world
	_ = uint64FromInt64(4)
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if !(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)) && b2InternalAssertFcn(tls, __ccgo_ts+1152, __ccgo_ts+9061, int32FromInt32(127)) != 0 {
		__builtin_trap(tls)
	}
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return uintptrFromInt32(0)
	}
	joint = b2GetJointFullId(tls, world, jointId)
	if !((*b2Joint)(unsafe.Pointer(joint)).Type1 == type1) && b2InternalAssertFcn(tls, __ccgo_ts+9084, __ccgo_ts+9061, int32FromInt32(134)) != 0 {
		__builtin_trap(tls)
	}
	jointSim = b2GetJointSim(tls, world, joint)
	if !((*b2JointSim)(unsafe.Pointer(jointSim)).Type1 == type1) && b2InternalAssertFcn(tls, __ccgo_ts+9104, __ccgo_ts+9061, int32FromInt32(136)) != 0 {
		__builtin_trap(tls)
	}
	return jointSim
}

func b2CreateJoint(tls *_Stack, world uintptr, bodyA uintptr, bodyB uintptr, userData uintptr, drawSize float32, type1 JointType, collideConnected uint8) (r b2JointPair) {
	var bodyIdA, bodyIdB, jointId, keyA, keyB, maxSetIndex, newCapacity, newCapacity1, setIndex, v1, v14, v18, v2, v22, v26, v3, v30, v34, v38, v42, v46, v5, v50, v7, v9 int32
	var edgeA, edgeB, joint, jointA, jointB, jointSim, mergedSet, set, set1, set2, v10, v13, v15, v17, v19, v21, v23, v25, v27, v29, v31, v33, v35, v37, v39, v41, v43, v45, v47, v49, v51, v6, v8, p12 uintptr
	var mergeIslands uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyIdA, bodyIdB, edgeA, edgeB, joint, jointA, jointB, jointId, jointSim, keyA, keyB, maxSetIndex, mergeIslands, mergedSet, newCapacity, newCapacity1, set, set1, set2, setIndex, v1, v10, v13, v14, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v29, v3, v30, v31, v33, v34, v35, v37, v38, v39, v41, v42, v43, v45, v46, v47, v49, v5, v50, v51, v6, v7, v8, v9, p12
	bodyIdA = (*b2Body)(unsafe.Pointer(bodyA)).Id
	bodyIdB = (*b2Body)(unsafe.Pointer(bodyB)).Id
	v1 = (*b2Body)(unsafe.Pointer(bodyA)).SetIndex
	v2 = (*b2Body)(unsafe.Pointer(bodyB)).SetIndex
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	maxSetIndex = v3
	// Create joint id and joint
	jointId = b2AllocId(tls, world+1080)
	if jointId == (*b2World)(unsafe.Pointer(world)).Joints.Count {
		v6 = world + 1104
		if (*b2JointArray)(unsafe.Pointer(v6)).Count == (*b2JointArray)(unsafe.Pointer(v6)).Capacity {
			if (*b2JointArray)(unsafe.Pointer(v6)).Capacity < int32(2) {
				v7 = int32(2)
			} else {
				v7 = (*b2JointArray)(unsafe.Pointer(v6)).Capacity + (*b2JointArray)(unsafe.Pointer(v6)).Capacity>>int32(1)
			}
			newCapacity = v7
			b2JointArray_Reserve(tls, v6, newCapacity)
		}
		*(*b2Joint)(unsafe.Pointer((*b2JointArray)(unsafe.Pointer(v6)).Data + uintptr((*b2JointArray)(unsafe.Pointer(v6)).Count)*72)) = b2Joint{}
		*(*int32)(unsafe.Pointer(v6 + 8)) += int32(1)
	}
	v8 = world + 1104
	v9 = jointId
	if !(0 <= v9 && v9 < (*b2JointArray)(unsafe.Pointer(v8)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
		__builtin_trap(tls)
	}
	v10 = (*b2JointArray)(unsafe.Pointer(v8)).Data + uintptr(v9)*72
	goto _11
_11:
	joint = v10
	(*b2Joint)(unsafe.Pointer(joint)).JointId = jointId
	(*b2Joint)(unsafe.Pointer(joint)).UserData = userData
	p12 = joint + 68
	*(*uint16_t)(unsafe.Pointer(p12)) = uint16_t(int32(*(*uint16_t)(unsafe.Pointer(p12))) + int32FromInt32(1))
	(*b2Joint)(unsafe.Pointer(joint)).SetIndex = -int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).ColorIndex = -int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).LocalIndex = -int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).IslandId = -int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).IslandPrev = -int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).IslandNext = -int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).DrawSize = drawSize
	(*b2Joint)(unsafe.Pointer(joint)).Type1 = type1
	(*b2Joint)(unsafe.Pointer(joint)).CollideConnected = collideConnected
	(*b2Joint)(unsafe.Pointer(joint)).IsMarked = boolUint8(false1 != 0)
	// Doubly linked list on bodyA
	(*(*b2JointEdge)(unsafe.Pointer(joint + 20))).BodyId = bodyIdA
	(*(*b2JointEdge)(unsafe.Pointer(joint + 20))).PrevKey = -int32(1)
	(*(*b2JointEdge)(unsafe.Pointer(joint + 20))).NextKey = (*b2Body)(unsafe.Pointer(bodyA)).HeadJointKey
	keyA = jointId<<int32(1) | 0
	if (*b2Body)(unsafe.Pointer(bodyA)).HeadJointKey != -int32(1) {
		v13 = world + 1104
		v14 = (*b2Body)(unsafe.Pointer(bodyA)).HeadJointKey >> int32(1)
		if !(0 <= v14 && v14 < (*b2JointArray)(unsafe.Pointer(v13)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v15 = (*b2JointArray)(unsafe.Pointer(v13)).Data + uintptr(v14)*72
		goto _16
	_16:
		jointA = v15
		edgeA = jointA + 20 + uintptr((*b2Body)(unsafe.Pointer(bodyA)).HeadJointKey&int32FromInt32(1))*12
		(*b2JointEdge)(unsafe.Pointer(edgeA)).PrevKey = keyA
	}
	(*b2Body)(unsafe.Pointer(bodyA)).HeadJointKey = keyA
	*(*int32)(unsafe.Pointer(bodyA + 72)) += int32(1)
	// Doubly linked list on bodyB
	(*(*b2JointEdge)(unsafe.Pointer(joint + 20 + 1*12))).BodyId = bodyIdB
	(*(*b2JointEdge)(unsafe.Pointer(joint + 20 + 1*12))).PrevKey = -int32(1)
	(*(*b2JointEdge)(unsafe.Pointer(joint + 20 + 1*12))).NextKey = (*b2Body)(unsafe.Pointer(bodyB)).HeadJointKey
	keyB = jointId<<int32(1) | int32(1)
	if (*b2Body)(unsafe.Pointer(bodyB)).HeadJointKey != -int32(1) {
		v17 = world + 1104
		v18 = (*b2Body)(unsafe.Pointer(bodyB)).HeadJointKey >> int32(1)
		if !(0 <= v18 && v18 < (*b2JointArray)(unsafe.Pointer(v17)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v19 = (*b2JointArray)(unsafe.Pointer(v17)).Data + uintptr(v18)*72
		goto _20
	_20:
		jointB = v19
		edgeB = jointB + 20 + uintptr((*b2Body)(unsafe.Pointer(bodyB)).HeadJointKey&int32FromInt32(1))*12
		(*b2JointEdge)(unsafe.Pointer(edgeB)).PrevKey = keyB
	}
	(*b2Body)(unsafe.Pointer(bodyB)).HeadJointKey = keyB
	*(*int32)(unsafe.Pointer(bodyB + 72)) += int32(1)
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_disabledSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_disabledSet) {
		v21 = world + 1064
		v22 = int32(b2_disabledSet)
		if !(0 <= v22 && v22 < (*b2SolverSetArray)(unsafe.Pointer(v21)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
			__builtin_trap(tls)
		}
		v23 = (*b2SolverSetArray)(unsafe.Pointer(v21)).Data + uintptr(v22)*88
		goto _24
	_24:
		// if either body is disabled, create in disabled set
		set = v23
		(*b2Joint)(unsafe.Pointer(joint)).SetIndex = int32(b2_disabledSet)
		(*b2Joint)(unsafe.Pointer(joint)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(set)).JointSims.Count
		v25 = set + 32
		if (*b2JointSimArray)(unsafe.Pointer(v25)).Count == (*b2JointSimArray)(unsafe.Pointer(v25)).Capacity {
			if (*b2JointSimArray)(unsafe.Pointer(v25)).Capacity < int32(2) {
				v26 = int32(2)
			} else {
				v26 = (*b2JointSimArray)(unsafe.Pointer(v25)).Capacity + (*b2JointSimArray)(unsafe.Pointer(v25)).Capacity>>int32(1)
			}
			newCapacity1 = v26
			b2JointSimArray_Reserve(tls, v25, newCapacity1)
		}
		*(*int32)(unsafe.Pointer(v25 + 8)) += int32(1)
		v27 = (*b2JointSimArray)(unsafe.Pointer(v25)).Data + uintptr((*b2JointSimArray)(unsafe.Pointer(v25)).Count-int32FromInt32(1))*196
		goto _28
	_28:
		jointSim = v27
		memset(tls, jointSim, 0, uint64(196))
		(*b2JointSim)(unsafe.Pointer(jointSim)).JointId = jointId
		(*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdA = bodyIdA
		(*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdB = bodyIdB
	} else {
		if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_staticSet) && (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_staticSet) {
			v29 = world + 1064
			v30 = int32(b2_staticSet)
			if !(0 <= v30 && v30 < (*b2SolverSetArray)(unsafe.Pointer(v29)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
				__builtin_trap(tls)
			}
			v31 = (*b2SolverSetArray)(unsafe.Pointer(v29)).Data + uintptr(v30)*88
			goto _32
		_32:
			// joint is connecting static bodies
			set1 = v31
			(*b2Joint)(unsafe.Pointer(joint)).SetIndex = int32(b2_staticSet)
			(*b2Joint)(unsafe.Pointer(joint)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(set1)).JointSims.Count
			v33 = set1 + 32
			if (*b2JointSimArray)(unsafe.Pointer(v33)).Count == (*b2JointSimArray)(unsafe.Pointer(v33)).Capacity {
				if (*b2JointSimArray)(unsafe.Pointer(v33)).Capacity < int32(2) {
					v34 = int32(2)
				} else {
					v34 = (*b2JointSimArray)(unsafe.Pointer(v33)).Capacity + (*b2JointSimArray)(unsafe.Pointer(v33)).Capacity>>int32(1)
				}
				newCapacity1 = v34
				b2JointSimArray_Reserve(tls, v33, newCapacity1)
			}
			*(*int32)(unsafe.Pointer(v33 + 8)) += int32(1)
			v35 = (*b2JointSimArray)(unsafe.Pointer(v33)).Data + uintptr((*b2JointSimArray)(unsafe.Pointer(v33)).Count-int32FromInt32(1))*196
			goto _36
		_36:
			jointSim = v35
			memset(tls, jointSim, 0, uint64(196))
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
				if !((*b2Body)(unsafe.Pointer(bodyA)).SetIndex >= int32(b2_firstSleepingSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex >= int32(b2_firstSleepingSet)) && b2InternalAssertFcn(tls, __ccgo_ts+9127, __ccgo_ts+9061, int32FromInt32(253)) != 0 {
					__builtin_trap(tls)
				}
				if !((*b2Body)(unsafe.Pointer(bodyA)).SetIndex != int32(b2_staticSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex != int32(b2_staticSet)) && b2InternalAssertFcn(tls, __ccgo_ts+3493, __ccgo_ts+9061, int32FromInt32(254)) != 0 {
					__builtin_trap(tls)
				}
				// joint should go into the sleeping set (not static set)
				setIndex = maxSetIndex
				v37 = world + 1064
				v38 = setIndex
				if !(0 <= v38 && v38 < (*b2SolverSetArray)(unsafe.Pointer(v37)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
					__builtin_trap(tls)
				}
				v39 = (*b2SolverSetArray)(unsafe.Pointer(v37)).Data + uintptr(v38)*88
				goto _40
			_40:
				set2 = v39
				(*b2Joint)(unsafe.Pointer(joint)).SetIndex = setIndex
				(*b2Joint)(unsafe.Pointer(joint)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(set2)).JointSims.Count
				v41 = set2 + 32
				if (*b2JointSimArray)(unsafe.Pointer(v41)).Count == (*b2JointSimArray)(unsafe.Pointer(v41)).Capacity {
					if (*b2JointSimArray)(unsafe.Pointer(v41)).Capacity < int32(2) {
						v42 = int32(2)
					} else {
						v42 = (*b2JointSimArray)(unsafe.Pointer(v41)).Capacity + (*b2JointSimArray)(unsafe.Pointer(v41)).Capacity>>int32(1)
					}
					newCapacity1 = v42
					b2JointSimArray_Reserve(tls, v41, newCapacity1)
				}
				*(*int32)(unsafe.Pointer(v41 + 8)) += int32(1)
				v43 = (*b2JointSimArray)(unsafe.Pointer(v41)).Data + uintptr((*b2JointSimArray)(unsafe.Pointer(v41)).Count-int32FromInt32(1))*196
				goto _44
			_44:
				jointSim = v43
				memset(tls, jointSim, 0, uint64(196))
				(*b2JointSim)(unsafe.Pointer(jointSim)).JointId = jointId
				(*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdA = bodyIdA
				(*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdB = bodyIdB
				if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex != (*b2Body)(unsafe.Pointer(bodyB)).SetIndex && (*b2Body)(unsafe.Pointer(bodyA)).SetIndex >= int32(b2_firstSleepingSet) && (*b2Body)(unsafe.Pointer(bodyB)).SetIndex >= int32(b2_firstSleepingSet) {
					// merge sleeping sets
					b2MergeSolverSets(tls, world, (*b2Body)(unsafe.Pointer(bodyA)).SetIndex, (*b2Body)(unsafe.Pointer(bodyB)).SetIndex)
					if !((*b2Body)(unsafe.Pointer(bodyA)).SetIndex == (*b2Body)(unsafe.Pointer(bodyB)).SetIndex) && b2InternalAssertFcn(tls, __ccgo_ts+9208, __ccgo_ts+9061, int32FromInt32(275)) != 0 {
						__builtin_trap(tls)
					}
					// fix potentially invalid set index
					setIndex = (*b2Body)(unsafe.Pointer(bodyA)).SetIndex
					v45 = world + 1064
					v46 = setIndex
					if !(0 <= v46 && v46 < (*b2SolverSetArray)(unsafe.Pointer(v45)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
						__builtin_trap(tls)
					}
					v47 = (*b2SolverSetArray)(unsafe.Pointer(v45)).Data + uintptr(v46)*88
					goto _48
				_48:
					mergedSet = v47
					// Careful! The joint sim pointer was orphaned by the set merge.
					v49 = mergedSet + 32
					v50 = (*b2Joint)(unsafe.Pointer(joint)).LocalIndex
					if !(0 <= v50 && v50 < (*b2JointSimArray)(unsafe.Pointer(v49)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(342)) != 0 {
						__builtin_trap(tls)
					}
					v51 = (*b2JointSimArray)(unsafe.Pointer(v49)).Data + uintptr(v50)*196
					goto _52
				_52:
					jointSim = v51
				}
				if !((*b2Joint)(unsafe.Pointer(joint)).SetIndex == setIndex) && b2InternalAssertFcn(tls, __ccgo_ts+9243, __ccgo_ts+9061, int32FromInt32(286)) != 0 {
					__builtin_trap(tls)
				}
			}
		}
	}
	(*b2JointSim)(unsafe.Pointer(jointSim)).ConstraintHertz = float32FromFloat32(60)
	(*b2JointSim)(unsafe.Pointer(jointSim)).ConstraintDampingRatio = float32FromFloat32(2)
	(*b2JointSim)(unsafe.Pointer(jointSim)).ConstraintSoftness = b2Softness{
		MassScale: float32FromFloat32(1),
	}
	if !((*b2JointSim)(unsafe.Pointer(jointSim)).JointId == jointId) && b2InternalAssertFcn(tls, __ccgo_ts+9271, __ccgo_ts+9061, int32FromInt32(297)) != 0 {
		__builtin_trap(tls)
	}
	if !((*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdA == bodyIdA) && b2InternalAssertFcn(tls, __ccgo_ts+9300, __ccgo_ts+9061, int32FromInt32(298)) != 0 {
		__builtin_trap(tls)
	}
	if !((*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdB == bodyIdB) && b2InternalAssertFcn(tls, __ccgo_ts+9329, __ccgo_ts+9061, int32FromInt32(299)) != 0 {
		__builtin_trap(tls)
	}
	if (*b2Joint)(unsafe.Pointer(joint)).SetIndex > int32(b2_disabledSet) {
		// Add edge to island graph
		mergeIslands = boolUint8(true1 != 0)
		b2LinkJoint(tls, world, joint, mergeIslands)
	}
	b2ValidateSolverSets(tls, world)
	return b2JointPair{
		Joint:    joint,
		JointSim: jointSim,
	}
}

func b2CreateDistanceJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	var bodyA, bodyB, joint, world uintptr
	var empty b2DistanceJoint
	var jointId JointId
	var pair b2JointPair
	var v1, v10, v11, v12, v13, v15, v2, v3, v5, v6, v7, v8 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, empty, joint, jointId, pair, world, v1, v10, v11, v12, v13, v15, v2, v3, v5, v6, v7, v8
	if !((*DistanceJointDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)) && b2InternalAssertFcn(tls, __ccgo_ts+730, __ccgo_ts+9061, int32FromInt32(355)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorldFromId(tls, worldId)
	if !(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)) && b2InternalAssertFcn(tls, __ccgo_ts+1152, __ccgo_ts+9061, int32FromInt32(358)) != 0 {
		__builtin_trap(tls)
	}
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return JointId{}
	}
	if !(b2Body_IsValid(tls, (*DistanceJointDef)(unsafe.Pointer(def)).BodyIdA) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+9358, __ccgo_ts+9061, int32FromInt32(365)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2Body_IsValid(tls, (*DistanceJointDef)(unsafe.Pointer(def)).BodyIdB) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+9389, __ccgo_ts+9061, int32FromInt32(366)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsValidFloat(tls, (*DistanceJointDef)(unsafe.Pointer(def)).Length) != 0 && (*DistanceJointDef)(unsafe.Pointer(def)).Length > float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+9420, __ccgo_ts+9061, int32FromInt32(367)) != 0 {
		__builtin_trap(tls)
	}
	bodyA = b2GetBodyFullId(tls, world, (*DistanceJointDef)(unsafe.Pointer(def)).BodyIdA)
	bodyB = b2GetBodyFullId(tls, world, (*DistanceJointDef)(unsafe.Pointer(def)).BodyIdB)
	pair = b2CreateJoint(tls, world, bodyA, bodyB, (*DistanceJointDef)(unsafe.Pointer(def)).UserData, float32FromFloat32(1), int32(b2_distanceJoint), (*DistanceJointDef)(unsafe.Pointer(def)).CollideConnected)
	joint = pair.JointSim
	(*b2JointSim)(unsafe.Pointer(joint)).Type1 = int32(b2_distanceJoint)
	(*b2JointSim)(unsafe.Pointer(joint)).LocalOriginAnchorA = (*DistanceJointDef)(unsafe.Pointer(def)).LocalAnchorA
	(*b2JointSim)(unsafe.Pointer(joint)).LocalOriginAnchorB = (*DistanceJointDef)(unsafe.Pointer(def)).LocalAnchorB
	empty = b2DistanceJoint{}
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint = empty
	v1 = (*DistanceJointDef)(unsafe.Pointer(def)).Length
	v2 = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.Length = v3
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.Hertz = (*DistanceJointDef)(unsafe.Pointer(def)).Hertz
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.DampingRatio = (*DistanceJointDef)(unsafe.Pointer(def)).DampingRatio
	v6 = (*DistanceJointDef)(unsafe.Pointer(def)).MinLength
	v7 = float32(float32FromFloat32(0.005) * b2_lengthUnitsPerMeter)
	if v6 > v7 {
		v10 = v6
	} else {
		v10 = v7
	}
	v8 = v10
	goto _9
_9:
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.MinLength = v8
	v11 = (*DistanceJointDef)(unsafe.Pointer(def)).MinLength
	v12 = (*DistanceJointDef)(unsafe.Pointer(def)).MaxLength
	if v11 > v12 {
		v15 = v11
	} else {
		v15 = v12
	}
	v13 = v15
	goto _14
_14:
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.MaxLength = v13
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.MaxMotorForce = (*DistanceJointDef)(unsafe.Pointer(def)).MaxMotorForce
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.MotorSpeed = (*DistanceJointDef)(unsafe.Pointer(def)).MotorSpeed
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.EnableSpring = (*DistanceJointDef)(unsafe.Pointer(def)).EnableSpring
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.EnableLimit = (*DistanceJointDef)(unsafe.Pointer(def)).EnableLimit
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.EnableMotor = (*DistanceJointDef)(unsafe.Pointer(def)).EnableMotor
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.Impulse = float32FromFloat32(0)
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.LowerImpulse = float32FromFloat32(0)
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.UpperImpulse = float32FromFloat32(0)
	(*b2JointSim)(unsafe.Pointer(joint)).ｆ__ccgo13_68.DistanceJoint.MotorImpulse = float32FromFloat32(0)
	// If the joint prevents collisions, then destroy all contacts between attached bodies
	if int32FromUint8((*DistanceJointDef)(unsafe.Pointer(def)).CollideConnected) == false1 {
		b2DestroyContactsBetweenBodies(tls, world, bodyA, bodyB)
	}
	jointId = JointId{
		Index1:     (*b2JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	return jointId
}

func b2CreateMotorJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	var bodyA, bodyB, joint, world uintptr
	var jointId JointId
	var pair b2JointPair
	var v1, v2, v3, v4, v6, v7 float32
	_, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, joint, jointId, pair, world, v1, v2, v3, v4, v6, v7
	if !((*MotorJointDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)) && b2InternalAssertFcn(tls, __ccgo_ts+730, __ccgo_ts+9061, int32FromInt32(408)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorldFromId(tls, worldId)
	if !(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)) && b2InternalAssertFcn(tls, __ccgo_ts+1152, __ccgo_ts+9061, int32FromInt32(411)) != 0 {
		__builtin_trap(tls)
	}
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return JointId{}
	}
	bodyA = b2GetBodyFullId(tls, world, (*MotorJointDef)(unsafe.Pointer(def)).BodyIdA)
	bodyB = b2GetBodyFullId(tls, world, (*MotorJointDef)(unsafe.Pointer(def)).BodyIdB)
	pair = b2CreateJoint(tls, world, bodyA, bodyB, (*MotorJointDef)(unsafe.Pointer(def)).UserData, float32FromFloat32(1), int32(b2_motorJoint), (*MotorJointDef)(unsafe.Pointer(def)).CollideConnected)
	joint = pair.JointSim
	(*b2JointSim)(unsafe.Pointer(joint)).Type1 = int32(b2_motorJoint)
	(*b2JointSim)(unsafe.Pointer(joint)).LocalOriginAnchorA = Vec2{}
	(*b2JointSim)(unsafe.Pointer(joint)).LocalOriginAnchorB = Vec2{}
	*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 68)) = b2MotorJoint{}
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LinearOffset = (*MotorJointDef)(unsafe.Pointer(def)).LinearOffset
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).AngularOffset = (*MotorJointDef)(unsafe.Pointer(def)).AngularOffset
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MaxForce = (*MotorJointDef)(unsafe.Pointer(def)).MaxForce
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MaxTorque = (*MotorJointDef)(unsafe.Pointer(def)).MaxTorque
	v1 = (*MotorJointDef)(unsafe.Pointer(def)).CorrectionFactor
	v2 = float32FromFloat32(0)
	v3 = float32FromFloat32(1)
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
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).CorrectionFactor = v4
	// If the joint prevents collisions, then destroy all contacts between attached bodies
	if int32FromUint8((*MotorJointDef)(unsafe.Pointer(def)).CollideConnected) == false1 {
		b2DestroyContactsBetweenBodies(tls, world, bodyA, bodyB)
	}
	jointId = JointId{
		Index1:     (*b2JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	return jointId
}

func b2CreateMouseJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	var bodyA, bodyB, joint, world uintptr
	var empty b2MouseJoint
	var jointId JointId
	var pair b2JointPair
	var transformA, transformB, v1, v5 Transform
	var vx, vy float32
	var v2, v3, v6, v7 Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, empty, joint, jointId, pair, transformA, transformB, vx, vy, world, v1, v2, v3, v5, v6, v7
	if !((*MouseJointDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)) && b2InternalAssertFcn(tls, __ccgo_ts+730, __ccgo_ts+9061, int32FromInt32(446)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorldFromId(tls, worldId)
	if !(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)) && b2InternalAssertFcn(tls, __ccgo_ts+1152, __ccgo_ts+9061, int32FromInt32(449)) != 0 {
		__builtin_trap(tls)
	}
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return JointId{}
	}
	bodyA = b2GetBodyFullId(tls, world, (*MouseJointDef)(unsafe.Pointer(def)).BodyIdA)
	bodyB = b2GetBodyFullId(tls, world, (*MouseJointDef)(unsafe.Pointer(def)).BodyIdB)
	transformA = b2GetBodyTransformQuick(tls, world, bodyA)
	transformB = b2GetBodyTransformQuick(tls, world, bodyB)
	pair = b2CreateJoint(tls, world, bodyA, bodyB, (*MouseJointDef)(unsafe.Pointer(def)).UserData, float32FromFloat32(1), int32(b2_mouseJoint), (*MouseJointDef)(unsafe.Pointer(def)).CollideConnected)
	joint = pair.JointSim
	(*b2JointSim)(unsafe.Pointer(joint)).Type1 = int32(b2_mouseJoint)
	v1 = transformA
	v2 = (*MouseJointDef)(unsafe.Pointer(def)).Target
	vx = v2.X - v1.P.X
	vy = v2.Y - v1.P.Y
	v3 = Vec2{
		X: float32(v1.Q.C*vx) + float32(v1.Q.S*vy),
		Y: float32(-v1.Q.S*vx) + float32(v1.Q.C*vy),
	}
	goto _4
_4:
	(*b2JointSim)(unsafe.Pointer(joint)).LocalOriginAnchorA = v3
	v5 = transformB
	v6 = (*MouseJointDef)(unsafe.Pointer(def)).Target
	vx = v6.X - v5.P.X
	vy = v6.Y - v5.P.Y
	v7 = Vec2{
		X: float32(v5.Q.C*vx) + float32(v5.Q.S*vy),
		Y: float32(-v5.Q.S*vx) + float32(v5.Q.C*vy),
	}
	goto _8
_8:
	(*b2JointSim)(unsafe.Pointer(joint)).LocalOriginAnchorB = v7
	empty = b2MouseJoint{}
	*(*b2MouseJoint)(unsafe.Add(unsafe.Pointer(joint), 68)) = empty
	(*(*b2MouseJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).TargetA = (*MouseJointDef)(unsafe.Pointer(def)).Target
	(*(*b2MouseJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).Hertz = (*MouseJointDef)(unsafe.Pointer(def)).Hertz
	(*(*b2MouseJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).DampingRatio = (*MouseJointDef)(unsafe.Pointer(def)).DampingRatio
	(*(*b2MouseJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MaxForce = (*MouseJointDef)(unsafe.Pointer(def)).MaxForce
	jointId = JointId{
		Index1:     (*b2JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	return jointId
}

func b2CreateFilterJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	var bodyA, bodyB, joint, world uintptr
	var collideConnected uint8
	var jointId JointId
	var pair b2JointPair
	_, _, _, _, _, _, _ = bodyA, bodyB, collideConnected, joint, jointId, pair, world
	if !((*FilterJointDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)) && b2InternalAssertFcn(tls, __ccgo_ts+730, __ccgo_ts+9061, int32FromInt32(482)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorldFromId(tls, worldId)
	if !(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)) && b2InternalAssertFcn(tls, __ccgo_ts+1152, __ccgo_ts+9061, int32FromInt32(485)) != 0 {
		__builtin_trap(tls)
	}
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return JointId{}
	}
	bodyA = b2GetBodyFullId(tls, world, (*FilterJointDef)(unsafe.Pointer(def)).BodyIdA)
	bodyB = b2GetBodyFullId(tls, world, (*FilterJointDef)(unsafe.Pointer(def)).BodyIdB)
	collideConnected = boolUint8(false1 != 0)
	pair = b2CreateJoint(tls, world, bodyA, bodyB, (*FilterJointDef)(unsafe.Pointer(def)).UserData, float32FromFloat32(1), int32(b2_filterJoint), collideConnected)
	joint = pair.JointSim
	(*b2JointSim)(unsafe.Pointer(joint)).Type1 = int32(b2_filterJoint)
	(*b2JointSim)(unsafe.Pointer(joint)).LocalOriginAnchorA = b2Vec2_zero
	(*b2JointSim)(unsafe.Pointer(joint)).LocalOriginAnchorB = b2Vec2_zero
	jointId = JointId{
		Index1:     (*b2JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	return jointId
}

func b2CreateRevoluteJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	var bodyA, bodyB, joint, world uintptr
	var empty b2RevoluteJoint
	var jointId JointId
	var pair b2JointPair
	var v1, v10, v11, v13, v14, v2, v3, v4, v6, v7, v8, v9 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, empty, joint, jointId, pair, world, v1, v10, v11, v13, v14, v2, v3, v4, v6, v7, v8, v9
	if !((*RevoluteJointDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)) && b2InternalAssertFcn(tls, __ccgo_ts+730, __ccgo_ts+9061, int32FromInt32(509)) != 0 {
		__builtin_trap(tls)
	}
	if !((*RevoluteJointDef)(unsafe.Pointer(def)).LowerAngle <= (*RevoluteJointDef)(unsafe.Pointer(def)).UpperAngle) && b2InternalAssertFcn(tls, __ccgo_ts+9472, __ccgo_ts+9061, int32FromInt32(510)) != 0 {
		__builtin_trap(tls)
	}
	if !((*RevoluteJointDef)(unsafe.Pointer(def)).LowerAngle >= float32(-float32FromFloat32(0.99)*float32FromFloat32(3.14159265359))) && b2InternalAssertFcn(tls, __ccgo_ts+9507, __ccgo_ts+9061, int32FromInt32(511)) != 0 {
		__builtin_trap(tls)
	}
	if !((*RevoluteJointDef)(unsafe.Pointer(def)).UpperAngle <= float32(float32FromFloat32(0.99)*float32FromFloat32(3.14159265359))) && b2InternalAssertFcn(tls, __ccgo_ts+9541, __ccgo_ts+9061, int32FromInt32(512)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorldFromId(tls, worldId)
	if !(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)) && b2InternalAssertFcn(tls, __ccgo_ts+1152, __ccgo_ts+9061, int32FromInt32(516)) != 0 {
		__builtin_trap(tls)
	}
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return JointId{}
	}
	bodyA = b2GetBodyFullId(tls, world, (*RevoluteJointDef)(unsafe.Pointer(def)).BodyIdA)
	bodyB = b2GetBodyFullId(tls, world, (*RevoluteJointDef)(unsafe.Pointer(def)).BodyIdB)
	pair = b2CreateJoint(tls, world, bodyA, bodyB, (*RevoluteJointDef)(unsafe.Pointer(def)).UserData, (*RevoluteJointDef)(unsafe.Pointer(def)).DrawSize, int32(b2_revoluteJoint), (*RevoluteJointDef)(unsafe.Pointer(def)).CollideConnected)
	joint = pair.JointSim
	(*b2JointSim)(unsafe.Pointer(joint)).Type1 = int32(b2_revoluteJoint)
	(*b2JointSim)(unsafe.Pointer(joint)).LocalOriginAnchorA = (*RevoluteJointDef)(unsafe.Pointer(def)).LocalAnchorA
	(*b2JointSim)(unsafe.Pointer(joint)).LocalOriginAnchorB = (*RevoluteJointDef)(unsafe.Pointer(def)).LocalAnchorB
	empty = b2RevoluteJoint{}
	*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68)) = empty
	v1 = (*RevoluteJointDef)(unsafe.Pointer(def)).ReferenceAngle
	v2 = -float32FromFloat32(3.14159265359)
	v3 = float32FromFloat32(3.14159265359)
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
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).ReferenceAngle = v4
	v8 = (*RevoluteJointDef)(unsafe.Pointer(def)).TargetAngle
	v9 = -float32FromFloat32(3.14159265359)
	v10 = float32FromFloat32(3.14159265359)
	if v8 < v9 {
		v13 = v9
	} else {
		if v8 > v10 {
			v14 = v10
		} else {
			v14 = v8
		}
		v13 = v14
	}
	v11 = v13
	goto _12
_12:
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).TargetAngle = v11
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).Hertz = (*RevoluteJointDef)(unsafe.Pointer(def)).Hertz
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).DampingRatio = (*RevoluteJointDef)(unsafe.Pointer(def)).DampingRatio
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LowerAngle = (*RevoluteJointDef)(unsafe.Pointer(def)).LowerAngle
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).UpperAngle = (*RevoluteJointDef)(unsafe.Pointer(def)).UpperAngle
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MaxMotorTorque = (*RevoluteJointDef)(unsafe.Pointer(def)).MaxMotorTorque
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MotorSpeed = (*RevoluteJointDef)(unsafe.Pointer(def)).MotorSpeed
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableSpring = (*RevoluteJointDef)(unsafe.Pointer(def)).EnableSpring
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableLimit = (*RevoluteJointDef)(unsafe.Pointer(def)).EnableLimit
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableMotor = (*RevoluteJointDef)(unsafe.Pointer(def)).EnableMotor
	// If the joint prevents collisions, then destroy all contacts between attached bodies
	if int32FromUint8((*RevoluteJointDef)(unsafe.Pointer(def)).CollideConnected) == false1 {
		b2DestroyContactsBetweenBodies(tls, world, bodyA, bodyB)
	}
	jointId = JointId{
		Index1:     (*b2JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	return jointId
}

func b2CreatePrismaticJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	var bodyA, bodyB, joint, world uintptr
	var empty b2PrismaticJoint
	var invLength, length float32
	var jointId JointId
	var n, v1, v2 Vec2
	var pair b2JointPair
	_, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, empty, invLength, joint, jointId, length, n, pair, world, v1, v2
	if !((*PrismaticJointDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)) && b2InternalAssertFcn(tls, __ccgo_ts+730, __ccgo_ts+9061, int32FromInt32(561)) != 0 {
		__builtin_trap(tls)
	}
	if !((*PrismaticJointDef)(unsafe.Pointer(def)).LowerTranslation <= (*PrismaticJointDef)(unsafe.Pointer(def)).UpperTranslation) && b2InternalAssertFcn(tls, __ccgo_ts+9574, __ccgo_ts+9061, int32FromInt32(562)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorldFromId(tls, worldId)
	if !(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)) && b2InternalAssertFcn(tls, __ccgo_ts+1152, __ccgo_ts+9061, int32FromInt32(566)) != 0 {
		__builtin_trap(tls)
	}
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return JointId{}
	}
	bodyA = b2GetBodyFullId(tls, world, (*PrismaticJointDef)(unsafe.Pointer(def)).BodyIdA)
	bodyB = b2GetBodyFullId(tls, world, (*PrismaticJointDef)(unsafe.Pointer(def)).BodyIdB)
	pair = b2CreateJoint(tls, world, bodyA, bodyB, (*PrismaticJointDef)(unsafe.Pointer(def)).UserData, float32FromFloat32(1), int32(b2_prismaticJoint), (*PrismaticJointDef)(unsafe.Pointer(def)).CollideConnected)
	joint = pair.JointSim
	(*b2JointSim)(unsafe.Pointer(joint)).Type1 = int32(b2_prismaticJoint)
	(*b2JointSim)(unsafe.Pointer(joint)).LocalOriginAnchorA = (*PrismaticJointDef)(unsafe.Pointer(def)).LocalAnchorA
	(*b2JointSim)(unsafe.Pointer(joint)).LocalOriginAnchorB = (*PrismaticJointDef)(unsafe.Pointer(def)).LocalAnchorB
	empty = b2PrismaticJoint{}
	*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68)) = empty
	v1 = (*PrismaticJointDef)(unsafe.Pointer(def)).LocalAxisA
	length = sqrtf(tls, float32(v1.X*v1.X)+float32(v1.Y*v1.Y))
	if length < float32FromFloat32(1.1920928955078125e-07) {
		v2 = Vec2{}
		goto _3
	}
	invLength = float32FromFloat32(1) / length
	n = Vec2{
		X: float32(invLength * v1.X),
		Y: float32(invLength * v1.Y),
	}
	v2 = n
	goto _3
_3:
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LocalAxisA = v2
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).ReferenceAngle = (*PrismaticJointDef)(unsafe.Pointer(def)).ReferenceAngle
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).TargetTranslation = (*PrismaticJointDef)(unsafe.Pointer(def)).TargetTranslation
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).Hertz = (*PrismaticJointDef)(unsafe.Pointer(def)).Hertz
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).DampingRatio = (*PrismaticJointDef)(unsafe.Pointer(def)).DampingRatio
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LowerTranslation = (*PrismaticJointDef)(unsafe.Pointer(def)).LowerTranslation
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).UpperTranslation = (*PrismaticJointDef)(unsafe.Pointer(def)).UpperTranslation
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MaxMotorForce = (*PrismaticJointDef)(unsafe.Pointer(def)).MaxMotorForce
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MotorSpeed = (*PrismaticJointDef)(unsafe.Pointer(def)).MotorSpeed
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableSpring = (*PrismaticJointDef)(unsafe.Pointer(def)).EnableSpring
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableLimit = (*PrismaticJointDef)(unsafe.Pointer(def)).EnableLimit
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableMotor = (*PrismaticJointDef)(unsafe.Pointer(def)).EnableMotor
	// If the joint prevents collisions, then destroy all contacts between attached bodies
	if int32FromUint8((*PrismaticJointDef)(unsafe.Pointer(def)).CollideConnected) == false1 {
		b2DestroyContactsBetweenBodies(tls, world, bodyA, bodyB)
	}
	jointId = JointId{
		Index1:     (*b2JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	return jointId
}

func b2CreateWeldJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	var bodyA, bodyB, joint, world uintptr
	var empty b2WeldJoint
	var jointId JointId
	var pair b2JointPair
	_, _, _, _, _, _, _ = bodyA, bodyB, empty, joint, jointId, pair, world
	if !((*WeldJointDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)) && b2InternalAssertFcn(tls, __ccgo_ts+730, __ccgo_ts+9061, int32FromInt32(611)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorldFromId(tls, worldId)
	if !(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)) && b2InternalAssertFcn(tls, __ccgo_ts+1152, __ccgo_ts+9061, int32FromInt32(614)) != 0 {
		__builtin_trap(tls)
	}
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return JointId{}
	}
	bodyA = b2GetBodyFullId(tls, world, (*WeldJointDef)(unsafe.Pointer(def)).BodyIdA)
	bodyB = b2GetBodyFullId(tls, world, (*WeldJointDef)(unsafe.Pointer(def)).BodyIdB)
	pair = b2CreateJoint(tls, world, bodyA, bodyB, (*WeldJointDef)(unsafe.Pointer(def)).UserData, float32FromFloat32(1), int32(b2_weldJoint), (*WeldJointDef)(unsafe.Pointer(def)).CollideConnected)
	joint = pair.JointSim
	(*b2JointSim)(unsafe.Pointer(joint)).Type1 = int32(b2_weldJoint)
	(*b2JointSim)(unsafe.Pointer(joint)).LocalOriginAnchorA = (*WeldJointDef)(unsafe.Pointer(def)).LocalAnchorA
	(*b2JointSim)(unsafe.Pointer(joint)).LocalOriginAnchorB = (*WeldJointDef)(unsafe.Pointer(def)).LocalAnchorB
	empty = b2WeldJoint{}
	*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 68)) = empty
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).ReferenceAngle = (*WeldJointDef)(unsafe.Pointer(def)).ReferenceAngle
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LinearHertz = (*WeldJointDef)(unsafe.Pointer(def)).LinearHertz
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LinearDampingRatio = (*WeldJointDef)(unsafe.Pointer(def)).LinearDampingRatio
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).AngularHertz = (*WeldJointDef)(unsafe.Pointer(def)).AngularHertz
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).AngularDampingRatio = (*WeldJointDef)(unsafe.Pointer(def)).AngularDampingRatio
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LinearImpulse = b2Vec2_zero
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).AngularImpulse = float32FromFloat32(0)
	// If the joint prevents collisions, then destroy all contacts between attached bodies
	if int32FromUint8((*WeldJointDef)(unsafe.Pointer(def)).CollideConnected) == false1 {
		b2DestroyContactsBetweenBodies(tls, world, bodyA, bodyB)
	}
	jointId = JointId{
		Index1:     (*b2JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	return jointId
}

func b2CreateWheelJoint(tls *_Stack, worldId WorldId, def uintptr) (r JointId) {
	var bodyA, bodyB, joint, world uintptr
	var invLength, length float32
	var jointId JointId
	var n, v1, v2 Vec2
	var pair b2JointPair
	_, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, invLength, joint, jointId, length, n, pair, world, v1, v2
	if !((*WheelJointDef)(unsafe.Pointer(def)).InternalValue == int32FromInt32(_B2_SECRET_COOKIE)) && b2InternalAssertFcn(tls, __ccgo_ts+730, __ccgo_ts+9061, int32FromInt32(653)) != 0 {
		__builtin_trap(tls)
	}
	if !((*WheelJointDef)(unsafe.Pointer(def)).LowerTranslation <= (*WheelJointDef)(unsafe.Pointer(def)).UpperTranslation) && b2InternalAssertFcn(tls, __ccgo_ts+9574, __ccgo_ts+9061, int32FromInt32(654)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorldFromId(tls, worldId)
	if !(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)) && b2InternalAssertFcn(tls, __ccgo_ts+1152, __ccgo_ts+9061, int32FromInt32(658)) != 0 {
		__builtin_trap(tls)
	}
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return JointId{}
	}
	bodyA = b2GetBodyFullId(tls, world, (*WheelJointDef)(unsafe.Pointer(def)).BodyIdA)
	bodyB = b2GetBodyFullId(tls, world, (*WheelJointDef)(unsafe.Pointer(def)).BodyIdB)
	pair = b2CreateJoint(tls, world, bodyA, bodyB, (*WheelJointDef)(unsafe.Pointer(def)).UserData, float32FromFloat32(1), int32(b2_wheelJoint), (*WheelJointDef)(unsafe.Pointer(def)).CollideConnected)
	joint = pair.JointSim
	(*b2JointSim)(unsafe.Pointer(joint)).Type1 = int32(b2_wheelJoint)
	(*b2JointSim)(unsafe.Pointer(joint)).LocalOriginAnchorA = (*WheelJointDef)(unsafe.Pointer(def)).LocalAnchorA
	(*b2JointSim)(unsafe.Pointer(joint)).LocalOriginAnchorB = (*WheelJointDef)(unsafe.Pointer(def)).LocalAnchorB
	*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68)) = b2WheelJoint{}
	v1 = (*WheelJointDef)(unsafe.Pointer(def)).LocalAxisA
	length = sqrtf(tls, float32(v1.X*v1.X)+float32(v1.Y*v1.Y))
	if length < float32FromFloat32(1.1920928955078125e-07) {
		v2 = Vec2{}
		goto _3
	}
	invLength = float32FromFloat32(1) / length
	n = Vec2{
		X: float32(invLength * v1.X),
		Y: float32(invLength * v1.Y),
	}
	v2 = n
	goto _3
_3:
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LocalAxisA = v2
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).PerpMass = float32FromFloat32(0)
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).AxialMass = float32FromFloat32(0)
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MotorImpulse = float32FromFloat32(0)
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LowerImpulse = float32FromFloat32(0)
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).UpperImpulse = float32FromFloat32(0)
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LowerTranslation = (*WheelJointDef)(unsafe.Pointer(def)).LowerTranslation
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).UpperTranslation = (*WheelJointDef)(unsafe.Pointer(def)).UpperTranslation
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MaxMotorTorque = (*WheelJointDef)(unsafe.Pointer(def)).MaxMotorTorque
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MotorSpeed = (*WheelJointDef)(unsafe.Pointer(def)).MotorSpeed
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).Hertz = (*WheelJointDef)(unsafe.Pointer(def)).Hertz
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).DampingRatio = (*WheelJointDef)(unsafe.Pointer(def)).DampingRatio
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableSpring = (*WheelJointDef)(unsafe.Pointer(def)).EnableSpring
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableLimit = (*WheelJointDef)(unsafe.Pointer(def)).EnableLimit
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableMotor = (*WheelJointDef)(unsafe.Pointer(def)).EnableMotor
	// If the joint prevents collisions, then destroy all contacts between attached bodies
	if int32FromUint8((*WheelJointDef)(unsafe.Pointer(def)).CollideConnected) == false1 {
		b2DestroyContactsBetweenBodies(tls, world, bodyA, bodyB)
	}
	jointId = JointId{
		Index1:     (*b2JointSim)(unsafe.Pointer(joint)).JointId + int32(1),
		World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
		Generation: (*b2Joint)(unsafe.Pointer(pair.Joint)).Generation,
	}
	return jointId
}

func b2DestroyJointInternal(tls *_Stack, world uintptr, joint uintptr, wakeBodies uint8) {
	var bodyA, bodyB, edgeA, edgeB, movedJoint, movedJointSim, nextEdge, nextEdge1, nextJoint, nextJoint1, prevEdge, prevEdge1, prevJoint, prevJoint1, set, v1, v11, v13, v15, v17, v19, v21, v23, v25, v27, v29, v3, v33, v35, v5, v7, v9 uintptr
	var edgeKeyA, edgeKeyB, idA, idB, jointId, localIndex, movedId, movedIndex, movedIndex1, setIndex, v10, v14, v18, v2, v22, v26, v30, v31, v34, v6 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, edgeA, edgeB, edgeKeyA, edgeKeyB, idA, idB, jointId, localIndex, movedId, movedIndex, movedIndex1, movedJoint, movedJointSim, nextEdge, nextEdge1, nextJoint, nextJoint1, prevEdge, prevEdge1, prevJoint, prevJoint1, set, setIndex, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v29, v3, v30, v31, v33, v34, v35, v5, v6, v7, v9
	jointId = (*b2Joint)(unsafe.Pointer(joint)).JointId
	edgeA = joint + 20 + uintptr(0)*12
	edgeB = joint + 20 + uintptr(1)*12
	idA = (*b2JointEdge)(unsafe.Pointer(edgeA)).BodyId
	idB = (*b2JointEdge)(unsafe.Pointer(edgeB)).BodyId
	v1 = world + 1024
	v2 = idA
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	bodyA = v3
	v5 = world + 1024
	v6 = idB
	if !(0 <= v6 && v6 < (*b2BodyArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodyArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*128
	goto _8
_8:
	bodyB = v7
	// Remove from body A
	if (*b2JointEdge)(unsafe.Pointer(edgeA)).PrevKey != -int32(1) {
		v9 = world + 1104
		v10 = (*b2JointEdge)(unsafe.Pointer(edgeA)).PrevKey >> int32(1)
		if !(0 <= v10 && v10 < (*b2JointArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v11 = (*b2JointArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*72
		goto _12
	_12:
		prevJoint = v11
		prevEdge = prevJoint + 20 + uintptr((*b2JointEdge)(unsafe.Pointer(edgeA)).PrevKey&int32FromInt32(1))*12
		(*b2JointEdge)(unsafe.Pointer(prevEdge)).NextKey = (*b2JointEdge)(unsafe.Pointer(edgeA)).NextKey
	}
	if (*b2JointEdge)(unsafe.Pointer(edgeA)).NextKey != -int32(1) {
		v13 = world + 1104
		v14 = (*b2JointEdge)(unsafe.Pointer(edgeA)).NextKey >> int32(1)
		if !(0 <= v14 && v14 < (*b2JointArray)(unsafe.Pointer(v13)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v15 = (*b2JointArray)(unsafe.Pointer(v13)).Data + uintptr(v14)*72
		goto _16
	_16:
		nextJoint = v15
		nextEdge = nextJoint + 20 + uintptr((*b2JointEdge)(unsafe.Pointer(edgeA)).NextKey&int32FromInt32(1))*12
		(*b2JointEdge)(unsafe.Pointer(nextEdge)).PrevKey = (*b2JointEdge)(unsafe.Pointer(edgeA)).PrevKey
	}
	edgeKeyA = jointId<<int32(1) | 0
	if (*b2Body)(unsafe.Pointer(bodyA)).HeadJointKey == edgeKeyA {
		(*b2Body)(unsafe.Pointer(bodyA)).HeadJointKey = (*b2JointEdge)(unsafe.Pointer(edgeA)).NextKey
	}
	*(*int32)(unsafe.Pointer(bodyA + 72)) -= int32(1)
	// Remove from body B
	if (*b2JointEdge)(unsafe.Pointer(edgeB)).PrevKey != -int32(1) {
		v17 = world + 1104
		v18 = (*b2JointEdge)(unsafe.Pointer(edgeB)).PrevKey >> int32(1)
		if !(0 <= v18 && v18 < (*b2JointArray)(unsafe.Pointer(v17)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v19 = (*b2JointArray)(unsafe.Pointer(v17)).Data + uintptr(v18)*72
		goto _20
	_20:
		prevJoint1 = v19
		prevEdge1 = prevJoint1 + 20 + uintptr((*b2JointEdge)(unsafe.Pointer(edgeB)).PrevKey&int32FromInt32(1))*12
		(*b2JointEdge)(unsafe.Pointer(prevEdge1)).NextKey = (*b2JointEdge)(unsafe.Pointer(edgeB)).NextKey
	}
	if (*b2JointEdge)(unsafe.Pointer(edgeB)).NextKey != -int32(1) {
		v21 = world + 1104
		v22 = (*b2JointEdge)(unsafe.Pointer(edgeB)).NextKey >> int32(1)
		if !(0 <= v22 && v22 < (*b2JointArray)(unsafe.Pointer(v21)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
			__builtin_trap(tls)
		}
		v23 = (*b2JointArray)(unsafe.Pointer(v21)).Data + uintptr(v22)*72
		goto _24
	_24:
		nextJoint1 = v23
		nextEdge1 = nextJoint1 + 20 + uintptr((*b2JointEdge)(unsafe.Pointer(edgeB)).NextKey&int32FromInt32(1))*12
		(*b2JointEdge)(unsafe.Pointer(nextEdge1)).PrevKey = (*b2JointEdge)(unsafe.Pointer(edgeB)).PrevKey
	}
	edgeKeyB = jointId<<int32(1) | int32(1)
	if (*b2Body)(unsafe.Pointer(bodyB)).HeadJointKey == edgeKeyB {
		(*b2Body)(unsafe.Pointer(bodyB)).HeadJointKey = (*b2JointEdge)(unsafe.Pointer(edgeB)).NextKey
	}
	*(*int32)(unsafe.Pointer(bodyB + 72)) -= int32(1)
	if (*b2Joint)(unsafe.Pointer(joint)).IslandId != -int32(1) {
		if !((*b2Joint)(unsafe.Pointer(joint)).SetIndex > int32(b2_disabledSet)) && b2InternalAssertFcn(tls, __ccgo_ts+9621, __ccgo_ts+9061, int32FromInt32(762)) != 0 {
			__builtin_trap(tls)
		}
		b2UnlinkJoint(tls, world, joint)
	} else {
		if !((*b2Joint)(unsafe.Pointer(joint)).SetIndex <= int32(b2_disabledSet)) && b2InternalAssertFcn(tls, __ccgo_ts+9654, __ccgo_ts+9061, int32FromInt32(767)) != 0 {
			__builtin_trap(tls)
		}
	}
	// Remove joint from solver set that owns it
	setIndex = (*b2Joint)(unsafe.Pointer(joint)).SetIndex
	localIndex = (*b2Joint)(unsafe.Pointer(joint)).LocalIndex
	if setIndex == int32(b2_awakeSet) {
		b2RemoveJointFromGraph(tls, world, (*(*b2JointEdge)(unsafe.Pointer(joint + 20))).BodyId, (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + 1*12))).BodyId, (*b2Joint)(unsafe.Pointer(joint)).ColorIndex, localIndex)
	} else {
		v25 = world + 1064
		v26 = setIndex
		if !(0 <= v26 && v26 < (*b2SolverSetArray)(unsafe.Pointer(v25)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
			__builtin_trap(tls)
		}
		v27 = (*b2SolverSetArray)(unsafe.Pointer(v25)).Data + uintptr(v26)*88
		goto _28
	_28:
		set = v27
		v29 = set + 32
		v30 = localIndex
		if !(0 <= v30 && v30 < (*b2JointSimArray)(unsafe.Pointer(v29)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(342)) != 0 {
			__builtin_trap(tls)
		}
		movedIndex = -int32(1)
		if v30 != (*b2JointSimArray)(unsafe.Pointer(v29)).Count-int32FromInt32(1) {
			movedIndex = (*b2JointSimArray)(unsafe.Pointer(v29)).Count - int32(1)
			*(*b2JointSim)(unsafe.Pointer((*b2JointSimArray)(unsafe.Pointer(v29)).Data + uintptr(v30)*196)) = *(*b2JointSim)(unsafe.Pointer((*b2JointSimArray)(unsafe.Pointer(v29)).Data + uintptr(movedIndex)*196))
		}
		*(*int32)(unsafe.Pointer(v29 + 8)) -= int32(1)
		v31 = movedIndex
		goto _32
	_32:
		movedIndex1 = v31
		if movedIndex1 != -int32(1) {
			// Fix moved joint
			movedJointSim = (*b2SolverSet)(unsafe.Pointer(set)).JointSims.Data + uintptr(localIndex)*196
			movedId = (*b2JointSim)(unsafe.Pointer(movedJointSim)).JointId
			v33 = world + 1104
			v34 = movedId
			if !(0 <= v34 && v34 < (*b2JointArray)(unsafe.Pointer(v33)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
				__builtin_trap(tls)
			}
			v35 = (*b2JointArray)(unsafe.Pointer(v33)).Data + uintptr(v34)*72
			goto _36
		_36:
			movedJoint = v35
			if !((*b2Joint)(unsafe.Pointer(movedJoint)).LocalIndex == movedIndex1) && b2InternalAssertFcn(tls, __ccgo_ts+3280, __ccgo_ts+9061, int32FromInt32(788)) != 0 {
				__builtin_trap(tls)
			}
			(*b2Joint)(unsafe.Pointer(movedJoint)).LocalIndex = localIndex
		}
	}
	// Free joint and id (preserve joint generation)
	(*b2Joint)(unsafe.Pointer(joint)).SetIndex = -int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).LocalIndex = -int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).ColorIndex = -int32(1)
	(*b2Joint)(unsafe.Pointer(joint)).JointId = -int32(1)
	b2FreeId(tls, world+1080, jointId)
	if wakeBodies != 0 {
		b2WakeBody(tls, world, bodyA)
		b2WakeBody(tls, world, bodyB)
	}
	b2ValidateSolverSets(tls, world)
}

func b2DestroyJoint(tls *_Stack, jointId JointId) {
	var joint, world uintptr
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	if !(int32FromUint8((*b2World)(unsafe.Pointer(world)).Locked) == int32FromInt32(false1)) && b2InternalAssertFcn(tls, __ccgo_ts+1152, __ccgo_ts+9061, int32FromInt32(812)) != 0 {
		__builtin_trap(tls)
	}
	if (*b2World)(unsafe.Pointer(world)).Locked != 0 {
		return
	}
	joint = b2GetJointFullId(tls, world, jointId)
	b2DestroyJointInternal(tls, world, joint, boolUint8(true1 != 0))
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
	return b2MakeBodyId(tls, world, (*(*b2JointEdge)(unsafe.Pointer(joint + 20))).BodyId)
}

func b2Joint_GetBodyB(tls *_Stack, jointId JointId) (r BodyId) {
	var joint, world uintptr
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	return b2MakeBodyId(tls, world, (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + 1*12))).BodyId)
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

func b2Joint_SetLocalAnchorA(tls *_Stack, jointId JointId, localAnchor Vec2) {
	var joint, jointSim, world uintptr
	_, _, _ = joint, jointSim, world
	if !(b2IsValidVec2(tls, localAnchor) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+9688, __ccgo_ts+9061, int32FromInt32(853)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	jointSim = b2GetJointSim(tls, world, joint)
	(*b2JointSim)(unsafe.Pointer(jointSim)).LocalOriginAnchorA = localAnchor
}

func b2Joint_GetLocalAnchorA(tls *_Stack, jointId JointId) (r Vec2) {
	var joint, jointSim, world uintptr
	_, _, _ = joint, jointSim, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	jointSim = b2GetJointSim(tls, world, joint)
	return (*b2JointSim)(unsafe.Pointer(jointSim)).LocalOriginAnchorA
}

func b2Joint_SetLocalAnchorB(tls *_Stack, jointId JointId, localAnchor Vec2) {
	var joint, jointSim, world uintptr
	_, _, _ = joint, jointSim, world
	if !(b2IsValidVec2(tls, localAnchor) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+9688, __ccgo_ts+9061, int32FromInt32(871)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	jointSim = b2GetJointSim(tls, world, joint)
	(*b2JointSim)(unsafe.Pointer(jointSim)).LocalOriginAnchorB = localAnchor
}

func b2Joint_GetLocalAnchorB(tls *_Stack, jointId JointId) (r Vec2) {
	var joint, jointSim, world uintptr
	_, _, _ = joint, jointSim, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	jointSim = b2GetJointSim(tls, world, joint)
	return (*b2JointSim)(unsafe.Pointer(jointSim)).LocalOriginAnchorB
}

func b2Joint_SetReferenceAngle(tls *_Stack, jointId JointId, angleInRadians float32) {
	var joint, jointSim, world uintptr
	_, _, _ = joint, jointSim, world
	if !(b2IsValidFloat(tls, angleInRadians) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+9717, __ccgo_ts+9061, int32FromInt32(889)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	jointSim = b2GetJointSim(tls, world, joint)
	switch (*b2Joint)(unsafe.Pointer(joint)).Type1 {
	case int32(b2_prismaticJoint):
		(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(jointSim), 68))).ReferenceAngle = angleInRadians
	case int32(b2_revoluteJoint):
		(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(jointSim), 68))).ReferenceAngle = angleInRadians
	case int32(b2_weldJoint):
		(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(jointSim), 68))).ReferenceAngle = angleInRadians
	default:
		break
	}
}

func b2Joint_GetReferenceAngle(tls *_Stack, jointId JointId) (r float32) {
	var joint, jointSim, world uintptr
	_, _, _ = joint, jointSim, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	jointSim = b2GetJointSim(tls, world, joint)
	switch (*b2Joint)(unsafe.Pointer(joint)).Type1 {
	case int32(b2_prismaticJoint):
		return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(jointSim), 68))).ReferenceAngle
	case int32(b2_revoluteJoint):
		return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(jointSim), 68))).ReferenceAngle
	case int32(b2_weldJoint):
		return (*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(jointSim), 68))).ReferenceAngle
	default:
		return float32FromFloat32(0)
	}
	return r
}

func b2Joint_SetLocalAxisA(tls *_Stack, jointId JointId, localAxis Vec2) {
	var aa, v11, v4, v8, v9 float32
	var joint, jointSim, world uintptr
	var v1, v2, v3 Vec2
	var v6 uint8
	_, _, _, _, _, _, _, _, _, _, _, _ = aa, joint, jointSim, world, v1, v11, v2, v3, v4, v6, v8, v9
	if !(b2IsValidVec2(tls, localAxis) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+9750, __ccgo_ts+9061, int32FromInt32(938)) != 0 {
		__builtin_trap(tls)
	}
	v1 = localAxis
	v2 = v1
	v3 = v1
	v4 = float32(v2.X*v3.X) + float32(v2.Y*v3.Y)
	goto _5
_5:
	aa = v4
	v8 = float32FromFloat32(1) - aa
	if v8 < float32FromInt32(0) {
		v11 = -v8
	} else {
		v11 = v8
	}
	v9 = v11
	goto _10
_10:
	v6 = boolUint8(v9 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07)))
	goto _7
_7:
	;
	if !(v6 != 0) && b2InternalAssertFcn(tls, __ccgo_ts+9777, __ccgo_ts+9061, int32FromInt32(939)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	jointSim = b2GetJointSim(tls, world, joint)
	switch (*b2Joint)(unsafe.Pointer(joint)).Type1 {
	case int32(b2_prismaticJoint):
		(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(jointSim), 68))).LocalAxisA = localAxis
	case int32(b2_wheelJoint):
		(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(jointSim), 68))).LocalAxisA = localAxis
	default:
		break
	}
}

func b2Joint_GetLocalAxisA(tls *_Stack, jointId JointId) (r Vec2) {
	var joint, jointSim, world uintptr
	_, _, _ = joint, jointSim, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	jointSim = b2GetJointSim(tls, world, joint)
	switch (*b2Joint)(unsafe.Pointer(joint)).Type1 {
	case int32(b2_prismaticJoint):
		return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(jointSim), 68))).LocalAxisA
	case int32(b2_wheelJoint):
		return (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(jointSim), 68))).LocalAxisA
	default:
		return b2Vec2_zero
	}
	return r
}

func b2Joint_SetCollideConnected(tls *_Stack, jointId JointId, shouldCollide uint8) {
	var alreadyAdded uint8
	var bodyA, bodyB, joint, shape, world, v1, v10, v12, v14, v16, v3, v5, v7 uintptr
	var newCapacity, shapeCountA, shapeCountB, shapeId, v11, v15, v17, v2, v6, v9 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = alreadyAdded, bodyA, bodyB, joint, newCapacity, shape, shapeCountA, shapeCountB, shapeId, world, v1, v10, v11, v12, v14, v15, v16, v17, v2, v3, v5, v6, v7, v9
	world = b2GetWorldLocked(tls, int32FromUint16(jointId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	joint = b2GetJointFullId(tls, world, jointId)
	if (*b2Joint)(unsafe.Pointer(joint)).CollideConnected == shouldCollide {
		return
	}
	(*b2Joint)(unsafe.Pointer(joint)).CollideConnected = shouldCollide
	v1 = world + 1024
	v2 = (*(*b2JointEdge)(unsafe.Pointer(joint + 20))).BodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	bodyA = v3
	v5 = world + 1024
	v6 = (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + 1*12))).BodyId
	if !(0 <= v6 && v6 < (*b2BodyArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodyArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*128
	goto _8
_8:
	bodyB = v7
	if shouldCollide != 0 {
		// need to tell the broad-phase to look for new pairs for one of the
		// two bodies. Pick the one with the fewest shapes.
		shapeCountA = (*b2Body)(unsafe.Pointer(bodyA)).ShapeCount
		shapeCountB = (*b2Body)(unsafe.Pointer(bodyB)).ShapeCount
		if shapeCountA < shapeCountB {
			v9 = (*b2Body)(unsafe.Pointer(bodyA)).HeadShapeId
		} else {
			v9 = (*b2Body)(unsafe.Pointer(bodyB)).HeadShapeId
		}
		shapeId = v9
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
			if (*b2Shape)(unsafe.Pointer(shape)).ProxyKey != -int32(1) {
				v14 = world + 40
				v15 = (*b2Shape)(unsafe.Pointer(shape)).ProxyKey
				alreadyAdded = b2AddKey(tls, v14+216, uint64FromInt32(v15+int32(1)))
				if int32FromUint8(alreadyAdded) == int32FromInt32(false1) {
					v16 = v14 + 232
					if (*b2IntArray)(unsafe.Pointer(v16)).Count == (*b2IntArray)(unsafe.Pointer(v16)).Capacity {
						if (*b2IntArray)(unsafe.Pointer(v16)).Capacity < int32(2) {
							v17 = int32(2)
						} else {
							v17 = (*b2IntArray)(unsafe.Pointer(v16)).Capacity + (*b2IntArray)(unsafe.Pointer(v16)).Capacity>>int32(1)
						}
						newCapacity = v17
						b2IntArray_Reserve(tls, v16, newCapacity)
					}
					*(*int32)(unsafe.Pointer((*b2IntArray)(unsafe.Pointer(v16)).Data + uintptr((*b2IntArray)(unsafe.Pointer(v16)).Count)*4)) = v15
					*(*int32)(unsafe.Pointer(v16 + 8)) += int32(1)
				}
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
	var bodyA, bodyB, joint, world, v1, v3, v5, v7 uintptr
	var v2, v6 int32
	_, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, joint, world, v1, v2, v3, v5, v6, v7
	world = b2GetWorldLocked(tls, int32FromUint16(jointId.World0))
	if world == uintptrFromInt32(0) {
		return
	}
	joint = b2GetJointFullId(tls, world, jointId)
	v1 = world + 1024
	v2 = (*(*b2JointEdge)(unsafe.Pointer(joint + 20))).BodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	bodyA = v3
	v5 = world + 1024
	v6 = (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + 1*12))).BodyId
	if !(0 <= v6 && v6 < (*b2BodyArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodyArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*128
	goto _8
_8:
	bodyB = v7
	b2WakeBody(tls, world, bodyA)
	b2WakeBody(tls, world, bodyB)
}

func b2Joint_GetConstraintForce(tls *_Stack, jointId JointId) (r Vec2) {
	var base, joint, world uintptr
	_, _, _ = base, joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	base = b2GetJointSim(tls, world, joint)
	switch (*b2Joint)(unsafe.Pointer(joint)).Type1 {
	case int32(b2_distanceJoint):
		goto _1
	case int32(b2_motorJoint):
		goto _2
	case int32(b2_mouseJoint):
		goto _3
	case int32(b2_filterJoint):
		goto _4
	case int32(b2_prismaticJoint):
		goto _5
	case int32(b2_revoluteJoint):
		goto _6
	case int32(b2_weldJoint):
		goto _7
	case int32(b2_wheelJoint):
		goto _8
	default:
		goto _9
	}
	goto _10
_1:
	;
	return b2GetDistanceJointForce(tls, world, base)
_2:
	;
	return b2GetMotorJointForce(tls, world, base)
_3:
	;
	return b2GetMouseJointForce(tls, world, base)
_4:
	;
	return b2Vec2_zero
_5:
	;
	return b2GetPrismaticJointForce(tls, world, base)
_6:
	;
	return b2GetRevoluteJointForce(tls, world, base)
_7:
	;
	return b2GetWeldJointForce(tls, world, base)
_8:
	;
	return b2GetWheelJointForce(tls, world, base)
_9:
	;
_13:
	;
	if bool(!(int32FromInt32(false1) != 0)) && b2InternalAssertFcn(tls, __ccgo_ts+4123, __ccgo_ts+9061, int32FromInt32(1094)) != 0 {
		__builtin_trap(tls)
	}
	goto _12
_12:
	;
	if 0 != 0 {
		goto _13
	}
	goto _11
_11:
	;
	return b2Vec2_zero
_10:
	;
	return r
}

func b2Joint_GetConstraintTorque(tls *_Stack, jointId JointId) (r float32) {
	var base, joint, world uintptr
	_, _, _ = base, joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	base = b2GetJointSim(tls, world, joint)
	switch (*b2Joint)(unsafe.Pointer(joint)).Type1 {
	case int32(b2_distanceJoint):
		goto _1
	case int32(b2_motorJoint):
		goto _2
	case int32(b2_mouseJoint):
		goto _3
	case int32(b2_filterJoint):
		goto _4
	case int32(b2_prismaticJoint):
		goto _5
	case int32(b2_revoluteJoint):
		goto _6
	case int32(b2_weldJoint):
		goto _7
	case int32(b2_wheelJoint):
		goto _8
	default:
		goto _9
	}
	goto _10
_1:
	;
	return float32FromFloat32(0)
_2:
	;
	return b2GetMotorJointTorque(tls, world, base)
_3:
	;
	return b2GetMouseJointTorque(tls, world, base)
_4:
	;
	return float32FromFloat32(0)
_5:
	;
	return b2GetPrismaticJointTorque(tls, world, base)
_6:
	;
	return b2GetRevoluteJointTorque(tls, world, base)
_7:
	;
	return b2GetWeldJointTorque(tls, world, base)
_8:
	;
	return b2GetWheelJointTorque(tls, world, base)
_9:
	;
_13:
	;
	if bool(!(int32FromInt32(false1) != 0)) && b2InternalAssertFcn(tls, __ccgo_ts+4123, __ccgo_ts+9061, int32FromInt32(1132)) != 0 {
		__builtin_trap(tls)
	}
	goto _12
_12:
	;
	if 0 != 0 {
		goto _13
	}
	goto _11
_11:
	;
	return float32FromFloat32(0)
_10:
	;
	return r
}

func b2Joint_GetLinearSeparation(tls *_Stack, jointId JointId) (r float32) {
	var axisA, axisA1, dp, pA, pB, perpA, perpA1, v10, v11, v2, v23, v3, v31, v32, v34, v35, v37, v38, v45, v46, v49, v52, v56, v57, v59, v6, v60, v62, v63, v7, v70, v71, v9 Vec2
	var base, distanceJoint, joint, prismaticJoint, weldJoint, wheelJoint, world uintptr
	var length, limitSeparation, limitSeparation1, perpendicularSeparation, perpendicularSeparation1, translation, translation1, x, y, v24, v26, v27, v29, v39, v41, v42, v44, v47, v50, v53, v64, v66, v67, v69, v72 float32
	var xfA, xfB, v1, v5 Transform
	var v30, v55 Rot
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axisA, axisA1, base, distanceJoint, dp, joint, length, limitSeparation, limitSeparation1, pA, pB, perpA, perpA1, perpendicularSeparation, perpendicularSeparation1, prismaticJoint, translation, translation1, weldJoint, wheelJoint, world, x, xfA, xfB, y, v1, v10, v11, v2, v23, v24, v26, v27, v29, v3, v30, v31, v32, v34, v35, v37, v38, v39, v41, v42, v44, v45, v46, v47, v49, v5, v50, v52, v53, v55, v56, v57, v59, v6, v60, v62, v63, v64, v66, v67, v69, v7, v70, v71, v72, v9
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	base = b2GetJointSim(tls, world, joint)
	xfA = b2GetBodyTransform(tls, world, (*(*b2JointEdge)(unsafe.Pointer(joint + 20))).BodyId)
	xfB = b2GetBodyTransform(tls, world, (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + 1*12))).BodyId)
	v1 = xfA
	v2 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorA
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	pA = v3
	v5 = xfB
	v6 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorB
	x = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y) + v5.P.X
	y = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y) + v5.P.Y
	v7 = Vec2{
		X: x,
		Y: y,
	}
	goto _8
_8:
	pB = v7
	v9 = pB
	v10 = pA
	v11 = Vec2{
		X: v9.X - v10.X,
		Y: v9.Y - v10.Y,
	}
	goto _12
_12:
	dp = v11
	switch (*b2Joint)(unsafe.Pointer(joint)).Type1 {
	case int32(b2_distanceJoint):
		goto _13
	case int32(b2_motorJoint):
		goto _14
	case int32(b2_mouseJoint):
		goto _15
	case int32(b2_filterJoint):
		goto _16
	case int32(b2_prismaticJoint):
		goto _17
	case int32(b2_revoluteJoint):
		goto _18
	case int32(b2_weldJoint):
		goto _19
	case int32(b2_wheelJoint):
		goto _20
	default:
		goto _21
	}
	goto _22
_13:
	;
	distanceJoint = base + 68
	v23 = dp
	v24 = sqrtf(tls, float32(v23.X*v23.X)+float32(v23.Y*v23.Y))
	goto _25
_25:
	length = v24
	if (*b2DistanceJoint)(unsafe.Pointer(distanceJoint)).EnableSpring != 0 {
		if (*b2DistanceJoint)(unsafe.Pointer(distanceJoint)).EnableLimit != 0 {
			if length < (*b2DistanceJoint)(unsafe.Pointer(distanceJoint)).MinLength {
				return (*b2DistanceJoint)(unsafe.Pointer(distanceJoint)).MinLength - length
			} else {
				if length > (*b2DistanceJoint)(unsafe.Pointer(distanceJoint)).MaxLength {
					return length - (*b2DistanceJoint)(unsafe.Pointer(distanceJoint)).MaxLength
				}
			}
			return float32FromFloat32(0)
		}
		return float32FromFloat32(0)
	}
	v26 = length - (*b2DistanceJoint)(unsafe.Pointer(distanceJoint)).Length
	if v26 < float32FromInt32(0) {
		v29 = -v26
	} else {
		v29 = v26
	}
	v27 = v29
	goto _28
_28:
	return v27
_14:
	;
	return float32FromFloat32(0)
_15:
	;
	return float32FromFloat32(0)
_16:
	;
	return float32FromFloat32(0)
_17:
	;
	prismaticJoint = base + 68
	v30 = xfA.Q
	v31 = (*b2PrismaticJoint)(unsafe.Pointer(prismaticJoint)).LocalAxisA
	v32 = Vec2{
		X: float32(v30.C*v31.X) - float32(v30.S*v31.Y),
		Y: float32(v30.S*v31.X) + float32(v30.C*v31.Y),
	}
	goto _33
_33:
	axisA = v32
	v34 = axisA
	v35 = Vec2{
		X: -v34.Y,
		Y: v34.X,
	}
	goto _36
_36:
	perpA = v35
	v37 = perpA
	v38 = dp
	v39 = float32(v37.X*v38.X) + float32(v37.Y*v38.Y)
	goto _40
_40:
	v41 = v39
	if v41 < float32FromInt32(0) {
		v44 = -v41
	} else {
		v44 = v41
	}
	v42 = v44
	goto _43
_43:
	perpendicularSeparation = v42
	limitSeparation = float32FromFloat32(0)
	if (*b2PrismaticJoint)(unsafe.Pointer(prismaticJoint)).EnableLimit != 0 {
		v45 = axisA
		v46 = dp
		v47 = float32(v45.X*v46.X) + float32(v45.Y*v46.Y)
		goto _48
	_48:
		translation = v47
		if translation < (*b2PrismaticJoint)(unsafe.Pointer(prismaticJoint)).LowerTranslation {
			limitSeparation = (*b2PrismaticJoint)(unsafe.Pointer(prismaticJoint)).LowerTranslation - translation
		}
		if (*b2PrismaticJoint)(unsafe.Pointer(prismaticJoint)).UpperTranslation < translation {
			limitSeparation = translation - (*b2PrismaticJoint)(unsafe.Pointer(prismaticJoint)).UpperTranslation
		}
	}
	return sqrtf(tls, float32(perpendicularSeparation*perpendicularSeparation)+float32(limitSeparation*limitSeparation))
_18:
	;
	v49 = dp
	v50 = sqrtf(tls, float32(v49.X*v49.X)+float32(v49.Y*v49.Y))
	goto _51
_51:
	return v50
_19:
	;
	weldJoint = base + 68
	if (*b2WeldJoint)(unsafe.Pointer(weldJoint)).LinearHertz == float32FromFloat32(0) {
		v52 = dp
		v53 = sqrtf(tls, float32(v52.X*v52.X)+float32(v52.Y*v52.Y))
		goto _54
	_54:
		return v53
	}
	return float32FromFloat32(0)
_20:
	;
	wheelJoint = base + 68
	v55 = xfA.Q
	v56 = (*b2WheelJoint)(unsafe.Pointer(wheelJoint)).LocalAxisA
	v57 = Vec2{
		X: float32(v55.C*v56.X) - float32(v55.S*v56.Y),
		Y: float32(v55.S*v56.X) + float32(v55.C*v56.Y),
	}
	goto _58
_58:
	axisA1 = v57
	v59 = axisA1
	v60 = Vec2{
		X: -v59.Y,
		Y: v59.X,
	}
	goto _61
_61:
	perpA1 = v60
	v62 = perpA1
	v63 = dp
	v64 = float32(v62.X*v63.X) + float32(v62.Y*v63.Y)
	goto _65
_65:
	v66 = v64
	if v66 < float32FromInt32(0) {
		v69 = -v66
	} else {
		v69 = v66
	}
	v67 = v69
	goto _68
_68:
	perpendicularSeparation1 = v67
	limitSeparation1 = float32FromFloat32(0)
	if (*b2WheelJoint)(unsafe.Pointer(wheelJoint)).EnableLimit != 0 {
		v70 = axisA1
		v71 = dp
		v72 = float32(v70.X*v71.X) + float32(v70.Y*v71.Y)
		goto _73
	_73:
		translation1 = v72
		if translation1 < (*b2WheelJoint)(unsafe.Pointer(wheelJoint)).LowerTranslation {
			limitSeparation1 = (*b2WheelJoint)(unsafe.Pointer(wheelJoint)).LowerTranslation - translation1
		}
		if (*b2WheelJoint)(unsafe.Pointer(wheelJoint)).UpperTranslation < translation1 {
			limitSeparation1 = translation1 - (*b2WheelJoint)(unsafe.Pointer(wheelJoint)).UpperTranslation
		}
	}
	return sqrtf(tls, float32(perpendicularSeparation1*perpendicularSeparation1)+float32(limitSeparation1*limitSeparation1))
_21:
	;
_76:
	;
	if bool(!(int32FromInt32(false1) != 0)) && b2InternalAssertFcn(tls, __ccgo_ts+4123, __ccgo_ts+9061, int32FromInt32(1252)) != 0 {
		__builtin_trap(tls)
	}
	goto _75
_75:
	;
	if 0 != 0 {
		goto _76
	}
	goto _74
_74:
	;
	return float32FromFloat32(0)
_22:
	;
	return r
}

func b2Joint_GetAngularSeparation(tls *_Stack, jointId JointId) (r float32) {
	var angle, c, relativeAngle, s, v15, v17, v19, v3 float32
	var base, joint, prismaticJoint, revoluteJoint, weldJoint, world uintptr
	var xfA, xfB Transform
	var v1, v2 Rot
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = angle, base, c, joint, prismaticJoint, relativeAngle, revoluteJoint, s, weldJoint, world, xfA, xfB, v1, v15, v17, v19, v2, v3
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	base = b2GetJointSim(tls, world, joint)
	xfA = b2GetBodyTransform(tls, world, (*(*b2JointEdge)(unsafe.Pointer(joint + 20))).BodyId)
	xfB = b2GetBodyTransform(tls, world, (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + 1*12))).BodyId)
	v1 = xfB.Q
	v2 = xfA.Q
	s = float32(v1.S*v2.C) - float32(v1.C*v2.S)
	c = float32(v1.C*v2.C) + float32(v1.S*v2.S)
	v3 = b2Atan2(tls, s, c)
	goto _4
_4:
	relativeAngle = v3
	switch (*b2Joint)(unsafe.Pointer(joint)).Type1 {
	case int32(b2_distanceJoint):
		goto _5
	case int32(b2_motorJoint):
		goto _6
	case int32(b2_mouseJoint):
		goto _7
	case int32(b2_filterJoint):
		goto _8
	case int32(b2_prismaticJoint):
		goto _9
	case int32(b2_revoluteJoint):
		goto _10
	case int32(b2_weldJoint):
		goto _11
	case int32(b2_wheelJoint):
		goto _12
	default:
		goto _13
	}
	goto _14
_5:
	;
	return float32FromFloat32(0)
_6:
	;
	return float32FromFloat32(0)
_7:
	;
	return float32FromFloat32(0)
_8:
	;
	return float32FromFloat32(0)
_9:
	;
	prismaticJoint = base + 68
	v15 = __builtin_remainderf(tls, relativeAngle-(*b2PrismaticJoint)(unsafe.Pointer(prismaticJoint)).ReferenceAngle, float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)))
	goto _16
_16:
	return v15
_10:
	;
	revoluteJoint = base + 68
	if (*b2RevoluteJoint)(unsafe.Pointer(revoluteJoint)).EnableLimit != 0 {
		v17 = __builtin_remainderf(tls, relativeAngle-(*b2RevoluteJoint)(unsafe.Pointer(revoluteJoint)).ReferenceAngle, float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)))
		goto _18
	_18:
		angle = v17
		if angle < (*b2RevoluteJoint)(unsafe.Pointer(revoluteJoint)).LowerAngle {
			return (*b2RevoluteJoint)(unsafe.Pointer(revoluteJoint)).LowerAngle - angle
		}
		if (*b2RevoluteJoint)(unsafe.Pointer(revoluteJoint)).UpperAngle < angle {
			return angle - (*b2RevoluteJoint)(unsafe.Pointer(revoluteJoint)).UpperAngle
		}
	}
	return float32FromFloat32(0)
_11:
	;
	weldJoint = base + 68
	if (*b2WeldJoint)(unsafe.Pointer(weldJoint)).AngularHertz == float32FromFloat32(0) {
		v19 = __builtin_remainderf(tls, relativeAngle-(*b2WeldJoint)(unsafe.Pointer(weldJoint)).ReferenceAngle, float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)))
		goto _20
	_20:
		return v19
	}
	return float32FromFloat32(0)
_12:
	;
	return float32FromFloat32(0)
_13:
	;
_23:
	;
	if bool(!(int32FromInt32(false1) != 0)) && b2InternalAssertFcn(tls, __ccgo_ts+4123, __ccgo_ts+9061, int32FromInt32(1322)) != 0 {
		__builtin_trap(tls)
	}
	goto _22
_22:
	;
	if 0 != 0 {
		goto _23
	}
	goto _21
_21:
	;
	return float32FromFloat32(0)
_14:
	;
	return r
}

func b2Joint_GetConstraintTuning(tls *_Stack, jointId JointId, hertz uintptr, dampingRatio uintptr) {
	var base, joint, world uintptr
	_, _, _ = base, joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	base = b2GetJointSim(tls, world, joint)
	*(*float32)(unsafe.Pointer(hertz)) = (*b2JointSim)(unsafe.Pointer(base)).ConstraintHertz
	*(*float32)(unsafe.Pointer(dampingRatio)) = (*b2JointSim)(unsafe.Pointer(base)).ConstraintDampingRatio
}

func b2Joint_SetConstraintTuning(tls *_Stack, jointId JointId, hertz float32, dampingRatio float32) {
	var base, joint, world uintptr
	_, _, _ = base, joint, world
	if !(b2IsValidFloat(tls, hertz) != 0 && hertz >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+9805, __ccgo_ts+9061, int32FromInt32(1338)) != 0 {
		__builtin_trap(tls)
	}
	if !(b2IsValidFloat(tls, dampingRatio) != 0 && dampingRatio >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+9846, __ccgo_ts+9061, int32FromInt32(1339)) != 0 {
		__builtin_trap(tls)
	}
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	base = b2GetJointSim(tls, world, joint)
	(*b2JointSim)(unsafe.Pointer(base)).ConstraintHertz = hertz
	(*b2JointSim)(unsafe.Pointer(base)).ConstraintDampingRatio = dampingRatio
}

func b2PrepareJoint(tls *_Stack, joint uintptr, context uintptr) {
	var a1, a2, a3, hertz1, omega, v1, v2, v3, v5, v6, v7 float32
	var v8 b2Softness
	_, _, _, _, _, _, _, _, _, _, _, _ = a1, a2, a3, hertz1, omega, v1, v2, v3, v5, v6, v7, v8
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
		goto _10
	case int32(b2_motorJoint):
		goto _11
	case int32(b2_mouseJoint):
		goto _12
	case int32(b2_filterJoint):
		goto _13
	case int32(b2_prismaticJoint):
		goto _14
	case int32(b2_revoluteJoint):
		goto _15
	case int32(b2_weldJoint):
		goto _16
	case int32(b2_wheelJoint):
		goto _17
	default:
		goto _18
	}
	goto _19
_10:
	;
	b2PrepareDistanceJoint(tls, joint, context)
	goto _19
_11:
	;
	b2PrepareMotorJoint(tls, joint, context)
	goto _19
_12:
	;
	b2PrepareMouseJoint(tls, joint, context)
	goto _19
_13:
	;
	goto _19
_14:
	;
	b2PreparePrismaticJoint(tls, joint, context)
	goto _19
_15:
	;
	b2PrepareRevoluteJoint(tls, joint, context)
	goto _19
_16:
	;
	b2PrepareWeldJoint(tls, joint, context)
	goto _19
_17:
	;
	b2PrepareWheelJoint(tls, joint, context)
	goto _19
_18:
	;
_22:
	;
	if bool(!(int32FromInt32(false1) != 0)) && b2InternalAssertFcn(tls, __ccgo_ts+4123, __ccgo_ts+9061, int32FromInt32(1388)) != 0 {
		__builtin_trap(tls)
	}
	goto _21
_21:
	;
	if 0 != 0 {
		goto _22
	}
	goto _20
_20:
	;
_19:
}

func b2WarmStartJoint(tls *_Stack, joint uintptr, context uintptr) {
	switch (*b2JointSim)(unsafe.Pointer(joint)).Type1 {
	case int32(b2_distanceJoint):
		goto _1
	case int32(b2_motorJoint):
		goto _2
	case int32(b2_mouseJoint):
		goto _3
	case int32(b2_filterJoint):
		goto _4
	case int32(b2_prismaticJoint):
		goto _5
	case int32(b2_revoluteJoint):
		goto _6
	case int32(b2_weldJoint):
		goto _7
	case int32(b2_wheelJoint):
		goto _8
	default:
		goto _9
	}
	goto _10
_1:
	;
	b2WarmStartDistanceJoint(tls, joint, context)
	goto _10
_2:
	;
	b2WarmStartMotorJoint(tls, joint, context)
	goto _10
_3:
	;
	b2WarmStartMouseJoint(tls, joint, context)
	goto _10
_4:
	;
	goto _10
_5:
	;
	b2WarmStartPrismaticJoint(tls, joint, context)
	goto _10
_6:
	;
	b2WarmStartRevoluteJoint(tls, joint, context)
	goto _10
_7:
	;
	b2WarmStartWeldJoint(tls, joint, context)
	goto _10
_8:
	;
	b2WarmStartWheelJoint(tls, joint, context)
	goto _10
_9:
	;
_13:
	;
	if bool(!(int32FromInt32(false1) != 0)) && b2InternalAssertFcn(tls, __ccgo_ts+4123, __ccgo_ts+9061, int32FromInt32(1428)) != 0 {
		__builtin_trap(tls)
	}
	goto _12
_12:
	;
	if 0 != 0 {
		goto _13
	}
	goto _11
_11:
	;
_10:
}

func b2SolveJoint(tls *_Stack, joint uintptr, context uintptr, useBias uint8) {
	switch (*b2JointSim)(unsafe.Pointer(joint)).Type1 {
	case int32(b2_distanceJoint):
		goto _1
	case int32(b2_motorJoint):
		goto _2
	case int32(b2_mouseJoint):
		goto _3
	case int32(b2_filterJoint):
		goto _4
	case int32(b2_prismaticJoint):
		goto _5
	case int32(b2_revoluteJoint):
		goto _6
	case int32(b2_weldJoint):
		goto _7
	case int32(b2_wheelJoint):
		goto _8
	default:
		goto _9
	}
	goto _10
_1:
	;
	b2SolveDistanceJoint(tls, joint, context, useBias)
	goto _10
_2:
	;
	b2SolveMotorJoint(tls, joint, context, useBias)
	goto _10
_3:
	;
	b2SolveMouseJoint(tls, joint, context)
	goto _10
_4:
	;
	goto _10
_5:
	;
	b2SolvePrismaticJoint(tls, joint, context, useBias)
	goto _10
_6:
	;
	b2SolveRevoluteJoint(tls, joint, context, useBias)
	goto _10
_7:
	;
	b2SolveWeldJoint(tls, joint, context, useBias)
	goto _10
_8:
	;
	b2SolveWheelJoint(tls, joint, context, useBias)
	goto _10
_9:
	;
_13:
	;
	if bool(!(int32FromInt32(false1) != 0)) && b2InternalAssertFcn(tls, __ccgo_ts+4123, __ccgo_ts+9061, int32FromInt32(1468)) != 0 {
		__builtin_trap(tls)
	}
	goto _12
_12:
	;
	if 0 != 0 {
		goto _13
	}
	goto _11
_11:
	;
_10:
}

func b2PrepareOverflowJoints(tls *_Stack, context uintptr) {
	var graph, joint, joints uintptr
	var i, jointCount int32
	_, _, _, _, _ = graph, i, joint, jointCount, joints
	graph = (*b2StepContext)(unsafe.Pointer(context)).Graph
	joints = (*(*b2GraphColor)(unsafe.Pointer(graph + uintptr(int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1))*56))).JointSims.Data
	jointCount = (*(*b2GraphColor)(unsafe.Pointer(graph + uintptr(int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1))*56))).JointSims.Count
	i = 0
	for {
		if !(i < jointCount) {
			break
		}
		joint = joints + uintptr(i)*196
		b2PrepareJoint(tls, joint, context)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b2WarmStartOverflowJoints(tls *_Stack, context uintptr) {
	var graph, joint, joints uintptr
	var i, jointCount int32
	_, _, _, _, _ = graph, i, joint, jointCount, joints
	graph = (*b2StepContext)(unsafe.Pointer(context)).Graph
	joints = (*(*b2GraphColor)(unsafe.Pointer(graph + uintptr(int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1))*56))).JointSims.Data
	jointCount = (*(*b2GraphColor)(unsafe.Pointer(graph + uintptr(int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1))*56))).JointSims.Count
	i = 0
	for {
		if !(i < jointCount) {
			break
		}
		joint = joints + uintptr(i)*196
		b2WarmStartJoint(tls, joint, context)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b2SolveOverflowJoints(tls *_Stack, context uintptr, useBias uint8) {
	var graph, joint, joints uintptr
	var i, jointCount int32
	_, _, _, _, _ = graph, i, joint, jointCount, joints
	graph = (*b2StepContext)(unsafe.Pointer(context)).Graph
	joints = (*(*b2GraphColor)(unsafe.Pointer(graph + uintptr(int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1))*56))).JointSims.Data
	jointCount = (*(*b2GraphColor)(unsafe.Pointer(graph + uintptr(int32FromInt32(_B2_GRAPH_COLOR_COUNT)-int32FromInt32(1))*56))).JointSims.Count
	i = 0
	for {
		if !(i < jointCount) {
			break
		}
		joint = joints + uintptr(i)*196
		b2SolveJoint(tls, joint, context, useBias)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b2DrawJoint(tls *_Stack, draw uintptr, world uintptr, joint uintptr) {
	var bodyA, bodyB, jointSim, v1, v3, v5, v7 uintptr
	var c1, c2, color b2HexColor1
	var colorIndex, v2, v6 int32
	var colors [12]b2HexColor1
	var p1, pA, pB, target, v10, v11, v14, v15, v17, v18, v20 Vec2
	var transformA, transformB, v13, v9 Transform
	var x, y, v19 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, c1, c2, color, colorIndex, colors, jointSim, p1, pA, pB, target, transformA, transformB, x, y, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v20, v3, v5, v6, v7, v9
	v1 = world + 1024
	v2 = (*(*b2JointEdge)(unsafe.Pointer(joint + 20))).BodyId
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	bodyA = v3
	v5 = world + 1024
	v6 = (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + 1*12))).BodyId
	if !(0 <= v6 && v6 < (*b2BodyArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodyArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*128
	goto _8
_8:
	bodyB = v7
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_disabledSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_disabledSet) {
		return
	}
	jointSim = b2GetJointSim(tls, world, joint)
	transformA = b2GetBodyTransformQuick(tls, world, bodyA)
	transformB = b2GetBodyTransformQuick(tls, world, bodyB)
	v9 = transformA
	v10 = (*b2JointSim)(unsafe.Pointer(jointSim)).LocalOriginAnchorA
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
	v14 = (*b2JointSim)(unsafe.Pointer(jointSim)).LocalOriginAnchorB
	x = float32(v13.Q.C*v14.X) - float32(v13.Q.S*v14.Y) + v13.P.X
	y = float32(v13.Q.S*v14.X) + float32(v13.Q.C*v14.Y) + v13.P.Y
	v15 = Vec2{
		X: x,
		Y: y,
	}
	goto _16
_16:
	pB = v15
	color = int32(b2_colorDarkSeaGreen)
	switch (*b2Joint)(unsafe.Pointer(joint)).Type1 {
	case int32(b2_distanceJoint):
		b2DrawDistanceJoint(tls, draw, jointSim, transformA, transformB)
	case int32(b2_mouseJoint):
		target = (*(*b2MouseJoint)(unsafe.Add(unsafe.Pointer(jointSim), 68))).TargetA
		c1 = int32(b2_colorGreen)
		(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, target, float32FromFloat32(4), c1, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pB, float32FromFloat32(4), c1, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		c2 = int32(b2_colorLightGray)
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, target, pB, c2, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	case int32(b2_filterJoint):
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, pA, pB, int32(b2_colorGold), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	case int32(b2_prismaticJoint):
		b2DrawPrismaticJoint(tls, draw, jointSim, transformA, transformB)
	case int32(b2_revoluteJoint):
		b2DrawRevoluteJoint(tls, draw, jointSim, transformA, transformB, (*b2Joint)(unsafe.Pointer(joint)).DrawSize)
	case int32(b2_wheelJoint):
		b2DrawWheelJoint(tls, draw, jointSim, transformA, transformB)
	default:
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, transformA.P, pA, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, pA, pB, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, transformB.P, pB, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	}
	if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawGraphColors != 0 {
		colors = [12]b2HexColor1{
			0:  int32(b2_colorRed),
			1:  int32(b2_colorOrange),
			2:  int32(b2_colorYellow),
			3:  int32(b2_colorGreen),
			4:  int32(b2_colorCyan),
			5:  int32(b2_colorBlue),
			6:  int32(b2_colorViolet),
			7:  int32(b2_colorPink),
			8:  int32(b2_colorChocolate),
			9:  int32(b2_colorGoldenRod),
			10: int32(b2_colorCoral),
		}
		colorIndex = (*b2Joint)(unsafe.Pointer(joint)).ColorIndex
		if colorIndex != -int32(1) {
			v17 = pA
			v18 = pB
			v19 = float32FromFloat32(0.5)
			v20 = Vec2{
				X: float32((float32FromFloat32(1)-v19)*v17.X) + float32(v19*v18.X),
				Y: float32((float32FromFloat32(1)-v19)*v17.Y) + float32(v19*v18.Y),
			}
			goto _21
		_21:
			p1 = v20
			(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, p1, float32FromFloat32(5), colors[colorIndex], (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		}
	}
}

/**@}*/

/**@}*/

func b2MotorJoint_SetLinearOffset(tls *_Stack, jointId JointId, linearOffset Vec2) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LinearOffset = linearOffset
}

func b2MotorJoint_GetLinearOffset(tls *_Stack, jointId JointId) (r Vec2) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	return (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LinearOffset
}

func b2MotorJoint_SetAngularOffset(tls *_Stack, jointId JointId, angularOffset float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).AngularOffset = angularOffset
}

func b2MotorJoint_GetAngularOffset(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	return (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).AngularOffset
}

func b2MotorJoint_SetMaxForce(tls *_Stack, jointId JointId, maxForce float32) {
	var joint uintptr
	var v1, v2, v3, v5 float32
	_, _, _, _, _ = joint, v1, v2, v3, v5
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	v1 = float32FromFloat32(0)
	v2 = maxForce
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MaxForce = v3
}

func b2MotorJoint_GetMaxForce(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	return (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MaxForce
}

func b2MotorJoint_SetMaxTorque(tls *_Stack, jointId JointId, maxTorque float32) {
	var joint uintptr
	var v1, v2, v3, v5 float32
	_, _, _, _, _ = joint, v1, v2, v3, v5
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	v1 = float32FromFloat32(0)
	v2 = maxTorque
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MaxTorque = v3
}

func b2MotorJoint_GetMaxTorque(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	return (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MaxTorque
}

func b2MotorJoint_SetCorrectionFactor(tls *_Stack, jointId JointId, correctionFactor float32) {
	var joint uintptr
	var v1, v2, v3, v4, v6, v7 float32
	_, _, _, _, _, _, _ = joint, v1, v2, v3, v4, v6, v7
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	v1 = correctionFactor
	v2 = float32FromFloat32(0)
	v3 = float32FromFloat32(1)
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
	(*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).CorrectionFactor = v4
}

func b2MotorJoint_GetCorrectionFactor(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_motorJoint))
	return (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).CorrectionFactor
}

func b2GetMotorJointForce(tls *_Stack, world uintptr, base uintptr) (r Vec2) {
	var force, v2, v3 Vec2
	var v1 float32
	_, _, _, _ = force, v1, v2, v3
	v1 = (*b2World)(unsafe.Pointer(world)).Inv_h
	v2 = (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(base), 68))).LinearImpulse
	v3 = Vec2{
		X: float32(v1 * v2.X),
		Y: float32(v1 * v2.Y),
	}
	goto _4
_4:
	force = v3
	return force
}

func b2GetMotorJointTorque(tls *_Stack, world uintptr, base uintptr) (r float32) {
	return float32((*b2World)(unsafe.Pointer(world)).Inv_h * (*(*b2MotorJoint)(unsafe.Add(unsafe.Pointer(base), 68))).AngularImpulse)
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
	var B, K, v55, v56 Mat22
	var a2, b2, c, c1, d, det, iA, iB, ka, mA, mB, s, v53, v58 float32
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var idA, idB, localIndexA, localIndexB, v10, v14, v18, v2, v22, v25, v26, v6 int32
	var rA, rB, v27, v28, v29, v32, v33, v35, v36, v37, v40, v41, v43, v44, v45, v47, v48, v49 Vec2
	var v31, v39, v51, v52 Rot
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = B, K, a2, b2, bodyA, bodyB, bodySimA, bodySimB, c, c1, d, det, iA, iB, idA, idB, joint, ka, localIndexA, localIndexB, mA, mB, rA, rB, s, setA, setB, world, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v28, v29, v3, v31, v32, v33, v35, v36, v37, v39, v40, v41, v43, v44, v45, v47, v48, v49, v5, v51, v52, v53, v55, v56, v58, v6, v7, v9
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_motorJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+10297, __ccgo_ts+10325, int32FromInt32(101)) != 0 {
		__builtin_trap(tls)
	}
	// chase body id to the solver set where the body lives
	idA = (*b2JointSim)(unsafe.Pointer(base)).BodyIdA
	idB = (*b2JointSim)(unsafe.Pointer(base)).BodyIdB
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	v1 = world + 1024
	v2 = idA
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	bodyA = v3
	v5 = world + 1024
	v6 = idB
	if !(0 <= v6 && v6 < (*b2BodyArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodyArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*128
	goto _8
_8:
	bodyB = v7
	if !((*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+4643, __ccgo_ts+10325, int32FromInt32(112)) != 0 {
		__builtin_trap(tls)
	}
	v9 = world + 1064
	v10 = (*b2Body)(unsafe.Pointer(bodyA)).SetIndex
	if !(0 <= v10 && v10 < (*b2SolverSetArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v11 = (*b2SolverSetArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*88
	goto _12
_12:
	setA = v11
	v13 = world + 1064
	v14 = (*b2Body)(unsafe.Pointer(bodyB)).SetIndex
	if !(0 <= v14 && v14 < (*b2SolverSetArray)(unsafe.Pointer(v13)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v15 = (*b2SolverSetArray)(unsafe.Pointer(v13)).Data + uintptr(v14)*88
	goto _16
_16:
	setB = v15
	localIndexA = (*b2Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b2Body)(unsafe.Pointer(bodyB)).LocalIndex
	v17 = setA
	v18 = localIndexA
	if !(0 <= v18 && v18 < (*b2BodySimArray)(unsafe.Pointer(v17)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	v19 = (*b2BodySimArray)(unsafe.Pointer(v17)).Data + uintptr(v18)*100
	goto _20
_20:
	bodySimA = v19
	v21 = setB
	v22 = localIndexB
	if !(0 <= v22 && v22 < (*b2BodySimArray)(unsafe.Pointer(v21)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	v23 = (*b2BodySimArray)(unsafe.Pointer(v21)).Data + uintptr(v22)*100
	goto _24
_24:
	bodySimB = v23
	mA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvMass
	iA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvInertia
	mB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvMass
	iB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvInertia
	(*b2JointSim)(unsafe.Pointer(base)).InvMassA = mA
	(*b2JointSim)(unsafe.Pointer(base)).InvMassB = mB
	(*b2JointSim)(unsafe.Pointer(base)).InvIA = iA
	(*b2JointSim)(unsafe.Pointer(base)).InvIB = iB
	joint = base + 68
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) {
		v25 = localIndexA
	} else {
		v25 = -int32(1)
	}
	(*b2MotorJoint)(unsafe.Pointer(joint)).IndexA = v25
	if (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) {
		v26 = localIndexB
	} else {
		v26 = -int32(1)
	}
	(*b2MotorJoint)(unsafe.Pointer(joint)).IndexB = v26
	v27 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorA
	v28 = (*b2BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v29 = Vec2{
		X: v27.X - v28.X,
		Y: v27.Y - v28.Y,
	}
	goto _30
_30:
	v31 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v32 = v29
	v33 = Vec2{
		X: float32(v31.C*v32.X) - float32(v31.S*v32.Y),
		Y: float32(v31.S*v32.X) + float32(v31.C*v32.Y),
	}
	goto _34
_34:
	(*b2MotorJoint)(unsafe.Pointer(joint)).AnchorA = v33
	v35 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorB
	v36 = (*b2BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v37 = Vec2{
		X: v35.X - v36.X,
		Y: v35.Y - v36.Y,
	}
	goto _38
_38:
	v39 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v40 = v37
	v41 = Vec2{
		X: float32(v39.C*v40.X) - float32(v39.S*v40.Y),
		Y: float32(v39.S*v40.X) + float32(v39.C*v40.Y),
	}
	goto _42
_42:
	(*b2MotorJoint)(unsafe.Pointer(joint)).AnchorB = v41
	v43 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Center
	v44 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Center
	v45 = Vec2{
		X: v43.X - v44.X,
		Y: v43.Y - v44.Y,
	}
	goto _46
_46:
	v47 = v45
	v48 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearOffset
	v49 = Vec2{
		X: v47.X - v48.X,
		Y: v47.Y - v48.Y,
	}
	goto _50
_50:
	(*b2MotorJoint)(unsafe.Pointer(joint)).DeltaCenter = v49
	v51 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v52 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	s = float32(v51.S*v52.C) - float32(v51.C*v52.S)
	c = float32(v51.C*v52.C) + float32(v51.S*v52.S)
	v53 = b2Atan2(tls, s, c)
	goto _54
_54:
	(*b2MotorJoint)(unsafe.Pointer(joint)).DeltaAngle = v53 - (*b2MotorJoint)(unsafe.Pointer(joint)).AngularOffset
	rA = (*b2MotorJoint)(unsafe.Pointer(joint)).AnchorA
	rB = (*b2MotorJoint)(unsafe.Pointer(joint)).AnchorB
	K.Cx.X = mA + mB + float32(float32(rA.Y*rA.Y)*iA) + float32(float32(rB.Y*rB.Y)*iB)
	K.Cx.Y = float32(float32(-rA.Y*rA.X)*iA) - float32(float32(rB.Y*rB.X)*iB)
	K.Cy.X = K.Cx.Y
	K.Cy.Y = mA + mB + float32(float32(rA.X*rA.X)*iA) + float32(float32(rB.X*rB.X)*iB)
	v55 = K
	a2 = v55.Cx.X
	b2 = v55.Cy.X
	c1 = v55.Cx.Y
	d = v55.Cy.Y
	det = float32(a2*d) - float32(b2*c1)
	if det != float32FromFloat32(0) {
		det = float32FromFloat32(1) / det
	}
	B = Mat22{
		Cx: Vec2{
			X: float32(det * d),
			Y: float32(-det * c1),
		},
		Cy: Vec2{
			X: float32(-det * b2),
			Y: float32(det * a2),
		},
	}
	v56 = B
	goto _57
_57:
	(*b2MotorJoint)(unsafe.Pointer(joint)).LinearMass = v56
	ka = iA + iB
	if ka > float32FromFloat32(0) {
		v58 = float32FromFloat32(1) / ka
	} else {
		v58 = float32FromFloat32(0)
	}
	(*b2MotorJoint)(unsafe.Pointer(joint)).AngularMass = v58
	if int32FromUint8((*b2StepContext)(unsafe.Pointer(context)).EnableWarmStarting) == false1 {
		(*b2MotorJoint)(unsafe.Pointer(joint)).LinearImpulse = b2Vec2_zero
		(*b2MotorJoint)(unsafe.Pointer(joint)).AngularImpulse = float32FromFloat32(0)
	}
}

func b2WarmStartMotorJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var bodyA, bodyB, joint, v1, v2 uintptr
	var iA, iB, mA, mB, v12, v18, v21, v27 float32
	var rA, rB, v11, v13, v14, v16, v17, v20, v22, v23, v25, v26, v4, v5, v8, v9 Vec2
	var v3, v7 Rot
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodyA, bodyB, iA, iB, joint, mA, mB, rA, rB, v1, v11, v12, v13, v14, v16, v17, v18, v2, v20, v21, v22, v23, v25, v26, v27, v3, v4, v5, v7, v8, v9
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	joint = base + 68
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Pointer(bp)) = b2_identityBodyState8
	if (*b2MotorJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v1 = bp
	} else {
		v1 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2MotorJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	bodyA = v1
	if (*b2MotorJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2MotorJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	bodyB = v2
	v3 = (*b2BodyState)(unsafe.Pointer(bodyA)).DeltaRotation
	v4 = (*b2MotorJoint)(unsafe.Pointer(joint)).AnchorA
	v5 = Vec2{
		X: float32(v3.C*v4.X) - float32(v3.S*v4.Y),
		Y: float32(v3.S*v4.X) + float32(v3.C*v4.Y),
	}
	goto _6
_6:
	rA = v5
	v7 = (*b2BodyState)(unsafe.Pointer(bodyB)).DeltaRotation
	v8 = (*b2MotorJoint)(unsafe.Pointer(joint)).AnchorB
	v9 = Vec2{
		X: float32(v7.C*v8.X) - float32(v7.S*v8.Y),
		Y: float32(v7.S*v8.X) + float32(v7.C*v8.Y),
	}
	goto _10
_10:
	rB = v9
	v11 = (*b2BodyState)(unsafe.Pointer(bodyA)).LinearVelocity
	v12 = mA
	v13 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearImpulse
	v14 = Vec2{
		X: v11.X - float32(v12*v13.X),
		Y: v11.Y - float32(v12*v13.Y),
	}
	goto _15
_15:
	(*b2BodyState)(unsafe.Pointer(bodyA)).LinearVelocity = v14
	v16 = rA
	v17 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearImpulse
	v18 = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	goto _19
_19:
	*(*float32)(unsafe.Pointer(bodyA + 8)) -= float32(iA * (v18 + (*b2MotorJoint)(unsafe.Pointer(joint)).AngularImpulse))
	v20 = (*b2BodyState)(unsafe.Pointer(bodyB)).LinearVelocity
	v21 = mB
	v22 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearImpulse
	v23 = Vec2{
		X: v20.X + float32(v21*v22.X),
		Y: v20.Y + float32(v21*v22.Y),
	}
	goto _24
_24:
	(*b2BodyState)(unsafe.Pointer(bodyB)).LinearVelocity = v23
	v25 = rB
	v26 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearImpulse
	v27 = float32(v25.X*v26.Y) - float32(v25.Y*v26.X)
	goto _28
_28:
	*(*float32)(unsafe.Pointer(bodyB + 8)) += float32(iB * (v27 + (*b2MotorJoint)(unsafe.Pointer(joint)).AngularImpulse))
}

func b2SolveMotorJoint(tls *_Stack, base uintptr, context uintptr, useBias uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var Cdot, angularBias, angularSeparation, c, iA, iB, impulse, invLength, length, mA, mB, maxImpulse, maxImpulse1, oldImpulse, s4, wA, wB, v10, v102, v11, v12, v14, v15, v40, v44, v5, v52, v7, v77, v87, v9, v93, v96 float32
	var Cdot1, b6, ds, impulse1, linearBias, linearSeparation, n, oldImpulse1, rA, rB, u, vA, vB, v100, v101, v17, v18, v21, v22, v24, v25, v26, v28, v29, v30, v32, v33, v34, v36, v37, v38, v41, v42, v45, v46, v48, v49, v50, v53, v54, v56, v57, v58, v60, v61, v62, v64, v65, v66, v69, v70, v72, v73, v74, v76, v79, v80, v82, v83, v84, v86, v88, v89, v91, v92, v95, v97, v98 Vec2
	var bodyA, bodyB, joint, v1, v2 uintptr
	var v16, v20, v3, v4 Rot
	var v68 Mat22
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = Cdot, Cdot1, angularBias, angularSeparation, b6, bodyA, bodyB, c, ds, iA, iB, impulse, impulse1, invLength, joint, length, linearBias, linearSeparation, mA, mB, maxImpulse, maxImpulse1, n, oldImpulse, oldImpulse1, rA, rB, s4, u, vA, vB, wA, wB, v1, v10, v100, v101, v102, v11, v12, v14, v15, v16, v17, v18, v2, v20, v21, v22, v24, v25, v26, v28, v29, v3, v30, v32, v33, v34, v36, v37, v38, v4, v40, v41, v42, v44, v45, v46, v48, v49, v5, v50, v52, v53, v54, v56, v57, v58, v60, v61, v62, v64, v65, v66, v68, v69, v7, v70, v72, v73, v74, v76, v77, v79, v80, v82, v83, v84, v86, v87, v88, v89, v9, v91, v92, v93, v95, v96, v97, v98
	_ = uint64FromInt64(4)
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_motorJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+10297, __ccgo_ts+10325, int32FromInt32(189)) != 0 {
		__builtin_trap(tls)
	}
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Pointer(bp)) = b2_identityBodyState8
	joint = base + 68
	if (*b2MotorJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v1 = bp
	} else {
		v1 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2MotorJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	bodyA = v1
	if (*b2MotorJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2MotorJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	bodyB = v2
	vA = (*b2BodyState)(unsafe.Pointer(bodyA)).LinearVelocity
	wA = (*b2BodyState)(unsafe.Pointer(bodyA)).AngularVelocity
	vB = (*b2BodyState)(unsafe.Pointer(bodyB)).LinearVelocity
	wB = (*b2BodyState)(unsafe.Pointer(bodyB)).AngularVelocity
	// angular constraint
	v3 = (*b2BodyState)(unsafe.Pointer(bodyB)).DeltaRotation
	v4 = (*b2BodyState)(unsafe.Pointer(bodyA)).DeltaRotation
	s4 = float32(v3.S*v4.C) - float32(v3.C*v4.S)
	c = float32(v3.C*v4.C) + float32(v3.S*v4.S)
	v5 = b2Atan2(tls, s4, c)
	goto _6
_6:
	angularSeparation = v5 + (*b2MotorJoint)(unsafe.Pointer(joint)).DeltaAngle
	v7 = __builtin_remainderf(tls, angularSeparation, float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)))
	goto _8
_8:
	angularSeparation = v7
	angularBias = float32(float32((*b2StepContext)(unsafe.Pointer(context)).Inv_h*(*b2MotorJoint)(unsafe.Pointer(joint)).CorrectionFactor) * angularSeparation)
	Cdot = wB - wA
	impulse = float32(-(*b2MotorJoint)(unsafe.Pointer(joint)).AngularMass * (Cdot + angularBias))
	oldImpulse = (*b2MotorJoint)(unsafe.Pointer(joint)).AngularImpulse
	maxImpulse = float32((*b2StepContext)(unsafe.Pointer(context)).H * (*b2MotorJoint)(unsafe.Pointer(joint)).MaxTorque)
	v9 = (*b2MotorJoint)(unsafe.Pointer(joint)).AngularImpulse + impulse
	v10 = -maxImpulse
	v11 = maxImpulse
	if v9 < v10 {
		v14 = v10
	} else {
		if v9 > v11 {
			v15 = v11
		} else {
			v15 = v9
		}
		v14 = v15
	}
	v12 = v14
	goto _13
_13:
	(*b2MotorJoint)(unsafe.Pointer(joint)).AngularImpulse = v12
	impulse = (*b2MotorJoint)(unsafe.Pointer(joint)).AngularImpulse - oldImpulse
	wA = wA - float32(iA*impulse)
	wB = wB + float32(iB*impulse)
	// linear constraint
	v16 = (*b2BodyState)(unsafe.Pointer(bodyA)).DeltaRotation
	v17 = (*b2MotorJoint)(unsafe.Pointer(joint)).AnchorA
	v18 = Vec2{
		X: float32(v16.C*v17.X) - float32(v16.S*v17.Y),
		Y: float32(v16.S*v17.X) + float32(v16.C*v17.Y),
	}
	goto _19
_19:
	rA = v18
	v20 = (*b2BodyState)(unsafe.Pointer(bodyB)).DeltaRotation
	v21 = (*b2MotorJoint)(unsafe.Pointer(joint)).AnchorB
	v22 = Vec2{
		X: float32(v20.C*v21.X) - float32(v20.S*v21.Y),
		Y: float32(v20.S*v21.X) + float32(v20.C*v21.Y),
	}
	goto _23
_23:
	rB = v22
	v24 = (*b2BodyState)(unsafe.Pointer(bodyB)).DeltaPosition
	v25 = (*b2BodyState)(unsafe.Pointer(bodyA)).DeltaPosition
	v26 = Vec2{
		X: v24.X - v25.X,
		Y: v24.Y - v25.Y,
	}
	goto _27
_27:
	v28 = rB
	v29 = rA
	v30 = Vec2{
		X: v28.X - v29.X,
		Y: v28.Y - v29.Y,
	}
	goto _31
_31:
	v32 = v26
	v33 = v30
	v34 = Vec2{
		X: v32.X + v33.X,
		Y: v32.Y + v33.Y,
	}
	goto _35
_35:
	ds = v34
	v36 = (*b2MotorJoint)(unsafe.Pointer(joint)).DeltaCenter
	v37 = ds
	v38 = Vec2{
		X: v36.X + v37.X,
		Y: v36.Y + v37.Y,
	}
	goto _39
_39:
	linearSeparation = v38
	v40 = float32((*b2StepContext)(unsafe.Pointer(context)).Inv_h * (*b2MotorJoint)(unsafe.Pointer(joint)).CorrectionFactor)
	v41 = linearSeparation
	v42 = Vec2{
		X: float32(v40 * v41.X),
		Y: float32(v40 * v41.Y),
	}
	goto _43
_43:
	linearBias = v42
	v44 = wB
	v45 = rB
	v46 = Vec2{
		X: float32(-v44 * v45.Y),
		Y: float32(v44 * v45.X),
	}
	goto _47
_47:
	v48 = vB
	v49 = v46
	v50 = Vec2{
		X: v48.X + v49.X,
		Y: v48.Y + v49.Y,
	}
	goto _51
_51:
	v52 = wA
	v53 = rA
	v54 = Vec2{
		X: float32(-v52 * v53.Y),
		Y: float32(v52 * v53.X),
	}
	goto _55
_55:
	v56 = vA
	v57 = v54
	v58 = Vec2{
		X: v56.X + v57.X,
		Y: v56.Y + v57.Y,
	}
	goto _59
_59:
	v60 = v50
	v61 = v58
	v62 = Vec2{
		X: v60.X - v61.X,
		Y: v60.Y - v61.Y,
	}
	goto _63
_63:
	Cdot1 = v62
	v64 = Cdot1
	v65 = linearBias
	v66 = Vec2{
		X: v64.X + v65.X,
		Y: v64.Y + v65.Y,
	}
	goto _67
_67:
	v68 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearMass
	v69 = v66
	u = Vec2{
		X: float32(v68.Cx.X*v69.X) + float32(v68.Cy.X*v69.Y),
		Y: float32(v68.Cx.Y*v69.X) + float32(v68.Cy.Y*v69.Y),
	}
	v70 = u
	goto _71
_71:
	b6 = v70
	impulse1 = Vec2{
		X: -b6.X,
		Y: -b6.Y,
	}
	oldImpulse1 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearImpulse
	maxImpulse1 = float32((*b2StepContext)(unsafe.Pointer(context)).H * (*b2MotorJoint)(unsafe.Pointer(joint)).MaxForce)
	v72 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearImpulse
	v73 = impulse1
	v74 = Vec2{
		X: v72.X + v73.X,
		Y: v72.Y + v73.Y,
	}
	goto _75
_75:
	(*b2MotorJoint)(unsafe.Pointer(joint)).LinearImpulse = v74
	v76 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearImpulse
	v77 = float32(v76.X*v76.X) + float32(v76.Y*v76.Y)
	goto _78
_78:
	if v77 > float32(maxImpulse1*maxImpulse1) {
		v79 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearImpulse
		length = sqrtf(tls, float32(v79.X*v79.X)+float32(v79.Y*v79.Y))
		if length < float32FromFloat32(1.1920928955078125e-07) {
			v80 = Vec2{}
			goto _81
		}
		invLength = float32FromFloat32(1) / length
		n = Vec2{
			X: float32(invLength * v79.X),
			Y: float32(invLength * v79.Y),
		}
		v80 = n
		goto _81
	_81:
		(*b2MotorJoint)(unsafe.Pointer(joint)).LinearImpulse = v80
		(*b2MotorJoint)(unsafe.Pointer(joint)).LinearImpulse.X *= maxImpulse1
		(*b2MotorJoint)(unsafe.Pointer(joint)).LinearImpulse.Y *= maxImpulse1
	}
	v82 = (*b2MotorJoint)(unsafe.Pointer(joint)).LinearImpulse
	v83 = oldImpulse1
	v84 = Vec2{
		X: v82.X - v83.X,
		Y: v82.Y - v83.Y,
	}
	goto _85
_85:
	impulse1 = v84
	v86 = vA
	v87 = mA
	v88 = impulse1
	v89 = Vec2{
		X: v86.X - float32(v87*v88.X),
		Y: v86.Y - float32(v87*v88.Y),
	}
	goto _90
_90:
	vA = v89
	v91 = rA
	v92 = impulse1
	v93 = float32(v91.X*v92.Y) - float32(v91.Y*v92.X)
	goto _94
_94:
	wA = wA - float32(iA*v93)
	v95 = vB
	v96 = mB
	v97 = impulse1
	v98 = Vec2{
		X: v95.X + float32(v96*v97.X),
		Y: v95.Y + float32(v96*v97.Y),
	}
	goto _99
_99:
	vB = v98
	v100 = rB
	v101 = impulse1
	v102 = float32(v100.X*v101.Y) - float32(v100.Y*v101.X)
	goto _103
_103:
	wB = wB + float32(iB*v102)
	(*b2BodyState)(unsafe.Pointer(bodyA)).LinearVelocity = vA
	(*b2BodyState)(unsafe.Pointer(bodyA)).AngularVelocity = wA
	(*b2BodyState)(unsafe.Pointer(bodyB)).LinearVelocity = vB
	(*b2BodyState)(unsafe.Pointer(bodyB)).AngularVelocity = wB
}

/**@}*/

/**@}*/

func b2MouseJoint_SetTarget(tls *_Stack, jointId JointId, target Vec2) {
	var base uintptr
	_ = base
	if !(b2IsValidVec2(tls, target) != 0) && b2InternalAssertFcn(tls, __ccgo_ts+10354, __ccgo_ts+10378, int32FromInt32(16)) != 0 {
		__builtin_trap(tls)
	}
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_mouseJoint))
	(*(*b2MouseJoint)(unsafe.Add(unsafe.Pointer(base), 68))).TargetA = target
}

func b2MouseJoint_GetTarget(tls *_Stack, jointId JointId) (r Vec2) {
	var base uintptr
	_ = base
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_mouseJoint))
	return (*(*b2MouseJoint)(unsafe.Add(unsafe.Pointer(base), 68))).TargetA
}

func b2MouseJoint_SetSpringHertz(tls *_Stack, jointId JointId, hertz float32) {
	var base uintptr
	_ = base
	if !(b2IsValidFloat(tls, hertz) != 0 && hertz >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+9805, __ccgo_ts+10378, int32FromInt32(29)) != 0 {
		__builtin_trap(tls)
	}
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_mouseJoint))
	(*(*b2MouseJoint)(unsafe.Add(unsafe.Pointer(base), 68))).Hertz = hertz
}

func b2MouseJoint_GetSpringHertz(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_mouseJoint))
	return (*(*b2MouseJoint)(unsafe.Add(unsafe.Pointer(base), 68))).Hertz
}

func b2MouseJoint_SetSpringDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	var base uintptr
	_ = base
	if !(b2IsValidFloat(tls, dampingRatio) != 0 && dampingRatio >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+9846, __ccgo_ts+10378, int32FromInt32(42)) != 0 {
		__builtin_trap(tls)
	}
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_mouseJoint))
	(*(*b2MouseJoint)(unsafe.Add(unsafe.Pointer(base), 68))).DampingRatio = dampingRatio
}

func b2MouseJoint_GetSpringDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_mouseJoint))
	return (*(*b2MouseJoint)(unsafe.Add(unsafe.Pointer(base), 68))).DampingRatio
}

func b2MouseJoint_SetMaxForce(tls *_Stack, jointId JointId, maxForce float32) {
	var base uintptr
	_ = base
	if !(b2IsValidFloat(tls, maxForce) != 0 && maxForce >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+10407, __ccgo_ts+10378, int32FromInt32(55)) != 0 {
		__builtin_trap(tls)
	}
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_mouseJoint))
	(*(*b2MouseJoint)(unsafe.Add(unsafe.Pointer(base), 68))).MaxForce = maxForce
}

func b2MouseJoint_GetMaxForce(tls *_Stack, jointId JointId) (r float32) {
	var base uintptr
	_ = base
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_mouseJoint))
	return (*(*b2MouseJoint)(unsafe.Add(unsafe.Pointer(base), 68))).MaxForce
}

func b2GetMouseJointForce(tls *_Stack, world uintptr, base uintptr) (r Vec2) {
	var force, v2, v3 Vec2
	var v1 float32
	_, _, _, _ = force, v1, v2, v3
	v1 = (*b2World)(unsafe.Pointer(world)).Inv_h
	v2 = (*(*b2MouseJoint)(unsafe.Add(unsafe.Pointer(base), 68))).LinearImpulse
	v3 = Vec2{
		X: float32(v1 * v2.X),
		Y: float32(v1 * v2.Y),
	}
	goto _4
_4:
	force = v3
	return force
}

func b2GetMouseJointTorque(tls *_Stack, world uintptr, base uintptr) (r float32) {
	return float32((*b2World)(unsafe.Pointer(world)).Inv_h * (*(*b2MouseJoint)(unsafe.Add(unsafe.Pointer(base), 68))).AngularImpulse)
}

func b2PrepareMouseJoint(tls *_Stack, base uintptr, context uintptr) {
	var B, K, v30, v31 Mat22
	var a1, a11, a21, a31, angularDampingRatio, angularHertz, b1, c, d, det, iB, mB, omega, v22, v23, v26, v27 float32
	var bodyB, bodySimB, joint, setB, world, v1, v11, v3, v5, v7, v9 uintptr
	var idB, localIndexB, v10, v13, v2, v6 int32
	var rB, v14, v15, v16, v19, v20, v33, v34, v35 Vec2
	var v18 Rot
	var v24, v28 b2Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = B, K, a1, a11, a21, a31, angularDampingRatio, angularHertz, b1, bodyB, bodySimB, c, d, det, iB, idB, joint, localIndexB, mB, omega, rB, setB, world, v1, v10, v11, v13, v14, v15, v16, v18, v19, v2, v20, v22, v23, v24, v26, v27, v28, v3, v30, v31, v33, v34, v35, v5, v6, v7, v9
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_mouseJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+10454, __ccgo_ts+10378, int32FromInt32(79)) != 0 {
		__builtin_trap(tls)
	}
	// chase body id to the solver set where the body lives
	idB = (*b2JointSim)(unsafe.Pointer(base)).BodyIdB
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	v1 = world + 1024
	v2 = idB
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	bodyB = v3
	if !((*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+3007, __ccgo_ts+10378, int32FromInt32(88)) != 0 {
		__builtin_trap(tls)
	}
	v5 = world + 1064
	v6 = (*b2Body)(unsafe.Pointer(bodyB)).SetIndex
	if !(0 <= v6 && v6 < (*b2SolverSetArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2SolverSetArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*88
	goto _8
_8:
	setB = v7
	localIndexB = (*b2Body)(unsafe.Pointer(bodyB)).LocalIndex
	v9 = setB
	v10 = localIndexB
	if !(0 <= v10 && v10 < (*b2BodySimArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	v11 = (*b2BodySimArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*100
	goto _12
_12:
	bodySimB = v11
	(*b2JointSim)(unsafe.Pointer(base)).InvMassB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvMass
	(*b2JointSim)(unsafe.Pointer(base)).InvIB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvInertia
	joint = base + 68
	if (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) {
		v13 = localIndexB
	} else {
		v13 = -int32(1)
	}
	(*b2MouseJoint)(unsafe.Pointer(joint)).IndexB = v13
	v14 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorB
	v15 = (*b2BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v16 = Vec2{
		X: v14.X - v15.X,
		Y: v14.Y - v15.Y,
	}
	goto _17
_17:
	v18 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v19 = v16
	v20 = Vec2{
		X: float32(v18.C*v19.X) - float32(v18.S*v19.Y),
		Y: float32(v18.S*v19.X) + float32(v18.C*v19.Y),
	}
	goto _21
_21:
	(*b2MouseJoint)(unsafe.Pointer(joint)).AnchorB = v20
	v22 = (*b2MouseJoint)(unsafe.Pointer(joint)).Hertz
	v23 = (*b2StepContext)(unsafe.Pointer(context)).H
	if v22 == float32FromFloat32(0) {
		v24 = b2Softness{}
		goto _25
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v22)
	a11 = float32(float32FromFloat32(2)*(*b2MouseJoint)(unsafe.Pointer(joint)).DampingRatio) + float32(v23*omega)
	a21 = float32(float32(v23*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v24 = b2Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _25
_25:
	(*b2MouseJoint)(unsafe.Pointer(joint)).LinearSoftness = v24
	angularHertz = float32FromFloat32(0.5)
	angularDampingRatio = float32FromFloat32(0.1)
	v26 = angularHertz
	v27 = (*b2StepContext)(unsafe.Pointer(context)).H
	if v26 == float32FromFloat32(0) {
		v28 = b2Softness{}
		goto _29
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v26)
	a11 = float32(float32FromFloat32(2)*angularDampingRatio) + float32(v27*omega)
	a21 = float32(float32(v27*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v28 = b2Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _29
_29:
	(*b2MouseJoint)(unsafe.Pointer(joint)).AngularSoftness = v28
	rB = (*b2MouseJoint)(unsafe.Pointer(joint)).AnchorB
	mB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvMass
	iB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvInertia
	K.Cx.X = mB + float32(float32(iB*rB.Y)*rB.Y)
	K.Cx.Y = float32(float32(-iB*rB.X) * rB.Y)
	K.Cy.X = K.Cx.Y
	K.Cy.Y = mB + float32(float32(iB*rB.X)*rB.X)
	v30 = K
	a1 = v30.Cx.X
	b1 = v30.Cy.X
	c = v30.Cx.Y
	d = v30.Cy.Y
	det = float32(a1*d) - float32(b1*c)
	if det != float32FromFloat32(0) {
		det = float32FromFloat32(1) / det
	}
	B = Mat22{
		Cx: Vec2{
			X: float32(det * d),
			Y: float32(-det * c),
		},
		Cy: Vec2{
			X: float32(-det * b1),
			Y: float32(det * a1),
		},
	}
	v31 = B
	goto _32
_32:
	(*b2MouseJoint)(unsafe.Pointer(joint)).LinearMass = v31
	v33 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Center
	v34 = (*b2MouseJoint)(unsafe.Pointer(joint)).TargetA
	v35 = Vec2{
		X: v33.X - v34.X,
		Y: v33.Y - v34.Y,
	}
	goto _36
_36:
	(*b2MouseJoint)(unsafe.Pointer(joint)).DeltaCenter = v35
	if int32FromUint8((*b2StepContext)(unsafe.Pointer(context)).EnableWarmStarting) == false1 {
		(*b2MouseJoint)(unsafe.Pointer(joint)).LinearImpulse = b2Vec2_zero
		(*b2MouseJoint)(unsafe.Pointer(joint)).AngularImpulse = float32FromFloat32(0)
	}
}

func b2WarmStartMouseJoint(tls *_Stack, base uintptr, context uintptr) {
	var dqB, v1 Rot
	var iB, mB, wB, v12, v6 float32
	var joint, stateB uintptr
	var rB, vB, v10, v11, v2, v3, v5, v7, v8 Vec2
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = dqB, iB, joint, mB, rB, stateB, vB, wB, v1, v10, v11, v12, v2, v3, v5, v6, v7, v8
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_mouseJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+10454, __ccgo_ts+10378, int32FromInt32(132)) != 0 {
		__builtin_trap(tls)
	}
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	joint = base + 68
	stateB = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2MouseJoint)(unsafe.Pointer(joint)).IndexB)*32
	vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	dqB = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v1 = dqB
	v2 = (*b2MouseJoint)(unsafe.Pointer(joint)).AnchorB
	v3 = Vec2{
		X: float32(v1.C*v2.X) - float32(v1.S*v2.Y),
		Y: float32(v1.S*v2.X) + float32(v1.C*v2.Y),
	}
	goto _4
_4:
	rB = v3
	v5 = vB
	v6 = mB
	v7 = (*b2MouseJoint)(unsafe.Pointer(joint)).LinearImpulse
	v8 = Vec2{
		X: v5.X + float32(v6*v7.X),
		Y: v5.Y + float32(v6*v7.Y),
	}
	goto _9
_9:
	vB = v8
	v10 = rB
	v11 = (*b2MouseJoint)(unsafe.Pointer(joint)).LinearImpulse
	v12 = float32(v10.X*v11.Y) - float32(v10.Y*v11.X)
	goto _13
_13:
	wB = wB + float32(iB*(v12+(*b2MouseJoint)(unsafe.Pointer(joint)).AngularImpulse))
	(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
	(*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
}

func b2SolveMouseJoint(tls *_Stack, base uintptr, context uintptr) {
	var Cdot, b3, bias, impulse1, n, oldImpulse, rB, separation, u, vB, v10, v11, v12, v14, v15, v16, v18, v19, v20, v23, v24, v26, v27, v28, v3, v31, v32, v34, v37, v38, v4, v41, v42, v44, v46, v47, v49, v50, v7, v8 Vec2
	var dqB, v2 Rot
	var iB, impulse, impulseScale, impulseScale1, invLength, length, mB, mag, massScale, massScale1, maxImpulse, wB, v1, v22, v35, v40, v45, v51, v6 float32
	var joint, stateB uintptr
	var v30 Mat22
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = Cdot, b3, bias, dqB, iB, impulse, impulse1, impulseScale, impulseScale1, invLength, joint, length, mB, mag, massScale, massScale1, maxImpulse, n, oldImpulse, rB, separation, stateB, u, vB, wB, v1, v10, v11, v12, v14, v15, v16, v18, v19, v2, v20, v22, v23, v24, v26, v27, v28, v3, v30, v31, v32, v34, v35, v37, v38, v4, v40, v41, v42, v44, v45, v46, v47, v49, v50, v51, v6, v7, v8
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	joint = base + 68
	stateB = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2MouseJoint)(unsafe.Pointer(joint)).IndexB)*32
	vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	// Softness with no bias to reduce rotation speed
	massScale = (*b2MouseJoint)(unsafe.Pointer(joint)).AngularSoftness.MassScale
	impulseScale = (*b2MouseJoint)(unsafe.Pointer(joint)).AngularSoftness.ImpulseScale
	if iB > float32FromFloat32(0) {
		v1 = -wB / iB
	} else {
		v1 = float32FromFloat32(0)
	}
	impulse = v1
	impulse = float32(massScale*impulse) - float32(impulseScale*(*b2MouseJoint)(unsafe.Pointer(joint)).AngularImpulse)
	*(*float32)(unsafe.Pointer(joint + 28)) += impulse
	wB = wB + float32(iB*impulse)
	maxImpulse = float32((*b2MouseJoint)(unsafe.Pointer(joint)).MaxForce * (*b2StepContext)(unsafe.Pointer(context)).H)
	dqB = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v2 = dqB
	v3 = (*b2MouseJoint)(unsafe.Pointer(joint)).AnchorB
	v4 = Vec2{
		X: float32(v2.C*v3.X) - float32(v2.S*v3.Y),
		Y: float32(v2.S*v3.X) + float32(v2.C*v3.Y),
	}
	goto _5
_5:
	rB = v4
	v6 = wB
	v7 = rB
	v8 = Vec2{
		X: float32(-v6 * v7.Y),
		Y: float32(v6 * v7.X),
	}
	goto _9
_9:
	v10 = vB
	v11 = v8
	v12 = Vec2{
		X: v10.X + v11.X,
		Y: v10.Y + v11.Y,
	}
	goto _13
_13:
	Cdot = v12
	v14 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
	v15 = rB
	v16 = Vec2{
		X: v14.X + v15.X,
		Y: v14.Y + v15.Y,
	}
	goto _17
_17:
	v18 = v16
	v19 = (*b2MouseJoint)(unsafe.Pointer(joint)).DeltaCenter
	v20 = Vec2{
		X: v18.X + v19.X,
		Y: v18.Y + v19.Y,
	}
	goto _21
_21:
	separation = v20
	v22 = (*b2MouseJoint)(unsafe.Pointer(joint)).LinearSoftness.BiasRate
	v23 = separation
	v24 = Vec2{
		X: float32(v22 * v23.X),
		Y: float32(v22 * v23.Y),
	}
	goto _25
_25:
	bias = v24
	massScale1 = (*b2MouseJoint)(unsafe.Pointer(joint)).LinearSoftness.MassScale
	impulseScale1 = (*b2MouseJoint)(unsafe.Pointer(joint)).LinearSoftness.ImpulseScale
	v26 = Cdot
	v27 = bias
	v28 = Vec2{
		X: v26.X + v27.X,
		Y: v26.Y + v27.Y,
	}
	goto _29
_29:
	v30 = (*b2MouseJoint)(unsafe.Pointer(joint)).LinearMass
	v31 = v28
	u = Vec2{
		X: float32(v30.Cx.X*v31.X) + float32(v30.Cy.X*v31.Y),
		Y: float32(v30.Cx.Y*v31.X) + float32(v30.Cy.Y*v31.Y),
	}
	v32 = u
	goto _33
_33:
	b3 = v32
	impulse1.X = float32(-massScale1*b3.X) - float32(impulseScale1*(*b2MouseJoint)(unsafe.Pointer(joint)).LinearImpulse.X)
	impulse1.Y = float32(-massScale1*b3.Y) - float32(impulseScale1*(*b2MouseJoint)(unsafe.Pointer(joint)).LinearImpulse.Y)
	oldImpulse = (*b2MouseJoint)(unsafe.Pointer(joint)).LinearImpulse
	(*b2MouseJoint)(unsafe.Pointer(joint)).LinearImpulse.X += impulse1.X
	(*b2MouseJoint)(unsafe.Pointer(joint)).LinearImpulse.Y += impulse1.Y
	v34 = (*b2MouseJoint)(unsafe.Pointer(joint)).LinearImpulse
	v35 = sqrtf(tls, float32(v34.X*v34.X)+float32(v34.Y*v34.Y))
	goto _36
_36:
	mag = v35
	if mag > maxImpulse {
		v37 = (*b2MouseJoint)(unsafe.Pointer(joint)).LinearImpulse
		length = sqrtf(tls, float32(v37.X*v37.X)+float32(v37.Y*v37.Y))
		if length < float32FromFloat32(1.1920928955078125e-07) {
			v38 = Vec2{}
			goto _39
		}
		invLength = float32FromFloat32(1) / length
		n = Vec2{
			X: float32(invLength * v37.X),
			Y: float32(invLength * v37.Y),
		}
		v38 = n
		goto _39
	_39:
		v40 = maxImpulse
		v41 = v38
		v42 = Vec2{
			X: float32(v40 * v41.X),
			Y: float32(v40 * v41.Y),
		}
		goto _43
	_43:
		(*b2MouseJoint)(unsafe.Pointer(joint)).LinearImpulse = v42
	}
	impulse1.X = (*b2MouseJoint)(unsafe.Pointer(joint)).LinearImpulse.X - oldImpulse.X
	impulse1.Y = (*b2MouseJoint)(unsafe.Pointer(joint)).LinearImpulse.Y - oldImpulse.Y
	v44 = vB
	v45 = mB
	v46 = impulse1
	v47 = Vec2{
		X: v44.X + float32(v45*v46.X),
		Y: v44.Y + float32(v45*v46.Y),
	}
	goto _48
_48:
	vB = v47
	v49 = rB
	v50 = impulse1
	v51 = float32(v49.X*v50.Y) - float32(v49.Y*v50.X)
	goto _52
_52:
	wB = wB + float32(iB*v51)
	(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
	(*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
}

func b2PrismaticJoint_EnableSpring(tls *_Stack, jointId JointId, enableSpring uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	if enableSpring != (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableSpring {
		(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableSpring = enableSpring
		(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).SpringImpulse = float32FromFloat32(0)
	}
}

func b2PrismaticJoint_IsSpringEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableSpring
}

func b2PrismaticJoint_SetSpringHertz(tls *_Stack, jointId JointId, hertz float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).Hertz = hertz
}

func b2PrismaticJoint_GetSpringHertz(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).Hertz
}

func b2PrismaticJoint_SetSpringDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).DampingRatio = dampingRatio
}

func b2PrismaticJoint_GetSpringDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).DampingRatio
}

func b2PrismaticJoint_SetTargetTranslation(tls *_Stack, jointId JointId, translation float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).TargetTranslation = translation
}

func b2PrismaticJoint_GetTargetTranslation(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).TargetTranslation
}

func b2PrismaticJoint_EnableLimit(tls *_Stack, jointId JointId, enableLimit uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	if enableLimit != (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableLimit {
		(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableLimit = enableLimit
		(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LowerImpulse = float32FromFloat32(0)
		(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).UpperImpulse = float32FromFloat32(0)
	}
}

func b2PrismaticJoint_IsLimitEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableLimit
}

func b2PrismaticJoint_GetLowerLimit(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LowerTranslation
}

func b2PrismaticJoint_GetUpperLimit(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).UpperTranslation
}

func b2PrismaticJoint_SetLimits(tls *_Stack, jointId JointId, lower float32, upper float32) {
	var joint uintptr
	var v1, v10, v2, v3, v5, v6, v7, v8 float32
	_, _, _, _, _, _, _, _, _ = joint, v1, v10, v2, v3, v5, v6, v7, v8
	if !(lower <= upper) && b2InternalAssertFcn(tls, __ccgo_ts+10482, __ccgo_ts+10497, int32FromInt32(99)) != 0 {
		__builtin_trap(tls)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	if lower != (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LowerTranslation || upper != (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).UpperTranslation {
		v1 = lower
		v2 = upper
		if v1 < v2 {
			v5 = v1
		} else {
			v5 = v2
		}
		v3 = v5
		goto _4
	_4:
		(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LowerTranslation = v3
		v6 = lower
		v7 = upper
		if v6 > v7 {
			v10 = v6
		} else {
			v10 = v7
		}
		v8 = v10
		goto _9
	_9:
		(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).UpperTranslation = v8
		(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LowerImpulse = float32FromFloat32(0)
		(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).UpperImpulse = float32FromFloat32(0)
	}
}

func b2PrismaticJoint_EnableMotor(tls *_Stack, jointId JointId, enableMotor uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	if enableMotor != (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableMotor {
		(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableMotor = enableMotor
		(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MotorImpulse = float32FromFloat32(0)
	}
}

func b2PrismaticJoint_IsMotorEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableMotor
}

func b2PrismaticJoint_SetMotorSpeed(tls *_Stack, jointId JointId, motorSpeed float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MotorSpeed = motorSpeed
}

func b2PrismaticJoint_GetMotorSpeed(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MotorSpeed
}

func b2PrismaticJoint_GetMotorForce(tls *_Stack, jointId JointId) (r float32) {
	var base, world uintptr
	_, _ = base, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	base = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return float32((*b2World)(unsafe.Pointer(world)).Inv_h * (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 68))).MotorImpulse)
}

func b2PrismaticJoint_SetMaxMotorForce(tls *_Stack, jointId JointId, force float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	(*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MaxMotorForce = force
}

func b2PrismaticJoint_GetMaxMotorForce(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	return (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MaxMotorForce
}

func b2PrismaticJoint_GetTranslation(tls *_Stack, jointId JointId) (r float32) {
	var axisA, d, pA, pB, v10, v11, v13, v14, v15, v17, v18, v2, v3, v6, v7 Vec2
	var joint, jointSim, world uintptr
	var transformA, transformB, v5, v9 Transform
	var translation, x, y, v19 float32
	var v1 Rot
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axisA, d, joint, jointSim, pA, pB, transformA, transformB, translation, world, x, y, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v3, v5, v6, v7, v9
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	jointSim = b2GetJointSimCheckType(tls, jointId, int32(b2_prismaticJoint))
	transformA = b2GetBodyTransform(tls, world, (*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdA)
	transformB = b2GetBodyTransform(tls, world, (*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdB)
	joint = jointSim + 68
	v1 = transformA.Q
	v2 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).LocalAxisA
	v3 = Vec2{
		X: float32(v1.C*v2.X) - float32(v1.S*v2.Y),
		Y: float32(v1.S*v2.X) + float32(v1.C*v2.Y),
	}
	goto _4
_4:
	axisA = v3
	v5 = transformA
	v6 = (*b2JointSim)(unsafe.Pointer(jointSim)).LocalOriginAnchorA
	x = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y) + v5.P.X
	y = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y) + v5.P.Y
	v7 = Vec2{
		X: x,
		Y: y,
	}
	goto _8
_8:
	pA = v7
	v9 = transformB
	v10 = (*b2JointSim)(unsafe.Pointer(jointSim)).LocalOriginAnchorB
	x = float32(v9.Q.C*v10.X) - float32(v9.Q.S*v10.Y) + v9.P.X
	y = float32(v9.Q.S*v10.X) + float32(v9.Q.C*v10.Y) + v9.P.Y
	v11 = Vec2{
		X: x,
		Y: y,
	}
	goto _12
_12:
	pB = v11
	v13 = pB
	v14 = pA
	v15 = Vec2{
		X: v13.X - v14.X,
		Y: v13.Y - v14.Y,
	}
	goto _16
_16:
	d = v15
	v17 = d
	v18 = axisA
	v19 = float32(v17.X*v18.X) + float32(v17.Y*v18.Y)
	goto _20
_20:
	translation = v19
	return translation
}

func b2PrismaticJoint_GetSpeed(tls *_Stack, jointId JointId) (r float32) {
	var axisA, cA, cB, d, rA, rB, vA, vB, vRel, v10, v11, v13, v14, v15, v18, v19, v21, v22, v23, v26, v27, v29, v30, v31, v33, v34, v35, v37, v38, v39, v41, v42, v46, v47, v49, v50, v51, v54, v55, v57, v58, v59, v61, v62, v63, v66, v67, v69, v70, v73, v74 Vec2
	var bodyA, bodyB, bodySimA, bodySimB, bodyStateA, bodyStateB, joint, jointSim, prismatic, world, v1, v3, v5, v7 uintptr
	var speed, wA, wB, v43, v44, v45, v53, v65, v71, v75 float32
	var transformA, transformB Transform
	var v17, v25, v9 Rot
	var v2, v6 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axisA, bodyA, bodyB, bodySimA, bodySimB, bodyStateA, bodyStateB, cA, cB, d, joint, jointSim, prismatic, rA, rB, speed, transformA, transformB, vA, vB, vRel, wA, wB, world, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v29, v3, v30, v31, v33, v34, v35, v37, v38, v39, v41, v42, v43, v44, v45, v46, v47, v49, v5, v50, v51, v53, v54, v55, v57, v58, v59, v6, v61, v62, v63, v65, v66, v67, v69, v7, v70, v71, v73, v74, v75, v9
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointFullId(tls, world, jointId)
	if !((*b2Joint)(unsafe.Pointer(joint)).Type1 == int32(b2_prismaticJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+10530, __ccgo_ts+10497, int32FromInt32(178)) != 0 {
		__builtin_trap(tls)
	}
	jointSim = b2GetJointSim(tls, world, joint)
	if !((*b2JointSim)(unsafe.Pointer(jointSim)).Type1 == int32(b2_prismaticJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+10563, __ccgo_ts+10497, int32FromInt32(180)) != 0 {
		__builtin_trap(tls)
	}
	v1 = world + 1024
	v2 = (*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdA
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	bodyA = v3
	v5 = world + 1024
	v6 = (*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdB
	if !(0 <= v6 && v6 < (*b2BodyArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodyArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*128
	goto _8
_8:
	bodyB = v7
	bodySimA = b2GetBodySim(tls, world, bodyA)
	bodySimB = b2GetBodySim(tls, world, bodyB)
	bodyStateA = b2GetBodyState(tls, world, bodyA)
	bodyStateB = b2GetBodyState(tls, world, bodyB)
	transformA = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform
	transformB = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform
	prismatic = jointSim + 68
	v9 = transformA.Q
	v10 = (*b2PrismaticJoint)(unsafe.Pointer(prismatic)).LocalAxisA
	v11 = Vec2{
		X: float32(v9.C*v10.X) - float32(v9.S*v10.Y),
		Y: float32(v9.S*v10.X) + float32(v9.C*v10.Y),
	}
	goto _12
_12:
	axisA = v11
	cA = (*b2BodySim)(unsafe.Pointer(bodySimA)).Center
	cB = (*b2BodySim)(unsafe.Pointer(bodySimB)).Center
	v13 = (*b2JointSim)(unsafe.Pointer(jointSim)).LocalOriginAnchorA
	v14 = (*b2BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v15 = Vec2{
		X: v13.X - v14.X,
		Y: v13.Y - v14.Y,
	}
	goto _16
_16:
	v17 = transformA.Q
	v18 = v15
	v19 = Vec2{
		X: float32(v17.C*v18.X) - float32(v17.S*v18.Y),
		Y: float32(v17.S*v18.X) + float32(v17.C*v18.Y),
	}
	goto _20
_20:
	rA = v19
	v21 = (*b2JointSim)(unsafe.Pointer(jointSim)).LocalOriginAnchorB
	v22 = (*b2BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v23 = Vec2{
		X: v21.X - v22.X,
		Y: v21.Y - v22.Y,
	}
	goto _24
_24:
	v25 = transformB.Q
	v26 = v23
	v27 = Vec2{
		X: float32(v25.C*v26.X) - float32(v25.S*v26.Y),
		Y: float32(v25.S*v26.X) + float32(v25.C*v26.Y),
	}
	goto _28
_28:
	rB = v27
	v29 = cB
	v30 = cA
	v31 = Vec2{
		X: v29.X - v30.X,
		Y: v29.Y - v30.Y,
	}
	goto _32
_32:
	v33 = rB
	v34 = rA
	v35 = Vec2{
		X: v33.X - v34.X,
		Y: v33.Y - v34.Y,
	}
	goto _36
_36:
	v37 = v31
	v38 = v35
	v39 = Vec2{
		X: v37.X + v38.X,
		Y: v37.Y + v38.Y,
	}
	goto _40
_40:
	d = v39
	if bodyStateA != 0 {
		v41 = (*b2BodyState)(unsafe.Pointer(bodyStateA)).LinearVelocity
	} else {
		v41 = b2Vec2_zero
	}
	vA = v41
	if bodyStateB != 0 {
		v42 = (*b2BodyState)(unsafe.Pointer(bodyStateB)).LinearVelocity
	} else {
		v42 = b2Vec2_zero
	}
	vB = v42
	if bodyStateA != 0 {
		v43 = (*b2BodyState)(unsafe.Pointer(bodyStateA)).AngularVelocity
	} else {
		v43 = float32FromFloat32(0)
	}
	wA = v43
	if bodyStateB != 0 {
		v44 = (*b2BodyState)(unsafe.Pointer(bodyStateB)).AngularVelocity
	} else {
		v44 = float32FromFloat32(0)
	}
	wB = v44
	v45 = wB
	v46 = rB
	v47 = Vec2{
		X: float32(-v45 * v46.Y),
		Y: float32(v45 * v46.X),
	}
	goto _48
_48:
	v49 = vB
	v50 = v47
	v51 = Vec2{
		X: v49.X + v50.X,
		Y: v49.Y + v50.Y,
	}
	goto _52
_52:
	v53 = wA
	v54 = rA
	v55 = Vec2{
		X: float32(-v53 * v54.Y),
		Y: float32(v53 * v54.X),
	}
	goto _56
_56:
	v57 = vA
	v58 = v55
	v59 = Vec2{
		X: v57.X + v58.X,
		Y: v57.Y + v58.Y,
	}
	goto _60
_60:
	v61 = v51
	v62 = v59
	v63 = Vec2{
		X: v61.X - v62.X,
		Y: v61.Y - v62.Y,
	}
	goto _64
_64:
	vRel = v63
	v65 = wA
	v66 = axisA
	v67 = Vec2{
		X: float32(-v65 * v66.Y),
		Y: float32(v65 * v66.X),
	}
	goto _68
_68:
	v69 = d
	v70 = v67
	v71 = float32(v69.X*v70.X) + float32(v69.Y*v70.Y)
	goto _72
_72:
	v73 = axisA
	v74 = vRel
	v75 = float32(v73.X*v74.X) + float32(v73.Y*v74.Y)
	goto _76
_76:
	speed = v71 + v75
	return speed
}

func b2GetPrismaticJointForce(tls *_Stack, world uintptr, base uintptr) (r Vec2) {
	var axialForce, inv_h, perpForce, v12, v8 float32
	var axisA, force, perpA, v10, v13, v14, v16, v17, v18, v2, v3, v5, v6, v9 Vec2
	var idA int32
	var joint uintptr
	var transformA Transform
	var v1 Rot
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axialForce, axisA, force, idA, inv_h, joint, perpA, perpForce, transformA, v1, v10, v12, v13, v14, v16, v17, v18, v2, v3, v5, v6, v8, v9
	idA = (*b2JointSim)(unsafe.Pointer(base)).BodyIdA
	transformA = b2GetBodyTransform(tls, world, idA)
	joint = base + 68
	v1 = transformA.Q
	v2 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).LocalAxisA
	v3 = Vec2{
		X: float32(v1.C*v2.X) - float32(v1.S*v2.Y),
		Y: float32(v1.S*v2.X) + float32(v1.C*v2.Y),
	}
	goto _4
_4:
	axisA = v3
	v5 = axisA
	v6 = Vec2{
		X: -v5.Y,
		Y: v5.X,
	}
	goto _7
_7:
	perpA = v6
	inv_h = (*b2World)(unsafe.Pointer(world)).Inv_h
	perpForce = float32(inv_h * (*b2PrismaticJoint)(unsafe.Pointer(joint)).Impulse.X)
	axialForce = float32(inv_h * ((*b2PrismaticJoint)(unsafe.Pointer(joint)).MotorImpulse + (*b2PrismaticJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b2PrismaticJoint)(unsafe.Pointer(joint)).UpperImpulse))
	v8 = perpForce
	v9 = perpA
	v10 = Vec2{
		X: float32(v8 * v9.X),
		Y: float32(v8 * v9.Y),
	}
	goto _11
_11:
	v12 = axialForce
	v13 = axisA
	v14 = Vec2{
		X: float32(v12 * v13.X),
		Y: float32(v12 * v13.Y),
	}
	goto _15
_15:
	v16 = v10
	v17 = v14
	v18 = Vec2{
		X: v16.X + v17.X,
		Y: v16.Y + v17.Y,
	}
	goto _19
_19:
	force = v18
	return force
}

func b2GetPrismaticJointTorque(tls *_Stack, world uintptr, base uintptr) (r float32) {
	return float32((*b2World)(unsafe.Pointer(world)).Inv_h * (*(*b2PrismaticJoint)(unsafe.Add(unsafe.Pointer(base), 68))).Impulse.Y)
}

// Linear constraint (point-to-line)
// d = p2 - p1 = x2 + r2 - x1 - r1
// C = dot(perp, d)
// Cdot = dot(d, cross(w1, perp)) + dot(perp, v2 + cross(w2, r2) - v1 - cross(w1, r1))
//      = -dot(perp, v1) - dot(cross(d + r1, perp), w1) + dot(perp, v2) + dot(cross(r2, perp), v2)
// J = [-perp, -cross(d + r1, perp), perp, cross(r2,perp)]
//
// Angular constraint
// C = a2 - a1 + a_initial
// Cdot = w2 - w1
// J = [0 0 -1 0 0 1]
//
// K = J * invM * JT
//
// J = [-a -s1 a s2]
//     [0  -1  0  1]
// a = perp
// s1 = cross(d + r1, a) = cross(p2 - x1, a)
// s2 = cross(r2, a) = cross(p2 - x2, a)

// Motor/Limit linear constraint
// C = dot(ax1, d)
// Cdot = -dot(ax1, v1) - dot(cross(d + r1, ax1), w1) + dot(ax1, v2) + dot(cross(r2, ax1), v2)
// J = [-ax1 -cross(d+r1,ax1) ax1 cross(r2,ax1)]

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
	var a11, a12, a21, a22, a31, c, iA, iB, k, mA, mB, omega, s, v53, v55, v71, v75, v77, v78, v79 float32
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var d, rA, rB, v27, v28, v29, v32, v33, v35, v36, v37, v40, v41, v44, v45, v47, v48, v49, v57, v58, v59, v61, v62, v63, v65, v66, v67, v69, v70, v73, v74 Vec2
	var idA, idB, localIndexA, localIndexB, v10, v14, v18, v2, v22, v25, v26, v6 int32
	var qA, qB, v31, v39, v43, v51, v52 Rot
	var v80 b2Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a12, a21, a22, a31, bodyA, bodyB, bodySimA, bodySimB, c, d, iA, iB, idA, idB, joint, k, localIndexA, localIndexB, mA, mB, omega, qA, qB, rA, rB, s, setA, setB, world, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v28, v29, v3, v31, v32, v33, v35, v36, v37, v39, v40, v41, v43, v44, v45, v47, v48, v49, v5, v51, v52, v53, v55, v57, v58, v59, v6, v61, v62, v63, v65, v66, v67, v69, v7, v70, v71, v73, v74, v75, v77, v78, v79, v80, v9
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_prismaticJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+10599, __ccgo_ts+10497, int32FromInt32(281)) != 0 {
		__builtin_trap(tls)
	}
	// chase body id to the solver set where the body lives
	idA = (*b2JointSim)(unsafe.Pointer(base)).BodyIdA
	idB = (*b2JointSim)(unsafe.Pointer(base)).BodyIdB
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	v1 = world + 1024
	v2 = idA
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	bodyA = v3
	v5 = world + 1024
	v6 = idB
	if !(0 <= v6 && v6 < (*b2BodyArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodyArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*128
	goto _8
_8:
	bodyB = v7
	if !((*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+4643, __ccgo_ts+10497, int32FromInt32(292)) != 0 {
		__builtin_trap(tls)
	}
	v9 = world + 1064
	v10 = (*b2Body)(unsafe.Pointer(bodyA)).SetIndex
	if !(0 <= v10 && v10 < (*b2SolverSetArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v11 = (*b2SolverSetArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*88
	goto _12
_12:
	setA = v11
	v13 = world + 1064
	v14 = (*b2Body)(unsafe.Pointer(bodyB)).SetIndex
	if !(0 <= v14 && v14 < (*b2SolverSetArray)(unsafe.Pointer(v13)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v15 = (*b2SolverSetArray)(unsafe.Pointer(v13)).Data + uintptr(v14)*88
	goto _16
_16:
	setB = v15
	localIndexA = (*b2Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b2Body)(unsafe.Pointer(bodyB)).LocalIndex
	v17 = setA
	v18 = localIndexA
	if !(0 <= v18 && v18 < (*b2BodySimArray)(unsafe.Pointer(v17)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	v19 = (*b2BodySimArray)(unsafe.Pointer(v17)).Data + uintptr(v18)*100
	goto _20
_20:
	bodySimA = v19
	v21 = setB
	v22 = localIndexB
	if !(0 <= v22 && v22 < (*b2BodySimArray)(unsafe.Pointer(v21)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	v23 = (*b2BodySimArray)(unsafe.Pointer(v21)).Data + uintptr(v22)*100
	goto _24
_24:
	bodySimB = v23
	mA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvMass
	iA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvInertia
	mB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvMass
	iB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvInertia
	(*b2JointSim)(unsafe.Pointer(base)).InvMassA = mA
	(*b2JointSim)(unsafe.Pointer(base)).InvMassB = mB
	(*b2JointSim)(unsafe.Pointer(base)).InvIA = iA
	(*b2JointSim)(unsafe.Pointer(base)).InvIB = iB
	joint = base + 68
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) {
		v25 = localIndexA
	} else {
		v25 = -int32(1)
	}
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).IndexA = v25
	if (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) {
		v26 = localIndexB
	} else {
		v26 = -int32(1)
	}
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).IndexB = v26
	qA = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	qB = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v27 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorA
	v28 = (*b2BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v29 = Vec2{
		X: v27.X - v28.X,
		Y: v27.Y - v28.Y,
	}
	goto _30
_30:
	v31 = qA
	v32 = v29
	v33 = Vec2{
		X: float32(v31.C*v32.X) - float32(v31.S*v32.Y),
		Y: float32(v31.S*v32.X) + float32(v31.C*v32.Y),
	}
	goto _34
_34:
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).AnchorA = v33
	v35 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorB
	v36 = (*b2BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v37 = Vec2{
		X: v35.X - v36.X,
		Y: v35.Y - v36.Y,
	}
	goto _38
_38:
	v39 = qB
	v40 = v37
	v41 = Vec2{
		X: float32(v39.C*v40.X) - float32(v39.S*v40.Y),
		Y: float32(v39.S*v40.X) + float32(v39.C*v40.Y),
	}
	goto _42
_42:
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).AnchorB = v41
	v43 = qA
	v44 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).LocalAxisA
	v45 = Vec2{
		X: float32(v43.C*v44.X) - float32(v43.S*v44.Y),
		Y: float32(v43.S*v44.X) + float32(v43.C*v44.Y),
	}
	goto _46
_46:
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).AxisA = v45
	v47 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Center
	v48 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Center
	v49 = Vec2{
		X: v47.X - v48.X,
		Y: v47.Y - v48.Y,
	}
	goto _50
_50:
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).DeltaCenter = v49
	v51 = qB
	v52 = qA
	s = float32(v51.S*v52.C) - float32(v51.C*v52.S)
	c = float32(v51.C*v52.C) + float32(v51.S*v52.S)
	v53 = b2Atan2(tls, s, c)
	goto _54
_54:
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).DeltaAngle = v53 - (*b2PrismaticJoint)(unsafe.Pointer(joint)).ReferenceAngle
	v55 = __builtin_remainderf(tls, (*b2PrismaticJoint)(unsafe.Pointer(joint)).DeltaAngle, float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)))
	goto _56
_56:
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).DeltaAngle = v55
	rA = (*b2PrismaticJoint)(unsafe.Pointer(joint)).AnchorA
	rB = (*b2PrismaticJoint)(unsafe.Pointer(joint)).AnchorB
	v57 = rB
	v58 = rA
	v59 = Vec2{
		X: v57.X - v58.X,
		Y: v57.Y - v58.Y,
	}
	goto _60
_60:
	v61 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).DeltaCenter
	v62 = v59
	v63 = Vec2{
		X: v61.X + v62.X,
		Y: v61.Y + v62.Y,
	}
	goto _64
_64:
	d = v63
	v65 = d
	v66 = rA
	v67 = Vec2{
		X: v65.X + v66.X,
		Y: v65.Y + v66.Y,
	}
	goto _68
_68:
	v69 = v67
	v70 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).AxisA
	v71 = float32(v69.X*v70.Y) - float32(v69.Y*v70.X)
	goto _72
_72:
	a12 = v71
	v73 = rB
	v74 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).AxisA
	v75 = float32(v73.X*v74.Y) - float32(v73.Y*v74.X)
	goto _76
_76:
	a22 = v75
	// effective masses
	k = mA + mB + float32(float32(iA*a12)*a12) + float32(float32(iB*a22)*a22)
	if k > float32FromFloat32(0) {
		v77 = float32FromFloat32(1) / k
	} else {
		v77 = float32FromFloat32(0)
	}
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).AxialMass = v77
	v78 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).Hertz
	v79 = (*b2StepContext)(unsafe.Pointer(context)).H
	if v78 == float32FromFloat32(0) {
		v80 = b2Softness{}
		goto _81
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v78)
	a11 = float32(float32FromFloat32(2)*(*b2PrismaticJoint)(unsafe.Pointer(joint)).DampingRatio) + float32(v79*omega)
	a21 = float32(float32(v79*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v80 = b2Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _81
_81:
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).SpringSoftness = v80
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
	var LA, LB, a11, a21, angleImpulse, axialImpulse, iA, iB, mA, mB, perpImpulse, s11, s21, v37, v41, v52, v56, v58, v62, v71, v76 float32
	var P, axisA, d, perpA, rA, rB, v11, v12, v13, v15, v16, v17, v19, v20, v21, v23, v24, v25, v28, v29, v31, v32, v33, v35, v36, v39, v4, v40, v43, v44, v46, v47, v48, v5, v50, v51, v54, v55, v59, v60, v63, v64, v66, v67, v68, v70, v72, v73, v75, v77, v78, v8, v9 Vec2
	var joint, stateA, stateB, v1, v2 uintptr
	var v27, v3, v7 Rot
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = LA, LB, P, a11, a21, angleImpulse, axialImpulse, axisA, d, iA, iB, joint, mA, mB, perpA, perpImpulse, rA, rB, s11, s21, stateA, stateB, v1, v11, v12, v13, v15, v16, v17, v19, v2, v20, v21, v23, v24, v25, v27, v28, v29, v3, v31, v32, v33, v35, v36, v37, v39, v4, v40, v41, v43, v44, v46, v47, v48, v5, v50, v51, v52, v54, v55, v56, v58, v59, v60, v62, v63, v64, v66, v67, v68, v7, v70, v71, v72, v73, v75, v76, v77, v78, v8, v9
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_prismaticJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+10599, __ccgo_ts+10497, int32FromInt32(351)) != 0 {
		__builtin_trap(tls)
	}
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Pointer(bp)) = b2_identityBodyState10
	joint = base + 68
	if (*b2PrismaticJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v1 = bp
	} else {
		v1 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2PrismaticJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v1
	if (*b2PrismaticJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2PrismaticJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v2
	v3 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v4 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).AnchorA
	v5 = Vec2{
		X: float32(v3.C*v4.X) - float32(v3.S*v4.Y),
		Y: float32(v3.S*v4.X) + float32(v3.C*v4.Y),
	}
	goto _6
_6:
	rA = v5
	v7 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v8 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).AnchorB
	v9 = Vec2{
		X: float32(v7.C*v8.X) - float32(v7.S*v8.Y),
		Y: float32(v7.S*v8.X) + float32(v7.C*v8.Y),
	}
	goto _10
_10:
	rB = v9
	v11 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
	v12 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaPosition
	v13 = Vec2{
		X: v11.X - v12.X,
		Y: v11.Y - v12.Y,
	}
	goto _14
_14:
	v15 = v13
	v16 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).DeltaCenter
	v17 = Vec2{
		X: v15.X + v16.X,
		Y: v15.Y + v16.Y,
	}
	goto _18
_18:
	v19 = rB
	v20 = rA
	v21 = Vec2{
		X: v19.X - v20.X,
		Y: v19.Y - v20.Y,
	}
	goto _22
_22:
	v23 = v17
	v24 = v21
	v25 = Vec2{
		X: v23.X + v24.X,
		Y: v23.Y + v24.Y,
	}
	goto _26
_26:
	d = v25
	v27 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v28 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).AxisA
	v29 = Vec2{
		X: float32(v27.C*v28.X) - float32(v27.S*v28.Y),
		Y: float32(v27.S*v28.X) + float32(v27.C*v28.Y),
	}
	goto _30
_30:
	axisA = v29
	v31 = d
	v32 = rA
	v33 = Vec2{
		X: v31.X + v32.X,
		Y: v31.Y + v32.Y,
	}
	goto _34
_34:
	v35 = v33
	v36 = axisA
	v37 = float32(v35.X*v36.Y) - float32(v35.Y*v36.X)
	goto _38
_38:
	// impulse is applied at anchor point on body B
	a11 = v37
	v39 = rB
	v40 = axisA
	v41 = float32(v39.X*v40.Y) - float32(v39.Y*v40.X)
	goto _42
_42:
	a21 = v41
	axialImpulse = (*b2PrismaticJoint)(unsafe.Pointer(joint)).SpringImpulse + (*b2PrismaticJoint)(unsafe.Pointer(joint)).MotorImpulse + (*b2PrismaticJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b2PrismaticJoint)(unsafe.Pointer(joint)).UpperImpulse
	v43 = axisA
	v44 = Vec2{
		X: -v43.Y,
		Y: v43.X,
	}
	goto _45
_45:
	// perpendicular constraint
	perpA = v44
	v46 = d
	v47 = rA
	v48 = Vec2{
		X: v46.X + v47.X,
		Y: v46.Y + v47.Y,
	}
	goto _49
_49:
	v50 = v48
	v51 = perpA
	v52 = float32(v50.X*v51.Y) - float32(v50.Y*v51.X)
	goto _53
_53:
	s11 = v52
	v54 = rB
	v55 = perpA
	v56 = float32(v54.X*v55.Y) - float32(v54.Y*v55.X)
	goto _57
_57:
	s21 = v56
	perpImpulse = (*b2PrismaticJoint)(unsafe.Pointer(joint)).Impulse.X
	angleImpulse = (*b2PrismaticJoint)(unsafe.Pointer(joint)).Impulse.Y
	v58 = axialImpulse
	v59 = axisA
	v60 = Vec2{
		X: float32(v58 * v59.X),
		Y: float32(v58 * v59.Y),
	}
	goto _61
_61:
	v62 = perpImpulse
	v63 = perpA
	v64 = Vec2{
		X: float32(v62 * v63.X),
		Y: float32(v62 * v63.Y),
	}
	goto _65
_65:
	v66 = v60
	v67 = v64
	v68 = Vec2{
		X: v66.X + v67.X,
		Y: v66.Y + v67.Y,
	}
	goto _69
_69:
	P = v68
	LA = float32(axialImpulse*a11) + float32(perpImpulse*s11) + angleImpulse
	LB = float32(axialImpulse*a21) + float32(perpImpulse*s21) + angleImpulse
	v70 = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	v71 = mA
	v72 = P
	v73 = Vec2{
		X: v70.X - float32(v71*v72.X),
		Y: v70.Y - float32(v71*v72.Y),
	}
	goto _74
_74:
	(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = v73
	*(*float32)(unsafe.Pointer(stateA + 8)) -= float32(iA * LA)
	v75 = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	v76 = mB
	v77 = P
	v78 = Vec2{
		X: v75.X + float32(v76*v77.X),
		Y: v75.Y + float32(v76*v77.Y),
	}
	goto _79
_79:
	(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = v78
	*(*float32)(unsafe.Pointer(stateB + 8)) += float32(iB * LB)
}

func b2SolvePrismaticJoint(tls *_Stack, base uintptr, context uintptr, useBias uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var C, C1, C2, Cdot, Cdot1, Cdot2, Cdot3, LA, LA1, LA2, LA3, LA4, LB, LB1, LB2, LB3, LB4, a11, a111, a12, a21, a211, a22, bias, bias1, bias2, c, deltaImpulse, det, iA, iB, impulse, impulse1, impulse2, impulseScale, impulseScale1, impulseScale2, impulseScale3, k11, k12, k22, mA, mB, massScale, massScale1, massScale2, massScale3, maxImpulse, oldImpulse, oldImpulse1, oldImpulse2, s11, s21, s3, translation, wA, wB, v104, v106, v107, v108, v110, v111, v116, v121, v131, v133, v134, v135, v137, v138, v143, v148, v161, v165, v173, v177, v181, v183, v195, v200, v205, v33, v41, v45, v53, v55, v60, v65, v75, v77, v78, v79, v80, v82, v83, v84, v89, v94 float32
	var C3, Cdot4, P, P1, P2, P3, P4, axisA, b9, bias3, d, impulse3, perpA, rA, rB, vA, vB, x, v100, v102, v103, v11, v112, v113, v115, v117, v118, v12, v120, v122, v123, v125, v126, v127, v129, v13, v130, v139, v140, v142, v144, v145, v147, v149, v15, v150, v152, v153, v155, v156, v157, v159, v16, v160, v163, v164, v167, v168, v169, v17, v171, v172, v175, v176, v184, v185, v187, v188, v189, v19, v192, v193, v196, v197, v199, v20, v201, v202, v204, v206, v207, v21, v23, v24, v25, v28, v29, v31, v32, v35, v36, v37, v39, v4, v40, v43, v44, v47, v48, v49, v5, v51, v52, v56, v57, v59, v61, v62, v64, v66, v67, v69, v70, v71, v73, v74, v8, v85, v86, v88, v9, v90, v91, v93, v95, v96, v98, v99 Vec2
	var K, v191 Mat22
	var joint, stateA, stateB, v1, v2 uintptr
	var v179, v180, v27, v3, v7 Rot
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, C1, C2, C3, Cdot, Cdot1, Cdot2, Cdot3, Cdot4, K, LA, LA1, LA2, LA3, LA4, LB, LB1, LB2, LB3, LB4, P, P1, P2, P3, P4, a11, a111, a12, a21, a211, a22, axisA, b9, bias, bias1, bias2, bias3, c, d, deltaImpulse, det, iA, iB, impulse, impulse1, impulse2, impulse3, impulseScale, impulseScale1, impulseScale2, impulseScale3, joint, k11, k12, k22, mA, mB, massScale, massScale1, massScale2, massScale3, maxImpulse, oldImpulse, oldImpulse1, oldImpulse2, perpA, rA, rB, s11, s21, s3, stateA, stateB, translation, vA, vB, wA, wB, x, v1, v100, v102, v103, v104, v106, v107, v108, v11, v110, v111, v112, v113, v115, v116, v117, v118, v12, v120, v121, v122, v123, v125, v126, v127, v129, v13, v130, v131, v133, v134, v135, v137, v138, v139, v140, v142, v143, v144, v145, v147, v148, v149, v15, v150, v152, v153, v155, v156, v157, v159, v16, v160, v161, v163, v164, v165, v167, v168, v169, v17, v171, v172, v173, v175, v176, v177, v179, v180, v181, v183, v184, v185, v187, v188, v189, v19, v191, v192, v193, v195, v196, v197, v199, v2, v20, v200, v201, v202, v204, v205, v206, v207, v21, v23, v24, v25, v27, v28, v29, v3, v31, v32, v33, v35, v36, v37, v39, v4, v40, v41, v43, v44, v45, v47, v48, v49, v5, v51, v52, v53, v55, v56, v57, v59, v60, v61, v62, v64, v65, v66, v67, v69, v7, v70, v71, v73, v74, v75, v77, v78, v79, v8, v80, v82, v83, v84, v85, v86, v88, v89, v9, v90, v91, v93, v94, v95, v96, v98, v99
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_prismaticJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+10599, __ccgo_ts+10497, int32FromInt32(396)) != 0 {
		__builtin_trap(tls)
	}
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Pointer(bp)) = b2_identityBodyState10
	joint = base + 68
	if (*b2PrismaticJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v1 = bp
	} else {
		v1 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2PrismaticJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v1
	if (*b2PrismaticJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2PrismaticJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v2
	vA = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	v3 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v4 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).AnchorA
	v5 = Vec2{
		X: float32(v3.C*v4.X) - float32(v3.S*v4.Y),
		Y: float32(v3.S*v4.X) + float32(v3.C*v4.Y),
	}
	goto _6
_6:
	// current anchors
	rA = v5
	v7 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v8 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).AnchorB
	v9 = Vec2{
		X: float32(v7.C*v8.X) - float32(v7.S*v8.Y),
		Y: float32(v7.S*v8.X) + float32(v7.C*v8.Y),
	}
	goto _10
_10:
	rB = v9
	v11 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
	v12 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaPosition
	v13 = Vec2{
		X: v11.X - v12.X,
		Y: v11.Y - v12.Y,
	}
	goto _14
_14:
	v15 = v13
	v16 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).DeltaCenter
	v17 = Vec2{
		X: v15.X + v16.X,
		Y: v15.Y + v16.Y,
	}
	goto _18
_18:
	v19 = rB
	v20 = rA
	v21 = Vec2{
		X: v19.X - v20.X,
		Y: v19.Y - v20.Y,
	}
	goto _22
_22:
	v23 = v17
	v24 = v21
	v25 = Vec2{
		X: v23.X + v24.X,
		Y: v23.Y + v24.Y,
	}
	goto _26
_26:
	d = v25
	v27 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v28 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).AxisA
	v29 = Vec2{
		X: float32(v27.C*v28.X) - float32(v27.S*v28.Y),
		Y: float32(v27.S*v28.X) + float32(v27.C*v28.Y),
	}
	goto _30
_30:
	axisA = v29
	v31 = axisA
	v32 = d
	v33 = float32(v31.X*v32.X) + float32(v31.Y*v32.Y)
	goto _34
_34:
	translation = v33
	v35 = d
	v36 = rA
	v37 = Vec2{
		X: v35.X + v36.X,
		Y: v35.Y + v36.Y,
	}
	goto _38
_38:
	v39 = v37
	v40 = axisA
	v41 = float32(v39.X*v40.Y) - float32(v39.Y*v40.X)
	goto _42
_42:
	// These scalars are for torques generated by axial forces
	a11 = v41
	v43 = rB
	v44 = axisA
	v45 = float32(v43.X*v44.Y) - float32(v43.Y*v44.X)
	goto _46
_46:
	a21 = v45
	// spring constraint
	if (*b2PrismaticJoint)(unsafe.Pointer(joint)).EnableSpring != 0 {
		// This is a real spring and should be applied even during relax
		C = translation - (*b2PrismaticJoint)(unsafe.Pointer(joint)).TargetTranslation
		bias = float32((*b2PrismaticJoint)(unsafe.Pointer(joint)).SpringSoftness.BiasRate * C)
		massScale = (*b2PrismaticJoint)(unsafe.Pointer(joint)).SpringSoftness.MassScale
		impulseScale = (*b2PrismaticJoint)(unsafe.Pointer(joint)).SpringSoftness.ImpulseScale
		v47 = vB
		v48 = vA
		v49 = Vec2{
			X: v47.X - v48.X,
			Y: v47.Y - v48.Y,
		}
		goto _50
	_50:
		v51 = axisA
		v52 = v49
		v53 = float32(v51.X*v52.X) + float32(v51.Y*v52.Y)
		goto _54
	_54:
		Cdot = v53 + float32(a21*wB) - float32(a11*wA)
		deltaImpulse = float32(float32(-massScale*(*b2PrismaticJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot+bias)) - float32(impulseScale*(*b2PrismaticJoint)(unsafe.Pointer(joint)).SpringImpulse)
		*(*float32)(unsafe.Pointer(joint + 16)) += deltaImpulse
		v55 = deltaImpulse
		v56 = axisA
		v57 = Vec2{
			X: float32(v55 * v56.X),
			Y: float32(v55 * v56.Y),
		}
		goto _58
	_58:
		P = v57
		LA = float32(deltaImpulse * a11)
		LB = float32(deltaImpulse * a21)
		v59 = vA
		v60 = mA
		v61 = P
		v62 = Vec2{
			X: v59.X - float32(v60*v61.X),
			Y: v59.Y - float32(v60*v61.Y),
		}
		goto _63
	_63:
		vA = v62
		wA = wA - float32(iA*LA)
		v64 = vB
		v65 = mB
		v66 = P
		v67 = Vec2{
			X: v64.X + float32(v65*v66.X),
			Y: v64.Y + float32(v65*v66.Y),
		}
		goto _68
	_68:
		vB = v67
		wB = wB + float32(iB*LB)
	}
	// Solve motor constraint
	if (*b2PrismaticJoint)(unsafe.Pointer(joint)).EnableMotor != 0 {
		v69 = vB
		v70 = vA
		v71 = Vec2{
			X: v69.X - v70.X,
			Y: v69.Y - v70.Y,
		}
		goto _72
	_72:
		v73 = axisA
		v74 = v71
		v75 = float32(v73.X*v74.X) + float32(v73.Y*v74.Y)
		goto _76
	_76:
		Cdot1 = v75 + float32(a21*wB) - float32(a11*wA)
		impulse = float32((*b2PrismaticJoint)(unsafe.Pointer(joint)).AxialMass * ((*b2PrismaticJoint)(unsafe.Pointer(joint)).MotorSpeed - Cdot1))
		oldImpulse = (*b2PrismaticJoint)(unsafe.Pointer(joint)).MotorImpulse
		maxImpulse = float32((*b2StepContext)(unsafe.Pointer(context)).H * (*b2PrismaticJoint)(unsafe.Pointer(joint)).MaxMotorForce)
		v77 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).MotorImpulse + impulse
		v78 = -maxImpulse
		v79 = maxImpulse
		if v77 < v78 {
			v82 = v78
		} else {
			if v77 > v79 {
				v83 = v79
			} else {
				v83 = v77
			}
			v82 = v83
		}
		v80 = v82
		goto _81
	_81:
		(*b2PrismaticJoint)(unsafe.Pointer(joint)).MotorImpulse = v80
		impulse = (*b2PrismaticJoint)(unsafe.Pointer(joint)).MotorImpulse - oldImpulse
		v84 = impulse
		v85 = axisA
		v86 = Vec2{
			X: float32(v84 * v85.X),
			Y: float32(v84 * v85.Y),
		}
		goto _87
	_87:
		P1 = v86
		LA1 = float32(impulse * a11)
		LB1 = float32(impulse * a21)
		v88 = vA
		v89 = mA
		v90 = P1
		v91 = Vec2{
			X: v88.X - float32(v89*v90.X),
			Y: v88.Y - float32(v89*v90.Y),
		}
		goto _92
	_92:
		vA = v91
		wA = wA - float32(iA*LA1)
		v93 = vB
		v94 = mB
		v95 = P1
		v96 = Vec2{
			X: v93.X + float32(v94*v95.X),
			Y: v93.Y + float32(v94*v95.Y),
		}
		goto _97
	_97:
		vB = v96
		wB = wB + float32(iB*LB1)
	}
	if (*b2PrismaticJoint)(unsafe.Pointer(joint)).EnableLimit != 0 {
		// Lower limit
		C1 = translation - (*b2PrismaticJoint)(unsafe.Pointer(joint)).LowerTranslation
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
		oldImpulse1 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).LowerImpulse
		v98 = vB
		v99 = vA
		v100 = Vec2{
			X: v98.X - v99.X,
			Y: v98.Y - v99.Y,
		}
		goto _101
	_101:
		v102 = axisA
		v103 = v100
		v104 = float32(v102.X*v103.X) + float32(v102.Y*v103.Y)
		goto _105
	_105:
		Cdot2 = v104 + float32(a21*wB) - float32(a11*wA)
		impulse1 = float32(float32(-(*b2PrismaticJoint)(unsafe.Pointer(joint)).AxialMass*massScale1)*(Cdot2+bias1)) - float32(impulseScale1*oldImpulse1)
		v106 = oldImpulse1 + impulse1
		v107 = float32FromFloat32(0)
		if v106 > v107 {
			v110 = v106
		} else {
			v110 = v107
		}
		v108 = v110
		goto _109
	_109:
		(*b2PrismaticJoint)(unsafe.Pointer(joint)).LowerImpulse = v108
		impulse1 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).LowerImpulse - oldImpulse1
		v111 = impulse1
		v112 = axisA
		v113 = Vec2{
			X: float32(v111 * v112.X),
			Y: float32(v111 * v112.Y),
		}
		goto _114
	_114:
		P2 = v113
		LA2 = float32(impulse1 * a11)
		LB2 = float32(impulse1 * a21)
		v115 = vA
		v116 = mA
		v117 = P2
		v118 = Vec2{
			X: v115.X - float32(v116*v117.X),
			Y: v115.Y - float32(v116*v117.Y),
		}
		goto _119
	_119:
		vA = v118
		wA = wA - float32(iA*LA2)
		v120 = vB
		v121 = mB
		v122 = P2
		v123 = Vec2{
			X: v120.X + float32(v121*v122.X),
			Y: v120.Y + float32(v121*v122.Y),
		}
		goto _124
	_124:
		vB = v123
		wB = wB + float32(iB*LB2)
		// Upper limit
		// Note: signs are flipped to keep C positive when the constraint is satisfied.
		// This also keeps the impulse positive when the limit is active.
		// sign flipped
		C2 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).UpperTranslation - translation
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
		oldImpulse2 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).UpperImpulse
		v125 = vA
		v126 = vB
		v127 = Vec2{
			X: v125.X - v126.X,
			Y: v125.Y - v126.Y,
		}
		goto _128
	_128:
		v129 = axisA
		v130 = v127
		v131 = float32(v129.X*v130.X) + float32(v129.Y*v130.Y)
		goto _132
	_132:
		// sign flipped
		Cdot3 = v131 + float32(a11*wA) - float32(a21*wB)
		impulse2 = float32(float32(-(*b2PrismaticJoint)(unsafe.Pointer(joint)).AxialMass*massScale2)*(Cdot3+bias2)) - float32(impulseScale2*oldImpulse2)
		v133 = oldImpulse2 + impulse2
		v134 = float32FromFloat32(0)
		if v133 > v134 {
			v137 = v133
		} else {
			v137 = v134
		}
		v135 = v137
		goto _136
	_136:
		(*b2PrismaticJoint)(unsafe.Pointer(joint)).UpperImpulse = v135
		impulse2 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).UpperImpulse - oldImpulse2
		v138 = impulse2
		v139 = axisA
		v140 = Vec2{
			X: float32(v138 * v139.X),
			Y: float32(v138 * v139.Y),
		}
		goto _141
	_141:
		P3 = v140
		LA3 = float32(impulse2 * a11)
		LB3 = float32(impulse2 * a21)
		// sign flipped
		v142 = vA
		v143 = mA
		v144 = P3
		v145 = Vec2{
			X: v142.X + float32(v143*v144.X),
			Y: v142.Y + float32(v143*v144.Y),
		}
		goto _146
	_146:
		vA = v145
		wA = wA + float32(iA*LA3)
		v147 = vB
		v148 = mB
		v149 = P3
		v150 = Vec2{
			X: v147.X - float32(v148*v149.X),
			Y: v147.Y - float32(v148*v149.Y),
		}
		goto _151
	_151:
		vB = v150
		wB = wB - float32(iB*LB3)
	}
	// Solve the prismatic constraint in block form
	v152 = axisA
	v153 = Vec2{
		X: -v152.Y,
		Y: v152.X,
	}
	goto _154
_154:
	perpA = v153
	v155 = d
	v156 = rA
	v157 = Vec2{
		X: v155.X + v156.X,
		Y: v155.Y + v156.Y,
	}
	goto _158
_158:
	v159 = v157
	v160 = perpA
	v161 = float32(v159.X*v160.Y) - float32(v159.Y*v160.X)
	goto _162
_162:
	// These scalars are for torques generated by the perpendicular constraint force
	s11 = v161
	v163 = rB
	v164 = perpA
	v165 = float32(v163.X*v164.Y) - float32(v163.Y*v164.X)
	goto _166
_166:
	s21 = v165
	v167 = vB
	v168 = vA
	v169 = Vec2{
		X: v167.X - v168.X,
		Y: v167.Y - v168.Y,
	}
	goto _170
_170:
	v171 = perpA
	v172 = v169
	v173 = float32(v171.X*v172.X) + float32(v171.Y*v172.Y)
	goto _174
_174:
	Cdot4.X = v173 + float32(s21*wB) - float32(s11*wA)
	Cdot4.Y = wB - wA
	bias3 = b2Vec2_zero
	massScale3 = float32FromFloat32(1)
	impulseScale3 = float32FromFloat32(0)
	if useBias != 0 {
		v175 = perpA
		v176 = d
		v177 = float32(v175.X*v176.X) + float32(v175.Y*v176.Y)
		goto _178
	_178:
		C3.X = v177
		v179 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
		v180 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
		s3 = float32(v179.S*v180.C) - float32(v179.C*v180.S)
		c = float32(v179.C*v180.C) + float32(v179.S*v180.S)
		v181 = b2Atan2(tls, s3, c)
		goto _182
	_182:
		C3.Y = v181 + (*b2PrismaticJoint)(unsafe.Pointer(joint)).DeltaAngle
		v183 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate
		v184 = C3
		v185 = Vec2{
			X: float32(v183 * v184.X),
			Y: float32(v183 * v184.Y),
		}
		goto _186
	_186:
		bias3 = v185
		massScale3 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
		impulseScale3 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
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
	v187 = Cdot4
	v188 = bias3
	v189 = Vec2{
		X: v187.X + v188.X,
		Y: v187.Y + v188.Y,
	}
	goto _190
_190:
	v191 = K
	v192 = v189
	a111 = v191.Cx.X
	a12 = v191.Cy.X
	a211 = v191.Cx.Y
	a22 = v191.Cy.Y
	det = float32(a111*a22) - float32(a12*a211)
	if det != float32FromFloat32(0) {
		det = float32FromFloat32(1) / det
	}
	x = Vec2{
		X: float32(det * (float32(a22*v192.X) - float32(a12*v192.Y))),
		Y: float32(det * (float32(a111*v192.Y) - float32(a211*v192.X))),
	}
	v193 = x
	goto _194
_194:
	b9 = v193
	impulse3.X = float32(-massScale3*b9.X) - float32(impulseScale3*(*b2PrismaticJoint)(unsafe.Pointer(joint)).Impulse.X)
	impulse3.Y = float32(-massScale3*b9.Y) - float32(impulseScale3*(*b2PrismaticJoint)(unsafe.Pointer(joint)).Impulse.Y)
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).Impulse.X += impulse3.X
	(*b2PrismaticJoint)(unsafe.Pointer(joint)).Impulse.Y += impulse3.Y
	v195 = impulse3.X
	v196 = perpA
	v197 = Vec2{
		X: float32(v195 * v196.X),
		Y: float32(v195 * v196.Y),
	}
	goto _198
_198:
	P4 = v197
	LA4 = float32(impulse3.X*s11) + impulse3.Y
	LB4 = float32(impulse3.X*s21) + impulse3.Y
	v199 = vA
	v200 = mA
	v201 = P4
	v202 = Vec2{
		X: v199.X - float32(v200*v201.X),
		Y: v199.Y - float32(v200*v201.Y),
	}
	goto _203
_203:
	vA = v202
	wA = wA - float32(iA*LA4)
	v204 = vB
	v205 = mB
	v206 = P4
	v207 = Vec2{
		X: v204.X + float32(v205*v206.X),
		Y: v204.Y + float32(v205*v206.Y),
	}
	goto _208
_208:
	vB = v207
	wB = wB + float32(iB*LB4)
	(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
	(*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
	(*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
}

func b2DrawPrismaticJoint(tls *_Stack, draw uintptr, base uintptr, transformA Transform, transformB Transform) {
	var axis, lower, pA, pB, perp, upper, v10, v11, v13, v15, v16, v18, v2, v20, v21, v23, v24, v26, v28, v29, v3, v31, v33, v34, v36, v38, v39, v41, v43, v44, v46, v48, v49, v51, v53, v54, v6, v7 Vec2
	var c1, c2, c3, c4, c5 b2HexColor1
	var joint uintptr
	var x, y, v14, v19, v27, v32, v37, v42, v47, v52 float32
	var v1, v5 Transform
	var v9 Rot
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axis, c1, c2, c3, c4, c5, joint, lower, pA, pB, perp, upper, x, y, v1, v10, v11, v13, v14, v15, v16, v18, v19, v2, v20, v21, v23, v24, v26, v27, v28, v29, v3, v31, v32, v33, v34, v36, v37, v38, v39, v41, v42, v43, v44, v46, v47, v48, v49, v5, v51, v52, v53, v54, v6, v7, v9
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_prismaticJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+10599, __ccgo_ts+10497, int32FromInt32(635)) != 0 {
		__builtin_trap(tls)
	}
	joint = base + 68
	v1 = transformA
	v2 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorA
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	pA = v3
	v5 = transformB
	v6 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorB
	x = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y) + v5.P.X
	y = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y) + v5.P.Y
	v7 = Vec2{
		X: x,
		Y: y,
	}
	goto _8
_8:
	pB = v7
	v9 = transformA.Q
	v10 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).LocalAxisA
	v11 = Vec2{
		X: float32(v9.C*v10.X) - float32(v9.S*v10.Y),
		Y: float32(v9.S*v10.X) + float32(v9.C*v10.Y),
	}
	goto _12
_12:
	axis = v11
	c1 = int32(b2_colorGray)
	c2 = int32(b2_colorGreen)
	c3 = int32(b2_colorRed)
	c4 = int32(b2_colorBlue)
	c5 = int32(b2_colorDimGray)
	(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, pA, pB, c5, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	if (*b2PrismaticJoint)(unsafe.Pointer(joint)).EnableLimit != 0 {
		v13 = pA
		v14 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).LowerTranslation
		v15 = axis
		v16 = Vec2{
			X: v13.X + float32(v14*v15.X),
			Y: v13.Y + float32(v14*v15.Y),
		}
		goto _17
	_17:
		lower = v16
		v18 = pA
		v19 = (*b2PrismaticJoint)(unsafe.Pointer(joint)).UpperTranslation
		v20 = axis
		v21 = Vec2{
			X: v18.X + float32(v19*v20.X),
			Y: v18.Y + float32(v19*v20.Y),
		}
		goto _22
	_22:
		upper = v21
		v23 = axis
		v24 = Vec2{
			X: -v23.Y,
			Y: v23.X,
		}
		goto _25
	_25:
		perp = v24
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, lower, upper, c1, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		v26 = lower
		v27 = float32FromFloat32(0.1)
		v28 = perp
		v29 = Vec2{
			X: v26.X - float32(v27*v28.X),
			Y: v26.Y - float32(v27*v28.Y),
		}
		goto _30
	_30:
		v31 = lower
		v32 = float32FromFloat32(0.1)
		v33 = perp
		v34 = Vec2{
			X: v31.X + float32(v32*v33.X),
			Y: v31.Y + float32(v32*v33.Y),
		}
		goto _35
	_35:
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, v29, v34, c2, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		v36 = upper
		v37 = float32FromFloat32(0.1)
		v38 = perp
		v39 = Vec2{
			X: v36.X - float32(v37*v38.X),
			Y: v36.Y - float32(v37*v38.Y),
		}
		goto _40
	_40:
		v41 = upper
		v42 = float32FromFloat32(0.1)
		v43 = perp
		v44 = Vec2{
			X: v41.X + float32(v42*v43.X),
			Y: v41.Y + float32(v42*v43.Y),
		}
		goto _45
	_45:
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, v39, v44, c3, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	} else {
		v46 = pA
		v47 = float32FromFloat32(1)
		v48 = axis
		v49 = Vec2{
			X: v46.X - float32(v47*v48.X),
			Y: v46.Y - float32(v47*v48.Y),
		}
		goto _50
	_50:
		v51 = pA
		v52 = float32FromFloat32(1)
		v53 = axis
		v54 = Vec2{
			X: v51.X + float32(v52*v53.X),
			Y: v51.Y + float32(v52*v53.Y),
		}
		goto _55
	_55:
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, v49, v54, c1, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	}
	(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pA, float32FromFloat32(5), c1, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pB, float32FromFloat32(5), c4, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
}

func b2RevoluteJoint_EnableSpring(tls *_Stack, jointId JointId, enableSpring uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	if enableSpring != (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableSpring {
		(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableSpring = enableSpring
		(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).SpringImpulse = float32FromFloat32(0)
	}
}

func b2RevoluteJoint_IsSpringEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableSpring
}

func b2RevoluteJoint_SetSpringHertz(tls *_Stack, jointId JointId, hertz float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).Hertz = hertz
}

func b2RevoluteJoint_GetSpringHertz(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).Hertz
}

func b2RevoluteJoint_SetSpringDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).DampingRatio = dampingRatio
}

func b2RevoluteJoint_GetSpringDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).DampingRatio
}

func b2RevoluteJoint_SetTargetAngle(tls *_Stack, jointId JointId, angle float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).TargetAngle = angle
}

func b2RevoluteJoint_GetTargetAngle(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).TargetAngle
}

func b2RevoluteJoint_GetAngle(tls *_Stack, jointId JointId) (r float32) {
	var angle, c, s, v3, v5 float32
	var jointSim, world uintptr
	var transformA, transformB Transform
	var v1, v2 Rot
	_, _, _, _, _, _, _, _, _, _, _ = angle, c, jointSim, s, transformA, transformB, world, v1, v2, v3, v5
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	jointSim = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	transformA = b2GetBodyTransform(tls, world, (*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdA)
	transformB = b2GetBodyTransform(tls, world, (*b2JointSim)(unsafe.Pointer(jointSim)).BodyIdB)
	v1 = transformB.Q
	v2 = transformA.Q
	s = float32(v1.S*v2.C) - float32(v1.C*v2.S)
	c = float32(v1.C*v2.C) + float32(v1.S*v2.S)
	v3 = b2Atan2(tls, s, c)
	goto _4
_4:
	angle = v3 - (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(jointSim), 68))).ReferenceAngle
	v5 = __builtin_remainderf(tls, angle, float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)))
	goto _6
_6:
	angle = v5
	return angle
}

func b2RevoluteJoint_EnableLimit(tls *_Stack, jointId JointId, enableLimit uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	if enableLimit != (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableLimit {
		(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableLimit = enableLimit
		(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LowerImpulse = float32FromFloat32(0)
		(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).UpperImpulse = float32FromFloat32(0)
	}
}

func b2RevoluteJoint_IsLimitEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableLimit
}

func b2RevoluteJoint_GetLowerLimit(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LowerAngle
}

func b2RevoluteJoint_GetUpperLimit(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).UpperAngle
}

func b2RevoluteJoint_SetLimits(tls *_Stack, jointId JointId, lower float32, upper float32) {
	var joint uintptr
	var v1, v10, v2, v3, v5, v6, v7, v8 float32
	_, _, _, _, _, _, _, _, _ = joint, v1, v10, v2, v3, v5, v6, v7, v8
	if !(lower <= upper) && b2InternalAssertFcn(tls, __ccgo_ts+10482, __ccgo_ts+10631, int32FromInt32(115)) != 0 {
		__builtin_trap(tls)
	}
	if !(lower >= float32(-float32FromFloat32(0.99)*float32FromFloat32(3.14159265359))) && b2InternalAssertFcn(tls, __ccgo_ts+10663, __ccgo_ts+10631, int32FromInt32(116)) != 0 {
		__builtin_trap(tls)
	}
	if !(upper <= float32(float32FromFloat32(0.99)*float32FromFloat32(3.14159265359))) && b2InternalAssertFcn(tls, __ccgo_ts+10687, __ccgo_ts+10631, int32FromInt32(117)) != 0 {
		__builtin_trap(tls)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	if lower != (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LowerAngle || upper != (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).UpperAngle {
		v1 = lower
		v2 = upper
		if v1 < v2 {
			v5 = v1
		} else {
			v5 = v2
		}
		v3 = v5
		goto _4
	_4:
		(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LowerAngle = v3
		v6 = lower
		v7 = upper
		if v6 > v7 {
			v10 = v6
		} else {
			v10 = v7
		}
		v8 = v10
		goto _9
	_9:
		(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).UpperAngle = v8
		(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LowerImpulse = float32FromFloat32(0)
		(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).UpperImpulse = float32FromFloat32(0)
	}
}

func b2RevoluteJoint_EnableMotor(tls *_Stack, jointId JointId, enableMotor uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	if enableMotor != (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableMotor {
		(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableMotor = enableMotor
		(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MotorImpulse = float32FromFloat32(0)
	}
}

func b2RevoluteJoint_IsMotorEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableMotor
}

func b2RevoluteJoint_SetMotorSpeed(tls *_Stack, jointId JointId, motorSpeed float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MotorSpeed = motorSpeed
}

func b2RevoluteJoint_GetMotorSpeed(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MotorSpeed
}

func b2RevoluteJoint_GetMotorTorque(tls *_Stack, jointId JointId) (r float32) {
	var joint, world uintptr
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return float32((*b2World)(unsafe.Pointer(world)).Inv_h * (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MotorImpulse)
}

func b2RevoluteJoint_SetMaxMotorTorque(tls *_Stack, jointId JointId, torque float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	(*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MaxMotorTorque = torque
}

func b2RevoluteJoint_GetMaxMotorTorque(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_revoluteJoint))
	return (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MaxMotorTorque
}

func b2GetRevoluteJointForce(tls *_Stack, world uintptr, base uintptr) (r Vec2) {
	var force, v2, v3 Vec2
	var v1 float32
	_, _, _, _ = force, v1, v2, v3
	v1 = (*b2World)(unsafe.Pointer(world)).Inv_h
	v2 = (*(*b2RevoluteJoint)(unsafe.Add(unsafe.Pointer(base), 68))).LinearImpulse
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
	revolute = base + 68
	torque = float32((*b2World)(unsafe.Pointer(world)).Inv_h * ((*b2RevoluteJoint)(unsafe.Pointer(revolute)).MotorImpulse + (*b2RevoluteJoint)(unsafe.Pointer(revolute)).LowerImpulse - (*b2RevoluteJoint)(unsafe.Pointer(revolute)).UpperImpulse))
	return torque
}

// Point-to-point constraint
// C = p2 - p1
// Cdot = v2 - v1
//      = v2 + cross(w2, r2) - v1 - cross(w1, r1)
// J = [-I -r1_skew I r2_skew ]
// Identity used:
// w k % (rx i + ry j) = w * (-ry i + rx j)

// Motor constraint
// Cdot = w2 - w1
// J = [0 0 -1 0 0 1]
// K = invI1 + invI2

func b2PrepareRevoluteJoint(tls *_Stack, base uintptr, context uintptr) {
	var a11, a21, a31, c, iA, iB, k, mA, mB, omega, s, v49, v51, v52, v53 float32
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var idA, idB, localIndexA, localIndexB, v10, v14, v18, v2, v22, v25, v26, v6 int32
	var v27, v28, v29, v32, v33, v35, v36, v37, v40, v41, v43, v44, v45 Vec2
	var v31, v39, v47, v48 Rot
	var v54 b2Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a21, a31, bodyA, bodyB, bodySimA, bodySimB, c, iA, iB, idA, idB, joint, k, localIndexA, localIndexB, mA, mB, omega, s, setA, setB, world, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v28, v29, v3, v31, v32, v33, v35, v36, v37, v39, v40, v41, v43, v44, v45, v47, v48, v49, v5, v51, v52, v53, v54, v6, v7, v9
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_revoluteJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+10710, __ccgo_ts+10631, int32FromInt32(204)) != 0 {
		__builtin_trap(tls)
	}
	// chase body id to the solver set where the body lives
	idA = (*b2JointSim)(unsafe.Pointer(base)).BodyIdA
	idB = (*b2JointSim)(unsafe.Pointer(base)).BodyIdB
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	v1 = world + 1024
	v2 = idA
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	bodyA = v3
	v5 = world + 1024
	v6 = idB
	if !(0 <= v6 && v6 < (*b2BodyArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodyArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*128
	goto _8
_8:
	bodyB = v7
	if !((*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+4643, __ccgo_ts+10631, int32FromInt32(215)) != 0 {
		__builtin_trap(tls)
	}
	v9 = world + 1064
	v10 = (*b2Body)(unsafe.Pointer(bodyA)).SetIndex
	if !(0 <= v10 && v10 < (*b2SolverSetArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v11 = (*b2SolverSetArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*88
	goto _12
_12:
	setA = v11
	v13 = world + 1064
	v14 = (*b2Body)(unsafe.Pointer(bodyB)).SetIndex
	if !(0 <= v14 && v14 < (*b2SolverSetArray)(unsafe.Pointer(v13)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v15 = (*b2SolverSetArray)(unsafe.Pointer(v13)).Data + uintptr(v14)*88
	goto _16
_16:
	setB = v15
	localIndexA = (*b2Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b2Body)(unsafe.Pointer(bodyB)).LocalIndex
	v17 = setA
	v18 = localIndexA
	if !(0 <= v18 && v18 < (*b2BodySimArray)(unsafe.Pointer(v17)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	v19 = (*b2BodySimArray)(unsafe.Pointer(v17)).Data + uintptr(v18)*100
	goto _20
_20:
	bodySimA = v19
	v21 = setB
	v22 = localIndexB
	if !(0 <= v22 && v22 < (*b2BodySimArray)(unsafe.Pointer(v21)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	v23 = (*b2BodySimArray)(unsafe.Pointer(v21)).Data + uintptr(v22)*100
	goto _24
_24:
	bodySimB = v23
	mA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvMass
	iA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvInertia
	mB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvMass
	iB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvInertia
	(*b2JointSim)(unsafe.Pointer(base)).InvMassA = mA
	(*b2JointSim)(unsafe.Pointer(base)).InvMassB = mB
	(*b2JointSim)(unsafe.Pointer(base)).InvIA = iA
	(*b2JointSim)(unsafe.Pointer(base)).InvIB = iB
	joint = base + 68
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) {
		v25 = localIndexA
	} else {
		v25 = -int32(1)
	}
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).IndexA = v25
	if (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) {
		v26 = localIndexB
	} else {
		v26 = -int32(1)
	}
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).IndexB = v26
	// initial anchors in world space
	v27 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorA
	v28 = (*b2BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v29 = Vec2{
		X: v27.X - v28.X,
		Y: v27.Y - v28.Y,
	}
	goto _30
_30:
	v31 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	v32 = v29
	v33 = Vec2{
		X: float32(v31.C*v32.X) - float32(v31.S*v32.Y),
		Y: float32(v31.S*v32.X) + float32(v31.C*v32.Y),
	}
	goto _34
_34:
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).AnchorA = v33
	v35 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorB
	v36 = (*b2BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v37 = Vec2{
		X: v35.X - v36.X,
		Y: v35.Y - v36.Y,
	}
	goto _38
_38:
	v39 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v40 = v37
	v41 = Vec2{
		X: float32(v39.C*v40.X) - float32(v39.S*v40.Y),
		Y: float32(v39.S*v40.X) + float32(v39.C*v40.Y),
	}
	goto _42
_42:
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).AnchorB = v41
	v43 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Center
	v44 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Center
	v45 = Vec2{
		X: v43.X - v44.X,
		Y: v43.Y - v44.Y,
	}
	goto _46
_46:
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).DeltaCenter = v45
	v47 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v48 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	s = float32(v47.S*v48.C) - float32(v47.C*v48.S)
	c = float32(v47.C*v48.C) + float32(v47.S*v48.S)
	v49 = b2Atan2(tls, s, c)
	goto _50
_50:
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).DeltaAngle = v49
	k = iA + iB
	if k > float32FromFloat32(0) {
		v51 = float32FromFloat32(1) / k
	} else {
		v51 = float32FromFloat32(0)
	}
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).AxialMass = v51
	v52 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).Hertz
	v53 = (*b2StepContext)(unsafe.Pointer(context)).H
	if v52 == float32FromFloat32(0) {
		v54 = b2Softness{}
		goto _55
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v52)
	a11 = float32(float32FromFloat32(2)*(*b2RevoluteJoint)(unsafe.Pointer(joint)).DampingRatio) + float32(v53*omega)
	a21 = float32(float32(v53*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v54 = b2Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _55
_55:
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).SpringSoftness = v54
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
	var axialImpulse, iA, iB, mA, mB, v12, v18, v21, v27 float32
	var joint, stateA, stateB, v1, v2 uintptr
	var rA, rB, v11, v13, v14, v16, v17, v20, v22, v23, v25, v26, v4, v5, v8, v9 Vec2
	var v3, v7 Rot
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axialImpulse, iA, iB, joint, mA, mB, rA, rB, stateA, stateB, v1, v11, v12, v13, v14, v16, v17, v18, v2, v20, v21, v22, v23, v25, v26, v27, v3, v4, v5, v7, v8, v9
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_revoluteJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+10710, __ccgo_ts+10631, int32FromInt32(263)) != 0 {
		__builtin_trap(tls)
	}
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Pointer(bp)) = b2_identityBodyState11
	joint = base + 68
	if (*b2RevoluteJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v1 = bp
	} else {
		v1 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2RevoluteJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v1
	if (*b2RevoluteJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2RevoluteJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v2
	v3 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v4 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).AnchorA
	v5 = Vec2{
		X: float32(v3.C*v4.X) - float32(v3.S*v4.Y),
		Y: float32(v3.S*v4.X) + float32(v3.C*v4.Y),
	}
	goto _6
_6:
	rA = v5
	v7 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v8 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).AnchorB
	v9 = Vec2{
		X: float32(v7.C*v8.X) - float32(v7.S*v8.Y),
		Y: float32(v7.S*v8.X) + float32(v7.C*v8.Y),
	}
	goto _10
_10:
	rB = v9
	axialImpulse = (*b2RevoluteJoint)(unsafe.Pointer(joint)).SpringImpulse + (*b2RevoluteJoint)(unsafe.Pointer(joint)).MotorImpulse + (*b2RevoluteJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b2RevoluteJoint)(unsafe.Pointer(joint)).UpperImpulse
	v11 = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	v12 = mA
	v13 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse
	v14 = Vec2{
		X: v11.X - float32(v12*v13.X),
		Y: v11.Y - float32(v12*v13.Y),
	}
	goto _15
_15:
	(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = v14
	v16 = rA
	v17 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse
	v18 = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	goto _19
_19:
	*(*float32)(unsafe.Pointer(stateA + 8)) -= float32(iA * (v18 + axialImpulse))
	v20 = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	v21 = mB
	v22 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse
	v23 = Vec2{
		X: v20.X + float32(v21*v22.X),
		Y: v20.Y + float32(v21*v22.Y),
	}
	goto _24
_24:
	(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = v23
	v25 = rB
	v26 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse
	v27 = float32(v25.X*v26.Y) - float32(v25.Y*v26.X)
	goto _28
_28:
	*(*float32)(unsafe.Pointer(stateB + 8)) += float32(iB * (v27 + axialImpulse))
}

func b2SolveRevoluteJoint(tls *_Stack, base uintptr, context uintptr, useBias uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var C, C1, C2, Cdot, Cdot1, Cdot2, Cdot3, a11, a12, a21, a22, bias, bias1, bias2, c, det, iA, iB, impulse, impulse1, impulse2, impulse3, impulseScale, impulseScale1, impulseScale2, impulseScale3, jointAngle, jointAngle1, jointAngleDelta, mA, mB, massScale, massScale1, massScale2, massScale3, maxImpulse, oldImpulse, oldImpulse1, oldImpulse2, s4, wA, wB, v10, v104, v11, v12, v14, v15, v18, v20, v22, v23, v24, v26, v27, v28, v29, v31, v40, v48, v5, v7, v76, v89, v9, v95, v98 float32
	var Cdot4, b8, bias3, dcA, dcB, impulse4, rA, rB, separation, vA, vB, x, v100, v102, v103, v33, v34, v37, v38, v41, v42, v44, v45, v46, v49, v50, v52, v53, v54, v56, v57, v58, v60, v61, v62, v64, v65, v66, v68, v69, v70, v72, v73, v74, v77, v78, v80, v81, v82, v85, v86, v88, v90, v91, v93, v94, v97, v99 Vec2
	var K, v84 Mat22
	var dqA, dqB, v16, v17, v3, v32, v36, v4 Rot
	var fixedRotation uint8
	var joint, stateA, stateB, v1, v2 uintptr
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, C1, C2, Cdot, Cdot1, Cdot2, Cdot3, Cdot4, K, a11, a12, a21, a22, b8, bias, bias1, bias2, bias3, c, dcA, dcB, det, dqA, dqB, fixedRotation, iA, iB, impulse, impulse1, impulse2, impulse3, impulse4, impulseScale, impulseScale1, impulseScale2, impulseScale3, joint, jointAngle, jointAngle1, jointAngleDelta, mA, mB, massScale, massScale1, massScale2, massScale3, maxImpulse, oldImpulse, oldImpulse1, oldImpulse2, rA, rB, s4, separation, stateA, stateB, vA, vB, wA, wB, x, v1, v10, v100, v102, v103, v104, v11, v12, v14, v15, v16, v17, v18, v2, v20, v22, v23, v24, v26, v27, v28, v29, v3, v31, v32, v33, v34, v36, v37, v38, v4, v40, v41, v42, v44, v45, v46, v48, v49, v5, v50, v52, v53, v54, v56, v57, v58, v60, v61, v62, v64, v65, v66, v68, v69, v7, v70, v72, v73, v74, v76, v77, v78, v80, v81, v82, v84, v85, v86, v88, v89, v9, v90, v91, v93, v94, v95, v97, v98, v99
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_revoluteJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+10710, __ccgo_ts+10631, int32FromInt32(291)) != 0 {
		__builtin_trap(tls)
	}
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Pointer(bp)) = b2_identityBodyState11
	joint = base + 68
	if (*b2RevoluteJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v1 = bp
	} else {
		v1 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2RevoluteJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v1
	if (*b2RevoluteJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2RevoluteJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v2
	vA = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	dqA = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	dqB = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	fixedRotation = boolUint8(iA+iB == float32FromFloat32(0))
	// Solve spring.
	if (*b2RevoluteJoint)(unsafe.Pointer(joint)).EnableSpring != 0 && int32FromUint8(fixedRotation) == false1 {
		v3 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
		v4 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
		s4 = float32(v3.S*v4.C) - float32(v3.C*v4.S)
		c = float32(v3.C*v4.C) + float32(v3.S*v4.S)
		v5 = b2Atan2(tls, s4, c)
		goto _6
	_6:
		jointAngle = v5 + (*b2RevoluteJoint)(unsafe.Pointer(joint)).DeltaAngle
		v7 = __builtin_remainderf(tls, jointAngle-(*b2RevoluteJoint)(unsafe.Pointer(joint)).TargetAngle, float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)))
		goto _8
	_8:
		jointAngleDelta = v7
		C = jointAngleDelta
		bias = float32((*b2RevoluteJoint)(unsafe.Pointer(joint)).SpringSoftness.BiasRate * C)
		massScale = (*b2RevoluteJoint)(unsafe.Pointer(joint)).SpringSoftness.MassScale
		impulseScale = (*b2RevoluteJoint)(unsafe.Pointer(joint)).SpringSoftness.ImpulseScale
		Cdot = wB - wA
		impulse = float32(float32(-massScale*(*b2RevoluteJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot+bias)) - float32(impulseScale*(*b2RevoluteJoint)(unsafe.Pointer(joint)).SpringImpulse)
		*(*float32)(unsafe.Pointer(joint + 8)) += impulse
		wA = wA - float32(iA*impulse)
		wB = wB + float32(iB*impulse)
	}
	// Solve motor constraint.
	if (*b2RevoluteJoint)(unsafe.Pointer(joint)).EnableMotor != 0 && int32FromUint8(fixedRotation) == false1 {
		Cdot1 = wB - wA - (*b2RevoluteJoint)(unsafe.Pointer(joint)).MotorSpeed
		impulse1 = float32(-(*b2RevoluteJoint)(unsafe.Pointer(joint)).AxialMass * Cdot1)
		oldImpulse = (*b2RevoluteJoint)(unsafe.Pointer(joint)).MotorImpulse
		maxImpulse = float32((*b2StepContext)(unsafe.Pointer(context)).H * (*b2RevoluteJoint)(unsafe.Pointer(joint)).MaxMotorTorque)
		v9 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).MotorImpulse + impulse1
		v10 = -maxImpulse
		v11 = maxImpulse
		if v9 < v10 {
			v14 = v10
		} else {
			if v9 > v11 {
				v15 = v11
			} else {
				v15 = v9
			}
			v14 = v15
		}
		v12 = v14
		goto _13
	_13:
		(*b2RevoluteJoint)(unsafe.Pointer(joint)).MotorImpulse = v12
		impulse1 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).MotorImpulse - oldImpulse
		wA = wA - float32(iA*impulse1)
		wB = wB + float32(iB*impulse1)
	}
	if (*b2RevoluteJoint)(unsafe.Pointer(joint)).EnableLimit != 0 && int32FromUint8(fixedRotation) == false1 {
		v16 = dqB
		v17 = dqA
		s4 = float32(v16.S*v17.C) - float32(v16.C*v17.S)
		c = float32(v16.C*v17.C) + float32(v16.S*v17.S)
		v18 = b2Atan2(tls, s4, c)
		goto _19
	_19:
		jointAngle1 = v18 + (*b2RevoluteJoint)(unsafe.Pointer(joint)).DeltaAngle - (*b2RevoluteJoint)(unsafe.Pointer(joint)).ReferenceAngle
		v20 = __builtin_remainderf(tls, jointAngle1, float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)))
		goto _21
	_21:
		jointAngle1 = v20
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
		Cdot2 = wB - wA
		oldImpulse1 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).LowerImpulse
		impulse2 = float32(float32(-massScale1*(*b2RevoluteJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot2+bias1)) - float32(impulseScale1*oldImpulse1)
		v22 = oldImpulse1 + impulse2
		v23 = float32FromFloat32(0)
		if v22 > v23 {
			v26 = v22
		} else {
			v26 = v23
		}
		v24 = v26
		goto _25
	_25:
		(*b2RevoluteJoint)(unsafe.Pointer(joint)).LowerImpulse = v24
		impulse2 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).LowerImpulse - oldImpulse1
		wA = wA - float32(iA*impulse2)
		wB = wB + float32(iB*impulse2)
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
		Cdot3 = wA - wB
		oldImpulse2 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).UpperImpulse
		impulse3 = float32(float32(-massScale2*(*b2RevoluteJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot3+bias2)) - float32(impulseScale2*oldImpulse2)
		v27 = oldImpulse2 + impulse3
		v28 = float32FromFloat32(0)
		if v27 > v28 {
			v31 = v27
		} else {
			v31 = v28
		}
		v29 = v31
		goto _30
	_30:
		(*b2RevoluteJoint)(unsafe.Pointer(joint)).UpperImpulse = v29
		impulse3 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).UpperImpulse - oldImpulse2
		// sign flipped on applied impulse
		wA = wA + float32(iA*impulse3)
		wB = wB - float32(iB*impulse3)
	}
	// Solve point-to-point constraint
	v32 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v33 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).AnchorA
	v34 = Vec2{
		X: float32(v32.C*v33.X) - float32(v32.S*v33.Y),
		Y: float32(v32.S*v33.X) + float32(v32.C*v33.Y),
	}
	goto _35
_35:
	// J = [-I -r1_skew I r2_skew]
	// r_skew = [-ry; rx]
	// K = [ mA+r1y^2*iA+mB+r2y^2*iB,  -r1y*iA*r1x-r2y*iB*r2x]
	//     [  -r1y*iA*r1x-r2y*iB*r2x, mA+r1x^2*iA+mB+r2x^2*iB]
	// current anchors
	rA = v34
	v36 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v37 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).AnchorB
	v38 = Vec2{
		X: float32(v36.C*v37.X) - float32(v36.S*v37.Y),
		Y: float32(v36.S*v37.X) + float32(v36.C*v37.Y),
	}
	goto _39
_39:
	rB = v38
	v40 = wB
	v41 = rB
	v42 = Vec2{
		X: float32(-v40 * v41.Y),
		Y: float32(v40 * v41.X),
	}
	goto _43
_43:
	v44 = vB
	v45 = v42
	v46 = Vec2{
		X: v44.X + v45.X,
		Y: v44.Y + v45.Y,
	}
	goto _47
_47:
	v48 = wA
	v49 = rA
	v50 = Vec2{
		X: float32(-v48 * v49.Y),
		Y: float32(v48 * v49.X),
	}
	goto _51
_51:
	v52 = vA
	v53 = v50
	v54 = Vec2{
		X: v52.X + v53.X,
		Y: v52.Y + v53.Y,
	}
	goto _55
_55:
	v56 = v46
	v57 = v54
	v58 = Vec2{
		X: v56.X - v57.X,
		Y: v56.Y - v57.Y,
	}
	goto _59
_59:
	Cdot4 = v58
	bias3 = b2Vec2_zero
	massScale3 = float32FromFloat32(1)
	impulseScale3 = float32FromFloat32(0)
	if useBias != 0 {
		dcA = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaPosition
		dcB = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
		v60 = dcB
		v61 = dcA
		v62 = Vec2{
			X: v60.X - v61.X,
			Y: v60.Y - v61.Y,
		}
		goto _63
	_63:
		v64 = rB
		v65 = rA
		v66 = Vec2{
			X: v64.X - v65.X,
			Y: v64.Y - v65.Y,
		}
		goto _67
	_67:
		v68 = v62
		v69 = v66
		v70 = Vec2{
			X: v68.X + v69.X,
			Y: v68.Y + v69.Y,
		}
		goto _71
	_71:
		v72 = v70
		v73 = (*b2RevoluteJoint)(unsafe.Pointer(joint)).DeltaCenter
		v74 = Vec2{
			X: v72.X + v73.X,
			Y: v72.Y + v73.Y,
		}
		goto _75
	_75:
		separation = v74
		v76 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate
		v77 = separation
		v78 = Vec2{
			X: float32(v76 * v77.X),
			Y: float32(v76 * v77.Y),
		}
		goto _79
	_79:
		bias3 = v78
		massScale3 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
		impulseScale3 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
	}
	K.Cx.X = mA + mB + float32(float32(rA.Y*rA.Y)*iA) + float32(float32(rB.Y*rB.Y)*iB)
	K.Cy.X = float32(float32(-rA.Y*rA.X)*iA) - float32(float32(rB.Y*rB.X)*iB)
	K.Cx.Y = K.Cy.X
	K.Cy.Y = mA + mB + float32(float32(rA.X*rA.X)*iA) + float32(float32(rB.X*rB.X)*iB)
	v80 = Cdot4
	v81 = bias3
	v82 = Vec2{
		X: v80.X + v81.X,
		Y: v80.Y + v81.Y,
	}
	goto _83
_83:
	v84 = K
	v85 = v82
	a11 = v84.Cx.X
	a12 = v84.Cy.X
	a21 = v84.Cx.Y
	a22 = v84.Cy.Y
	det = float32(a11*a22) - float32(a12*a21)
	if det != float32FromFloat32(0) {
		det = float32FromFloat32(1) / det
	}
	x = Vec2{
		X: float32(det * (float32(a22*v85.X) - float32(a12*v85.Y))),
		Y: float32(det * (float32(a11*v85.Y) - float32(a21*v85.X))),
	}
	v86 = x
	goto _87
_87:
	b8 = v86
	impulse4.X = float32(-massScale3*b8.X) - float32(impulseScale3*(*b2RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse.X)
	impulse4.Y = float32(-massScale3*b8.Y) - float32(impulseScale3*(*b2RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse.Y)
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse.X += impulse4.X
	(*b2RevoluteJoint)(unsafe.Pointer(joint)).LinearImpulse.Y += impulse4.Y
	v88 = vA
	v89 = mA
	v90 = impulse4
	v91 = Vec2{
		X: v88.X - float32(v89*v90.X),
		Y: v88.Y - float32(v89*v90.Y),
	}
	goto _92
_92:
	vA = v91
	v93 = rA
	v94 = impulse4
	v95 = float32(v93.X*v94.Y) - float32(v93.Y*v94.X)
	goto _96
_96:
	wA = wA - float32(iA*v95)
	v97 = vB
	v98 = mB
	v99 = impulse4
	v100 = Vec2{
		X: v97.X + float32(v98*v99.X),
		Y: v97.Y + float32(v98*v99.Y),
	}
	goto _101
_101:
	vB = v100
	v102 = rB
	v103 = impulse4
	v104 = float32(v102.X*v103.Y) - float32(v102.Y*v103.X)
	goto _105
_105:
	wB = wB + float32(iB*v104)
	(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
	(*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
	(*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
}

func b2DrawRevoluteJoint(tls *_Stack, draw uintptr, base uintptr, transformA Transform, transformB Transform, drawSize float32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var L, angle, c, jointAngle, lowerAngle, s, upperAngle, x, y, v11, v19 float32
	var c1, c2, c3, color b2HexColor1
	var cs CosSin
	var joint uintptr
	var pA, pB, pC, r, ref, rhi, rlo, v15, v16, v17, v2, v25, v26, v27, v29, v3, v30, v31, v35, v36, v37, v6, v7 Vec2
	var rot, rotHi, rotLo, rotRef, v10, v13, v21, v23, v33, v9 Rot
	var v1, v5 Transform
	var _ /* buffer at bp+0 */ [32]uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = L, angle, c, c1, c2, c3, color, cs, joint, jointAngle, lowerAngle, pA, pB, pC, r, ref, rhi, rlo, rot, rotHi, rotLo, rotRef, s, upperAngle, x, y, v1, v10, v11, v13, v15, v16, v17, v19, v2, v21, v23, v25, v26, v27, v29, v3, v30, v31, v33, v35, v36, v37, v5, v6, v7, v9
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_revoluteJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+10710, __ccgo_ts+10631, int32FromInt32(492)) != 0 {
		__builtin_trap(tls)
	}
	joint = base + 68
	v1 = transformA
	v2 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorA
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	pA = v3
	v5 = transformB
	v6 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorB
	x = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y) + v5.P.X
	y = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y) + v5.P.Y
	v7 = Vec2{
		X: x,
		Y: y,
	}
	goto _8
_8:
	pB = v7
	c1 = int32(b2_colorGray)
	c2 = int32(b2_colorGreen)
	c3 = int32(b2_colorRed)
	L = drawSize
	// draw->drawPoint(pA, 3.0f, b2_colorGray40, draw->context);
	// draw->drawPoint(pB, 3.0f, b2_colorLightBlue, draw->context);
	(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawCircleFcn})))(tls, pB, L, c1, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	v9 = transformB.Q
	v10 = transformA.Q
	s = float32(v9.S*v10.C) - float32(v9.C*v10.S)
	c = float32(v9.C*v10.C) + float32(v9.S*v10.S)
	v11 = b2Atan2(tls, s, c)
	goto _12
_12:
	angle = v11
	cs = b2ComputeCosSin(tls, angle)
	v13 = Rot{
		C: cs.Cosine,
		S: cs.Sine,
	}
	goto _14
_14:
	rot = v13
	r = Vec2{
		X: float32(L * rot.C),
		Y: float32(L * rot.S),
	}
	v15 = pB
	v16 = r
	v17 = Vec2{
		X: v15.X + v16.X,
		Y: v15.Y + v16.Y,
	}
	goto _18
_18:
	pC = v17
	(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, pB, pC, c1, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	if (*b2DebugDraw)(unsafe.Pointer(draw)).DrawJointExtras != 0 {
		v19 = __builtin_remainderf(tls, angle-(*b2RevoluteJoint)(unsafe.Pointer(joint)).ReferenceAngle, float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)))
		goto _20
	_20:
		jointAngle = v19
		__builtin_snprintf(tls, bp, uint64(32), __ccgo_ts+10741, vaList(bp+40, float64(float32(float32FromFloat32(180)*jointAngle)/float32FromFloat32(3.14159265359))))
		(*(*func(*_Stack, Vec2, uintptr, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawStringFcn})))(tls, pC, bp, int32(b2_colorWhite), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	}
	lowerAngle = (*b2RevoluteJoint)(unsafe.Pointer(joint)).LowerAngle + (*b2RevoluteJoint)(unsafe.Pointer(joint)).ReferenceAngle
	upperAngle = (*b2RevoluteJoint)(unsafe.Pointer(joint)).UpperAngle + (*b2RevoluteJoint)(unsafe.Pointer(joint)).ReferenceAngle
	if (*b2RevoluteJoint)(unsafe.Pointer(joint)).EnableLimit != 0 {
		cs = b2ComputeCosSin(tls, lowerAngle)
		v21 = Rot{
			C: cs.Cosine,
			S: cs.Sine,
		}
		goto _22
	_22:
		rotLo = v21
		rlo = Vec2{
			X: float32(L * rotLo.C),
			Y: float32(L * rotLo.S),
		}
		cs = b2ComputeCosSin(tls, upperAngle)
		v23 = Rot{
			C: cs.Cosine,
			S: cs.Sine,
		}
		goto _24
	_24:
		rotHi = v23
		rhi = Vec2{
			X: float32(L * rotHi.C),
			Y: float32(L * rotHi.S),
		}
		v25 = pB
		v26 = rlo
		v27 = Vec2{
			X: v25.X + v26.X,
			Y: v25.Y + v26.Y,
		}
		goto _28
	_28:
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, pB, v27, c2, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		v29 = pB
		v30 = rhi
		v31 = Vec2{
			X: v29.X + v30.X,
			Y: v29.Y + v30.Y,
		}
		goto _32
	_32:
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, pB, v31, c3, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		cs = b2ComputeCosSin(tls, (*b2RevoluteJoint)(unsafe.Pointer(joint)).ReferenceAngle)
		v33 = Rot{
			C: cs.Cosine,
			S: cs.Sine,
		}
		goto _34
	_34:
		rotRef = v33
		ref = Vec2{
			X: float32(L * rotRef.C),
			Y: float32(L * rotRef.S),
		}
		v35 = pB
		v36 = ref
		v37 = Vec2{
			X: v35.X + v36.X,
			Y: v35.Y + v36.Y,
		}
		goto _38
	_38:
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, pB, v37, int32(b2_colorBlue), (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	}
	color = int32(b2_colorGold)
	(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, transformA.P, pA, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, pA, pB, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, transformB.P, pB, color, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	// char buffer[32];
	// sprintf(buffer, "%.1f", b2Length(joint->impulse));
	// draw->DrawString(pA, buffer, draw->context);
}

func b2PrepareJointsTask(tls *_Stack, startIndex int32, endIndex int32, context uintptr) {
	var i int32
	var joint, joints uintptr
	_, _, _ = i, joint, joints
	joints = (*b2StepContext)(unsafe.Pointer(context)).Joints
	i = startIndex
	for {
		if !(i < endIndex) {
			break
		}
		joint = *(*uintptr)(unsafe.Pointer(joints + uintptr(i)*8))
		b2PrepareJoint(tls, joint, context)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b2WarmStartJointsTask(tls *_Stack, startIndex int32, endIndex int32, context uintptr, colorIndex int32) {
	var color, joint, joints uintptr
	var i int32
	_, _, _, _ = color, i, joint, joints
	color = (*b2StepContext)(unsafe.Pointer(context)).Graph + uintptr(colorIndex)*56
	joints = (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Data
	if !(0 <= startIndex && startIndex < (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Count) && b2InternalAssertFcn(tls, __ccgo_ts+12405, __ccgo_ts+12460, int32FromInt32(159)) != 0 {
		__builtin_trap(tls)
	}
	if !(startIndex <= endIndex && endIndex <= (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Count) && b2InternalAssertFcn(tls, __ccgo_ts+12484, __ccgo_ts+12460, int32FromInt32(160)) != 0 {
		__builtin_trap(tls)
	}
	i = startIndex
	for {
		if !(i < endIndex) {
			break
		}
		joint = joints + uintptr(i)*196
		b2WarmStartJoint(tls, joint, context)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b2SolveJointsTask(tls *_Stack, startIndex int32, endIndex int32, context uintptr, colorIndex int32, useBias uint8) {
	var color, joint, joints uintptr
	var i int32
	_, _, _, _ = color, i, joint, joints
	color = (*b2StepContext)(unsafe.Pointer(context)).Graph + uintptr(colorIndex)*56
	joints = (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Data
	if !(0 <= startIndex && startIndex < (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Count) && b2InternalAssertFcn(tls, __ccgo_ts+12405, __ccgo_ts+12460, int32FromInt32(177)) != 0 {
		__builtin_trap(tls)
	}
	if !(startIndex <= endIndex && endIndex <= (*b2GraphColor)(unsafe.Pointer(color)).JointSims.Count) && b2InternalAssertFcn(tls, __ccgo_ts+12484, __ccgo_ts+12460, int32FromInt32(178)) != 0 {
		__builtin_trap(tls)
	}
	i = startIndex
	for {
		if !(i < endIndex) {
			break
		}
		joint = joints + uintptr(i)*196
		b2SolveJoint(tls, joint, context, useBias)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b2TransferJoint(tls *_Stack, world uintptr, targetSet uintptr, sourceSet uintptr, joint uintptr) {
	var color, movedJoint, movedJointSim, sourceSim, targetSim, v1, v11, v13, v17, v19, v3, v5, v7, v9 uintptr
	var colorIndex, localIndex, movedId, movedIndex, movedIndex1, newCapacity, v10, v14, v15, v18, v2, v6 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = color, colorIndex, localIndex, movedId, movedIndex, movedIndex1, movedJoint, movedJointSim, newCapacity, sourceSim, targetSim, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v3, v5, v6, v7, v9
	if !(targetSet != sourceSet) && b2InternalAssertFcn(tls, __ccgo_ts+14968, __ccgo_ts+14063, int32FromInt32(560)) != 0 {
		__builtin_trap(tls)
	}
	localIndex = (*b2Joint)(unsafe.Pointer(joint)).LocalIndex
	colorIndex = (*b2Joint)(unsafe.Pointer(joint)).ColorIndex
	if (*b2SolverSet)(unsafe.Pointer(sourceSet)).SetIndex == int32(b2_awakeSet) {
		if !(0 <= colorIndex && colorIndex < int32(_B2_GRAPH_COLOR_COUNT)) && b2InternalAssertFcn(tls, __ccgo_ts+3038, __ccgo_ts+14063, int32FromInt32(569)) != 0 {
			__builtin_trap(tls)
		}
		color = world + 328 + uintptr(colorIndex)*56
		v1 = color + 32
		v2 = localIndex
		if !(0 <= v2 && v2 < (*b2JointSimArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(342)) != 0 {
			__builtin_trap(tls)
		}
		v3 = (*b2JointSimArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*196
		goto _4
	_4:
		sourceSim = v3
	} else {
		if !(colorIndex == -int32FromInt32(1)) && b2InternalAssertFcn(tls, __ccgo_ts+14991, __ccgo_ts+14063, int32FromInt32(576)) != 0 {
			__builtin_trap(tls)
		}
		v5 = sourceSet + 32
		v6 = localIndex
		if !(0 <= v6 && v6 < (*b2JointSimArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(342)) != 0 {
			__builtin_trap(tls)
		}
		v7 = (*b2JointSimArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*196
		goto _8
	_8:
		sourceSim = v7
	}
	// Create target and copy. Fix joint.
	if (*b2SolverSet)(unsafe.Pointer(targetSet)).SetIndex == int32(b2_awakeSet) {
		b2AddJointToGraph(tls, world, sourceSim, joint)
		(*b2Joint)(unsafe.Pointer(joint)).SetIndex = int32(b2_awakeSet)
	} else {
		(*b2Joint)(unsafe.Pointer(joint)).SetIndex = (*b2SolverSet)(unsafe.Pointer(targetSet)).SetIndex
		(*b2Joint)(unsafe.Pointer(joint)).LocalIndex = (*b2SolverSet)(unsafe.Pointer(targetSet)).JointSims.Count
		(*b2Joint)(unsafe.Pointer(joint)).ColorIndex = -int32(1)
		v9 = targetSet + 32
		if (*b2JointSimArray)(unsafe.Pointer(v9)).Count == (*b2JointSimArray)(unsafe.Pointer(v9)).Capacity {
			if (*b2JointSimArray)(unsafe.Pointer(v9)).Capacity < int32(2) {
				v10 = int32(2)
			} else {
				v10 = (*b2JointSimArray)(unsafe.Pointer(v9)).Capacity + (*b2JointSimArray)(unsafe.Pointer(v9)).Capacity>>int32(1)
			}
			newCapacity = v10
			b2JointSimArray_Reserve(tls, v9, newCapacity)
		}
		*(*int32)(unsafe.Pointer(v9 + 8)) += int32(1)
		v11 = (*b2JointSimArray)(unsafe.Pointer(v9)).Data + uintptr((*b2JointSimArray)(unsafe.Pointer(v9)).Count-int32FromInt32(1))*196
		goto _12
	_12:
		targetSim = v11
		memcpy(tls, targetSim, sourceSim, uint64(196))
	}
	// Destroy source.
	if (*b2SolverSet)(unsafe.Pointer(sourceSet)).SetIndex == int32(b2_awakeSet) {
		b2RemoveJointFromGraph(tls, world, (*(*b2JointEdge)(unsafe.Pointer(joint + 20))).BodyId, (*(*b2JointEdge)(unsafe.Pointer(joint + 20 + 1*12))).BodyId, colorIndex, localIndex)
	} else {
		v13 = sourceSet + 32
		v14 = localIndex
		if !(0 <= v14 && v14 < (*b2JointSimArray)(unsafe.Pointer(v13)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(342)) != 0 {
			__builtin_trap(tls)
		}
		movedIndex = -int32(1)
		if v14 != (*b2JointSimArray)(unsafe.Pointer(v13)).Count-int32FromInt32(1) {
			movedIndex = (*b2JointSimArray)(unsafe.Pointer(v13)).Count - int32(1)
			*(*b2JointSim)(unsafe.Pointer((*b2JointSimArray)(unsafe.Pointer(v13)).Data + uintptr(v14)*196)) = *(*b2JointSim)(unsafe.Pointer((*b2JointSimArray)(unsafe.Pointer(v13)).Data + uintptr(movedIndex)*196))
		}
		*(*int32)(unsafe.Pointer(v13 + 8)) -= int32(1)
		v15 = movedIndex
		goto _16
	_16:
		movedIndex1 = v15
		if movedIndex1 != -int32(1) {
			// fix swapped element
			movedJointSim = (*b2SolverSet)(unsafe.Pointer(sourceSet)).JointSims.Data + uintptr(localIndex)*196
			movedId = (*b2JointSim)(unsafe.Pointer(movedJointSim)).JointId
			v17 = world + 1104
			v18 = movedId
			if !(0 <= v18 && v18 < (*b2JointArray)(unsafe.Pointer(v17)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1361, int32FromInt32(341)) != 0 {
				__builtin_trap(tls)
			}
			v19 = (*b2JointArray)(unsafe.Pointer(v17)).Data + uintptr(v18)*72
			goto _20
		_20:
			movedJoint = v19
			(*b2Joint)(unsafe.Pointer(movedJoint)).LocalIndex = localIndex
		}
	}
}

/**@}*/

/**@}*/

func b2WeldJoint_SetLinearHertz(tls *_Stack, jointId JointId, hertz float32) {
	var joint uintptr
	_ = joint
	if !(b2IsValidFloat(tls, hertz) != 0 && hertz >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+9805, __ccgo_ts+15185, int32FromInt32(16)) != 0 {
		__builtin_trap(tls)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_weldJoint))
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LinearHertz = hertz
}

func b2WeldJoint_GetLinearHertz(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_weldJoint))
	return (*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LinearHertz
}

func b2WeldJoint_SetLinearDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	var joint uintptr
	_ = joint
	if !(b2IsValidFloat(tls, dampingRatio) != 0 && dampingRatio >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+9846, __ccgo_ts+15185, int32FromInt32(29)) != 0 {
		__builtin_trap(tls)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_weldJoint))
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LinearDampingRatio = dampingRatio
}

func b2WeldJoint_GetLinearDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_weldJoint))
	return (*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LinearDampingRatio
}

func b2WeldJoint_SetAngularHertz(tls *_Stack, jointId JointId, hertz float32) {
	var joint uintptr
	_ = joint
	if !(b2IsValidFloat(tls, hertz) != 0 && hertz >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+9805, __ccgo_ts+15185, int32FromInt32(42)) != 0 {
		__builtin_trap(tls)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_weldJoint))
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).AngularHertz = hertz
}

func b2WeldJoint_GetAngularHertz(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_weldJoint))
	return (*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).AngularHertz
}

func b2WeldJoint_SetAngularDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	var joint uintptr
	_ = joint
	if !(b2IsValidFloat(tls, dampingRatio) != 0 && dampingRatio >= float32FromFloat32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+9846, __ccgo_ts+15185, int32FromInt32(55)) != 0 {
		__builtin_trap(tls)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_weldJoint))
	(*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).AngularDampingRatio = dampingRatio
}

func b2WeldJoint_GetAngularDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_weldJoint))
	return (*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).AngularDampingRatio
}

func b2GetWeldJointForce(tls *_Stack, world uintptr, base uintptr) (r Vec2) {
	var force, v2, v3 Vec2
	var v1 float32
	_, _, _, _ = force, v1, v2, v3
	v1 = (*b2World)(unsafe.Pointer(world)).Inv_h
	v2 = (*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(base), 68))).LinearImpulse
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
	return float32((*b2World)(unsafe.Pointer(world)).Inv_h * (*(*b2WeldJoint)(unsafe.Add(unsafe.Pointer(base), 68))).AngularImpulse)
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

func b2PrepareWeldJoint(tls *_Stack, base uintptr, context uintptr) {
	var a11, a21, a31, c, iA, iB, ka, mA, mB, omega, s, v49, v51, v53, v54, v55, v58, v59 float32
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var idA, idB, localIndexA, localIndexB, v10, v14, v18, v2, v22, v25, v26, v6 int32
	var qA, qB, v31, v39, v47, v48 Rot
	var v27, v28, v29, v32, v33, v35, v36, v37, v40, v41, v43, v44, v45 Vec2
	var v56, v60 b2Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a21, a31, bodyA, bodyB, bodySimA, bodySimB, c, iA, iB, idA, idB, joint, ka, localIndexA, localIndexB, mA, mB, omega, qA, qB, s, setA, setB, world, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v28, v29, v3, v31, v32, v33, v35, v36, v37, v39, v40, v41, v43, v44, v45, v47, v48, v49, v5, v51, v53, v54, v55, v56, v58, v59, v6, v60, v7, v9
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_weldJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+15213, __ccgo_ts+15185, int32FromInt32(93)) != 0 {
		__builtin_trap(tls)
	}
	// chase body id to the solver set where the body lives
	idA = (*b2JointSim)(unsafe.Pointer(base)).BodyIdA
	idB = (*b2JointSim)(unsafe.Pointer(base)).BodyIdB
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	v1 = world + 1024
	v2 = idA
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	bodyA = v3
	v5 = world + 1024
	v6 = idB
	if !(0 <= v6 && v6 < (*b2BodyArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodyArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*128
	goto _8
_8:
	bodyB = v7
	if !((*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+4643, __ccgo_ts+15185, int32FromInt32(104)) != 0 {
		__builtin_trap(tls)
	}
	v9 = world + 1064
	v10 = (*b2Body)(unsafe.Pointer(bodyA)).SetIndex
	if !(0 <= v10 && v10 < (*b2SolverSetArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v11 = (*b2SolverSetArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*88
	goto _12
_12:
	setA = v11
	v13 = world + 1064
	v14 = (*b2Body)(unsafe.Pointer(bodyB)).SetIndex
	if !(0 <= v14 && v14 < (*b2SolverSetArray)(unsafe.Pointer(v13)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v15 = (*b2SolverSetArray)(unsafe.Pointer(v13)).Data + uintptr(v14)*88
	goto _16
_16:
	setB = v15
	localIndexA = (*b2Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b2Body)(unsafe.Pointer(bodyB)).LocalIndex
	v17 = setA
	v18 = localIndexA
	if !(0 <= v18 && v18 < (*b2BodySimArray)(unsafe.Pointer(v17)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	v19 = (*b2BodySimArray)(unsafe.Pointer(v17)).Data + uintptr(v18)*100
	goto _20
_20:
	bodySimA = v19
	v21 = setB
	v22 = localIndexB
	if !(0 <= v22 && v22 < (*b2BodySimArray)(unsafe.Pointer(v21)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	v23 = (*b2BodySimArray)(unsafe.Pointer(v21)).Data + uintptr(v22)*100
	goto _24
_24:
	bodySimB = v23
	mA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvMass
	iA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvInertia
	mB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvMass
	iB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvInertia
	(*b2JointSim)(unsafe.Pointer(base)).InvMassA = mA
	(*b2JointSim)(unsafe.Pointer(base)).InvMassB = mB
	(*b2JointSim)(unsafe.Pointer(base)).InvIA = iA
	(*b2JointSim)(unsafe.Pointer(base)).InvIB = iB
	joint = base + 68
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) {
		v25 = localIndexA
	} else {
		v25 = -int32(1)
	}
	(*b2WeldJoint)(unsafe.Pointer(joint)).IndexA = v25
	if (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) {
		v26 = localIndexB
	} else {
		v26 = -int32(1)
	}
	(*b2WeldJoint)(unsafe.Pointer(joint)).IndexB = v26
	qA = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	qB = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v27 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorA
	v28 = (*b2BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v29 = Vec2{
		X: v27.X - v28.X,
		Y: v27.Y - v28.Y,
	}
	goto _30
_30:
	v31 = qA
	v32 = v29
	v33 = Vec2{
		X: float32(v31.C*v32.X) - float32(v31.S*v32.Y),
		Y: float32(v31.S*v32.X) + float32(v31.C*v32.Y),
	}
	goto _34
_34:
	(*b2WeldJoint)(unsafe.Pointer(joint)).AnchorA = v33
	v35 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorB
	v36 = (*b2BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v37 = Vec2{
		X: v35.X - v36.X,
		Y: v35.Y - v36.Y,
	}
	goto _38
_38:
	v39 = qB
	v40 = v37
	v41 = Vec2{
		X: float32(v39.C*v40.X) - float32(v39.S*v40.Y),
		Y: float32(v39.S*v40.X) + float32(v39.C*v40.Y),
	}
	goto _42
_42:
	(*b2WeldJoint)(unsafe.Pointer(joint)).AnchorB = v41
	v43 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Center
	v44 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Center
	v45 = Vec2{
		X: v43.X - v44.X,
		Y: v43.Y - v44.Y,
	}
	goto _46
_46:
	(*b2WeldJoint)(unsafe.Pointer(joint)).DeltaCenter = v45
	v47 = qB
	v48 = qA
	s = float32(v47.S*v48.C) - float32(v47.C*v48.S)
	c = float32(v47.C*v48.C) + float32(v47.S*v48.S)
	v49 = b2Atan2(tls, s, c)
	goto _50
_50:
	(*b2WeldJoint)(unsafe.Pointer(joint)).DeltaAngle = v49 - (*b2WeldJoint)(unsafe.Pointer(joint)).ReferenceAngle
	v51 = __builtin_remainderf(tls, (*b2WeldJoint)(unsafe.Pointer(joint)).DeltaAngle, float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)))
	goto _52
_52:
	(*b2WeldJoint)(unsafe.Pointer(joint)).DeltaAngle = v51
	ka = iA + iB
	if ka > float32FromFloat32(0) {
		v53 = float32FromFloat32(1) / ka
	} else {
		v53 = float32FromFloat32(0)
	}
	(*b2WeldJoint)(unsafe.Pointer(joint)).AxialMass = v53
	if (*b2WeldJoint)(unsafe.Pointer(joint)).LinearHertz == float32FromFloat32(0) {
		(*b2WeldJoint)(unsafe.Pointer(joint)).LinearSoftness = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness
	} else {
		v54 = (*b2WeldJoint)(unsafe.Pointer(joint)).LinearHertz
		v55 = (*b2StepContext)(unsafe.Pointer(context)).H
		if v54 == float32FromFloat32(0) {
			v56 = b2Softness{}
			goto _57
		}
		omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v54)
		a11 = float32(float32FromFloat32(2)*(*b2WeldJoint)(unsafe.Pointer(joint)).LinearDampingRatio) + float32(v55*omega)
		a21 = float32(float32(v55*omega) * a11)
		a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
		v56 = b2Softness{
			BiasRate:     omega / a11,
			MassScale:    float32(a21 * a31),
			ImpulseScale: a31,
		}
		goto _57
	_57:
		(*b2WeldJoint)(unsafe.Pointer(joint)).LinearSoftness = v56
	}
	if (*b2WeldJoint)(unsafe.Pointer(joint)).AngularHertz == float32FromFloat32(0) {
		(*b2WeldJoint)(unsafe.Pointer(joint)).AngularSoftness = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness
	} else {
		v58 = (*b2WeldJoint)(unsafe.Pointer(joint)).AngularHertz
		v59 = (*b2StepContext)(unsafe.Pointer(context)).H
		if v58 == float32FromFloat32(0) {
			v60 = b2Softness{}
			goto _61
		}
		omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v58)
		a11 = float32(float32FromFloat32(2)*(*b2WeldJoint)(unsafe.Pointer(joint)).AngularDampingRatio) + float32(v59*omega)
		a21 = float32(float32(v59*omega) * a11)
		a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
		v60 = b2Softness{
			BiasRate:     omega / a11,
			MassScale:    float32(a21 * a31),
			ImpulseScale: a31,
		}
		goto _61
	_61:
		(*b2WeldJoint)(unsafe.Pointer(joint)).AngularSoftness = v60
	}
	if int32FromUint8((*b2StepContext)(unsafe.Pointer(context)).EnableWarmStarting) == false1 {
		(*b2WeldJoint)(unsafe.Pointer(joint)).LinearImpulse = b2Vec2_zero
		(*b2WeldJoint)(unsafe.Pointer(joint)).AngularImpulse = float32FromFloat32(0)
	}
}

func b2WarmStartWeldJoint(tls *_Stack, base uintptr, context uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var iA, iB, mA, mB, v12, v18, v21, v27 float32
	var joint, stateA, stateB, v1, v2 uintptr
	var rA, rB, v11, v13, v14, v16, v17, v20, v22, v23, v25, v26, v4, v5, v8, v9 Vec2
	var v3, v7 Rot
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = iA, iB, joint, mA, mB, rA, rB, stateA, stateB, v1, v11, v12, v13, v14, v16, v17, v18, v2, v20, v21, v22, v23, v25, v26, v27, v3, v4, v5, v7, v8, v9
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Pointer(bp)) = b2_identityBodyState16
	joint = base + 68
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
	v4 = (*b2WeldJoint)(unsafe.Pointer(joint)).AnchorA
	v5 = Vec2{
		X: float32(v3.C*v4.X) - float32(v3.S*v4.Y),
		Y: float32(v3.S*v4.X) + float32(v3.C*v4.Y),
	}
	goto _6
_6:
	rA = v5
	v7 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v8 = (*b2WeldJoint)(unsafe.Pointer(joint)).AnchorB
	v9 = Vec2{
		X: float32(v7.C*v8.X) - float32(v7.S*v8.Y),
		Y: float32(v7.S*v8.X) + float32(v7.C*v8.Y),
	}
	goto _10
_10:
	rB = v9
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
	v16 = rA
	v17 = (*b2WeldJoint)(unsafe.Pointer(joint)).LinearImpulse
	v18 = float32(v16.X*v17.Y) - float32(v16.Y*v17.X)
	goto _19
_19:
	*(*float32)(unsafe.Pointer(stateA + 8)) -= float32(iA * (v18 + (*b2WeldJoint)(unsafe.Pointer(joint)).AngularImpulse))
	v20 = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	v21 = mB
	v22 = (*b2WeldJoint)(unsafe.Pointer(joint)).LinearImpulse
	v23 = Vec2{
		X: v20.X + float32(v21*v22.X),
		Y: v20.Y + float32(v21*v22.Y),
	}
	goto _24
_24:
	(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = v23
	v25 = rB
	v26 = (*b2WeldJoint)(unsafe.Pointer(joint)).LinearImpulse
	v27 = float32(v25.X*v26.Y) - float32(v25.Y*v26.X)
	goto _28
_28:
	*(*float32)(unsafe.Pointer(stateB + 8)) += float32(iB * (v27 + (*b2WeldJoint)(unsafe.Pointer(joint)).AngularImpulse))
}

func b2SolveWeldJoint(tls *_Stack, base uintptr, context uintptr, useBias uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var C, Cdot, a11, a12, a21, a22, bias, c, det, iA, iB, impulse, impulseScale, impulseScale1, mA, mB, massScale, massScale1, s4, wA, wB, v31, v35, v43, v5, v68, v74, v77, v83 float32
	var C1, Cdot1, b7, bias1, dcA, dcB, impulse1, rA, rB, vA, vB, x, v12, v13, v15, v16, v17, v19, v20, v21, v23, v24, v25, v27, v28, v29, v32, v33, v36, v37, v39, v40, v41, v44, v45, v47, v48, v49, v51, v52, v53, v55, v56, v57, v60, v61, v63, v64, v65, v67, v69, v70, v72, v73, v76, v78, v79, v8, v81, v82, v9 Vec2
	var K, v59 Mat22
	var joint, stateA, stateB, v1, v2 uintptr
	var v11, v3, v4, v7 Rot
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, C1, Cdot, Cdot1, K, a11, a12, a21, a22, b7, bias, bias1, c, dcA, dcB, det, iA, iB, impulse, impulse1, impulseScale, impulseScale1, joint, mA, mB, massScale, massScale1, rA, rB, s4, stateA, stateB, vA, vB, wA, wB, x, v1, v11, v12, v13, v15, v16, v17, v19, v2, v20, v21, v23, v24, v25, v27, v28, v29, v3, v31, v32, v33, v35, v36, v37, v39, v4, v40, v41, v43, v44, v45, v47, v48, v49, v5, v51, v52, v53, v55, v56, v57, v59, v60, v61, v63, v64, v65, v67, v68, v69, v7, v70, v72, v73, v74, v76, v77, v78, v79, v8, v81, v82, v83, v9
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_weldJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+15213, __ccgo_ts+15185, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Pointer(bp)) = b2_identityBodyState16
	joint = base + 68
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
	vA = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	// angular constraint
	bias = float32FromFloat32(0)
	massScale = float32FromFloat32(1)
	impulseScale = float32FromFloat32(0)
	if useBias != 0 || (*b2WeldJoint)(unsafe.Pointer(joint)).AngularHertz > float32FromFloat32(0) {
		v3 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
		v4 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
		s4 = float32(v3.S*v4.C) - float32(v3.C*v4.S)
		c = float32(v3.C*v4.C) + float32(v3.S*v4.S)
		v5 = b2Atan2(tls, s4, c)
		goto _6
	_6:
		C = v5 + (*b2WeldJoint)(unsafe.Pointer(joint)).DeltaAngle
		bias = float32((*b2WeldJoint)(unsafe.Pointer(joint)).AngularSoftness.BiasRate * C)
		massScale = (*b2WeldJoint)(unsafe.Pointer(joint)).AngularSoftness.MassScale
		impulseScale = (*b2WeldJoint)(unsafe.Pointer(joint)).AngularSoftness.ImpulseScale
	}
	Cdot = wB - wA
	impulse = float32(float32(-massScale*(*b2WeldJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot+bias)) - float32(impulseScale*(*b2WeldJoint)(unsafe.Pointer(joint)).AngularImpulse)
	*(*float32)(unsafe.Pointer(joint + 52)) += impulse
	wA = wA - float32(iA*impulse)
	wB = wB + float32(iB*impulse)
	// linear constraint
	v7 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v8 = (*b2WeldJoint)(unsafe.Pointer(joint)).AnchorA
	v9 = Vec2{
		X: float32(v7.C*v8.X) - float32(v7.S*v8.Y),
		Y: float32(v7.S*v8.X) + float32(v7.C*v8.Y),
	}
	goto _10
_10:
	rA = v9
	v11 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v12 = (*b2WeldJoint)(unsafe.Pointer(joint)).AnchorB
	v13 = Vec2{
		X: float32(v11.C*v12.X) - float32(v11.S*v12.Y),
		Y: float32(v11.S*v12.X) + float32(v11.C*v12.Y),
	}
	goto _14
_14:
	rB = v13
	bias1 = b2Vec2_zero
	massScale1 = float32FromFloat32(1)
	impulseScale1 = float32FromFloat32(0)
	if useBias != 0 || (*b2WeldJoint)(unsafe.Pointer(joint)).LinearHertz > float32FromFloat32(0) {
		dcA = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaPosition
		dcB = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
		v15 = dcB
		v16 = dcA
		v17 = Vec2{
			X: v15.X - v16.X,
			Y: v15.Y - v16.Y,
		}
		goto _18
	_18:
		v19 = rB
		v20 = rA
		v21 = Vec2{
			X: v19.X - v20.X,
			Y: v19.Y - v20.Y,
		}
		goto _22
	_22:
		v23 = v17
		v24 = v21
		v25 = Vec2{
			X: v23.X + v24.X,
			Y: v23.Y + v24.Y,
		}
		goto _26
	_26:
		v27 = v25
		v28 = (*b2WeldJoint)(unsafe.Pointer(joint)).DeltaCenter
		v29 = Vec2{
			X: v27.X + v28.X,
			Y: v27.Y + v28.Y,
		}
		goto _30
	_30:
		C1 = v29
		v31 = (*b2WeldJoint)(unsafe.Pointer(joint)).LinearSoftness.BiasRate
		v32 = C1
		v33 = Vec2{
			X: float32(v31 * v32.X),
			Y: float32(v31 * v32.Y),
		}
		goto _34
	_34:
		bias1 = v33
		massScale1 = (*b2WeldJoint)(unsafe.Pointer(joint)).LinearSoftness.MassScale
		impulseScale1 = (*b2WeldJoint)(unsafe.Pointer(joint)).LinearSoftness.ImpulseScale
	}
	v35 = wB
	v36 = rB
	v37 = Vec2{
		X: float32(-v35 * v36.Y),
		Y: float32(v35 * v36.X),
	}
	goto _38
_38:
	v39 = vB
	v40 = v37
	v41 = Vec2{
		X: v39.X + v40.X,
		Y: v39.Y + v40.Y,
	}
	goto _42
_42:
	v43 = wA
	v44 = rA
	v45 = Vec2{
		X: float32(-v43 * v44.Y),
		Y: float32(v43 * v44.X),
	}
	goto _46
_46:
	v47 = vA
	v48 = v45
	v49 = Vec2{
		X: v47.X + v48.X,
		Y: v47.Y + v48.Y,
	}
	goto _50
_50:
	v51 = v41
	v52 = v49
	v53 = Vec2{
		X: v51.X - v52.X,
		Y: v51.Y - v52.Y,
	}
	goto _54
_54:
	Cdot1 = v53
	K.Cx.X = mA + mB + float32(float32(rA.Y*rA.Y)*iA) + float32(float32(rB.Y*rB.Y)*iB)
	K.Cy.X = float32(float32(-rA.Y*rA.X)*iA) - float32(float32(rB.Y*rB.X)*iB)
	K.Cx.Y = K.Cy.X
	K.Cy.Y = mA + mB + float32(float32(rA.X*rA.X)*iA) + float32(float32(rB.X*rB.X)*iB)
	v55 = Cdot1
	v56 = bias1
	v57 = Vec2{
		X: v55.X + v56.X,
		Y: v55.Y + v56.Y,
	}
	goto _58
_58:
	v59 = K
	v60 = v57
	a11 = v59.Cx.X
	a12 = v59.Cy.X
	a21 = v59.Cx.Y
	a22 = v59.Cy.Y
	det = float32(a11*a22) - float32(a12*a21)
	if det != float32FromFloat32(0) {
		det = float32FromFloat32(1) / det
	}
	x = Vec2{
		X: float32(det * (float32(a22*v60.X) - float32(a12*v60.Y))),
		Y: float32(det * (float32(a11*v60.Y) - float32(a21*v60.X))),
	}
	v61 = x
	goto _62
_62:
	b7 = v61
	impulse1 = Vec2{
		X: float32(-massScale1*b7.X) - float32(impulseScale1*(*b2WeldJoint)(unsafe.Pointer(joint)).LinearImpulse.X),
		Y: float32(-massScale1*b7.Y) - float32(impulseScale1*(*b2WeldJoint)(unsafe.Pointer(joint)).LinearImpulse.Y),
	}
	v63 = (*b2WeldJoint)(unsafe.Pointer(joint)).LinearImpulse
	v64 = impulse1
	v65 = Vec2{
		X: v63.X + v64.X,
		Y: v63.Y + v64.Y,
	}
	goto _66
_66:
	(*b2WeldJoint)(unsafe.Pointer(joint)).LinearImpulse = v65
	v67 = vA
	v68 = mA
	v69 = impulse1
	v70 = Vec2{
		X: v67.X - float32(v68*v69.X),
		Y: v67.Y - float32(v68*v69.Y),
	}
	goto _71
_71:
	vA = v70
	v72 = rA
	v73 = impulse1
	v74 = float32(v72.X*v73.Y) - float32(v72.Y*v73.X)
	goto _75
_75:
	wA = wA - float32(iA*v74)
	v76 = vB
	v77 = mB
	v78 = impulse1
	v79 = Vec2{
		X: v76.X + float32(v77*v78.X),
		Y: v76.Y + float32(v77*v78.Y),
	}
	goto _80
_80:
	vB = v79
	v81 = rB
	v82 = impulse1
	v83 = float32(v81.X*v82.Y) - float32(v81.Y*v82.X)
	goto _84
_84:
	wB = wB + float32(iB*v83)
	(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
	(*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
	(*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
}

func b2WheelJoint_EnableSpring(tls *_Stack, jointId JointId, enableSpring uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	if enableSpring != (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableSpring {
		(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableSpring = enableSpring
		(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).SpringImpulse = float32FromFloat32(0)
	}
}

func b2WheelJoint_IsSpringEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	return (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableSpring
}

func b2WheelJoint_SetSpringHertz(tls *_Stack, jointId JointId, hertz float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).Hertz = hertz
}

func b2WheelJoint_GetSpringHertz(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	return (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).Hertz
}

func b2WheelJoint_SetSpringDampingRatio(tls *_Stack, jointId JointId, dampingRatio float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).DampingRatio = dampingRatio
}

func b2WheelJoint_GetSpringDampingRatio(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	return (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).DampingRatio
}

func b2WheelJoint_EnableLimit(tls *_Stack, jointId JointId, enableLimit uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	if (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableLimit != enableLimit {
		(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LowerImpulse = float32FromFloat32(0)
		(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).UpperImpulse = float32FromFloat32(0)
		(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableLimit = enableLimit
	}
}

func b2WheelJoint_IsLimitEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	return (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableLimit
}

func b2WheelJoint_GetLowerLimit(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	return (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LowerTranslation
}

func b2WheelJoint_GetUpperLimit(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	return (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).UpperTranslation
}

func b2WheelJoint_SetLimits(tls *_Stack, jointId JointId, lower float32, upper float32) {
	var joint uintptr
	var v1, v10, v2, v3, v5, v6, v7, v8 float32
	_, _, _, _, _, _, _, _, _ = joint, v1, v10, v2, v3, v5, v6, v7, v8
	if !(lower <= upper) && b2InternalAssertFcn(tls, __ccgo_ts+10482, __ccgo_ts+15240, int32FromInt32(88)) != 0 {
		__builtin_trap(tls)
	}
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	if lower != (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LowerTranslation || upper != (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).UpperTranslation {
		v1 = lower
		v2 = upper
		if v1 < v2 {
			v5 = v1
		} else {
			v5 = v2
		}
		v3 = v5
		goto _4
	_4:
		(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LowerTranslation = v3
		v6 = lower
		v7 = upper
		if v6 > v7 {
			v10 = v6
		} else {
			v10 = v7
		}
		v8 = v10
		goto _9
	_9:
		(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).UpperTranslation = v8
		(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).LowerImpulse = float32FromFloat32(0)
		(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).UpperImpulse = float32FromFloat32(0)
	}
}

func b2WheelJoint_EnableMotor(tls *_Stack, jointId JointId, enableMotor uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	if (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableMotor != enableMotor {
		(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MotorImpulse = float32FromFloat32(0)
		(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableMotor = enableMotor
	}
}

func b2WheelJoint_IsMotorEnabled(tls *_Stack, jointId JointId) (r uint8) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	return (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).EnableMotor
}

func b2WheelJoint_SetMotorSpeed(tls *_Stack, jointId JointId, motorSpeed float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MotorSpeed = motorSpeed
}

func b2WheelJoint_GetMotorSpeed(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	return (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MotorSpeed
}

func b2WheelJoint_GetMotorTorque(tls *_Stack, jointId JointId) (r float32) {
	var joint, world uintptr
	_, _ = joint, world
	world = b2GetWorld(tls, int32FromUint16(jointId.World0))
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	return float32((*b2World)(unsafe.Pointer(world)).Inv_h * (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MotorImpulse)
}

func b2WheelJoint_SetMaxMotorTorque(tls *_Stack, jointId JointId, torque float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	(*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MaxMotorTorque = torque
}

func b2WheelJoint_GetMaxMotorTorque(tls *_Stack, jointId JointId) (r float32) {
	var joint uintptr
	_ = joint
	joint = b2GetJointSimCheckType(tls, jointId, int32(b2_wheelJoint))
	return (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(joint), 68))).MaxMotorTorque
}

func b2GetWheelJointForce(tls *_Stack, world uintptr, base uintptr) (r Vec2) {
	var axialForce, perpForce, v4, v8 float32
	var axisA, force, perpA, v1, v10, v12, v13, v14, v2, v5, v6, v9 Vec2
	var joint uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axialForce, axisA, force, joint, perpA, perpForce, v1, v10, v12, v13, v14, v2, v4, v5, v6, v8, v9
	joint = base + 68
	// This is a frame behind
	axisA = (*b2WheelJoint)(unsafe.Pointer(joint)).AxisA
	v1 = axisA
	v2 = Vec2{
		X: -v1.Y,
		Y: v1.X,
	}
	goto _3
_3:
	perpA = v2
	perpForce = float32((*b2World)(unsafe.Pointer(world)).Inv_h * (*b2WheelJoint)(unsafe.Pointer(joint)).PerpImpulse)
	axialForce = float32((*b2World)(unsafe.Pointer(world)).Inv_h * ((*b2WheelJoint)(unsafe.Pointer(joint)).SpringImpulse + (*b2WheelJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b2WheelJoint)(unsafe.Pointer(joint)).UpperImpulse))
	v4 = perpForce
	v5 = perpA
	v6 = Vec2{
		X: float32(v4 * v5.X),
		Y: float32(v4 * v5.Y),
	}
	goto _7
_7:
	v8 = axialForce
	v9 = axisA
	v10 = Vec2{
		X: float32(v8 * v9.X),
		Y: float32(v8 * v9.Y),
	}
	goto _11
_11:
	v12 = v6
	v13 = v10
	v14 = Vec2{
		X: v12.X + v13.X,
		Y: v12.Y + v13.Y,
	}
	goto _15
_15:
	force = v14
	return force
}

func b2GetWheelJointTorque(tls *_Stack, world uintptr, base uintptr) (r float32) {
	return float32((*b2World)(unsafe.Pointer(world)).Inv_h * (*(*b2WheelJoint)(unsafe.Add(unsafe.Pointer(base), 68))).MotorImpulse)
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
	var a11, a12, a21, a22, a31, iA, iB, ka, km, kp, mA, mB, omega, s1, s2, v68, v72, v74, v81, v85, v87, v88, v89, v92 float32
	var axisA, d, perpA, rA, rB, v27, v28, v29, v32, v33, v35, v36, v37, v40, v41, v44, v45, v47, v48, v49, v51, v52, v53, v55, v56, v57, v59, v60, v62, v63, v64, v66, v67, v70, v71, v75, v76, v77, v79, v80, v83, v84 Vec2
	var bodyA, bodyB, bodySimA, bodySimB, joint, setA, setB, world, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var idA, idB, localIndexA, localIndexB, v10, v14, v18, v2, v22, v25, v26, v6 int32
	var qA, qB, v31, v39, v43 Rot
	var v90 b2Softness
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a11, a12, a21, a22, a31, axisA, bodyA, bodyB, bodySimA, bodySimB, d, iA, iB, idA, idB, joint, ka, km, kp, localIndexA, localIndexB, mA, mB, omega, perpA, qA, qB, rA, rB, s1, s2, setA, setB, world, v1, v10, v11, v13, v14, v15, v17, v18, v19, v2, v21, v22, v23, v25, v26, v27, v28, v29, v3, v31, v32, v33, v35, v36, v37, v39, v40, v41, v43, v44, v45, v47, v48, v49, v5, v51, v52, v53, v55, v56, v57, v59, v6, v60, v62, v63, v64, v66, v67, v68, v7, v70, v71, v72, v74, v75, v76, v77, v79, v80, v81, v83, v84, v85, v87, v88, v89, v9, v90, v92
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_wheelJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+15269, __ccgo_ts+15240, int32FromInt32(185)) != 0 {
		__builtin_trap(tls)
	}
	// chase body id to the solver set where the body lives
	idA = (*b2JointSim)(unsafe.Pointer(base)).BodyIdA
	idB = (*b2JointSim)(unsafe.Pointer(base)).BodyIdB
	world = (*b2StepContext)(unsafe.Pointer(context)).World
	v1 = world + 1024
	v2 = idA
	if !(0 <= v2 && v2 < (*b2BodyArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2BodyArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*128
	goto _4
_4:
	bodyA = v3
	v5 = world + 1024
	v6 = idB
	if !(0 <= v6 && v6 < (*b2BodyArray)(unsafe.Pointer(v5)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
		__builtin_trap(tls)
	}
	v7 = (*b2BodyArray)(unsafe.Pointer(v5)).Data + uintptr(v6)*128
	goto _8
_8:
	bodyB = v7
	if !((*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) || (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet)) && b2InternalAssertFcn(tls, __ccgo_ts+4643, __ccgo_ts+15240, int32FromInt32(196)) != 0 {
		__builtin_trap(tls)
	}
	v9 = world + 1064
	v10 = (*b2Body)(unsafe.Pointer(bodyA)).SetIndex
	if !(0 <= v10 && v10 < (*b2SolverSetArray)(unsafe.Pointer(v9)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v11 = (*b2SolverSetArray)(unsafe.Pointer(v9)).Data + uintptr(v10)*88
	goto _12
_12:
	setA = v11
	v13 = world + 1064
	v14 = (*b2Body)(unsafe.Pointer(bodyB)).SetIndex
	if !(0 <= v14 && v14 < (*b2SolverSetArray)(unsafe.Pointer(v13)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+402, int32FromInt32(57)) != 0 {
		__builtin_trap(tls)
	}
	v15 = (*b2SolverSetArray)(unsafe.Pointer(v13)).Data + uintptr(v14)*88
	goto _16
_16:
	setB = v15
	localIndexA = (*b2Body)(unsafe.Pointer(bodyA)).LocalIndex
	localIndexB = (*b2Body)(unsafe.Pointer(bodyB)).LocalIndex
	v17 = setA
	v18 = localIndexA
	if !(0 <= v18 && v18 < (*b2BodySimArray)(unsafe.Pointer(v17)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	v19 = (*b2BodySimArray)(unsafe.Pointer(v17)).Data + uintptr(v18)*100
	goto _20
_20:
	bodySimA = v19
	v21 = setB
	v22 = localIndexB
	if !(0 <= v22 && v22 < (*b2BodySimArray)(unsafe.Pointer(v21)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(193)) != 0 {
		__builtin_trap(tls)
	}
	v23 = (*b2BodySimArray)(unsafe.Pointer(v21)).Data + uintptr(v22)*100
	goto _24
_24:
	bodySimB = v23
	mA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvMass
	iA = (*b2BodySim)(unsafe.Pointer(bodySimA)).InvInertia
	mB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvMass
	iB = (*b2BodySim)(unsafe.Pointer(bodySimB)).InvInertia
	(*b2JointSim)(unsafe.Pointer(base)).InvMassA = mA
	(*b2JointSim)(unsafe.Pointer(base)).InvMassB = mB
	(*b2JointSim)(unsafe.Pointer(base)).InvIA = iA
	(*b2JointSim)(unsafe.Pointer(base)).InvIB = iB
	joint = base + 68
	if (*b2Body)(unsafe.Pointer(bodyA)).SetIndex == int32(b2_awakeSet) {
		v25 = localIndexA
	} else {
		v25 = -int32(1)
	}
	(*b2WheelJoint)(unsafe.Pointer(joint)).IndexA = v25
	if (*b2Body)(unsafe.Pointer(bodyB)).SetIndex == int32(b2_awakeSet) {
		v26 = localIndexB
	} else {
		v26 = -int32(1)
	}
	(*b2WheelJoint)(unsafe.Pointer(joint)).IndexB = v26
	qA = (*b2BodySim)(unsafe.Pointer(bodySimA)).Transform.Q
	qB = (*b2BodySim)(unsafe.Pointer(bodySimB)).Transform.Q
	v27 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorA
	v28 = (*b2BodySim)(unsafe.Pointer(bodySimA)).LocalCenter
	v29 = Vec2{
		X: v27.X - v28.X,
		Y: v27.Y - v28.Y,
	}
	goto _30
_30:
	v31 = qA
	v32 = v29
	v33 = Vec2{
		X: float32(v31.C*v32.X) - float32(v31.S*v32.Y),
		Y: float32(v31.S*v32.X) + float32(v31.C*v32.Y),
	}
	goto _34
_34:
	(*b2WheelJoint)(unsafe.Pointer(joint)).AnchorA = v33
	v35 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorB
	v36 = (*b2BodySim)(unsafe.Pointer(bodySimB)).LocalCenter
	v37 = Vec2{
		X: v35.X - v36.X,
		Y: v35.Y - v36.Y,
	}
	goto _38
_38:
	v39 = qB
	v40 = v37
	v41 = Vec2{
		X: float32(v39.C*v40.X) - float32(v39.S*v40.Y),
		Y: float32(v39.S*v40.X) + float32(v39.C*v40.Y),
	}
	goto _42
_42:
	(*b2WheelJoint)(unsafe.Pointer(joint)).AnchorB = v41
	v43 = qA
	v44 = (*b2WheelJoint)(unsafe.Pointer(joint)).LocalAxisA
	v45 = Vec2{
		X: float32(v43.C*v44.X) - float32(v43.S*v44.Y),
		Y: float32(v43.S*v44.X) + float32(v43.C*v44.Y),
	}
	goto _46
_46:
	(*b2WheelJoint)(unsafe.Pointer(joint)).AxisA = v45
	v47 = (*b2BodySim)(unsafe.Pointer(bodySimB)).Center
	v48 = (*b2BodySim)(unsafe.Pointer(bodySimA)).Center
	v49 = Vec2{
		X: v47.X - v48.X,
		Y: v47.Y - v48.Y,
	}
	goto _50
_50:
	(*b2WheelJoint)(unsafe.Pointer(joint)).DeltaCenter = v49
	rA = (*b2WheelJoint)(unsafe.Pointer(joint)).AnchorA
	rB = (*b2WheelJoint)(unsafe.Pointer(joint)).AnchorB
	v51 = rB
	v52 = rA
	v53 = Vec2{
		X: v51.X - v52.X,
		Y: v51.Y - v52.Y,
	}
	goto _54
_54:
	v55 = (*b2WheelJoint)(unsafe.Pointer(joint)).DeltaCenter
	v56 = v53
	v57 = Vec2{
		X: v55.X + v56.X,
		Y: v55.Y + v56.Y,
	}
	goto _58
_58:
	d = v57
	axisA = (*b2WheelJoint)(unsafe.Pointer(joint)).AxisA
	v59 = axisA
	v60 = Vec2{
		X: -v59.Y,
		Y: v59.X,
	}
	goto _61
_61:
	perpA = v60
	v62 = d
	v63 = rA
	v64 = Vec2{
		X: v62.X + v63.X,
		Y: v62.Y + v63.Y,
	}
	goto _65
_65:
	v66 = v64
	v67 = perpA
	v68 = float32(v66.X*v67.Y) - float32(v66.Y*v67.X)
	goto _69
_69:
	// perpendicular constraint (keep wheel on line)
	s1 = v68
	v70 = rB
	v71 = perpA
	v72 = float32(v70.X*v71.Y) - float32(v70.Y*v71.X)
	goto _73
_73:
	s2 = v72
	kp = mA + mB + float32(float32(iA*s1)*s1) + float32(float32(iB*s2)*s2)
	if kp > float32FromFloat32(0) {
		v74 = float32FromFloat32(1) / kp
	} else {
		v74 = float32FromFloat32(0)
	}
	(*b2WheelJoint)(unsafe.Pointer(joint)).PerpMass = v74
	v75 = d
	v76 = rA
	v77 = Vec2{
		X: v75.X + v76.X,
		Y: v75.Y + v76.Y,
	}
	goto _78
_78:
	v79 = v77
	v80 = axisA
	v81 = float32(v79.X*v80.Y) - float32(v79.Y*v80.X)
	goto _82
_82:
	// spring constraint
	a12 = v81
	v83 = rB
	v84 = axisA
	v85 = float32(v83.X*v84.Y) - float32(v83.Y*v84.X)
	goto _86
_86:
	a22 = v85
	ka = mA + mB + float32(float32(iA*a12)*a12) + float32(float32(iB*a22)*a22)
	if ka > float32FromFloat32(0) {
		v87 = float32FromFloat32(1) / ka
	} else {
		v87 = float32FromFloat32(0)
	}
	(*b2WheelJoint)(unsafe.Pointer(joint)).AxialMass = v87
	v88 = (*b2WheelJoint)(unsafe.Pointer(joint)).Hertz
	v89 = (*b2StepContext)(unsafe.Pointer(context)).H
	if v88 == float32FromFloat32(0) {
		v90 = b2Softness{}
		goto _91
	}
	omega = float32(float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)) * v88)
	a11 = float32(float32FromFloat32(2)*(*b2WheelJoint)(unsafe.Pointer(joint)).DampingRatio) + float32(v89*omega)
	a21 = float32(float32(v89*omega) * a11)
	a31 = float32FromFloat32(1) / (float32FromFloat32(1) + a21)
	v90 = b2Softness{
		BiasRate:     omega / a11,
		MassScale:    float32(a21 * a31),
		ImpulseScale: a31,
	}
	goto _91
_91:
	(*b2WheelJoint)(unsafe.Pointer(joint)).SpringSoftness = v90
	km = iA + iB
	if km > float32FromFloat32(0) {
		v92 = float32FromFloat32(1) / km
	} else {
		v92 = float32FromFloat32(0)
	}
	(*b2WheelJoint)(unsafe.Pointer(joint)).MotorMass = v92
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
	var LA, LB, a11, a21, axialImpulse, iA, iB, mA, mB, s11, s21, v40, v44, v52, v56, v58, v62, v71, v76 float32
	var P, axisA, d, perpA, rA, rB, v11, v12, v13, v15, v16, v17, v19, v20, v21, v23, v24, v25, v28, v29, v31, v32, v34, v35, v36, v38, v39, v4, v42, v43, v46, v47, v48, v5, v50, v51, v54, v55, v59, v60, v63, v64, v66, v67, v68, v70, v72, v73, v75, v77, v78, v8, v9 Vec2
	var joint, stateA, stateB, v1, v2 uintptr
	var v27, v3, v7 Rot
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = LA, LB, P, a11, a21, axialImpulse, axisA, d, iA, iB, joint, mA, mB, perpA, rA, rB, s11, s21, stateA, stateB, v1, v11, v12, v13, v15, v16, v17, v19, v2, v20, v21, v23, v24, v25, v27, v28, v29, v3, v31, v32, v34, v35, v36, v38, v39, v4, v40, v42, v43, v44, v46, v47, v48, v5, v50, v51, v52, v54, v55, v56, v58, v59, v60, v62, v63, v64, v66, v67, v68, v7, v70, v71, v72, v73, v75, v76, v77, v78, v8, v9
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_wheelJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+15269, __ccgo_ts+15240, int32FromInt32(267)) != 0 {
		__builtin_trap(tls)
	}
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Pointer(bp)) = b2_identityBodyState17
	joint = base + 68
	if (*b2WheelJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v1 = bp
	} else {
		v1 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2WheelJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v1
	if (*b2WheelJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2WheelJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v2
	v3 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v4 = (*b2WheelJoint)(unsafe.Pointer(joint)).AnchorA
	v5 = Vec2{
		X: float32(v3.C*v4.X) - float32(v3.S*v4.Y),
		Y: float32(v3.S*v4.X) + float32(v3.C*v4.Y),
	}
	goto _6
_6:
	rA = v5
	v7 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v8 = (*b2WheelJoint)(unsafe.Pointer(joint)).AnchorB
	v9 = Vec2{
		X: float32(v7.C*v8.X) - float32(v7.S*v8.Y),
		Y: float32(v7.S*v8.X) + float32(v7.C*v8.Y),
	}
	goto _10
_10:
	rB = v9
	v11 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
	v12 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaPosition
	v13 = Vec2{
		X: v11.X - v12.X,
		Y: v11.Y - v12.Y,
	}
	goto _14
_14:
	v15 = v13
	v16 = (*b2WheelJoint)(unsafe.Pointer(joint)).DeltaCenter
	v17 = Vec2{
		X: v15.X + v16.X,
		Y: v15.Y + v16.Y,
	}
	goto _18
_18:
	v19 = rB
	v20 = rA
	v21 = Vec2{
		X: v19.X - v20.X,
		Y: v19.Y - v20.Y,
	}
	goto _22
_22:
	v23 = v17
	v24 = v21
	v25 = Vec2{
		X: v23.X + v24.X,
		Y: v23.Y + v24.Y,
	}
	goto _26
_26:
	d = v25
	v27 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v28 = (*b2WheelJoint)(unsafe.Pointer(joint)).AxisA
	v29 = Vec2{
		X: float32(v27.C*v28.X) - float32(v27.S*v28.Y),
		Y: float32(v27.S*v28.X) + float32(v27.C*v28.Y),
	}
	goto _30
_30:
	axisA = v29
	v31 = axisA
	v32 = Vec2{
		X: -v31.Y,
		Y: v31.X,
	}
	goto _33
_33:
	perpA = v32
	v34 = d
	v35 = rA
	v36 = Vec2{
		X: v34.X + v35.X,
		Y: v34.Y + v35.Y,
	}
	goto _37
_37:
	v38 = v36
	v39 = axisA
	v40 = float32(v38.X*v39.Y) - float32(v38.Y*v39.X)
	goto _41
_41:
	a11 = v40
	v42 = rB
	v43 = axisA
	v44 = float32(v42.X*v43.Y) - float32(v42.Y*v43.X)
	goto _45
_45:
	a21 = v44
	v46 = d
	v47 = rA
	v48 = Vec2{
		X: v46.X + v47.X,
		Y: v46.Y + v47.Y,
	}
	goto _49
_49:
	v50 = v48
	v51 = perpA
	v52 = float32(v50.X*v51.Y) - float32(v50.Y*v51.X)
	goto _53
_53:
	s11 = v52
	v54 = rB
	v55 = perpA
	v56 = float32(v54.X*v55.Y) - float32(v54.Y*v55.X)
	goto _57
_57:
	s21 = v56
	axialImpulse = (*b2WheelJoint)(unsafe.Pointer(joint)).SpringImpulse + (*b2WheelJoint)(unsafe.Pointer(joint)).LowerImpulse - (*b2WheelJoint)(unsafe.Pointer(joint)).UpperImpulse
	v58 = axialImpulse
	v59 = axisA
	v60 = Vec2{
		X: float32(v58 * v59.X),
		Y: float32(v58 * v59.Y),
	}
	goto _61
_61:
	v62 = (*b2WheelJoint)(unsafe.Pointer(joint)).PerpImpulse
	v63 = perpA
	v64 = Vec2{
		X: float32(v62 * v63.X),
		Y: float32(v62 * v63.Y),
	}
	goto _65
_65:
	v66 = v60
	v67 = v64
	v68 = Vec2{
		X: v66.X + v67.X,
		Y: v66.Y + v67.Y,
	}
	goto _69
_69:
	P = v68
	LA = float32(axialImpulse*a11) + float32((*b2WheelJoint)(unsafe.Pointer(joint)).PerpImpulse*s11) + (*b2WheelJoint)(unsafe.Pointer(joint)).MotorImpulse
	LB = float32(axialImpulse*a21) + float32((*b2WheelJoint)(unsafe.Pointer(joint)).PerpImpulse*s21) + (*b2WheelJoint)(unsafe.Pointer(joint)).MotorImpulse
	v70 = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	v71 = mA
	v72 = P
	v73 = Vec2{
		X: v70.X - float32(v71*v72.X),
		Y: v70.Y - float32(v71*v72.Y),
	}
	goto _74
_74:
	(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = v73
	*(*float32)(unsafe.Pointer(stateA + 8)) -= float32(iA * LA)
	v75 = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	v76 = mB
	v77 = P
	v78 = Vec2{
		X: v75.X + float32(v76*v77.X),
		Y: v75.Y + float32(v76*v77.Y),
	}
	goto _79
_79:
	(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = v78
	*(*float32)(unsafe.Pointer(stateB + 8)) += float32(iB * LB)
}

func b2SolveWheelJoint(tls *_Stack, base uintptr, context uintptr, useBias uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var C, C1, C2, C3, Cdot, Cdot1, Cdot2, Cdot3, Cdot4, LA, LA1, LA2, LA3, LB, LB1, LB2, LB3, a11, a21, bias, bias1, bias2, bias3, iA, iB, impulse, impulse1, impulse2, impulse3, impulse4, impulseScale, impulseScale1, impulseScale2, impulseScale3, mA, mB, massScale, massScale1, massScale2, massScale3, maxImpulse, oldImpulse, oldImpulse1, oldImpulse2, s11, s21, translation, wA, wB, v109, v111, v112, v113, v115, v116, v121, v126, v135, v143, v147, v155, v157, v162, v167, v33, v41, v45, v47, v48, v49, v50, v52, v53, v60, v62, v67, v72, v82, v84, v85, v86, v88, v89, v94, v99 float32
	var P, P1, P2, P3, axisA, d, perpA, rA, rB, vA, vB, v100, v101, v103, v104, v105, v107, v108, v11, v117, v118, v12, v120, v122, v123, v125, v127, v128, v13, v130, v131, v133, v134, v137, v138, v139, v141, v142, v145, v146, v149, v15, v150, v151, v153, v154, v158, v159, v16, v161, v163, v164, v166, v168, v169, v17, v19, v20, v21, v23, v24, v25, v28, v29, v31, v32, v35, v36, v37, v39, v4, v40, v43, v44, v5, v54, v55, v56, v58, v59, v63, v64, v66, v68, v69, v71, v73, v74, v76, v77, v78, v8, v80, v81, v9, v90, v91, v93, v95, v96, v98 Vec2
	var fixedRotation uint8
	var joint, stateA, stateB, v1, v2 uintptr
	var v27, v3, v7 Rot
	var _ /* dummyState at bp+0 */ b2BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, C1, C2, C3, Cdot, Cdot1, Cdot2, Cdot3, Cdot4, LA, LA1, LA2, LA3, LB, LB1, LB2, LB3, P, P1, P2, P3, a11, a21, axisA, bias, bias1, bias2, bias3, d, fixedRotation, iA, iB, impulse, impulse1, impulse2, impulse3, impulse4, impulseScale, impulseScale1, impulseScale2, impulseScale3, joint, mA, mB, massScale, massScale1, massScale2, massScale3, maxImpulse, oldImpulse, oldImpulse1, oldImpulse2, perpA, rA, rB, s11, s21, stateA, stateB, translation, vA, vB, wA, wB, v1, v100, v101, v103, v104, v105, v107, v108, v109, v11, v111, v112, v113, v115, v116, v117, v118, v12, v120, v121, v122, v123, v125, v126, v127, v128, v13, v130, v131, v133, v134, v135, v137, v138, v139, v141, v142, v143, v145, v146, v147, v149, v15, v150, v151, v153, v154, v155, v157, v158, v159, v16, v161, v162, v163, v164, v166, v167, v168, v169, v17, v19, v2, v20, v21, v23, v24, v25, v27, v28, v29, v3, v31, v32, v33, v35, v36, v37, v39, v4, v40, v41, v43, v44, v45, v47, v48, v49, v5, v50, v52, v53, v54, v55, v56, v58, v59, v60, v62, v63, v64, v66, v67, v68, v69, v7, v71, v72, v73, v74, v76, v77, v78, v8, v80, v81, v82, v84, v85, v86, v88, v89, v9, v90, v91, v93, v94, v95, v96, v98, v99
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_wheelJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+15269, __ccgo_ts+15240, int32FromInt32(308)) != 0 {
		__builtin_trap(tls)
	}
	mA = (*b2JointSim)(unsafe.Pointer(base)).InvMassA
	mB = (*b2JointSim)(unsafe.Pointer(base)).InvMassB
	iA = (*b2JointSim)(unsafe.Pointer(base)).InvIA
	iB = (*b2JointSim)(unsafe.Pointer(base)).InvIB
	// dummy state for static bodies
	*(*b2BodyState)(unsafe.Pointer(bp)) = b2_identityBodyState17
	joint = base + 68
	if (*b2WheelJoint)(unsafe.Pointer(joint)).IndexA == -int32(1) {
		v1 = bp
	} else {
		v1 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2WheelJoint)(unsafe.Pointer(joint)).IndexA)*32
	}
	stateA = v1
	if (*b2WheelJoint)(unsafe.Pointer(joint)).IndexB == -int32(1) {
		v2 = bp
	} else {
		v2 = (*b2StepContext)(unsafe.Pointer(context)).States + uintptr((*b2WheelJoint)(unsafe.Pointer(joint)).IndexB)*32
	}
	stateB = v2
	vA = (*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity
	wA = (*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity
	vB = (*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity
	wB = (*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity
	fixedRotation = boolUint8(iA+iB == float32FromFloat32(0))
	v3 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v4 = (*b2WheelJoint)(unsafe.Pointer(joint)).AnchorA
	v5 = Vec2{
		X: float32(v3.C*v4.X) - float32(v3.S*v4.Y),
		Y: float32(v3.S*v4.X) + float32(v3.C*v4.Y),
	}
	goto _6
_6:
	// current anchors
	rA = v5
	v7 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaRotation
	v8 = (*b2WheelJoint)(unsafe.Pointer(joint)).AnchorB
	v9 = Vec2{
		X: float32(v7.C*v8.X) - float32(v7.S*v8.Y),
		Y: float32(v7.S*v8.X) + float32(v7.C*v8.Y),
	}
	goto _10
_10:
	rB = v9
	v11 = (*b2BodyState)(unsafe.Pointer(stateB)).DeltaPosition
	v12 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaPosition
	v13 = Vec2{
		X: v11.X - v12.X,
		Y: v11.Y - v12.Y,
	}
	goto _14
_14:
	v15 = v13
	v16 = (*b2WheelJoint)(unsafe.Pointer(joint)).DeltaCenter
	v17 = Vec2{
		X: v15.X + v16.X,
		Y: v15.Y + v16.Y,
	}
	goto _18
_18:
	v19 = rB
	v20 = rA
	v21 = Vec2{
		X: v19.X - v20.X,
		Y: v19.Y - v20.Y,
	}
	goto _22
_22:
	v23 = v17
	v24 = v21
	v25 = Vec2{
		X: v23.X + v24.X,
		Y: v23.Y + v24.Y,
	}
	goto _26
_26:
	d = v25
	v27 = (*b2BodyState)(unsafe.Pointer(stateA)).DeltaRotation
	v28 = (*b2WheelJoint)(unsafe.Pointer(joint)).AxisA
	v29 = Vec2{
		X: float32(v27.C*v28.X) - float32(v27.S*v28.Y),
		Y: float32(v27.S*v28.X) + float32(v27.C*v28.Y),
	}
	goto _30
_30:
	axisA = v29
	v31 = axisA
	v32 = d
	v33 = float32(v31.X*v32.X) + float32(v31.Y*v32.Y)
	goto _34
_34:
	translation = v33
	v35 = d
	v36 = rA
	v37 = Vec2{
		X: v35.X + v36.X,
		Y: v35.Y + v36.Y,
	}
	goto _38
_38:
	v39 = v37
	v40 = axisA
	v41 = float32(v39.X*v40.Y) - float32(v39.Y*v40.X)
	goto _42
_42:
	a11 = v41
	v43 = rB
	v44 = axisA
	v45 = float32(v43.X*v44.Y) - float32(v43.Y*v44.X)
	goto _46
_46:
	a21 = v45
	// motor constraint
	if (*b2WheelJoint)(unsafe.Pointer(joint)).EnableMotor != 0 && int32FromUint8(fixedRotation) == false1 {
		Cdot = wB - wA - (*b2WheelJoint)(unsafe.Pointer(joint)).MotorSpeed
		impulse = float32(-(*b2WheelJoint)(unsafe.Pointer(joint)).MotorMass * Cdot)
		oldImpulse = (*b2WheelJoint)(unsafe.Pointer(joint)).MotorImpulse
		maxImpulse = float32((*b2StepContext)(unsafe.Pointer(context)).H * (*b2WheelJoint)(unsafe.Pointer(joint)).MaxMotorTorque)
		v47 = (*b2WheelJoint)(unsafe.Pointer(joint)).MotorImpulse + impulse
		v48 = -maxImpulse
		v49 = maxImpulse
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
		(*b2WheelJoint)(unsafe.Pointer(joint)).MotorImpulse = v50
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
		v54 = vB
		v55 = vA
		v56 = Vec2{
			X: v54.X - v55.X,
			Y: v54.Y - v55.Y,
		}
		goto _57
	_57:
		v58 = axisA
		v59 = v56
		v60 = float32(v58.X*v59.X) + float32(v58.Y*v59.Y)
		goto _61
	_61:
		Cdot1 = v60 + float32(a21*wB) - float32(a11*wA)
		impulse1 = float32(float32(-massScale*(*b2WheelJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot1+bias)) - float32(impulseScale*(*b2WheelJoint)(unsafe.Pointer(joint)).SpringImpulse)
		*(*float32)(unsafe.Pointer(joint + 16)) += impulse1
		v62 = impulse1
		v63 = axisA
		v64 = Vec2{
			X: float32(v62 * v63.X),
			Y: float32(v62 * v63.Y),
		}
		goto _65
	_65:
		P = v64
		LA = float32(impulse1 * a11)
		LB = float32(impulse1 * a21)
		v66 = vA
		v67 = mA
		v68 = P
		v69 = Vec2{
			X: v66.X - float32(v67*v68.X),
			Y: v66.Y - float32(v67*v68.Y),
		}
		goto _70
	_70:
		vA = v69
		wA = wA - float32(iA*LA)
		v71 = vB
		v72 = mB
		v73 = P
		v74 = Vec2{
			X: v71.X + float32(v72*v73.X),
			Y: v71.Y + float32(v72*v73.Y),
		}
		goto _75
	_75:
		vB = v74
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
		v76 = vB
		v77 = vA
		v78 = Vec2{
			X: v76.X - v77.X,
			Y: v76.Y - v77.Y,
		}
		goto _79
	_79:
		v80 = axisA
		v81 = v78
		v82 = float32(v80.X*v81.X) + float32(v80.Y*v81.Y)
		goto _83
	_83:
		Cdot2 = v82 + float32(a21*wB) - float32(a11*wA)
		impulse2 = float32(float32(-massScale1*(*b2WheelJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot2+bias1)) - float32(impulseScale1*(*b2WheelJoint)(unsafe.Pointer(joint)).LowerImpulse)
		oldImpulse1 = (*b2WheelJoint)(unsafe.Pointer(joint)).LowerImpulse
		v84 = oldImpulse1 + impulse2
		v85 = float32FromFloat32(0)
		if v84 > v85 {
			v88 = v84
		} else {
			v88 = v85
		}
		v86 = v88
		goto _87
	_87:
		(*b2WheelJoint)(unsafe.Pointer(joint)).LowerImpulse = v86
		impulse2 = (*b2WheelJoint)(unsafe.Pointer(joint)).LowerImpulse - oldImpulse1
		v89 = impulse2
		v90 = axisA
		v91 = Vec2{
			X: float32(v89 * v90.X),
			Y: float32(v89 * v90.Y),
		}
		goto _92
	_92:
		P1 = v91
		LA1 = float32(impulse2 * a11)
		LB1 = float32(impulse2 * a21)
		v93 = vA
		v94 = mA
		v95 = P1
		v96 = Vec2{
			X: v93.X - float32(v94*v95.X),
			Y: v93.Y - float32(v94*v95.Y),
		}
		goto _97
	_97:
		vA = v96
		wA = wA - float32(iA*LA1)
		v98 = vB
		v99 = mB
		v100 = P1
		v101 = Vec2{
			X: v98.X + float32(v99*v100.X),
			Y: v98.Y + float32(v99*v100.Y),
		}
		goto _102
	_102:
		vB = v101
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
		v103 = vA
		v104 = vB
		v105 = Vec2{
			X: v103.X - v104.X,
			Y: v103.Y - v104.Y,
		}
		goto _106
	_106:
		v107 = axisA
		v108 = v105
		v109 = float32(v107.X*v108.X) + float32(v107.Y*v108.Y)
		goto _110
	_110:
		// sign flipped on Cdot
		Cdot3 = v109 + float32(a11*wA) - float32(a21*wB)
		impulse3 = float32(float32(-massScale2*(*b2WheelJoint)(unsafe.Pointer(joint)).AxialMass)*(Cdot3+bias2)) - float32(impulseScale2*(*b2WheelJoint)(unsafe.Pointer(joint)).UpperImpulse)
		oldImpulse2 = (*b2WheelJoint)(unsafe.Pointer(joint)).UpperImpulse
		v111 = oldImpulse2 + impulse3
		v112 = float32FromFloat32(0)
		if v111 > v112 {
			v115 = v111
		} else {
			v115 = v112
		}
		v113 = v115
		goto _114
	_114:
		(*b2WheelJoint)(unsafe.Pointer(joint)).UpperImpulse = v113
		impulse3 = (*b2WheelJoint)(unsafe.Pointer(joint)).UpperImpulse - oldImpulse2
		v116 = impulse3
		v117 = axisA
		v118 = Vec2{
			X: float32(v116 * v117.X),
			Y: float32(v116 * v117.Y),
		}
		goto _119
	_119:
		P2 = v118
		LA2 = float32(impulse3 * a11)
		LB2 = float32(impulse3 * a21)
		// sign flipped on applied impulse
		v120 = vA
		v121 = mA
		v122 = P2
		v123 = Vec2{
			X: v120.X + float32(v121*v122.X),
			Y: v120.Y + float32(v121*v122.Y),
		}
		goto _124
	_124:
		vA = v123
		wA = wA + float32(iA*LA2)
		v125 = vB
		v126 = mB
		v127 = P2
		v128 = Vec2{
			X: v125.X - float32(v126*v127.X),
			Y: v125.Y - float32(v126*v127.Y),
		}
		goto _129
	_129:
		vB = v128
		wB = wB - float32(iB*LB2)
	}
	// point to line constraint
	v130 = axisA
	v131 = Vec2{
		X: -v130.Y,
		Y: v130.X,
	}
	goto _132
_132:
	perpA = v131
	bias3 = float32FromFloat32(0)
	massScale3 = float32FromFloat32(1)
	impulseScale3 = float32FromFloat32(0)
	if useBias != 0 {
		v133 = perpA
		v134 = d
		v135 = float32(v133.X*v134.X) + float32(v133.Y*v134.Y)
		goto _136
	_136:
		C3 = v135
		bias3 = float32((*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.BiasRate * C3)
		massScale3 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.MassScale
		impulseScale3 = (*b2JointSim)(unsafe.Pointer(base)).ConstraintSoftness.ImpulseScale
	}
	v137 = d
	v138 = rA
	v139 = Vec2{
		X: v137.X + v138.X,
		Y: v137.Y + v138.Y,
	}
	goto _140
_140:
	v141 = v139
	v142 = perpA
	v143 = float32(v141.X*v142.Y) - float32(v141.Y*v142.X)
	goto _144
_144:
	s11 = v143
	v145 = rB
	v146 = perpA
	v147 = float32(v145.X*v146.Y) - float32(v145.Y*v146.X)
	goto _148
_148:
	s21 = v147
	v149 = vB
	v150 = vA
	v151 = Vec2{
		X: v149.X - v150.X,
		Y: v149.Y - v150.Y,
	}
	goto _152
_152:
	v153 = perpA
	v154 = v151
	v155 = float32(v153.X*v154.X) + float32(v153.Y*v154.Y)
	goto _156
_156:
	Cdot4 = v155 + float32(s21*wB) - float32(s11*wA)
	impulse4 = float32(float32(-massScale3*(*b2WheelJoint)(unsafe.Pointer(joint)).PerpMass)*(Cdot4+bias3)) - float32(impulseScale3*(*b2WheelJoint)(unsafe.Pointer(joint)).PerpImpulse)
	*(*float32)(unsafe.Pointer(joint + 8)) += impulse4
	v157 = impulse4
	v158 = perpA
	v159 = Vec2{
		X: float32(v157 * v158.X),
		Y: float32(v157 * v158.Y),
	}
	goto _160
_160:
	P3 = v159
	LA3 = float32(impulse4 * s11)
	LB3 = float32(impulse4 * s21)
	v161 = vA
	v162 = mA
	v163 = P3
	v164 = Vec2{
		X: v161.X - float32(v162*v163.X),
		Y: v161.Y - float32(v162*v163.Y),
	}
	goto _165
_165:
	vA = v164
	wA = wA - float32(iA*LA3)
	v166 = vB
	v167 = mB
	v168 = P3
	v169 = Vec2{
		X: v166.X + float32(v167*v168.X),
		Y: v166.Y + float32(v167*v168.Y),
	}
	goto _170
_170:
	vB = v169
	wB = wB + float32(iB*LB3)
	(*b2BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
	(*b2BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
	(*b2BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
	(*b2BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
}

func b2DrawWheelJoint(tls *_Stack, draw uintptr, base uintptr, transformA Transform, transformB Transform) {
	var axis, lower, pA, pB, perp, upper, v10, v11, v13, v15, v16, v18, v2, v20, v21, v23, v24, v26, v28, v29, v3, v31, v33, v34, v36, v38, v39, v41, v43, v44, v46, v48, v49, v51, v53, v54, v6, v7 Vec2
	var c1, c2, c3, c4, c5 b2HexColor1
	var joint uintptr
	var x, y, v14, v19, v27, v32, v37, v42, v47, v52 float32
	var v1, v5 Transform
	var v9 Rot
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = axis, c1, c2, c3, c4, c5, joint, lower, pA, pB, perp, upper, x, y, v1, v10, v11, v13, v14, v15, v16, v18, v19, v2, v20, v21, v23, v24, v26, v27, v28, v29, v3, v31, v32, v33, v34, v36, v37, v38, v39, v41, v42, v43, v44, v46, v47, v48, v49, v5, v51, v52, v53, v54, v6, v7, v9
	if !((*b2JointSim)(unsafe.Pointer(base)).Type1 == int32(b2_wheelJoint)) && b2InternalAssertFcn(tls, __ccgo_ts+15269, __ccgo_ts+15240, int32FromInt32(519)) != 0 {
		__builtin_trap(tls)
	}
	joint = base + 68
	v1 = transformA
	v2 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorA
	x = float32(v1.Q.C*v2.X) - float32(v1.Q.S*v2.Y) + v1.P.X
	y = float32(v1.Q.S*v2.X) + float32(v1.Q.C*v2.Y) + v1.P.Y
	v3 = Vec2{
		X: x,
		Y: y,
	}
	goto _4
_4:
	pA = v3
	v5 = transformB
	v6 = (*b2JointSim)(unsafe.Pointer(base)).LocalOriginAnchorB
	x = float32(v5.Q.C*v6.X) - float32(v5.Q.S*v6.Y) + v5.P.X
	y = float32(v5.Q.S*v6.X) + float32(v5.Q.C*v6.Y) + v5.P.Y
	v7 = Vec2{
		X: x,
		Y: y,
	}
	goto _8
_8:
	pB = v7
	v9 = transformA.Q
	v10 = (*b2WheelJoint)(unsafe.Pointer(joint)).LocalAxisA
	v11 = Vec2{
		X: float32(v9.C*v10.X) - float32(v9.S*v10.Y),
		Y: float32(v9.S*v10.X) + float32(v9.C*v10.Y),
	}
	goto _12
_12:
	axis = v11
	c1 = int32(b2_colorGray)
	c2 = int32(b2_colorGreen)
	c3 = int32(b2_colorRed)
	c4 = int32(b2_colorDimGray)
	c5 = int32(b2_colorBlue)
	(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, pA, pB, c5, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	if (*b2WheelJoint)(unsafe.Pointer(joint)).EnableLimit != 0 {
		v13 = pA
		v14 = (*b2WheelJoint)(unsafe.Pointer(joint)).LowerTranslation
		v15 = axis
		v16 = Vec2{
			X: v13.X + float32(v14*v15.X),
			Y: v13.Y + float32(v14*v15.Y),
		}
		goto _17
	_17:
		lower = v16
		v18 = pA
		v19 = (*b2WheelJoint)(unsafe.Pointer(joint)).UpperTranslation
		v20 = axis
		v21 = Vec2{
			X: v18.X + float32(v19*v20.X),
			Y: v18.Y + float32(v19*v20.Y),
		}
		goto _22
	_22:
		upper = v21
		v23 = axis
		v24 = Vec2{
			X: -v23.Y,
			Y: v23.X,
		}
		goto _25
	_25:
		perp = v24
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, lower, upper, c1, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		v26 = lower
		v27 = float32FromFloat32(0.1)
		v28 = perp
		v29 = Vec2{
			X: v26.X - float32(v27*v28.X),
			Y: v26.Y - float32(v27*v28.Y),
		}
		goto _30
	_30:
		v31 = lower
		v32 = float32FromFloat32(0.1)
		v33 = perp
		v34 = Vec2{
			X: v31.X + float32(v32*v33.X),
			Y: v31.Y + float32(v32*v33.Y),
		}
		goto _35
	_35:
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, v29, v34, c2, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
		v36 = upper
		v37 = float32FromFloat32(0.1)
		v38 = perp
		v39 = Vec2{
			X: v36.X - float32(v37*v38.X),
			Y: v36.Y - float32(v37*v38.Y),
		}
		goto _40
	_40:
		v41 = upper
		v42 = float32FromFloat32(0.1)
		v43 = perp
		v44 = Vec2{
			X: v41.X + float32(v42*v43.X),
			Y: v41.Y + float32(v42*v43.Y),
		}
		goto _45
	_45:
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, v39, v44, c3, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	} else {
		v46 = pA
		v47 = float32FromFloat32(1)
		v48 = axis
		v49 = Vec2{
			X: v46.X - float32(v47*v48.X),
			Y: v46.Y - float32(v47*v48.Y),
		}
		goto _50
	_50:
		v51 = pA
		v52 = float32FromFloat32(1)
		v53 = axis
		v54 = Vec2{
			X: v51.X + float32(v52*v53.X),
			Y: v51.Y + float32(v52*v53.Y),
		}
		goto _55
	_55:
		(*(*func(*_Stack, Vec2, Vec2, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawSegmentFcn})))(tls, v49, v54, c1, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	}
	(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pA, float32FromFloat32(5), c1, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
	(*(*func(*_Stack, Vec2, float32, HexColor, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2DebugDraw)(unsafe.Pointer(draw)).DrawPointFcn})))(tls, pB, float32FromFloat32(5), c4, (*b2DebugDraw)(unsafe.Pointer(draw)).Context)
}

func b2Joint_IsValid(tls *_Stack, id JointId) (r uint8) {
	var joint, world uintptr
	var jointId int32
	_, _, _ = joint, jointId, world
	if int32(_B2_MAX_WORLDS) <= int32FromUint16(id.World0) {
		return boolUint8(false1 != 0)
	}
	world = uintptr(unsafe.Pointer(&b2_worlds)) + uintptr(id.World0)*1792
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
	if !((*b2Joint)(unsafe.Pointer(joint)).JointId == jointId) && b2InternalAssertFcn(tls, __ccgo_ts+8933, __ccgo_ts+15342, int32FromInt32(1705)) != 0 {
		__builtin_trap(tls)
	}
	return boolUint8(int32FromUint16(id.Generation) == int32FromUint16((*b2Joint)(unsafe.Pointer(joint)).Generation))
}
