package openapi

var _ refContainer = &MediaType{}

type MediaType struct {
	parent refContainer
	name   string
	Schema *Schema
}

func (m *MediaType) GetName() string {
	return m.parent.GetName() + m.name
}

func (m *MediaType) getBasePath() string {
	return m.parent.getBasePath()
}

func (m *MediaType) getRef() string {
	return ""
}

func (m *MediaType) getChildren() map[string]Traversable {
	return map[string]Traversable{
		"Model": m.Schema,
	}
}

func (m *MediaType) setChild(i string, t Traversable) {
	m.Schema = t.(*Schema)
}
