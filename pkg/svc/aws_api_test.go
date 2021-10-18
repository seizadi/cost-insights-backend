package svc

import (
	"strconv"
	"testing"
	
	ceTypes "github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/spf13/viper"
)

type TestMetricAmount struct {
	costRound bool
	expected  float64
}

var tests = []TestMetricAmount{

	{true, 50.00},
	{false, 50.05},
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

	type testData struct {
		startDate string
		endDate string
		amount string
		unit string
	}
	
	type testIterations struct {
		testData []testData
		expectedSupportCost map[AwsAccountType]float64
	}
	
	var testCases = []testIterations{
		{
			testData: []testData {
				{"2020-08-01", "2020-08-02", "1500.00", "USD"},
				{"2020-08-02", "2020-08-03", "1000.00", "USD"},
				{"2020-08-03", "2020-08-04","4000.00", "USD"},
				
			},
			expectedSupportCost: map[AwsAccountType]float64{
				DeveloperAccount: 195.00,
				BusinessAccount: 650.00,
					EnterpriseAccount: 15000.00,
			},
		},
		{
			testData: []testData {
				{"2020-08-01", "2020-08-02", "15.00", "USD"},
				{"2020-08-02", "2020-08-03", "10.00", "USD"},
				{"2020-08-03", "2020-08-04","40.00", "USD"},
				{"2020-08-04", "2020-08-05", "10.00", "USD"},
			},
			expectedSupportCost: map[AwsAccountType]float64{
				DeveloperAccount: 29.00,
				BusinessAccount: 100.00,
				EnterpriseAccount: 15000.00,
			},
		},
		{
			testData: []testData{
				{"2020-08-01", "2020-08-02", "150000.00", "USD"},
				{"2020-08-02", "2020-08-03", "100000.00", "USD"},
				{"2020-08-03", "2020-08-04", "400000.00", "USD"},
				{"2020-08-04", "2020-08-05", "100000.00", "USD"},
			},
			expectedSupportCost: map[AwsAccountType]float64{
				DeveloperAccount:  22500.00,
				BusinessAccount:   29400.00,
				EnterpriseAccount: 52000.00,
			},
		},
		{
			testData: []testData {
				{startDate: "2020-09-01", endDate:"2020-10-01", amount: "755828.00", unit: "USD"},
			},
			expectedSupportCost: map[AwsAccountType]float64{
				DeveloperAccount: 22674.84,
				BusinessAccount: 29574.84,
				EnterpriseAccount: 52291.40,
			},
		},
		
	}
	
	for _ ,test := range testCases {
		
		resultByTime := []ceTypes.ResultByTime{}
		for index, _ := range test.testData {
			result := ceTypes.ResultByTime{
				TimePeriod: &ceTypes.DateInterval{
					Start: &test.testData[index].startDate,
					End:   &test.testData[index].endDate,
				},
				Total: map[string]ceTypes.MetricValue{
					string(ceTypes.MetricNetAmortizedCost): ceTypes.MetricValue{
						Amount: &test.testData[index].amount,
						Unit:   &test.testData[index].unit,
					},
				},
			}
			resultByTime = append(resultByTime, result)
		}
		
		// 9/15/2021 - https://aws.amazon.com/premiumsupport/pricing/
		
		// Test whether an error is thrown
		// Test if the number of aggregation items outputted matches the number of values given
		// Test to see if the expected sum matches that outputted by the summing function
		// Test to see if changing the business type alters the aggregation sum value
		
		for _, account := range AwsAccounts {
			
			viper.Set("account.type", string(account.AccountType))
			viper.Set("support.cost", true)
			
			results, err := aggregationForAWS(resultByTime)
			if err != nil {
				t.Error("Unexpected error:", err)
			}
			if len(results) != len(resultByTime) {
				t.Errorf("Output %d not equal to expected %d", len(results), len(resultByTime))
			}
			
			supportCost, _ := SupportCostForAWS(account, resultByTime)
			if supportCost != test.expectedSupportCost[account.AccountType] {
				t.Errorf("Output %f not equal to expected %f", supportCost, test.expectedSupportCost[account.AccountType])
			}
			
			for index, result := range results {
				
				cost,_ := strconv.ParseFloat(test.testData[index].amount, 64)
				expectedCost := cost + supportCost/float64(len(results))
				
				if expectedCost != result.Amount {
					t.Errorf("Expected Cost %f not equal to expected %f", expectedCost, result.Amount)
				}
			}
		}
	}
		
		// 9/15/2021 - https://aws.amazon.com/premiumsupport/pricing/
		
		// Test whether an error is thrown
		// Test if the number of aggregation items outputted matches the number of values given
		// Test to see if the expected sum matches that outputted by the summing function
		// Test to see if changing the business type alters the aggregation sum value
	
}
