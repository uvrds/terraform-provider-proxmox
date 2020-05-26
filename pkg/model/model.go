package model

//type virtualMachine struct {
//	name   string
//	vmType vmtype.Alias
//}

type ProviderConfig struct {
	ipamEndpoints []string
	ipamUsername  string
	ipamPassword  string
}
