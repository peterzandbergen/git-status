package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

func allIsClean(st repoDirStatus) bool {
	for _, fs := range st.Status {
		if fs.Staging != git.Unmodified {
			return false
		}
		if fs.Worktree != git.Unmodified {
			return false
		}
	}
	return true
}

func repoStatus(dir string) (git.Status, error) {
	repo, err := git.PlainOpen(dir)
	if err != nil {
		return nil, fmt.Errorf("error opening repo: %w", err)
	}
	wt, err := repo.Worktree()
	if err != nil {
		return nil, fmt.Errorf("error opening worktree: %w", err)
	}
	status, err := wt.Status()
	if err != nil {
		return nil, fmt.Errorf("error getting status: %w", err)
	}
	return status, nil
}

type repoDirStatus struct {
	git.Status

	err  error
	path string
}

func getRepoDirStatus(path string) repoDirStatus {
	res := repoDirStatus{
		path: path,
	}
	state, err := repoStatus(path)
	if err != nil {
		res.err = err
	}
	res.Status = state
	return res
}

func reposStatus(reposDir string) ([]repoDirStatus, error) {
	// get dir
	dirInfo, err := os.Lstat(reposDir)
	if err != nil {
		return nil, err
	}
	// Get the directories
	if !dirInfo.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", reposDir)
	}
	// Open the directory
	df, err := os.Open(reposDir)
	if err != nil {
		return nil, fmt.Errorf("error opening dir: %w", err)
	}
	defer df.Close()
	dirs, err := df.Readdir(0)
	if err != nil {
		return nil, fmt.Errorf("error reading dir: %w", err)
	}
	var res []repoDirStatus
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		res = append(res, getRepoDirStatus(filepath.Join(reposDir, dir.Name())))
	}
	return res, nil
}

func run2(args []string, verbose bool) int {
	if len(args) < 2 {
		fmt.Println("Directory name missing")
		return 1
	}
	dir := args[1]
	states, err := reposStatus(dir)
	if err != nil {
		fmt.Printf("error reading states: %s", err)
		return 2
	}
	for _, status := range states {
		if !allIsClean(status) {
			fmt.Printf("%s is not clean\n", status.path)
			if verbose {
				fmt.Printf("%s\n", status.Status.String())
			}
			
		}
		
	}
	_ = states
	return 0
}

func run1() {
	if len(os.Args) < 2 {
		fmt.Println("Directory name missing")
		os.Exit(1)
	}
	dir := os.Args[1]
	repo, err := git.PlainOpen(dir)
	if err != nil {
		fmt.Printf("Error opening repo: %s\n", err)
		os.Exit(2)
	}
	wt, err := repo.Worktree()
	if err != nil {
		fmt.Printf("Error getting worktree: %s", err)
		os.Exit(3)
	}
	status, err := wt.Status()
	if err != nil {
		fmt.Printf("Error getting worktree status: %s", err)
		os.Exit(4)
	}
	if !status.IsClean() {
		fmt.Printf("Worktree is not clean: %s", status.String())
		os.Exit(4)
	}
	os.Exit(0)
}

func main() {
	os.Exit(run2(os.Args, len(os.Args) == 3))
}
