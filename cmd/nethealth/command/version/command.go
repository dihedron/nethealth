package version

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/dihedron/nethealth/cmd/nethealth/build"
	"github.com/dihedron/nethealth/cmd/nethealth/command/base"
)

// Version is the command that prints information about the application
// or plugin to the console; it support both compact and verbose mode.
type Version struct {
	base.Command
	// Verbose prints extensive information about this application or plugin.
	Verbose bool `short:"v" long:"verbose" description:"Print extensive information about this application."`
}

type ShortInfo struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Copyright   string `json:"copyright,omitempty"`
	Version     string `json:"version,omitempty"`
}

type DetailedInfo struct {
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Copyright   string       `json:"copyright,omitempty"`
	Version     VersionInfo  `json:"version,omitempty"`
	Compiler    CompilerInfo `json:"compiler,omitempty"`
	Build       BuildInfo    `json:"build,omitempty"`
	Git         GitInfo      `json:"git,omitempty"`
}

type VersionInfo struct {
	Major string `json:"major,omitempty"`
	Minor string `json:"minor,omitempty"`
	Patch string `json:"patch,omitempty"`
}

func (v *VersionInfo) String() string {
	return fmt.Sprintf("%s.%s.%s", v.Major, v.Minor, v.Patch)
}

type BuildInfo struct {
	Time string `json:"time,omitempty"`
	Date string `json:"date,omitempty"`
}

type CompilerInfo struct {
	Version      string `json:"version,omitempty"`
	OS           string `json:"os,omitempty"`
	Architecture string `json:"arch,omitempty"`
}

type GitInfo struct {
	Tag    string `json:"tag,omitempty"`
	Commit string `json:"commit,omitempty"`
	Hash   string `json:"hash,omitempty"`
}

// Execute is the real implementation of the Version command.
func (cmd *Version) Execute(args []string) error {
	logger := cmd.InitLogger(true)
	logger.Debugf("running version command (BuildName: %s)", build.Name)
	if cmd.AutomationFriendly {
		var info interface{}
		if !cmd.Verbose {
			// short
			info = &ShortInfo{
				Name:        build.Name,
				Description: build.Description,
				Copyright:   build.Copyright,
				Version:     fmt.Sprintf("v%s.%s.%s", build.VersionMajor, build.VersionMinor, build.VersionPatch),
			}
		} else {
			// verbose
			info = &DetailedInfo{
				Name:        build.Name,
				Description: build.Description,
				Copyright:   build.Copyright,
				Version: VersionInfo{
					Major: build.VersionMajor,
					Minor: build.VersionMinor,
					Patch: build.VersionPatch,
				},
				Build: BuildInfo{
					Time: build.BuildTime,
					Date: build.BuildDate,
				},
				Compiler: CompilerInfo{
					Version:      build.GoVersion,
					OS:           build.GoOS,
					Architecture: build.GoArch,
				},
				Git: GitInfo{
					Tag:    build.GitTag,
					Hash:   build.GitHash,
					Commit: build.GitCommit,
				},
			}
		}
		data, err := json.Marshal(info)
		if err != nil {
			logger.Errorf("error marshalling plugin info to JSON: %v", err)
			return err
		}
		logger.Debugf("marshalling data to JSON: %s", string(data))
		fmt.Println(string(data))
	} else {
		if !cmd.Verbose {
			fmt.Printf("\n  %s %s - %s - %s\n\n", path.Base(os.Args[0]), build.GitTag, build.Copyright, build.Description)
		} else {
			fmt.Printf("\n  %s %s - %s - %s\n\n", path.Base(os.Args[0]), build.GitTag, build.Copyright, build.Description)
			fmt.Printf("  - Name             : %s\n", build.Name)
			fmt.Printf("  - Description      : %s\n", build.Description)
			fmt.Printf("  - Copyright        : %s\n", build.Copyright)
			fmt.Printf("  - Major Version    : %s\n", build.VersionMajor)
			fmt.Printf("  - Minor Version    : %s\n", build.VersionMinor)
			fmt.Printf("  - Patch Version    : %s\n", build.VersionPatch)
			fmt.Printf("  - Built on         : %s\n", build.BuildDate)
			fmt.Printf("  - Built at         : %s\n", build.BuildTime)
			fmt.Printf("  - Compiler         : %s\n", build.GoVersion)
			fmt.Printf("  - Operating System : %s\n", build.GoOS)
			fmt.Printf("  - Architecture     : %s\n", build.GoArch)
			fmt.Printf("  - Git Tag          : %s\n", build.GitTag)
			fmt.Printf("  - Build Hash       : %s\n", build.GitHash)
			fmt.Printf("  - Build Commit     : %s\n", build.GitCommit)
		}
	}
	logger.Debugf("command done")
	return nil
}
