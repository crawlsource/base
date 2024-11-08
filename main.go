package main

import (
	"crawlsource_base/Common/Configs"
	"fmt"
	"github.com/gocolly/colly/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {

	Configs.LoadConfigs()
	_, err := gorm.Open(sqlite.Open(Configs.Env.String("databases.sqlite")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database SQLITE")
	}

	BaseURL := "https://ponisha.ir"
	VisitURL := fmt.Sprintf("%s/dashboard/find-projects?filterSkillsByUserId=38245")
	cookies := []*http.Cookie{
		{
			Name:  "accessToken",
			Value: "U2FsdGVkX18rLanocN3SkhWutHYFRK6KLNC/GB9iO9l0f1nRvmz4WC7IdWiG074CWd4LqDs6jLkhCf0wAKQLm9GLbCpJoEtjhvu2kXPw1+1s2z17JJIekLegH9InInyhP2TPpqShnxTyQMqwGpJXouvnLXveQNVt5PGfNNlYertPjbZ9mmGK6QIfAotUr12+vujcZf4bRPHwCFyJoHvXvrFfm3OSEPki/MJsqLN2O0AnOMKMoLIzFvrHzrAJFTIKFhbJ+R+255b/l84OxBTMgP04tU2oYgPFXAUtOHa7h4Coq9RgSOHQC0Yx5B7LFReAzrsRDPCg6z3mPwPAF+rGeQCoBCkylcyZ9SHrbaWwrU6XzkDOWDcPHcNgeWr2qzT6GkT2ZznT326RSdUBdFTm1Xf9BcTUHi4aU3HkzmiHmCKBydHIugBpzove1UYb/dwa4HboHwl8OFsotR7HXuObLbe4S/GvTOcGraYAfHnpevSJyWcb8diRxCQXpPYl2No5WbCHekte48QxrxGNHIKhc+FT83jT5BtWp6+Qe1zTP3G5mXDxyi3Lg1YUt8w9vkexWmHT9yAbd268K4KS34cT09TYgevlAVZ4TwwP5SRH5IhVAX9qg8vZdQT/PdiSwOT+280Zy5gGky03pzltPalVkZ9RHxj/6+j+C/8BDS61awFAs1PsDJWmCggNUZA2Kgj68lLGH1SFjjtZs3Lje2Ko5ShUsYLMxNo6gOaxYcRVP+Lp1pA/R8tmeDrj2zXvbVIvEr5Vy4VI0h2VU6Juuzmg9i3AqKL8wWeMb8VE3SuYanuvT3YzkbPHg4dzYnVlBL4QqmL+JK8ZvUJelhztL3a+199mnJIIgWwXx2ngfi+k3CJq1PDzj/FMWuPfV0Q3WjzmFnywZ30DDX9MBl3nGapCvkPL/KS12kZoxzjh6JZpZuVdBc6eeCOJEqn7Vwn4QrPyiuZ8eIy1zUgd+NPQTP3zShUaN/j4T3b26Swop+g0NntBAgifcPxFn3DkEhUVA6HPVuLs+7P97PwM8ZbK+sYg9WSsDw8tDX1g/O5DVmYg/WON8yDzWw0ogSoqsqcgybO+MOQcsjJJszmVqOAX6k/KiB/HyNTJ/gCdJdTdnY1yE2p0L0juO8VYY4/LKRP0Sr8uZT6y1C1fP8pD/f2gR4H+xwabh+CaWScnW0FyuQwyeI/8mD/AnvRold9x5tV1P20s/ARfZFI/ZpLq7/QMgKyVcOmdvCYV+wZspC3m8NOw3S8u2Lbn3rXbouNg9URuNkdlFGtqCsMe8iiGf2mDmctkYirgX373PZXpWNnns96EuSWEy0YhfuLkIptynXyXziJCFtuClKhlsINGnv5KmXrohW+JBGyqvEmhw+Gjj/sz7DjSP1IY0x/HQgkS3stfLDT8RY1aeMm7WILbj5v3QxBZcUOHPYPXeV4DiqkDyGSohzo=",
		},
	}
	c := colly.NewCollector()
	c.OnHTML("a.MuiBox-root", func(element *colly.HTMLElement) {
		fmt.Printf("%s : %s \n", element.ChildText("span"), element.Attr("href"))
	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.SetCookies(BaseURL, cookies)
	c.Visit(VisitURL)
	fmt.Println("Finish")

}
