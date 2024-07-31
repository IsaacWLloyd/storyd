package main

/*POST   /lobbies              # Create a new lobby
GET    /lobbies/{id}         # Get lobby details
POST   /lobbies/{id}/join    # Join a lobby
POST   /games/{id}/sentences # Submit a sentence*/

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/IsaacLloyd/internal/game/lobby"
)

type Server struct {
	router   *mux.Router
	upgrader websocket.Upgrader
	lobbies  map[string]*Lobby
}

func NewServer() *Server {
	s := &Server{
		router: mux.NewRouter(),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		lobbies: make(map[string]*Lobby),
	}

	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	s.router.HandleFunc("/lobbies/{id}", s.GetLobby).Methods("GET")
	s.router.HandleFunc("/lobbies", s.CreateLobby).Methods("POST")
	s.router.HandleFunc("/lobbies/{id}/join", s.JoinLobby).Methods("POST")
	s.router.HandleFunc("/ws", s.handleWebSocket)
}

func (s *Server) GetLobby(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]

	lobby, exists := s.lobbies[id]
	if !exists {
		http.Error(w, "Lobby not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lobby)
}

func (s *Server) CreateLobby(w http.ResponseWriter, r *http.Request) {
	var lobbyResponce struct {
		ID string `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&lobbyRequest); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

	if _, exists := s.lobbies[lobbyRequest.ID]; exists {
        http.Error(w, "Lobby already exists", http.StatusConflict)
        return
    }

	lobby := lobby.NewLobby(lobbyRequest.ID)
    s.lobbies[lobbyRequest.ID] = lobby

	response := map[string]interface{}{
        "message": "Lobby created successfully",
        "lobbyId": lobby.ID,
    }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func main() {
	gameServer := server.NewServer()
	

	fmt.Println("starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
