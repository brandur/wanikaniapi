package wanikaniapi_test

import (
	"testing"

	"github.com/brandur/wanikaniapi"
	"github.com/brandur/wanikaniapi/wktesting"
	assert "github.com/stretchr/testify/require"
)

func TestPageFullyReviews(t *testing.T) {
	client := wktesting.TestClient()

	var reviews []*wanikaniapi.Review
	err := client.PageFully(func(id *wanikaniapi.ID) (*wanikaniapi.PageObject, error) {
		page, err := client.ListReviews(&wanikaniapi.ReviewListParams{
			ListParams: &wanikaniapi.ListParams{
				PageAfterID: id,
			},
		})
		if err != nil {
			return nil, err
		}

		reviews = append(reviews, page.Data...)
		return page.PageObject, nil
	})
	assert.NoError(t, err)

	t.Logf("num reviews = %v", len(reviews))
}
