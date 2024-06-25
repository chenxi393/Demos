package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func ESTest() {
	// 配置 Elasticsearch 客户端
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://:9200",
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 测试连接
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// 打印 Elasticsearch 信息
			fmt.Printf("Elasticsearch Info: %s\n", r["version"].(map[string]interface{})["number"])
		}
	}

	wg := sync.WaitGroup{}
	var err1, err2, err3 int32

	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for index := range 10 {
				// 要插入的文档
				doc := map[string]interface{}{
					"title":   strconv.FormatInt(int64(index), 10) + "Test Document",
					"content": "egyaubdiakbdoabkduthsthhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhabljdaudiavdyyavdvi",
					"id":      index,
					"id1":     index,
					"id2":     index,
					"id3":     index,
					"id4":     index,
					"id5":     index,
				}

				// 将文档编码为 JSON
				var buf bytes.Buffer
				if err := json.NewEncoder(&buf).Encode(doc); err != nil {
					log.Fatalf("Error encoding document: %s", err)
				}

				// 创建索引请求
				req := esapi.IndexRequest{
					Index:   "hello",
					Body:    strings.NewReader(buf.String()),
					Refresh: "true",
				}

				// 执行请求
				res, err = req.Do(context.Background(), es)
				if err != nil {
					atomic.AddInt32(&err1, 1)
					log.Printf("Error getting response: %s", err)
					continue
				}
				defer res.Body.Close()

				// 检查响应
				if res.IsError() {
					atomic.AddInt32(&err2, 1)
					log.Printf("[%s] Error indexing document ID=%d", res.Status(), 1)
				} else {
					var r map[string]interface{}
					if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
						atomic.AddInt32(&err3, 1)
					} else {
						// 打印响应结果
						fmt.Printf("Document indexed with ID=%v\n", r["_id"])
					}
				}

			}
		}()
	}
	wg.Wait()

	fmt.Println("err1: ", err1, "err2: ", err2, "err3: ", err3)
}
