package cli

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/afero"

	customerimporter "github.com/jeslopcru/email-customer-challenge/internal"

	"github.com/spf13/cobra"
)

type CobraFn func(cmd *cobra.Command, args []string)

// InitCustomerCmd initialize customerimporter command.
func InitCustomerCmd() *cobra.Command {
	customersCmd := &cobra.Command{
		Use:   "customers",
		Long:  "reads from the given customers.csv file and returns a sorted email domains along with the number of customers with e-mail addresses for each domain",
		Short: "read customer.csv and returns number of customer of each domain",
		Run:   runCustomerFn(),
	}

	return customersCmd
}

func runCustomerFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		csvFilepath, err := filepath.Abs("data/customers.csv")
		if err != nil {
			log.Fatal("Cannot read file path")
		}

		fileOpen, err := afero.NewOsFs().Open(csvFilepath)
		if err != nil {
			log.Fatalln("Couldn't open the csv file", err)
		}

		importer := customerimporter.New(fileOpen, 5, 2)

		emailGroups, err := importer.Run()
		if err != nil {
			log.Fatal("Read failed")
		}
		defer fileOpen.Close()

		for el := emailGroups.Front(); el != nil; el = el.Next() {
			fmt.Println(el.Key, el.Value)
		}
	}
}
