package logging

import (
	"fmt"
	"os"
	"sync"
)

type File struct {
	f         *os.File
	name      string
	perm      os.FileMode
	size      int64
	maxSize   int64
	nextNum   int
	debug     bool
	formatter Formatter
	mutex     sync.Mutex
}

func MustNewFile(name string, perm os.FileMode, maxSize int64, debug bool, formatter Formatter) *File {
	file, err := NewFile(name, perm, maxSize, debug, formatter)
	if err != nil {
		panic(err)
	}
	return file
}

func fileSizeOrZero(filename string) int64 {
	info, err := os.Stat(filename)
	if err != nil {
		return 0
	}
	return info.Size()
}

func fileExistsf(format string, v ...interface{}) bool {
	_, err := os.Stat(fmt.Sprintf(format, v...))
	return err == nil
}

func NewFile(name string, perm os.FileMode, maxSize int64, debug bool, formatter Formatter) (file *File, err error) {
	file = &File{
		name:      name,
		perm:      perm,
		maxSize:   maxSize,
		nextNum:   1,
		debug:     debug,
		formatter: formatter,
	}

	if maxSize > 0 {
		file.size = fileSizeOrZero(name)

		for fileExistsf("%s.%d", name, file.nextNum) {
			file.nextNum++
		}

		if file.size >= file.maxSize {
			nextName := fmt.Sprintf("%s.%d", file.name, file.nextNum)
			err = os.Rename(file.name, nextName)
			if err != nil {
				return nil, err
			}
			file.nextNum++
			file.size = 0
		}
	}

	file.f, err = os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (file *File) split() {
	err := file.f.Close()
	if err != nil {
		panic(err)
	}
	nextName := fmt.Sprintf("%s.%d", file.name, file.nextNum)
	err = os.Rename(file.name, nextName)
	if err != nil {
		panic(err)
	}
	file.f, err = os.OpenFile(file.name, os.O_CREATE|os.O_WRONLY, file.perm)
	if err != nil {
		panic(err)
	}
	file.nextNum++
	file.size = 0
}

func (file *File) Close() error {
	return file.f.Close()
}

func (file *File) Printf(msg string, v ...interface{}) {
	file.mutex.Lock()
	defer file.mutex.Unlock()

	n, _ := fmt.Fprintln(file.f, file.formatter.Format(msg, v...))
	file.f.Sync()

	file.size += int64(n)
	if file.size >= file.maxSize {
		file.split()
	}
}

func (file *File) Debugf(msg string, v ...interface{}) {
	if !file.debug {
		return
	}

	file.mutex.Lock()
	defer file.mutex.Unlock()

	n, _ := fmt.Fprintln(file.f, file.formatter.Format(msg, v...))
	file.f.Sync()

	file.size += int64(n)
	if file.size >= file.maxSize {
		file.split()
	}
}

func (file *File) UnresolvedErrorf(err error, msg string, v ...interface{}) {
	file.mutex.Lock()
	defer file.mutex.Unlock()

	n, _ := fmt.Fprintln(file.f, file.formatter.FormatError(err, msg, v...))
	file.f.Sync()

	file.size += int64(n)
	if file.size >= file.maxSize {
		file.split()
	}
}
