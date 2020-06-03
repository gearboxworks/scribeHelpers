package toolDocker

import (
	"fmt"
	"github.com/docker/docker/client"
	"net/url"
)


//func (gear *Docker) IsErrContainerNotFound(err error) bool {
//	return client.IsErrContainerNotFound(err)
//}


func (d *Docker) IsErrConnectionFailed(err error) bool {
	return client.IsErrConnectionFailed(err)
}


func (d *Docker) IsErrNotFound(err error) bool {
	return client.IsErrNotFound(err)
}


func (d *Docker) IsErrPluginPermissionDenied(err error) bool {
	return client.IsErrPluginPermissionDenied(err)
}


func (d *Docker) IsErrUnauthorized(err error) bool {
	return client.IsErrUnauthorized(err)
}


func (d *Docker) ParseHostURL(format string, args ...interface{}) (*url.URL, error) {
	return client.ParseHostURL(fmt.Sprintf(format, args...))
}


func ParseHostURL(format string, args ...interface{}) (*url.URL, error) {
	return client.ParseHostURL(fmt.Sprintf(format, args...))
}
