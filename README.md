# IoT Web App

This Go app shows the hostname it is running on. In K3s, it also shows the Kubernetes Pod UID as **Pod ID**.

## Run Docker and K3s

Install Docker with Compose, GNU Make, Vagrant, and VirtualBox. From the project root, run:

```sh
make up
```

This builds `iot-webapp:v1.0.0`, starts the local `webapp1` container, and boots the Vagrant VM. Vagrant provisions K3s on first boot. Make imports the image into K3s, deploys `webapp1`, and waits for the rollout.

- [Local Docker app](http://localhost:8080) shows `webapp1` as the hostname.
- [K3s app](http://192.168.56.111:30080) shows the pod's hostname and UID.

To compare the values from the K3s page with the pod name and UID, run:

```sh
cd infra
unset VAGRANT_CWD VAGRANT_DOTFILE_PATH
vagrant ssh -c 'sudo k3s kubectl get pods -l app=webapp1 -o custom-columns=HOSTNAME:.metadata.name,POD_ID:.metadata.uid'
```

Run the following Make targets from the project root:

| Command | Action |
| --- | --- |
| `make vm-up` | Boot the VirtualBox VM and provision K3s on first boot. |
| `make deploy` | Build the image and deploy it to K3s. |
| `make app-up` | Build and start the Docker Compose app. |
| `make down` | Remove the Compose container and halt the VM. |

## Run with Docker directly

Use this in place of Docker Compose for the local instance. If Compose is running, stop it first with `docker compose down`.

```sh
docker build -t iot-webapp:v1.0.0 . && \
docker run --rm -d \
  --name webapp1 \
  --hostname webapp1 \
  -p 8080:8080 \
  iot-webapp:v1.0.0
```

Open [the app](http://localhost:8080). To stop the container, run `docker stop webapp1`.
