package updater

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	OS   = runtime.GOOS
	ARCH = runtime.GOARCH
)

// downloadFile Downloads file from the given url and saves to filepath
func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// buildFileName builds the correct file name for your Operating system and architecture
func buildFilename() string {
	caser := cases.Title(language.English)
	os := caser.String(OS)

	arch := parseArch(ARCH)

	return fmt.Sprintf("lsh_%s_%s", os, arch)
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

// untarFile expects a io.Reader with the file and a path which it will untar to
func untarFile(path string, r io.Reader) error {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {

		case err == io.EOF:
			return nil

		case err != nil:
			return err

		case header == nil:
			continue
		}

		target := filepath.Join(path, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}
		case tar.TypeReg:
			fileDir := filepath.Dir(target)

			if err := os.MkdirAll(fileDir, 0770); err != nil {
				return fmt.Errorf("untar failed while creating folder: %s", err.Error())
			}
			outFile, err := os.Create(target)
			if err != nil {
				return fmt.Errorf("untar failed while creating file: %s", err.Error())
			}
			err = os.Chmod(outFile.Name(), 0755)
			if err != nil {
				return fmt.Errorf("untar failed, permission error: %s", err.Error())
			}
			if _, err := io.Copy(outFile, tr); err != nil {
				return fmt.Errorf("untar failed while copying data to file: %s", err.Error())

			}
			outFile.Close()
		}
	}
}
