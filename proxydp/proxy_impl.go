package proxydp

import (
	"log"
)

// ProxyImplGit 与原始类实现相同的接口
type ProxyImplGit struct {
}

// Clone implementation
func (git ProxyImplGit) Clone(url string) error {
	log.Println("StartProxy Clone Git")
	g := Git{}
	err := g.Clone(url)
	log.Println("EndProxy Clone Git")
	return err
}

// Push implementation
func (git ProxyImplGit) Push(url string) error {
	log.Println("StartProxy Push Git")
	g := Git{}
	err := g.Clone(url)
	log.Println("EndProxy Push Git")
	return err
}
