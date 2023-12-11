package main

import (
	"fmt"
	"log"
	"time"
)

// Funcion que calcula la serie de fibonacci
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// clase memoria
type Memory struct {
	f     Function
	cache map[int]FunctionResult
}

// un tipo de function
type Function func(key int) (interface{}, error)

// una clase FunctionResult
type FunctionResult struct {
	value interface{}
	err   error
}

// funcion de nueva cache
func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

// funcion de obtener
func (m *Memory) Get(key int) (interface{}, error) {
	result, exists := m.cache[key]

	if !exists {
		result.value, result.err = m.f(key)
		m.cache[key] = result
	}
	return result.value, result.err

}

func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func mains() {
	cache := NewCache(GetFibonacci)
	fibo := []int{42, 40, 41, 42, 38}
	for _, n := range fibo {
		start := time.Now()
		val, err := cache.Get(n)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%d, %s, %d\n", n, time.Since(start), val)
	}
}
