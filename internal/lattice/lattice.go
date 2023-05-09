package lattice

import "go-rla/internal/models"

type RLA interface {
	Propose(data models.Data) error
}
