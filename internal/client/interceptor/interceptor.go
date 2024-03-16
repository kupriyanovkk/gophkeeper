package interceptor

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Interceptor struct {
	protectedMethods map[string]bool
}

// Unary returns a unary client interceptor.
//
// It takes in context.Context, method string, req and reply interface{}, *grpc.ClientConn, grpc.UnaryInvoker, and ...grpc.CallOption as parameters and returns an error.
func (i *Interceptor) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		if i.protectedMethods[method] {
			if md, ok := metadata.FromOutgoingContext(ctx); ok {
				if len(md.Get("authorization")) == 0 {
					return errors.New("authorization required")
				}
			}
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// NewInterceptor creates a new Interceptor with the given protected methods.
//
// It takes in a map of protected methods and returns a pointer to Interceptor.
func NewInterceptor(protectedMethods map[string]bool) *Interceptor {
	return &Interceptor{
		protectedMethods: protectedMethods,
	}
}
