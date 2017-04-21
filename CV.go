package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	uuid "github.com/nu7hatch/gouuid"
)

type cvData struct {
	ImageName string `bson:"imagename"`
	Name      string `bson:"name"`
	Email     string `bson:"email"`
	Phone     string `bson:"phone"`
	Web       string `bson:"web"`
	Addr      string `bson:"addr"`
	Gender    string `bson:"gender"`

	// Education below
	XPer  string `bson:"xper"`
	Xname string `bson:"xiname"`
	Xyear string `bson:"xyear"`

	XIIper  string `bson:"xiiper"`
	XIIname string `bson:"xiiname"`
	XIIyear string `bson:"xiiyear"`

	Gradper  string `bson:"gradper"`
	Gradname string `bson:"gradname"`
	Gradyear string `bson:"gradyear"`

	PGper  string `bson:"pgper"`
	PGname string `bson:"pgname"`
	PGyear string `bson:"pgyear"`
	// Education above

	Objective string `bson:"objective"`

	// Job part below
	JobTitle1 string `bson:"jobtitle1"`
	ComName1  string `bson:"comname1"`
	JobDates1 string `bson:"jobdates1"`

	JobTitle2 string `bson:"jobtitle2"`
	ComName2  string `bson:"comname2"`
	JobDates2 string `bson:"jobdates2"`

	JobTitle3 string `bson:"jobtitle3"`
	ComName3  string `bson:"comname3"`
	JobDates3 string `bson:"jobdates3"`

	JobTitle4 string `bson:"jobtitle4"`
	ComName4  string `bson:"comname4"`
	JobDates4 string `bson:"jobdates4"`

	JobTitle5 string `bson:"jobtitle5"`
	ComName5  string `bson:"comname5"`
	JobDates5 string `bson:"jobdates5"`

	JobTitle6 string `bson:"jobtitle6"`
	ComName6  string `bson:"comname6"`
	JobDates6 string `bson:"jobdates6"`

	JobTitle7 string `bson:"jobtitle7"`
	ComName7  string `bson:"comname7"`
	JobDates7 string `bson:"jobdates7"`

	JobTitle8 string `bson:"jobtitle8"`
	ComName8  string `bson:"comname8"`
	JobDates8 string `bson:"jobdates8"`
	//Job part above

	Skill1 string `bson:"skill1"`
	Skill2 string `bson:"skill2"`
	Skill3 string `bson:"skill3"`
	Skill4 string `bson:"skill4"`
	Skill5 string `bson:"skill5"`
	Skill6 string `bson:"skill6"`
	Skill7 string `bson:"skill7"`
	Skill8 string `bson:"skill8"`
	Skill9 string `bson:"skill9"`
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./*.html"))
}

func main() {
	http.Handle("./",
		http.StripPrefix("/",
			http.FileServer(http.Dir("./"))))

	http.HandleFunc("/", index)
	http.HandleFunc("/new.html", newSub)

	fmt.Println("Listening at 8888")
	http.ListenAndServe(":8888", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Println(r.URL.Path)
}

func newSub(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "new.html", nil)

	if r.Method == http.MethodPost {
		//TODO: Handle Post request
	}

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Println(r.URL.Path)
}

func uploadPage(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		// To recieve a file, for html its going to be input type="file" name="file"
		src, hdr, err := r.FormFile("imageFile")
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}
		defer src.Close()

		nameUpdate, _ := uuid.NewV4()
		hdr.Filename = nameUpdate.String()
		filePath := "./assets/images/" + hdr.Filename + ".png"

		//writing file by creating one
		dst, err := os.Create(filePath)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}
		defer dst.Close()

		// copy the uploaded file
		_, err = io.Copy(dst, src)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}

		err = uploadedFileDb(filePath)

		if err != nil {
			fmt.Fprintf(w, `{"response":"Failed :("}`)
			os.Remove(filePath)
		}
		fmt.Fprintf(w, `{"response":"Success"}`)
	}

	log.Println(r.URL.Path)
}

func uploadedFileDb(path string) error {

	//TODO: Handle uploaded file in DB

	return nil
}
