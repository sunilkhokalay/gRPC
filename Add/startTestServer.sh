echo "Enter server IP"
read IP
go build Test.go
sudo docker build --build-arg "SERVER_IP=$IP" -f DockerFile_ServerTest . -t docker_server_test
sudo docker run -it docker_server_test
