/*
 * Copyright @2024
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gosplash

import "github.com/Heyra-Joker/gosplash/types"

type Stats struct {
	Client *Client
}

// TotalsReq Get a list of counts for all of Unsplash request
type TotalsReq struct {
}

// TotalsReply Get a list of counts for all of Unsplash reply
type TotalsReply struct {
	types.Stats
}

func (TotalsReq) API() string {
	return "https://api.unsplash.com/stats/total"
}

func (TotalsReq) Params() map[string]string {
	return nil
}

// Totals Get a list of counts for all of Unsplash.
// document: https://unsplash.com/documentation#totals
func (s *Stats) Totals() (reply *TotalsReply, response *OmitResponse, err error) {
	response, err = s.Client.request("GET", TotalsReq{}, &reply, nil)
	return
}

// MonthReq Get the overall Unsplash stats for the past 30 days request
type MonthReq struct {
}

// MonthReqReply Get the overall Unsplash stats for the past 30 days reply
type MonthReqReply struct {
	types.Stats
}

func (MonthReq) API() string {
	return "https://api.unsplash.com/stats/month"
}

func (MonthReq) Params() map[string]string {
	return nil
}

// Month Get the overall Unsplash stats for the past 30 days.
// document: https://unsplash.com/documentation#month
func (s *Stats) Month() (reply *TotalsReply, response *OmitResponse, err error) {
	response, err = s.Client.request("GET", TotalsReq{}, &reply, nil)
	return
}
