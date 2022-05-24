/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/4/18 1:53
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/4/18 1:53
 */

package temp_func

import "testing"

func TestToPath(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			"test0",
			"TestUser",
			"test-user",
		},
		{
			"test1",
			"test",
			"test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToPath(tt.args); got != tt.want {
				t.Errorf("ToPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
