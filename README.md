# Checktool for etcd backups

Based on [the etcd backup tool](https://github.com/crewjam/etcd-aws) for Etcd running in AWS.
This tool backups the contents of an etcd cluster running in AWS to S3. This tool allows to download the resulting file of the backup and restore it to an arbitrary etcd instance.

Purpose is to create an Jenkins job which can check the etcd backup for validity and usability to restore a vanilla etcd cluster.

## Installation

### Local
```
go get github.com/fourscouts/checktool
```

### Docker
todo

## Example usage

Launch a local ETCD docker image
```
docker run --name etcd \
-p 2379:2379 \
quay.io/coreos/etcd:v3.0.15 /usr/local/bin/etcd \
-advertise-client-urls 'http://0.0.0.0:2379,http://0.0.0.0:4001' \
-listen-client-urls 'http://0.0.0.0:2379,http://0.0.0.0:4001'
```

You can now restore an etcd backup by running:
````
checktool -backup-path=/tmp/etcd-backup.gz -etcd-url=127.0.0.1

````


## commandline options
```
Usage of ./checktool:
  -backup-path string
    	The path to the backup file. Environment variable: ETCD_BACKUP_PATH (default "/tmp/etcd-backup.gz")
  -etcd-url string
    	URL of the etcd node. Environment variable: ETCD_URL (default "127.0.0.1")

```
