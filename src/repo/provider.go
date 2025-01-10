package repo

import (
	"github.com/google/wire"
)

// ProviderSetRepo is data providers.
var ProviderSetRepo = wire.NewSet(
	NewIUserGetterImpl,
	ProvideSecretRepo,
	ProviderConfigMapRepo,
	ProviderEventRepo,
	ProviderRsRep,
	ProviderDeploymentRepo,
	ProviderPodRepo,
	ProviderNamespaceRepo,
)
