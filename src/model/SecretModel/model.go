package SecretModel

var SECRET_TYPE map[string]string

func init() {
	SECRET_TYPE = map[string]string{
		"Opaque":                              "自定义类型",
		"kubernetes.io/service-account-token": "服务账号令牌",
		"kubernetes.io/dockercfg":             "docker配置",
		"kubernetes.io/dockerconfigjson":      "docker配置(JSON)",
		"kubernetes.io/basic-auth":            "Basic认证凭据",
		"kubernetes.io/ssh-auth":              " SSH凭据",
		"kubernetes.io/tls":                   "TLS凭据",
		"bootstrap.kubernetes.io/token":       "启动引导令牌数据",
	}
}

type SecretModel struct {
	Name       string
	Namespace  string
	CreateTime string
	Type       string
}

type PostSecretModel struct {
	Name      string
	Namespace string
	Type      string
	Data      map[string]string
}

func New(attrs ...SecretAttrFunc) *SecretModel {
	secret := &SecretModel{}
	SecretAttrFuncs(attrs).apply(secret)
	return secret
}

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
