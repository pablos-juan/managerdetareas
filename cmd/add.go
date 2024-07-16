/*
Copyright © 2024 github.com/pablos-juan/managerdetareas
*/
package cmd

import (
	"strings"

	"github.com/pablos-juan/managerdetareas/utilities"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "todo",
	Short: "Añade una tarea a la lista",
	Long:  `Añade una tarea nueva a la lista de tareas pendientes.`,
	Run: func(cmd *cobra.Command, args []string) {
		nombre := strings.Join(args, " ")
		addTask(nombre)
		utilities.PrintTasks()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

// addTask es una función que agrega una nueva tarea a la lista de tareas.
// Recibe el nombre de la tarea como parámetro y crea una nueva instancia de la estructura Task con el nombre proporcionado.
// Luego carga la lista de tareas existente, agrega la nueva tarea a la lista y guarda la lista actualizada.
func addTask(name string) {
	nt := utilities.Task{
		Name: name,
		Done: false,
	}

	tasks, err := utilities.Load()
	if err != nil {
		panic(err)
	}
	tasks = append(tasks, nt)

	err = utilities.Save(tasks)
	if err != nil {
		panic(err)
	}
}
