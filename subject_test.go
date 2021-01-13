package wanikaniapi_test

import (
	"testing"

	"github.com/brandur/wanikaniapi"
	"github.com/brandur/wanikaniapi/wktesting"
	assert "github.com/stretchr/testify/require"
)

func TestSubjectListAndGet(t *testing.T) {
	client := wktesting.TestClient()
	var sampleID wanikaniapi.ID

	{
		subjects, err := client.SubjectList(&wanikaniapi.SubjectListParams{})
		assert.NoError(t, err)
		assert.Greater(t, len(subjects.Data), 0)

		sampleID = subjects.Data[0].ID
	}

	{
		subject, err := client.SubjectGet(&wanikaniapi.SubjectGetParams{ID: &sampleID})
		assert.NoError(t, err)
		assert.NotNil(t, subject)
	}
}
