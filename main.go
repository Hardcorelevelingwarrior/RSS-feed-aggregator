package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Hardcorelevelingwarrior/RSS-feed-aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main()  {
	godotenv.Load()
	dbURL := os.Getenv("CONN")
	db, err := sql.Open("postgres",dbURL)
	if err != nil {return }
	dbQueries := database.New(db)
	cfg := apiConfig{
		DB: dbQueries,
	}
	port := os.Getenv("PORT")
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{}))
	//Add sub-router /v1
	a := chi.NewRouter()
	router.Mount("/v1",a)
	//Router for /v1
	a.Get("/readiness", func(w http.ResponseWriter, r *http.Request) {
		respondWithJSON(w, 200, map[string]interface{}{"status": "ok"})
	})
	a.Get("/err",func(w http.ResponseWriter, r *http.Request) {
		respondWithError(w,500,"Internal Server Error")
	})
	a.Post("/users",func(w http.ResponseWriter, r *http.Request) {
		request := request {}
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&request)
		param := database.CreateUserParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name: request.Name,
		}
		newUser := database.User{}
		newUser,err = cfg.DB.CreateUser(r.Context(),param)
		if err != nil {respondWithError(w,403,err.Error());return}
		respondWithJSON(w,201,newUser)
	})
	a.Get("/users",func(w http.ResponseWriter, r *http.Request) {
		api_key := strings.TrimPrefix(r.Header.Get("Authorization")," ApiKey ") 
		newUsers,err := cfg.DB.GetUsersByAPIkey(r.Context(),api_key)
		if err != nil {respondWithError(w,404,err.Error());return}
		respondWithJSON(w,200,newUsers)
	})








	//Start server
	srv := http.Server{
		Addr: ":" +port,
		Handler: router,
	}
	log.Fatal(srv.ListenAndServe())
	
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}){
	w.WriteHeader(status)
	w.Header().Add("Content-Type","application/json")
	response, err := json.Marshal(payload)
	if err != nil {
		return 
	}
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w,code,map[string]string{"error": msg})
}

type apiConfig struct {
	DB *database.Queries
}

type request struct {
	Name string `json:"name,omitempty"`
}