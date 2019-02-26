package main

import (
	"flag"
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

var (
	port   = flag.Int("port", 8080, "port number")
	secure = flag.Bool("secure", false, "enable SSL")
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
	if *secure {
		fmt.Printf("Serving at :%d with SSL\n", *port)
		app.Run(iris.AutoTLS(fmt.Sprintf("0.0.0.0:%d", *port), "htapen.com", "pengxiao@outlook.com"))
	} else {
		// Serve using a host:port form.
		var laddr = iris.Addr(fmt.Sprintf("0.0.0.0:%d", *port))
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
