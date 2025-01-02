package db

// THE CONS SHOULD BE TAKEN FROM THE TABLE "categories" Might no be needed Unless we would get an option of adding Categories
const (
	// The following are the possible values for the "category" field in the "categories" table
	GENRE         = 1
	PLAY_STYLE    = 2
	PLATFORM      = 3
	COMMUNICATION = 4
	GOALS         = 5
	SESSION       = 6
	VIBE          = 7
	LANGUAGE      = 8
	// The following are the possible values for the "status" field in the "user_match" table 
	NEW = "new" // this is used if the user has not seen the connection yet
	SHOWN = "shown" 	// this is used if the user has seen the connection but has not made a decision
	REMOVED = "removed"  // this is used if the user want to not see the connection, but it the status can be changed to new after a certain time period and if there is a lack of new connections
	
	REQUESTED = "requested"	// this is used if the user has requested the connection but the other party has not accepted yet
	CONNECTED = "connected"	// this is used if the user has accepted the connection and is now connected
	BLOCKED = "blocked"	// this is used if the user has blocked the connection all communication is blocked
	DELETED = "deleted"	// this is used if the user has deleted the connection but there is still a possibility of the other party to reach out via message ect. 
	
	DEMO_USER_COUNT = 10

)
