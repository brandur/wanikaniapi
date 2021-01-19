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
			IfModifiedSince: wanikaniapi.Time(subjects1.LastModified),
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("last modified 1: %v\n", subjects1.LastModified)
	fmt.Printf("not modified 1: %v\n", subjects1.NotModified)

	fmt.Printf("last modified 1: %v\n", subjects2.LastModified)
	fmt.Printf("not modified 2: %v\n", subjects2.NotModified)
}
