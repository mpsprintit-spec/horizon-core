package wikipedia

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskSearchPermission() bool {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Println("Knowledge belum tersedia.")
	fmt.Print("Apakah Horizon boleh mencari di Wikipedia? (Y/N): ")

	answer, _ := reader.ReadString('\n')

	answer = strings.TrimSpace(strings.ToUpper(answer))

	return answer == "Y"
}

func AskSavePermission() bool {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Print("Apakah ingin menyimpan ke Knowledge Horizon? (Y/N): ")

	answer, _ := reader.ReadString('\n')

	answer = strings.TrimSpace(strings.ToUpper(answer))

	return answer == "Y"
}
