package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"log"
	"net/http"
	"strings"
	"html/template"

	"GOPHERCISES/3_CYO_Adventure/cyoa"
)

func main() {

	absPath, err := filepath.Abs("../../gopher.json")
	check(err)
	portName := flag.Int("port", 8080, "the port to start the server on")
	fileName := flag.String("file", absPath, "the JSON file containing the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *fileName)

	f, err := os.Open(*fileName)
	check(err)

	story, err := cyoa.JsonStory(f)
	check(err)

	tpl := template.Must(template.New("").Parse(overrideHandlerTmplt))

	h := cyoa.NewHandler(story, 
		cyoa.WithPathFunction(alternatePathFunction), 
		cyoa.WithTemplate(tpl)
		) // example of passing in variadic parameter options for handler
	port := fmt.Sprintf(":%v", *portName)
	fmt.Println("Starting the server on port ", port)
	log.Fatal(http.ListenAndServe(port, h))
}

func alternatePathFunction(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	path = path[7:]
	return path
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var overrideHandlerTmplt = `
<!DOCTYPE>
<html>
    <head>
        <meta charset="UTF8">
        <title>Choose Your Own Adventure</title>
    </head>

    <body>
	<section class="page">
        <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
            <p>{{.}}</p>
        {{end}}

        <ul>
        {{range .Options}}
            <li>
                <a href="/story/{{.Arc}}">{{.Text}}</a>
            </li>
        {{end}}
        </ul>
		</section>
		<style>
		body {
		  font-family: helvetica, arial;
		}
		h1 {
		  text-align:center;
		  position:relative;
		}
		.page {
		  width: 80%;
		  max-width: 500px;
		  margin: auto;
		  margin-top: 40px;
		  margin-bottom: 40px;
		  padding: 80px;
		  background: #FFFCF6;
		  border: 1px solid #eee;
		  box-shadow: 0 10px 6px -6px #777;
		}
		ul {
		  border-top: 1px dotted #ccc;
		  padding: 10px 0 0 0;
		  -webkit-padding-start: 0;
		}
		li {
		  padding-top: 10px;
		}
		a,
		a:visited {
		  text-decoration: none;
		  color: #6295b5;
		}
		a:active,
		a:hover {
		  color: #7792a2;
		}
		p {
		  text-indent: 1em;
		}
	  </style>

    </body>
</html>
`

