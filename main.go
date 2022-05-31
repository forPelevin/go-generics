package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

func main() {
	ints := []int{1, 2, 3}
	doubledInts := Map(ints, func(item int) int {
		return item * 2
	})

	// prints [2 4 6]
	fmt.Println(doubledInts)

	words := []string{"hello", "world"}
	capitalizedWords := Map(words, func(item string) string {
		return strings.ToUpper(item)
	})

	// prints [HELLO WORLD]
	fmt.Println(capitalizedWords)

	people := []person{{name: "linda", age: 18}, {name: "john", age: 22}}
	modifiedPeople := Map(people, func(p person) person {
		p.name = strings.Title(p.name)
		p.age += 1

		return p
	})

	// prints [{Linda 19} {John 23}]
	fmt.Println(modifiedPeople)
}

// Map modifies every item of list and returns a new modified slice.
func Map[T any](list []T, modify func(item T) T) []T {
	if list == nil {
		return nil
	}

	if modify == nil {
		return list
	}

	mapped := make([]T, len(list))
	for i, item := range list {
		mapped[i] = modify(item)
	}

	return mapped
}

// MapTyped modifies every item of list and returns a new modified slice. It works only with Integer values.
func MapTyped(list []int, modify func(item int) int) []int {
	if list == nil {
		return nil
	}

	if modify == nil {
		return list
	}

	mapped := make([]int, len(list))
	for i, item := range list {
		mapped[i] = modify(item)
	}

	return mapped
}

// MapAny modifies every item of list and returns a new modified slice. It works with Any type, so you should cast types by yourself.
func MapAny(list []any, modify func(item any) any) []any {
	if list == nil {
		return nil
	}

	if modify == nil {
		return list
	}

	mapped := make([]any, len(list))
	for i, item := range list {
		mapped[i] = modify(item)
	}

	return mapped
}
