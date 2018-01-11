package websim

import (
	"github.com/gin-gonic/gin"
	"math/rand"
  "strconv"
  "time"
)

func Timer(c *gin.Context) {
  s, _ := strconv.Atoi(c.Param("bucketstart"))
  e, _ := strconv.Atoi(c.Param("bucketend"))
	randsleep := s + rand.Intn(e-s)
  time.Sleep(time.Duration(randsleep) * time.Millisecond)
	c.JSON(200, gin.H{
		"message":     "OK",
		"bucketstart": c.Param("bucketstart"),
		"bucketend":   c.Param("bucketend"),
		"randsleep_ms":   randsleep,
	})
}


func TimerSize(c *gin.Context) {
  s, _ := strconv.Atoi(c.Param("bucketstart"))
  e, _ := strconv.Atoi(c.Param("bucketend"))
  m, _ := strconv.Atoi(c.Param("mb"))
  var mydata []byte
  for i := 0; i < (m * 1000000) ; i++ {
    mydata = append(mydata, 0)
  }
	randsleep := s + rand.Intn(e-s)
  time.Sleep(time.Duration(randsleep) * time.Millisecond)
	c.JSON(200, gin.H{
		"message":     "OK",
		"bucketstart": c.Param("bucketstart"),
		"bucketend":   c.Param("bucketend"),
		"randsleep_ms":   randsleep,
    "size_bytes": len(mydata),
    "data": mydata,
	})
}

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}
