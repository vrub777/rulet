package Services

type URL struct {
}

const (
	CategoryIcons        = "/Categores"
	HostName             = "localhost:8090"
	HostStaticName       = "localhost:8091"
	Img                  = "/Img"
	Protocol             = "http://"
	PathStaticJs         = "/js"
	PathStaticCss        = "/css"
	PathStaticImg        = "/image"
	PathStaticIcoCatalog = "/ico-catalog"
	Path404              = "/404"
)

func (url *URL) GetHostNameWithProtocol() string {
	return Protocol + HostName
}

func (url *URL) GetHostStaticFull() string {
	return Protocol + HostStaticName
}

func (url *URL) GetFullPathJs() string {
	return url.GetHostStaticFull() + PathStaticJs
}

func (url *URL) GetFullPathCss() string {
	return url.GetHostStaticFull() + PathStaticCss
}

func (url *URL) GetFullPathImg() string {
	return url.GetHostStaticFull() + PathStaticImg
}

func (url *URL) GetFullPathIcoCatalog() string {
	return url.GetFullPathImg() + PathStaticIcoCatalog
}

func (url *URL) Get404() string {
	return url.GetHostNameWithProtocol() + Path404
}

func (url *URL) GetPathCatalogIcons() string {
	return url.GetHostNameWithProtocol() + Img + CategoryIcons
}
