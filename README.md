# Checktool for etcd backups

Based on [the etcd backup tool](https://github.com/crewjam/etcd-aws) for Etcd running in AWS.
This tool backups the contents of an etcd cluster running in AWS to S3. This tool allows to download the resulting file of the backup and restore it to an arbitrary etcd instance.

Purpose is to create an Jenkins job which can check the etcd backup for validity and usability to restore a vanilla etcd cluster.


## commandline options
```
  -backup-file string
        The name of the backup file. Environment variable: ETCD_BACKUP_KEY (default "etcd-backup.gz")
  -etcd-url string
        URL of the etcd node. Environment variable: ETCD_URL (default "127.0.0.1")
```
