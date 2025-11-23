package utils

import (
	"crypto/tls"
	"fmt"
	"gozero-sso-service/internal/config"
	"os"
	"path/filepath"
	"time"

	"github.com/secsy/goftp"
)

func UploadFileToFTPS(config *config.FtpInfo, localFilePath string) error {
	address := fmt.Sprintf("%s:%s", config.Hostname, config.Port)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ClientSessionCache: tls.NewLRUClientSessionCache(2),
	}

	configFTP := goftp.Config{
		User:               config.Username,
		Password:           config.Password,
		TLSConfig:          tlsConfig,
		TLSMode:            goftp.TLSExplicit,
		ConnectionsPerHost: 1,
		Timeout:            20 * time.Second,
	}

	client, err := goftp.DialConfig(configFTP, address)
	if err != nil {
		return err
	}

	defer client.Close()

	uploadFile, err := os.Open(localFilePath)
	if err != nil {
		return err
	}

	defer uploadFile.Close()

	_, fileName := filepath.Split(localFilePath)

	err = client.Store(fmt.Sprintf("%s/%s", config.BasePath, fileName), uploadFile)
	if err != nil {
		return err
	}
	return nil
}
