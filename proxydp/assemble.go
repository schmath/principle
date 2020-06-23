package proxydp

// AssembleGit 装配(对象实例化和注入)
type AssembleGit struct {
}

// API 装配层
func (assemble *AssembleGit) API() error {
	client := CleintGit{
		IGit: ProxyExtendGit{},
	}
	client.ClientClone("https://www.baidu.com")

	client = CleintGit{
		IGit: ProxyImplGit{},
	}
	client.ClientClone("https://www.163.com")

	return nil
}
