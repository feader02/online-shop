package api

import (
	"fmt"
	"github.com/feader02/online-shop/internal/entities"
	"github.com/feader02/online-shop/internal/utils"
	"strconv"
	"strings"
)

func (a *App) GetProducts(pageNum int, pageSize int, search string, prType string, priceRadius string) ([]entities.Product, error) {
	var products []entities.Product
	queryArgs := make([]interface{}, 0)

	offset := (pageNum - 1) * pageSize
	offsetSQL := fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)

	searchSQL := ""
	if search != "" {
		trigrams := utils.CreateTrigrams(search)
		searchSQL = "WHERE ("
		for i, trigram := range trigrams {
			if i > 0 {
				searchSQL += " OR "
			}
			searchSQL += "(name LIKE ? OR description LIKE ?)"
			queryArgs = append(queryArgs, "%"+trigram+"%", "%"+trigram+"%")
		}
		searchSQL += ")"
	}

	typeSQL := ""
	if prType != "" {
		prTypeSlice := strings.Split(prType, ",")
		if searchSQL == "" {
			typeSQL = "WHERE ("
		} else {
			typeSQL = " AND ("
		}
		for i, v := range prTypeSlice {
			queryArgs = append(queryArgs, v)
			if i == 0 {
				typeSQL += "type = ?"
			} else {
				typeSQL += " OR type = ?"
			}
		}
		typeSQL += ")"
	}

	priceRadiusSQL := ""
	if priceRadius != "" {
		priceRange := strings.Split(priceRadius, "-")
		if len(priceRange) == 2 {
			minPrice, err := strconv.ParseFloat(priceRange[0], 64)
			if err != nil {
				return nil, err
			}
			maxPrice, err := strconv.ParseFloat(priceRange[1], 64)
			if err != nil {
				return nil, err
			}
			queryArgs = append(queryArgs, minPrice, maxPrice)
			if searchSQL != "" || typeSQL != "" {
				priceRadiusSQL = " AND price BETWEEN ? AND ?"
			} else {
				priceRadiusSQL = " WHERE price BETWEEN ? AND ?"
			}
		}
	}

	SQLRequest := searchSQL + typeSQL + priceRadiusSQL + offsetSQL
	err := a.storage.Get(&products, "Product", SQLRequest, queryArgs...)
	if err != nil {
		return nil, err
	}

	return products, nil
}
