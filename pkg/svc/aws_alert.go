package svc

import (
	"context"
	"time"
	
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	ceTypes "github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	
	"github.com/seizadi/cost-insights-backend/pkg/pb"
	"github.com/seizadi/cost-insights-backend/pkg/types"
	"github.com/seizadi/cost-insights-backend/pkg/utils"
)

func (m costInsightsAwsServer) ProjectGrowthAlert() (*pb.Entity, error) {
	entity := pb.Entity{}
	entity.Type= "ProjectGrowthAlert"
	// TODO - Alert on each project for now do it for the aggregate
	entity.Project = "All Projects (AWS Accounts)"
	
	// TODO - Configure the Alert Interval - default to 90 days
	// FIXME - Change from hard coded value to Time.Now()
	intervals := "R2/P90D/2021-09-01"
	interval, err := utils.ParseIntervals(intervals)
	if err != nil {
		return &pb.Entity{}, err
	}
	
	startDate, err := utils.InclusiveStartDateOf(interval.Duration, interval.EndDate)
	if err != nil {
		return &pb.Entity{}, err
	}
	
	// TODO - Alert on each project for now do it for the aggregate
	// Daily cost grouped by cloud product (AWS Service)
	groupKey := "SERVICE"
	resp, err := m.client.GetCostAndUsage(context.TODO(), &costexplorer.GetCostAndUsageInput{
		TimePeriod: &ceTypes.DateInterval{Start: &startDate, End: &interval.EndDate},
		Metrics: []string{string(DEFAULT_COST_METHOD)},
		// TODO - Need Account(i.e. Project) to filter
		//Filter: &ceTypes.Expression{
		//	Dimensions: &ceTypes.DimensionValues{
		//		Key: ceTypes.DimensionLinkedAccount,
		//		Values: []string{"ACCOUNT_ID"},
		//	},
		//},
		Granularity: ceTypes.GranularityDaily,
		GroupBy: []ceTypes.GroupDefinition{
			{Key: &groupKey, Type: ceTypes.GroupDefinitionTypeDimension},
		},
	})
	if err != nil {
		return &pb.Entity{}, err
	}
	
	entities, err := getEntityAwsProducts(resp.ResultsByTime)
	if err != nil {
		return &pb.Entity{}, err
	}
	
	entity.Products = entities
	
	// TODO - CostInsights Plugin should accept Start and End date from duration without truncating it to month
	start, err := time.Parse(types.DEFAULT_DATE_FORMAT, startDate)
	if err != nil {
		return &pb.Entity{}, err
	}
	
	end, err := time.Parse(types.DEFAULT_DATE_FORMAT, interval.EndDate)
	if err != nil {
		return &pb.Entity{}, err
	}
	entity.PeriodStart = start.Format(types.ALERT_DATE_FORMAT)
	entity.PeriodEnd = end.Format(types.ALERT_DATE_FORMAT)
	
	entity.PeriodStart = "2020-Q2"
	entity.PeriodEnd = "2020-Q3"
	
	var startAggregate float64
	var endAggregate float64
	for _, e := range entity.Products {
		startAggregate = startAggregate + e.Aggregation[0]
		endAggregate = endAggregate + e.Aggregation[1]
	}
	entity.Aggregation = []float64{startAggregate, endAggregate}
	entity.Change = utils.ChangeOfEntity(entity.Aggregation)
	
	return &entity, nil
}

func ProjectGrowthAlert() (*pb.Entity, error) {
	entity := pb.Entity{
		Type: "ProjectGrowthAlert",
		Project:          "example-project",
		PeriodStart: "2020-02",
		PeriodEnd: "2020-03",
		Aggregation: []float64{60000, 120000},
		Change: &pb.ChangeStatistic{
			Ratio:  1,
			Amount: 60000,
		},
		Products: []*pb.Entity{
			&pb.Entity{
				Id:          "Compute Engine",
				Aggregation: []float64{50000, 118000},
			},
			&pb.Entity{
				Id:          "Cloud Dataflow",
				Aggregation: []float64{1200, 1500},
			},
			&pb.Entity{
				Id:          "Cloud Storage",
				Aggregation: []float64{800, 500},
			},
		},
	}
	
	return &entity, nil
}

func UnlabeledAlert() (*pb.Entity, error) {
	entity := pb.Entity{
		Type: "UnlabeledDataflowAlert",
		PeriodStart: "2020-09-1",
		PeriodEnd: "2020-09-30",
		LabeledCost: 6200,
		UnlabeledCost: 7000,
		Projects: []*pb.Entity{
			&pb.Entity{
				Id:          "example-project-1",
				UnlabeledCost: 5000,
				LabeledCost: 3000,
			},
			&pb.Entity{
				Id:          "example-project-2",
				UnlabeledCost: 2000,
				LabeledCost: 3200,
			},
		},
	}
	
	return &entity, nil
}
