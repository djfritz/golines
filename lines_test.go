package lines_test

import (
	"testing"
	"os"
	"code.google.com/p/golines"
)

func TestNumLines(t *testing.T) {
	f, _ := os.Open("test.input")
	l := lines.NewLiner(f)
	n := l.NumLines()
	if n != 10 {
		t.Errorf("mismatched length: %d\n", n)
	}
}

func TestReadLine(t *testing.T) {
	f, _ := os.Open("test.input")
	l := lines.NewLiner(f)
	line, err := l.ReadLine()
	if err != nil {
		t.Errorf(err.Error())
	}
	if line != "line 0" {
		t.Errorf("invalid line: %v\n", line)
	}
	line, err = l.ReadLine()
	if line != "line 1" {
		t.Errorf("invalid line: %v\n", line)
	}
}

func TestReadLines(t *testing.T) {
	f, _ := os.Open("test.input")
	l := lines.NewLiner(f)
	line, err := l.ReadLines(2)
	if err != nil {
		t.Errorf(err.Error())
	}
	if line[0] != "line 0" {
		t.Errorf("invalid line: %v\n", line[0])
	}
	if line[1] != "line 1" {
		t.Errorf("invalid line: %v\n", line[1])
	}
}

func TestSeekLine(t *testing.T) {
	f, _ := os.Open("test.input")
	l := lines.NewLiner(f)
	err := l.SeekLine(5)
	if err != nil {
		t.Fatalf(err.Error())
	}
	line, err := l.ReadLines(2)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if line[0] != "line 5" {
		t.Errorf("invalid line: %v\n", line[0])
	}
	if line[1] != "line 6" {
		t.Errorf("invalid line: %v\n", line[1])
	}
}

func TestReadAll(t *testing.T) {
	f, _ := os.Open("test.input")
	l := lines.NewLiner(f)
	line, err := l.ReadAll()
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(line) != 10 {
		t.Errorf("invalid length: %d\n", len(line))
	}
}
