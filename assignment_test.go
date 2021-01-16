package wanikaniapi_test

import (
	"net/http"
	"testing"

	"github.com/brandur/wanikaniapi"
	"github.com/brandur/wanikaniapi/wktesting"
	assert "github.com/stretchr/testify/require"
)

func TestAssignmentList(t *testing.T) {
	client := wktesting.TestClient()
	client.RecordMode = true

	_, err := client.AssignmentList(&wanikaniapi.AssignmentListParams{
		Levels:  []int{1, 2, 3},
		Started: wanikaniapi.Bool(true),
	})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, []byte(nil), req.Body)
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/v2/assignments", req.Path)
	assert.Equal(t, "levels=1,2,3&started=true", wktesting.MustQueryUnescape(req.Query))
}

func TestAssignmentGet(t *testing.T) {
	client := wktesting.TestClient()
	client.RecordMode = true

	_, err := client.AssignmentGet(&wanikaniapi.AssignmentGetParams{ID: wanikaniapi.IDPtr(123)})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, []byte(nil), req.Body)
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/v2/assignments/123", req.Path)
	assert.Equal(t, "", req.Query)
}
