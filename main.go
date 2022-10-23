package main

import(
	"github.com/go-chi/chi"
	"log"
	"net/htpp"
)

func main(){
	// Create new router
	router := chi.NewRouter()

	router.Use(Logger) // register middleware with router

	router.Get("/", MainHandler)
	router.Get("/api/{parameter}", GetWithParamHandler)
	router.Post("/api/add", PostHandler)

	err := http.ListenAndServe(":8080", router); err != nil {
		log.Println(err)
	}

}

// [GET] localhost:8080/
func MainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := w.Write([]byte("Hello from Chi Router !")); err != nil {
		log.Println(err)
	}

}

// [GET] localhost:8080/api/{parameter}
func GetWithParamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := chi.URLParam(r, "parameter")
	_, err := w.Write([]byte("Hello with parameter: " + param)); err != nil {
		log.Println(err)
	}
}

// [POST] localhost:8080/api/add
func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := w.Write([]byte("Hello from Chi Router [POST] !")); err != nil {
		log.Println(err)
	}	
}

// Chi Middleware
func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)) {
		log.Println(r.URL.Path)
		handler.ServeHTTP(w, r)
	}
}