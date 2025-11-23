package utils

import (
	"crypto/tls"
	"fmt"
	"gozero-sso-service/internal/config"
	"os"
	"sync"
	"time"

	"github.com/secsy/goftp"
)

type FtpsUploader struct {
	client  *goftp.Client
	config  goftp.Config
	address string
	mu      sync.Mutex
}

func NewFtpsUploader(info *config.FtpInfo) (*FtpsUploader, error) {
	address := fmt.Sprintf("%s:%s", info.Hostname, info.Port)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ClientSessionCache: tls.NewLRUClientSessionCache(2),
	}

	uploader := &FtpsUploader{
		address: address,
		config: goftp.Config{
			User:               info.Username,
			Password:           info.Password,
			TLSConfig:          tlsConfig,
			TLSMode:            goftp.TLSExplicit,
			ConnectionsPerHost: 1,
			Timeout:            20 * time.Second,
		},
	}

	return uploader, uploader.connect()
}

func (u *FtpsUploader) connect() error {
	client, err := goftp.DialConfig(u.config, u.address)
	if err != nil {
		return err
	}
	u.client = client
	return nil
}

func (u *FtpsUploader) Upload(localPath, remotePath string) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	// If a client is nil or broken → reconnect
	if u.client == nil {
		if err := u.connect(); err != nil {
			return err
		}
	}

	file, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Try upload
	if err := u.client.Store(remotePath, file); err != nil {
		// If failed, try once more after reconnect
		_ = u.client.Close()
		u.client = nil

		if err2 := u.connect(); err2 != nil {
			return fmt.Errorf("reconnect failed: %w", err2)
		}

		// second try
		file.Seek(0, 0)
		if err3 := u.client.Store(remotePath, file); err3 != nil {
			return fmt.Errorf("upload failed after reconnect: %w", err3)
		}
	}

	return nil
}
