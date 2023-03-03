package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"strings"
)

const PROJECTID string = "petecloudfilesystem"

var ctx = context.Background()

func authenticate() *storage.Client {
	//Create client
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Printf("Failed to create gcp client: %v\n", err)
		return nil
	}
	return client
}

func createBucket(name string, client *storage.Client) int {
	//Choose name
	//TODO: BETTER WAY TO CREATE THE NAME
	var sb strings.Builder
	sb.WriteString("petecfs-")
	sb.WriteString(name)
	sb.WriteString("-00000000")
	bkt := client.Bucket(sb.String())

	//Check if it exists
	_, err := bkt.Attrs(ctx)
	if err == nil {
		fmt.Println("Bucket already exists")
		return -1
	}

	//Create bucket
	if err = bkt.Create(ctx, PROJECTID, nil); err != nil {
		fmt.Printf("Error: %v\n", err)
		return -1
	}
	return 0
}

func main() {
	fmt.Println("ola mundo")

	var client *storage.Client

	if client = authenticate(); client == nil {
		return
	}

	if err := createBucket("ubuntu", client); err == -1 {
		return
	}

	fmt.Println("Bucket created")

	//TODO: Close the client
}
