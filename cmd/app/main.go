package main

import (
	"net/http"

	"github.com/gene-qxsi/CRM-M/internal/http-server/handlers/show_handlers"
	"github.com/gene-qxsi/CRM-M/internal/http-server/handlers/user_handlers"
	"github.com/gene-qxsi/CRM-M/internal/http-server/handlers/user_handlers_html"
	"github.com/gene-qxsi/CRM-M/internal/services"
	"github.com/gene-qxsi/CRM-M/internal/storage"
	"github.com/go-chi/chi/v5"
)

func main() {
	storage := storage.New()
	service := services.New(storage)
	uh := user_handlers.New(service)      // обработчик json
	sh := show_handlers.New(service)      // рендер страниц
	hh := user_handlers_html.New(service) // обработчик html

	router := chi.NewRouter()

	// API - json. html + JavaScript
	router.Get("/api1/users/{id}", uh.GetUser)
	router.Get("/api1/users", uh.GetUsers)
	router.Post("/api1/users", uh.CreateUser)
	router.Delete("/api1/users/{id}", uh.DeleteUser)

	// API - html. only html
	router.Get("/api2/users/{id}", hh.GetUser)
	router.Get("/api2/users", hh.GetUsers)
	router.Post("/api2/register", hh.RegisterUser)
	router.Post("/api2/login", hh.LoginUser)
	router.Post("/api2/logout", hh.LogoutUser)
	router.Post("/api2/users", hh.CreateUser)
	router.Delete("/api2/users/{id}", hh.DeleteUser)

	// рендеринг страниц
	router.Get("/", sh.ShowIndex)
	router.Get("/register", sh.ShowRegister)
	router.Get("/admin_info", sh.ShowAdminInfo)
	router.Get("/login", sh.ShowLogin)
	router.Get("/logout", sh.ShowLogout)
	http.ListenAndServe(":8080", router)

	// TODO: настроить html и js страницы
}
