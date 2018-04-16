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

func TestCheckHash(t *testing.T) {
	s := "secret"
	v := map[string]string{
		"md5" : "5ebe2294ecd0e0f08eab7690d2a6ee69",
		"sha" : "e5e9fa1ba31ecd1ae84f75caaa474f3a663f05f4",
		"sha1" : "e5e9fa1ba31ecd1ae84f75caaa474f3a663f05f4",
		"sha256" : "2bb80d537b1da3e38bd30361aa855686bde0eacd7162fef6a25fe97bf527a25b",
		"sha512" : "bd2b1aaf7ef4f09be9f52ce2d8d599674d81aa9d6a4421696dc4d93dd0619d682ce56b4d64a9ef097761ced99e0f67265b5f76085e5b0ee7ca4696b2ad6fe2b2",
	}
	for h, d := range v {
		if ok, err := CheckHash(s, "{" + h + "}" + d); !ok {
			if err != nil {
				t.Fatal(err)
			} else {
				t.Fatal()
			}
		}
	}
}