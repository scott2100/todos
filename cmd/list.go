package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
	"time"
	"todolist/utils/error"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Long:  `List all uncompleted todos.`,
	Run: func(cmd *cobra.Command, args []string) {
		listAll, err := cmd.Flags().GetBool("all")
		error.HandleError(err)
		file, err := os.Open("todos.csv")
		error.HandleError(err)
		defer file.Close()

		r, err := csv.NewReader(file).ReadAll()
		error.HandleError(err)

		w := tabwriter.NewWriter(os.Stdout, 30, 30, 0, ' ', tabwriter.TabIndent)

		if len(r) > 1 {
			_, err := fmt.Fprintln(w, "ID\tTasks\tCreated\tCompleted")
			error.HandleError(err)
		}
		for _, row := range r[1:] {
			isComplete := row[3]
			createdDateTime, err := time.Parse(time.RFC3339, row[2])
			error.HandleError(err)
			createdDateTimeString := timediff.TimeDiff(createdDateTime)
			if isComplete == "false" && !listAll {
				_, err := fmt.Fprint(w, row[0], "\t", row[1], "\t", createdDateTimeString, "\t", row[3], "\n")
				error.HandleError(err)
			} else if listAll {
				_, err := fmt.Fprint(w, row[0], "\t", row[1], "\t", createdDateTimeString, "\t", row[3], "\n")
				error.HandleError(err)
			}
		}
		err = w.Flush()
		error.HandleError(err)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().BoolP("all", "a", false, "List all completed and uncompleted tasks.")
}
