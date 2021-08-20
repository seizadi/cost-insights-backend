package utils

import (
	"errors"
	"regexp"
	"strconv"
	"time"
	
	errorsPkg "github.com/pkg/errors"
	"github.com/sajari/regression"
	
	"github.com/seizadi/cost-insight-backend/pkg/pb"
	"github.com/seizadi/cost-insight-backend/pkg/types"
)

func ChangeOf(aggregation []*pb.DateAggregation) *pb.ChangeStatistic {
	var firstAmount int32 = 0
	var lastAmount int32 = 0
	if len(aggregation) > 0 {
		firstAmount = aggregation[0].Amount
		lastAmount = aggregation[len(aggregation) - 1].Amount
	}
	
	// if either the first or last amounts are zero, the rate of increase/decrease is infinite
	if firstAmount == 0 || lastAmount == 0 {
		return &pb.ChangeStatistic{
			Amount: lastAmount - firstAmount,
		}
	}
	
	return &pb.ChangeStatistic{
		Ratio: float32(lastAmount - firstAmount) / float32(firstAmount),
		Amount: lastAmount - firstAmount,
	}
}

func TrendlineOf(aggregation []*pb.DateAggregation) (*pb.Trendline, error) {
	trend := pb.Trendline{}
	r := new(regression.Regression)
	r.SetVar(0, "amount")
	for i:=0; i < len(aggregation); i++ {
		t, err := time.Parse(types.DEFAULT_DATE_FORMAT, aggregation[i].Date)
		if err != nil {
			return &trend, errorsPkg.Wrap(err, "failed to parse date: " + aggregation[i].Date)
		}
		r.Train(regression.DataPoint(float64(t.UnixNano() / int64(time.Millisecond))/1000, []float64{float64(aggregation[i].Amount)}))
	}
	err := r.Run()
	if err != nil {
		return &trend, errorsPkg.Wrap(err, "failed to train on data")
	}
	// Returns form:  Predicted = 67.0341 + amount*-0.0001
	formula := r.Formula
	// RegEx for Floating Point is: [+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?
	regEx := regexp.MustCompile(`Predicted = (?P<intercept>[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?) \+ amount\*(?P<slope>[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?)`)
	matches := regEx.FindStringSubmatch(formula)
	names := regEx.SubexpNames()
	if len(matches) != 3 {
		return &trend, errors.New("failed to parse formula: " + formula)
	}
	
	for i, match := range matches {
		if i != 0 {
			if names[i] == "intercept" {
				intercept, err := strconv.ParseFloat(match, 32)
				if err != nil {
					return &trend, errors.New("failed to parse float: " + names[i])
				}
				trend.Intercept = float32(intercept)
			} else if names[i] == "slope" {
				slope, err := strconv.ParseFloat(match, 32)
				if err != nil {
					return &trend, errors.New("failed to parse float: " + names[i])
				}
				trend.Slope = float32(slope)
			}
		}
	}
	return &trend, nil
}

func GetProductCost(product string, intervals string, cost int32) (*pb.ProductCost, error) {
	aggregation, err := AggregationFor(intervals, cost)
	if err != nil {
		return &pb.ProductCost{}, errors.New("failed to get cost for product: " + product)
	}
	return &pb.ProductCost{
		Id: product,
		Aggregation: aggregation,
	}, nil
}

func GetGroupedProducts(intervals string) ([]*pb.ProductCost, error){
	var cost []*pb.ProductCost
	productCost, err := GetProductCost("Cloud Dataflow", intervals, types.CLOUD_DATAFLOW_COST)
	if err != nil {
		return cost, err
	}
	cost = append(cost, productCost)
	
	productCost, err = GetProductCost("Compute Engine", intervals, types.COMPUTE_ENGINE_COST)
	if err != nil {
		return cost, err
	}
	cost = append(cost, productCost)
	
	productCost, err = GetProductCost("Cloud Storage", intervals, types.CLOUD_STORAGE_COST)
	if err != nil {
		return cost, err
	}
	cost = append(cost, productCost)
	
	productCost, err = GetProductCost("BigQuery", intervals, types.BIG_QUERY_COST)
	if err != nil {
		return cost, err
	}
	cost = append(cost, productCost)
	
	productCost, err = GetProductCost("Cloud SQL", intervals, types.CLOUD_SQL_COST)
	if err != nil {
		return cost, err
	}
	cost = append(cost, productCost)
	
	productCost, err = GetProductCost("Cloud Spanner", intervals, types.CLOUD_SPANNER_COST)
	if err != nil {
		return cost, err
	}
	cost = append(cost, productCost)
	
	productCost, err = GetProductCost("Cloud Pub/Sub", intervals, types.CLOUD_Pub_Sub_COST)
	if err != nil {
		return cost, err
	}
	cost = append(cost, productCost)
	
	productCost, err = GetProductCost("Cloud Bigtable", intervals, types.CLOUD_BIGTABLE_COST)
	if err != nil {
		return cost, err
	}
	cost = append(cost, productCost)

	return cost, nil
}

func GetProjectCost(project string, intervals string, cost int32) (*pb.ProjectCost, error) {
	aggregation, err := AggregationFor(intervals, cost)
	if err != nil {
		return &pb.ProjectCost{}, errors.New("failed to get cost for project: " + project)
	}
	return &pb.ProjectCost{
		Id: project,
		Aggregation: aggregation,
	}, nil
}

func GetGroupedProjects(intervals string) ([]*pb.ProjectCost, error) {
	var cost []*pb.ProjectCost
	productCost, err := GetProjectCost("project-a", intervals, types.PROJECT_A_COST)
	if err != nil {
		return cost, err
	}
	cost = append(cost, productCost)
	
	productCost, err = GetProjectCost("project-b", intervals, types.PROJECT_B_COST)
	if err != nil {
		return cost, err
	}
	cost = append(cost, productCost)
	
	productCost, err = GetProjectCost("project-c", intervals, types.PROJECT_C_COST)
	if err != nil {
		return cost, err
	}
	cost = append(cost, productCost)
	
	return cost, nil
}
