# wrapper
Negroni Context Handler Wrapper(Golang's context.Context)

**Example:**

```
#!go

package main

import (
    "fmt"
    "net/http"
    "time"

    "golang.org/x/net/context"

    "github.com/going/wrapper"
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "github.com/stretchr/graceful"
)

func main() {
    router := NewRouter("v1")
    n := negroni.New()
    n.Use(negroni.NewLogger())
    n.UseHandler(router)
    fmt.Println("-> Starting ....")
    graceful.Run(":8085", 10*time.Second, n)
}

func NewRouter(version string) *mux.Router {
    r := mux.NewRouter()
    v := r.PathPrefix(fmt.Sprintf("/%s", version)).Subrouter()
    v.Methods("GET").Path("/ping").Handler(wrapper.Wrapper(ping))

    ctx := context.Background()
    ctx = context.WithValue(ctx, "PONG", 5)
    v.Methods("GET").Path("/pong").Handler(wrapper.Handler(ctx, pong))

    return v
}

func ping(ctx context.Context, w http.ResponseWriter, r *http.Request) {
    ctx = ctx.WithValue(ctx, "PING", 10)
    ...
}

func pong(ctx context.Context, w http.ResponseWriter, r *http.Request) {
    value := ctx.Value("PONG").(int)
    ...
}
```
