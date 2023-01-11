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
			name: "mock a true flag",
			args: args{
				r:          bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1"),
				countLines: true,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := count(tt.args.r, tt.args.countLines); got != tt.want {
				t.Errorf("count() = %v, want %v", got, tt.want)
			}
		})
	}
}
