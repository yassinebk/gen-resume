# Resume Gen

A simple CLI to generate a template based on a yaml file.

## Why ?

In today's market, the best way to secure interviews is by tailoring your resume to match the company or job offer. To achieve that, I found myself struggling with LaTeX, dealing with its frustrating syntax and slow compilation times (I use Overleaf since I switch laptops frequently).

After a few cups of coffee, I thought: Why not build something myself?

- Always accessible
- Fast
- Customizable
- Efficient: In the sense that I can retain previous data, generate new versions, and store everything in a single file containing all relevant information (career-related details, achievements, links)—keeping it all in one place so nothing gets forgotten.o now you only have to run the command:
-

```shell
genresume -i input.yaml -o myresume.pdf
```

Upload and Good Luck in your job search :D .

## Installation

```shell
go install github.com/yassinebk/resumegen@latest
```

### Deps

Make sure you have `wkhtmltopdf` installed on your machine and available in your PATH.

## Docker

You can use the docker image directly with the following command ( make sure to prepend `/tmp` as that's the path we are mounting the volume to )

```shell
docker run -v $(pwd):/tmp -it --rm ghcr.io/yassinebk/genresume:latest -i /tmp/data.yaml -o /tmp/resume.pdf
```
