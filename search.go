package tonic

// search represents search operations run on a db
type search struct {
	db              *DB
	whereConditions []map[string]interface{}
}

func (s *search) clone() *search {
	clone := *s
	return &clone
}
