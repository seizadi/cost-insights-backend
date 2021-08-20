package utils

import (
	"errors"
	"math"
	"math/rand"
	"regexp"
	"time"
	
	isoDuration "github.com/senseyeio/duration"
	
	"github.com/seizadi/aws-cost/pkg/pb"
	"github.com/seizadi/aws-cost/pkg/types"
)

const (
	// MONTH  = 1
	QUARTER  = 3
	// SEMESTER = 6
)

// FIXME - Much of this can be cleanuped up if we return the time resource with the time string
// so that we don't have to recreate the time resource and avoid error checking

// ParseIntervals
// @param intervals An ISO 8601 repeating interval string, such as R2/P30D/2020-09-01
// https://en.wikipedia.org/wiki/ISO_8601#Repeating_intervals
//
func ParseIntervals(intervals string) (types.IntervalFields, error) {
	retIntervalFields := types.IntervalFields{}
	r := regexp.MustCompile(`\/(?P<duration>P\d+[DM])\/(?P<date>\d{4}-\d{2}-\d{2})`)
	matches := r.FindStringSubmatch(intervals)
	names := r.SubexpNames()
	if ( len(matches) != 3) {
		return retIntervalFields, errors.New("invalid intervals: " +  intervals)
	}

	for i, match := range matches {
		if i != 0 {
			if names[i] == "duration" {
				retIntervalFields.Duration = types.Duration(match)
			} else if names[i] == "date" {
				retIntervalFields.EndDate = match
			}
		}
	}
	return retIntervalFields, nil
}

func LastPeriod(t time.Time, period time.Month) (start, end time.Time) {
	y, m, _ := t.Date()
	loc := t.Location()
	start = time.Date(y, m-period, 1, 0, 0, 0, 0, loc)
	end = time.Date(y, m, 1, 0, 0, 0, -1, loc)
	return start, end
}

func CurrPeriod(t time.Time, period time.Month) (start, end time.Time) {
	y, m, _ := t.Date()
	loc := t.Location()
	start = time.Date(y, m, 1, 0, 0, 0, 0, loc)
	end = time.Date(y, m+period, 1, 0, 0, 0, -1, loc)
	return start, end
}

// TODO - Not sure we need this code for Time pakcage!
func QuarterEndDate(inclusiveEndDate string) (string, error) {
	endDate, err := time.Parse(types.DEFAULT_DATE_FORMAT, inclusiveEndDate)
	if err != nil {
		return "", err
	}
	_, end := CurrPeriod(endDate, QUARTER)
	endOfQuarter := end.Format(types.DEFAULT_DATE_FORMAT)
	if endOfQuarter == inclusiveEndDate {
		return endOfQuarter, nil
	}
	start, _ := CurrPeriod(endDate, QUARTER)
	return start.AddDate(0, 0, -1).Format(types.DEFAULT_DATE_FORMAT), nil
}

// InclusiveStartDateOf
// Derive the start date of a given period, assuming two repeating intervals.
//
// @param duration see comment on Duration enum
// @param inclusiveEndDate from CostInsightsApi.getLastCompleteBillingDate
//
// TODO - Could support a wider range of the ISO durations than the fixed ones here.
// We can also eliminate two sets of case statments by combining duratin:
// 		days := d.W*7 + d.D
//		t = t.AddDate(d.Y, d.M, days)
func InclusiveStartDateOf( duration types.Duration,  inclusiveEndDate string) (string, error) {
	switch (duration) {
	case types.P7D, types.P30D, types.P90D:
		t, err := time.Parse(types.DEFAULT_DATE_FORMAT, inclusiveEndDate)
		if err != nil {
			return "", err
		}
		d, err := isoDuration.ParseISO8601(string(duration))
		if err != nil {
			return "", err
		}
		return t.AddDate(0, 0, -2*d.D).Format(types.DEFAULT_DATE_FORMAT), nil
	case types.P3M:
		q, err := QuarterEndDate(inclusiveEndDate)
		if err != nil {
			return "", err
		}
		qt, err := time.Parse(types.DEFAULT_DATE_FORMAT, q)
		if err != nil {
			return "", err
		}
		d, err := isoDuration.ParseISO8601(string(duration))
		if err != nil {
			return "", err
		}
		return qt.AddDate(0, -2*d.M, 0).Format(types.DEFAULT_DATE_FORMAT), nil
	}
	return "", errors.New(string("duration: " + duration + " unknown"))
}

func ExclusiveEndDateOf(duration types.Duration, inclusiveEndDate string) (string, error) {
	switch (duration) {
	case types.P7D, types.P30D, types.P90D:
		t, err := time.Parse(types.DEFAULT_DATE_FORMAT, inclusiveEndDate)
		if err != nil {
			return "", err
		}
		return t.AddDate(0, 0, 1).Format(types.DEFAULT_DATE_FORMAT), nil
	case types.P3M:
		q, err := QuarterEndDate(inclusiveEndDate)
		if err != nil {
			return "", err
		}
		qt, err := time.Parse(types.DEFAULT_DATE_FORMAT, q)
		if err != nil {
			return "", err
		}
		return qt.AddDate(0, 0, 1).Format(types.DEFAULT_DATE_FORMAT), nil
	}
	return "", errors.New(string("duration: " + duration + " unknown"))
}

func InclusiveEndDateOf(duration types.Duration, inclusiveEndDate string) (string, error) {
	d, err := ExclusiveEndDateOf(duration, inclusiveEndDate)
	if err != nil {
		return "", err
	}
	t, err := time.Parse(types.DEFAULT_DATE_FORMAT, d)
	if err != nil {
		return "", err
	}
	return t.AddDate(0, 0, -1).Format(types.DEFAULT_DATE_FORMAT), nil
}

func nextDelta(baseline int32) float64 {
	const varianceFromBaseline = 0.15
	// Let's give positive vibes in trendlines - higher change for positive delta with >0.5 value
	const positiveTrendChance = 0.55
	const normalization = positiveTrendChance - 1
	return float64(baseline) * (rand.Float64() + normalization) * varianceFromBaseline
}

func AggregationFor(intervals string, baseline int32) ([]*pb.DateAggregation, error) {
	retDateAggregation := []*pb.DateAggregation{}
	r, err := ParseIntervals(intervals)
	if err != nil {
		return retDateAggregation, err
	}
	
	inclusiveEndDate, err := InclusiveEndDateOf(r.Duration, r.EndDate)
	if err != nil {
		return retDateAggregation, err
	}
	
	endDate, err := time.Parse(types.DEFAULT_DATE_FORMAT, r.EndDate)
	if err != nil {
		return retDateAggregation, err
	}
	
	iEndDate, err := InclusiveStartDateOf(r.Duration, inclusiveEndDate)
	if err != nil {
		return retDateAggregation, err
	}
	
	iEndDateT, err := time.Parse(types.DEFAULT_DATE_FORMAT, iEndDate)
	if err != nil {
		return retDateAggregation, err
	}
	
	days := endDate.Sub(iEndDateT).Hours() / 24 // Number of days to create values
	
	for i := 0; i < int(days); i++ {
		var billAmount int32
		if i == 0 {
			billAmount = baseline
		} else {
			billAmount = retDateAggregation[i-1].Amount
		}
		start, err := InclusiveStartDateOf(r.Duration, inclusiveEndDate)
		if err != nil {
			return retDateAggregation, err
		}
		date, err := time.Parse(types.DEFAULT_DATE_FORMAT, start)
		if err != nil {
			return retDateAggregation, err
		}
		value := pb.DateAggregation {
			Date: date.AddDate(0, 0, i).Format(types.DEFAULT_DATE_FORMAT),
			Amount: int32(math.Round(math.Max(0, float64(billAmount)+nextDelta(baseline)))),
		}
		retDateAggregation = append(retDateAggregation, &value)
	}
	return retDateAggregation, nil
}


