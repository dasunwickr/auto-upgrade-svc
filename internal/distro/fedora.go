package distro

import (
	"fmt"
	"os/exec"
	"strings"
	"text/template"

	"auto-upgrade/internal/updater"
)

type fedoraUpdater struct{}

func NewFedoraUpdater() updater.Updater {
	return &fedoraUpdater{}
}

func (f *fedoraUpdater) UpdateAndUpgrade() error {
	cmd := exec.Command("dnf", "upgrade", "-y")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("dnf upgrade failed: %w\nOutput: %s", err, output)
	}
	return nil
}

func (f *fedoraUpdater) ServiceFileContent() (string, error) {
	const serviceTemplate = `[Unit]
Description=Automatic DNF Upgrade at Boot
After=network.target

[Service]
Type=oneshot
ExecStart=/usr/local/bin/auto-upgrade
StandardOutput=journal
StandardError=journal
RemainAfterExit=true

[Install]
WantedBy=multi-user.target
`
	tmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		return "", err
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, nil); err != nil {
		return "", err
	}
	return buf.String(), nil
}
