package Models

type ViewListCategory struct {
	ListFirstLavelCategory []*ViewFirstCategory
	UrlJs                  string
	UrlCss                 string
	UrlImgIco              string
	CountIco               int
	UrlAddPhoto            string
	UrlRemovePhoto         string
	ViewFileUploadPage     ViewFileUploadPage
}
