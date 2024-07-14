package utils

import (
	"encoding/xml"
	"io"
	"net/http"

	models "GitHub.com/weatherAPI/models/xml"
)

// FetchXMLData fetches XML data from the given URL and checks if it matches the expected structure
func FetchXMLData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err // Return the specific error
	}
	defer resp.Body.Close()

	// Read the XML data
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Validate the XML structure
	var weatherData models.Data
	if err := xml.Unmarshal(data, &weatherData); err != nil {
		return nil, err
	}

	return data, nil
}

// ParseXML parses the XML data into a Data struct
func ParseXML(data []byte) (models.Data, error) {
	var weatherData models.Data
	err := xml.Unmarshal(data, &weatherData)
	return weatherData, err
}
