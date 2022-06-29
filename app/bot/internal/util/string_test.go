package util

import "testing"

func TestUtf8StringSection(t *testing.T) {
	before, after := 5, 5
	for _, tt := range []struct {
		start, end int
		in, out    string
	}{
		{0, 6, "高数一点没学，期末三天成功速通😋", "|高数|一点没学，"},
		{36, 42, "已经毕业了，也用不上这个东东。", "用不上这个|东东|。"},
		{27, 33, "高数一点没学，期末三天成功速通😋", "没学，期末|三天|成功速通😋"},
		{12, 17, "This is the first document we’ve added", " the |first| docu"},
		{6, 12, "咱俩交换一下吧。", "咱俩|交换|一下吧。"},
		{65, 71, "工信处女干事每月经过下属科室都要亲口交代24口交换机等技术性器件的安装工作", "交代24口|交换|机等技术性"},
	} {
		if out := Highlight(tt.in, tt.start, tt.end, before, after, "|"); out != tt.out {
			t.Errorf("Highlight(%s, %d, %d, %d, %d) = %s, want %s", tt.in, tt.start, tt.end, before, after, out, tt.out)
		}
	}
}
