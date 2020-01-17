package route

import (
	"encoding/json"
	"github.com/3pings/clWallApi/config"
	"net/http"
	"os"
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
	AppSpecs struct {
		AppVersion      string `json:"app_version"`
		ServingHostname string `json:"service_hostname"`
	} `json:"appspecs"`
}

func GetServices(w http.ResponseWriter, r *http.Request) {

	//Connect to DB
	db := config.DBConn()

	//Set Header Info
	w.Header().Set("Content-Type", "application/json")

	//Set Variables
	var ss []Services
	var s Services
	hn, _ := os.Hostname()
	appVer := "1.0"

	//Get Event Data
	event, err := db.Query("select name, start_date, venue_name, venue_address, venue_city, event_url, logo_url from events")
	if err != nil {
		panic(err.Error())
	}
	event.Next()

	err = event.Scan(&s.Event.Name, &s.Event.StartDate, &s.Event.VenueName, &s.Event.VenueAddress, &s.Event.VenueCity, &s.Event.EventUrl, &s.Event.LogoUrl)
	if err != nil {
		panic(err.Error())
	}
	event.Close()

	//Get Weather Data
	weather, err := db.Query("select description, icon, temp, temp_min, temp_max, humidity, city from weather order by timestamp DESC")
	if err != nil {
		panic(err.Error())
	}
	weather.Next()

	err = weather.Scan(&s.Weather.Description, &s.Weather.Icon, &s.Weather.Temp, &s.Weather.TempMin, &s.Weather.TempMax, &s.Weather.Humidity, &s.Weather.City)
	if err != nil {
		panic(err.Error())
	}
	weather.Close()

	//Get Incident Data
	incident, err := db.Query("select severity, coordinates, description from incidents")
	if err != nil {
		panic(err.Error())
	}
	incident.Next()

	err = incident.Scan(&s.Incident.Severity, &s.Incident.Coordinates, &s.Incident.Description)
	if err != nil {
		panic(err.Error())
	}
	incident.Close()

	s.AppSpecs.ServingHostname = hn
	s.AppSpecs.AppVersion = appVer

	ss = append(ss, s)
	defer db.Close()

	json.NewEncoder(w).Encode(ss[0])

}
