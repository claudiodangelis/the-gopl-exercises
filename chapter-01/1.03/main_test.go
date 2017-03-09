package main

import "testing"

func BenchmarkEcho1(b *testing.B) {
	args := []string{"hello", "world", "how", "is", "it", "going", "?", "I", "hope", "everthing", "is", "good"}
	for i := 0; i < b.N; i++ {
		echo1(args)
	}
}

func BenchmarkEcho2(b *testing.B) {
	args := []string{"hello", "world", "how", "is", "it", "going", "?", "I", "hope", "everthing", "is", "good"}
	for i := 0; i < b.N; i++ {
		echo2(args)
	}
}
