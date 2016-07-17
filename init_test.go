package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockedIO(t *testing.T) {

	assert := assert.New(t)

	var s string

	i := mockedIO{}

	i.ScanString(&s)

	assert.Equal("asd", s, "Mocked Scan should return asd")

}

type selection1IO struct {
	b *bytes.Buffer
	i *int
}

func (io selection1IO) ScanString(a *string) (n int, err error) {
	*a = "asd"
	return 0, nil
}

func (io selection1IO) ScanInt(a *int) (n int, err error) {
	switch *io.i {
	case 0:
		*a = 3
		break
	case 1:
		*a = -1
		break
	case 2:
		*a = 1
		break
	}

	*io.i = *io.i + 1
	return 0, nil
}

func (io selection1IO) Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(io.b, a...)
}

func (io selection1IO) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(io.b, format, a...)
}

func TestSelection1IO(t *testing.T) {

	assert := assert.New(t)

	var b bytes.Buffer
	i := 0

	io := selection1IO{&b, &i}

	io.Print("asd")
	io.Printf("%s-%d\n", "123", 456)

	assert.Equal("asd123-456\n", string(b.Bytes()))
}

type withname struct {
	Name string
}

func TestSelection(te *testing.T) {

	assert := assert.New(te)

	var b bytes.Buffer
	i := 0

	io := selection1IO{&b, &i}

	projects := []withname{
		withname{
			Name: "Project1",
		},
		withname{
			Name: "Project2",
		},
		withname{
			Name: "Project3",
		},
	}

	t := projects
	s := make([]interface{}, len(t))
	for i, v := range t {
		s[i] = v
	}

	sel := selection(s, io, func(i interface{}) string {
		return i.(withname).Name
	})

	assert.Equal(1, sel)
	assert.Equal("[0]: Project1\n[1]: Project2\n[2]: Project3\nPlease select project index: Wrong index given\nPlease select project index: Wrong index given\nPlease select project index: ", string(b.Bytes()))

}
