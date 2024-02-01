package main

import "fmt"

func main() {
	chainId := "534352"
	msgSentTable := fmt.Sprintf("msg_sent_%s", chainId)
	msgHashTable := fmt.Sprintf("msg_hash_%s", chainId)

	updateSql := `
				update %s a  set msg_hash=b.msg_hash, msg_hash_relation=true 
				from %s  b 
				where  a.tx_hash=b.tx_hash  and a.msg_hash_relation=false
				`
	updateSql = fmt.Sprintf(updateSql, msgSentTable, msgHashTable)

	fmt.Println(updateSql)
}
