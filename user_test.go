package wanikaniapi_test

import (
	"testing"

	"github.com/brandur/wanikaniapi"
	"github.com/brandur/wanikaniapi/wktesting"
	assert "github.com/stretchr/testify/require"
)

func TestUserGet(t *testing.T) {
	client := wktesting.TestClient()

	user, err := client.UserGet(&wanikaniapi.UserGetParams{})
	assert.NoError(t, err)
	t.Logf("user: %+v", user)
}
