package b2

import (
	"fmt"
	"math"
	"testing"
)

// BenchmarkPyramid steps a large pyramid of boxes, which keeps the contact solver busy.
func BenchmarkPyramid(b *testing.B) {
	def := DefaultWorldDef()
	def.EnableSleep = 0
	world := CreateWorld(def)
	defer world.DestroyWorld()

	groundDef := DefaultBodyDef()
	groundDef.Position = Pos{X: 0, Y: -1}
	ground := world.CreateBody(groundDef)
	ground.CreatePolygonShape(DefaultShapeDef(), MakeBox(100, 1))

	const baseCount = 60
	box := MakeSquare(0.5)
	for row := range baseCount {
		for col := range baseCount - row {
			bodyDef := DefaultBodyDef()
			bodyDef.Type1 = DynamicBody
			bodyDef.Position = Pos{X: float32(col) + 0.5*float32(row) - 0.5*baseCount, Y: 0.5 + float32(row)}
			world.CreateBody(bodyDef).CreatePolygonShape(DefaultShapeDef(), box)
		}
	}

	// settle the pyramid so every iteration solves the same contacts
	for range 10 {
		world.Step(1.0/60.0, 4)
	}

	b.ResetTimer()
	for range b.N {
		world.Step(1.0/60.0, 4)
	}
}

// BenchmarkExample runs the first 4 seconds of the scene in the example directory: a
// small pyramid hit by a heavy, fast box.
func BenchmarkExample(b *testing.B) {
	benchmarkScene(b, Capacity{}, 241, createExample)
}

func createExample(world World) func(int) {
	const layers = 20

	floorDef := DefaultShapeDef()
	floorDef.Material.Restitution = 0.1
	floorDef.Material.Friction = 0.9

	ground := world.CreateBody(DefaultBodyDef())
	ground.CreateSegmentShape(floorDef, Segment{Point1: Vec2{X: -2 * layers}, Point2: Vec2{X: 2 * layers}})
	ground.CreateSegmentShape(floorDef, Segment{Point1: Vec2{X: -2 * layers}, Point2: Vec2{X: -2 * layers, Y: 100}})
	ground.CreateSegmentShape(floorDef, Segment{Point1: Vec2{X: 2 * layers}, Point2: Vec2{X: 2 * layers, Y: 100}})

	boxDef := DefaultShapeDef()
	boxDef.Material.Restitution = 0.25
	boxDef.Material.Friction = 0.5

	bodyDef := DefaultBodyDef()
	bodyDef.Type1 = DynamicBody

	for l := range layers {
		for i := range l {
			bodyDef.Position = Pos{X: float32(i) + 0.5 - float32(l)/2, Y: layers - float32(l) - 0.5}
			world.CreateBody(bodyDef).CreatePolygonShape(boxDef, MakeSquare(0.5))
		}
	}

	bigDef := DefaultShapeDef()
	bigDef.Density = 10
	bigDef.Material.Restitution = 0.1
	bigDef.Material.Friction = 0.9

	bodyDef.Position = Pos{X: -1.8 * layers, Y: 2.6 * layers}
	bodyDef.LinearVelocity = Vec2{X: 75, Y: -100}
	world.CreateBody(bodyDef).CreatePolygonShape(bigDef, MakeSquare(2))

	return nil
}

// The benchmarks below are ported from box2d shared/benchmarks.c and benchmark/main.c, in
// the release sizes. Compare ms/step with the upstream benchmark, which prints the total
// milliseconds of the timed steps.

// benchmarkScene runs the scene built by create once per iteration, like the box2d
// benchmark: the world gets the capacity of the scene, the first step is not timed, and
// the others are. create may return a function that runs before every step, like the step
// functions of the box2d benchmarks.
func benchmarkScene(b *testing.B, capacity Capacity, stepCount int, create func(world World) func(step int)) {
	for _, workerCount := range []int32{1, 4} {
		b.Run(fmt.Sprintf("workers=%d", workerCount), func(b *testing.B) {
			b.ReportAllocs()

			for range b.N {
				b.StopTimer()

				def := DefaultWorldDef()
				def.WorkerCount = workerCount
				def.Capacity = capacity
				world := CreateWorld(def)

				beforeStep := create(world)

				for step := range stepCount {
					if step == 1 {
						b.StartTimer()
					}

					if beforeStep != nil {
						beforeStep(step)
					}

					world.Step(1.0/60.0, 4)
				}

				b.StopTimer()
				world.DestroyWorld()
				b.StartTimer()
			}

			b.ReportMetric(float64(b.Elapsed())/1e6/float64(b.N*(stepCount-1)), "ms/step")
		})
	}
}

// BenchmarkJointGrid steps a grid of circles connected by revolute joints. Most of the
// time goes into the joint solver.
func BenchmarkJointGrid(b *testing.B) {
	capacity := Capacity{StaticShapeCount: 7, DynamicShapeCount: 11000, StaticBodyCount: 7, DynamicBodyCount: 11000, ArenaByteCount: 120 * 1024}
	benchmarkScene(b, capacity, 500, createJointGrid)
}

func createJointGrid(world World) func(int) {
	world.EnableSleeping(0)

	const n = 100

	shapeDef := DefaultShapeDef()
	shapeDef.Density = 1
	shapeDef.Filter.CategoryBits = 2
	shapeDef.Filter.MaskBits = uint64(^uint32(2))

	circle := Circle{Radius: 0.4}

	jointDef := DefaultRevoluteJointDef()
	jointDef.Base.DrawScale = 0.4

	bodyDef := DefaultBodyDef()
	bodies := make([]Body, 0, n*n)

	for k := range n {
		for i := range n {
			bodyDef.Type1 = DynamicBody
			if k >= n/2-3 && k <= n/2+3 && i == 0 {
				bodyDef.Type1 = StaticBody
			}

			bodyDef.Position = Pos{X: float32(k), Y: -float32(i)}
			body := world.CreateBody(bodyDef)
			body.CreateCircleShape(shapeDef, circle)

			if i > 0 {
				jointDef.Base.BodyIdA = bodies[len(bodies)-1].Id
				jointDef.Base.BodyIdB = body.Id
				jointDef.Base.LocalFrameA.P = Vec2{Y: -0.5}
				jointDef.Base.LocalFrameB.P = Vec2{Y: 0.5}
				world.CreateRevoluteJoint(jointDef)
			}

			if k > 0 {
				jointDef.Base.BodyIdA = bodies[len(bodies)-n].Id
				jointDef.Base.BodyIdB = body.Id
				jointDef.Base.LocalFrameA.P = Vec2{X: 0.5}
				jointDef.Base.LocalFrameB.P = Vec2{X: -0.5}
				world.CreateRevoluteJoint(jointDef)
			}

			bodies = append(bodies, body)
		}
	}

	return nil
}

// BenchmarkJunkyard pushes a heap of small polygons around with a kinematic body, so
// contacts begin and end all the time.
func BenchmarkJunkyard(b *testing.B) {
	capacity := Capacity{StaticShapeCount: 290, DynamicShapeCount: 8800, StaticBodyCount: 1, DynamicBodyCount: 8900, ContactCount: 38000, ArenaByteCount: 3700 * 1024}
	benchmarkScene(b, capacity, 800, createJunkyard)
}

func createJunkyard(world World) func(int) {
	const gridSize = 1

	groundDef := DefaultBodyDef()
	ground := world.CreateBody(groundDef)
	shapeDef := DefaultShapeDef()

	x := float32(-80 * gridSize)
	for range 161 {
		ground.CreatePolygonShape(shapeDef, MakeOffsetBox(0.55*gridSize, 0.5*gridSize, Vec2{X: x}, IdentityRot))
		x += gridSize
	}

	for _, x := range []float32{-80 * gridSize, 80 * gridSize} {
		y := float32(gridSize)
		for range 50 {
			ground.CreatePolygonShape(shapeDef, MakeOffsetBox(0.5*gridSize, 0.55*gridSize, Vec2{X: x, Y: y}, IdentityRot))
			y += gridSize
		}
	}

	const columnCount = 200
	const rowCount = 40
	const radius = 0.25

	// Fibonacci sphere algorithm
	phi := float32(math.Pi * (math.Sqrt(5) - 1))
	points := make([]Vec2, 5)
	for i := range points {
		cs := ComputeCosSin(phi * float32(i))
		points[i] = Vec2{X: radius * cs.Cosine, Y: radius * cs.Sine}
	}

	hull, _ := ComputeHull(points)
	polygon := MakePolygon(hull, 0)

	bodyDef := DefaultBodyDef()
	bodyDef.Type1 = DynamicBody

	side := float32(-0.1)
	const yStart = 15

	for i := range columnCount {
		x := 1.5 * float32(2*i-columnCount) * radius

		for j := range rowCount {
			y := 4*float32(j)*radius + yStart

			bodyDef.Position = Pos{X: x + side, Y: y}
			side = -side

			world.CreateBody(bodyDef).CreatePolygonShape(shapeDef, polygon)
		}
	}

	bodyDef.Type1 = KinematicBody
	bodyDef.Position = Pos{}
	pusher := world.CreateBody(bodyDef)
	pusher.CreatePolygonShape(shapeDef, MakeOffsetBox(2, 4, Vec2{Y: 4}, IdentityRot))

	return func(step int) {
		const timeStep = 1.0 / 60.0
		cs := ComputeCosSin(0.2 * timeStep * float32(step))
		target := WorldTransform{P: Pos{X: 60 * cs.Sine}, Q: IdentityRot}
		pusher.SetTargetTransform(target, timeStep, 1)
	}
}

// BenchmarkQueries runs ray casts, shape casts and AABB overlaps through the Go API before
// every step, in a sparse grid of static boxes around a pile of dynamic boxes.
func BenchmarkQueries(b *testing.B) {
	capacity := Capacity{StaticShapeCount: 28000, DynamicShapeCount: 2200, StaticBodyCount: 28000, DynamicBodyCount: 2300, ContactCount: 9800, ArenaByteCount: 930 * 1024}
	benchmarkScene(b, capacity, 200, createQueries)
}

func createQueries(world World) func(int) {
	// repeatable random numbers, like QueryRandom
	randomState := uint32(1234)
	random := func(lower, upper float32) float32 {
		randomState = 1664525*randomState + 1013904223
		unit := float32(randomState>>8) * (1.0 / 16777216.0)
		return lower + (upper-lower)*unit
	}

	const extent = 500
	const cellCount = 500
	const fill = 0.1

	// the pile keeps this region of the grid clear
	const pileX = 0.4 * extent
	const pileY = 0.3 * extent
	clear := AABB{LowerBound: Vec2{X: pileX - 10, Y: pileY - 10}, UpperBound: Vec2{X: pileX + 50, Y: pileY + 60}}

	bodyDef := DefaultBodyDef()
	shapeDef := DefaultShapeDef()

	for i := range cellCount {
		y := float32(i)
		for j := range cellCount {
			x := float32(j)

			fillTest := random(0, 1)
			ratio := random(1, 5)
			halfWidth := random(0.05, 0.25)
			orientation := random(0, 1)
			category := int(random(0, 2.999))

			if fillTest > fill {
				continue
			}

			if clear.LowerBound.X <= x && x <= clear.UpperBound.X && clear.LowerBound.Y <= y && y <= clear.UpperBound.Y {
				continue
			}

			bodyDef.Position = Pos{X: x, Y: y}
			body := world.CreateBody(bodyDef)

			box := MakeBox(halfWidth, ratio*halfWidth)
			if orientation > 0.5 {
				box = MakeBox(ratio*halfWidth, halfWidth)
			}

			shapeDef.Filter.CategoryBits = 1 << category
			body.CreatePolygonShape(shapeDef, box)
		}
	}

	shapeDef.Filter.CategoryBits = 1

	// ground for the pile
	bodyDef.Position = Pos{}
	world.CreateBody(bodyDef).CreateSegmentShape(shapeDef, Segment{
		Point1: Vec2{X: pileX - 10, Y: pileY},
		Point2: Vec2{X: pileX + 50, Y: pileY},
	})

	// static sensors sitting in the pile
	sensorDef := DefaultShapeDef()
	sensorDef.IsSensor = 1
	for i := range 8 {
		bodyDef.Position = Pos{X: pileX + 4*float32(i), Y: pileY + 2}
		world.CreateBody(bodyDef).CreatePolygonShape(sensorDef, MakeSquare(2))
	}

	// the pile, dropped from just above the ground so it settles quickly
	const columnCount = 40
	const rowCount = 50
	const spacing = 0.55

	bodyDef.Type1 = DynamicBody
	for i := range rowCount {
		for j := range columnCount {
			bodyDef.Position = Pos{X: pileX + spacing*float32(j), Y: pileY + 1 + spacing*float32(i)}
			world.CreateBody(bodyDef).CreatePolygonShape(shapeDef, MakeSquare(0.25))
		}
	}

	// kinematic paddle sweeping through the bottom of the pile
	bodyDef.Type1 = KinematicBody
	bodyDef.Position = Pos{X: pileX - 5, Y: pileY + 1.2}
	bodyDef.LinearVelocity = Vec2{X: 5}
	paddle := world.CreateBody(bodyDef)
	paddle.CreatePolygonShape(shapeDef, MakeBox(1, 1))

	const paddleMinX = pileX - 5
	const paddleMaxX = pileX + 30

	// rays and casts cross the whole world, overlaps are centered on the origins
	const queryCount = 1000
	origins := make([]Pos, queryCount)
	translations := make([]Vec2, queryCount)
	for i := range queryCount {
		x1, y1 := random(0, extent), random(0, extent)
		x2, y2 := random(0, extent), random(0, extent)
		origins[i] = Pos{X: x1, Y: y1}
		translations[i] = Vec2{X: x2 - x1, Y: y2 - y1}
	}

	circle := MakeProxy([]Vec2{{}}, 0.1)

	filters := [2]QueryFilter{DefaultQueryFilter(), DefaultQueryFilter()}
	filters[1].MaskBits = 1

	return func(int) {
		// reverse the paddle at the ends of its sweep
		position := paddle.GetPosition()
		if position.X > paddleMaxX {
			paddle.SetLinearVelocity(Vec2{X: -5})
		} else if position.X < paddleMinX {
			paddle.SetLinearVelocity(Vec2{X: 5})
		}

		for _, filter := range filters {
			for i := range queryCount {
				world.CastRayClosest(origins[i], translations[i], filter)
			}

			for i := range queryCount {
				world.CastShape(origins[i], circle, translations[i], filter, func(_ Shape, _ Pos, _ Vec2, fraction float32) float32 {
					return fraction
				})
			}

			aabb := AABB{LowerBound: Vec2{X: -5, Y: -5}, UpperBound: Vec2{X: 5, Y: 5}}
			for i := range queryCount {
				world.OverlapAABB(origins[i], aabb, filter, func(Shape) bool {
					return true
				})
			}
		}
	}
}
