# idea

a collaborative storywriting game where users take turns writing a sentence that must include an increasingly harder series of words in an increasingly short time limit. 

I am going to write it in Go potentially with the structs defined in protobuf so that I can easily add a web frontend later. To start it will be
played in cli.

## features

- CLI client
- multiplayer functionality
- player submitted words
- lobby system like jackbox
- shortening timer, last person standing

## tentative structure

word-game/
├── cmd/
│   └── server/
│   |   └── main.go
│   └── client/
│       └── main.go
├── internal/
│   ├── game/
│   │   ├── game.go
│   │   ├── player.go
│   │   ├── lobby.go
│   │   ├── word_bank.go
│   │   └── timer.go
│   ├── server/
│   │   ├── server.go
│   │   ├── handler.go
│   │   └── websocket.go
│   ├── client/
│   │   ├── client.go
│   │   └── ui.go│   
│   └── summary/
│       └── generator.go
├── pkg/
│   └── protocol/
│       └── messages.go
└── configs/
    └── game_config.yaml

## file roles

+ cmd/server/main.go

  - Entry point for the server application
  - Initializes and starts the game server

+ cmd/client/main.go

  - This is the entry point for the CLI client application.
  - It initializes the client, handles command-line arguments, and starts the client UI.

+ internal/game/game.go

  - Defines the Game struct and its methods
  - Manages game state, turns, word usage, and difficulty progression
  - Implements the core game logic for the sentence-writing phases


+ internal/game/player.go

  - Defines the Player struct and its methods
  - Manages player state, scoring, and turn actions


+ internal/game/lobby.go

  - Defines the Lobby struct and its methods
  - Handles player joining/leaving, game start conditions, and lobby lifecycle


+ internal/game/word_bank.go

  - Manages the word banks for different difficulty levels
  - Provides functions to get random words based on the current game phase


+ internal/game/timer.go

  - Implements the timing mechanism for turns
  - Handles the progressively shortening timer logic


+ internal/server/server.go

  - Defines the main Server struct
  - Manages multiple lobbies and games
  - Handles high-level server operations


+ internal/server/handler.go

  - Implements HTTP handlers for non-WebSocket endpoints
  - Handles initial game creation, joining, and other RESTful operations


+ internal/server/websocket.go

  - Manages WebSocket connections for real-time game communication
  - Handles message routing between clients and game logic

+ internal/client/client.go

  - Implements the core client logic.
  - Manages the connection to the server.
  - Handles sending and receiving messages.
  - Implements the client-side game state management.


+ internal/client/ui.go

  - Implements the CLI user interface.
  - Handles user input and output.
  - Renders the game state, lobby information, and prompts to the console.

+ internal/summary/generator.go

  - Generates the sharable game summary document
  - Formats game results and notable moments


+ pkg/protocol/messages.go

  - Defines the message structures for client-server communication
  - Includes message types for game events, player actions, etc.


+ configs/game_config.yaml

  - Configuration file for game settings
  - Includes parameters like max players, timer durations, word bank sizes, etc.


## development plan 

0. init - initalize go project and refresh on how some stuff works
0.5. create a server and client in the two main.go to test.
1. internal/game - design necessary structs
   + game.go
   + player.go
   + lobby.go
   + word_bank.go
   + timer.go
2. internal/server - implement necessary webserver stuff
3. cmd/ - flesh out the 2 main.go
4. internal/client - implement CLI interface.
5. internal/game - implement functions to operate on structs
