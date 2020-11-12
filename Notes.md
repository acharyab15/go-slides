* Jon Calhoun
    - Really cool guy who is easy to learn from
    - Reminds me of Cal, of course our Cal is more handsome, but mannerisms kind of

* A basic web application

package main -- Declares what package the code is a part of. the main package, we want to have application start by running the main() function.

imports -- tell the compiler which packages you intend to use in the code. 
    - fmt: format strings and write them to different places. 
    - net/http: create application that responds to web requests, make web requests to other servers.

handlerfunc(...) -- process incoming web requests. 
    - everytime someone visits the website, the code in handlerFunc(...) gets run and determines what to return to the visitor. 
    - Handlers take two arguments 
        - http.ResponseWriter(w): allows us to modify the response that we want to send to whoever visits the website. write to response body. name response writer.
        - http.Request(r): access data from the web request. e.g get users email address and password during sign up etc.
    - fmt.Fprint to write the HTML string Welcome to my awesome site to the http.ResponseWriter.

main().
    - handlerFunc is the function to handle all web requsts going to our server with path /.
    - any path like "some-other-path" would still work
    - http.ListenAndServe starts up a web server listening on port 3000 using the default http handlers


* Adding a contact and 404 page
    - We get a URL from the Request object. 
    - It has a Path that returns the path of the URL. 
    - Based on the path, we render a certain page. 
    - Otherwise, 404 using WriteHeader.

* Routers
    - net/http.ServeMux: Straightforward router, great for simple web apps. Lacks some features and need to write a lot on our own. We can add to it for applications, but not worth it.
    - github.com/julienschmidt/httprouter: popular router. simple, fast, efficient. outperforms many other routers. 
    - gorilla/mux.Router: one of the most popular routers. supports everything most people need. slightly heavy, but not that much

* What are templates?
    - Mad Libs are an example of a template used to create dynamic content.
    - one person asks for specific words (noun, verb, place, part of a car) without giving them context and once gathered, it is read aloud.
    - template in Go very similar, but instead of noun/adverb, variables are used.
    - template for our galleries. images showing up will change based on specific gallery etc.

* Why templates?
    - We have been putting html inside of strings. That won't work to scale it
    - Templates can be used to prevent duplication. A navbar template that can then be used in each page.
    - Add simple logic. Show a sign out button, if a user is logged in. Show sign in and sign up if logged out.

* Creating a template
    - Parse the hello.gohtml template file. Try to open the template file and validate it.
    - Create a data variable, which is an anonymous struct, with a Name field. Set Name to John Smith.
    - Execute the template, passing two things
        Where we want to write the output
        The data to be used when executing the template

* MVC
    Architectural pattern designed to help organize applications by separating code based on what it is responsible for

* Views
    - Given a specific page that we want to render, and data for that page, view

* Controllers
    - User accounts are represented in the database as user objects.
    - User model representative of that data in our database.
    - Handles stuff like updating user's contact information in the database.

* MVC walkthrough
    - A user submits an update to their contact information
    - The router decides that the UserController should be used
    - The UserController uses the User model to update the user’s contact info
    - The User model returns the updated data to the UserControllero
    - The UserController uses the ShowUser view to render the correct template
    - The end user sees the updated User page in their browser

* creating-first-views
    - git checkout creating-first-views








