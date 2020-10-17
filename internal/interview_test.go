package customerimporter_test

import (
	"path/filepath"
	"testing"

	"github.com/spf13/afero"

	customerimporter "github.com/jeslopcru/email-customer-challenge/internal"
)

func TestNewCustomer_ErrorHandling(t *testing.T) {
	cases := []struct {
		fixture         string
		returnErr       bool
		name            string
		numberOfColumns int
		indexOfEmail    int
	}{
		{
			fixture:         "testdata/empty.csv",
			returnErr:       false,
			name:            "EmptyFile",
			numberOfColumns: 6,
			indexOfEmail:    5,
		},
		{
			fixture:         "testdata/valid.csv",
			returnErr:       false,
			name:            "ValidFile",
			numberOfColumns: 9,
			indexOfEmail:    5,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			csvFilepath, _ := filepath.Abs(tc.fixture)
			fileOpen, _ := afero.NewOsFs().Open(csvFilepath)

			importer := customerimporter.New(fileOpen, tc.numberOfColumns, tc.indexOfEmail)
			_, err := importer.Run()
			returnedErr := err != nil

			if returnedErr != tc.returnErr {
				t.Fatalf("Expected returnErr: %v, got: %v", tc.returnErr, returnedErr)
			}
		})
	}
}

func TestRun(t *testing.T) {
	t.Run("should ignore not email data", func(t *testing.T) {
		csvFilepath, _ := filepath.Abs("testdata/valid.csv")
		fileOpen, _ := afero.NewOsFs().Open(csvFilepath)

		importer := customerimporter.New(fileOpen, 5, 1)
		result, _ := importer.Run()
		if result.Len() > 1 {
			t.Fatalf("Expected empty result: got: %v", result.Len())
		}
	})

	t.Run("should works", func(t *testing.T) {
		csvFilepath, _ := filepath.Abs("testdata/valid.csv")
		fileOpen, _ := afero.NewOsFs().Open(csvFilepath)

		importer := customerimporter.New(fileOpen, 5, 2)
		result, _ := importer.Run()
		if result.Len() != 1 {
			t.Fatalf("Expected one result: got: %v", result.Len())
		}
	})
}
