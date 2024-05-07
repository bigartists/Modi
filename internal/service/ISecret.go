package service

import (
	"github.com/gin-gonic/gin"
	"modi/core/result"
	Model "modi/internal/model/SecretModel"
)

type ISecret interface {
	GetSecretByNs(ns string) *result.ErrorResult
	PostSecret(secret *Model.PostSecretModel, c *gin.Context) *result.ErrorResult
}
