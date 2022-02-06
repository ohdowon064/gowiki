package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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

func handler(w http.ResponseWriter, r *http.Request) {
	/*
		- http.ResponseWriter: HTTP 서버의 응답을 수집한다. 해당 응답에 무언가를 쓰기위해서 파라미터로 받음
		- http.Request: client HTTP 요청을 나타낸다.
		- r.URL.Path[1:]: 요청 URL의 Path 요소, root url "/"의 다음부터 슬라이싱
			- 예를 들어, http://localhost:8080/monkeys에 요청하면
			- Hi, there, I love monkeys!를 반환

		- fmt에서 print앞에 F가 붙으면 파일 입출력을 뜻한다.
		- w 즉, http.ResponseWriter에 문자열을 입력한다는 뜻이다.
	*/
	fmt.Fprintf(w, "Hi, there, I love %s!", r.URL.Path[1:]) // w.Write([]byte("Hi~"))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	t, _ := template.ParseFiles("view.html")
	t.Execute(w, p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	/*
		- template(템플릿): 재사용 가능한 양식
		- template.ParseFiles: edit.html를 읽고 *template.Template을 반환한다.
		- t.Execute: 템플릿을 실행시킨다. -> 생성된 HTML을 http.ResponseWriter (w변수)에 작성한다.
		- edit.html에서 .Title, .Body 식별자들은 p.Title, p.Body에 대응된다.
		- {{ printf "%s" .Body }}: 함수호출의 결과값(bytes stream이 아닌 string으로)이 된다.

		- html/template 패키지는 only safe and correct-looking HTML만으로 template action에 의해 생성되는 것을 보장한다.
		- 사용자가 입력한 <> 태그사인을 자동으로 부등호로 이스케이프해준다.
	*/
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

func main() {
	/*
		log.Fatal로 http.ListenAndServe를 넘기는 이유
		- ListenAndServe는 에러발생 시 항상 에러를 반환한다.
		- 따라서 해당 에러를 기록하기 위해서 log.Fatal 사용
	*/
	http.HandleFunc("/", handler)          // web root url "/"로 들어오는 모든 요청에 대해 hanlder로 처리하도록 http 패키지에 지시한다.
	http.HandleFunc("/view/", viewHandler) // "/view/"로 오는 요청을 처리한다.
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil)) // 8080포트에서 listen(요청대기), 프로그램이 종료될 때가지 블록된다.
}
