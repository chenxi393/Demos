package log

import (
	"io"
	stdlog "log"
	"net/http"
	"os"
)

var log *stdlog.Logger

// 为什么要这样定义 创建新类型
// 因为我们要给fileLog创建方法
// 而基础类型式是不能创建方法的

type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

/*
	为什么filelog可以传进去
	type Writer interface {
	    Write(p []byte) (n int, err error)
	}

因为它实现了io.Write 接口的所有方法 可以赋值的
这个函数是创键Logger写入位置为destination
*/
func Run(destination string) {
	// LstdFlags 就是日期+时间
	log = stdlog.New(fileLog(destination), "[go] - ", stdlog.LstdFlags)
}

func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			msg, err := io.ReadAll(r.Body)
			if err != nil || len(msg) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}

func write(message string) {
	log.Printf("%v\n", message)
}
