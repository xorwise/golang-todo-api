package utils

import "os"

func IsMediaDirExists() bool {
	if _, err := os.Stat("media"); os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateMediaDir() error {
	if err := os.Mkdir("media", os.ModePerm); err != nil {
		return err
	}
	return nil
}
