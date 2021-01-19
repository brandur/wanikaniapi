package wanikaniapi_test

import (
	"fmt"
	"os"
	"time"

	"github.com/brandur/wanikaniapi"
)

//
// Note that tests in this file are compiled but not run because they don't
// have the magic `Output` comment.
//

func ExampleClient_initialize() {
	_ = wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
		APIToken: os.Getenv("WANI_KANI_API_TOKEN"),
	})
}

func ExampleClient_makingAPIRequests() {
	client := wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
		APIToken: os.Getenv("WANI_KANI_API_TOKEN"),
	})

	voiceActors, err := client.VoiceActorList(&wanikaniapi.VoiceActorListParams{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("all voice actors: %+v\n", voiceActors)
}

func ExampleClient_settingAPIParameters() {
	client := wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
		APIToken: os.Getenv("WANI_KANI_API_TOKEN"),
	})

	voiceActors, err := client.VoiceActorList(&wanikaniapi.VoiceActorListParams{
		IDs:          []wanikaniapi.WKID{1, 2, 3},
		UpdatedAfter: wanikaniapi.Time(time.Now()),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("all voice actors: %+v\n", voiceActors)
}

func ExampleClient_singlePage() {
	client := wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
		APIToken: os.Getenv("WANI_KANI_API_TOKEN"),
	})

	subjects, err := client.SubjectList(&wanikaniapi.SubjectListParams{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("next page URL: %+v\n", subjects.Pages.NextURL)
}

func ExampleClient_PageFully() {
	client := wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
		APIToken: os.Getenv("WANI_KANI_API_TOKEN"),
	})

	var subjects []*wanikaniapi.Subject
	err := client.PageFully(func(id *wanikaniapi.WKID) (*wanikaniapi.PageObject, error) {
		page, err := client.SubjectList(&wanikaniapi.SubjectListParams{
			ListParams: wanikaniapi.ListParams{
				PageAfterID: id,
			},
		})
		if err != nil {
			return nil, err
		}

		subjects = append(subjects, page.Data...)
		return &page.PageObject, nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("num subjects: %v\n", len(subjects))
}

func ExampleClient_logging() {
	_ = wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
		Logger: &wanikaniapi.LeveledLogger{Level: wanikaniapi.LevelDebug},
	})
}

func ExampleClient_handlingErrors() {
	client := wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
		APIToken: os.Getenv("WANI_KANI_API_TOKEN"),
	})

	_, err := client.SubjectList(&wanikaniapi.SubjectListParams{})
	if err != nil {
		if apiErr, ok := err.(*wanikaniapi.APIError); ok {
			fmt.Printf("WaniKani API error; status: %v, message: %s\n",
				apiErr.StatusCode, apiErr.Message)
		} else {
			fmt.Printf("other error: %+v\n", err)
		}
	}
}

func ExampleClient_conditionalRequestsIfModifiedSince() {
	client := wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
		APIToken: os.Getenv("WANI_KANI_API_TOKEN"),
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

	fmt.Printf("not modified 1: %v\n", subjects1.NotModified)
	fmt.Printf("not modified 2: %v\n", subjects2.NotModified)
}

func ExampleClient_conditionalRequestsIfNoneMatch() {
	client := wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
		APIToken: os.Getenv("WANI_KANI_API_TOKEN"),
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

	fmt.Printf("not modified 1: %v\n", subjects1.NotModified)
	fmt.Printf("not modified 2: %v\n", subjects2.NotModified)
}
