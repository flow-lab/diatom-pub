# diatom

Diatom test

```shell
# build and run in docker
docker compose --profile dev up --force-recreate --build -d
# curl http://localhost/backend/api/api.yaml
# curl http://localhost/backend/api/authors/a5dcd6de-4012-408c-a079-9945b3ad7c5a -s -w "\nHTTP_RESPONSE_STATUS_CODE: %{http_code}\n"

# show logs
docker compose --profile dev logs -f
``````