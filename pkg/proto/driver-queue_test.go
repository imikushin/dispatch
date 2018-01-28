///////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package proto

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDriverPriorityQueue_Push(t *testing.T) {
	dq := &DriverPriorityQueue{}
	heap.Init(dq)
	d1 := NewDriver("1", 1, 1)
	d2 := NewDriver("2", 2, 0)
	heap.Push(dq, d1)
	heap.Push(dq, d2)

	assert.Equal(t, d1, (*dq)[0])
}
