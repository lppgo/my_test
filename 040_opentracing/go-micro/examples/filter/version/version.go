// Package version provides a way of calling a version of a service
package version

import (
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/selector"
	"go-micro.dev/v4/registry"
)

// Filter will filter the version of the service
func Filter(v string) client.CallOption {
	filter := func(services []*registry.Service) []*registry.Service {
		var filtered []*registry.Service

		for _, service := range services {
			if service.Version == v {
				filtered = append(filtered, service)
			}
		}

		return filtered
	}

	return client.WithSelectOption(selector.WithFilter(filter))
}
