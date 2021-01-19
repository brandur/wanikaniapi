package wanikaniapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
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

// Bool is a helper function that returns a pointer to the given value. This is
// useful for setting values in parameter structs.
func Bool(b bool) *bool {
	return &b
}

// ID is a helper function that returns a pointer to the given value. This is
// useful for setting values in parameter structs.
func ID(id WKID) *WKID {
	return &id
}

// Int is a helper function that returns a pointer to the given value. This is
// useful for setting values in parameter structs.
func Int(i int) *int {
	return &i
}

// String is a helper function that returns a pointer to the given value. This is
// useful for setting values in parameter structs.
func String(s string) *string {
	return &s
}

// Time is a helper function that returns a pointer to the given value. This is
// useful for setting values in parameter structs.
func Time(t time.Time) *WKTime {
	wkT := WKTime(t)
	return &wkT
}

//////////////////////////////////////////////////////////////////////////////
//
//
//
// Exported constants/types
//
//
//
//////////////////////////////////////////////////////////////////////////////

// Constants for the various object types that may be returned in the object
// field of WaniKani's API resources.
const (
	ObjectTypeAssignment             = WKObjectType("assignment")
	ObjectTypeCollection             = WKObjectType("collection")
	ObjectTypeKanji                  = WKObjectType("kanji")
	ObjectTypeLevelProgression       = WKObjectType("level_progression")
	ObjectTypeRadical                = WKObjectType("radical")
	ObjectTypeReport                 = WKObjectType("report")
	ObjectTypeReset                  = WKObjectType("reset")
	ObjectTypeReview                 = WKObjectType("review")
	ObjectTypeReviewStatistic        = WKObjectType("review_statistic")
	ObjectTypeSpacedRepetitionSystem = WKObjectType("spaced_repetition_system")
	ObjectTypeStudyMaterial          = WKObjectType("study_material")
	ObjectTypeUser                   = WKObjectType("user")
	ObjectTypeVocabulary             = WKObjectType("vocabulary")
	ObjectTypeVoiceActor             = WKObjectType("voice_actor")
)

const WaniKaniAPIURL = "https://api.wanikani.com"
const WaniKaniRevision = "20170710"

type Client struct {
	APIToken string
	Logger   LeveledLoggerInterface

	// RecordMode stubs out any actual HTTP calls, and instead starts storing
	// request data to RecordedRequests.
	RecordMode bool

	RecordedRequests  []*RecordedRequest
	RecordedResponses []*RecordedResponse

	baseURL    string
	httpClient *http.Client
}

func NewClient(config *ClientConfig) *Client {
	var logger LeveledLoggerInterface

	if config.Logger == nil {
		logger = &LeveledLogger{Level: LevelError}
	} else {
		logger = config.Logger
	}

	return &Client{
		APIToken: config.APIToken,
		Logger:   logger,

		baseURL:    WaniKaniAPIURL,
		httpClient: &http.Client{},
	}
}

func (c *Client) PageFully(onPage func(*WKID) (*PageObject, error)) error {
	var nextPageAfterID *WKID

	for {
		page, err := onPage(nextPageAfterID)
		if err != nil {
			return fmt.Errorf("error paginating fully: %w", err)
		}
		if page == nil {
			c.Logger.Debugf("Page function returned nil; breaking pagination")
			return err
		}

		{
			var nextPageAfterIDDisplay WKID
			if nextPageAfterID != nil {
				nextPageAfterIDDisplay = *nextPageAfterID
			}
			c.Logger.Debugf("Got page; id=%+v, per_page=%v, next_url=%v",
				nextPageAfterIDDisplay, page.Pages.PerPage, page.Pages.NextURL)
		}

		if page.Pages.NextURL == "" {
			break
		}

		u, err := url.Parse(page.Pages.NextURL)
		if err != nil {
			return fmt.Errorf("error parsing next page URL: %w", err)
		}

		queryValues, err := url.ParseQuery(u.RawQuery)
		if err != nil {
			return fmt.Errorf("error parsing next page query string: %w", err)
		}

		pageAfterIDStr := queryValues.Get("page_after_id")
		if pageAfterIDStr == "" {
			return fmt.Errorf("no `page_after_id` in next page query string")
		}

		pageAfterIDInt, err := strconv.Atoi(pageAfterIDStr)
		if err != nil {
			return fmt.Errorf("couldn't parse `page_after_id` in next page query string")
		}

		pageAfterID := WKID(pageAfterIDInt)
		nextPageAfterID = &pageAfterID
	}

	return nil
}

func (c *Client) request(method, path, query string, reqData interface{}, respData interface{}) error {
	if c.APIToken == "" && !c.RecordMode {
		return fmt.Errorf("wanikaniapi.Client.APIToken must be set to make a live API call")
	}

	url := c.baseURL + path
	if query != "" {
		url += "?" + query
	}

	c.Logger.Debugf("Requesting URL: %v (revision: %v)", url, WaniKaniRevision)

	var reqBytes []byte
	var reqReader io.Reader
	if reqData != nil {
		var err error
		reqBytes, err = json.Marshal(reqData)
		if err != nil {
			return err
		}
		reqReader = bytes.NewReader(reqBytes)
	}

	req, err := http.NewRequest(method, url, reqReader)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	req.Header.Set("Wanikani-Revision", WaniKaniRevision)
	if reqReader != nil {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	var respBytes []byte
	var statusCode int
	if c.RecordMode {
		c.RecordedRequests = append(c.RecordedRequests, &RecordedRequest{
			Body:   reqBytes,
			Header: req.Header,
			Method: method,
			Path:   path,
			Query:  query,
		})

		statusCode = http.StatusOK
		if len(c.RecordedResponses) > 0 {
			var resp *RecordedResponse
			resp, c.RecordedResponses = c.RecordedResponses[0], c.RecordedResponses[1:]

			respBytes = resp.Body
			statusCode = resp.StatusCode
		}
		if respBytes == nil {
			respBytes = []byte("{}")
		}
	} else {
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		statusCode = resp.StatusCode
		respBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil
		}
	}

	if statusCode != http.StatusOK {
		var apiErr APIError
		err := json.Unmarshal(respBytes, &apiErr)
		if err != nil {
			return fmt.Errorf("error unmarshaling error response: %v", err)
		}

		return apiErr
	}

	err = json.Unmarshal(respBytes, respData)
	if err != nil {
		return fmt.Errorf("error unmarshaling response: %v", err)
	}

	return nil
}

// APIError represents an HTTP status API error that came back from WaniKani's
// API. It may be caused by a variety of problems like a bad access token
// resulting in a 401 Unauthorized or making too many requests resulting in a
// 429 Too Many Requests.
//
// Inspect Code and Error for details, and see the API reference for more
// information:
//
// https://docs.api.wanikani.com/20170710/#errors
type APIError struct {
	// Code is the HTTP status code that came back with the API error.
	Code  int    `json:"code"`

	// Error is the error message that came back with the API error.
	//
	// This is called Message instead of Error so as not to conflict with the
	// Error function on Go's error interface.
	Message string `json:"error"`
}

// Error returns the error message that came back with the API error.
func (e APIError) Error() string {
	return e.Message
}

type ClientConfig struct {
	APIToken string
	Logger   LeveledLoggerInterface
}

type ListParams struct {
	PageAfterID  *WKID
	PageBeforeID *WKID
}

func (p *ListParams) encodeToURLValues() url.Values {
	values := url.Values{}

	if p != nil {
		if p.PageAfterID != nil {
			values.Add("page_after_id", strconv.FormatInt(int64(*p.PageAfterID), 10))
		}
		if p.PageBeforeID != nil {
			values.Add("page_before_id", strconv.FormatInt(int64(*p.PageBeforeID), 10))
		}
	}

	return values
}

type ListParamsInterface interface {
	EncodeToQuery() string
}

type Object struct {
	DataUpdatedAt time.Time    `json:"data_updated_at"`
	ID            WKID         `json:"id"`
	ObjectType    WKObjectType `json:"object"`
	URL           string       `json:"url"`
}

type PageObject struct {
	Object

	TotalCount int64 `json:"total_count"`

	Pages struct {
		NextURL     string `json:"next_url"`
		PerPage     int    `json:"per_page"`
		PreviousURL string `json:"previous_url"`
	} `json:"pages"`
}

// RecordedRequest is a request recorded when RecordMode is on.
type RecordedRequest struct {
	Body   []byte
	Header http.Header
	Method string
	Path   string
	Query  string
}

// RecordedResponse is a reponse injected when RecordMode is on.
type RecordedResponse struct {
	Body       []byte
	StatusCode int
}

// WKID represents a WaniKani API identifier.
type WKID int64

// WKObjectType represents a type of object in the WaniKani API.
type WKObjectType string

// WKTime is a type based on time.Time that lets us precisely control the JSON
// marshaling for use in API parameters to endpoints.
type WKTime time.Time

// Encode encodes the time to the RFC3339 format that WaniKani expects.
func (t WKTime) Encode() string {
	return time.Time(t).Format(time.RFC3339)
}

// MarshalJSON overrides JSON marshaling for WKTime so that it can be put in
// the RFC3339 format that WaniKani expects.
func (t WKTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Encode() + `"`), nil
}

//////////////////////////////////////////////////////////////////////////////
//
//
//
// Internal functions
//
//
//
//////////////////////////////////////////////////////////////////////////////

func joinIDs(ids []WKID, separator string) string {
	var s string

	for i, n := range ids {
		if i != 0 {
			s += ","
		}

		s += strconv.FormatInt(int64(n), 10)
	}

	return s
}

func joinInts(ints []int, separator string) string {
	var s string

	for i, n := range ints {
		if i != 0 {
			s += ","
		}

		s += strconv.Itoa(n)
	}

	return s
}

func joinObjectTypes(types []WKObjectType, separator string) string {
	var s string

	for i, typ := range types {
		if i != 0 {
			typ += ","
		}

		s += string(typ)
	}

	return s
}

func joinStrings(strs []string, separator string) string {
	var s string

	for i, str := range strs {
		if i != 0 {
			str += ","
		}

		s += str
	}

	return s
}
