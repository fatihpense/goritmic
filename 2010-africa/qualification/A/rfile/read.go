package rfile

import (
	"io/ioutil"
	//"strconv"
	paketim "./aaa"
	"strings"
)

func PrintE(a string) {

	// üstte paketim yazdık artık
	//ccc.PrintD(a)
	paketim.PrintD(a)

}

// It would be better for such a function to return error, instead of handling
// it on their own.
func Readlns(fname string) (nums []string, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	// Assign cap to avoid resize on every append.
	nums = make([]string, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.

		//n, err := strconv.Atoi(l)
		//if err != nil { return nil, err }
		nums = append(nums, l) //n
	}

	return nums, nil
}
