package storage

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/mod/semver"
	"net/url"
	ioos "os"
	"path"
	"strings"
)

type FS struct {
	path string
	url  *url.URL
}

func NewFs(path string, url *url.URL) *FS {
	return &FS{path, url}
}

func (f *FS) GetPublicUrl() *url.URL {
	return f.url
}

func (f *FS) Tree(namespace, name string) (map[string][]*ProviderInfo, error) {
	root := path.Join(f.path, namespace, name)
	result := make(map[string][]*ProviderInfo)

	dirs, err := ioos.ReadDir(root)

	if err != nil {
		return nil, err
	}

	for _, d := range dirs {

		if !semver.IsValid("v"+d.Name()) || !d.IsDir() {
			continue
		}
		//todo validate dir structure. SHA256SUMS.sig, SHA256SUMS, *.zip

		result[d.Name()] = []*ProviderInfo{}
		files, err := ioos.ReadDir(path.Join(root, d.Name()))
		if err != nil {
			log.Error(err)
			continue
		}

		//todo reduce count of loops.
		sums := make(map[string]string)
		for _, f := range files {
			if !strings.HasSuffix(f.Name(), "_SHA256SUMS") {
				continue
			}

			sums, err = parseSHA256SUMS(path.Join(root, d.Name(), f.Name()))
			if err != nil {
				log.Fatal(err)
			}
		}

		for _, f := range files {
			if f.IsDir() {
				continue
			}
			pf := &ProviderInfo{Filename: f.Name()}

			switch {
			case strings.Contains(f.Name(), "_SHA256SUMS.sig"):
				pf.Type = SIG
			case strings.HasSuffix(f.Name(), "_SHA256SUMS"):
				pf.Type = SUM
			case strings.HasSuffix(f.Name(), ".zip"):
				if s := strings.Split(strings.TrimSuffix(f.Name(), ".zip"), "_"); len(s) == 4 {
					pf.Type, pf.Os, pf.Arch = BINARY, s[2], s[3]
				}
				fmt.Printf("FNAME: %#v, SUMS: %#v\n", f.Name(), sums)
				pf.SHA256SUM = sums[f.Name()]
			default:
				log.Error("Unknown type of file: ", f.Name())
			}

			result[d.Name()] = append(result[d.Name()], pf)
		}
	}

	return result, nil
}

func parseSHA256SUMS(filepath string) (map[string]string, error) {

	file, err := ioos.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := make(map[string]string)
	for scanner.Scan() {
		splitted := strings.FieldsFunc(scanner.Text(), func(r rune) bool { return r == ' ' })
		result[splitted[1]] = splitted[0]
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
