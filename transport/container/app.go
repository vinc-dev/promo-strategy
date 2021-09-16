package container

import "github.com/vinc-dev/promo-strategy/transport/http/container"

// AppContainer handle all requirement for app to run properly
type AppContainer struct {
	Middleware *container.MiddlewareContainer
	Services   *ServiceContainer
}

// CreateAppContainer construct all requirement for app
func CreateAppContainer() *AppContainer {
	return &AppContainer{
		Middleware: container.CreateMiddlewareContainer(),
		Services:   CreateServiceContainer(CreateRepositoryContainer()),
	}
}
