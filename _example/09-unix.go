package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/guonaihong/gout"
)

func main() {
	path := "./unix.sock"
	defer os.Remove(path)

	ctx, cancel := context.WithCancel(context.Background())
	srv := server(path)
	defer func() {
		srv.Shutdown(ctx)
		cancel()
	}()

	c := http.Client{}
	s := ""
	err := gout.New(&c).
		Debug(true).
		UnixSocket(path).
		POST("http://xxx/test/unix/").
		SetHeader(gout.H{"h1": "v1", "h2": "v2"}).
		BindBody(&s).Do()

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("result = %s\n", s)
}

func server(path string) *http.Server { _ = "STUB: not implemented"; return nil }
