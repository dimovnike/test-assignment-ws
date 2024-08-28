package main

import (
	"slices"
	"testing"
)

func TestQueue_Deqeue(t *testing.T) {
	tests := []struct {
		name string
		q    Queue[int]
		want []int
	}{
		{
			name: "test1",
		},
		{
			name: "test2",
			q: func() Queue[int] {
				var q Queue[int]
				q.Queue(3, 2, 1)
				return q
			}(),
			want: []int{3, 2, 1},
		},
		{
			name: "test3",
			q: func() Queue[int] {
				var q Queue[int]
				q.Queue(3, 2, 1)
				q.Deqeue()
				return q
			}(),
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int
			valid := true

			for {
				var v int
				v, valid = tt.q.Deqeue()
				if !valid {
					break
				}

				got = append(got, v)
			}

			if !slices.Equal(got, tt.want) {
				t.Errorf("Queue.Deqeue() got = %v, want %v", got, tt.want)
			}
		})
	}
}
