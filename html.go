package stringx

import (
	"fmt"
	"regexp"
	"strings"
)

func RemoveStylesOfHtmlTag(html string, keptStyles ...string) string {
	styleTagR := regexp.MustCompile(`style="([^=>]*;?)*"`)
	keptStyleR := regexp.MustCompile(fmt.Sprintf("^%s", strings.Join(keptStyles, "|")))
	hasKeptStyles := len(keptStyles) != 0

	matches := styleTagR.FindAllStringSubmatch(html, -1)
	var toCleanStyles = make(map[string]struct{})

	for _, groups := range matches {
		if len(groups)!=2{
			continue
		}

		styles := strings.Split(groups[1], ";")
		for _, style := range styles{
			trimStr := strings.ToLower(strings.TrimSpace(style))
			if trimStr == "" || (hasKeptStyles && keptStyleR.MatchString(trimStr)) {
				continue
			}

			toCleanStyles[style] = struct{}{}
		}
	}

	cleanedHtml := html
	for cleanStyle := range toCleanStyles{
		escapedStyle := strings.Replace(cleanStyle, "(", `\(`, -1)
		escapedStyle = strings.Replace(escapedStyle, ")", `\)`, -1)
		cleanR := regexp.MustCompile(fmt.Sprintf("%s[;]?", escapedStyle))
		cleanedHtml = cleanR.ReplaceAllString(cleanedHtml, "")
	}

	return cleanedHtml
}
