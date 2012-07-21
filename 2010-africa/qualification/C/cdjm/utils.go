//v 0.1
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

func Input(fname string) (nums []string, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	nums = make([]string, 0, len(lines))
	for _, l := range lines {

		if len(l) == 0 {
			continue
		}

		nums = append(nums, l)
	}

	return nums, nil
}
