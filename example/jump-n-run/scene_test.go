package main

import (
	"testing"

	b2 "github.com/oliverbestmann/box2d-go/v3"
)

// step simulates the given seconds, with jump the jump button is held the whole time.
func step(s *Scene, seconds float64, throttle float32, jump bool) {
	for i := range int(seconds / timeStep) {
		s.Step(throttle, jump && i == 0, jump)
	}
}

// jumpHeight returns how high the player gets when holding jump for the given seconds.
func jumpHeight(held float64) float32 {
	s := NewScene()
	defer s.Destroy()

	step(s, 2, 0, false)
	y := s.Player.Body.GetPosition().Y
	maxY := y

	for i := range 60 {
		hold := float64(i)*timeStep < held
		s.Step(0, i == 0, hold)
		maxY = max(maxY, s.Player.Body.GetPosition().Y)
	}

	return maxY - y
}

func TestJumpHeight(t *testing.T) {
	full, short := jumpHeight(1), jumpHeight(0.05)
	t.Logf("full jump %.2fm, short hop %.2fm", full, short)

	if full < 2.5 || full > 3.2 {
		t.Fatalf("full jump height %v", full)
	}
	if short > full/2 {
		t.Fatalf("releasing jump early does not cut the jump: %v vs %v", short, full)
	}
}

func TestOneWayPlatform(t *testing.T) {
	s := NewScene()
	defer s.Destroy()

	// start below the right ledge moving up, the player passes through and lands on top
	s.Player.Body.SetTransform(b2.Pos{X: 34.5, Y: 4.5}, b2.IdentityRot)
	s.Player.Body.SetLinearVelocity(b2.Vec2{Y: 15})

	step(s, 1.5, 0, false)
	if y := s.Player.Body.GetPosition().Y; y < 6 || !s.Player.OnGround {
		t.Fatalf("player did not land on the ledge: y=%v on ground=%v", y, s.Player.OnGround)
	}
}

func TestScene(t *testing.T) {
	s := NewScene()
	defer s.Destroy()

	welds := len(s.Welds)

	// the player lands and hovers on the ground, the tower stays intact
	step(s, 2, 0, false)
	if !s.Player.OnGround {
		t.Fatalf("player not on ground at %v", s.Player.Body.GetPosition())
	}
	if len(s.Welds) != welds {
		t.Fatalf("tower broke under its own weight: %d of %d welds left", len(s.Welds), welds)
	}

	// walking right moves the player
	x := s.Player.Body.GetPosition().X
	step(s, 0.5, 1, false)
	if dx := s.Player.Body.GetPosition().X - x; dx < 1 {
		t.Fatalf("player moved only %v", dx)
	}

	// jumping leaves the ground
	y := s.Player.Body.GetPosition().Y
	step(s, 0.3, 0, true)
	if dy := s.Player.Body.GetPosition().Y - y; dy < 1 {
		t.Fatalf("player jumped only %v", dy)
	}

	// shooting the tower breaks welds
	s.Player.Body.SetTransform(b2.Pos{X: 16, Y: 2.5}, b2.IdentityRot)
	step(s, 2, 0, false)
	for range 10 {
		s.Shoot(b2.Pos{X: 20, Y: 2})
		step(s, 0.2, 0, false)
	}
	step(s, 1, 0, false)
	if len(s.Welds) == welds {
		t.Fatalf("no weld broke")
	}
	t.Logf("%d of %d welds left, %d projectiles", len(s.Welds), welds, len(s.Projectiles))
}

func TestPlatformCarriesPlayer(t *testing.T) {
	s := NewScene()
	defer s.Destroy()

	// put the player on the shuttle and check it follows it
	shuttle := s.platforms[1].body
	step(s, 0.1, 0, false)
	p := shuttle.GetPosition()
	s.Player.Body.SetTransform(b2.Pos{X: p.X, Y: p.Y + 1.5}, b2.IdentityRot)

	step(s, 1, 0, false)
	for range 120 {
		s.Step(0, false, false)
		dx := s.Player.Body.GetPosition().X - shuttle.GetPosition().X
		if dx < -0.5 || dx > 0.5 {
			t.Fatalf("player drifted %v from the shuttle", dx)
		}
	}
}
