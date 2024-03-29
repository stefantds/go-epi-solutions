package offline_sampling_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	utils "github.com/stefantds/go-epi-judge/test_utils"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/offline_sampling"
	"github.com/stefantds/go-epi-judge/test_utils/stats"
)

type solutionFunc = func(int, []int)

var solutions = []solutionFunc{
	RandomSampling,
}

func TestRandomSampling(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "offline_sampling.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		K       int
		A       []int
		Details string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.K,
			&tc.A,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				if err := randomSamplingWrapper(s, tc.K, tc.A); err != nil {
					t.Errorf("%v\ntest case:\n%+v\n", err, tc)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func randomSamplingWrapper(solution solutionFunc, k int, a []int) error {
	return stats.RunFuncWithRetries(
		func() bool {
			return randomSamplingRunner(solution, k, a)
		},
		errors.New("the results don't match the expected distribution"),
	)
}

func randomSamplingRunner(solution solutionFunc, k int, a []int) bool {
	const N = 1000000
	copyA := make([]int, len(a))
	copy(copyA, a)

	results := make([][]int, N)

	for i := 0; i < N; i++ {
		copyA := make([]int, len(a))
		copy(copyA, a)

		solution(k, copyA)

		result := make([]int, k)
		copy(result, copyA[:k])
		results[i] = result
	}

	totalPossibleOutcomes := stats.BinomialCoefficient(len(a), k)

	sort.Ints(copyA)

	combinations := make([][]int, totalPossibleOutcomes)
	for i := 0; i < totalPossibleOutcomes; i++ {
		combinations[i] = stats.ComputeCombinationIdx(copyA, k, i)
	}

	sort.Slice(combinations, func(i, j int) bool {
		return utils.LexIntsCompare(combinations[i], combinations[j])
	})

	sequence := make([]int, len(results))
	for i, r := range results {
		sort.Ints(r)
		pos := sort.Search(
			len(combinations),
			func(j int) bool { return !utils.LexIntsCompare(combinations[j], r) },
		)
		if pos < len(combinations) && reflect.DeepEqual(combinations[pos], r) {
			sequence[i] = pos
		} else {
			panic("result not in known combinations")
		}
	}

	return stats.CheckSequenceIsUniformlyRandom(sequence, totalPossibleOutcomes, 0.01)
}
