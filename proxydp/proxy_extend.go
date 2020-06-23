package proxydp

import "log"

// ProxyExtendGit 继承原始类Git,重写原来的函数
type ProxyExtendGit struct {
	Git
}

// Clone Override Git{}.Clone()
func (extend ProxyExtendGit) Clone(url string) error {
	log.Println("StartExtend Git-Clone")
	// 调用父类
	err := extend.Git.Clone(url)
	log.Println("EndExtend Git-Clone")
	return err
}

// Push Override Git{}.PuPush()
func (extend ProxyExtendGit) Push(url string) error {
	log.Println("StartExtend Git-Push ")
	// 调用父类
	err := extend.Git.Push(url)
	log.Println("EndExtend Git-Push")
	return err
}
