package http

type Scope string
type Scopes []Scope

const (
	// AdminScope - Allows full access to API resources
	WorkerScope Scope = "Worker"

	// GuestScope - Allows access to guests resources
	GuestScope Scope = "Guest"
)

func (e Scope) String() string {
	switch e {
	case WorkerScope:
		return "Worker"
	case GuestScope:
		return "Guest"
	default:
		return "error"
	}
}
