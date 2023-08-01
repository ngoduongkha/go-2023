package main

import "testing"

func Test_countDifferentIntegers(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				word: "a123bc34d8ef34",
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				word: "leet1234code234",
			},
			want: 2,
		},
		{
			name: "Test 3",
			args: args{
				word: "a1b01c001",
			},
			want: 1,
		},
		{
			name: "Test 4",
			args: args{
				word: "as8c6as7d67a",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDifferentIntegers(tt.args.word); got != tt.want {
				t.Errorf("countDifferentIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
