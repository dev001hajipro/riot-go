package common

import "testing"

func TestLOLQueue(t *testing.T) {
	if RankedSolo5x5 != "RANKED_SOLO_5x5" {
		t.Fail()
	}
	if RankedFlexSR != "RANKED_FLEX_SR" {
		t.Fail()
	}
	if RankedFlexTT != "RANKED_FLEX_TT" {
		t.Fail()
	}
}
