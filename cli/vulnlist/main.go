package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spiegel-im-spiegel/go-myjvn"
	"github.com/spiegel-im-spiegel/go-myjvn/option"
	"github.com/spiegel-im-spiegel/go-myjvn/rss"
	"github.com/spiegel-im-spiegel/go-myjvn/status"
	"github.com/spiegel-im-spiegel/go-myjvn/vuldef"
)

func run(start, end time.Time) {
	api := myjvn.New()

	startItem := 1
	maxItem := 0
	vulnInfo := (*vuldef.VULDEF)(nil)
	for {
		opt := option.New(
			option.WithRangeDatePublicMode(option.NoRange),
			option.WithRangeDatePublishedPeriod(start, end),
			option.WithRangeDateFirstPublishedMode(option.NoRange),
			option.WithStartItem(startItem),
			//option.WithSeverity(option.SeverityHigh),
		)
		var rssData *rss.JVNRSS
		if startItem == 1 {
			xml, err := api.VulnOverviewListXML(opt)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			stat, err := status.Unmarshal(xml)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			if err := stat.GetError(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			maxItem = stat.Status.TotalRes
			if maxItem == 0 {
				fmt.Fprintln(os.Stderr, "no data")
				return
			}
			rssData, err = rss.Unmarshal(xml)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		} else {
			var err error
			rssData, err = api.VulnOverviewList(opt)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		}
		index := 0
		size := len(rssData.Items)
		for {
			ids := []string{}
			for i := 0; i < vuldef.MaxItems; i++ {
				if (index + i) >= size {
					break
				}
				ids = append(ids, rssData.Items[index+i].Identifier)
			}
			vuln, err := api.VulnDetailInfo(ids)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			if vulnInfo != nil {
				vulnInfo.Append(vuln)
			} else {
				vulnInfo = vuln
			}
			index += len(ids)
			if index >= size {
				break
			}
		}
		startItem += index
		if startItem > maxItem {
			break
		}
	}
	if vulnInfo != nil {
		json, err := vulnInfo.JSON("")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Println(string(json))
	} else {
		fmt.Fprintln(os.Stderr, "no data")
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}
	start, err := time.Parse("2006-01-02", os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	end, err := time.Parse("2006-01-02", os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	run(start, end)
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
