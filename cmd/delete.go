/*
Copyright © 2024 github.com/pablos-juan/managerdetareas
*/
package cmd

import (
	"strconv"

	"github.com/pablos-juan/managerdetareas/utilities"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
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
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().BoolP("all", "a", false, "Borra todas las tareas")
}

func deleteAllTasks() {
	err := utilities.Save([]utilities.Task{})
	if err != nil {
		panic(err)
	}
}

func deleteTaskIn(args []string) {
	i := getIndex(args)
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

func getIndex(args []string) int {
	if len(args) == 1 {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			utilities.ErrorMessage("Error. '%s' no es un número!", args[0])
			return -1
		}

		return i
	}

	utilities.ErrorMessage("Error. Necesita un número entero como argumento!")
	return -1
}
