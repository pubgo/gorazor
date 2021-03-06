package cases

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func Login(msg string) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n<!DOCTYPE html>\n<html>\n<head>\n  <meta charset=\"utf-8\" />\n  <link rel=\"stylesheet\" href=\"/css/bootstrap.min.css\">\n  <title>登陆后台</title>\n  <style>\n  body {\n    padding-top: 40px;\n    padding-bottom: 40px;\n    background-color: #eee;\n  }\n\n  .form-signin {\n    max-width: 330px;\n    padding: 15px;\n    margin: 0 auto;\n  }\n  .form-signin .form-signin-heading,\n  .form-signin .checkbox {\n    margin-bottom: 10px;\n  }\n  .form-signin .checkbox {\n    font-weight: normal;\n  }\n  .form-signin .form-control {\n    position: relative;\n    height: auto;\n    -webkit-box-sizing: border-box;\n    -moz-box-sizing: border-box;\n    box-sizing: border-box;\n    padding: 10px;\n    font-size: 16px;\n  }\n  .form-signin .form-control:focus {\n    z-index: 2;\n  }\n  .form-signin input[type=\"email\"] {\n    margin-bottom: -1px;\n    border-bottom-right-radius: 0;\n    border-bottom-left-radius: 0;\n  }\n  .form-signin input[type=\"password\"] {\n    margin-bottom: 10px;\n    border-top-left-radius: 0;\n    border-top-right-radius: 0;\n  }\n  </style>\n\n</head>\n<body>\n  <div class=\"container\">\n    <form class=\"form-signin\" role=\"form\" method=\"post\" action=\"/admin/login\">\n      <h2 class=\"form-signin-heading\">请登陆</h2>\n      ")
	if msg != "" {

		_buffer.WriteString("<div class=\"alert alert-danger\">")
		_buffer.WriteString(gorazor.HTMLEscape(msg))
		_buffer.WriteString("</div>")

	}
	_buffer.WriteString("\n\n    <input type=\"text\" name=\"username\" class=\"form-control\" placeholder=\"用户名\" required autofocus>\n    <input type=\"password\" class=\"form-control\" placeholder=\"密码\" required>\n    <label class=\"checkbox\">\n      <input type=\"checkbox\" value=\"remember-me\"> 记住我\n    </label>\n    <button class=\"btn btn-lg btn-primary btn-block\" type=\"submit\">登陆</button>\n  </form>\n  </div>\n  <script src=\"/js/jquery.min.js\"></script>\n  <script src=\"/js/bootstrap.min.js\"></script>\n</body>\n</html>")

	return _buffer.String()
}
