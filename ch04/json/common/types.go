package common

// work group
type WorkGroup struct {
	Name    string
	Members []Member
}

// group member
type Member struct {
	Name     string
	Position string
}
