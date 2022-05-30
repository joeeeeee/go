package main

import (
	"github.com/joe/iam/internal/apiserver"
)

func main() {
	apiserver.NewApp("iam-apiserver").Run()
}