# Live Customer Database
The project is meant to simulate live customer traffic on a primary and replica pair of MySQL instances. 

These 3 pieces of technology are presently driven by `docker-compose` but will eventually be converted to Kubernetes.

# Pre-Requisites

Since this application is designed to work with `docker-compose` you must install Docker and it's corresponding tool "docker-compose".

The following documentation from Docker can help with installation of both pieces of software:

<b>
<a href="https://docs.docker.com/get-docker/">Install Docker</a>


<a href="https://docs.docker.com/compose/install/">Install Docker-Compose</a>
</b>

# How-To Install

<details><summary>Step 1: Clone Repo & Build Docker Images</summary>
<p>

### Clone Repo & Build Docker Images

The following command will clone the repository to your local dev environment.

```bash
git clone https://github.com/excircle/liveDB.git
```
Once cloned, change directories into the folder called `images` and build the Docker image needed for the liveDB app.

```bash
cd images
sudo docker build -t livedb/mysql57 .
cd .. #To take you back to the 'liveDB' root directory "
```
Once complete you'll be able to use this image with the following docker-compose file.
</p>
</details>

<details><summary>Step 2: Launch docker-compose to setup MySQL instances.</summary>
<p>

### Configure 'docker-compose.yml' and Launch

The 
</p>
</details>