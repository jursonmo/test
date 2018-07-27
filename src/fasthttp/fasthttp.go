package main

import(
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
	"github.com/golang/snappy"
)

type MyHandler struct {
	foobar string
}

// request handler in net/http style, i.e. method bound to MyHandler struct.
func (h *MyHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	// notice that we may access MyHandler properties here - see h.foobar.
	fmt.Fprintf(ctx, "Hello, world! Requested path is %q. Foobar is %q",
		ctx.Path(), h.foobar)
}

// request handler in fasthttp style, i.e. just plain function.
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
}

func main() {
// pass bound struct method to fasthttp
	myHandler := &MyHandler{
		foobar: "foobar",
	}
	fasthttp.ListenAndServe(":8080", myHandler.HandleFastHTTP)

	// pass plain function to fasthttp
	fasthttp.ListenAndServe(":8081", fastHTTPHandler)

	time.Sleep(1)

	fmt.Println("source len:", len(test))

    got := snappy.Encode(nil, []byte(test))
    fmt.Println("compressed len:", len(got))

    a, _ := snappy.Decode(nil, got)
    fmt.Println("uncompressed len:", len(a))
}


const test = `{"Ag(T+D)":{"instID":"Ag(T+D)","name":"白银延期","last":"4141","upDown":"21","upDownRate":"0.51","quoteDate":"20170328"
	,"quoteTime":"22:34:29"},"Au(T+D)":{"instID":"Au(T+D)","name":"黄金延期","last":"280.55","upDown":"0.88","upDownRate":"0.31"
		,"quoteDate":"20170328","quoteTime":"22:34:15"},"mAu(T+D)":{"instID":"mAu(T+D)","name":"Mini黄金延期","last":"280.5","upDown":"0.7"
			,"upDownRate":"0.25","quoteDate":"20170328","quoteTime":"22:34:10"}}`
