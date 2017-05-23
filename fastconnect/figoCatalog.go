package fastconnect

import (
	"encoding/json"
	"log"

	"github.com/TobiEiss/fiGo"
)

// ReadIndividualCatalogEntry returns catlo entry for serviceID
func ReadIndividualCatalogEntry(connection fiGo.IConnection, accessToken string, catalogCategory string, countryCode string, serviceID string) (CatalogEntry, error) {
	var catalogEntry CatalogEntry
	answerByte, err := connection.ReadIndividualCatalogEntry(accessToken, catalogCategory, countryCode, serviceID)
	if err != nil {
		return catalogEntry, err
	}

	log.Println(string(answerByte))

	// try to unmarshal
	err = json.Unmarshal(answerByte, &catalogEntry)
	if err != nil {
		return catalogEntry, err
	}

	return catalogEntry, err
}
