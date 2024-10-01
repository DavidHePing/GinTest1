DockerHubUrl="localhost:5000"
Tag=$1


ImageName="$DockerHubUrl/gin-test1:${Tag}"

(
    helm upgrade --install gin-test1 helm-chart/ \
    --set service.nodePort=30080 \
    --set service.imageName=$ImageName

    # docker-compose pull
    # docker-compose up -d
)