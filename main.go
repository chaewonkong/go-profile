package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

const employeeJSON = `
{
	"id": 1,
	"name": "John Doe",
	"age": 30,
	"email": "johndoe@example.com",
	"isEmployed": true
}   
`

// fibRecurr 재귀적으로 Fibonacci 수를 계산: O(2^n)
func fibRecurr(n int) int {
	if n <= 1 {
		return n
	}
	return fibRecurr(n-1) + fibRecurr(n-2)
}

// fibCache cache를 이용해 Fibonacci 수를 계산: O(n)
func fibCache(n int) int {
	// Create a slice to store Fibonacci numbers up to n
	fibCache := make([]int, n+1)

	// Initialize the first two numbers in the Fibonacci sequence
	fibCache[0] = 0
	fibCache[1] = 1

	// Compute Fibonacci numbers up to n, storing each result
	for i := 2; i <= n; i++ {
		fibCache[i] = fibCache[i-1] + fibCache[i-2]
	}

	// Return the nth Fibonacci number
	return fibCache[n]
}

func handler(w http.ResponseWriter, r *http.Request) {
	algorithm := r.PathValue("algorithm")
	p := r.PathValue("num")
	num, err := strconv.Atoi(p)
	if err != nil {
		http.Error(w, "number should be integer string", http.StatusBadRequest)
		return
	}

	if algorithm == "recursive" {
		res := fibRecurr(num)
		fmt.Fprintf(w, "result: %d", res)
		return
	}

	res := fibCache(num)

	fmt.Fprintf(w, "result: %d", res)
}

func main() {
	// Fibonacci calculator
	http.HandleFunc("GET /fib/{algorithm}/{num}", handler)

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
