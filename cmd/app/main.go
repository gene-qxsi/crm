package main

import (
	"net/http"

	"github.com/gene-qxsi/CRM-M/internal/http-server/handlers/show_handlers"
	"github.com/gene-qxsi/CRM-M/internal/http-server/handlers/user_handlers"
	"github.com/gene-qxsi/CRM-M/internal/services"
	"github.com/gene-qxsi/CRM-M/internal/storage"
	"github.com/go-chi/chi/v5"
)

func main() {
	storage := storage.New()
	service := services.New(storage)
	uh := user_handlers.New(service) // обработчик json
	sh := show_handlers.New(service) // рендер страниц

	router := chi.NewRouter()

	setJSHandlers(router, uh)   // API json
	setShowHandlers(router, sh) // show html pages

	router.Handle("/internal/static/js/*", http.StripPrefix("/internal/static/js/", http.FileServer(http.Dir("internal/static/js"))))

	http.ListenAndServe(":8080", router)

}

func setJSHandlers(router *chi.Mux, uh *user_handlers.Userhandler) {
	router.Get("/users/{id}", uh.GetUser)
	router.Get("/users", uh.GetUsers)
	router.Post("/users", uh.CreateUser)
	router.Post("/register", uh.RegisterUser)
	router.Delete("/users/{id}", uh.DeleteUser)
	router.Post("/login", uh.LoginUser)
	router.Post("/logout", uh.LogoutUser)
}

func setShowHandlers(router *chi.Mux, sh *show_handlers.Show) {
	router.Get("/", sh.ShowIndex)
	router.Get("/register_view", sh.ShowRegister)
	router.Get("/admin_info_view", sh.ShowAdminInfo)
	router.Get("/login_view", sh.ShowLogin)
	router.Get("/logout_view", sh.ShowLogout)
}
