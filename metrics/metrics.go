package metrics

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
	
	"github.com/seizadi/cost-insights-backend/pkg/pb"
	"github.com/seizadi/cost-insights-backend/pkg/types"
	"github.com/seizadi/cost-insights-backend/pkg/utils"
)


// CustomMetric struct which contains a date
// a DAR (Daily Average Request) and DAC (Daily Average Client)
type CustomMetric struct {
	Date   string `json:"Date"`
	DAR    float64    `json:"DAR"`
	DAC   float64 `json:"DAC"`
}

type Budget struct {
	Date   string `json:"Date"`
	BudgetCAPEX    float64    `json:"BudgetCAPEX"`
	BudgetOPEX   float64 `json:"BudgetOPEX"`
	BudgetTotal   float64 `json:"BudgetTotal"`
}

func getMockMetrics() (*[]CustomMetric, error) {
	// Open our jsonFile
	metricsFile, err := os.Open("metrics/metrics.json")
	if err != nil {
		return nil, err
	}
	defer metricsFile.Close()
	
	byteValue, err := ioutil.ReadAll(metricsFile)
	if err != nil {
		return nil, err
	}
	
	var metrics []CustomMetric
	err = json.Unmarshal(byteValue, &metrics)
	if err != nil {
		return nil, err
	}
	
	return &metrics, nil
}

func getMockBudget() (*[]Budget, error) {
	// Open our jsonFile
	budgetFile, err := os.Open("metrics/budget.json")
	if err != nil {
		return nil, err
	}
	defer budgetFile.Close()
	
	byteValue, err := ioutil.ReadAll(budgetFile)
	if err != nil {
		return nil, err
	}
	
	var budget []Budget
	err = json.Unmarshal(byteValue, &budget)
	if err != nil {
		return nil, err
	}
	
	return &budget, nil
}

func getMetricKeyIndex(metrics *[]CustomMetric) map[string]int {
	keys := make(map[string]int)
	
	for index, metric := range *metrics {
		keys[metric.Date] = index
	}
	return keys
}

func getMetricValue(metricType string, curValue float64, curDate string, keys map[string]int, mockMetrics *[]CustomMetric) float64 {
	if index, ok := keys[curDate]; ok {
		metrics := *mockMetrics
		metric := metrics[index]
		if metricType == "DAR" {
			return metric.DAR
		} else if metricType == "DAC" {
			return metric.DAC
		}
	}
	return curValue
}

func getBudgetKeyIndex(budget *[]Budget) map[string]int {
	keys := make(map[string]int)
	
	for index, v := range *budget {
		keys[v.Date] = index
	}
	return keys
}

func getBudgetValue(budgetType string, curValue float64, curDate string, keys map[string]int, mockBudget *[]Budget) float64 {
	if index, ok := keys[curDate]; ok {
		items := *mockBudget
		item := items[index]
		if budgetType == "BudgetTotal" {
			return item.BudgetTotal
		} else if budgetType == "BudgetCAPEX" {
			return item.BudgetCAPEX
		} else if budgetType == "BudgetOPEX" {
			return item.BudgetOPEX
		}
	}
	return curValue
}

func GetMetrics(metricType string, intervals string) ([]*pb.DateAggregation, error){
	retDateAggregation := []*pb.DateAggregation{}
	var mockMetrics *[]CustomMetric
	var mockMBudget *[]Budget
	var keys map[string]int
	var err error
	
	if metricType == "BudgetTotal" {
		mockMBudget, err = getMockBudget()
		keys = getBudgetKeyIndex(mockMBudget)
	} else {
		mockMetrics, err = getMockMetrics()
		keys = getMetricKeyIndex(mockMetrics)
	}
	if err != nil {
		return retDateAggregation, err
	}

	
	
	r, err := utils.ParseIntervals(intervals)
	if err != nil {
		return retDateAggregation, err
	}
	
	inclusiveEndDate, err := utils.InclusiveEndDateOf(r.Duration, r.EndDate)
	if err != nil {
		return retDateAggregation, err
	}
	
	endDate, err := time.Parse(types.DEFAULT_DATE_FORMAT, r.EndDate)
	if err != nil {
		return retDateAggregation, err
	}
	
	iEndDate, err := utils.InclusiveStartDateOf(r.Duration, inclusiveEndDate)
	if err != nil {
		return retDateAggregation, err
	}
	
	iEndDateT, err := time.Parse(types.DEFAULT_DATE_FORMAT, iEndDate)
	if err != nil {
		return retDateAggregation, err
	}
	
	days := endDate.Sub(iEndDateT).Hours() / 24 // Number of days to create values
	
	var startValue float64
	var curValue float64
	
	for i := 0; i < int(days); i++ {
		start, err := utils.InclusiveStartDateOf(r.Duration, inclusiveEndDate)
		if err != nil {
			return retDateAggregation, err
		}
		date, err := time.Parse(types.DEFAULT_DATE_FORMAT, start)
		if err != nil {
			return retDateAggregation, err
		}
		
		curDate := date.AddDate(0, 0, i).Format(types.DEFAULT_DATE_FORMAT)
		if metricType == "BudgetTotal" {
			curValue = getBudgetValue(metricType, curValue, curDate, keys, mockMBudget)
		} else {
			curValue = getMetricValue(metricType, curValue, curDate, keys, mockMetrics)
		}
		
		if startValue == 0 && curValue != 0 {
			startValue = curValue
		}
		
		value := pb.DateAggregation {
			Date: curDate,
			Amount: curValue,
		}
		retDateAggregation = append(retDateAggregation, &value)
	}
	
	// Set any zero value to the startValue
	for i := 0; i < int(days); i++ {
		if retDateAggregation[i].Amount == 0 {
			retDateAggregation[i].Amount = startValue
		}
	}
		
		return retDateAggregation, nil
}
