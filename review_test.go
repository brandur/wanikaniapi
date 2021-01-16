package wanikaniapi_test

import (
	"net/http"
	"testing"

	"github.com/brandur/wanikaniapi"
	"github.com/brandur/wanikaniapi/wktesting"
	assert "github.com/stretchr/testify/require"
)

func TestReviewList(t *testing.T) {
	client := wktesting.TestClient()
	client.RecordMode = true

	_, err := client.ReviewList(&wanikaniapi.ReviewListParams{
		AssignmentIDs: []wanikaniapi.ID{1, 2, 3},
		SubjectIDs: []wanikaniapi.ID{4, 5, 6},
	})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, []byte(nil), req.Body)
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/v2/reviews", req.Path)
	assert.Equal(t, "assignment_ids=1,2,3&subject_ids=4,5,6", wktesting.QueryUnescape(req.Query))
}

func TestReviewGet(t *testing.T) {
	client := wktesting.TestClient()
	client.RecordMode = true

	_, err := client.ReviewGet(&wanikaniapi.ReviewGetParams{ID: wanikaniapi.IDPtr(123)})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, []byte(nil), req.Body)
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/v2/reviews/123", req.Path)
	assert.Equal(t, "", req.Query)
}
