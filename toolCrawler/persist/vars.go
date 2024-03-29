package persist

import (
	"fmt"
	"github.com/gearboxworks/scribeHelpers/toolCrawler/global"
)

var tables = []global.Tablename{
	"hosts",
	"resources",
	"queue",
	"visited",
}
var ddl = []global.Sql{
	"CREATE TABLE IF NOT EXISTS hosts (id INTEGER PRIMARY KEY, scheme TEXT, domain TEXT, port INTEGER)",
	"CREATE UNIQUE INDEX IF NOT EXISTS idx_hosts ON hosts (scheme,domain,port)",

	"CREATE TABLE IF NOT EXISTS resources (id INTEGER PRIMARY KEY, hash INTEGER, host_id INTEGER, urlpath TEXT)",
	"CREATE UNIQUE INDEX IF NOT EXISTS idx_resources_hash ON resources (hash)",
	"CREATE UNIQUE INDEX IF NOT EXISTS idx_resources_url ON resources (host_id,urlpath)",

	"CREATE TABLE IF NOT EXISTS queue (id INTEGER PRIMARY KEY, resource_hash INTEGER, timestamp INTEGER)",
	"CREATE UNIQUE INDEX IF NOT EXISTS idx_queue_resource_hash ON queue (resource_hash,timestamp)",

	"CREATE TABLE IF NOT EXISTS visited (id INTEGER PRIMARY KEY, resource_hash INTEGER, timestamp INTEGER, headers TEXT, body BLOB, cookies BLOB)",
	"CREATE UNIQUE INDEX IF NOT EXISTS idx_visited ON visited (resource_hash,timestamp)",
}

const (
	SelectQueueCountDml      = "select-queue-count"
	SelectQueueItemDml       = "select-queue-item"
	SelectQueueItemByHashDml = "select-queue-item-by-hash"

	SelectResourceCountDml     = "select-resource-count"
	SelectResourceCountByIdDml = "select-resource-count-by-id"
	SelectResourceDml          = "select-resource"
	SelectResourceByIdDml      = "select-resource-by-id"
	SelectResourceByHashDml    = "select-resource-by-hash"
	SelectHostByUrlDml         = "select-host-by-url"
	SelectHostByIdDml          = "select-host-by-id"
	SelectHostBySDPDml         = "select-host-by-sdp"

	SelectVisitedStatsByHashDml = "select-visited-stats-by-hash"

	InsertHostDml      = "insert-host"
	InsertResourceDml  = "insert-resource"
	InsertVisitedDml   = "insert-visited"
	InsertQueueItemDml = "insert-queue-item"

	DeleteQueueItemDml        = "delete-queue-item"
	DeleteQueueItemsByHashDml = "delete-queue-items-by-hash"
)

type sqlMap map[global.Name]global.Sql

var dml = sqlMap{
	SelectResourceDml:      "SELECT id,hash,host_id,urlpath FROM resources",
	SelectResourceCountDml: "SELECT COUNT() FROM resources",
	SelectQueueItemDml:     "SELECT IFNULL(id,0) AS id,IFNULL(resource_hash,0) AS resource_hash FROM queue",
}

func init() {
	dml = mergesqlmap(dml, sqlMap{
		SelectQueueItemByHashDml: dml[SelectQueueItemDml] + " WHERE resource_hash = ?",
		SelectQueueCountDml:      " SELECT COUNT(*) FROM queue",

		SelectResourceByIdDml:      dml[SelectResourceDml] + " WHERE id = ?",
		SelectResourceByHashDml:    dml[SelectResourceDml] + " WHERE hash = ?",
		SelectResourceCountByIdDml: dml[SelectResourceDml] + " WHERE id = ?",

		SelectHostByIdDml:  "SELECT id,scheme,domain,port FROM hosts WHERE id = ?",
		SelectHostByUrlDml: "SELECT id FROM hosts WHERE scheme || '://' || domain || CASE WHEN port IN (0,80) THEN '' ELSE ':'||CAST(port AS text) END || '/' LIKE ?",
		SelectHostBySDPDml: "SELECT id FROM hosts WHERE scheme=? AND domain=? AND port=?",

		SelectVisitedStatsByHashDml: "SELECT COUNT(*),CASE WHEN count(*)=0 THEN 0 ELSE MAX(timestamp) END AS timestamp FROM visited WHERE resource_hash = ?",

		InsertHostDml:      "INSERT INTO hosts (scheme,domain,port) VALUES (?,?,?)",
		InsertResourceDml:  "INSERT INTO resources (hash,host_id,urlpath) VALUES (?,?,?)",
		InsertVisitedDml:   "INSERT INTO visited (resource_hash,timestamp,headers,body,cookies) VALUES(?,?,?,?,?)",
		InsertQueueItemDml: "INSERT INTO queue (resource_hash,timestamp) VALUES (?,?)",

		DeleteQueueItemDml:        "DELETE FROM queue WHERE id = ?",
		DeleteQueueItemsByHashDml: "DELETE FROM queue WHERE resource_hash = ?",
	})
}

func mergesqlmap(m1, m2 sqlMap) sqlMap {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}

var ErrNonIndexableUrl = fmt.Errorf("URL cannot be indexed")
