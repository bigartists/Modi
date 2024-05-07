package service

import (
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"modi/client"
	"modi/core/result"
	"modi/internal/kind"
	Model "modi/internal/model/SecretModel"
	"modi/pkg/utils"
)

var SecretServiceGetter ISecret

type ISecretServiceGetterImpl struct {
}

func (I ISecretServiceGetterImpl) PostSecret(secret *Model.PostSecretModel, c *gin.Context) *result.ErrorResult {
	_, err := client.K8sClient.CoreV1().Secrets(secret.Namespace).Create(
		c,
		&v1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      secret.Name,
				Namespace: secret.Namespace,
			},
			Type:       v1.SecretType(secret.Type),
			StringData: secret.Data,
		},
		metav1.CreateOptions{},
	)

	if err != nil {
		return result.Result(false, err)
	} else {
		return result.Result(true, nil)
	}
}

func (I ISecretServiceGetterImpl) GetSecretByNs(ns string) *result.ErrorResult {
	var list []*v1.Secret

	if ns == "" {
		list = kind.SecretMapInstance.ListAll()
	} else {
		list = kind.SecretMapInstance.ListAllByNs(ns)

	}

	var ret []*Model.SecretModel
	for _, item := range list {
		ret = append(ret, Model.New(
			Model.WithName(item.Name),
			Model.WithNamespace(item.Namespace),
			Model.WithCreateTime(
				utils.FormatTime(item.CreationTimestamp),
			),
			Model.WithType(Model.SECRET_TYPE[string(item.Type)]),
		))
	}

	return result.Result(ret, nil)
}

func init() {
	SecretServiceGetter = &ISecretServiceGetterImpl{}
}
