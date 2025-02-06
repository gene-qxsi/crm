package show_handlers

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gene-qxsi/CRM-M/internal/models"
	"github.com/gene-qxsi/CRM-M/internal/services"
)

type Show struct {
	service *services.UserService
}

type IndexData struct {
	Role            string
	Name            string
	IsAuthenticated bool
}

type AdminData struct {
	Users           []models.User
	IsAuthenticated bool
}

func New(s *services.UserService) *Show {
	return &Show{service: s}
}

func (s *Show) ShowIndex(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http-server.handlers.show_handlers.show_handler.ShowIndex"

	tmpl, err := template.ParseFiles("internal/pages/html/index.html")
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.Write([]byte(err.Error()))
		return
	}

	var data IndexData
	r.ParseForm()
	cook, err := r.Cookie("session_id")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			data = IndexData{Name: "", IsAuthenticated: false}
		} else {
			log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
			w.Write([]byte(err.Error()))
			return
		}
	} else {
		id, err := strconv.Atoi(cook.Value)
		log.Println(id)
		if err != nil {
			log.Printf("❌ ERROR: Невозможно преобразовать значение куки в ID: %s. PATH: %s\n", err, op)
			w.Write([]byte("Ошибка авторизации"))
			return
		}

		user, err := s.service.GetUser(id)
		if err != nil {
			log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
			w.Write([]byte(err.Error()))
			return
		}
		log.Println(user.Role)
		data = IndexData{Role: user.Role, Name: user.Name, IsAuthenticated: true}
	}

	tmpl.Execute(w, data)
}

func (s *Show) ShowRegister(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http-server.handlers.show_handlers.show_handler.ShowRegistration"
	tmpl, err := template.ParseFiles("internal/pages/html/register.html")
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.Write([]byte(err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func (s *Show) ShowAdminInfo(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http-server.handlers.show_handlers.show_handler.ShowAdminInfo"
	tmpl, err := template.ParseFiles("internal/pages/html/admin_info.html")
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = r.Cookie("session_id")
	isAuthenticated := err == nil

	users, err := s.service.GetUsers()
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.Write([]byte(err.Error()))
		return
	}

	data := AdminData{Users: users, IsAuthenticated: isAuthenticated}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.Write([]byte(err.Error()))
		return
	}
}

func (s *Show) ShowLogin(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http-server.handlers.show_handlers.show_handler.ShowLogin"
	tmpl, err := template.ParseFiles("internal/pages/html/login.html")
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.Write([]byte(err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func (s *Show) ShowLogout(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http-server.handlers.show_handlers.show_handler.ShowLogout"
	tmpl, err := template.ParseFiles("internal/pages/html/logout.html")
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.Write([]byte(err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}
