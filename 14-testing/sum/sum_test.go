package sum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSumInt(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	want := 15
	got := SumInt(input)
	if got != want {
		// test would continue on if test case fail
		t.Errorf("sum of 1 to 5 should be %v; got %v", want, got)
		//Fatalf stop the test if it fails at this point.
		//t.Fatalf()
	}

	want = 0
	got = SumInt(nil)
	if got != want {
		// test would continue on if test case fail
		t.Errorf("sum of nil should be %v; got %v", want, got)
		//Fatalf stop the test if it fails at this point.
		//t.Fatalf()
	}
}

func TestSumIntV2(t *testing.T) {
	//type args struct {
	//	vs []int
	//}
	tt := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "one to five numbers",
			input: []int{1, 2, 3, 4, 5},
			want:  15,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := SumInt(tc.input)
			require.Equal(t, tc.want, got)
		})

	}

}
