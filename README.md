# kamp

Amplify your developer impact by developing directly in Kubernetes.

## kamp run


#### Run an arbitrary docker container in Kubernetes (default to /bin/bash)

```bash
kamp run <image>:<tag>
```

#### Run an arbitrary docker container in Kubernetes with a custom command

```bash
kamp run <image>:<tag> -c /bin/csh
```

#### Run an arbitrary docker container in Kubernetes with a custom command and namespace

```bash
kamp run <image>:<tag> -c /bin/csh -n <namespace>
```

#### Run an arbitrary docker container in Kubernetes with a CIFS volume mount from your local host

```bash
kamp run <image>:<tag> -v $GOPATH/src/github.com/Nivenly/kamp:/root/kamp
```

## kamp log

#### Tail all kubernetes logs for a given namespace

```bash
kamp log * -n <namespace>
```

#### Tail all logs that match a regex query for a given namespace

```bash
kamp log *apiserver* -n <namespace>
```

## kamp push

#### Tag the container that is running in Kubernetes, and push it to a docker registry

```bash
kamp push <newimage>:<newtag>
```


