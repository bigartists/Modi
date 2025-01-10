package controllers

import "github.com/google/wire"

var ProviderSetCtrl = wire.NewSet(
	//NewUserController,
	NewAuthController,
	NewSecretController,
	NewConfigMapController,
)
