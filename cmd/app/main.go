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

	// router.Handle("/internal/static/js/*", http.StripPrefix("/internal/static/js/", http.FileServer(http.Dir("internal/static/js"))))

	http.ListenAndServe(":8080", router)

}

func setJSHandlers(router *chi.Mux, uh *user_handlers.Userhandler) {
	router.Get("/api/users/{id}", uh.GetUser)
	router.Get("/api/users", uh.GetUsers)
	router.Post("/api/users", uh.CreateUser)
	router.Post("/api/register", uh.RegisterUser)
	router.Delete("/api/users/{id}", uh.DeleteUser)
	router.Post("/api/login", uh.LoginUser)
	router.Post("/api/logout", uh.LogoutUser)
}

func setShowHandlers(router *chi.Mux, sh *show_handlers.Show) {

	// router.Get("/", sh.ShowIndex)
	// router.Get("/register", sh.ShowRegister)
	// router.Get("/admin_info", sh.ShowAdminInfo)
	// router.Get("/login", sh.ShowLogin)
	// router.Get("/logout", sh.ShowLogout)
}
