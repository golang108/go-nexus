package nexusrm

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const restListRepositories = "service/rest/v1/repositories"

// RepositoryType enumerates the types of repositories in RM
/*
type RepositoryType int

const (
	proxy = iota
	hosted
	group
)

func (r RepositoryType) String() string {
	switch r {
	case proxy:
		return "proxy"
	case hosted:
		return "hosted"
	case group:
		return "group"
	default:
		return ""
	}
}
*/

// Repository collects the information returned by RM about a repository
type Repository struct {
	Name       string `json:"name"`
	Format     string `json:"format"`
	Type       string `json:"type"`
	URL        string `json:"url"`
	Attributes struct {
		Proxy struct {
			RemoteURL string `json:"remoteUrl"`
		} `json:"proxy"`
	} `json:"attributes,omitempty"`
}

// Equals compares two Repository objects
func (a *Repository) Equals(b *Repository) (_ bool) {
	if a == b {
		return true
	}

	if a.Name != b.Name {
		return
	}

	if a.Format != b.Format {
		return
	}

	if a.Type != b.Type {
		return
	}

	if a.URL != b.URL {
		return
	}

	if a.Attributes.Proxy.RemoteURL != b.Attributes.Proxy.RemoteURL {
		return
	}

	return true
}

// GetRepositories returns a list of components in the indicated repository
func GetRepositories(rm RM) ([]Repository, error) {
	doError := func(err error) error {
		return fmt.Errorf("could not find repositories: %v", err)
	}

	body, resp, err := rm.Get(restListRepositories)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, doError(err)
	}

	repos := make([]Repository, 0)
	if err := json.Unmarshal(body, &repos); err != nil {
		return nil, doError(err)
	}

	return repos, nil
}
