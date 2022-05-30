package main

import (
	"github.com/joe/iam/internal/authzserver"
)

func main() {
	authzserver.NewApp("iam-authz-server").Run()
}
