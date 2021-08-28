# Cost-Insights-Backend Docs

## Summary

We want to enable the 
[CostInsights plugin](https://backstage.io/blog/2020/10/22/cost-insights-plugin) built into the 
Backstage Dev Portal using this backend project. The CostIsnight plugin frontend does not have have
backend implementation. The CostInsights backend will use the Cloud Provider cost management APIs,
e.g. [AWS Billing and Cost Management API](https://docs.aws.amazon.com/aws-cost-anagement/latest/APIReference/API_GetCostAndUsage.html).

We will try to extend the CloudProvider tags/labels using Kubernetes labels
to extract cost information for each application team.

General flow of the implementation:

* Reference the [Cost Insight AWS Doc](https://github.com/backstage/backstage/blob/master/plugins/cost-insights/contrib/aws-cost-explorer-api.md)
  as a guide.
* Implement the [CostInsightsApi](https://github.com/backstage/backstage/blob/master/plugins/cost-insights/src/api/CostInsightsApi.ts) as stubs that return static data and test with plugin
* Integrate CostInsights Backend using Backstage Plugin and Proxy so that APIs are native to Backstage. This should provide similar function to current CostInsights API Mock but talking to backend server.
* Integrate with [AWS Billing and Cost Management API](https://docs.aws.amazon.com/aws-cost-anagement/latest/APIReference/API_GetCostAndUsage.html) and test with backend plugin


## Data Model
The data model for CostInsights that will be supported by the backend is shown below, if we have
to relate this to the 
[CMDB data model](https://seizadi.github.io/cmdb/model/) you have to reference this table:

| CMDB        | CostInsights |
| :---------: | -----------: |
| Account	  | Project      |
| App	      | Service      |
| KubeCluster | Product      |

I have kept the terminology adopted by CostInsights since it is in line with Backsatge terminology.
```mermaid
erDiagram
    Metric {
        string date
        number amount
    }
    CloudProvider ||--|{ User : ""
    CloudProvider ||--|{ Group : ""
    CloudProvider ||--|{ Project : ""
    User }|--|{ Group : ""
    Group }|--|{ Project : ""    
    Project }|--|{ Product : ""
    Product }|--|{ Metric : ""
    Alert }|--|{ Project : ""
    Alert }|--|| Group : ""
    Group }|--|{ Service : ""
    Product }|--|{ Service : ""
    Alert }|--|{ Service : ""
```

## Data Flows
In the diagrams we will use CI acronym for CostInsights.

```mermaid
sequenceDiagram
    User->>CI Plugin Frontend: Open CI Backstage page
    CI Plugin Frontend->>CI Plugin Backend: CI ClientAPI
    CI Plugin Backend->>Backstage Proxy: API Request
    Backstage Proxy->>CI Backend: API Request
    CI Backend->>Cloud Provider: Cloud Provider API
    Cloud Provider->>CI Backend: Response
    CI Backend->>Backstage Proxy: Map Cloud Provider to CI API
    Backstage Proxy->>CI Plugin Backend: CI Response
    CI Plugin Backend->>CI Plugin Frontend: Map CI Backend to ClientAPI
    CI Plugin Frontend->>User: View CI Page with Cloud Provider Costs
```

## Development
The project uses a Go microservice, this is a departure with Backstage that is developed using nodeJS
for the services. We will use the
[Backstage bacend plugin](https://backstage.io/docs/plugins/backend-plugin)
to implement the 
[CostInsights API](https://github.com/backstage/backstage/blob/master/plugins/cost-insights/src/api/CostInsightsApi.ts).
We will reference the
[CostInsights API template](https://github.com/backstage/backstage/blob/master/plugins/cost-insights/src/example/templates/CostInsightsClient.ts).
for our implementation.

```bash
yarn create-plugin --backend  --no-private cost-insights
```
## MkDocs References
* [How to get diagrams in MkDocs](https://chrieke.medium.com/the-best-mkdocs-plugins-and-customizations-fc820eb19759)
* [How to setup MkDocs on Mac and Github](https://suedbroecker.net/2021/01/25/how-to-install-mkdocs-on-mac-and-setup-the-integration-to-github-pages/)
* [Mermaid](https://mermaid-js.github.io/mermaid)
* [Mermaid2 plugin](https://github.com/fralau/mkdocs-mermaid2-plugin#declaring-the-superfences-extension)
* [Mermaid Live Editor](https://mermaid-js.github.io/mermaid-live-editor/#/edit/eyJjb2RlIjoiZ3JhcGggVERcbiAgICBBW0NocmlzdG1hc10gLS0-fEdldCBtb25leXwgQihHbyBzaG9wcGluZylcbiAgICBCIC0tPiBDe0xldCBtZSB0aGlua31cbiAgICBDIC0tPnxPbmV8IERbTGFwdG9wXVxuICAgIEMgLS0-fFR3b3wgRVtpUGhvbmVdXG4gICAgQyAtLT58VGhyZWV8IEZbZmE6ZmEtY2FyIENhcl0iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOmZhbHNlfQ)
* [Mermaid Diagram Syntax](https://mermaid-js.github.io/mermaid/#/flowchart?id=flowcharts-basic-syntax)
