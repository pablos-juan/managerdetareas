/*
Copyright © 2024 NAME HERE <pabloccsanchez@gmail.com>
*/
package cmd

import (
	"strings"

	"github.com/pablos-juan/managerdetareas/utilities"
	"github.com/spf13/cobra"
)

func addTask(name string) {
	nt := utilities.Task{
		Name: name,
		Done: false,
	}

	err := utilities.Save(nt)
	if err != nil {
		panic(err)
	}
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Añade una tarea a la lista",
	Long:  `Añade una tarea nueva a la lista de tareas pendientes.`,
	Run: func(cmd *cobra.Command, args []string) {
		nombre := strings.Join(args, " ")
		addTask(nombre)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
