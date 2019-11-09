package main

import (
	"context"
	"log"
	"fmt"
	"go_p2pb2b"
)

func main() {
	ctx := context.Background()
	client := go_p2pb2b.new_client("https://api.p2pb2b/io/api/v1", "", ctx)
	res, err := client.get_markets()
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Println(res)
}
