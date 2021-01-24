package main

import (
	"os"
	"testing"
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

