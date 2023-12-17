package calendar

import (
	"fmt"
	"os"
	"path"
)

func getFileContent(day int, part int) (string, error) {
	content := ""

	current_dir, err := os.Getwd()
	if err != nil {
		return content, err
	}

	file_name := fmt.Sprintf("day%d_part%d.txt", day, part)
	data_path := path.Join(current_dir, "data", file_name)
	byte_content, err := os.ReadFile(data_path)
	if err != nil {
		return content, err
	}

	return string(byte_content), nil
}
