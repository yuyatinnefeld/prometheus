# Scrape EC2 Instance with Prometheus

## Setup EC2
0. create a secrity group
```bash
name: allow-all-tcp
security group: allow-all-tcp (Type: SSH, Range: 22, CIDR:0.0.0.0/0) + (Type: All TCP, Range: 0-65535, CIDR:0.0.0.0/0)
```
1. spinning up the EC2
```bash
server: e.g. ubuntu server 22.04
security group: allow-all-tcp
key-pair: ec2_key_pair
```
2. install prometheus export in EC2 instance
```bash
# 35.157.111.164 [Public IP]
# user: ubuntu -> if linux then user: ec2-user

chmod 400 ec2_key_pair.pem
ssh -i ec2_key_pair.pem ubuntu@35.157.111.164

ubuntu@ip-172-31-34-250:~$ sudo apt update
ubuntu@ip-172-31-34-250:~$ sudo apt install prometheus-node-exporter
ubuntu@ip-172-31-34-250:~$ sudo systemctl status prometheus-node-exporter
```
3. create a new user in IAM + group (name: prometheus, iam: ec2-read-only access)
4. create an access key
## Setup Prometheus
4. update promtheus.yml
5. run prometheus `docker-compose up -d`
6. verify service discovery `open http://localhost:9090/service-discovery?search=`
7. verify target `open http://localhost:9090/targets`
8. verify the metircs page `open http://35.157.111.164:9100/metrics`