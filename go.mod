module github.com/wwqdrh/ds-connect

go 1.18

replace (
	github.com/wwqdrh/nettool v0.0.0-20221024025010-9990cfd11043 => ../nettool
	github.com/wwqdrh/ostool v0.0.0-20221024032333-a4c63e32e8b4 => ../ostool
	github.com/wwqdrh/clitool v0.0.0-20221025031738-dfe3c08ac76f => ../clitool
)

require (
	github.com/docker/docker v20.10.19+incompatible
	github.com/fsouza/go-dockerclient v1.9.0
	github.com/spf13/cobra v1.6.1
	github.com/wwqdrh/logger v0.0.9
	github.com/wwqdrh/nettool v0.0.0-20221024025010-9990cfd11043
	github.com/wwqdrh/ostool v0.0.0-20221024032333-a4c63e32e8b4
)

require (
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/Microsoft/go-winio v0.6.0 // indirect
	github.com/Microsoft/hcsshim v0.9.3 // indirect
	github.com/containerd/cgroups v1.0.3 // indirect
	github.com/containerd/containerd v1.6.6 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/mitchellh/go-ps v1.0.0 // indirect
	github.com/moby/sys/mount v0.3.3 // indirect
	github.com/moby/sys/mountinfo v0.6.2 // indirect
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.3-0.20211202183452-c5a74bcca799 // indirect
	github.com/opencontainers/runc v1.1.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/wwqdrh/clitool v0.0.0-20221025031738-dfe3c08ac76f // indirect
	go.opencensus.io v0.23.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/sys v0.0.0-20220908164124-27713097b956 // indirect
	golang.org/x/tools v0.1.12 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/xjasonlyu/tun2socks/v2 v2.4.1 => github.com/linfan/tun2socks/v2 v2.4.2-0.20220501081747-6f4a45525a7c
