package wanikaniapi

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type ID int64

type ListParams struct {
	PageAfterID  *ID
	PageBeforeID *ID
}

func (p *ListParams) encodeToURLValues() url.Values {
	values := url.Values{}

	if p.PageAfterID != nil {
		values.Add("page_after_id", strconv.Itoa(int(*p.PageAfterID)))
	}
	if p.PageBeforeID != nil {
		values.Add("page_before_id", strconv.Itoa(int(*p.PageBeforeID)))
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

func Bool(b bool) *bool {
	return &b
}

func Int(i int) *int {
	return &i
}

func Time(t time.Time) *time.Time {
	return &t
}

func (c *Client) PageFully(onPage func(*ID) (*PageObject, error)) error {
	var nextPageAfterID *ID

	for {
		page, err := onPage(nextPageAfterID)
		if err != nil {
			return fmt.Errorf("error paginating fully: %w", err)
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
