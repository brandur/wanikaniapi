package wanikaniapi_test

import (
	"net/http"
	"testing"

	"github.com/brandur/wanikaniapi"
	"github.com/brandur/wanikaniapi/wktesting"
	assert "github.com/stretchr/testify/require"
)

func TestResetList(t *testing.T) {
	client := wktesting.LocalClient()

	_, err := client.ResetList(&wanikaniapi.ResetListParams{
		IDs: []wanikaniapi.WKID{1, 2, 3},
	})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, "", string(req.Body))
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/v2/resets", req.Path)
	assert.Equal(t, "ids=1,2,3", wktesting.MustQueryUnescape(req.Query))
}

func TestResetGet(t *testing.T) {
	client := wktesting.LocalClient()

	_, err := client.ResetGet(&wanikaniapi.ResetGetParams{ID: wanikaniapi.ID(123)})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, "", string(req.Body))
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/v2/resets/123", req.Path)
	assert.Equal(t, "", req.Query)
}
