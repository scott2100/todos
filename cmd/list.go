package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Long:  `List all uncompleted todos.`,
	Run: func(cmd *cobra.Command, args []string) {
		//log.Println("list called")
		listAll, err := cmd.Flags().GetBool("all")
		check(err)
		file, err := os.Open("todos.csv")
		check(err)
		defer file.Close()

		r, err := csv.NewReader(file).ReadAll()
		check(err)

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', tabwriter.TabIndent)
		fmt.Fprintln(w, "ID\tDESCRIPTION\tCREATED\tCOMPLETE\t")
		for _, row := range r[1:] {
			isComplete := row[3]
			if isComplete == "false" && !listAll {
				fmt.Fprint(w, row[0], "\t", row[1], "\t", row[2], "\t", row[3], "\n")
			} else if listAll {
				fmt.Fprint(w, row[0], "\t", row[1], "\t", row[2], "\t", row[3], "\n")
			}
		}

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().BoolP("all", "a", false, "List all completed and uncompleted tasks.")
}
