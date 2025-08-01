package types

// RollupConfig represents configuration for a rollup property.
type RollupConfig struct {
	RelationPropertyName string         `json:"relation_property_name"`
	RelationPropertyID   PropertyID     `json:"relation_property_id"`
	RollupPropertyName   string         `json:"rollup_property_name"`
	RollupPropertyID     PropertyID     `json:"rollup_property_id"`
	Function             RollupFunction `json:"function"`
}

// RollupFunction represents aggregation functions available for rollup properties.
type RollupFunction string

// RollupFunction represents aggregation functions available for rollup properties.
//
// See: https://developers.notion.com/reference/property-object#rollup
const (
	RollupFunctionCount            RollupFunction = "count"
	RollupFunctionCountValues      RollupFunction = "count_values"
	RollupFunctionEmpty            RollupFunction = "empty"
	RollupFunctionNotEmpty         RollupFunction = "not_empty"
	RollupFunctionUnique           RollupFunction = "unique"
	RollupFunctionShowUnique       RollupFunction = "show_unique"
	RollupFunctionPercentEmpty     RollupFunction = "percent_empty"
	RollupFunctionPercentNotEmpty  RollupFunction = "percent_not_empty"
	RollupFunctionSum              RollupFunction = "sum"
	RollupFunctionAverage          RollupFunction = "average"
	RollupFunctionMedian           RollupFunction = "median"
	RollupFunctionMin              RollupFunction = "min"
	RollupFunctionMax              RollupFunction = "max"
	RollupFunctionRange            RollupFunction = "range"
	RollupFunctionEarliestDate     RollupFunction = "earliest_date"
	RollupFunctionLatestDate       RollupFunction = "latest_date"
	RollupFunctionDateRange        RollupFunction = "date_range"
	RollupFunctionChecked          RollupFunction = "checked"
	RollupFunctionUnchecked        RollupFunction = "unchecked"
	RollupFunctionPercentChecked   RollupFunction = "percent_checked"
	RollupFunctionPercentUnchecked RollupFunction = "percent_unchecked"
	RollupFunctionCountPerGroup    RollupFunction = "count_per_group"
	RollupFunctionPercentPerGroup  RollupFunction = "percent_per_group"
	RollupFunctionShowOriginal     RollupFunction = "show_original"
)
