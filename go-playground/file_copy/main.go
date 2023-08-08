package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	n1, err := copy1("hello.log", "hello_copy1.log")
	if err != nil {
		fmt.Println(err)
	}
	n2, err := copy2("hello.log", "hello_copy2.log")
	if err != nil {
		fmt.Println(err)
	}
	n3, err := copy3("hello.log", "hello_copy3.log", 1024)
	if err != nil {
		fmt.Println(err)
	}
	if n1 != n2 || n2 != n3 {
		fmt.Println("copy error", n1, n2, n3)
	} else {
		fmt.Println("copy success", n1, n2, n3)
	}
}

func copy1(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	return io.Copy(destination, source)
}

func copy2(src, dst string) (int64, error) {
	source, err := os.ReadFile(src)
	if err != nil {
		return 0, err
	}

	_, err = os.Stat(dst)
	if err == nil {
		return 0, fmt.Errorf("%s already exists", dst)
	}

	err = os.WriteFile(dst, source, 0644)
	if err != nil {
		return 0, err
	}
	return int64(len(source)), nil
}

func copy3(src, dst string, bufSize int64) (int64, error) {
	buf := make([]byte, bufSize)

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	_, err = os.Stat(dst)
	if err == nil {
		return 0, fmt.Errorf("%s already exists", dst)
	}

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	var total int64
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}
		total += int64(n)
		if n == 0 {
			// EOF
			break
		}
		if _, err := destination.Write(buf[:n]); err != nil {
			return 0, err
		}
	}
	return total, nil
}
