package models

type MetData struct {
	DomainCountryCode string  `xml:"domain_countryIsoCode2`
	DomainShortTitle  string  `xml:"domain_shortTitle"`
	TsUpdated         string  `xml:"tsUpdated"`
	Day               string  `xml:"valid_day"`
	Valid             string  `xml:"valid"`
	Icon              string  `xml:"nn_icon"`
	TempLow           int     `xml:"tnsyn_degreesC"`
	TempMax           int     `xml:"txsyn_degreesC"`
	TempUnit          string  `xml:"tnsyn_var_unit"`
	WindIcon          string  `xml:"ddff_icon"`
	WindDir           string  `xml:"dd_shortText"`
	WindDirLong       string  `xml:"dd_longText"`
	WindSpeed         float64 `xml:"ff_val"`
	WindUnit          string  `xml:"ff_var_unit"`
}