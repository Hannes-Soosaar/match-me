# This is the match me project

## Required

- **Docker**: Required for running the PostgreSQL database in a container. Tested with version 27.1.2. 
(Docker Compose is included with Docker, so no separate installation is needed).
- **Make**: For running predefined commands like `make start`, `make stop-db`, etc.
- **Node.js and NPM**: Required for the frontend. Tested with node version 23.1.0 & NPM version 10.9.0.
- **Go**: Required for the backend. Tested with version 1.23.3.
- **Concurrently**: A development dependency for running multiple commands (frontend and backend) simultaneously.


## Setting up the project

1. Clone the repository: 
```git clone https://gitea.kood.tech/karl-hendrikkahn/match-me.git```.
2. Install backend dependencies (for Go):
```cd backend```
```go mod tidy```
3. Install frontend dependencies:
```cd frontend```
```npm install```
4. Start the database, backend, and frontend servers all at the same time by running ```make start``` in the project root directory.
5. Close the backend and frontend servers by pressing ```Ctrl+C``` in the terminal.
6. Close the database server with ```make stop-db```.

#### Other helpful commands
Check if the docker container is running with ```docker ps```.
Stop the database server by running ```make stop-db``` in the project root directory.
Reset the database with ```make reset-db``` in the project root directory.