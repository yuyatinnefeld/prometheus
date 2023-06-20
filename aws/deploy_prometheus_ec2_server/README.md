# Use EC2 as Prometheus Server

## Copy Terraform Source
```bash
open https://github.com/yuyatinnefeld/aws/tree/master/terraform/ec2
```

## Gnerate AWS components via Terraform
```bash
terraform init
terraform plan
terraform apply --auto-approve
```
## Use SSH to connect a EC2 instance server
```bash
chmod 400 ec2-keypair.pem
ssh -i ec2-keypair.pem ubuntu@10.0.53.26 # Public IPv4 address
ubuntu@10.0.53.26:~$ 
```

## Install prometheus
- https://prometheus.io/download/

```bash
# download tar file
export PROMETHEUS_VERSION="prometheus-2.44.0.linux-amd64"

wget https://github.com/prometheus/prometheus/releases/download/v2.44.0/${PROMETHEUS_VERSION}.tar.gz

# untar
tar xvf ${PROMETHEUS_VERSION}.tar.gz

# move to prometheus dir
cd ${PROMETHEUS_VERSION}

# start prometheus
./prometheus

# open 2nd termianl and verify the prometheus is running
ssh -i ec2-keypair.pem ubuntu@54.93.189.128 # Public IPv4 address
curl http://localhost:9090/metrics
```

## Manage prometheus server with systemd

```bash
ssh -i ec2-keypair.pem ubuntu@54.93.189.128 # Public IPv4 address
bash ./systemd-setup.sh

# start promtheus server for testing
sudo -u prometheus /usr/local/bin/prometheus \
    --config.file /etc/prometheus/prometheus.yml \
    --storage.tsdb.path /var/lib/prometheus/ \
    --web.console.templates=/etc/prometheus/consoles \
    --web.console.libraries=/etc/prometheus/console_libraries

curl http://localhost:9090/metrics

# kill the process
ubuntu@ip-10-0-53-26:~$ ps -x
    PID TTY      STAT   TIME COMMAND
   1223 ?        Ss     0:00 /lib/systemd/systemd --user
   ....
   1326 pts/0    Sl+    0:00 ./prometheus

ubuntu@ip-10-0-53-26:~$ kill 1326

# create an unit service file for prometheus server
sudo mkdir 
sudo cp prometheus.service /etc/systemd/system/prometheus.service
sudo systemctl daemon-reload
sudo systemctl start prometheus
sudo systemctl stop prometheus
```
