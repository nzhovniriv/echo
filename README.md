# echo

#### Docker Reference

```bash
# Build an image (echo) from a Dockerfile
docker build --tag echo .

# Create and run a new container (echo) from the image (echo)
docker run --name echo -p 8888:8888 echo

# Remove the container (echo)
docker rm echo

# Remove the image (echo)
docker rmi echo
```