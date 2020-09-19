package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"
//befor import "C" cannot be empty line
import (
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/jordan-wright/email"
)

func FileExist(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//export SendEmail
func SendEmail(from, to, bcc, cc, subject, text, html, file, addr, identity, username, password, host, output *C.char) {
	e := email.NewEmail()
	e.From = C.GoString(from)//a <a@qq.com>

	e.To = nil
	tos := strings.Split(C.GoString(to), "|") //b@qq.com|c@qq.com
	if len(tos) > 0 {
		for _,v := range tos {
			if len(v) > 0 {
				e.To = append(e.To, v)
			}
		}
	}

	e.Bcc = nil
	bccs := strings.Split(C.GoString(bcc), "|") //c@qq.com|d@qq.com
	if len(bccs) > 0 {
		for _,v := range bccs {
			if len(v) > 0 {
				e.Bcc = append(e.Bcc, v)
			}
		}
	}

	e.Cc = nil
	ccs := strings.Split(C.GoString(cc), "|") //e@qq.com|f@qq.com
	if len(ccs) > 0 {
		for _,v := range ccs {
			if len(v) > 0 {
				e.Cc = append(e.Cc, v)
			}
		}
	}

	e.Subject = C.GoString(subject)
	e.Text = []byte(C.GoString(text)) //Text Body
	e.HTML = []byte(C.GoString(html)) //<h1>HTML is supported</h1>

	files := strings.Split(C.GoString(file), "|") //file1.txt|file2.txt
	for _, v := range files {
		if FileExist(v) == true {
			e.AttachFile(v)
		}
	}

	err := e.Send(C.GoString(addr), smtp.PlainAuth(C.GoString(identity), C.GoString(username), C.GoString(password), C.GoString(host)))
	if err == nil {
		fmt.Println("ok")
	} else {
		fmt.Println(err.Error())
		C.strcpy(output, C.CString(err.Error()))
	}
}

func main() {
	// Need a main function to make CGO compile package as C shared library
}
