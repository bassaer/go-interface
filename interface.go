package main

import "fmt"

type Db interface {
	Read() string
}

type Cassandra struct {
}

type MySQL struct {
}

func (c *Cassandra) Read() string {
	return "this is Cassandra"
}

func (m *MySQL) Read() string {
	return "this is MySQL"
}

func Out(d Db) {
	fmt.Println(d.Read())
}

func main() {
	c := &Cassandra{}
	m := &MySQL{}
	Out(c)
	Out(m)
}
