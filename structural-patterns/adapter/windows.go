package adapter

import "fmt"

// Windows 为实现Computer接口，需要适配器
type Windows struct {

}

func (w *Windows) insertIntoUSBPort()  {
	fmt.Println("Windows机器USB连接器插入。")
}