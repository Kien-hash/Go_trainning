package userdata

import (
	"math/rand"
	types "resful-server/datastruct"
	"testing"
	"time"
)

func TestSelectPasswords(t *testing.T) {
	// Testcase 1: Username was registed on database
	testcase1 := SelectPasswords("kien")
	if testcase1.Passwords == "" {
		t.Errorf("Output expect is: %v", testcase1)
	}

	// Testcase 2: Username was not registed on database
	userExpected := types.User{ID: 0, Username: "", Passwords: "", Firstname: "", Lastname: ""}
	testcase2 := SelectPasswords("\"zxcdasf\"zxcvasdf")
	if testcase2 != userExpected {
		t.Errorf("Output is unexpected: %v", testcase2)
	}

	// Testcase 2: Username is not null but empty
	testcase3 := SelectPasswords("")
	if testcase2 != userExpected {
		t.Errorf("Output is unexpected: %v", testcase3)
	}
}

func TestInsertUser(t *testing.T) {
	// Testcase 1: username wasn't registed on database
	username := randomString(10)
	user := types.User{ID: 0, Username: username, Passwords: "asdfvczx", Firstname: "asdfcvz", Lastname: "adsfcxvad"}
	testcase1 := InsertUser(user)
	if testcase1 != "Register successfully" {
		t.Errorf("Output is unexpected: %s", testcase1)
	}

	// Testcase 2: Username is not null but empty
	user.Username = ""
	testcase2 := InsertUser(user)
	if testcase2 != "Username is empty!" {
		t.Errorf("Output is unexpected: %s", testcase2)
	}

	// Testcase 3: Username was registed on database
	user.Username = "kien"
	testcase3 := InsertUser(user)
	if testcase3 == "Username is empty!" || testcase3 == "Register successfully" {
		t.Errorf("Output is unexpected: %s", testcase3)
	}
}

func randomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

// PASS
// coverage: 90.3% of statements
// ok      _/home/kien/Workspace/Git/kien-training/Chat_implement_/resful-server/userdata  0.069s
// Không thể cover 100% do không bắt được lỗi do truy vấn
