package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"time"
	"os"
	"bufio"
	"errors"

	log "github.com/Sirupsen/logrus"
	"github.com/coreos/go-etcd/etcd"

)

// loadEtcdNode reads `r` containing JSON objects representing etcd nodes and
// loads them into server.
func loadEtcdNode(etcdClient *etcd.Client, r io.Reader) error {
	jsonReader := json.NewDecoder(r)
	for {
		node := etcd.Node{}
		if err := jsonReader.Decode(&node); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if node.Key == "" && node.Dir {
			continue // skip root
		}

		// compute a new TTL
		ttl := 0
		if node.Expiration != nil {
			ttl := node.Expiration.Sub(time.Now()).Seconds()
			if ttl < 0 {
				// expired, skip
				continue
			}
		}

		if node.Dir {
			_, err := etcdClient.SetDir(node.Key, uint64(ttl))
			if err != nil {
				return fmt.Errorf("%s: %s", node.Key, err)
			}
		} else {
			_, err := etcdClient.Set(node.Key, node.Value, uint64(ttl))
			if err != nil {
				return fmt.Errorf("%s: %s", node.Key, err)
			}
		}
	}
	return nil
}


func checkNodeState(key string, etcdClient *etcd.Client) (int, error) {
	response, err := etcdClient.Get(key, false, false)
	if err != nil {
		return 0, err
	}

	childNodes := response.Node.Nodes
	count := 0

	// enumerate all the child nodes.
	for range childNodes {
		count += 1
	}

	return count, nil
}


func restoreBackup(etcdLocalURL string, backupPath string, backupFile string) error {


	etcdClient := etcd.NewClient([]string{fmt.Sprintf("http://%s:2379", etcdLocalURL)})


	var valueCount int
	var err error

	valueCount, err = checkNodeState("registry", etcdClient)

	if valueCount > 0 {
		return errors.New("etcd dir tree already populated")
	}

	file, err := os.Open(fmt.Sprintf("%s%s", backupPath, backupFile))

	if err != nil {
		return err
	}

	gzipReader, err := gzip.NewReader(bufio.NewReader(file))
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	if err := loadEtcdNode(etcdClient, gzipReader); err != nil {
		return err
	}

	log.Printf("restore: complete")
	file.Close()

	return nil
}
