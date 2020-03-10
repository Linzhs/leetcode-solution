package api

type pair struct {
	word  string
	level int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	allComboDict := make(map[string]*[]string, len(wordList))

	// 预处理 生成词典 fuzzyWord => word list
	for _, word := range wordList {
		for i := 0; i < len(beginWord); i++ {
			fuzzyWord := word[0:i] + "*"
			if i < len(beginWord)-1 {
				fuzzyWord += word[i+1:]
			}

			list, ok := allComboDict[fuzzyWord]
			if ok {
				*list = append(*list, word)
			} else {
				allComboDict[fuzzyWord] = &[]string{word}
			}
		}
	}

	queue := make([]pair, 0, len(wordList))
	queue = append(queue, pair{
		word:  beginWord,
		level: 1,
	})

	visited := make(map[string]struct{})
	visited[beginWord] = struct{}{}

	// bfs
	for len(queue) > 0 {
		elem := queue[0]
		queue = queue[1:]

		word := elem.word
		level := elem.level
		for i := 0; i < len(beginWord); i++ {
			fuzzyWord := word[0:i] + "*"
			if i < len(beginWord)-1 {
				fuzzyWord += word[i+1:]
			}
			list, ok := allComboDict[fuzzyWord]
			if !ok {
				continue
			}

			for _, v := range *list {
				if v == endWord {
					return level + 1
				}

				if _, ok := visited[v]; !ok {
					visited[v] = struct{}{}
					queue = append(queue, pair{
						word:  v,
						level: level + 1,
					})
				}
			}
		}
	}

	return 0
}
