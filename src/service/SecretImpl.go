package service

import (
	"github.com/bigartists/Modi/client"
	"github.com/bigartists/Modi/src/model/SecretModel"
	"github.com/bigartists/Modi/src/repo"
	"github.com/bigartists/Modi/src/result"
	"github.com/bigartists/Modi/src/utils"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ISecret interface {
	GetSecretByNs(ns string) *result.ErrorResult
	PostSecret(secret *SecretModel.PostSecretModel, c *gin.Context) *result.ErrorResult
}

type SecretService struct {
	secretRepo *repo.SecretRepo
}

func NewSecretService(secretRepo *repo.SecretRepo) ISecret {
	return &SecretService{secretRepo: secretRepo}
}

func (this *SecretService) PostSecret(secret *SecretModel.PostSecretModel, c *gin.Context) *result.ErrorResult {
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

func (this *SecretService) GetSecretByNs(ns string) *result.ErrorResult {
	var list []*v1.Secret

	if ns == "" {
		list = this.secretRepo.ListAll()
	} else {
		list = this.secretRepo.ListAllByNs(ns)

	}

	var ret []*SecretModel.SecretModel
	for _, item := range list {
		ret = append(ret, SecretModel.New(
			SecretModel.WithName(item.Name),
			SecretModel.WithNamespace(item.Namespace),
			SecretModel.WithCreateTime(
				utils.FormatTime(item.CreationTimestamp),
			),
			SecretModel.WithType(SecretModel.SECRET_TYPE[string(item.Type)]),
		))
	}

	return result.Result(ret, nil)
}
