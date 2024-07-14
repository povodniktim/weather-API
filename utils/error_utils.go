package utils

import "errors"

// Custom errors
var (
    ErrCityNotFound = errors.New("city not found")
    ErrFetchXMLData = errors.New("error fetching XML data")
    ErrConvertToCity = errors.New("error converting XML data to City struct")
)
