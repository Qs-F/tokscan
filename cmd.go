package tokscan

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func Run() error {
	args, err := Parse(os.Args[1:])
	if err != nil {
		return err
	}

	// Key is file path, value is list of tokens used in that file
	ret := make(map[string][]string)

	for _, dir := range args.Dirs {
		if err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			usedTokens, err := Scan(file, AllTokens())
			if err != nil {
				return err
			}
			ret[file.Name()] = usedTokens
			return nil
		}); err != nil {
			return err
		}
	}
	s, err := json.Marshal(ret)
	if err != nil {
		return err
	}
	fmt.Println(string(s))
	return nil
}
