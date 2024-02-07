package updater

import (
	"fmt"
	"log"
	"os"

	"github.com/latitudesh/lsh/internal/version"
)

func Update() error {
	log.Println("Update started!")

	release, err := LatestLshRelease()
	if err != nil {
		return err
	}

	if release.TagName == version.Version {
		log.Println()
	}

	tempDir, err := os.MkdirTemp("", "temp")
	if err != nil {
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	filename := buildFilename()
	downloadURL := fmt.Sprintf("https://github.com/latitudesh/lsh/releases/download/%s/%s.tar.gz", release.TagName, filename)
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

	currentBinPath := fmt.Sprintf("%s/.lsh/lsh", homeDir)
	newBinPath := fmt.Sprintf("%s/%s/lsh", tempDir, filename)
	oldBinPath := fmt.Sprintf("%s/.lsh/lsh-old", homeDir)

	// Rename current lsh binary to lsh-old
	err = os.Rename(currentBinPath, oldBinPath)
	if err != nil {
		return err
	}

	// Rename and move downloaded lsh binary
	os.Rename(newBinPath, currentBinPath)
	if err != nil {
		return err
	}

	// Remove old binary
	os.Remove(oldBinPath)
	if err != nil {
		return err
	}

	os.RemoveAll(tempDir)
	log.Println("Update finished successfully!")

	return nil
}
