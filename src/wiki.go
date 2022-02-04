package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte // body에는 이후 사용될 io 라이브러리때문에 []byte 사용
}

func (p *Page) save() error {
	/*
		- Page 구조체 메서드로 선언.
		- save 메소드는 텍스트파일에 페이지의 Body를 저장한다.
		- Title을 파일 이름으로 사용한다.
		- WriteFile을 리턴한다. -> 추후 에러발생시 핸들링을 위해서 에러값이 반환된다.
			- 에러가 없으면(잘되면) pointer의 zero value인 nil을 반환한다.
		- 0600은 8진법 정수로써, 해당 파일은 현재 유저에게만 R/W 권한을 준다는 의미.
	*/

	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
	// os.WriteFile: byte slice를 파일로 저장하는 표준 라이브러리
}

func loadPage(title string) (*Page, error) {
	/*
		- 파일이름(title)을 파라미터로 받아 파일 내용을 읽어온다.
		- 읽어온 내용을 새로운 body변수에 넣고 Page 구조체 생성
		- 파리미터로 받은 title, 파일을 읽어서 만든 body로 생성된 Page 리터럴 포인터를 반환
			- 리터럴은 값 그자체를 의미한다.

		- 리터럴이란?
			- 소스코드의 고정된 값을 대표하는 용어, 문자 그대로의 값
			page = Page{Title: title, Body: body}
			return &page
			- 이렇게 했으면 Page 변수 포인터를 반환하는 것이지만
			- 바로 return &Page{} 이렇게 했기 때문에 Page 리터럴 포인터를 반환한 것이다.

		- 표준 라이브러리 os.ReadFile 함수는 []byte, error를 반환한다.
		- _는 blank identifier를 의미한다. -> 안쓴다는 뜻, 아무것에도 할당하지않는다.
		- error를 함께 반환하여 에러 핸들링을 한다.
			- 두번째 반환값(error)을 체크한다.
			- nil인 경우 성공, nil 아니면 에러발생으로 핸들링 가능
	*/
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save() // TestPage.txt 파일이 저장된다.

	p2, _ := loadPage("TestPage") // p2 구조체로 파일내용이 읽어진다.
	fmt.Println(string(p2.Body))
}
