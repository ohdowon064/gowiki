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

/*
	- save 메소드는 텍스트파일에 페이지의 Body를 저장한다.
	- Title을 파일 이름으로 사용한다.
	- WriteFile을 리턴한다. -> 추후 에러발생시 핸들링을 위해서 에러값이 반환된다.
		- 에러가 없으면(잘되면) pointer의 zero value인 nil을 반환한다.
	- 0600은 8진법 정수로써, 해당 파일은 현재 유저에게만 R/W 권한을 준다는 의미.
*/
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
