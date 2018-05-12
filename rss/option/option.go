package option

import (
	"net/url"
	"strconv"
	"time"
)

//Option is option data for "getVulnOverviewList" API
type Option struct {
	severity                Severity //filter condition
	rangeDatePublic         *Period  //range mode for public date
	rangeDatePublished      *Period  //range mode for published date
	rangeDateFirstPublished *Period  //range mode for first published date
	startItem               int      //start point for entries
	keyword                 string   //filter condition
}

//OptFunc is self-referential function for functional options pattern
type OptFunc func(*Option)

// New returns a new Option instance
func New(opts ...OptFunc) *Option {
	o := &Option{severity: SeverityNone, rangeDatePublic: NewPeriod(), rangeDatePublished: NewPeriod(), rangeDateFirstPublished: NewPeriod(), startItem: 1}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

//WithStartItem returns function for setting startItem
func WithStartItem(start int) OptFunc {
	return func(o *Option) {
		o.SetStartItem(start)
	}
}

//WithSeverity returns function for setting severity
func WithSeverity(se Severity) OptFunc {
	return func(o *Option) {
		o.SetSeverity(se)
	}
}

//WithKeyword returns function for setting keyword
func WithKeyword(k string) OptFunc {
	return func(o *Option) {
		o.SetKeyword(k)
	}
}

//WithRangeDatePublicMode returns function for setting rangeDatePublic
func WithRangeDatePublicMode(mode RangeMode) OptFunc {
	return func(o *Option) {
		o.SetRangeDatePublicMode(mode)
	}
}

//WithRangeDatePublicPeriod returns function for setting rangeDatePublic
func WithRangeDatePublicPeriod(start, end time.Time) OptFunc {
	return func(o *Option) {
		o.SetRangeDatePublicPeriod(start, end)
	}
}

//WithRangeDatePublishedMode returns function for setting rangeDatePublished
func WithRangeDatePublishedMode(mode RangeMode) OptFunc {
	return func(o *Option) {
		o.SetRangeDatePublishedMode(mode)
	}
}

//WithRangeDatePublishedPeriod returns function for setting rangeDatePublished
func WithRangeDatePublishedPeriod(start, end time.Time) OptFunc {
	return func(o *Option) {
		o.SetRangeDatePublishedPeriod(start, end)
	}
}

//WithRangeDateFirstPublishedMode returns function for setting rangeDateFirstPublished
func WithRangeDateFirstPublishedMode(mode RangeMode) OptFunc {
	return func(o *Option) {
		o.SetRangeDateFirstPublishedMode(mode)
	}
}

//WithRangeDateFirstPublishedPeriod returns function for setting rangeDateFirstPublished
func WithRangeDateFirstPublishedPeriod(start, end time.Time) OptFunc {
	return func(o *Option) {
		o.SetRangeDateFirstPublishedPeriod(start, end)
	}
}

//SetStartItem sets startItem option
func (o *Option) SetStartItem(start int) *Option {
	if o != nil {
		o.startItem = start
	}
	return o
}

//SetSeverity sets severity option
func (o *Option) SetSeverity(se Severity) *Option {
	if o != nil {
		o.severity = se
	}
	return o
}

//SetKeyword sets keyword option
func (o *Option) SetKeyword(k string) *Option {
	if o != nil {
		o.keyword = k
	}
	return o
}

//SetRangeDatePublicMode sets rangeDatePublic option
func (o *Option) SetRangeDatePublicMode(mode RangeMode) *Option {
	if o != nil {
		o.rangeDatePublic.SetMode(mode)
	}
	return o
}

//SetRangeDatePublicPeriod sets datePublicStart and datePublicEnd option
func (o *Option) SetRangeDatePublicPeriod(start, end time.Time) *Option {
	if o != nil {
		o.rangeDatePublic.SetPeriod(start, end)
	}
	return o
}

//SetRangeDatePublishedMode sets rangeDatePublished option
func (o *Option) SetRangeDatePublishedMode(mode RangeMode) *Option {
	if o != nil {
		o.rangeDatePublished.SetMode(mode)
	}
	return o
}

//SetRangeDatePublishedPeriod sets datePublishedStart and datePublishedEnd option
func (o *Option) SetRangeDatePublishedPeriod(start, end time.Time) *Option {
	if o != nil {
		o.rangeDatePublished.SetPeriod(start, end)
	}
	return o
}

//SetRangeDateFirstPublishedMode sets rangeDateFirstPublished option
func (o *Option) SetRangeDateFirstPublishedMode(mode RangeMode) *Option {
	if o != nil {
		o.rangeDateFirstPublished.SetMode(mode)
	}
	return o
}

//SetRangeDateFirstPublishedPeriod sets dateFirstPublishedStart and dateFirstPublishedEnd option
func (o *Option) SetRangeDateFirstPublishedPeriod(start, end time.Time) *Option {
	if o != nil {
		o.rangeDateFirstPublished.SetPeriod(start, end)
	}
	return o
}

//AddQuery makes query of URL
func (o Option) AddQuery(query url.Values) {
	query.Add("startItem", strconv.Itoa(o.startItem))
	if o.severity != SeverityNone {
		query.Add("severity", o.severity.String())
	}
	if len(o.keyword) > 0 {
		query.Add("keyword", o.keyword)
	}
	if o.rangeDatePublic.IsPeriodSSet() {
		query.Add("rangeDatePublic", NoRange.String())
		query.Add("datePublicStartY", strconv.Itoa(o.rangeDatePublic.start.Year()))
		query.Add("datePublicStartM", strconv.Itoa(int(o.rangeDatePublic.start.Month())))
		query.Add("datePublicStartD", strconv.Itoa(o.rangeDatePublic.start.Day()))
		query.Add("datePublicEndY", strconv.Itoa(o.rangeDatePublic.end.Year()))
		query.Add("datePublicEndM", strconv.Itoa(int(o.rangeDatePublic.end.Month())))
		query.Add("datePublicEndD", strconv.Itoa(o.rangeDatePublic.end.Day()))
	} else {
		query.Add("rangeDatePublic", o.rangeDatePublic.mode.String())
	}
	if o.rangeDatePublished.IsPeriodSSet() {
		query.Add("rangeDatePublished", NoRange.String())
		query.Add("datePublishedStartY", strconv.Itoa(o.rangeDatePublished.start.Year()))
		query.Add("datePublishedStartM", strconv.Itoa(int(o.rangeDatePublished.start.Month())))
		query.Add("datePublishedStartD", strconv.Itoa(o.rangeDatePublished.start.Day()))
		query.Add("datePublishedEndY", strconv.Itoa(o.rangeDatePublished.end.Year()))
		query.Add("datePublishedEndM", strconv.Itoa(int(o.rangeDatePublished.end.Month())))
		query.Add("datePublishedEndD", strconv.Itoa(o.rangeDatePublished.end.Day()))
	} else {
		query.Add("rangeDatePublished", o.rangeDatePublished.mode.String())
	}
	if o.rangeDateFirstPublished.IsPeriodSSet() {
		query.Add("rangeDateFirstPublished", NoRange.String())
		query.Add("dateFirstPublishedStartY", strconv.Itoa(o.rangeDateFirstPublished.start.Year()))
		query.Add("dateFirstPublishedStartM", strconv.Itoa(int(o.rangeDateFirstPublished.start.Month())))
		query.Add("dateFirstPublishedStartD", strconv.Itoa(o.rangeDateFirstPublished.start.Day()))
		query.Add("dateFirstPublishedEndY", strconv.Itoa(o.rangeDateFirstPublished.end.Year()))
		query.Add("dateFirstPublishedEndM", strconv.Itoa(int(o.rangeDateFirstPublished.end.Month())))
		query.Add("dateFirstPublishedEndD", strconv.Itoa(o.rangeDateFirstPublished.end.Day()))
	} else {
		query.Add("rangeDateFirstPublished", o.rangeDateFirstPublished.mode.String())
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
