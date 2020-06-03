package gearConfig

/*
Command execution is complex and there's several steps to the logic.
Essentially, the ENTRYPOINT image definition is converted to an S6 service.

During build:
1. The ENTRYPOINT definition within a Docker image needs to be pulled in to a Gearbox image.
2. This is used to start any service that was defined with ENTRYPOINT within the original image.
3. GearBuild.Run will contain index 0 of the ENTRYPOINT array.
4. GearBuild.Args will contain slice [1:] of the ENTRYPOINT array.

During runtime(boot):
1. GEARBOX_ENTRYPOINT, (aka GearBuild.Run), will be checked and executed as part of an S6 service.
2. GEARBOX_ENTRYPOINT_ARGS, (GearBuild.Args), will be appended and the whole service started.

During runtime(interactive command via symlink):
1. ARG 1 of the command line will be checked against GearRun.Commands for every container on the system.
2. When found, will execute.

During runtime(other interactive commands):
1.


GearBuild.Run
	This will default to GEARBOX_ENTRYPOINT env from within the image build process.
It is generated from the command: `docker inspect --format '{{ with .ContainerConfig.Entrypoint}} {{ index . 0 }}{{ end }}'`

GearBuild.Args
	This will default to GEARBOX_ENTRYPOINT_ARGS env from within the image build process.
It is generated from the command: `docker inspect --format '{{ join .ContainerConfig.Entrypoint " " }}'`
Any additional arguments provided by the user will be appended to this at runtime.
*/
