package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Untar expects a io.Reader with the file and a path which it will untar to
func Untar(path string, r io.Reader) error {
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
