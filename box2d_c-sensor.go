package b2

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

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

func b2SensorQueryCallback(tls *_Stack, proxyId int32, userData uint64, context uintptr) (r1 uint8) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var C, v16 Transform
	var customFilterFcn, otherShape, queryContext, sensor, sensorShape, shapeRef, world, v18, v19, v20, v23, v26 uintptr
	var d, v13, v14 Vec2
	var idA, idB ShapeId
	var newCapacity, newSize, oldCapacity, oldSize, sensorShapeId, shapeId, v21, v22, v25 int32
	var otherTransform, v6, v7 WorldTransform
	var output DistanceOutput
	var overlaps, shouldCollide, v4 uint8
	var r, v10, v12, v8, v9 Rot
	var v1 bool
	var v2, v3 Filter
	var _ /* cache at bp+168 */ SimplexCache
	var _ /* input at bp+0 */ DistanceInput
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, customFilterFcn, d, idA, idB, newCapacity, newSize, oldCapacity, oldSize, otherShape, otherTransform, output, overlaps, queryContext, r, sensor, sensorShape, sensorShapeId, shapeId, shapeRef, shouldCollide, world, v1, v10, v12, v13, v14, v16, v18, v19, v2, v20, v21, v22, v23, v25, v26, v3, v4, v6, v7, v8, v9
	_ = uint64FromInt64(4)
	shapeId = int32FromUint64(userData)
	queryContext = context
	sensorShape = (*b2SensorQueryContext)(unsafe.Pointer(queryContext)).SensorShape
	sensorShapeId = (*b2Shape)(unsafe.Pointer(sensorShape)).Id
	if shapeId == sensorShapeId {
		return boolUint8(true1 != 0)
	}
	world = (*b2SensorQueryContext)(unsafe.Pointer(queryContext)).World
	if v1 = !!(0 <= shapeId && shapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+1908, __ccgo_ts+21483, int32FromInt32(60))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	otherShape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*280
	// Are sensor events enabled on the other shape?
	if int32FromUint8((*b2Shape)(unsafe.Pointer(otherShape)).EnableSensorEvents) == false1 {
		return boolUint8(true1 != 0)
	}
	// Skip shapes on the same body
	if (*b2Shape)(unsafe.Pointer(otherShape)).BodyId == (*b2Shape)(unsafe.Pointer(sensorShape)).BodyId {
		return boolUint8(true1 != 0)
	}
	// Check filter
	v2 = (*b2Shape)(unsafe.Pointer(sensorShape)).Filter
	v3 = (*b2Shape)(unsafe.Pointer(otherShape)).Filter
	if v2.GroupIndex == v3.GroupIndex && v2.GroupIndex != 0 {
		v4 = boolUint8(v2.GroupIndex > 0)
		goto _5
	}
	v4 = boolUint8(v2.MaskBits&v3.CategoryBits != uint64(0) && v2.CategoryBits&v3.MaskBits != uint64(0))
	goto _5
_5:
	if int32FromUint8(v4) == false1 {
		return boolUint8(true1 != 0)
	}
	// Custom user filter
	if (*b2Shape)(unsafe.Pointer(sensorShape)).EnableCustomFiltering != 0 || (*b2Shape)(unsafe.Pointer(otherShape)).EnableCustomFiltering != 0 {
		customFilterFcn = (*b2World)(unsafe.Pointer((*b2SensorQueryContext)(unsafe.Pointer(queryContext)).World)).CustomFilterFcn
		if customFilterFcn != uintptrFromInt32(0) {
			idA = ShapeId{
				Index1:     sensorShapeId + int32(1),
				World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
				Generation: (*b2Shape)(unsafe.Pointer(sensorShape)).Generation,
			}
			idB = ShapeId{
				Index1:     shapeId + int32(1),
				World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
				Generation: (*b2Shape)(unsafe.Pointer(otherShape)).Generation,
			}
			shouldCollide = (*(*func(*_Stack, ShapeId, ShapeId, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{customFilterFcn})))(tls, idA, idB, (*b2World)(unsafe.Pointer((*b2SensorQueryContext)(unsafe.Pointer(queryContext)).World)).CustomFilterContext)
			if int32FromUint8(shouldCollide) == false1 {
				return boolUint8(true1 != 0)
			}
		}
	}
	// The relative pose is differenced in double so sensor overlap stays exact far from the origin
	otherTransform = b2GetBodyTransform(tls, world, (*b2Shape)(unsafe.Pointer(otherShape)).BodyId)
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyA = b2MakeShapeDistanceProxy(tls, sensorShape)
	(*(*DistanceInput)(unsafe.Add(nil, bp))).ProxyB = b2MakeShapeDistanceProxy(tls, otherShape)
	v6 = (*b2SensorQueryContext)(unsafe.Pointer(queryContext)).Transform
	v7 = otherTransform
	v8 = v6.Q
	v9 = v7.Q
	r.S = float32(v8.C*v9.S) - float32(v8.S*v9.C)
	r.C = float32(v8.C*v9.C) + float32(v8.S*v9.S)
	v10 = r
	goto _11
_11:
	C.Q = v10
	d = Vec2{
		X: v7.P.X - v6.P.X,
		Y: v7.P.Y - v6.P.Y,
	}
	v12 = v6.Q
	v13 = d
	v14 = Vec2{
		X: float32(v12.C*v13.X) + float32(v12.S*v13.Y),
		Y: float32(-v12.S*v13.X) + float32(v12.C*v13.Y),
	}
	goto _15
_15:
	C.P = v14
	v16 = C
	goto _17
_17:
	(*(*DistanceInput)(unsafe.Add(nil, bp))).Transform = v16
	(*(*DistanceInput)(unsafe.Add(nil, bp))).UseRadii = boolUint8(true1 != 0)
	*(*SimplexCache)(unsafe.Add(nil, bp+168)) = SimplexCache{}
	output = b2ComputeShapeDistance(tls, bp, bp+168, uintptrFromInt32(0), 0)
	overlaps = boolUint8(output.Distance < float32(float32FromFloat32(10)*float32FromFloat32(1.1920928955078125e-07)))
	if int32FromUint8(overlaps) == false1 {
		return boolUint8(true1 != 0)
	}
	// Record the overlap
	sensor = (*b2SensorQueryContext)(unsafe.Pointer(queryContext)).Sensor
	v18 = sensor + 32
	v19 = sensor + 32 + 8
	v20 = sensor + 32 + 12
	v21 = int32(8)
	if *(*int32)(unsafe.Add(nil, v19)) >= *(*int32)(unsafe.Add(nil, v20)) {
		oldCapacity = *(*int32)(unsafe.Add(nil, v20))
		oldSize = oldCapacity * v21
		if oldCapacity == 0 {
			v22 = int32(8)
		} else {
			v22 = int32(2) * oldCapacity
		}
		newCapacity = v22
		newSize = newCapacity * v21
		*(*uintptr)(unsafe.Add(nil, v18)) = b2GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v18)), uint64FromInt32(oldSize), uint64FromInt32(newSize))
		*(*int32)(unsafe.Add(nil, v20)) = newCapacity
	}
	v26 = v19
	v25 = *(*int32)(unsafe.Pointer(v26))
	*(*int32)(unsafe.Pointer(v26)) = *(*int32)(unsafe.Pointer(v26)) + 1
	v23 = *(*uintptr)(unsafe.Add(nil, v18)) + uintptr(v25*v21)
	goto _24
_24:
	shapeRef = v23
	(*b2Visitor)(unsafe.Pointer(shapeRef)).ShapeId = shapeId
	(*b2Visitor)(unsafe.Pointer(shapeRef)).Generation = (*b2Shape)(unsafe.Pointer(otherShape)).Generation
	return boolUint8(true1 != 0)
}

func b2SensorTask(tls *_Stack, startIndex int32, endIndex int32, threadIndex int32, context uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var blockIndex, v9 uint32_t
	var body, overlapData, s1, s2, sensor, sensorShape, taskContext, trees, world, v8 uintptr
	var count1, count2, elementSize, hitCount, i, i1, i2, newCapacity, newSize, oldSize, overlapCount, sensorIndex, uniqueCount, v6 int32
	var queryBounds AABB
	var temp b2DynamicArray_b2Visitor
	var transform WorldTransform
	var v1, v3 bool
	var _ /* queryContext at bp+0 */ b2SensorQueryContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = blockIndex, body, count1, count2, elementSize, hitCount, i, i1, i2, newCapacity, newSize, oldSize, overlapCount, overlapData, queryBounds, s1, s2, sensor, sensorIndex, sensorShape, taskContext, temp, transform, trees, uniqueCount, world, v1, v3, v6, v8, v9
	world = context
	taskContext = (*b2World)(unsafe.Pointer(world)).SensorTaskContexts.Data + uintptr(threadIndex)*16
	if v1 = !!(startIndex < endIndex); !v1 {
		b2InternalAssert(tls, __ccgo_ts+17428, __ccgo_ts+21483, int32FromInt32(142))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	trees = world + 40
	sensorIndex = startIndex
	for {
		if !(sensorIndex < endIndex) {
			break
		}
		if v1 = !!(0 <= sensorIndex && sensorIndex < (*b2World)(unsafe.Pointer(world)).Sensors.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+21507, __ccgo_ts+21483, int32FromInt32(147))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		sensor = (*b2World)(unsafe.Pointer(world)).Sensors.Data + uintptr(sensorIndex)*56
		if v3 = !!(0 <= (*b2Sensor)(unsafe.Pointer(sensor)).ShapeId && (*b2Sensor)(unsafe.Pointer(sensor)).ShapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v3 {
			b2InternalAssert(tls, __ccgo_ts+21574, __ccgo_ts+21483, int32FromInt32(148))
		}
		_ = v3 || bool(int32FromInt32(0) != 0)
		sensorShape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b2Sensor)(unsafe.Pointer(sensor)).ShapeId)*280
		// Swap overlap arrays
		temp = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1
		(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2
		(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2 = temp
		(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count = 0
		// Append sensor hits
		hitCount = (*b2Sensor)(unsafe.Pointer(sensor)).Hits.Count
		i = 0
		for {
			if !(i < hitCount) {
				break
			}
			elementSize = int32(8)
			if (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count >= (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Capacity {
				oldSize = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Capacity * elementSize
				if (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Capacity == 0 {
					v6 = int32(8)
				} else {
					v6 = int32(2) * (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Capacity
				}
				newCapacity = v6
				newSize = newCapacity * elementSize
				(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data = b2GrowAlloc(tls, (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data, uint64FromInt32(oldSize), uint64FromInt32(newSize))
				(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Capacity = newCapacity
			}
			*(*b2Visitor)(unsafe.Add(nil, (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data+uintptr((*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count)*8)) = *(*b2Visitor)(unsafe.Add(nil, (*b2Sensor)(unsafe.Pointer(sensor)).Hits.Data+uintptr(i)*8))
			(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count += int32(1)
			goto _5
		_5:
			;
			i = i + 1
		}
		// Clear the hits
		(*b2Sensor)(unsafe.Pointer(sensor)).Hits.Count = 0
		if v1 = !!(0 <= (*b2Shape)(unsafe.Pointer(sensorShape)).BodyId && (*b2Shape)(unsafe.Pointer(sensorShape)).BodyId < (*b2World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
			b2InternalAssert(tls, __ccgo_ts+21648, __ccgo_ts+21483, int32FromInt32(166))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		body = (*b2World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b2Shape)(unsafe.Pointer(sensorShape)).BodyId)*104
		if (*b2Body)(unsafe.Pointer(body)).SetIndex == int32(b2_disabledSet) || int32FromUint8((*b2Shape)(unsafe.Pointer(sensorShape)).EnableSensorEvents) == false1 {
			if (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Count != 0 {
				// This sensor is dropping all overlaps because it has been disabled.
				v8 = taskContext
				v9 = uint32FromInt32(sensorIndex)
				blockIndex = v9 / uint32(64)
				if v3 = !!(blockIndex < (*b2BitSet)(unsafe.Pointer(v8)).BlockCount); !v3 {
					b2InternalAssert(tls, __ccgo_ts+7356, __ccgo_ts+7388, int32FromInt32(29))
				}
				_ = v3 || bool(int32FromInt32(0) != 0)
				*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v8)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v9 % uint32FromInt32(64))
			}
			goto _2
		}
		transform = b2GetBodyTransformQuick(tls, world, body)
		*(*b2SensorQueryContext)(unsafe.Add(nil, bp)) = b2SensorQueryContext{
			World:       world,
			TaskContext: taskContext,
			Sensor:      sensor,
			SensorShape: sensorShape,
			Transform:   transform,
		}
		if v1 = !!((*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex == sensorIndex); !v1 {
			b2InternalAssert(tls, __ccgo_ts+21730, __ccgo_ts+21483, int32FromInt32(187))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		queryBounds = (*b2Shape)(unsafe.Pointer(sensorShape)).Aabb
		// Query all trees
		b2DynamicTree_OverlapAABB(tls, trees+uintptr(0)*104, queryBounds, (*b2Shape)(unsafe.Pointer(sensorShape)).Filter.MaskBits, __ccgo_fp(b2SensorQueryCallback), bp)
		b2DynamicTree_OverlapAABB(tls, trees+uintptr(1)*104, queryBounds, (*b2Shape)(unsafe.Pointer(sensorShape)).Filter.MaskBits, __ccgo_fp(b2SensorQueryCallback), bp)
		b2DynamicTree_OverlapAABB(tls, trees+uintptr(2)*104, queryBounds, (*b2Shape)(unsafe.Pointer(sensorShape)).Filter.MaskBits, __ccgo_fp(b2SensorQueryCallback), bp)
		// Sort the overlaps to enable finding begin and end events.
		qsort(tls, (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data, uint64FromInt32((*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count), uint64(8), b2CompareVisitors)
		// Remove duplicates from overlaps2 (sorted). Duplicates are possible due to the hit events appended earlier.
		uniqueCount = 0
		overlapCount = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count
		overlapData = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data
		i1 = 0
		for {
			if !(i1 < overlapCount) {
				break
			}
			if uniqueCount == 0 || (*(*b2Visitor)(unsafe.Add(nil, overlapData+uintptr(i1)*8))).ShapeId != (*(*b2Visitor)(unsafe.Add(nil, overlapData+uintptr(uniqueCount-int32(1))*8))).ShapeId {
				*(*b2Visitor)(unsafe.Add(nil, overlapData+uintptr(uniqueCount)*8)) = *(*b2Visitor)(unsafe.Add(nil, overlapData+uintptr(i1)*8))
				uniqueCount = uniqueCount + int32(1)
			}
			goto _12
		_12:
			;
			i1 = i1 + 1
		}
		(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count = uniqueCount
		count1 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Count
		count2 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count
		if count1 != count2 {
			// something changed
			v8 = taskContext
			v9 = uint32FromInt32(sensorIndex)
			blockIndex = v9 / uint32(64)
			if v1 = !!(blockIndex < (*b2BitSet)(unsafe.Pointer(v8)).BlockCount); !v1 {
				b2InternalAssert(tls, __ccgo_ts+7356, __ccgo_ts+7388, int32FromInt32(29))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v8)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v9 % uint32FromInt32(64))
		} else {
			i2 = 0
			for {
				if !(i2 < count1) {
					break
				}
				s1 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Data + uintptr(i2)*8
				s2 = (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data + uintptr(i2)*8
				if (*b2Visitor)(unsafe.Pointer(s1)).ShapeId != (*b2Visitor)(unsafe.Pointer(s2)).ShapeId || int32FromUint16((*b2Visitor)(unsafe.Pointer(s1)).Generation) != int32FromUint16((*b2Visitor)(unsafe.Pointer(s2)).Generation) {
					// something changed
					v8 = taskContext
					v9 = uint32FromInt32(sensorIndex)
					blockIndex = v9 / uint32(64)
					if v1 = !!(blockIndex < (*b2BitSet)(unsafe.Pointer(v8)).BlockCount); !v1 {
						b2InternalAssert(tls, __ccgo_ts+7356, __ccgo_ts+7388, int32FromInt32(29))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					*(*uint64_t)(unsafe.Add(nil, (*b2BitSet)(unsafe.Pointer(v8)).Bits+uintptr(blockIndex)*8)) |= uint64FromInt32(1) << (v9 % uint32FromInt32(64))
					break
				}
				goto _16
			_16:
				;
				i2 = i2 + 1
			}
		}
		goto _2
	_2:
		;
		sensorIndex = sensorIndex + 1
	}
}

func b2OverlapSensors(tls *_Stack, world uintptr) {
	var bitSet, bits, r1, r11, r2, r21, refs1, refs2, sensor, sensorShape uintptr
	var blockCount, ctz, k, v5 uint32_t
	var count1, count2, elementSize, elementSize1, elementSize2, elementSize3, elementSize4, elementSize5, i, i1, index1, index2, minRange, newCapacity, newCapacity1, newCapacity2, newCapacity3, newCapacity4, newCapacity5, newSize, newSize1, newSize2, newSize3, newSize4, newSize5, oldSize, oldSize1, oldSize2, oldSize3, oldSize4, oldSize5, sensorCount, sensorIndex, v9 int32
	var event, event2, event4 SensorEndTouchEvent
	var event1, event3, event5 SensorBeginTouchEvent
	var sensorId, visitorId, visitorId1, visitorId2, visitorId3, visitorId4, visitorId5 ShapeId
	var word uint64_t
	var v1, v7 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bitSet, bits, blockCount, count1, count2, ctz, elementSize, elementSize1, elementSize2, elementSize3, elementSize4, elementSize5, event, event1, event2, event3, event4, event5, i, i1, index1, index2, k, minRange, newCapacity, newCapacity1, newCapacity2, newCapacity3, newCapacity4, newCapacity5, newSize, newSize1, newSize2, newSize3, newSize4, newSize5, oldSize, oldSize1, oldSize2, oldSize3, oldSize4, oldSize5, r1, r11, r2, r21, refs1, refs2, sensor, sensorCount, sensorId, sensorIndex, sensorShape, visitorId, visitorId1, visitorId2, visitorId3, visitorId4, visitorId5, word, v1, v5, v7, v9
	sensorCount = (*b2World)(unsafe.Pointer(world)).Sensors.Count
	if sensorCount == 0 {
		return
	}
	if v1 = !!((*b2World)(unsafe.Pointer(world)).WorkerCount > int32FromInt32(0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+17527, __ccgo_ts+21483, int32FromInt32(247))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	i = 0
	for {
		if !(i < (*b2World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		b2SetBitCountAndClear(tls, (*b2World)(unsafe.Pointer(world)).SensorTaskContexts.Data+uintptr(i)*16, uint32FromInt32(sensorCount))
		goto _2
	_2:
		;
		i = i + 1
	}
	// Parallel-for sensors overlaps
	minRange = int32(16)
	b2ParallelFor(tls, world, __ccgo_fp(b2SensorTask), sensorCount, minRange, world)
	bitSet = (*b2World)(unsafe.Pointer(world)).SensorTaskContexts.Data
	i1 = int32(1)
	for {
		if !(i1 < (*b2World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		b2InPlaceUnion(tls, bitSet, (*b2World)(unsafe.Pointer(world)).SensorTaskContexts.Data+uintptr(i1)*16)
		goto _3
	_3:
		;
		i1 = i1 + 1
	}
	// Iterate sensors bits and publish events
	// Process sensor state changes. Iterate over set bits
	bits = (*b2BitSet)(unsafe.Pointer(bitSet)).Bits
	blockCount = (*b2BitSet)(unsafe.Pointer(bitSet)).BlockCount
	k = uint32(0)
	for {
		if !(k < blockCount) {
			break
		}
		word = *(*uint64_t)(unsafe.Add(nil, bits+uintptr(k)*8))
		for word != uint64(0) {
			v5 = uint32FromInt32(__builtin_ctzll(tls, word))
			goto _6
		_6:
			ctz = v5
			sensorIndex = int32FromUint32(uint32FromInt32(64)*k + ctz)
			if v1 = !!(0 <= sensorIndex && sensorIndex < (*b2World)(unsafe.Pointer(world)).Sensors.Count); !v1 {
				b2InternalAssert(tls, __ccgo_ts+21507, __ccgo_ts+21483, int32FromInt32(281))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			sensor = (*b2World)(unsafe.Pointer(world)).Sensors.Data + uintptr(sensorIndex)*56
			if v7 = !!(0 <= (*b2Sensor)(unsafe.Pointer(sensor)).ShapeId && (*b2Sensor)(unsafe.Pointer(sensor)).ShapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v7 {
				b2InternalAssert(tls, __ccgo_ts+21574, __ccgo_ts+21483, int32FromInt32(282))
			}
			_ = v7 || bool(int32FromInt32(0) != 0)
			sensorShape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b2Sensor)(unsafe.Pointer(sensor)).ShapeId)*280
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
			index1 = 0
			index2 = 0
			for index1 < count1 && index2 < count2 {
				r1 = refs1 + uintptr(index1)*8
				r2 = refs2 + uintptr(index2)*8
				if (*b2Visitor)(unsafe.Pointer(r1)).ShapeId == (*b2Visitor)(unsafe.Pointer(r2)).ShapeId {
					if int32FromUint16((*b2Visitor)(unsafe.Pointer(r1)).Generation) < int32FromUint16((*b2Visitor)(unsafe.Pointer(r2)).Generation) {
						// end
						visitorId = ShapeId{
							Index1:     (*b2Visitor)(unsafe.Pointer(r1)).ShapeId + int32(1),
							World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
							Generation: (*b2Visitor)(unsafe.Pointer(r1)).Generation,
						}
						event = SensorEndTouchEvent{
							SensorShapeId:  sensorId,
							VisitorShapeId: visitorId,
						}
						elementSize = int32(16)
						if (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count >= (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity {
							oldSize = (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity * elementSize
							if (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity == 0 {
								v9 = int32(8)
							} else {
								v9 = int32(2) * (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity
							}
							newCapacity = v9
							newSize = newCapacity * elementSize
							(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data = b2GrowAlloc(tls, (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data, uint64FromInt32(oldSize), uint64FromInt32(newSize))
							(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity = newCapacity
						}
						*(*SensorEndTouchEvent)(unsafe.Add(nil, (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data+uintptr((*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count)*16)) = event
						(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count += int32(1)
						index1 = index1 + int32(1)
					} else {
						if int32FromUint16((*b2Visitor)(unsafe.Pointer(r1)).Generation) > int32FromUint16((*b2Visitor)(unsafe.Pointer(r2)).Generation) {
							// begin
							visitorId1 = ShapeId{
								Index1:     (*b2Visitor)(unsafe.Pointer(r2)).ShapeId + int32(1),
								World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
								Generation: (*b2Visitor)(unsafe.Pointer(r2)).Generation,
							}
							event1 = SensorBeginTouchEvent{
								SensorShapeId:  sensorId,
								VisitorShapeId: visitorId1,
							}
							elementSize1 = int32(16)
							if (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Count >= (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity {
								oldSize1 = (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity * elementSize1
								if (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity == 0 {
									v9 = int32(8)
								} else {
									v9 = int32(2) * (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity
								}
								newCapacity1 = v9
								newSize1 = newCapacity1 * elementSize1
								(*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Data = b2GrowAlloc(tls, (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Data, uint64FromInt32(oldSize1), uint64FromInt32(newSize1))
								(*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity = newCapacity1
							}
							*(*SensorBeginTouchEvent)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Data+uintptr((*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Count)*16)) = event1
							(*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Count += int32(1)
							index2 = index2 + int32(1)
						} else {
							// persisted
							index1 = index1 + int32(1)
							index2 = index2 + int32(1)
						}
					}
				} else {
					if (*b2Visitor)(unsafe.Pointer(r1)).ShapeId < (*b2Visitor)(unsafe.Pointer(r2)).ShapeId {
						// end
						visitorId2 = ShapeId{
							Index1:     (*b2Visitor)(unsafe.Pointer(r1)).ShapeId + int32(1),
							World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
							Generation: (*b2Visitor)(unsafe.Pointer(r1)).Generation,
						}
						event2 = SensorEndTouchEvent{
							SensorShapeId:  sensorId,
							VisitorShapeId: visitorId2,
						}
						elementSize2 = int32(16)
						if (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count >= (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity {
							oldSize2 = (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity * elementSize2
							if (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity == 0 {
								v9 = int32(8)
							} else {
								v9 = int32(2) * (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity
							}
							newCapacity2 = v9
							newSize2 = newCapacity2 * elementSize2
							(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data = b2GrowAlloc(tls, (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data, uint64FromInt32(oldSize2), uint64FromInt32(newSize2))
							(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity = newCapacity2
						}
						*(*SensorEndTouchEvent)(unsafe.Add(nil, (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data+uintptr((*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count)*16)) = event2
						(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count += int32(1)
						index1 = index1 + int32(1)
					} else {
						// begin
						visitorId3 = ShapeId{
							Index1:     (*b2Visitor)(unsafe.Pointer(r2)).ShapeId + int32(1),
							World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
							Generation: (*b2Visitor)(unsafe.Pointer(r2)).Generation,
						}
						event3 = SensorBeginTouchEvent{
							SensorShapeId:  sensorId,
							VisitorShapeId: visitorId3,
						}
						elementSize3 = int32(16)
						if (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Count >= (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity {
							oldSize3 = (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity * elementSize3
							if (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity == 0 {
								v9 = int32(8)
							} else {
								v9 = int32(2) * (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity
							}
							newCapacity3 = v9
							newSize3 = newCapacity3 * elementSize3
							(*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Data = b2GrowAlloc(tls, (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Data, uint64FromInt32(oldSize3), uint64FromInt32(newSize3))
							(*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity = newCapacity3
						}
						*(*SensorBeginTouchEvent)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Data+uintptr((*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Count)*16)) = event3
						(*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Count += int32(1)
						index2 = index2 + int32(1)
					}
				}
			}
			for index1 < count1 {
				// end
				r11 = refs1 + uintptr(index1)*8
				visitorId4 = ShapeId{
					Index1:     (*b2Visitor)(unsafe.Pointer(r11)).ShapeId + int32(1),
					World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
					Generation: (*b2Visitor)(unsafe.Pointer(r11)).Generation,
				}
				event4 = SensorEndTouchEvent{
					SensorShapeId:  sensorId,
					VisitorShapeId: visitorId4,
				}
				elementSize4 = int32(16)
				if (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count >= (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity {
					oldSize4 = (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity * elementSize4
					if (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity == 0 {
						v9 = int32(8)
					} else {
						v9 = int32(2) * (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity
					}
					newCapacity4 = v9
					newSize4 = newCapacity4 * elementSize4
					(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data = b2GrowAlloc(tls, (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data, uint64FromInt32(oldSize4), uint64FromInt32(newSize4))
					(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity = newCapacity4
				}
				*(*SensorEndTouchEvent)(unsafe.Add(nil, (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data+uintptr((*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count)*16)) = event4
				(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count += int32(1)
				index1 = index1 + int32(1)
			}
			for index2 < count2 {
				// begin
				r21 = refs2 + uintptr(index2)*8
				visitorId5 = ShapeId{
					Index1:     (*b2Visitor)(unsafe.Pointer(r21)).ShapeId + int32(1),
					World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
					Generation: (*b2Visitor)(unsafe.Pointer(r21)).Generation,
				}
				event5 = SensorBeginTouchEvent{
					SensorShapeId:  sensorId,
					VisitorShapeId: visitorId5,
				}
				elementSize5 = int32(16)
				if (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Count >= (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity {
					oldSize5 = (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity * elementSize5
					if (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity == 0 {
						v9 = int32(8)
					} else {
						v9 = int32(2) * (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity
					}
					newCapacity5 = v9
					newSize5 = newCapacity5 * elementSize5
					(*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Data = b2GrowAlloc(tls, (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Data, uint64FromInt32(oldSize5), uint64FromInt32(newSize5))
					(*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Capacity = newCapacity5
				}
				*(*SensorBeginTouchEvent)(unsafe.Add(nil, (*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Data+uintptr((*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Count)*16)) = event5
				(*b2World)(unsafe.Pointer(world)).SensorBeginEvents.Count += int32(1)
				index2 = index2 + int32(1)
			}
			// Clear the smallest set bit
			word = word & (word - uint64(1))
		}
		goto _4
	_4:
		;
		k = k + 1
	}
}

func b2DestroySensor(tls *_Stack, world uintptr, sensorShape uintptr) {
	var elementSize1, i, movedIndex, newCapacity, newSize, oldSize, v3, v6, v7 int32
	var event SensorEndTouchEvent
	var movedSensor, otherSensorShape, ref, sensor, v4, v5 uintptr
	var v1, v11, v8 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = elementSize1, event, i, movedIndex, movedSensor, newCapacity, newSize, oldSize, otherSensorShape, ref, sensor, v1, v11, v3, v4, v5, v6, v7, v8
	if v1 = !!(0 <= (*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex && (*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex < (*b2World)(unsafe.Pointer(world)).Sensors.Count); !v1 {
		b2InternalAssert(tls, __ccgo_ts+21770, __ccgo_ts+21483, int32FromInt32(374))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	sensor = (*b2World)(unsafe.Pointer(world)).Sensors.Data + uintptr((*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex)*56
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
				Index1:     (*b2Visitor)(unsafe.Pointer(ref)).ShapeId + int32(1),
				World0:     (*b2World)(unsafe.Pointer(world)).WorldId,
				Generation: (*b2Visitor)(unsafe.Pointer(ref)).Generation,
			},
		}
		elementSize1 = int32(16)
		if (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count >= (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity {
			oldSize = (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity * elementSize1
			if (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity == 0 {
				v3 = int32(8)
			} else {
				v3 = int32(2) * (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity
			}
			newCapacity = v3
			newSize = newCapacity * elementSize1
			(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data = b2GrowAlloc(tls, (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data, uint64FromInt32(oldSize), uint64FromInt32(newSize))
			(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Capacity = newCapacity
		}
		*(*SensorEndTouchEvent)(unsafe.Add(nil, (*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Data+uintptr((*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count)*16)) = event
		(*(*b2DynamicArray_b2SensorEndTouchEvent)(unsafe.Add(nil, world+2304+uintptr((*b2World)(unsafe.Pointer(world)).EndEventArrayIndex)*16))).Count += int32(1)
		goto _2
	_2:
		;
		i = i + 1
	}
	// Destroy sensor
	b2Free(tls, (*b2Sensor)(unsafe.Pointer(sensor)).Hits.Data, uint64FromInt32((*b2Sensor)(unsafe.Pointer(sensor)).Hits.Capacity)*uint64(8))
	(*b2Sensor)(unsafe.Pointer(sensor)).Hits.Data = uintptrFromInt32(0)
	(*b2Sensor)(unsafe.Pointer(sensor)).Hits.Count = 0
	(*b2Sensor)(unsafe.Pointer(sensor)).Hits.Capacity = 0
	b2Free(tls, (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Data, uint64FromInt32((*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Capacity)*uint64(8))
	(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Data = uintptrFromInt32(0)
	(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Count = 0
	(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps1.Capacity = 0
	b2Free(tls, (*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data, uint64FromInt32((*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Capacity)*uint64(8))
	(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Data = uintptrFromInt32(0)
	(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Count = 0
	(*b2Sensor)(unsafe.Pointer(sensor)).Overlaps2.Capacity = 0
	v4 = (*b2World)(unsafe.Pointer(world)).Sensors.Data
	v5 = world + 2208 + 8
	v3 = (*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex
	v6 = int32(56)
	if v1 = !!(0 <= v3 && v3 < *(*int32)(unsafe.Add(nil, v5)) && bool(__ccgo_ts+2029 != 0)); !v1 {
		b2InternalAssert(tls, __ccgo_ts+2055, __ccgo_ts+2115, int32FromInt32(153))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	*(*int32)(unsafe.Add(nil, v5)) = *(*int32)(unsafe.Add(nil, v5)) - 1
	if v3 != *(*int32)(unsafe.Add(nil, v5)) {
		memcpy(tls, v4+uintptr(v3*v6), v4+uintptr(*(*int32)(unsafe.Add(nil, v5))*v6), uint64FromInt32(v6))
		v7 = *(*int32)(unsafe.Add(nil, v5))
		goto _10
	}
	v7 = -int32(1)
	goto _10
_10:
	movedIndex = v7
	if movedIndex != -int32(1) {
		if v8 = !!(0 <= (*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex && (*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex < (*b2World)(unsafe.Pointer(world)).Sensors.Count); !v8 {
			b2InternalAssert(tls, __ccgo_ts+21770, __ccgo_ts+21483, int32FromInt32(405))
		}
		_ = v8 || bool(int32FromInt32(0) != 0)
		// Fixup moved sensor
		movedSensor = (*b2World)(unsafe.Pointer(world)).Sensors.Data + uintptr((*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex)*56
		if v11 = !!(0 <= (*b2Sensor)(unsafe.Pointer(movedSensor)).ShapeId && (*b2Sensor)(unsafe.Pointer(movedSensor)).ShapeId < (*b2World)(unsafe.Pointer(world)).Shapes.Count); !v11 {
			b2InternalAssert(tls, __ccgo_ts+21863, __ccgo_ts+21483, int32FromInt32(406))
		}
		_ = v11 || bool(int32FromInt32(0) != 0)
		otherSensorShape = (*b2World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b2Sensor)(unsafe.Pointer(movedSensor)).ShapeId)*280
		(*b2Shape)(unsafe.Pointer(otherSensorShape)).SensorIndex = (*b2Shape)(unsafe.Pointer(sensorShape)).SensorIndex
	}
}
