package main

import (
	"testing"
	"os"
)

func pFile() string {
	wd, _ := os.Getwd()
	return wd + string(os.PathSeparator) + ".gitignore"
}
func pDir() string {
	wd, _ := os.Getwd()
	return wd + string(os.PathSeparator) + "examples"
}

func TestCwd(t *testing.T) {
	cwd, _ := Cwd()
	if cwd == "" {
		t.Fatal("Failed to deremine current working directory")
	}
}

func TestStrFmt(t *testing.T) {
	s1 := "foo is %{fooVar}, while bar is %{barVar}!"
	f1 := StrFmt(s1, FmtArgs{"fooVar":"f00", "barVar":"b4r"})
	e1 := "foo is f00, while bar is b4r!"
	if f1 != e1 {
		t.Fatal("["+f1+"] != ["+e1+"]")
	}
}

func TestPathInfo(t *testing.T) {
	d := pFile()
	if exists, info := PathInfo(d); exists {
		if info == nil {
			t.Fatal("PathInfo("+d+"): Returned as existing, but os.FileInfo is nil")
		}
	}
}

func TestFileExists(t *testing.T) {
	d := pFile()
	if !FileExists(d) {
		t.Fatal("FileExists("+d+"): Reported as not existing")
	}
	if FileExists("does-not-exist") {
		t.Fatal("FileExists(-): Reported as existing")
	}
}

func TestDirExists(t *testing.T) {
	d := pDir()
	if !DirExists(d) {
		t.Fatal("DirExists("+d+"): Reported as not existing")
	}
	if DirExists("does-not-exists") {
		t.Fatal("DirExists(-): Reported as existing")
	}
}