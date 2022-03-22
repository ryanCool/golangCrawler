package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync"
)

func initRouter(router *gin.Engine) {
	router.POST("craw", crawGithub)
}

func crawGithub(ctx *gin.Context) {
	wg := sync.WaitGroup{}
	var stars string
	var fork, contributor int
	var err error
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			stars, fork, contributor, err = craw()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{})
			}
		}()
	}

	wg.Wait()

	ctx.JSON(http.StatusOK, map[string]interface{}{"stars": stars, "forks": fork, "contributor": contributor})
}

func craw() (string, int, int, error) {
	resp, err := soup.Get("https://github.com/golang-jwt/jwt")
	if err != nil {
		panic(err)
	}
	doc := soup.HTMLParse(resp)
	starsNo := ""
	forkNo, contriNo := 0, 0
	contriNo, err = FindContriNo(&doc)
	if err != nil {
		return "", 0, 0, err
	}

	starsFound, forkFound := false, false
	links := doc.FindAll("div", "class", "mt-2")
	for _, link := range links {
		if !forkFound {
			forkNo, err = FindForkNo(&link)
			if err == nil {
				forkFound = true
			}
		}

		if !starsFound {
			starsNo, err = FindStarsNo(&link)
			if err == nil {
				starsFound = true
			}
		}
	}


	return starsNo, forkNo, contriNo, nil
}

func FindContriNo(doc *soup.Root) (int, error) {
	ls := doc.FindAll("div", "class", "BorderGrid-row")
	for _, link := range ls {
		d := link.Find("a", "href", "/golang-jwt/jwt/graphs/contributors")
		if d.Error != nil {
			continue
		} else {
			dd := d.Find("span", "data-view-component", "true")
			return strconv.Atoi(dd.Text())
		}
	}
	return 0, fmt.Errorf("not found")
}

func FindStarsNo(link *soup.Root) (string, error) {
	d := link.Find("a", "href", "/golang-jwt/jwt/stargazers")
	if d.Error != nil {
		return "0", d.Error
	}

	starsStr := link.Find("strong").Text()

	return starsStr, nil
}

func FindForkNo(link *soup.Root) (int, error) {
	d := link.Find("a", "href", "/golang-jwt/jwt/network/members")
	if d.Error != nil {
		return 0, d.Error
	}

	forkNoStr := link.Find("strong").Text()
	forkNo, err := strconv.Atoi(forkNoStr)
	if err != nil {
		fmt.Println("convert fork number err:", err)
		return 0, err
	}

	return forkNo, nil
}
