package main

import "testing"

func Benchmark_Print(b *testing.B) {
	InitSeverIndex()
	b.ResetTimer() // 重置时间 忽略前面的过程

	for i := 0; i < b.N; i++ {
		Select()
	}
}

func Benchmark_Print_Parallel(b *testing.B) {
	InitSeverIndex()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Select()
		}
	})
}

func Benchmark_Fast_Print(b *testing.B) {
	InitSeverIndex()
	b.ResetTimer() // 重置时间 忽略前面的过程

	for i := 0; i < b.N; i++ {
		FastSelect()
	}
}

func Benchmark_Fast_Print_Parallel(b *testing.B) {
	InitSeverIndex()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			FastSelect()
		}
	})
}

func Benchmark_Fib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(10)
	}
}

func Benchmark_NoPreAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoPreAlloc(50)
	}
}

func Benchmark_PreAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PreAlloc(50)
	}
}
