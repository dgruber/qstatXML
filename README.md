qstatXML
========

Simple example how to create an application which creates customized qstat output of Univa Grid Engine based on qstat -xml

It is based on the gestatus subpackage of the DRMAA1 Go implemenation.

In order to use it you need to set the library path to the C DRMAA library of Univa Grid Engine.

Example:

    source $SGE_ROOT/$SGE_CELL/common/settings.sh
    export LD_LIBRARY_PATH=$SGE_ROOT/lib/lx-amd64

    ./qstatXML <jobid>

	Job Name: sleep
	Job Number: 5
	Job Script: sleep
	Job Args: [ 123456]
	Job Owner: daniel
	Job Group: daniel
	Job UID: 1000
	Job GID: 1000
	Job accounting string: sge
	Job is now: false
	Job is binary: true
	Job has reservation: false
	Job is array job: false
	Job merges stderr false
	Job has 'no shell' requested: false
	Job has memory binding: false
	Job memory binding: no_bind
	Job submission time: 2014-09-17 10:49:55 +0200 CEST
	Job start time: 2014-09-17 10:49:58 +0200 CEST
	Job run time: 56m8.400789s
	Job deadline: 1970-01-01 01:00:00 +0100 CET
	Job mail options: 0
	Job AR: 0
	Job POSIX priority: 0
	Job class name: 
	Job mailing adresses: [ daniel@mint14]
	Job destination queue instance list: [                all.q@u1010 all.q@mint14]
	Job slots list: [                4 3]
	Job destination host List: [                u1010 mint14]
	Job slots sum: 7
	Job tasks: 1
	Job PE request: mytestpe 7-7:1
	Job resource usage cpu: 0.000000
	Job resource usage mem: 0.000000
	Job resource usage io: 0.000012
	Job resource usage iow: 0.000000
	Job resource usage vmem: 13983744.000000
	Job resource usage maxvmem: 13983744.000000
