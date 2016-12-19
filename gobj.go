package gobj

import (
	"fmt"
	"strconv"
	"strings"
)

// OBJFile describes the content of an OBJ geometry definition file.
type OBJFile struct {
	verts []Vertex
	polys []Polygon
	aabb  AABB
}

// Verts returns the slice of vertices contained in the OBJ file.
func (of OBJFile) Verts() []Vertex {
	return of.verts
}

// Polys returns the slice of polygons contained in the OBJ file.
func (of OBJFile) Polys() []Polygon {
	return of.polys
}

// AABB returns the minimum axis-aligned bouding box containing every vertices
// contained in the OBJ file.
func (of OBJFile) AABB() AABB {
	return of.aabb
}

func (of *OBJFile) parseVertex(kw string, data []string) error {
	v := Vertex{}
	err := v.Set(data)
	if err != nil {
		return err
	}
	of.verts = append(of.verts, v)
	return nil
}

func (of *OBJFile) parseFace(kw string, data []string) error {
	var p Polygon // polygonal face currently filled
	for _, s := range data {
		// read the indices of the face vertices
		sidx := strings.Split(s, "/")[0]
		vidx, err := strconv.Atoi(sidx)
		if err != nil {
			return fmt.Errorf("invalid vertex coordinate value \"%s\"", s)
		}
		p = append(p, of.verts[vidx-1])
	}

	// extend the mesh bounding box with the polygon's one
	of.aabb.extend(p.AABB())
	of.polys = append(of.polys, p)
	return nil
}

// DumpInfo dumps some informations about the OBJ file.
func (of *OBJFile) DumpInfo() string {
	nfo := fmt.Sprintln("num verts:", len(of.verts))
	nfo += fmt.Sprintln("num polys:", len(of.polys))
	nfo += fmt.Sprintln("aabb     :", of.aabb)
	return nfo
}

// updateMin checks if a > b, then a will be set to the value of b.
func updateMin(a *float64, b float64) {
	if b < *a {
		*a = b
	}
}

// updateMax checks if a < b, then a will be set to the value of b.
func updateMax(a *float64, b float64) {
	if *a < b {
		*a = b
	}
}
