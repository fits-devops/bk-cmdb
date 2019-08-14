package service

import (
	"net/http"
)


func (s *testService) initClass() {
	s.addAction(http.MethodPost, "/create/class", s.CreateOneClass, nil)
	s.addAction(http.MethodPost, "/createmany/class", s.CreateManyClass, nil)
	s.addAction(http.MethodPost, "/setmany/class", s.SetManyClass, nil)
	s.addAction(http.MethodPost, "/set/class", s.SetOneClass, nil)
	s.addAction(http.MethodPut, "/update/class", s.UpdateClass, nil)
	s.addAction(http.MethodDelete, "/delete/class", s.DeleteClass, nil)
	s.addAction(http.MethodPost, "/read/class", s.SearchClass, nil)
}

func (s *testService) initStudent() {
	s.addAction(http.MethodPost, "/create/student", s.CreateStudent, nil)

}

func (s *testService) initService() {
	s.initClass()
	s.initStudent()
}
