package main

import (
	"reflect"
	"testing"
)

func TestGetSpeeds(t *testing.T) {
	tests := []struct {
		name         string
		currentSpeed speed
		want         []speed
	}{
		{
			name:         "test1",
			currentSpeed: speed{2, 1},
			want:         []speed{{1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}, {3, 0}, {3, 1}, {3, 2}},
		},
		{
			name:         "test1",
			currentSpeed: speed{0, 0},
			want:         []speed{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSpeeds(tt.currentSpeed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSpeeds() = %v, want %v", got, tt.want)
			}
		})
	}
}
