package tool

// Tool ...
type Tool struct {
	name string
}

// NewTool ...
func NewTool(name string) *Tool {
	return &Tool{
		name: name,
	}
}

func (t Tool) Name() string {
	return t.name
}
