dockerName=blog_server
version=v1.0
registry=huanggaoqing/blog_server

sudo docker tag $dockerName:$version $registry:$dockerName-$version
docker push $registry:$dockerName-$version