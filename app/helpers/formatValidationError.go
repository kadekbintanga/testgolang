package helpers

import(
	"strings"
	"fmt"
)

func FormatValidationErrorInput(err error) string{
	fmt.Println("Input Validation Error")
	if strings.Contains(err.Error(), "Author"){
		return "Author is required"
	}else if strings.Contains(err.Error(), "Title"){
		return "Title is required"
	}else if strings.Contains(err.Error(), "Body"){
		return "Body is required"
	}else if strings.Contains(err.Error(), "unmarshal"){
		if  strings.Contains(err.Error(), "InputArticle.author"){
			return "author must has string type"
		}else if strings.Contains(err.Error(), "InputArticle.title"){
			return "title must has string type"
		}else if strings.Contains(err.Error(), "InputArticle.body"){
			return "body must has string type"
		}else{
			return "Something went wrong"
		}
	}else{
		return "Something went wrong"
	}
}

func FormatValidationErrorSql(err error) string{
	fmt.Println("SQL Validation Error")
	if strings.Contains(err.Error(), "Duplicate"){
		if strings.Contains(err.Error(), "articles.title"){
			return "Title has been used"
		}else{
			return "Something went wrong"
		}
	}else{
		return "Something went wrong"
	}
}