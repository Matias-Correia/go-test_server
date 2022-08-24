Go Server: Record Logs
===================================================

## Deployment

* Production
    * Launching:
       ```shell
       docker-compose up --build
       ```

## Purpose

This go server was built to receive logs from a modified IPFS BitSwap version in order to evaluate the changes
Each database entrie has:
- `LogID` int(11)
- `BlockId` varchar(128)
- `LocalPeer` varchar(128)
- `RemotePeer` varchar(128)
- `SentAt` datetime(3)
- `ReceivedAt` datetime(3)
- `BlockRequestedAt` datetime(3)
- `Duplicate` tinyint(1)

The `LogID` being the primary key identifying the logs, `BlockId` idenfies the transmited block id, `LocalPeer` and `RemotePeer` identify the ID of the sender and receiver respectively, `SentAt`, `ReceivedAt` and `BlockRequestedAt` can be null and each records the timestamp of the action performed. For example if we are logging a Block Request than `SentAt` and `ReceivedAt` should be null and `BlockRequestedAt` has the timestamp from when the block was requested. Lastly `Duplicate` tells us if the action is a duplicate or not.

To access the database run:
    ```shell
    docker exec -ti go-test_server_db_1 /bin/bash -c 'exec mariadb -u root -p"$MARIADB_ROOT_PASSWORD"'
    ```

