package service

import "errors"

var ErrProductsNotFound = errors.New("products were not found")
var ErrEnoughtProducts = errors.New("not enought products to compare, provide at lease 2 products")
var ErrEnoughtProductsNotFound = errors.New("not enought products were not found")
var ErrMaxProductComparison = errors.New("cannot compare more than 10 items")
var ErrNoIds = errors.New("no products id were given")
var ErrEmptyIds = errors.New("ids cannot contain empty values")
var ErrEmptyFields = errors.New("fields cannot contain empty values")
var ErrIdsDuplicated = errors.New("products id are duplicated")
