package out

import "mime/multipart"

type MinIOPort interface {
	UploadImage(file multipart.File, fileHeader *multipart.FileHeader) (string, error)
}