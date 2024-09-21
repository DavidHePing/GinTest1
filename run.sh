docker build -t my-go-app . 
docker rm -f go-test1
docker run -d -p 8080:8080 --name go-test1 my-go-app 