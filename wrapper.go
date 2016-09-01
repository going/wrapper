/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          wrapper.go
 * Description:   handlers wrapper
 */

package wrapper

import (
	"net/http"

	"context"
)

type ContextHandler interface {
	ServeHTTPContext(context.Context, http.ResponseWriter, *http.Request)
}

type ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

func (h ContextHandlerFunc) ServeHTTPContext(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
	h(ctx, rw, req)
}

type ContextAdapter struct {
	Context context.Context
	Handler ContextHandler
}

func (ca *ContextAdapter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ca.Handler.ServeHTTPContext(ca.Context, rw, req)
}

func Handler(ctx context.Context, handler ContextHandlerFunc) *ContextAdapter {
	return &ContextAdapter{
		Context: ctx,
		Handler: handler,
	}
}

func Wrapper(handler ContextHandlerFunc) *ContextAdapter {
	ctx := context.Background()
	return &ContextAdapter{
		Context: ctx,
		Handler: handler,
	}
}
