# aws-cost

We want to build a Cost Explorer into the Backstage Dev Portal, the 
[Backstage Cost Insight Plugin](https://backstage.io/blog/2020/10/22/cost-insights-plugin) will be the
starting point. The Backstage plugin is an overlay on top of Cloud Provider cost management, 
e.g. [AWS Billing and Cost Management API](https://docs.aws.amazon.com/aws-cost-anagement/latest/APIReference/API_GetCostAndUsage.html).
We will try to extend the AWS Cost using Kubernetes labels 
to AWS tags to extract cost information for each application team. 
Backstage currently does not provide a CostInsightsApi client out of the box. This project is
intended to implement a CostInsightsApi client that can be called by Backstage Cost Insight Plugin.

General flow of the implementation:

   * Reference the [Cost Insight AWS Doc](https://github.com/backstage/backstage/blob/master/plugins/cost-insights/contrib/aws-cost-explorer-api.md)
as a guide.
   * Implement the [CostInsightApi](https://github.com/backstage/backstage/blob/master/plugins/cost-insights/src/api/CostInsightsApi.ts) as stubs that return static data and test with plugin
   * Integrate with [AWS Billing and Cost Management API](https://docs.aws.amazon.com/aws-cost-anagement/latest/APIReference/API_GetCostAndUsage.html) and test with plugin


## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software and how to install them.

```
Give examples
```

### Installing

A step-by-step series of examples that tell you have to get a development environment running.

Say what the step will be.

```
Give the example
```

And repeat.

```
until finished
```

End with an example of getting some data out of the system or using it for a little demo.



## Deployment

Add additional notes about how to deploy this application. Maybe list some common pitfalls or debugging strategies.

## Running the tests

Explain how to run the automated tests for this system.

```
Give an example
```

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/seizadi/aws-cost/tags).

## Testing

```bash
curl http://localhost:8080/aws-cost/v1/version
curl http://localhost:8080/aws-cost/v1/last_complete_billing_date
curl http://localhost:8080/aws-cost/v1/user_groups
curl http://localhost:8080/aws-cost/v1/user_groups?user_id=some_id
curl http://localhost:8080/aws-cost/v1/group_projects
curl http://localhost:8080/aws-cost/v1/group_projects?group=group_id
curl http://localhost:8080/aws-cost/v1/daily_metric_data?intervals="R2/P30D/2020-09-01"
curl http://localhost:8080/aws-cost/v1/group_daily_cost?group=group_id&intervals="R2/P30D/2020-09-01"
curl http://localhost:8080/aws-cost/v1/product_insights?product=computeEngine&intervals="R2/P30D/2020-09-01"
curl http://localhost:8080/aws-cost/v1/product_insights?product=cloudDataflow&intervals="R2/P30D/2020-09-01"
curl http://localhost:8080/aws-cost/v1/product_insights?product=cloudStorage&intervals="R2/P30D/2020-09-01"
curl http://localhost:8080/aws-cost/v1/product_insights?product=bigQuery&intervals="R2/P30D/2020-09-01"
curl http://localhost:8080/aws-cost/v1/product_insights?product=events&intervals="R2/P30D/2020-09-01"
curl http://localhost:8080/aws-cost/v1/alerts?group=group_id | jq
```
