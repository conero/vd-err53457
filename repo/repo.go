package repo

import (
	"gitee.com/conero/uymas/fs"
	"gitee.com/conero/uymas/logger"
	"net/url"
	"strings"
	"time"
)

var (
	Logger *logger.Logger
)

const (
	RepoSvn = "SVN"
	RepoGit = "GIT"
)

type RevisionInfo struct {
	Version  string
	Author   string
	Uuid     string
	Datetime time.Time
}

const (
	DiffModeAdd    = "A"
	DiffModeDel    = "D"
	DiffModeModify = "M"
)

type DiffPath struct {
	Filename string
	Mode     string // A,M,D
	IsDir    bool
}

type Repo interface {
	Latest() (RevisionInfo, error)
	Patch(string, string) ([]DiffPath, error)
	BaseUrl() string
}

func NewRepo(vUrl string, vType string) Repo {
	var tTy string
	vUrl, tTy = StdRepoUrl(vUrl)
	if vType == "" && tTy != "" {
		vType = tTy
	}
	vType = strings.ToUpper(vType)
	switch vType {
	case RepoSvn:
		return &Svn{vUrl: vUrl}
	}
	return nil
}

func StdRepoUrl(vUrl string) (rpUrl, vType string) {
	u, er := url.Parse(vUrl)
	if er == nil {
		if strings.Index(vUrl, ".git") > -1 {
			vType = RepoGit
		} else if u.Scheme == "svn" || u.Scheme == "svn+ssh" {
			vType = RepoSvn
		}
	} else {
		vUrl = fs.StdPathName(vUrl)
		if strings.Index(vUrl, ".git") > -1 {
			vType = RepoGit
		} else {
			if vUrl[len(vUrl)-1:] != "/" {
				vUrl += "/"
			}
		}
	}
	rpUrl = vUrl
	return
}

func getLogger() *logger.Logger {
	if Logger == nil {
		Logger = logger.NewLogger(logger.Config{})
	}
	return Logger
}
