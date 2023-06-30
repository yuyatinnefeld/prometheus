# Install Alert Manager

```bash
wget https://github.com/prometheus/alertmanager/releases/download/v0.21.0/alertmanager-0.21.0.linux-amd64.tar.gz
tar xzf alertmanager-0.21.0.linux-amd64.tar.gz
```

# Setup

```bash
groupadd -f alertmanager
useradd -g alertmanager --no-create-home --shell /bin/false alertmanager
mkdir -p /etc/alertmanager/templates
mkdir /var/lib/alertmanager
chown alertmanager:alertmanager /etc/alertmanager
chown alertmanager:alertmanager /var/lib/alertmanager

mv alertmanager-0.21.0.linux-amd64 alertmanager-files

cp alertmanager-files/alertmanager /usr/bin/
cp alertmanager-files/amtool /usr/bin/
chown alertmanager:alertmanager /usr/bin/alertmanager
chown alertmanager:alertmanager /usr/bin/amtool

cp alertmanager-files/alertmanager.yml /etc/alertmanager/alertmanager.yml
chown alertmanager:alertmanager /etc/alertmanager/alertmanager.yml

vi /usr/lib/systemd/system/alertmanager.service

```

# Configure alertmanager.service

```bash
[Unit]
Description=AlertManager
Wants=network-online.target
After=network-online.target

[Service]
User=alertmanager
Group=alertmanager
Type=simple
ExecStart=/usr/bin/alertmanager \
    --config.file /etc/alertmanager/alertmanager.yml \
    --storage.path /var/lib/alertmanager/

[Install]
WantedBy=multi-user.target
```

# Update systemctl
```bash
chmod 664 /usr/lib/systemd/system/alertmanager.service
systemctl daemon-reload
systemctl start alertmanager
systemctl status alertmanager
```