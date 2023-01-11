package main

import (
	"bytes"
	"io"
	"testing"
)

// Test_count tests the count function set to count words
func Test_count(t *testing.T) {
	type args struct {
		r          io.Reader
		countLines bool
		countBytes bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "default",
			args: args{
				r: bytes.NewBufferString("word1 word2 word3 word4\n"),
			},
			want: 4,
		},
		{
			name: "mock a -l flag",
			args: args{
				r:          bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1"),
				countLines: true,
			},
			want: 3,
		},
		{
			name: "mock a -b flag",
			args: args{
				r:          bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1"),
				countBytes: true,
			},
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := count(tt.args.r, tt.args.countLines, tt.args.countBytes); got != tt.want {
				t.Errorf("count() = %v, want %v", got, tt.want)
			}
		})
	}
}
