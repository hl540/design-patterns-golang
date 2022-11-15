package prototype

import "fmt"

// Folder 具体原型
type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) print(s string) {
	fmt.Println(s + f.name)
	for _, i := range f.children {
		i.print(s + s)
	}
}

func (f *Folder) clone() Inode {
	var tempChildren []Inode
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	return &Folder{
		children: tempChildren,
		name:     f.name + "_clone",
	}
}
