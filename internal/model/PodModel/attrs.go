package PodModel

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

func (this PodAttrFuncs) apply(pod *PodImpl) {
	for _, f := range this {
		f(pod)
	}
}
