package logic

import "testing"

func TestGetAccessToken(t *testing.T) {
	accessToken := GetDingAccessToken()
	t.Log(accessToken)
}
