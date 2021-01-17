package wanikaniapi_test

import (
	"net/http"
	"testing"

	"github.com/brandur/wanikaniapi"
	"github.com/brandur/wanikaniapi/wktesting"
	assert "github.com/stretchr/testify/require"
)

func TestSpacedRepetitionSystemList(t *testing.T) {
	client := wktesting.LocalClient()

	_, err := client.SpacedRepetitionSystemList(&wanikaniapi.SpacedRepetitionSystemListParams{
		IDs: []wanikaniapi.ID{1, 2, 3},
	})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, "", string(req.Body))
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/v2/spaced_repetition_systems", req.Path)
	assert.Equal(t, "ids=1,2,3", wktesting.MustQueryUnescape(req.Query))
}

func TestSpacedRepetitionSystemGet(t *testing.T) {
	client := wktesting.LocalClient()

	_, err := client.SpacedRepetitionSystemGet(&wanikaniapi.SpacedRepetitionSystemGetParams{
		ID: wanikaniapi.IDPtr(123),
	})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, "", string(req.Body))
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/v2/spaced_repetition_systems/123", req.Path)
	assert.Equal(t, "", req.Query)
}
