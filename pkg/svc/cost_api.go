package svc

import (
	"context"
	
	"github.com/golang/protobuf/ptypes/empty"
	
	"github.com/seizadi/cost-insights-backend/pkg/pb"
)

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
// Implements the Cost Insights API
//
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

const (
	DEFAULT_DATE_FORMAT = "%v-%02v-%02v" // "2021-06-16"
)

type CostInsightsApiServer interface{
	GetLastCompleteBillingDate(context.Context, *empty.Empty) (*pb.LastCompleteBillingDateResponse, error)
	GetUserGroups(context.Context, *pb.UserGroupsRequest) (*pb.UserGroupsResponse, error)
	GetGroupProjects(context.Context, *pb.GroupProjectsRequest) (*pb.GroupProjectsResponse, error)
	GetDailyMetricData(context.Context, *pb.DailyMetricDataRequest) (*pb.DailyMetricDataResponse, error)
	GetGroupDailyCost(context.Context, *pb.GroupDailyCostRequest) (*pb.GroupDailyCostResponse, error)
	GetProjectDailyCost(context.Context, *pb.ProjectDailyCostRequest) (*pb.ProjectDailyCostResponse, error)
	GetProductInsights(context.Context, *pb.ProductInsightsRequest) (*pb.Entity, error)
	GetAlerts(context.Context, *pb.AlertRequest) (*pb.AlertResponse, error)
}

