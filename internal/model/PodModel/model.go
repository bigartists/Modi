package PodModel

type PodImpl struct {
	Name           string
	Images         string
	NodeName       string
	Namespace      string
	DeploymentName string
	CreateTime     string
	Phase          string // pod 当前所属的阶段
	Message        string
	IsReady        bool // 判断pod 是否就绪
}

func New(attrs ...PodAttrFunc) *PodImpl {
	pod := &PodImpl{}
	PodAttrFuncs(attrs).apply(pod)
	return pod
}
