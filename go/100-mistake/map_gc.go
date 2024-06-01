package main

import (
	"fmt"
	"runtime"
)

func MapGC() {
	// Init
	n := 1_000_000
	m := make(map[int][128]byte)
	printAllocMB()

	// Add elements
	for i := 0; i < n; i++ {
		m[i] = randBytes()
	}
	printAllocMB()

	// Remove elements
	for i := 0; i < n; i++ {
		delete(m, i)
	}

	// End
	runtime.GC()
	printAllocMB()
	runtime.KeepAlive(m)
}

func randBytes() [128]byte {
	return [128]byte{}
}

func printAllocMB() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}
 

/*
当删除了所有 kv 后，内存占用依然有 293 MB，这实际上是创建长度为 100w 的 map 所消耗的内存大小。map的buckets创建后就不会缩容。map中的buckets只会增加不会减少。
- 当 val 大小 <= 128B 时，val 其实是直接放在 bucket 里的。 会把 bmap 标记为不含指针，这样可以避免 gc 时扫描整个 hmap。
- 当 val 大小 >= 128B 时，val会转为指针存储。
解决：
- 定期拷贝map到另一个map；
- 将 val 类型改成指针，可以减少buckets占用的内存；


*/