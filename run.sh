DockerHubUrl="localhost:5000"
Tag=$1


ImageName="$DockerHubUrl/gin-test1:${Tag}"

(
    export ImageName=$ImageName
    export NodePort=30080
    
    helm upgrade --install gin-test1 helm-chart/

    # docker-compose pull
    # docker-compose up -d
)