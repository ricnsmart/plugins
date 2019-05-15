module github.com/ricnsmart/plugins

go 1.12

require (
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db // indirect
	github.com/google/go-cmp v0.2.0 // indirect
	github.com/stretchr/testify v1.3.0 // indirect
	github.com/tidwall/pretty v0.0.0-20180105212114-65a9db5fad51 // indirect
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.0.0
	golang.org/x/crypto v0.0.0-20190426145343-a29dc8fdc734 // indirect
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.38.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190510104115-cbcb75029529
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190510132918-efd6b22b2522
	golang.org/x/image => github.com/golang/image v0.0.0-20190507092727-e4e5bf290fec
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190509164839-32b2708ab171
	golang.org/x/net => github.com/golang/net v0.0.0-20190509222800-a4d6f7feada5
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190402181905-9f3314589c9a
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190509141414-a5b02f93d862
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190511041617-99f201b6807e
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.5.0
	google.golang.org/appengine => github.com/golang/appengine v1.5.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190508193815-b515fa19cec8
	google.golang.org/grpc => github.com/grpc/grpc-go v1.20.1
	labix.org/v2/mgo => github.com/go-mgo/mgo v0.0.0-20180705113738-7446a0344b78
	launchpad.net/gocheck => github.com/go-check/check v0.0.0-20180628173108-788fd7840127
)
