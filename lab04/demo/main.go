package main

import "example.com/rest-demo/service"

func main() {
	s := service.NewService()
	s.StartWebService()
}
