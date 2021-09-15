package storage

import "net/url"

type PTypes int

const (
	BINARY PTypes = iota
	SIG
	SUM
)

type Storage interface {
	Tree(namespace, name string) (map[string][]*ProviderInfo, error)

	GetPublicUrl() *url.URL
}

type ProviderInfo struct {
	Type      PTypes
	Filename  string
	Os        string
	Arch      string
	SHA256SUM string
}
