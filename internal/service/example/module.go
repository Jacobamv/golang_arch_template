package example

import (
	"context"

	storage "github.com/Jacobamv/golang_arch_template/internal/storage/example"
	"github.com/Jacobamv/golang_arch_template/pkg/gateways/example"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewBExample)

type BExample interface {
	Example(ctx context.Context) (err error)
}

type provider struct {
	// Logger

	logger zerolog.Logger

	exampleGateway example.GExample

	exampleStorage storage.SExample
}

type Params struct {
	fx.In

	// Logger
	Logger         zerolog.Logger
	ExampleGateway example.GExample
	ExampleStorage storage.SExample
}

// NewBExample ...
func NewBExample(params Params) BExample {
	return &provider{
		// Logger
		logger: params.Logger,

		exampleGateway: params.ExampleGateway,
		exampleStorage: params.ExampleStorage,
	}
}
