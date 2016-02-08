package yado

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

type YadoInfo struct {
	Name    string
	Url     string
	Vacants []string
}

const base_url string = "https://as.its-kenpo.or.jp"
const yado_list_page_title string = "直営・通年・夏季保養施設(空き照会)"

func GetYadoInfo() []YadoInfo {

	// 1. get url list for yados
	m := getYadoUrls()
	// 2. get vacants information from each yado page
	ret := make([]YadoInfo, 0, 10)
	for name, url := range m {
		vacants := getYadoVacants(url)
		ret = append(ret, YadoInfo{Name: name, Url: url, Vacants: vacants})
	}
	// 3. return all information

	return ret
}

func getYadoUrls() map[string]string {
	doc, err := goquery.NewDocument(base_url)
	if err != nil {
		log.Fatal(err)
	}
	ret := map[string]string{}
	doc.Find("#container > .request-box > .service_category").Each(func(i int, s *goquery.Selection) {
		a := s.Find("a")
		if a.Text() == yado_list_page_title {
			url, _ := a.Attr("href")
			doc, err := goquery.NewDocument(base_url + url)
			if err != nil {
				log.Fatal(err)
			}
			doc.Find(".request-box > form > ul > li > a").Each(func(i int, s *goquery.Selection) {
				ret[s.Text()], _ = s.Attr("href")
			})
		}
	})
	return ret
}

func getYadoVacants(url string) []string {
	doc, err := goquery.NewDocument(base_url + url)
	vacants := make([]string, 1, 10)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".request-box > ul > li > a").Each(func(i int, s *goquery.Selection) {
		url, _ = s.Attr("href")
		fmt.Printf("%s %s\n", s.Text(), url)
		doc, err := goquery.NewDocument(base_url + url)
		if err != nil {
			log.Fatal(err)
		}
		doc.Find("#apply_join_time > option").Each(func(i int, s *goquery.Selection) {
			date, _ := s.Attr("value")
			vacants = append(vacants, date)
		})
	})
	return vacants
}
