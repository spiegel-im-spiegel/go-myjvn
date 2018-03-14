package option

import (
	"net/url"
	"strconv"
	"time"
)

//Option is option data for "getVulnOverviewList" API
type Option struct {
	severity  Severity  //filter condition
	rangeDate RangeMode //range mode
	dateStart time.Time //start date for search (should rangeDate == NoRange)
	dateEnd   time.Time //end date for search (should rangeDate == NoRange)
	startItem int       //start point for entries
}

//OptFunc is self-referential function for functional options pattern
type OptFunc func(*Option)

// New returns a new Option instance
func New(opts ...OptFunc) *Option {
	now := time.Now()
	o := &Option{severity: SeverityNone, rangeDate: RangeWeek, dateStart: now, dateEnd: now, startItem: 1}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

//WithStartItem returns function for setting startItem
func WithStartItem(start int) OptFunc {
	return func(o *Option) {
		o.startItem = start
	}
}

//WithRangeMode returns function for setting rangeDate
func WithRangeMode(mode RangeMode) OptFunc {
	return func(o *Option) {
		if o.rangeDate == NoRange {
			return
		}
		if mode == RangeMonth {
			o.rangeDate = mode
			return
		}
		o.rangeDate = RangeWeek
	}
}

//WithRangeDate returns function for setting rangeDate
func WithRangeDate(start, end time.Time) OptFunc {
	if start.After(end) {
		start, end = end, start
	}
	return func(o *Option) {
		o.dateStart = start
		o.dateEnd = end
		o.rangeDate = NoRange
	}
}

//WithSeverity returns function for setting severity
func WithSeverity(se Severity) OptFunc {
	return func(o *Option) {
		o.severity = se
	}
}

//AddQuery makes query of URL
func (o Option) AddQuery(query url.Values) {
	query.Add("startItem", strconv.Itoa(o.startItem))
	if o.severity != SeverityNone {
		query.Add("severity", o.severity.String())
	}
	if o.rangeDate == NoRange {
		query.Add("rangeDatePublic", NoRange.String())
		query.Add("datePublicStartY", strconv.Itoa(o.dateStart.Year()))
		query.Add("datePublicStartM", strconv.Itoa(int(o.dateStart.Month())))
		query.Add("datePublicStartD", strconv.Itoa(o.dateStart.Day()))
		query.Add("datePublicEndY", strconv.Itoa(o.dateEnd.Year()))
		query.Add("datePublicEndM", strconv.Itoa(int(o.dateEnd.Month())))
		query.Add("datePublicEndD", strconv.Itoa(o.dateEnd.Day()))
		query.Add("rangeDatePublished", NoRange.String())
		query.Add("datePublishedStartY", strconv.Itoa(o.dateStart.Year()))
		query.Add("datePublishedStartM", strconv.Itoa(int(o.dateStart.Month())))
		query.Add("datePublishedStartD", strconv.Itoa(o.dateStart.Day()))
		query.Add("datePublishedEndY", strconv.Itoa(o.dateEnd.Year()))
		query.Add("datePublishedEndM", strconv.Itoa(int(o.dateEnd.Month())))
		query.Add("datePublishedEndD", strconv.Itoa(o.dateEnd.Day()))
		query.Add("rangeDateFirstPublished", NoRange.String())
		query.Add("dateFirstPublishedStartY", strconv.Itoa(o.dateStart.Year()))
		query.Add("dateFirstPublishedStartM", strconv.Itoa(int(o.dateStart.Month())))
		query.Add("dateFirstPublishedStartD", strconv.Itoa(o.dateStart.Day()))
		query.Add("dateFirstPublishedEndY", strconv.Itoa(o.dateEnd.Year()))
		query.Add("dateFirstPublishedEndM", strconv.Itoa(int(o.dateEnd.Month())))
		query.Add("dateFirstPublishedEndD", strconv.Itoa(o.dateEnd.Day()))
	} else {
		query.Add("rangeDatePublic", o.rangeDate.String())
		query.Add("rangeDatePublished", o.rangeDate.String())
		query.Add("rangeDateFirstPublished", o.rangeDate.String())
	}
}

/* Copyright 2018 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
