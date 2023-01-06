package tokscan

import (
	"bufio"
	"io"
	"strings"
)

func Scan(r io.ReadCloser, target string) (bool, error) {
	defer r.Close()
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
