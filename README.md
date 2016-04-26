[![Circle CI](https://circleci.com/gh/asarturas/dodin/tree/master.svg?style=svg)](https://circleci.com/gh/asarturas/dodin/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/asarturas/dodin)](https://goreportcard.com/report/github.com/asarturas/dodin)
# dodin - dynamic inventory provider for Ansible

Currently only Digital Ocean is supported, but contributions are welcome to enable it for other providers too. Ping me and I will help you to start up the development environment.

Features:
- [x] Group droplets by name with regular expression;
- [x] Unlimited custom groups;
- [x] Add variables to groups;
- [ ] Display full information for each droplet with `--host hostname` flag (by droplet name);
- [ ] Ability to generate static inventory file with extra flag.

## What is dodin for?

Ansible provides neat option to use to use [dynamic inventories][dynamic_inventories], instead of static inventory files.
It is useful when managing dynamic infrastructure, where static inventory file would need to be constantly modified in order to keep it up with nodes in the cloud.

## How to install dodin?

1. There is no binary distributed, so you have to have a golang runtime in order to use it;
2. Download sources and compile it via `go get github.com/asarturas/dodin/cmd/dodin-digital-ocean`.

## How to use dodin?

1. In your project directory create `dodin-digital-ocean.ini` file, where you define group hosts with regular expression and list group variables in a same way as you do with static inventories:
   
   ```
   [master]
   members=master\-[0-9]{2}
   
   [master:vars]
   ansible_ssh_user=core
   ansible_python_interpreter="PATH=/home/core/bin:$PATH python"
   
   [minion]
   members=minion\-[0-9]{2}
   
   [minion:vars]
   ansible_ssh_user=core
   ansible_python_interpreter="PATH=/home/core/bin:$PATH python"
   ```
   
2. Export your digital ocean api token like this `export DO_API_TOKEN=1234567890`;
3. Check dodin works correctly via `$GOPATH/bin/dodin-digital-ocean`, it should output json with all your nodes in all group;
4. Ping cluster hosts via `ansible cluster -m ping -i "$GOPATH/bin/dodin-digital-ocean"`.

## Why use dodin?

Unfortunately standard provided scripts are not very flexible.
For instance, when you have coreos cluster with 3 master and 3 minion hosts,
then without script modification you would not be able to group nodes logically nor you will be able to add custom variables the way you would do in static inventory file:
```
[master]
master-01
master-02
master-03

[master:vars]
ansible_ssh_user=core
ansible_python_interpreter="PATH=/home/core/bin:$PATH python"

[minion]
minion-01
minion-02
minion-03

[minion:vars]
ansible_ssh_user=core
ansible_python_interpreter="PATH=/home/core/bin:$PATH python"
```
For such infrastructure standard digital ocean dynamic inventory script digital_ocean.py would group them operating system, size, region and put each in separate group by name and id.
The output would be like this:
```
{
    "12190784": ["178.62.29.42"],
    "distro_CoreOS": ["178.62.29.42", "188.166.150.62", "178.62.95.186", "178.62.96.225", "178.62.117.116", "46.101.5.85"],
    "master-01": ["178.62.29.42"],
    "size_512mb": ["178.62.29.42", "188.166.150.62", "178.62.95.186", "178.62.96.225", "178.62.117.116", "46.101.5.85"],
    "region_lon1": ["178.62.29.42", "188.166.150.62", "178.62.95.186", "178.62.96.225", "178.62.117.116", "46.101.5.85"],
    "12190825": ["178.62.117.116"],
    "12190827": ["46.101.5.85"],
    "status_active": ["178.62.29.42", "188.166.150.62", "178.62.95.186", "178.62.96.225", "178.62.117.116", "46.101.5.85"],
    "12190824": ["178.62.96.225"],
    "12190787": ["188.166.150.62"],
    "image_coreos-alpha": ["178.62.29.42", "188.166.150.62", "178.62.95.186", "178.62.96.225", "178.62.117.116", "46.101.5.85"],
    "minion-03": ["46.101.5.85"],
    "minion-01": ["178.62.96.225"],
    "master-02": ["188.166.150.62"],
    "master-03": ["178.62.95.186"],
    "image_16335999": ["178.62.29.42", "188.166.150.62", "178.62.95.186", "178.62.96.225", "178.62.117.116", "46.101.5.85"],
    "minion-02": ["178.62.117.116"],
    "12190789": ["178.62.95.186"]
}
```
Dodin allows you to replicate that example static inventory file with simple ini config, which looks very similar to original.
The only difference that instead of listing all the hostnames in group, we're specifying regular expression pattern
to match node name:
```
[master]
members=master\-[0-9]{2}

[master:vars]
ansible_ssh_user=core
ansible_python_interpreter="PATH=/home/core/bin:$PATH python"

[minion]
members=minion\-[0-9]{2}

[minion:vars]
ansible_ssh_user=core
ansible_python_interpreter="PATH=/home/core/bin:$PATH python"
```
This given example would put into master group any node named master-00 to master-99.
Similarly it would put minion-00 to minion-99.
Note that variables are defined the same way as they would in static file.
The dodin-digital-ocean would output:
```
{
    "master": {
        "hosts":
            ["178.62.29.42","188.166.150.62","178.62.95.186"],
        "vars": {
            "ansible_python_interpreter":"PATH=/home/core/bin:$PATH python",
            "ansible_ssh_user":"core"
        }
    },
    "minion": {
        "hosts":
            ["178.62.96.225","178.62.117.116","46.101.5.85"],
        "vars": {
            "ansible_python_interpreter":"PATH=/home/core/bin:$PATH python",
            "ansible_ssh_user":"core"
        }
    }
}
```
Which is equivalent of the first one, but in dynamic format, expected by ansible-playbook.

## Disclaimer

The software is provided as is. I use it and am happy to help you setup or to resolve any issues with it, but you are taking all the responsility for using it.

[dynamic_inventories]: http://docs.ansible.com/ansible/intro_dynamic_inventory.html
