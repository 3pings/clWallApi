package route

import (
	"encoding/json"
	"fmt"
	"github.com/3pings/clWallApi/config"
	"net/http"
)

type Services struct {
	Weather struct {
		Description string  `json:"description"`
		Icon        string  `json:"icon"`
		Temp        float64 `json:"temp"`
		TempMin     float64 `json:"temp_min"`
		TempMax     float64 `json:"temp_max"`
		Humidity    float64 `json:"humidity"`
		City        string  `json:"city"`
	} `json:"weather"`
	Event struct {
		Name         string `json:"name"`
		StartDate    string `json:"start_date"`
		VenueName    string `json:"venue_name"`
		VenueAddress string `json:"venue_address"`
		VenueCity    string `json:"venue_city"`
		EventUrl     string `json:"event_url"`
		LogoUrl      string `json:"logo_url"`
	} `json:"event"`
	Incident struct {
		Severity    int    `json:"severity"`
		Coordinates string `json:"coordinates"`
		Description string `json:"description"`
	} `json:"incident"`
}

func GetServices(w http.ResponseWriter, r *http.Request) {

	//Set Header Info
	w.Header().Set("Content-Type", "application/json")

	//Set Variables
	var ss []Services
	var s Services

	//Get Event Data
	event, err := config.DB.Query("select name, start_date, venue_name, venue_address, venue_city, event_url, logo_url from events")
	if err != nil {
		panic(err.Error())
	}
	event.Next()

	err = event.Scan(&s.Event.Name, &s.Event.StartDate, &s.Event.VenueName, &s.Event.VenueAddress, &s.Event.VenueCity, &s.Event.EventUrl, &s.Event.LogoUrl)
	if err != nil {
		panic(err.Error())
	}

	//Get Weather Data
	weather, err := config.DB.Query("select description, icon, temp, temp_min, temp_max, humidity, city from weather order by timestamp DESC")
	if err != nil {
		panic(err.Error())
	}
	weather.Next()
	err = weather.Scan(&s.Weather.Description, &s.Weather.Icon, &s.Weather.Temp, &s.Weather.TempMin, &s.Weather.TempMax, &s.Weather.Humidity, &s.Weather.City)
	if err != nil {
		panic(err.Error())
	}

	//Get Incident Data
	incident, err := config.DB.Query("select severity, coordinates, description from incidents")
	if err != nil {
		panic(err.Error())
	}
	incident.Next()
	err = incident.Scan(&s.Incident.Severity, &s.Incident.Coordinates, &s.Incident.Description)
	if err != nil {
		panic(err.Error())
	}
	ss = append(ss, s)

	fmt.Println(ss)

	json.NewEncoder(w).Encode(ss)

}
