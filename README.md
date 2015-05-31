Example Code for "20 Minutes to Production, Zero Downtime"
==========================================================

Prerequirements
--------------------

You will need Vagrant and Docker installed. This has been primarily tested on Ubuntu.

Make sure the following is in /etc/hosts:

```
10.10.10.101 api-blue.qcon.demo
10.10.10.101 api-green.qcon.demo
10.10.10.101 api.qcon.demo
```

The included docker registry is insecure, so you'll need to let your local docker daemon know it is ok to use it. You need to add `--insecure-registry 10.10.10.103:5000` to your daemon's startup parameters. The way you do this varies by OS, but here are a few pointers:

- On Ubuntu 15:04+ (with systemd): http://stackoverflow.com/questions/29430024/docker-registry-with-insecure-registry-and-docker-1-5
- On Ubuntu (with upstart): Edit /etc/default/docker
- With boot2docker: Search "insecure-registry" here: https://github.com/boot2docker/boot2docker


### Starting up the demo environment

Start up the demo environment.

```
demo_env/bin/start
```

This will use Vagrant to start up a cluster of CoreOS VMs and schedule the Docker images we need to the VMs using fleet. Specifically, vulcand, docker-registry and our qcon-gtin service.


### Running the demo

For each new terminal window, source the following script to add the Vagrant ssh keys, set the Fleet tunneling variable, and use a simplified prompt (for live demos): 

```
source source_me
```

Test at:

http://api-blue.qcon.demo/stores/health/gtins/check
http://api.qcon.demo/stores/health/gtins/check

#### Update the qcon-gtin service

Modify whatever you'd like in the `main.go` source for the service. For the demo, we're going to add this:

```
 if len(gtinParam) == 13 {
   gtinParam = strings.TrimPrefix(gtinParam, "0")
 }
```

Note: If you want to be able to deploy multiple versions in your cluster, give this new build a new version by ticking the version number in `Makefile` and `gtin@.service` 

Build:
``` 
make
```

Packing up in Docker and push to Docker registry:
```
make release
```

Deploy the GTIN service (with zero downtime):
```
cd units
fleetctl destroy gtin.proxy.public@1.service
fleetctl destroy gtin.proxy.private@1.service
fleetctl destroy gtin@1.service
fleetctl start gtin@1.service
fleetctl start gtin.proxy.private@1.service
fleetctl start gtin.proxy.public@1.service
```

Or, use the alias:
```
redeploy-service gtin 2
redeploy-service gtin 3
```

