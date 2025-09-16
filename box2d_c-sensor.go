package b2

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

func b2SensorArray_Create(tls *_Stack, capacity int32) (r b2SensorArray) {
	var a b2SensorArray
	_ = a
	a = b2SensorArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(40)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2SensorArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2SensorArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2SensorArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2SensorArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2SensorArray)(unsafe.Pointer(a)).Capacity)*uint64(40)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(40)))
	(*b2SensorArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2SensorArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2SensorArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2SensorArray)(unsafe.Pointer(a)).Capacity)*uint64(40)))
	(*b2SensorArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2SensorArray)(unsafe.Pointer(a)).Count = 0
	(*b2SensorArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2SensorTaskContextArray_Create(tls *_Stack, capacity int32) (r b2SensorTaskContextArray) {
	var a b2SensorTaskContextArray
	_ = a
	a = b2SensorTaskContextArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(16)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2SensorTaskContextArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2SensorTaskContextArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2SensorTaskContextArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2SensorTaskContextArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2SensorTaskContextArray)(unsafe.Pointer(a)).Capacity)*uint64(16)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(16)))
	(*b2SensorTaskContextArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2SensorTaskContextArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2SensorTaskContextArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2SensorTaskContextArray)(unsafe.Pointer(a)).Capacity)*uint64(16)))
	(*b2SensorTaskContextArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2SensorTaskContextArray)(unsafe.Pointer(a)).Count = 0
	(*b2SensorTaskContextArray)(unsafe.Pointer(a)).Capacity = 0
}

// Sensor shapes need to
// - detect begin and end overlap events
// - events must be reported in deterministic order
// - maintain an active list of overlaps for query

// Assumption
// - sensors don't detect shapes on the same body

// Algorithm
// Query all sensors for overlaps
// Check against previous overlaps

// Data structures
// Each sensor has an double buffered array of overlaps
// These overlaps use a shape reference with index and generation

func b2SensorQueryCallback(tls *_Stack, proxyId int32, userData uint64, context uintptr) (r uint8) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var newCapacity, sensorShapeId, shapeId, v10, v2 int32
	var otherShape, queryContext, sensor, sensorShape, shapeRef, world, v1, v11, v3, v9 uintptr
	var otherTransform Transform
	var output DistanceOutput
	var overlaps, v7 uint8
	var v5, v6 Filter
	var _ /* cache at bp+184 */ SimplexCache
	var _ /* input at bp+0 */ DistanceInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = newCapacity, otherShape, otherTransform, output, overlaps, queryContext, sensor, sensorShape, sensorShapeId, shapeId, shapeRef, world, v1, v10, v11, v2, v3, v5, v6, v7, v9
	_ = uint64FromInt64(4)
	shapeId = int32FromUint64(userData)
	queryContext = context
	sensorShape = (*b2SensorQueryContext)(unsafe.Pointer(queryContext)).SensorShape
	sensorShapeId = (*b2Shape)(unsafe.Pointer(sensorShape)).Id
	if shapeId == sensorShapeId {
		return boolUint8(true1 != 0)
	}
	world = (*b2SensorQueryContext)(unsafe.Pointer(queryContext)).World
	v1 = world + 1248
	v2 = shapeId
	if !(0 <= v2 && v2 < (*b2ShapeArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2ShapeArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*288
	goto _4
_4:
	otherShape = v3
	// Are sensor events enabled on the other shape?
	if int32FromUint8((*b2Shape)(unsafe.Pointer(otherShape)).EnableSensorEvents) == false1 {
		return boolUint8(true1 != 0)
	}
	// Skip shapes on the same body
	if (*b2Shape)(unsafe.Pointer(otherShape)).BodyId == (*b2Shape)(unsafe.Pointer(sensorShape)).BodyId {
		return boolUint8(true1 != 0)
	}
	// Check filter
	v5 = (*b2Shape)(unsafe.Pointer(sensorShape)).Filter
	v6 = (*b2Shape)(unsafe.Pointer(otherShape)).Filter
	if v5.GroupIndex == v6.GroupIndex && v5.GroupIndex != 0 {
		v7 = boolUint8(v5.GroupIndex > 0)
		goto _8
	}
	v7 = boolUint8(v5.MaskBits&v6.CategoryBits != uint64(0) && v5.CategoryBits&v6.MaskBits != uint64(0))
	goto _8
_8:
	if int32FromUint8(v7) == false1 {
		return boolUint8(true1 != 0)
	}
	otherTransform = b2GetBodyTransform(tls, world, (*b2Shape)(unsafe.Pointer(otherShape)).BodyId)
	(*(*DistanceInput)(unsafe.Pointer(bp))).ProxyA = b2MakeShapeDistanceProxy(tls, sensorShape)
	(*(*DistanceInput)(unsafe.Pointer(bp))).ProxyB = b2MakeShapeDistanceProxy(tls, otherShape)
	(*(*DistanceInput)(unsafe.Pointer(bp))).TransformA = (*b2SensorQueryContext)(unsafe.Pointer(queryContext)).Transform
	(*(*DistanceInput)(unsafe.Pointer(bp))).TransformB = otherTransform
	(*(*DistanceInput)(unsafe.Pointer(bp))).UseRadii = boolUint8(true1 != 0)
	*(*SimplexCache)(unsafe.Pointer(bp + 184)) = SimplexCache{}
	output = b2ShapeDistance(tls, bp, bp+184, uintptrFromInt32(0), 0)
	overlaps = boolUint8(output.Distance < float32(float32FromFloat32(10)*float32FromFloat32(1.1920928955078125e-07)))
	if int32FromUint8(overlaps) == false1 {
		return boolUint8(true1 != 0)
	}
	// Record the overlap
	sensor = (*b2SensorQueryContext)(unsafe.Pointer(queryContext)).Sensor
	v9 = sensor + 16
	if (*b2ShapeRefArray)(unsafe.Pointer(v9)).Count == (*b2ShapeRefArray)(unsafe.Pointer(v9)).Capacity {
		if (*b2ShapeRefArray)(unsafe.Pointer(v9)).Capacity < int32(2) {
			v10 = int32(2)
		} else {
			v10 = (*b2ShapeRefArray)(unsafe.Pointer(v9)).Capacity + (*b2ShapeRefArray)(unsafe.Pointer(v9)).Capacity>>int32(1)
		}
		newCapacity = v10
		b2ShapeRefArray_Reserve(tls, v9, newCapacity)
	}
	*(*int32)(unsafe.Pointer(v9 + 8)) += int32(1)
	v11 = (*b2ShapeRefArray)(unsafe.Pointer(v9)).Data + uintptr((*b2ShapeRefArray)(unsafe.Pointer(v9)).Count-int32FromInt32(1))*8
	goto _12
_12:
	shapeRef = v11
	(*b2ShapeRef)(unsafe.Pointer(shapeRef)).ShapeId = shapeId
	(*b2ShapeRef)(unsafe.Pointer(shapeRef)).Generation = (*b2Shape)(unsafe.Pointer(otherShape)).Generation
	return boolUint8(true1 != 0)
}

func b2SensorTask(tls *_Stack, startIndex int32, endIndex int32, threadIndex uint32, context uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var blockIndex, v15, v17, v20 uint32_t
	var body, s1, s2, sensor, sensorShape, taskContext, trees, world, v10, v12, v14, v16, v19, v2, v4, v6, v8 uintptr
	var count1, count2, i, sensorIndex, v11, v3, v7 int32
	var queryBounds AABB
	var temp b2ShapeRefArray
	var transform Transform
	var _ /* queryContext at bp+0 */ b2SensorQueryContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = blockIndex, body, count1, count2, i, queryBounds, s1, s2, sensor, sensorIndex, sensorShape, taskContext, temp, transform, trees, world, v10, v11, v12, v14, v15, v16, v17, v19, v2, v20, v3, v4, v6, v7, v8
	world = context
	if !(int32FromUint32(threadIndex) < (*b2World)(unsafe.Pointer(world)).WorkerCount) && b2InternalAssertFcn(tls, __ccgo_ts+10751, __ccgo_ts+10789, int32FromInt32(140)) != 0 {
		__builtin_trap(tls)
	}
	taskContext = (*b2World)(unsafe.Pointer(world)).SensorTaskContexts.Data + uintptr(threadIndex)*16
	if !(startIndex < endIndex) && b2InternalAssertFcn(tls, __ccgo_ts+10813, __ccgo_ts+10789, int32FromInt32(143)) != 0 {
		__builtin_trap(tls)
	}
	trees = world + 40
	sensorIndex = startIndex
	for {
		if !(sensorIndex < endIndex) {
			break
		}
		v2 = world + 1280
		v3 = sensorIndex
		if !(0 <= v3 && v3 < (*b2SensorArray)(unsafe.Pointer(v2)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+10835, int32FromInt32(35)) != 0 {
			__builtin_trap(tls)
		}
		v4 = (*b2SensorArray)(unsafe.Pointer(v2)).Data + uintptr(v3)*40
		goto _5
	_5:
		sensor = v4
		v6 = world + 1248
		v7 = (*b2Sensor)(unsafe.Pointer(sensor)).ShapeId
		if !(0 <= v7 && v7 < (*b2ShapeArray)(unsafe.Pointer(v6)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v8 = (*b2ShapeArray)(unsafe.Pointer(v6)).Data + uintptr(v7)*288
		goto _9
	_9:
		sensorShape = v8
		// swap overlap arrays
		temp = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1
		(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2
		(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2 = temp
		(*b2ShapeRefArray)(unsafe.Pointer(sensor + 16)).Count = 0
		v10 = world + 1024
		v11 = (*b2Shape)(unsafe.Pointer(sensorShape)).BodyId
		if !(0 <= v11 && v11 < (*b2BodyArray)(unsafe.Pointer(v10)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+380, int32FromInt32(192)) != 0 {
			__builtin_trap(tls)
		}
		v12 = (*b2BodyArray)(unsafe.Pointer(v10)).Data + uintptr(v11)*128
		goto _13
	_13:
		body = v12
		if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_disabledSet) || int32FromUint8((*b2Shape)(unsafe.Pointer(sensorShape)).EnableSensorEvents) == false1 {
			if (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Count != 0 {
				v14 = taskContext
				v15 = uint32FromInt32(sensorIndex)
				blockIndex = v15 / uint32(64)
				if !(blockIndex < (*b2BitSet)(unsafe.Pointer(v14)).BlockCount) && b2InternalAssertFcn(tls, __ccgo_ts+10859, __ccgo_ts+10891, int32FromInt32(28)) != 0 {
					__builtin_trap(tls)
				}
				*(*uint64_t)(unsafe.Pointer((*b2BitSet)(unsafe.Pointer(v14)).Bits + uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v15 % uint32FromInt32(64))
			}
			goto _1
		}
		transform = b2GetBodyTransformQuick(tls, world, body)
		*(*b2SensorQueryContext)(unsafe.Pointer(bp)) = b2SensorQueryContext{
			World:       world,
			TaskContext: taskContext,
			Sensor:      sensor,
			SensorShape: sensorShape,
			Transform:   transform,
		}
		if !((*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex == sensorIndex) && b2InternalAssertFcn(tls, __ccgo_ts+10915, __ccgo_ts+10789, int32FromInt32(177)) != 0 {
			__builtin_trap(tls)
		}
		queryBounds = (*b2Shape)(unsafe.Pointer(sensorShape)).Aabb
		// Query all trees
		b2DynamicTree_Query(tls, trees+uintptr(0)*72, queryBounds, (*b2Shape)(unsafe.Pointer(sensorShape)).Filter.MaskBits, __ccgo_fp(b2SensorQueryCallback), bp)
		b2DynamicTree_Query(tls, trees+uintptr(1)*72, queryBounds, (*b2Shape)(unsafe.Pointer(sensorShape)).Filter.MaskBits, __ccgo_fp(b2SensorQueryCallback), bp)
		b2DynamicTree_Query(tls, trees+uintptr(2)*72, queryBounds, (*b2Shape)(unsafe.Pointer(sensorShape)).Filter.MaskBits, __ccgo_fp(b2SensorQueryCallback), bp)
		// Sort the overlaps to enable finding begin and end events.
		qsort(tls, (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data, uint64FromInt32((*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count), uint64(8), b2CompareShapeRefs)
		count1 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Count
		count2 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count
		if count1 != count2 {
			// something changed
			v16 = taskContext
			v17 = uint32FromInt32(sensorIndex)
			blockIndex = v17 / uint32(64)
			if !(blockIndex < (*b2BitSet)(unsafe.Pointer(v16)).BlockCount) && b2InternalAssertFcn(tls, __ccgo_ts+10859, __ccgo_ts+10891, int32FromInt32(28)) != 0 {
				__builtin_trap(tls)
			}
			*(*uint64_t)(unsafe.Pointer((*b2BitSet)(unsafe.Pointer(v16)).Bits + uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v17 % uint32FromInt32(64))
		} else {
			i = 0
			for {
				if !(i < count1) {
					break
				}
				s1 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Data + uintptr(i)*8
				s2 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data + uintptr(i)*8
				if (*b2ShapeRef)(unsafe.Pointer(s1)).ShapeId != (*b2ShapeRef)(unsafe.Pointer(s2)).ShapeId || int32FromUint16((*b2ShapeRef)(unsafe.Pointer(s1)).Generation) != int32FromUint16((*b2ShapeRef)(unsafe.Pointer(s2)).Generation) {
					// something changed
					v19 = taskContext
					v20 = uint32FromInt32(sensorIndex)
					blockIndex = v20 / uint32(64)
					if !(blockIndex < (*b2BitSet)(unsafe.Pointer(v19)).BlockCount) && b2InternalAssertFcn(tls, __ccgo_ts+10859, __ccgo_ts+10891, int32FromInt32(28)) != 0 {
						__builtin_trap(tls)
					}
					*(*uint64_t)(unsafe.Pointer((*b2BitSet)(unsafe.Pointer(v19)).Bits + uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v20 % uint32FromInt32(64))
					break
				}
				goto _18
			_18:
				;
				i = i + 1
			}
		}
		goto _1
	_1:
		;
		sensorIndex = sensorIndex + 1
	}
}

func b2OverlapSensors(tls *_Stack, world uintptr) {
	var bitSet, bits, r1, r11, r2, r21, refs1, refs2, sensor, sensorShape, userSensorTask, v10, v12, v14, v16, v18, v20, v22, v24, v6, v8 uintptr
	var blockCount, ctz, k, v4 uint32_t
	var count1, count2, i, i1, index11, index2, minRange, newCapacity, newCapacity1, sensorCount, sensorIndex, v11, v15, v17, v19, v21, v23, v25, v7 int32
	var event, event2, event4 SensorEndTouchEvent
	var event1, event3, event5 SensorBeginTouchEvent
	var sensorId, visitorId, visitorId1, visitorId2, visitorId3, visitorId4, visitorId5 ShapeId
	var word uint64_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bitSet, bits, blockCount, count1, count2, ctz, event, event1, event2, event3, event4, event5, i, i1, index11, index2, k, minRange, newCapacity, newCapacity1, r1, r11, r2, r21, refs1, refs2, sensor, sensorCount, sensorId, sensorIndex, sensorShape, userSensorTask, visitorId, visitorId1, visitorId2, visitorId3, visitorId4, visitorId5, word, v10, v11, v12, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v4, v6, v7, v8
	sensorCount = (*b2World)(unsafe.Pointer(world)).Sensors.Count
	if sensorCount == 0 {
		return
	}
	if !((*b2World)(unsafe.Pointer(world)).WorkerCount > int32FromInt32(0)) && b2InternalAssertFcn(tls, __ccgo_ts+10955, __ccgo_ts+10789, int32FromInt32(223)) != 0 {
		__builtin_trap(tls)
	}
	i = 0
	for {
		if !(i < (*b2World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		b2SetBitCountAndClear(tls, (*b2World)(unsafe.Pointer(world)).SensorTaskContexts.Data+uintptr(i)*16, uint32FromInt32(sensorCount))
		goto _1
	_1:
		;
		i = i + 1
	}
	// Parallel-for sensors overlaps
	minRange = int32(16)
	userSensorTask = (*(*func(*_Stack, uintptr, int32, int32, uintptr, uintptr) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).EnqueueTaskFcn})))(tls, __ccgo_fp(b2SensorTask), sensorCount, minRange, world, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
	*(*int32)(unsafe.Pointer(world + 1776)) += int32(1)
	if userSensorTask != uintptrFromInt32(0) {
		(*(*func(*_Stack, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b2World)(unsafe.Pointer(world)).FinishTaskFcn})))(tls, userSensorTask, (*b2World)(unsafe.Pointer(world)).UserTaskContext)
	}
	bitSet = (*b2World)(unsafe.Pointer(world)).SensorTaskContexts.Data
	i1 = int32(1)
	for {
		if !(i1 < (*b2World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		b2InPlaceUnion(tls, bitSet, (*b2World)(unsafe.Pointer(world)).SensorTaskContexts.Data+uintptr(i1)*16)
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
	// Iterate sensors bits and publish events
	// Process contact state changes. Iterate over set bits
	bits = (*b2BitSet)(unsafe.Pointer(bitSet)).Bits
	blockCount = (*b2BitSet)(unsafe.Pointer(bitSet)).BlockCount
	k = uint32(0)
	for {
		if !(k < blockCount) {
			break
		}
		word = *(*uint64_t)(unsafe.Pointer(bits + uintptr(k)*8))
		for word != uint64(0) {
			v4 = uint32FromInt32(__builtin_ctzll(tls, word))
			goto _5
		_5:
			ctz = v4
			sensorIndex = int32FromUint32(uint32FromInt32(64)*k + ctz)
			v6 = world + 1280
			v7 = sensorIndex
			if !(0 <= v7 && v7 < (*b2SensorArray)(unsafe.Pointer(v6)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+10835, int32FromInt32(35)) != 0 {
				__builtin_trap(tls)
			}
			v8 = (*b2SensorArray)(unsafe.Pointer(v6)).Data + uintptr(v7)*40
			goto _9
		_9:
			sensor = v8
			v10 = world + 1248
			v11 = (*b2Sensor)(unsafe.Pointer(sensor)).ShapeId
			if !(0 <= v11 && v11 < (*b2ShapeArray)(unsafe.Pointer(v10)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
				__builtin_trap(tls)
			}
			v12 = (*b2ShapeArray)(unsafe.Pointer(v10)).Data + uintptr(v11)*288
			goto _13
		_13:
			sensorShape = v12
			sensorId = ShapeId{
				Index1:     (*b2Sensor)(unsafe.Pointer(sensor)).ShapeId + int32(1),
				World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
				Generation: (*b2Shape)(unsafe.Pointer(sensorShape)).Generation,
			}
			count1 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Count
			count2 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count
			refs1 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Data
			refs2 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data
			// overlaps1 can have overlaps that end
			// overlaps2 can have overlaps that begin
			index11 = 0
			index2 = 0
			for index11 < count1 && index2 < count2 {
				r1 = refs1 + uintptr(index11)*8
				r2 = refs2 + uintptr(index2)*8
				if (*b2ShapeRef)(unsafe.Pointer(r1)).ShapeId == (*b2ShapeRef)(unsafe.Pointer(r2)).ShapeId {
					if int32FromUint16((*b2ShapeRef)(unsafe.Pointer(r1)).Generation) < int32FromUint16((*b2ShapeRef)(unsafe.Pointer(r2)).Generation) {
						// end
						visitorId = ShapeId{
							Index1:     (*b2ShapeRef)(unsafe.Pointer(r1)).ShapeId + int32(1),
							World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
							Generation: (*b2ShapeRef)(unsafe.Pointer(r1)).Generation,
						}
						event = SensorEndTouchEvent{
							SensorShapeId:  sensorId,
							VisitorShapeId: visitorId,
						}
						v14 = world + 1376 + uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16
						if (*b2SensorEndTouchEventArray)(unsafe.Pointer(v14)).Count == (*b2SensorEndTouchEventArray)(unsafe.Pointer(v14)).Capacity {
							if (*b2SensorEndTouchEventArray)(unsafe.Pointer(v14)).Capacity < int32(2) {
								v15 = int32(2)
							} else {
								v15 = (*b2SensorEndTouchEventArray)(unsafe.Pointer(v14)).Capacity + (*b2SensorEndTouchEventArray)(unsafe.Pointer(v14)).Capacity>>int32(1)
							}
							newCapacity1 = v15
							b2SensorEndTouchEventArray_Reserve(tls, v14, newCapacity1)
						}
						*(*SensorEndTouchEvent)(unsafe.Pointer((*b2SensorEndTouchEventArray)(unsafe.Pointer(v14)).Data + uintptr((*b2SensorEndTouchEventArray)(unsafe.Pointer(v14)).Count)*16)) = event
						*(*int32)(unsafe.Pointer(v14 + 8)) += int32(1)
						index11 = index11 + int32(1)
					} else {
						if int32FromUint16((*b2ShapeRef)(unsafe.Pointer(r1)).Generation) > int32FromUint16((*b2ShapeRef)(unsafe.Pointer(r2)).Generation) {
							// begin
							visitorId1 = ShapeId{
								Index1:     (*b2ShapeRef)(unsafe.Pointer(r2)).ShapeId + int32(1),
								World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
								Generation: (*b2ShapeRef)(unsafe.Pointer(r2)).Generation,
							}
							event1 = SensorBeginTouchEvent{
								SensorShapeId:  sensorId,
								VisitorShapeId: visitorId1,
							}
							v16 = world + 1344
							if (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v16)).Count == (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v16)).Capacity {
								if (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v16)).Capacity < int32(2) {
									v17 = int32(2)
								} else {
									v17 = (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v16)).Capacity + (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v16)).Capacity>>int32(1)
								}
								newCapacity = v17
								b2SensorBeginTouchEventArray_Reserve(tls, v16, newCapacity)
							}
							*(*SensorBeginTouchEvent)(unsafe.Pointer((*b2SensorBeginTouchEventArray)(unsafe.Pointer(v16)).Data + uintptr((*b2SensorBeginTouchEventArray)(unsafe.Pointer(v16)).Count)*16)) = event1
							*(*int32)(unsafe.Pointer(v16 + 8)) += int32(1)
							index2 = index2 + int32(1)
						} else {
							// persisted
							index11 = index11 + int32(1)
							index2 = index2 + int32(1)
						}
					}
				} else {
					if (*b2ShapeRef)(unsafe.Pointer(r1)).ShapeId < (*b2ShapeRef)(unsafe.Pointer(r2)).ShapeId {
						// end
						visitorId2 = ShapeId{
							Index1:     (*b2ShapeRef)(unsafe.Pointer(r1)).ShapeId + int32(1),
							World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
							Generation: (*b2ShapeRef)(unsafe.Pointer(r1)).Generation,
						}
						event2 = SensorEndTouchEvent{
							SensorShapeId:  sensorId,
							VisitorShapeId: visitorId2,
						}
						v18 = world + 1376 + uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16
						if (*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Count == (*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Capacity {
							if (*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Capacity < int32(2) {
								v19 = int32(2)
							} else {
								v19 = (*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Capacity + (*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Capacity>>int32(1)
							}
							newCapacity1 = v19
							b2SensorEndTouchEventArray_Reserve(tls, v18, newCapacity1)
						}
						*(*SensorEndTouchEvent)(unsafe.Pointer((*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Data + uintptr((*b2SensorEndTouchEventArray)(unsafe.Pointer(v18)).Count)*16)) = event2
						*(*int32)(unsafe.Pointer(v18 + 8)) += int32(1)
						index11 = index11 + int32(1)
					} else {
						// begin
						visitorId3 = ShapeId{
							Index1:     (*b2ShapeRef)(unsafe.Pointer(r2)).ShapeId + int32(1),
							World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
							Generation: (*b2ShapeRef)(unsafe.Pointer(r2)).Generation,
						}
						event3 = SensorBeginTouchEvent{
							SensorShapeId:  sensorId,
							VisitorShapeId: visitorId3,
						}
						v20 = world + 1344
						if (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v20)).Count == (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v20)).Capacity {
							if (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v20)).Capacity < int32(2) {
								v21 = int32(2)
							} else {
								v21 = (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v20)).Capacity + (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v20)).Capacity>>int32(1)
							}
							newCapacity = v21
							b2SensorBeginTouchEventArray_Reserve(tls, v20, newCapacity)
						}
						*(*SensorBeginTouchEvent)(unsafe.Pointer((*b2SensorBeginTouchEventArray)(unsafe.Pointer(v20)).Data + uintptr((*b2SensorBeginTouchEventArray)(unsafe.Pointer(v20)).Count)*16)) = event3
						*(*int32)(unsafe.Pointer(v20 + 8)) += int32(1)
						index2 = index2 + int32(1)
					}
				}
			}
			for index11 < count1 {
				// end
				r11 = refs1 + uintptr(index11)*8
				visitorId4 = ShapeId{
					Index1:     (*b2ShapeRef)(unsafe.Pointer(r11)).ShapeId + int32(1),
					World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
					Generation: (*b2ShapeRef)(unsafe.Pointer(r11)).Generation,
				}
				event4 = SensorEndTouchEvent{
					SensorShapeId:  sensorId,
					VisitorShapeId: visitorId4,
				}
				v22 = world + 1376 + uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16
				if (*b2SensorEndTouchEventArray)(unsafe.Pointer(v22)).Count == (*b2SensorEndTouchEventArray)(unsafe.Pointer(v22)).Capacity {
					if (*b2SensorEndTouchEventArray)(unsafe.Pointer(v22)).Capacity < int32(2) {
						v23 = int32(2)
					} else {
						v23 = (*b2SensorEndTouchEventArray)(unsafe.Pointer(v22)).Capacity + (*b2SensorEndTouchEventArray)(unsafe.Pointer(v22)).Capacity>>int32(1)
					}
					newCapacity1 = v23
					b2SensorEndTouchEventArray_Reserve(tls, v22, newCapacity1)
				}
				*(*SensorEndTouchEvent)(unsafe.Pointer((*b2SensorEndTouchEventArray)(unsafe.Pointer(v22)).Data + uintptr((*b2SensorEndTouchEventArray)(unsafe.Pointer(v22)).Count)*16)) = event4
				*(*int32)(unsafe.Pointer(v22 + 8)) += int32(1)
				index11 = index11 + int32(1)
			}
			for index2 < count2 {
				// begin
				r21 = refs2 + uintptr(index2)*8
				visitorId5 = ShapeId{
					Index1:     (*b2ShapeRef)(unsafe.Pointer(r21)).ShapeId + int32(1),
					World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
					Generation: (*b2ShapeRef)(unsafe.Pointer(r21)).Generation,
				}
				event5 = SensorBeginTouchEvent{
					SensorShapeId:  sensorId,
					VisitorShapeId: visitorId5,
				}
				v24 = world + 1344
				if (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v24)).Count == (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v24)).Capacity {
					if (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v24)).Capacity < int32(2) {
						v25 = int32(2)
					} else {
						v25 = (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v24)).Capacity + (*b2SensorBeginTouchEventArray)(unsafe.Pointer(v24)).Capacity>>int32(1)
					}
					newCapacity = v25
					b2SensorBeginTouchEventArray_Reserve(tls, v24, newCapacity)
				}
				*(*SensorBeginTouchEvent)(unsafe.Pointer((*b2SensorBeginTouchEventArray)(unsafe.Pointer(v24)).Data + uintptr((*b2SensorBeginTouchEventArray)(unsafe.Pointer(v24)).Count)*16)) = event5
				*(*int32)(unsafe.Pointer(v24 + 8)) += int32(1)
				index2 = index2 + int32(1)
			}
			// Clear the smallest set bit
			word = word & (word - uint64(1))
		}
		goto _3
	_3:
		;
		k = k + 1
	}
}

func b2DestroySensor(tls *_Stack, world uintptr, sensorShape uintptr) {
	var event SensorEndTouchEvent
	var i, movedIndex, movedIndex1, newCapacity, v10, v13, v17, v2, v7, v9 int32
	var movedSensor, otherSensorShape, ref, sensor, v1, v12, v14, v16, v18, v3, v6, v8 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = event, i, movedIndex, movedIndex1, movedSensor, newCapacity, otherSensorShape, ref, sensor, v1, v10, v12, v13, v14, v16, v17, v18, v2, v3, v6, v7, v8, v9
	v1 = world + 1280
	v2 = (*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex
	if !(0 <= v2 && v2 < (*b2SensorArray)(unsafe.Pointer(v1)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+10835, int32FromInt32(35)) != 0 {
		__builtin_trap(tls)
	}
	v3 = (*b2SensorArray)(unsafe.Pointer(v1)).Data + uintptr(v2)*40
	goto _4
_4:
	sensor = v3
	i = 0
	for {
		if !(i < (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count) {
			break
		}
		ref = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data + uintptr(i)*8
		event = SensorEndTouchEvent{
			SensorShapeId: ShapeId{
				Index1:     (*b2Shape)(unsafe.Pointer(sensorShape)).Id + int32(1),
				World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
				Generation: (*b2Shape)(unsafe.Pointer(sensorShape)).Generation,
			},
			VisitorShapeId: ShapeId{
				Index1:     (*b2ShapeRef)(unsafe.Pointer(ref)).ShapeId + int32(1),
				World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
				Generation: (*b2ShapeRef)(unsafe.Pointer(ref)).Generation,
			},
		}
		v6 = world + 1376 + uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16
		if (*b2SensorEndTouchEventArray)(unsafe.Pointer(v6)).Count == (*b2SensorEndTouchEventArray)(unsafe.Pointer(v6)).Capacity {
			if (*b2SensorEndTouchEventArray)(unsafe.Pointer(v6)).Capacity < int32(2) {
				v7 = int32(2)
			} else {
				v7 = (*b2SensorEndTouchEventArray)(unsafe.Pointer(v6)).Capacity + (*b2SensorEndTouchEventArray)(unsafe.Pointer(v6)).Capacity>>int32(1)
			}
			newCapacity = v7
			b2SensorEndTouchEventArray_Reserve(tls, v6, newCapacity)
		}
		*(*SensorEndTouchEvent)(unsafe.Pointer((*b2SensorEndTouchEventArray)(unsafe.Pointer(v6)).Data + uintptr((*b2SensorEndTouchEventArray)(unsafe.Pointer(v6)).Count)*16)) = event
		*(*int32)(unsafe.Pointer(v6 + 8)) += int32(1)
		goto _5
	_5:
		;
		i = i + 1
	}
	// Destroy sensor
	b2ShapeRefArray_Destroy(tls, sensor)
	b2ShapeRefArray_Destroy(tls, sensor+16)
	v8 = world + 1280
	v9 = (*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex
	if !(0 <= v9 && v9 < (*b2SensorArray)(unsafe.Pointer(v8)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+10835, int32FromInt32(35)) != 0 {
		__builtin_trap(tls)
	}
	movedIndex = -int32(1)
	if v9 != (*b2SensorArray)(unsafe.Pointer(v8)).Count-int32FromInt32(1) {
		movedIndex = (*b2SensorArray)(unsafe.Pointer(v8)).Count - int32(1)
		*(*b2Sensor)(unsafe.Pointer((*b2SensorArray)(unsafe.Pointer(v8)).Data + uintptr(v9)*40)) = *(*b2Sensor)(unsafe.Pointer((*b2SensorArray)(unsafe.Pointer(v8)).Data + uintptr(movedIndex)*40))
	}
	*(*int32)(unsafe.Pointer(v8 + 8)) -= int32(1)
	v10 = movedIndex
	goto _11
_11:
	movedIndex1 = v10
	if movedIndex1 != -int32(1) {
		v12 = world + 1280
		v13 = (*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex
		if !(0 <= v13 && v13 < (*b2SensorArray)(unsafe.Pointer(v12)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+10835, int32FromInt32(35)) != 0 {
			__builtin_trap(tls)
		}
		v14 = (*b2SensorArray)(unsafe.Pointer(v12)).Data + uintptr(v13)*40
		goto _15
	_15:
		// Fixup moved sensor
		movedSensor = v14
		v16 = world + 1248
		v17 = (*b2Sensor)(unsafe.Pointer(movedSensor)).ShapeId
		if !(0 <= v17 && v17 < (*b2ShapeArray)(unsafe.Pointer(v16)).Count) && b2InternalAssertFcn(tls, __ccgo_ts+349, __ccgo_ts+1384, int32FromInt32(138)) != 0 {
			__builtin_trap(tls)
		}
		v18 = (*b2ShapeArray)(unsafe.Pointer(v16)).Data + uintptr(v17)*288
		goto _19
	_19:
		otherSensorShape = v18
		(*b2Shape)(unsafe.Pointer(otherSensorShape)).SensorIndex = (*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex
	}
}

func b2SensorBeginTouchEventArray_Create(tls *_Stack, capacity int32) (r b2SensorBeginTouchEventArray) {
	var a b2SensorBeginTouchEventArray
	_ = a
	a = b2SensorBeginTouchEventArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(16)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2SensorBeginTouchEventArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2SensorBeginTouchEventArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2SensorBeginTouchEventArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2SensorBeginTouchEventArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2SensorBeginTouchEventArray)(unsafe.Pointer(a)).Capacity)*uint64(16)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(16)))
	(*b2SensorBeginTouchEventArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2SensorBeginTouchEventArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2SensorBeginTouchEventArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2SensorBeginTouchEventArray)(unsafe.Pointer(a)).Capacity)*uint64(16)))
	(*b2SensorBeginTouchEventArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2SensorBeginTouchEventArray)(unsafe.Pointer(a)).Count = 0
	(*b2SensorBeginTouchEventArray)(unsafe.Pointer(a)).Capacity = 0
}

func b2SensorEndTouchEventArray_Create(tls *_Stack, capacity int32) (r b2SensorEndTouchEventArray) {
	var a b2SensorEndTouchEventArray
	_ = a
	a = b2SensorEndTouchEventArray{}
	if capacity > 0 {
		a.Data = b2Alloc(tls, int32FromUint64(uint64FromInt32(capacity)*uint64(16)))
		a.Capacity = capacity
	}
	return a
}

/* Reserve */
func b2SensorEndTouchEventArray_Reserve(tls *_Stack, a uintptr, newCapacity int32) {
	if newCapacity <= (*b2SensorEndTouchEventArray)(unsafe.Pointer(a)).Capacity {
		return
	}
	(*b2SensorEndTouchEventArray)(unsafe.Pointer(a)).Data = b2GrowAlloc(tls, (*b2SensorEndTouchEventArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2SensorEndTouchEventArray)(unsafe.Pointer(a)).Capacity)*uint64(16)), int32FromUint64(uint64FromInt32(newCapacity)*uint64(16)))
	(*b2SensorEndTouchEventArray)(unsafe.Pointer(a)).Capacity = newCapacity
}

/* Destroy */
func b2SensorEndTouchEventArray_Destroy(tls *_Stack, a uintptr) {
	b2Free(tls, (*b2SensorEndTouchEventArray)(unsafe.Pointer(a)).Data, int32FromUint64(uint64FromInt32((*b2SensorEndTouchEventArray)(unsafe.Pointer(a)).Capacity)*uint64(16)))
	(*b2SensorEndTouchEventArray)(unsafe.Pointer(a)).Data = uintptrFromInt32(0)
	(*b2SensorEndTouchEventArray)(unsafe.Pointer(a)).Count = 0
	(*b2SensorEndTouchEventArray)(unsafe.Pointer(a)).Capacity = 0
}
