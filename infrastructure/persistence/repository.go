package persistence

import (
	"collegare/domain"
	"collegare/domain/models"

	"github.com/gofrs/uuid/v5"
)

type MockCallFlowRepository struct{}

func NewMockCallFlowRepository() *MockCallFlowRepository {
	return &MockCallFlowRepository{}
}

func (r *MockCallFlowRepository) FindAll() []models.CallFlow {
	panic("Not Implemented")
}

func (r *MockCallFlowRepository) Find(id uuid.UUID) models.CallFlow {
	return models.CallFlow{
		Id:   id,
		Name: "Dummy flow",
		Script: models.Script{
			Start: 1,
			Nodes: map[int]models.Node{
				0: {
					Id:   1,
					Loop: true,
					Name: "",
					Track: models.Track{
						Id:       uuid.Max,
						Filename: "wilhelm.opus",
					},
					Options: map[string]int{
						domain.One:   1,
						domain.Two:   2,
						domain.Three: 3,
						domain.Four:  4,
						domain.Zero:  0,
					},
				},
				1: {
					Id:   1,
					Loop: true,
					Name: "",
					Track: models.Track{
						Id:       uuid.Max,
						Filename: "c.opus",
					},
					Options: map[string]int{
						domain.One:   1,
						domain.Two:   2,
						domain.Three: 3,
						domain.Four:  4,
						domain.Zero:  0,
					},
				},
				2: {
					Id:   1,
					Loop: true,
					Name: "",
					Track: models.Track{
						Id:       uuid.Max,
						Filename: "aeolian.opus",
					},
					Options: map[string]int{
						domain.One:   1,
						domain.Two:   2,
						domain.Three: 3,
						domain.Four:  4,
						domain.Zero:  0,
					},
				},
				3: {
					Id:   1,
					Loop: true,
					Name: "",
					Track: models.Track{
						Id:       uuid.Max,
						Filename: "ionian.opus",
					},
					Options: map[string]int{
						domain.One:   1,
						domain.Two:   2,
						domain.Three: 3,
						domain.Four:  4,
						domain.Zero:  0,
					},
				},
				4: {
					Id:   1,
					Loop: true,
					Name: "",
					Track: models.Track{
						Id:       uuid.Max,
						Filename: "mixolydian.opus",
					},
					Options: map[string]int{
						domain.One:   1,
						domain.Two:   2,
						domain.Three: 3,
						domain.Four:  4,
						domain.Zero:  0,
					},
				},
			},
		},
	}
}
