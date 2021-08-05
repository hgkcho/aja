package fooooo

// Foo ...
type Foo struct {
	name string
}

// NewFoo ...
func NewFoo(name string) *Foo {
	return &Foo{
		name: name,
	}
}

func (t Foo) Name() string {
	return t.name
}
