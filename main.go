package main

import (
	"goCloud/client"
	"goCloud/server"
)

func main(){
	go server.InitServer()


	client.Client()
}
