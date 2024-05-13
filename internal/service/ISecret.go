package service

import (
	"github.com/gin-gonic/gin"
	Model "modi/internal/model/SecretModel"
	"modi/internal/result"
)

type ISecret interface {
	GetSecretByNs(ns string) *result.ErrorResult
	PostSecret(secret *Model.PostSecretModel, c *gin.Context) *result.ErrorResult
}
