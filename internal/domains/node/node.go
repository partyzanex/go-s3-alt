package node

import "time"

type Node struct {
	ID        ID
	Online    bool
	Timestamp time.Time
	Address   string
}
