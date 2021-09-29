package tree

import (
	"sort"
	"testing"
)

func TestSearch(t *testing.T) {
	tr := &Tree{}
	tr.Insert("s1", 5, 12)
	tr.Insert("s2", 1, 20)
	tr.Insert("s3", 11, 21)
	tr.Insert("s4", 2, 5)
	tr.Insert("s5", 7, 40)
	tr.Insert("s6", 0, 4)
	tr.Insert("s7", 15, 30)

	tr.root.InorderTraversal()

	testcases := []struct {
		qmin     uint
		qmax     uint
		expected []string
	}{
		{
			qmin:     3,
			qmax:     10,
			expected: []string{"s1", "s2", "s4", "s5", "s6"},
		},
		{
			qmin:     0,
			qmax:     5,
			expected: []string{"s2", "s4", "s6"},
		},
		{
			qmin:     20,
			qmax:     40,
			expected: []string{"s3", "s5", "s7"},
		},
		{
			qmin:     0,
			qmax:     40,
			expected: []string{"s1", "s2", "s3", "s4", "s5", "s6", "s7"},
		},
		{
			qmin:     0,
			qmax:     0,
			expected: []string{},
		},
		{
			qmin:     40,
			qmax:     90,
			expected: []string{},
		},
	}

	for _, tc := range testcases {
		tarr := tr.Search(tc.qmin, tc.qmax)
		sarr := make([]string, 0, len(tarr))
		for _, v := range tarr {
			sarr = append(sarr, v.Name)
		}

		sort.Strings(sarr)

		if len(sarr) != len(tc.expected) {
			t.Logf("expected %v, actual %v for tc = %+v", tc.expected, sarr, tc)
			t.Fail()
			continue
		}

		for i := range sarr {
			if sarr[i] != tc.expected[i] {
				t.Logf("expected %v, actual %v for tc = %+v", tc.expected, sarr, tc)
				t.Fail()
				break
			}
		}

	}

}

func TestIsOverlapping(t *testing.T) {

	n := &TreeNode{
		MinTime: 3,
		MaxTime: 20,
	}

	testcases := []struct {
		minTime  uint
		maxTime  uint
		expected bool
	}{
		{
			minTime:  3,
			maxTime:  20,
			expected: true,
		},
		{
			minTime:  3,
			maxTime:  19,
			expected: true,
		},
		{
			minTime:  2,
			maxTime:  20,
			expected: true,
		},
		{
			minTime:  2,
			maxTime:  21,
			expected: true,
		},
		{
			minTime:  4,
			maxTime:  19,
			expected: true,
		},
		{
			minTime:  2,
			maxTime:  3,
			expected: false,
		},
		{
			minTime:  20,
			maxTime:  21,
			expected: false,
		},
		{
			minTime:  1,
			maxTime:  2,
			expected: false,
		},
		{
			minTime:  21,
			maxTime:  22,
			expected: false,
		},
	}

	for _, tc := range testcases {
		actual := n.isOverlapping(tc.minTime, tc.maxTime)
		if actual != tc.expected {
			t.Logf("expected %v, actual %v for tc = %+v", tc.expected, actual, tc)
			t.Fail()
		}
	}
}
