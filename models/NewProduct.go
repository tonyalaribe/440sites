package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/structs"
	"github.com/tonyalaribe/shop440/features/products"
)

func NewProduct(product products.Product) error {
	m := structs.New(product).Map()
	m["date"] = product.DateCreated.Format("2006-01-02")

	fileContentByte := []byte{}
	jsonFrontmatterByte, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		log.Println(err)
		return err
	}

	fileContentByte = append(fileContentByte, jsonFrontmatterByte...)
	fileContentByte = append(fileContentByte, []byte("\n")...)
	fileContentByte = append(fileContentByte, []byte(product.Description)...)

	filePath := product.Slug
	err = ioutil.WriteFile(filePath, fileContentByte, os.ModePerm)
	if err != nil {
		log.Println(err)
		return err
	}

	err = Command("hugo", CommandSiteDirFn)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
