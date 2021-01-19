# wanikaniapi [![Build Status](https://github.com/brandur/wanikaniapi/workflows/wanikaniapi%20CI/badge.svg)](https://github.com/brandur/wanikaniapi/actions) [![Go Reference](https://pkg.go.dev/badge/github.com/brandur/wanikaniapi.svg)](https://pkg.go.dev/github.com/brandur/wanikaniapi)

A Go client for [WaniKani's API](https://docs.api.wanikani.com/).

**WARNING:** This project is under active development. Most API calls work fine, but it's incomplete in some ways like error handling and documentation.

## Usage

See the [full API reference on Go.dev](https://pkg.go.dev/github.com/brandur/wanikaniapi).

Contents:

* [Client initialization](#client-initialization)
* [Making API requests](#making-api-requests)
* [Pagination](#pagination)
* [Logging](#logging)
* [Handling errors](#handling-errors)
* [Conditional requests](#conditional-requests)
* [Automatic retries](#automatic-retries)

### Client initialization

All API requests are made through [`wanikaniapi.Client`](https://pkg.go.dev/github.com/brandur/wanikaniapi#Client). Make sure to include an API token:

``` go
package main

import (
    "os"

    "github.com/brandur/wanikaniapi"
)

func main() {
    client := wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
        APIToken: os.Getenv("WANI_KANI_API_TOKEN"),
    })

    ...
}
```

### Making API requests

Use an initialized client to make API requests:

``` go
package main

import (
    "os"

    "github.com/brandur/wanikaniapi"
)

func main() {
    client := wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
        APIToken: os.Getenv("WANI_KANI_API_TOKEN"),
    })

    voiceActors, err := client.VoiceActorList(&wanikaniapi.VoiceActorListParams{})
    if err != nil {
        panic(err)
    }

    ...
}
```

Function naming follows the pattern of `<API resource><Action>` like `AssignmentList`. Most resources support `*Get` and `*List`, and some support mutating operations like `*Create` or `*Start`.

### Setting API parameters

Go makes no distinction between a value that was left unset versus one set to an empty value (e.g. `""` for a string), so API parameters use pointers so it can be determined which values were meant to be sent and which ones weren't.

The package provides a set of helper functions to make setting pointers easy:

``` go
package main

import (
    "os"

    "github.com/brandur/wanikaniapi"
)

func main() {
    client := wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
        APIToken: os.Getenv("WANI_KANI_API_TOKEN"),
    })

    voiceActors, err := client.VoiceActorList(&wanikaniapi.VoiceActorListParams{
        IDs:          []wanikaniapi.WKID{1, 2, 3},
        UpdatedAfter: wanikaniapi.Time(time.Now()),
    })
    if err != nil {
        panic(err)
    }

    ...
}
```

The following helpers are available:

* [`Bool`](https://pkg.go.dev/github.com/brandur/wanikaniapi#Bool)
* [`ID`](https://pkg.go.dev/github.com/brandur/wanikaniapi#ID)
* [`Int`](https://pkg.go.dev/github.com/brandur/wanikaniapi#Int)
* [`String`](https://pkg.go.dev/github.com/brandur/wanikaniapi#String)
* [`Time`](https://pkg.go.dev/github.com/brandur/wanikaniapi#Time)

No helpers are needed for setting slices like `IDs` because slices are `nil` by default.

### Nil versus non-nil on API response structs

Values in API responses may be a pointer or non-pointer based on whether they're defined as nullable or not nullable by the WaniKani API:

``` go
type LevelProgressionData struct {
    AbandonedAt *time.Time `json:"abandoned_at"`
    CreatedAt   time.Time  `json:"created_at"`

    ...
```

`CreatedAt` always has a value and is therefore `time.Time`. `AbandonedAt` may be set or unset, and is therefore `*time.Time` instead.

### Pagination

List endpoints return list objects which contain only a single page worth of data, although they do have a pointer to where the next page's worth can be fetched:

``` go
package main

import (
    "fmt"
    "os"

    "github.com/brandur/wanikaniapi"
)

func main() {
    client := wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
        APIToken: os.Getenv("WANI_KANI_API_TOKEN"),
    })

    subjects, err := client.SubjectList(&wanikaniapi.SubjectListParams{})
    if err != nil {
        panic(err)
    }

    fmt.Printf("next page URL: %+v\n", subjects.Pages.NextURL)
}
```

Use the [`PageFully`](https://pkg.go.dev/github.com/brandur/wanikaniapi#Client.PageFully) helper to fully paginate an endpoint:

``` go
package main

import (
    "fmt"
    "os"

    "github.com/brandur/wanikaniapi"
)

func main() {
    client := wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
        APIToken: os.Getenv("WANI_KANI_API_TOKEN"),
    })

    var subjects []*wanikaniapi.Subject
    err := client.PageFully(func(id *wanikaniapi.WKID) (*wanikaniapi.PageObject, error) {
        page, err := client.SubjectList(&wanikaniapi.SubjectListParams{
            ListParams: wanikaniapi.ListParams{
                PageAfterID: id,
            },
        })
        if err != nil {
            return nil, err
        }

        subjects = append(subjects, page.Data...)
        return &page.PageObject, nil
    })
    if err != nil {
        panic(err)
    }

    fmt.Printf("num subjects: %v\n", len(subjects))
}
```

But remember to cache aggressively to minimize load on WaniKani. See [conditional requests](#conditional-requests) below.

### Logging

Configure a logger by passing a `Logger` parameter while initializing a client:

``` go
package main

import (
    "github.com/brandur/wanikaniapi"
)

func main() {
    client := wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
        Logger: &wanikaniapi.LeveledLogger{Level: wanikaniapi.LevelDebug},
    })

    ...
}
```

`Logger` expects a [`LeveledLoggerInterface`](https://pkg.go.dev/github.com/brandur/wanikaniapi#LeveledLoggerInterface):

``` go
type LeveledLoggerInterface interface {
    Debugf(format string, v ...interface{})
    Errorf(format string, v ...interface{})
    Infof(format string, v ...interface{})
    Warnf(format string, v ...interface{})
}
```

The package includes a basic logger called [`LeveledLogger`](https://pkg.go.dev/github.com/brandur/wanikaniapi#LeveledLogger) that implements it.

Some popular loggers like [Logrus](https://github.com/sirupsen/logrus/) and Zap's [SugaredLogger](https://godoc.org/go.uber.org/zap#SugaredLogger) also support this interface out-of-the-box so it's possible to set `DefaultLeveledLogger` to a `*logrus.Logger` or `*zap.SugaredLogger` directly. For others it may be necessary to write a shim layer to support them.

### Handling errors

API errors are returned as the special error struct [`*APIError`](https://pkg.go.dev/github.com/brandur/wanikaniapi#APIError):

``` go
package main

import (
    "github.com/brandur/wanikaniapi"
)

func main() {
    client := wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
        Logger: &wanikaniapi.LeveledLogger{Level: wanikaniapi.LevelDebug},
    })

    _, err := client.SubjectList(&wanikaniapi.SubjectListParams{})
    if err != nil {
        if apiErr, ok := err.(*wanikaniapi.APIError); ok {
            fmt.Printf("WaniKani API error; status: %v, message: %s\n",
                apiErr.StatusCode, apiErr.Message)
        } else {
            fmt.Printf("other error: %+v\n", err)
        }
    }

    ...
}
```

API calls may still return non-`APIError` errors for non-API problems (e.g. network error, TLS error, unmarshaling error, etc.).

### Conditional requests

TODO

### Automatic retries

The client can be configured to automatically retry errors that are known to be safe to retry:

``` go
package main

import (
    "github.com/brandur/wanikaniapi"
)

func main() {
    client := wanikaniapi.NewClient(&wanikaniapi.ClientConfig{
        MaxRetries: 2,
    })

    ...
}
```

## Development

### Run tests

Run the test suite:

``` sh
go test .
```

Tests generally compare recorded requests so that they don't have to make live API calls, but there are a few tests for the trickier cases which will only run when an API token is set:

``` sh
export WANI_KANI_API_TOKEN=
go test .
```
