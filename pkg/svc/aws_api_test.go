package svc

import (
	"testing"

	ceTypes "github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/spf13/viper"
)

type TestMetricAmount struct {
	costRound bool
	expected  float64
}

var tests = []TestMetricAmount{

	TestMetricAmount{true, 50.00},
	TestMetricAmount{false, 50.05},
}

func TestGetAwsMetricAmount(t *testing.T) {

	amount := "50.05"
	unit := "USD"
	metric := ceTypes.MetricValue{
		Amount: &amount,
		Unit:   &unit,
	}
	for _, test := range tests {
		viper.Set("cost.round", test.costRound)
		output := getAwsMetricAmount(metric)
		if output != test.expected {
			t.Errorf("Output %f not equal to expected %f", output, test.expected)
		}
	}

}

func TestAggregationForAws(t *testing.T) {
	
	intervals := [][]string {
		{"2020-08-01", "2020-08-02"},
		{"2020-08-02", "2020-08-03"},
		{"2020-08-03", "2020-08-04"},
	}
	
	values := [][]string {
		{"15.00", "USD"},
		{"10.00", "USD"},
		{"100.00", "USD"},
	}
	
	// 9/15/2021 - https://aws.amazon.com/premiumsupport/pricing/
	// Set the account type to developer then the aggregation should be Greater of $29.00 or 3% added
	// Set the account type to business then
	
	
	resultByTime := []ceTypes.ResultByTime{
		{
			TimePeriod: &ceTypes.DateInterval{
				Start: &intervals[0][0],
				End: &intervals[0][1],
				
			},
			Total: map[string]ceTypes.MetricValue{
				string(ceTypes.MetricNetAmortizedCost) : ceTypes.MetricValue{
					Amount: &values[0][0],
					Unit: &values[0][1],
				},
			},
			
		},
		{
			TimePeriod: &ceTypes.DateInterval{
				Start: &intervals[1][0],
				End: &intervals[1][1],
				
			},
			Total: map[string]ceTypes.MetricValue{
				string(ceTypes.MetricNetAmortizedCost) : ceTypes.MetricValue{
					Amount: &values[1][0],
					Unit: &values[1][1],
				},
			},
			
		},
		
	}
	
		results, err := aggregationForAWS(resultByTime)
		if err != nil {
			t.Error("Unexpected error:", err)
		}
		if len(results) != len(resultByTime) {
			t.Errorf("Output %d not equal to expected %d", len(results), len(resultByTime))
		}
		
}
