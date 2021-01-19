package wanikaniapi_test

import (
	"net/http"
	"testing"

	"github.com/brandur/wanikaniapi"
	"github.com/brandur/wanikaniapi/wktesting"
	assert "github.com/stretchr/testify/require"
)

func TestReviewCreate(t *testing.T) {
	client := wktesting.LocalClient()

	_, err := client.ReviewCreate(&wanikaniapi.ReviewCreateParams{
		IncorrectMeaningAnswers: wanikaniapi.Int(2),
		SubjectID:               wanikaniapi.ID(123),
	})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, `{"review":{"incorrect_meaning_answers":2,"subject_id":123}}`, string(req.Body))
	assert.Equal(t, http.MethodPost, req.Method)
	assert.Equal(t, "/v2/reviews", req.Path)
	assert.Equal(t, "", req.Query)
}

func TestReviewList(t *testing.T) {
	client := wktesting.LocalClient()

	_, err := client.ReviewList(&wanikaniapi.ReviewListParams{
		AssignmentIDs: []wanikaniapi.WKID{1, 2, 3},
		SubjectIDs:    []wanikaniapi.WKID{4, 5, 6},
	})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, "", string(req.Body))
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/v2/reviews", req.Path)
	assert.Equal(t, "assignment_ids=1,2,3&subject_ids=4,5,6", wktesting.MustQueryUnescape(req.Query))
}

func TestReviewGet(t *testing.T) {
	client := wktesting.LocalClient()

	_, err := client.ReviewGet(&wanikaniapi.ReviewGetParams{ID: wanikaniapi.ID(123)})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, "", string(req.Body))
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/v2/reviews/123", req.Path)
	assert.Equal(t, "", req.Query)
}
