package lines

import (
	"os"
	"bytes"
	"fmt"
	"io"
)

const SIZE_SMALL = 5000
const SIZE_BIG  = 4096

type Liner struct {
	f *os.File
}

func NewLiner(f *os.File) *Liner {
	return &Liner{
		f: f,
	}
}

func (l *Liner) NumLines() int {
	pos, err := l.f.Seek(0, os.SEEK_CUR)
	if err != nil {
		return -1
	}
	b := make([]byte, SIZE_BIG)
	sum := 0
	for {
		n, _ := l.f.Read(b)
		if n == 0 {
			break
		}
		sum += bytes.Count(b[:n], []byte{'\n'})
	}
	pos, err = l.f.Seek(0, int(pos))
	return sum
}

// Return a line without the trailing delimiter
func (l *Liner) ReadLine() (string, error) {
	var ret string
	b := make([]byte, SIZE_SMALL)
	pos, err := l.f.Seek(0, os.SEEK_CUR)
	if err != nil {
		return "", err
	}
	
	for {
		n, err := l.f.Read(b)
		
		if n > 0 {
			i := bytes.Index(b[:n], []byte{'\n'})
			if i != -1 {
				ret += string(b[:i])
				break
			}
			ret += string(b[:n])
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			return ret, err
		}
	}		
	l.f.Seek(int64(len(ret) + 1), int(pos))
	return ret, nil
}

func (l *Liner) SeekLine(line int) error {
	_, err := l.f.Seek(0, os.SEEK_SET)
	if err != nil {
		return err
	}
	for i := 0; i < line; i++ {
		li, err := l.ReadLine()
		fmt.Println(li)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *Liner) ReadLines(count int) ([]string, error) {
	var ret []string
	for i := 0; i < count; i++ {
		line, err := l.ReadLine()
		if len(line) > 0 {
			ret = append(ret, line)
		}
		if err != nil {
			return ret, err
		}
	}
	return ret, nil
}

func (l *Liner) ReadAll() ([]string, error) {
	var ret []string
	var err error
	n := l.NumLines()
	if n > 0 {
		ret, err = l.ReadLines(n)
	}
	return ret, err
}


