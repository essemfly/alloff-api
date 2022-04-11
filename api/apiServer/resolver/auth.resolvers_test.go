package resolver

import (
	"fmt"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/google/uuid"
	"github.com/lessbutter/alloff-api/api/apiServer"
	"github.com/lessbutter/alloff-api/api/apiServer/middleware"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/storage/mongo"
	"github.com/stretchr/testify/require"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestAuthResolvers(t *testing.T) {
	testConf := config.GetConfiguration("local")
	conn := mongo.NewMongoDB(testConf)
	conn.RegisterRepos()

	testUserMobile := "01073881067"
	testUserUUID := "aeb06898-5183-4fca-9e37-851999f26f5a"
	testUserJwt, _ := middleware.GenerateToken(testUserMobile, testUserUUID)

	h := handler.NewDefaultServer(apiServer.NewExecutableSchema(apiServer.Config{Resolvers: &Resolver{}}))
	c := client.New(h)

	// Test CreateUser
	t.Run("Create User", func(t *testing.T) {
		var resp interface{}

		rand.Seed(time.Now().UnixNano())
		CODE_CHARSET := []rune("123467890")
		b := make([]rune, 8)
		for i := range b {
			b[i] = CODE_CHARSET[rand.Intn(len(CODE_CHARSET))]
		}
		mobile := "010" + string(b)
		uuid := uuid.New()
		queryStr := fmt.Sprintf(`
			mutation NewUser {
  				createUser(
    				input: {
      					uuid: "%s",
    					mobile: "%s"
					}
				)
			}`, uuid, mobile)

		c.MustPost(queryStr, &resp)

		user, _ := ioc.Repo.Users.GetByMobile(mobile)
		require.Equal(t, mobile, user.Mobile)
		require.Equal(t, uuid.String(), user.Uuid)
		require.Equal(t, reflect.TypeOf(time.Now()), reflect.TypeOf(user.Created))
		require.Equal(t, reflect.TypeOf(time.Now()), reflect.TypeOf(user.Updated))
	})

	// Test Login
	t.Run("Login", func(t *testing.T) {
		var resp map[string]interface{}

		queryStr := fmt.Sprintf(`
			mutation Login {
				login(
					input: {
						mobile: "%s",
						uuid: "%s",
					}
				)
			}`, testUserMobile, testUserUUID)

		c.MustPost(queryStr, &resp)
		actualToken := resp["login"]
		require.Equal(t, testUserJwt, actualToken)
	})

	// Test UpdateUserInfo
	t.Run("UpdateUserInfo", func(t *testing.T) {
		var resp struct {
			updateUserInfo struct {
				id                    string
				uuid                  string
				mobile                string
				name                  string
				email                 string
				baseAddress           string
				detailAddress         string
				postcode              string
				personalCustomsNumber string
			}
		}

		name := "테스트"
		email := "test@gqltest.com"
		baseAddress := "서울특별시"
		detailAddress := "독도 302호"
		postcode := "10351"
		personalCustomsNumber := "P170021193583"

		queryStr := fmt.Sprintf(`
			mutation UpdateUserInfo {
				updateUserInfo(
					input: {
						uuid: "%s"
						mobile: "%s"
						name: "%s"
						email: "%s"
						baseAddress: "%s"
						detailAddress: "%s"
						postcode: "%s"
						personalCustomsNumber: "%s"
						}
					) {
					id, uuid, mobile, name, email, baseAddress, detailAddress, postcode, personalCustomsNumber
					}
				}
			`, testUserUUID, testUserMobile, name, email, baseAddress, detailAddress, postcode, personalCustomsNumber)
		c.MustPost(queryStr, &resp)
	})
}
