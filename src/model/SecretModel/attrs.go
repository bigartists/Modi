package SecretModel

type SecretAttrFunc func(*SecretModel)
type SecretAttrFuncs []SecretAttrFunc

func WithName(name string) SecretAttrFunc {
	return func(secret *SecretModel) {
		secret.Name = name
	}
}

func WithCreateTime(createTime string) SecretAttrFunc {
	return func(secret *SecretModel) {
		secret.CreateTime = createTime
	}
}

func WithNamespace(namespace string) SecretAttrFunc {
	return func(secret *SecretModel) {
		secret.Namespace = namespace
	}
}

func WithType(t string) SecretAttrFunc {
	return func(secret *SecretModel) {
		secret.Type = t
	}
}

func (this SecretAttrFuncs) apply(pod *SecretModel) {
	for _, f := range this {
		f(pod)
	}
}
