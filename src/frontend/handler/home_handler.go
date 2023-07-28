/**
 * @description: the handlers for home page
 * @author: Xie Yuting
 * @date: 2023-07-28
 */

package frontend

import (
	"fmt"
	"html/template"
	"net/http"
)

// var (
// 	templates = template.Must(template.New("home").ParseGlob("../template/home.html"))
// )

type HomeHandler struct{}

func (h *HomeHandler) GetHomePage(w http.ResponseWriter, r *http.Request) {
	// dir, _ := os.Getwd()
	// fmt.Println("****** current wd: ", dir, " ******")
	// filenames, _ := filepath.Glob("template/*.html")
	// for _, filename := range filenames {
	// 	fmt.Println("***** ", filename, " *****")
	// }
	templates := template.Must(template.New("").ParseGlob("src/frontend/template/*.html"))
	if err := templates.ExecuteTemplate(w, "home", map[string]interface{}{
		"name": "Xie Yuting",
	}); err != nil {
		fmt.Println("An error occurred: ", err)
	}

	fmt.Println(w)
}
