package models

import "encoding/xml"

type Data struct {
	XMLName		xml.Name	`xml:"data"`
	ID       	string    	`xml:"id,attr"`
	Language 	string    	`xml:"language"`
	IconUrlBase string	   	`xml:"icon_url_base"`
	IconFormat  string	   	`xml:"icon_format"`
	MetData		[]MetData	`xml:"metData"`
}