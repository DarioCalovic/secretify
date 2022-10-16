package secretify

// MaxFileSize default max file size
const MaxFileSize = 5242880 // 5 MB
// AllowedFileExtensions default allowed file extensions
const AllowedFileExtensions = "" // comma separated, e.g. "jpeg,jpg,gif,png,eps,raw,tif,tiff,bmp,log,txt,key,pem,csv"

type File struct {
	Base
	Identifier string `gorm:"unique"`
	Path       string
	Name       string
	Type       string
	Size       uint
}
