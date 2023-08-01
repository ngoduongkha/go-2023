package main

import "testing"

func Test_countRectangles(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				matrix: [][]int{
					{1, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{1, 0, 0, 1, 1, 1, 0},
					{0, 1, 0, 1, 1, 1, 0},
					{0, 1, 0, 0, 0, 0, 0},
					{0, 1, 0, 1, 1, 0, 0},
					{0, 0, 0, 1, 1, 0, 0},
					{0, 0, 0, 0, 0, 0, 1},
				},
			},
			want: 6,
		},
		{
			name: "Test 2",
			args: args{
				matrix: [][]int{
					{1, 0, 0, 0, 0, 0, 0},
					{1, 0, 0, 0, 0, 0, 0},
					{1, 0, 0, 1, 1, 1, 0},
					{0, 1, 0, 1, 1, 1, 0},
					{0, 1, 0, 0, 0, 0, 0},
					{0, 1, 0, 1, 1, 0, 0},
					{0, 0, 0, 1, 1, 0, 0},
					{0, 0, 0, 0, 0, 0, 1},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRectangles(tt.args.matrix); got != tt.want {
				t.Errorf("countRectangles() = %v, want %v", got, tt.want)
			}
		})
	}
}
