package main

import (
	"log"
	"os"
)

func main() {
	log.SetPrefix("mcpkg: ")
	log.SetFlags(0)

	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("Expected arguments.\n" + os.Args[0] + " server install/search")
	}

	switch args[0] {
		case "server":

		switch args[1] {
			case "install":
				paperInstall(args[2:])
		
			case "search":
				paperSearch(args[2:])

			default:
				log.Fatalln("Expected install/search")
		}

	default:
		log.Fatalln("Expected server/mod")
	}
}
