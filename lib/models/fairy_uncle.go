package models

import (
	"encoding/json"
	"fmt"
	"github.com/hachi-n/fairy_uncle/lib/db"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
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

func (r FairyUncleRecord) Find(key string) *FairyUncle {
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

func (r FairyUncleRecord) Select(options map[string]string) ([]*FairyUncle, error) {
	fairyUncles, err := r.selectTimeFilter(options)
	return fairyUncles, err
}

func (r FairyUncleRecord) selectTimeFilter(options map[string]string) ([]*FairyUncle, error) {
	var fairyUncles []*FairyUncle

	// Fix me
	for _, v := range r {
		scheduledTime, err := time.Parse("04 15", v.Time)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		actionTime, err := time.Parse(time.RFC3339, options["actionTime"])
		y, m, d := actionTime.Date()

		realScheduleTime := scheduledTime.AddDate(y, int(m)-1, d-1)

		if err != nil {
			return nil, err
		}

		i, err := strconv.Atoi(options["tolerance"])
		if err != nil {
			return nil, err
		}
		normalizedTolerance := time.Duration(i)

		if actionTime.Before(realScheduleTime.Add(normalizedTolerance*time.Minute)) &&
			actionTime.After(realScheduleTime.Add(-normalizedTolerance*time.Minute)) {
			fairyUncles = append(fairyUncles, v)
		}

		// Debug
		//if realScheduleTime.Before(realScheduleTime.Add(normalizedTolerance*time.Minute)) &&
		//	realScheduleTime.After(realScheduleTime.Add(-normalizedTolerance*time.Minute)) {
		//	fairyUncles = append(fairyUncles, v)
		//}
	}

	return fairyUncles, nil
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

func (r *FairyUncle) Notify() error {
	// Fix me
	err := exec.Command("terminal-notifier",
		"-title", r.Title,
		"-subtitle", r.SubTitle,
		"-message", r.Message,
		"-sound", r.Sound,
		"-icon", r.Icon,
		"-open", r.Open,
	).Start()
	return err
}
