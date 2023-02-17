module github.com/mongodb/grip

go 1.16

require (
	github.com/andygrunwald/go-jira v1.14.0
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf
	github.com/dghubble/oauth1 v0.7.2
	github.com/fuyufjh/splunk-hec-go v0.3.4-0.20190414090710-10df423a9f36
	github.com/google/go-github v17.0.0+incompatible
	github.com/mattn/go-xmpp v0.0.0-20210723025538-3871461df959
	github.com/montanaflynn/stats v0.0.0-20180911141734-db72e6cae808
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.8.1
	github.com/trivago/tgo v1.0.7
	golang.org/x/oauth2 v0.0.0-20211005180243-6b3c2da341f1
)

require (
	github.com/evergreen-ci/utility v0.0.0-20230216205613-b8156d58f1e5
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	// TODO (EVG-18584): gopsutil cannot be upgraded any further because the newer releases rely on a go1.17-only
	// feature. This should not be upgraded until the completion of EVG-18584.
	github.com/shirou/gopsutil/v3 v3.23.1
	github.com/slack-go/slack v0.12.1
	google.golang.org/appengine v1.6.7 // indirect
)
