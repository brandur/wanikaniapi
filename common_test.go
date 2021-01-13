package wanikaniapi_test

import (
	"testing"

	"github.com/brandur/wanikaniapi"
	"github.com/brandur/wanikaniapi/wktesting"
	assert "github.com/stretchr/testify/require"
)

func TestPageFully(t *testing.T) {
	client := wktesting.TestClient()

	var assignments []*wanikaniapi.Assignment
	err := client.PageFully(func(id *wanikaniapi.ID) (*wanikaniapi.PageObject, error) {
		page, err := client.AssignmentList(&wanikaniapi.AssignmentListParams{
			ListParams: &wanikaniapi.ListParams{
				PageAfterID: id,
			},
		})
		if err != nil {
			return nil, err
		}

		assignments = append(assignments, page.Data...)
		return page.PageObject, nil
	})
	assert.NoError(t, err)

	t.Logf("num assignments = %v", len(assignments))
}
