package user

import (
	"testing"

	"github.com/marcboeker/mega/graph/model"
	"github.com/marcboeker/mega/test"

	"github.com/stretchr/testify/suite"
)

type IdentityTestSuite struct {
	suite.Suite
	svc Service
	app *test.App
}

func (s *IdentityTestSuite) SetupTest() {
	var err error
	s.app, err = test.Initialize()
	if err != nil {
		s.Suite.FailNow(err.Error())
	}
	s.svc = New(s.app.Client)
}

func (s *IdentityTestSuite) TearDownTest() {
	s.app.Close()
}

func (s *IdentityTestSuite) TestCreateAndRetrieveUser() {
	user, err := s.svc.Create(s.app.Ctx, model.AddUserInput{Name: "Kermit", Age: 10})
	s.Suite.NoError(err)
	s.Suite.NotEmpty(user.ID)

	user2, err := s.svc.Get(s.app.Ctx, user.ID)
	s.Suite.NoError(err)
	s.Suite.Equal(user2.ID, user.ID)
}

func TestIdentityTestSuite(t *testing.T) {
	suite.Run(t, new(IdentityTestSuite))
}
