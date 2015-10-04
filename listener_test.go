package ircnotifier_test

import (
	"testing"
	"bytes"
)

func Test(t *testing.T) {
	bla := []byte(`{"object_kind":"push","before":"95790bf891e76fee5e1747ab589903a6a1f80f22","after":"da1560886d4f094c3e6c9ef40349f7d38b5d27d7","ref":"refs/heads/master","user_id":4,"user_name":"John Smith","user_email":"john@example.com","project_id":15,"repository":{"name":"Diaspora","url":"git@example.com:mike/diasporadiaspora.git","description":"","homepage":"http://example.com/mike/diaspora","git_http_url":"http://example.com/mike/diaspora.git","git_ssh_url":"git@example.com:mike/diaspora.git","visibility_level":0},"commits":[{"id":"b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327","message":"Update Catalan translation to e38cb41.","timestamp":"2011-12-12T14:27:31+02:00","url":"http://example.com/mike/diaspora/commit/b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327","author":{"name":"Jordi Mallach","email":"jordi@softcatala.org"}},{"id":"da1560886d4f094c3e6c9ef40349f7d38b5d27d7","message":"fixed readme","timestamp":"2012-01-03T23:36:29+02:00","url":"http://example.com/mike/diaspora/commit/da1560886d4f094c3e6c9ef40349f7d38b5d27d7","author":{"name":"GitLab dev user","email":"gitlabdev@dv6700.(none)"}}],"total_commits_count":4}`)
	var akk ircnotifier.PushEvent
	json.Unmarshal(bla, &akk)
	if akk.UserName != "John Smith" {
		t.Error("UserName does not match")
func TestDecode(t *testing.T) {
	json := []byte(`{"object_kind":"push","before":"95790bf891e76fee5e1747ab589903a6a1f80f22","after":"da1560886d4f094c3e6c9ef40349f7d38b5d27d7","ref":"refs/heads/master","user_id":4,"user_name":"John Smith","user_email":"john@example.com","project_id":15,"repository":{"name":"Diaspora","url":"git@example.com:mike/diasporadiaspora.git","description":"","homepage":"http://example.com/mike/diaspora","git_http_url":"http://example.com/mike/diaspora.git","git_ssh_url":"git@example.com:mike/diaspora.git","visibility_level":0},"commits":[{"id":"b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327","message":"Update Catalan translation to e38cb41.","timestamp":"2011-12-12T14:27:31+02:00","url":"http://example.com/mike/diaspora/commit/b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327","author":{"name":"Jordi Mallach","email":"jordi@softcatala.org"}},{"id":"da1560886d4f094c3e6c9ef40349f7d38b5d27d7","message":"fixed readme","timestamp":"2012-01-03T23:36:29+02:00","url":"http://example.com/mike/diaspora/commit/da1560886d4f094c3e6c9ef40349f7d38b5d27d7","author":{"name":"GitLab dev user","email":"gitlabdev@dv6700.(none)"}}],"total_commits_count":4}`)
	reader := bytes.NewReader(json);
	testee := &NotificationListener{}
	pushEvent, err := testee.parseRequestBody(reader)
	if err != nil {
		t.Errorf("Error occured %s", err.Error())
	}
	if pushEvent.UserName != "John Smith" {
		t.Errorf("Expected: %s Given: %s", "John Smith", pushEvent.UserName)
	}
	if pushEvent.Repository.Name != "Diaspora" {
		t.Errorf("Expected: %s Given: %s", "Diaspora", pushEvent.Repository.Name)
	}
	if pushEvent.Commits[1].Id != "da1560886d4f094c3e6c9ef40349f7d38b5d27d7" {
		t.Errorf("Expected: %s Given: %s", "da1560886d4f094c3e6c9ef40349f7d38b5d27d7", pushEvent.Commits[1].Id)
	}

}