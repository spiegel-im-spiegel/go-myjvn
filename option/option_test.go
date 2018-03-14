package option

import (
	"net/url"
	"testing"
	"time"
)

func TestAddQuery(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	testCases := []struct {
		start     time.Time
		end       time.Time
		mode      RangeMode
		se        Severity
		startItem int
		query     string
	}{
		{start: time.Date(2018, 2, 3, 0, 0, 0, 0, jst), end: time.Date(2019, 12, 31, 0, 0, 0, 0, jst), startItem: 1, mode: NoRange, se: SeverityNone, query: "dateFirstPublishedEndD=31&dateFirstPublishedEndM=12&dateFirstPublishedEndY=2019&dateFirstPublishedStartD=3&dateFirstPublishedStartM=2&dateFirstPublishedStartY=2018&datePublicEndD=31&datePublicEndM=12&datePublicEndY=2019&datePublicStartD=3&datePublicStartM=2&datePublicStartY=2018&datePublishedEndD=31&datePublishedEndM=12&datePublishedEndY=2019&datePublishedStartD=3&datePublishedStartM=2&datePublishedStartY=2018&rangeDateFirstPublished=n&rangeDatePublic=n&rangeDatePublished=n&startItem=1"},
		{start: time.Date(2019, 12, 31, 0, 0, 0, 0, jst), end: time.Date(2018, 2, 3, 0, 0, 0, 0, jst), startItem: 1, mode: RangeWeek, se: SeverityNone, query: "dateFirstPublishedEndD=31&dateFirstPublishedEndM=12&dateFirstPublishedEndY=2019&dateFirstPublishedStartD=3&dateFirstPublishedStartM=2&dateFirstPublishedStartY=2018&datePublicEndD=31&datePublicEndM=12&datePublicEndY=2019&datePublicStartD=3&datePublicStartM=2&datePublicStartY=2018&datePublishedEndD=31&datePublishedEndM=12&datePublishedEndY=2019&datePublishedStartD=3&datePublishedStartM=2&datePublishedStartY=2018&rangeDateFirstPublished=n&rangeDatePublic=n&rangeDatePublished=n&startItem=1"},
	}

	for _, tc := range testCases {
		opt := New(
			WithStartItem(tc.startItem),
			WithRangeDate(tc.start, tc.end),
			WithRangeMode(tc.mode),
			WithSeverity(tc.se),
		)
		query := url.Values{}
		opt.AddQuery(query)
		str := query.Encode()
		if str != tc.query {
			t.Errorf("AddQuery() = \"%v\", want \"%v\".", str, tc.query)
		}
	}
}

func TestAddQuery2(t *testing.T) {
	testCases := []struct {
		mode      RangeMode
		startItem int
		se        Severity
		query     string
	}{
		{mode: RangeWeek, startItem: 1, se: SeverityHigh, query: "rangeDateFirstPublished=w&rangeDatePublic=w&rangeDatePublished=w&severity=h&startItem=1"},
		{mode: RangeMonth, startItem: 1, se: SeverityHigh, query: "rangeDateFirstPublished=m&rangeDatePublic=m&rangeDatePublished=m&severity=h&startItem=1"},
		{mode: NoRange, startItem: 1, se: SeverityHigh, query: "rangeDateFirstPublished=w&rangeDatePublic=w&rangeDatePublished=w&severity=h&startItem=1"},
	}

	for _, tc := range testCases {
		opt := New(
			WithStartItem(tc.startItem),
			WithRangeMode(tc.mode),
			WithSeverity(tc.se),
		)
		query := url.Values{}
		opt.AddQuery(query)
		str := query.Encode()
		if str != tc.query {
			t.Errorf("AddQuery() = \"%v\", want \"%v\".", str, tc.query)
		}
	}
}
