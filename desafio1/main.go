package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	content := `
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8" />
				<meta http-equiv="X-UA-Compatible" content="IE=edge" />
				<meta name="viewport" content="width=device-width, initial-scale=1.0" />
				<title>Document</title>

				<link rel="preconnect" href="https://fonts.googleapis.com">
				<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
				<link href="https://fonts.googleapis.com/css2?family=Zen+Loop&display=swap" rel="stylesheet">

			</head>
			<style>
				body {
					background-image: url('https://www.pixelstalk.net/wp-content/uploads/2016/03/Logo-Supernatural-Wallpaper.png');
					background-repeat: no-repeat;
  					background-attachment: fixed;
  					background-position: center;
					text-align: center;
					color: white;
					font-size: 48px;
					font-weight: bold;
				}

				h1 {
					background: -webkit-linear-gradient(#eee, #333);
  					-webkit-background-clip: text;
  					-webkit-text-fill-color: transparent;
				}
			</style>
			<body>
				<h1>Imersão Full Cycle</h1>
			</body>
		</html>	
	`
	fmt.Fprintf(w, content)
}
