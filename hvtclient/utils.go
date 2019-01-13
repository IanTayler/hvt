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

package hvtclient

import (
	"fmt"
	"os"
)

func DefaultHvtClient() *HvtClient {
	accessToken := os.Getenv("HVT_ACCESS_TOKEN")
	accountID := os.Getenv("HVT_ACCOUNT_ID")
	user := os.Getenv("USER")
	if user == "" {
		user = "UNK"
	}
	host := os.Getenv("HOST")
	if host == "" {
		host = "UNK"
	}
	username := fmt.Sprintf("%s@%s", user, host)
	return NewHvtClient(accessToken, accountID, username)
}
