package values

import (
	"encoding/xml"
	"time"
)

//Time is wrapper class of time.Time
type Time struct {
	time.Time
}

//UnmarshalXML returns result of Unmarshal for xml.Unmarshal()
func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}
	tm, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return err
	}
	*t = Time{tm}
	return nil
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
