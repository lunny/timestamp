// Copyright 2018 Lunny Xiao. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package timestamp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeStamp(t *testing.T) {
	ti := time.Now()
	tm := TimeStamp(ti.Unix())
	assert.EqualValues(t, ti.Unix(), tm)
	assert.EqualValues(t, ti.Format("2006-01-02 15:04:05"), tm.Format("2006-01-02 15:04:05"))
}

func TestTimeStampNano(t *testing.T) {
	ti := time.Now()
	tm := TimeStampNano(ti.UnixNano())
	assert.EqualValues(t, ti.UnixNano(), tm)
	assert.EqualValues(t, ti.Format("2006-01-02 15:04:05"), tm.Format("2006-01-02 15:04:05"))
}
