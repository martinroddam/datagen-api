package main

import (
	"strings"

	"github.com/manveru/faker"
	log "github.com/sirupsen/logrus"
)

// GenerateRecords produces the number of specified Record and returns in a RecordSet
func GenerateRecords(numRecords int) (RecordSet, error) {

	log.Info("Starting data generation...")

	fake, err := faker.New("en")
	if err != nil {
		return RecordSet{}, err
	}

	var recordSet RecordSet
	var records []Record
	var stats Stats
	stats.CountryCounts = newCountryCountMap()

	for i := 0; i < numRecords; i++ {

		var rec Record
		var n Name

		rec.FullName = formatName(n.New(fake))
		rec.Address = fake.StreetAddress()
		rec.City = fake.City()
		rec.State = fake.State()
		rec.Postcode = fake.PostCode()
		rec.Country = fake.Country()
		rec.Email = fake.Email()
		rec.HomePhone = fake.PhoneNumber()
		rec.MobilePhone = fake.CellPhoneNumber()

		stats = evaluateCountry(fake.Country(), stats)

		records = append(records, rec)
		stats.RecordCount++
	}

	recordSet.Records = records
	recordSet.Statistics = stats

	return recordSet, nil
}

func formatName(name Name) string {
	return strings.ToUpper(strings.TrimSpace(name.Prefix) + " " + strings.TrimSpace(name.First) + " " + strings.TrimSpace(name.Last))
}

func newCountryCountMap() map[string]int {
	m := make(map[string]int)
	for _, country := range countCountries {
		m[country] = 0
	}
	return m
}

// New generates a new Name object
func (n *Name) New(fake *faker.Faker) Name {
	n.Prefix = fake.NamePrefix()
	n.First = fake.FirstName()
	n.Last = fake.LastName()
	return *n
}

func evaluateCountry(country string, stats Stats) Stats {
	if requiresCounting(country) {
		stats.CountryCounts[country]++
		return stats
	}
	return stats
}

func requiresCounting(country string) bool {
	for _, c := range countCountries {
		if c == country {
			return true
		}
	}
	return false
}
