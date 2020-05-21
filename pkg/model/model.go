package model

type VMType interface {
	name() string
	ordinal() int
	values() *[]string
}

type VMTypes uint

const (
	LXC VMTypes = iota
	KVM
)

var vmTypesStrings = []string{
	"LXC",
	"KVM",
}

func (t VMTypes) name() string {
	return vmTypesStrings[t]
}

func (t VMTypes) ordinal() int {
	return int(t)
}

func (t VMTypes) values() *[]string {
	return &vmTypesStrings
}
