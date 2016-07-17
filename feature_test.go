package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileOneLine(t *testing.T) {

	assert := assert.New(t)

	s1, s2 := readFile("/tmp/git-redlab-tests/oneline")

	assert.Equal("LINE1", s1)
	assert.Equal("", s2)

}

func TestReadFileTwoLines(t *testing.T) {

	assert := assert.New(t)

	s1, s2 := readFile("/tmp/git-redlab-tests/twolines")

	assert.Equal("LINE1", s1)
	assert.Equal("LINE2", s2)

}

func TestReadFileTwoLinesWithSeparation(t *testing.T) {

	assert := assert.New(t)

	s1, s2 := readFile("/tmp/git-redlab-tests/twolineswithseparation")

	assert.Equal("LINE1", s1)
	assert.Equal("LINE2", s2)

}

func TestReadFileThreeLines(t *testing.T) {

	assert := assert.New(t)

	s1, s2 := readFile("/tmp/git-redlab-tests/threelines")

	assert.Equal("LINE1", s1)
	assert.Equal("LINE2\nLINE3", s2)

}
