package helpers

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
)

func ZipFolder(folderPath string) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create a zip file header
		zipFileHeader, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Set the name of the file within the zip archive
		zipFileHeader.Name, err = filepath.Rel(folderPath, path)
		if err != nil {
			return err
		}

		// Write the file header to the zip archive
		writer, err := zipWriter.CreateHeader(zipFileHeader)
		if err != nil {
			return err
		}

		// If the current file is a directory, don't write anything, just create the directory in the archive
		if info.IsDir() {
			return nil
		}

		// Open the file for reading
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Copy the file contents to the zip archive
		_, err = io.Copy(writer, file)
		return err
	})

	if err != nil {
		return nil, err
	}

	// Close the zip writer
	err = zipWriter.Close()
	if err != nil {
		return nil, err
	}

	return &buf, nil
}