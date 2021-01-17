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

func TestUserUpdate(t *testing.T) {
	client := wktesting.LocalClient()

	_, err := client.UserUpdate(&wanikaniapi.UserUpdateParams{
		Preferences: &wanikaniapi.UserUpdatePreferencesParams{
			LessonsAutoplayAudio:       wanikaniapi.Bool(true),
			ReviewsDisplaySRSIndicator: wanikaniapi.Bool(true),
		},
	})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, `{"study_material":{"preferences":{"lessons_autoplay_audio":true,"reviews_display_srs_indicator":true}}}`, string(req.Body))
	assert.Equal(t, http.MethodPut, req.Method)
	assert.Equal(t, "/v2/user", req.Path)
	assert.Equal(t, "", req.Query)
}
