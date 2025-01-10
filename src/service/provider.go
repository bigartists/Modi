package service

import "github.com/google/wire"

var ProviderSetService = wire.NewSet(
	NewUserServiceImpl,
	NewSecretService,
	// wire.Bind(接口指针, 具体类型指针)
	// 这行代码告诉 wire：当需要 ISecret 接口时可以使用 *SecretService 类型来满足这个需求因为 *SecretService 实现了 ISecret 接口
	//wire.Bind(new(ISecret), new(*SecretService)),
	ProviderPodService,
	ProviderConfigmapService,
	ProviderDeploymentService,
)
