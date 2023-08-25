package test

import (
	"github.com/Adil-ly/goArtifact.git"
	"testing"
)

// TestArtifactUtilIsInSliceWithInt
//
//	@Description: int测试
//	@param t
func TestArtifactUtilIsInSliceWithInt(t *testing.T) {
	dataSlice := []int{1, 2, 35, 6, 7, 8}
	dataInt := 7
	res := goArtifact.NewArtifactUtil().IsInSliceWithInt(dataInt, dataSlice)
	t.Logf("结果：%t", res)
}