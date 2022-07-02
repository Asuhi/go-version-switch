# go-version-switch

## compile

`brew install go`
`go build`

## add version

`brew install go@[version]`
`sudo mv /opt/homebrew/Cellar/go@[verion]/[version] /opt/homebrew/Cellar/go/` 

## use

list all go version 
gvs will read `/opt/homebrew/Cellar/go/` path and print all dir
`gvs list`

change version
gvs will create soft links to `/usr/local/bin/`
`sudo gvs sw [version]`
