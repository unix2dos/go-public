//go:build wireinject
// +build wireinject

package example

import "github.com/google/wire"

func GetPostService() (*PostService, func(), error) {
	panic(wire.Build(
		NewPostService,
		wire.Build(new(IPostUsecase), new(*postUsecase)),
		NewPostUsecase,
		NewPostRepo,
	))
}
