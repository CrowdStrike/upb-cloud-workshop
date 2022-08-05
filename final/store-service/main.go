package main

import "exam-store/service"

func main() {
	s := service.NewService()
	s.StartWebService()
}
