package main
 
import (
	"testing"
)
 
var sample = []int{1, 2, 5, -3, 0, 4, -101, 21, 55, -3, 0, 4, 10, 2, 5, -3, 0, 4, 1, 2, 5, -30, 0, 4, 1, 20, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 10, 4, 1, 2, 5, -3, 0, 4, 1, 2, -5, -3, 0, 9, 1, 2, 5, -3, 8, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4}

func BenchmarkMergeSort(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = mergesort(sample)
	}
	b.StopTimer()
}

func BenchmarkBubbleSort(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = bubblesort(sample)
	}
	b.StopTimer()
}

func BenchmarkShakeSort(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = shakesort(sample)
	}
	b.StopTimer()
}

func BenchmarkInsertionSort(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = insertionsort(sample)
	}
	b.StopTimer()
}

func BenchmarkGnomeSort(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = gnomesort(sample)
	}
	b.StopTimer()
}

func BenchmarkMergeSort2(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = mergesort2(sample)
	}
	b.StopTimer()
}

func BenchmarkTreeSort(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = treesort(sample)
	}
	b.StopTimer()
}
