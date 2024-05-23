# module-nine

- terminal 1
sudo apt install docker.io
sudo docker pull prom/prometheus
sudo docker network create network
sudo docker volume create prometheus-data
sudo docker container run --name prometheus -v prometheus.yml -v prometheus-data:/prometheus -p 9090:9090 --network network prom/prometheus

- terminal 2
sudo docker pull grafana/grafana
sudo docker volume create grafana-data
sudo docker container run --name grafana -v grafana-data -p 3000:3000 --network network grafana/grafana

- terminal 3
git clone https://github.com/brumah/module-nine
sudo docker build -t module-nine .
sudo docker run -p 8080:8080 module-nine

- make requests to port 8080
- open port 9090 & 3000
- view parameters on prometheus
- add prometheus data source on grafana
- add visualization
