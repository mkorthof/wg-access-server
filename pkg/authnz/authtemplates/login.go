package authtemplates

import (
	_ "embed"
	"html/template"
	"io"

	"github.com/place1/wg-access-server/pkg/authnz/authruntime"
)

type LoginPage struct {
	Providers []*authruntime.Provider
}

// if we got embedded file, use that instead of the var below
func RenderLoginPage(w io.Writer, data LoginPage) error {
	var loginContent string
	if (loginFile != nil) {
		loginContent = string(loginFile)
	} else {
		loginContent = loginPage
	}
	tpl, err := template.New("login-page").Parse(loginContent)
	if err != nil {
		return err
	}
	return tpl.Execute(w, data)
}

//go:embed /login.html.tmpl
var loginFile []byte

const loginPage string = `
<!DOCTYPE html>
<html>
<head>
	<title>WireGuard Access Server - Sign In</title>
	<style>
		* {
			font-family: monospace;
			font-size: 16px;
			-webkit-font-smoothing: antialiased;
			-moz-osx-font-smoothing: grayscale;
		}

		body {
			background-color: #3899c9;
		}

		.top {
			position: absolute;
			top: 15%;
			left: 50%;
			transform: translate(-50%, -50%);
			background-color: #fff;
			width: 550px;
			padding: 20px;
			box-shadow: 0 10px 20px rgba(0, 0, 0, 0.19), 0 6px 6px rgba(0, 0, 0, 0.23);
			text-align: center;
		}

		.top h1 {
			font-size: 30px;
			margin-top: 5px;
			margin-bottom: 5px;
		}

		.filter-blue {
			filter: invert(27%) sepia(31%) saturate(30%) hue-rotate(358deg) brightness(98%) contrast(93%);
		}

		.form {
			position: absolute;
			top: 50%;
			left: 50%;
			transform: translate(-50%, -50%);
			background-color: #fff;
			width: 285px;
			padding: 40px;
			box-shadow: 0 10px 20px rgba(0, 0, 0, 0.19), 0 6px 6px rgba(0, 0, 0, 0.23);
		}

		.form h2 {
			margin: 0 0 35px;
			text-align: center;
			line-height: 1;
			color: black;
			font-size: 22px;
			font-weight: 400;
		}

		.form input {
			outline: none;
			display: block;
			width: 100%;
			padding: 10px 15px;
			border: 1px solid #ccc;
			color: #ccc;
			box-sizing: border-box;
		}

		.form a {
			display: block;
		}

		.form > * {
			margin: 0 0 20px;
		}

		.form > *:last-child {
			margin-bottom: 0px;
		}

		.form input:focus {
			color: #333;
			border: 1px solid #44c4e7;
		}

		.form button {
			cursor: pointer;
			background: #44c4e7;
			width: 100%;
			padding: 10px 15px;
			border: 0;
			color: #fff;
			text-transform: capitalize;
		}

		.form button:hover {
			background: #369cb8;
		}

		.form hr {
			position: relative;
			width: 55%;
			margin-left: auto;
			margin-right: auto;
			overflow: visible;
			background-color: #4d4d4d;
		}

		.form hr:after {
			content: "";
			position: absolute;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -50%);
			width: 4px;
			height: 4px;
			background-color: #4d4d4d;
			border-radius: 50%;
		}
	</style>
</head>
<body>
	<section class="top">
		<img width="100" class="filter-blue" src="data:image/svg+xml;base64,PHN2ZyByb2xlPSJpbWciIHZpZXdCb3g9IjAgMCAyNCAyNCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48dGl0bGU+V2lyZUd1YXJkIGljb248L3RpdGxlPjxwYXRoIGQ9Ik0yMy45OCAxMS42NDVTMjQuNTMzIDAgMTEuNzM1IDBDLjQxOCAwIC4wNjQgMTEuMTcuMDY0IDExLjE3Uy0xLjYgMjQgMTEuOTk3IDI0QzI1LjA0IDI0IDIzLjk4IDExLjY0NSAyMy45OCAxMS42NDV6TTguMTU1IDcuNTc2YzIuNC0xLjQ3IDUuNDY5LS41NzEgNi42MTggMS42MzguMjE4LjQxOS4yNDYgMS4wNjMuMTA4IDEuNTAzLS40NzcgMS41MTYtMS42MDEgMi4zNjYtMy4xNDUgMi43MjguNDU1LS4zOS44MTctLjgzMi45MzMtMS40NDJhMi4xMTIgMi4xMTIgMCAwIDAtLjM2NC0xLjY3NyAyLjE0IDIuMTQgMCAwIDAtMi40NjUtLjc1Yy0uOTUuMzYtMS40NyAxLjIyOC0xLjM3NyAyLjI5NC4wODcuOTkuODM5IDEuNjMyIDIuMjQ1IDEuODc2LS4yMS4xMTEtLjM3Mi4xOTMtLjUzLjI4MWE1LjExMyA1LjExMyAwIDAgMC0xLjY0NCAxLjQzYy0uMTQzLjE5Mi0uMjQuMjA4LS40NTguMDc1LTIuODI3LTEuNzI5LTMuMDA5LTYuMDY3LjA3OC03Ljk1NnpNNi4wNCAxOC4yNThjLS40NTUuMTE2LS44OTUuMjg2LTEuMzU5LjQzOC4yMjctMS41MzIgMi4wMjEtMi45NDMgMy41MzktMi43ODJhMy45MSAzLjkxIDAgMCAwLS43NCAyLjA3MmMtLjUwNC4wOTMtLjk4LjE1NS0xLjQ0LjI3MnpNMTUuNzAzIDMuM2MuNDQ4LjAxNy44OTguMDEgMS4zNDcuMDJhMi4zMjQgMi4zMjQgMCAwIDEgLjMzNC4wNDcgMy4yNDkgMy4yNDkgMCAwIDEtLjM0LjQzNGMtLjE2LjE1LS4zNDEuMjk2LS41NzMuMDY5LS4wNTUtLjA1NS0uMTg3LS4wNDItLjI4My0uMDQ0LS40NDctLjAwNS0uODk0LS4wMi0xLjM0LS4wMDNhOC4zMjMgOC4zMjMgMCAwIDAtMS4xNTQuMTE4Yy0uMDcyLjAxMy0uMTc4LjI1LS4xNDYuMzM4LjA3OC4yMDcuMTkxLjQzNS4zNTkuNTY3LjYxOS40OSAxLjI3Ny45MjggMS45IDEuNDEzLjYwNC40NzIgMS4xNjcuOTkgMS41MSAxLjcuNDQ2LjkyOC40NiAxLjkuMjY3IDIuODc3LS4zMjIgMS42My0xLjE0NyAyLjk4LTIuNDgzIDMuOTYyLS41MzguMzk1LTEuMjA1LjYyLTEuODIxLjkwMy0uNTQzLjI1LTEuMS40NjUtMS42NDQuNzEyLS45OC40NDYtMS41MyAxLjUxLTEuMzY5IDIuNjE1LjE0OSAxLjAxNSAxLjA0IDEuODYyIDIuMDU5IDIuMDM3IDEuMjIzLjIxIDIuNDg2LS41ODYgMi43ODUtMS44My4zMzYtMS4zOTctLjQyMy0yLjY0Ni0xLjg0NS0zLjAyNGwtLjI1Ni0uMDY2Yy4zOC0uMTcuNzA4LS4yOTEgMS4wMTItLjQ1OHEuNzkzLS40MzcgMS41NTgtLjkyNWMuMTUtLjA5Ni4yMzEtLjA5Ni4zNi4wMTQuOTc3Ljg0NiAxLjU2IDEuODk4IDEuNzI0IDMuMTg3LjI3IDIuMTM1LS43NCA0LjA5Ni0yLjY0NiA1LjEwMS0yLjk0OCAxLjU1NS02LjU1Ny0uMjE1LTcuMjA4LTMuNDg0LS41NTgtMi44IDEuNDE4LTUuMzQgMy43OTctNS44MyAxLjAyMy0uMjExIDEuOTU4LS42MzcgMi42ODUtMS40MjUuNDctLjUwOC42OTctLjk0NC43NzUtMS4xNDFhMy4xNjUgMy4xNjUgMCAwIDAgLjIxNy0xLjE1OCAyLjcxIDIuNzEgMCAwIDAtLjIzNy0uOTkyYy0uMjQ4LS41NjYtMS4yLTEuNDY2LTEuNDM1LTEuNjU2bC0yLjI0LTEuNzU0Yy0uMDc5LS4wNjUtLjE2OC0uMDYtLjM2LS4wNDctLjIzLjAxNi0uODE1LjA0OC0xLjA2Ny0uMDE4LjIwNC0uMTU1Ljc2LS4zOCAxLS41Ni0uNzI2LS40OS0xLjU1NC0uMzE0LTIuMzE1LS40Ni4xNzYtLjMyOCAxLjA0Ni0uODMxIDEuNTQxLS44ODhhNy4zMjMgNy4zMjMgMCAwIDAtLjEzNS0uODIyYy0uMDMtLjExMS0uMTU0LS4yMi0uMjYzLS4yODMtLjI2Mi0uMTU0LS41NDEtLjI4MS0uODQzLS40MzRhMS43NTUgMS43NTUgMCAwIDEgLjkwNi0uMjggMy4zODUgMy4zODUgMCAwIDEgLjkwOC4wODhjLjU0LjEyMy45Ny4wNDIgMS4zOTktLjMyNC0uMzM4LS4xMzYtLjY3Ni0uMjYtMS4wMDMtLjQwN2E5Ljg0MyA5Ljg0MyAwIDAgMS0uOTQyLS40OTNjLjg1LjExOCAxLjY3MS40MzcgMi41NC4zMmwuMDIyLS4xMTgtMi4wMTgtLjQ3YzEuMjAzLS4xMSAyLjMyMy0uMTI4IDMuMzg0LjM4OC4yOTkuMTQ2LjYxLjI2Ni44OTcuNDMyLjE0LjA4LjIzMy4yNC4zNDguMzY1LjA5LjA5OC4xNjQuMjMuMjc2LjI5LjQyNC4yMjUuODkuMjM0IDEuMzY2LjIyM2wuMDEtLjE2Yy40NzkuMTUgMS4wMTcuNzAyIDEuMDE3IDEuMTA1LS43NzYgMC0xLjU1LS4wMDMtMi4zMjUuMDA0LS4wODMgMC0uMTY1LjA2MS0uMjQ3LjA5NC4wNzguMDQ2LjE1NS4xMjguMjM1LjEzMXogTTE0LjcwMyAyLjE1M2EuMTE4LjExOCAwIDAgMC0uMDE2LjE5LjE3OS4xNzkgMCAwIDAgLjI0Ni4wNjVjLjA3NS0uMDM4LjE0OC0uMDc4LjIzOC0uMTI1LS4wNzItLjA2Mi0uMTMtLjExNC0uMTktLjE2My0uMTA2LS4wODctLjE5My0uMDMyLS4yNzguMDMzeiIvPjwvc3ZnPg=="/>
		<h1>WireGuard Access Server</h1>
	</section>
	<section class="form">
		<h2>Sign In</h2>

			{{range $i, $p := .Providers}}
				<a href="/signin/{{$i}}">
					<button>{{$p.Type}}</button>
				</a>
			{{end}}

			<!--
			<form autocomplete="off">
				<input placeholder="Username" type="text" id="username"></input>
				<input placeholder="Password" type="password" id="password"></input>
				<button id="submit">Login</button>
			</form>
			<hr />
			-->

	</section>
</body>
</html>
`