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

package cmd

import (
	"fmt"

	"github.com/IanTayler/hvt/hvtclient"
	"github.com/spf13/cobra"
)

func showHandler(cmd *cobra.Command, args []string) {
	hvtClient := hvtclient.DefaultHvtClient()
	timeEntryList, _ := hvtClient.ListTimeEntries("2019-01-01", "2019-01-11")
	fmt.Printf("%+v\n", *timeEntryList)
}

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show time entries for your account.",
	Long: `Use this command to see which time entries have already been logged
	to projects in your account, including information about which projects
	and tasks those time entries are assigned to.`,
	Run: showHandler,
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
