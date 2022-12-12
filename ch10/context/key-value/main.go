package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	search(ctx, "name")
	ctx = context.WithValue(ctx, "name", "Kut Zhang")
	search(ctx, "name")

	ctx = context.TODO()
	search(ctx, "name")
	ctx = context.WithValue(ctx, "name", "Magick Liu")
	search(ctx, "name")
}

func search(ctx context.Context, key string) {
	v := ctx.Value(key)

	if v != nil {
		fmt.Printf("found key %q: %s\n", key, v)
	} else {
		fmt.Printf("key not found: %q\n", key)
	}
}
