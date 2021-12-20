/*
 * @Author: lwnmengjing
 * @Date: 2021/12/16 7:39 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/12/16 7:39 下午
 */

package pkg

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type TemplateConfig struct {
	Service     string        `yaml:"service"`
	TemplateUrl string        `yaml:"templateUrl"`
	CreateRepo  bool          `yaml:"createRepo"`
	Destination string        `yaml:"destination"`
	Github      *GithubConfig `yaml:"github"`
	Data        interface{}   `yaml:"data"`
	Ignore      []string      `yaml:"ignore"`
}

// Generator generate operator
type Generator struct {
	TemplatePath    string
	DestinationPath string
	Cfg             interface{}
	GithubConfig    *GithubConfig
	accessToken     string
	Ignore          []string
}

// Generate example
//func Generate(url, destinationPath string, cfg interface{}, githubConfig *GithubConfig, accessToken string) error {
//	templatePath := filepath.Base(url)
func Generate(c *TemplateConfig) (err error) {
	templatePath := filepath.Base(c.TemplateUrl)
	var accessToken string
	if c.Github != nil {
		accessToken = c.Github.Token
	}
	err = GitClone(c.TemplateUrl, templatePath, false, accessToken)
	if err != nil {
		log.Println(err)
		return err
	}
	if !c.CreateRepo {
		c.Github = nil
	}
	defer os.RemoveAll(templatePath)
	//delete destinationPath
	_ = os.RemoveAll(c.Destination)
	_ = os.RemoveAll(filepath.Join(templatePath, ".git"))
	t := &Generator{
		TemplatePath:    templatePath,
		DestinationPath: c.Destination,
		Cfg:             c.Data,
		GithubConfig:    c.Github,
		accessToken:     accessToken,
		Ignore:          c.Ignore,
	}
	if err = t.Traverse(); err != nil {
		log.Println(err)
		return err
	}
	if err = t.CreateGithubRepo(); err != nil {
		log.Println(err)
		return err
	}
	if err = t.CommitGithubRepo(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Traverse traverse all dir
func (e *Generator) Traverse() error {
	return filepath.WalkDir(e.TemplatePath, e.TraverseFunc)
}

// TraverseFunc traverse callback
func (e *Generator) TraverseFunc(path string, f os.DirEntry, err error) error {
	templatePath := path
	t := template.New(path)
	t = template.Must(t.Parse(path))
	var buffer bytes.Buffer
	if err = t.Execute(&buffer, e.Cfg); err != nil {
		log.Println(err)
		return err
	}
	path = strings.ReplaceAll(buffer.String(), e.TemplatePath, e.DestinationPath)

	if f.IsDir() {
		// dir
		if !PathExist(path) {
			return PathCreate(path)
		}
		return nil
	}
	if len(e.Ignore) > 0 {
		for i := range e.Ignore {
			if e.Ignore[i] == templatePath || e.Ignore[i] == filepath.Base(templatePath) {
				//find file, then copy
				_, err = FileCopy(templatePath, path)
				if err != nil {
					log.Println(err)
				}
				return err
			}
		}
	}
	var rb []byte
	if rb, err = ioutil.ReadFile(templatePath); err != nil {
		log.Println(err)
		return err
	}
	buffer = bytes.Buffer{}
	t = template.New(path + "[file]")
	t = template.Must(t.Parse(string(rb)))
	if err = t.Execute(&buffer, e.Cfg); err != nil {
		log.Println(err)
		return err
	}
	// create file
	err = FileCreate(buffer, path)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// CreateGithubRepo create github repo
func (e *Generator) CreateGithubRepo() error {
	if e.GithubConfig == nil {
		return nil
	}
	repo, err := CreateGithubRepo(
		e.GithubConfig.Organization,
		e.GithubConfig.Name,
		e.GithubConfig.Description,
		e.accessToken, true)
	if err != nil {
		log.Println(err)
		return err
	}
	//owner := e.GithubConfig.Organization
	//if repo.Organization == nil || *repo.Organization.Name == "" {
	//	owner = *repo.Owner.Login
	//}
	//err = AddActionSecretsGithubRepo(
	//	owner,
	//	e.GithubConfig.Name,
	//	e.accessToken, e.GithubConfig.Secrets)
	//if err != nil {
	//	log.Println(err)
	//	return err
	//}
	err = PathCreate(e.DestinationPath)
	if err != nil {
		log.Println(err)
		return err
	}
	err = GitRemote(repo.GetCloneURL(), e.DestinationPath)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// CommitGithubRepo commit repo to github
func (e *Generator) CommitGithubRepo() (err error) {
	if e.GithubConfig == nil {
		return nil
	}
	return CommitAndPushGithubRepo(e.DestinationPath, e.accessToken)
}
