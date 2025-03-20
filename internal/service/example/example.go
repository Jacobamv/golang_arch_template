package example

import "context"

func (p *provider) Example(ctx context.Context) (err error) {
	// gateway
	err = p.exampleGateway.Ping(ctx)

	//storage

	err = p.exampleStorage.Example(ctx)

	return
}
