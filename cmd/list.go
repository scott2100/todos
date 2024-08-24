package cmd

import (
	"fmt"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
	"time"
	"todolist/todo"
	"todolist/utils/error"
	"todolist/utils/file"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Long:  `List all uncompleted todos.`,
	Run:   listTodos,
}

func listTodos(cmd *cobra.Command, args []string) {
	listAll, err := cmd.Flags().GetBool("all")
	error.HandleError(err)
	todosList := file.ReadFile()

	w := tabwriter.NewWriter(os.Stdout, 30, 30, 0, ' ', tabwriter.TabIndent)
	defer w.Flush()

	if len(todosList) > 0 {
		printHeader(w)
	}

	printRows(w, todosList, listAll)
}

func printRows(w *tabwriter.Writer, todosList []todo.Todo, listAll bool) {
	for _, t := range todosList {
		createdDateTimeString := timediff.TimeDiff(t.Created)
		completedDateTimeString := t.Completed.Format(time.RFC822)
		if t.Completed.IsZero() {
			_, err := fmt.Fprint(w, t.ID, "\t", t.Description, "\t", createdDateTimeString, "\t", "Not Complete", "\n")
			error.HandleError(err)
		} else if listAll == true {
			_, err := fmt.Fprint(w, t.ID, "\t", t.Description, "\t", createdDateTimeString, "\t", completedDateTimeString, "\n")
			error.HandleError(err)
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
