package types

const DEFAULT_DATE_FORMAT = "2006-01-02"
const ALERT_DATE_FORMAT = "2006-01"

type Duration string
const (
	P7D Duration = "P7D"
	P30D Duration = "P30D"
	P90D Duration = "P90D"
	P3M Duration = "P3M"
)

type IntervalFields struct{
	Duration Duration
	EndDate string
}

