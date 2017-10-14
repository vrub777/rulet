package Run

import (
	Controller "Users/Controllers"
	"github.com/gin-gonic/gin"
)

type AppService struct {
}

func (s *AppService) Run() {
	router := gin.Default()
	router.Static("/img", "./Img")

	ac := Controller.AuthController{}
	regc := Controller.RegistrationController{}
	listUsersc := Controller.ListUsersController{}
	errorc := Controller.ErrorsController{}
	userc := Controller.UserController{}
	categoryc := Controller.CategoryController{}
	//headc := Controller.HeadController{}

	router.GET("/auth", ac.ShowSimpleAuthUser)
	router.GET("/auth/:backURL", ac.ShowSimpleAuthUser)
	router.POST("/auth", ac.PostSimpleAuthUser)
	router.POST("/auth/:backURL", ac.PostSimpleAuthUser)
	router.POST("/authOut", ac.OutAuthUser)
	router.GET("/reg", regc.ShowPageRegistration)
	router.POST("/reg", regc.PostPageRegistration)

	router.GET("/listusers", listUsersc.ShowPage)
	router.POST("/listusers/:action/:id", listUsersc.PostAction)
	router.GET("/addUser", userc.ShowAddUser)
	router.POST("/addUser", userc.PostAddUser)
	router.GET("/editUser/:id", userc.ShowEditUser)
	router.POST("/editUser/:id", userc.PostEditUser)

	router.GET("/listcategores", categoryc.ShowListCategory)
	router.GET("/showfileupload", categoryc.ShowFileUploadPage)
	//router.POST("/listcategores/edit-first/:id/:name", categoryc.EditCategory)
	router.POST("/listcategores/edit-first", categoryc.EditCategory)
	//router.POST("/listcategores/uploadimg", categoryc.UploadImg)
	router.POST("/listcategores/geticon/:idCat", categoryc.GetUploadIconJson)
	router.POST("/listcategores/uploadajaximg/:idCat", categoryc.UploadAjaxImg)
	router.POST("/listcategores/removeimgcategory/:idImg", categoryc.RemoveImg)

	//router.GET("/blockhead", headc.ShowByHttpHead)
	router.NotFound404(errorc.E404)

	router.Run(":8090")
}
