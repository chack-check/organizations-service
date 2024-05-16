package files

import "github.com/chack-check/organizations-service/domain/files/models"

type MockFilesAdapter struct{}

func (adapter MockFilesAdapter) ValidateUploadingFile(file models.UploadingFile) bool {
	return true
}

func (adapter MockFilesAdapter) SaveFile(file models.UploadingFile) (*models.SavedFile, error) {
	original := file.GetOriginal()
	var convertedUrl *string
	var convertedFilename *string
	if converted := file.GetConverted(); converted != nil {
		url := converted.GetUrl()
		filename := converted.GetFilename()
		convertedUrl = &url
		convertedFilename = &filename
	}

	savedFile := models.NewSavedFile(original.GetUrl(), original.GetFilename(), convertedUrl, convertedFilename)
	return &savedFile, nil
}
