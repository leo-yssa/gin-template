package dto

import "mime/multipart"

type Announcement struct {
	File *multipart.FileHeader `form:"file"`
}