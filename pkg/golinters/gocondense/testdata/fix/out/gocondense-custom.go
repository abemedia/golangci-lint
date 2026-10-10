//golangcitest:config_path testdata/gocondense-custom.yml
//golangcitest:expected_exitcode 0
package testdata

import "fmt"

func gocondense() {
	// this is 89 columns when condensed, within the custom max-len
	fmt.Println("alpha", "bravo", "charlie", "delta", "echo", "foxtrot", "golf", "hotel")

	if true {
		// this is 101 columns when condensed with the custom tab-width,
		// so it checks both settings are applied properly
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
}
