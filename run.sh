DockerHubUrl="localhost:5000"
Tag=$1


ImageName="$DockerHubUrl/gin-test1:${Tag}"

(
    docker pull $ImageName

    helm upgrade --install gin-test1 helm-chart/ \
    --set service.nodePort=30080 \
    --set image.repository=$ImageName

    # docker-compose pull
    # docker-compose up -d
)