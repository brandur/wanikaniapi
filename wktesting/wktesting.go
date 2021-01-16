package wktesting

import (
	"encoding/json"
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

// MustMarshalJSON is the same as url.QueryUnescape except it panics on error.
func MustMarshalJSON(v interface{}) []byte {
	marshaled, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return marshaled
}

// MustQueryUnescape is the same as url.QueryUnescape except it panics on
// error.
func MustQueryUnescape(s string) string {
	unescaped, err := url.QueryUnescape(s)
	if err != nil {
		panic(err)
	}
	return unescaped
}

// LocalClient returns a WaniKani API client with RecordMode on. This means
// that instead of making live API calls, it saves them to RecoredRequests.
// This is suitable for most tests and allows us to avoid mutating our
// associated WaniKani account or exhausting its rate limit.
func LocalClient() *wanikaniapi.Client {
	client := wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
		APIToken: WaniKaniAPIToken,
		Logger:   &wanikaniapi.LeveledLogger{Level: wanikaniapi.LevelDebug},
	})
	client.RecordMode = true
	return client
}
