package wanikaniapi_test

import (
	"testing"

	"github.com/brandur/wanikaniapi"
	"github.com/brandur/wanikaniapi/wktesting"
	assert "github.com/stretchr/testify/require"
)

func TestPageFully(t *testing.T) {
	client := wktesting.TestClient()

	var i int
	var subjects []*wanikaniapi.Subject
	err := client.PageFully(func(id *wanikaniapi.ID) (*wanikaniapi.PageObject, error) {
		// Quit after convincing ourselves that we can page by going through a
		// few pages. This saves time and API calls.
		i++
		if i > 2 {
			return nil, nil
		}

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
	assert.NoError(t, err)

	t.Logf("num subjects paged before quitting: %v", len(subjects))
}
