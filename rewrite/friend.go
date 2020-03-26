package main

import "database/sql"

const (
	FriendTable = "friends"

	FriendRequester = "requester"
	FriendSetup     = "setup"
	FriendFriend    = "friend"
	FriendSlot      = "slot"
)

var FriendMatchingQuery = SQLQuery([]string{
	FriendFriend,
	FriendSlot,
}, FriendTable, FriendRequester)

func FriendMatchings(tx *sql.Tx, player string) [PlayerFriendGameCount]string {
	rows, err := tx.Query(FriendMatchingQuery, player)
	if err != nil {
		Panic(err)
	}
	var friends [PlayerFriendGameCount]string
	for rows.Next() {
		var friend string
		var slot uint8
		err = rows.Scan(&friend, &slot)
		if err != nil {
			Panic(err)
		}
		friends[slot] = friend
	}
	err = rows.Err()
	if err != nil {
		Panic(err)
	}
	return friends
}
