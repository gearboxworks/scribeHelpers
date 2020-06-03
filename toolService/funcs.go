package toolService

/*
# Windows query services:
	# https://godoc.org/golang.org/x/sys/windows/svc/mgr
	# https://stackoverflow.com/questions/13878921/how-to-get-all-windows-service-names-starting-with-a-common-word
	# https://stackoverflow.com/questions/12172997/how-to-collect-each-service-name-and-its-status-in-windows
sc query | findstr SERVICE_NAME

for /f "tokens=2" %s in ('sc query state^= all ^| find "SERVICE_NAME"') do
    @(for /f "tokens=4" %t in ('sc query %s ^| find "STATE     "') do @echo %s is %t)

sc queryex type= service state= all | find /i "NATION"

# Darwin:

# Linux:
systemctl show --property ActiveState docker

# Generic:
docker info --format '{{json .}}'

*/
