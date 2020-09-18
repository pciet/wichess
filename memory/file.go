package memory

import "os"

func filePath(filename string) string { return Folder + "/" + filename }

func fileExists(filename string) bool {
	info, err := os.Stat(filePath(filename))
	if err != nil {
		return false
	}
	if info.IsDir() {
		return false
	}
	return true
}
