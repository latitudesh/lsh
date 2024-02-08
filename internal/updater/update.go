package updater

import (
	"fmt"
	"log"
	"os"
)

// Update expects a version string and updates the cli to it
func Update(version string) error {
	log.Println("Update started!")

	updateFile := NewUpdateFile()

	tempDir, err := os.MkdirTemp("", "temp")
	if err != nil {
		return err
	}
	DownloadPath := fmt.Sprintf("%s/%s", tempDir, updateFile.Name)

	log.Println("Downloading new version...")
	err = downloadFile(DownloadPath, updateFile.Url(version))
	if err != nil {
		return err
	}
	log.Println("Download finished successfully!")

	f, err := os.Open(fmt.Sprintf("%s/%s", tempDir, updateFile.Name))
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
	newExecPath := fmt.Sprintf("%s/%s/lsh", tempDir, updateFile.Dir)

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