package rest

import (
	"github.com/gorilla/mux"
)

func InitHandlers() *mux.Router {
	route := mux.NewRouter()

	// Наименование всех handler фуникций должно начинаться с маленькой буквы,
	// вызов данных выукнций должен быть доступен только внутри пакета.

	// Требуется описывать методы на каждый handler. Можно объеденить
	// в одном handler обработку нескольких методов {GET|POST|UPDATE|DELETE}.
	// При таком подхоте требуется чётко выделить в функциях резделение логики
	// или разбить её на большее количество функций.

	// Запись ответа во всех handler рекомендуется записывать
	// через функцию - writeRequest.
	route.HandleFunc("/api/", tmpHandler).Methods("GET")
	route.HandleFunc("/api/data", setData).Methods("POST")

	return route
}
