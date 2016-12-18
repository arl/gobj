package gobj

import "testing"

func check(t *testing.T, e error) {
	if e != nil {
		t.Fatalf("%s", e)
	}
}

func TestLoadOBJFile(t *testing.T) {
	obj, err := Load("testdata/test.obj")
	check(t, err)

	numVerts := len(obj.Vertices())
	numPolys := len(obj.Polygons())
	bb := obj.AABB()

	if numVerts != 897 {
		t.Fatalf("want 897 vertices, got %d", numVerts)
	}
	if numPolys != 1335 {
		t.Fatalf("want 1335 polygons, got %d", numPolys)
	}
	wantAABB := AABB{0.000252, 56.501854,
		-0.010531, 3.000328,
		-0.001253, 32.506012,
	}
	if bb != wantAABB {
		t.Fatalf("want aabb %v, got %v", wantAABB, bb)
	}
}
