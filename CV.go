package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"

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
	DoB       string `bson:"dob"`
	Country   string `bson:"country"`

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
	http.Handle("/assets/",
		http.StripPrefix("/assets",
			http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", index)
	http.HandleFunc("/new.html", newSub)
	http.HandleFunc("/contact.html", contact)
	http.HandleFunc("/policy.html", policy)
	http.HandleFunc("/help.html", help)

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
	var data cvData

	if r.Method == http.MethodPost {
		data.ImageName = uploadPage(w, r)
		data.Name = r.FormValue("flname")
		data.Email = r.FormValue("umail")
		data.Phone = r.FormValue("ucontact")
		data.Web = r.FormValue("web")
		data.DoB = r.FormValue("dob")
		data.Addr = r.FormValue("addr")
		data.Gender = r.FormValue("gender")
		data.Country = r.FormValue("con")

		data.XPer = r.FormValue("xper")
		data.Xname = r.FormValue("xname")
		data.Xyear = r.FormValue("xyear")

		data.XIIper = r.FormValue("xiiper")
		data.XIIname = r.FormValue("xiiname")
		data.XIIyear = r.FormValue("xiiyear")

		data.Objective = r.FormValue("uobj")

		data.ComName1 = r.FormValue("comname1")
		data.ComName2 = r.FormValue("comname2")
		data.ComName3 = r.FormValue("comname3")
		data.ComName4 = r.FormValue("comname4")
		data.ComName5 = r.FormValue("comname5")
		data.ComName6 = r.FormValue("comname6")
		data.ComName7 = r.FormValue("comname7")
		data.ComName8 = r.FormValue("comname8")

		data.JobTitle1 = r.FormValue("jobtitle1")
		data.JobTitle2 = r.FormValue("jobtitle2")
		data.JobTitle3 = r.FormValue("jobtitle3")
		data.JobTitle4 = r.FormValue("jobtitle4")
		data.JobTitle5 = r.FormValue("jobtitle5")
		data.JobTitle6 = r.FormValue("jobtitle6")
		data.JobTitle7 = r.FormValue("jobtitle7")
		data.JobTitle8 = r.FormValue("jobtitle8")

		data.JobDates1 = r.FormValue("jobdate1")
		data.JobDates2 = r.FormValue("jobdate2")
		data.JobDates3 = r.FormValue("jobdate3")
		data.JobDates4 = r.FormValue("jobdate4")
		data.JobDates5 = r.FormValue("jobdate5")
		data.JobDates6 = r.FormValue("jobdate6")
		data.JobDates7 = r.FormValue("jobdate7")
		data.JobDates8 = r.FormValue("jobdate8")

		data.Skill1 = r.FormValue("skill1")
		data.Skill2 = r.FormValue("skill2")
		data.Skill3 = r.FormValue("skill3")
		data.Skill4 = r.FormValue("skill4")
		data.Skill5 = r.FormValue("skill5")
		data.Skill6 = r.FormValue("skill6")
		data.Skill7 = r.FormValue("skill7")
		data.Skill8 = r.FormValue("skill8")
		data.Skill9 = r.FormValue("skill9")

		err := dataInDb(data)
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}

	err := tpl.ExecuteTemplate(w, "new.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Println(r.URL.Path)
}

func uploadPage(w http.ResponseWriter, r *http.Request) string {

	if r.Method == http.MethodPost {
		// To recieve a file, for html its going to be input type="file" name="file"
		src, hdr, err := r.FormFile("pic")
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
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
		}
		defer dst.Close()

		// copy the uploaded file
		_, err = io.Copy(dst, src)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
		}

		if err != nil {
			os.Remove(filePath)
		}
		return filePath
	}

	log.Println(r.URL.Path)
	return ""
}

func dataInDb(data cvData) error {
	session, err := mgo.Dial("mongodb://localhost/")

	if err != nil {
		log.Println(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("CVDB").C("CVCol")
	err = c.Insert(data)

	return err
}
func help(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "help.html", nil)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Println(r.URL.Path)
}
func policy(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "policy.html", nil)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Println(r.URL.Path)
}
func contact(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.html", nil)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Println(r.URL.Path)
}
