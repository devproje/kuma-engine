package mode

// EngineMode the mode of the engine
type EngineMode int

const (
	// Release mode: general production mode
	Release EngineMode = iota
	// Debug mode: log debug only (DO NOT USE IN PRODUCTION)
	Debug
	// Remove mode: bot will remove all commands
	Remove
)
