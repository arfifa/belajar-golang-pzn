package helper

import (
	"arfifa/belajar-golang-restful-api/model/domain"
	"arfifa/belajar-golang-restful-api/model/web"
)

func ToCategoryReponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id: category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryReponse(category))
	}

	return categoryResponses
}