package utils

import (
	"github.com/run-bigpig/skplayer/internal/types"
	"golang.org/x/net/html"
	"slices"
	"strings"
)

func DealPlayUrl(playUrl string) []*types.Play {
	var playArr []*types.Play
	playerArr := strings.Split(playUrl, "$$$")
	if len(playerArr) > 0 {
		playlist := strings.Split(playerArr[0], "#")
		for _, v := range playlist {
			play := strings.Split(v, "$")
			if len(play) < 2 {
				continue
			}
			playArr = append(playArr, &types.Play{
				Name: play[0],
				Url:  play[1],
			})
		}
	}
	return playArr
}

// SliceRemove 移除切片中指定下标的元素
func SliceRemove[S ~[]E, E any](s S, index int) S {
	if index < 0 || index >= len(s) {
		return s
	}
	j := index + 1
	if index == len(s)-1 {
		j = len(s)
	}
	return slices.Delete(s, index, j)
}

// RemoveHTMLTags 移除字符串中的html标签
func RemoveHTMLTags(input string) string {
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		return input
	}
	var f func(*html.Node)
	var output strings.Builder
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			output.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return strings.TrimSpace(output.String())
}
