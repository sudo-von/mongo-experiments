[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<br />
<p align="center">
  <a>
    <img src="https://static.wikia.nocookie.net/shingeki-no-kyojin/images/6/64/Eren_Jaeger_854_%28Anime%29.png/revision/latest/scale-to-width-down/340?cb=20210305233525&path-prefix=es" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">mongo-experiments</h3>

  <p align="center">
    Personal project where i'm gonna be using Docker, DockerCompose, MongoDB and Golang to improve my MongoDB/Golang skills. 
    <br />
    <a href="https://github.com/sudo-von/mongo-experiments"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/sudo-von/mongo-experiments">View Demo</a>
    ·
    <a href="https://github.com/sudo-von/mongo-experiments/issues">Report Bug</a>
    ·
    <a href="https://github.com/sudo-von/mongo-experiments/issues">Request Feature</a>
  </p>
</p>


<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#services">Services</a></li>
    <li><a href="#todo">To-do</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

I decided to create a repository that would allow me to practice basic / advanced queries in MongoDB so i created a REST API in Golang to test the basic  query operations (Create, Read, Update, Delete).

Due my limited time, the advanced queries won't be in the REST API as endpoints but feel free to collaborate in the case that you want to improve this. For now, if you want to found the advanced queries they will be in a separate file in the workbook folder since i will use them directly into the MongoDB console. 

I used Docker to create containers for all these services so that anyone could run the project and collaborate.

### Built With

* [Docker](https://www.docker.com/)
* [Docker-compose](https://docs.docker.com/compose/)
* [Golang](https://golang.org/)
* [Go-chi](https://github.com/go-chi/chi)
* [MongoDB](https://www.mongodb.com/)
* [MongoExpress](https://github.com/mongo-express/mongo-express)

### Prerequisites

* git
* docker
* docker-compose

### Installation

1. <b>Clone</b> the repo.
   ```sh
   git clone https://github.com/sudo-von/mongo-experiments.git
   ```
2. <b>Start</b> the services.
   ```sh
   docker-compose up --build
   ```
3. By default, <b>there will always be 3 documents</b> in the database collection because i didn't want to use volumes, these documents will be deleted and inserted every time you build the container, if you want to add more documents you can do it in the mongo/init/users.json file, also, if you put a json with a new structure and file name, it will be inserted as well as long as it is within the same default path at mongo/init/.
4. <b>Verify</b> that the services are running.
  ```
  Successfully tagged mongo-experiments_api:latest
  Starting mongo         ... done
  Starting mongo-express ... done
  Starting mongo-seed    ... done
  Starting api           ... done
  ```

<!-- USAGE EXAMPLES -->
## Usage

1. Make requests to the endpoints you want to test using an <b>HTTP Client</b> like <b>Postman/Insomnia</b>.
![insomnia](img/insomnia.jpg)

2. Explore the <b>workbook folder</b>.
3. If you want to learn more about <b>MongoDB theory</b> then read the theory file.

## Services

| Port | Service | Description |
| ------ | ------ | ------ |
| 3000 | api | Golang REST API service.|
| 27917 | mongo | MongoDB service. |
| 8081 | mongo-express | Web-based MongoDB admin interface written with Node.js, Express and Bootstrap3.|

## TO-DO

##### Docker containers

<ul>
    <li><del>MongoDB</del></li>
    <li><del>Mongo seed</del></li>
    <li><del>Mongo express</del></li>
    <li><del>Golang REST API</del></li>
</ul>

##### Rest API endpoints

<ul>
    <li><del>GET /users</del></li>
    <li><del>POST /users</del></li>
    <li><del>PUT /users</del></li>
    <li><del>DELETE /users</del></li>
</ul>

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<!-- CONTACT -->
## Contact

Von - [@sudo_von](https://twitter.com/sudo_von) - sudo.von.contact@gmail.com

Project Link: [https://github.com/sudo-von/mongo-experiments](https://github.com/sudo-von/mongo-experiments)


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/sudo-von/mongo-experiments.svg?style=for-the-badge
[contributors-url]: https://github.com/sudo-von/mongo-experiments/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/sudo-von/mongo-experiments.svg?style=for-the-badge
[forks-url]: https://github.com/sudo-von/mongo-experiments/network/members
[stars-shield]: https://img.shields.io/github/stars/sudo-von/mongo-experiments.svg?style=for-the-badge
[stars-url]: https://github.com/sudo-von/mongo-experiments/stargazers
[issues-shield]: https://img.shields.io/github/issues/sudo-von/mongo-experiments.svg?style=for-the-badge
[issues-url]: https://github.com/sudo-von/mongo-experiments/issues
[license-shield]: https://img.shields.io/github/license/sudo-von/mongo-experiments.svg?style=for-the-badge
[license-url]: https://github.com/sudo-von/mongo-experiments/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/jes%C3%BAs-%C3%A1ngel-rodr%C3%ADguez-mart%C3%ADnez-84991a1b4/