/*
Copyright © 2024 github.com/pablos-juan/managerdetareas
*/
package cmd

import (
	"strconv"

	"github.com/pablos-juan/managerdetareas/utilities"
	"github.com/spf13/cobra"
)

func deleteTask(i int) {
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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Borra una tarea de la lista",
	Long:  `Borra una tarea de la lista de tareas pendientes.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			i, err := strconv.Atoi(args[0])
			if err != nil {
				utilities.ErrorMessage("Error. '%s' no es un número!", args[0])
				return
			}

			deleteTask(i)
			return
		}
		utilities.ErrorMessage("Error. Necesita un número entero como argumento!")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
