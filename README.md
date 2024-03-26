# jobileu

<p align="center"> 
  <a href="https://go.dev/" target="_blank">
    <img width="80%" src="https://blog.materialize.pro/wp-content/uploads/2021/10/GOLANG.png"/>
  </a> 
</p>

<h4 align="center" >🚀 🟦 Jobileu 🟦 🚀</h4>

<h4 align="center">
  GO API To find Jobs Opportunities
</h4>

<p align="center">
  |&nbsp;&nbsp;
  <a style="color: #8a4af3;" href="#project">Overview</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a style="color: #8a4af3;" href="#techs">Technologies</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a style="color: #8a4af3;" href="#app">Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;
  <a style="color: #8a4af3;" href="#routes">Routes</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a style="color: #8a4af3;" href="#run-project">Run</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;
  <a style="color: #8a4af3;" href="#author">Author</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
</p>

#

<h1 align="center">
  
  <a href="https://github.com/Samuel-Ricardo">
    <img src="https://img.shields.io/static/v1?label=&message=Samuel%20Ricardo&color=black&style=for-the-badge&logo=GITHUB"/>
  </a>

  <a herf="https://www.instagram.com/samuel_ricardo.ex/">
    <img src='https://img.shields.io/static/v1?label=&message=Samuel.ex&color=black&style=for-the-badge&logo=instagram'/> 
  </a>

  <a herf='https://www.linkedin.com/in/samuel-ricardo/'>
    <img src='https://img.shields.io/static/v1?label=&message=Samuel%20Ricardo&color=black&style=for-the-badge&logo=LinkedIn'/> 
  </a>

</h1>

<br>

<p id="project"/>

<h2> | 🛰️ About:  </h2>

<p align="justfy">
  Jobileu API use the powerfull GO Lang to register Jobs opportunites, i use GIN to build a REST Server and storage data with SQLite, the API documentation use SWAGGER and all Application runs in DOcker Containers
</p>

<br>

<h2 id="techs">
| 🏗️ - Technologies and Concepts Studied:
</h2>

> <a href='https://go.dev/'> <img width="40px" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original.svg" /> </a>

<br>

- GO
- GIN
- SWAGGER
- SQLite
- Docker
- GORM
- Logger
- REST
- Swagger
- Validation
- Error Handler

> Among Others...
> <br>

#

<h2 id="app">
  💻 | Application:
</h2>

<img src="https://www.split.io/wp-content/uploads/go-sdk-blog.jpg"/>

<br>

<p>
  You can submit a new Job opportunities, search or delete a specific job, update a job data and list all jobs and data schemas to standardize requests, improving resilience.
</p>

<p>
  With SQLite i storage data locally and use GORM to communicate with database, the ambient is containerized with Docker and the Application is documented with SWAGGER.
</p>

<p>
  I Setup default requests and response that validate data before use case  
</p>

<br>

#

<h2 id="run-project"> 
   👨‍💻 | How to use
</h2>

<br>

### Open your Git Terminal and clone this repository

```git
  $ git clone "git@github.com:Samuel-Ricardo/jobileu.git"
```

### Make Pull

```git
  $ git pull "git@github.com:Samuel-Ricardo/jobileu.git"
```

<br>

This application use `Docker` so you dont need to install and cofigurate anything other than docker on your machine.

> <a target="_blank" href="https://www.docker.com/"> <img width="48px" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-plain-wordmark.svg" /> </a>

<br>

Navigate to project folder and run it using `docker-compose`

```bash

  # After setup docker environment just run this commmand on root project folder:

  $ docker-compose up --build   # For First Time run this command

  $ docker-compose up           # to run project


```

```bash

  #Apps Running on:

  $ API: http://localhost:8080

  See more: ./docker-compose.yaml

```

<br>

#
