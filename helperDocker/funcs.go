package helperDocker

import (
	"fmt"
	"github.com/docker/docker/client"
	"net/url"
)


//func (gear *DockerGear) IsErrContainerNotFound(err error) bool {
//	return client.IsErrContainerNotFound(err)
//}


func (gear *DockerGear) IsErrConnectionFailed(err error) bool {
	return client.IsErrConnectionFailed(err)
}


func (gear *DockerGear) IsErrNotFound(err error) bool {
	return client.IsErrNotFound(err)
}


func (gear *DockerGear) IsErrPluginPermissionDenied(err error) bool {
	return client.IsErrPluginPermissionDenied(err)
}


func (gear *DockerGear) IsErrUnauthorized(err error) bool {
	return client.IsErrUnauthorized(err)
}


func (gear *DockerGear) ParseHostURL(format string, args ...interface{}) (*url.URL, error) {
	return client.ParseHostURL(fmt.Sprintf(format, args...))
}


func ParseHostURL(format string, args ...interface{}) (*url.URL, error) {
	return client.ParseHostURL(fmt.Sprintf(format, args...))
}
