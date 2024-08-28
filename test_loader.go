package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Test defines a test cases
type Test struct {
	Grid  Grid
	Start Point
	End   Point
}

// parseLine parses a line into list of n integers
func parseLine(l string, n int) ([]int, error) {
	l = strings.Trim(strings.TrimRight(l, "\n"), " ")
	nums := []int{}

	for i, numberStr := range strings.Split(l, " ") {
		number, err := strconv.ParseUint(numberStr, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("failed to parse %dth number: %w", i+1, err)
		}

		nums = append(nums, int(number))
	}

	if len(nums) != n {
		return nil, errors.New("bad number of tokens found")
	}

	return nums, nil
}

// Load loads a test case
func (t *Test) Load(s *bufio.Reader) error {
	// read and parse the following lines according to definitions
	lineDefs := []struct {
		name    string
		numbers int
	}{
		{"width/height", 2},
		{"start/end", 4},
		{"number of obstacles", 1},
	}

	var nObstacles int

	for n, ldef := range lineDefs {
		line, err := s.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read %s line: %w", ldef.name, err)
		}

		nums, err := parseLine(line, ldef.numbers)
		if err != nil {
			return fmt.Errorf("failed to parse %s line: %w", ldef.name, err)
		}

		switch n {
		case 0:
			t.Grid.Width, t.Grid.Height = nums[0], nums[1]
		case 1:
			t.Start = Point{nums[0], nums[1]}
			t.End = Point{nums[2], nums[3]}
		case 2:
			nObstacles = nums[0]
		}
	}

	// read each obstacle
	for i := 0; i < nObstacles; i++ {
		line, err := s.ReadString('\n')
		if err != nil {
			if i != nObstacles-1 || err != io.EOF {
				return fmt.Errorf("failed to read %dth obstacle: %w", i, err)
			}
		}

		nums, err := parseLine(line, 4)
		if err != nil {
			return fmt.Errorf("failed to parse %dth obstacle: %w", i, err)
		}

		t.Grid.Obstacles = append(t.Grid.Obstacles, Obstacle{nums[0], nums[1], nums[2], nums[3]})
	}

	return nil
}
