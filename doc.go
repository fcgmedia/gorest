/*
GoREST is a RESTful style web-services framework.

Creating services in Go is straight forward, GoREST takes this a step further by adding a layer that makes tedious tasks much more automated and avoids regular pitfalls. This gives you the opportunity to focus more on the task at hand: minor low-level HTTP handling.

Example:

	func main() {
		gorest.RegisterService(new(HelloService))
		http.Handle("/", gorest.Handle())
		http.ListenAndServe(":8787", nil)
	}

	type HelloService struct {
		gorest.RestService `root:"/tutorial/"`
		helloWorld  gorest.EndPoint `method:"GET" path:"/hello-world/" output:"string"`
		sayHello    gorest.EndPoint `method:"GET" path:"/hello/{name:string}" output:"string"`
	}

	func (serv HelloService) HelloWorld() string {
		return "Hello World"
	}

	func (serv HelloService) SayHello(name string) string {
		return "Hello " + name
	}
*/
package gorest
