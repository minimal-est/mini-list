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
		http.Error(w, "JSON으로 직렬화하는 과정에서 실패했습니다.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")

	svg := `<svg xmlns="http://www.w3.org/2000/svg" width="650" height="150">`
	svg += `<style>text { fill:white; font-family: sans-serif; font-size: 14px; stroke: black; stroke-width: 2; paint
-order
: stroke;}</style>`
	svg += fmt.Sprintf(`<text x="10" y="20">📌 Latest posts by %s</text>`, author)
	for i, post := range result.Data.Content {
		if i >= 5 {
			break
		}
		date := post.CreatedAt
		if len(date) >= 10 {
			date = date[:10]
		}

		title := post.Title

		y := 40 + i*20
		svg += fmt.Sprintf(`<text x="10" y="%d">- [%s][%s] %s [%d]</text>`, y, date, post.FolderName, title, post.HitCount)
	}
	svg += `</svg>`

	w.Write([]byte(svg))
}
