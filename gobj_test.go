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
	nfo := obj.DumpInfo()

	want := `num vertices : 897
num triangles: 1335
bounding box : x[0.000252, 56.501854], y[-0.010531, 3.000328], z[-0.001253, 32.506012]
`
	if want != nfo {
		t.Fatalf("want:\n%s\ngot:\n%s\n", want, nfo)
	}
}
