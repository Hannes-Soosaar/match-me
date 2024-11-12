.PHONY: start

start:
# Start frontend and backend at the same time
	./frontend/node_modules/.bin/concurrently \
		"cd frontend && npm start" \
		"cd backend && go run main.go"