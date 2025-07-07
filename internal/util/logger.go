// internal/util/logger.go
package util

import (
	"log"
)

var logChan = make(chan string, 100)

// AsyncLog 非阻塞日志写入
func AsyncLog(msg string) {
	select {
	case logChan <- msg:
	default:
		// 丢弃日志，防止阻塞主流程
	}
}

// StartLogger 后台 goroutine 打印日志
func StartLogger() {
	go func() {
		for msg := range logChan {
			log.Println(msg)
		}
	}()
}
