package PodModel

type PodImpl struct {
	Name       string
	Images     string
	NodeName   string
	CreateTime string
	Phase      string // pod 当前所属的阶段
	Message    string
}

func New(attrs ...PodAttrFunc) *PodImpl {
	pod := &PodImpl{}
	PodAttrFuncs(attrs).apply(pod)
	return pod
}
