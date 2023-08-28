package test

import (
	"github.com/Adil-ly/goArtifact"
	"testing"
)

// TestArtifactUtilIsInSliceWithInt
//
//	@Description: int测试
//	@param t
func TestArtifactUtilIsInSliceWithInt(t *testing.T) {
	dataIntSlice := []int{1, 2, 35, 6, 7, 8}
	dataInt := 8
	util := goArtifact.NewArtifactUtil()
	resWithInt := util.IsInSliceWithInt(dataInt, dataIntSlice)
	t.Logf("int结果：%t", resWithInt)

	dataInt8Slice := []int8{1, 2, 35, 6, 7, 8}
	var dataInt8 int8 = 8
	resWithInt8 := util.IsInSliceWithInt8(dataInt8, dataInt8Slice)
	t.Logf("int8结果：%t", resWithInt8)

	dataInt16Slice := []int16{1, 2, 35, 6, 7, 8}
	var dataInt16 int16 = 8
	resWithInt16 := util.IsInSliceWithInt16(dataInt16, dataInt16Slice)
	t.Logf("int16结果：%t", resWithInt16)

	dataInt32Slice := []int32{1, 2, 35, 6, 7, 8}
	var dataInt32 int32 = 8
	resWithInt32 := util.IsInSliceWithInt32(dataInt32, dataInt32Slice)
	t.Logf("int32结果：%t", resWithInt32)

	dataInt64Slice := []int64{1, 2, 35, 6, 7, 8}
	var dataInt64 int64 = 8
	resWithInt64 := util.IsInSliceWithInt64(dataInt64, dataInt64Slice)
	t.Logf("int64结果：%t", resWithInt64)
}

func TestArtifactUtilIsInSliceWithUint(t *testing.T) {
	dataUintSlice := []uint{1, 2, 35, 6, 7, 8}
	var dataUint uint = 8
	util := goArtifact.NewArtifactUtil()
	resWithUint := util.IsInSliceWithUint(dataUint, dataUintSlice)
	t.Logf("uint结果：%t", resWithUint)

	dataUint8Slice := []uint8{1, 2, 35, 6, 7, 8}
	var dataUint8 uint8 = 8
	resWithUint8 := util.IsInSliceWithUint8(dataUint8, dataUint8Slice)
	t.Logf("uint8结果：%t", resWithUint8)

	dataUint16Slice := []uint16{1, 2, 35, 6, 7, 8}
	var dataUint16 uint16 = 8
	resWithUint16 := util.IsInSliceWithUint16(dataUint16, dataUint16Slice)
	t.Logf("uint16结果：%t", resWithUint16)

	dataUint32Slice := []uint32{1, 2, 35, 6, 7, 8}
	var dataUint32 uint32 = 8
	resWithUint32 := util.IsInSliceWithUint32(dataUint32, dataUint32Slice)
	t.Logf("uint32结果：%t", resWithUint32)

	dataUint64Slice := []uint64{1, 2, 35, 6, 7, 8}
	var dataUint64 uint64 = 8
	resWithUint64 := util.IsInSliceWithUint64(dataUint64, dataUint64Slice)
	t.Logf("uint64结果：%t", resWithUint64)
}
