package main

import (
	"fmt"
	"log"

	// "github.com/gorilla/mux"

	"github.com/gin-gonic/gin"
	"github.com/rparmer/remote-falcon-api/internal/config"
	"github.com/rparmer/remote-falcon-api/internal/database"
	"github.com/rparmer/remote-falcon-api/internal/route"
	"github.com/rparmer/remote-falcon-api/internal/services/remote"
)

//	func getHello(w http.ResponseWriter, r *http.Request) {
//		t := config.GetConfig()
//		fmt.Println(t.Database)
//		io.WriteString(w, "Hello, HTTP!\n")
//	}
// var secret = []byte("secret")

func main() {
	// r := mux.NewRouter()
	c := config.GetConfig()
	mode := gin.ReleaseMode

	if c.Server.Debug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	r := gin.New()
	// r.Use(sessions.Sessions("mysession", cookie.NewStore(secret)))
	// r.Use(auth.AuthRequired)
	router := r.Group("/remotefalcon/api")
	route.AddRoutes(router)

	// db, _ := database.New()
	db, _ := database.New()
	// database.Migrate(db)
	remoteService := remote.NewService(db)
	err := remoteService.SaveJuke(remote.RemoteJuke{
		RemoteId: 4,
	})
	if err != nil {
		fmt.Println(err)
	}

	remote, _ := remoteService.GetJuke(4)
	fmt.Println(remote.RemoteId)
	// activeViewerService := model.NewActiveViewerRepository(db)
	// activeViewerService.DeleteAllByRemoteToken("1234")
	// r.HandleFunc("/", getHello).Methods("POST", "GET")
	r.GET("/", func(c *gin.Context) {
		// c.String(200, "lakfjlafjlaf\nakdfalfd")

		// r, err := remoteService.GetRemoteByName("parmer family lights")
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// c.JSON(200, r)
		// c.IndentedJSON(200, models.GetRemote())
	})

	// var remote Remote
	// remote := models.GetRemote()

	// fmt.Println(remote)

	fmt.Println("db connected!!")

	fmt.Println("listening on port 8080...")
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
	// err := http.ListenAndServe(":8080", r)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
