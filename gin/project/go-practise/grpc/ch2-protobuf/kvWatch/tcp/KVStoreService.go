package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"sync"
	"time"
)

type KVStoreService struct {
	m      map[string]string
	filter map[string]func(string)
	mu     sync.Mutex
}

func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(string)),
	}
}

func (srv *KVStoreService) Get(key string, value *string) error {
	srv.mu.Lock()

	defer srv.mu.Unlock()

	if v, ok := srv.m[key]; ok {
		*value = v
		return nil
	}

	return fmt.Errorf("key not found")
}

func (srv *KVStoreService) Set(kv [2]string, reply *string) error {
	srv.mu.Lock()

	defer srv.mu.Unlock()
	key, value := kv[0], kv[1]

	if oldValue, ok := srv.m[key]; ok {
		if oldValue == value {
			return nil
		}

		for _, fn := range srv.filter {
			fn(value)
		}
	}
	srv.m[key] = value

	*reply = "success"

	return nil
}

const HelloServiceName = "store"

func RegisterService(srv *KVStoreService) error {
	return rpc.RegisterName(HelloServiceName, srv)
}

func main() {
	// 将对象类型中所有满足RPC规则的对象方法注册为RPC函数
	err2 := RegisterService(NewKVStoreService())

	if err2 != nil {
		log.Fatal("register error :", err2)
	}

	// 建立TCP链接
	listen, err := net.Listen("tcp", ":8088")

	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listen.Accept()

		if err != nil {
			log.Fatal("Accept error:", err)
		}

		// 给对方提供RPC服务
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

func (srv *KVStoreService) Watch(timeoutSecond int, keyChange *string) error {

	id := fmt.Sprintf("Watch-%s-%03d", time.Now(), rand.Int())

	ch := make(chan string, 10)

	srv.filter[id] = func(key string) {
		ch <- key
	}

	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return fmt.Errorf("time out")
	case key := <- ch :
		*keyChange = key
		return nil
	}
}
