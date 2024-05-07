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
