package repository

import(
	"testgolang/app/models"
	"testgolang/config"
	"strings"
)

type ArticleRepository interface{
	CreateArticle(Article models.Article)(models.Article, error)
	GetArticle(author string, search string)([]models.Article, error)
}

func NewArticleRepository() ArticleRepository{
	return &dbConnection{
		connection: config.ConnectDB(),
	}
}

func (db *dbConnection) CreateArticle(Article models.Article)(models.Article, error){
	err := db.connection.Save(&Article).Error
	if err != nil {
		return Article, err
	}
	return Article, nil
}

func (db *dbConnection) GetArticle(author string, search string)([]models.Article, error){
	var Article []models.Article
	connection := db.connection.Debug().Model(&Article)
	if author != ""{
		connection = connection.Where("author = ?", author)
	}
	if search != ""{
		search = "%" + strings.ToLower(search) + "%"
		connection = connection.Where("( LOWER(title) LIKE ? OR LOWER(body) LIKE ? )", search,search)
	}
	connection = connection.Order("created_at desc").Find(&Article)
	err := connection.Error
	if err != nil {
		return Article, err
	}
	return Article, nil
}