package comm

import (
	"github.com/utlai/utl/internal/comm/domain"
)

type InstructionService struct {
}

func NewInstruction() *InstructionService {
	return &InstructionService{}
}

func (s *InstructionService) Parer(instruction domain.RasaResp) (result map[string]interface{}) {
	for _, e := range instruction.Entities {
		result[e.Entity] = e.Value
	}

	return
}
