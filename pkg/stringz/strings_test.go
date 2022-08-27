package stringz

import (
	"reflect"
	"testing"
)

func TestArrayContains(t *testing.T) {
	type args struct {
		haystack []string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test string array contains-true",
			args: args{
				haystack: []string{"lilei", "hanmeimei"},
				needle:   "lilei",
			},
			want: true,
		},
		{
			name: "Test string array contains-false",
			args: args{
				haystack: []string{"lilei", "hanmeimei"},
				needle:   "tom",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayContains(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("ArrayContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloneSlice(t *testing.T) {
	type args struct {
		src []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test clone slice",
			args: args{
				src: []string{"Hello", "World"},
			},
			want: []string{"Hello", "World"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CloneSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CloneSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImplode(t *testing.T) {
	type args struct {
		haystack  []string
		separator string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test clone slice",
			args: args{
				haystack:  []string{"Hello", "World"},
				separator: ",",
			},
			want: "Hello,World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Implode(tt.args.haystack, tt.args.separator); got != tt.want {
				t.Errorf("Implode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExplode(t *testing.T) {
	type args struct {
		haystack  string
		separator string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test clone slice",
			args: args{
				haystack:  "Hello,World",
				separator: ",",
			},
			want: []string{"Hello", "World"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Explode(tt.args.haystack, tt.args.separator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Explode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceTemplate(t *testing.T) {
	type args struct {
		template string
		args     []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test replace template",
			args: args{
				template: "hello, %s",
				args:     []any{"world"},
			},
			want: "hello, world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceTemplate(tt.args.template, tt.args.args...); got != tt.want {
				t.Errorf("ReplaceTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
