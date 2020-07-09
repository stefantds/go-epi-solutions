package huffman_coding_test

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/huffman_coding"
)

func TestHuffmanEncoding(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "huffman_coding.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Symbols        charWithFrequencyDecoder
		ExpectedResult float64
		Details        string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Symbols,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := HuffmanEncoding(tc.Symbols.Values)
			if !reflect.DeepEqual(result, tc.ExpectedResult) {
				t.Errorf("expected %v, got %v", tc.ExpectedResult, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

type charWithFrequencyDecoder struct {
	Values []CharWithFrequency
}

func (o *charWithFrequencyDecoder) DecodeRecord(record string) error {
	allData := make([][2]interface{}, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	values := make([]CharWithFrequency, len(allData))
	for i := 0; i < len(allData); i++ {
		values[i].C = []rune(allData[i][0].(string))[0]
		values[i].Freq = allData[i][1].(float64)
	}

	o.Values = values
	return nil
}