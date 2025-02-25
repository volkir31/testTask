package request

import (
	"strconv"
	"testTask/src/misc"
	"time"
)

type GetFact struct {
	PeriodStart     time.Time   `json:"period_start"`
	PeriodEnd       time.Time   `json:"period_end"`
	PeriodKey       misc.Period `json:"period_key"`
	IndicatorToMoId int         `json:"indicator_to_mo_id"`
}

func (req *GetFact) ToFormData() map[string][]string {
	result := make(map[string][]string)
	result["period_start"] = []string{req.PeriodStart.Format(time.DateOnly)}
	result["period_end"] = []string{req.PeriodEnd.Format(time.DateOnly)}
	result["period_key"] = []string{req.PeriodKey.String()}
	result["indicator_to_mo_id"] = []string{strconv.Itoa(req.IndicatorToMoId)}

	return result
}

func (req *GetFact) Url() string {
	return "https://development.kpi-drive.ru/_api/indicators/get_facts"
}
