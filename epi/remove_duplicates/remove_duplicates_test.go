package remove_duplicates_test

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/remove_duplicates"
)

func TestEliminateDuplicate(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "remove_duplicates.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Names          namesDecoder
		ExpectedResult []string
		Details        string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Names,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := eliminateDuplicateWrapper(tc.Names.Value)
			if !equal(result, tc.ExpectedResult) {
				t.Errorf("expected %v, got %v", tc.ExpectedResult, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func eliminateDuplicateWrapper(names []Name) []Name {
	EliminateDuplicate(names)
	return names
}

type namesDecoder struct {
	Value []Name
}

func (d *namesDecoder) DecodeRecord(record string) error {
	allData := make([][2]string, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}
	d.Value = make([]Name, len(allData))

	for i, n := range allData {
		d.Value[i].FirstName = n[0]
		d.Value[i].LastName = n[1]
	}

	return nil
}

func equal(result []Name, expected []string) bool {
	if len(expected) != len(result) {
		return false
	}

	sort.Strings(expected)
	sort.Slice(result, func(i, j int) bool {
		return result[i].FirstName < result[j].FirstName
	})

	for i := 0; i < len(result); i++ {
		if expected[i] != result[i].FirstName {
			return false
		}
	}

	return true
}