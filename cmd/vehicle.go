package cmd

import (
	"context"
	"fmt"
	"github.com/issy/dvla-tools/pkg/ves"
	"github.com/spf13/cobra"
	"os"
)

var vehicleRegistration string

var vehicleDetailsCmd = &cobra.Command{
	Use:       "vehicle",
	Short:     "Fetch info about a certain vehicle",
	ValidArgs: []string{"registration"},
	Run: func(cmd *cobra.Command, args []string) {
		if vehicleRegistration == "" {
			if len(args) == 1 {
				vehicleRegistration = args[0]
			} else if len(args) == 2 {
				vehicleRegistration = args[0] + args[1]
			} else {
				panic("Must set registration")
			}
		}

		vehicleRequest := ves.VehicleRequest{RegistrationNumber: &vehicleRegistration} // VehicleRequest | Registration number of the vehicle to find details for
		xCorrelationId := "xCorrelationId_example"                                     // string | Consumer Correlation ID (optional)

		xApiKey := readVESApiKeyFromEnvVar()
		configuration := ves.NewConfiguration()
		apiClient := ves.NewAPIClient(configuration)
		resp, r, err := apiClient.VehicleApi.GetVehicleDetailsByRegistrationNumber(context.Background()).XApiKey(xApiKey).VehicleRequest(vehicleRequest).XCorrelationId(xCorrelationId).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `VehicleApi.GetVehicleDetailsByRegistrationNumber``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		// response from `GetVehicleDetailsByRegistrationNumber`: Vehicle
		fmt.Println(fmt.Sprintf("Make: %s", resp.GetMake()))
		fmt.Println(fmt.Sprintf("Colour: %s", resp.GetColour()))
		fmt.Println(fmt.Sprintf("Fuel type: %s", resp.GetFuelType()))
		fmt.Println(fmt.Sprintf("MOT status: %s", resp.GetMotStatus()))
		fmt.Println(fmt.Sprintf("Tax status: %s", resp.GetTaxStatus()))
	},
}

func init() {
	rootCmd.AddCommand(vehicleDetailsCmd)
	motTestsCmd.PersistentFlags().StringVarP(&vehicleRegistration, "vehicleRegistration", "v", "", "Registration of vehicle")
}
