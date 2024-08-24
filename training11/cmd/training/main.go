package main

import (
	"context"

	"example.com/training/log"
)

func main() {
	ctx := context.Background()

	log.Demo(ctx)
}
