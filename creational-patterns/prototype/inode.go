package prototype

// Inode 原型接口
type Inode interface {
	print(string)
	clone() Inode
}
