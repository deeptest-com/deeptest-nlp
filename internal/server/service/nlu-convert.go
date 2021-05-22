package service

import (
	_fileUtils "github.com/utlai/utl/internal/pkg/libs/file"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/repo"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
	"gopkg.in/yaml.v3"
	"path/filepath"
)

type NluConvertService struct {
	NluSynonymRepo     *repo.NluSynonymRepo     `inject:""`
	NluSynonymItemRepo *repo.NluSynonymItemRepo `inject:""`
	NluLookupRepo      *repo.NluLookupRepo      `inject:""`
	NluLookupItemRepo  *repo.NluLookupItemRepo  `inject:""`
	NluRegexRepo       *repo.NluRegexRepo       `inject:""`

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
	_fileUtils.RmDir(filepath.Join(projectDir, "synonym"))

	synonyms := s.NluSynonymRepo.ListByProjectId(projectId)
	for _, synonym := range synonyms {
		nluSynonym := domain.NluSynonym{Version: serverConst.NluVersion}
		synonymDef := domain.NluSynonymDef{Synonym: synonym.Name}

		synonymItems := s.NluSynonymItemRepo.ListBySynonymId(projectId)
		for _, item := range synonymItems {
			synonymDef.Examples += item.Content + "\n"
		}
		nluSynonym.SynonymDef = append(nluSynonym.SynonymDef, synonymDef)

		filePath := filepath.Join(projectDir, "synonym", synonym.Name+".yml")
		bytes, _ := yaml.Marshal(&nluSynonym)
		_fileUtils.WriteFile(string(bytes), filePath)
	}

	return
}
func (s *NluConvertService) ConvertLookup(projectId uint, projectDir string, nluDomain *domain.NluDomain) {
	_fileUtils.RmDir(filepath.Join(projectDir, "lookup"))

	lookups := s.NluLookupRepo.ListByProjectId(projectId)
	for _, lookup := range lookups {
		nluLookup := domain.NluLookup{Version: serverConst.NluVersion}
		lookupDef := domain.NluLookupDef{Lookup: lookup.Name}

		lookupItems := s.NluSynonymItemRepo.ListBySynonymId(projectId)
		for _, item := range lookupItems {
			lookupDef.Examples += item.Content + "\n"
		}
		nluLookup.LookupDef = append(nluLookup.LookupDef, lookupDef)

		filePath := filepath.Join(projectDir, "lookup", lookup.Name+".yml")
		bytes, _ := yaml.Marshal(&nluLookup)
		_fileUtils.WriteFile(string(bytes), filePath)
	}

	return
}
func (s *NluConvertService) ConvertRegex(projectId uint, projectDir string, nluDomain *domain.NluDomain) {
	_fileUtils.RmDir(filepath.Join(projectDir, "regex"))

	regexes := s.NluRegexRepo.ListByProjectId(projectId)
	for _, regex := range regexes {
		nluRegex := domain.NluRegex{Version: serverConst.NluVersion}

		regexDef := domain.NluRegexDef{Regex: regex.Name}

		lookupItems := s.NluSynonymItemRepo.ListBySynonymId(projectId)
		for _, item := range lookupItems {
			regexDef.Examples += item.Content + "\n"
		}
		nluRegex.RegexDef = append(nluRegex.RegexDef, regexDef)

		filePath := filepath.Join(projectDir, "regex", regex.Name+".yml")
		bytes, _ := yaml.Marshal(&nluRegex)
		_fileUtils.WriteFile(string(bytes), filePath)
	}

	return
}

func (s *NluConvertService) ConvertIntent(projectId uint, projectDir string) (files []string) {
	_fileUtils.RmDir(filepath.Join(projectDir, "intent"))

	tasks := s.NluTaskRepo.ListByProjectId(projectId)
	for _, task := range tasks {
		intents := s.NluIntentRepo.ListByTaskId(task.ID)

		for _, intent := range intents {
			nluIntent := domain.NluIntent{}

			sents := s.NluSentRepo.ListByIntentId(intent.ID)

			for _, sent := range sents {
				_logUtils.Info(sent.Text)
			}

			intentFilePath := filepath.Join(projectDir, "intent", intent.Name+".yml")
			bytes, _ := yaml.Marshal(&nluIntent)
			_fileUtils.WriteFile(intentFilePath, string(bytes))
		}
	}

	return
}
