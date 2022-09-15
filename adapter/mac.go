package adapter

import "fmt"

// Mac 实现client可以使用的Computer接口
type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("mac机器Lightning连接器插入。")
}
