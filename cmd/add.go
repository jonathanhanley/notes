package cmd

import (
	"fmt"
	"github.com/noculture/notes/models"
	"github.com/spf13/cobra"
	"log"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a note",
	Long: "Add a note to a notebook from the terminal. Use `notes add \"text\"` to jot in the default notebook or" +
		"`notes add NotebookName \"text\"` to add notes to other notebooks",
	Run: func(cmd *cobra.Command, args []string) {
		db := setupDatabase()
		var err error

		switch len(args) {
		case 0:
			fmt.Println(" :warning: You need to add some text")
		case 1:
			err = db.AddNote("Default", models.Note{Content: args[0]})
			fmt.Println(" :pencil2: Note added")
		default:
			err = db.AddNote(args[0], models.Note{Content: args[1]})
			fmt.Println(" :pencil2: Note added")
		}
		if err != nil {
			log.Panic()
		}

	},
}

func init() {
	root.AddCommand(addCommand)
}
