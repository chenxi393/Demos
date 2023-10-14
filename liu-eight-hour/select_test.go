package main


import(
	"fmt"
	"time"
)
func fetchData(ch <-chan int) {
    select {
    case data := <-ch:
        fmt.Println("Received data:", data)
    case <-time.After(2 * time.Second):
        fmt.Println("Timeout occurred")
    }
}

func main() {
    ch := make(chan int)

    go func() {
        time.Sleep(2 * time.Second)
        ch <- 10
    }()

    fetchData(ch)

    time.Sleep(3 * time.Second) // 等待goroutine执行完毕
}

/*
在上述例子中，如果从channel接收数据的操作在3秒内未完成，就会执行time.After返回的通道操作，提示超时。

需要注意的是，select语句只能用于channel的操作，不能用于普通的数据类型。

此外，还需要注意以下几点：

如果有多个channel操作准备就绪，select会随机选择一个case执行，因此不要依赖于选择的顺序。
select语句会阻塞等待，因此需要确保至少有一个channel操作可以触发，否则可能导致死锁。
如果没有default语句，且所有case都未准备就绪，select语句会阻塞等待所有的case中的某个操作就绪。

*/
