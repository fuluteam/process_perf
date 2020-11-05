package perf

import "testing"

/**
 * @author [kevinyang]
 * @email [yangchujie6@mail.com]
 * @create date 2020-11-04 17:04
 * @modify date 2020-11-04 17:04
 * @desc [单元测试文件]
 */

func TestStartCollect(t *testing.T) {
	type args struct {
		pid int32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "测试不存在的进程ID",
			args:    args{pid: 999999999},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StartCollect(tt.args.pid); (err != nil) != tt.wantErr {
				t.Errorf("StartCollect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
