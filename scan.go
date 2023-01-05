package tokscan

import (
	"bufio"
	"io"
	"strings"
)

func Scan(r io.Reader, target string) (bool, error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), target) {
			return true, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return false, err
	}
	return false, nil
}
