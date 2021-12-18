package main

import (
	"context"
	"fmt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	sheet("googleDocApi/demo/inner-geography-333815-6af830f78885.json", "1QhH_SGNZnJFBEQT0U4dPqmPJMKKRcXVQnSO3FR5cfCM")
}
func dr(path, id string) {
	ctx := context.Background()
	docsvc, err := drive.NewService(ctx, option.WithCredentialsFile(path))
	if err != nil {
		fmt.Println(err)
		return
	}
	d := drive.NewDrivesService(docsvc)
	document, err := d.Get(id).Do()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(document.Kind)
}

// path 凭证文件路径  id  google sheet表格id
//https://docs.google.com/document/d/195j9eDD3ccgjQRttHhJPymLJUCOUjs-jmwTrekvdjFE/edit
func sheet(path, id string) {
	ctx := context.Background()
	docsvc, err := sheets.NewService(ctx, option.WithCredentialsFile(path))
	if err != nil {
		fmt.Println(err)
		return
	}
	d := sheets.NewSpreadsheetsValuesService(docsvc)
	document, err := d.Get(id, "A1:Z200").Do()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(document.Values)
}
