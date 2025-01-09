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

type PodAttrFunc func(*PodImpl)
type PodAttrFuncs []PodAttrFunc

func WithName(name string) PodAttrFunc {
	return func(pod *PodImpl) {
		pod.Name = name
	}
}

func WithImages(images string) PodAttrFunc {
	return func(pod *PodImpl) {
		pod.Images = images
	}
}

func WithNodeName(nodeName string) PodAttrFunc {
	return func(pod *PodImpl) {
		pod.NodeName = nodeName
	}
}

func WithCreateTime(createTime string) PodAttrFunc {
	return func(pod *PodImpl) {
		pod.CreateTime = createTime
	}
}

func WithPhase(phase string) PodAttrFunc {
	return func(pod *PodImpl) {
		pod.Phase = phase
	}
}

func WithMessage(message string) PodAttrFunc {
	return func(pod *PodImpl) {
		pod.Message = message
	}
}

func WithIsReady(isReady bool) PodAttrFunc {
	return func(pod *PodImpl) {
		pod.IsReady = isReady
	}
}

func WithNamespace(namespace string) PodAttrFunc {
	return func(pod *PodImpl) {
		pod.Namespace = namespace
	}
}

func WithDeploymentName(deploymentName string) PodAttrFunc {
	return func(pod *PodImpl) {
		pod.DeploymentName = deploymentName
	}
}

func (this PodAttrFuncs) apply(pod *PodImpl) {
	for _, f := range this {
		f(pod)
	}
}
