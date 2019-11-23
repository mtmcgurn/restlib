// Copyright 2019 Reaction Engineering International. All rights reserved.
// Use of this source code is governed by the MIT license in the file LICENSE.txt.

package cache

type Cache interface {
	Get(key string, item interface{}) bool

	Set(key string, item interface{}) error

	GetString(key string) (string, bool)

	SetString(key string, value string)
}
