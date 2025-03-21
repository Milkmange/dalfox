package utils

import (
	"testing"

	"github.com/hahwul/dalfox/v2/pkg/model"
)

func Test_IndexOf(t *testing.T) {
	type args struct {
		element string
		data    []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Element found",
			args: args{element: "b", data: []string{"a", "b", "c"}},
			want: 1,
		},
		{
			name: "Element not found",
			args: args{element: "d", data: []string{"a", "b", "c"}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexOf(tt.args.element, tt.args.data); got != tt.want {
				t.Errorf("IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DuplicatedResult(t *testing.T) {
	type args struct {
		result []model.PoC
		rst    model.PoC
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Duplicate found",
			args: args{
				result: []model.PoC{{Type: "type1"}, {Type: "type2"}},
				rst:    model.PoC{Type: "type1"},
			},
			want: true,
		},
		{
			name: "Duplicate not found",
			args: args{
				result: []model.PoC{{Type: "type1"}, {Type: "type2"}},
				rst:    model.PoC{Type: "type3"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DuplicatedResult(tt.args.result, tt.args.rst); got != tt.want {
				t.Errorf("DuplicatedResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ContainsFromArray(t *testing.T) {
	type args struct {
		slice []string
		item  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Item found",
			args: args{slice: []string{"a", "b", "c"}, item: "b"},
			want: true,
		},
		{
			name: "Item not found",
			args: args{slice: []string{"a", "b", "c"}, item: "d"},
			want: false,
		},
		{
			name: "Item found with parentheses",
			args: args{slice: []string{"a", "b", "c"}, item: "b(something)"},
			want: true,
		},
		{
			name: "Item not found with parentheses",
			args: args{slice: []string{"a", "b", "c"}, item: "d(something)"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsFromArray(tt.args.slice, tt.args.item); got != tt.want {
				t.Errorf("ContainsFromArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CheckPType(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid type",
			args: args{str: "validType"},
			want: true,
		},
		{
			name: "Invalid type toBlind",
			args: args{str: "toBlind"},
			want: false,
		},
		{
			name: "Invalid type toGrepping",
			args: args{str: "toGrepping"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPType(tt.args.str); got != tt.want {
				t.Errorf("CheckPType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IsAllowType(t *testing.T) {
	type args struct {
		contentType string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Allowed type - text/html",
			args: args{
				contentType: "text/html",
			},
			want: true,
		},
		{
			name: "Not allowed type - application/json",
			args: args{
				contentType: "application/json",
			},
			want: false,
		},
		{
			name: "Not allowed type - text/javascript",
			args: args{
				contentType: "text/javascript",
			},
			want: false,
		},
		{
			name: "Allowed type with charset - text/html; charset=UTF-8",
			args: args{
				contentType: "text/html; charset=UTF-8",
			},
			want: true,
		},
		{
			name: "Not allowed type with charset - application/json; charset=UTF-8",
			args: args{
				contentType: "application/json; charset=UTF-8",
			},
			want: false,
		},
		{
			name: "Allowed type - application/xml",
			args: args{
				contentType: "application/xml",
			},
			want: true,
		},
		{
			name: "Not allowed type - image/jpeg",
			args: args{
				contentType: "image/jpeg",
			},
			want: false,
		},
		{
			name: "Not allowed type - image/png",
			args: args{
				contentType: "image/png",
			},
			want: false,
		},
		{
			name: "Not allowed type - text/plain",
			args: args{
				contentType: "text/plain",
			},
			want: false,
		},
		{
			name: "Not allowed type - text/css",
			args: args{
				contentType: "text/css",
			},
			want: false,
		},
		{
			name: "Not allowed type - application/rss+xml",
			args: args{
				contentType: "application/rss+xml",
			},
			want: false,
		},
		{
			name: "Allowed type - text/xml",
			args: args{
				contentType: "text/xml",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAllowType(tt.args.contentType); got != tt.want {
				t.Errorf("isAllowType() = %v, want %v", got, tt.want)
			}
		})
	}
}
