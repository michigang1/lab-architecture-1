package main

import "log"

func main() {
	if err := runApplication(); err != nil {
		log.Panicf("Failed to run application: %v", err)
	}
	log.Print("log from imlewel")
	log.Print("log from imlewel 2")
}
