.PHONY: start

start:
# Start frontend and backend at the same time
	./frontend/node_modules/.bin/concurrently --kill-others-on-fail \
		"make start-db"\
		"sh -c 'sleep 3 && cd backend && go run main.go'"\
		"cd frontend && npm start" 

make-executable:
	chmod +x reset_db.sh

# Resets the database, meaning it deletes all the data and initialises it again with the default data from init.sql
reset-db: make-executable
	./reset_db.sh

# Starts the database
start-db:
	docker-compose up -d

# Stops the database
stop-db:
	docker-compose down