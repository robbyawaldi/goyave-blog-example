package article

import (
	"net/http"

	"github.com/System-Glitch/goyave-blog-example/database/model"
	"github.com/System-Glitch/goyave/v3"
	"github.com/System-Glitch/goyave/v3/database"
)

// Index list all articles.
func Index(response *goyave.Response, request *goyave.Request) { // TODO paginate
	articles := []model.Article{}
	result := database.Conn().Find(&articles)
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, articles)
	}
}

// Show a single article.
func Show(response *goyave.Response, request *goyave.Request) {
	article := model.Article{}
	result := database.Conn().Where("slug = ?", request.Params["slug"]).First(&article)
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, article)
	}
}

// Store a new article.
func Store(response *goyave.Response, request *goyave.Request) {
	article := model.Article{
		Title:    request.String("title"),
		Contents: request.String("contents"),
		AuthorID: request.User.(*model.User).ID,
	}
	if err := database.Conn().Create(&article).Error; err != nil {
		response.Error(err)
	} else {
		response.JSON(http.StatusCreated, map[string]interface{}{
			"id":   article.ID,
			"slug": article.Slug,
		})
	}
}

// Update an existing article. Only the author of the article can do that.
func Update(response *goyave.Response, request *goyave.Request) {
	article := model.Article{}
	db := database.Conn()
	result := db.Select("id")
	if slug, ok := request.Params["slug"]; ok {
		result = result.Where("slug = ?", slug).First(&article)
	} else {
		result = result.First(&article, request.Params["id"])
	}
	if response.HandleDatabaseError(result) {
		if err := db.Model(&article).Updates(request.Data).Error; err != nil {
			response.Error(err)
		}
	}
}

// Destroy an existing article. Only the author of the article can do that.
func Destroy(response *goyave.Response, request *goyave.Request) {
	article := model.Article{}
	db := database.Conn()
	result := db.Select("id")
	if slug, ok := request.Params["slug"]; ok {
		result = result.Where("slug = ?", slug).First(&article)
	} else {
		result = result.First(&article, request.Params["id"])
	}
	// TODO allow search by id too for patch and delete
	if response.HandleDatabaseError(result) {
		if err := db.Delete(&article).Error; err != nil {
			response.Error(err)
		}
	}
}

// TODO article tests