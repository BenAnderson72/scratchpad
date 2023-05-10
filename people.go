package scratchpad

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"
)

// create person struct with birthdate,firstname,surname,email
type Person struct {
	Birthdate time.Time
	Firstname string
	Surname   string
	Email     string
}

func (p *Person) String() string {
	return fmt.Sprintf("Date: %s, Firstname: %s, Surname: %s, Email: %s", p.Birthdate, p.Firstname, p.Surname, p.Email)
}

// create a constructor for Person object
func NewPerson(birthdate time.Time, firstname string, surname string, email string) *Person {
	return &Person{birthdate, firstname, surname, email}
}

// method to read people.csv and create an array of Person objects
func ReadPeopleCSV() ([]*Person, error) {

	filename := "./people.csv"

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = 4

	people := make([]*Person, 0)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// convert record[0] to date
		birthdate, err := time.Parse("2006-01-02", record[0])
		if err == nil {

			person := NewPerson(birthdate, record[1], record[2], record[3])
			people = append(people, person)

		}
	}
	return people, nil
}
