package wanikaniapi_test

import (
	"net/http"
	"testing"

	"github.com/brandur/wanikaniapi"
	"github.com/brandur/wanikaniapi/wktesting"
	assert "github.com/stretchr/testify/require"
)

func TestSubjectList(t *testing.T) {
	client := wktesting.TestClient()
	client.RecordMode = true

	_, err := client.SubjectList(&wanikaniapi.SubjectListParams{
		Hidden: wanikaniapi.Bool(true),
		Levels: []int{1, 2, 3},
	})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, []byte(nil), req.Body)
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/v2/subjects", req.Path)
	assert.Equal(t, "hidden=true&levels=1,2,3", wktesting.MustQueryUnescape(req.Query))
}

func TestSubjectGet(t *testing.T) {
	client := wktesting.TestClient()
	client.RecordMode = true

	_, err := client.SubjectGet(&wanikaniapi.SubjectGetParams{ID: wanikaniapi.IDPtr(123)})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, []byte(nil), req.Body)
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/v2/subjects/123", req.Path)
	assert.Equal(t, "", req.Query)
}
