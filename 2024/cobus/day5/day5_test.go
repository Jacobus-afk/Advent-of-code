package main

import "testing"

func TestFindValidOrdering(t *testing.T) {
  orderingRuleMaps := buildOrderingRuleMap(PageOrderingRules)
	t.Run("valid page order", func(t *testing.T) {
		got := checkValidOrdering([]int{75, 47, 64, 53, 29}, orderingRuleMaps)

		want := true

		if got != want {
			t.Errorf("got: %t, want: %t", got, want)
		}
	})

	t.Run("invalid page order", func(t *testing.T) {
		got := checkValidOrdering([]int{75, 97, 47, 61, 53}, orderingRuleMaps)

		want := false

		if got != want {
			t.Errorf("got: %t, want: %t", got, want)
		}
	})

  t.Run("sum of valid middle pages", func(t *testing.T) {
    got := GetValidMiddlePages([][]int{
      {75,47,61,53,29},
      {97,61,53,29,13},
      {75,29,13},
      {75,97,47,61,53},
      {61,13,29},
      {97,13,75,29,47},
    }, orderingRuleMaps)

    want := 143

    if got != want {
      t.Errorf("got %d, want %d", got, want)
    }

  })
}
