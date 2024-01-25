package domain

import "fmt"

type Cell struct {
	ID      string
	Name    string
	Address string
}

func CreateCellID(tenantID, shardID string) string {
	return fmt.Sprintf("TENANT#%s#SHARD%s", tenantID, shardID)
}

/* -- cell-mapping
id							name		address
TENANT#customer#SHARD#1		CELL#1		http://customer.shard1.httpconnector.cell1:8080
TENANT#customer#SHARD#1		CELL#2		http://customer.shard1.httpconnector.cell2:8080
TENANT#customer#SHARD#1		CELL#3		http://customer.shard1.httpconnector.cell3:8080

TENANT#customer#SHARD#2		CELL#1		http://http-connector.cell-1:8080
TENANT#customer#SHARD#2		CELL#2		http://http-connector.cell-2:8080
TENANT#customer#SHARD#2		CELL#3		http://http-connector.cell-3:8080

TENANT#partner#SHARD#1		CELL#1		http://http-connector.cell-2:8080
TENANT#partner#SHARD#1		CELL#2		http://http-connector.cell-2:8080
*/
