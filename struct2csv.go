package gocsv

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func GetHeader(strct interface{}) []string {
	var header []string
	kind := reflect.TypeOf(strct).Kind()
	switch kind {
	case reflect.Slice:
		elem := reflect.ValueOf(strct).Index(0)

		for j := 0; j < elem.NumField(); j++ {
			tag := elem.Type().Field(j).Tag.Get("csv")
			header = append(header, tag)
			fmt.Println(tag)
		}
	case reflect.Struct:
		elem := reflect.ValueOf(strct)
		for j := 0; j < elem.NumField(); j++ {
			tag := elem.Type().Field(j).Tag.Get("csv")
			header = append(header, tag)
			fmt.Println(tag)
		}
	}
	return header
}

func WriteToCSV(fileName string, strct interface{}) error {
	csvFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)

	_ = csvWriter.Write(GetHeader(strct))
	csvWriter.Flush()

	switch reflect.TypeOf(strct).Kind() {
	case reflect.Slice:
		slice := reflect.ValueOf(strct)

		for i := 0; i < slice.Len(); i++ {
			elem := slice.Index(i)

			var str []string

			for i := 0; i < elem.NumField(); i++ {
				value := elem.Field(i)
				switch elem.Field(i).Kind() {
				case reflect.String:
					str = append(str, value.Interface().(string))
				case reflect.Int:
					str = append(str, strconv.Itoa(value.Interface().(int)))
				}

			}
			_ = csvWriter.Write(str)
			csvWriter.Flush()
		}
	case reflect.Struct:
		elem := reflect.ValueOf(strct)
		var str []string
		for i := 0; i < elem.NumField(); i++ {
			value := elem.Field(i)
			switch elem.Field(i).Kind() {
			case reflect.String:
				str = append(str, value.Interface().(string))
			case reflect.Int:
				str = append(str, strconv.Itoa(value.Interface().(int)))
			}
		}
		_ = csvWriter.Write(str)
		csvWriter.Flush()
	}
	return nil
}
