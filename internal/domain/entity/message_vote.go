// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package entity

type QueryVoteModel struct {
	RecordId     int    `json:"record_id"`
	ReceiverId   int    `json:"receiver_id"`
	DialogType   int    `json:"dialog_type"`
	MsgType      int    `json:"msg_type"`
	VoteId       int    `json:"vote_id"`
	AnswerMode   int    `json:"answer_mode"`
	AnswerOption string `json:"answer_option"`
	AnswerNum    int    `json:"answer_num"`
	VoteStatus   int    `json:"vote_status"`
}
