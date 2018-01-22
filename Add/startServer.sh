go build server.go
sudo docker build -f DockerFile_server . -t docker_server
sudo docker run -it -p 50051:50051 docker_server

