package systemd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func EnsureService(servicePath, serviceContent string) error {
	if _, err := os.Stat(servicePath); os.IsNotExist(err) {
		if err := ioutil.WriteFile(servicePath, []byte(serviceContent), 0644); err != nil {
			return fmt.Errorf("failed to write systemd service file: %w", err)
		}
		if err := exec.Command("systemctl", "daemon-reload").Run(); err != nil {
			return fmt.Errorf("failed to reload systemd daemon: %w", err)
		}
		if err := exec.Command("systemctl", "enable", "auto-upgrade.service").Run(); err != nil {
			return fmt.Errorf("failed to enable service: %w", err)
		}
	}
	return nil
}
