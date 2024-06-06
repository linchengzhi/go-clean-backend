package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/linchengzhi/go-clean-backend/domain/cerror"
	"github.com/linchengzhi/go-clean-backend/domain/dto"
	"github.com/linchengzhi/go-clean-backend/domain/entity"
	"github.com/linchengzhi/go-clean-backend/domain/types"
	"github.com/linchengzhi/go-clean-backend/util"
)

func SetSessionMiddleware(conf *dto.Config, g *gin.Engine) {
	store, err := NewRedisStore(conf)
	if err != nil {
		panic(err)
	}
	g.Use(sessions.Sessions("mysession", store))
}

// NewSessionStore creates a new session store
func NewRedisStore(conf *dto.Config) (sessions.Store, error) {
	store, err := redis.NewStore(10, "tcp", conf.Redis.Addr, conf.Redis.Password, []byte("secret"))
	if err != nil {
		return nil, err
	}
	store.Options(sessions.Options{
		MaxAge:   types.SessionMaxAge,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	})
	return store, nil
}

// set session
func SetSession(c *gin.Context, user *entity.User) {
	session := sessions.Default(c)
	session.Set("user", user)
	session.Save()
}

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			util.RespondErr(c, cerror.ErrLogout)
			c.Abort()
		} else {
			//设置到context
			c.Set("user", user)
			c.Set("user_id", user.(entity.User).Id)
		}
	}
}

func DelSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user")
}
