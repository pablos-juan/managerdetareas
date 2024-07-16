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
		utilities.PrintTasks()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
