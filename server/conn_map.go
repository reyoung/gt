package server

import (
	"net"
	"sync"
)

type connMap struct {
	conns  map[uint32]net.Conn
	mutex  sync.RWMutex
	serial uint32
}

func (c *connMap) Put(conn net.Conn) uint32 {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.serial++
	c.conns[c.serial] = conn
	return c.serial
}

func (c *connMap) Get(serial uint32) net.Conn {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.conns[serial]
}

func (c *connMap) Delete(serial uint32) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.conns, serial)
}

func newConnMap() *connMap {
	return &connMap{
		conns: make(map[uint32]net.Conn),
	}
}
