package tokscan

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
)

func Run() (string, error) {
	args, err := Parse(os.Args[1:])
	if err != nil {
		return "", err
	}
	return Find(args)
}

func Find(args *Args) (string, error) {
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
			return "", err
		}
	}
	s, err := json.Marshal(ret)
	if err != nil {
		return "", err
	}
	return string(s), nil
}
