<h1>backend</h1>
1.In order to run the backend service in the system successfully, Go and MySQL are needed in the system. These two programs can be installed at their official websites: https://go.dev/dl/, https://www.mysql.com/downloads/.
After installed Golang and MySQL, user should create database at MySQL, and set charset as utf8-mb4. Here is an example name ‘LMSystem’, we can go to the config.ini file in the conf file and set the user name, password and database name.
<br/>
<br/>
2.After cloning, use below code to run the backend server:
<br/>
(1)$ cd backend
<br/>
(2)$ go mod tidy
<br/>
(3)$ go run main.go
<h1>frontend</h1>
1.Make sure you computer already install node.js.<br/>
If not, you can downloaded from <a href="https://nodejs.org/en">here</a><br/>
2.After cloning, use below code to run the frontend server:
<br/>
  (1)$ cd frontend
<br/>
  (2)$ npm install -i
<br/>
  (3)$ npm start
