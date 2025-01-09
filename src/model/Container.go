package model

type ContainerModel struct {
	Name         string `json:"name"`
	Image        string `json:"image"`
	State        string `json:"state"`
	CreateTime   string `json:"create_time"`
	Ready        string `json:"ready"`
	RestartCount string `json:"restart_count"`
	Message      string `json:"message"`
}
