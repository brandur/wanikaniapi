package main

import (
	"fmt"
	"os"

	"github.com/brandur/wanikaniapi"
)

func main() {
	client := wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
		APIToken: os.Getenv("WANI_KANI_API_TOKEN"),
		Logger:   &wanikaniapi.LeveledLogger{Level: wanikaniapi.LevelDebug},
	})

	subjects1, err := client.SubjectList(&wanikaniapi.SubjectListParams{})
	if err != nil {
		panic(err)
	}

	subjects2, err := client.SubjectList(&wanikaniapi.SubjectListParams{
		Params: wanikaniapi.Params{
			IfNoneMatch: wanikaniapi.String(subjects1.ETag),
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("etag 1: %v\n", subjects1.ETag)
	fmt.Printf("not modified 1: %v\n", subjects1.NotModified)

	fmt.Printf("etag 1: %v\n", subjects2.ETag)
	fmt.Printf("not modified 2: %v\n", subjects2.NotModified)
}
