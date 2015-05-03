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
vagrant ssh-config
vagrant ssh-add {key from ssh-config}
export FLEETCTL_TUNNEL={ip and port from ssh-config}

cd ..
fleetctl start vulcand@{1..3}
cd ..
```

Launch service:

```
./bin/launch_service
```

Test at:

http://api-blue.qcon.demo/stores/health/gtins/check

http://api.qcon.demo/stores/health/gtins/check


