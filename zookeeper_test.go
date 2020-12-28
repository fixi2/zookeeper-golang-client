package zookeeper_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	zookeeper "zookeeper-test"

	"github.com/go-zookeeper/zk"
)

var (
	conn           *zk.Conn
	paths, servers []string
)

func TestGetzNode(t *testing.T) {

	for _, path := range paths {
		fmt.Printf("[ %s ]\n", path)

		data, stat, err := zookeeper.GetZnode(conn, path)
		if err != nil {
			log.Println(err)
			continue
		}

		zookeeper.Print("data", data)
		zookeeper.Print("stat", stat)
	}
}

func TestChildren(t *testing.T) {

	for _, path := range paths {
		fmt.Printf("[ %s ]\n", path)

		data, stat, err := zookeeper.GetChildren(conn, path)
		if err != nil {
			log.Println(err)
			continue
		}

		zookeeper.Print("data", data)
		zookeeper.Print("stat", stat)
	}
}

func TestExist(t *testing.T) {

	for _, path := range paths {
		fmt.Printf("[ %s ]\n", path)

		exist, stat, err := zookeeper.Exist(conn, path)
		if err != nil {
			log.Println(err)
			continue
		}

		zookeeper.Print("exist", exist)
		zookeeper.Print("stat", stat)
	}
}

func TestState(t *testing.T) {

	state := conn.State()
	zookeeper.Print("state", state)
}

func TestServer(t *testing.T) {

	server := conn.Server()
	zookeeper.Print("server", server)
}

func TestGetACL(t *testing.T) {

	for _, path := range paths {
		fmt.Printf("[ %s ]\n", path)

		acl, stat, err := conn.GetACL(path)
		if err != nil {
			log.Println(err)
		}

		zookeeper.Print("acl", acl)
		zookeeper.Print("stat", stat)
	}
}

func TestMakeACL(t *testing.T) {

	acl := zk.AuthACL(zk.PermAll)
	zookeeper.Print("acl", acl)

	acl = zk.DigestACL(zk.PermAll, "user", "1234")
	zookeeper.Print("acl", acl)

	acl = zk.WorldACL(zk.PermAll)
	zookeeper.Print("acl", acl)
}

type node struct {
	path  string
	data  []byte
	flags int32
	acl   []zk.ACL
}

func TestCreateNode(t *testing.T) {

	nodes := []node{
		{
			path:  "/create_test2",
			data:  []byte("ct_data"),
			flags: 0,
			acl:   zk.WorldACL(zk.PermAll),
		},
		{
			path:  "/create_test2/ct_child",
			data:  []byte("ct_child_data"),
			flags: 0,
			acl:   zk.WorldACL(zk.PermAll),
		},
		{
			path:  "/create_test2/ct_child",
			data:  []byte("ct_child_data_2"),
			flags: 0,
			acl:   zk.WorldACL(zk.PermAll),
		}, // error : node already exist
	}

	for _, node := range nodes {
		str, err := conn.Create(node.path, node.data, node.flags, node.acl)
		if err != nil {
			log.Println(err)
			continue
		}

		zookeeper.Print("str", str)
	}
}

func TestCreateSequentialNode(t *testing.T) {

	var path string
	// path = "/sequential/seq_" // /sequential 없는 상황에서 node not exist
	// path = "/sequential" // 잘 생성 된다.
	path = "/create_test/ct_child" // 기존에 ch_child 있어도 잘 됨

	for i := 0; i < 10; i++ {
		str, err := conn.Create(path, []byte{}, zk.FlagSequence, zk.WorldACL(zk.PermAll))
		if err != nil {
			log.Println(err)
		}

		zookeeper.Print("path", str)
	}
}

func TestContainerNode(t *testing.T) {

	var path string
	path = "/container"

	str, err := conn.CreateContainer(path, []byte{}, zk.FlagTTL, zk.WorldACL(zk.PermAll))
	if err != nil {
		log.Println(err)
	}

	zookeeper.Print("path", str)

	data, stat, err := conn.Get(path)
	if err != nil {
		log.Println(err)
	}

	zookeeper.Print("data", data)
	zookeeper.Print("stat", stat)
}

func TestProtectedEphemeralSequential(t *testing.T) {

	var path string
	path = "/proceted_sequential"

	str, err := conn.CreateProtectedEphemeralSequential(path, []byte{}, zk.WorldACL(zk.PermAll))
	if err != nil {
		log.Println(err)
	}
	zookeeper.Print("str", str)

	str, err = conn.CreateProtectedEphemeralSequential(path, []byte{}, zk.WorldACL(zk.PermAll))
	if err != nil {
		log.Println(err)
	}
	zookeeper.Print("str", str)

	str, err = conn.CreateProtectedEphemeralSequential(path, []byte{}, zk.WorldACL(zk.PermAll))
	if err != nil {
		log.Println(err)
	}
	zookeeper.Print("str", str)

	children, stat, err := zookeeper.GetChildren(conn, "/")
	if err != nil {
		log.Println(err)
	}

	zookeeper.Print("children", children)
	zookeeper.Print("stat", stat)

}

func TestFLWRuok(t *testing.T) {

	state := zk.FLWRuok(servers, time.Second*5)

	zookeeper.Print("state", state)
}

func TestFLWCons(t *testing.T) {

	serverClients, state := zk.FLWCons(servers, time.Second)

	fmt.Println("serverClients:")
	for _, serverClient := range serverClients {
		fmt.Printf("%+v\n", serverClient)
	}
	zookeeper.Print("state", state)
}

func TestFLWSrvr(t *testing.T) {

	serverStats, state := zk.FLWSrvr(servers, time.Second)

	fmt.Println("serverStats:")
	for _, serverStat := range serverStats {
		fmt.Printf("%+v\n", serverStat)
	}
	zookeeper.Print("state", state)
}

type CustomLogger struct {}

func (cl CustomLogger) Printf(s string, v ...interface{}) {
	s2 := fmt.Sprintf(s, v...)
	fmt.Printf("Custom Logger: %s\n", s2)
}

func TestWithLogger(t *testing.T) {

	conn2, _, err := zk.Connect(servers, time.Second * 5, zk.WithLogger(CustomLogger{}))
	if err != nil {
		log.Println(err)
		return
	}

	_, err = conn2.Create("/temp", []byte{}, 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		log.Println(err)
	}
}

func TestSet(t *testing.T) {

	data, stat, err := conn.Get("/temp")
	zookeeper.Print("data", data)
	zookeeper.Print("stat", stat)

	stat, err = conn.Set("/temp", []byte("set_data"), 100)
	if err != nil {
		log.Println(err)
	}

	zookeeper.Print("stat", stat)

	stat, err = conn.Set("/temp", []byte("set_dataa"), 1)
	if err != nil {
		log.Println(err)
	}

	zookeeper.Print("stat", stat)
}

func init() {
	var err error

	servers = []string{"10.113.78.147:2181", "10.113.79.117:2181", "10.113.97.243:2181"}

	paths = []string{
		"/zookeeper",
		"/temp",
		"/temp/temp_child",
		"/invalid_path",
	}

	conn, _, err = zk.Connect(servers, time.Second*5)
	if err != nil {
		log.Fatalln(err)
	}

}
