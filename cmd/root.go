package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "percolation",
		Short:         "Monte carlo simulation for percolation",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	rootCmd.AddCommand(RunSimulation())
	rootCmd.AddCommand(RunSimulationGui())
	return rootCmd
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
