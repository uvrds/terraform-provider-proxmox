package client

func (api *API) GetStatusVM(node string, id string) error {
	path := "/nodes/" + node + "/qemu/" + id + "/status/current"
	err := api.get(path, nil)
	if err != nil {
		return err
	}
	return nil
}

type Lxc struct {
	VMID       string
	Ostemplate string
	Storage    string
	Node       string
	Name       string
}

func (api *API) CreateLxc(data Lxc) error {

	options := map[string]string{

		"ostemplate": data.Ostemplate,
		"vmid":       data.VMID,
		"storage":    data.Storage,
	}

	path := "/nodes/" + data.Node + "/lxc"
	err := api.post(path, options)
	if err != nil {
		return err
	}
	return nil
}
