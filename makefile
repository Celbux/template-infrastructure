start-data-api:
	go run ./services/dataapo/main.go --web-api-host=0.0.0.0:8082 --web-read-timeout=5s --web-write-timeout=0s --web-shutdown-timeout=5s --datastore-project-id=dev8celbux --datastore-setting=LOCAL_WITH_CLOUD_DB --datastore-credentials-file-path=./key.json

start-user-api:
	go run ./services/user/main.go --web-api-host=0.0.0.0:8080 --web-read-timeout=5s --web-write-timeout=0s --web-shutdown-timeout=5s --datastore-project-id=dev8celbux --datastore-setting=LOCAL_WITH_CLOUD_DB --datastore-credentials-file-path=./key.json
