Go Server: Record Logs
===================================================

## Deployment

* Production
    * The IP to be used is something like 192.x.x.x ( e.g. 192.168.1.100)
    * Add in `/etc/hosts` the entry `192.x.x.x	ses.pt`
    * Launching:
       ```shell
       docker-compose up --build
       ```

### Purpose

* This go server was built to receive logs from a modified IPFS BitSwap version in order to evaluate the changes
* The database where the log entries are saved records the blockID, RequestDelay, BlockDelay and whether the block was sucessfully delivered or not
* To access the database run:
    ```shell
    docker exec -ti go-test_server_db_1 /bin/bash -c 'exec mariadb -u root -p"$MARIADB_ROOT_PASSWORD"'
    ```

