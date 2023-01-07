package tokscan

import (
	"bufio"
	"io"
	"strings"
)

func Scan(r io.ReadCloser, targets []string) ([]string, error) {
	defer r.Close()
	var ret []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		for _, target := range targets {
			if len(targets) == len(ret) {
				for _, target := range targets {
					if !in(target, ret) {
						goto SEARCH
					}
				}
				return ret, nil
			}

		SEARCH:
			if in(target, ret) {
				continue
			}
			if strings.Contains(scanner.Text(), target) {
				ret = append(ret, target)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return ret, nil
}

func in(target string, slice []string) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}
