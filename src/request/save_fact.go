package request

import (
	"strconv"
	"time"
)

type SaveFact struct {
	GetFact
	IndicatorToMoFactId int       `json:"indicator_to_mo_fact_id"`
	Value               int       `json:"value"`
	FactTime            time.Time `json:"fact_time"`
	IsPlan              bool      `json:"is_plan"`
	AuthUserId          int       `json:"auth_user_id"`
	Comment             string    `json:"comment"`
}

func (req *SaveFact) ToFormData() map[string][]string {
	result := make(map[string][]string)
	for key, value := range req.GetFact.ToFormData() {
		result[key] = value
	}
	isPlan := 0
	if req.IsPlan {
		isPlan = 1
	}
	result["indicator_to_mo_fact_id"] = []string{strconv.Itoa(req.IndicatorToMoFactId)}
	result["value"] = []string{strconv.Itoa(req.Value)}
	result["fact_time"] = []string{req.FactTime.Format(time.DateOnly)}
	result["is_plan"] = []string{strconv.Itoa(isPlan)}
	result["auth_user_id"] = []string{strconv.Itoa(req.AuthUserId)}
	result["comment"] = []string{req.Comment}

	return result
}

func (req *SaveFact) Url() string {
	return "https://development.kpi-drive.ru/_api/facts/save_fact"
}
