package cmd

import (
	"log"
	"strconv"

	"github.com/gopxl/pixel/pixelgl"
	"github.com/lucasdellatorre/percolation/internal/guisimulation"
	"github.com/spf13/cobra"
)

func RunSimulationGui() *cobra.Command {
	return &cobra.Command{
		Use:     "gui",
		Short:   "Run the simulation GUI with a specific matrix size n",
		Example: "N=4, Runs the simulation with a 4x4 matrix",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			value := args[0]

			n, err := strconv.Atoi(value)

			if err != nil {
				log.Fatalln("The n value is invalid")
			}

			pixelgl.Run(func() {
				guisimulation.Run(n)
			})
		},
	}

}
