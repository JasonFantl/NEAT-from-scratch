package neat

import (
	"io"
	"os"
)

func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func AddToFile(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, "\n"+data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func Abs(val float64) float64 {
	if val >= 0 {
		return val
	}
	return -val
}
