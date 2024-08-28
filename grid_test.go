package main

import (
	"reflect"
	"testing"
)

// prepare a list of test grids to use in multiple tests
var testGrids = map[string]struct {
	grid   Grid
	visual string
}{
	"no obstacles 5": {
		grid: Grid{
			Width:  5,
			Height: 5,
		},
		visual: "" +
			".....\n" +
			".....\n" +
			".....\n" +
			".....\n" +
			".....\n",
	},
	"no obstacles 10": {
		grid: Grid{
			Width:  10,
			Height: 10,
		},
		visual: "" +
			"..........\n" +
			"..........\n" +
			"..........\n" +
			"..........\n" +
			"..........\n" +
			"..........\n" +
			"..........\n" +
			"..........\n" +
			"..........\n" +
			"..........\n",
	},
	"all obstacles": {
		grid: Grid{
			Width:     5,
			Height:    5,
			Obstacles: []Obstacle{{0, 4, 0, 4}},
		},
		visual: "" +
			"XXXXX\n" +
			"XXXXX\n" +
			"XXXXX\n" +
			"XXXXX\n" +
			"XXXXX\n",
	},
	"obstacle1": {
		grid: Grid{
			Width:     5,
			Height:    5,
			Obstacles: []Obstacle{{0, 2, 0, 2}},
		},
		visual: "" +
			"XXX..\n" +
			"XXX..\n" +
			"XXX..\n" +
			".....\n" +
			".....\n",
	},
	"obstacle2": {
		grid: Grid{
			Width:  5,
			Height: 5,
			Obstacles: []Obstacle{
				{4, 4, 0, 0},
				{1, 2, 1, 2},
				{3, 4, 3, 3},
				{0, 0, 4, 4},
			},
		},
		visual: "" +
			"....X\n" +
			".XX..\n" +
			".XX..\n" +
			"...XX\n" +
			"X....\n",
	},
	"obstacles 10 / 1": {
		grid: Grid{
			Width:  10,
			Height: 10,
			Obstacles: []Obstacle{
				{2, 4, 1, 1},
				{0, 9, 3, 3},
				{1, 9, 4, 4},
				{2, 6, 5, 6},
				{3, 6, 7, 7},
				{0, 0, 8, 8},
				{5, 9, 8, 8},
				{0, 0, 6, 6},
				{0, 0, 7, 7},
			},
		},
		visual: "" +
			"..........\n" +
			"..XXX.....\n" +
			"..........\n" +
			"XXXXXXXXXX\n" +
			".XXXXXXXXX\n" +
			"..XXXXX...\n" +
			"X.XXXXX...\n" +
			"X..XXXX...\n" +
			"X....XXXXX\n" +
			"..........\n",
	},
	"obstacles 10 / 2": {
		grid: Grid{
			Width:  10,
			Height: 10,
			Obstacles: []Obstacle{
				{1, 8, 1, 8},
			},
		},
		// 0 0, 9 9
		visual: "" +
			"..........\n" +
			".XXXXXXXX.\n" +
			".XXXXXXXX.\n" +
			".XXXXXXXX.\n" +
			".XXXXXXXX.\n" +
			".XXXXXXXX.\n" +
			".XXXXXXXX.\n" +
			".XXXXXXXX.\n" +
			".XXXXXXXX.\n" +
			"..........\n",
	},
	"example 7 hops": {
		// from the PDF example
		// 5 5
		// 4 0 4 4
		// 1
		// 1 4 2 3
		grid: Grid{
			Width:  5,
			Height: 5,
			Obstacles: []Obstacle{
				{1, 4, 2, 3},
			},
		},
		visual: "" +
			".....\n" +
			".....\n" +
			".XXXX\n" +
			".XXXX\n" +
			".....\n",
	},
}

func TestGrid_Print(t *testing.T) {
	for name, tt := range testGrids {
		t.Run(name, func(t *testing.T) {
			if got := tt.grid.Print(); got != tt.visual {
				t.Errorf("Grid.Print() =\n%v, want =\n%v", got, tt.visual)
			}
		})
	}
}

func TestGrid_MinHops(t *testing.T) {
	type args struct {
		start Point
		end   Point
	}
	tests := []struct {
		name string
		grid Grid
		args args
		want int
	}{
		{
			name: "no obstacles 5/1",
			grid: testGrids["no obstacles 5"].grid,
			args: args{
				start: Point{0, 0},
				end:   Point{4, 4},
			},
			want: 3,
		},
		{
			name: "no obstacles 5/2",
			grid: testGrids["no obstacles 5"].grid,
			args: args{
				start: Point{0, 0},
				end:   Point{3, 3},
			},
			want: 2,
		},
		{
			name: "no obstacles 10/1",
			grid: testGrids["no obstacles 10"].grid,
			args: args{
				start: Point{0, 0},
				end:   Point{9, 9},
			},
			want: 4,
		},
		{
			name: "all obstacles",
			grid: testGrids["all obstacles"].grid,
			args: args{
				start: Point{0, 0},
				end:   Point{3, 3},
			},
			want: -1,
		},
		{
			name: "obstacle2 / 1",
			grid: testGrids["obstacle2"].grid,
			args: args{
				start: Point{0, 0},
				end:   Point{4, 4},
			},
			want: 4,
		},
		{
			name: "obstacle2 / 2",
			grid: testGrids["obstacle2"].grid,
			args: args{
				start: Point{0, 0},
				end:   Point{0, 3},
			},
			want: 2,
		},
		{
			name: "obstacles 10 / 1",
			grid: testGrids["obstacles 10 / 1"].grid,
			args: args{
				start: Point{0, 0},
				end:   Point{0, 9},
			},
			want: 6,
		},
		{
			name: "obstacles 10 / 2",
			grid: testGrids["obstacles 10 / 2"].grid,
			args: args{
				start: Point{0, 0},
				end:   Point{9, 9},
			},
			want: 8,
		},
		{
			name: "example 7 hops",
			grid: testGrids["example 7 hops"].grid,
			args: args{
				start: Point{4, 0},
				end:   Point{4, 4},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.grid.Height == 0 && tt.grid.Width == 0 {
				t.Errorf("grid is empty")
			}
			if got := tt.grid.MinHops(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grid.MinHops() = %v, want %v", got, tt.want)
			}
		})
	}
}
