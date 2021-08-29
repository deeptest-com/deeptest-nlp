package service

import (
	"github.com/kataras/iris/v12/websocket"
	consts "github.com/utlai/utl/internal/comm/const"
	_fileUtils "github.com/utlai/utl/internal/pkg/libs/file"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
	serverVari "github.com/utlai/utl/internal/server/utils/var"
	"gopkg.in/yaml.v3"
	"path/filepath"
)

type NluPatternService struct {
	ProjectRepo       *repo.ProjectRepo  `inject:""`
	NluServiceService *NluServiceService `inject:""`
	WebSocketService  *WebSocketService  `inject:""`
	Namespace         string
	*websocket.NSConn `stateless:"true"`
}

func NewNluPatternService() *NluPatternService {
	return &NluPatternService{Namespace: serverConst.WsNamespace}
}

func (s *NluPatternService) Reload(projectId uint) (project model.Project) {
	project = s.ProjectRepo.GetDetail(projectId)

	dir := filepath.Join(project.Path, consts.Pattern.ToString())
	files, _ := _fileUtils.ListDir(dir)

	for _, f := range files {
		content := _fileUtils.ReadFileBuf(f)

		task := domain.NluTaskForPattern{}
		yaml.Unmarshal(content, &task)

		serverVari.PatternData[projectId] = append(serverVari.PatternData[projectId], task)
	}

	return
}
