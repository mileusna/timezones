package timezones

import (
	"os"
	"strings"
	"time"
)

var zoneDir string

// List returns all timezones from most common OS directories on maOS and Linux with timezones defined
func List() []string {
	return list([]string{
		"/usr/share/zoneinfo/",
		"/usr/share/lib/zoneinfo/",
		"/usr/lib/locale/TZ/",
		"/system/usr/share/zoneinfo/",
	})
}

// ListDir returns timezones in specified directory, if your directory is not discovered with timezones.List()
func ListDir(dir string) []string {
	if !strings.HasPrefix(dir, "/") {
		dir += "/"
	}
	return list([]string{dir})
}

func list(zoneDirs []string) []string {
	tzones := []string{}
	for _, zoneDir = range zoneDirs {
		tzones = append(tzones, readDir("")...)
	}
	return tzones
}

func readDir(path string) []string {
	tzones := []string{}
	files, _ := os.ReadDir(zoneDir + path)
	for _, f := range files {
		if f.IsDir() {
			tzones = append(tzones, readDir(path+"/"+f.Name())...)
		} else {
			if _, err := time.LoadLocation((path + "/" + f.Name())[1:]); err != nil {
				continue // not TZ file
			}
			tzones = append(tzones, (path + "/" + f.Name())[1:])
		}
	}
	return tzones
}
