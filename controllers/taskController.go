package controllers
import(
	"net/http"
	"task-manager/database"
	"task-manager/models"
	"github.com/gin-gonic/gin"
)
func CreateTask(c *gin.Context){
	var task models.Task
	if err:=c.ShouldBindJSON(&task);err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	_,err := database.DB.NamedExec(`INSERT INTO TASKS (title, description, status, priority, deadline, project_id,user_id) values (:title, :description,:status,:priority,:deadline,:project_id,:user_id)`,&task)
	if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Error creating task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"Task created successfully"})
}