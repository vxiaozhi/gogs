if [ "$1" == "build" ]; then
    docker build -f .devcontainer/Dockerfile -t gin_nav_svr_dev .
else
    docker run -it --rm  -v $PWD:/app  -w /app  gin_nav_svr_dev bash
fi