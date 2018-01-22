go build client.go
sudo docker build -f DockerFile_client . -t docker_client
sudo docker run -it docker_client
