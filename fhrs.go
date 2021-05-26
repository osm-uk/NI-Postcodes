package main

import "encoding/xml"

type FHRSEstablishment struct {
	XMLName xml.Name `xml:"FHRSEstablishment"`
	Header  struct {
		ExtractDate string `xml:"ExtractDate"`
		ItemCount   string `xml:"ItemCount"`
		ReturnCode  string `xml:"ReturnCode"`
	} `xml:"Header"`
	EstablishmentCollection struct {
		EstablishmentDetail []struct {
			FHRSID   string `xml:"FHRSID"`
			PostCode string `xml:"PostCode"`
			Geocode  struct {
				Longitude string `xml:"Longitude"`
				Latitude  string `xml:"Latitude"`
			} `xml:"Geocode"`
		} `xml:"EstablishmentDetail"`
	} `xml:"EstablishmentCollection"`
}
