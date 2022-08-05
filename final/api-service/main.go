package main

import "exam-api/service"

func main() {
	s := service.NewService()
	s.StartWebService()
}
