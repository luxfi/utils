// Copyright (C) 2019-2025, Lux Industries, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package perms provides common file permission constants.
package perms

import "os"

// File permissions
const (
	ReadOnly         = 0o400
	ReadWrite        = 0o640
	ReadWriteExecute = 0o750
)

// Create creates a file with the given permissions.
func Create(name string, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
}
