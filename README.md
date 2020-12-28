# zookeeper-test
zookeeper golang client test

https://github.com/go-zookeeper/zk

테스트하는 Zookeeper Cluster zNode 초기 상태
```
/
├── zookeeper
│   ├── config
│   └── quota
├── temp
└── temp_child
```

### Tip

### CreateContainer()
- flag에 TTL이포함되어야 한다. 그렇지 않으면 Invalid Flag Error

### ACL
PermAll로 설정한 경우 (DigestACL은 user:passwd 설정) 리턴되는 ACL 형태
- AuthACL()
  - [{Perms:31 Scheme:auth ID:}]
- DigestACL()
  - [{Perms:31 Scheme:digest ID:user:MkkAAWC5ibFpYWKu1Zr/JdiwisA=}]
- WorldACL()
  - [{Perms:31 Scheme:world ID:anyone}]

### Method Postfix "W"의 의미
`GetW(), ExistW(), ChildrenW()`처럼 기존에 존재하는 메서드인 `Get(), Exist(), Children()`의 뒤에 W가 붙은 형태는 추가적으로 Watch를 설정하는 메서드


### Create()
- Parent Node까지 한번에 생성은 불가능하다.
  - "/parent" 노드가 없는 상황에서, "/parent/child" 노드 생성은 불가능
    
### FLW
FLW는 Four-letter-word로 zookeeper에서 정의한 4글자 명령어이다. conf, ruok, srvr, stat.. 등이 있다.

### Sync()
- Zookeeper 앙상블에서 Leader, Follower간 Sync의 인자로 준 path 노드를 동기화하는 메서드