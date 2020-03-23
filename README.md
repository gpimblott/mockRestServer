# MockRestServer
This is a simple REST server that responds with mock responses based on the requested URI.

The server returns the contents of a file located in the directory specified by the environment variable DATA_DIR.  
For example

* http://localhost:8090/test will return the response in {DATA_DIR}/test.json
* http://localhost:8090/test/data will return the response in {DATA_DIR}/test/data.json

The following environment variables need to be configured

* PORT - port that the server will run on ( default 8090 )
* DATA_DIR - directory containing the root of the responses ( default ./www/ )

## Docker

The included Docker and docker-compose files allow server to be run easily

docker build . -t mockserver
docker container run --publish 8090:8090 --detach --name mockserver mockserver:latest