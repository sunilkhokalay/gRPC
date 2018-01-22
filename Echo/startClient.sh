echo "Enter server IP"
read IP
go build client.go
sudo docker build --build-arg "SERVER_IP=$IP" -f DockerFile_client . -t docker_client
sudo docker run -it docker_client
