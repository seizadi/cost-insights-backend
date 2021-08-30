package svc

import (
	"context"
	"errors"
	"time"
	
	"github.com/golang/protobuf/ptypes/empty"
	
	"github.com/seizadi/cost-insights-backend/pkg/pb"
	"github.com/seizadi/cost-insights-backend/pkg/types"
	"github.com/seizadi/cost-insights-backend/pkg/utils"
)

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
// Implements the Cost Insights API
//
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Default implementation of the AwsCost server interface
type costInsightsMockServer struct{}

// NewCostInsightsApiMockServer
// returns an instance of the default server interface
func NewCostInsightsApiMockServer() (pb.CostInsightsApiServer, error) {
	return &costInsightsMockServer{}, nil
}

// GetLastCompleteBillingDate
// returns the most current date for which billing data is complete, in YYYY-MM-DD format. This helps
// define the intervals used in other API methods to avoid showing incomplete cost. The costs for
// today, for example, will not be complete. This ideally comes from the cloud provider.
//
// Implements CostInsightsApiClient getLastCompleteBillingDate(): Promise<string>;
func (costInsightsMockServer) GetLastCompleteBillingDate(context.Context, *empty.Empty) (*pb.LastCompleteBillingDateResponse, error) {
	date := time.Now().AddDate(0, 0, -1).Format(types.DEFAULT_DATE_FORMAT)
	return &pb.LastCompleteBillingDateResponse{Date: date}, nil
}

// GetUserGroups
// Get a list of groups the given user belongs to. These may be LDAP groups or similar
// organizational groups. Cost Insights is designed to show costs based on group membership;
// if a user has multiple groups, they are able to switch between groups to see costs for each.
//
// This method should be removed once the Backstage identity plugin provides the same concept.
//
// @param userId The login id for the current user
//
// Implements CostInsightsApiClient getUserGroups(userId: string): Promise<Group[]>;
//

func (costInsightsMockServer) GetUserGroups(context.Context, *pb.UserGroupsRequest) (*pb.UserGroupsResponse, error) {
	groups := []*pb.Group{
		{Id: "pied-piper"},
	}
	return &pb.UserGroupsResponse{Groups: groups}, nil
}

// GetUserGroups
// Get a list of cloud billing entities that belong to this group (projects in GCP, AWS has a
// similar concept in billing accounts). These act as filters for the displayed costs, users can
// choose whether they see all costs for a group, or those from a particular owned project.
//
// @param group The group id from getUserGroups or query parameters
// Implements CostInsightsApiClient getGroupProjects(group: string): Promise<Project[]>;
func (costInsightsMockServer) GetGroupProjects(context.Context, *pb.GroupProjectsRequest) (*pb.GroupProjectsResponse, error) {
	projects := []*pb.Project{
		{ Id: "project-a" },
		{ Id: "project-b" },
		{ Id: "project-c" },
	}
	return &pb.GroupProjectsResponse{Projects: projects}, nil
}

// GetGroupDailyCost
// Get daily cost aggregations for a given group and interval time frame.
//
// The return type includes an array of daily cost aggregations as well as statistics about the
// change in cost over the intervals. Calculating these statistics requires us to bucket costs
// into two or more time periods, hence a repeating interval format rather than just a start and
// end date.
//
// The rate of change in this comparison allows teams to reason about their cost growth (or
// reduction) and compare it to metrics important to the business.
//
// @param group The group id from getUserGroups or query parameters
// @param intervals An ISO 8601 repeating interval string, such as R2/P30D/2020-09-01
//   https://en.wikipedia.org/wiki/ISO_8601#Repeating_intervals
//
// Implements CostInsightsApiClient getGroupDailyCost(group: string, intervals: string): Promise<Cost>;
func (costInsightsMockServer) GetGroupDailyCost(ctx context.Context, req *pb.GroupDailyCostRequest) (*pb.GroupDailyCostResponse, error) {
	cost := pb.GroupDailyCostResponse{}
	cost.Format = "number"
	aggregation, err := utils.AggregationFor(req.Intervals, types.GROUP_COST)
	if err != nil {
		return &pb.GroupDailyCostResponse{}, err
	}
	cost.Aggregation = aggregation
	cost.Change = utils.ChangeOf(aggregation)
	trendline, err := utils.TrendlineOf(aggregation)
	if err != nil {
		return &pb.GroupDailyCostResponse{}, err
	}
	cost.Trendline = trendline
	
	// Optional field providing cost groupings / breakdowns keyed by the type. In this example,
	// daily cost grouped by cloud product OR by project / billing account.
	var groupedCosts pb.GroupedCosts
	productCost, err := utils.GetGroupedProducts(req.Intervals)
	if err != nil {
		return &cost, err
	}
	groupedCosts.Product = productCost
	
	projectCost, err := utils.GetGroupedProjects(req.Intervals)
	if err != nil {
		return &cost, err
	}
	groupedCosts.Project = projectCost
	cost.GroupedCosts = &groupedCosts
	
	return &cost, nil
}

// GetDailyMetricData
//    * Get aggregations for a particular metric and interval time frame. Teams
//    * can see metrics important to their business in comparison to the growth
//    * (or reduction) of a project or group's daily costs.
//    *
//    * @param metric A metric from the cost-insights configuration in app-config.yaml.
//    * @param intervals An ISO 8601 repeating interval string, such as R2/P30D/2020-09-01
//    *   https://en.wikipedia.org/wiki/ISO_8601#Repeating_intervals
//    */
// Implements CostInsightsApiClient getDailyMetricData(metric: string, intervals: string): Promise<MetricData>;
func (costInsightsMockServer) GetDailyMetricData(ctx context.Context, req *pb.DailyMetricDataRequest) (*pb.DailyMetricDataResponse, error) {
	cost := pb.DailyMetricDataResponse{}
	cost.Format = "number"
	aggregation, err := utils.AggregationFor(req.Intervals, types.DAILY_COST)
	if err != nil {
		return &pb.DailyMetricDataResponse{}, err
	}
	cost.Aggregation = aggregation
	cost.Change = utils.ChangeOf(aggregation)
	trendline, err := utils.TrendlineOf(aggregation)
	if err != nil {
		return &pb.DailyMetricDataResponse{}, err
	}
	cost.Trendline = trendline
	
	return &cost, nil
}

// GetProjectDailyCost
// Get daily cost aggregations for a given billing entity (project in GCP, AWS has a similar
// concept in billing accounts) and interval time frame.
//
// The return type includes an array of daily cost aggregations as well as statistics about the
// change in cost over the intervals. Calculating these statistics requires us to bucket costs
// into two or more time periods, hence a repeating interval format rather than just a start and
// end date.
//
// The rate of change in this comparison allows teams to reason about the project's cost growth
// (or reduction) and compare it to metrics important to the business.
//
// @param project The project id from getGroupProjects or query parameters
// @param intervals An ISO 8601 repeating interval string, such as R2/P30D/2020-09-01
//   https://en.wikipedia.org/wiki/ISO_8601#Repeating_intervals
//
// Implements CostInsightsApiClient getProjectDailyCost(project: string, intervals: string): Promise<Cost>;
func (costInsightsMockServer) GetProjectDailyCost(ctx context.Context, req *pb.ProjectDailyCostRequest) (*pb.ProjectDailyCostResponse, error) {
	cost := pb.ProjectDailyCostResponse{}
	cost.Format = "number"
	aggregation, err := utils.AggregationFor(req.Intervals, types.GROUP_COST)
	if err != nil {
		return &pb.ProjectDailyCostResponse{}, err
	}
	cost.Aggregation = aggregation
	cost.Change = utils.ChangeOf(aggregation)
	trendline, err := utils.TrendlineOf(aggregation)
	if err != nil {
		return &pb.ProjectDailyCostResponse{}, err
	}
	cost.Trendline = trendline
	
	var groupedCosts pb.GroupedCosts
	projectCost, err := utils.GetGroupedProjects(req.Intervals)
	if err != nil {
		return &cost, err
	}
	groupedCosts.Project = projectCost
	
	return &cost, nil
}

// GetProductInsights
// Get cost aggregations for a particular cloud product and interval time frame. This includes
// total cost for the product, as well as a breakdown of particular entities that incurred cost
// in this product. The type of entity depends on the product - it may be deployed services,
// storage buckets, managed database instances, etc.
//
// If project is supplied, this should only return product costs for the given billing entity
// (project in GCP).
//
// The time period is supplied as a Duration rather than intervals, since this is always expected
// to return data for two bucketed time period (e.g. month vs month, or quarter vs quarter).
//
// @param options Options to use when fetching insights for a particular cloud product and
//                interval time frame.
//
// Implements CostInsightsApiClient getProductInsights(options: ProductInsightsOptions): Promise<Entity>;
func (costInsightsMockServer) GetProductInsights(ctx context.Context, req *pb.ProductInsightsRequest) (*pb.Entity, error) {
	switch (req.Product) {
	case "computeEngine":
		return utils.MockComputeEngineInsights(), nil;
	case "cloudDataflow":
		return utils.MockCloudDataflowInsights(), nil;
	case "cloudStorage":
		return utils.MockCloudStorageInsights(), nil;
	case "bigQuery":
		return utils.MockBigQueryInsights(), nil;
	case "events":
		return utils.MockEventsInsights(), nil;
	default:
		return &pb.Entity{}, errors.New("failed to get insights for " + req.Product + " product must match product property in configuration(app-info.yaml)")
	}
	return &pb.Entity{}, nil
}

// GetAlerts
//
// Get current cost alerts for a given group. These show up as Action Items for the group on the
// Cost Insights page. Alerts may include cost-saving recommendations, such as infrastructure
// migrations, or cost-related warnings, such as an unexpected billing anomaly.
//
// Implements CostInsightsApiClient getAlerts(group: string): Promise<Alert[]>;
func (costInsightsMockServer) GetAlerts(ctx context.Context, req *pb.AlertRequest) (*pb.AlertResponse, error) {
	return &pb.AlertResponse{Alerts: utils.MockAlerts()}, nil;
}

