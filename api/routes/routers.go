package routes

import (
	"infinityguard/headle"
	"net/http"
)



func StartServer (port string){
	http.HandleFunc("/", headle.Pong)
	http.ListenAndServe(":"+port, nil)
}