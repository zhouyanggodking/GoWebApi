package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	myFolder := "myfolder"
	myFile := "myfile.txt"
	// create folder
	//createFolder(myFolder)
	//
	//os.Chdir(myFolder)
	//
	//// create file
	//createFile(myFile)
	//
	//os.Chdir("..")

	// copy myfolder to destfolder

	destFolder := path.Join("destfolder", "sub1", "sub2")

	os.MkdirAll(destFolder, os.ModePerm)

	//
	//createFolder(destFolder)
	//
	//err := copyFile(myFolder + "\\" + myFile, destFolder + "\\test")
	//fmt.Print(err)

	//fds, _ := ioutil.ReadDir(myFolder)
	//
	//for _, fd := range fds {
	//	fmt.Println(fd.IsDir())
	//	fmt.Println(fd.Name())
	//}

	//copyDir(myFolder, destFolder)

	file, _ := os.Stat(path.Join(myFolder, myFile))

	fmt.Print(file.Size()) // check for file size
}

func createFolder(folderName string) error{
	if _, err := os.Stat(folderName); os.IsNotExist(err) { // folder not exists
		return os.Mkdir(folderName, os.ModePerm)
	}
	return nil
}

func createFile(filename string) (*os.File, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return os.Create(filename)
	}
	return nil, os.ErrExist
}

func copyFile(src, dest string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dest); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dest, srcinfo.Mode())
}

// Dir copies a whole directory recursively
func copyDir(src, dest string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	// if src not exists, just return err
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	// create dest fold with the same permission
	if err = os.MkdirAll(dest, srcinfo.Mode()); err != nil {
		return err
	}
	// read all entries in src folder
	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dest, fd.Name())

		if fd.IsDir() { // copy dir recursively
			if err = copyDir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else { // copy plain file
			if err = copyFile(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
