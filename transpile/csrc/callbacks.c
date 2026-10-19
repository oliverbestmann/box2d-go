// Signatures of every box2d callback that the Go code implements or calls.
//
// check() assigns each function to its box2d callback type, so the C compiler rejects
// this file when box2d changes a callback. ccgo translates the functions to Go, and
// callbacks.go in the package root checks the Go closures against those translations.

#include "box2d/box2d.h"

void b2go_TaskCallback( int startIndex, int endIndex, uint32_t workerIndex, void* taskContext ) {}
void* b2go_EnqueueTaskCallback( b2TaskCallback* task, int itemCount, int minRange, void* taskContext, void* userContext ) { return 0; }
void b2go_FinishTaskCallback( void* userTask, void* userContext ) {}
float b2go_FrictionCallback( float frictionA, int userMaterialIdA, float frictionB, int userMaterialIdB ) { return 0; }
float b2go_RestitutionCallback( float restitutionA, int userMaterialIdA, float restitutionB, int userMaterialIdB ) { return 0; }
bool b2go_CustomFilterFcn( b2ShapeId shapeIdA, b2ShapeId shapeIdB, void* context ) { return 0; }
bool b2go_PreSolveFcn( b2ShapeId shapeIdA, b2ShapeId shapeIdB, b2Manifold* manifold, void* context ) { return 0; }
bool b2go_OverlapResultFcn( b2ShapeId shapeId, void* context ) { return 0; }
float b2go_CastResultFcn( b2ShapeId shapeId, b2Vec2 point, b2Vec2 normal, float fraction, void* context ) { return 0; }
int b2go_AssertFcn( const char* condition, const char* fileName, int lineNumber ) { return 0; }

void b2go_DrawPolygonFcn( const b2Vec2* vertices, int vertexCount, b2HexColor color, void* context ) {}
void b2go_DrawSolidPolygonFcn( b2Transform transform, const b2Vec2* vertices, int vertexCount, float radius, b2HexColor color, void* context ) {}
void b2go_DrawCircleFcn( b2Vec2 center, float radius, b2HexColor color, void* context ) {}
void b2go_DrawSolidCircleFcn( b2Transform transform, float radius, b2HexColor color, void* context ) {}
void b2go_DrawSolidCapsuleFcn( b2Vec2 p1, b2Vec2 p2, float radius, b2HexColor color, void* context ) {}
void b2go_DrawSegmentFcn( b2Vec2 p1, b2Vec2 p2, b2HexColor color, void* context ) {}
void b2go_DrawTransformFcn( b2Transform transform, void* context ) {}
void b2go_DrawPointFcn( b2Vec2 p, float size, b2HexColor color, void* context ) {}
void b2go_DrawStringFcn( b2Vec2 p, const char* s, b2HexColor color, void* context ) {}

void b2go_check( void )
{
	b2TaskCallback* task = b2go_TaskCallback;
	b2EnqueueTaskCallback* enqueueTask = b2go_EnqueueTaskCallback;
	b2FinishTaskCallback* finishTask = b2go_FinishTaskCallback;
	b2FrictionCallback* friction = b2go_FrictionCallback;
	b2RestitutionCallback* restitution = b2go_RestitutionCallback;
	b2CustomFilterFcn* customFilter = b2go_CustomFilterFcn;
	b2PreSolveFcn* preSolve = b2go_PreSolveFcn;
	b2OverlapResultFcn* overlapResult = b2go_OverlapResultFcn;
	b2CastResultFcn* castResult = b2go_CastResultFcn;
	b2AssertFcn* assert = b2go_AssertFcn;

	b2DebugDraw draw = {
		.DrawPolygonFcn = b2go_DrawPolygonFcn,
		.DrawSolidPolygonFcn = b2go_DrawSolidPolygonFcn,
		.DrawCircleFcn = b2go_DrawCircleFcn,
		.DrawSolidCircleFcn = b2go_DrawSolidCircleFcn,
		.DrawSolidCapsuleFcn = b2go_DrawSolidCapsuleFcn,
		.DrawSegmentFcn = b2go_DrawSegmentFcn,
		.DrawTransformFcn = b2go_DrawTransformFcn,
		.DrawPointFcn = b2go_DrawPointFcn,
		.DrawStringFcn = b2go_DrawStringFcn,
	};
}
