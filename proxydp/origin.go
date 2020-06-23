package proxydp

import (
	"fmt"
	"strings"
)

// Git IGet Implementation
type Git struct {
}

// Clone 克隆
func (git *Git) Clone(url string) error {
	if strings.HasPrefix(url, "https") {
		fmt.Println("clone from " + url)
		return nil
	}
	fmt.Println("failed to clone from " + url)
	return fmt.Errorf("failed to clone from %s", url)
}

// Push 推送
func (git *Git) Push(url string) error {
	if strings.HasPrefix(url, "https") {
		fmt.Println("Push to " + url)
		return nil
	}
	fmt.Println("failed Push to " + url)
	return fmt.Errorf("failed Push to %s", url)
}
