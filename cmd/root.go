/*
Copyright © 2024 github.com/pablos-juan/managerdetareas
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mann",
	Short: "Gestor de tareas",
	Long:  `A diferencia de un manager real, esta app no hace nada, pero es un buen ejemplo de como hacer un CLI con Cobra`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
