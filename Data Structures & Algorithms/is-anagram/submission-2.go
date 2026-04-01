import(
	"maps"
)

func isAnagram(s string, t string) bool {
	sValueMap := map[byte]int{}
	tValueMap := map[byte]int{}

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		sValueMap[s[i]]++
		tValueMap[t[i]]++
	}


	return maps.Equal(sValueMap, tValueMap)
}
