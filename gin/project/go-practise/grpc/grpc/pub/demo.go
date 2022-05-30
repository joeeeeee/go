package main

import (
	"fmt"
	"github.com/moby/moby/pkg/pubsub"
	"strings"
	"time"
)

func mains() {
	publisher := pubsub.NewPublisher(100*time.Microsecond, 10)

	golang := publisher.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "golang:") {
				return true
			}
		}
		return false
	})

	docker := publisher.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "docker:") {
				return true
			}
		}
		return false
	})

	go func() {
		publisher.Publish("hi")
	}()

	go func() {
		publisher.Publish("golang: hello golang")
	}()

	go func() {
		publisher.Publish("docker: hello docker")
	}()

	time.Sleep(1)

	go func() {
		fmt.Println("golang topic:", <-golang)
	}()

	go func() {
		fmt.Println("docker topic:", <-docker)
	}()

	time.Sleep(10 * time.Second)
}
