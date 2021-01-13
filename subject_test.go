package wanikaniapi_test

import (
	"testing"

	"github.com/brandur/wanikaniapi"
	"github.com/brandur/wanikaniapi/wktesting"
	assert "github.com/stretchr/testify/require"
)

func TestListSubjects(t *testing.T) {
	client := wktesting.TestClient()
	subjects, err := client.SubjectList(&wanikaniapi.SubjectListParams{
		ListParams: &wanikaniapi.ListParams{},
	})
	assert.NoError(t, err)

	t.Logf("result = %+v", subjects.Data[0])
	t.Logf("result = %+v", subjects.Data[0].KanjiData)
	t.Logf("result = %+v", subjects.Data[0].RadicalData)
}
