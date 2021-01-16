package wktesting

import (
	"net/url"
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

// QueryUnescape is the same as url.QueryUnescape except it panics on error.
func QueryUnescape(s string) string {
	unescaped, err := url.QueryUnescape(s)
	if err != nil {
		panic(err)
	}
	return unescaped
}

// TestClient returns a WaniKani API client suitable for use in tests.
func TestClient() *wanikaniapi.Client {
	return wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
		APIToken: WaniKaniAPIToken,
		Logger:   &wanikaniapi.LeveledLogger{Level: wanikaniapi.LevelDebug},
	})
}
