package models

import (
	"testing"
)

func TestCreateSavedFileWithoutConverted(t *testing.T) {
	file := NewSavedFile(
		"originalUrl",
		"originalFilename",
		nil,
		nil,
	)

	if file.GetOriginalUrl() != "originalUrl" {
		t.Fatalf("Error creating saved file: originalUrl %s != \"%s\"", file.GetOriginalUrl(), "originalUrl")
	}
	if file.GetOriginalFilename() != "originalFilename" {
		t.Fatalf("Error creating saved file: originalFilename %s != \"%s\"", file.GetOriginalFilename(), "originalFilename")
	}
	if file.GetConvertedUrl() != nil {
		t.Fatalf("Error creating saved file: convertedUrl %v != %v", file.GetConvertedUrl(), nil)
	}
	if file.GetConvertedFilename() != nil {
		t.Fatalf("Error creating saved file: convertedFilename %v != %v", file.GetConvertedFilename(), nil)
	}
}

func TestCreateSavedFileWithConverted(t *testing.T) {
	var convertedUrl = "convertedUrl"
	var convertedFilename = "convertedFilename"

	file := NewSavedFile(
		"originalUrl",
		"originalFilename",
		&convertedUrl,
		&convertedFilename,
	)

	if file.GetOriginalUrl() != "originalUrl" {
		t.Fatalf("Error creating saved file: originalUrl %s != \"%s\"", file.GetOriginalUrl(), "originalUrl")
	}
	if file.GetOriginalFilename() != "originalFilename" {
		t.Fatalf("Error creating saved file: originalFilename %s != \"%s\"", file.GetOriginalFilename(), "originalFilename")
	}
	if file.GetConvertedUrl() != &convertedUrl {
		t.Fatalf("Error creating saved file: convertedUrl %v != %v", file.GetConvertedUrl(), &convertedUrl)
	}
	if file.GetConvertedFilename() != &convertedFilename {
		t.Fatalf("Error creating saved file: convertedFilename %v != %v", file.GetConvertedFilename(), &convertedFilename)
	}
}
