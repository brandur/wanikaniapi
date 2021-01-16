package wanikaniapi_test

import (
	"net/http"
	"testing"

	"github.com/brandur/wanikaniapi"
	"github.com/brandur/wanikaniapi/wktesting"
	assert "github.com/stretchr/testify/require"
)

func TestUserGet(t *testing.T) {
	client := wktesting.LocalClient()

	_, err := client.UserGet(&wanikaniapi.UserGetParams{})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, []byte(nil), req.Body)
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/v2/user", req.Path)
	assert.Equal(t, "", req.Query)
}
