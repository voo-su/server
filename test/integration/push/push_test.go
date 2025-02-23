package push

import (
	"testing"
	"voo.su/pkg/push"
)

func TestFirebasePush(t *testing.T) {
	firebase := push.NewFIREBASEPush("", "", "", "")

	payloadInfo := &push.PayloadInfo{
		Title:   "title",
		Content: "content",
		Badge:   1,
	}

	_, err := firebase.Push("", push.NewFIREBASEPayload(payloadInfo, ""))
	if err != nil {
		t.Error(err)
	}
}
