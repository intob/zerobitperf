package zerobitperf

import (
	"crypto/rand"
	"testing"
)

const N = 1000000 // 1M randomized tests
const L = 32      // Number of random bytes to test
const Z = 8       // Fixed number of leading zero bytes

func makeTestdata() [][]byte {
	td := make([][]byte, N)
	for i := 0; i < N; i++ {
		td[i] = make([]byte, L)
		rand.Read(td[i])
		for j := 0; j < Z; j++ {
			td[i][j] = 0
		}
	}
	return td
}

func Benchmark_1(b *testing.B) {
	td := makeTestdata()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nzerobit_1(td[i%N])
	}
}

func Benchmark_2(b *testing.B) {
	td := makeTestdata()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nzerobit_2(td[i%N])
	}
}

func Benchmark_3(b *testing.B) {
	td := makeTestdata()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nzerobit_3(td[i%N])
	}
}

func Benchmark_4(b *testing.B) {
	td := makeTestdata()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nzerobit_4(td[i%N])
	}
}

func Benchmark_5(b *testing.B) {
	td := makeTestdata()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nzerobit_5(td[i%N])
	}
}
