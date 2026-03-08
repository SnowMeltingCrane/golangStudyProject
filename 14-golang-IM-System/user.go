package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

func NewUser(conn net.Conn) *User {

	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}

	go user.ListenMessager()

	return user
}

// 监听当前user channel的方法
func (this *User) ListenMessager() {

	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}

}
