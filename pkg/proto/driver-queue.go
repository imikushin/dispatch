///////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package proto

type DriverPriorityQueue []*Driver

func (pq DriverPriorityQueue) Len() int { return len(pq) }

func (pq DriverPriorityQueue) Less(i, j int) bool {
	return pq[i].Cost < pq[j].Cost
}

func (pq DriverPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *DriverPriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Driver))
}

func (pq *DriverPriorityQueue) Pop() interface{} {
	i := len(*pq) - 1
	defer func() { *pq = (*pq)[0:i] }()
	return (*pq)[i]
}
