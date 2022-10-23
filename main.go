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

	err := http.ListenAndServe(":8080", router); err != nil {
		log.Println(err)
	}

}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello from Chi Router !")); err != nil {
		log.Println(err)
	}

}

func GetWithParamHandler(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "parameter")
	_, err := w.Write([]byte("Hello with parameter: " + param)); err != nil {
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