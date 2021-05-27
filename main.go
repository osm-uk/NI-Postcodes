package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, district := range []string{"801", "802", "803", "804", "805", "807", "809", "811", "813", "814", "816"} {
		fmt.Println(district)

		resp, err := http.Get("http://ratings.food.gov.uk/OpenDataFiles/FHRS" + district + "en-GB.xml")
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer resp.Body.Close()

		var fhrsData = FHRSEstablishment{}
		err = xml.NewDecoder(resp.Body).Decode(&fhrsData)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if fhrsData.Header.ReturnCode != "Success" {
			fmt.Println(fhrsData.Header.ReturnCode)
			continue
		}

		file, err := os.Create("data/" + district + ".csv")
		if err != nil {
			fmt.Println(err)
			break
		}
		defer file.Close()

		for _, establishment := range fhrsData.EstablishmentCollection.EstablishmentDetail {

			if establishment.PostCode == "" {
				continue
			}
			if establishment.Geocode.Latitude == "" {
				continue
			}
			if establishment.Geocode.Longitude == "" {
				continue
			}

			_, err = io.WriteString(file,
				strings.ToUpper(establishment.PostCode)+","+
					establishment.Geocode.Latitude+","+
					establishment.Geocode.Longitude+"\n")
		}

		err = file.Sync()
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
