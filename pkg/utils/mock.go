package utils

import (
	"time"
	
	"github.com/seizadi/cost-insights-backend/pkg/pb"
	"github.com/seizadi/cost-insights-backend/pkg/types"
)

func MockComputeEngineInsights() *pb.Entity {
	entity := pb.Entity{
		Id:          "computeEngine",
		Aggregation: []float64{80000, 90000},
		Change: &pb.ChangeStatistic{
			Ratio:  0.125,
			Amount: 10000,
		},
		Entities: &pb.Record{
			Service: []*pb.Entity{
				&pb.Entity{
					Id:          "service-a",
					Aggregation: []float64{20000, 10000},
					Change: &pb.ChangeStatistic{
						Ratio:  -0.5,
						Amount: -10000,
					},
					Entities: &pb.Record{
						SKU: []*pb.Entity{
							&pb.Entity{
								Id:          "Mock SKU A",
								Aggregation: []float64{4000, 2000},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.5,
									Amount: -2000,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Mock SKU B",
								Aggregation: []float64{7000, 6000},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.14285714285714285,
									Amount: -1000,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Mock SKU C",
								Aggregation: []float64{9000, 2000},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.7777777777777778,
									Amount: -7000,
								},
								Entities: &pb.Record{},
							},
						},
						Deployment: []*pb.Entity{
							&pb.Entity{
								Id:          "Compute Engine",
								Aggregation: []float64{7000, 6000},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.5,
									Amount: -2000,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Kubernetes",
								Aggregation: []float64{4000, 2000},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.14285714285714285,
									Amount: -1000,
								},
								Entities: &pb.Record{},
							},
						},
					},
					
				},
				&pb.Entity{
					Id:          "service-b",
					Aggregation: []float64{10000, 20000},
					Change: &pb.ChangeStatistic{
						Ratio:  1,
						Amount: 10000,
					},
					Entities: &pb.Record{
						SKU: []*pb.Entity{
							&pb.Entity{
								Id:          "Mock SKU A",
								Aggregation: []float64{1000, 2000},
								Change: &pb.ChangeStatistic{
									Ratio:  1,
									Amount: 1000,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Mock SKU B",
								Aggregation: []float64{4000, 8000},
								Change: &pb.ChangeStatistic{
									Ratio:  1,
									Amount: 4000,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Mock SKU C",
								Aggregation: []float64{5000, 10000},
								Change: &pb.ChangeStatistic{
									Ratio:  1,
									Amount: 5000,
								},
								Entities: &pb.Record{},
							},
						},
						Deployment: []*pb.Entity{
							&pb.Entity{
								Id:          "Compute Engine",
								Aggregation: []float64{7000, 6000},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.5,
									Amount: -2000,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Kubernetes",
								Aggregation: []float64{4000, 2000},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.14285714285714285,
									Amount: -1000,
								},
								Entities: &pb.Record{},
							},
						},
					},
					
				},
				&pb.Entity{
					Id:          "service-c",
					Aggregation: []float64{0, 10000},
					Change: &pb.ChangeStatistic{
						Amount: 10000,
					},
					Entities: &pb.Record{},
				},
				
			},
		},
	}
	return &entity
}

func MockCloudDataflowInsights() *pb.Entity{
	entity := pb.Entity{
		Id:          "cloudDataflow",
		Aggregation: []float64{100000, 158000},
		Change: &pb.ChangeStatistic{
			Ratio:  0.58,
			Amount: 50000,
		},
		Entities: &pb.Record{
			Pipeline: []*pb.Entity{
				&pb.Entity{
					Aggregation: []float64{10000, 12000},
					Change: &pb.ChangeStatistic{
						Ratio:  0.2,
						Amount: -2000,
					},
					Entities: &pb.Record{
						SKU: []*pb.Entity{
							&pb.Entity{
								Id:          "Mock SKU A",
								Aggregation: []float64{3000, 4000},
								Change: &pb.ChangeStatistic{
									Ratio:  0.333333,
									Amount: 12000,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Mock SKU B",
								Aggregation: []float64{7000, 8000},
								Change: &pb.ChangeStatistic{
									Ratio:  0.14285714,
									Amount: 1000,
								},
								Entities: &pb.Record{},
							},
						},
						Deployment: []*pb.Entity{
							&pb.Entity{
								Id:          "Compute Engine",
								Aggregation: []float64{7000, 6000},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.5,
									Amount: -2000,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Kubernetes",
								Aggregation: []float64{4000, 2000},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.14285714285714285,
									Amount: -1000,
								},
								Entities: &pb.Record{},
							},
						},
					},
					
				},
				&pb.Entity{
					Id: "pipeline-a",
					Aggregation: []float64{60000, 70000},
					Change: &pb.ChangeStatistic{
						Ratio:  0.16666666666666666,
						Amount: 10000,
					},
					Entities: &pb.Record{
						SKU: []*pb.Entity{
							&pb.Entity{
								Id:          "Mock SKU A",
								Aggregation: []float64{20000, 15000},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.25,
									Amount: -5000,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Mock SKU B",
								Aggregation: []float64{30000, 35000},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.16666666666666666,
									Amount: -5000,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Mock SKU C",
								Aggregation: []float64{10000, 20000},
								Change: &pb.ChangeStatistic{
									Ratio:  1,
									Amount: 10000,
								},
								Entities: &pb.Record{},
							},
						},
						Deployment: []*pb.Entity{
							&pb.Entity{
								Id:          "Compute Engine",
								Aggregation: []float64{7000, 6000},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.5,
									Amount: -2000,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Kubernetes",
								Aggregation: []float64{4000, 2000},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.14285714285714285,
									Amount: -1000,
								},
								Entities: &pb.Record{},
							},
						},
					},
					
				},
				&pb.Entity{
					Id: "pipeline-b",
					Aggregation: []float64{12000, 8000},
					Change: &pb.ChangeStatistic{
						Ratio:  -0.33333,
						Amount: -4000,
					},
					Entities: &pb.Record{
						SKU: []*pb.Entity{
							&pb.Entity{
								Id:          "Mock SKU A",
								Aggregation: []float64{4000, 4000},
								Change: &pb.ChangeStatistic{
									Ratio:  0,
									Amount: 0,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Mock SKU B",
								Aggregation: []float64{8000, 4000},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.5,
									Amount: -4000,
								},
								Entities: &pb.Record{},
							},
						},
					},
					
				},
				&pb.Entity{
					Id: "pipeline-c",
					Aggregation: []float64{0, 10000},
					Change: &pb.ChangeStatistic{
						Amount: 10000,
					},
					Entities: &pb.Record{},
				},
			},
		},
	}
	return &entity
}

func MockCloudStorageInsights() *pb.Entity{
	entity := pb.Entity{
		Id:          "cloudStorage",
		Aggregation: []float64{45000, 45000},
		Change: &pb.ChangeStatistic{
			Ratio:  0,
			Amount: 0,
		},
		Entities: &pb.Record{
			Bucket: []*pb.Entity{
				&pb.Entity{
					Id: "bucket-a",
					Aggregation: []float64{15000, 20000},
					Change: &pb.ChangeStatistic{
						Ratio:  0.333,
						Amount: 5000,
					},
					Entities: &pb.Record{
						SKU: []*pb.Entity{
							&pb.Entity{
								Id:          "Mock SKU A",
								Aggregation: []float64{10000, 11000},
								Change: &pb.ChangeStatistic{
									Ratio:  0.1,
									Amount: 1000,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Mock SKU B",
								Aggregation: []float64{2000, 5000},
								Change: &pb.ChangeStatistic{
									Ratio:  1.5,
									Amount: 3000,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Mock SKU C",
								Aggregation: []float64{3000, 4000},
								Change: &pb.ChangeStatistic{
									Ratio:  0.3333,
									Amount: 1000,
								},
								Entities: &pb.Record{},
							},
						},
						Deployment: []*pb.Entity{
							&pb.Entity{
								Id:          "Compute Engine",
								Aggregation: []float64{7000, 6000},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.5,
									Amount: -2000,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Kubernetes",
								Aggregation: []float64{4000, 2000},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.14285714285714285,
									Amount: -1000,
								},
								Entities: &pb.Record{},
							},
						},
					},
					
				},
				&pb.Entity{
					Id: "bucket-b",
					Aggregation: []float64{30000, 25000},
					Change: &pb.ChangeStatistic{
						Ratio:  -0.16666,
						Amount: -5000,
					},
					Entities: &pb.Record{
						SKU: []*pb.Entity{
							&pb.Entity{
								Id:          "Mock SKU A",
								Aggregation: []float64{12000, 13000},
								Change: &pb.ChangeStatistic{
									Ratio:  0.08333333333333333,
									Amount: 1000,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Mock SKU B",
								Aggregation: []float64{16000, 12000},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.25,
									Amount: -4000,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Mock SKU C",
								Aggregation: []float64{2000, 0},
								Change: &pb.ChangeStatistic{
									Amount: -2000,
								},
								Entities: &pb.Record{},
							},
						},
					},
					
				},
				&pb.Entity{
					Id: "bucket-c",
					Aggregation: []float64{0, 0},
					Change: &pb.ChangeStatistic{
						Amount: 0,
					},
					Entities: &pb.Record{},
				},
			},
		},
	}
	return &entity
}

func MockBigQueryInsights() *pb.Entity{
	entity := pb.Entity{
		Id:          "bigQuery",
		Aggregation: []float64{10000, 30000},
		Change: &pb.ChangeStatistic{
			Ratio:  3,
			Amount: 20000,
		},
		Entities: &pb.Record{
			Dataset: []*pb.Entity{
				&pb.Entity{
					Id: "dataset-a",
					Aggregation: []float64{5000, 10000},
					Change: &pb.ChangeStatistic{
						Ratio:  1,
						Amount: 5000,
					},
					Entities: &pb.Record{},
				},
				&pb.Entity{
					Id: "dataset-b",
					Aggregation: []float64{5000, 10000},
					Change: &pb.ChangeStatistic{
						Ratio:  1,
						Amount: 5000,
					},
					Entities: &pb.Record{},
				},
				&pb.Entity{
					Id: "dataset-c",
					Aggregation: []float64{0, 10000},
					Change: &pb.ChangeStatistic{
						Amount: 10000,
					},
					Entities: &pb.Record{},
				},
			},
		},
	}
	return &entity
}

func MockEventsInsights() *pb.Entity{
	entity := pb.Entity{
		Id:          "events",
		Aggregation: []float64{20000, 10000},
		Change: &pb.ChangeStatistic{
			Ratio:  -0.5,
			Amount: -10000,
		},
		Entities: &pb.Record{
			Event: []*pb.Entity{
				&pb.Entity{
					Id:          "event-a",
					Aggregation: []float64{15000, 7000},
					Change: &pb.ChangeStatistic{
						Ratio:  -0.53333333333,
						Amount: -8000,
					},
					Entities: &pb.Record{
						Product: []*pb.Entity{
							&pb.Entity{
								Id:          "Mock Product A",
								Aggregation: []float64{5000, 2000},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.6,
									Amount: -3000,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Mock Product B",
								Aggregation: []float64{7000, 2500},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.64285714285,
									Amount: -4500,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Mock Product C",
								Aggregation: []float64{3000, 2500},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.16666666666,
									Amount: -500,
								},
								Entities: &pb.Record{},
							},
						},
					},
				},
				&pb.Entity{
					Id:          "event-b",
					Aggregation: []float64{5000, 3000},
					Change: &pb.ChangeStatistic{
						Ratio:  -0.4,
						Amount: -2000,
					},
					Entities: &pb.Record{
						Product: []*pb.Entity{
							&pb.Entity{
								Id:          "Mock Product A",
								Aggregation: []float64{2000, 1000},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.5,
									Amount: -1000,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Mock Product B",
								Aggregation: []float64{1000, 1500},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.5,
									Amount: 500,
								},
								Entities: &pb.Record{},
							},
							&pb.Entity{
								Id:          "Mock Product C",
								Aggregation: []float64{2000, 500},
								Change: &pb.ChangeStatistic{
									Ratio:  -0.75,
									Amount: -1500,
								},
								Entities: &pb.Record{},
							},
						},
					},
				},
			},
		},
	}
	return &entity
}

func MockAlerts() []*pb.Entity{
	alerts := []*pb.Entity{}
	entity1 := pb.Entity{
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
	alerts = append(alerts, &entity1)
	entity2 := pb.Entity{
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
	alerts = append(alerts, &entity2)
	entity3 := pb.Entity{
		Type: "KubernetesMigrationAlert",
		StartDate: time.Now().AddDate(0, 0, -30).Format(types.DEFAULT_DATE_FORMAT),
		EndDate: time.Now().Format(types.DEFAULT_DATE_FORMAT),
		Aggregation: []float64{60000, 120000},
		Change: &pb.ChangeStatistic{
			Ratio:  0,
			Amount: 0,
		},
		Services: []*pb.Entity{
			&pb.Entity{
				Id:          "service-a",
				Aggregation: []float64{20000, 10000},
				Change: &pb.ChangeStatistic{
					Ratio:  -0.5,
					Amount: -10000,
				},
				Entities: &pb.Record{},
			},
			&pb.Entity{
				Id:          "service-b",
				Aggregation: []float64{30000, 15000},
				Change: &pb.ChangeStatistic{
					Ratio:  -0.5,
					Amount: -15000,
				},
				Entities: &pb.Record{},
			},
		},
	}
	alerts = append(alerts, &entity3)
	return alerts
}

