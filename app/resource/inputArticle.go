package resource

type InputArticle struct{
	Author		string		`json:"author" binding:"required"`
	Title		string		`json:"title" binding:"required"`
	Body		string		`json:"body" binding:"required"`
}