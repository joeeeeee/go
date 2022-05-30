package iam_pump

import (
	"math/rand"
	"time"
)

func main()  {
	rand.Seed(time.Now().UTC().UnixNano())
	pump
}