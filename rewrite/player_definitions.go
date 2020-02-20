package main

const NewPlayerPieceCount = 3

type PlayerRecord struct {
	Wins   int
	Losses int
	Draws  int
}

type PlayerFriendStatus struct {
	ActiveFriends   [player_friend_game_count]string
	MatchingFriends [player_friend_game_count]string
	ActiveTurn      [player_friend_game_count]bool
}
