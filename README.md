## fire
This is a remote execute tools, function like sshd+ssh but offers some interesting function.

including 

* set timeout to a process.
* view which program are running.
* kill program which are just running.
* a simple file server inside

Use `fire -d` to start a daemon server, and use `fire -H host ...` to tell daemon server to run commands.

## How to use 
*step by step tutorial*

### simple *hello world*
1. `fire -d --port=8119` to start a daemon server
2. `fire -m run echo hello world`, if you see hello world, then that's everything goes well.

### show running program
1. `fire -m ps` and will see output like this(ID maybe different).

```
ID                   TIME       CMD        RUNNING
5tb2zu1a34           10:35:44   echo       false
```

### start a long running program
1. `JOB_ID=$(fire -b sh -c "while true; do sleep 1s && echo hello; done")`, use -b to start program in background.
2. var $JOB_ID specify the running program.
3. `fire -m ps` will see program is still running.
```
ID                   TIME       CMD        RUNNING
72c7tni2gl           10:43:19   sh         true
5tb2zu1a34           10:35:44   echo       false
```
4. `fire -m kill 72`, to kill program. (NOTICE, you just need to use the prefix of ID)

### a complex usage
```
#
# connect to example1.com and start a program, with timeout 2s
#
JOB_ID=`fire -H example1.com -t 2s -b sh -c "echo start; sleep 5s; echo done"`
#
# wait program exit
#
fire -m wait $JOB_ID
```

### help usage
    fire [OPTIONS] args ...

    Help Options:
      -h, --help=        Show this help message

    Type:
      -m, --type=        type [run|ps|wait] (run)

    Run:
      -H, --host=        host to connect (localhost)
      -t, --timeout=     time out [s|m|h] (0s)
      -b, --background   run in background
      -e, --env=         add env to runner,multi support. eg -e PATH=/bin -e TMPDIR=/tmp
          --dialtimeout= dial timeout,unit seconds (2s)

    Serve:
      -d, --daemon       run as server (false)
      -p, --port=        port to connect or server
          --fs=          open a http file server (/tmp/:/home/)
          

### unknown problem
BUG(ssx): heartbeat: When open heartbeat. When the target machine dies. fire -d quit. Don't know why.
