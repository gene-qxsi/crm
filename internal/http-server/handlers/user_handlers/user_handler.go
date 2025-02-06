package user_handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gene-qxsi/CRM-M/internal/models"
	"github.com/gene-qxsi/CRM-M/internal/services"
	"github.com/go-chi/chi/v5"
)

type Userhandler struct {
	service *services.UserService
}

func New(service *services.UserService) *Userhandler {
	return &Userhandler{service: service}
}

// TODO добавить API обработки ошибок, заменить существуюшщий
func (h *Userhandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http-server.handlers.user_handler.CreateUser"
	if h.service == nil {
		log.Println(fmt.Errorf("ERROR: %s. ERROR PATH: %s", "h.service == nil, требуется инициализация", op))
		w.Write([]byte("ошибка 500"))
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.Write([]byte("ошибка в теле запроса " + err.Error()))
		return
	}

	_, err = h.service.CreateUser(user)
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.Write([]byte("не удалось создать учетную запись: " + err.Error()))
		return
	}

	w.Write([]byte("учетная запись успешно создана"))
	log.Println("✅ операция CreateUser - успешно выполнена")
}

func (h *Userhandler) GetUser(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http-server.handlers.user_handler.GetUser"

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

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		w.Write([]byte(err.Error()))
		return
	}

	log.Println("✅ операция GetUser - успешно выполнена")
}

func (h *Userhandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http-server.handlers.user_handler.GetUsers"

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

	w.Header().Set("Content-Type", "application/json")
	for _, u := range users {
		err = json.NewEncoder(w).Encode(u)
		if err != nil {
			log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
			w.Write([]byte(err.Error()))
			return
		}
	}

	log.Println("✅ операция GetUsers - успешно выполнена")
}

func (h *Userhandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http-server.handlers.user_handler.DeleteUser"

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
