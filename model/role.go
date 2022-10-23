package model
type Role interface{
	GetAuthByGroupID(int) bool
}