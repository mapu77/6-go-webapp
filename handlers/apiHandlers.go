package handlers

import (
	"net/http"
	"github.com/parnurzeal/gorequest"
	"fmt"
	"html/template"
	"encoding/json"
	"strconv"
)

const apiURL = "https://ad-go-api.herokuapp.com/api/v1"

type Flight struct {
	Id            string `json:"id"`
	Code          string `json:"code"`
	Company       string `json:"company"`
	DepartureTime string `json:"departure_time"`
	DepartureCity string `json:"departure_city"`
	ArrivalTime   string `json:"arrival_time"`
	ArrivalCity   string `json:"arrival_city"`
}

type Hotel struct {
	Id			string	`json:"id"`
	Name 		string	`json:"name"`
	HotelChain 	string	`json:"hotel_chain"`
	Rooms		int		`json:"rooms"`
	Street		string	`json:"street"`
	ZipCode		 string	`json:"zip_code"`
	City 		string	`json:"city"`
	Country		string	`json:"country"`
}


type Return struct {
	Username  string
	Flights   []Flight
	Companies []string
	Cities    []string
	Chains	  []string
	Hotels    []Hotel
}

func NewFlight(w http.ResponseWriter, r *http.Request) (int, error) {
	r.ParseForm()
	f := Flight{Code: r.PostFormValue("Code"),
		Company: r.PostFormValue("Company"),
		DepartureTime: r.PostFormValue("DepartureTime"),
		DepartureCity: r.PostFormValue("DepartureCity"),
		ArrivalTime: r.PostFormValue("ArrivalTime"),
		ArrivalCity: r.PostFormValue("ArrivalCity"),}
	request := gorequest.New()
	resp, body, errs := request.Post(apiURL + "/flights").
		Send(f).
		End()
	fmt.Print(resp, body, errs)
	if resp.StatusCode != 500 {
		u := User{Username: getCookieUsername(w, r)}

		t, _ := template.ParseFiles("templates/success.html", "templates/menu.html")
		return http.StatusOK, t.ExecuteTemplate(w, "success.html", u)

	} else {
		return 500, nil
	}
	return resp.StatusCode, nil
}

func GetFlights(w http.ResponseWriter, r *http.Request) (int, error) {
	values := r.URL.Query()
	var codeQuery string
	if len(values["code"]) != 0 {
		code := values["code"][0]
		codeQuery = "code=" + code + "&"
	}
	var companyQuery string
	if len(values["company"]) != 0 {
		company := values["company"][0]
		if company != "Todas" {
			companyQuery = "company=" + company + "&"
		}
	}
	var departureCityQuery string
	if len(values["departureCity"]) != 0 {
		departureCity := values["departureCity"][0]
		if departureCity != "Todas" {
			departureCityQuery = "departureCity=" + departureCity + "&"
		}
	}
	var arrivalCityQuery string
	if len(values["arrivalCity"]) != 0 {
		arrivalCity := values["arrivalCity"][0]
		if arrivalCity != "Todas" {
			arrivalCityQuery = "arrivalCity=" + arrivalCity + "&"
		}
	}
	url := apiURL + "/flights?" + codeQuery + companyQuery + departureCityQuery + arrivalCityQuery
	resp, body, errs := gorequest.New().Get(url).End()
	fmt.Print("resp:", resp)
	fmt.Print("body:", body)
	fmt.Print("errs:", errs)

	if resp.StatusCode != 500 {
		var flights []Flight
		err := json.Unmarshal([]byte(body), &flights)
		if err != nil {
			fmt.Print("ERRORRR")
		}
		//fmt.Print(flights)
		u := User{Username: getCookieUsername(w, r)}

		r := Return{Username: u.Username, Flights: flights}

		t, _ := template.ParseFiles("templates/tablaVuelos.html", "templates/menu.html")
		return http.StatusOK, t.ExecuteTemplate(w, "tablaVuelos.html", r)

	} else {
		return 500, nil
	}
	return resp.StatusCode, nil
}

func GetHotels(w http.ResponseWriter, r *http.Request) (int, error) {
	values := r.URL.Query()
	var nombreQuery string
	if len(values["nombre"]) != 0 {
		nombre := values["nombre"][0]
		nombreQuery = "name=" + nombre + "&"
	}
	var ciudadQuery string
	if len(values["ciudad"]) != 0 {
		ciudad := values["ciudad"][0]
		if ciudad != "Todas" {
			ciudadQuery = "city=" + ciudad + "&"
		}
	}
	var cadenaQuery string
	if len(values["cadena"]) != 0 {
		cadena := values["cadena"][0]
		if cadena != "Todas" {
			cadenaQuery = "hotelChain=" + cadena + "&"
		}
	}

	url := apiURL + "/hotels?" + nombreQuery + ciudadQuery + cadenaQuery
	resp, body, errs := gorequest.New().Get(url).End()
	fmt.Print("resp:", resp)
	fmt.Print("body:", body)
	fmt.Print("errs:", errs)

	if resp.StatusCode != 500 {
		var hotels []Hotel
		err := json.Unmarshal([]byte(body), &hotels)
		if err != nil {
			fmt.Print("ERRORRR")
		}
		//fmt.Print(flights)
		u := User{Username: getCookieUsername(w, r)}

		r := Return{Username: u.Username, Hotels: hotels}

		t, _ := template.ParseFiles("templates/tablaHoteles.html", "templates/menu.html")
		return http.StatusOK, t.ExecuteTemplate(w, "tablaHoteles.html", r)

	} else {
		return 500, nil
	}
	return resp.StatusCode, nil
}

func NewHotel(w http.ResponseWriter, r *http.Request) (int, error) {
	r.ParseForm()
	rooms, err := strconv.Atoi(r.PostFormValue("Rooms"))
	if err != nil {
		return 400, nil
	}
	h := Hotel{	Name:r.PostFormValue("Name"),
				HotelChain:r.PostFormValue("HotelChain"),
				Rooms:rooms,
				Street:r.PostFormValue("Street"),
				ZipCode:r.PostFormValue("ZipCode"),
				City:r.PostFormValue("City"),
				Country:r.PostFormValue("Country"),}
	request := gorequest.New()
	resp, body, errs := request.Post(apiURL + "/hotels").
		Send(h).
		End()
	fmt.Print(resp, body, errs)
	if resp.StatusCode != 500 || resp.StatusCode < 400{
		u := User{Username: getCookieUsername(w, r)}

		t, _ := template.ParseFiles("templates/success.html", "templates/menu.html")
		return http.StatusOK, t.ExecuteTemplate(w, "success.html", u)

	} else if resp.StatusCode == 500 {
		return 500, nil
	} else {
		return 400, nil
	}
	return resp.StatusCode, nil
}
