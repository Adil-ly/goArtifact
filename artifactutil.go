package goArtifact

type ArtifactUtil struct {
}

func NewArtifactUtil() *ArtifactUtil {
	return &ArtifactUtil{}
}

// IsInSliceWithInt
//
//	@Description: 判断当前Int的是否存在于切片中
//	@receiver a
//	@param target
//	@param targetSlice
//	@return bool
func (a *ArtifactUtil) IsInSliceWithInt(target int, targetSlice []int) bool {
	for _, item := range targetSlice {
		if target == item {
			return true
		}
	}
	return false
}

// IsInSliceWithInt8
//
//	@Description: 判断当前Int8的是否存在于切片中
//	@receiver a
//	@param target
//	@param targetSlice
//	@return bool
func (a *ArtifactUtil) IsInSliceWithInt8(target int8, targetSlice []int8) bool {
	for _, item := range targetSlice {
		if target == item {
			return true
		}
	}
	return false
}

// IsInSliceWithInt32
//
//	@Description: 判断当前Int32的是否存在于切片中
//	@receiver a
//	@param target
//	@param targetSlice
//	@return bool
func (a *ArtifactUtil) IsInSliceWithInt32(target int32, targetSlice []int32) bool {
	for _, item := range targetSlice {
		if target == item {
			return true
		}
	}
	return false
}

// IsInSliceWithInt64
//
//	@Description: 判断当前Int64的是否存在于切片中
//	@receiver a
//	@param target
//	@param targetSlice
//	@return bool
func (a *ArtifactUtil) IsInSliceWithInt64(target int64, targetSlice []int64) bool {
	for _, item := range targetSlice {
		if target == item {
			return true
		}
	}
	return false
}

// IsInSliceWithString
//
//	@Description: 判断当前string的是否存在于切片中
//	@receiver a
//	@param target
//	@param targetSlice
//	@return bool
func (a *ArtifactUtil) IsInSliceWithString(target string, targetSlice []string) bool {
	for _, item := range targetSlice {
		if target == item {
			return true
		}
	}
	return false
}

// ReDupWithString
//
//	@Description: 字符串切片的去重
//	@receiver a
//	@param params
//	@return []string
func (a *ArtifactUtil) ReDupWithString(params []string) []string {
	res := make([]string, 0, len(params))
	tmpMap := map[string]struct{}{}
	for _, item := range params {
		if _, exists := tmpMap[item]; !exists {
			tmpMap[item] = struct{}{}
			res = append(res, item)
		}
	}
	return res
}

// ReDupWithInt
//
//	@Description:Int切片的去重
//	@receiver a
//	@param params
//	@return []int
func (a *ArtifactUtil) ReDupWithInt(params []int) []int {
	res := make([]int, 0, len(params))
	tmpMap := map[int]struct{}{}
	for _, item := range params {
		if _, exists := tmpMap[item]; !exists {
			tmpMap[item] = struct{}{}
			res = append(res, item)
		}
	}
	return res
}

// ReDupWithInt8
//
//	@Description: Int8切片的去重
//	@receiver a
//	@param params
//	@return []int8
func (a *ArtifactUtil) ReDupWithInt8(params []int8) []int8 {
	res := make([]int8, 0, len(params))
	tmpMap := map[int8]struct{}{}
	for _, item := range params {
		if _, exists := tmpMap[item]; !exists {
			tmpMap[item] = struct{}{}
			res = append(res, item)
		}
	}
	return res
}

// ReDupWithInt32
//
//	@Description: Int32切片的去重
//	@receiver a
//	@param params
//	@return []int32
func (a *ArtifactUtil) ReDupWithInt32(params []int32) []int32 {
	res := make([]int32, 0, len(params))
	tmpMap := map[int32]struct{}{}
	for _, item := range params {
		if _, exists := tmpMap[item]; !exists {
			tmpMap[item] = struct{}{}
			res = append(res, item)
		}
	}
	return res
}

// ReDupWithInt64
//
//	@Description: Int64切片的去重
//	@receiver a
//	@param params
//	@return []int64
func (a *ArtifactUtil) ReDupWithInt64(params []int64) []int64 {
	res := make([]int64, 0, len(params))
	tmpMap := map[int64]struct{}{}
	for _, item := range params {
		if _, exists := tmpMap[item]; !exists {
			tmpMap[item] = struct{}{}
			res = append(res, item)
		}
	}
	return res
}
