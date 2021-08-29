package service

import (
	"fmt"
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
	"strings"
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

func (s *NluPatternService) Reload(id uint) (project model.Project) {
	if serverVari.PatternData[id] == nil {
		serverVari.PatternData[id] = map[string][]string{}
	}

	project = s.ProjectRepo.GetDetail(id)

	dir := filepath.Join(project.Path, consts.Pattern.ToString())
	files, _ := _fileUtils.ListDir(dir)

	for _, f := range files {
		content := _fileUtils.ReadFileBuf(f)

		task := domain.NluPatternTask{}
		yaml.Unmarshal(content, &task)

		for _, pattern := range task.Intents {
			for _, line := range strings.Split(pattern.Examples, "\n") {
				key := fmt.Sprintf("%d-%s", pattern.Id, pattern.Name)
				arr := strings.Split(line, "-")

				if len(arr) == 1 {
					continue
				}

				value := strings.TrimSpace(arr[1])
				serverVari.PatternData[id][key] = append(serverVari.PatternData[id][key], value)
			}
		}
	}

	return
}
