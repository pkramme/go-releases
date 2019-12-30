package main

import (
	"flag"
	"fmt"
	"net/http"
	"encoding/json"
)

const APIBasePath = "https://api.github.com/"

type LicenseT struct {
	Name string `json:"name"`
}

type RepositoryT struct {
	Releases ReleasesT
	Fullname string `json:"full_name"`
	Description string `json:"description"`
	HtmlURL string `json:"html_url"`
	License LicenseT `json:"license"`
}

type ReleaseT struct {
	TagName string `json:"tag_name"`
	Name string `json:"name"`
	PublishedAt string `json:"published_at"`
}

func (r *ReleasesT) Get(Username string, Reponame string) (err error) {
        res, err := http.Get(APIBasePath + "repos/" + Username + "/" + Reponame + "/releases")
        if err != nil {
                return
        }
        err = json.NewDecoder(res.Body).Decode(&r)
        if err != nil {
                return
        }
        res.Body.Close()
	return nil
}

func (r *ReleasesT) PrettyPrint() {
	fmt.Println("Releases:")
        for _, v := range *r {
                fmt.Println(" -", v.TagName, "-", v.Name, "(", v.PublishedAt, ")")
        }

}

func (r *RepositoryT) Get(Username string, Reponame string) (err error) {
        res, err := http.Get(APIBasePath + "repos/" + Username + "/" + Reponame)
        if err != nil {
                return
        }
        err = json.NewDecoder(res.Body).Decode(&r)
        if err != nil {
                return
        }
        res.Body.Close()
        return nil
}

func (r *RepositoryT) PrettyPrint() {
	fmt.Println(r.Fullname, "(", r.HtmlURL, ")")
	fmt.Println("~", r.Description, "~")
	fmt.Println("Licensed under", r.License.Name)
}

type ReleasesT []ReleaseT

func main() {
	UsernameFlag := flag.String("u", "", "Github Username")
	RepoFlag := flag.String("r", "", "Github Repository")
	flag.Parse()
	var repo RepositoryT
	repo.Get(*UsernameFlag, *RepoFlag)
	repo.Releases.Get(*UsernameFlag, *RepoFlag)
	repo.PrettyPrint()
	repo.Releases.PrettyPrint()
}
