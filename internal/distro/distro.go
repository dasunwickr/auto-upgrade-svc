package distro

import (
	"errors"
	"os"
	"strings"

	"github.com/dasunwickr/auto-upgrade-svc/internal/updater"
)

func DetectUpdater() (updater.Updater, error) {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return nil, err
	}
	content := string(data)

	switch {
	case strings.Contains(content, "Fedora"):
		return NewFedoraUpdater(), nil
	case strings.Contains(content, "Ubuntu"):
		// return NewUbuntuUpdater(), nil // placeholder
		return nil, errors.New("Ubuntu updater not implemented yet")
	default:
		return nil, errors.New("unsupported Linux distribution")
	}
}
