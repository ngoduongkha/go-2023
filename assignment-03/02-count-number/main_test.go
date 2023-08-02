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
		{
			name: "Test 5",
			args: args{
				word: "abc123def456",
			},
			want: 2,
		},
		{
			name: "Test 6",
			args: args{
				word: "1a2b3c4d5e6",
			},
			want: 6,
		},
		{
			name: "Test 7",
			args: args{
				word: "12ab34cd56",
			},
			want: 3,
		},
		{
			name: "Test 8",
			args: args{
				word: "leading0s0123456",
			},
			want: 2,
		},
		{
			name: "Test 9",
			args: args{
				word: "001002003004005006",
			},
			want: 1,
		},
		{
			name: "Test 10",
			args: args{
				word: "abc",
			},
			want: 0,
		},
		{
			name: "Test 11",
			args: args{
				word: "",
			},
			want: 0,
		},
		{
			name: "Test 12",
			args: args{
				word: "          ",
			},
			want: 0,
		},
		{
			name: "Test 13",
			args: args{
				word: "  123 2 sx 21 c212x",
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
