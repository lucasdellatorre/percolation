package cmd

import (
	"log"
	"strconv"

	"github.com/lucasdellatorre/percolation/internal/simulation"
	"github.com/spf13/cobra"
)

func RunSimulation() *cobra.Command {
	return &cobra.Command{
		Use:     "run",
		Short:   "Run the simulation with a specific matrix size n",
		Example: "N=4, Runs the simulation with a 4x4 matrix",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			value := args[0]

			n, err := strconv.Atoi(value)

			if err != nil {
				log.Fatalln("The n value is invalid")
			}

			simulation.Run(n)
		},
	}

}
