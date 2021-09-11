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
