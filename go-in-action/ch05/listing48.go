package main

//import "fmt"
//
//type notifier interface {
//	notify()
//}
//
//type user struct {
//	name string
//	email string
//}
//
//func (u *user) notify() {
//	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
//}
//
//type admin struct {
//	name string
//	email string
//}
//
//func (u *admin) notify() {
//	fmt.Printf("Sending admin email to %s<%s>\n", u.name, u.email)
//}
//
//func main()  {
//	bill := user{"Bill", "bill@gmail.com"}
//	sendNotification(&bill)
//
//	lisa := user{"Lisa", "lisa@gmail.com"}
//	sendNotification(&lisa)
//}
//
//func sendNotification(n notifier)  {
//	n.notify()
//}
