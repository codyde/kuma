package registry

import "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/resources/model"

var global = NewTypeRegistry()

func Global() TypeRegistry {
	return global
}

func RegisterType(res model.Resource) {
	if err := global.RegisterType(res); err != nil {
		panic(err)
	}
}

func RegistryListType(resList model.ResourceList) {
	if err := global.RegisterListType(resList); err != nil {
		panic(err)
	}
}