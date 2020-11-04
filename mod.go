package fairy_uncle

import (
	"fmt"
	"github.com/hachi-n/fairy_uncle/lib/models"
)

func AddRecord(key, jsonStr string) error {
	fairyUncleData := models.LoadFairyUncleData()
	fairyUncle, err := models.NewFairyUncle([]byte(jsonStr))
	if err != nil {
		return err
	}

	record := models.FairyUncleRecord{key: fairyUncle}
	if err := fairyUncleData.Insert(record); err != nil {
		return err
	}

	fmt.Printf("Insert key: %s\n", key)
	return nil
}

func RemoveRecord(key string) error {
	fairyUncleData := models.LoadFairyUncleData()
	if err := fairyUncleData.Delete(key); err != nil {
		return err
	}

	fmt.Printf("Deleted key: %s\n", key)
	return nil
}

func ShowRecord(key string) error {
	fairyUncleData := models.LoadFairyUncleData()
	fairyUncle := fairyUncleData.Select(key)
	if fairyUncle == nil {
		return fmt.Errorf("No Record Error. Key: %s\n", key)
	}

	fmt.Println(fairyUncle)
	return nil
}

func ListConfig() error {
	fairyUncleData := models.LoadFairyUncleData()

	if len(*fairyUncleData) == 0 {
		return fmt.Errorf("No Data. Please Add data.\n")
	}

	fmt.Println(fairyUncleData)
	return nil
}

func Notify() error {
	fairyUncleData := models.LoadFairyUncleData()
	_ = fairyUncleData

	return nil
}
