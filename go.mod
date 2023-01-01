module github.com/mongodb/grip

go 1.16

require (
	github.com/andygrunwald/go-jira v0.0.0-20170512141550-c8c6680f245f
	github.com/coreos/go-systemd v0.0.0-20160607160209-6dc8b843c670
	github.com/dghubble/oauth1 v0.7.2
	github.com/fuyufjh/splunk-hec-go v0.3.4-0.20190414090710-10df423a9f36
	github.com/google/go-github v17.0.0+incompatible
	github.com/mattn/go-xmpp v0.0.0-20161121012536-f4550b539938
	github.com/montanaflynn/stats v0.0.0-20180911141734-db72e6cae808
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.8.1
	github.com/trivago/tgo v1.0.7
	golang.org/x/oauth2 v0.0.0-20211005180243-6b3c2da341f1
)

require (
	github.com/fatih/structs v1.1.0 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	// TODO (EVG-18584): gopsutil cannot be upgraded any further because the newer releases rely on a go1.17-only
	// feature. This should not be upgraded until the completion of EVG-18584.
	github.com/shirou/gopsutil/v3 v3.21.12
	github.com/slack-go/slack v0.11.4
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	google.golang.org/appengine v1.6.7 // indirect
)
