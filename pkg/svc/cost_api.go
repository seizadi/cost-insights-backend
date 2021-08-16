package svc

import (
	"context"
	"fmt"
	"time"
	
	"github.com/golang/protobuf/ptypes/empty"
	
	"github.com/seizadi/aws-cost/pkg/pb"
)

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
// Implements the Cost Insight API
//
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

const (
	// version is the current version of the service
	data = "some data here!!!"
	DEFAULT_DATE_FORMAT = "%v-%02v-%02v" // "2021-06-16"
)

// Default implementation of the AwsCost server interface
type costInsightServer struct{}

// NewCostInsightApiServerServer
// returns an instance of the default server interface
func NewCostInsightApiServerServer() (pb.CostInsightsApiServer, error) {
	return &costInsightServer{}, nil
}

// GetLastCompleteBillingDate
// returns the most current date for which billing data is complete, in YYYY-MM-DD format. This helps
// define the intervals used in other API methods to avoid showing incomplete cost. The costs for
// today, for example, will not be complete. This ideally comes from the cloud provider.
//
// Implements CostInsightApiClient getLastCompleteBillingDate(): Promise<string>;
func (costInsightServer) GetLastCompleteBillingDate(context.Context, *empty.Empty) (*pb.LastCompleteBillingDateResponse, error) {
	year, month, day := time.Now().AddDate(0, 0, -1).Date()
	date := fmt.Sprintf(DEFAULT_DATE_FORMAT, year, int(month), day)
	return &pb.LastCompleteBillingDateResponse{Date: date}, nil
}

