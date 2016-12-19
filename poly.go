package gobj

// Polygon represents a polygonal face-element.
type Polygon []Vertex

// Scale applies f scale factor to every coord of the polygon vertices.
func (p *Polygon) Scale(f float64) {
	for i := range *p {
		(*p)[i].Scale(f)
	}
}

// AABB computes and returns the axis-aligned bounding-box
// of the polygon.
func (p *Polygon) AABB() AABB {
	bb := NewAABB()
	for _, v := range *p {
		updateMin(&bb.MinX, v.X())
		updateMin(&bb.MinY, v.Y())
		updateMin(&bb.MinZ, v.Z())
		updateMax(&bb.MaxX, v.X())
		updateMax(&bb.MaxY, v.Y())
		updateMax(&bb.MaxZ, v.Z())
	}
	return bb
}
