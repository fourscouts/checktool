package main

import (
	"flag"
	"os"
	log "github.com/Sirupsen/logrus"
)

func main() {

	defaultEtcdLocalURL := "127.0.0.1"
	if u := os.Getenv("ETCD_URL"); u != "" {
		defaultEtcdLocalURL = u
	}

	EtcdLocalURL := flag.String("etcd-url", defaultEtcdLocalURL,
		"URL of the etcd node. "+
			"Environment variable: ETCD_URL")

	defaultBackupFile := "etcd-backup.gz"
	if d := os.Getenv("ETCD_BACKUP_FILE"); d != "" {
		defaultBackupFile = d
	}
	backupFile := flag.String("backup-file", defaultBackupFile,
		"The name of the backup file. "+
			"Environment variable: ETCD_BACKUP_KEY")

	flag.Parse()

	if err := restoreBackup(*EtcdLocalURL, *backupFile); err != nil {
		log.Fatalf("ERROR: %s", err)
	}

}