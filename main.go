package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("failed to read number of tests: %v", err)
	}

	numTests, err := parseLine(line, 1)
	if err != nil {
		log.Fatalf("failed to parse number of tests: %v", err)
	}

	// read and execute all tests
	for i := 0; i < numTests[0]; i++ {
		var test Test
		if err := test.Load(reader); err != nil {
			log.Fatalf("failed to load %dth test: %v", i+1, err)
		}

		// run the test
		minHops := test.Grid.MinHops(test.Start, test.End)
		if minHops == -1 {
			fmt.Println("No solution.")
			continue
		}

		fmt.Printf("Optimal solution takes %d hops.\n", minHops)
	}
}
