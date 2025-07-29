<div align="left" style="position: relative;">
<img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" align="right" width="30%" style="margin: -20px 0 0 20px;">
<h1>AUTO-UPGRADE-SVC</h1>
<p align="left">
	<em>A simple, extensible system service for automatically updating and upgrading Linux packages at boot.</em>
</p>
<p align="left">
	<img src="https://img.shields.io/github/license/dasunwickr/auto-upgrade-svc?style=default&logo=opensourceinitiative&logoColor=white&color=007dff" alt="license">
	<img src="https://img.shields.io/github/last-commit/dasunwickr/auto-upgrade-svc?style=default&logo=git&logoColor=white&color=007dff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/dasunwickr/auto-upgrade-svc?style=default&color=007dff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/dasunwickr/auto-upgrade-svc?style=default&color=007dff" alt="repo-language-count">
</p>
</div>
<br clear="right">

## Quick Links

- [Overview](#-overview)
- [Features](#-features)
- [Project Structure](#-project-structure)
  - [Project Index](#-project-index)
- [Getting Started](#-getting-started)
  - [Prerequisites](#-prerequisites)
  - [Installation](#-installation)
  - [Usage](#-usage)
  - [Testing](#-testing)
- [Project Roadmap](#-project-roadmap)
- [Contributing](#-contributing)
- [License](#-license)
- [Acknowledgments](#-acknowledgments)

---

## Overview

`auto-upgrade-svc` is a lightweight Go-based system service designed to automatically update and upgrade packages on supported Linux distributions during system boot.
Currently, Fedora is supported with plans to expand to other distros.

The project is designed with scalability and modularity in mind to support multiple Linux package managers and system init systems.

---

## Features

- Automatically detects Linux distribution and runs the appropriate package update commands.
- Supports Fedora out of the box.
- Generates and manages a systemd service unit to run updates at boot.
- Modular internal architecture to add support for other distros and init systems easily.
- Runs with sudo privileges only when needed, keeping security in mind.

---

## Project Structure

```sh
└── auto-upgrade-svc/
    ├── README.md
    ├── cmd
    │   └── auto-upgrade
    │       └── main.go
    ├── go.mod
    └── internal
        ├── distro
        ├── systemd
        └── updater
````

### Project Index

<details open>
	<summary><b><code>AUTO-UPGRADE-SVC/</code></b></summary>
	<details>
		<summary><b>Root</b></summary>
		<table>
		<tr>
			<td><b><a href='https://github.com/dasunwickr/auto-upgrade-svc/blob/master/go.mod'>go.mod</a></b></td>
			<td>Defines module path and dependencies</td>
		</tr>
		</table>
	</details>
	<details>
		<summary><b>cmd/auto-upgrade</b></summary>
		<table>
		<tr>
			<td><b><a href='https://github.com/dasunwickr/auto-upgrade-svc/blob/master/cmd/auto-upgrade/main.go'>main.go</a></b></td>
			<td>Program entrypoint, initializes updater and systemd manager</td>
		</tr>
		</table>
	</details>
	<details>
		<summary><b>internal/updater</b></summary>
		<table>
		<tr>
			<td><b><a href='https://github.com/dasunwickr/auto-upgrade-svc/blob/master/internal/updater/updater.go'>updater.go</a></b></td>
			<td>Updater interface and Fedora implementation</td>
		</tr>
		</table>
	</details>
	<details>
		<summary><b>internal/systemd</b></summary>
		<table>
		<tr>
			<td><b><a href='https://github.com/dasunwickr/auto-upgrade-svc/blob/master/internal/systemd/manager.go'>manager.go</a></b></td>
			<td>Systemd service manager to create and enable units</td>
		</tr>
		</table>
	</details>
	<details>
		<summary><b>internal/distro</b></summary>
		<table>
		<tr>
			<td><b><a href='https://github.com/dasunwickr/auto-upgrade-svc/blob/master/internal/distro/distro.go'>distro.go</a></b></td>
			<td>Detects Linux distro and returns appropriate updater</td>
		</tr>
		<tr>
			<td><b><a href='https://github.com/dasunwickr/auto-upgrade-svc/blob/master/internal/distro/fedora.go'>fedora.go</a></b></td>
			<td>Fedora-specific distro detection and updater</td>
		</tr>
		</table>
	</details>
</details>

---

## Getting Started

### Prerequisites

* **Go** (1.20 or newer recommended) installed on your development system.
* Linux system with Fedora for initial testing.

---

### Installation

Build the project from source:

```sh
git clone https://github.com/dasunwickr/auto-upgrade-svc.git
cd auto-upgrade-svc
go build ./cmd/auto-upgrade
```

This produces a binary named `auto-upgrade` in the current directory.

---

### Usage

Run the binary manually or install the systemd service it creates:

```sh
sudo ./auto-upgrade
```

To run directly with `go run`:

```sh
go run ./cmd/auto-upgrade
```

The service will:

* Update package metadata and upgrade packages automatically on system boot.
* Create and enable the systemd service unit during the first run.

---

### Testing

Run the automated test suite with:

```sh
go test ./...
```

---

## Project Roadmap

* [x] Initial Fedora support with DNF update & upgrade commands.
* [ ] Add support for Ubuntu/Debian (APT).
* [ ] Add support for Arch Linux (Pacman).
* [ ] Add support for other init systems (SysV, OpenRC).
* [ ] Implement automatic rollback on failed updates.
* [ ] Add logging and metrics support.

---

## Contributing

Contributions are welcome! Please follow the standard GitHub flow:

* Fork the repository.
* Create a feature branch.
* Make your changes and test locally.
* Submit a pull request.

See the [Contributing Guidelines](https://github.com/dasunwickr/auto-upgrade-svc/blob/main/CONTRIBUTING.md) for more details.

---

## License

This project is licensed under the **GNU General Public License v3 (GPLv3)**.

You can view the full license [here](https://www.gnu.org/licenses/gpl-3.0.en.html).

---

## Acknowledgments

* Inspired by best practices in Linux package management and system services.
* Thanks to the open source community for tooling and libraries.

---
