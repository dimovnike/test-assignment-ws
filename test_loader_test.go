package main

import (
	"bufio"
	"bytes"
	"reflect"
	"testing"
)

func TestTest_Load(t *testing.T) {
	tests := []struct {
		name    string
		reader  *bufio.Reader
		want    Test
		wantErr bool
	}{
		{
			name:    "empty buf",
			reader:  bufio.NewReader(&bytes.Buffer{}),
			wantErr: true,
		},
		{
			name: "case 1 no obstacles",
			reader: func() *bufio.Reader {
				buf := bytes.Buffer{}
				buf.WriteString("" +
					"5 5\n" +
					"1 2 3 4\n" +
					"0\n",
				)
				return bufio.NewReader(&buf)
			}(),
			want: Test{
				Grid: Grid{
					Width:  5,
					Height: 5,
				},
				Start: Point{1, 2},
				End:   Point{3, 4},
			},
		},
		{
			name: "case 2, 1 obstacle, from example",
			reader: func() *bufio.Reader {
				buf := bytes.Buffer{}
				buf.WriteString("" +
					"5 5\n" +
					"4 0 4 4\n" +
					"1\n" +
					"1 4 2 3\n",
				)
				return bufio.NewReader(&buf)
			}(),
			want: Test{
				Grid: Grid{
					Width:  5,
					Height: 5,
					Obstacles: []Obstacle{
						{1, 4, 2, 3},
					},
				},
				Start: Point{4, 0},
				End:   Point{4, 4},
			},
		},
		{
			name: "case 2, 2 obstacles",
			reader: func() *bufio.Reader {
				buf := bytes.Buffer{}
				buf.WriteString("" +
					"6 7\n" +
					"0 0 1 1\n" +
					"2\n" +
					"1 1 2 2\n" +
					"3 3 4 4\n",
				)
				return bufio.NewReader(&buf)
			}(),
			want: Test{
				Grid: Grid{
					Width:  6,
					Height: 7,
					Obstacles: []Obstacle{
						{1, 1, 2, 2},
						{3, 3, 4, 4},
					},
				},
				Start: Point{0, 0},
				End:   Point{1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Test{}
			if err := tr.Load(tt.reader); (err != nil) != tt.wantErr {
				t.Errorf("Test.Load() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tr, tt.want) {
				t.Errorf("Test.Load() = %v, want %v", tr, tt.want)
			}
		})
	}
}
