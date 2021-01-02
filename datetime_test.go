package datetime

import (
	"testing"
	"time"

	"github.com/keyhorn/assert"
	"github.com/keyhorn/assert/matcher"
)

var equalTo = matcher.EqualTo

func TestToday(t *testing.T) {
	today := Today()
	assert.That(t, today.Hour(), equalTo(0))
	assert.That(t, today.Minute(), equalTo(0))
	assert.That(t, today.Second(), equalTo(0))
	assert.That(t, today.Millisecond(), equalTo(0))
	assert.That(t, today.Microsecond(), equalTo(0))
	assert.That(t, today.Nanosecond(), equalTo(0))
}

func TestFormat(t *testing.T) {
	dt := New(2021, time.January, 2, 13, 4, 5, 123456789, time.Now().Location())
	var items = []struct {
		actual   string
		expected string
	}{
		{actual: dt.Format("G"), expected: "AD"},
		{actual: dt.Format("GG"), expected: "A.D."},
		{actual: dt.Format("yyyy"), expected: "2021"},
		{actual: dt.Format("yy"), expected: "21"},
		{actual: dt.Format("MM"), expected: "01"},
		{actual: dt.Format("M"), expected: "1"},
		{actual: dt.Format("dd"), expected: "02"},
		{actual: dt.Format("d"), expected: "2"},
		{actual: dt.Format("E"), expected: "Saturday"},
		{actual: dt.Format("e"), expected: "Sat"},
		{actual: dt.Format("HH"), expected: "13"},
		{actual: dt.Format("hh"), expected: "01"},
		{actual: dt.Format("mm"), expected: "04"},
		{actual: dt.Format("m"), expected: "4"},
		{actual: dt.Format("ss"), expected: "05"},
		{actual: dt.Format("s"), expected: "5"},
		{actual: dt.Format("a"), expected: "PM"},
	}

	for _, item := range items {
		assert.That(t, item.actual, equalTo(item.expected))
	}
}

func TestFormatCommonEra1(t *testing.T) {
	now := Now()
	assert.That(t, now.Format("G yyyy"), equalTo("AD 2021"))
	assert.That(t, now.Format("GG yyyy"), equalTo("A.D. 2021"))
}

func TestFormatCommonEra2(t *testing.T) {
	now := New(-1000, time.April, 10, 0, 0, 0, 0, time.Now().Location())
	assert.That(t, now.Format("G yyyy"), equalTo("BC 1000"))
	assert.That(t, now.Format("GG yyyy"), equalTo("B.C. 1000"))
}

func TestFormatCommonEra3(t *testing.T) {
	now := New(-1000, time.April, 10, 0, 0, 0, 0, time.Now().Location())
	assert.That(t, now.Format("yyyy"), equalTo("-1000"))
}
