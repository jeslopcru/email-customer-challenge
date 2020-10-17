package customerimporter

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"strings"

	"github.com/spf13/afero"

	"github.com/elliotchance/orderedmap"
)

type CustomerImporter struct {
	file               afero.File
	numberOfColumns    int
	indexOfEmailColumn int
}

func New(file afero.File, numberOfColumns int, indexOfEmailColumn int) *CustomerImporter {
	if indexOfEmailColumn > numberOfColumns {
		log.Fatal("indexOfEmailColumn should be less than numberOfColumns")
	}

	return &CustomerImporter{
		file:               file,
		numberOfColumns:    numberOfColumns,
		indexOfEmailColumn: indexOfEmailColumn,
	}
}

func (ci *CustomerImporter) Run() (*orderedmap.OrderedMap, error) {
	reader := csv.NewReader(bufio.NewReader(ci.file))

	emailGroups := orderedmap.NewOrderedMap()

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if len(line) != ci.numberOfColumns {
			continue
		}

		splitEmail := ci.obtainEmail(line)

		const EmailParts = 2
		if len(splitEmail) != EmailParts {
			continue
		}
		domain := splitEmail[1]
		ci.addElement(emailGroups, domain)
	}

	return emailGroups, nil
}

func (ci *CustomerImporter) addElement(emailGroups *orderedmap.OrderedMap, domain string) {
	value, _ := emailGroups.Get(domain)

	if value == nil {
		emailGroups.Set(domain, 1)
	} else {
		emailGroups.Set(domain, value.(int)+1)
	}
}

func (ci *CustomerImporter) obtainEmail(line []string) []string {
	emailField := line[ci.indexOfEmailColumn]
	splitEmail := strings.Split(emailField, "@")
	return splitEmail
}
