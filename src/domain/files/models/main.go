package models

type SavedFile struct {
	originalUrl       string
	originalFilename  string
	convertedUrl      *string
	convertedFilename *string
}

func (file SavedFile) GetOriginalUrl() string {
	return file.originalUrl
}

func (file SavedFile) GetOriginalFilename() string {
	return file.originalFilename
}

func (file SavedFile) GetConvertedUrl() *string {
	return file.convertedUrl
}

func (file SavedFile) GetConvertedFilename() *string {
	return file.convertedFilename
}

type UploadingFileMeta struct {
	url            string
	filename       string
	signature      string
	systemFiletype string
}

func (meta *UploadingFileMeta) GetUrl() string {
	return meta.url
}

func (meta *UploadingFileMeta) GetFilename() string {
	return meta.filename
}

func (meta *UploadingFileMeta) GetSignature() string {
	return meta.signature
}

func (meta *UploadingFileMeta) GetSystemFiletype() string {
	return meta.systemFiletype
}

type UploadingFile struct {
	original  UploadingFileMeta
	converted *UploadingFileMeta
}

func (file *UploadingFile) GetOriginal() UploadingFileMeta {
	return file.original
}

func (file *UploadingFile) GetConverted() *UploadingFileMeta {
	return file.converted
}

func NewUploadingFileMeta(url string, filename string, signature string, systemFiletype string) UploadingFileMeta {
	return UploadingFileMeta{
		url:            url,
		filename:       filename,
		signature:      signature,
		systemFiletype: systemFiletype,
	}
}

func NewUplaodingFile(original UploadingFileMeta, converted *UploadingFileMeta) UploadingFile {
	return UploadingFile{
		original:  original,
		converted: converted,
	}
}

func NewSavedFile(originalUrl string, originalFilename string, convertedUrl *string, convertedFilename *string) SavedFile {
	return SavedFile{
		originalUrl:       originalUrl,
		originalFilename:  originalFilename,
		convertedUrl:      convertedUrl,
		convertedFilename: convertedFilename,
	}
}
