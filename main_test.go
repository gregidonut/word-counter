package main

import (
	"bytes"
	"fmt"
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
			want: 35,
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

const (
	foxTextFile = "./testdata/fox.txt"
	jerryMDFile = "./testdata/jerry.md"
)

func Test_actualRun(t *testing.T) {
	type args struct {
		countLines bool
		countBytes bool
		files      []string
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "OneFileNoOtherFlags",
			args: args{
				files: []string{foxTextFile},
			},
			wantW: fmt.Sprintf("%s: %d\ntotal: %d\n", foxTextFile, 9, 9),
		},
		{
			name: "TwoFilesNoOtherFlags",
			args: args{
				files: []string{foxTextFile, jerryMDFile},
			},
			wantW: fmt.Sprintf("%s: %d\n%s: %d\ntotal: %d\n", foxTextFile, 9, jerryMDFile, 10, 19),
		},
		{
			name: "OneFileCountLines",
			args: args{
				files:      []string{foxTextFile},
				countLines: true,
			},
			wantW: fmt.Sprintf("%s: %d\ntotal: %d\n", foxTextFile, 1, 1),
		},
		{
			name: "TwoFilesCountLines",
			args: args{
				files:      []string{foxTextFile, jerryMDFile},
				countLines: true,
			},
			wantW: fmt.Sprintf("%s: %d\n%s: %d\ntotal: %d\n", foxTextFile, 1, jerryMDFile, 3, 4),
		},
		{
			name: "OneFileCountBytes",
			args: args{
				files:      []string{foxTextFile},
				countBytes: true,
			},
			wantW: fmt.Sprintf("%s: %d\ntotal: %d\n", foxTextFile, 44, 44),
		},
		{
			name: "TwoFilesCountBytes",
			args: args{
				files:      []string{foxTextFile, jerryMDFile},
				countBytes: true,
			},
			wantW: fmt.Sprintf("%s: %d\n%s: %d\ntotal: %d\n", foxTextFile, 44, jerryMDFile, 50, 94),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			actualRun(w, tt.args.countLines, tt.args.countBytes, tt.args.files)

			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("actualRun() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
