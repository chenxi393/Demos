package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	//reader *kafka.Reader
	topic  = "user_click"
)

func writeKafka(ctx context.Context) {
	writer := kafka.Writer{
		Addr:                   kafka.TCP("localhost:9094"),
		Topic:                  topic,
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: true, // 交给运维的人配
	}
	defer writer.Close()
	for i := 0; i < 3; i++ {
		err := writer.WriteMessages(
			ctx,
			kafka.Message{
				Key:   []byte("1"),
				Value: []byte("你好啊 kfaka"),
			},
			kafka.Message{
				Key:   []byte("2"),
				Value: []byte("你好啊 kfaka"),
			},
			kafka.Message{
				Key:   []byte("3"),
				Value: []byte("你好啊 kfaka"),
			},
			kafka.Message{
				Key:   []byte("4"),
				Value: []byte("你好啊 kfaka"),
			},
			kafka.Message{
				Key:   []byte("5"),
				Value: []byte("你好啊 kfaka"),
			},
		)
		if err != nil {
			if err == kafka.LeaderNotAvailable {
				time.Sleep(1 * time.Second)
			} else {
				fmt.Printf("批量写入kafka失败: %v\n", err)
			}
		} else {
			log.Println("写入成功")
			break
		}
	}
}

func readKafka(ctx context.Context, readTopic string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9094"},
		Topic:          readTopic,
		CommitInterval: 1 * time.Second,
		GroupID:        "team1",
		StartOffset:    kafka.FirstOffset,
	})
	// defer reader.Close()
	// 这里无线循环其实没啥用
	for {
		messages, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Printf("读kafka失败: %v\n", err)
			break
		}
		fmt.Printf("topic = %s, patition= %d ,offset = %d ,key= %s value=%s \n",
			messages.Topic, messages.Partition, messages.Offset, messages.Key, messages.Value)
	}
}

// 监听信号2和15 当收到信号时 关闭reader
// 微服务常用
func listenSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM) //注册2和15
	sig := <-c
	fmt.Printf("接受到信号 %s \n", sig.String())
	//if reader != nil {
	//	reader.Close()
	//}
	os.Exit(0)
}
func main() {
	ctx := context.Background()
	writeKafka(ctx)
	//go listenSignal()
	go readKafka(ctx,"my-topic")
	readKafka(ctx, topic)

	// to produce messages
	// topic := "my-topic"
	// partition := 0

	// conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9094", topic, partition)
	// if err != nil {
	// 	log.Fatal("failed to dial leader:", err)
	// }

	// conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	// _, err = conn.WriteMessages(
	// 	kafka.Message{Value: []byte("one!")},
	// 	kafka.Message{Value: []byte("two!")},
	// 	kafka.Message{Value: []byte("three!")},
	// )
	// if err != nil {
	// 	log.Fatal("failed to write messages:", err)
	// }

	// if err := conn.Close(); err != nil {
	// 	log.Fatal("failed to close writer:", err)
	// }
}
