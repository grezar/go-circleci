package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/grezar/go-circleci"
)

func main() {
	config := &circleci.Config{
		Token: os.Getenv("CIRCLE_TOKEN"),
	}

	client, err := circleci.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	pipelines, err := client.Pipelines.List(context.Background(), circleci.PipelineListOptions{
		OrgSlug: circleci.String("gh/grezar"),
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, pipeline := range pipelines.Items {
		fmt.Println(pipeline)
	}
}
