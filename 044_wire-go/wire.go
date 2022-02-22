//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"wire-go/provider"
)

// InitializeEvent 声明injector的函数签名
func InitializeEvent(msg string) provider.Event {
	// wire.Build(provider.NewEvent, provider.NewGreeter, provider.NewMessage)

	wire.Build(EventSet) // 组合provider

	return provider.Event{} //返回值没有实际意义，只需符合函数签名即可
}

// EventSet Event通常是一起使用的一个集合，使用wire.NewSet进行组合
var EventSet = wire.NewSet(provider.NewEvent, provider.NewMessage, provider.NewGreeter)
