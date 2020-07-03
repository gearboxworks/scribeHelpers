package toolDocker

import (
	"fmt"
	"github.com/docker/docker/client"
	"net/url"
)


//func (gear *TypeDocker) IsErrContainerNotFound(err error) bool {
//	return client.IsErrContainerNotFound(err)
//}


func (d *TypeDocker) IsErrConnectionFailed(err error) bool {
	return client.IsErrConnectionFailed(err)
}


func (d *TypeDocker) IsErrNotFound(err error) bool {
	return client.IsErrNotFound(err)
}


func (d *TypeDocker) IsErrPluginPermissionDenied(err error) bool {
	return client.IsErrPluginPermissionDenied(err)
}


func (d *TypeDocker) IsErrUnauthorized(err error) bool {
	return client.IsErrUnauthorized(err)
}


func (d *TypeDocker) ParseHostURL(format string, args ...interface{}) (*url.URL, error) {
	return client.ParseHostURL(fmt.Sprintf(format, args...))
}


func ParseHostURL(format string, args ...interface{}) (*url.URL, error) {
	return client.ParseHostURL(fmt.Sprintf(format, args...))
}
