package main

import (
	"flag"
	"os"
	log "github.com/Sirupsen/logrus"
)

var etcdLocalURL *string
var backupPath *string

func main() {
	
	if err := restoreBackup(*etcdLocalURL, *backupPath); err != nil {
		log.Fatalf("ERROR: %s", err)
	}

}

func init () {

	var defaultEtcdLocalURL string = "127.0.0.1"
	if u := os.Getenv("ETCD_URL"); u != "" {
		defaultEtcdLocalURL = u
	}

	etcdLocalURL = flag.String("etcd-url", defaultEtcdLocalURL,
		"URL of the etcd node. " +
			"Environment variable: ETCD_URL")

	var defaultBackupPath string = "/tmp/etcd-backup.gz"
	if p := os.Getenv("ETCD_BACKUP_PATH"); p != "" {
		defaultBackupPath = p
	}

	backupPath = flag.String("backup-path", defaultBackupPath,
		"The path to the backup file. " +
			"Environment variable: ETCD_BACKUP_PATH")

	flag.Parse()
}