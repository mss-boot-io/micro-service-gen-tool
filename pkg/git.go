/*
 * @Author: lwnmengjing
 * @Date: 2021/12/16 9:07 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/12/16 9:07 下午
 */

package pkg

import (
	"context"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
	"log"
)

type GithubConfig struct {
	Name         string            `yaml:"name"`
	Organization string            `yaml:"organization"`
	Description  string            `yaml:"description"`
	Secrets      map[string]string `yaml:"secrets"`
	Token        string            `yaml:"token"`
}

// GitRemote from remote git
func GitRemote(url, directory string) error {
	r, err := git.PlainInit(directory, false)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{url},
	})
	if err != nil {
		log.Println(err)
		return err
	}
	err = r.CreateBranch(&config.Branch{
		Name: "main",
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GitClone clone git repo
func GitClone(url, directory string, noCheckout bool, accessToken string) error {
	auth := &http.BasicAuth{}
	if accessToken != "" {
		//fixme username not valid
		auth.Username = "username"
		auth.Password = accessToken
	}
	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:               url,
		NoCheckout:        noCheckout,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Auth:              auth,
	})
	return err
}

// CreateGithubRepo create github repo
func CreateGithubRepo(organization, name, description, token string, private bool) (*github.Repository, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	r := &github.Repository{Name: &name, Private: &private, Description: &description}
	repo, _, err := client.Repositories.Create(ctx, organization, r)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Printf("Successfully created new repo: %s\n", repo.GetName())
	return repo, nil
}

// AddActionSecretsGithubRepo add action secret
//func AddActionSecretsGithubRepo(organization, name, token string, data map[string]string) error {
//	ctx := context.Background()
//	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
//	tc := oauth2.NewClient(ctx, ts)
//	client := github.NewClient(tc)
//	var err error
//	for k, v := range data {
//		input := github.EncryptedSecret{
//			Name: k,
//			EncryptedValue: v,
//		}
//		_, err = client.Actions.CreateOrUpdateRepoSecret(ctx, organization, name, &input)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

// CommitAndPushGithubRepo commit and push github repo
func CommitAndPushGithubRepo(directory, accessToken string) error {
	r, err := git.PlainOpen(directory)
	if err != nil {
		log.Println(err)
		return err
	}
	w, err := r.Worktree()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = w.Add(".")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = w.Commit(":tada: init project", &git.CommitOptions{})
	if err != nil {
		log.Println(err)
		return err
	}
	return r.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: "123",
			Password: accessToken,
		},
	})
}
