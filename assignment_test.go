package wanikaniapi_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/brandur/wanikaniapi"
	"github.com/brandur/wanikaniapi/wktesting"
	assert "github.com/stretchr/testify/require"
)

func TestAssignmentList(t *testing.T) {
	client := wktesting.LocalClient()

	_, err := client.AssignmentList(&wanikaniapi.AssignmentListParams{
		Levels:  []int{1, 2, 3},
		Started: wanikaniapi.Bool(true),
	})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, "", string(req.Body))
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/v2/assignments", req.Path)
	assert.Equal(t, "levels=1,2,3&started=true", wktesting.MustQueryUnescape(req.Query))
}

func TestAssignmentGet(t *testing.T) {
	client := wktesting.LocalClient()

	_, err := client.AssignmentGet(&wanikaniapi.AssignmentGetParams{ID: wanikaniapi.ID(123)})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, "", string(req.Body))
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/v2/assignments/123", req.Path)
	assert.Equal(t, "", req.Query)
}

func TestAssignmentStart(t *testing.T) {
	client := wktesting.LocalClient()

	startedAt := time.Now()
	_, err := client.AssignmentStart(&wanikaniapi.AssignmentStartParams{
		ID:        wanikaniapi.ID(123),
		StartedAt: wanikaniapi.Time(startedAt),
	})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t,
		fmt.Sprintf(`{"started_at":"%v"}`, startedAt.Format(time.RFC3339)),
		string(req.Body),
	)
	assert.Equal(t, http.MethodPost, req.Method)
	assert.Equal(t, "/v2/assignments/123/start", req.Path)
	assert.Equal(t, "", req.Query)
}
