docker build -t lnq99/sn-gateway:1.0 --platform=linux/amd64 gateway-service
docker build -t lnq99/sn-profiles:1.0 --platform=linux/amd64 profiles-service
docker build -t lnq99/sn-posts:1.0 --platform=linux/amd64 posts-service
docker build -t lnq99/sn-stats:1.0 --platform=linux/amd64 stats-service
docker build -t lnq99/sn-web:1.0 --platform=linux/amd64 client

docker push lnq99/sn-gateway:1.0
docker push lnq99/sn-profiles:1.0
docker push lnq99/sn-posts:1.0
docker push lnq99/sn-stats:1.0
docker push lnq99/sn-web:1.0