package wktesting

import (
	"os"

	"github.com/brandur/wanikaniapi"
)

// WaniKaniAPIToken is an API token to be used in testing.
var WaniKaniAPIToken string

func init() {
	WaniKaniAPIToken = os.Getenv("WANI_KANI_API_TOKEN")
	if WaniKaniAPIToken == "" {
		panic("tests need WANI_KANI_API_TOKEN in env")
	}
}

// TestClient returns a WaniKani API client suitable for use in tests.
func TestClient() *wanikaniapi.Client {
	return wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
		APIToken: WaniKaniAPIToken,
		Logger:   &wanikaniapi.LeveledLogger{Level: wanikaniapi.LevelDebug},
	})
}
