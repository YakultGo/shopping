package util

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

// GetFreePort 获取可用的端口号
func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func GetRandomCode() string {
	src := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(src)
	code := rand.Intn(1000000)
	return fmt.Sprintf("%06d", code)
}
