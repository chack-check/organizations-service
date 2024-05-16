package files

import "github.com/chack-check/organizations-service/domain/files/models"

type FilesPort interface {
	ValidateUploadingFile(file models.UploadingFile) bool
	SaveFile(file models.UploadingFile) (*models.SavedFile, error)
}
