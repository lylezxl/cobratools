package versions

import (
	"fmt"
	"runtime"
)

var (
	gitCommit    string = "$Format:%H$"
	buildDate    string = "1970-01-01T00:00:00Z"
	gitVerison   string = "lose"
	gitTreeState string = "not a git tree"
	project       string = "none"
)

type Info struct {
	Major        string `json:"major"`
	Minor 		string `json:"minor"`
	Gitcommit    string `json:"GitCommit"`
	BuildDate    string `json:"BuildDate"`
	GitVerison   string `json:"GitVerison"`
	GitTreeState string `json:"git_tree_state"`
	GoVersion    string `json:"go_version"`
    Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func getVersionInfo() Info {
	return Info{
		Gitcommit:    gitCommit,
		BuildDate:    buildDate,
		GitVerison:   gitVerison,
		GitTreeState: gitTreeState,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}




