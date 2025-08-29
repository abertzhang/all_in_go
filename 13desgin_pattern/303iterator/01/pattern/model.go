package pattern
type Part struct {
	Title string
	Number int
}


type PartCollection struct {
	Part
	Parts []*Part
}
func (p *PartCollection) CreateIterator() IIterator {
	return &PartIterator{Parts: p.Parts}
}

type PartIterator struct {
Index int
Parts []*Part
}
func (p *PartIterator) HasNext() bool {
	return p.Index < len(p.Parts)
}
func (p *PartIterator) Next() interface{} {
	if p.HasNext() {
		part := p.Parts[p.Index]
		p.Index++
		return part
	}
	return nil
}