package pkger

import (
	"context"

	"go-micro.dev/v4/config/source"
)

type pkgerPathKey struct{}

// WithPath sets the path to pkger
func WithPath(p string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, pkgerPathKey{}, p)
	}
}
