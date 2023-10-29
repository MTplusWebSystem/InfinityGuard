package main

import (
	"fmt"
	"infinityguard/routes"
)

func main () {
	const port = "8080"
	fmt.Printf("Servidor em execução \n em localhost:%s", port)
	routes.StartServer(port)
}


