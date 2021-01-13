package wanikaniapi_test

import (
	"testing"

	"github.com/brandur/wanikaniapi"
	"github.com/brandur/wanikaniapi/wktesting"
	assert "github.com/stretchr/testify/require"
)

func TestAssignmentListAndGet(t *testing.T) {
	client := wktesting.TestClient()
	var sampleID wanikaniapi.ID

	{
		assignments, err := client.AssignmentList(&wanikaniapi.AssignmentListParams{})
		assert.NoError(t, err)
		assert.Greater(t, len(assignments.Data), 0)

		sampleID = assignments.Data[0].ID
	}

	{
		assignment, err := client.AssignmentGet(&wanikaniapi.AssignmentGetParams{ID: &sampleID})
		assert.NoError(t, err)
		assert.NotNil(t, assignment)
	}
}
