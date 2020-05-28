package vmtype

// Alias hide the real type of the enum
// and users can use it to define the var for accepting enum
type Alias = string

type list struct {
	KVM Alias
	LXC Alias
}

// Enum for public use
var Enum = &list{
	KVM: "KVM",
	LXC: "LXC",
}
