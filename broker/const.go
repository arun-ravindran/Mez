package broker

// private type for Context keys
type contextKey int

const (
	clientIDKey contextKey = iota

	//Polling time for subscribe (edgenode and edgeserver)
	maxPollTime = 1100000
)
