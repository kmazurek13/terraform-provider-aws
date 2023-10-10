// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package docdb

import (
	"time"
)

const (
	propagationTimeout = 2 * time.Minute
)

const (
	engineDocDB = "docdb"
)

func engine_Values() []string {
	return []string{
		engineDocDB,
	}
}

const (
	errCodeInvalidParameterValue = "InvalidParameterValue"
)
