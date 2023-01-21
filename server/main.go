package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/metatexx/go-app-todo/frontend"
	"log"
	"net/http"
)

func main() {
	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", &frontend.ToDoList{})

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	http.Handle("/", &app.Handler{
		Name:        "todo List",
		Description: "An example todo App",
		Styles:      []string{"/web/styles.css"},
		Image:       "/web/icon-512.png",
		Icon: app.Icon{
			Default:    "/web/icon-192.png",
			Large:      "/web/icon-512.png",
			AppleTouch: "/web/icon-192.png",
		},
	})
	fmt.Println("Listening on http://127.0.0.1:8001")
	if err := http.ListenAndServe("127.0.0.1:8001", nil); err != nil {
		log.Fatal(err)
	}
}
