package jsonz

import (
	"io"
	"strings"
	"testing"
)

type Book struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Authors []string `json:"authors"`
	Press   string   `json:"press"`
}

var jsonData = `{
  "id": "9787111558422",
  "name": "The Go Programming Language",
  "authors": [
    "Alan A.A.Donovan",
    "Brian W. Kergnighan"
  ],
  "press": "Pearson Education"
}`

func TestStringE(t *testing.T) {
	type args[T any] struct {
		body T
	}

	type Test[T any] struct {
		name    string
		args    args[T]
		want    string
		wantErr bool
	}

	var book Book
	_ = UnmarshalStruct([]byte(jsonData), &book)

	tests := []Test[Book]{
		{
			name: "Test String()",
			args: args[Book]{
				body: book,
			},
			want:    "{\"id\":\"9787111558422\",\"name\":\"The Go Programming Language\",\"authors\":[\"Alan A.A.Donovan\",\"Brian W. Kergnighan\"],\"press\":\"Pearson Education\"}",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringE(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("String() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrettyE(t *testing.T) {
	type args[T any] struct {
		body T
	}

	type Test[T any] struct {
		name    string
		args    args[T]
		want    string
		wantErr bool
	}

	var book Book
	_ = UnmarshalStruct([]byte(jsonData), &book)

	tests := []Test[Book]{
		{
			name: "Test Pretty()",
			args: args[Book]{
				body: book,
			},
			want: `{
	"id": "9787111558422",
	"name": "The Go Programming Language",
	"authors": [
		"Alan A.A.Donovan",
		"Brian W. Kergnighan"
	],
	"press": "Pearson Education"
}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrettyE(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pretty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pretty() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshalStruct(t *testing.T) {
	type args struct {
		data   []byte
		target any
	}

	var book Book

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test UnmarshalStruct()",
			args: args{
				data:   []byte(jsonData),
				target: &book,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UnmarshalStruct(tt.args.data, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecodeStruct(t *testing.T) {
	type args struct {
		reader io.Reader
		target any
	}

	var book Book

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test DecodeStruct()",
			args: args{
				reader: strings.NewReader(jsonData),
				target: &book,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DecodeStruct(tt.args.reader, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("DecodeStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
