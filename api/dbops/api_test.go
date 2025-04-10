package dbops

import "testing"

var (
	tempvid string
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}
func testAddUser(t *testing.T) {
	err := AddUserCredential("test", "123")
	if err != nil {
		t.Errorf("error of AddUser: %v\n", err)
	}
}
func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("test")
	if pwd != "123" || err != nil {
		t.Errorf("error of GetUser")
	}
}
func testDeleteUser(t *testing.T) {
	err := DeleteUser("test", "123")
	if err != nil {
		t.Errorf("error of DeleteUser: %v\n", err)
	}
}
func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("test")
	if err != nil {
		t.Errorf("error of RegetUser: %v\n", err)
	}
	if pwd != "" {
		t.Errorf("deleteing user test failed")
	}
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Delete", testDeleteUser)
	t.Run("Reget", testRegetUser)
}
func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "video 1")

	if err != nil {
		t.Errorf("error of AddNewVideo: %v\n", err)
	}
	tempvid = vi.Id
}
func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("error of GetVideoInfo: %v\n", err)
	}
}
func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("error of DeleteVideo: %v\n", err)
	}
}
func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil {
		t.Errorf("error of RegetVideo: %v\n", err)
	}
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DeleteVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}
