package main

import (
	"flag"
	"fmt"
	"net/url"

	"github.com/kataras/iris"
	"github.com/kataras/iris/core/host"
	"github.com/kataras/iris/middleware/logger"
)

var (
	port     = flag.Int("port", 80, "port number")
	secure   = flag.Bool("secure", false, "enable SSL")
	certFile = flag.String("cert", "/etc/letsencrypt/live/htapen.com/fullchain.pem", "cert file")
	keyFile  = flag.String("key", "/etc/letsencrypt/live/htapen.com/privkey.pem", "private key file")
)

func main() {
	flag.Parse()
	start()
}

type omniServer struct {
	app *iris.Application
}

func start() {
	rxs := omniServer{}
	rxs.initServer()
}

func (server *omniServer) initServer() {

	app := iris.New()
	server.app = app
	customLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		// Query appends the url query to the Path.
		Query: true,

		//Columns: true,

		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKeys: []string{"logger_message"},

		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		MessageHeaderKeys: []string{"User-Agent"},
	})

	app.Use(customLogger)

	server.initHandlers()

	// Now listening on: http://localhost:{port}
	// Application started. Press CTRL+C to shut down.
	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	if *secure {
		// Redirect 80 request to 443
		target, _ := url.Parse("https://" + addr)
		go host.NewProxy("0.0.0.0:80", target).ListenAndServe()

		fmt.Printf("Serving at :%d with SSL\n", *port)
		app.Run(iris.TLS(addr, *certFile, *keyFile))
	} else {
		// Serve using a host:port form.
		var laddr = iris.Addr(addr)
		app.Run(laddr, iris.WithCharset("UTF-8"))
	}
}

func (server *omniServer) initHandlers() {
	server.app.Get("/{p:path}", handleOtherwise)
}

func handleOtherwise(ctx iris.Context) {
	param := ctx.Params().Get("p")
	ctx.WriteString(param)
}
