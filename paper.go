package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
)

const PAPER_ENDPOINT string = "https://api.papermc.io/v2"

type Projects struct {
	Projects []string `json:"projects"`
}

type ProjectInfo struct {
	Project_Id     string   `json:"project_id"`
	Project_name   string   `json:"project_name"`
	Version_groups []string `json:"version_groups"`
	Versions       []string `json:"versions"`
}

type VersionInfo struct {
	Project_id   string        `json:"project_id"`
	Project_name string        `json:"project_name"`
	Version      string        `json:"version"`
	Builds       []json.Number `json:"builds"`
}

type Download struct {
	Name   string `json:"name"`
	Sha256 string `json:"sha256"`
}

type Downloads struct {
	Application     Download `json:"application"`
	Mojang_mappings Download `json:"mojang_mappings"`
}

type BuildInfo struct {
	Project_id   string      `json:"project_id"`
	Project_name string      `json:"project_name"`
	Version      string      `json:"version"`
	Build        json.Number `json:"build"`
	Time         string      `json:"time"`
	Channel      string      `json:"channel"`
	Promoted     bool        `json:"promoted"`
	Changes      []struct{}  `json:"changes"`
	Downloads    Downloads   `json:"downloads"`
}

func getProjects() Projects {
	resp, _ := get(PAPER_ENDPOINT + "/projects")
	bodyObj, _ := io.ReadAll(resp.Body)
	bodyStr := string(bodyObj)
	body := Projects{}

	err := json.Unmarshal([]byte(bodyStr), &body)

	if err != nil {
		log.Println(err.Error())
		log.Fatalln("getProjects() json unmarshal failed.")
	}

	return body
}

func getProjectInfo(project string) ProjectInfo {
	resp, _ := get(PAPER_ENDPOINT + "/projects/" + project)
	bodyObj, _ := io.ReadAll(resp.Body)
	bodyStr := string(bodyObj)
	body := ProjectInfo{}

	err := json.Unmarshal([]byte(bodyStr), &body)

	if err != nil {
		log.Println(err.Error())
		log.Fatalln("getProjectInfo() json unmarshal failed.")
	}

	return body
}

func getVersionInfo(project string, version string) VersionInfo {
	resp, _ := get(PAPER_ENDPOINT + "/projects/" + project + "/versions/" + version)
	bodyObj, _ := io.ReadAll(resp.Body)
	bodyStr := string(bodyObj)
	body := VersionInfo{}

	err := json.Unmarshal([]byte(bodyStr), &body)

	if err != nil {
		log.Println(err.Error())
		log.Fatalln("getVersionInfo() json unmarshal failed.")
	}

	return body
}

func getBuildInfo(project string, version string, build string) BuildInfo {
	resp, _ := get(PAPER_ENDPOINT + "/projects/" + project + "/versions/" + version + "/builds/" + build)
	bodyObj, _ := io.ReadAll(resp.Body)
	bodyStr := string(bodyObj)
	body := BuildInfo{}

	err := json.Unmarshal([]byte(bodyStr), &body)

	if err != nil {
		log.Println(err.Error())
		log.Fatalln("getBuildInfo() json unmarshal failed.")
	}

	return body
}

func downloadBuild(project string, version string, build string, download string) {
	resp, _ := get(PAPER_ENDPOINT + "/projects/" + project + "/versions/" + version + "/builds/" + build + "/downloads/" + download)
	out, err := os.Create(download)

	if err != nil {
		log.Println(err.Error())
		log.Fatalln("downloadBuild() File create failed.")
	}

	io.Copy(out, resp.Body)
}

func paperSearch(args []string) {

	if len(args) > 0 {
		switch args[0] {
		case "paper", "travertine", "waterfall", "velocity", "folia":
			info := getProjectInfo(args[0])

			if len(args) == 1 {
				for _, version := range info.Versions {
					if version == info.Versions[len(info.Versions)-1] {
						fmt.Println("\x1b[32m" + info.Project_name + "/" + version + " -- Latest\x1b[0m")
					} else {
						fmt.Println(info.Project_name + "/" + version)
					}
				}
			} else if args[1] == "latest" {
				versionInfo := getVersionInfo(info.Project_Id, info.Versions[len(info.Versions)-1])
				var builds []string

				for _, i := range versionInfo.Builds {
					builds = append(builds, string(i))
				}

				fmt.Printf("%s/%s\n\t%s\n", info.Project_name, versionInfo.Version, strings.Join(builds, ", "))
			} else if slices.Contains(info.Versions, args[1]) {
				version := args[1]
				versionInfo := getVersionInfo(info.Project_Id, version)

				var builds []string

				for _, i := range versionInfo.Builds {
					builds = append(builds, string(i))
				}

				fmt.Printf("%s/%s\n\t%s\n", info.Project_name, version, strings.Join(builds, ", "))
			} else {
				log.Fatalf("'%s' | Version not found", args[1])
			}
		}
	} else {
		for _, project := range getProjects().Projects {
			fmt.Println(project)
		}
	}
}

func paperInstall(args []string) {
	if len(args) < 3 {
		log.Fatalln("Expected arguments.\n" + os.Args[0] + " mod install <project> <version> <build>")
	}

	downloadBuild(args[0], args[1], args[2], args[0]+"-"+args[1]+"-"+args[2]+".jar")
}
