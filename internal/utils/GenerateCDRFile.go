package utils

import (
	"bufio"
	"fmt"
	"gozero-sso-service/internal/config"
	"gozero-sso-service/internal/dataaccess/model"
	"os"
	"time"
)

func getFilename(is8x79 bool, config *config.Config) string {
	t := time.Now().Format("20060102150405000")
	filename := "VNN" + t + ".bil"
	if is8x79 {
		return fmt.Sprintf("%s/8x79/%s", config.Job.UploadPath, filename)
	} else {
		return fmt.Sprintf("%s/other/%s", config.Job.UploadPath, filename)
	}
}

func GenerateVT8x79CDRFile(config *config.Config, items []model.VT8x79) (string, error) {
	filename := getFilename(true, config)
	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)

	for _, item := range items {
		line := fmt.Sprintf("%s:%s:%s:%s:%s:%v:%v\n",
			item.UserID,
			item.ServiceID,
			item.Keyword,
			item.SubmitDate, // first send_time
			item.SubmitDate, // second send_time
			item.MessageType,
			item.NumberMessage,
		)

		if _, err := writer.WriteString(line); err != nil {
			return "", err
		}
	}

	if err := writer.Flush(); err != nil {
		return "", err
	}

	return filename, nil
}

func GenerateVTOtherCDRFile(config *config.Config, items []model.VTOther) (string, error) {
	filename := getFilename(false, config)
	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)

	for _, item := range items {
		line := fmt.Sprintf("%s:%s:%s:%s:%s:%v:%v\n",
			item.UserID,
			item.ServiceID,
			item.Keyword,
			item.SubmitDate, // first send_time
			item.SubmitDate, // second send_time
			item.MessageType,
			item.NumberMessage,
		)

		if _, err := writer.WriteString(line); err != nil {
			return "", err
		}
	}

	if err := writer.Flush(); err != nil {
		return "", err
	}

	return filename, nil
}
