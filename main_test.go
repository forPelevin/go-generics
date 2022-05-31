package main

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name   string
		list   []any
		modify func(item any) any
		want   []any
	}{
		{
			name: "nil list",
			want: nil,
		},
		{
			name:   "nil modify func",
			list:   []any{1, 2, 3},
			modify: nil,
			want:   []any{1, 2, 3},
		},
		{
			name: "empty list",
			list: []any{},
			modify: func(item any) any {
				return item
			},
			want: []any{},
		},
		{
			name: "list of ints",
			list: []any{1, 2, 3},
			modify: func(item any) any {
				return item.(int) * 2
			},
			want: []any{2, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.list, tt.modify); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapTyped(t *testing.T) {
	tests := []struct {
		name   string
		list   []int
		modify func(item int) int
		want   []int
	}{
		{
			name: "nil list",
			want: nil,
		},
		{
			name:   "nil modify func",
			list:   []int{1, 2, 3},
			modify: nil,
			want:   []int{1, 2, 3},
		},
		{
			name: "empty list",
			list: []int{},
			modify: func(item int) int {
				return item
			},
			want: []int{},
		},
		{
			name: "list of ints",
			list: []int{1, 2, 3},
			modify: func(item int) int {
				return item * 2
			},
			want: []int{2, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapTyped(tt.list, tt.modify); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapTyped() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapAny(t *testing.T) {
	tests := []struct {
		name   string
		list   []any
		modify func(item any) any
		want   []any
	}{
		{
			name: "nil list",
			want: nil,
		},
		{
			name:   "nil modify func",
			list:   []any{1, 2, 3},
			modify: nil,
			want:   []any{1, 2, 3},
		},
		{
			name: "empty list",
			list: []any{},
			modify: func(item any) any {
				return item
			},
			want: []any{},
		},
		{
			name: "list of ints",
			list: []any{1, 2, 3},
			modify: func(item any) any {
				return item.(int) * 2
			},
			want: []any{2, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapAny(tt.list, tt.modify); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGenericMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Map([]int{1, 2, 3}, func(item int) int {
			return item * 2
		})
	}
}

func BenchmarkTypedMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MapTyped([]int{1, 2, 3}, func(item int) int {
			return item * 2
		})
	}
}

func BenchmarkAnyMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MapAny([]any{1, 2, 3}, func(item any) any {
			return item.(int) * 2
		})
	}
}
