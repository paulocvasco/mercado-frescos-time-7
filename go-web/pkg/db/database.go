package db

import "os"

func Save(path string, data []byte) error {
	var err error
	var file *os.File

	defer file.Close()
	// check if exists a file to save data
	_, err = os.Stat(path)
	if err != nil {
		file, err = os.Create(path)
		if err != nil {
			return err
		}
	} else {
		file, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return err
		}
	}
	// save data
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func Load(path string) (data []byte, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	data, err = os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}
