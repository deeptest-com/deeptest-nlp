package repo

import (
	"github.com/utlai/utl/internal/comm/domain"
	"github.com/utlai/utl/internal/server/model"
	"gorm.io/gorm"
	"time"
)

type AgentRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewAgentRepo() *AgentRepo {
	return &AgentRepo{}
}

func (r AgentRepo) Register(to domain.Agent) (err error) {
	agent := model.Agent{
		Ip:               to.Ip,
		Port:             to.Port,
		MacAddress:       to.MacAddress,
		Status:           to.Status,
		LastRegisterTime: time.Now(),
	}

	po := r.GetByMac(to.MacAddress)

	if po.ID == 0 {
		r.DB.Model(&model.Agent{}).Create(&agent)
	} else {
		r.DB.Model(&model.Agent{}).
			Where("mac_address=?", to.MacAddress).
			Updates(&agent)
	}

	return
}

func (r AgentRepo) Query() (agents []model.Agent) {
	r.DB.Model(&model.Agent{}).Where("NOT deletedAND NOT disabled").Find(&agents)
	return
}

func (r AgentRepo) GetById(id uint) (agent model.Agent) {
	r.DB.Model(&model.Agent{}).Where("ID=?", id).First(&agent)
	return
}
func (r AgentRepo) GetByMac(mac string) (agent model.Agent) {
	r.DB.Model(&model.Agent{}).Where("mac_address=? AND NOT deleted", mac).First(&agent)
	return
}

func (r AgentRepo) Save(po *model.Agent) {
	r.DB.Model(&model.Agent{}).Omit("").Create(po)
	return
}
