package router

import (
	api "9900project/api/v1"
	"9900project/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	//r.POST("chat", api.Chat)
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "pong"})
		})

		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			//user function
			authed.PUT("user", api.UpdateUser)
			authed.POST("user/avatar", api.UploadAvatar)
			authed.POST("user/password", api.ChangePassword)
			//authed.POST("user/sending-email", api.SendEmail)
			//authed.POST("user/valid-email", api.ValidEmail)

			//forum function
			authed.POST("forum", api.CreateForum)
			authed.GET("forum/:course_number", api.ShowForumList)
			//post function
			authed.POST("post", api.CreatePost)
			//authed.GET("post", api.GetPost2)
			//authed.POST("post_vote", api.PostVote)
			authed.GET("post/:id", api.GetPostByForumId)
			authed.GET("posts/:course", api.GetPostByCourseNumber)
			authed.GET("post_information/:id", api.GetPostInformationByForumId)
			authed.POST("post_search/:course", api.SearchPostByInfo)
			//comment function
			authed.POST("comment", api.CreateComment)
			authed.GET("comment/:id", api.GetCommentByPostId)

			//show and enroll course
			authed.POST("teacher-course", api.CreateCourse)
			authed.GET("course", api.GetCoursesById)
			authed.POST("course-number", api.GetCoursesByNumber)
			authed.POST("student-course", api.SelectCourse)
			authed.GET("student-course", api.GetCoursesSelectById)
			authed.GET("student-select-course", api.StudentSelectCourse)
			authed.DELETE("student_course/:courseNumber", api.DropCourseById)
			authed.POST("student-course-statistics", api.Statistics)
			authed.GET("course/:courseNumber", api.GetTeacher)

			//course material
			authed.POST("material", api.CreateMaterial)
			authed.GET("material/:course_number/:file_category", api.ShowMaterial)
			authed.PUT("material", api.UpdateMaterial)
			authed.DELETE("material", api.DeleteMaterial)
			authed.POST("material/:course_number", api.SearchMaterialByInfo)

			//assignment
			authed.POST("assignment", api.CreateAssignment)
			authed.PUT("assignment", api.UpdateAssignment)
			authed.GET("assignment/:course_number", api.ShowAssignment)
			authed.DELETE("assignment", api.DeleteAssignment)

			//assignment solution and grade
			authed.POST("assignment_solution", api.CreateAssMark)
			authed.DELETE("assignment_solution", api.DeleteAssMark)
			authed.GET("assignment_solution/:course_number/:assignment_id", api.ShowAssMark)
			authed.GET("assignment_solution/:course_number", api.ShowAssMarkForStudent)
			authed.GET("assignment_submission/:course_number/:assignment_id", api.ShowSubmission)
			authed.PUT("assignment_grade", api.UpdateAssMark)
			authed.PUT("assignment_solution", api.UpdateAssMarkGroup)
			authed.GET("assignment_mark/:group_id", api.GetAssMarkByGroupId)
			//quiz
			authed.POST("quiz", api.CreateQuiz)
			authed.GET("quiz/:course_number", api.GetQuiz)
			authed.POST("quiz_question", api.CreateQuizQuestion)
			authed.GET("quiz_question/:quiz_id", api.GetQuizQuestions)
			authed.GET("quiz_sum/:course_number", api.GetQuizSum)
			authed.POST("quiz_mark", api.CreateQuizMark)
			authed.POST("get_quiz_mark", api.GetQuizMark)

			//sum_score
			authed.GET("score/:course_number", api.GetScore)

			//notification
			authed.GET("notification", api.GetNotifications)
			authed.PUT("notification/:notification_id", api.UpdatetNotification)

			//tutor
			authed.POST("tutor", api.CreateTutor)
			authed.GET("tutor/:course_number", api.GetTutor)
			authed.DELETE("tutor/:id", api.DeleteTutor)
			authed.POST("users", api.GetUserByName)

			//group
			authed.POST("group", api.CreateGroup)
			authed.GET("group/:course_number", api.GetGroup)
			authed.DELETE("group/:id", api.DeleteGroup)
			authed.PUT("group/:id", api.UpdateTutor)
			authed.GET("assignment_group/:course_number", api.GetGroupByUserId)

			//recommand course based on studyoption
			authed.POST("user_mandatory_course", api.GetMandatoryCourse)
		}
	}
	return r
}
