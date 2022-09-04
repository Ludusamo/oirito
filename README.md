# oirito

## Docker Testing

### Build

`docker build -t oirito .`

### Run

`docker run -p 5000:5000 --env-file ./.env --rm --name oirito oirito`

You have to create a `.env` file with the necessary environment variables set.
