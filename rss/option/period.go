package option

import "time"

//Period is setting info for range date
type Period struct {
	mode  RangeMode //range mode
	start time.Time //start date for search (should mode == NoRange)
	end   time.Time //end date for search (should mode == NoRange)
}

//NewPeriod returns nre Period instance
func NewPeriod() *Period {
	return &Period{mode: RangeWeek, start: time.Time{}, end: time.Time{}}
}

//SetMode sets modo
func (p *Period) SetMode(mode RangeMode) *Period {
	if p != nil {
		if mode != NoRange && p.IsPeriodSSet() {
			return p
		}
		switch mode {
		case RangeWeek, RangeMonth, NoRange:
			p.mode = mode
		default:
		}
		p.start = time.Time{}
		p.end = time.Time{}
	}
	return p
}

//SetPeriod sets modo to NoRange and sets period
func (p *Period) SetPeriod(start, end time.Time) *Period {
	if p != nil {
		p.mode = NoRange
		if start.After(end) {
			start, end = end, start
		}
		p.start = start
		p.end = end
	}
	return p
}

//IsPeriodSSet returns true if sets period
func (p *Period) IsPeriodSSet() bool {
	if p != nil {
		if p.mode == NoRange && !p.start.IsZero() && !p.end.IsZero() {
			return true
		}
	}
	return false
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
