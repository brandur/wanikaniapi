package wktesting

import (
	"encoding/json"
	"net/url"
	"os"

	"github.com/brandur/wanikaniapi"
)

//////////////////////////////////////////////////////////////////////////////
//
//
//
// Exported functions
//
//
//
//////////////////////////////////////////////////////////////////////////////

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

// LiveClient checks to see if WANI_KANI_API_TOKEN is set in the environment.
// If it is, it returns a WaniClient API client suitable for making live API
// calls. If it isn't, it returns nil. It's up to the calling test to return in
// the latter case to avoid failure.
func LiveClient() *wanikaniapi.Client {
	apiToken := os.Getenv("WANI_KANI_API_TOKEN")
	if apiToken == "" {
		logger.Infof("no WANI_KANI_API_TOKEN in env; skipping test")
		return nil
	}

	client := wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
		APIToken: apiToken,
		Logger:   logger,
	})
	return client
}

// LocalClient returns a WaniKani API client with RecordMode on. This means
// that instead of making live API calls, it saves them to RecoredRequests.
// This is suitable for most tests and allows us to avoid mutating our
// associated WaniKani account or exhausting its rate limit.
func LocalClient() *wanikaniapi.Client {
	client := wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
		Logger: logger,
	})
	client.RecordMode = true
	return client
}

//////////////////////////////////////////////////////////////////////////////
//
//
//
// Internal
//
//
//
//////////////////////////////////////////////////////////////////////////////

var logger = &wanikaniapi.LeveledLogger{Level: wanikaniapi.LevelDebug}
