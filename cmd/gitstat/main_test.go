package main

import (
	"os"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func TestDir(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("error")
	}
	_ = wd
	if exitcode := run2([]string{"cmd", "../../../"}, true); exitcode != 0 {
		t.Errorf("error executing run2: %d", exitcode)
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

func TestLog(t *testing.T) {
	// open the repo
	repo, err := git.PlainOpen("../../")
	if err != nil {
		t.Fatalf("error opening repo: %s", err)
	}
	entries, err := repo.Log(&git.LogOptions{
		All: true,
	})
	if err != nil {
		t.Fatalf("error getting log: %s", err)
	}
	_ = entries
	// l, err := entries.Next()
	// if err != nil {
	// 	t.Fatalf("error getting log: %s", err)
	// }
	// _ = l
	for l, err := entries.Next(); err == nil; l, err = entries.Next() {
		t.Logf("%s", l.String())
	}
}

func TestReferences(t *testing.T) {
	// open the repo
	repo, err := git.PlainOpen("../../")
	if err != nil {
		t.Fatalf("error opening repo: %s", err)
	}
	entries, err := repo.References()
	if err != nil {
		t.Fatalf("error getting log: %s", err)
	}
	_ = entries
	for l, err := entries.Next(); err == nil; l, err = entries.Next() {
		t.Logf("%s", l.String())
	}
}

func TestWithRemote(t *testing.T) {
	repo, err := git.PlainOpen("../../")
	if err != nil {
		t.Fatalf("error opening repo: %s", err)
	}
	head, err := repo.Head()
	if err != nil {
		t.Fatalf("error getting head: %s", err)
	}
	_ = head
	refs, err := repo.References()
	if err != nil {
		t.Fatalf("error getting refs: %s", err)
	}
	_ = refs
	originMasterHash, err := repo.ResolveRevision(plumbing.Revision("origin/master"))
	if err != nil {
		t.Fatalf("error getting origin/master: %s", err)
	}
	originMasterCommit, err := repo.CommitObject(*originMasterHash)
	if err != nil {
		t.Fatalf("error getting origin/master commit: %s", err)
	}
	headCommit, err := repo.CommitObject(head.Hash())
	if err != nil {
		t.Fatalf("error getting head commit: %s", err)
	}
	isAncestor, err := headCommit.IsAncestor(originMasterCommit)
	t.Logf("headCommit.IsAncestor(originMasterCommit): %v", isAncestor)
	isAncestor, err = originMasterCommit.IsAncestor(headCommit)
	t.Logf("originMasterCommit.IsAncestor(headCommit): %v", isAncestor)
}
