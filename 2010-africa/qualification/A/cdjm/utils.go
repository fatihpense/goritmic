package cdjm

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func Output(filename string, cases []string) {

	for i, v := range cases {
		cases[i] = "Case #" + strconv.Itoa(i+1) + ": " + v

	}

	ioutil.WriteFile(filename, []byte(strings.Join(cases, "\n")), 777)
	//WriteFile(filename string, data []byte, perm os.FileMode) error
}
