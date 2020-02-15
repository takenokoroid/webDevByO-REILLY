package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffertracer :=New(&buf)
	if travcer==nil{
		t.Error("Newからの戻り値がnilです")
	}else{
		tracer.Trace("こんにちは。traceパッケージ")
		if buf.String()!="こんにちは。traceパッケージ\n"{
			t.Errorf("'%s'とい誤った文字列が出力されました",buf.String())
		}
	}
}
