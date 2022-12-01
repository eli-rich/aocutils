package main

import "fmt"

func ArrayMap[A any, R any](arr []A, fn func(A) R) []R {
	// Applies the supplied function to every item in the array
	result := make([]R, len(arr))
	for i, v := range arr {
		result[i] = fn(v)
	}
	return result
}

func ArrayFilter[A any](arr []A, fn func(A) bool) []A {
	// Filters the array by checking the supplied function
	var result []A
	for _, v := range arr {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func ArrayReduce[A any, R any](arr []A, fn func(R, A) R, initial R) R {
	// Reduces the array to a single value dictated by the supplied function and the initial value
	var result R = initial
	for _, v := range arr {
		result = fn(result, v)
	}
	return result
}

func ArrayFind[A any](arr []A, fn func(A) bool) A {
	// Finds the first item in the array matching the specified function
	var result A
	for _, v := range arr {
		if fn(v) {
			return v
		}
	}
	return result
}

func ArrayFindIndex[A any](arr []A, fn func(A) bool) int {
	// Finds the index of the first item in the array matching the specified function
	for i, v := range arr {
		if fn(v) {
			return i
		}
	}
	return -1
}

func ArrayEvery[A any](arr []A, fn func(A) bool) bool {
	// Checks if every element in the array passed the specified function
	for _, v := range arr {
		if !fn(v) {
			return false
		}
	}
	return true
}

func ArraySome[A any](arr []A, fn func(A) bool) bool {
	// Checks if any element in the array passes the specified function
	for _, v := range arr {
		if fn(v) {
			return true
		}
	}
	return false
}

func ArrayReverse[A any](arr []A) []A {
	// Returns a new array that has been reversed
	length := len(arr) - 1
	result := make([]A, length+1)
	for i, j := length, 0; i > -1; i, j = i-1, j+1 {
		fmt.Println(j)
		result[j] = arr[i]
	}
	return result
}
