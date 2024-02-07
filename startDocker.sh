sudo docker stop setu-engine && sudo docker rm setu-engine
sudo docker build -t setu-engine .
sudo docker run -d -p 8070:8070 --name setu-engine setu-engine