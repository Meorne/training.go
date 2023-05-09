package tool

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FindReplaceFile(srcFile, oldName, newName string) (int, int, string, string, error) {
	rData, rErr := os.ReadFile(srcFile)
	if rErr != nil {
		return 0, 0, "", "", rErr
	}
	if len(rData) == 0 {
		return 0, 0, "", "", fmt.Errorf("file is empty %v", srcFile)
	}

	oData, oErr := os.Open(srcFile)
	if oErr != nil {
		return 0, 0, "", "", oErr
	}

	nbOcc := 0
	lines := make([]string, 0)
	sResult := make([]string, 0)
	scanner := bufio.NewScanner(oData)
	i := 0
	for scanner.Scan() {
		i++
		line := scanner.Text()
		found, res, occ := ProcessLine(line, oldName, newName)

		if found == true {
			lines = append(lines, fmt.Sprintf("%d", i))
		}
		nbOcc += occ

		sResult = append(sResult, res)

	}
	return len(lines), nbOcc, strings.Join(lines, " - "), strings.Join(sResult, "\n"), nil
}

func WriteNewFile(finalText string) error {
	file, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(finalText)
	if err != nil {
		return err
	}

	fmt.Println("== output file has been created ==")
	return nil
}
