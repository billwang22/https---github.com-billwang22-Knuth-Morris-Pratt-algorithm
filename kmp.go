package main

import (
	"fmt"
)

func findMaxComm(perfix []string, suffix []string, m_len int) int {
	maxLen := 0
	if m_len == 0 {
		//fmt.Println("return maxLen....")
		//fmt.Println(maxLen)
		return maxLen
	}
	for _, value := range perfix {
		for i := 0; i < m_len; i++ {
			if value == suffix[i] {
				comLen := len(value)
				if comLen > maxLen {
					maxLen = comLen
				}
			}
		}
	}
	//fmt.Println("return maxLen....")
	//fmt.Println(maxLen)
	return maxLen
}

func generatePartMatch(keyWord string) []int {
	//fmt.Println(keyWord)
	p_len := len(keyWord)
	keyWordArray := []rune(keyWord)
	//partMatch := []int{}
	partMatch := make([]int, p_len)

	//fmt.Println(p_len)
	for i := 1; i <= p_len; i++ {
		subWord := string(keyWordArray[0:i])
		fmt.Println(subWord)
		perfix := []string{}
		suffix := []string{}
		subLen := len(subWord)
		for j := 1; j < i; j++ {
			perfix = append(perfix, subWord[0:j])
			suffix = append(suffix, subWord[j:subLen])
			//fmt.Println(perfix)
			//fmt.Println(suffix)
		}
		//maxlen := findMaxComm(perfix, suffix, subLen)
		//partMatch.append(partMatch, maxlen)
		// fmt.Println("----------------")

		// fmt.Println(i)
		// fmt.Println(subWord)
		// fmt.Println(perfix)
		// fmt.Println(suffix)
		// fmt.Println(subLen - 1)
		// fmt.Println("----------------")
		partMatch[i-1] = findMaxComm(perfix, suffix, subLen-1)

	}
	return partMatch
}

func KMP(bigstring string, keyWord string) bool {
	foundFlag := false
	partMatchTable := generatePartMatch(keyWord)
	fmt.Println(partMatchTable)
	i, j := 0, 0
	matchLen := 0
	bigLen := len(bigstring)
	bigstringArray := []rune(bigstring)
	keyWordArray := []rune(keyWord)
	for i = 0; i < bigLen; {
		// fmt.Println("big string ----------------")
		// fmt.Println(i)
		index := bigstringArray[i]
		crword := keyWordArray[j]
		// fmt.Println(string(index))
		// fmt.Println("small string ----------------")
		// fmt.Println(j)
		// fmt.Println(string(crword))
		if string(index) == string(crword) {
			matchLen += 1
			i++
			j++
		} else {
			if matchLen != 0 {
				noMatchPos := j
				fmt.Println("no match pos ----------------")
				fmt.Println(noMatchPos)
				j = j - (matchLen - partMatchTable[noMatchPos-1])
				fmt.Println("reloaction pos ----------------")
				fmt.Println(j)
				matchLen = j
			} else {
				i++
				j = 0
			}
		}
		if matchLen == len(keyWord) {
			foundFlag = true
			fmt.Println("Found The String")
			j = 0
			matchLen = 0
		}

	}
	return foundFlag

}

func main() {
	KMP("BBC ABCDAB ABCDABCDABDE", "ABCDABD")
}
