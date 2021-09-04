package serverService

import (
	"github.com/utlai/utl/internal/server/repo"
	"gopkg.in/yaml.v2"
)

type NluCompileService struct {
	NluSynonymRepo     *repo.NluSynonymRepo     `inject:""`
	NluSynonymItemRepo *repo.NluSynonymItemRepo `inject:""`
	NluLookupRepo      *repo.NluLookupRepo      `inject:""`
	NluLookupItemRepo  *repo.NluLookupItemRepo  `inject:""`
	NluRegexRepo       *repo.NluRegexRepo       `inject:""`
	NluRegexItemRepo   *repo.NluRegexItemRepo   `inject:""`

	ProjectRepo   *repo.ProjectRepo   `inject:""`
	NluTaskRepo   *repo.NluTaskRepo   `inject:""`
	NluIntentRepo *repo.NluIntentRepo `inject:""`
	NluSentRepo   *repo.NluSentRepo   `inject:""`
	NluSlotRepo   *repo.NluSlotRepo   `inject:""`

	NluCompileRasaService    *NluCompileRasaService    `inject:""`
	NluCompilePatternService *NluCompilePatternService `inject:""`
}

func NewNluCompileService() *NluCompileService {
	return &NluCompileService{}
}

func (s *NluCompileService) CompileProject(id uint) {
	//if serverConf.Config.Analyzer == consts.Rasa {
	s.NluCompileRasaService.RasaCompile(id)
	//} else if serverConf.Config.Analyzer == consts.Pattern {
	s.NluCompilePatternService.PatternCompile(id)
	//}

	return
}

func (s *NluCompileService) ChangeArrToFlow(obj interface{}) (ret string) {
	bytes, _ := yaml.Marshal(&obj)
	m := yaml.MapSlice{}
	yaml.Unmarshal(bytes, &m)

	d, _ := yaml.Marshal(&m)
	ret = string(d)
	return
}
