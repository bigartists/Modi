package informer

import "github.com/google/wire"

// pkg/informer/provider.go
var InformerSet = wire.NewSet(
	NewInformerManager,
)
