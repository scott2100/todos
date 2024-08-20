package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark todo as completed",
	Long:  `Mark todo as completed`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("complete called")

		file, err := os.OpenFile("todos.csv", os.O_WRONLY|os.O_APPEND, 0644)
		check(err)

		r, err := csv.NewReader(file).ReadAll()
		check(err)
		file.Close()

		file, _ = os.OpenFile("todos.csv", os.O_WRONLY|os.O_APPEND, 0644)
		w := csv.NewWriter(file)

		tw := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', tabwriter.TabIndent)
		fmt.Fprintln(tw, "ID\tDESCRIPTION\tCREATED\tCOMPLETE\t")
		for _, row := range r[1:] {
			id := row[0]
			if id == strings.Join(args, "") {
				w.Write([]string{row[0], row[1], row[2], strconv.FormatBool(true)})
			}
		}

		tw.Flush()
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
