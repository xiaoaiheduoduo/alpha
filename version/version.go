package version

import (
	"fmt"
	"runtime"
)

type Info struct {
	GitVersion string `json:"git_version"`
	GitCommit  string `json:"git_commit"`
	BuildDate  string `json:"build_date"`
	GoVersion  string `json:"go_version"`
	Compiler   string `json:"compiler"`
	Platform   string `json:"platform"`
}

func (info Info) String() string {
	return info.GitVersion
}

var info *Info

func init() {
	info = &Info{
		GitVersion: gitVersion,
		GitCommit:  gitCommit,
		BuildDate:  buildDate,
		GoVersion:  runtime.Version(),
		Compiler:   runtime.Compiler,
		Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

func Get() *Info {
	return info
}
