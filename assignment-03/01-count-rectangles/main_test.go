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
		{
			name: "Test 3",
			args: args{
				matrix: [][]int{
					{0},
				},
			},
			want: 0,
		},
		{
			name: "Test 4",
			args: args{
				matrix: [][]int{
					{},
				},
			},
			want: 0,
		},
		{
			name: "Test 5",
			args: args{
				matrix: [][]int{
					{1},
				},
			},
			want: 1,
		},
		{
			name: "Test 6",
			args: args{
				matrix: [][]int{
					{1, 1, 1, 1, 1, 1, 1},
				},
			},
			want: 1,
		},
		{
			name: "Test 7",
			args: args{
				matrix: [][]int{
					{1, 0, 0, 0, 0, 0, 0},
					{0, 1, 0, 0, 0, 0, 0},
					{0, 0, 1, 0, 0, 0, 0},
					{0, 0, 0, 1, 0, 0, 0},
					{0, 0, 0, 0, 1, 0, 0},
					{0, 0, 0, 0, 0, 1, 0},
					{0, 0, 0, 0, 0, 0, 1},
				},
			},
			want: 7,
		},
		{
			name: "Test 8",
			args: args{
				matrix: [][]int{
					{1, 1, 1},
					{1, 1, 0},
				},
			},
			want: 1,
		},
		{
			name: "Test 9",
			args: args{
				matrix: [][]int{
					{1, 0, 1},
					{1, 0, 1},
				},
			},
			want: 2,
		},
		{
			name: "Test 10",
			args: args{
				matrix: [][]int{
					{1},
					{1},
					{1},
				},
			},
			want: 1,
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
