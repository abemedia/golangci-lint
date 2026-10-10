//golangcitest:config_path testdata/gocondense.yml
//golangcitest:expected_exitcode 0
package testdata

import "fmt"

func gocondense() {
	fmt.Println(
		"foo",
		"bar",
	)

	_ = []int{
		1,
		2,
		3,
	}

	// this is 89 columns when condensed, so it stays on multiple lines
	fmt.Println(
		"alpha",
		"bravo",
		"charlie",
		"delta",
		"echo",
		"foxtrot",
		"golf",
		"hotel",
	)
}
