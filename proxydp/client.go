package proxydp

// CleintGit Git的调用者，依赖接口
type CleintGit struct {
	IGit
}

// ClientClone 高层用户接口层
func (client *CleintGit) ClientClone(url string) error {
	client.IGit.Clone(url)
	return nil
}
