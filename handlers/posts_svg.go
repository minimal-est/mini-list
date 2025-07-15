package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"mini-list/types"
	"net/http"
	"time"
)

func PostsSvgHandler(w http.ResponseWriter, r *http.Request) {
	author := chi.URLParam(r, "archive")

	apiURL := fmt.Sprintf("https://minimalest.kr/api/archive/%s/post/preview", author)
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(apiURL)
	if err != nil {
		http.Error(w, "포스트를 불러오는 데 실패했습니다.", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	result := &types.Response{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		http.Error(w, "JSON 직렬화 실패", http.StatusInternalServerError)
		return
	}

	maxLen := len("📌 Latest posts by " + author)
	lines := []string{}
	for i, post := range result.Data.Content {
		if i >= 5 {
			break
		}
		date := post.CreatedAt
		if len(date) > 10 {
			date = date[:10]
		}
		title := post.Title
		line := fmt.Sprintf("- [%s][%s] %s [%d]", date, post.FolderName, title, post.HitCount)
		lines = append(lines, line)
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	svgWidth := maxLen*5 + 20
	svgHeight := 40 + len(lines)*20

	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")

	svg := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d">`, svgWidth, svgHeight)
	svg += `<style>text { font-family: sans-serif; font-size: 14px; }</style>`
	svg += `<rect width="100%" height="100%" rx="10" fill="#f9f9f9"/>`

	header := fmt.Sprintf("📌 Latest posts by %s", author)
	svg += `<text x="10" y="20" stroke="white" stroke-width="3" fill="none">` + header + `</text>`
	svg += `<text x="10" y="20" fill="#111">` + header + `</text>`

	for i, line := range lines {
		y := 40 + i*20
		svg += fmt.Sprintf(`<text x="10" y="%d" stroke="white" stroke-width="2" fill="none">%s</text>`, y, line)
		svg += fmt.Sprintf(`<text x="10" y="%d" fill="#333">%s</text>`, y, line)
	}

	svg += `</svg>`
	w.Write([]byte(svg))
}
