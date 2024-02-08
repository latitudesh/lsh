package updater

import (
	"log"
	"os"
	"path/filepath"

	"github.com/latitudesh/lsh/internal/utils"
)

// Update expects a version string and updates the cli to it
func Update(version string) error {
	log.Println("Update started!")

	updateFile := NewUpdateFile()

	tempDir, err := os.MkdirTemp("", "temp")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)

	destinationPath := filepath.Join(tempDir, updateFile.Name)

	log.Println("Downloading new version...")
	err = updateFile.Download(destinationPath, version)
	if err != nil {
		return err
	}
	log.Println("Download finished successfully!")

	err = decompress(tempDir, updateFile)
	if err != nil {
		return err
	}

	currentExecPath, err := os.Executable()
	if err != nil {
		return err
	}
	oldExecPath := currentExecPath + "-old"
	newExecPath := filepath.Join(tempDir, updateFile.Dir, updateFile.Executable)

	// Rename current lsh binary to lsh-old
	err = os.Rename(currentExecPath, oldExecPath)
	if err != nil {
		return err
	}

	// Rename and move downloaded lsh binary
	err = os.Rename(newExecPath, currentExecPath)
	if err != nil {
		return err
	}

	// Remove old binary (linux|macOs only)
	if OS != "windows" {
		err = os.Remove(oldExecPath)
		if err != nil {
			return err
		}
	}

	log.Println("Update finished successfully!")

	return nil
}

func decompress(destination string, file *UpdateFile) error {
	source := filepath.Join(destination, file.Name)
	if file.Extension == ".tar.gz" {
		err := utils.Untar(source, destination)
		if err != nil {
			return err
		}

		return nil
	}

	err := utils.Unzip(source, destination)
	if err != nil {
		return err
	}
	return nil
}
