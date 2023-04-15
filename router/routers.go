package router

import (
	"myGram/controllers"
	"myGram/docs"
	_ "myGram/docs"
	"myGram/middleware"
	"myGram/repository"
	"myGram/service"
	"os"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @title myGram API Documentation
// @version 1.0.0
// @description This is the documentation for myGram API
// @termsOfService http://swagger.io/terms/
// @contact.name myGram API Support
// @contact.url http://swagger.io/support
// @license.name Apache 2.0
func StartRouter(r *gin.Engine, db *gorm.DB) {

	repoUser := repository.NewUserRepository(db)
	srcUser := service.NewUserService(repoUser)
	ctrlUser := controllers.NewUserController(srcUser)
	// User
	userRouter := r.Group("/user")
	{
		// Register User
		userRouter.POST("/register", ctrlUser.Register)
		// Login User
		userRouter.POST("/login", ctrlUser.Login)
	}
	repoPhoto := repository.NewPhotoRepository(db)
	srcPhoto := service.NewPhotoService(repoPhoto)
	ctrlPhoto := controllers.NewPhotoController(srcPhoto)
	// Photo
	photoRouter := r.Group("/photo")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.GET("/", ctrlPhoto.GetAllPhotos)
		photoRouter.GET("/:photoId", ctrlPhoto.GetPhotoByID)
		photoRouter.POST("/", ctrlPhoto.CreatePhoto)
		photoRouter.PUT("/:photoId", middleware.PhotoAuthorization(), ctrlPhoto.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.PhotoAuthorization(), ctrlPhoto.DeletePhoto)
	}

	repoComment := repository.NewCommentRepository(db)
	srcComment := service.NewCommentService(repoComment)
	ctrlComment := controllers.NewCommentController(srcComment)
	// Comment
	commentRouter := r.Group("/comment")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.GET("/", ctrlComment.GetAllComments)
		commentRouter.GET("/:commentId", ctrlComment.GetCommentByID)
		commentRouter.POST("/", ctrlComment.CreateComment)
		commentRouter.PUT("/:commentId", middleware.CommentAuthorization(), ctrlComment.UpdateComment)
		commentRouter.DELETE("/:commentId", middleware.CommentAuthorization(), ctrlComment.DeleteComment)
	}
	repoSosmed := repository.NewSosmedRepository(db)
	srcSosmed := service.NewSosmedService(repoSosmed)
	ctrlSosmed := controllers.NewSosmedController(srcSosmed)
	// Social Media
	socialMediaRouter := r.Group("/socialmedia")
	{
		socialMediaRouter.Use(middleware.Authentication())
		socialMediaRouter.GET("/", ctrlSosmed.GetAllSosmed)
		socialMediaRouter.GET("/:socialMediaId", ctrlSosmed.GetSosmedByID)
		socialMediaRouter.POST("/", ctrlSosmed.CreateSosmed)
		socialMediaRouter.PUT("/:socialMediaId", middleware.SocialMediaAuthorization(), ctrlSosmed.UpdateSosmed)
		socialMediaRouter.DELETE("/:socialMediaId", middleware.SocialMediaAuthorization(), ctrlSosmed.DeleteSosmed)
	}
	docs.SwaggerInfo.Host = "mygram-golang-production-988c.up.railway.app:" + os.Getenv("PORT")
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	// routing docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
