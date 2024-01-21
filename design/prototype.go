// Package main
// @Description: 原型模式
package main

import (
	"fmt"
	"sync"
)

type Inode interface {
	print(string)
	clone() (Inode, error)
}

type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}
func (f *File) clone() (Inode, error) {
	return &File{f.name + "_clone"}, nil
}

// Folder
// @Description:
type Folder struct {
	children     []Inode // files
	name         string
	lockChildren sync.RWMutex // 孩子结点的读写锁
}

// print
//
//	@Description:
//	@receiver f
//	@param indentation
func (f *Folder) print(indentation string) {
	// read lock
	f.lockChildren.RLock()
	defer f.lockChildren.RUnlock()

	fmt.Println(indentation + f.name)
	for _, child := range f.children {
		child.print(indentation + indentation)
	}

}

func (f *Folder) clone() (Inode, error) {
	// write lock
	f.lockChildren.Lock()
	defer f.lockChildren.Unlock()
	// recursion CLONE
	folder := &Folder{name: f.name}
	var childrenTmp []Inode
	for _, child := range f.children {
		c, _ := child.clone()
		childrenTmp = append(childrenTmp, c)
	}
	folder.children = childrenTmp
	return folder, nil
}

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	cloneFolder, _ := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}
