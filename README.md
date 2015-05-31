Example Code for "20 Minutes to Production, Zero Downtime"
==========================================================

Set up for demo
--------------------

Make sure the following is in /etc/hosts:

```
10.10.10.101 api-blue.qcon.demo
10.10.10.101 api-green.qcon.demo
10.10.10.101 api.qcon.demo
```

Start up Vagrant/CoreOS cluster:

```
cd demo_env/vagrant
vagrant up
cd ../..
```

Now, for each terminal window, add the Vagrant ssh keys, set the Fleet tunneling variable, and use a simplified prompt (for live demos): 

```
source source_me
```

Let's start the services for the demo:
```
bin/launch_demo_env
```

The included docker registry is insecure, so you'll need to let your local docker daemon know it is ok to use it. You need to add `--insecure-registry 10.10.10.103:5000` to your daemon's startup parameters. The way you do this varies by OS, but here are a few pointers:

- On Ubuntu 15:04+ (with systemd): http://stackoverflow.com/questions/29430024/docker-registry-with-insecure-registry-and-docker-1-5
- On Ubuntu (with upstart): Edit /etc/default/docker
- With boot2docker: Search "insecure-registry" here: https://github.com/boot2docker/boot2docker

Finally, let's create the docker image, push it to our registry, and deploy it:
```
make release
bin/launch_service
```

Test at:

http://api-blue.qcon.demo/stores/health/gtins/check

http://api.qcon.demo/stores/health/gtins/check


