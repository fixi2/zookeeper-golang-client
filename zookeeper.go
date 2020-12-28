package zookeeper

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"log"
	"time"
)

func ConnectZooKeeper(connectAddr ...string) (*zk.Conn, error) {
	conn, _, err := zk.Connect(connectAddr, time.Second*10)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}

	return conn, err
}

func GetZnode(conn *zk.Conn, path string) ([]byte, *zk.Stat, error) {

	data, stat, err := conn.Get(path)

	return data, stat, err
}

func GetChildren(conn *zk.Conn, path string) ([]string, *zk.Stat, error) {

	data, stat, err := conn.Children(path)

	return data, stat, err
}

func Exist(conn *zk.Conn, path string) (bool, *zk.Stat, error) {

	exist, stat, err := conn.Exists(path)

	return exist, stat, err
}

func Print(name string, value interface{}) {

	switch v := value.(type) {
	case []byte:
		fmt.Printf("%s :\n%+v\n", name, string(v))
	default:
		fmt.Printf("%s :\n%+v\n", name, v)
	}
}