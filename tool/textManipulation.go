package tool

import "regexp"

// ProcessLine replace old name in line by new one
func ProcessLine(line, oldName, newName string) (found bool, res string, occ int) {
	reOldName := regexp.MustCompile("(?i)" + oldName + "([\\s,.])")
	processedLine := reOldName.ReplaceAllString(line, newName+"$1")
	nbOcc := len(reOldName.FindAllStringIndex(line, -1))
	return nbOcc > 0, processedLine, nbOcc
}
