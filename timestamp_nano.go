// Copyright 2018 Lunny Xiao. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package timestamp

import (
	"time"
)

// TimeStampNano defines a timestamp to nano
type TimeStampNano int64

// TimeStampNow returns now int64
func TimeStampNanoNow() TimeStampNano {
	return TimeStampNano(time.Now().UnixNano())
}

// Add adds seconds and return sum
func (ts TimeStampNano) Add(seconds int64) TimeStampNano {
	return ts.AddDuration(time.Duration(seconds) * time.Second)
}

// AddDuration adds time.Duration and return sum
func (ts TimeStampNano) AddDuration(interval time.Duration) TimeStampNano {
	return ts + TimeStampNano(interval)
}

// Year returns the time's year
func (ts TimeStampNano) Year() int {
	return ts.AsTime().Year()
}

// AsTime convert timestamp as time.Time in Local locale
func (ts TimeStampNano) AsTime() (tm time.Time) {
	tm = time.Unix(int64(ts)/1e9, int64(ts)%1e9).In(DefaultLocation)
	return
}

// AsTimePtr convert timestamp as *time.Time in Local locale
func (ts TimeStampNano) AsTimePtr() *time.Time {
	tm := time.Unix(int64(ts)/1e9, int64(ts)%1e9).In(DefaultLocation)
	return &tm
}

// Format formats timestamp as
func (ts TimeStampNano) Format(f string) string {
	return ts.AsTime().Format(f)
}

// FormatLong formats as RFC1123Z
func (ts TimeStampNano) FormatLong() string {
	return ts.Format(time.RFC1123Z)
}

// FormatShort formats as short
func (ts TimeStampNano) FormatShort() string {
	return ts.Format("Jan 02, 2006")
}

// IsZero is zero time
func (ts TimeStampNano) IsZero() bool {
	return ts.AsTime().IsZero()
}
