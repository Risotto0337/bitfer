# Bitfer

Bitfer is a cross-distro, source-based package manager for Linux systems.

It is designed to work on:
- Slackware
- openSUSE
- Arch Linux
- Debian / Ubuntu
- Fedora
- and other Linux distributions with build tools

Bitfer installs software by:
- fetching package recipes from GitHub
- downloading source code
- building it locally
- installing into /opt/bitfer

---

# 🚀 Features

- Install packages from GitHub repositories
- Build software from source automatically
- Works across multiple Linux distributions
- Multi-repository support
- Simple CLI interface
- YAML-based build recipes
- Easy Git-based updates

---

# ⚙️ How Bitfer Works

bitfer install mpv
        ↓
load repository index.json
        ↓
find package recipe (bitfer.yaml)
        ↓
download source code
        ↓
run build steps
        ↓
install to /opt/bitfer

---

# 💻 Usage

## Install a package

bitfer install mpv

---

## Remove a package

bitfer remove mpv

---

# 📦 Example Packages

Bitfer can build:

- mpv
- ffmpeg
- firefox
- wine
- qemu
- mesa
- rust
- pipewire
- supertuxkart
- fzf
- ripgrep
- bat

---

# 🧱 Installation

## Method 1 (recommended)

curl -L https://raw.githubusercontent.com/Risotto0337/bitfer/main/install.sh | bash

---

## Method 2 (manual)

git clone https://github.com/Risotto0337/bitfer.git
cd bitfer
go build -o bitfer cmd/bitfer/main.go
sudo mv bitfer /usr/local/bin/

---

# 🌐 Supported Linux Distributions

Works on any Linux system with build tools.

Tested / Targeted:
- Slackware Linux
- openSUSE

Expected:
- Arch Linux
- Debian / Ubuntu
- Fedora
- Gentoo
- Alpine Linux

---

# 🧰 Requirements

- gcc / g++
- make
- git
- curl
- bash
- tar / gzip
- python3
- cmake / meson (for some packages)

---

# 📁 Repository Format

index.json
package/
  └── bitfer.yaml

---

## Example index.json

{
  "mpv": {
    "name": "mpv",
    "path": "mpv/bitfer.yaml"
  },
  "ffmpeg": {
    "name": "ffmpeg",
    "path": "ffmpeg/bitfer.yaml"
  }
}

---

## Example bitfer.yaml

name: mpv

source:
  url: https://github.com/mpv-player/mpv/archive/refs/tags/v0.38.0.tar.gz

build:
  steps:
    - ./bootstrap.py
    - ./waf configure
    - ./waf build
    - ./waf install DESTDIR=$DESTDIR

---

# 🔄 Updating Repo

git add .
git commit -m "update"
git push

---

# 🧪 Status

- Core system: working
- GitHub repos: working
- Multi-repo support: working
- Dependency solver: not yet
- Binary cache: not yet

---

# ⚠️ Notes

- Builds from source
- Requires build tools
- Not a binary package manager like apt or zypper

---

# 🤝 Contributing

Add packages by creating:
- index.json entry
- bitfer.yaml recipe

---

