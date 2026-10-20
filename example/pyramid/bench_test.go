package main

import (
	"testing"
	"time"
)

// BenchmarkPhysics runs the first 5 seconds of the example scene with each engine. In the
// paced runs every step waits for the next 60 Hz frame, like the example does, so worker
// threads go idle between steps.
func BenchmarkPhysics(b *testing.B) {
	engines := []struct {
		name       string
		new        func(gravity float64, concurrent bool) Physics
		concurrent bool
	}{
		{"box2d", b2New, false},
		{"box2d-concurrent", b2New, true},
		{"chipmunk", cpNew, false},
	}

	for _, paced := range []bool{false, true} {
		for _, engine := range engines {
			name := engine.name
			if paced {
				name += "/paced"
			}

			b.Run(name, func(b *testing.B) {
				const stepCount = 300

				var physicsTime time.Duration
				for range b.N {
					v := viz{makePhysics: engine.new, concurrent: engine.concurrent}
					v.initialize()

					next := time.Now()
					for range stepCount {
						if paced {
							next = next.Add(time.Second / 60)
							time.Sleep(time.Until(next))
						}

						start := time.Now()
						v.physics.Step(1/60.0, 4)
						physicsTime += time.Since(start)
					}

					v.physics.Destroy()
				}

				b.ReportMetric(float64(physicsTime)/1e6/float64(b.N*stepCount), "ms/step")
			})
		}
	}
}
