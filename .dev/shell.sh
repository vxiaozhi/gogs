if [ "$1" == "build_image" ]; then
    docker build -f .dev/Dockerfile -t gin_nav_svr_dev .
elif [ "$1" == "compile" ]; then
    docker run -it --rm  -v $PWD/:/app  -w /app -p 8300:3000 gin_nav_svr_dev go build -o .dev/gogs
elif [ "$1" == "run_dev" ]; then
    docker run -it --rm  -v $PWD/.dev:/app  -w /app -p 8300:3000 gin_nav_svr_dev ./gogs -c app.ini

elif [ "$1" == "run_pro" ]; then
     docker run --name=gogs -d --rm -p 10022:22 -p 8300:3000 -v ./gogs_data:/data gogs/gogs
else
    docker run -it --rm  -v $PWD:/app  -w /app -p 8300:3000 gin_nav_svr_dev bash
fi