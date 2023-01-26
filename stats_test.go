package gocyclo_test

import (
	"reflect"
	"testing"

	"github.com/dwarakauttarkar/gocyclo"
)

func TestAverageComplexity(t *testing.T) {
	tests := []struct {
		stats gocyclo.Stats
		want  float64
	}{
		{gocyclo.Stats{
			{CyclomaticComplexity: 2},
		}, 2},
		{gocyclo.Stats{
			{CyclomaticComplexity: 2},
			{CyclomaticComplexity: 3},
		}, 2.5},
		{gocyclo.Stats{
			{CyclomaticComplexity: 2},
			{CyclomaticComplexity: 3},
			{CyclomaticComplexity: 4},
		}, 3},
		{gocyclo.Stats{
			{CyclomaticComplexity: 2},
			{CyclomaticComplexity: 3},
			{CyclomaticComplexity: 3},
			{CyclomaticComplexity: 3},
		}, 2.75},
	}
	for _, tt := range tests {
		got := tt.stats.AverageComplexity()
		if got != tt.want {
			t.Errorf("Average complexity for %q, got: %g, want: %g", tt.stats, got, tt.want)
		}
	}
}

func TestTotalComplexity(t *testing.T) {
	tests := []struct {
		stats gocyclo.Stats
		want  uint64
	}{
		{gocyclo.Stats{
			{CyclomaticComplexity: 2},
		}, 2},
		{gocyclo.Stats{
			{CyclomaticComplexity: 2},
			{CyclomaticComplexity: 3},
		}, 5},
		{gocyclo.Stats{
			{CyclomaticComplexity: 2},
			{CyclomaticComplexity: 3},
			{CyclomaticComplexity: 4},
		}, 9},
		{gocyclo.Stats{
			{CyclomaticComplexity: 2},
			{CyclomaticComplexity: 3},
			{CyclomaticComplexity: 3},
			{CyclomaticComplexity: 3},
		}, 11},
	}
	for _, tt := range tests {
		got := tt.stats.TotalCyclomaticComplexity()
		if got != tt.want {
			t.Errorf("Total complexity for %q, got: %d, want: %d", tt.stats, got, tt.want)
		}
	}
}

func TestSortAndFilter(t *testing.T) {
	tests := []struct {
		stats gocyclo.Stats
		top   int
		over  int
		want  gocyclo.Stats
	}{
		{
			stats: gocyclo.Stats{
				{CyclomaticComplexity: 1},
				{CyclomaticComplexity: 4},
				{CyclomaticComplexity: 2},
				{CyclomaticComplexity: 3},
			},
			top: -1, over: 0,
			want: gocyclo.Stats{
				{CyclomaticComplexity: 4},
				{CyclomaticComplexity: 3},
				{CyclomaticComplexity: 2},
				{CyclomaticComplexity: 1},
			},
		},
		{
			stats: gocyclo.Stats{
				{CyclomaticComplexity: 1},
				{CyclomaticComplexity: 2},
				{CyclomaticComplexity: 3},
				{CyclomaticComplexity: 4},
			},
			top: 2, over: 0,
			want: gocyclo.Stats{
				{CyclomaticComplexity: 4},
				{CyclomaticComplexity: 3},
			},
		},
		{
			stats: gocyclo.Stats{
				{CyclomaticComplexity: 1},
				{CyclomaticComplexity: 2},
				{CyclomaticComplexity: 4},
				{CyclomaticComplexity: 4},
				{CyclomaticComplexity: 5},
			},
			top: -1, over: 3,
			want: gocyclo.Stats{
				{CyclomaticComplexity: 5},
				{CyclomaticComplexity: 4},
				{CyclomaticComplexity: 4},
			},
		},
		{
			stats: gocyclo.Stats{
				{CyclomaticComplexity: 1},
				{CyclomaticComplexity: 2},
				{CyclomaticComplexity: 3},
				{CyclomaticComplexity: 4},
				{CyclomaticComplexity: 5},
			},
			top: 2, over: 2,
			want: gocyclo.Stats{
				{CyclomaticComplexity: 5},
				{CyclomaticComplexity: 4},
			},
		},
	}
	for _, tt := range tests {
		got := tt.stats.SortAndFilter(tt.top, tt.over)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Sort and filter (top %d over %d) for %q, got: %q, want: %q",
				tt.top, tt.over, tt.stats, got, tt.want)
		}
	}
}
