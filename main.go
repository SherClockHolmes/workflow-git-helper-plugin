package main

import (
	"flag"
	"log"

	"github.com/SherClockHolmes/workflow-git-helper-plugin/git"
)

func main() {
	branchFlag := flag.String("branch", "", "Checkout git respository on this branch")
	flag.Parse()

	// Get first command line argument
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("Missing first argument: repository")
	}

	gc, err := git.New(args[0], *branchFlag)
	if err != nil {
		log.Fatal(err)
	}

	stdout, err := gc.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(stdout))
}
