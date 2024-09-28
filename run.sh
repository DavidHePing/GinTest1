DockerHubUrl="localhost:5000"
Tag=$1

helm upgrade --install helm-demo .

# ImageName="$DockerHubUrl/gin-test1:${Tag}"

# (
#     export ImageName=$ImageName

#     docker-compose pull
#     docker-compose up -d
# )