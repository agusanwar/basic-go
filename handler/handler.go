package handler

import (
	"go-web/entity"
	"log"
	"net/http"
	"path"
	"strconv"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// kondisi jika bukan root
	log.Print(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// view home
	tmpl, err := template.ParseFiles(path.Join("view", "index.html"), path.Join("view", "layout.html"))

	 // // data
	// data := map[string] interface{} {
	// 	"title": "golang | web",
	// 	"content": "Programming Language with goloang",
	// }
	
	// jika error
	if err != nil {
		log.Println(err)
		http.Error(w,"Error is view", http.StatusInternalServerError)
		return
	}

	// pass data to view slise
	data := []entity.Product{
		{ID: 1, Name: "Macbook Pro", Price: 1000000, Stock: 10},
		{ID: 2, Name: "Macbook Air", Price: 800000, Stock: 5},
		{ID: 3, Name: "Macbook M!", Price: 1000000, Stock: 10},
	}

	// jika tidak error
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w,"Error is view", http.StatusInternalServerError)
		return
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About page"))
}
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("profile page"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNum, err := strconv.Atoi(id)

	if err != nil || idNum < 1 {
		http.NotFound(w, r)
		return
	}
	// fmt.Fprintf(w, "Product page: %s", id)

	// data
	// data := map[string] interface{} {
	// 	"id": idNum,
	// }

	// pass data to view
	data := entity.Product{
		ID: 1,
		Name: "Macbook Pro",
		Price: 1000000,
		Stock: 10,
	}
	// view product
	tmpl, err := template.ParseFiles(path.Join("view", "product.html"), path.Join("view", "layout.html"))
	// jika error
	if err != nil {
		log.Println(err)
		http.Error(w,"Error is view", http.StatusInternalServerError)
		return
	}
	// jika tidak error
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w,"Error is view", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method // GET or POST

	switch method {
	case "GET":
		w.Write([]byte("GET"))
	case "POST":
		w.Write([]byte("POST"))
	default:
		http.Error(w, "Error method", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		// view product
		tmpl, err := template.ParseFiles(path.Join("view", "form.html"), path.Join("view", "layout.html"))
		// jika error
		if err != nil {
			log.Println(err)
			http.Error(w,"Error is view", http.StatusInternalServerError)
			return
		}
		// jika tidak error
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w,"Error is view", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "Error method", http.StatusBadRequest)
}

func ProsesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		err := r.ParseForm()
		// jika error
		if err != nil {
			log.Println(err)
			http.Error(w,"Error is view", http.StatusInternalServerError)
			return
		}
			name := r.Form.Get("name")
			message := r.Form.Get("message")

			data := map[string]interface{}{
				"name": name,
				"message": message,
			}
			tmpl, err := template.ParseFiles(path.Join("view", "resault.html"), path.Join("view", "layout.html"))
			// jika error
			if err != nil {
				log.Println(err)
				http.Error(w,"Error is view", http.StatusInternalServerError)
				return
			}
			// jika tidak error
			err = tmpl.Execute(w, data)
			if err != nil {
				log.Println(err)
				http.Error(w,"Error is view", http.StatusInternalServerError)
				return
			}
			return
	}
	http.Error(w, "Error method", http.StatusBadRequest)
}