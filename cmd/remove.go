/*
Copyright © 2024 github.com/pablos-juan/managerdetareas
*/
package cmd

import (
	"github.com/pablos-juan/managerdetareas/utilities"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var removeCmd = &cobra.Command{
	Use:   "delete",
	Short: "Borra una tarea de la lista",
	Long:  `Borra una tarea de la lista de tareas pendientes.`,
	Run: func(cmd *cobra.Command, args []string) {
		if all, _ := cmd.Flags().GetBool("all"); all {
			deleteAllTasks()
			return
		}

		deleteTaskIn(args)
		utilities.PrintTasks()
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().BoolP("all", "a", false, "Borra todas las tareas")
}

func deleteAllTasks() {
	err := utilities.Save([]utilities.Task{})
	if err != nil {
		panic(err)
	}
}

func deleteTaskIn(args []string) {
	i := utilities.GetIndex(args)
	if i < 0 {
		return
	}

	tasks, err := utilities.Load()
	if err != nil {
		panic(err)
	}

	if i < 0 || i > len(tasks) {
		utilities.ErrorMessage("Error. '%d' no está dentro de las tareas!", i)
		return
	}

	tasks = append(tasks[:i-1], tasks[i:]...)
	err = utilities.Save(tasks)
	if err != nil {
		panic(err)
	}
}
