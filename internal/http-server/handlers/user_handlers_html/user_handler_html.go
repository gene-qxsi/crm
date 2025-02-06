package user_handlers_html

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gene-qxsi/CRM-M/internal/models"
	"github.com/gene-qxsi/CRM-M/internal/services"
	"github.com/go-chi/chi/v5"
)

// НЕ РЕКОМЕНДУЕТСЯ К ИСПОЛЬЗОВАНИЮ

type UserhandlerHTML struct {
	service *services.UserService
}

func New(service *services.UserService) *UserhandlerHTML {
	return &UserhandlerHTML{service: service}
}

func (h *UserhandlerHTML) CreateUser(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http-server.handlers.userhandlershtml.CreateUser"
	if h.service == nil {
		log.Println(fmt.Errorf("ERROR: %s. ERROR PATH: %s", "h.service == nil, требуется инициализация", op))
		w.Write([]byte("ошибка 500"))
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.Write([]byte("ошибка в теле запроса " + err.Error()))
		return
	}

	name := r.FormValue("name")
	age, _ := strconv.Atoi(r.FormValue("age"))
	password := r.FormValue("password")

	user := models.User{Name: name, Age: age, Password: password}

	_, err = h.service.CreateUser(user)
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.Write([]byte("не удалось создать учетную запись: " + err.Error()))
		return
	}

	w.Write([]byte("учетная запись успешно создана"))
	log.Println("✅ операция CreateUser - успешно выполнена")
}

func (h *UserhandlerHTML) GetUser(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http-server.handlers.userhandlershtml.GetUser"

	if h.service == nil {
		log.Println(fmt.Errorf("ERROR: %s. ERROR PATH: %s", "h.service == nil, требуется инициализация", op))
		w.Write([]byte("ошибка 500"))
		return
	}

	_id := chi.URLParam(r, "id")
	if _id == "" {
		log.Println(fmt.Errorf("ERROR: %s. ERROR PATH: %s", "параметр id не установлен", op))
		w.Write([]byte("параметр id не установлен"))
		return
	}

	id, err := strconv.Atoi(_id)
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.Write([]byte(err.Error()))
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
			w.Write([]byte(fmt.Sprintf("пользователя с id = %d не существует", id)))
			return
		}
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	data := url.Values{}
	data.Add("id", strconv.Itoa(user.ID))
	data.Add("name", user.Name)
	data.Add("age", strconv.Itoa(user.Age))
	data.Add("password", user.Password)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(data.Encode()))
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.Write([]byte(err.Error()))
		return
	}

	log.Println("✅ операция GetUser - успешно выполнена")
}

func (h *UserhandlerHTML) GetUsers(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http-server.handlers.userhandlershtml.GetUsers"

	if h.service == nil {
		log.Println(fmt.Errorf("ERROR: %s. ERROR PATH: %s", "h.service == nil, требуется инициализация", op))
		w.Write([]byte("ошибка 500"))
		return
	}

	users, err := h.service.GetUsers()
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	val := url.Values{}
	for _, u := range users {
		val.Add("id", strconv.Itoa(u.ID))
		val.Add("name", u.Name)
		val.Add("age", strconv.Itoa(u.Age))
		val.Add("password", u.Password)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(val.Encode()))
	log.Println("✅ операция GetUsers - успешно выполнена")
}

func (h *UserhandlerHTML) DeleteUser(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http-server.handlers.userhandlershtml.DeleteUser"

	if h.service == nil {
		log.Println(fmt.Errorf("ERROR: %s. ERROR PATH: %s", "h.service == nil, требуется инициализация", op))
		w.Write([]byte("ошибка 500"))
		return
	}

	_id := chi.URLParam(r, "id")

	id, err := strconv.Atoi(_id)
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	log.Println("✅ операция DeleteUsers - успешно выполнена")
}

func (h *UserhandlerHTML) LoginUser(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http-server.handlers.userhandlershtml.LoginUser"

	err := r.ParseForm()
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	name := r.FormValue("name")
	password := r.FormValue("password")

	user, err := h.service.GetUserByNameAndPassword(name, password)
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "учетная запись не найден", http.StatusUnauthorized)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "session_id",
		Value: fmt.Sprintf("%d", user.ID),
		Path:  "/",
	})

	http.Redirect(w, r, "/app", http.StatusSeeOther)
	log.Println("✅ операция LoginUser - успешно выполнена")
}

func (h *UserhandlerHTML) RegisterUser(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http-server.handlers.userhandlershtml.RegistrationUser"

	err := r.ParseForm()
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	name := r.FormValue("name")
	age, _ := strconv.Atoi(r.FormValue("age"))
	password := r.FormValue("password")

	user := models.User{Name: name, Age: age, Password: password}

	id, err := h.service.CreateUser(user)
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "не удалось создать учетную запись", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "session_id",
		Value: fmt.Sprintf("%d", id),
		Path:  "/",
	})

	http.Redirect(w, r, "/app", http.StatusSeeOther)
	log.Println("✅ операция RegisterUser - успешно выполнена")
}

func (h *UserhandlerHTML) LogoutUser(w http.ResponseWriter, r *http.Request) {
	// const op = "internal.http-server.handlers.userhandlershtml.LogoutUser"

	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	http.Redirect(w, r, "/app", http.StatusSeeOther)
	log.Println("✅ операция LogoutUser - успешно выполнена")
}
