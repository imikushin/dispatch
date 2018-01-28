///////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package proto

import (
	"container/heap"
	"sync"

	log "github.com/sirupsen/logrus"
)

type Scheduler struct {
	incomingInvocations <-chan Invocation
	driverQueue *DriverPriorityQueue

	driversAvailable *sync.Cond
}

type Invocation struct {
	ID int
}

func (inv Invocation) Start(driverName string) {
	log.Debugf("Driver '%s', running invocation: %v", driverName, inv.ID)
	// TODO insert some timeout to test work queue saturation
}

type Driver struct {
	Name      string
	Cost      int
	WorkQueue chan Invocation

	HasSpace *sync.Cond
}

func NewDriver(name string, cost int, bufferSize int) *Driver {
	return &Driver{
		Name:      name,
		Cost:      cost,
		WorkQueue: make(chan Invocation, bufferSize),
		HasSpace:  sync.NewCond(&sync.Mutex{}),
	}
}

func (d *Driver) Run() {
	for inv := range d.WorkQueue {
		inv.Start(d.Name)
		d.HasSpace.Broadcast() // the driver has free space in its work queue
	}
}

func (s *Scheduler) Add(driver *Driver) {
	s.driversAvailable.L.Lock()
	defer s.driversAvailable.L.Unlock()

	heap.Push(s.driverQueue, driver)
	s.driversAvailable.Signal()
}

func (s *Scheduler) addWhenAvailable(driver *Driver) {
	driver.HasSpace.L.Lock()
	defer driver.HasSpace.L.Unlock()

	driver.HasSpace.Wait()

	s.Add(driver)
}

func (s *Scheduler) Run() {
	for inv := range s.incomingInvocations {
		s.schedule(inv)
	}
}

func (s *Scheduler) schedule(inv Invocation) {
	s.driversAvailable.L.Lock()
	defer s.driversAvailable.L.Unlock()

	// wait until there are available drivers in the driver queue
	for len(*s.driverQueue) == 0 {
		s.driversAvailable.Wait()
	}

	// pick the lowest cost driver
	driver := (*s.driverQueue)[0]

	// put the invocation into the driver's work queue unless it is full
	select {
	case driver.WorkQueue <- inv:
	default:
		// remove the driver from driver queue, only put it back when its queue has free space
		heap.Pop(s.driverQueue)
		go s.addWhenAvailable(driver)
	}
}
