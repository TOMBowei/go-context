package main

import (
	"context"
	"fmt"
)

func step1(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "name", "小波")
	return child
}

func step2(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "age", 18)
	return child
}

func step3(ctx context.Context) {
	fmt.Printf("name %s\n", ctx.Value("name"))
	fmt.Printf("age %d\n", ctx.Value("age"))
}

func main() {
	grandpa := context.TODO()
	father := step1(grandpa)
	grandson := step2(father)
	step3(grandson)
}
