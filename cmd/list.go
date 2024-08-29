package cmd

import (
	"fmt"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
	"time"
	"todolist/utils/database"
	"todolist/utils/error"
	"todolist/utils/file"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Long:  `List all uncompleted todos.`,
	Run:   listTodos,
	Args:  cobra.NoArgs,
}

func listTodos(cmd *cobra.Command, args []string) {
	listAll, err := cmd.Flags().GetBool("all")
	error.HandleError(err)

	db := database.OpenDBConnection()
	defer db.Close()

	selectSql := `SELECT id, description, created, completed FROM todos`
	row, err := db.Query(selectSql)
	error.HandleError(err)
	defer row.Close()

	countSql := `SELECT COUNT(*) FROM todos`
	rowCount, err := db.Query(countSql)
	error.HandleError(err)
	defer rowCount.Close()

	w := tabwriter.NewWriter(os.Stdout, 20, 20, 0, ' ', tabwriter.TabIndent)
	defer w.Flush()

	if rowCount.Next() {
		printHeader(w)
	}

	for row.Next() {
		var id int
		var description string
		var created time.Time
		var completed time.Time
		row.Scan(&id, &description, &created, &completed)
		createdDateTimeString := timediff.TimeDiff(created)

		completedDateTimeString := completed.Format(time.RFC822)
		if completed.IsZero() {
			completedDateTimeString = "Not Complete"
			fmt.Fprint(w, id, "\t", description, "\t", createdDateTimeString, "\t", completedDateTimeString, "\n")
		} else if listAll == true {
			fmt.Fprint(w, id, "\t", description, "\t", createdDateTimeString, "\t", completedDateTimeString, "\n")
		}
	}
}

func printHeader(w *tabwriter.Writer) {
	_, err := fmt.Fprintln(w, file.Header)
	error.HandleError(err)
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().BoolP("all", "a", false, "List all completed and uncompleted tasks.")
}
