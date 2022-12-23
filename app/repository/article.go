package repository

import(
	"testgolang/app/models"
	"testgolang/config"
)

type ArticleRepository interface{

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