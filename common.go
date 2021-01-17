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

// IDPtr is a helper function that returns a pointer to the given value. This is
// useful for setting values in parameter structs.
func IDPtr(id ID) *ID {
	return &id
}

// Int is a helper function that returns a pointer to the given value. This is
// useful for setting values in parameter structs.
func Int(i int) *int {
	return &i
}

// Time is a helper function that returns a pointer to the given value. This is
// useful for setting values in parameter structs.
func Time(t time.Time) *time.Time {
	return &t
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

const (
	ObjectTypeAssignment             = ObjectType("assignment")
	ObjectTypeKanji                  = ObjectType("kanji")
	ObjectTypeLevelProgression       = ObjectType("level_progression")
	ObjectTypeRadical                = ObjectType("radical")
	ObjectTypeReset                  = ObjectType("reset")
	ObjectTypeReview                 = ObjectType("review")
	ObjectTypeReviewStatistic        = ObjectType("review_statistic")
	ObjectTypeSpacedRepetitionSystem = ObjectType("spaced_repetition_system")
	ObjectTypeUser                   = ObjectType("user")
	ObjectTypeVocabulary             = ObjectType("vocabulary")
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
	RecordedResponses [][]byte

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

func (c *Client) PageFully(onPage func(*ID) (*PageObject, error)) error {
	var nextPageAfterID *ID

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
			var nextPageAfterIDDisplay ID
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

		pageAfterID := ID(pageAfterIDInt)
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
	if c.RecordMode {
		c.RecordedRequests = append(c.RecordedRequests, &RecordedRequest{
			Body:   reqBytes,
			Header: req.Header,
			Method: method,
			Path:   path,
			Query:  query,
		})

		if len(c.RecordedResponses) > 0 {
			respBytes, c.RecordedResponses = c.RecordedResponses[0], c.RecordedResponses[1:]
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

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("Unexpected status from API: %v", resp.Status)
		}

		respBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil
		}
	}

	err = json.Unmarshal(respBytes, respData)
	if err != nil {
		return fmt.Errorf("error unmarshaling response: %v", err)
	}

	return nil
}

type ClientConfig struct {
	APIToken string
	Logger   LeveledLoggerInterface
}
type ID int64

type ListParams struct {
	PageAfterID  *ID
	PageBeforeID *ID
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
	DataUpdatedAt time.Time  `json:"data_updated_at"`
	ID            ID         `json:"id"`
	ObjectType    ObjectType `json:"object"`
	URL           string     `json:"url"`
}

type ObjectType string

type PageObject struct {
	Object

	TotalCount ID `json:"total_count"`

	Pages struct {
		NextURL     string `json:"next_url"`
		PerPage     int    `json:"per_page"`
		PreviousURL string `json:"previous_url"`
	} `json:"pages"`
}

type RecordedRequest struct {
	Body   []byte
	Header http.Header
	Method string
	Path   string
	Query  string
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

func joinIDs(ids []ID, separator string) string {
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

func joinObjectTypes(types []ObjectType, separator string) string {
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
