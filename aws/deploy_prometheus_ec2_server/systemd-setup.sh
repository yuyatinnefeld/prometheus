# create a prom user
sudo useradd --no-create-home --shell /bin/false prometheus

# create dirs
sudo mkdir /etc/prometheus
sudo mkdir /var/lib/prometheus

# update permissions
sudo chown prometheus:prometheus /etc/prometheus
sudo chown prometheus:prometheus /var/lib/prometheus

# move the prom files
cd ${PROMETHEUS_VERSION}
sudo cp prometheus /usr/local/bin
sudo cp promtool /usr/local/bin
sudo cp -r consoles /etc/prometheus
sudo cp -r console_libraries /etc/prometheus
sudo cp prometheus.yml /etc/prometheus/prometheus.yml

# update permission
sudo chown prometheus:prometheus /usr/local/bin/prometheus
sudo chown prometheus:prometheus /usr/local/bin/promtool
sudo chown -R prometheus:prometheus /etc/prometheus/consoles
sudo chown -R prometheus:prometheus /etc/prometheus/console_libraries
sudo chown prometheus:prometheus /etc/prometheus/prometheus.yml