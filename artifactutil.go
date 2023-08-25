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
