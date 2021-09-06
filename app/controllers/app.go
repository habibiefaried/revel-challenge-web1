package controllers

import (
	"github.com/revel/revel"
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
	if (tfile == ""){
		return c.Redirect("/include?file=1.php")
	} else {
		return c.RenderText(tfile)
	}
}