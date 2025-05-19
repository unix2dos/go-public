//go:build wireinject
// +build wireinject

package wire

import "github.com/google/wire"

func GetPostService() (*PostService, func(), error) {
	panic(wire.Build(
		NewPostService,
		wire.Build(new(IPostUsecase), new(*postUsecase)),
		NewPostUsecase,
		NewPostRepo,
	))
}
