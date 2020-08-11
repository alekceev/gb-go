package main

// 3. Дополните функцию hello() http-сервера так, чтобы принять и отобразить на странице один GET-параметр, например name.
// При этом в браузере запрос будет выглядеть так: http://localhost/hello?name=World.
// Значение этого параметра надо получить в функции и отобразить при выводе html-кода.

import (
	"log"
	"net/http"
	"text/template"
)

func hello(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	if len(name) == 0 {
		name = "World"
	}
	res.Header().Set("Content-Type", "text/html")

	t := template.Must(template.New("").Parse(`<doctype html>
		<html>
			<head>
				<title>Hello {{.Name}}!</title>
			</head>
			<body>
				Hello {{.Name}}!
			</body>
		</html>`))

	err := t.Execute(res, struct{ Name string }{Name: name})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8000", nil)
}
