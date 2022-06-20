package repo

import "testing"

// SVN：
//		`https://host/svn/Expo_Cloud/branches/bigdata20220412`,
//		`svn://gitee.com/conero/lang`
//		`svn+ssh://gitee.com/conero/lang`
// Git：
//		`http://host/lff19910329/big-protal.git`
//		`git@gitee.com:conero/lang.git`
func TestStdRepoUrl(t *testing.T) {
	tUrl := "https://host/svn/Expo_Cloud/branches/bigdata20220412"
	rfUrl := "https://host/svn/Expo_Cloud/branches/bigdata20220412"
	rfTy := ""

	if ru, rt := StdRepoUrl(tUrl); ru != rfUrl || rt != rfTy {
		t.Errorf("Address (%v) translation error, unexpected structure.", tUrl)
	}

}
