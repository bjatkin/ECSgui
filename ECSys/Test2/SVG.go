package gui

type SVG struct {
	ID         string
	Attributes map[string]string
	Tags       []Tag
}

func (s *SVG) SetAttribute(name, value string) *SVG {
	s.Attributes[name] = value
	return s
}

type Tag struct {
	ID         string
	Type       string
	Attributes map[string]string
	Children   []Tag
	Parent     *Tag
}

func (t *Tag) Append(tags ...Tag) *Tag {
	for _, tag := range tags {
		tag.Parent = t
		t.Children = append(t.Children, tag)
	}
	return t
}
