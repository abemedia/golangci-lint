//golangcitest:config_path testdata/gocondense.yml
package testdata

import "fmt"

func gocondense() {
	// want +1 "File is not properly formatted"
	fmt.Println(
		"foo",
		"bar",
	)
}
