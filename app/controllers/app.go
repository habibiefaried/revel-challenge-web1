package controllers

import (
	"github.com/revel/revel"
	"io/ioutil"
	"strings"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "To QA & Sec: Tolong test web ini sebelum masuk prod ya!"
	return c.Render(greeting)
}

func (c App) GitConfig() revel.Result {
	return c.RenderText(`[core]
        repositoryformatversion = 0
        filemode = true
        bare = false
        logallrefupdates = true
[remote "origin"]
        url = git@github.com:habibiefaried/revel-challenge-web1.git
        fetch = +refs/heads/*:refs/remotes/origin/*
[branch "master"]
        remote = origin
        merge = refs/heads/master
`)
}

func (c App) GitHead() revel.Result {
	return c.RenderText("ref: refs/heads/master")
}

func (c App) Include() revel.Result {
	tfile := c.Params.Query.Get("file")
	log := c.Log.New("include",tfile)
	if (tfile == ""){
		return c.Redirect("/include?file=src/revel-challenge-web1/public/index.php")
	} else {
		log.Debug("include", tfile)
		if (strings.Contains(tfile, "php")) {
			b, err := ioutil.ReadFile(tfile)
			if err != nil {
				log.Errorf("Failed to load :%s",err.Error())
				return c.RenderText(err.Error())
			} else {
				return c.RenderText(string(b))
			}
		} else {
			return c.RenderText("Hanya boleh file php")
		}
	}
}