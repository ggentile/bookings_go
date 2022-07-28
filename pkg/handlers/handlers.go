package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ggentile/bookings_go/pkg/config"
	"github.com/ggentile/bookings_go/pkg/models"
	"github.com/ggentile/bookings_go/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
	/*
		n, err := fmt.Fprintf(w, "hello, world!")
		if err != nil {
			log.Println(err)
		}
		fmt.Println(fmt.Sprintf("number of bytes written: %d", n))
	*/
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
	/*
		sum, err := addValues(2, 2)
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
	*/
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var x float32
	var y float32
	x = 100.0
	y = 10.0
	f, err := divideValues(x, y)

	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("forbidden to divide by 0")
		return 0, err
	}

	result := x / y
	return result, nil
}

func addValues(x, y int) (int, error) {
	//var result int
	//return result

	return x + y, nil
}
