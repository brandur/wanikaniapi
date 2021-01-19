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
		IDs:          []wanikaniapi.ID{1, 2, 3},
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
	err := client.PageFully(func(id *wanikaniapi.ID) (*wanikaniapi.PageObject, error) {
		page, err := client.SubjectList(&wanikaniapi.SubjectListParams{
			ListParams: &wanikaniapi.ListParams{
				PageAfterID: id,
			},
		})
		if err != nil {
			return nil, err
		}

		subjects = append(subjects, page.Data...)
		return page.PageObject, nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("num subjects: %v", len(subjects))
}

func ExampleClient_logging() {
	_ = wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
		Logger: &wanikaniapi.LeveledLogger{Level: wanikaniapi.LevelDebug},
	})
}
