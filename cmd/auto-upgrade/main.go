package main

import (
	"log"
	"os"

	"github.com/dasunwickr/auto-upgrade-svc/internal/distro"
	"github.com/dasunwickr/auto-upgrade-svc/internal/systemd"
)

func main() {
	if os.Geteuid() != 0 {
		log.Fatal("Must be run as root")
	}

	updater, err := distro.DetectUpdater()
	if err != nil {
		log.Fatalf("Failed to detect distro updater: %v", err)
	}

	serviceContent, err := updater.ServiceFileContent()
	if err != nil {
		log.Fatalf("Failed to generate systemd service content: %v", err)
	}

	const servicePath = "/etc/systemd/system/auto-upgrade.service"

	if err := systemd.EnsureService(servicePath, serviceContent); err != nil {
		log.Fatalf("Failed to setup systemd service: %v", err)
	}

	if err := updater.UpdateAndUpgrade(); err != nil {
		log.Fatalf("Upgrade failed: %v", err)
	}

	log.Println("Upgrade completed successfully.")
}
