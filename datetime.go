package datetime

import (
	"math"
	"strings"
	"time"
)

// A Datetime represents an instant in time with nanosecond precision.
type Datetime struct {
	time time.Time
}

// Now returns the current local time.
func Now() Datetime {
	return Datetime{
		time: time.Now(),
	}
}

// Today returns the current local time.
func Today() Datetime {
	t := time.Now()
	return Datetime{
		time: time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()),
	}
}

// New returns the Time corresponding to
//	yyyy-mm-dd hh:mm:ss + nsec nanoseconds
// in the appropriate zone for that time in the given location.
func New(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) Datetime {
	return Datetime{
		time: time.Date(year, month, day, hour, min, sec, nsec, loc),
	}
}

// String returns the English name of the month ("January", "February", ...).
func (d Datetime) String() string {
	return d.time.String()
}

// GetTime returns time.Time.
func (d Datetime) GetTime() time.Time {
	return d.time
}

// Year returns the year in which t occurs.
func (d Datetime) Year() int {
	return d.time.Year()
}

// Month returns the month of the year specified by t.
func (d Datetime) Month() time.Month {
	return d.time.Month()
}

// Day returns the day of the month specified by t.
func (d Datetime) Day() int {
	return d.time.Day()
}

// Weekday returns the day of the week specified by t.
func (d Datetime) Weekday() time.Weekday {
	return d.time.Weekday()
}

// Hour returns the hour within the day specified by t, in the range [0, 23].
func (d Datetime) Hour() int {
	return d.time.Hour()
}

// Minute returns the minute offset within the hour specified by t, in the range [0, 59].
func (d Datetime) Minute() int {
	return d.time.Minute()
}

// Second returns the second offset within the minute specified by t, in the range [0, 59].
func (d Datetime) Second() int {
	return d.time.Second()
}

// Millisecond returns the millisecond offset within the second specified by t,
// in the range [0, 999].
func (d Datetime) Millisecond() int {
	return d.time.Nanosecond() / 1000000
}

// Microsecond returns the microsecond offset within the second specified by t,
// in the range [0, 999999].
func (d Datetime) Microsecond() int {
	return d.time.Nanosecond() / 1000
}

// Nanosecond returns the nanosecond offset within the second specified by t,
// in the range [0, 999999999].
func (d Datetime) Nanosecond() int {
	return d.time.Nanosecond()
}

// Timestamp returns the timestamp specified by t.
func (d Datetime) Timestamp() int64 {
	return d.time.UnixNano() / int64(time.Millisecond)
}

// 文字	日付または時刻のコンポーネント	表示	例
// G	紀元	テキスト	AD
// y	年	年	1996; 96
// Y	暦週の基準年	年	2009; 09
// M	年における月(状況依存)	月	July; Jul; 07
// L	年における月(スタンドアロン形式)	月	July; Jul; 07
// w	年における週	数値	27
// W	月における週	数値	2
// D	年における日	数値	189
// d	月における日	数値	10
// F	月における曜日	数値	2
// E	曜日の名前	テキスト	Tuesday; Tue
// u	曜日の番号(1 =月曜、...、7 =日曜)	数値	1
// a	午前/午後	テキスト	PM
// H	一日における時(0 - 23)	数値	0
// k	一日における時(1 - 24)	数値	24
// K	午前/午後の時(0 - 11)	数値	0
// h	午前/午後の時(1 - 12)	数値	12
// m	分	数値	30
// s	秒	数値	55
// S	ミリ秒	数値	978
// z	タイムゾーン	一般的なタイムゾーン	Pacific Standard Time; PST; GMT-08:00
// Z	タイムゾーン	RFC 822タイムゾーン	-0800
// X	タイムゾーン	ISO 8601タイムゾーン	-08; -0800; -08:00

// Format returns a textual representation of the time value formatted
// according to layout, which defines the format by showing how the reference
// time, defined to be
//	Mon MMM d hh:mm:ss Z z yyyy
// would be displayed if it were the value; it serves as an example of the
// desired output. The same display rules will then be applied to the time
// value.
//
// Format    | Description                  | Example
// ----------+------------------------------+-----------
// GG        | Common era                   | A.D.; B.C.
// G         | Common era                   | AD; BC
// yyyy      | Year                         | 2020
// yy        | Year                         | 20
// MM        | Month in year                | 04; 12
// M         | Month in year                | 4; 12
// dd        | Day in month                 | 01; 31
// d         | Day in month                 | 1; 31
// E         | Day of the week              | Monday
// e         | Day of the week              | Mon
// HH        | Hour in day (1-24)           | 01; 24
// hh        | Hour in am/pm (1-12)         | 01; 12
// mm        | Minute in hour               | 01; 60
// m         | Minute in hour               | 1; 60
// ss        | Second in minute             | 01; 60
// s         | Second in minute             | 1; 60
// SSS       | Millisecond                  | 123
// SSSSSS    | Microsecond                  | 123456
// SSSSSSSSS | Nanosecond                   | 123456789
// a         | Am/pm marker                 | AM; PM
func (d Datetime) Format(format string) string {
	t := d.time
	f := format

	// replace common era
	f, t = replaceCommonEra(f, t, "GG", "B.C.", "A.D.")
	f, t = replaceCommonEra(f, t, "G", "BC", "AD")

	f = d.replaceFormat(f)
	return t.Format(f)
}

func replaceCommonEra(f string, t time.Time, oldStr, newStrBC, newStrAD string) (string, time.Time) {
	_f := f
	_t := t

	if strings.Contains(_f, oldStr) {
		// relace common era
		if _t.Year() < 0 {
			// B.C.
			_f = strings.Replace(_f, oldStr, newStrBC, -1)
		} else {
			// A.D.
			_f = strings.Replace(_f, oldStr, newStrAD, -1)
		}
		// year change plus
		_t = time.Date(int(math.Abs(float64(_t.Year()))), _t.Month(), _t.Day(), _t.Hour(), _t.Minute(), _t.Second(), _t.Nanosecond(), _t.Location())
	}
	return _f, _t
}

func (d Datetime) replaceFormat(format string) string {
	var _format = format
	_format = d.replaceFormatYear(_format)
	_format = d.replaceFormatMonth(_format)
	_format = d.replaceFormatDay(_format)
	_format = d.replaceFormatHour(_format)
	_format = d.replaceFormatMinute(_format)
	_format = d.replaceFormatSecond(_format)
	_format = d.replaceFormatAMPM(_format)
	_format = d.replaceFormatTZ(_format)
	_format = d.replaceFormatWeekday(_format)
	return _format
}

func (d Datetime) replaceFormatYear(format string) string {
	var _format = format
	if strings.Contains(_format, "yyyy") {
		_format = strings.Replace(_format, "yyyy", "2006", -1)
	}
	if strings.Contains(_format, "yy") {
		_format = strings.Replace(_format, "yy", "06", -1)
	}
	return _format
}

func (d Datetime) replaceFormatMonth(format string) string {
	var _format = format
	if strings.Contains(_format, "MMMM") {
		_format = strings.Replace(_format, "MMMM", "January", -1)
	}
	if strings.Contains(_format, "MMM") {
		_format = strings.Replace(_format, "MMM", "Jan", -1)
	}
	if strings.Contains(_format, "MM") {
		_format = strings.Replace(_format, "MM", "01", -1)
	}
	if strings.Contains(_format, "M") {
		_format = strings.Replace(_format, "M", "1", -1)
	}
	return _format
}

func (d Datetime) replaceFormatDay(format string) string {
	var _format = format
	if strings.Contains(_format, "dd") {
		_format = strings.Replace(_format, "dd", "02", -1)
	}
	if strings.Contains(_format, "d") {
		_format = strings.Replace(_format, "d", "2", -1)
	}
	return _format
}

func (d Datetime) replaceFormatWeekday(format string) string {
	var _format = format
	if strings.Contains(_format, "E") {
		_format = strings.Replace(_format, "E", "Monday", -1)
	}
	if strings.Contains(_format, "e") {
		_format = strings.Replace(_format, "e", "Mon", -1)
	}
	return _format
}

func (d Datetime) replaceFormatHour(format string) string {
	var _format = format
	if strings.Contains(_format, "HH") {
		_format = strings.Replace(_format, "HH", "15", -1)
	}
	if strings.Contains(_format, "hh") {
		_format = strings.Replace(_format, "hh", "03", -1)
	}
	return _format
}

func (d Datetime) replaceFormatMinute(format string) string {
	var _format = format
	if strings.Contains(_format, "mm") {
		_format = strings.Replace(_format, "mm", "04", -1)
	}
	if strings.Contains(_format, "m") {
		_format = strings.Replace(_format, "m", "4", -1)
	}
	return _format
}

func (d Datetime) replaceFormatSecond(format string) string {
	var _format = format
	if strings.Contains(_format, "ss") {
		_format = strings.Replace(_format, "ss", "05", -1)
	}
	if strings.Contains(_format, "s") {
		_format = strings.Replace(_format, "s", "5", -1)
	}
	if strings.Contains(_format, "SSSSSSSSS") {
		_format = strings.Replace(_format, "SSSSSSSSS", "000000000", -1)
	}
	if strings.Contains(_format, "SSSSSS") {
		_format = strings.Replace(_format, "SSSSSS", "000000", -1)
	}
	if strings.Contains(_format, "SSS") {
		_format = strings.Replace(_format, "SSS", "000", -1)
	}
	return _format
}

func (d Datetime) replaceFormatAMPM(format string) string {
	var _format = format
	if strings.Contains(_format, "a") {
		_format = strings.Replace(_format, "a", "PM", -1)
	}
	// if strings.Contains(_format, "a") {
	// 	_format = strings.Replace(_format, "a", "pm", -1)
	// }
	return _format
}

func (d Datetime) replaceFormatTZ(format string) string {
	var _format = format
	if strings.Contains(_format, "z") {
		_format = strings.Replace(_format, "z", "GMT", -1)
	}
	if strings.Contains(_format, "Z") {
		_format = strings.Replace(_format, "Z", "-0700", -1)
	}
	return _format
}
