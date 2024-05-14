package service

import (
	"github.com/gin-gonic/gin"
	Model "modi/src/model/SecretModel"
	"modi/src/result"
)

type ISecret interface {
	GetSecretByNs(ns string) *result.ErrorResult
	PostSecret(secret *Model.PostSecretModel, c *gin.Context) *result.ErrorResult
}
