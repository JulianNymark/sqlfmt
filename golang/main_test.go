package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSamples(t *testing.T) {
	// for each sample, test it
	files, err := ioutil.ReadDir("./samples")
	if err != nil {
		log.Fatal(err)
	}

	inputFiles := []string{}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "in_") {
			inputFiles = append(inputFiles, file.Name())
		}
	}

	// for each in_X, start a test comparing it to out_X
	for _, infile := range inputFiles {
		re := regexp.MustCompile("in_")
		outfile := re.ReplaceAllString(infile, "out_")
		_TestSample(t, infile, outfile)
	}
}

func _TestSample(t *testing.T, infile string, outfile string) {
	assert := assert.New(t)

	fd, _ := os.Open("./samples/" + outfile)
	expected, _ := ioutil.ReadAll(fd)
	input_r, _ := os.Open("./samples/" + infile)

	tokens := tokenizeInput(input_r)
	formatted := formatTokens(tokens)

	assert.Equal(string(expected), formatted)
}
