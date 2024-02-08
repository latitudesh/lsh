package updater

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	OS   = runtime.GOOS
	ARCH = runtime.GOARCH
)

type UpdateFile struct {
	Name      string
	Dir       string
	Extension string
}

func NewUpdateFile() *UpdateFile {
	dir := buildDirName()
	ext := buildExtesion(OS)
	return &UpdateFile{
		Name:      dir + ext,
		Dir:       dir,
		Extension: ext,
	}
}

func (uf *UpdateFile) Url(version string) string {
	return fmt.Sprintf("https://github.com/latitudesh/lsh/releases/download/%s/%s", version, uf.Name)
}

func (uf *UpdateFile) Download(dstPath, version string) error {
	resp, err := http.Get(uf.Url(version))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// buildDirName builds the correct file name for your Operating system and architecture
func buildDirName() string {
	caser := cases.Title(language.English)
	os := caser.String(OS)

	arch := parseArch(ARCH)

	return fmt.Sprintf("lsh_%s_%s", os, arch)
}

func buildExtesion(os string) string {
	if os == "windows" {
		return ".zip"
	}
	return ".tar.gz"
}

// parseArch expects a GOARCH and returns the arch in our github release format
func parseArch(goArch string) string {
	if goArch == "amd64" {
		return "x86_64"
	}
	if goArch == "386" {
		return "i386"
	}

	return goArch
}
