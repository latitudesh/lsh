package updater

import (
	"fmt"
	"log"
	"os"
)

// Update expects a version string and updates the cli to it
func Update(version string) error {
	log.Println("Update started!")

	tempDir, err := os.MkdirTemp("", "temp")
	if err != nil {
		return err
	}

	filename := buildFilename()
	downloadURL := fmt.Sprintf("https://github.com/latitudesh/lsh/releases/download/%s/%s.tar.gz", version, filename)
	filePath := fmt.Sprintf("%s/%s.tar.gz", tempDir, filename)

	log.Println("Downloading new version...")
	err = downloadFile(filePath, downloadURL)
	if err != nil {
		return err
	}
	log.Println("Download finished successfully!")

	f, err := os.Open(fmt.Sprintf("%s/%s.tar.gz", tempDir, filename))
	if err != nil {
		return err
	}

	err = untarFile(tempDir, f)
	if err != nil {
		return err
	}

	currentExecPath, err := os.Executable()
	if err != nil {
		return err
	}

	oldExecPath := fmt.Sprintf("%s-old", currentExecPath)
	newExecPath := fmt.Sprintf("%s/%s/lsh", tempDir, filename)

	// Rename current lsh binary to lsh-old
	err = os.Rename(currentExecPath, oldExecPath)
	if err != nil {
		return err
	}

	// Rename and move downloaded lsh binary
	os.Rename(newExecPath, currentExecPath)
	if err != nil {
		return err
	}

	// Remove old binary
	os.Remove(oldExecPath)
	if err != nil {
		return err
	}

	os.RemoveAll(tempDir)
	log.Println("Update finished successfully!")

	return nil
}
