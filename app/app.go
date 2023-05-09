package app

import (
	"net/http"
	"webserver/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":5000", nil); err != nil {
		panic(err)
	}
}
