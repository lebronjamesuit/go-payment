package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const app_version = "1.0.0"
const css_version = "1.1"

// input param: no
// output param: error
// Type pointer applicaiton has method server()
func (app *application) serve() error {

	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", app.config.port),
		Handler:     app.routes(),
		IdleTimeout: 30 * time.Second,
	}

	app.inforLog.Println("Starting server now")

	return srv.ListenAndServe()

}

func main() {

	fmt.Print("hello main")

	var config configuration

	flag.IntVar(&config.port, "port", 4002, "usage port")
	flag.StringVar(&config.env, "env", "dev enviroment", "usage env")
	flag.StringVar(&config.api, "api", "http://localhost:4002", "usage api")

	flag.Parse()

	config.stripe.key = os.Getenv("STRIPE_KEY")
	config.stripe.secrect = os.Getenv("STRIPE_SECRET")

	inforLog := log.New(os.Stdout, "LOG_PREFIX start server port 4002 ok", 1)

	// create a instance
	app := &application{
		config:   config,
		inforLog: inforLog,
		version:  app_version,
	}

	error := app.serve()
	if error != nil {
		fmt.Println("Bi loi roi")
	}

}

type configuration struct {
	port int
	env  string
	api  string

	stripe struct {
		secrect string
		key     string
	}
}

type application struct {
	config   configuration
	inforLog *log.Logger

	version string
}
