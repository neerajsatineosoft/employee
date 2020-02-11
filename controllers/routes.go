package controllers

import "github.com/employee/middlewares"

func (s *Server) InitializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Employee routes
	//s.Router.HandleFunc("/registration", middlewares.SetMiddlewareJSON(s.Registration)).Methods("POST")
	//s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")

	// s.Router.HandleFunc("/getemployees", middlewares.SetMiddlewareJSON(s.GetEmployee)).Methods("GET")
	// s.Router.HandleFunc("/Employees/{id}", middlewares.SetMiddlewareJSON(s.GetEmployees)).Methods("GET")
	// s.Router.HandleFunc("/Employees/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateEmployees))).Methods("PUT")
	// s.Router.HandleFunc("/Employees/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteEmployee)).Methods("DELETE")

}
