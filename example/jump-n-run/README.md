# Jump and run

This README describes the physics code of a small jump and run sandbox. The game uses box2d-go and draws with the ebitengine `vector` package. It is meant for readers who want to change the physics or reuse the character controller.

Run it from the `example` directory:

```sh
go run ./jump-n-run
```

Move with A/D or the arrow keys and jump with W, Space, or Up. Hold the left mouse button to shoot at the cursor. R resets the scene.

## Files

- **`mover.go`** contains the character controller.
- **`scene.go`** builds the level and steps the world. It does not import ebitengine, so the tests run without a window.
- **`main.go`** reads input, calls `Scene.Step`, and draws.
- **`scene_test.go`** runs the scene headless and checks movement, jumps, one-way platforms, moving platforms, and the breakable tower.

## World and time step

The world is 40 by 22.5 meters with gravity of 10 m/s² pointing down. The origin is the bottom left corner and y points up. The screen shows the whole world at 32 pixels per meter, so the game has no camera.

`main.go` sets ebitengine to 120 updates per second, and every update runs one physics step of 1/120 s with 4 substeps (box2d splits each step into smaller solver passes for stability). On a 60 Hz display ebitengine runs two updates per frame, so the game speed does not depend on the display.

`Scene.Step` runs in this order:

1. Move the kinematic platforms to their next position.
2. Apply a jump if one was requested, then update the character.
3. Step the world.
4. Destroy welds that broke and projectiles that expired.

## The character

`Mover` is a Go port of `DynamicMover` from the box2d 3.2 samples (`samples/dynamic_mover.cpp` in the box2d repository). The character is a dynamic body (a body that the solver moves in response to forces and collisions) with a capsule shape 1.6 m tall. Because it is a dynamic body, crates and balls push the character, and the character pushes them back with its real mass.

The body has rotation locked, sleep disabled, and zero surface friction. The joints below move the character, and surface friction would work against them.

### Mover joint

A mover joint connects the character to an empty static body. Each update sets the joint's target velocity, and the joint applies up to a maximum force to reach it. The force limit is only set on the x axis, so the joint does not affect vertical motion. The limit is 70 N on walkable ground, 40 N in the air, and 0 N on ground that is too steep. A small air force keeps the character's momentum while jumping, and the zero force lets it slide down slopes.

### Pogo joint

The character does not stand on its capsule. A pogo joint (a one-axis spring between the character and the ground) holds the center of the capsule's bottom circle 0.9 m above the ground. The capsule floats 0.6 m above the ground, which lets the character walk over small steps and bumps.

Every update casts a ray down from the center of the capsule's bottom circle. The ray is 0.9 m long in the air and 1.8 m long on the ground, so the spring also catches the ground after a step down. If the ray hits, the update destroys the old pogo joint and creates a new one anchored at the hit point, because the character may now stand on a different body. The new joint starts with the impulse and velocity of the old one, so the spring does not reset (this is called warm starting).

The spring has a frequency of 5 Hz and a damping ratio of 0.8. Its push and pull forces are limited to 100 times the character's weight. While jumping, the pull limit is zero, so the spring cannot pull the character back down.

The game draws the pogo as a plum line from the capsule to the anchor. Without a hit it draws the ray in gray.

### Walking

Walking follows the ground and air movement model of Quake. Each update:

1. Subtracts the velocity of the ground under the character, so the next two steps work relative to the ground.
2. Applies friction on the ground. Above 3 m/s the speed drops in proportion to itself. Below 3 m/s it drops by a fixed 24 m/s per second, so the character stops quickly. Below 0.1 m/s the horizontal speed is set to zero.
3. Accelerates towards the input direction, up to 7 m/s. In the air the acceleration is 80 percent of the ground acceleration.
4. Adds the ground velocity back and passes the result to the mover joint.

The first and last steps are the one change from the box2d sample. Without them the friction step slows the character relative to the world, and it slides off platforms that move sideways.

A surface counts as walkable when the y component of its normal is at least 0.7, which means a slope of about 45 degrees or less.

### Jumping

A jump only works on walkable ground. It applies an upward impulse that sets the vertical speed to 13 m/s, plus the vertical speed of the ground, so jumps from a rising elevator go higher.

Three things shorten the jump arc:

- The character uses 3 times the world gravity while rising with the jump button held.
- It uses 5 times the world gravity while falling or after the button is released.
- Releasing the button while rising halves the upward speed once.

A full jump reaches 2.8 m and a quick tap reaches about 0.9 m. The jump ends when the ray hits the ground again and the character moves towards it. The first 1/24 s of a jump never ends it, so the character can leave the ground first.

The pogo force limits use the 3 times gravity weight, also while falling.

## Moving platforms

The two moving platforms are kinematic bodies (bodies with infinite mass that move at a set velocity and ignore forces). Each update computes the next position on a cosine curve between two points and calls `SetTargetTransform`, which sets the velocity that reaches that position in one step. The pogo joint attaches to the platform like any other body, and the ground velocity correction in the walking steps keeps the character on it.

The bar below the ceiling is a kinematic body with a fixed angular velocity of 0.6 rad/s.

## One-way platforms

Both moving platforms and the two thin ledges are one-way: the character jumps up through them from below and lands on top. Their shapes have the user data value `userDataOneWay`.

The character's shape enables pre-solve events, and `Scene` registers two callbacks:

- **`preSolve`** runs before the solver uses a contact. If the contact is between the character and a one-way shape and the contact normal does not point up from the platform to the character (normal y of 0.7 or less), it sets the manifold point count to zero, which disables the contact for this step.
- **`preContinuous`** applies the same rule to continuous collision (the check that stops fast bodies from passing through thin shapes between steps).

Crates, balls, and projectiles do not enable pre-solve events, so one-way platforms are solid for them. One side effect is that the elevator passes through the character when it moves down onto it, instead of pushing it into the floor.

## Breakable tower

The tower in the middle is 3 by 8 bricks, each 0.8 m square. Weld joints connect every brick to its right and upper neighbor. Each weld has a `ForceThreshold` of 400 N. When the force in a weld goes above the threshold, box2d reports a joint event, and `Scene.Step` destroys that joint. The bricks stay in the world and fall apart once enough welds are gone.

The threshold depends on the time step, because the same hit creates a larger force over a shorter step. The value of 400 N comes from headless runs at 120 Hz: the tower survives its own weight, and one shot breaks several welds.

The two loose towers and the pyramid have no joints and only rely on friction (0.7) to stay upright.

## Projectiles and collision filters

A projectile is a ball with a radius of 0.15 m and a density of 20 kg/m², which makes it heavy for its size. It starts at 30 m/s towards the cursor. It is marked as a bullet, so box2d runs continuous collision against dynamic bodies too and the projectile does not pass through bricks. The game fires at most one projectile every 0.15 s and removes each one after 5 seconds.

Collision filters use three category bits:

| Category | Bit | Collides with |
|---|---|---|
| World | 1 | everything |
| Player | 2 | world |
| Projectile | 4 | world, projectiles |

The pogo ray uses the character's filter, so it never hits the character itself or a projectile.

## Tuning

The fields on `Mover` hold all the movement values. If the jump feels floaty, raise `GravityScale` and `FallGravityScale`. If you raise them, also raise `JumpSpeed`, or the jump gets lower. The jump height in meters is `JumpSpeed² / (2 × 10 × GravityScale)`. Some ledges in the level need close to the full 2.8 m.
