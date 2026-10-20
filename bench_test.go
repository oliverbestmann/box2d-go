package b2

import "testing"

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
