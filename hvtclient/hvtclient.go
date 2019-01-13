// Copyright © 2019 Ian Tayler <iangtayler@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

//Package hvtclient provides the library code for communication with Harvest.
package hvtclient

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//BaseAPIURL is the root URL for the Harvest API.
const BaseAPIURL = "https://api.harvestapp.com/v2"

//HvtInterface defines methods that a Harvest client should implement.
type HvtInterface interface {
}

//HvtClient is an implementation of a Harvest client.
type HvtClient struct {
	accessToken string
	accountID   string
	username    string
	httpClient  *http.Client
}

//NewHvtClient create a Harvest client that will authenticate requests.
func NewHvtClient(accessToken, accountID, username string) *HvtClient {
	httpClient := &http.Client{}
	return &HvtClient{
		accessToken: accessToken,
		accountID:   accountID,
		httpClient:  httpClient,
		username:    username,
	}
}

//NewAuthRequest create an *http.Request with authentication headers.
func (h *HvtClient) NewAuthRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, reqErr := http.NewRequest(method, BaseAPIURL+path, body)
	if reqErr != nil {
		return nil, reqErr
	}
	req.Header.Add("Authorization", "Bearer "+h.accessToken)
	req.Header.Add("Harvest-Account-Id", h.accountID)
	req.Header.Add("User-Agent", fmt.Sprintf(
		"hvt User: %s, Dev: Ian Tayler (iangtayler@gmail.com)",
		h.username,
	))
	return req, nil
}

//ListTimeEntries get time entries for the authenticated user in a period of time.
func (h *HvtClient) ListTimeEntries(from, to string) (*TimeEntryList, error) {
	req, _ := h.NewAuthRequest("GET", "/time_entries", nil)
	q := req.URL.Query()
	if from != "" {
		q.Add("from", from)
	}
	if to != "" {
		q.Add("to", to)
	}
	req.URL.RawQuery = q.Encode()
	resp, reqErr := h.httpClient.Do(req)
	if reqErr != nil {
		return nil, reqErr
	}
	defer resp.Body.Close()
	timeEntryList := &TimeEntryList{}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	unmarshalErr := json.Unmarshal(bodyBytes, &timeEntryList)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return timeEntryList, nil
}

//PostTimeEntry create a time entry by posting it to the Harvest API.
func (h *HvtClient) PostTimeEntry(projectID int64, taskID int64, spentDate string, hours string) error {
	urlValues := url.Values{
		"project_id": {strconv.FormatInt(projectID, 10)},
		"task_id":    {strconv.FormatInt(taskID, 10)},
		"spent_date": {spentDate},
		"hours":      {hours}}
	bodyReader := strings.NewReader(urlValues.Encode())
	req, _ := h.NewAuthRequest("POST", BaseAPIURL+"/time_entries", bodyReader)
	resp, clientErr := h.httpClient.Do(req)
	if clientErr != nil {
		return clientErr
	}
	defer resp.Body.Close()
	return nil
}

//Project represents a Harvest project.
type Project struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

//TimeEntry represents a Harvest time entry.
type TimeEntry struct {
	ID        int64   `json:"id"`
	SpentDate string  `json:"spent_date"`
	Hours     float64 `json:"hours"`
	Notes     string  `json:"notes"`
	Project   Project `json:"project"`
}

//TimeEntryList models the response for HvtClient.ListTimeEntries.
type TimeEntryList struct {
	TimeEntries []TimeEntry `json:"time_entries"`
}
