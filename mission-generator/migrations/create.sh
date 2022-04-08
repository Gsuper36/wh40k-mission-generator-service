#bin/sh

timestamp=$(date +%s)
migration_name="${timestamp}_${1}"
touch "${migration_name}.up.sql"
touch "${migration_name}.down.sql"