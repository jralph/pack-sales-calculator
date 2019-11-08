package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"pack-sales-calculator/calculator"
)

func main() {
	port := flag.Int("port", 9787, "The port to run the server on.")
	flag.Parse()

	mode := flag.Args()[0]

	if mode == "cli" {
		runAsCliTool(flag.Args()[0:])
	}

	if mode == "api" {
		runAsWebApi(flag.Args()[0:], *port)
	}
}

func runAsCliTool(args []string) {
	orderString := args[1]
	availablePacks := args[2]

	order, err := strconv.Atoi(orderString)
	if err != nil {
		panic(err)
	}

	packStrings := strings.Split(availablePacks, ",")

	var packs []int
	for _, pack := range packStrings {
		size, err := strconv.Atoi(pack)
		if err != nil {
			panic(err)
		}

		packs = append(packs, size)
	}

	fmt.Println(calculator.PackCalculator(order, packs))
}

func runAsWebApi(args []string, port int) {
	
}