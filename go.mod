module github.com/mongodb/grip

go 1.20

require (
	github.com/andygrunwald/go-jira v1.14.0
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf
	github.com/dghubble/oauth1 v0.7.2
	github.com/fuyufjh/splunk-hec-go v0.3.4-0.20190414090710-10df423a9f36
	github.com/google/go-github/v53 v53.0.0
	github.com/mattn/go-xmpp v0.0.0-20210723025538-3871461df959
	github.com/montanaflynn/stats v0.0.0-20180911141734-db72e6cae808
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.8.3
	github.com/trivago/tgo v1.0.7
	go.opentelemetry.io/otel v1.16.0
	go.opentelemetry.io/otel/trace v1.16.0
)

require (
	github.com/ProtonMail/go-crypto v0.0.0-20230217124315-7d5c6f04bbb8 // indirect
	github.com/PuerkitoBio/rehttp v1.1.0 // indirect
	github.com/cloudflare/circl v1.3.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/golang-jwt/jwt v3.2.1+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/sabhiram/go-gitignore v0.0.0-20210923224102-525f6e181f06 // indirect
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	github.com/tklauser/numcpus v0.4.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	go.opentelemetry.io/otel/sdk v1.15.1 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/oauth2 v0.8.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/evergreen-ci/utility v0.0.0-20230519193518-4d91d64f59fb
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	// TODO (EVG-18584): gopsutil cannot be upgraded any further because the newer releases rely on a go1.17-only
	// feature. This should not be upgraded until the completion of EVG-18584.
	github.com/shirou/gopsutil/v3 v3.22.3
	github.com/slack-go/slack v0.12.1
	golang.org/x/sys v0.8.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
)
