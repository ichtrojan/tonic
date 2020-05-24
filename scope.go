package tonic

// Scope contain current operation's info on the database
type Scope struct {
	db     *DB
	Search *search
}
