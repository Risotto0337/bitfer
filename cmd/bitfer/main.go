package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
	"bitfer/core"
)

/*
	Bitfer v0.3 core
	- multi-repo support
	- GitHub raw compatible
	- source build system
*/

type Recipe struct {
	Name string `yaml:"name"`

	Source struct {
		URL string `yaml:"url"`
	} `yaml:"source"`

	Build struct {
		Steps []string `yaml:"steps"`
	} `yaml:"build"`
}

type RepoIndex map[string]struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func readFileOrHTTP(path string) ([]byte, error) {
	if strings.HasPrefix(path, "http") {
		resp, err := http.Get(path)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return ioutil.ReadAll(resp.Body)
	}
	return ioutil.ReadFile(path)
}

/* ---------------- MULTI-REPO LOADER ---------------- */

func loadRepoConfig() []string {
	data, err := ioutil.ReadFile("/etc/bitfer/repos.json")
	if err != nil {
		return []string{}
	}

	var cfg struct {
		Repos []string `json:"repos"`
	}

	json.Unmarshal(data, &cfg)
	return cfg.Repos
}

func loadAllIndexes() RepoIndex {
	merged := RepoIndex{}
	repos := loadRepoConfig()

	for _, url := range repos {
		data, err := readFileOrHTTP(url)
		if err != nil {
			continue
		}

		var index RepoIndex
		json.Unmarshal(data, &index)

		for k, v := range index {
			merged[k] = v
		}
	}

	return merged
}

/* ---------------- PACKAGE RESOLVER ---------------- */

func resolvePackage(pkg string) string {
	index := loadAllIndexes()

	if entry, ok := index[pkg]; ok {
		return entry.Path
	}

	return ""
}

/* ---------------- RECIPE LOADER ---------------- */

func loadRecipe(path string) Recipe {
	data, err := readFileOrHTTP(path)
	if err != nil {
		panic(err)
	}

	var r Recipe
	yaml.Unmarshal(data, &r)
	return r
}

/* ---------------- MAIN ---------------- */

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("  bitfer install <pkg>")
		fmt.Println("  bitfer remove <pkg>")
		return
	}

	cmd := os.Args[1]
	pkg := os.Args[2]

	/* -------- REMOVE -------- */
	if cmd == "remove" {
		core.Remove(pkg)
		fmt.Println("Removed:", pkg)
		return
	}

	/* -------- INSTALL -------- */
	if cmd == "install" {

		path := resolvePackage(pkg)

		if path == "" {
			fmt.Println("Package not found:", pkg)
			return
		}

		fmt.Println("Fetching recipe:", path)

		r := loadRecipe(path)

		workdir := "/tmp/bitfer-build"
		os.MkdirAll(workdir, 0755)

		fmt.Println("Fetching source...")
		core.Fetch(r.Source.URL, workdir)

		fmt.Println("Building...")
		core.Build(r.Build.Steps, workdir)

		fmt.Println("Installing...")
		core.Install(workdir+"/pkg", r.Name)

		fmt.Println("DONE:", r.Name)
	}
}
