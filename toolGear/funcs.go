package toolGear

import (
	"fmt"
	"github.com/docker/docker/client"
	"net/url"
)


//func (gear *TypeDockerGear) IsErrContainerNotFound(err error) bool {
//	return client.IsErrContainerNotFound(err)
//}


func (gear *TypeDockerGear) IsErrConnectionFailed(err error) bool {
	return client.IsErrConnectionFailed(err)
}


func (gear *TypeDockerGear) IsErrNotFound(err error) bool {
	return client.IsErrNotFound(err)
}


func (gear *TypeDockerGear) IsErrPluginPermissionDenied(err error) bool {
	return client.IsErrPluginPermissionDenied(err)
}


func (gear *TypeDockerGear) IsErrUnauthorized(err error) bool {
	return client.IsErrUnauthorized(err)
}


func (gear *TypeDockerGear) ParseHostURL(format string, args ...interface{}) (*url.URL, error) {
	return client.ParseHostURL(fmt.Sprintf(format, args...))
}


func ParseHostURL(format string, args ...interface{}) (*url.URL, error) {
	return client.ParseHostURL(fmt.Sprintf(format, args...))
}
