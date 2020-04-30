package versions

import (
	"fmt"
	"runtime"
)

var (
	GitCommit    string = "$Format:%H$"
	BuildDate    string = "1970-01-01T00:00:00Z"
	GitBranch    string = "lose"
	GitTreeState string = "not a git tree"
	App          string = "none"
)

type Info struct {
	Major        string `json:"major"`
	Minor        string `json:"minor"`
	Gitcommit    string `json:"GitCommit"`
	BuildDate    string `json:"BuildDate"`
	GitBranch    string `json:"git_branch"`
	GitTreeState string `json:"git_tree_state"`
	GoVersion    string `json:"go_version"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func getVersionInfo() Info {
	return Info{
		Gitcommit:    GitCommit,
		BuildDate:    BuildDate,
		GitBranch:    GitBranch,
		GitTreeState: GitTreeState,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
