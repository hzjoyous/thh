package template

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"thh/base"
	"thh/helpers"
)

var commandList = make(map[string]base.Console)

func GetAllConsoles() map[string]base.Console {
	return commandList
}

func addConsole(signature string, description string, handle func()) {
	c := base.Console{Signature: signature, Description: description, Handle: handle}
	commandList[c.Signature] = c
}
func init() {
	addConsole("template", "this is a template", mainAction)
	addConsole("runPath", "runPath", getRunPath)
	addConsole("t1", "t1", t1)
}

func getRunPath() {
	binary, _ := os.Executable()
	root := filepath.Dir(filepath.Dir(binary))
	fmt.Println(root)
}

func t1() {

	var buf bytes.Buffer

	buf.WriteString(`#include <iostream>
using namespace std;
int main(){
	cout << "1";
	int x;
	cin >> x;
	switch(x){
	`)

	for i := 0; i <= 99999; i++ {
		t := i
		t2 := 0

		buf.WriteString(`case `)

		buf.WriteString(helpers.ToString(t))
		buf.WriteString(`:
		`)

		buf.WriteString(`cout<<""<<`)
		nlen := 1
		switch {
		case i > 10000:
			nlen = 5
			break
		case i > 1000:
			nlen = 4
			break
		case i > 100:
			nlen = 3
			break
		case i > 10:
			nlen = 2
			break
		}
		buf.WriteString(helpers.ToString(nlen))
		buf.WriteString(`<<"位数"`)
		buf.WriteString(`<<endl;
		`)

		buf.WriteString(`cout<<"个位:"<<`)
		t = i % 10
		buf.WriteString(helpers.ToString(t))
		buf.WriteString(`<<endl;
		`)
		t2 += t
		if i > 10 {

			t = i / 10 % 10
			t2 = t2*10 + t
			buf.WriteString(`cout<<"十位:"<<`)
			buf.WriteString(helpers.ToString(t))
			buf.WriteString(`<<endl;
		`)
		}
		if i > 100 {
			t = i / 100 % 10
			t2 = t2*10 + t
			buf.WriteString(`cout<<"百位:"<<`)
			buf.WriteString(helpers.ToString(t))
			buf.WriteString(`<<endl;
		`)
		}
		if i > 1000 {
			t = i / 1000 % 10
			t2 = t2*10 + t
			buf.WriteString(`cout<<"千位:"<<`)
			buf.WriteString(helpers.ToString(t))
			buf.WriteString(`<<endl;
		`)
		}
		if i > 10000 {
			t = i / 10000 % 10
			t2 = t2*10 + t
			buf.WriteString(`cout<<"万位:"<<`)
			buf.WriteString(helpers.ToString(t))
			buf.WriteString(`<<endl;
		`)
		}
		buf.WriteString(`cout<<"反转:"<<`)
		buf.WriteString(helpers.ToString(t2))
		buf.WriteString(`<<endl;
		`)

		buf.WriteString(`break;
	`)

	}
	buf.WriteString(`
	}
	return 1;
}`)
	_ = helpers.FilePutContents("./tmp/main.cpp", buf.Bytes(), true)
}

func mainAction() {
	sort1()
	fmt.Println("this is template main")
}

func sort1() {
	numberList := []int{5, 4, 3, 2, 1}
	length := len(numberList)
	endTag := false
	for i := 0; i < length; i++ {
		endTag = true
		for j := 0; j < length-i-1; j++ {
			if numberList[j] < numberList[j+1] {
				endTag = false
				t := numberList[j+1]
				numberList[j+1] = numberList[j]
				numberList[j] = t
			}
		}
		fmt.Println(numberList)
		if endTag == true {
			break
		}
	}

}
