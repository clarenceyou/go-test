GO Test
--------

## Docker image

### Build image
    docker build -t clarenceyou/gotest .

### Tag image
    docker tag [image-id] clarenceyou/gotest:1.0.0

### Run container
    docker run -d clarenceyou/gotest:1.0.0


minishift start --iso-url centos
minishift start --iso-url centos --docker-opt log-driver=journald
minishift start --iso-url centos --docker-opt log-driver=journald --logging

oc adm policy add-cluster-role-to-user cluster-reader system:serviceaccount:logging:aggregated-logging-elasticsearch

oc adm policy add-cluster-role-to-user cluster-reader system:serviceaccount:logging:aggregated-logging-kibana 

oc adm policy add-cluster-role-to-user cluster-admin developer

MINI_TOKEN=$(oc whoami -t)
eval $(minishift docker-env)
docker login -u developer -p ${MINI_TOKEN} 172.30.1.1:5000 
cd $GOPATH/src/yougroupteam/you-gps
docker build -t gotest:latest .
docker tag gotest:latest 172.30.1.1:5000/myproject/gotest:latest
docker push 172.30.1.1:5000/myproject/gotest:latest
oc deploy --latest gotest


curl --cacert $ES_CA --key $ES_CLIENT_KEY --cert $ES_CLIENT_CERT -s -w "%{http_code}" -XGET "https://$ES_HOST:$ES_PORT/_cat/health?v"
curl --cacert $ES_CA --key $ES_CLIENT_KEY --cert $ES_CLIENT_CERT -s -w "%{http_code}" -XGET "https://$ES_HOST:$ES_PORT/_cat/indices?v"

curl --cacert $ES_CA --key $ES_CLIENT_KEY --cert $ES_CLIENT_CERT -s -w "%{http_code}" -XPOST "https://$ES_HOST:$ES_PORT/_msearch"