package service

import (
	_fileUtils "github.com/utlai/utl/internal/pkg/libs/file"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/repo"
	"gopkg.in/yaml.v3"
	"path/filepath"
)

type NluConvertService struct {
	NluSynonymRepo *repo.NluSynonymRepo `inject:""`
	NluLookupRepo  *repo.NluLookupRepo  `inject:""`
	NluRegexRepo   *repo.NluRegexRepo   `inject:""`

	ProjectRepo   *repo.ProjectRepo   `inject:""`
	NluTaskRepo   *repo.NluTaskRepo   `inject:""`
	NluIntentRepo *repo.NluIntentRepo `inject:""`
	NluSentRepo   *repo.NluSentRepo   `inject:""`
}

func NewNluConvertService() *NluConvertService {
	return &NluConvertService{}
}

func (s *NluConvertService) ConvertProject(id uint) (files []string) {
	nluDomain := domain.NluDomain{}

	project := s.ProjectRepo.Get(id)
	projectDir := project.Path

	s.ConvertIntent(id, projectDir)
	s.ConvertSynonym(id, projectDir, &nluDomain)
	s.ConvertLookup(id, projectDir, &nluDomain)
	s.ConvertRegex(id, projectDir, &nluDomain)

	domainFilePath := filepath.Join(projectDir, "domain.yml")
	bytes, _ := yaml.Marshal(&nluDomain)
	_fileUtils.WriteFile(string(bytes), domainFilePath)

	return
}

func (s *NluConvertService) ConvertSynonym(projectId uint, projectDir string, nluDomain *domain.NluDomain) {
	synonyms := s.NluSynonymRepo.ListByProjectId(projectId)

	for _, item := range synonyms {
		// convert po to domain

		filePath := filepath.Join(projectDir, "synonym", item.Name)
		bytes, _ := yaml.Marshal(&item)
		_fileUtils.WriteFile(string(bytes), filePath)
	}

	return
}
func (s *NluConvertService) ConvertLookup(projectId uint, projectDir string, nluDomain *domain.NluDomain) {
	lookups := s.NluLookupRepo.ListByProjectId(projectId)

	for _, item := range lookups {
		// convert po to domain

		filePath := filepath.Join(projectDir, "lookup", item.Name)
		bytes, _ := yaml.Marshal(&item)
		_fileUtils.WriteFile(string(bytes), filePath)
	}

	return
}
func (s *NluConvertService) ConvertRegex(projectId uint, projectDir string, nluDomain *domain.NluDomain) {
	regexes := s.NluRegexRepo.ListByProjectId(projectId)

	for _, item := range regexes {
		// convert po to domain

		filePath := filepath.Join(projectDir, "regex", item.Name)
		bytes, _ := yaml.Marshal(&item)
		_fileUtils.WriteFile(string(bytes), filePath)
	}

	return
}

func (s *NluConvertService) ConvertIntent(projectId uint, projectDir string) (files []string) {
	tasks := s.NluTaskRepo.ListByProjectId(projectId)
	for _, task := range tasks {
		intents := s.NluIntentRepo.ListByTaskId(task.ID)

		for _, intent := range intents {
			nluIntent := domain.NluIntent{}

			sents := s.NluSentRepo.ListByIntentId(intent.ID)

			for _, sent := range sents {
				_logUtils.Info(sent.Text)
			}

			intentFilePath := filepath.Join(projectDir, "intent", "")
			bytes, _ := yaml.Marshal(&nluIntent)
			_fileUtils.WriteFile(intentFilePath, string(bytes))
		}
	}
}
