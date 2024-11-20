package artifactutil

import (
	"testing"
)

// TestArtifactUtilIsInSliceWithInt
//
//	@Description:
//	@param t
func TestArtifactUtilIsInSliceWithInt(t *testing.T) {
	dataIntSlice := []int{1, 2, 35, 6, 7, 8}
	dataInt := 8

	resWithInt := IsInSlice(dataInt, dataIntSlice)
	t.Logf("int：%t", resWithInt)

	dataInt8Slice := []int8{1, 2, 35, 6, 7, 8}
	var dataInt8 int8 = 8
	resWithInt8 := IsInSlice(dataInt8, dataInt8Slice)
	t.Logf("int8：%t", resWithInt8)

	dataInt16Slice := []int16{1, 2, 35, 6, 7, 8}
	var dataInt16 int16 = 8
	resWithInt16 := IsInSlice(dataInt16, dataInt16Slice)
	t.Logf("int16：%t", resWithInt16)

	dataInt32Slice := []int32{1, 2, 35, 6, 7, 8}
	var dataInt32 int32 = 8
	resWithInt32 := IsInSlice(dataInt32, dataInt32Slice)
	t.Logf("int32：%t", resWithInt32)

	dataInt64Slice := []int64{1, 2, 35, 6, 7, 8}
	var dataInt64 int64 = 8
	resWithInt64 := IsInSlice(dataInt64, dataInt64Slice)
	t.Logf("int64：%t", resWithInt64)
}

func TestArtifactUtilReDup(t *testing.T) {
	dataUintSlice := []uint{1, 8, 2, 35, 6, 7, 8, 8}

	resWithUint := ReDup(dataUintSlice)
	t.Logf("uint：%v", resWithUint)

	dataUint8Slice := []uint8{1, 2, 35, 6, 7, 8}

	resWithUint8 := ReDup(dataUint8Slice)
	t.Logf("uint8：%v", resWithUint8)

	dataUint16Slice := []uint16{1, 2, 35, 6, 7, 8}

	resWithUint16 := ReDup(dataUint16Slice)
	t.Logf("uint16：%v", resWithUint16)

	dataUint32Slice := []uint32{1, 2, 35, 6, 7, 8}

	resWithUint32 := ReDup(dataUint32Slice)
	t.Logf("uint32：%v", resWithUint32)

	dataUint64Slice := []uint64{1, 2, 35, 6, 7, 8}
	resWithUint64 := ReDup(dataUint64Slice)
	t.Logf("uint64：%v", resWithUint64)
}

func TestArtifactGetMapValue(t *testing.T) {
	dataMap := map[string]any{
		"2": 1,
		"3": 2,
		"4": 1,
	}

	t.Logf("map：%v", GetMapValue(dataMap))
	t.Logf("map：%v", GetMapValueDup(dataMap))
}
