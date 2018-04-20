// Copyright (C) 2014-2018 Goodrain Co., Ltd.
// RAINBOND, Application Management Platform

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package sources

import (
	"io"
	"testing"

	"github.com/goodrain/rainbond/pkg/event"
)

func init() {
	event.NewManager(event.EventConfig{
		DiscoverAddress: []string{"172.17.0.1:2379"},
	})
}
func TestGitClone(t *testing.T) {
	csi := CodeSourceInfo{
		RepositoryURL: "https://github.com/goodrain/rainbond-docs.git",
		Branch:        "master",
	}
	//logger := event.GetManager().GetLogger("system")
	res, err := GitClone(csi, "/tmp/rainbonddoc2", nil, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", res)
}

func TestGitPull(t *testing.T) {
	csi := CodeSourceInfo{
		RepositoryURL: "git@code.goodrain.com:goodrain/test.git",
		Branch:        "master2",
	}
	//logger := event.GetManager().GetLogger("system")
	res, err := GitPull(csi, "/tmp/master2", nil, 1)
	if err != nil {
		t.Fatal(err)
	}
	commit, err := GetLastCommit(res)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", commit)
}

func TestGitPullOrClone(t *testing.T) {
	csi := CodeSourceInfo{
		RepositoryURL: "git@code.goodrain.com:goodrain/goodrain_web.git",
		Branch:        "publiccloud",
	}
	//logger := event.GetManager().GetLogger("system")
	res, err := GitCloneOrPull(csi, "/tmp/goodrainweb", nil, 1)
	if err != nil {
		t.Fatal(err)
	}
	//get last commit
	commit, err := GetLastCommit(res)
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}
	t.Logf("%+v", commit)
}

func TestGetCodeCacheDir(t *testing.T) {
	csi := CodeSourceInfo{
		RepositoryURL: "git@121.196.222.148:summersoft/yycx_push.git",
		Branch:        "test",
	}
	t.Log(csi.GetCodeSourceDir())
}
