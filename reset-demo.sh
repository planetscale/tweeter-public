#!/bin/bash

# Reset pscaledb
## Skip pending reverts
drnumout=$(pscale deploy-request list tweeter --org bmorrison-ps -f json)
drnum=$(jq -n "$drnumout" | jq -r '.[0].number')
pscale dr skip-revert tweeter $drnum

## Drop likes col
echo "alter table tweets drop likes;" | pscale shell tweeter add-likes
pscale dr create tweeter add-likes
drnumout=$(pscale deploy-request list tweeter --org bmorrison-ps -f json)
drnum=$(jq -n "$drnumout" | jq -r '.[0].number')
pscale dr deploy $drnum
pscale dr skip-revert tweeter $drnum

## Remove existing branch
pscale branch rm tweeter add-likes --force
echo "delete from schema_migrations where version = '20230832120038';" | pscale shell tweeter main

# Reset git
git checkout clean
git branch -D main
git push origin -d main
git checkout -b main
git push --set-upstream origin main

# Reset fly app
fly deploy
