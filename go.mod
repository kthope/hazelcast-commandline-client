module github.com/hazelcast/hazelcast-commandline-client

go 1.15

require (
	github.com/alecthomas/chroma v0.9.2
	github.com/c-bata/go-prompt v0.2.5
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510
	github.com/hazelcast/hazelcast-go-client v1.1.1-0.20211013140338-f97b91075f8d
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	gopkg.in/yaml.v2 v2.4.0
)

//github.com/c-bata/go-prompt v0.2.5 => github.com/yuce/go-prompt v0.2.7-0.20220122092443-6bb274e42657
replace github.com/c-bata/go-prompt v0.2.5 => /home/yuce/Work/Projects/go-prompt
