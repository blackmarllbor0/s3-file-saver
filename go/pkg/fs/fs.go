package fs

import "os"

func CreateDirIfNotExists(dirName string) error {
	if _, err := os.Stat(dirName); err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(dirName, os.ModePerm)
		}

		return err
	}

	return nil
}
