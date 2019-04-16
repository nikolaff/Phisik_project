package main

import (
	"encoding/json"
	"net/http"
	"os"
)

//Структура для входных параметров
//Имена переменныех должны начинаться с большой буквы
type Data struct {
	X float64 `json: "x"`
	Y float64 `json: "y"`
}

//Структура данных результата
type RData struct {
	Phii float64 `json: "r"`
}

//обработчик для функции Ajax запроса расчет потенциала поля по координатам
func myfuncAjax(w http.ResponseWriter, r *http.Request) {
	var d Data
	var rp RData
	//Читаем тело запроса r.Body и преобразуем данные в структуру Data
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	rp.Phii = PhiMain(d.X, d.Y) // Функция вычисления

	// Сознаем json response from struct RData
	a, err := json.Marshal(rp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Задаем заголовок данным для возврата ответа
	w.Header().Set("Content-Type", "application/json")
	//Возвращаем результат в виде json структуры
	w.Write(a)
}

func main() {
	/*
	   	http.Handle("/", http.FileServer(http.Dir("static")))
	   	// Определяем маршрут и обработчик для функции Ajax запроса
	   	http.HandleFunc("/myfunc", myfuncAjax)

	   	http.ListenAndServe(":8081", nil)
	   }
	*/
	// Heroku прокидывает порт для приложения в переменную окружения PORT

	port := os.Getenv("PORT")

	http.Handle("/", http.FileServer(http.Dir("static")))
	// Определяем маршрут и обработчик для функции Ajax запроса
	http.HandleFunc("/myfunc", myfuncAjax)
	http.ListenAndServe(":"+port, nil)
}

//log.Fatal()  // Спросить
