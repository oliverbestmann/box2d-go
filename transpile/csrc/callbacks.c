// Signatures of every box2d callback that the Go code implements or calls.
//
// check() assigns each function to its box2d callback type, so the C compiler rejects
// this file when box2d changes a callback. ccgo translates the functions to Go, and
// callbacks.go in the package root checks the Go closures against those translations.

#include "box2d/box2d.h"

// for b2ThreadFunction, implemented in Go by threads.go
#include "../../box2d-c/src/core.h"

void b2go_ThreadFunction( void* context ) {}
float b2go_FrictionCallbackFcn( float frictionA, uint64_t userMaterialIdA, float frictionB, uint64_t userMaterialIdB ) { return 0; }
float b2go_RestitutionCallbackFcn( float restitutionA, uint64_t userMaterialIdA, float restitutionB, uint64_t userMaterialIdB ) { return 0; }
bool b2go_CustomFilterCallbackFcn( b2ShapeId shapeIdA, b2ShapeId shapeIdB, void* context ) { return 0; }
void b2go_PreSolveCallbackFcn( b2ShapeId shapeIdA, b2ShapeId shapeIdB, b2Manifold* manifold, void* context ) {}
bool b2go_PreContinuousCallbackFcn( b2ShapeId shapeIdA, b2ShapeId shapeIdB, b2Pos point, b2Vec2 normal, void* context ) { return 0; }
bool b2go_OverlapResultCallbackFcn( b2ShapeId shapeId, void* context ) { return 0; }
float b2go_CastResultCallbackFcn( b2ShapeId shapeId, b2Pos point, b2Vec2 normal, float fraction, void* context ) { return 0; }
bool b2go_PlaneResultCallbackFcn( b2ShapeId shapeId, const b2PlaneResult* plane, void* context ) { return 0; }
int b2go_AssertCallbackFcn( const char* condition, const char* fileName, int lineNumber ) { return 0; }
void b2go_LogCallbackFcn( const char* message ) {}

void b2go_DrawPolygonFcn( b2WorldTransform transform, const b2Vec2* vertices, int vertexCount, b2HexColor color, void* context ) {}
void b2go_DrawSolidPolygonFcn( b2WorldTransform transform, const b2Vec2* vertices, int vertexCount, float radius, b2HexColor color, void* context ) {}
void b2go_DrawCircleFcn( b2Pos center, float radius, b2HexColor color, void* context ) {}
void b2go_DrawSolidCircleFcn( b2WorldTransform transform, b2Vec2 center, float radius, b2HexColor color, void* context ) {}
void b2go_DrawSolidCapsuleFcn( b2Pos p1, b2Pos p2, float radius, b2HexColor color, void* context ) {}
void b2go_DrawLineFcn( b2Pos p1, b2Pos p2, b2HexColor color, void* context ) {}
void b2go_DrawTransformFcn( b2WorldTransform transform, void* context ) {}
void b2go_DrawPointFcn( b2Pos p, float size, b2HexColor color, void* context ) {}
void b2go_DrawStringFcn( b2Pos p, const char* s, b2HexColor color, void* context ) {}
void b2go_DrawBoundsFcn( b2AABB aabb, b2HexColor color, void* context ) {}

void b2go_check( void )
{
	b2ThreadFunction* thread = b2go_ThreadFunction;
	b2FrictionCallbackFcn* friction = b2go_FrictionCallbackFcn;
	b2RestitutionCallbackFcn* restitution = b2go_RestitutionCallbackFcn;
	b2CustomFilterCallbackFcn* customFilter = b2go_CustomFilterCallbackFcn;
	b2PreSolveCallbackFcn* preSolve = b2go_PreSolveCallbackFcn;
	b2PreContinuousCallbackFcn* preContinuous = b2go_PreContinuousCallbackFcn;
	b2OverlapResultCallbackFcn* overlapResult = b2go_OverlapResultCallbackFcn;
	b2CastResultCallbackFcn* castResult = b2go_CastResultCallbackFcn;
	b2PlaneResultCallbackFcn* planeResult = b2go_PlaneResultCallbackFcn;
	b2AssertCallbackFcn* assert = b2go_AssertCallbackFcn;
	b2LogCallbackFcn* log = b2go_LogCallbackFcn;

	b2DebugDraw draw = {
		.DrawPolygonFcn = b2go_DrawPolygonFcn,
		.DrawSolidPolygonFcn = b2go_DrawSolidPolygonFcn,
		.DrawCircleFcn = b2go_DrawCircleFcn,
		.DrawSolidCircleFcn = b2go_DrawSolidCircleFcn,
		.DrawSolidCapsuleFcn = b2go_DrawSolidCapsuleFcn,
		.DrawLineFcn = b2go_DrawLineFcn,
		.DrawTransformFcn = b2go_DrawTransformFcn,
		.DrawPointFcn = b2go_DrawPointFcn,
		.DrawStringFcn = b2go_DrawStringFcn,
		.DrawBoundsFcn = b2go_DrawBoundsFcn,
	};
}
