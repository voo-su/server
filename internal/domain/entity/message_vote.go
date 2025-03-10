package entity

type QueryVoteModel struct {
	MessageId    int    `json:"message_id"`
	ReceiverId   int    `json:"receiver_id"`
	ChatType     int    `json:"chat_type"`
	MsgType      int    `json:"msg_type"`
	VoteId       int    `json:"vote_id"`
	AnswerMode   int    `json:"answer_mode"`
	AnswerOption string `json:"answer_option"`
	AnswerNum    int    `json:"answer_num"`
	VoteStatus   int    `json:"vote_status"`
}
