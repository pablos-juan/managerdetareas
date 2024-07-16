/*
Copyright © 2024 github.com/pablos-juan/managerdetareas
*/
package cmd

import (
	"github.com/pablos-juan/managerdetareas/utilities"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Marca una tarea como completada",
	Long:  `Marca una tarea como completada, moviendola al final de la lista de tareas pendientes.`,
	Run: func(cmd *cobra.Command, args []string) {
		completeTaskIn(args)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

// completeTaskIn completa una tarea en la lista de tareas.
// Recibe un slice de strings que contiene los argumentos.
// Si el índice de la tarea es válido, se marca como completada.
// Si ocurre algún error al cargar o guardar las tareas, se produce un pánico.
func completeTaskIn(args []string) {
	i := utilities.GetIndex(args)
	if i < 0 {
		return
	}

	tasks, err := utilities.Load()
	if err != nil {
		panic(err)
	}
	tasks[i-1].Done = true

	err = utilities.Save(tasks)
	if err != nil {
		panic(err)
	}
}
