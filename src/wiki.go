package main

import "os"

/*
	1. 메서드 전달 및 저장을 위한 data structure
	2. net/http 패키지 사용하여 웹어플리케이션 구축
	3. html/template 패키지를 사용하여 HTML 템플릿 생성
	4. regexp 패키지로 사용자 input validate
	5. closure 사용
*/

type Page struct {
	Title string
	Body  []byte // body에는 이후 사용될 io 라이브러리때문에 []byte 사용
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
	// os.WriteFile: byte slice를 파일로 저장하는 표준 라이브러리
}

func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := os.ReadFile(filename)
	return &Page{Title: title, Body: body}
}
