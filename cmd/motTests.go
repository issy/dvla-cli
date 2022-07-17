package cmd

import (
	"fmt"
	"github.com/issy/dvla-tools/pkg/mot"
	"github.com/spf13/cobra"
)

var registration string

var motTestsCmd = &cobra.Command{
	Use:       "tests",
	Short:     "Fetch MOT tests for a certain vehicle",
	ValidArgs: []string{"registration"},
	Run: func(cmd *cobra.Command, args []string) {
		if registration == "" {
			if len(args) == 1 {
				registration = args[0]
			} else if len(args) == 2 {
				registration = args[0] + args[1]
			} else {
				panic("Must set registration")
			}
		}
		client := mot.NewClient(readMOTApiKeyFromEnvVar())
		car, err := client.GetCar(registration)
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("Model: %s %s\nFuel type: %s\nColour: %s\nDate registered: %s\n%d tests found", car.Make, car.Model, car.FuelType, car.PrimaryColour, car.FirstUsedDate, len(car.MotTests)))
		for _, motTest := range car.MotTests {
			fmt.Println(fmt.Sprintf("Date: %s", motTest.CompletedDate))
			fmt.Println(fmt.Sprintf("Result: %s", motTest.TestResult))
			fmt.Println(fmt.Sprintf("Odometer reading: %s%s", motTest.OdometerValue, motTest.OdometerUnit))
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(motTestsCmd)
	motTestsCmd.PersistentFlags().StringVarP(&registration, "registration", "r", "", "Registration of vehicle")
}
