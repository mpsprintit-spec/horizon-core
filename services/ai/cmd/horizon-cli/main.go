package main

import (
	"fmt"
	"os"

	"github.com/project-horizon/horizon-core/services/ai/temporary/websearch"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("===================================")
		fmt.Println(" Horizon CLI")
		fmt.Println("===================================")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  horizon learn <word>")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "learn":

		if len(os.Args) < 3 {
			fmt.Println("Usage:")
			fmt.Println("  horizon learn <word>")
			os.Exit(1)
		}

		word := os.Args[2]

		if err := websearch.LearnWord(word); err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(1)
		}

	default:

		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)

	}

}
