package Controllers

import (
	m "Users/Models"
	s "Users/Services"
	"fmt"
	"github.com/gin-gonic/gin"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

type CategoryController struct {
	HeadController
	TemplateHtml
}

func (cat *CategoryController) ShowListCategory(c *gin.Context) {
	url := s.URL{}
	if !cat.isAuth(c.Request) {
		c.Redirect(http.StatusFound, url.Get404())
		return
	}
	if !cat.isAccess(c.Request) {
		c.Redirect(http.StatusFound, url.Get404())
		return
	}
	cat.ShowHead(c, "Список категорий")

	categoryService := s.Categoryzator{}
	viewCategoryModel := categoryService.GetViewListCategores()

	viewListCategory := m.ViewListCategory{
		UrlJs:                  url.GetFullPathJs(),
		UrlCss:                 url.GetFullPathCss(),
		UrlImgIco:              url.GetFullPathImg(),
		CountIco:               1,
		UrlAddPhoto:            url.GetHostNameWithProtocol() + "/listcategores/uploadimg",
		UrlRemovePhoto:         url.GetHostNameWithProtocol() + "/listcategores/removefile",
		ListFirstLavelCategory: viewCategoryModel.ListFirstLavelCategory,
	}

	cat.showHtmlWithoutHeaderRR(c.Writer, viewListCategory, "categoresList")
}

func (cat *CategoryController) EditCategory(c *gin.Context) {
	idStr := c.Request.FormValue("id") //c.Params.ByName("id")
	id, errConvert := strconv.Atoi(idStr)

	fmt.Println(id)
	if errConvert != nil {
		c.JSON(200, gin.H{"status": "", "error": "Некорректные данные клиента"})
	}
	name := c.Request.FormValue("name") //c.Params.ByName("name")
	categoryService := s.Categoryzator{}
	categoryModel := m.UpdateFirstCategoryModel{Id: id, Name: name}
	categoryService.UpdateFirstCategory(categoryModel)

	c.JSON(200, gin.H{"status": "Ok"})
}

func (cat *CategoryController) UploadAjaxImg(c *gin.Context) {
	idCategory := c.Params.ByName("idCat")
	id, errConvert := strconv.Atoi(idCategory)

	if errConvert != nil || id <= 0 {
		return
	}

	cat.putImageInDirectory(c, "Categores", id)
}

func (cat *CategoryController) GetUploadIconJson(c *gin.Context) {
	idCategoryValue := c.Params.ByName("idCat")
	idCategory, _ := strconv.Atoi(idCategoryValue)
	categorizator := s.Categoryzator{}
	urlIcon := categorizator.GetIcoUrlByIdCategory(idCategory)
	c.JSON(200, gin.H{"path": urlIcon})
}

func (cat *CategoryController) putImageInDirectory(c *gin.Context, path string, idCategory int) {
	var (
		status int
		err    error
	)
	fmt.Printf("0 \n") //TODO
	defer func() {
		if nil != err {
			http.Error(c.Writer, err.Error(), status)
		}
	}()
	// parse request
	const _24K = (1 << 10) * 24
	if err = c.Request.ParseMultipartForm(_24K); nil != err {
		status = http.StatusInternalServerError
		return
	}

	for _, fheaders := range c.Request.MultipartForm.File {
		for _, hdr := range fheaders {
			var infile multipart.File
			if infile, err = hdr.Open(); nil != err {
				status = http.StatusInternalServerError
				return
			}
			defer infile.Close()
			sizeFile, _ := infile.Seek(0, 2)
			if sizeFile > (1 << 24) {
				fmt.Printf("Размер файла большой!")
				c.JSON(200, gin.H{"status": "", "error": "Размер файла слишком большой"})
				return
			}

			infile.Seek(0, 0)
			if isNotValidSize(hdr) {
				fmt.Printf("Size файла большой! \n")
				c.JSON(200, gin.H{"status": "", "error": "Разрешение изображения слишком большое"})
				return
			}
			fmt.Printf("Прошло валидацию!")
			// open destination
			var outfile *os.File
			var pathImgCategory = "./Img/" + path + "/"
			var nameIcoFile = strconv.Itoa(idCategory) + ".jpg"
			if outfile, err = os.Create(pathImgCategory + nameIcoFile); nil != err {
				status = http.StatusInternalServerError
				return
			}
			if _, err = io.Copy(outfile, infile); nil != err {
				status = http.StatusInternalServerError
				return
			}
			categorizator := s.Categoryzator{}
			categorizator.UpdateIcoNameInCategory(idCategory, nameIcoFile)
			defer outfile.Close()
			c.JSON(200, gin.H{"status": "Ok", "error": ""})
		}
	}
}

func (cat *CategoryController) RemoveImg(c *gin.Context) {
	idVal := c.Params.ByName("idImg")
	idCategory, _ := strconv.Atoi(idVal)

	if idCategory <= 0 {
		return
	}

	var err error
	categorizator := s.Categoryzator{}
	var nameIcoCategory = categorizator.GetIcoFilePathByIdCategory(idCategory)
	err = os.Remove(nameIcoCategory)

	if err != nil {
		return
	}

	categorizator.UpdateIcoNameInCategory(idCategory, "")
}

func isNotValidSize(r *multipart.FileHeader) bool {
	infile, _ := r.Open()
	im, _, err := image.DecodeConfig(infile)
	defer infile.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}
	if im.Height < 10 || im.Height > 500 || im.Width < 10 || im.Width > 500 {
		return true
	}
	return false
}

func (cat *CategoryController) ShowFileUploadPage(c *gin.Context) {
	url := s.URL{}
	cat.ShowHead(c, "Загрузка файла")
	viewModel := m.ViewFileUploadPage{
		IdCat:  1,
		UrlJs:  url.GetFullPathJs(),
		UrlCss: url.GetFullPathCss(),
	}
	cat.showHtmlWithoutHeader(c.Writer, viewModel, "fileUpload")
}

func (cat *CategoryController) isAccess(req *http.Request) bool {
	userService := s.Userator{}
	user := userService.GetByRequest(req)
	idUser := user.Id
	userRoleService := s.UserRoler{}

	if userRoleService.IsAdmin(idUser) || userRoleService.IsOperator(idUser) {
		return true
	}

	return false
}

func (cat *CategoryController) isAuth(req *http.Request) bool {
	auth := s.Authorithator{}
	var isAuth = auth.IsAuth(req)

	if isAuth {
		return true
	}

	return false
}
