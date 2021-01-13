package wanikaniapi_test

import (
	"testing"

	"github.com/brandur/wanikaniapi"
	"github.com/brandur/wanikaniapi/wktesting"
	assert "github.com/stretchr/testify/require"
)

func TestReviewListAndGet(t *testing.T) {
	client := wktesting.TestClient()
	var sampleID wanikaniapi.ID

	{
		reviews, err := client.ReviewList(&wanikaniapi.ReviewListParams{})
		assert.NoError(t, err)
		assert.Greater(t, len(reviews.Data), 0)

		sampleID = reviews.Data[0].ID
	}

	{
		review, err := client.ReviewGet(&wanikaniapi.ReviewGetParams{ID: &sampleID})
		assert.NoError(t, err)
		assert.NotNil(t, review)
	}
}
