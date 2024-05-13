package service

import (
	"github.com/gin-gonic/gin"
	Model "modi/pkg/model/SecretModel"
	"modi/pkg/result"
)

type ISecret interface {
	GetSecretByNs(ns string) *result.ErrorResult
	PostSecret(secret *Model.PostSecretModel, c *gin.Context) *result.ErrorResult
}
