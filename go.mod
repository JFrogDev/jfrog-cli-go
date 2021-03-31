module github.com/jfrog/jfrog-cli

go 1.14

require (
	github.com/buger/jsonparser v0.0.0-20180910192245-6acdf747ae99
	github.com/codegangsta/cli v1.20.0
	github.com/frankban/quicktest v1.11.3 // indirect
	github.com/jfrog/gocmd v0.1.19
	github.com/jfrog/gofrog v1.0.6
	github.com/jfrog/jfrog-cli-core v1.4.2
	github.com/jfrog/jfrog-client-go v0.20.2
	github.com/jszwec/csvutil v1.4.0
	github.com/mholt/archiver v2.1.0+incompatible
	github.com/pierrec/lz4 v2.6.0+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
	github.com/vbauerster/mpb/v4 v4.7.0
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/crypto v0.0.0-20201208171446-5f87f3452ae9
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/jfrog/jfrog-client-go => github.com/gailazar300/jfrog-client-go v0.18.1-0.20210331160120-7602aa14ef3e

replace github.com/jfrog/jfrog-cli-core => github.com/jfrog/jfrog-cli-core v1.4.3-0.20210330182328-a8186dce25ab

replace github.com/jfrog/gocmd => github.com/jfrog/gocmd v0.1.20-0.20210330134550-459be3fc0ee8

// replace github.com/jfrog/gofrog => github.com/jfrog/gofrog v1.0.6
