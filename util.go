package main

import (
	"strings"
	"fmt"
	"os"
	"path/filepath"
	"hash"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"encoding/hex"
)

type FmtArgs map[string]interface{}

func StrFmt(format string, arguments FmtArgs) string {
	args, i := make([]string, len(arguments)*2), 0
	for k, v := range arguments {
		args[i] = "%{" + k + "}"
		args[i+1] = fmt.Sprint(v)
		i += 2
	}
	return strings.NewReplacer(args...).Replace(format)
}

var cwd string

func Cwd() (string, error) {
	if cwd == "" {
		path, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			return "", err
		} else {
			cwd = path
		}
		if cwd == "" {
			panic("Failed to determined current working directory")
		}
	}
	return cwd, nil
}

func PathInfo(path string) (bool, os.FileInfo) {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err == nil {
		return true, stat
	}
	return false, nil
}

func FileExists(filePath string) bool {
	exists, stat := PathInfo(filePath)
	return exists && !stat.IsDir()
}

func DirExists(dirPath string) bool {
	exists, stat := PathInfo(dirPath)
	return exists && stat.IsDir()
}

func CheckHash(raw string, etalon string) (bool, error) {
	var hr hash.Hash
	if etalon[:5] == "{md5}" {
		hr = md5.New()
		etalon = etalon[5:]
	} else if etalon[:5] == "{sha}" {
		hr = sha1.New()
		etalon = etalon[5:]
	} else if etalon[:6] == "{sha1}" {
		hr = sha1.New()
		etalon = etalon[6:]
	} else if etalon[:8] == "{sha256}" {
		hr = sha256.New()
		etalon = etalon[8:]
	} else if etalon[:8] == "{sha512}" {
		hr = sha512.New()
		etalon = etalon[8:]
	} else {
		return false, errors.New("invalid hasher prefix")
	}
	hr.Write([]byte(raw))
	return hex.EncodeToString(hr.Sum(nil)) == etalon, nil
}
