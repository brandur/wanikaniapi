package wanikaniapi_test

import (
	"net/http"
	"testing"

	"github.com/brandur/wanikaniapi"
	"github.com/brandur/wanikaniapi/wktesting"
	assert "github.com/stretchr/testify/require"
)

func TestStudyMaterialCreate(t *testing.T) {
	client := wktesting.LocalClient()

	_, err := client.StudyMaterialCreate(&wanikaniapi.StudyMaterialCreateParams{
		MeaningNote: wanikaniapi.String("hard"),
		SubjectID:   wanikaniapi.IDPtr(123),
	})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, `{"study_material":{"meaning_note":"hard","subject_id":123}}`, string(req.Body))
	assert.Equal(t, http.MethodPost, req.Method)
	assert.Equal(t, "/v2/study_materials", req.Path)
	assert.Equal(t, "", req.Query)
}

func TestStudyMaterialList(t *testing.T) {
	client := wktesting.LocalClient()

	_, err := client.StudyMaterialList(&wanikaniapi.StudyMaterialListParams{
		Hidden: wanikaniapi.Bool(true),
		IDs:    []wanikaniapi.ID{1, 2, 3},
	})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, []byte(nil), req.Body)
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/v2/study_materials", req.Path)
	assert.Equal(t, "hidden=true&ids=1,2,3", wktesting.MustQueryUnescape(req.Query))
}

func TestStudyMaterialGet(t *testing.T) {
	client := wktesting.LocalClient()

	_, err := client.StudyMaterialGet(&wanikaniapi.StudyMaterialGetParams{ID: wanikaniapi.IDPtr(123)})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, []byte(nil), req.Body)
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/v2/study_materials/123", req.Path)
	assert.Equal(t, "", req.Query)
}
