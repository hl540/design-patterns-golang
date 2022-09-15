package adapter

import "fmt"

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("客户将 Lightning 连接器插入计算机。")
	com.InsertIntoLightningPort()
}
