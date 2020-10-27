//routes.go
package main

func initializeRoutes(){
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)
}
//
//func showIndexPage(c *gin.Context) {
//	c.HTML(
//			http.StatusOK,
//			"index.html",
//			gin.H{
//				"title":"Home Page",
//			},
//
//			)
//}