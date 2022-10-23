package main

import(
	"github.com/go-chi/chi"
	"log"
	"net/htpp"
)

func main(){
	// Create new router
	router := chi.NewRouter()

	router.Get("/", MainHandler)

	err := http.ListenAndServe(":8080", router); err != nil{
		log.Println(err)
	}

}

func MainHandler(w http.ResponseWriter, r *http.Request){
	_, err := w.Write([]byte("Hello from Chi Router !")); err != nil{
		log.Println(err)
	}

}