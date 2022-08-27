/*
 * Copyright 2022 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package stringz

import (
	"testing"
)

func Test_implode(t *testing.T) {
	type args struct {
		alice     []string
		separator string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test string implode",
			args: args{
				alice:     []string{"1", "2", "3"},
				separator: "&",
			},
			want: "1&2&3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := implode(tt.args.alice, tt.args.separator); got != tt.want {
				t.Errorf("implode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringBuffer_ToSortStrings(t *testing.T) {
	type fields struct {
		buffer []string
	}
	type args struct {
		separator string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Test StringBuffer ToSortStrings",
			fields: fields{
				buffer: []string{
					"product_id=389238",
					"user_id=29389",
					"content=newproductmask",
					"environment=test",
				},
			},
			args: args{
				separator: "&",
			},
			want: "content=newproductmask&environment=test&product_id=389238&user_id=29389",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := StringBuffer{
				buffer: tt.fields.buffer,
			}
			if got := sb.ToSortStrings(tt.args.separator); got != tt.want {
				t.Errorf("ToSortStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringBuffer_ToStrings(t *testing.T) {
	type fields struct {
		buffer []string
	}
	type args struct {
		separator string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Test StringBuffer ToSortStrings",
			fields: fields{
				buffer: []string{
					"product_id=389238",
					"user_id=29389",
					"content=newproductmask",
					"environment=test",
				},
			},
			args: args{
				separator: "&",
			},
			want: "product_id=389238&user_id=29389&content=newproductmask&environment=test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := StringBuffer{
				buffer: tt.fields.buffer,
			}
			if got := sb.ToStrings(tt.args.separator); got != tt.want {
				t.Errorf("ToStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringBuffer_Length(t *testing.T) {
	type fields struct {
		buffer []string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Test StringBuffer length",
			fields: fields{
				buffer: []string{"Hello", "World"},
			},
			want: 2,
		},
		{
			name: "Test StringBuffer length",
			fields: fields{
				buffer: make([]string, 0),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := StringBuffer{
				buffer: tt.fields.buffer,
			}
			if got := sb.Length(); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}
