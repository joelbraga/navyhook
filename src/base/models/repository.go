package basemodels


type CRepo struct {
	Name string `json:"name,omitempty"`
	Hooks Hooks `json:"hooks,omitempty"`
}

type Repositories []CRepo

type Hook struct {
	Name string `json:"name,omitempty"`
	Exec bool `json:"exec,omitempty"`
	RemoveFolder bool `json:"remove_folder,omitempty"`
	NavyFolder string `json:"navy_folder,omitempty"`
}

type Hooks []Hook