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
	po := model.Agent{
		Ip:               to.Ip,
		Port:             to.Port,
		Status:           to.Status,
		LastRegisterTime: time.Now(),
	}

	if po.ID == 0 {
		r.DB.Model(&model.Agent{}).Create(&po)
	} else {
		r.DB.Model(&model.Agent{}).
			Where("mac_address=?", to.MacAddress).
			Updates(&po)
	}

	return
}

func (r AgentRepo) GetById(id uint) (vm model.Agent) {
	r.DB.Model(&model.Agent{}).Where("ID=?", id).First(&vm)
	return
}
func (r AgentRepo) GetByMac(mac string) (vm model.Agent) {
	r.DB.Model(&model.Agent{}).Where("mac=?", mac).First(&vm)
	return
}

func (r AgentRepo) Save(po *model.Agent) {
	r.DB.Model(&model.Agent{}).Omit("").Create(po)
	return
}
