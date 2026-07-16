package wikipedia

import "fmt"

func ShowSummary(result *WikiResponse) {

	fmt.Println()
	fmt.Println("====================================")
	fmt.Println(result.Title)
	fmt.Println("====================================")
	fmt.Println()
	fmt.Println(result.Extract)
	fmt.Println()
}
