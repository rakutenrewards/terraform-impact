package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
}

func ListDirsIn(dirPath string) []string {
	infos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}
	}

	var dirs []string
	for _, info := range infos {
		if info.IsDir() {
			dirs = append(dirs, info.Name())
		}
	}

	return dirs
}

func ListFilesIn(dirPath string) []string {
	infos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}
	}

	var files []string
	for _, info := range infos {
		if !info.IsDir() {
			files = append(files, info.Name())
		}
	}

	return files
}

func SamePath(a string, b string) bool {
	return filepath.Clean(a) == filepath.Clean(b)
}

func TraceSymlinkFile(filePath string) []string {
	if !exists(filePath) || IsDir(filePath) {
		return []string{}
	}

	result := []string{filePath}
	symlink, err := os.Readlink(filePath)
	if err == nil {
		symlinkPath := filepath.Join(filepath.Dir(filePath), symlink)
		result = append(result, TraceSymlinkFile(symlinkPath)...)
	}

	return result
}

func exists(filePath string) bool {
	_, err := os.Stat(filePath)

	return err == nil
}
