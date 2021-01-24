package main

import (
	"os"
	"testing"

	"github.com/go-git/go-git/v5"
)

func TestDir(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("error")
	}
	_ = wd
	if exitcode := run2([]string{"cmd", "../../../"}, true); exitcode != 0 {
		t.Errorf("error executing run2: %d",exitcode)
	}
}

func TestMyDir(t *testing.T) {
	status := getRepoDirStatus("../../")
	if status.err != nil {
		t.Fatalf("error: %s", status.err)
	}
	if !allIsClean(status) {
		t.Errorf("status not clean: \n%s", status.Status.String())
	}
}

func TestConfigRemote(t *testing.T) {
	// open the repo
	repo, err := git.PlainOpen("../../")
	if err != nil {
		t.Fatalf("error opening repo: %s", err)
	}
	cfg, err := repo.Config()
	if err != nil {
		t.Fatalf("error getting config: %s", err)
	}
	_ = cfg

}