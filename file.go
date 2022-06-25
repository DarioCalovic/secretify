package secretify

const MaxFileSize = 5242880 // 5 MB
// const AllowedFileExtensions = "jpeg,jpg,gif,png,eps,raw,tif,tiff,bmp,log,txt,key,pem,csv" // comma separated
const AllowedFileExtensions = "" // comma separated

type File struct {
	Base
	Identifier string `gorm:"unique"`
	Path       string
	Name       string
	Type       string
	Size       uint
}
