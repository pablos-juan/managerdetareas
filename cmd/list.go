/*
Copyright © 2024 github.com/pablos-juan/managerdetareas
*/
package cmd

import (
	"github.com/pablos-juan/managerdetareas/utilities"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "Muestra la lista de tareas pendientes",
	Long:  `Muestra la lista de tareas pendientes en forma de tabla.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := utilities.Load()
		if err != nil {
			panic(err)
		}

		if incomplete, _ := cmd.Flags().GetBool("incomplete"); incomplete {
			ft := utilities.Filter(tasks, func(t utilities.Task) bool {
				return !t.Done
			})
			utilities.PrintTable(ft)
			return
		}

		if done, _ := cmd.Flags().GetBool("done"); done {
			ft := utilities.Filter(tasks, func(t utilities.Task) bool {
				return t.Done
			})
			utilities.PrintTable(ft)
			return
		}

		utilities.PrintTasks()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("incomplete", "i", false, "Muestra solo las tareas que no se han completado")
	listCmd.Flags().BoolP("done", "d", false, "Muestra solo las tareas completadas")
}
