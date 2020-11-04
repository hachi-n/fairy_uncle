package models

import (
	"encoding/json"
	"fmt"
	"github.com/hachi-n/fairy_uncle/lib/db"
	"os"
	"strings"
)

type FairyUncleRecord map[string]*FairyUncle

func LoadFairyUncleData() *FairyUncleRecord {
	d, err := db.LoadDatabase()
	if err != nil {
		fmt.Println("load database error")
		os.Exit(1)
	}

	fairyUncleRecord := new(FairyUncleRecord)

	if len(d.Data) <= 0 {
		m := make(map[string]*FairyUncle)
		*fairyUncleRecord = m
		return fairyUncleRecord
	}

	err = json.Unmarshal(d.Data, fairyUncleRecord)
	if err != nil {
		fmt.Println("load database unmarshal error")
		os.Exit(1)
	}

	return fairyUncleRecord
}

func (r *FairyUncleRecord) String() string {
	b, err := r.ToJSON()
	if err != nil {
		return "Undefined"
	}
	return string(b)
}

func (r *FairyUncleRecord) ToJSON() ([]byte, error) {
	return json.MarshalIndent(r, "", strings.Repeat(" ", 4))
}

func (r FairyUncleRecord) Select(key string) *FairyUncle {
	return r[key]
}

func (r FairyUncleRecord) Insert(fairyUncleRecord FairyUncleRecord) error {
	for k, v := range fairyUncleRecord {
		r[k] = v
	}
	return r.save()
}

func (r FairyUncleRecord) Delete(key string) error {
	delete(r, key)
	return r.save()
}

func (r *FairyUncleRecord) save() error {
	b, err := r.ToJSON()
	if err != nil {
		return err
	}
	if err = db.DB.Store(b); err != nil {
		return err
	}
	return nil
}

type FairyUncle struct {
	Title    string
	Time     string
	SubTitle string
	Message  string
	Icon     string
	Open     string
	Sound    string
}

func NewFairyUncle(jsonBytes []byte) (*FairyUncle, error) {
	fairyUncle := new(FairyUncle)
	if err := json.Unmarshal(jsonBytes, fairyUncle); err != nil {
		return fairyUncle, err
	}
	return fairyUncle, nil
}

func (r *FairyUncle) String() string {
	b, err := json.MarshalIndent(r, "", strings.Repeat(" ", 4))
	if err != nil {
		return "Undefined"
	}
	return string(b)
}
