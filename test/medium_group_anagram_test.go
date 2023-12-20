package test

import (
	"fmt"
	"github.com/bloomberg/go-testgroup"
	"testing"
	"time"
)

func groupAnagramLogic(arr []string) [][]string {
	var result [][]string
	storeAnagram := map[string][]string{}

	for _, s := range arr {
		var count [26]int
		var breal string
		for _, c := range s {
			count[c-'a']++
		}
		breal = "s"
		fmt.Println(breal)
	}
	fmt.Println(storeAnagram)
	//for i := 0; i < len(arr); i++ {
	//	if len(storeAnagram) == 0 {
	//		storeAnagram[arr[i]] = append(storeAnagram[arr[i]], arr[i])
	//	}
	//	temp := strings.Split(arr[i], "")
	//	fmt.Println(temp)
	//}
	result = append(result, arr)

	return result
}

type groupAnagrams struct{}

func TestGroupAnagram(t *testing.T) {
	testgroup.RunSerially(t, &groupAnagrams{})
	//testgroup.RunInParallel(t, &groupAnagrams{})
}

func (group *groupAnagrams) GroupAnagram1(t *testgroup.T) {

	var expected [][]string
	expected = append(expected, []string{"eat", "tea", "tan", "ate", "nat", "bat"})

	payload := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	t.Equal(expected, groupAnagramLogic(payload))
}

func (group *groupAnagrams) GroupAnagram2(t *testgroup.T) {
	var expected [][]string
	expected = append(expected, []string{""})

	payload := []string{""}

	t.Equal(expected, groupAnagramLogic(payload))
}

func (group *groupAnagrams) GroupAnagram3(t *testgroup.T) {
	var expected [][]string
	expected = append(expected, []string{"a"})
	time.Sleep(3 * time.Second)
	payload := []string{"a"}

	t.Equal(expected, groupAnagramLogic(payload))
}
