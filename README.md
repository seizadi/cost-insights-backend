# CostInsights Backend

CostInsights Backend is intended to complement CostInsights Plugin Frontend to support 
cost management of workloads running on Cloud Providers like AWS or GCP.
See project docs for more detail.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine 
for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

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

We use [SemVer](http://semver.org/) for versioning. For the versions available, 
see the [tags on this repository](https://github.com/seizadi/cost-insights-backend/tags).

## Testing

```bash
curl http://localhost:8080/cost-insights-backend/v1/version
curl http://localhost:8080/cost-insights-backend/v1/last_complete_billing_date
curl http://localhost:8080/cost-insights-backend/v1/user_groups
curl http://localhost:8080/cost-insights-backend/v1/user_groups?user_id=some_id
curl http://localhost:8080/cost-insights-backend/v1/group_projects
curl http://localhost:8080/cost-insights-backend/v1/group_projects?group=group_id
curl http://localhost:8080/cost-insights-backend/v1/daily_metric_data?intervals="R2/P30D/2021-06-01"
curl http://localhost:8080/cost-insights-backend/v1/group_daily_cost?group=group_id&intervals="R2/P30D/2021-06-01"
curl http://localhost:8080/cost-insights-backend/v1/project_daily_cost?project=project-a&intervals="R2/P30D/2021-06-01"
curl http://localhost:8080/cost-insights-backend/v1/product_insights?product=computeEngine&intervals="R2/P30D/2021-06-01"
curl http://localhost:8080/cost-insights-backend/v1/product_insights?product=cloudDataflow&intervals="R2/P30D/2021-06-01"
curl http://localhost:8080/cost-insights-backend/v1/product_insights?product=cloudStorage&intervals="R2/P30D/2021-06-01"
curl http://localhost:8080/cost-insights-backend/v1/product_insights?product=bigQuery&intervals="R2/P30D/2021-06-01"
curl http://localhost:8080/cost-insights-backend/v1/product_insights?product=events&intervals="R2/P30D/2021-06-01"
curl http://localhost:8080/cost-insights-backend/v1/alerts?group=group_id
```

## Development

### MkDocs
The docs automatically get built on github using github Actions.
You can build and serve them locally for development.
```bash
mkdocs serve
```
