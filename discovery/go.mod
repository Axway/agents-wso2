module github.com/Axway/agents-wso2/discovery

go 1.15

require (
	github.com/Axway/agent-sdk v0.0.20-0.20210128210152-8f30d2013836
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/getkin/kin-openapi v0.36.0 // indirect
	github.com/go-openapi/swag v0.19.13 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20200217142428-fce0ec30dd00 // indirect
	github.com/magiconair/properties v1.8.4 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/sirupsen/logrus v1.7.0
	github.com/smartystreets/assertions v1.2.0 // indirect
	github.com/spf13/afero v1.5.1 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.1.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/tidwall/gjson v1.6.7
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	golang.org/x/text v0.3.5 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace (
	github.com/Shopify/sarama => github.com/elastic/sarama v0.0.0-20191122160421-355d120d0970
	github.com/dop251/goja => github.com/andrewkroh/goja v0.0.0-20190128172624-dd2ac4456e20
	github.com/fsnotify/fsevents => github.com/fsnotify/fsevents v0.1.1
)
